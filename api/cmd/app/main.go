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

package main

import (
	"github.com/ZupIT/horusec-devkit/pkg/services/tracer"
	"github.com/ZupIT/horusec-devkit/pkg/utils/logger"

	"github.com/ZupIT/horusec-platform/api/config/providers"
	"github.com/ZupIT/horusec-platform/api/internal/enums"
)

// @title Horusec-API
// @description Service responsible for analysis operations from Horusec-CLI.
// @termsOfService http://swagger.io/terms/

// @contact.name Horusec
// @contact.url https://github.com/ZupIT/horusec-platform
// @contact.email horusec@zup.com.br

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name X-Horusec-Authorization
func main() {
	j, err := tracer.NewJaeger()
	if err != nil {
		logger.LogDebugWithLevel(err.Error() + ` JAEGER_SERVICE_NAME set to "api" as default`)
		j = &tracer.Jaeger{Name: "api"}
	}
	router, err := providers.Initialize(enums.DefaultPort, *j)
	if err != nil {
		panic(err)
	}
	router.ListenAndServe()
}
