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

const Wrapper = styled.div`
  background-color: ${({ theme }) => theme.colors.dataTable.background};
  border-radius: 4px;
  height: auto;
  width: 100%;
  padding: 18px 15px 0px 15px;
  position: relative;
  grid-area: vulDetails;
`;

const Title = styled.h2`
  color: ${({ theme }) => theme.colors.dataTable.title};
  font-size: ${({ theme }) => theme.metrics.fontSize.xlarge};
  border-radius: 4px;
  padding: 0 10px 0px 10px;
  font-weight: normal;
  display: block;
  min-height: 60px;
`;

export default {
  Wrapper,
  Title,
};
