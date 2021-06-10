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

import React, { useEffect, useRef } from 'react';
import Styled from './styled';
import { ObjectLiteral } from 'helpers/interfaces/ObjectLiteral';
import { get } from 'lodash';
import { useTranslation } from 'react-i18next';
import { Logout, Icon, Helper } from 'components';
import { useHistory } from 'react-router';

const Header: React.FC = () => {
  const { t } = useTranslation();
  const history = useHistory();
  const headerRef = useRef(null);

  const getTitleByURL = (): any => {
    const path = window.location.pathname;

    const titles: ObjectLiteral = {
      '/home/dashboard/repositories': {
        text: t('HEADER.TITLE.DASHBOARDREPOSITORY'),
        aria: t('DASHBOARD_SCREEN.ARIA_TITLE_REPOSITORY'),
        icon: 'pie',
        helper:
          'https://docs.horusec.io/docs/web/services/manager/introduction/#analytics-dashboard',
      },
      '/home/dashboard/workspace': {
        text: t('HEADER.TITLE.DASHBOARDWORKSPACE'),
        aria: t('DASHBOARD_SCREEN.ARIA_TITLE_WORKSPACE'),
        icon: 'pie',
        helper:
          'https://docs.horusec.io/docs/web/services/manager/introduction/#analytics-dashboard',
      },
      '/home/vulnerabilities': {
        text: t('HEADER.TITLE.VULNERABILITIES'),
        aria: t('HEADER.ARIA.VULNERABILITIES'),
        icon: 'shield',
        helper:
          'https://docs.horusec.io/docs/web/services/manager/vulnerabilities-management/',
      },
      '/home/repositories': {
        text: t('HEADER.TITLE.REPOSITORIES'),
        aria: t('HEADER.ARIA.REPOSITORIES'),
        icon: 'columns',
        helper:
          'https://docs.horusec.io/docs/web/services/manager/repository-management/',
      },
      '/home/webhooks': {
        text: t('HEADER.TITLE.WEBHOOKS'),
        aria: t('HEADER.ARIA.WEBHOOKS'),
        icon: 'webhook',
        helper:
          'https://docs.horusec.io/docs/web/services/manager/webhooks-management/',
      },
      '/home/workspaces': {
        text: t('HEADER.TITLE.WORKSPACES'),
        aria: t('HEADER.ARIA.WORKSPACES'),
        icon: 'grid',
        helper:
          'https://docs.horusec.io/docs/web/services/manager/workspace-management',
      },
      '/home/settings': {
        text: t('HEADER.TITLE.CONFIGURATION'),
        aria: t('HEADER.ARIA.CONFIGURATION'),
        icon: 'config',
        helper:
          'https://docs.horusec.io/docs/web/services/manager/account-management/',
      },
      '/home/add-workspace': {
        text: '',
        aria: t('HEADER.ARIA.ADD_WORKSPACE'),
        icon: 'zup',
        helper: 'https://docs.horusec.io/docs/web/overview/',
      },
      default: {
        text: '',
        aria: '',
        icon: 'zup',
        helper: 'https://docs.horusec.io/docs/web/overview/',
      },
    };

    let title = titles.default;

    Object.entries(titles).forEach((item) => {
      const [key, value] = item;
      if (path.includes(key)) title = value;
    });

    return title;
  };

  useEffect(() => {
    headerRef?.current?.focus();
  }, [history.location.pathname]);

  return (
    <Styled.Wrapper>
      <Styled.Header ref={headerRef} tabIndex={0}>
        <Styled.Title>
          <Icon name={getTitleByURL().icon} size="20px" />
          <Styled.Text aria-label={getTitleByURL().aria}>
            {getTitleByURL().text}
          </Styled.Text>
        </Styled.Title>

        <Styled.List tabIndex={0}>
          <Styled.Item>
            <Helper url={getTitleByURL().helper} />
          </Styled.Item>

          <Styled.Item
            tabIndex={0}
            aria-label={t('HEADER.ARIA.CONFIG')}
            active={history.location.pathname.includes('/home/settings')}
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

export default Header;
