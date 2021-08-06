// Copyright 2021 ZUP IT SERVICOS EM TECNOLOGIA E INOVACAO SA
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package administrator

import (
	entity "github.com/ZupIT/horusec-platform/auth/internal/entities/account"
)

type administrators struct {
	newest *entity.Account
	all    []*entity.Account
}

func newAdministrators(newest *entity.Account, all []*entity.Account) *administrators {
	return &administrators{newest: newest, all: all}
}

func (a *administrators) Current() *entity.Account {
	for _, oldest := range a.all {
		if oldest.Email == a.newest.Email {
			return mergeAccounts(oldest, a.newest)
		}
	}
	return nil
}

func (a administrators) Oldest() []*entity.Account {
	var accounts []*entity.Account
	for _, adm := range a.all {
		if adm.Email != a.newest.Email {
			accounts = append(accounts, adm)
		}
	}
	return accounts
}

func mergeAccounts(oldest, newest *entity.Account) *entity.Account {
	return &entity.Account{
		AccountID:          oldest.AccountID,
		Email:              oldest.Email,
		Password:           newest.Password,
		Username:           newest.Username,
		IsConfirmed:        newest.IsConfirmed,
		IsApplicationAdmin: true,
		CreatedAt:          oldest.CreatedAt,
		UpdatedAt:          newest.UpdatedAt,
	}
}
