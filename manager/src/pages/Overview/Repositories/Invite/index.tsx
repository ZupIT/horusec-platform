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
import { Repository } from 'helpers/interfaces/Repository';
import {
  SearchBar,
  Select,
  Permissions,
  Datatable,
  Datasource,
} from 'components';
import { Account } from 'helpers/interfaces/Account';
import coreService from 'services/core';
import useResponseMessage from 'helpers/hooks/useResponseMessage';
import { getCurrentUser } from 'helpers/localStorage/currentUser';
import { findIndex, cloneDeep } from 'lodash';
import useFlashMessage from 'helpers/hooks/useFlashMessage';
import { Checkbox, IconButton } from '@material-ui/core';
import { Link, useHistory, useParams } from 'react-router-dom';
import { ArrowBack } from '@material-ui/icons';

function RepositoryInvite() {
  const { t } = useTranslation();
  const { dispatchMessage } = useResponseMessage();
  const { showSuccessFlash } = useFlashMessage();
  const currentUser = getCurrentUser();

  const [userAccounts, setUserAccounts] = useState<Account[]>([]);
  const [filteredUserAccounts, setFilteredUserAccounts] = useState<Account[]>(
    []
  );
  const [accountsInRepository, setAccountsInRepository] = useState<string[]>(
    []
  );
  const [isLoading, setLoading] = useState(false);
  const [permissionsIsOpen, setPermissionsOpen] = useState(false);

  const [repoToInvite, setRepoToInvite] = useState<Repository>(null);

  const { workspaceId, repositoryId } =
    useParams<{ workspaceId: string; repositoryId: string }>();
  const history = useHistory();

  const roles = [
    {
      label: t('PERMISSIONS.ADMIN'),
      value: 'admin',
    },
    {
      label: t('PERMISSIONS.SUPERVISOR'),
      value: 'supervisor',
    },
    {
      label: t('PERMISSIONS.MEMBER'),
      value: 'member',
    },
  ];

  function getOneRepository() {
    setLoading(true);
    coreService
      .getOneRepository(workspaceId, repositoryId)
      .then((result) => {
        setRepoToInvite(result.data.content);
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

  const fetchUsersInRepository = (allUsersInWorkspace: Account[]) => {
    coreService
      .getUsersInRepository(repoToInvite.workspaceID, repoToInvite.repositoryID)
      .then((result) => {
        const accountIds: string[] = [];
        const allUsers = cloneDeep(allUsersInWorkspace);

        // eslint-disable-next-line array-callback-return
        result?.data?.content.map((account: Account) => {
          accountIds.push(account.accountID);

          const index = findIndex(allUsers, {
            accountID: account.accountID,
          });

          allUsers[index] = account;
        });
        setAccountsInRepository(accountIds);
        setFilteredUserAccounts(allUsers);
        setUserAccounts(allUsers);
        setLoading(false);
      })
      .catch((err) => {
        dispatchMessage(err?.response?.data);
      });
  };

  const fetchAllUsersInWorkspace = () => {
    setLoading(true);
    coreService
      .getUsersInWorkspace(repoToInvite.workspaceID)
      .then((result) => {
        fetchUsersInRepository(result?.data?.content);
      })
      .catch((err) => {
        dispatchMessage(err?.response?.data);
      });
  };

  const onSearchUser = (search: string) => {
    if (search) {
      const filtered = userAccounts.filter((user) =>
        user.email.toLocaleLowerCase().includes(search.toLocaleLowerCase())
      );

      setFilteredUserAccounts(filtered);
    } else {
      setFilteredUserAccounts(userAccounts);
    }
  };

  const inviteUserToRepository = (account: Account) => {
    coreService
      .includeUserInRepository(
        repoToInvite.workspaceID,
        repoToInvite.repositoryID,
        account.email,
        account.role,
        account.accountID,
        account.username
      )
      .then(() => {
        showSuccessFlash(t('REPOSITORIES_SCREEN.SUCCESS_ADD_USER'));
        setAccountsInRepository([...accountsInRepository, account.accountID]);
      })
      .catch((err) => {
        dispatchMessage(err?.response?.data);
      });
  };

  const removeUserOfRepository = (account: Account) => {
    coreService
      .removeUserOfRepository(
        repoToInvite.workspaceID,
        repoToInvite.repositoryID,
        account.accountID
      )
      .then(() => {
        showSuccessFlash(t('REPOSITORIES_SCREEN.SUCCESS_REMOVE_USER'));
        const filteredIds = accountsInRepository.filter(
          (item) => item !== account.accountID
        );
        setAccountsInRepository(filteredIds);
      })
      .catch((err) => {
        dispatchMessage(err?.response?.data);
      });
  };

  const handleInviteUser = (isChecked: boolean, account: Account) => {
    if (isChecked) inviteUserToRepository(account);
    else removeUserOfRepository(account);
  };

  const handleChangeUserRole = (role: string, account: Account) => {
    coreService
      .updateUserRoleInRepository(
        repoToInvite.workspaceID,
        repoToInvite.repositoryID,
        account.accountID,
        role
      )
      .then(() => {
        setFilteredUserAccounts((state) =>
          state.map((el) =>
            el.accountID === account.accountID ? { ...el, role } : el
          )
        );
      })
      .catch((err) => {
        dispatchMessage(err?.response?.data);
      });
  };

  useEffect(() => {
    if (repoToInvite) {
      fetchAllUsersInWorkspace();
    }
    // eslint-disable-next-line
  }, [repoToInvite]);

  return (
    <Styled.Wrapper>
      <Styled.Content>
        <Styled.SubTitle>
          {t('REPOSITORIES_SCREEN.INVITE_USER_BELOW')}
        </Styled.SubTitle>

        <SearchBar
          placeholder={t('REPOSITORIES_SCREEN.SEARCH_USER_EMAIL_BELOW')}
          onSearch={(value) => onSearchUser(value)}
        />

        <Datatable
          columns={[
            {
              label: t('REPOSITORIES_SCREEN.ACTION'),
              property: 'action',
              type: 'custom',
              cssClass: ['flex-row-center'],
            },
            {
              label: t('REPOSITORIES_SCREEN.USER'),
              property: 'username',
              type: 'text',
            },
            {
              label: t('REPOSITORIES_SCREEN.EMAIL'),
              property: 'email',
              type: 'text',
            },
            {
              label: t('REPOSITORIES_SCREEN.PERMISSION'),
              property: 'permission',
              type: 'custom',
            },
            { label: '', property: 'help', type: 'custom' },
          ]}
          datasource={filteredUserAccounts.map((row) => {
            const repo: Datasource = {
              ...row,
              id: row.accountID,
              help: (
                <Styled.HelpIcon
                  name="help"
                  size="18px"
                  onClick={() => setPermissionsOpen(true)}
                />
              ),
              action: (
                <Checkbox
                  value={accountsInRepository.includes(row.accountID)}
                  disabled={row.email === currentUser?.email}
                  onChange={(_event, checked) => handleInviteUser(checked, row)}
                />
              ),
              permission: (
                <Select
                  disabled={
                    row.email === currentUser?.email ||
                    !accountsInRepository.includes(row.accountID)
                  }
                  value={row.role}
                  options={roles}
                  onChangeValue={(value) => handleChangeUserRole(value, row)}
                />
              ),
            };
            return repo;
          })}
          emptyListText={t('REPOSITORIES_SCREEN.NO_USERS_TO_INVITE')}
        />
      </Styled.Content>
      <Permissions
        isOpen={permissionsIsOpen}
        onClose={() => setPermissionsOpen(false)}
        rolesType="REPOSITORY"
      />
    </Styled.Wrapper>
  );
}

export default RepositoryInvite;
