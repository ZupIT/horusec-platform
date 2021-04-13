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
