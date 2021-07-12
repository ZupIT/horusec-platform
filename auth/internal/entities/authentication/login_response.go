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

package authentication

import (
	"time"

	"github.com/google/uuid"
)

type LoginResponse struct {
	AccountID          uuid.UUID `json:"accountID"`
	AccessToken        string    `json:"accessToken"`
	RefreshToken       string    `json:"refreshToken"`
	Username           string    `json:"username"`
	Email              string    `json:"email"`
	ExpiresAt          time.Time `json:"expiresAt"`
	ExpiresIn          int       `json:"expiresIn"`
	RefreshExpiresIn   int       `json:"refreshExpiresIn"`
	IsApplicationAdmin bool      `json:"isApplicationAdmin"`
}
