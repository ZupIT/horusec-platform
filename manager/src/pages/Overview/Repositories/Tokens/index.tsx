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
import { Repository } from 'helpers/interfaces/Repository';
import { useTranslation } from 'react-i18next';
import { Button, Dialog, Datatable, Datasource } from 'components';
import coreService from 'services/core';
import useResponseMessage from 'helpers/hooks/useResponseMessage';
import { RepositoryToken } from 'helpers/interfaces/RepositoryToken';
import AddToken from './Add';
import useFlashMessage from 'helpers/hooks/useFlashMessage';
import { formatToHumanDate } from 'helpers/formatters/date';
import { Link, useHistory, useParams } from 'react-router-dom';
import useWorkspace from 'helpers/hooks/useWorkspace';
import { IconButton } from '@material-ui/core';
import { ArrowBack } from '@material-ui/icons';

function RepositoryTokens() {
  const { t } = useTranslation();
  const { dispatchMessage } = useResponseMessage();
  const { showSuccessFlash } = useFlashMessage();

  const [tokens, setTokens] = useState<RepositoryToken[]>([]);
  const [isLoading, setLoading] = useState(false);
  const [deleteIsLoading, setDeleteIsLoading] = useState(false);

  const [tokenToDelete, setTokenToDelete] = useState<RepositoryToken>(null);
  const [addTokenVisible, setAddTokenVisible] = useState(false);

  const [repoToManagerTokens, setRepoToManagerTokens] =
    useState<Repository>(null);

  const { workspaceId, repositoryId } =
    useParams<{ workspaceId: string; repositoryId: string }>();
  const history = useHistory();

  function getOneRepository() {
    setLoading(true);
    coreService
      .getOneRepository(workspaceId, repositoryId)
      .then((result) => {
        setRepoToManagerTokens(result.data.content);
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
    if (workspaceId && repositoryId) getOneRepository();
    //eslint-disable-next-line
  }, [workspaceId, repositoryId]);

  const fetchData = () => {
    setLoading(true);
    coreService
      .getAllTokensOfRepository(
        repoToManagerTokens.workspaceID,
        repoToManagerTokens.repositoryID
      )
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
      .removeTokenOfRepository(
        tokenToDelete.workspaceID,
        tokenToDelete.repositoryID,
        tokenToDelete.tokenID
      )
      .then(() => {
        showSuccessFlash(t('REPOSITORIES_SCREEN.REMOVE_SUCCESS_TOKEN'));
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
    if (repoToManagerTokens) {
      fetchData();
    }
    //eslint-disable-next-line
  }, [repoToManagerTokens]);

  return (
    <Styled.Wrapper>
      <Styled.Header>
        <Styled.TitleContent>
          <Link to="/overview/repositories">
            <IconButton size="small">
              <ArrowBack />
            </IconButton>
          </Link>
          <Styled.Title>{t('REPOSITORIES_SCREEN.TOKENS')}</Styled.Title>
        </Styled.TitleContent>

        <Button
          text={t('REPOSITORIES_SCREEN.ADD_TOKEN')}
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
              label: t('REPOSITORIES_SCREEN.TOKEN'),
              property: 'token',
              type: 'text',
            },
            {
              label: t('REPOSITORIES_SCREEN.DESCRIPTION'),
              property: 'description',
              type: 'text',
            },
            {
              label: t('REPOSITORIES_SCREEN.EXPIRES'),
              property: 'expiresAt',
              type: 'text',
            },
            {
              label: t('REPOSITORIES_SCREEN.ACTION'),
              property: 'actions',
              type: 'actions',
            },
          ]}
          datasource={tokens.map((row) => {
            const repo: Datasource = {
              ...row,
              id: row.tokenID,
              token: '*********' + row.suffixValue,
              expiresAt: row.isExpirable
                ? formatToHumanDate(row.expiresAt)
                : t('GENERAL.NOT_EXPIRABLE'),
              actions: [
                {
                  title: t('REPOSITORIES_SCREEN.DELETE'),
                  icon: 'delete',
                  function: () => setTokenToDelete(row),
                },
              ],
            };
            return repo;
          })}
          emptyListText={t('REPOSITORIES_SCREEN.NO_TOKENS')}
        />
      </Styled.Content>

      <Dialog
        message={t('REPOSITORIES_SCREEN.CONFIRM_DELETE_TOKEN')}
        confirmText={t('REPOSITORIES_SCREEN.YES')}
        loadingConfirm={deleteIsLoading}
        defaultButton
        hasCancel
        isVisible={!!tokenToDelete}
        onCancel={() => setTokenToDelete(null)}
        onConfirm={handleConfirmDeleteToken}
      />

      <AddToken
        isVisible={addTokenVisible}
        currentRepository={repoToManagerTokens}
        onCancel={() => setAddTokenVisible(false)}
        onConfirm={() => {
          setAddTokenVisible(false);
          fetchData();
        }}
      />
    </Styled.Wrapper>
  );
}

export default RepositoryTokens;
