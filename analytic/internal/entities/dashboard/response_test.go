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
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestSetTotalAuthors(t *testing.T) {
	t.Run("should success set total authors", func(t *testing.T) {
		response := &Response{}

		assert.NoError(t, response.SetTotalAuthors(18, nil))
		assert.Equal(t, 18, response.TotalAuthors)
	})

	t.Run("should return error when it is not nil", func(t *testing.T) {
		response := &Response{}

		assert.Error(t, response.SetTotalAuthors(0, errors.New("test")))
	})
}

func TestSetTotalRepositories(t *testing.T) {
	t.Run("should success set total repositories", func(t *testing.T) {
		response := &Response{}

		assert.NoError(t, response.SetTotalRepositories(18, nil))
		assert.Equal(t, 18, response.TotalRepositories)
	})

	t.Run("should return error when it is not nil", func(t *testing.T) {
		response := &Response{}

		assert.Error(t, response.SetTotalRepositories(0, errors.New("test")))
	})
}

func TestSetChartBySeverity(t *testing.T) {
	t.Run("should success set chart by severity", func(t *testing.T) {
		response := &Response{}

		assert.NoError(t, response.SetChartBySeverity(&Vulnerability{}, nil))
		assert.NotNil(t, 18, response.VulnerabilityBySeverity)
	})

	t.Run("should return error when it is not nil", func(t *testing.T) {
		response := &Response{}

		assert.Error(t, response.SetChartBySeverity(&Vulnerability{}, errors.New("test")))
	})
}

func TestSetChartByAuthor(t *testing.T) {
	authors := []*VulnerabilitiesByAuthor{
		{
			Author:        "test",
			Vulnerability: Vulnerability{},
		},
	}

	t.Run("should success set chart by author", func(t *testing.T) {
		response := &Response{}

		assert.NoError(t, response.SetChartByAuthor(authors, nil))
		assert.NotNil(t, 18, response.VulnerabilityBySeverity)
	})

	t.Run("should return error when it is not nil", func(t *testing.T) {
		response := &Response{}

		assert.Error(t, response.SetChartByAuthor(authors, errors.New("test")))
	})
}

func TestSetChartByRepository(t *testing.T) {
	repositories := []*VulnerabilitiesByRepository{
		{
			RepositoryName: "test",
			Vulnerability:  Vulnerability{},
		},
	}

	t.Run("should success set chart by repository", func(t *testing.T) {
		response := &Response{}

		assert.NoError(t, response.SetChartByRepository(repositories, nil))
		assert.NotNil(t, 18, response.VulnerabilityBySeverity)
	})

	t.Run("should return error when it is not nil", func(t *testing.T) {
		response := &Response{}

		assert.Error(t, response.SetChartByRepository(repositories, errors.New("test")))
	})
}

func TestSetChartByLanguage(t *testing.T) {
	languages := []*VulnerabilitiesByLanguage{
		{
			Language:      "test",
			Vulnerability: Vulnerability{},
		},
	}

	t.Run("should success set chart by language", func(t *testing.T) {
		response := &Response{}

		assert.NoError(t, response.SetChartByLanguage(languages, nil))
		assert.NotNil(t, 18, response.VulnerabilityBySeverity)
	})

	t.Run("should return error when it is not nil", func(t *testing.T) {
		response := &Response{}

		assert.Error(t, response.SetChartByLanguage(languages, errors.New("test")))
	})
}

func TestSetChartByTime(t *testing.T) {
	times := []*VulnerabilitiesByTime{
		{
			Vulnerability: Vulnerability{},
		},
	}

	t.Run("should success set chart by time", func(t *testing.T) {
		response := &Response{}

		assert.NoError(t, response.SetChartByTime(times, nil))
		assert.NotNil(t, 18, response.VulnerabilityBySeverity)
	})

	t.Run("should return error when it is not nil", func(t *testing.T) {
		response := &Response{}

		assert.Error(t, response.SetChartByTime(times, errors.New("test")))
	})
}
