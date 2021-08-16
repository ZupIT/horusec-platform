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

import React, { useState, useEffect } from 'react';
import Styled from './styled';
import { useTranslation } from 'react-i18next';
import { Button, HomeCard, Icon } from 'components';
import { SearchBar } from 'components';
import { Workspace } from 'helpers/interfaces/Workspace';
import coreService from 'services/core';
import useResponseMessage from 'helpers/hooks/useResponseMessage';
import HandleWorkspace from '../Workspace/HandleWorkspace';

const Welcome: React.FC = () => {
  const { dispatchMessage } = useResponseMessage();
  const { t } = useTranslation();

  const [allWorkspaces, setAllWorkspaces] = useState<Workspace[]>([]);
  const [filteredWorkspaces, setFilteredWorkspaces] =
    useState<Workspace[]>(allWorkspaces);

  const [isLoading, setLoading] = useState<boolean>(false);
  const [isOpenAddWorkspaceModal, setOpenAddWorkspaceModal] =
    useState<boolean>(false);

  const onSearch = (search: string) => {
    if (search) {
      const filtered = allWorkspaces.filter((workspace) =>
        workspace.name.toLocaleLowerCase().includes(search.toLocaleLowerCase())
      );

      setFilteredWorkspaces(filtered);
    } else {
      setFilteredWorkspaces(allWorkspaces);
    }
  };

  const fetchAllWorkspaces = () => {
    setLoading(true);
    coreService
      .getAllWorkspaces()
      .then((result) => {
        const workspaces = (result?.data?.content as Workspace[]) || [];
        setAllWorkspaces(workspaces);
        setFilteredWorkspaces(workspaces);
      })
      .catch((err) => {
        dispatchMessage(err?.response?.data);
        setAllWorkspaces([]);
        setFilteredWorkspaces([]);
      })
      .finally(() => {
        setLoading(false);
      });
  };

  useEffect(() => {
    fetchAllWorkspaces();
    // eslint-disable-next-line
  }, []);

  return (
    <>
      <Styled.Title>{t('HOME_SCREEN.WELCOME')}</Styled.Title>

      <Styled.Subtitle>{t('HOME_SCREEN.MY_WORKSPACES')}</Styled.Subtitle>

      <Styled.SearchWrapper>
        <SearchBar
          onSearch={onSearch}
          placeholder={t('HOME_SCREEN.SEARCH_WORKSPACE')}
        />

        <Button
          text={t('HOME_SCREEN.ADD_WORKSPACE')}
          icon="add"
          rounded
          id="addWorkspace"
          width="180px"
          pulsing={!isLoading && allWorkspaces.length <= 0}
          onClick={() => setOpenAddWorkspaceModal(true)}
        />
      </Styled.SearchWrapper>

      <Styled.ListWrapper>
        <Styled.Phrase>{t('HOME_SCREEN.CHOOSE_WORKSPACE')}</Styled.Phrase>

        {isLoading ? (
          <Styled.Message>
            <Icon name="loading" size="50px" />
            <Styled.MessageText>
              {t('HOME_SCREEN.LOADING_WORKSPACES')}
            </Styled.MessageText>
          </Styled.Message>
        ) : null}

        {!isLoading && filteredWorkspaces.length <= 0 ? (
          <Styled.Message>
            <Styled.MessageText>
              {t('HOME_SCREEN.EMPTY_WORKSPACES')}
            </Styled.MessageText>
          </Styled.Message>
        ) : null}

        <Styled.List>
          {filteredWorkspaces.map((workspace) => (
            <HomeCard workspace={workspace} key={workspace.workspaceID} />
          ))}
        </Styled.List>
      </Styled.ListWrapper>

      <HandleWorkspace
        isVisible={isOpenAddWorkspaceModal}
        onConfirm={() => {
          setOpenAddWorkspaceModal(false);
          fetchAllWorkspaces();
        }}
        onCancel={() => setOpenAddWorkspaceModal(false)}
      />
    </>
  );
};

export default Welcome;
