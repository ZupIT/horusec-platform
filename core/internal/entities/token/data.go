package token

import (
	"encoding/json"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/google/uuid"

	"github.com/ZupIT/horusec-devkit/pkg/utils/crypto"

	tokenEnums "github.com/ZupIT/horusec-platform/core/internal/enums/token"
)

type Data struct {
	Description  string    `json:"description"`
	RepositoryID uuid.UUID `json:"repositoryID" swaggerignore:"true"`
	WorkspaceID  uuid.UUID `json:"workspaceID" swaggerignore:"true"`
	TokenID      uuid.UUID `json:"tokenID" swaggerignore:"true"`
	IsExpirable  bool      `json:"isExpirable"`
	ExpiresAt    time.Time `json:"expiresAt"`
}

func (d *Data) Validate() error {
	return validation.ValidateStruct(d,
		validation.Field(&d.WorkspaceID, is.UUID),
		validation.Field(&d.RepositoryID, is.UUID),
		validation.Field(&d.TokenID, is.UUID),
		validation.Field(&d.Description, validation.Required, validation.Length(1, 255)),
		validation.Field(&d.ExpiresAt, validation.By(d.validateExpiresAt)),
	)
}

func (d *Data) validateExpiresAt(_ interface{}) error {
	if d.IsExpirable && d.ExpiresAt.Before(time.Now()) {
		return tokenEnums.ErrorInvalidTokenExpiresAt
	}

	return nil
}

func (d *Data) SetWorkspaceID(workspaceID uuid.UUID) *Data {
	d.WorkspaceID = workspaceID

	return d
}

func (d *Data) SetIDs(workspaceID, repositoryID, tokenID uuid.UUID) *Data {
	d.WorkspaceID = workspaceID
	d.RepositoryID = repositoryID
	d.TokenID = tokenID

	return d
}

func (d *Data) ToToken() (*Token, string) {
	token := uuid.NewString()

	return &Token{
		TokenID:      uuid.New(),
		WorkspaceID:  d.WorkspaceID,
		RepositoryID: d.getRepositoryID(),
		Description:  d.Description,
		SuffixValue:  d.getSuffixValue(token),
		Value:        d.getTokenHash(token),
		IsExpirable:  d.IsExpirable,
		CreatedAt:    time.Now(),
		ExpiresAt:    d.ExpiresAt,
	}, token
}

func (d *Data) getSuffixValue(token string) string {
	return token[31:]
}

func (d *Data) getTokenHash(token string) string {
	return crypto.GenerateSHA256(token)
}

func (d *Data) getRepositoryID() *uuid.UUID {
	if d.RepositoryID == uuid.Nil {
		return nil
	}

	return &d.RepositoryID
}

func (d *Data) ToByes() []byte {
	bytes, _ := json.Marshal(d)

	return bytes
}
