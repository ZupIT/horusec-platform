package account

import "time"

const (
	DatabaseTableAccount           = "accounts"
	DuplicatedConstraintEmail      = "duplicate key value violates unique constraint \"accounts_email_key\""
	DuplicatedConstraintUsername   = "duplicate key value violates unique constraint \"uk_accounts_username\""
	DuplicatedConstraintPrimaryKey = "duplicate key value violates unique constraint \"accounts_pkey\""
	ID                             = "accountID"
	ResetPasswordCharset           = "1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	ResetPasswordCodeDuration      = time.Minute * 10
)
