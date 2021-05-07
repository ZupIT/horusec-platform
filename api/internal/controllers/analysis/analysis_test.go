package analysis

import (
	"errors"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	repoAnalysis "github.com/ZupIT/horusec-platform/api/internal/repositories/analysis"
	"github.com/ZupIT/horusec-platform/api/internal/repositories/repository"

	"github.com/ZupIT/horusec-devkit/pkg/entities/analysis"
	"github.com/ZupIT/horusec-devkit/pkg/entities/vulnerability"
	analysisEnum "github.com/ZupIT/horusec-devkit/pkg/enums/analysis"
	"github.com/ZupIT/horusec-devkit/pkg/enums/confidence"
	"github.com/ZupIT/horusec-devkit/pkg/enums/languages"
	"github.com/ZupIT/horusec-devkit/pkg/enums/severities"
	vulnerabilityEnum "github.com/ZupIT/horusec-devkit/pkg/enums/vulnerability"
	"github.com/ZupIT/horusec-devkit/pkg/services/database/enums"
	"github.com/ZupIT/horusec-devkit/pkg/services/database/response"

	appConfiguration "github.com/ZupIT/horusec-devkit/pkg/services/app"
	"github.com/ZupIT/horusec-devkit/pkg/services/broker"
)

func TestController_GetAnalysis(t *testing.T) {
	t.Run("Should return analysis existing from database", func(t *testing.T) {
		brokerMock := &broker.Mock{}
		mockAppConfig := &appConfiguration.Mock{}
		repoRepositoryMock := &repository.Mock{}
		repoAnalysisMock := &repoAnalysis.Mock{}
		repoAnalysisMock.On("FindAnalysisByID").Return(response.NewResponse(0, nil, &analysis.Analysis{
			ID:         uuid.New(),
			Status:     analysisEnum.Success,
			Errors:     "",
			CreatedAt:  time.Now(),
			FinishedAt: time.Now(),
		}))
		controller := NewAnalysisController(
			brokerMock,
			mockAppConfig,
			repoRepositoryMock,
			repoAnalysisMock,
		)
		res, err := controller.GetAnalysis(uuid.New())
		assert.NoError(t, err)
		assert.NotEmpty(t, res)
	})
	t.Run("Should return error when get analysis from database", func(t *testing.T) {
		brokerMock := &broker.Mock{}
		mockAppConfig := &appConfiguration.Mock{}
		repoRepositoryMock := &repository.Mock{}
		repoAnalysisMock := &repoAnalysis.Mock{}
		repoAnalysisMock.On("FindAnalysisByID").Return(response.NewResponse(0, errors.New("unexpected error"), nil))
		controller := NewAnalysisController(
			brokerMock,
			mockAppConfig,
			repoRepositoryMock,
			repoAnalysisMock,
		)
		res, err := controller.GetAnalysis(uuid.New())
		assert.Error(t, err)
		assert.Empty(t, res)
	})
	t.Run("Should return error not found records when get analysis from database", func(t *testing.T) {
		brokerMock := &broker.Mock{}
		mockAppConfig := &appConfiguration.Mock{}
		repoRepositoryMock := &repository.Mock{}
		repoAnalysisMock := &repoAnalysis.Mock{}
		repoAnalysisMock.On("FindAnalysisByID").Return(response.NewResponse(0, enums.ErrorNotFoundRecords, nil))
		controller := NewAnalysisController(
			brokerMock,
			mockAppConfig,
			repoRepositoryMock,
			repoAnalysisMock,
		)
		res, err := controller.GetAnalysis(uuid.New())
		assert.Error(t, err)
		assert.Equal(t, err, enums.ErrorNotFoundRecords)
		assert.Empty(t, res)
	})
	t.Run("Should not found data return not found records error", func(t *testing.T) {
		brokerMock := &broker.Mock{}
		mockAppConfig := &appConfiguration.Mock{}
		repoRepositoryMock := &repository.Mock{}
		repoAnalysisMock := &repoAnalysis.Mock{}
		repoAnalysisMock.On("FindAnalysisByID").Return(response.NewResponse(0, nil, nil))
		controller := NewAnalysisController(
			brokerMock,
			mockAppConfig,
			repoRepositoryMock,
			repoAnalysisMock,
		)
		res, err := controller.GetAnalysis(uuid.New())
		assert.Error(t, err)
		assert.Equal(t, err, enums.ErrorNotFoundRecords)
		assert.Empty(t, res)
	})
}

func TestController_SaveAnalysis(t *testing.T) {
	t.Run("Should save analysis with success simple", func(t *testing.T) {
		brokerMock := &broker.Mock{}
		appConfigMock := &appConfiguration.Mock{}
		appConfigMock.On("IsEmailsDisabled").Return(true)
		repoRepositoryMock := &repository.Mock{}
		repoAnalysisMock := &repoAnalysis.Mock{}
		repoAnalysisMock.On("CreateFullAnalysisResponse").Return(nil)
		repoAnalysisMock.On("CreateFullAnalysisArguments").Return(func(any *analysis.Analysis) {})
		repoAnalysisMock.On("FindAnalysisByID").Return(response.NewResponse(0, nil, &analysis.Analysis{
			ID:         uuid.New(),
			Status:     analysisEnum.Success,
			Errors:     "",
			CreatedAt:  time.Now(),
			FinishedAt: time.Now(),
		}))
		controller := NewAnalysisController(
			brokerMock,
			appConfigMock,
			repoRepositoryMock,
			repoAnalysisMock,
		)
		res, err := controller.SaveAnalysis(&analysis.Analysis{
			ID:             uuid.New(),
			WorkspaceID:    uuid.New(),
			WorkspaceName:  uuid.NewString(),
			RepositoryID:   uuid.New(),
			RepositoryName: uuid.NewString(),
			Status:         analysisEnum.Success,
			Errors:         "",
			CreatedAt:      time.Now(),
			FinishedAt:     time.Now(),
		})
		assert.NoError(t, err)
		assert.NotEqual(t, res, uuid.Nil)
	})
	t.Run("Should save analysis with success when exists vulnerability", func(t *testing.T) {
		brokerMock := &broker.Mock{}
		appConfigMock := &appConfiguration.Mock{}
		appConfigMock.On("IsEmailsDisabled").Return(true)
		repoRepositoryMock := &repository.Mock{}
		repoAnalysisMock := &repoAnalysis.Mock{}
		repoAnalysisMock.On("CreateFullAnalysisResponse").Return(nil)
		repoAnalysisMock.On("CreateFullAnalysisArguments").Return(func(any *analysis.Analysis) {})
		repoAnalysisMock.On("FindAnalysisByID").Return(response.NewResponse(0, nil, &analysis.Analysis{
			ID:         uuid.New(),
			Status:     analysisEnum.Success,
			Errors:     "",
			CreatedAt:  time.Now(),
			FinishedAt: time.Now(),
		}))
		controller := NewAnalysisController(
			brokerMock,
			appConfigMock,
			repoRepositoryMock,
			repoAnalysisMock,
		)
		res, err := controller.SaveAnalysis(&analysis.Analysis{
			ID:             uuid.New(),
			WorkspaceID:    uuid.New(),
			WorkspaceName:  uuid.NewString(),
			RepositoryID:   uuid.New(),
			RepositoryName: uuid.NewString(),
			Status:         analysisEnum.Success,
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
		})
		assert.NoError(t, err)
		assert.NotEqual(t, res, uuid.Nil)
	})
	t.Run("Should save analysis with success when exists vulnerability with duplicated hash and sent only one vulnerability to save", func(t *testing.T) {
		brokerMock := &broker.Mock{}
		appConfigMock := &appConfiguration.Mock{}
		appConfigMock.On("IsEmailsDisabled").Return(true)
		repoRepositoryMock := &repository.Mock{}
		repoAnalysisMock := &repoAnalysis.Mock{}
		repoAnalysisMock.On("CreateFullAnalysisResponse").Return(nil)
		repoAnalysisMock.On("CreateFullAnalysisArguments").Return(func(arguments *analysis.Analysis) {
			assert.Len(t, arguments.AnalysisVulnerabilities, 1)
		})
		repoAnalysisMock.On("FindAnalysisByID").Return(response.NewResponse(0, nil, &analysis.Analysis{
			ID:         uuid.New(),
			Status:     analysisEnum.Success,
			Errors:     "",
			CreatedAt:  time.Now(),
			FinishedAt: time.Now(),
		}))
		controller := NewAnalysisController(
			brokerMock,
			appConfigMock,
			repoRepositoryMock,
			repoAnalysisMock,
		)
		dataToSave := &analysis.Analysis{
			ID:             uuid.New(),
			WorkspaceID:    uuid.New(),
			WorkspaceName:  uuid.NewString(),
			RepositoryID:   uuid.New(),
			RepositoryName: uuid.NewString(),
			Status:         analysisEnum.Success,
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
		res, err := controller.SaveAnalysis(dataToSave)
		assert.NoError(t, err)
		assert.NotEqual(t, res, uuid.Nil)
	})
	t.Run("Should save analysis with success and not create repository because already exists", func(t *testing.T) {
		brokerMock := &broker.Mock{}
		appConfigMock := &appConfiguration.Mock{}
		appConfigMock.On("IsEmailsDisabled").Return(true)
		repoRepositoryMock := &repository.Mock{}
		repoRepositoryMock.On("FindRepository").Return(uuid.New(), nil)
		repoAnalysisMock := &repoAnalysis.Mock{}
		repoAnalysisMock.On("CreateFullAnalysisResponse").Return(nil)
		repoAnalysisMock.On("CreateFullAnalysisArguments").Return(func(any *analysis.Analysis) {})
		repoAnalysisMock.On("FindAnalysisByID").Return(response.NewResponse(0, nil, &analysis.Analysis{
			ID:         uuid.New(),
			Status:     analysisEnum.Success,
			Errors:     "",
			CreatedAt:  time.Now(),
			FinishedAt: time.Now(),
		}))
		controller := NewAnalysisController(
			brokerMock,
			appConfigMock,
			repoRepositoryMock,
			repoAnalysisMock,
		)
		res, err := controller.SaveAnalysis(&analysis.Analysis{
			ID:             uuid.New(),
			WorkspaceID:    uuid.New(),
			WorkspaceName:  uuid.NewString(),
			RepositoryID:   uuid.Nil,
			RepositoryName: uuid.NewString(),
			Status:         analysisEnum.Success,
			Errors:         "",
			CreatedAt:      time.Now(),
			FinishedAt:     time.Now(),
		})
		assert.NoError(t, err)
		assert.NotEqual(t, res, uuid.Nil)
	})
	t.Run("Should return error unknown when find repository", func(t *testing.T) {
		brokerMock := &broker.Mock{}
		appConfigMock := &appConfiguration.Mock{}
		repoRepositoryMock := &repository.Mock{}
		repoRepositoryMock.On("FindRepository").Return(uuid.Nil, errors.New("unexpected error"))
		repoRepositoryMock.On("CreateRepository").Return(nil)
		repoAnalysisMock := &repoAnalysis.Mock{}
		controller := NewAnalysisController(
			brokerMock,
			appConfigMock,
			repoRepositoryMock,
			repoAnalysisMock,
		)
		res, err := controller.SaveAnalysis(&analysis.Analysis{
			ID:             uuid.New(),
			WorkspaceID:    uuid.New(),
			WorkspaceName:  uuid.NewString(),
			RepositoryID:   uuid.Nil,
			RepositoryName: uuid.NewString(),
			Status:         analysisEnum.Success,
			Errors:         "",
			CreatedAt:      time.Now(),
			FinishedAt:     time.Now(),
		})
		assert.Error(t, err)
		assert.Equal(t, res, uuid.Nil)
	})
	t.Run("Should return error unknown when find repository", func(t *testing.T) {
		brokerMock := &broker.Mock{}
		appConfigMock := &appConfiguration.Mock{}
		repoRepositoryMock := &repository.Mock{}
		repoRepositoryMock.On("FindRepository").Return(uuid.Nil, errors.New("unexpected error"))
		repoRepositoryMock.On("CreateRepository").Return(nil)
		repoAnalysisMock := &repoAnalysis.Mock{}
		controller := NewAnalysisController(
			brokerMock,
			appConfigMock,
			repoRepositoryMock,
			repoAnalysisMock,
		)
		res, err := controller.SaveAnalysis(&analysis.Analysis{
			ID:             uuid.New(),
			WorkspaceID:    uuid.New(),
			WorkspaceName:  uuid.NewString(),
			RepositoryID:   uuid.Nil,
			RepositoryName: uuid.NewString(),
			Status:         analysisEnum.Success,
			Errors:         "",
			CreatedAt:      time.Now(),
			FinishedAt:     time.Now(),
		})
		assert.Error(t, err)
		assert.Equal(t, res, uuid.Nil)
	})
	t.Run("Should return error when create new repository", func(t *testing.T) {
		brokerMock := &broker.Mock{}
		appConfigMock := &appConfiguration.Mock{}
		repoRepositoryMock := &repository.Mock{}
		errCreateRepository := errors.New("unexpected error")
		repoRepositoryMock.On("FindRepository").Return(uuid.Nil, enums.ErrorNotFoundRecords)
		repoRepositoryMock.On("CreateRepository").Return(errCreateRepository)
		repoAnalysisMock := &repoAnalysis.Mock{}
		controller := NewAnalysisController(
			brokerMock,
			appConfigMock,
			repoRepositoryMock,
			repoAnalysisMock,
		)
		res, err := controller.SaveAnalysis(&analysis.Analysis{
			ID:             uuid.New(),
			WorkspaceID:    uuid.New(),
			WorkspaceName:  uuid.NewString(),
			RepositoryID:   uuid.Nil,
			RepositoryName: uuid.NewString(),
			Status:         analysisEnum.Success,
			Errors:         "",
			CreatedAt:      time.Now(),
			FinishedAt:     time.Now(),
		})
		assert.Error(t, err)
		assert.Equal(t, err, errCreateRepository)
		assert.Equal(t, res, uuid.Nil)
	})
	t.Run("Should return error when create analysis", func(t *testing.T) {
		brokerMock := &broker.Mock{}
		appConfigMock := &appConfiguration.Mock{}
		appConfigMock.On("IsEmailsDisabled").Return(true)
		repoRepositoryMock := &repository.Mock{}
		repoAnalysisMock := &repoAnalysis.Mock{}
		repoAnalysisMock.On("CreateFullAnalysisResponse").Return(errors.New("unexpected error"))
		repoAnalysisMock.On("CreateFullAnalysisArguments").Return(func(any *analysis.Analysis) {})
		repoAnalysisMock.On("FindAnalysisByID").Return(response.NewResponse(0, nil, &analysis.Analysis{
			ID:         uuid.New(),
			Status:     analysisEnum.Success,
			Errors:     "",
			CreatedAt:  time.Now(),
			FinishedAt: time.Now(),
		}))
		controller := NewAnalysisController(
			brokerMock,
			appConfigMock,
			repoRepositoryMock,
			repoAnalysisMock,
		)
		res, err := controller.SaveAnalysis(&analysis.Analysis{
			ID:             uuid.New(),
			WorkspaceID:    uuid.New(),
			WorkspaceName:  uuid.NewString(),
			RepositoryID:   uuid.New(),
			RepositoryName: uuid.NewString(),
			Status:         analysisEnum.Success,
			Errors:         "",
			CreatedAt:      time.Now(),
			FinishedAt:     time.Now(),
		})
		assert.Error(t, err)
		assert.Equal(t, res, uuid.Nil)
	})
	t.Run("Should save analysis with success simple and publish to webhook queue", func(t *testing.T) {
		brokerMock := &broker.Mock{}
		brokerMock.On("Publish").Return(nil)
		appConfigMock := &appConfiguration.Mock{}
		appConfigMock.On("IsEmailsDisabled").Return(false)
		repoRepositoryMock := &repository.Mock{}
		repoAnalysisMock := &repoAnalysis.Mock{}
		repoAnalysisMock.On("CreateFullAnalysisResponse").Return(nil)
		repoAnalysisMock.On("CreateFullAnalysisArguments").Return(func(any *analysis.Analysis) {})
		repoAnalysisMock.On("FindAnalysisByID").Return(response.NewResponse(0, nil, &analysis.Analysis{
			ID:         uuid.New(),
			Status:     analysisEnum.Success,
			Errors:     "",
			CreatedAt:  time.Now(),
			FinishedAt: time.Now(),
		}))
		controller := NewAnalysisController(
			brokerMock,
			appConfigMock,
			repoRepositoryMock,
			repoAnalysisMock,
		)
		res, err := controller.SaveAnalysis(&analysis.Analysis{
			ID:             uuid.New(),
			WorkspaceID:    uuid.New(),
			WorkspaceName:  uuid.NewString(),
			RepositoryID:   uuid.New(),
			RepositoryName: uuid.NewString(),
			Status:         analysisEnum.Success,
			Errors:         "",
			CreatedAt:      time.Now(),
			FinishedAt:     time.Now(),
		})
		assert.NoError(t, err)
		assert.NotEqual(t, res, uuid.Nil)
	})
	t.Run("Should save analysis with error when get analysis to publish in broker queue", func(t *testing.T) {
		brokerMock := &broker.Mock{}
		brokerMock.On("Publish").Return(errors.New("unexpected error"))
		appConfigMock := &appConfiguration.Mock{}
		appConfigMock.On("IsEmailsDisabled").Return(false)
		repoRepositoryMock := &repository.Mock{}
		repoAnalysisMock := &repoAnalysis.Mock{}
		repoAnalysisMock.On("CreateFullAnalysisResponse").Return(nil)
		repoAnalysisMock.On("CreateFullAnalysisArguments").Return(func(any *analysis.Analysis) {})
		repoAnalysisMock.On("FindAnalysisByID").Return(response.NewResponse(0, errors.New("unexpected error"), nil))
		controller := NewAnalysisController(
			brokerMock,
			appConfigMock,
			repoRepositoryMock,
			repoAnalysisMock,
		)
		res, err := controller.SaveAnalysis(&analysis.Analysis{
			ID:             uuid.New(),
			WorkspaceID:    uuid.New(),
			WorkspaceName:  uuid.NewString(),
			RepositoryID:   uuid.New(),
			RepositoryName: uuid.NewString(),
			Status:         analysisEnum.Success,
			Errors:         "",
			CreatedAt:      time.Now(),
			FinishedAt:     time.Now(),
		})
		assert.Error(t, err)
		assert.Equal(t, res, uuid.Nil)
	})
	t.Run("Should save analysis with error when publish in broker queue", func(t *testing.T) {
		brokerMock := &broker.Mock{}
		brokerMock.On("Publish").Return(errors.New("unexpected error"))
		appConfigMock := &appConfiguration.Mock{}
		appConfigMock.On("IsEmailsDisabled").Return(false)
		repoRepositoryMock := &repository.Mock{}
		repoAnalysisMock := &repoAnalysis.Mock{}
		repoAnalysisMock.On("CreateFullAnalysisResponse").Return(nil)
		repoAnalysisMock.On("CreateFullAnalysisArguments").Return(func(any *analysis.Analysis) {})
		repoAnalysisMock.On("FindAnalysisByID").Return(response.NewResponse(0, nil, &analysis.Analysis{
			ID:         uuid.New(),
			Status:     analysisEnum.Success,
			Errors:     "",
			CreatedAt:  time.Now(),
			FinishedAt: time.Now(),
		}))
		controller := NewAnalysisController(
			brokerMock,
			appConfigMock,
			repoRepositoryMock,
			repoAnalysisMock,
		)
		res, err := controller.SaveAnalysis(&analysis.Analysis{
			ID:             uuid.New(),
			WorkspaceID:    uuid.New(),
			WorkspaceName:  uuid.NewString(),
			RepositoryID:   uuid.New(),
			RepositoryName: uuid.NewString(),
			Status:         analysisEnum.Success,
			Errors:         "",
			CreatedAt:      time.Now(),
			FinishedAt:     time.Now(),
		})
		assert.Error(t, err)
		assert.Equal(t, res, uuid.Nil)
	})
}
