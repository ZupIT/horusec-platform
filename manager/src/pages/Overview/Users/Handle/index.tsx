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
import { Dialog, Permissions } from 'components';
import { useTranslation } from 'react-i18next';
import Styled from './styled';
import { useTheme } from 'styled-components';
import coreService from 'services/core';
import useResponseMessage from 'helpers/hooks/useResponseMessage';
import useFlashMessage from 'helpers/hooks/useFlashMessage';
import * as Yup from 'yup';
import { Formik } from 'formik';
import SearchSelect from 'components/SearchSelect';
import { RouteParams } from 'helpers/interfaces/RouteParams';
import { useParams } from 'react-router-dom';
import { Account } from 'helpers/interfaces/Account';

interface Props {
  isVisible: boolean;
  onCancel: () => void;
  onConfirm: () => void;
  user?: Account;
}

interface Role {
  label: string;
  value: string;
}

const HandleUser: React.FC<Props> = ({
  isVisible,
  onCancel,
  onConfirm,
  user,
}) => {
  const { t } = useTranslation();
  const { colors } = useTheme();
  const { dispatchMessage } = useResponseMessage();
  const { showSuccessFlash } = useFlashMessage();
  const { workspaceId, repositoryId } = useParams<RouteParams>();
  const [usersOfWorkspace, setUsersWorkspace] = useState([]);
  const [isLoading, setLoading] = useState(false);
  const [permissionsIsOpen, setPermissionsIsOpen] = useState(false);

  const isRepository = !!repositoryId;
  const isEditing = !!user;

  const getRoles = () => {
    const roles: Role[] = [
      {
        label: t('PERMISSIONS.ADMIN'),
        value: 'admin',
      },
      {
        label: t('PERMISSIONS.MEMBER'),
        value: 'member',
      },
    ];

    if (isRepository)
      roles.push({
        label: t('PERMISSIONS.SUPERVISOR'),
        value: 'supervisor',
      });

    return roles;
  };

  const ValidationScheme = Yup.object({
    email: Yup.lazy(() => {
      if (!isRepository) {
        return Yup.string().required();
      }
      return Yup.string().notRequired();
    }),
    role: Yup.string().oneOf(['admin', 'member', 'supervisor']).required(),
    user: Yup.lazy(() => {
      if (isRepository) {
        return Yup.object({
          accountID: Yup.string(),
          email: Yup.string(),
          username: Yup.string(),
          role: Yup.string(),
        }).required();
      }
      return Yup.object().notRequired();
    }),
  });

  type InitialValue = Yup.InferType<typeof ValidationScheme>;

  const initialValues: InitialValue = {
    email: user?.email || '',
    role: user?.role || '',
    user: user || {},
  };

  const renderUserField = () => {
    return isRepository ? (
      <Styled.RoleWrapper>
        <SearchSelect
          options={usersOfWorkspace}
          label={t('USERS_SCREEN.HANDLE_MODAL.USER')}
          name="user"
          width="350px"
        />
      </Styled.RoleWrapper>
    ) : (
      <Styled.Field
        label={t('USERS_SCREEN.HANDLE_MODAL.EMAIL')}
        name="email"
        type="text"
      />
    );
  };

  const handleInviteUser = (value: any, actions: any) => {
    coreService
      .inviteUser(
        workspaceId,
        value.user?.email || value.email,
        value.role,
        repositoryId,
        value.user?.accountID,
        value.user?.username
      )
      .then(() => {
        showSuccessFlash(t('USERS_SCREEN.HANDLE_MODAL.SUCCESS'));
        onConfirm();
        actions.resetForm();
      })
      .catch((err) => {
        dispatchMessage(err?.response?.data);
      })
      .finally(() => {
        setLoading(false);
      });
  };

  const handleEditUser = (value: any, actions: any) => {
    coreService
      .updateUserRole(workspaceId, repositoryId, user.accountID, value.role)
      .then(() => {
        showSuccessFlash(t('USERS_SCREEN.HANDLE_MODAL.SUCCESS'));
        onConfirm();
        actions.resetForm();
      })
      .catch((err) => {
        dispatchMessage(err?.response?.data);
      })
      .finally(() => {
        setLoading(false);
      });
  };

  useEffect(() => {
    const fetchAllUsersInWorkspace = () => {
      coreService
        .getUsers(workspaceId, null)
        .then((result) => {
          const users = result?.data?.content?.map((item: Account) => {
            return {
              label: `${item.username} - ${item.email}`,
              value: item,
            };
          });

          setUsersWorkspace(users || []);
        })
        .catch((err) => {
          dispatchMessage(err?.response?.data);
        });
    };

    if (isRepository) fetchAllUsersInWorkspace();
    // eslint-disable-next-line
  }, []);

  return (
    <Formik
      initialValues={initialValues}
      validationSchema={ValidationScheme}
      enableReinitialize
      onSubmit={(value, actions) => {
        setLoading(true);

        isEditing
          ? handleEditUser(value, actions)
          : handleInviteUser(value, actions);
      }}
    >
      {(props) => (
        <Dialog
          isVisible={isVisible}
          message={
            isEditing
              ? t('USERS_SCREEN.HANDLE_MODAL.TITLE_EDIT')
              : t('USERS_SCREEN.HANDLE_MODAL.TITLE')
          }
          onCancel={() => {
            onCancel();
            props.resetForm();
          }}
          onConfirm={props.submitForm}
          confirmText={t('USERS_SCREEN.HANDLE_MODAL.SAVE')}
          disableConfirm={!props.isValid}
          disabledColor={colors.button.disableInDark}
          loadingConfirm={isLoading}
          width={450}
          hasCancel
        >
          <Styled.SubTitle>
            {isEditing
              ? t('USERS_SCREEN.HANDLE_MODAL.EDIT_SUBTITLE')
              : t('USERS_SCREEN.HANDLE_MODAL.INVITE_SUBTITLE')}
          </Styled.SubTitle>

          {isEditing ? (
            <Styled.SubTitle>
              {user.username} - {user.email}
            </Styled.SubTitle>
          ) : (
            renderUserField()
          )}

          <Styled.RoleWrapper>
            <SearchSelect
              options={getRoles()}
              label={t('USERS_SCREEN.HANDLE_MODAL.ROLE')}
              name="role"
              width="350px"
            />

            <Styled.HelpIcon
              name="help"
              size="20px"
              onClick={() => setPermissionsIsOpen(true)}
            />
          </Styled.RoleWrapper>

          <Permissions
            isOpen={permissionsIsOpen}
            onClose={() => setPermissionsIsOpen(false)}
            rolesType={isRepository ? 'REPOSITORY' : 'COMPANY'}
          />
        </Dialog>
      )}
    </Formik>
  );
};

export default HandleUser;
