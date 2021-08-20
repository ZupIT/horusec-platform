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

package repository

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPaginatedContent(t *testing.T) {
	t.Run("Should return search and check if not empty", func(t *testing.T) {
		p := (&PaginatedContent{}).SetSearch("abc")
		assert.Equal(t, "%abc%", p.GetSearch())
	})
	t.Run("Should return search and check is was empty", func(t *testing.T) {
		p := &PaginatedContent{}
		assert.Equal(t, "%%", p.GetSearch())
	})
	t.Run("Should return offset correctly 0 with page 0", func(t *testing.T) {
		p := (&PaginatedContent{}).SetSearch("").SetPage("0").SetSize("10").SetEnable(true)
		assert.Equal(t, 0, p.GetOffset())
	})
	t.Run("Should return offset correctly 0 with page 1", func(t *testing.T) {
		p := (&PaginatedContent{}).SetSearch("").SetPage("1").SetSize("10").SetEnable(true)
		assert.Equal(t, 0, p.GetOffset())
	})
	t.Run("Should return offset correctly 15", func(t *testing.T) {
		p := (&PaginatedContent{}).SetSearch("").SetPage("2").SetSize("10").SetEnable(true)
		assert.Equal(t, 15, p.GetOffset())
	})
	t.Run("Should return offset correctly 30", func(t *testing.T) {
		p := (&PaginatedContent{}).SetSearch("").SetPage("3").SetSize("10").SetEnable(true)
		assert.Equal(t, 30, p.GetOffset())
	})
	t.Run("Should return offset correctly 50 with size 50", func(t *testing.T) {
		p := (&PaginatedContent{}).SetSearch("").SetPage("2").SetSize("50").SetEnable(true)
		assert.Equal(t, 50, p.GetOffset())
	})
	t.Run("Should return offset empty when return error on parse size", func(t *testing.T) {
		p := (&PaginatedContent{}).SetSearch("").SetSize("abc")
		assert.Equal(t, 0, p.GetOffset())
	})
	t.Run("Should return offset empty when return error on parse page", func(t *testing.T) {
		p := (&PaginatedContent{}).SetSearch("").SetSize("40").SetPage("abc")
		assert.Equal(t, 0, p.GetOffset())
	})
}
