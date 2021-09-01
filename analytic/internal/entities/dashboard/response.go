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

package dashboard

type Response struct {
	TotalAuthors                int            `json:"totalAuthors"`
	TotalRepositories           int            `json:"totalRepositories"`
	VulnerabilityBySeverity     *BySeverities  `json:"vulnerabilityBySeverity"`
	VulnerabilitiesByAuthor     []ByAuthor     `json:"vulnerabilitiesByAuthor"`
	VulnerabilitiesByRepository []ByRepository `json:"vulnerabilitiesByRepository"`
	VulnerabilitiesByLanguage   []ByLanguage   `json:"vulnerabilitiesByLanguage"`
	VulnerabilitiesByTime       []ByTime       `json:"vulnerabilityByTime"`
}

func (r *Response) SetTotalAuthors(totalAuthors int, err error) error {
	if err == nil {
		r.TotalAuthors = totalAuthors
	}

	return err
}

func (r *Response) SetTotalRepositories(totalRepositories int, err error) error {
	if err == nil {
		r.TotalRepositories = totalRepositories
	}

	return err
}

func (r *Response) SetChartBySeverity(vulnerability *Vulnerability, err error) error {
	if err == nil && vulnerability != nil {
		r.VulnerabilityBySeverity = vulnerability.ToResponseBySeverities()
	}

	return err
}

func (r *Response) SetChartByAuthor(vulns []*VulnerabilitiesByAuthor, err error) error {
	if err == nil {
		for index := range vulns {
			r.VulnerabilitiesByAuthor = append(r.VulnerabilitiesByAuthor, vulns[index].ToResponseByAuthor())
		}
	}

	return err
}

func (r *Response) SetChartByRepository(vulns []*VulnerabilitiesByRepository, err error) error {
	if err == nil {
		for index := range vulns {
			r.VulnerabilitiesByRepository = append(r.VulnerabilitiesByRepository, vulns[index].ToResponseByRepository())
		}
	}

	return err
}

func (r *Response) SetChartByLanguage(vulns []*VulnerabilitiesByLanguage, err error) error {
	if err == nil {
		for index := range vulns {
			r.VulnerabilitiesByLanguage = append(r.VulnerabilitiesByLanguage, vulns[index].ToResponseByLanguage())
		}
	}

	return err
}

func (r *Response) SetChartByTime(vulns []*VulnerabilitiesByTime, err error) error {
	if err == nil {
		for index := range vulns {
			r.VulnerabilitiesByTime = append(r.VulnerabilitiesByTime, vulns[index].ToResponseByTime())
		}
	}

	return err
}
