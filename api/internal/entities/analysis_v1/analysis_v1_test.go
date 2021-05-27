package analysisv1

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAnalysisCLIDataV1_ParseDataV1ToV2(t *testing.T) {
	t.Run("Should parse analsyis v1 to v2 without problems", func(t *testing.T) {
		path, err := os.Getwd()
		assert.NoError(t, err)
		jsonFilePath := filepath.Join(path, "analysis_v1_mock.json")
		content, err := os.ReadFile(jsonFilePath)
		assert.NoError(t, err)
		entity := &AnalysisCLIDataV1{}
		err = json.Unmarshal(content, entity)
		assert.NoError(t, err)
		assert.NotEmpty(t, entity.ParseDataV1ToV2())
	})
}
