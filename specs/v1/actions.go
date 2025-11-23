package v1


// PatchOperationType enumerates the supported types of patch operations.
// These constants ensure type safety and prevent the use of invalid "magic strings".
type PatchOperationType string

const (
	// PatchOpAppend represents an operation to append a new value.
	PatchOpAppend PatchOperationType = "append"
	// PatchOpPrepend represents an operation to append a new value.
	PatchOpPrepend PatchOperationType = "prepend"
	// PatchOpRemove represents an operation to remove an existing value.
	PatchOpRemove PatchOperationType = "remove"
	// PatchOpReplace represents an operation to replace an existing value.
	PatchOpReplace PatchOperationType = "replace"
)

// TargetType defines the type of the target for a patch operation.
type TargetType string

const (
	// TargetTypeManifest indicates that the patch operation targets a manifest.
	TargetTypeManifest TargetType = "manifest"
	// TargetTypeConfig indicates that the patch operation targets a configuration object.
	TargetTypeConfig TargetType = "config"
	// TargetTypeIndex indicates that the patch operation targets an index.
	TargetTypeIndex TargetType = "index"
)
