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

import React, { Suspense, lazy } from 'react';
import { BrowserRouter, Redirect, Route, Switch } from 'react-router-dom';
import { isLogged } from 'helpers/localStorage/tokens';
import { PrivateRoute } from 'components';
import { MANAGER_BASE_PATH } from 'config/basePath';

const Routes = () => (
  <BrowserRouter basename={MANAGER_BASE_PATH}>
    <Suspense fallback="">
      <Switch>
        <Route exact path="/">
          {isLogged() ? <Redirect to="/home" /> : <Redirect to="/auth" />}
        </Route>

        <Route path="/auth" component={lazy(() => import('pages/Auth'))} />

        <PrivateRoute
          exact={false}
          path="/home"
          component={lazy(() => import('pages/Home'))}
        />

        <PrivateRoute
          exact={false}
          path="/overview"
          component={lazy(() => import('pages/Overview'))}
        />

        <PrivateRoute
          path="/settings"
          exact={true}
          component={lazy(() => import('pages/Settings'))}
        />

        <Route component={lazy(() => import('pages/NotFound'))} />
      </Switch>
    </Suspense>
  </BrowserRouter>
);

export default Routes;
