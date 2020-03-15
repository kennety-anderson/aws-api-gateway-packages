package body

import (
	"bytes"
	"encoding/json"
)

// Create method generate new body to response api gateway
func Create(data map[string]interface{}) string {
	var buf bytes.Buffer

	body, _ := json.Marshal(data)

	json.HTMLEscape(&buf, body)

	return buf.String()
}
