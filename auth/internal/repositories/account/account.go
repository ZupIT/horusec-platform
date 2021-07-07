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

package account

import (
	"github.com/google/uuid"

	"github.com/ZupIT/horusec-devkit/pkg/services/database"

	accountEntities "github.com/ZupIT/horusec-platform/auth/internal/entities/account"
	accountEnums "github.com/ZupIT/horusec-platform/auth/internal/enums/account"
	accountUseCases "github.com/ZupIT/horusec-platform/auth/internal/usecases/account"
)

type IRepository interface {
	GetAccount(accountID uuid.UUID) (*accountEntities.Account, error)
	GetAccountByEmail(email string) (*accountEntities.Account, error)
	GetAccountByUsername(username string) (*accountEntities.Account, error)
	CreateAccount(account *accountEntities.Account) (*accountEntities.Account, error)
	Update(account *accountEntities.Account) (*accountEntities.Account, error)
	Delete(accountID uuid.UUID) error
}

type Repository struct {
	databaseRead  database.IDatabaseRead
	databaseWrite database.IDatabaseWrite
	useCases      accountUseCases.IUseCases
}

func NewAccountRepository(connection *database.Connection, useCases accountUseCases.IUseCases) IRepository {
	return &Repository{
		databaseRead:  connection.Read,
		databaseWrite: connection.Write,
		useCases:      useCases,
	}
}

func (r *Repository) GetAccount(accountID uuid.UUID) (*accountEntities.Account, error) {
	account := &accountEntities.Account{}

	return account, r.databaseRead.Find(account, r.useCases.FilterAccountByID(accountID),
		accountEnums.DatabaseTableAccount).GetError()
}

func (r *Repository) GetAccountByEmail(email string) (*accountEntities.Account, error) {
	account := &accountEntities.Account{}

	return account, r.databaseRead.Find(account, r.useCases.FilterAccountByEmail(email),
		accountEnums.DatabaseTableAccount).GetError()
}

func (r *Repository) GetAccountByUsername(username string) (*accountEntities.Account, error) {
	account := &accountEntities.Account{}

	return account, r.databaseRead.Find(account, r.useCases.FilterAccountByUsername(username),
		accountEnums.DatabaseTableAccount).GetError()
}

func (r *Repository) CreateAccount(account *accountEntities.Account) (*accountEntities.Account, error) {
	return account, r.databaseWrite.Create(account, accountEnums.DatabaseTableAccount).GetError()
}

func (r *Repository) Update(account *accountEntities.Account) (*accountEntities.Account, error) {
	return account, r.databaseWrite.Update(account, r.useCases.FilterAccountByID(account.AccountID),
		accountEnums.DatabaseTableAccount).GetError()
}

func (r *Repository) Delete(accountID uuid.UUID) error {
	return r.databaseWrite.Delete(r.useCases.FilterAccountByID(accountID), accountEnums.DatabaseTableAccount).GetError()
}
