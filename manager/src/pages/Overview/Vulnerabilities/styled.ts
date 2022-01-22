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

import styled, { css } from 'styled-components';
import { SearchBar as SearchBarComp } from 'components';

interface LangueProps {
  color: string;
}

interface SearchBarProps {
  isSearching: boolean;
}

interface VulDetailProps {
  isOpen: boolean;
}

const Wrapper = styled.div`
  padding: 35px 15px;
  width: 100%;
  height: 95%;
  display: flex;
  flex-direction: column;
`;

const Options = styled.div`
  background-color: ${({ theme }) => theme.colors.background.secundary};
  border-radius: 4px;
  padding: 22px;
  display: flex;
  align-items: center;
`;

const SearchWrapper = styled.div<SearchBarProps>`
  width: 390px;
  transition: ease all 0.7s;

  ${({ isSearching }) =>
    isSearching &&
    css`
      width: 100%;
    `};
`;

const SelectsWrapper = styled.div<SearchBarProps>`
  width: 100%;
  transition: ease all 0.7s;
  display: flex;

  ${({ isSearching }) =>
    isSearching &&
    css`
      margin: 0;
      padding: 0;
      width: 0;

      .filter {
        width: 0;
        opacity: 0;
        pointer-events: none;
      }
    `};
`;

const Content = styled.div`
  background-color: ${({ theme }) => theme.colors.background.secundary};
  border-radius: 4px;
  padding: 22px;
  margin-top: 15px;
  height: 75%;
  display: flex;
  flex-direction: column;
`;

const ScrollList = styled.ul`
  overflow-y: auto;
  margin-bottom: 10px;
  height: 90%;
  list-style: none;
  padding-right: 10px;

  ::-webkit-scrollbar {
    width: 10px;
  }

  ::-webkit-scrollbar-thumb {
    background: ${({ theme }) => theme.colors.background.primary};
    border-radius: 2px;
  }

  ::-webkit-scrollbar-track {
    background-color: ${({ theme }) => theme.colors.background.highlight};
  }
`;

const File = styled.li`
  background-color: ${({ theme }) => theme.colors.background.highlight};
  border-radius: 2px;
  width: 100%;
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  padding: 10px 20px;
  margin-bottom: 12px;

  :hover {
    box-shadow: 0 4px 6px rgba(33, 33, 33, 0.8);
    cursor: pointer;

    div.view {
      width: 18px;
      opacity: 1;
    }
  }
`;

const FileColumn = styled.div`
  display: flex;
  align-items: flex-start;
  justify-content: center;
  flex-direction: column;
`;

const FileRow = styled.div`
  display: flex;
  align-items: center;
  flex-direction: row;
  justify-content: space-between;
`;

const FileLanguage = styled.div<LangueProps>`
  display: block;
  margin-right: 15px;
  font-weight: bold;
  font-size: ${({ theme }) => theme.metrics.fontSize.xsmall};
  color: ${({ color }) => color};
  margin-right: 20px;
`;

const FileName = styled.div`
  display: block;
  color: ${({ theme }) => theme.colors.text.opaque};
`;

const FileVulCount = styled.div`
  display: block;
  color: ${({ theme }) => theme.colors.text.primary};
  background: #a30d0033;
  border: 1px solid ${({ theme }) => theme.colors.flashMessage.error};
  padding: 4px 15px;
  width: 20px;
  height: 20px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: ${({ theme }) => theme.metrics.fontSize.xsmall};
  font-weight: bold;
  border-radius: 4px;
  margin-right: 5px;
`;

const Date = styled.div`
  color: ${({ theme }) => theme.colors.text.primary};
  font-size: ${({ theme }) => theme.metrics.fontSize.xsmall};
  margin-right: 15px;
`;

const View = styled.div`
  color: ${({ theme }) => theme.colors.text.primary};
  font-size: ${({ theme }) => theme.metrics.fontSize.large};
  margin-left: 10px;
  display: block;
  width: 0px;
  white-space: nowrap;
  opacity: 0;
  transition: all ease 1s;
  text-align: end;
  font-weight: bold;
`;

const LoadingWrapper = styled.li`
  width: 100%;
  padding: 10px 0;
  display: flex;
  align-items: center;
  justify-content: center;
  background-color: ${({ theme }) => theme.colors.background.highlight};
`;

const LoadingText = styled.span`
  color: ${({ theme }) => theme.colors.text.primary};
  font-size: ${({ theme }) => theme.metrics.fontSize.medium};
`;

const SearchBar = styled(SearchBarComp)``;

const HeaderVulList = styled.div`
  width: 100%;
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin-bottom: 25px;
`;

const FileTitle = styled.span`
  display: block;
  color: ${({ theme }) => theme.colors.text.primary};
  font-size: ${({ theme }) => theme.metrics.fontSize.medium};
  font-weight: bold;
`;

const Back = styled.div`
  display: flex;
  color: ${({ theme }) => theme.colors.text.primary};
  align-items: center;
  cursor: pointer;

  &:hover {
    transform: scale(1.2);
  }
`;

const BackText = styled.span`
  display: block;
  margin-left: 7px;
  font-size: ${({ theme }) => theme.metrics.fontSize.medium};
`;

const VulnerabilitiesList = styled.ul`
  list-style: none;
  margin-top: 15px;
`;

const Vulnerability = styled.li`
  width: 100%;
  border: 1px solid #787878;
  padding: 15px;
  margin-bottom: 15px;
  border-radius: 3px;

  a {
    color: ${({ theme }) => theme.colors.text.link};

    :hover {
      text-decoration: underline;
    }
  }
`;

const VulDetailWrapper = styled.div`
  display: flex;
  align-items: center;
  justify-content: space-between;
`;

const VulDetail = styled.p<VulDetailProps>`
  color: ${({ theme }) => theme.colors.text.opaque};
  word-wrap: normal;
  display: block;
  overflow: hidden;
  max-height: 20px;
  line-height: 20px;
  width: 95%;

  ${({ isOpen }) =>
    isOpen &&
    css`
      max-height: max-content;
      margin-bottom: 20px;
    `};
`;

const Ellipsis = styled.span`
  color: ${({ theme }) => theme.colors.text.primary};
  font-size: ${({ theme }) => theme.metrics.fontSize.xxlarge};
  background-color: ${({ theme }) => theme.colors.background.highlight};
  border-radius: 2px;
  width: 25px;
  text-align: center;
  line-height: 16px;
  height: 26px;
  transition: all ease 1s;
  cursor: pointer;

  &:hover {
    transform: scale(1.1);
    background-color: #444446;
  }
`;

const Info = styled.span`
  display: block;
  margin-top: 10px;
  margin-bottom: 5px;
  font-size: ${({ theme }) => theme.metrics.fontSize.xsmall};
`;

const Code = styled.code`
  display: block;
  padding: 8px;
  margin-top: 14px;
  background-color: ${({ theme }) => theme.colors.background.overlap};
`;

const CodeInfoWrapper = styled.span`
  display: flex;
  justify-content: space-between;
`;

const CodeInfo = styled.span`
  display: block;
  font-size: 12px;
  margin-top: 2px;
`;

const SelectOptionsWrapper = styled.span`
  margin-top: 10px;
  display: flex;
  align-items: center;
`;

const SearchTitle = styled.span`
  color: ${({ theme }) => theme.colors.text.opaque};
  margin-bottom: 20px;
`;

const UpdateContent = styled.div`
  background-color: ${({ theme }) => theme.colors.background.secundary};
  border-radius: 4px;
  padding: 10px 22px;
  margin-top: 10px;
  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: space-between;
`;

const UpdateCount = styled.span`
  color: ${({ theme }) => theme.colors.text.opaque};
`;

const UpdateBtns = styled.div`
  display: flex;
`;

const FileInfo = styled.div`
  margin-top: 5px;
  display: flex;
  align-items: center;
  justify-content: space-between;
`;

const FileInfoText = styled.span`
  margin-top: 5px;
  display: flex;
  align-items: center;
  color: ${({ theme }) => theme.colors.text.opaque};
  font-size: ${({ theme }) => theme.metrics.fontSize.xsmall};
`;

export default {
  UpdateContent,
  FileInfoText,
  FileInfo,
  FileRow,
  UpdateBtns,
  UpdateCount,
  SearchTitle,
  SelectOptionsWrapper,
  Info,
  CodeInfoWrapper,
  CodeInfo,
  HeaderVulList,
  Code,
  Ellipsis,
  VulDetailWrapper,
  VulDetail,
  VulnerabilitiesList,
  Vulnerability,
  BackText,
  Back,
  Wrapper,
  Options,
  SearchWrapper,
  SelectsWrapper,
  Content,
  File,
  FileLanguage,
  FileName,
  FileVulCount,
  FileColumn,
  Date,
  View,
  LoadingWrapper,
  LoadingText,
  ScrollList,
  SearchBar,
  FileTitle,
};
