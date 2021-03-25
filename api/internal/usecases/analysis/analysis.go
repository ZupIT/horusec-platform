package analysis

import (
	"io"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"

	analysisEntity "github.com/ZupIT/horusec-devkit/pkg/entities/analysis"
	"github.com/ZupIT/horusec-devkit/pkg/entities/cli"
	"github.com/ZupIT/horusec-devkit/pkg/entities/vulnerability"
	analysisEnum "github.com/ZupIT/horusec-devkit/pkg/enums/analysis"
	"github.com/ZupIT/horusec-devkit/pkg/enums/languages"
	"github.com/ZupIT/horusec-devkit/pkg/enums/severities"
	"github.com/ZupIT/horusec-devkit/pkg/enums/tools"
	vulnerabilityEnum "github.com/ZupIT/horusec-devkit/pkg/enums/vulnerability"
	"github.com/ZupIT/horusec-devkit/pkg/utils/parser"
)

type Interface interface {
	DecodeAnalysisDataFromIoRead(body io.ReadCloser) (analysisData *cli.AnalysisData, err error)
}

type UseCases struct {
}

func NewAnalysisUseCases() Interface {
	return &UseCases{}
}

func (au *UseCases) DecodeAnalysisDataFromIoRead(body io.ReadCloser) (
	analysisData *cli.AnalysisData, err error) {
	if err := parser.ParseBodyToEntity(body, &analysisData); err != nil {
		return nil, err
	}
	return analysisData, au.validateAnalysis(analysisData.Analysis)
}

func (au *UseCases) validateAnalysis(analysis *analysisEntity.Analysis) error {
	return validation.ValidateStruct(analysis,
		validation.Field(&analysis.ID, validation.Required, is.UUID),
		validation.Field(&analysis.Status, validation.Required,
			validation.In(au.sliceAnalysisStatus()...)),
		validation.Field(&analysis.CreatedAt, validation.Required, validation.NilOrNotEmpty),
		validation.Field(&analysis.FinishedAt, validation.Required, validation.NilOrNotEmpty),
		validation.Field(&analysis.AnalysisVulnerabilities,
			validation.By(au.validateVulnerabilities(analysis.AnalysisVulnerabilities))),
	)
}

func (au *UseCases) validateVulnerabilities(
	analysisVulnerabilities []analysisEntity.RelationshipAnalysisVuln) validation.RuleFunc {
	return func(value interface{}) error {
		for key := range analysisVulnerabilities {
			if err := au.setupValidationVulnerabilities(&analysisVulnerabilities[key].Vulnerability); err != nil {
				return err
			}
		}
		return nil
	}
}

func (au *UseCases) setupValidationVulnerabilities(vul *vulnerability.Vulnerability) error {
	return validation.ValidateStruct(vul,
		validation.Field(&vul.SecurityTool, validation.Required,
			validation.In(au.sliceTools()...)),
		validation.Field(&vul.VulnHash, validation.Required),
		validation.Field(&vul.Language, validation.Required,
			validation.In(au.sliceLanguages()...)),
		validation.Field(&vul.Severity, validation.Required,
			validation.In(au.sliceSeverities()...)),
		validation.Field(&vul.Type, validation.Required,
			validation.In(au.sliceVulnerabilitiesType()...)),
	)
}

// nolint
func (au *UseCases) sliceTools() []interface{} {
	return []interface{}{
		tools.GoSec,
		tools.SecurityCodeScan,
		tools.GitLeaks,
		tools.Brakeman,
		tools.NpmAudit,
		tools.Safety,
		tools.Bandit,
		tools.YarnAudit,
		tools.TfSec,
		tools.HorusecEngine,
		tools.Semgrep,
		tools.Flawfinder,
		tools.PhpCS,
		tools.ShellCheck,
		tools.BundlerAudit,
		tools.Sobelow,
		tools.MixAudit,
	}
}

// nolint
func (au *UseCases) sliceLanguages() []interface{} {
	return []interface{}{
		languages.Go,
		languages.CSharp,
		languages.Dart,
		languages.Ruby,
		languages.Python,
		languages.Java,
		languages.Kotlin,
		languages.Javascript,
		languages.Leaks,
		languages.HCL,
		languages.PHP,
		languages.Typescript,
		languages.C,
		languages.HTML,
		languages.Generic,
		languages.Yaml,
		languages.Shell,
		languages.Elixir,
		languages.Unknown,
	}
}

func (au *UseCases) sliceSeverities() []interface{} {
	return []interface{}{
		severities.Critical,
		severities.High,
		severities.Medium,
		severities.Low,
		severities.Unknown,
		severities.Info,
	}
}

func (au *UseCases) sliceVulnerabilitiesType() []interface{} {
	return []interface{}{
		vulnerabilityEnum.FalsePositive,
		vulnerabilityEnum.RiskAccepted,
		vulnerabilityEnum.Vulnerability,
		vulnerabilityEnum.Corrected,
	}
}

func (au *UseCases) sliceAnalysisStatus() []interface{} {
	return []interface{}{
		analysisEnum.Running,
		analysisEnum.Success,
		analysisEnum.Error,
	}
}
