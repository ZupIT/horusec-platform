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

import React, { useState } from 'react';
import { useTranslation } from 'react-i18next';
import { Dialog, Button, Icon } from 'components';
import Styled from './styled';
import { clearCurrentUser } from 'helpers/localStorage/currentUser';
import accountService from 'services/auth';
import { Header } from 'components';
import EditAccount from './Edit';
import ChangePassword from './ChangePassword';
import useResponseMessage from 'helpers/hooks/useResponseMessage';
import useFlashMessage from 'helpers/hooks/useFlashMessage';
import { clearTokens } from 'helpers/localStorage/tokens';
import { useHistory } from 'react-router-dom';
import useLanguage from 'helpers/hooks/useLanguage';
import { getCurrentConfig } from 'helpers/localStorage/horusecConfig';
import { authTypes } from 'helpers/enums/authTypes';

const Settings: React.FC = () => {
  const { t } = useTranslation();
  const { dispatchMessage } = useResponseMessage();
  const { showSuccessFlash } = useFlashMessage();
  const { allLanguages, currentLanguage, setUserLanguage } = useLanguage();
  const history = useHistory();
  const { authType } = getCurrentConfig();

  const [deleteDialogIsOpen, setOpenDeleteDialog] = useState(false);
  const [deleteInProgress, setDeleteInProgress] = useState(false);

  const [editDialogIsOpen, setOpenEditDialog] = useState(false);
  const [changePassDialogIsOpen, setOpenChangePassDialog] = useState(false);

  const handleConfirmDelete = () => {
    setDeleteInProgress(true);
    accountService
      .deleteAccount()
      .then(() => {
        history.replace('/auth');
        clearCurrentUser();
        clearTokens();
        showSuccessFlash(t('SETTINGS_SCREEN.SUCCESS_DELETE'));
      })
      .catch((err) => {
        dispatchMessage(err?.response?.data);
      });
  };

  return (
    <>
      <Header />

      <Styled.Wrapper>
        <Styled.Content>
          <Styled.Title>{t('SETTINGS_SCREEN.LANGUAGE')}</Styled.Title>

          <Styled.Subtitle>
            {t('SETTINGS_SCREEN.CHANGE_LANGUAGE')}
          </Styled.Subtitle>

          <Styled.LanguageList>
            {allLanguages.map((language, index) => (
              <Styled.LanguageItem
                key={index}
                onClick={() => setUserLanguage(language)}
                id={language.i18nValue}
                active={currentLanguage?.name === language.name}
              >
                <Icon name={language.icon} size="30px" />

                <Styled.LanguageName>
                  {language.description}
                </Styled.LanguageName>
              </Styled.LanguageItem>
            ))}
          </Styled.LanguageList>
        </Styled.Content>

        {authType !== authTypes.KEYCLOAK ? (
          <>
            <Styled.Content>
              <Styled.Title>{t('SETTINGS_SCREEN.TITLE')}</Styled.Title>

              <Styled.Subtitle>
                {t('SETTINGS_SCREEN.ACCOUNT_SUBTITLE')}
              </Styled.Subtitle>

              <Styled.BtnsWrapper>
                <Button
                  text={t('SETTINGS_SCREEN.CHANGE_USER_DATA')}
                  icon="edit"
                  onClick={() => setOpenEditDialog(true)}
                  outline
                />

                <Button
                  text={t('SETTINGS_SCREEN.CHANGE_PASS')}
                  icon="lock"
                  onClick={() => setOpenChangePassDialog(true)}
                  outline
                />
              </Styled.BtnsWrapper>
            </Styled.Content>

            <Styled.Content>
              <Styled.Title isDanger>
                {t('SETTINGS_SCREEN.DELETE_ACCOUNT')}
              </Styled.Title>

              <Styled.Subtitle>
                {t('SETTINGS_SCREEN.SURE_DELETE')}
              </Styled.Subtitle>

              <Button
                text={t('SETTINGS_SCREEN.DELETE')}
                onClick={() => setOpenDeleteDialog(true)}
                outline
              />
            </Styled.Content>
          </>
        ) : null}

        <Dialog
          message={t('SETTINGS_SCREEN.CONFIRM_DELETE')}
          confirmText={t('SETTINGS_SCREEN.YES')}
          loadingConfirm={deleteInProgress}
          defaultButton
          hasCancel
          isVisible={deleteDialogIsOpen}
          onCancel={() => setOpenDeleteDialog(false)}
          onConfirm={handleConfirmDelete}
        />

        <EditAccount
          isVisible={editDialogIsOpen}
          onCancel={() => setOpenEditDialog(false)}
          onConfirm={() => setOpenEditDialog(false)}
        />

        <ChangePassword
          isVisible={changePassDialogIsOpen}
          onCancel={() => setOpenChangePassDialog(false)}
          onConfirm={() => setOpenChangePassDialog(false)}
        />
      </Styled.Wrapper>
    </>
  );
};

export default Settings;
