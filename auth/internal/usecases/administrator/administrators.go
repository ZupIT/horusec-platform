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
	for _, adm := range a.all {
		if adm.Email == a.newest.Email {
			adm.Username = a.newest.Username
			adm.Password = a.newest.Password
			adm.CreatedAt = a.newest.CreatedAt
			return adm
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
