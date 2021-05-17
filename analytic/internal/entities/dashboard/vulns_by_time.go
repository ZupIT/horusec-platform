package dashboard

type VulnerabilitiesByTime struct {
	Vulnerability
}

func (v *VulnerabilitiesByTime) ToResponseByTime() ByTime {
	return ByTime{
		Time:         v.CreatedAt,
		BySeverities: v.ToResponseBySeverities(),
	}
}
