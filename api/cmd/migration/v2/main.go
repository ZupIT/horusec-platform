package main

import (
	"fmt"

	"github.com/ZupIT/horusec-devkit/pkg/entities/vulnerability"
	"github.com/ZupIT/horusec-devkit/pkg/services/database"
	databaseconfig "github.com/ZupIT/horusec-devkit/pkg/services/database/config"
	"github.com/ZupIT/horusec-devkit/pkg/utils/crypto"
	"github.com/ZupIT/horusec-devkit/pkg/utils/logger"

	"github.com/ZupIT/horusec-platform/api/cmd/migration/v2/enums"
)

func main() {
	logger.LogInfo("Starting migration...")
	logger.LogInfo("Connect on database")
	databaseConfig := databaseconfig.NewDatabaseConfig()
	conn, err := database.NewDatabaseReadAndWrite(databaseConfig)
	if err != nil {
		logger.LogPanic("{HORUSEC} Error on connection with database", err)
	}
	logger.LogInfo("Checking if exists migrations")
	exists, err := existsMigrations(conn)
	if err != nil {
		logger.LogPanic("{Horusec} Error on get migrations from database", err)
	}
	if exists {
		logger.LogInfo("Migration has been applied with success!")
		return
	}
	logger.LogInfo("Running migration", enums.MigrationV1ToV2Name)
	conn.Write = conn.Write.StartTransaction()
	runMigration(conn.Write, conn.Read)
	logger.LogInfo("Saving migration", enums.MigrationV1ToV2Name)
	saveMigration(conn.Write)
	if err := conn.Write.CommitTransaction().GetError(); err != nil {
		logger.LogPanic("{Horusec} Error on apply transaction", err)
	}
	logger.LogInfo("Migration has been applied with success!")
}

func existsMigrations(conn *database.Connection) (bool, error) {
	var migrations []string
	res := conn.Read.Find(&migrations, map[string]interface{}{}, enums.MigrationTable)
	if res.GetErrorExceptNotFound() != nil {
		return false, res.GetErrorExceptNotFound()
	}
	for _, v := range migrations {
		if v == enums.MigrationV1ToV2Name {
			return true, nil
		}
	}
	return false, nil
}

func runMigration(connWrite database.IDatabaseWrite, connRead database.IDatabaseRead) {
	logger.LogInfo("Getting all vulnerabilities saved on database...")
	vulns := []vulnerability.Vulnerability{}
	r := connRead.Find(&vulns, map[string]interface{}{}, (&vulnerability.Vulnerability{}).GetTable())
	if r.GetErrorExceptNotFound() != nil {
		logger.LogPanic("{HORUSEC} Error on get vulnerabilities with database", r.GetErrorExceptNotFound())
	}

	logger.LogInfo("Recreating vulnerability hash...")
	for key := range vulns {
		vulns[key].VulnHash = recreateVulnerabilityHash(vulns[key])
	}

	for key := range vulns {
		item := vulns[key]
		logger.LogInfo(fmt.Sprintf("Updating vulnerability hash: [%v/%v]", key, len(vulns)-1))
		r := connWrite.Update(
			map[string]interface{}{"vuln_hash": item.VulnHash},
			map[string]interface{}{"vulnerability_id": item.VulnerabilityID},
			(&vulnerability.Vulnerability{}).GetTable(),
		)
		if r.GetErrorExceptNotFound() != nil {
			logger.LogError("{HORUSEC} Error on update vulnerability", r.GetErrorExceptNotFound())
		}
	}
}

func saveMigration(connWrite database.IDatabaseWrite) {
	res := connWrite.Create(map[string]interface{}{"name": enums.MigrationV1ToV2Name}, enums.MigrationTable)
	if res.GetError() != nil {
		logger.LogPanic("{Horusec} Error on save migration in the database", res.GetError())
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
