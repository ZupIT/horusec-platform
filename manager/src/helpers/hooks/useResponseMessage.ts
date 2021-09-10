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
import React from 'react';
import { useTranslation } from 'react-i18next';
import { APIResponse } from 'helpers/interfaces/APIResponse';
import useFlashMessage from 'helpers/hooks/useFlashMessage';
import { ObjectLiteral } from 'helpers/interfaces/ObjectLiteral';

const useResponseMessage = () => {
  const { t } = useTranslation();
  const { showErrorFlash } = useFlashMessage();

  const dispatchMessage = React.useCallback(
    (response: APIResponse): void => {
      const messages: ObjectLiteral = {
        '{HORUSEC AUTH} invalid username or password': t(
          'API_ERRORS.ERROR_LOGIN'
        ),
        '{HORUSEC AUTH} account email not confirmed': t(
          'API_ERRORS.UNCONFIRMED_EMAIL'
        ),
        '{ACCOUNT} invalid reset password data': t(
          'API_ERRORS.INVALID_RECOVERY_CODE'
        ),
        '{ACCOUNT} email already in use': t('API_ERRORS.EMAIL_IN_USE'),
        '{ERROR_REPOSITORY} database not found records': t(
          'API_ERRORS.NO_RECORDS'
        ),
        '{ERROR_COMPANY} this account already in this company': t(
          'API_ERRORS.ALREADY_IN_COMPANY'
        ),
        'you do not have enough privileges for this action': t(
          'API_ERRORS.PRIVILEGES'
        ),
        '{ACCOUNT} username already in use': t('API_ERRORS.USERNAME_IN_USE'),
        '{CORE_REPOSITORY} repository name already in use': t(
          'API_ERRORS.REPO_NAME_IN_USE'
        ),
        '{CORE_REPOSITORY} this user does not belong to this workspace': t(
          'API_ERRORS.WORKSPACE_USER_NOT_BELONG'
        ),
        '{HORUSEC} webhook already exists to repository selected': t(
          'WEBHOOK_SCREEN.ALREADY_TO_REPOSITORY'
        ),
        '{CORE_WORKSPACE} not found user in workspace': t(
          'API_ERRORS.NOT_FOUND_USER'
        ),
        'email: must be a valid email address.': t('API_ERRORS.INVALID_EMAIL'),
        'admin ldap group should be a valid one for this user': t(
          'API_ERRORS.LDAP_GROUP'
        ),
        '{ACCOUNT} password is not valid': t('API_ERRORS.SAME_PASSWORD'),
        generic: t('API_ERRORS.GENERIC_ERROR'),
      };

      const msg = messages[response?.content] || messages.generic;

      showErrorFlash(msg);
    },
    [t, showErrorFlash]
  );

  return {
    dispatchMessage,
  };
};

export default useResponseMessage;
