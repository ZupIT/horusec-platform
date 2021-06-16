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

import styled from 'styled-components';

const AriaTitle = styled.h1`
  font-size: 0;
`;

const Wrapper = styled.div`
  width: 100%;
  margin: 10px auto;
  padding: 0 15px;

  overflow-y: scroll;

  display: grid;

  column-gap: 10px;
  row-gap: 15px;

  grid-template-columns: 1fr 3fr;

  grid-template-areas:
    'filters filters'
    'allVul allVul'
    'totalDevelopers vulByDeveloper'
    'totalRepositories vulByRepository'
    'vulByLanguage vulByLanguage'
    'vulTimeline vulTimeline'
    'vulDetails vulDetails';

  ::-webkit-scrollbar {
    width: 10px;
  }

  ::-webkit-scrollbar-thumb {
    background: ${({ theme }) => theme.colors.background.highlight};
    border-radius: 2px;
  }

  ::-webkit-scrollbar-track {
    background-color: ${({ theme }) => theme.colors.scrollbar};
  }
`;

const FilterWrapper = styled.div`
  display: flex;
  flex-wrap: nowrap;
  align-items: center;
  align-content: center;
  background-color: ${({ theme }) => theme.colors.background.secundary};
  border-radius: 4px;
  grid-area: filters;
`;

export default { Wrapper, AriaTitle, FilterWrapper };
