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

import React, { useState } from 'react';
import { Calendar, Checkbox, Dialog } from 'components';
import { useTranslation } from 'react-i18next';
import Styled from './styled';
import { useTheme } from 'styled-components';
import coreService from 'services/core';
import useResponseMessage from 'helpers/hooks/useResponseMessage';

import SuccessAddToken from './Success';
import validateExpiresAt from 'helpers/validators/validateExpiresAt';
import { Formik, FormikHelpers } from 'formik';
import * as Yup from 'yup';
import { RouteParams } from 'helpers/interfaces/RouteParams';

interface Props {
  isVisible: boolean;
  onCancel: () => void;
  onConfirm: () => void;
  currentParams: RouteParams;
}

const MIN_DATE = new Date(Date.now() + 86400000);

const AddToken: React.FC<Props> = ({
  isVisible,
  onCancel,
  onConfirm,
  currentParams,
}) => {
  const { t } = useTranslation();
  const { colors } = useTheme();
  const { dispatchMessage } = useResponseMessage();

  const [isLoading, setLoading] = useState(false);
  const [tokenCreated, setTokenCreated] = useState<string>(null);

  const ValidationScheme = Yup.object({
    description: Yup.string().required(),
    isExpirable: Yup.boolean().optional(),
    expiresAt: Yup.date()
      .test('boolean', t('TOKENS_SCREEN.INVALID_EXPIRES_AT'), (date) =>
        validateExpiresAt(date.toString())
      )
      .optional(),
  });

  type InitialValue = Yup.InferType<typeof ValidationScheme>;

  const initialValues: InitialValue = {
    description: '',
    isExpirable: false,
    expiresAt: MIN_DATE,
  };

  const handleConfirmSave = (
    values: InitialValue,
    actions: FormikHelpers<InitialValue>
  ) => {
    setLoading(true);

    const data = {
      description: values.description,
      isExpirable: values.isExpirable,
      expiresAt: values.expiresAt ? new Date(values.expiresAt) : null,
    };

    if (values.isExpirable === false) {
      delete data.isExpirable;
      delete data.expiresAt;
    }

    coreService
      .createToken(
        {
          workspaceID: currentParams.workspaceId,
          repositoryID: currentParams.repositoryId,
        },
        data
      )
      .then((res) => {
        onConfirm();
        actions.resetForm();
        setTokenCreated(res?.data?.content);
      })
      .catch((err) => {
        dispatchMessage(err?.response?.data);
      })
      .finally(() => {
        setLoading(false);
      });
  };

  return (
    <>
      <Formik
        enableReinitialize
        initialValues={initialValues}
        onSubmit={handleConfirmSave}
        validationSchema={ValidationScheme}
      >
        {(props) => (
          <Dialog
            isVisible={isVisible}
            message={t('TOKENS_SCREEN.CREATE_NEW_TOKEN')}
            onCancel={() => {
              onCancel();
              props.resetForm();
            }}
            onConfirm={props.submitForm}
            confirmText={t('TOKENS_SCREEN.SAVE')}
            disableConfirm={!props.isValid}
            disabledColor={colors.button.disableInDark}
            loadingConfirm={isLoading}
            width={600}
            defaultButton
            hasCancel
          >
            <Styled.SubTitle>
              {t('TOKENS_SCREEN.CREATE_TOKEN_BELOW')}
            </Styled.SubTitle>

            <Styled.Form>
              <Styled.Field
                label={t('TOKENS_SCREEN.DESCRIPTION')}
                name="description"
              />

              <Styled.ContainerCheckbox>
                <Checkbox
                  name="isExpirable"
                  label={t('TOKENS_SCREEN.IS_EXPIRABLE')}
                />
              </Styled.ContainerCheckbox>

              {props.values.isExpirable && (
                <Calendar
                  name="expiresAt"
                  title={t('TOKENS_SCREEN.EXPIRES_AT')}
                  minDate={MIN_DATE}
                />
              )}
            </Styled.Form>
          </Dialog>
        )}
      </Formik>
      {tokenCreated && (
        <SuccessAddToken
          tokenValue={tokenCreated}
          onConfirm={() => setTokenCreated(null)}
        />
      )}
    </>
  );
};

export default AddToken;
