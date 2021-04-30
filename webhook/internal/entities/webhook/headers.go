package webhook

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type HeaderType []Headers

type Headers struct {
	Key   string `json:"key" example:"x-authorization"`
	Value string `json:"value" example:"my-header-value"`
}

func (h HeaderType) Value() (driver.Value, error) {
	return json.Marshal(h)
}
func (h *HeaderType) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("[]byte assertion failed")
	}

	return json.Unmarshal(b, h)
}

func (h HeaderType) GetMapHeaders() map[string]string {
	headers := map[string]string{}
	for _, header := range h {
		headers[header.Key] = header.Value
	}
	return headers
}
