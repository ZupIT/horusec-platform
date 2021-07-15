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

const Wrapper = styled.section`
  display: flex;
  flex-direction: column;
  width: 100%;
`;

const Content = styled.div`
  height: calc(100vh - 36px);
  width: 100%;
  padding: 3%;
`;

const Title = styled.h1`
  color: ${({ theme }) => theme.colors.chart.title};
  font-size: ${({ theme }) => theme.metrics.fontSize.title};
  letter-spacing: 1px;
`;

const SearchWrapper = styled.div`
  display: flex;
  align-items: center;
  background-color: ${({ theme }) => theme.colors.background.secundary};
  padding: 10px 15px;
  margin-top: 20px;
  border-radius: 5px;
`;

const ListWrapper = styled.div`
  background-color: ${({ theme }) => theme.colors.background.secundary};
  padding: 30px 20px;
  border-radius: 5px;
  margin-top: 30px;
`;

const Phrase = styled.h3`
  color: ${({ theme }) => theme.colors.text.primary};
  font-size: ${({ theme }) => theme.metrics.fontSize.large};
  margin-bottom: 20px;
  font-weight: normal;
`;

const Message = styled.div`
  display: flex;
  align-items: center;
  justify-content: center;
`;

const MessageText = styled(Phrase)`
  margin: 0;
  color: ${({ theme }) => theme.colors.text.secundary};
`;

const List = styled.ul`
  display: flex;
  margin-top: 40px;
  list-style: none;
  flex-wrap: wrap;
  max-height: 320px;
  overflow-y: auto;

  ::-webkit-scrollbar {
    width: 6px;
  }

  ::-webkit-scrollbar-thumb {
    background: ${({ theme }) => theme.colors.scrollbar};
    border-radius: 4px;
  }
`;

export default {
  Content,
  Wrapper,
  Title,
  SearchWrapper,
  ListWrapper,
  Phrase,
  Message,
  MessageText,
  List,
};
