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
import { Dialog, Select } from 'components';
import { useTranslation } from 'react-i18next';
import { useTheme } from 'styled-components';
import Styled from './styled';
import coreService from 'services/core';
import { Repository } from 'helpers/interfaces/Repository';
import { get } from 'lodash';
import { isValidURL } from 'helpers/validators';
import useResponseMessage from 'helpers/hooks/useResponseMessage';
import webhookService from 'services/webhook';
import { Webhook, WebhookHeader } from 'helpers/interfaces/Webhook';
import useFlashMessage from 'helpers/hooks/useFlashMessage';
import { FieldArray, Formik, FormikHelpers } from 'formik';
import SearchSelect from 'components/SearchSelect';
import * as Yup from 'yup';
import { Workspace } from 'helpers/interfaces/Workspace';
import { useParams } from 'react-router-dom';
import { RouteParams } from 'helpers/interfaces/RouteParams';
interface Props {
  isVisible: boolean;
  isNew: boolean;
  webhookInitial: Webhook;
  onCancel: () => void;
  onConfirm: () => void;
}

const webhookHttpMethods = [{ value: 'POST' }, { value: 'GET' }];

const HandleWebhook: React.FC<Props> = ({
  isVisible,
  isNew,
  onCancel,
  onConfirm,
  webhookInitial,
}) => {
  const { t } = useTranslation();
  const { colors } = useTheme();

  const [currentWorkspace, setCurrentWorkspace] = useState<Workspace>(null);
  const [allRepositories, setAllRepositories] = useState<Repository[]>([]);

  const { dispatchMessage } = useResponseMessage();
  const { showSuccessFlash } = useFlashMessage();

  const [isLoading, setLoading] = useState(false);

  const { workspaceId, repositoryId } = useParams<RouteParams>();

  useEffect(() => {
    function getCurrentWorkspace() {
      coreService.getOneWorkspace(workspaceId).then((result) => {
        setCurrentWorkspace(result.data.content);
      });
    }
    function loadRepositories() {
      coreService.getAllRepositories(workspaceId).then((result) => {
        setAllRepositories(result.data.content);
      });
    }

    if (workspaceId) {
      getCurrentWorkspace();
      loadRepositories();
    }
  }, [workspaceId]);

  const updateWebhook = (
    values: InitialValue,
    action: FormikHelpers<InitialValue>
  ) => {
    setLoading(true);

    webhookService
      .update(
        currentWorkspace?.workspaceID,
        values.repositoryID,
        webhookInitial.webhookID,
        values.url,
        values.httpMethod,
        values.headers,
        values.description
      )
      .then(() => {
        showSuccessFlash(t('WEBHOOK_SCREEN.SUCCESS_UPDATE'));
        action.resetForm();
        onConfirm();
      })
      .catch((err) => {
        dispatchMessage(err?.response?.data);
      })
      .finally(() => {
        setLoading(false);
      });
  };

  const createWebhook = (
    values: InitialValue,
    action: FormikHelpers<InitialValue>
  ) => {
    setLoading(true);

    webhookService
      .create(
        currentWorkspace?.workspaceID,
        values.repositoryID,
        values.url,
        values.httpMethod,
        values.headers,
        values.description
      )
      .then(() => {
        showSuccessFlash(t('WEBHOOK_SCREEN.SUCCESS_CREATE'));
        action.resetForm();
        onConfirm();
      })
      .catch((err) => {
        dispatchMessage(err?.response?.data);
      })
      .finally(() => {
        setLoading(false);
      });
  };

  const handleConfirmSave = (
    values: InitialValue,
    action: FormikHelpers<InitialValue>
  ) => {
    if (isNew) createWebhook(values, action);
    else updateWebhook(values, action);
  };

  const ValidationScheme = Yup.object({
    description: Yup.string().optional(),
    url: Yup.string()
      .test('validUrl', t('WEBHOOK_SCREEN.INVALID_URL'), isValidURL)
      .required(t('WEBHOOK_SCREEN.INVALID_URL')),
    headers: Yup.array<WebhookHeader[]>().required(),
    repositoryID: Yup.string().required(),
    httpMethod: Yup.string().required(),
  });

  type InitialValue = Yup.InferType<typeof ValidationScheme>;

  const initialValues: InitialValue = {
    description: webhookInitial?.description || '',
    url: webhookInitial?.url || '',
    headers: webhookInitial?.headers || [{ key: '', value: '' }],
    repositoryID: webhookInitial?.repositoryID || repositoryId || '',
    httpMethod: webhookInitial?.method || 'POST',
  };

  return (
    <Formik
      initialValues={initialValues}
      enableReinitialize={true}
      validationSchema={ValidationScheme}
      onSubmit={(values, actions) => {
        setLoading(true);
        handleConfirmSave(values, actions);
      }}
    >
      {(props) => (
        <Dialog
          isVisible={isVisible}
          message={isNew ? t('WEBHOOK_SCREEN.ADD') : t('WEBHOOK_SCREEN.EDIT')}
          onCancel={() => {
            props.resetForm();
            onCancel();
          }}
          onConfirm={props.submitForm}
          confirmText={t('WEBHOOK_SCREEN.SAVE')}
          disableConfirm={!props.isValid}
          disabledColor={colors.button.disableInDark}
          loadingConfirm={isLoading}
          width={600}
          hasCancel
        >
          <Styled.Form>
            <Styled.Label>{t('WEBHOOK_SCREEN.DESCRIPTION_LABEL')}</Styled.Label>

            <Styled.Field
              label={t('WEBHOOK_SCREEN.DESCRIPTION')}
              name="description"
              width="100%"
            />

            {!repositoryId ? (
              <>
                <Styled.Label>
                  {t('WEBHOOK_SCREEN.REPOSITORY_LABEL')}
                </Styled.Label>

                <SearchSelect
                  options={allRepositories.map((el) => ({
                    label: el.name,
                    value: el.repositoryID,
                  }))}
                  label={t('WEBHOOK_SCREEN.REPOSITORY')}
                  width="100%"
                  name="repositoryID"
                />
              </>
            ) : null}

            <Styled.Label>{t('WEBHOOK_SCREEN.URL_LABEL')}</Styled.Label>

            <Styled.Wrapper>
              <SearchSelect
                label={t('WEBHOOK_SCREEN.TABLE.METHOD')}
                name="httpMethod"
                isDisabled={true}
                width="50%"
                options={webhookHttpMethods.map((el) => ({
                  label: el.value,
                  value: el.value,
                }))}
                style={{
                  backgroundColor: get(colors.methods, props.values.httpMethod),
                }}
              />

              <Styled.Field
                label={t('WEBHOOK_SCREEN.URL')}
                name="url"
                width="100%"
              />
            </Styled.Wrapper>

            <Styled.Label>{t('WEBHOOK_SCREEN.HEADERS_LABEL')}</Styled.Label>

            <FieldArray name="headers">
              {({ push, remove }) => {
                const { headers } = props.values;

                return headers.map((header, index) => (
                  <Styled.Wrapper key={index}>
                    <Styled.Field
                      label={t('WEBHOOK_SCREEN.KEY')}
                      name={`headers.${index}.key`}
                      width="200px"
                    />

                    <Styled.Field
                      label={t('WEBHOOK_SCREEN.VALUE')}
                      name={`headers.${index}.value`}
                      width="200px"
                    />

                    {index + 1 === headers.length && headers.length !== 1 ? (
                      <Styled.OptionIcon
                        name="delete"
                        size="20px"
                        onClick={() => remove(index)}
                      />
                    ) : null}

                    {index + 1 === headers.length && headers.length !== 5 ? (
                      <Styled.OptionIcon
                        name="plus"
                        size="20px"
                        onClick={() => push({ key: '', value: '' })}
                      />
                    ) : null}
                  </Styled.Wrapper>
                ));
              }}
            </FieldArray>
          </Styled.Form>
        </Dialog>
      )}
    </Formik>
  );
};

export default HandleWebhook;
