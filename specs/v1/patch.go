package v1

import (
	"time"

	"github.com/WYGIN/rebaze/specs"
)

// PatchOperation describes a single, atomic step within a patch. Each operation
// is clearly defined by its type, target, and associated values.
type PatchOperation struct {
	ID         string                      `json:"id"`
	Op         PatchOperationType `json:"op"`
	Target     string                      `json:"target"`
	TargetType TargetType         `json:"targetType"`
	Digest     string                      `json:"digest,omitempty"`
	OldDigest  string                      `json:"oldDigest,omitempty"`
	Value      any                         `json:"value,omitempty"`
	OldValue   any                         `json:"oldValue,omitempty"`
	Reason     string                      `json:"reason,omitempty"`
	// Annotations contains arbitrary metadata relating to the targeted content.
	Annotations map[string]string `json:"annotations,omitempty"`
	// --- Field for bubbled operations ---

	// If non-empty, this field contains a list of sub-operations, making this
	// a "bubbled" or grouped operation. The fields for atomic operations
	// above should be omitted in this case.
	Operations []PatchOperation `json:"operations,omitempty"`
}

// PatchSpec is the core structure that defines a patch artifact for an OCI image
// or other OCI artifact. It consolidates all components of a patch into a single,
// verifiable specification.
type PatchSpec struct {
	specs.Versioned
	// MediaType specifies the type of this document data structure e.g. `application/vnd.oci.image.index.v1+json`
	MediaType string `json:"mediaType,omitempty"`

	// ArtifactType specifies the IANA media type of artifact when the manifest is used for an artifact.
	ArtifactType string    `json:"artifactType,omitempty"`
	Created      time.Time `json:"created"`
	Author       string    `json:"author,omitempty"`
	Description  string    `json:"description,omitempty"`
	// Subject is an optional link from the image manifest to another manifest forming an association between the image manifest and the other manifest.
	Subject *Descriptor `json:"subject,omitempty"`
	// Data is an embedding of the targeted content. This is encoded as a base64
	// string when marshalled to JSON (automatically, by encoding/json). If
	// present, Data can be used directly to avoid fetching the targeted content.
	Data   []byte            `json:"data,omitempty"`
	Audit  []PatchAuditEvent `json:"audit,omitempty"`
	Users  []User            `json:"users"`
	Patchs []PatchOperation  `json:"patches"`
	// Annotations contains arbitrary metadata for the image index.
	Annotations map[string]string `json:"annotations,omitempty"`
}
