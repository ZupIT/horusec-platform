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
import Styled from './styled';
import useAuth from 'helpers/hooks/useAuth';
import { useTranslation } from 'react-i18next';
import { getRefreshToken } from 'helpers/localStorage/tokens';

const Logout: React.FC = () => {
  const { logout } = useAuth();
  const { t } = useTranslation();
  const refreshToken = getRefreshToken();

  return (
    <Styled.Button
      aria-label={t('HEADER.ARIA.LOGOUT')}
      tabIndex={0}
      onClick={() => logout(refreshToken)}
    >
      <Styled.LogoutIcon size="16px" name="logout" />

      <Styled.Text>{t('SIDE_MENU.LOGOUT')}</Styled.Text>
    </Styled.Button>
  );
};

export default Logout;
