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
	"fmt"
	"strconv"

	"github.com/ZupIT/horusec-devkit/pkg/utils/logger"
)

type PaginatedContent struct {
	// ToRemove: Paginated method is optional for now because is not generate breaking changes
	Enable bool
	Page   int
	Size   int
	Search string
}

const (
	defaultSizeRepositories = 15
)

func (p *PaginatedContent) SetSize(sizeString string) *PaginatedContent {
	sizeNumber, err := strconv.Atoi(sizeString)
	if err != nil {
		logger.LogWarn("{WARN} Can not get size from query string: ", err)
		return p
	}
	if sizeNumber < defaultSizeRepositories {
		sizeNumber = defaultSizeRepositories
	}
	p.Size = sizeNumber
	return p
}

func (p *PaginatedContent) SetSearch(search string) *PaginatedContent {
	p.Search = search
	return p
}

func (p *PaginatedContent) SetEnable(enable bool) *PaginatedContent {
	p.Enable = enable
	return p
}

func (p *PaginatedContent) SetPage(pageString string) *PaginatedContent {
	pageNumber, err := strconv.Atoi(pageString)
	if err != nil {
		logger.LogWarn("{WARN} Can not get page from query string: ", err)
		return p
	}
	p.Page = pageNumber
	return p
}

func (p *PaginatedContent) GetSearch() string {
	return fmt.Sprintf("%s%s%s", "%", p.Search, "%")
}

func (p *PaginatedContent) GetOffset() int {
	if p.Page > 1 {
		return p.Size * (p.Page - 1)
	}
	return 0
}
