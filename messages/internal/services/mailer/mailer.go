// Copyright 2020 ZUP IT SERVICOS EM TECNOLOGIA E INOVACAO SA
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

package mailer

import (
	"crypto/tls"

	"gopkg.in/gomail.v2"

	"github.com/ZupIT/horusec-devkit/pkg/utils/logger"

	mailerConfig "github.com/ZupIT/horusec-platform/messages/internal/services/mailer/config"
	"github.com/ZupIT/horusec-platform/messages/internal/services/mailer/enums"
)

type IService interface {
	SendEmail(msg *gomail.Message) error
	Noop() error
	IsAvailable() bool
	GetFromHeader() string
}

type Service struct {
	config mailerConfig.IConfig
	dialer *gomail.Dialer
}

func NewMailerService() IService {
	config := mailerConfig.NewMailerConfig()
	if err := config.Validate(); err != nil {
		logger.LogPanic(enums.MessageInvalidMailerConfig, err)
	}

	mailer := &Service{config: config}
	return mailer.setupDialer()
}

func (s *Service) setupDialer() IService {
	s.dialer = gomail.NewDialer(s.config.GetHost(), s.config.GetPort(), s.config.GetUsername(), s.config.GetPassword())
	s.dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true} //nolint:gosec // is necessary to send without use tls check

	return s
}

func (s *Service) SendEmail(msg *gomail.Message) error {
	return s.dialer.DialAndSend(msg)
}

func (s *Service) Noop() error {
	_, err := s.dialer.Dial()

	return err
}

func (s *Service) IsAvailable() bool {
	return s.Noop() == nil
}

func (s *Service) GetFromHeader() string {
	return s.config.GetFrom()
}
