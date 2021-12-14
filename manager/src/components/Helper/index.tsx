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

import React from 'react';
import { useTranslation } from 'react-i18next';
import Styled from './styled';
import { Menu, MenuItem } from '@material-ui/core';
import packageJSON from '../../../package.json';

interface HelperInterface {
  pageHelperUrl: string;
}

const Helper: React.FC<HelperInterface> = ({ pageHelperUrl }) => {
  const { repository, bugs } = packageJSON;
  const { t } = useTranslation();
  const [anchorEl, setAnchorEl] = React.useState<null | HTMLElement>(null);

  const handleClick = (event: React.MouseEvent<HTMLButtonElement>) => {
    setAnchorEl(event.currentTarget);
  };

  const handleClose = () => {
    setAnchorEl(null);
  };

  const openExternalLink = (url: string) => {
    window.open(url, '_blank');
    handleClose();
  };

  return (
    <>
      <Styled.Button
        tabIndex={0}
        aria-label={t('HEADER.ARIA.HELP')}
        aria-controls="simple-menu"
        aria-haspopup="true"
        onClick={handleClick}
      >
        <Styled.HelpIcon name="help" size="20px" />

        <Styled.Text>{t('SIDE_MENU.HELPER')}</Styled.Text>
      </Styled.Button>

      <Menu
        id="simple-menu"
        anchorEl={anchorEl}
        open={Boolean(anchorEl)}
        onClose={handleClose}
      >
        <MenuItem onClick={() => openExternalLink(pageHelperUrl)}>
          <Styled.IconItemMenu size="18px" name="documentation" />
          {t('GENERAL.ABOUT_PAGE')}
        </MenuItem>

        <MenuItem onClick={() => openExternalLink(repository.url)}>
          <Styled.IconItemMenu size="18px" name="github" />
          {t('GENERAL.GITHUB')}
        </MenuItem>

        <MenuItem onClick={() => openExternalLink(bugs.url)}>
          <Styled.IconItemMenu size="18px" name="forum" />
          {t('GENERAL.FORUM')}
        </MenuItem>
      </Menu>
    </>
  );
};

export default Helper;
