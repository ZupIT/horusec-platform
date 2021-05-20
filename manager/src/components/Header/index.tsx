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
import Styled from './styled';
import { ObjectLiteral } from 'helpers/interfaces/ObjectLiteral';
import { get } from 'lodash';
import { useTranslation } from 'react-i18next';
import { Logout, Icon, Helper } from 'components';
import { useHistory } from 'react-router';

const Footer: React.FC = () => {
  const { t } = useTranslation();
  const history = useHistory();

  const getTitleByURL = () => {
    const path = window.location.pathname;

    const titles: ObjectLiteral = {
      '/home/dashboard/repositories': {
        text: t('HEADER.TITLE.DASHBOARDREPOSITORY'),
        icon: 'pie',
        helper:
          'https://docs.horusec.io/docs/web/services/manager/introduction/#analytics-dashboard',
      },
      '/home/dashboard/workspace': {
        text: t('HEADER.TITLE.DASHBOARDWORKSPACE'),
        icon: 'pie',
        helper:
          'https://docs.horusec.io/docs/web/services/manager/introduction/#analytics-dashboard',
      },
      '/home/vulnerabilities': {
        text: t('HEADER.TITLE.VULNERABILITIES'),
        icon: 'shield',
        helper:
          'https://docs.horusec.io/docs/web/services/manager/vulnerabilities-management/',
      },
      '/home/repositories': {
        text: t('HEADER.TITLE.REPOSITORIES'),
        icon: 'columns',
        helper:
          'https://docs.horusec.io/docs/web/services/manager/repository-management/',
      },
      '/home/webhooks': {
        text: t('HEADER.TITLE.WEBHOOKS'),
        icon: 'webhook',
        helper:
          'https://docs.horusec.io/docs/web/services/manager/webhooks-management/',
      },
      '/home/workspaces': {
        text: t('HEADER.TITLE.WORKSPACES'),
        icon: 'grid',
        helper:
          'https://docs.horusec.io/docs/web/services/manager/workspace-management',
      },
      '/home/settings': {
        text: t('HEADER.TITLE.CONFIGURATION'),
        icon: 'config',
        helper:
          'https://docs.horusec.io/docs/web/services/manager/account-management/',
      },
      '/home/add-workspace': {
        text: '',
        icon: 'zup',
        helper: 'https://docs.horusec.io/docs/web/overview/',
      },
    };

    return get(titles, path, t('HEADER.TITLE.DEFAULT'));
  };

  return (
    <Styled.Wrapper>
      <Styled.Header>
        <Styled.Title>
          <Icon name={getTitleByURL().icon} size="20px" />
          <Styled.Text>{getTitleByURL().text}</Styled.Text>
        </Styled.Title>

        <Styled.List>
          <Styled.Item>
            <Helper url={getTitleByURL().helper} />
          </Styled.Item>

          <Styled.Item
            active={window.location.pathname === '/home/settings'}
            onClick={() => history.replace('/home/settings')}
          >
            <Styled.ConfigIcon name="config" size="15" />

            <Styled.ConfigText>{t('SIDE_MENU.CONFIG')}</Styled.ConfigText>
          </Styled.Item>

          <Styled.Item>
            <Logout />
          </Styled.Item>
        </Styled.List>
      </Styled.Header>
    </Styled.Wrapper>
  );
};

export default Footer;
