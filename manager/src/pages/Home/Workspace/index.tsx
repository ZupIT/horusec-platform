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
import HandleWorkspace from './HandleWorkspace';
import HandleRepository from '../Workspace/HandleRepository';
import { Repository } from 'helpers/interfaces/Repository';
import { useParams, useHistory } from 'react-router-dom';
import { RouteParams } from 'helpers/interfaces/RouteParams';

const Home: React.FC = () => {
  const { dispatchMessage } = useResponseMessage();
  const { t } = useTranslation();
  const { workspaceId } = useParams<RouteParams>();
  const history = useHistory();

  const [currentWorkspace, setCurrentWorkspace] = useState<Workspace>();
  const [repositories, setRepositories] = useState<Repository[]>([]);
  const [filteredRepositories, setFilteredRepositories] =
    useState<Repository[]>(repositories);

  const [repositoryToEdit, setRepositoryToEdit] = useState<Repository>();

  const [isLoading, setLoading] = useState<boolean>(false);

  const [isOpenWorkspaceEditModal, setOpenWorkspaceEditModal] =
    useState<boolean>(false);
  const [isOpenRepositoryModal, setOpenRepositoryModal] =
    useState<boolean>(false);

  const fetchWorkspaceData = () => {
    setLoading(true);
    coreService
      .getOneWorkspace(workspaceId)
      .then((result) => {
        const workspace = result.data.content as Workspace;
        setCurrentWorkspace(workspace);
        fetchAllRepositoriesByWorkspace();
      })
      .catch((err) => {
        dispatchMessage(err?.response?.data);
        setLoading(false);
      });
  };

  const fetchAllRepositoriesByWorkspace = () => {
    setLoading(true);
    coreService
      .getAllRepositories(workspaceId)
      .then((result) => {
        const listOfRepositories = result.data.content as Repository[];
        setRepositories(listOfRepositories);
        setFilteredRepositories(listOfRepositories);
      })
      .catch((err) => {
        dispatchMessage(err?.response?.data);
        setRepositories([]);
      })
      .finally(() => {
        setLoading(false);
      });
  };

  const onSearch = (search: string) => {
    if (search) {
      const filtered = repositories.filter((repo) =>
        repo.name.toLocaleLowerCase().includes(search.toLocaleLowerCase())
      );

      setFilteredRepositories(filtered);
    } else {
      setFilteredRepositories(repositories);
    }
  };

  useEffect(() => {
    fetchWorkspaceData();
    // eslint-disable-next-line
  }, []);

  return (
    <>
      <Styled.Head>
        <Styled.TitleWrapper>
          <Styled.Title>
            <Styled.Icon name="grid" size="22px" />
            {currentWorkspace?.name}
          </Styled.Title>

          <Button
            text={t('HOME_SCREEN.HANDLER')}
            width="150px"
            outline
            rounded
            icon="tool"
            style={{ marginLeft: '30px' }}
            onClick={() => setOpenWorkspaceEditModal(true)}
          />

          <Button
            text={t('HOME_SCREEN.OVERVIEW')}
            width="150px"
            outline
            rounded
            icon="goto"
            style={{ marginLeft: '10px' }}
            onClick={() =>
              history.push(`/overview/workspace/${workspaceId}/dashboard`)
            }
          />
        </Styled.TitleWrapper>

        <Button
          text={t('HOME_SCREEN.BACK_WORKSPACES')}
          ghost
          icon="back"
          onClick={() => history.push('/home')}
        />
      </Styled.Head>

      <Styled.SearchWrapper>
        <SearchBar
          onSearch={onSearch}
          placeholder={t('HOME_SCREEN.SEARCH_REPO')}
        />

        <Button
          onClick={() => {
            setRepositoryToEdit(null);
            setOpenRepositoryModal(true);
          }}
          text={t('HOME_SCREEN.ADD_REPO')}
          pulsing={!isLoading && repositories.length <= 0}
          width="180px"
          icon="add"
          rounded
        />
      </Styled.SearchWrapper>

      <Styled.ListWrapper>
        <Styled.Phrase>{t('HOME_SCREEN.CHOOSE_REPO')}</Styled.Phrase>

        {isLoading ? (
          <Styled.Message>
            <Icon name="loading" size="50px" />
            <Styled.MessageText>
              {t('HOME_SCREEN.LOADING_REPOSITORIES')}
            </Styled.MessageText>
          </Styled.Message>
        ) : null}

        {!isLoading && filteredRepositories.length <= 0 ? (
          <Styled.Message>
            <Styled.MessageText>
              {t('HOME_SCREEN.EMPTY_REPOSITORIES')}
            </Styled.MessageText>
          </Styled.Message>
        ) : null}

        <Styled.List>
          {filteredRepositories.map((repo) => (
            <HomeCard
              repository={repo}
              key={repo.repositoryID}
              onHandle={() => {
                setOpenRepositoryModal(true);
                setRepositoryToEdit(repo);
              }}
              onOverview={() =>
                history.push(
                  `/overview/workspace/${workspaceId}/repository/${repo.repositoryID}/dashboard`
                )
              }
            />
          ))}
        </Styled.List>
      </Styled.ListWrapper>

      <HandleWorkspace
        isVisible={isOpenWorkspaceEditModal}
        workspaceToEdit={currentWorkspace}
        onConfirm={() => {
          setOpenWorkspaceEditModal(false);
          fetchWorkspaceData();
        }}
        onCancel={() => setOpenWorkspaceEditModal(false)}
      />

      <HandleRepository
        workspaceID={workspaceId}
        isVisible={isOpenRepositoryModal}
        repositoryToEdit={repositoryToEdit}
        onConfirm={() => {
          setOpenRepositoryModal(false);
          setRepositoryToEdit(null);
          fetchAllRepositoriesByWorkspace();
        }}
        onCancel={() => {
          setRepositoryToEdit(null);
          setOpenRepositoryModal(false);
        }}
      />
    </>
  );
};

export default Home;
