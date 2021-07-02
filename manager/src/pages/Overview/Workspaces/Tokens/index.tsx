/**
 * Copyright 2020 ZUP IT SERVICOS EM TECNOLOGIA E INOVACAO SA
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
import Styled from './styled';
import { useTranslation } from 'react-i18next';
import { Button, Dialog, Datatable, Datasource } from 'components';
import coreService from 'services/core';
import useResponseMessage from 'helpers/hooks/useResponseMessage';
import { RepositoryToken } from 'helpers/interfaces/RepositoryToken';
import AddToken from './Add';
import useFlashMessage from 'helpers/hooks/useFlashMessage';
import { formatToHumanDate } from 'helpers/formatters/date';
import { Workspace } from 'helpers/interfaces/Workspace';
import { Link, useHistory, useParams } from 'react-router-dom';
import { IconButton } from '@material-ui/core';
import { ArrowBack } from '@material-ui/icons';

function WorkspaceTokens() {
  const { t } = useTranslation();
  const { dispatchMessage } = useResponseMessage();
  const { showSuccessFlash } = useFlashMessage();

  const [tokens, setTokens] = useState<RepositoryToken[]>([]);
  const [isLoading, setLoading] = useState(false);
  const [deleteIsLoading, setDeleteIsLoading] = useState(false);

  const [tokenToDelete, setTokenToDelete] = useState<RepositoryToken>(null);
  const [addTokenVisible, setAddTokenVisible] = useState(false);
  const [selectedWorkspace, setSelectedWorkspace] = useState<Workspace>(null);

  const { workspaceId } = useParams<{ workspaceId: string }>();
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
    coreService
      .getAllTokensOfWorkspace(selectedWorkspace.workspaceID)
      .then((result) => {
        setTokens(result?.data?.content);
      })
      .catch((err) => {
        dispatchMessage(err?.response?.data);
      })
      .finally(() => {
        setLoading(false);
      });
  };

  const handleConfirmDeleteToken = () => {
    setDeleteIsLoading(true);
    coreService
      .removeTokenOfWorkspace(tokenToDelete.workspaceID, tokenToDelete.tokenID)
      .then(() => {
        showSuccessFlash(t('WORKSPACES_SCREEN.REMOVE_SUCCESS_TOKEN'));
        setTokenToDelete(null);
        fetchData();
      })
      .catch((err) => {
        dispatchMessage(err?.response?.data);
      })
      .finally(() => {
        setDeleteIsLoading(false);
      });
  };

  useEffect(() => {
    if (selectedWorkspace) {
      fetchData();
    }
    //eslint-disable-next-line
  }, [selectedWorkspace]);

  return (
    <Styled.Wrapper>
      <Styled.Header>
        <Styled.TitleContent>
          <Link to="/overview/workspace">
            <IconButton size="small">
              <ArrowBack />
            </IconButton>
          </Link>
          <Styled.Title>{t('WORKSPACES_SCREEN.TOKENS')}</Styled.Title>
        </Styled.TitleContent>

        <Button
          text={t('WORKSPACES_SCREEN.ADD_TOKEN')}
          rounded
          width={150}
          icon="plus"
          onClick={() => setAddTokenVisible(true)}
        />
      </Styled.Header>

      <Styled.Content>
        <Datatable
          columns={[
            {
              label: t('WORKSPACES_SCREEN.TOKEN'),
              property: 'token',
              type: 'text',
            },
            {
              label: t('WORKSPACES_SCREEN.DESCRIPTION'),
              property: 'description',
              type: 'text',
            },
            {
              label: t('WORKSPACES_SCREEN.EXPIRES'),
              property: 'expiresAt',
              type: 'text',
            },
            {
              label: t('WORKSPACES_SCREEN.ACTION'),
              property: 'actions',
              type: 'actions',
            },
          ]}
          datasource={tokens.map((row) => {
            const data: Datasource = {
              ...row,
              id: row.tokenID,
              token: '********' + row.suffixValue,
              expiresAt: row.isExpirable
                ? formatToHumanDate(row.expiresAt)
                : t('GENERAL.NOT_EXPIRABLE'),
              actions: [
                {
                  title: t('WORKSPACES_SCREEN.DELETE'),
                  icon: 'delete',
                  function: () => setTokenToDelete(row),
                },
              ],
            };
            return data;
          })}
          isLoading={isLoading}
          emptyListText={t('WORKSPACES_SCREEN.NO_TOKENS')}
        />
      </Styled.Content>

      <Dialog
        message={t('WORKSPACES_SCREEN.CONFIRM_DELETE_TOKEN')}
        confirmText={t('WORKSPACES_SCREEN.YES')}
        loadingConfirm={deleteIsLoading}
        defaultButton
        hasCancel
        isVisible={!!tokenToDelete}
        onCancel={() => setTokenToDelete(null)}
        onConfirm={handleConfirmDeleteToken}
      />

      <AddToken
        isVisible={addTokenVisible}
        selectedWorkspace={selectedWorkspace}
        onCancel={() => setAddTokenVisible(false)}
        onConfirm={() => {
          setAddTokenVisible(false);
          fetchData();
        }}
      />
    </Styled.Wrapper>
  );
}

export default WorkspaceTokens;
