package analysis

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"testing"
	"time"

	analysisv1 "github.com/ZupIT/horusec-platform/api/internal/entities/analysis_v1"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/ZupIT/horusec-devkit/pkg/entities/analysis"
	"github.com/ZupIT/horusec-devkit/pkg/entities/cli"
	"github.com/ZupIT/horusec-devkit/pkg/entities/vulnerability"
	analysisEnum "github.com/ZupIT/horusec-devkit/pkg/enums/analysis"
	"github.com/ZupIT/horusec-devkit/pkg/enums/confidence"
	"github.com/ZupIT/horusec-devkit/pkg/enums/languages"
	"github.com/ZupIT/horusec-devkit/pkg/enums/severities"
	"github.com/ZupIT/horusec-devkit/pkg/enums/tools"
	vulnerabilityEnum "github.com/ZupIT/horusec-devkit/pkg/enums/vulnerability"
	"github.com/ZupIT/horusec-devkit/pkg/utils/parser"
	"github.com/ZupIT/horusec-devkit/pkg/utils/parser/enums"
)

func TestUseCases_DecodeAnalysisDataFromIoRead(t *testing.T) {
	t.Run("Should decode with success", func(t *testing.T) {
		analysisEntity := &analysis.Analysis{
			ID:         uuid.New(),
			Status:     analysisEnum.Success,
			Errors:     "",
			CreatedAt:  time.Now(),
			FinishedAt: time.Now(),
		}
		analysisData := &cli.AnalysisData{
			Analysis: analysisEntity,
		}
		body, err := parser.ParseEntityToIOReadCloser(analysisData)
		assert.NoError(t, err)
		r, _ := http.NewRequest(http.MethodPost, "/test", body)
		response, err := NewAnalysisUseCases().DecodeAnalysisDataFromIoRead(r)
		assert.NoError(t, err)
		assert.NotEmpty(t, response)
	})
	t.Run("Should return error when body not exists", func(t *testing.T) {
		r, _ := http.NewRequest(http.MethodPost, "/test", nil)
		_, err := NewAnalysisUseCases().DecodeAnalysisDataFromIoRead(r)
		assert.Equal(t, err, enums.ErrorBodyEmpty)
	})
	t.Run("Should return error when body is empty", func(t *testing.T) {
		body := ioutil.NopCloser(bytes.NewBufferString(""))
		r, _ := http.NewRequest(http.MethodPost, "/test", body)
		_, err := NewAnalysisUseCases().DecodeAnalysisDataFromIoRead(r)
		assert.Equal(t, err, enums.ErrorBodyEmpty)
	})
	t.Run("Should return error when body is wrong", func(t *testing.T) {
		body := ioutil.NopCloser(bytes.NewBufferString("some incorrect body"))
		r, _ := http.NewRequest(http.MethodPost, "/test", body)
		_, err := NewAnalysisUseCases().DecodeAnalysisDataFromIoRead(r)
		assert.Equal(t, err, enums.ErrorBodyInvalid)
	})
	t.Run("Should return unknown error when parse body", func(t *testing.T) {
		body := ioutil.NopCloser(bytes.NewBufferString("00000"))
		r, _ := http.NewRequest(http.MethodPost, "/test", body)
		_, err := NewAnalysisUseCases().DecodeAnalysisDataFromIoRead(r)
		assert.Error(t, err)
	})
	t.Run("Should return error when not found analysis in body", func(t *testing.T) {
		body := ioutil.NopCloser(bytes.NewBufferString("{}"))
		r, _ := http.NewRequest(http.MethodPost, "/test", body)
		_, err := NewAnalysisUseCases().DecodeAnalysisDataFromIoRead(r)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "analysis: cannot be blank")
	})
	t.Run("Should decode with success with vulnerabilities", func(t *testing.T) {
		analysisEntity := &analysis.Analysis{
			ID:         uuid.New(),
			Status:     analysisEnum.Success,
			Errors:     "",
			CreatedAt:  time.Now(),
			FinishedAt: time.Now(),
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
						SecurityTool:    tools.HorusecEngine,
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
		analysisData := &cli.AnalysisData{
			Analysis: analysisEntity,
		}
		body, err := parser.ParseEntityToIOReadCloser(analysisData)
		assert.NoError(t, err)
		r, _ := http.NewRequest(http.MethodPost, "/test", body)
		response, err := NewAnalysisUseCases().DecodeAnalysisDataFromIoRead(r)
		assert.NoError(t, err)
		assert.NotEmpty(t, response)
	})
	t.Run("Should return error when not exists confidence", func(t *testing.T) {
		analysisEntity := &analysis.Analysis{
			ID:         uuid.New(),
			Status:     analysisEnum.Success,
			Errors:     "",
			CreatedAt:  time.Now(),
			FinishedAt: time.Now(),
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
		analysisData := &cli.AnalysisData{
			Analysis: analysisEntity,
		}
		body, err := parser.ParseEntityToIOReadCloser(analysisData)
		assert.NoError(t, err)
		r, _ := http.NewRequest(http.MethodPost, "/test", body)
		_, err = NewAnalysisUseCases().DecodeAnalysisDataFromIoRead(r)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "securityTool: must be a valid value")
	})
	t.Run("Should run useCase with v1 and return error with wrong body", func(t *testing.T) {
		body := ioutil.NopCloser(bytes.NewBufferString("some incorrect body"))
		r, _ := http.NewRequest(http.MethodPost, "/test", body)
		r.Header.Set("X-Horusec-CLI-Version", "v1.10.3")
		_, err := NewAnalysisUseCases().DecodeAnalysisDataFromIoRead(r)
		assert.Equal(t, err, enums.ErrorBodyInvalid)
	})
	t.Run("Should run useCase with v1 and not return error", func(t *testing.T) {
		content, err := os.ReadFile("../../entities/analysis_v1/analysis_v1_mock.json")
		assert.NoError(t, err)
		entity := &analysisv1.AnalysisCLIDataV1{}
		err = json.Unmarshal(content, entity)
		assert.NoError(t, err)
		ioReader, err := parser.ParseEntityToIOReadCloser(entity)
		assert.NoError(t, err)
		r, _ := http.NewRequest(http.MethodPost, "/test", ioReader)
		r.Header.Set("X-Horusec-CLI-Version", "v1.10.3")
		dataToCheck, err := NewAnalysisUseCases().DecodeAnalysisDataFromIoRead(r)
		assert.NoError(t, err)
		assert.NotEmpty(t, dataToCheck)
		assert.NotEmpty(t, dataToCheck.Analysis)
		assert.GreaterOrEqual(t, len(dataToCheck.Analysis.AnalysisVulnerabilities), 1)
		for _, item := range dataToCheck.Analysis.AnalysisVulnerabilities {
			assert.NotEmpty(t, item)
			assert.NotEmpty(t, item.VulnerabilityID)
		}
	})
}
