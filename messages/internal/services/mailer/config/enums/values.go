package enums

const (
	EnvSMTPUsername  = "HORUSEC_SMTP_USERNAME"
	EnvSMTPPassword  = "HORUSEC_SMTP_PASSWORD" //nolint:gosec //false positive
	EnvSMTPHost      = "HORUSEC_SMTP_HOST"
	EnvSMTPPort      = "HORUSEC_SMTP_PORT"
	EnvSMTPEmailFrom = "HORUSEC_EMAIL_FROM"
)
