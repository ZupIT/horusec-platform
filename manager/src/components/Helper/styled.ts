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
import { Icon } from 'components';

const HelpIcon = styled(Icon)`
  color: ${({ theme }) => theme.colors.text.opaque};

  :hover {
    transform: scale(1.1);
  }
`;

const Button = styled.button`
  background: none;
  border: none;
  cursor: pointer;
  display: flex;
  align-items: center;

  i {
    margin: 0;
  }

  :hover {
    span,
    svg,
    * {
      color: ${({ theme }) => theme.colors.active};
    }
  }
`;

const Text = styled.span`
  color: ${({ theme }) => theme.colors.text.opaque};
  transition: all ease 0.2s;
  font-size: ${({ theme }) => theme.metrics.fontSize.small};
  margin-left: 5px;
`;

const IconItemMenu = styled(Icon)`
  margin-right: 5px;
`;

export default { HelpIcon, Button, Text, IconItemMenu };
