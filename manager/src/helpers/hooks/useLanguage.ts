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

import { useState, useEffect } from 'react';
import { useTranslation } from 'react-i18next';
import { Language as LanguageItem } from 'helpers/interfaces/Language';
import {
  getCurrentLanguage,
  setCurrentLanguage,
} from 'helpers/localStorage/currentLanguage';

import { Locale } from 'date-fns';
import enUS from 'date-fns/locale/en-US';
import ptBR from 'date-fns/locale/pt-BR';
import es from 'date-fns/locale/es';
import { get } from 'lodash';

const allLanguages: LanguageItem[] = [
  {
    name: 'en - US',
    i18nValue: 'enUS',
    htmlValue: 'en',
    icon: 'united-states',
    dateFormat: 'MM/dd/yyyy',
    description: 'English',
  },
  {
    name: 'pt - BR',
    i18nValue: 'ptBR',
    htmlValue: 'pt-BR',
    icon: 'brazil',
    dateFormat: 'dd/MM/yyyy',
    description: 'Português - Brasil',
  },
  {
    name: 'es',
    i18nValue: 'es',
    htmlValue: 'es',
    icon: 'spain',
    dateFormat: 'yyy/dd/mm',
    description: 'Español',
  },
];

const useLanguage = () => {
  const [currentLanguage, setLanguage] = useState(allLanguages[0]);
  const [currentLocale, setLocale] = useState<Locale>();
  const { i18n } = useTranslation();

  const handleLocale = (lang: LanguageItem) => {
    const locales = {
      enUS,
      ptBR,
      es,
    };

    setLocale(get(locales, lang.i18nValue, locales.enUS));
  };

  const setUserLanguage = (lang: LanguageItem) => {
    setLanguage(lang);
    setCurrentLanguage(lang);
    handleLocale(lang);

    i18n.changeLanguage(lang.i18nValue);
    window.document.documentElement.lang = lang.htmlValue;
  };

  useEffect(() => {
    const defaultLanguage = getCurrentLanguage();

    setUserLanguage(defaultLanguage || allLanguages[0]);

    // eslint-disable-next-line
  }, []);

  return {
    allLanguages,
    currentLanguage,
    setUserLanguage,
    currentLocale,
  };
};

export default useLanguage;
