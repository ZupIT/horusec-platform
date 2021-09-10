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

import { createTheme } from '@material-ui/core';
import { getCurrentTheme } from 'helpers/localStorage/currentTheme';

const theme = getCurrentTheme();

const themeMatUi = createTheme({
  palette: {
    primary: {
      // primary color
      main: theme.colors.active, // black
    },
    secondary: {
      main: theme.colors.secondary,
    },

    error: {
      main: theme.colors.input.error,
    },
  },

  overrides: {
    MuiSelect: {
      select: {
        paddingLeft: '5px',
      },
      icon: {
        color: theme.colors.button.text,
        '&$disabled': {
          color: 'transparent',
        },
      },
    },
    MuiFilledInput: {
      root: {
        backgroundColor: 'transparent',
        borderRadius: '10px !important',
        '&$disabled': {
          backgroundColor: 'transparent',
        },
      },
      input: {
        padding: '6px 0px',
      },
      underline: {
        '&:before': {
          borderBottom: 'none !important',
        },
        '&:after': {
          borderBottom: 'none !important',
        },
        '&:hover:not($disabled):before': {
          borderBottom: 'none !important',
        },
      },
    },
    MuiInputBase: {
      root: {
        color: theme.colors.input.text,
      },
      input: {
        '&::-webkit-calendar-picker-indicator': {
          filter: 'invert(1)',
        },
        '&:-webkit-autofill': {
          '-webkit-text-fill-color': 'white',
          '-webkit-box-shadow': '0 0 0 30px rgb(28 28 30) inset !important',
        },
        '&:disabled': {
          color: `${theme.colors.input.disabled} !important`,
        },
      },
    },
    MuiInput: {
      underline: {
        '&:before': {
          borderColor: `${theme.colors.input.border} !important`,
        },
        '&:after': {
          borderColor: `${theme.colors.input.border} !important`,
        },
        '&:hover:not($disabled):before': {
          borderColor: `${theme.colors.input.border} !important`,
        },
      },
    },
    MuiFormLabel: {
      root: {
        color: theme.colors.input.label,
        '&$focused': {
          color: theme.colors.input.label,
        },
        '&$disabled': {
          color: `${theme.colors.input.disabled} !important`,
        },
      },
      focused: {},
    },
    MuiInputLabel: {
      root: {
        color: theme.colors.input.label,
      },
    },
    MuiIconButton: {
      root: {
        color: theme.colors.button.text,
      },
    },

    MuiCheckbox: {
      root: {
        color: theme.colors.checkbox.border,
      },
      colorSecondary: {
        '&$checked': {
          color: theme.colors.checkbox.checked.secundary,
        },
      },
      checked: {},
    },
    MuiPaper: {
      root: {
        color: 'white',
        backgroundColor: theme.colors.background.highlight,

        '& .MuiPickersBasePicker-container': {
          backgroundColor: theme.colors.background.highlight,
        },
        '& .MuiPickersDay-day': {
          color: 'white',
        },
        '& .MuiPickersCalendarHeader-dayLabel': {
          color: 'white',
        },
        '& .MuiPickersCalendarHeader-iconButton': {
          background: 'none',
          color: '#fff',
        },
      },
    },
    MuiTableBody: {
      root: {
        backgroundColor: theme.colors.dataTable.row.background,
        '& > .MuiTableRow-root:first-child': {
          borderTopLeftRadius: '5px',
          borderBottomLeftRadius: '5px',
        },
        '& > .MuiTableRow-root:last-child': {
          borderTopRightRadius: '5px',
          borderBottomRightRadius: '5px',
        },
        gap: '10px',
      },
    },
    MuiTableCell: {
      root: {
        color: theme.colors.dataTable.column.text + '!important',
        borderBottomColor: theme.colors.dataTable.background,
      },
    },
    MuiTablePagination: {
      caption: {
        color: theme.colors.dataTable.row.text,
      },
      select: {
        color: theme.colors.dataTable.row.text,
      },
      actions: {
        color: theme.colors.dataTable.row.text,
      },
    },
    // MuiIconButton: {
    //   root: {
    //     color: theme.colors.icon.primary
    //   }
    // }
  },
});

export default themeMatUi;
