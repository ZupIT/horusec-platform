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
import { cloneDeep } from 'lodash';
import useWorkspace from 'helpers/hooks/useWorkspace';

interface Props {
  isVisible: boolean;
  webhookToCopy: Webhook;
  onCancel: () => void;
  onConfirm: () => void;
}

const webhookHttpMethods = [{ value: 'POST' }, { value: 'GET' }];

const AddWebhook: React.FC<Props> = ({
  isVisible,
  onCancel,
  onConfirm,
  webhookToCopy,
}) => {
  const { t } = useTranslation();
  const { colors } = useTheme();
  const { currentWorkspace } = useWorkspace();

  const { dispatchMessage } = useResponseMessage();
  const { showSuccessFlash } = useFlashMessage();

  const [isLoading, setLoading] = useState(false);
  const [httpMethod, setHttpMethod] = useState(webhookHttpMethods[0].value);
  const [repositories, setRepositories] = useState<Repository[]>([]);
  const [selectedRepository, setSelectedRepository] = useState<Repository>();
  const [url, setUrl] = useState<Field>({ value: '', isValid: false });
  const [headers, setHeaders] = useState<WebhookHeader[]>([
    { key: '', value: '' },
  ]);
  const [description, setDescription] = useState<Field>({
    value: '',
    isValid: false,
  });

  const resetFields = () => {
    setHeaders([{ key: '', value: '' }]);
    setSelectedRepository(null);
    setDescription({ isValid: false, value: '' });
    setUrl({ isValid: false, value: '' });
  };

  const handleCancel = () => {
    resetFields();
    onCancel();
  };

  const handleConfirmSave = () => {
    setLoading(true);

    webhookService
      .create(
        currentWorkspace?.workspaceID,
        selectedRepository.repositoryID,
        url.value,
        httpMethod,
        headers,
        description.value
      )
      .then(() => {
        showSuccessFlash(t('WEBHOOK_SCREEN.SUCCESS_CREATE'));
        resetFields();
        onConfirm();
      })
      .catch((err) => {
        dispatchMessage(err?.response?.data);
      })
      .finally(() => {
        setLoading(false);
      });
  };

  const handleSetHeader = (index: number, key: string, value: string) => {
    const headersCopy = cloneDeep(headers);
    const header = { key, value };
    headersCopy[index] = header;
    setHeaders(headersCopy);
  };

  const handleRemoveHeader = () => {
    const headersCopy = cloneDeep(headers);
    headersCopy.pop();
    setHeaders(headersCopy);
  };

  useEffect(() => {
    setHeaders(webhookToCopy?.headers || [{ key: '', value: '' }]);
    setDescription({ value: webhookToCopy?.description, isValid: true });
    setUrl({ value: webhookToCopy?.url, isValid: true });
  }, [webhookToCopy]);

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

  return (
    <Dialog
      isVisible={isVisible}
      message={t('WEBHOOK_SCREEN.ADD')}
      onCancel={handleCancel}
      onConfirm={handleConfirmSave}
      confirmText={t('WEBHOOK_SCREEN.SAVE')}
      disableConfirm={!url.isValid || !selectedRepository}
      disabledColor={colors.button.disableInDark}
      loadingConfirm={isLoading}
      width={600}
      hasCancel
    >
      <Styled.Form>
        <Styled.Label>{t('WEBHOOK_SCREEN.DESCRIPTION_LABEL')}</Styled.Label>

        <Styled.Field
          label={t('WEBHOOK_SCREEN.DESCRIPTION')}
          onChangeValue={(field: Field) => setDescription(field)}
          name="description"
          type="text"
          width="100%"
          initialValue={description.value}
        />

        <Styled.Label>{t('WEBHOOK_SCREEN.RESPOSITORY_LABEL')}</Styled.Label>

        <Select
          width="100%"
          options={repositories.map((el) => ({ label: el.name, value: el }))}
          label={t('WEBHOOK_SCREEN.REPOSITORY')}
          value={selectedRepository}
          onChangeValue={(value) => setSelectedRepository(value)}
        />

        <Styled.Label>{t('WEBHOOK_SCREEN.URL_LABEL')}</Styled.Label>

        <Styled.Wrapper>
          <Styled.URLSelect
            width="100px"
            value={httpMethod}
            options={webhookHttpMethods.map((el) => ({
              label: el.value,
              value: el.value,
            }))}
            onChangeValue={(item) => setHttpMethod(item)}
            disabled
            color={get(colors.methods, httpMethod.toLocaleLowerCase())}
          />

          <Styled.Field
            label={t('WEBHOOK_SCREEN.URL')}
            onChangeValue={(field: Field) => setUrl(field)}
            name="url"
            type="text"
            width="400px"
            validation={isValidURL}
            invalidMessage={t('WEBHOOK_SCREEN.INVALID_URL')}
            initialValue={url.value}
          />
        </Styled.Wrapper>

        <Styled.Label>{t('WEBHOOK_SCREEN.HEADERS_LABEL')}</Styled.Label>

        {headers.map((header, index) => (
          <Styled.Wrapper key={index}>
            <Styled.Field
              label={t('WEBHOOK_SCREEN.KEY')}
              name={`key-${index}`}
              onChangeValue={({ value }) =>
                handleSetHeader(index, value, headers[index].value)
              }
              width="200px"
              initialValue={headers[index]?.key}
            />

            <Styled.Field
              label={t('WEBHOOK_SCREEN.VALUE')}
              name={`value-${index}`}
              onChangeValue={({ value }) =>
                handleSetHeader(index, headers[index].key, value)
              }
              width="200px"
              initialValue={headers[index]?.value}
            />

            {index + 1 === headers.length && headers.length !== 1 ? (
              <Styled.OptionIcon
                name="delete"
                size="20px"
                onClick={handleRemoveHeader}
              />
            ) : null}

            {index + 1 === headers.length && headers.length !== 5 ? (
              <Styled.OptionIcon
                name="plus"
                size="20px"
                onClick={() => setHeaders([...headers, { key: '', value: '' }])}
              />
            ) : null}
          </Styled.Wrapper>
        ))}
      </Styled.Form>
    </Dialog>
  );
};

export default AddWebhook;
