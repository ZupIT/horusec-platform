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

import styled, { css } from 'styled-components';

interface TitlesProps {
  isDanger?: boolean;
}

interface LanguageProps {
  active?: boolean;
}

const Wrapper = styled.section`
  padding: 35px 15px;
  width: 100%;
`;

const Content = styled.div`
  margin-top: 25px;
  padding: 25px 15px 10px 25px;
  background-color: ${({ theme }) => theme.colors.background.secundary};
  border-radius: 4px;
  position: relative;
`;

const Title = styled.h1<TitlesProps>`
  color: ${({ theme }) => theme.colors.text.secundary};
  font-weight: normal;
  font-size: ${({ theme }) => theme.metrics.fontSize.xlarge};

  ${({ isDanger }) =>
    isDanger &&
    css`
      color: ${({ theme }) => theme.colors.input.error};
    `};
`;

const Subtitle = styled.span`
  display: block;
  margin: 20px 0;
  color: ${({ theme }) => theme.colors.text.secundary};
  font-weight: normal;
  font-size: ${({ theme }) => theme.metrics.fontSize.medium};
`;

const BtnsWrapper = styled.div`
  display: flex;

  button {
    margin-right: 15px;
  }
`;

const LanguageList = styled.ul`
  list-style: none;
  margin: 0;
  padding: 0;
  cursor: pointer;
  display: flex;
  align-items: center;
`;

const LanguageItem = styled.li<LanguageProps>`
  display: flex;
  align-items: center;
  margin-right: 20px;
  border: 1px solid ${({ theme }) => theme.colors.secondary};
  border-radius: 3px;
  padding: 3px 8px;
  min-width: 130px;

  ${({ active }) =>
    active &&
    css`
      border: 1px solid ${({ theme }) => theme.colors.button.secundary};
      span {
        color: ${({ theme }) => theme.colors.button.secundary};
      }
    `};

  :hover {
    border: 1px solid ${({ theme }) => theme.colors.button.secundary};
    span {
      color: ${({ theme }) => theme.colors.button.secundary};
    }
  }
`;

const LanguageName = styled.span`
  display: block;
  margin-left: 10px;
  color: ${({ theme }) => theme.colors.secondary};
`;

export default {
  Wrapper,
  Content,
  Title,
  Subtitle,
  BtnsWrapper,
  LanguageList,
  LanguageItem,
  LanguageName,
};
