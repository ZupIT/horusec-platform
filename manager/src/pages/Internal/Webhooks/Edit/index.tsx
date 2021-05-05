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
import { Field } from 'helpers/interfaces/Field';
import coreService from 'services/core';
import { Repository } from 'helpers/interfaces/Repository';
import { get } from 'lodash';
import { isValidURL } from 'helpers/validators';
import useResponseMessage from 'helpers/hooks/useResponseMessage';
import webhookService from 'services/webhook';
import { Webhook, WebhookHeader } from 'helpers/interfaces/Webhook';
import useFlashMessage from 'helpers/hooks/useFlashMessage';
import useWorkspace from 'helpers/hooks/useWorkspace';
import { FieldArray, Formik, FormikHelpers } from 'formik';
import * as Yup from 'yup';
import SearchSelect from 'components/SearchSelect';

interface Props {
  isVisible: boolean;
  onCancel: () => void;
  onConfirm: () => void;
  webhookToEdit: Webhook;
}

const webhookHttpMethods = [{ value: 'POST' }, { value: 'GET' }];

const AddWebhook: React.FC<Props> = ({
  isVisible,
  onCancel,
  onConfirm,
  webhookToEdit,
}) => {
  const { t } = useTranslation();
  const { colors } = useTheme();
  const { currentWorkspace } = useWorkspace();

  const { dispatchMessage } = useResponseMessage();
  const { showSuccessFlash } = useFlashMessage();

  const [isLoading, setLoading] = useState(false);
  const [httpMethod, setHttpMethod] = useState(webhookToEdit?.method);
  const [repositories, setRepositories] = useState<Repository[]>([]);

  const handleConfirmSave = (
    values: InitialValue,
    action: FormikHelpers<InitialValue>
  ) => {
    setLoading(true);

    webhookService
      .update(
        currentWorkspace?.workspaceID,
        webhookToEdit?.webhookID,
        values.repositoryID,
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

  useEffect(() => {
    const fetchRepositories = () => {
      coreService
        .getAllRepositories(currentWorkspace?.workspaceID)
        .then((result) => {
          setRepositories(result.data.content);
        });
    };

    fetchRepositories();
  }, [currentWorkspace]);

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
    description: webhookToEdit?.description || '',
    url: webhookToEdit?.url || '',
    headers: webhookToEdit?.headers || [{ key: '', value: '' }],
    repositoryID: '',
    httpMethod: '',
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
      {(props) => {
        <Dialog
          isVisible={isVisible}
          message={t('WEBHOOK_SCREEN.EDIT')}
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

            <Styled.Label>{t('WEBHOOK_SCREEN.RESPOSITORY_LABEL')}</Styled.Label>

            <SearchSelect
              options={repositories.map((el) => ({
                label: el.name,
                value: el.repositoryID,
              }))}
              label={t('WEBHOOK_SCREEN.REPOSITORY')}
              width="100%"
              name="repositoryID"
            />

            <Styled.Label>{t('WEBHOOK_SCREEN.URL_LABEL')}</Styled.Label>

            <Styled.Wrapper>
              <SearchSelect
                label={t('WEBHOOK_SCREEN.TABLE.METHOD')}
                name="httpMethod"
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
        </Dialog>;
      }}
    </Formik>
  );
};

export default AddWebhook;
