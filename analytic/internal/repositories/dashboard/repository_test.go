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

package dashboard

import (
	"errors"
	"testing"

	"github.com/ZupIT/horusec-devkit/pkg/services/database"
	"github.com/ZupIT/horusec-devkit/pkg/services/database/response"
	"github.com/stretchr/testify/assert"

	"github.com/ZupIT/horusec-platform/analytic/internal/entities/dashboard"
)

func TestGetDashboardTotalDevelopers(t *testing.T) {
	t.Run("should return total developers without error", func(t *testing.T) {
		databaseReadMock := &database.Mock{}
		databaseReadMock.On("Raw").Return(response.NewResponse(0, nil, 1))

		connection := &database.Connection{
			Read:  databaseReadMock,
			Write: &database.Mock{},
		}

		repository := NewRepoDashboard(connection)

		total, err := repository.GetDashboardTotalDevelopers(&dashboard.Filter{})
		assert.NoError(t, err)
		assert.NotNil(t, total)
	})
}

func TestGetDashboardVulnBySeverity(t *testing.T) {
	t.Run("should return get vulns by severity without errors", func(t *testing.T) {
		databaseReadMock := &database.Mock{}
		databaseReadMock.On("Raw").Return(
			response.NewResponse(0, nil, &dashboard.Vulnerability{}))

		connection := &database.Connection{
			Read:  databaseReadMock,
			Write: &database.Mock{},
		}

		repository := NewRepoDashboard(connection)

		result, err := repository.GetDashboardVulnBySeverity(&dashboard.Filter{})
		assert.NoError(t, err)
		assert.NotNil(t, result)
	})
}

func TestGetDashboardVulnByAuthor(t *testing.T) {
	t.Run("should return get vulns by author without errors", func(t *testing.T) {
		databaseReadMock := &database.Mock{}
		databaseReadMock.On("Raw").Return(
			response.NewResponse(0, nil, &[]*dashboard.VulnerabilitiesByAuthor{{}}))

		connection := &database.Connection{
			Read:  databaseReadMock,
			Write: &database.Mock{},
		}

		repository := NewRepoDashboard(connection)

		_, err := repository.GetDashboardVulnByAuthor(&dashboard.Filter{})
		assert.NoError(t, err)
	})

	t.Run("should return empty response when no data was found", func(t *testing.T) {
		databaseReadMock := &database.Mock{}
		databaseReadMock.On("Raw").Return(
			response.NewResponse(0, nil, &[]*dashboard.VulnerabilitiesByAuthor{}))

		connection := &database.Connection{
			Read:  databaseReadMock,
			Write: &database.Mock{},
		}

		repository := NewRepoDashboard(connection)

		resp, err := repository.GetDashboardVulnByAuthor(&dashboard.Filter{})
		assert.NoError(t, err)
		assert.Len(t, resp, 1)
	})

	t.Run("should return error when something went wrong", func(t *testing.T) {
		databaseReadMock := &database.Mock{}
		databaseReadMock.On("Raw").Return(
			response.NewResponse(0, errors.New("test"), &[]*dashboard.VulnerabilitiesByAuthor{{}}))

		connection := &database.Connection{
			Read:  databaseReadMock,
			Write: &database.Mock{},
		}

		repository := NewRepoDashboard(connection)

		_, err := repository.GetDashboardVulnByAuthor(&dashboard.Filter{})
		assert.Error(t, err)
	})
}

func TestGetDashboardVulnByLanguage(t *testing.T) {
	t.Run("should return response without errors", func(t *testing.T) {
		databaseReadMock := &database.Mock{}
		databaseReadMock.On("Raw").Return(
			response.NewResponse(0, nil, []*dashboard.VulnerabilitiesByLanguage{{}}))

		connection := &database.Connection{
			Read:  databaseReadMock,
			Write: &database.Mock{},
		}

		repository := NewRepoDashboard(connection)

		_, err := repository.GetDashboardVulnByLanguage(&dashboard.Filter{})
		assert.NoError(t, err)
	})

	t.Run("should return empty response when no data was found", func(t *testing.T) {
		databaseReadMock := &database.Mock{}
		databaseReadMock.On("Raw").Return(
			response.NewResponse(0, nil, []*dashboard.VulnerabilitiesByLanguage{}))

		connection := &database.Connection{
			Read:  databaseReadMock,
			Write: &database.Mock{},
		}

		repository := NewRepoDashboard(connection)

		resp, err := repository.GetDashboardVulnByLanguage(&dashboard.Filter{})
		assert.NoError(t, err)
		assert.Len(t, resp, 1)
	})

	t.Run("should return error when something went wrong", func(t *testing.T) {
		databaseReadMock := &database.Mock{}
		databaseReadMock.On("Raw").Return(
			response.NewResponse(0, errors.New("test"), []*dashboard.VulnerabilitiesByLanguage{{}}))

		connection := &database.Connection{
			Read:  databaseReadMock,
			Write: &database.Mock{},
		}

		repository := NewRepoDashboard(connection)

		_, err := repository.GetDashboardVulnByLanguage(&dashboard.Filter{})
		assert.Error(t, err)
	})
}

func TestGetDashboardVulnByTime(t *testing.T) {
	t.Run("should return get vulns by time without errors", func(t *testing.T) {
		databaseReadMock := &database.Mock{}
		databaseReadMock.On("Raw").Return(
			response.NewResponse(0, nil, []*dashboard.VulnerabilitiesByTime{}))

		connection := &database.Connection{
			Read:  databaseReadMock,
			Write: &database.Mock{},
		}

		repository := NewRepoDashboard(connection)

		_, err := repository.GetDashboardVulnByTime(&dashboard.Filter{})
		assert.NoError(t, err)
	})
}
