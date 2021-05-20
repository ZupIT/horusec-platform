package authentication

import "github.com/lib/pq"

type AuthzGroups struct {
	AuthzMember     pq.StringArray `json:"authzMember" gorm:"type:text[]"`
	AuthzAdmin      pq.StringArray `json:"authzAdmin" gorm:"type:text[]"`
	AuthzSupervisor pq.StringArray `json:"authzSupervisor" gorm:"type:text[]"`
}
