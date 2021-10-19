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

package hash

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/google/uuid"

	"github.com/ZupIT/horusec-devkit/pkg/entities/analysis"
	"github.com/ZupIT/horusec-devkit/pkg/entities/vulnerability"
	"github.com/ZupIT/horusec-devkit/pkg/services/database"
	"github.com/ZupIT/horusec-devkit/pkg/services/database/config"
	"github.com/ZupIT/horusec-devkit/pkg/utils/crypto"
	"github.com/ZupIT/horusec-devkit/pkg/utils/logger"
)

const MigrationName = "20211019_hash_migration"
const MigrationTable = "horusec_migrations"

type Migration struct {
	databaseConnection *database.Connection
}

type Vulnerability struct {
	*vulnerability.Vulnerability
	AnalysisID uuid.UUID
}

func StartMigration() {
	databaseConnection, err := database.NewDatabaseReadAndWrite(config.NewDatabaseConfig())
	if err != nil {
		logger.LogPanic("HASH MIGRATION - error on connection with database", err)
	}

	migration := Migration{
		databaseConnection: databaseConnection,
	}

	if migration.alreadyMigrated() {
		return
	}

	migration.migrateHashesToRemoveVulnID()
	migration.setMigrationAsApplied()
}

func (m *Migration) alreadyMigrated() bool {
	var migrations []string

	err := m.databaseConnection.Read.Find(&migrations, map[string]interface{}{}, MigrationTable).GetErrorExceptNotFound()
	if err != nil {
		logger.LogError("HASH MIGRATION - failed to check if migration is already applied", err)
	}

	for _, v := range migrations {
		if v == MigrationName {
			return true
		}
	}

	return false
}

func (m *Migration) migrateHashesToRemoveVulnID() {
	vulnerabilities, err := m.getVulnerabilities()
	if err != nil {
		logger.LogWarn("HASH MIGRATION - failed to get vulnerabilities to update", err)
	}

	for _, vuln := range vulnerabilities {
		m.updateVulnerability(vuln)
	}
}

func (m *Migration) getVulnerabilities() ([]*Vulnerability, error) {
	var vulnerabilities []*Vulnerability

	return vulnerabilities,
		m.databaseConnection.Read.Raw(m.querySelectLatestVulnerabilities(), &vulnerabilities).GetErrorExceptNotFound()
}

func (m *Migration) querySelectLatestVulnerabilities() string {
	return `
		SELECT vuln.vulnerability_id, vuln.line, vuln.code, vuln.details, vuln.file, vuln.commit_email, vuln.vuln_hash,
			   av.analysis_id
		FROM vulnerabilities AS vuln
		INNER JOIN analysis_vulnerabilities AS av
		ON vuln.vulnerability_id = av.vulnerability_id
		INNER JOIN (
			SELECT DISTINCT ON (repository_id) MAX (an.created_at) AS max_time, an.analysis_id 
			FROM analysis AS an
			GROUP BY an.repository_id, an.analysis_id
			ORDER BY an.repository_id, an.created_at DESC
		) AS latest_analysis 
		ON latest_analysis.analysis_id = av.analysis_id
		WHERE vuln.details LIKE ANY (array['%HS-JAVA%', '%HS-JS%', '%HS-CSHARP%', 
		'%HS-DART%', '%HS-JVM%', '%HS-KOTLIN%', '%HS-KUBERNETES%', '%HS-LEAKS%', '%HS-NGINX%',
		'%HS-JAVASCRIPT%', '%HS-SWIFT%'])
	`
}

func (m *Migration) updateVulnerability(vuln *Vulnerability) {
	expectHash := m.generateExpectedHash(vuln)

	if vuln.VulnHash != expectHash {
		correctVulnerability, err := m.getVulnerabilityByHash(expectHash)
		if err != nil {
			return
		}

		vulnerabilityNToN, err := m.getVulnerabilityNToN(vuln.AnalysisID, vuln.VulnerabilityID)
		if err != nil {
			return
		}

		vulnerabilityNToN.VulnerabilityID = correctVulnerability.VulnerabilityID
		err = m.databaseConnection.Write.Update(vulnerabilityNToN,
			m.filterManyToManyByID(vuln.AnalysisID, vuln.VulnerabilityID), vulnerabilityNToN.GetTable()).GetError()
		if err != nil {
			logger.LogError("HASH MIGRATION - failed to update vulnerability n to n", err)
			return
		}

		logger.LogInfo(fmt.Sprintf("HASH MIGRATION - hash %s migrated", vuln.VulnHash))
	}
}

func (m *Migration) getVulnerabilityByHash(hash string) (*vulnerability.Vulnerability, error) {
	vuln := &vulnerability.Vulnerability{}

	return vuln, m.databaseConnection.Read.Find(vuln, m.filterVulnerabilityByHash(hash), vuln.GetTable()).GetError()
}

func (m *Migration) filterVulnerabilityByHash(hash string) map[string]interface{} {
	return map[string]interface{}{"vuln_hash": hash}
}

func (m *Migration) getVulnerabilityNToN(analysisID, vulnerabilityID uuid.UUID) (*analysis.AnalysisVulnerabilities, error) {
	vulnerabilityNToN := &analysis.AnalysisVulnerabilities{}

	return vulnerabilityNToN, m.databaseConnection.Read.Find(vulnerabilityNToN,
		m.filterManyToManyByID(analysisID, vulnerabilityID), vulnerabilityNToN.GetTable()).GetError()
}

func (m *Migration) filterManyToManyByID(analysisID, vulnerabilityID uuid.UUID) map[string]interface{} {
	return map[string]interface{}{"vulnerability_id": vulnerabilityID, "analysis_id": analysisID}
}

func (m *Migration) generateExpectedHash(vuln *Vulnerability) string {
	return crypto.GenerateSHA256(
		m.parseCodeToOneLine(vuln.Code),
		vuln.Line,
		m.removeHorusecIDFromDetails(vuln.Details),
		vuln.File,
		vuln.CommitEmail,
	)
}

func (m *Migration) parseCodeToOneLine(code string) string {
	const oneCodeLineRegex = `\r?\n?\t`

	return strings.ReplaceAll(regexp.MustCompile(oneCodeLineRegex).ReplaceAllString(code, " "), " ", "")
}

func (m *Migration) removeHorusecIDFromDetails(details string) string {
	const horusecIDRegex = `HS-(JAVA|JS|CSHARP|DART|JVM|KOTLIN|KUBERNETES|LEAKS|NGINX|JAVASCRIPT|SWIFT)-[0-9]+:\s`

	return regexp.MustCompile(horusecIDRegex).ReplaceAllString(details, "")
}

func (m *Migration) setMigrationAsApplied() {
	err := m.databaseConnection.Write.Create(map[string]interface{}{"name": MigrationName}, MigrationTable).GetError()
	if err != nil {
		logger.LogPanic("HASH MIGRATION- failed to set that migration was applied on database", err)
	}
}
