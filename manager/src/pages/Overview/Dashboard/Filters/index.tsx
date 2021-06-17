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

import React, { useEffect } from 'react';
import Styled from './styled';
import { useTranslation } from 'react-i18next';
import { Calendar } from 'components';
import { FilterValues } from 'helpers/interfaces/FilterValues';
import useWorkspace from 'helpers/hooks/useWorkspace';
import { ObjectLiteral } from 'helpers/interfaces/ObjectLiteral';
import { Formik, FormikProps } from 'formik';
import * as Yup from 'yup';
import SearchSelect from 'components/SearchSelect';
import useRepository from 'helpers/hooks/useRepository';
interface FilterProps {
  onApply: (values: FilterValues) => void;
  type: 'workspace' | 'repository';
}

const Filters: React.FC<FilterProps> = ({ type, onApply }) => {
  const { t } = useTranslation();
  const { currentWorkspace } = useWorkspace();
  const { currentRepository, setCurrentRepository, allRepositories } =
    useRepository();

  const formikRef = React.createRef<FormikProps<FilterValues>>();

  const fixedRanges = [
    {
      label: t('DASHBOARD_SCREEN.BEGINNING'),
      value: 'beginning',
    },
    {
      label: t('DASHBOARD_SCREEN.CUSTOM_RANGE'),
      value: 'customRange',
    },
    {
      label: t('DASHBOARD_SCREEN.TODAY'),
      value: 'today',
    },
    {
      label: t('DASHBOARD_SCREEN.LAST_WEEK'),
      value: 'lastWeek',
    },
    {
      label: t('DASHBOARD_SCREEN.LAST_MONTH'),
      value: 'lastMonth',
    },
  ];

  const today = new Date();
  const lastWeek = new Date(new Date().setDate(today.getDate() - 7));
  const lastMonth = new Date(new Date().setDate(today.getDate() - 30));
  const firstYear = new Date(2019, 1, 1);

  const ValidationScheme = Yup.object({
    period: Yup.string().label(t('DASHBOARD_SCREEN.PERIOD')).notRequired(),
    initialDate: Yup.string()
      .label(t('DASHBOARD_SCREEN.START_DATE'))
      .notRequired()
      .nullable(),
    finalDate: Yup.string()
      .label(t('DASHBOARD_SCREEN.FINAL_DATE'))
      .notRequired()
      .nullable(),
    repositoryID: Yup.string()
      .label(t('DASHBOARD_SCREEN.REPOSITORY'))
      .when('type', {
        is: 'repository',
        then: Yup.string().required(),
      }),
    workspaceID: Yup.string().required(),
    type: Yup.string().oneOf(['workspace', 'repository']).required(),
  });

  useEffect(() => {
    function setValues() {
      formikRef.current.setFieldValue(
        'repositoryID',
        currentRepository?.repositoryID
      );
      formikRef.current.setFieldValue(
        'workspaceID',
        currentWorkspace?.workspaceID
      );
    }
    setValues();
  }, [
    currentRepository?.repositoryID,
    currentWorkspace?.workspaceID,
    formikRef,
  ]);

  useEffect(() => {
    if (type === 'repository' && !currentRepository?.repositoryID) {
      return;
    } else {
      onApply(initialValues);
    }
    // eslint-disable-next-line
  }, [currentWorkspace, currentRepository]);

  const getRangeOfPeriod: ObjectLiteral = {
    beginning: [firstYear, today],
    customRange: [today, today],
    today: [today, today],
    lastWeek: [lastWeek, today],
    lastMonth: [lastMonth, today],
  };

  const initialValues: FilterValues = {
    period: fixedRanges[0].value,
    initialDate: getRangeOfPeriod[fixedRanges[0].value][0],
    finalDate: getRangeOfPeriod[fixedRanges[0].value][1],
    repositoryID: currentRepository?.repositoryID,
    workspaceID: currentWorkspace?.workspaceID,
    type,
  };

  return (
    <Formik
      initialValues={initialValues}
      validationSchema={ValidationScheme}
      innerRef={formikRef}
      onSubmit={(values) => {
        if (values.period !== 'customRange') {
          values.initialDate = getRangeOfPeriod[values.period][0];
          values.finalDate = getRangeOfPeriod[values.period][1];
        } else {
          values.initialDate = new Date(values.initialDate);
          values.finalDate = new Date(values.finalDate);
        }
        if (values?.repositoryID) setCurrentRepository(values.repositoryID);
        onApply(values);
      }}
    >
      {(props) => (
        <Styled.Container
          tabIndex={0}
          aria-label={t('DASHBOARD_SCREEN.ARIA_FILTER')}
        >
          <Styled.Wrapper>
            <SearchSelect
              name="period"
              label={t('DASHBOARD_SCREEN.PERIOD')}
              options={fixedRanges}
            />
          </Styled.Wrapper>

          <div
            style={{
              display:
                props.values.period === fixedRanges[1].value ? 'flex' : 'none',
            }}
          >
            <Styled.CalendarWrapper>
              <Calendar
                name="initialDate"
                label={t('DASHBOARD_SCREEN.START_DATE')}
              />
            </Styled.CalendarWrapper>

            <Styled.CalendarWrapper>
              <Calendar
                name="finalDate"
                label={t('DASHBOARD_SCREEN.FINAL_DATE')}
              />
            </Styled.CalendarWrapper>
          </div>
          {type === 'repository' ? (
            <Styled.Wrapper>
              <SearchSelect
                name="repositoryID"
                label={t('DASHBOARD_SCREEN.REPOSITORY')}
                options={allRepositories.map((el) => ({
                  label: el.name,
                  value: el.repositoryID,
                }))}
              />
            </Styled.Wrapper>
          ) : null}
          <Styled.ApplyButton
            text={t('DASHBOARD_SCREEN.APPLY')}
            rounded
            width={130}
            height={38}
            type="submit"
          />
        </Styled.Container>
      )}
    </Formik>
  );
};

export default React.memo(Filters);
