package authentication

type AuthzGroups struct {
	AuthzMember     []string `json:"authzMember"`
	AuthzAdmin      []string `json:"authzAdmin"`
	AuthzSupervisor []string `json:"authzSupervisor"`
}
