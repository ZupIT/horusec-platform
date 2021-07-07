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
	"encoding/json"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/google/uuid"

	utils "github.com/ZupIT/horusec-devkit/pkg/utils/validation"
)

type ChangePasswordData struct {
	Password  string    `json:"password"`
	AccountID uuid.UUID `json:"accountID" swaggerignore:"true"`
}

func (c *ChangePasswordData) Validate() error {
	return validation.ValidateStruct(c,
		validation.Field(&c.Password, utils.PasswordValidationRules()...),
	)
}

func (c *ChangePasswordData) SetAccountID(accountID uuid.UUID) *ChangePasswordData {
	c.AccountID = accountID

	return c
}

func (c *ChangePasswordData) ToBytes() []byte {
	bytes, _ := json.Marshal(c)

	return bytes
}
