/**
 * Copyright 2021 ZUP IT SERVICOS EM TECNOLOGIA E INOVACAO SA
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

import React, { useState, useEffect } from 'react';
import { useTranslation } from 'react-i18next';
import Styled from './styled';
import { Button, Dialog, SearchBar, Datatable, DataSource } from 'components';
import { Webhook } from 'helpers/interfaces/Webhook';
import { useTheme } from 'styled-components';
import { get } from 'lodash';
import coreService from 'services/core';
import webhookService from 'services/webhook';
import useResponseMessage from 'helpers/hooks/useResponseMessage';
import useFlashMessage from 'helpers/hooks/useFlashMessage';

import HandleWebhook from './Handle';
import { useHistory, useParams } from 'react-router-dom';
import { Workspace } from 'helpers/interfaces/Workspace';
import { RouteParams } from 'helpers/interfaces/RouteParams';

interface Props {
  type: 'workspace' | 'repository';
}

const Webhooks: React.FC<Props> = ({ type }) => {
  const { t } = useTranslation();
  const { colors } = useTheme();

  const { dispatchMessage } = useResponseMessage();
  const { showSuccessFlash } = useFlashMessage();

  const [webhooks, setWebhooks] = useState<Webhook[]>([]);
  const [filteredWebhooks, setFilteredWebhooks] = useState<Webhook[]>([]);
  const [webhookToDelete, setWebhookToDelete] = useState<Webhook>();
  const [webhookToEdit, setWebhookToEdit] = useState<Webhook>();
  const [webhookToCopy, setWebhookToCopy] = useState<Webhook>();

  const [isLoading, setLoading] = useState(false);
  const [deleteIsLoading, setDeleteIsLoading] = useState(false);
  const [addWebhookVisible, setAddWebhookVisible] = useState(false);

  const [selectedWorkspace, setSelectedWorkspace] = useState<Workspace>(null);
  const { workspaceId, repositoryId } = useParams<RouteParams>();
  const history = useHistory();

  function getOneWorkspace() {
    setLoading(true);
    coreService
      .getOneWorkspace(workspaceId)
      .then((result) => {
        setSelectedWorkspace(result.data.content);
      })
      .catch((err) => {
        dispatchMessage(err?.response?.data);
        history.goBack();
      })
      .finally(() => {
        setLoading(false);
      });
  }

  useEffect(() => {
    if (workspaceId) getOneWorkspace();
    //eslint-disable-next-line
  }, [workspaceId]);

  const fetchData = () => {
    setLoading(true);

    webhookService
      .getAll(selectedWorkspace?.workspaceID)
      .then((result) => {
        setWebhooks(result?.data?.content);
        if (repositoryId) {
          const data = result?.data?.content.filter(
            (item: any) => item.repositoryID === repositoryId
          );
          setFilteredWebhooks(data);
          setWebhooks(data);
        } else {
          setWebhooks(result?.data?.content);
          setFilteredWebhooks(result?.data?.content);
        }
      })
      .catch((err) => {
        dispatchMessage(err?.response?.data);
      })
      .finally(() => {
        setLoading(false);
      });
  };

  const handleConfirmDelete = () => {
    setDeleteIsLoading(true);

    webhookService
      .remove(webhookToDelete.webhookID, webhookToDelete.workspaceID)
      .then(() => {
        showSuccessFlash(t('WEBHOOK_SCREEN.SUCCESS_DELETE'));
      })
      .catch((err) => {
        dispatchMessage(err?.response?.data);
      })
      .finally(() => {
        setDeleteIsLoading(false);
        fetchData();
        setWebhookToDelete(null);
      });
  };

  const onSearchWebhook = (search: string) => {
    if (search) {
      const filtered = webhooks.filter((webhook) =>
        webhook?.url.toLocaleLowerCase().includes(search.toLocaleLowerCase())
      );

      setFilteredWebhooks(filtered);
    } else {
      setFilteredWebhooks(webhooks);
    }
  };

  useEffect(() => {
    if (selectedWorkspace) fetchData();
    // eslint-disable-next-line
  }, [selectedWorkspace]);

  return (
    <Styled.Wrapper>
      <Styled.Options>
        <SearchBar
          placeholder={t('WEBHOOK_SCREEN.SEARCH')}
          onSearch={(value) => onSearchWebhook(value)}
        />

        <Button
          text={t('WEBHOOK_SCREEN.ADD')}
          rounded
          width={200}
          icon="plus"
          onClick={() => setAddWebhookVisible(true)}
        />
      </Styled.Options>

      <Styled.Content>
        <Datatable
          columns={[
            {
              label: t('WEBHOOK_SCREEN.TABLE.METHOD'),
              property: 'method',
              type: 'custom',
              cssClass: ['flex-center'],
            },
            {
              label: t('WEBHOOK_SCREEN.TABLE.URL'),
              property: 'url',
              type: 'text',
            },
            {
              label: t('WEBHOOK_SCREEN.TABLE.DESCRIPTION'),
              property: 'description',
              type: 'text',
            },
            {
              label: t('WEBHOOK_SCREEN.TABLE.REPOSITORY'),
              property: 'repository',
              type: 'text',
            },
            {
              label: t('WEBHOOK_SCREEN.TABLE.ACTION'),
              property: 'actions',
              type: 'actions',
            },
          ]}
          dataSource={filteredWebhooks.map((row) => {
            const data: DataSource = {
              ...row,
              id: row.webhookID,
              repository: row?.repository?.name,
              method: (
                <Styled.Tag
                  color={get(
                    colors.methods,
                    row?.method?.toLowerCase(),
                    colors.methods.unknown
                  )}
                >
                  {row.method}
                </Styled.Tag>
              ),
              actions: [
                {
                  title: t('WEBHOOK_SCREEN.TABLE.DELETE'),
                  icon: 'delete',
                  function: () => setWebhookToDelete(row),
                },
                {
                  title: t('WEBHOOK_SCREEN.TABLE.EDIT'),
                  icon: 'edit',
                  function: () => setWebhookToEdit(row),
                },
                {
                  title: t('WEBHOOK_SCREEN.TABLE.COPY'),
                  icon: 'copy',
                  function: () => {
                    setWebhookToCopy(row);
                    setAddWebhookVisible(true);
                  },
                },
              ],
            };
            return data;
          })}
          isLoading={isLoading}
          emptyListText={t('WEBHOOK_SCREEN.TABLE.EMPTY')}
        />
      </Styled.Content>

      <HandleWebhook
        isVisible={!!webhookToEdit}
        isNew={false}
        onCancel={() => setWebhookToEdit(null)}
        webhookInitial={webhookToEdit}
        onConfirm={() => {
          setWebhookToEdit(null);
          fetchData();
        }}
      />

      <HandleWebhook
        isVisible={addWebhookVisible}
        webhookInitial={webhookToCopy}
        isNew={true}
        onCancel={() => {
          setWebhookToCopy(null);
          setAddWebhookVisible(false);
        }}
        onConfirm={() => {
          setWebhookToCopy(null);
          setAddWebhookVisible(false);
          fetchData();
        }}
      />

      <Dialog
        message={t('WEBHOOK_SCREEN.CONFIRM_DELETE')}
        confirmText={t('WEBHOOK_SCREEN.YES')}
        loadingConfirm={deleteIsLoading}
        defaultButton
        hasCancel
        isVisible={!!webhookToDelete}
        onCancel={() => setWebhookToDelete(null)}
        onConfirm={handleConfirmDelete}
      />
    </Styled.Wrapper>
  );
};

export default Webhooks;
