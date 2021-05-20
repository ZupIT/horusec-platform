package dashboard

type VulnerabilitiesByRepository struct {
	RepositoryName string `json:"repository_name" gorm:"Column:repository_name"`
	Vulnerability
}

func (v *VulnerabilitiesByRepository) ToResponseByRepository() ByRepository {
	return ByRepository{
		RepositoryName: v.RepositoryName,
		BySeverities:   v.ToResponseBySeverities(),
	}
}
