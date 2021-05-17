package main

import (
	"fmt"

	"github.com/ZupIT/horusec-devkit/pkg/entities/vulnerability"
	"github.com/ZupIT/horusec-devkit/pkg/services/database"
	databaseconfig "github.com/ZupIT/horusec-devkit/pkg/services/database/config"
	"github.com/ZupIT/horusec-devkit/pkg/utils/crypto"
	"github.com/ZupIT/horusec-devkit/pkg/utils/logger"
)

func main() {
	logger.LogInfo("Starting migration...")
	logger.LogInfo("Connect on database")
	databaseConfig := databaseconfig.NewDatabaseConfig()
	conn, err := database.NewDatabaseReadAndWrite(databaseConfig)
	if err != nil {
		logger.LogPanic("{HORUSEC} Error on connection with database", err)
	}

	logger.LogInfo("Getting all vulnerabilities saved on database...")
	vulns := []vulnerability.Vulnerability{}
	r := conn.Read.Find(&vulns, map[string]interface{}{}, (&vulnerability.Vulnerability{}).GetTable())
	if r.GetErrorExceptNotFound() != nil {
		logger.LogPanic("{HORUSEC} Error on get vulnerabilities with database", err)
	}

	logger.LogInfo("Recreating vulnerability hash...")
	for key := range vulns {
		vulns[key].VulnHash = recreateVulnerabilityHash(vulns[key])
	}

	for key := range vulns {
		item := vulns[key]
		logger.LogInfo(fmt.Sprintf("Updating vulnerability hash: [%v/%v]", key, len(vulns)-1))
		r := conn.Write.Update(
			map[string]interface{}{"vuln_hash": item.VulnHash},
			map[string]interface{}{"vulnerability_id": item.VulnerabilityID},
			(&vulnerability.Vulnerability{}).GetTable(),
		)
		if r.GetErrorExceptNotFound() != nil {
			logger.LogError("{HORUSEC} Error on update vulnerability", err)
		}
	}
}

func recreateVulnerabilityHash(vuln vulnerability.Vulnerability) string {
	return crypto.GenerateSHA256(
		vuln.Code,
		vuln.Line,
		vuln.Details,
		vuln.File,
		vuln.CommitEmail,
	)
}
