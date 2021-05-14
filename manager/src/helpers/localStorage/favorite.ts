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

import { localStorageKeys } from 'helpers/enums/localStorageKeys';
import { get } from 'lodash';
import { getCurrentUser } from './currentUser';

const getAllFavorites = () => {
  try {
    const localData = window.localStorage.getItem(localStorageKeys.FAVORITE);
    return localData ? JSON.parse(localData) : {};
  } catch (e) {
    return {};
  }
};

const getFavoriteWorkspace = () => {
  try {
    const { email } = getCurrentUser();
    const allFavorites = getAllFavorites();
    return allFavorites ? get(allFavorites, email, null) : null;
  } catch (e) {
    return null;
  }
};

const setFavoriteWorkspace = (workspaceID: string) => {
  const allFavorites = getAllFavorites();
  const { email } = getCurrentUser();

  allFavorites[email] = workspaceID;

  window.localStorage.setItem(
    localStorageKeys.FAVORITE,
    JSON.stringify(allFavorites)
  );
};

export { getFavoriteWorkspace, setFavoriteWorkspace };
