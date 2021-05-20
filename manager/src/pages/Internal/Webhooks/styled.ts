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

interface TagProps {
  color: string;
}

const Wrapper = styled.section`
  padding: 35px;
  width: 100%;
`;

const Options = styled.div`
  background-color: ${({ theme }) => theme.colors.background.secundary};
  border-radius: 4px;
  padding: 22px;
  display: flex;
  align-items: center;
`;

const Content = styled.div`
  margin-top: 25px;
  padding: 25px 15px 10px 25px;
  background-color: ${({ theme }) => theme.colors.background.secundary};
  border-radius: 4px;
`;

const Tag = styled.span<TagProps>`
  display: block;
  text-align: center;
  text-transform: uppercase;
  height: 23px;
  line-height: 25px;
  width: 84px;
  border-radius: 64px;

  ${({ color }) => css`
    background-color: ${color};
  `};
`;

export default {
  Wrapper,
  Options,
  Tag,
  Content,
};
