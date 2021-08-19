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

import React from 'react';
import { SideMenu, Footer, Header } from 'components';
import Styled from './styled';
import { keycloakInstance } from 'config/keycloak';
import { clearTokens } from 'helpers/localStorage/tokens';
import { clearCurrentUser } from 'helpers/localStorage/currentUser';
import { WorkspaceProvider } from 'contexts/Workspace';
import { RepositoryProvider } from 'contexts/Repository';

function InternalLayout({ children }: { children: JSX.Element }) {
  keycloakInstance.onAuthRefreshError = () => {
    clearTokens();
    clearCurrentUser();
    keycloakInstance.logout();
  };

  return (
    <WorkspaceProvider>
      <RepositoryProvider>
        <>
          <Styled.Wrapper>
            <SideMenu />

            <Styled.Content>
              <Styled.HeaderWrapper>
                <Header />
              </Styled.HeaderWrapper>
              {children}
            </Styled.Content>
          </Styled.Wrapper>

          <Footer />
        </>
      </RepositoryProvider>
    </WorkspaceProvider>
  );
}

export default InternalLayout;
