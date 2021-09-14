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

import styled from 'styled-components';

const Wrapper = styled.div`
  display: flex;
  flex-direction: column;
  align-items: center;
  padding-bottom: 50px;
  height: 96.3vh;

  @media (max-width: 768px) {
    margin-top: 40px;
  }
`;

const LogoContent = styled.div`
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  flex: 1;
  width: 252px;
`;

const Content = styled.div`
  margin-top: 80px;
`;

const Logo = styled.img`
  width: 266px;

  position: absolute;
  top: 25%;
  left: 50%;
  transform: translate(-50%, -50%);

  @media (max-width: 768px) {
    width: 220px;
    top: 20%;
  }
`;

const Footer = styled.footer`
  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: space-around;
  width: 100%;
  margin-right: 70px;

  @media (max-width: 768px) {
    flex-direction: column-reverse;
    margin-top: 30px;
    margin-right: 0;
  }
`;

const ContactWrapper = styled.div``;

const ContactItem = styled.a`
  display: flex;
  flex-direction: row;
  margin-top: 15px;
  align-items: center;
  justify-content: center;
  cursor: pointer;

  :hover {
    span,
    i {
      color: ${({ theme }) => theme.colors.active};
    }
  }
`;

const ContactText = styled.span`
  display: block;
  color: ${({ theme }) => theme.colors.text.primary};
  margin-left: 5px;
`;

export default {
  Wrapper,
  Logo,
  Content,
  Footer,
  LogoContent,
  ContactWrapper,
  ContactItem,
  ContactText,
};
