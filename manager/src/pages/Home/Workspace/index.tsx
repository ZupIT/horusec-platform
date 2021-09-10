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

import React, { useState, useEffect, useRef } from 'react';
import Styled from './styled';
import { useTranslation } from 'react-i18next';
import { Button, HomeCard, Icon } from 'components';
import { SearchBar } from 'components';
import coreService from 'services/core';
import useResponseMessage from 'helpers/hooks/useResponseMessage';
import HandleWorkspace from './HandleWorkspace';
import HandleRepository from '../Workspace/HandleRepository';
import { Repository } from 'helpers/interfaces/Repository';
import { useHistory } from 'react-router-dom';
import usePermissions from 'helpers/hooks/usePermissions';
import useParamsRoute from 'helpers/hooks/useParamsRoute';
import { debounce } from 'lodash';

const Home: React.FC = () => {
  const PAGE_SIZE = 15;
  const loadMoreRef = useRef(null);

  const { dispatchMessage } = useResponseMessage();
  const { t } = useTranslation();
  const history = useHistory();
  const { ACTIONS, isAuthorizedAction } = usePermissions();
  const { workspaceId, workspace, getWorkspace } = useParamsRoute();

  const [repositories, setRepositories] = useState<Repository[]>([]);

  const [repositoryToEdit, setRepositoryToEdit] = useState<Repository>();

  const [currentPage, setCurrentPage] = useState(0);
  const [isFinishedItens, setFinishedItens] = useState(false);
  const [isSearch, setIsSearch] = useState(false);

  const [isOpenWorkspaceEditModal, setOpenWorkspaceEditModal] =
    useState<boolean>(false);
  const [isOpenRepositoryModal, setOpenRepositoryModal] =
    useState<boolean>(false);

  const fetchAllRepositoriesByWorkspace = (
    page?: number,
    clear?: boolean,
    search?: string
  ) => {
    if (!page) page = currentPage;

    setCurrentPage(page);

    if (page > 0) {
      if (clear) setRepositories([]);

      setIsSearch(!!search);

      coreService
        .getAllRepositories(workspaceId, page, search)
        .then((result) => {
          const listOfRepositories = result.data.content as Repository[];

          if (listOfRepositories.length < PAGE_SIZE) {
            setFinishedItens(true);

            if (!listOfRepositories.length) return;
          }

          if (search || clear) {
            setRepositories(listOfRepositories);
          } else {
            setRepositories([...repositories, ...listOfRepositories]);
          }
        })
        .catch((err) => {
          dispatchMessage(err?.response?.data);
          setRepositories([]);
        });
    }
  };

  const onSearch = debounce((search: string) => {
    setFinishedItens(false);
    fetchAllRepositoriesByWorkspace(1, true, search);
  }, 1000);

  useEffect(() => {
    if (!isSearch) fetchAllRepositoriesByWorkspace();
    // eslint-disable-next-line
  }, [currentPage]);

  useEffect(() => {
    const options: any = {
      root: null,
      rootMargin: '20px',
      threshold: 1.0,
    };

    const observer = new IntersectionObserver((entities) => {
      const target = entities[0];

      if (target.isIntersecting) {
        setCurrentPage((old) => old + 1);
      }
    }, options);

    if (loadMoreRef.current) {
      observer.observe(loadMoreRef.current);
    }

    // eslint-disable-next-line
  }, []);

  return (
    <>
      <Styled.Head>
        <Styled.TitleWrapper id="title-workspace-wrapper">
          <Styled.Title>
            <Styled.Icon name="grid" size="22px" />
            {workspace?.name}
          </Styled.Title>

          {isAuthorizedAction(ACTIONS.HANDLE_WORKSPACE) && (
            <Button
              text={t('HOME_SCREEN.HANDLER')}
              width="150px"
              outline
              rounded
              icon="tool"
              style={{ marginLeft: '30px' }}
              onClick={() => setOpenWorkspaceEditModal(true)}
            />
          )}

          {isAuthorizedAction(ACTIONS.VIEW_WORKSPACE) && (
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
          )}
        </Styled.TitleWrapper>

        <Button
          text={t('HOME_SCREEN.BACK_WORKSPACES')}
          ghost
          icon="back"
          onClick={() => history.push('/home')}
        />
      </Styled.Head>

      <Styled.Description>{workspace?.description}</Styled.Description>

      <Styled.SearchWrapper>
        <SearchBar
          onSearch={onSearch}
          placeholder={t('HOME_SCREEN.SEARCH_REPO')}
        />

        {isAuthorizedAction(ACTIONS.CREATE_REPOSITORY) && (
          <Button
            onClick={() => {
              setRepositoryToEdit(null);
              setOpenRepositoryModal(true);
            }}
            text={t('HOME_SCREEN.ADD_REPO')}
            pulsing={repositories.length <= 0}
            width="180px"
            icon="add"
            rounded
          />
        )}
      </Styled.SearchWrapper>

      <Styled.ListWrapper>
        <Styled.Phrase>{t('HOME_SCREEN.CHOOSE_REPO')}</Styled.Phrase>

        {repositories.length <= 0 ? (
          <Styled.Message>
            <Styled.MessageText>
              {t('HOME_SCREEN.EMPTY_REPOSITORIES')}
            </Styled.MessageText>
          </Styled.Message>
        ) : null}

        <Styled.List>
          {repositories.map((repo) => (
            <HomeCard
              repository={repo}
              key={repo.repositoryID}
              onHandle={() => {
                setOpenRepositoryModal(true);
                setRepositoryToEdit(repo);
              }}
              onOverview={() =>
                history.push(
                  `/overview/workspace/${repo.workspaceID}/repository/${repo.repositoryID}/dashboard`
                )
              }
            />
          ))}

          {!isFinishedItens && repositories.length >= 0 ? (
            <Styled.LoadingMessage ref={loadMoreRef}>
              <Icon name="loading" size="50px" />
              <Styled.MessageText>
                {t('HOME_SCREEN.LOADING_REPOSITORIES')}
              </Styled.MessageText>
            </Styled.LoadingMessage>
          ) : null}
        </Styled.List>
      </Styled.ListWrapper>

      <HandleWorkspace
        isVisible={isOpenWorkspaceEditModal}
        workspaceToEdit={workspace}
        onConfirm={() => {
          setOpenWorkspaceEditModal(false);
          getWorkspace();
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
          fetchAllRepositoriesByWorkspace(1, true);
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
