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
import { useTranslation } from 'react-i18next';
import { PaginationInfo } from 'helpers/interfaces/Pagination';
import { TablePagination } from '@material-ui/core';

interface Props {
  onChange: (pagination: PaginationInfo) => void;
  pagination: PaginationInfo;
}

const Pagination: React.FC<Props> = ({ onChange, pagination }) => {
  const { t } = useTranslation();

  const handleChangePage = (event: unknown, newPage: number) => {
    onChange({ ...pagination, currentPage: newPage + 1 });
  };

  const handleChangeRowsPerPage = (
    event: React.ChangeEvent<HTMLInputElement>
  ) => {
    onChange({
      ...pagination,
      currentPage: 1,
      pageSize: Number(event.target.value),
    });
  };

  return (
    <TablePagination
      rowsPerPageOptions={[10, 50, 100]}
      component="div"
      count={pagination.totalItems}
      rowsPerPage={pagination.pageSize}
      page={pagination.currentPage - 1}
      onChangePage={handleChangePage}
      onChangeRowsPerPage={handleChangeRowsPerPage}
      labelRowsPerPage={t('GENERAL.PAGINATION.ITENS_PAGE')}
      labelDisplayedRows={({ from, to, count }) =>
        `${from}-${to} ${t('GENERAL.PAGINATION.OF')} ${
          count !== -1 ? count : to
        }`
      }
    />
  );
};

export default Pagination;
