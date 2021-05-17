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
	if err == nil {
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
