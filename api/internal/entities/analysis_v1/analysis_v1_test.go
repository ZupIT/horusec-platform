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
