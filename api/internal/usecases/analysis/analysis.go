package analysis

import (
	netHTTP "net/http"
	"strings"

	analysisv1 "github.com/ZupIT/horusec-platform/api/internal/entities/analysis_v1"

	"github.com/ZupIT/horusec-devkit/pkg/enums/confidence"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"

	"github.com/ZupIT/horusec-devkit/pkg/utils/parser/enums"

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
	DecodeAnalysisDataFromIoRead(r *netHTTP.Request) (analysisData *cli.AnalysisData, err error)
}

type UseCases struct {
	versionsOfV1 []string
}

func NewAnalysisUseCases() Interface {
	return &UseCases{
		versionsOfV1: []string{"v1.7.0", "v1.8.0", "v1.8.1", "v1.8.2", "v1.8.3", "v1.8.4",
			"v1.9.0", "v1.10.0", "v1.10.1", "v1.10.2", "v1.10.3"},
	}
}

func (au *UseCases) DecodeAnalysisDataFromIoRead(r *netHTTP.Request) (
	analysisData *cli.AnalysisData, err error) {
	if r.Body == nil {
		return nil, enums.ErrorBodyEmpty
	}
	analysisData, err = au.parseBodyToAnalysis(r)
	if err != nil {
		return nil, err
	}
	return analysisData, au.validateAnalysisData(analysisData)
}

func (au *UseCases) parseBodyToAnalysis(r *netHTTP.Request) (analysisData *cli.AnalysisData, err error) {
	if au.isVersion1(r.Header.Get("X-Horusec-CLI-Version")) {
		analysisDataV1 := &analysisv1.AnalysisCLIDataV1{}
		err = parser.ParseBodyToEntity(r.Body, analysisDataV1)
		if err != nil {
			return nil, err
		}
		analysisData = analysisDataV1.ParseDataV1ToV2()
	} else {
		err = parser.ParseBodyToEntity(r.Body, &analysisData)
		if err != nil {
			return nil, err
		}
	}
	return analysisData, nil
}

func (au *UseCases) validateAnalysisData(analysisData *cli.AnalysisData) error {
	err := validation.ValidateStruct(analysisData,
		validation.Field(&analysisData.Analysis, validation.Required),
	)
	if err != nil {
		return err
	}
	return au.validateAnalysisToCLIv2(analysisData.Analysis)
}

func (au *UseCases) validateAnalysisToCLIv2(analysis *analysisEntity.Analysis) error {
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
	analysisVulnerabilities []analysisEntity.AnalysisVulnerabilities) validation.RuleFunc {
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
		validation.Field(&vul.SecurityTool, validation.Required, validation.In(au.sliceTools()...)),
		validation.Field(&vul.VulnHash, validation.Required),
		validation.Field(&vul.Confidence, validation.Required, validation.In(au.sliceConfidence()...)),
		validation.Field(&vul.Language, validation.Required, validation.In(au.sliceLanguages()...)),
		validation.Field(&vul.Severity, validation.Required, validation.In(au.sliceSeverities()...)),
		validation.Field(&vul.Type, validation.Required, validation.In(au.sliceVulnerabilitiesType()...)),
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
		languages.Nginx,
	}
}

func (au *UseCases) sliceSeverities() []interface{} {
	return []interface{}{
		severities.Critical,
		severities.High,
		severities.Medium,
		severities.Low,
		severities.Info,
		severities.Unknown,
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

func (au *UseCases) sliceConfidence() []interface{} {
	return []interface{}{
		confidence.High,
		confidence.Medium,
		confidence.Low,
	}
}

func (au *UseCases) isVersion1(versionSent string) bool {
	for _, versionV1 := range au.versionsOfV1 {
		if strings.EqualFold(versionV1, versionSent) {
			return true
		}
	}
	return false
}
