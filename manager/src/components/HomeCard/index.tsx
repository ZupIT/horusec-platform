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
import { Repository } from 'helpers/interfaces/Repository';
import { Workspace } from 'helpers/interfaces/Workspace';
import React from 'react';
import { useTranslation } from 'react-i18next';
import { useHistory } from 'react-router-dom';
import Styled from './styled';
import usePermissions from 'helpers/hooks/usePermissions';
import { getCurrentConfig } from 'helpers/localStorage/horusecConfig';
import { authTypes } from 'helpers/enums/authTypes';

interface Props {
  workspace?: Workspace;
  repository?: Repository;
  onHandle?: () => void;
  onOverview?: () => void;
}

const HomeCard: React.FC<Props> = ({
  workspace,
  repository,
  onHandle,
  onOverview,
}) => {
  const { t } = useTranslation();
  const history = useHistory();

  const { ACTIONS, isAuthorizedAction } = usePermissions();

  const isWorkspace = !!workspace;
  const { authType } = getCurrentConfig();

  const context = workspace || repository;

  return (
    <Styled.Card>
      <Styled.Title>
        <Styled.Icon name={isWorkspace ? 'grid' : 'columns'} size="24px" />
        {context.name}
      </Styled.Title>

      {context?.description ? (
        <Styled.Description>{context?.description}</Styled.Description>
      ) : null}

      <Styled.Info>
        {isWorkspace && authType !== authTypes.LDAP ? (
          <Styled.InfoItem>
            <Styled.InfoIcon size="14px" name="columns" />
            {workspace?.repositoriesCount || 0} {t('HOME_SCREEN.REPOSITORIES')}
          </Styled.InfoItem>
        ) : null}

        <Styled.InfoItem>
          <Styled.InfoIcon size="14px" name="user" />
          {t(`HOME_SCREEN.${context.role.toUpperCase()}`)}
        </Styled.InfoItem>
      </Styled.Info>

      <Styled.OptionsBar>
        {isWorkspace ? (
          <Styled.Option
            onClick={() =>
              history.push(`/home/workspace/${context.workspaceID}`)
            }
          >
            <Styled.InfoIcon name="check" size="16px" />
            {t('HOME_SCREEN.SELECT')}
          </Styled.Option>
        ) : (
          <>
            {isAuthorizedAction(ACTIONS.HANDLE_REPOSITORY, context) && (
              <Styled.Option onClick={onHandle}>
                <Styled.InfoIcon name="tool" size="16px" />
                {t('HOME_SCREEN.HANDLER')}
              </Styled.Option>
            )}

            <Styled.Option onClick={onOverview}>
              <Styled.InfoIcon name="goto" size="16px" />
              {t('HOME_SCREEN.OVERVIEW')}
            </Styled.Option>
          </>
        )}
      </Styled.OptionsBar>
    </Styled.Card>
  );
};

export default HomeCard;
