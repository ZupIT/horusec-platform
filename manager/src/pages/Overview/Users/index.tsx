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
import useResponseMessage from 'helpers/hooks/useResponseMessage';
import useFlashMessage from 'helpers/hooks/useFlashMessage';
import { SearchBar, Button, Datatable, DataSource, Dialog } from 'components';
import { getCurrentUser } from 'helpers/localStorage/currentUser';
import { Account } from 'helpers/interfaces/Account';
import HandleUser from './Handle';
import coreService from 'services/core';
import { useParams } from 'react-router-dom';
import { RouteParams } from 'helpers/interfaces/RouteParams';

function Users() {
  const { t } = useTranslation();
  const { dispatchMessage } = useResponseMessage();
  const { showSuccessFlash } = useFlashMessage();
  const currentUser = getCurrentUser();
  const { workspaceId, repositoryId } = useParams<RouteParams>();

  const [users, setUsers] = useState<Account[]>([]);
  const [filteredUsers, setFilteredUsers] = useState<Account[]>([]);
  const [selectedUser, setSelectedUser] = useState<Account>(null);

  const [isLoading, setLoading] = useState(false);
  const [isVisibleHandleUser, setHandleUserVisible] = useState(false);
  const [isVisibleRemoveUser, setVisibleRemoveUser] = useState(false);

  const fetchUsersList = () => {
    setLoading(true);
    coreService
      .getUsers(workspaceId, repositoryId)
      .then((result) => {
        setUsers(result?.data?.content);
        setFilteredUsers(result?.data?.content);
      })
      .catch((err) => {
        dispatchMessage(err?.response?.data);
      })
      .finally(() => {
        setLoading(false);
      });
  };

  const handleConfirmRemoveUser = () => {
    setLoading(true);
    coreService
      .removeUser(workspaceId, repositoryId, selectedUser.accountID)
      .then(() => {
        showSuccessFlash(t('USERS_SCREEN.REMOVE_SUCCESS'));
        fetchUsersList();
      })
      .catch((err) => {
        dispatchMessage(err?.response?.data);
      })
      .finally(() => {
        setLoading(false);
        setVisibleRemoveUser(false);
        setSelectedUser(null);
      });
  };

  const onSearch = (search: string) => {
    if (search) {
      const filtered = users.filter(
        (user) =>
          user.email.toLocaleLowerCase().includes(search.toLocaleLowerCase()) ||
          user.username.toLocaleLowerCase().includes(search.toLocaleLowerCase())
      );
      setFilteredUsers(filtered);
    } else {
      setFilteredUsers(users);
    }
  };

  useEffect(() => {
    fetchUsersList();
    // eslint-disable-next-line
  }, []);

  return (
    <>
      <Styled.Wrapper>
        <Styled.Header>
          <SearchBar
            placeholder={t('USERS_SCREEN.SEARCH')}
            onSearch={onSearch}
          />

          <Button
            text={t('USERS_SCREEN.INVITE')}
            rounded
            width={180}
            icon="plus"
            onClick={() => setHandleUserVisible(true)}
          />
        </Styled.Header>

        <Styled.Content>
          <Datatable
            columns={[
              {
                label: t('USERS_SCREEN.TABLE.USER'),
                property: 'username',
                type: 'text',
              },
              {
                label: t('USERS_SCREEN.TABLE.EMAIL'),
                property: 'email',
                type: 'text',
              },
              {
                label: t('USERS_SCREEN.TABLE.PERMISSION'),
                property: 'permission',
                type: 'text',
              },
              {
                label: t('USERS_SCREEN.TABLE.ACTION'),
                property: 'permission',
                type: 'actions',
              },
            ]}
            dataSource={filteredUsers.map((row) => {
              const data: DataSource = {
                ...row,
                id: row.accountID,

                permission: t(
                  `USERS_SCREEN.ROLE.${row.role.toLocaleUpperCase()}`
                ),
                actions: [],
              };

              if (row.email !== currentUser?.email) {
                data.actions = [
                  {
                    title: t('USERS_SCREEN.TABLE.EDIT'),
                    icon: 'edit',
                    function: () => {
                      setSelectedUser(row);
                      setHandleUserVisible(true);
                    },
                  },
                  {
                    title: t('USERS_SCREEN.TABLE.REMOVE'),
                    icon: 'delete',
                    function: () => {
                      setSelectedUser(row);
                      setVisibleRemoveUser(true);
                    },
                  },
                ];
              }
              return data;
            })}
            isLoading={isLoading}
            emptyListText={t('USERS_SCREEN.TABLE.EMPTY')}
          />
        </Styled.Content>
      </Styled.Wrapper>

      <HandleUser
        isVisible={isVisibleHandleUser}
        onCancel={() => {
          setHandleUserVisible(false);
          setSelectedUser(null);
        }}
        onConfirm={() => {
          setHandleUserVisible(false);
          setSelectedUser(null);
          fetchUsersList();
        }}
        user={selectedUser}
      />

      <Dialog
        message={t('USERS_SCREEN.CONFIRM_REMOVE')}
        confirmText={t('USERS_SCREEN.YES')}
        loadingConfirm={isLoading}
        defaultButton
        hasCancel
        isVisible={isVisibleRemoveUser}
        onCancel={() => setVisibleRemoveUser(false)}
        onConfirm={handleConfirmRemoveUser}
      />
    </>
  );
}

export default Users;
