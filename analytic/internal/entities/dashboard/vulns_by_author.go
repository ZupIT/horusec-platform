package dashboard

type VulnerabilitiesByAuthor struct {
	Author string `json:"author" gorm:"Column:author"`
	Vulnerability
}

func (v *VulnerabilitiesByAuthor) ToResponseByAuthor() ByAuthor {
	return ByAuthor{
		Author:       v.Author,
		BySeverities: v.ToResponseBySeverities(),
	}
}
