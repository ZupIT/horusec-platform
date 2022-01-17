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
import Styled from './styled';
import { useTranslation } from 'react-i18next';
import { Button, Dialog, Datatable, DataSource } from 'components';
import coreService from 'services/core';
import useResponseMessage from 'helpers/hooks/useResponseMessage';
import { RepositoryToken } from 'helpers/interfaces/RepositoryToken';
import AddToken from './Add';
import useFlashMessage from 'helpers/hooks/useFlashMessage';
import { formatToHumanDate } from 'helpers/formatters/date';
import { useParams } from 'react-router-dom';
import { RouteParams } from 'helpers/interfaces/RouteParams';

const Tokens: React.FC = () => {
  const { t } = useTranslation();
  const { dispatchMessage } = useResponseMessage();
  const { showSuccessFlash } = useFlashMessage();

  const [tokens, setTokens] = useState<RepositoryToken[]>([]);
  const [isLoading, setLoading] = useState(false);
  const [deleteIsLoading, setDeleteIsLoading] = useState(false);

  const [tokenToDelete, setTokenToDelete] = useState<RepositoryToken>(null);
  const [addTokenVisible, setAddTokenVisible] = useState(false);

  const { workspaceId, repositoryId } = useParams<RouteParams>();

  const fetchData = () => {
    setLoading(true);
    coreService
      .getAllTokens(workspaceId, repositoryId)
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
      .removeToken(
        {
          workspaceID: workspaceId,
          repositoryID: repositoryId,
        },
        tokenToDelete.tokenID
      )
      .then(() => {
        showSuccessFlash(t('TOKENS_SCREEN.REMOVE_SUCCESS_TOKEN'));
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
    fetchData();
    //eslint-disable-next-line
   }, [workspaceId, repositoryId]);

  return (
    <Styled.Wrapper>
      <Styled.Header>
        <Button
          text={t('TOKENS_SCREEN.ADD_TOKEN')}
          rounded
          width={200}
          icon="plus"
          onClick={() => setAddTokenVisible(true)}
        />
      </Styled.Header>

      <Styled.Content>
        <Datatable
          columns={[
            {
              label: t('TOKENS_SCREEN.TOKEN'),
              property: 'token',
              type: 'text',
            },
            {
              label: t('TOKENS_SCREEN.DESCRIPTION'),
              property: 'description',
              type: 'text',
            },
            {
              label: t('TOKENS_SCREEN.EXPIRES'),
              property: 'expiresAt',
              type: 'text',
            },
            {
              label: t('TOKENS_SCREEN.ACTION'),
              property: 'actions',
              type: 'actions',
            },
          ]}
          dataSource={tokens.map((row) => {
            const repo: DataSource = {
              ...row,
              id: row.tokenID,
              token: '*********' + row.suffixValue,
              expiresAt: row.isExpirable
                ? formatToHumanDate(row.expiresAt)
                : t('GENERAL.NOT_EXPIRABLE'),
              actions: [
                {
                  title: t('TOKENS_SCREEN.DELETE'),
                  icon: 'delete',
                  function: () => setTokenToDelete(row),
                },
              ],
            };
            return repo;
          })}
          emptyListText={t('TOKENS_SCREEN.NO_TOKENS')}
          isLoading={isLoading}
        />
      </Styled.Content>

      <Dialog
        message={t('TOKENS_SCREEN.CONFIRM_DELETE_TOKEN')}
        confirmText={t('TOKENS_SCREEN.YES')}
        loadingConfirm={deleteIsLoading}
        defaultButton
        hasCancel
        isVisible={!!tokenToDelete}
        onCancel={() => setTokenToDelete(null)}
        onConfirm={handleConfirmDeleteToken}
      />

      <AddToken
        isVisible={addTokenVisible}
        currentParams={{ workspaceId, repositoryId }}
        onCancel={() => setAddTokenVisible(false)}
        onConfirm={() => {
          setAddTokenVisible(false);
          fetchData();
        }}
      />
    </Styled.Wrapper>
  );
};

export default Tokens;
