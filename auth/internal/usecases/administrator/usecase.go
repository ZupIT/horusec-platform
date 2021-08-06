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
	accounts := make([]*entity.Account, 0, 0)
	err := u.read.Find(&accounts, where{"is_application_admin": true}, enum.DatabaseTableAccount).GetError()
	if err != nil {
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
