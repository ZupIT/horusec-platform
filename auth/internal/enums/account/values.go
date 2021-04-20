package account

const (
	DatabaseTableAccount           = "accounts"
	DuplicatedConstraintEmail      = "duplicate key value violates unique constraint \"accounts_email_key\""
	DuplicatedConstraintUsername   = "duplicate key value violates unique constraint \"uk_accounts_username\""
	DuplicatedConstraintPrimaryKey = "duplicate key value violates unique constraint \"accounts_pkey\""
)
