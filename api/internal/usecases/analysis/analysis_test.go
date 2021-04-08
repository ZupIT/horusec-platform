package analysis

import (
	"bytes"
	"io/ioutil"
	"testing"
	"time"

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
		response, err := NewAnalysisUseCases().DecodeAnalysisDataFromIoRead(body)
		assert.NoError(t, err)
		assert.NotEmpty(t, response)
	})
	t.Run("Should return error when body not exists", func(t *testing.T) {
		_, err := NewAnalysisUseCases().DecodeAnalysisDataFromIoRead(nil)
		assert.Equal(t, err, enums.ErrorBodyEmpty)
	})
	t.Run("Should return error when body is empty", func(t *testing.T) {
		body := ioutil.NopCloser(bytes.NewBufferString(""))
		_, err := NewAnalysisUseCases().DecodeAnalysisDataFromIoRead(body)
		assert.Equal(t, err, enums.ErrorBodyEmpty)
	})
	t.Run("Should return error when body is wrong", func(t *testing.T) {
		body := ioutil.NopCloser(bytes.NewBufferString("some incorrect body"))
		_, err := NewAnalysisUseCases().DecodeAnalysisDataFromIoRead(body)
		assert.Equal(t, err, enums.ErrorBodyInvalid)
	})
	t.Run("Should return unknown error when parse body", func(t *testing.T) {
		body := ioutil.NopCloser(bytes.NewBufferString("00000"))
		_, err := NewAnalysisUseCases().DecodeAnalysisDataFromIoRead(body)
		assert.Error(t, err)
	})
	t.Run("Should return error when not found analysis in body", func(t *testing.T) {
		body := ioutil.NopCloser(bytes.NewBufferString("{}"))
		_, err := NewAnalysisUseCases().DecodeAnalysisDataFromIoRead(body)
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
						Confidence:      confidence.High.ToString(),
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
		response, err := NewAnalysisUseCases().DecodeAnalysisDataFromIoRead(body)
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
						Confidence:      confidence.High.ToString(),
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
		_, err = NewAnalysisUseCases().DecodeAnalysisDataFromIoRead(body)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "securityTool: must be a valid value")
	})
}
