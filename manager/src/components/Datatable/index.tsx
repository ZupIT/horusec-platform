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
import Styled from './styled';
import { Button, Icon, Pagination } from 'components';
import { PaginationInfo } from 'helpers/interfaces/Pagination';
import ReactTooltip, { TooltipProps } from 'react-tooltip';
import {
  IconButton,
  Menu,
  MenuItem,
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableHead,
  TableRow,
  Toolbar,
} from '@material-ui/core';
import { MoreHoriz } from '@material-ui/icons';
import { divide, kebabCase } from 'lodash';
import PopupState, { bindTrigger, bindMenu } from 'material-ui-popup-state';

export interface TableColumn {
  label: string;
  property: string;
  type: 'text' | 'custom' | 'actions';
  cssClass?: string[];
}

export interface DataSource {
  [x: string]: any;
  id?: string | number;
  buttons?: {
    [x: string]: { name: string; size: string; function: () => void };
  };
  actions?: {
    icon: string;
    title: string;
    function: (element?: any) => void;
  }[];
}

interface DatatableInterface {
  title?: string;
  columns: TableColumn[];
  dataSource: DataSource[];
  total?: number;
  paginate?: {
    pagination: PaginationInfo;
    onChange: (pagination: PaginationInfo) => void;
  };
  emptyListText?: string;
  isLoading?: boolean;
  tooltip?: TooltipProps;
  fixed?: boolean;
  buttons?: {
    title: string;
    icon?: string;
    disabled?: boolean;
    show?: boolean;
    function: (...args: any) => void;
  }[];
}

const Datatable: React.FC<DatatableInterface> = (props) => {
  const {
    columns,
    dataSource,
    emptyListText,
    isLoading,
    paginate,
    tooltip,
    fixed = true,
    buttons = [],
    title,
  } = props;

  return (
    <Styled.Content>
      <TableContainer>
        <Toolbar
          disableGutters
          style={{ minHeight: 0, justifyContent: 'space-between' }}
        >
          {title ? <Styled.Title>{title}</Styled.Title> : <div></div>}
          <Styled.ButtonWrapper>
            {buttons.map((button, key) =>
              button.show ? (
                <Button
                  key={key}
                  text={button.title}
                  icon={button.icon}
                  onClick={button.function}
                  width="auto"
                  hidden={button.show}
                />
              ) : null
            )}
          </Styled.ButtonWrapper>
        </Toolbar>

        <Table>
          <TableHead>
            <TableRow>
              {columns.map((el, index) => (
                <TableCell key={index}>{el.label}</TableCell>
              ))}
            </TableRow>
          </TableHead>
          <TableBody>
            {!dataSource || dataSource.length <= 0 ? (
              <TableRow>
                <TableCell colSpan={columns.length} align="center">
                  {emptyListText}
                </TableCell>
              </TableRow>
            ) : (
              dataSource.map((row, dataId) => (
                <TableRow key={`${row.id || 'item'}-${dataId}`}>
                  {columns.map((column, columnId) => {
                    const renderTooltipProps = (tip: string) => {
                      return tooltip
                        ? {
                            'data-for': tooltip.id,
                            'data-tip': tip,
                          }
                        : null;
                    };

                    if (column.type === 'text') {
                      return (
                        <TableCell
                          tabIndex={0}
                          key={columnId}
                          className={column.cssClass?.join(' ')}
                          {...renderTooltipProps(row[column.property])}
                        >
                          {row[column.property] || '-'}
                        </TableCell>
                      );
                    }

                    if (column.type === 'custom') {
                      return (
                        <TableCell
                          tabIndex={0}
                          key={columnId}
                          className={column.cssClass?.join(' ')}
                          style={{ overflow: 'visible' }}
                        >
                          {row[column.property]}
                        </TableCell>
                      );
                    }

                    if (column.type === 'actions') {
                      return (
                        <TableCell
                          tabIndex={0}
                          key={columnId}
                          className={column.cssClass?.join(' ')}
                        >
                          {row[column.type].length >= 1 ? (
                            <div className="row">
                              <PopupState
                                variant="popover"
                                popupId={`popup-menu-${dataId}`}
                              >
                                {(popupState) => (
                                  <React.Fragment>
                                    <IconButton {...bindTrigger(popupState)}>
                                      <MoreHoriz />
                                    </IconButton>
                                    <Menu {...bindMenu(popupState)}>
                                      {row[column.type].map(
                                        (
                                          action: DataSource,
                                          actionId: React.Key
                                        ) => (
                                          <MenuItem
                                            key={actionId}
                                            onClick={() => {
                                              action.function();
                                              popupState.close();
                                            }}
                                          >
                                            <Button
                                              id={`action-${kebabCase(
                                                action.title
                                              )}-${columnId}-${dataId}`}
                                              rounded
                                              outline
                                              opaque
                                              text={action.title}
                                              width={'100%'}
                                              height={30}
                                              icon={action.icon}
                                            />
                                          </MenuItem>
                                        )
                                      )}
                                    </Menu>
                                  </React.Fragment>
                                )}
                              </PopupState>
                            </div>
                          ) : (
                            '-'
                          )}
                        </TableCell>
                      );
                    }

                    return null;
                  })}
                </TableRow>
              ))
            )}
          </TableBody>
        </Table>
      </TableContainer>
      {dataSource && dataSource.length > 0 && paginate ? (
        <Pagination
          pagination={paginate.pagination}
          onChange={paginate.onChange}
        />
      ) : null}
      {tooltip ? <ReactTooltip {...tooltip} /> : null}
    </Styled.Content>
  );
};

export default Datatable;
