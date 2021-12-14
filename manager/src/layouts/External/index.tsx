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
import HorusecLogo from 'assets/logos/horusec.svg';
import { Signature, Language, Icon } from 'components';
import { useTranslation } from 'react-i18next';
import packageJSON from '../../../package.json';

function ExternalLayout({ children }: { children: JSX.Element }) {
  const { repository, bugs } = packageJSON;
  const { t } = useTranslation();

  return (
    <Styled.Wrapper>
      <Styled.LogoContent>
        <h1>
          <Styled.Logo src={HorusecLogo} alt="Horusec Logo" />
        </h1>

        <Styled.Content>{children}</Styled.Content>
      </Styled.LogoContent>

      <Styled.Footer>
        <Language />

        <Signature />

        <Styled.ContactWrapper>
          <Styled.ContactItem href={repository.url} target="_blank">
            <Icon size="18px" name="github" />
            <Styled.ContactText>{t('GENERAL.GITHUB')}</Styled.ContactText>
          </Styled.ContactItem>

          <Styled.ContactItem href={bugs.url} target="_blank">
            <Icon size="18px" name="forum" />
            <Styled.ContactText>{t('GENERAL.FORUM')}</Styled.ContactText>
          </Styled.ContactItem>
        </Styled.ContactWrapper>
      </Styled.Footer>
    </Styled.Wrapper>
  );
}

export default ExternalLayout;
