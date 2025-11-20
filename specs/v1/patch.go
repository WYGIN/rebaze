package v1

import (
	"time"

	"github.com/opencontainers/image-spec/specs-go/v1"
)

// Descriptor mirrors OCI Descriptor spec (from opencontainers/image-spec/specs-go/v1/descriptor.go).
type Descriptor = v1.Descriptor

// PatchOperationType enumerates supported patch operations.
type PatchOperationType string

const (
	PatchOpAdd      PatchOperationType = "add"
	PatchOpRemove   PatchOperationType = "remove"
	PatchOpReplace  PatchOperationType = "replace"
	PatchOpMove     PatchOperationType = "move"
	PatchOpAnnotate PatchOperationType = "annotate"
)

// PatchOperation describes a single patch step.
type PatchOperation struct {
	Op         PatchOperationType `json:"op"`                  // add, remove, replace, move, annotate
	Target     string             `json:"target"`              // JSON pointer path
	TargetType string             `json:"targetType"`          // "manifest", "config", or "index"
	Digest     string             `json:"digest,omitempty"`    // for add/replace (target content)
	OldDigest  string             `json:"oldDigest,omitempty"` // for remove/replace (previous content, mandatory in those cases)
	From       string             `json:"from,omitempty"`      // for move ops, JSON pointer source
	Value      any                `json:"value,omitempty"`     // for config value patching
	OldValue   any                `json:"oldValue,omitempty"`  // previous value for audit/rollback
	Reason     string             `json:"reason,omitempty"`    // description of why
	Metadata   map[string]string  `json:"metadata,omitempty"`  // arbitrary operation-level metadata
}

// PatchPolicy describes a validation policy (e.g. OPA/Rego).
type PatchPolicy struct {
	Type        string            `json:"type"`
	Rego        string            `json:"rego,omitempty"`
	Description string            `json:"description,omitempty"`
	Metadata    map[string]string `json:"metadata,omitempty"`
}

// PatchAuditEvent describes an audit log entry.
type PatchAuditEvent struct {
	Timestamp time.Time       `json:"timestamp"`
	User      string          `json:"user,omitempty"`
	Action    string          `json:"action"`
	Operation *PatchOperation `json:"operation,omitempty"`
	Details   map[string]any  `json:"details,omitempty"`
}

// PatchSignature describes a cryptographic patch signature.
type PatchSignature struct {
	Type      string `json:"type"`      // e.g. "notation", "cosign"
	Reference string `json:"reference"` // signature digest or OCI ref
}

// PatchRollback describes rollback instructions.
type PatchRollback struct {
	Instructions []PatchOperation `json:"instructions"`
}

// PatchSpec describes a patch artifact for an OCI image or artifact.
type PatchSpec struct {
	SchemaVersion int               `json:"schemaVersion"` // Always 1
	MediaType     string            `json:"mediaType"`     // application/vnd.oci.image.patch.v1+json
	ID            string            `json:"id"`            // Unique id for this patch
	Created       time.Time         `json:"created"`       // RFC3339 timestamp
	Author        string            `json:"author,omitempty"`
	Description   string            `json:"description,omitempty"`
	Subject       Descriptor        `json:"subject"`    // Target image/index/config
	Operations    []PatchOperation  `json:"operations"` // All patch operations, ordered
	Rollback      *PatchRollback    `json:"rollback,omitempty"`
	Policies      []PatchPolicy     `json:"policies,omitempty"`
	Audit         []PatchAuditEvent `json:"audit,omitempty"`
	Signatures    []PatchSignature  `json:"signatures,omitempty"`
	Metadata      map[string]string `json:"metadata,omitempty"` // Arbitrary K/V metadata for tools
}
