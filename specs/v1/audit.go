package v1

import "time"

// PatchAuditEvent records an entry in the audit log, capturing details
// of when, by whom, and what action was performed.
type PatchAuditEvent struct {
	Timestamp time.Time `json:"timestamp"`
	UserID    string    `json:"user"`
	Action    string    `json:"action"`
	PatchID   string    `json:"patch"`
	// URLs specifies a list of URLs from which this object MAY be downloaded
	URLs []string `json:"urls,omitempty"`
	// Data is an embedding of the targeted content. This is encoded as a base64
	// string when marshalled to JSON (automatically, by encoding/json). If
	// present, Data can be used directly to avoid fetching the targeted content.
	Data        []byte            `json:"data,omitempty"`
	Annotations map[string]string `json:"annotations,omitempty"`
}
