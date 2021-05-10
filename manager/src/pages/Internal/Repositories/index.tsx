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
import { SearchBar, Button, Dialog, Datatable, Datasource } from 'components';
import { useTranslation } from 'react-i18next';
import coreService from 'services/core';
import { Repository } from 'helpers/interfaces/Repository';
import useResponseMessage from 'helpers/hooks/useResponseMessage';

import HandleRepository from './Handle';
import InviteToRepository from './Invite';
import Tokens from './Tokens';
import useFlashMessage from 'helpers/hooks/useFlashMessage';
import useWorkspace from 'helpers/hooks/useWorkspace';
import { getCurrentConfig } from 'helpers/localStorage/horusecConfig';
import { authTypes } from 'helpers/enums/authTypes';
import { getCurrentUser } from 'helpers/localStorage/currentUser';
import useRepository from 'helpers/hooks/useRepository';

const Repositories: React.FC = () => {
  const { t } = useTranslation();
  const { isAdminOfWorkspace } = useWorkspace();
  const { dispatchMessage } = useResponseMessage();
  const { showSuccessFlash } = useFlashMessage();
  const { authType } = getCurrentConfig();

  const { allRepositories, fetchAllRepositories } = useRepository();

  const [handleRepositoryVisible, setHandleRepositoryVisible] = useState(false);
  const [deleteIsLoading, setDeleteLoading] = useState(false);

  const [repoToManagerTokens, setRepoToManagerTokens] = useState<Repository>(
    null
  );
  const [repoToDelete, setRepoToDelete] = useState<Repository>(null);
  const [repoToEdit, setRepoToEdit] = useState<Repository>(null);
  const [repoToInvite, setRepoToInvite] = useState<Repository>(null);

  const [searchQuery, setSearchQuery] = useState('');

  const filteredRepositories = () =>
    allRepositories.filter((repo) =>
      repo.name.toLocaleLowerCase().includes(searchQuery.toLocaleLowerCase())
    );

  const handleConfirmDeleteRepo = () => {
    setDeleteLoading(true);
    coreService
      .deleteRepository(repoToDelete.workspaceID, repoToDelete.repositoryID)
      .then(() => {
        showSuccessFlash(t('REPOSITORIES_SCREEN.REMOVE_SUCCESS_REPO'));
        fetchAllRepositories(repoToDelete?.workspaceID);
        setRepoToDelete(null);
      })
      .catch((err) => {
        dispatchMessage(err?.response?.data);
      })
      .finally(() => {
        setDeleteLoading(false);
      });
  };

  const setVisibleHandleModal = (
    isVisible: boolean,
    repository?: Repository
  ) => {
    setHandleRepositoryVisible(isVisible);
    setRepoToEdit(repository || null);
  };

  const handleConfirmRepositoryEdit = (repository: Repository) => {
    setVisibleHandleModal(false);
    fetchAllRepositories(repository?.workspaceID);
  };

  return (
    <Styled.Wrapper>
      <Styled.Options>
        <SearchBar
          placeholder={t('REPOSITORIES_SCREEN.SEARCH_REPO')}
          onSearch={(value) => setSearchQuery(value)}
        />
        {isAdminOfWorkspace ? (
          <Button
            text={t('REPOSITORIES_SCREEN.CREATE_REPO')}
            rounded
            width={180}
            icon="plus"
            onClick={() => setVisibleHandleModal(true)}
          />
        ) : null}
      </Styled.Options>

      <Styled.Content>
        <Datatable
          columns={[
            {
              label: t('REPOSITORIES_SCREEN.NAME'),
              property: 'name',
              type: 'text',
            },
            {
              label: t('REPOSITORIES_SCREEN.DESCRIPTION'),
              property: 'description',
              type: 'text',
            },
            {
              label: t('REPOSITORIES_SCREEN.ACTION'),
              property: 'actions',
              type: 'actions',
            },
          ]}
          datasource={filteredRepositories().map((row) => {
            const repo: Datasource = {
              ...row,
              id: row.repositoryID,
              actions: [],
            };

            if (row.role === 'admin' || getCurrentUser().isApplicationAdmin) {
              repo.actions.push({
                title: t('REPOSITORIES_SCREEN.EDIT'),
                icon: 'edit',
                function: () => setVisibleHandleModal(true, row),
              });

              if (isAdminOfWorkspace) {
                repo.actions.push({
                  title: t('REPOSITORIES_SCREEN.DELETE'),
                  icon: 'delete',
                  function: () => setRepoToDelete(row),
                });

                if (authType !== authTypes.LDAP) {
                  repo.actions.push({
                    title: t('REPOSITORIES_SCREEN.INVITE'),
                    icon: 'users',
                    function: () => setRepoToInvite(row),
                  });
                }
              }

              repo.actions.push({
                title: t('REPOSITORIES_SCREEN.TOKENS'),
                icon: 'lock',
                function: () => setRepoToManagerTokens(row),
              });
            }
            return repo;
          })}
          emptyListText={t('REPOSITORIES_SCREEN.NO_REPOSITORIES')}
        />
      </Styled.Content>

      <Dialog
        message={t('REPOSITORIES_SCREEN.CONFIRM_DELETE_REPO')}
        confirmText={t('REPOSITORIES_SCREEN.YES')}
        loadingConfirm={deleteIsLoading}
        defaultButton
        hasCancel
        isVisible={!!repoToDelete}
        onCancel={() => setRepoToDelete(null)}
        onConfirm={handleConfirmDeleteRepo}
      />

      <HandleRepository
        isVisible={handleRepositoryVisible}
        repositoryToEdit={repoToEdit}
        onCancel={() => setVisibleHandleModal(false)}
        onConfirm={handleConfirmRepositoryEdit}
      />

      <InviteToRepository
        isVisible={!!repoToInvite}
        repoToInvite={repoToInvite}
        onClose={() => setRepoToInvite(null)}
      />

      <Tokens
        isVisible={!!repoToManagerTokens}
        repoToManagerTokens={repoToManagerTokens}
        onClose={() => setRepoToManagerTokens(null)}
      />
    </Styled.Wrapper>
  );
};

export default Repositories;
