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

import http from 'config/axios';
import axios from 'axios';
import { SERVICE_AUTH } from '../config/endpoints';
import {
  setCurrentUser,
  clearCurrentUser,
} from 'helpers/localStorage/currentUser';
import { AxiosResponse, AxiosError } from 'axios';
import {
  clearTokens,
  getRefreshToken,
  setTokens,
} from 'helpers/localStorage/tokens';
import { LoginParams } from 'helpers/interfaces/LoginParams';

const login = (params: LoginParams): Promise<AxiosResponse<any>> => {
  return http.post(`${SERVICE_AUTH}/auth/authenticate/login`, params);
};

const logout = (refreshToken: string): Promise<AxiosResponse<any>> =>
  http.post(`${SERVICE_AUTH}/auth/account/logout`, { refreshToken });

const createAccount = (
  username: string,
  password: string,
  email: string
): Promise<AxiosResponse<any>> => {
  return http.post(`${SERVICE_AUTH}/auth/account/create-account-horusec`, {
    username,
    email,
    password,
  });
};

const update = (
  username: string,
  email: string
): Promise<AxiosResponse<any>> => {
  return http.patch(`${SERVICE_AUTH}/auth/account/update`, {
    username,
    email,
  });
};

const deleteAccount = (): Promise<AxiosResponse<any>> => {
  return http.delete(`${SERVICE_AUTH}/auth/account/delete`);
};

const createAccountFromKeycloak = (
  accessToken: string
): Promise<AxiosResponse<any>> => {
  return http.post(`${SERVICE_AUTH}/auth/account/create-account-keycloak`, {
    accessToken,
  });
};

const sendCode = (email: string): Promise<AxiosResponse<any>> => {
  return http.post(`${SERVICE_AUTH}/auth/account/send-reset-code`, { email });
};

const validateCode = (
  email: string,
  code: string
): Promise<AxiosResponse<any>> => {
  return http.post(`${SERVICE_AUTH}/auth/account/check-reset-code`, {
    email,
    code,
  });
};

const changePassword = (
  token: string,
  password: string
): Promise<AxiosResponse<any>> => {
  return http.post(
    `${SERVICE_AUTH}/auth/account/change-password`,
    { password },
    {
      headers: {
        'X-Horusec-Authorization': `Bearer ${token}`,
      },
    }
  );
};

const updatePassword = (password: string): Promise<AxiosResponse<any>> => {
  return http.post(`${SERVICE_AUTH}/auth/account/change-password`, {
    password,
  });
};

const verifyUniqueUsernameEmail = (
  email: string,
  username: string
): Promise<AxiosResponse<any>> => {
  return http.post(`${SERVICE_AUTH}/auth/account/verify-already-used`, {
    email,
    username,
  });
};

const callRenewToken = async (): Promise<string> => {
  const refreshToken = getRefreshToken();

  const handleLogout = () => {
    clearCurrentUser();
    clearTokens();
    window.location.replace('/auth');
  };

  if (refreshToken) {
    return new Promise((resolve, reject) => {
      axios
        .post(`${SERVICE_AUTH}/auth/account/refresh-token`, { refreshToken })
        .then((result: AxiosResponse<any>) => {
          const {
            username,
            isApplicationAdmin,
            email,
            expiresAt,
            refreshToken,
            accessToken,
          } = result.data.content;

          setTokens(accessToken, refreshToken, expiresAt);
          setCurrentUser({ username, isApplicationAdmin, email });

          resolve(accessToken);
        })
        .catch((err: AxiosError) => {
          reject(err);
          handleLogout();
        });
    });
  } else {
    handleLogout();
  }
};

const getHorusecConfig = (): Promise<AxiosResponse<any>> => {
  return axios.get(`${SERVICE_AUTH}/auth/authenticate/config`);
};

export default {
  login,
  logout,
  createAccount,
  sendCode,
  validateCode,
  changePassword,
  callRenewToken,
  verifyUniqueUsernameEmail,
  getHorusecConfig,
  createAccountFromKeycloak,
  update,
  deleteAccount,
  updatePassword,
};
