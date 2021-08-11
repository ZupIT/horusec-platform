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
	"fmt"

	"github.com/ZupIT/horusec-devkit/pkg/services/database"
	"github.com/ZupIT/horusec-devkit/pkg/utils/logger"

	entity "github.com/ZupIT/horusec-platform/auth/internal/entities/account"
	enum "github.com/ZupIT/horusec-platform/auth/internal/enums/account"
)

type where map[string]interface{}

type UseCase struct {
	read  database.IDatabaseRead
	write database.IDatabaseWrite
}

func NewUseCase(connection *database.Connection) *UseCase {
	return &UseCase{read: connection.Read, write: connection.Write}
}

func (u *UseCase) CreateOrUpdate(account *entity.Account) error {
	accounts, err := u.findAll()
	if err != nil {
		return err
	}

	admins := newAdministrators(account, accounts)
	if err := u.delete(admins.Oldest()); err != nil {
		return err
	}

	if current := admins.Current(); current != nil {
		return u.update(current)
	}

	return u.create(account)
}

func (u *UseCase) findAll() ([]*entity.Account, error) {
	accounts := make([]*entity.Account, 0)
	res := u.read.Find(&accounts, where{"is_application_admin": true}, enum.DatabaseTableAccount)
	if err := res.GetErrorExceptNotFound(); err != nil {
		return nil, err
	}
	return accounts, nil
}

func (u *UseCase) update(account *entity.Account) error {
	username := account.Username
	email := account.Email
	err := u.write.Update(account, where{"account_id": account.AccountID}, enum.DatabaseTableAccount).GetError()
	if err != nil {
		return fmt.Errorf("failed to update account user %q with email %q: %w", username, email, err)
	}
	logger.LogDebugWithLevel(fmt.Sprintf("user %q with email %q was updated successfully", username, email))
	return nil
}

func (u *UseCase) create(account *entity.Account) error {
	username := account.Username
	email := account.Email
	err := u.write.Create(account, enum.DatabaseTableAccount).GetError()
	if err != nil {
		return fmt.Errorf("failed to create account user %q with email %q: %w", username, email, err)
	}
	logger.LogDebugWithLevel(fmt.Sprintf("user %q with email %q was created successfully", username, email))
	return nil
}

func (u *UseCase) delete(accounts []*entity.Account) error {
	for _, account := range accounts {
		id := account.AccountID
		username := account.Username
		email := account.Email
		if err := u.write.Delete(where{"account_id": id}, enum.DatabaseTableAccount).GetError(); err != nil {
			return fmt.Errorf("failed to delete account user %q with email %q: %w", username, email, err)
		}
		logger.LogDebugWithLevel(fmt.Sprintf("user %q with email %q was deleted successfully", username, email))
	}
	return nil
}
