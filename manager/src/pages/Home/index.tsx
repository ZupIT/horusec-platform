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
import { Header } from 'components';
import { Switch, useRouteMatch } from 'react-router-dom';
import { PrivateRoute } from 'components';
import Workspaces from './Welcome';
import Repositories from './Workspace';
import { WorkspaceProvider } from 'contexts/Workspace';

const Home: React.FC = () => {
  const { path } = useRouteMatch();

  return (
    <WorkspaceProvider>
      <Styled.Wrapper>
        <Header />

        <Styled.Content>
          <Switch>
            <PrivateRoute exact path={path} component={() => <Workspaces />} />

            <PrivateRoute
              exact
              path={`${path}/workspace/:workspaceId`}
              component={() => <Repositories />}
            />
          </Switch>
        </Styled.Content>
      </Styled.Wrapper>
    </WorkspaceProvider>
  );
};

export default Home;
