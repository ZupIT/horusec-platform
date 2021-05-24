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

import React from 'react';
import { useTranslation } from 'react-i18next';
import Styled from './styled';

interface HelperInterface {
  url: string;
}

const Helper: React.FC<HelperInterface> = ({ url }) => {
  const { t } = useTranslation();

  return (
    <Styled.Link
      tabIndex={0}
      aria-label={t('HEADER.ARIA.HELP')}
      href={url}
      target="_blank"
    >
      <Styled.HelpIcon name="help" size="20px" />

      <Styled.Text>{t('SIDE_MENU.HELPER')}</Styled.Text>
    </Styled.Link>
  );
};

export default Helper;
