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

package analysis

import (
	"errors"
	"testing"
	"time"

	"github.com/ZupIT/horusec-devkit/pkg/entities/vulnerability"
	"github.com/ZupIT/horusec-devkit/pkg/enums/confidence"
	"github.com/ZupIT/horusec-devkit/pkg/enums/languages"
	"github.com/ZupIT/horusec-devkit/pkg/enums/severities"
	vulnerabilityEnum "github.com/ZupIT/horusec-devkit/pkg/enums/vulnerability"
	"github.com/ZupIT/horusec-devkit/pkg/services/database/enums"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/ZupIT/horusec-devkit/pkg/entities/analysis"
	analysisEnums "github.com/ZupIT/horusec-devkit/pkg/enums/analysis"
	"github.com/ZupIT/horusec-devkit/pkg/services/database"
	"github.com/ZupIT/horusec-devkit/pkg/services/database/response"
)

func TestAnalysis_FindAnalysisByID(t *testing.T) {
	t.Run("Should find analysis by id with success", func(t *testing.T) {
		data := &analysis.Analysis{
			ID:             uuid.New(),
			RepositoryID:   uuid.New(),
			RepositoryName: uuid.New().String(),
			WorkspaceID:    uuid.New(),
			WorkspaceName:  uuid.New().String(),
			Status:         analysisEnums.Success,
			Errors:         "",
			CreatedAt:      time.Now(),
			FinishedAt:     time.Now(),
		}
		mockRead := &database.Mock{}
		mockRead.On("FindPreload").Return(response.NewResponse(0, nil, data))
		connectionMock := &database.Connection{
			Read: mockRead,
		}
		res := NewRepositoriesAnalysis(connectionMock).FindAnalysisByID(uuid.New())
		assert.NoError(t, res.GetError())
		assert.NotEmpty(t, res.GetData())
		assert.NotEqual(t, res.GetData().(*analysis.Analysis).ID, uuid.Nil)
	})
	t.Run("Should find analysis by id with success", func(t *testing.T) {
		data := &analysis.Analysis{
			ID:             uuid.New(),
			RepositoryID:   uuid.New(),
			RepositoryName: uuid.New().String(),
			WorkspaceID:    uuid.New(),
			WorkspaceName:  uuid.New().String(),
			Status:         analysisEnums.Success,
			Errors:         "",
			CreatedAt:      time.Now(),
			FinishedAt:     time.Now(),
		}
		mockRead := &database.Mock{}
		mockRead.On("FindPreload").Return(response.NewResponse(0, nil, data))
		connectionMock := &database.Connection{
			Read: mockRead,
		}
		res := NewRepositoriesAnalysis(connectionMock).FindAnalysisByID(uuid.New())
		assert.NoError(t, res.GetError())
		assert.NotEmpty(t, res.GetData())
		assert.NotEqual(t, res.GetData().(*analysis.Analysis).ID, uuid.Nil)
	})
}

func TestAnalysis_CreateFullAnalysis(t *testing.T) {
	t.Run("Should create analysis with success", func(t *testing.T) {
		mockWrite := &database.Mock{}
		mockWrite.On("StartTransaction").Return(mockWrite)
		mockWrite.On("CommitTransaction").Return(response.NewResponse(0, nil, nil))
		mockWrite.On("Create").Return(response.NewResponse(0, nil, nil))
		connectionMock := &database.Connection{
			Write: mockWrite,
		}
		data := &analysis.Analysis{
			ID:             uuid.New(),
			RepositoryID:   uuid.New(),
			RepositoryName: uuid.New().String(),
			WorkspaceID:    uuid.New(),
			WorkspaceName:  uuid.New().String(),
			Status:         analysisEnums.Success,
			Errors:         "",
			CreatedAt:      time.Now(),
			FinishedAt:     time.Now(),
		}
		err := NewRepositoriesAnalysis(connectionMock).CreateFullAnalysis(data)
		assert.NoError(t, err)
	})
	t.Run("Should create analysis and vulnerabilities with success", func(t *testing.T) {
		mockWrite := &database.Mock{}
		mockRead := &database.Mock{}
		mockRead.On("Raw").Return(response.NewResponse(1, enums.ErrorNotFoundRecords, nil))
		mockWrite.On("StartTransaction").Return(mockWrite)
		mockWrite.On("CommitTransaction").Return(response.NewResponse(0, nil, nil))
		mockWrite.On("Create").Return(response.NewResponse(0, nil, nil))
		connectionMock := &database.Connection{
			Write: mockWrite,
			Read:  mockRead,
		}
		data := &analysis.Analysis{
			ID:             uuid.New(),
			WorkspaceID:    uuid.New(),
			WorkspaceName:  uuid.NewString(),
			RepositoryID:   uuid.New(),
			RepositoryName: uuid.NewString(),
			Status:         analysisEnums.Success,
			Errors:         "",
			CreatedAt:      time.Now(),
			FinishedAt:     time.Now(),
			AnalysisVulnerabilities: []analysis.AnalysisVulnerabilities{
				{
					VulnerabilityID: uuid.New(),
					AnalysisID:      uuid.New(),
					CreatedAt:       time.Now(),
					Vulnerability: vulnerability.Vulnerability{
						VulnerabilityID: uuid.New(),
						Line:            "1",
						Column:          "1",
						Confidence:      confidence.High,
						File:            "/deployments/cert.pem",
						Code:            "-----BEGIN CERTIFICATE-----",
						Details:         "Asymmetric Private Key \n Found SSH and/or x.509 Cerficates among the files of your project, make sure you want this kind of information inside your Git repo, since it can be missused by someone with access to any kind of copy.  For more information checkout the CWE-312 (https://cwe.mitre.org/data/definitions/312.html) advisory.",
						SecurityTool:    "Wrong security tool",
						Language:        languages.Leaks,
						Severity:        severities.Critical,
						VulnHash:        "1234567890",
						Type:            vulnerabilityEnum.Vulnerability,
						CommitAuthor:    "Wilian Gabriel",
						CommitEmail:     "wilian.silva@zup.com.br",
						CommitHash:      "9876543210",
						CommitMessage:   "Initial Commit",
						CommitDate:      "2021-03-31T10:58:42Z",
					},
				},
			},
		}
		err := NewRepositoriesAnalysis(connectionMock).CreateFullAnalysis(data)
		assert.NoError(t, err)
	})
	t.Run("Should create analysis and many to many with success but not create vulnerability already exists", func(t *testing.T) {
		mockWrite := &database.Mock{}
		mockRead := &database.Mock{}
		mockRead.On("Raw").Return(response.NewResponse(1, nil, map[string]interface{}{"vulnerability_id": uuid.New().String()}))
		mockWrite.On("StartTransaction").Return(mockWrite)
		mockWrite.On("CommitTransaction").Return(response.NewResponse(0, nil, nil))
		mockWrite.On("Create").Return(response.NewResponse(0, nil, nil))
		mockWrite.On("Update").Return(response.NewResponse(0, nil, nil))
		connectionMock := &database.Connection{
			Write: mockWrite,
			Read:  mockRead,
		}
		data := &analysis.Analysis{
			ID:             uuid.New(),
			WorkspaceID:    uuid.New(),
			WorkspaceName:  uuid.NewString(),
			RepositoryID:   uuid.New(),
			RepositoryName: uuid.NewString(),
			Status:         analysisEnums.Success,
			Errors:         "",
			CreatedAt:      time.Now(),
			FinishedAt:     time.Now(),
			AnalysisVulnerabilities: []analysis.AnalysisVulnerabilities{
				{
					VulnerabilityID: uuid.New(),
					AnalysisID:      uuid.New(),
					CreatedAt:       time.Now(),
					Vulnerability: vulnerability.Vulnerability{
						VulnerabilityID: uuid.New(),
						Line:            "1",
						Column:          "1",
						Confidence:      confidence.High,
						File:            "/deployments/cert.pem",
						Code:            "-----BEGIN CERTIFICATE-----",
						Details:         "Asymmetric Private Key \n Found SSH and/or x.509 Cerficates among the files of your project, make sure you want this kind of information inside your Git repo, since it can be missused by someone with access to any kind of copy.  For more information checkout the CWE-312 (https://cwe.mitre.org/data/definitions/312.html) advisory.",
						SecurityTool:    "Wrong security tool",
						Language:        languages.Leaks,
						Severity:        severities.Critical,
						VulnHash:        "1234567890",
						Type:            vulnerabilityEnum.Vulnerability,
						CommitAuthor:    "Wilian Gabriel",
						CommitEmail:     "wilian.silva@zup.com.br",
						CommitHash:      "9876543210",
						CommitMessage:   "Initial Commit",
						CommitDate:      "2021-03-31T10:58:42Z",
					},
				},
			},
		}
		err := NewRepositoriesAnalysis(connectionMock).CreateFullAnalysis(data)
		assert.NoError(t, err)
	})
	t.Run("Should create analysis and many to many with success but not update vulnerability if already exists but return invalid vuln_id", func(t *testing.T) {
		mockWrite := &database.Mock{}
		mockRead := &database.Mock{}
		mockRead.On("Raw").Return(response.NewResponse(1, nil, map[string]interface{}{"vulnerability_id": "invalid uuid"}))
		mockWrite.On("StartTransaction").Return(mockWrite)
		mockWrite.On("RollbackTransaction").Return(response.NewResponse(0, nil, nil))
		mockWrite.On("Create").Return(response.NewResponse(0, nil, nil))
		mockWrite.On("Update").Return(response.NewResponse(0, nil, nil))
		connectionMock := &database.Connection{
			Write: mockWrite,
			Read:  mockRead,
		}
		data := &analysis.Analysis{
			ID:             uuid.New(),
			WorkspaceID:    uuid.New(),
			WorkspaceName:  uuid.NewString(),
			RepositoryID:   uuid.New(),
			RepositoryName: uuid.NewString(),
			Status:         analysisEnums.Success,
			Errors:         "",
			CreatedAt:      time.Now(),
			FinishedAt:     time.Now(),
			AnalysisVulnerabilities: []analysis.AnalysisVulnerabilities{
				{
					VulnerabilityID: uuid.New(),
					AnalysisID:      uuid.New(),
					CreatedAt:       time.Now(),
					Vulnerability: vulnerability.Vulnerability{
						VulnerabilityID: uuid.New(),
						Line:            "1",
						Column:          "1",
						Confidence:      confidence.High,
						File:            "/deployments/cert.pem",
						Code:            "-----BEGIN CERTIFICATE-----",
						Details:         "Asymmetric Private Key \n Found SSH and/or x.509 Cerficates among the files of your project, make sure you want this kind of information inside your Git repo, since it can be missused by someone with access to any kind of copy.  For more information checkout the CWE-312 (https://cwe.mitre.org/data/definitions/312.html) advisory.",
						SecurityTool:    "Wrong security tool",
						Language:        languages.Leaks,
						Severity:        severities.Critical,
						VulnHash:        "1234567890",
						Type:            vulnerabilityEnum.Vulnerability,
						CommitAuthor:    "Wilian Gabriel",
						CommitEmail:     "wilian.silva@zup.com.br",
						CommitHash:      "9876543210",
						CommitMessage:   "Initial Commit",
						CommitDate:      "2021-03-31T10:58:42Z",
					},
				},
			},
		}
		err := NewRepositoriesAnalysis(connectionMock).CreateFullAnalysis(data)
		assert.Error(t, err)
	})
	t.Run("Should create analysis and many to many with success but not create vulnerability already exists", func(t *testing.T) {
		mockWrite := &database.Mock{}
		mockRead := &database.Mock{}
		mockRead.On("Raw").Return(response.NewResponse(1, errors.New("unexpected error"), nil))
		mockWrite.On("StartTransaction").Return(mockWrite)
		mockWrite.On("RollbackTransaction").Return(response.NewResponse(0, nil, nil))
		mockWrite.On("Create").Return(response.NewResponse(0, nil, nil))
		connectionMock := &database.Connection{
			Write: mockWrite,
			Read:  mockRead,
		}
		data := &analysis.Analysis{
			ID:             uuid.New(),
			WorkspaceID:    uuid.New(),
			WorkspaceName:  uuid.NewString(),
			RepositoryID:   uuid.New(),
			RepositoryName: uuid.NewString(),
			Status:         analysisEnums.Success,
			Errors:         "",
			CreatedAt:      time.Now(),
			FinishedAt:     time.Now(),
			AnalysisVulnerabilities: []analysis.AnalysisVulnerabilities{
				{
					VulnerabilityID: uuid.New(),
					AnalysisID:      uuid.New(),
					CreatedAt:       time.Now(),
					Vulnerability: vulnerability.Vulnerability{
						VulnerabilityID: uuid.New(),
						Line:            "1",
						Column:          "1",
						Confidence:      confidence.High,
						File:            "/deployments/cert.pem",
						Code:            "-----BEGIN CERTIFICATE-----",
						Details:         "Asymmetric Private Key \n Found SSH and/or x.509 Cerficates among the files of your project, make sure you want this kind of information inside your Git repo, since it can be missused by someone with access to any kind of copy.  For more information checkout the CWE-312 (https://cwe.mitre.org/data/definitions/312.html) advisory.",
						SecurityTool:    "Wrong security tool",
						Language:        languages.Leaks,
						Severity:        severities.Critical,
						VulnHash:        "1234567890",
						Type:            vulnerabilityEnum.Vulnerability,
						CommitAuthor:    "Wilian Gabriel",
						CommitEmail:     "wilian.silva@zup.com.br",
						CommitHash:      "9876543210",
						CommitMessage:   "Initial Commit",
						CommitDate:      "2021-03-31T10:58:42Z",
					},
				},
			},
		}
		err := NewRepositoriesAnalysis(connectionMock).CreateFullAnalysis(data)
		assert.Error(t, err)
	})
	t.Run("Should create analysis with error", func(t *testing.T) {
		mockWrite := &database.Mock{}
		mockWrite.On("StartTransaction").Return(mockWrite)
		mockWrite.On("RollbackTransaction").Return(response.NewResponse(0, nil, nil))
		mockWrite.On("Create").Return(response.NewResponse(0, errors.New("unexpected error"), nil))
		connectionMock := &database.Connection{
			Write: mockWrite,
		}
		data := &analysis.Analysis{
			ID:             uuid.New(),
			RepositoryID:   uuid.New(),
			RepositoryName: uuid.New().String(),
			WorkspaceID:    uuid.New(),
			WorkspaceName:  uuid.New().String(),
			Status:         analysisEnums.Success,
			Errors:         "",
			CreatedAt:      time.Now(),
			FinishedAt:     time.Now(),
		}
		err := NewRepositoriesAnalysis(connectionMock).CreateFullAnalysis(data)
		assert.Error(t, err)
	})
	t.Run("Should create analysis with success but create vulnerability with error", func(t *testing.T) {
		mockWrite := &database.Mock{}
		mockRead := &database.Mock{}
		mockRead.On("Raw").Return(response.NewResponse(1, enums.ErrorNotFoundRecords, nil))
		mockWrite.On("StartTransaction").Return(mockWrite)
		mockWrite.On("RollbackTransaction").Return(response.NewResponse(0, nil, nil))
		mockWrite.On("Create").Return(response.NewResponse(0, nil, nil)).Once()
		mockWrite.On("Create").Return(response.NewResponse(0, errors.New("unexpected error"), nil)).Once()
		connectionMock := &database.Connection{
			Write: mockWrite,
			Read:  mockRead,
		}
		data := &analysis.Analysis{
			ID:             uuid.New(),
			WorkspaceID:    uuid.New(),
			WorkspaceName:  uuid.NewString(),
			RepositoryID:   uuid.New(),
			RepositoryName: uuid.NewString(),
			Status:         analysisEnums.Success,
			Errors:         "",
			CreatedAt:      time.Now(),
			FinishedAt:     time.Now(),
			AnalysisVulnerabilities: []analysis.AnalysisVulnerabilities{
				{
					VulnerabilityID: uuid.New(),
					AnalysisID:      uuid.New(),
					CreatedAt:       time.Now(),
					Vulnerability: vulnerability.Vulnerability{
						VulnerabilityID: uuid.New(),
						Line:            "1",
						Column:          "1",
						Confidence:      confidence.High,
						File:            "/deployments/cert.pem",
						Code:            "-----BEGIN CERTIFICATE-----",
						Details:         "Asymmetric Private Key \n Found SSH and/or x.509 Cerficates among the files of your project, make sure you want this kind of information inside your Git repo, since it can be missused by someone with access to any kind of copy.  For more information checkout the CWE-312 (https://cwe.mitre.org/data/definitions/312.html) advisory.",
						SecurityTool:    "Wrong security tool",
						Language:        languages.Leaks,
						Severity:        severities.Critical,
						VulnHash:        "1234567890",
						Type:            vulnerabilityEnum.Vulnerability,
						CommitAuthor:    "Wilian Gabriel",
						CommitEmail:     "wilian.silva@zup.com.br",
						CommitHash:      "9876543210",
						CommitMessage:   "Initial Commit",
						CommitDate:      "2021-03-31T10:58:42Z",
					},
				},
			},
		}
		err := NewRepositoriesAnalysis(connectionMock).CreateFullAnalysis(data)
		assert.Error(t, err)
	})
	t.Run("Should create analysis with success but create many to many with error", func(t *testing.T) {
		mockWrite := &database.Mock{}
		mockRead := &database.Mock{}
		mockRead.On("Raw").Return(response.NewResponse(1, enums.ErrorNotFoundRecords, nil))
		mockWrite.On("StartTransaction").Return(mockWrite)
		mockWrite.On("RollbackTransaction").Return(response.NewResponse(0, nil, nil))
		mockWrite.On("Create").Return(response.NewResponse(0, nil, nil)).Once()
		mockWrite.On("Create").Return(response.NewResponse(0, nil, nil)).Once()
		mockWrite.On("Create").Return(response.NewResponse(0, errors.New("unexpected error"), nil)).Once()
		connectionMock := &database.Connection{
			Write: mockWrite,
			Read:  mockRead,
		}
		data := &analysis.Analysis{
			ID:             uuid.New(),
			WorkspaceID:    uuid.New(),
			WorkspaceName:  uuid.NewString(),
			RepositoryID:   uuid.New(),
			RepositoryName: uuid.NewString(),
			Status:         analysisEnums.Success,
			Errors:         "",
			CreatedAt:      time.Now(),
			FinishedAt:     time.Now(),
			AnalysisVulnerabilities: []analysis.AnalysisVulnerabilities{
				{
					VulnerabilityID: uuid.New(),
					AnalysisID:      uuid.New(),
					CreatedAt:       time.Now(),
					Vulnerability: vulnerability.Vulnerability{
						VulnerabilityID: uuid.New(),
						Line:            "1",
						Column:          "1",
						Confidence:      confidence.High,
						File:            "/deployments/cert.pem",
						Code:            "-----BEGIN CERTIFICATE-----",
						Details:         "Asymmetric Private Key \n Found SSH and/or x.509 Cerficates among the files of your project, make sure you want this kind of information inside your Git repo, since it can be missused by someone with access to any kind of copy.  For more information checkout the CWE-312 (https://cwe.mitre.org/data/definitions/312.html) advisory.",
						SecurityTool:    "Wrong security tool",
						Language:        languages.Leaks,
						Severity:        severities.Critical,
						VulnHash:        "1234567890",
						Type:            vulnerabilityEnum.Vulnerability,
						CommitAuthor:    "Wilian Gabriel",
						CommitEmail:     "wilian.silva@zup.com.br",
						CommitHash:      "9876543210",
						CommitMessage:   "Initial Commit",
						CommitDate:      "2021-03-31T10:58:42Z",
					},
				},
			},
		}
		err := NewRepositoriesAnalysis(connectionMock).CreateFullAnalysis(data)
		assert.Error(t, err)
	})
	t.Run("Should run FindAllVulnerabilitiesByHashesAndRepository", func(t *testing.T) {
		mockRead := &database.Mock{}
		connectionMock := &database.Connection{
			Write: nil,
			Read:  mockRead,
		}
		mockRead.On("Raw").Return(response.NewResponse(1, nil, nil))
		res := NewRepositoriesAnalysis(connectionMock).FindAllVulnerabilitiesByHashesAndRepository([]string{"something"}, uuid.New())
		assert.NoError(t, res.GetError())
	})
	t.Run("Should run Exec", func(t *testing.T) {
		mockWrite := &database.Mock{}
		connectionMock := &database.Connection{
			Write: mockWrite,
			Read:  nil,
		}
		mockWrite.On("Exec").Return(nil)
		err := NewRepositoriesAnalysis(connectionMock).SaveTreatCompatibility(map[string]uuid.UUID{"something": uuid.New()}, &analysis.Analysis{})
		assert.NoError(t, err)
	})
	t.Run("Should run SaveTreatCompatibility with success", func(t *testing.T) {
		mockWrite := &database.Mock{}
		connectionMock := &database.Connection{
			Write: mockWrite,
			Read:  nil,
		}
		mockWrite.On("Exec").Return(nil)
		vulnerabilityID := uuid.New()
		mapHashes := map[string]uuid.UUID{
			"oldHash":  uuid.New(),
			"oldHash1": uuid.New(),
		}
		newAnalysis := &analysis.Analysis{
			AnalysisVulnerabilities: []analysis.AnalysisVulnerabilities{
				{
					VulnerabilityID: vulnerabilityID,
					AnalysisID:      uuid.New(),
					CreatedAt:       time.Now(),
					Vulnerability: vulnerability.Vulnerability{
						VulnerabilityID:  vulnerabilityID,
						Line:             "1",
						Column:           "1",
						Confidence:       confidence.High,
						File:             "/deployments/cert.pem",
						Code:             "-----BEGIN CERTIFICATE-----",
						Details:          "Asymmetric Private Key \n Found SSH and/or x.509 Cerficates among the files of your project, make sure you want this kind of information inside your Git repo, since it can be missused by someone with access to any kind of copy.  For more information checkout the CWE-312 (https://cwe.mitre.org/data/definitions/312.html) advisory.",
						SecurityTool:     "Wrong security tool",
						Language:         languages.Leaks,
						Severity:         severities.Critical,
						VulnHash:         "1234567890",
						Type:             vulnerabilityEnum.Vulnerability,
						CommitAuthor:     "Wilian Gabriel",
						CommitEmail:      "wilian.silva@zup.com.br",
						CommitHash:       "9876543210",
						CommitMessage:    "Initial Commit",
						CommitDate:       "2021-03-31T10:58:42Z",
						DeprecatedHashes: []string{"oldHash", "oldHash1"},
					},
				},
			},
		}
		err := NewRepositoriesAnalysis(connectionMock).SaveTreatCompatibility(mapHashes, newAnalysis)
		assert.NoError(t, err)
	})

	t.Run("Should run SaveTreatCompatibility and return not error but not run any query", func(t *testing.T) {
		mockWrite := &database.Mock{}
		connectionMock := &database.Connection{
			Write: mockWrite,
			Read:  nil,
		}
		mockWrite.On("Exec").Return(nil)
		vulnerabilityID := uuid.New()
		mapHashes := map[string]uuid.UUID{
			"other-hash-not-exists": uuid.New(),
		}
		newAnalysis := &analysis.Analysis{
			AnalysisVulnerabilities: []analysis.AnalysisVulnerabilities{
				{
					VulnerabilityID: vulnerabilityID,
					AnalysisID:      uuid.New(),
					CreatedAt:       time.Now(),
					Vulnerability: vulnerability.Vulnerability{
						VulnerabilityID:  vulnerabilityID,
						Line:             "1",
						Column:           "1",
						Confidence:       confidence.High,
						File:             "/deployments/cert.pem",
						Code:             "-----BEGIN CERTIFICATE-----",
						Details:          "Asymmetric Private Key \n Found SSH and/or x.509 Cerficates among the files of your project, make sure you want this kind of information inside your Git repo, since it can be missused by someone with access to any kind of copy.  For more information checkout the CWE-312 (https://cwe.mitre.org/data/definitions/312.html) advisory.",
						SecurityTool:     "Wrong security tool",
						Language:         languages.Leaks,
						Severity:         severities.Critical,
						VulnHash:         "1234567890",
						Type:             vulnerabilityEnum.Vulnerability,
						CommitAuthor:     "Wilian Gabriel",
						CommitEmail:      "wilian.silva@zup.com.br",
						CommitHash:       "9876543210",
						CommitMessage:    "Initial Commit",
						CommitDate:       "2021-03-31T10:58:42Z",
						DeprecatedHashes: []string{"oldHash", "oldHash1"},
					},
				},
			},
		}
		err := NewRepositoriesAnalysis(connectionMock).SaveTreatCompatibility(mapHashes, newAnalysis)
		assert.NoError(t, err)
	})
}
