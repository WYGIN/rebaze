package v1

import (
	"encoding/json"

	digest "github.com/opencontainers/go-digest"
)

var _ layerStep = (*LayerAppend)(nil)
var _ layerStep = (*LayerPrepend)(nil)
var _ layerStep = (*LayerRemove)(nil)
var _ layerStep = (*LayerReplace)(nil)

type layerOp struct {
	Operation PatchOperationType `json:"op"`
	data      layerStep          `json:"-"`
}

type layerStep interface {
	patchOp() PatchOperationType
}

type LayerAppend struct {
	Reference   digest.Digest     `json:"reference"`
	Anchor      digest.Digest     `json:"anchor"`
	Annotations map[string]string `json:"annotations,omitempty"`
}

func (l *LayerAppend) patchOp() PatchOperationType {
	return PatchOpAppend
}

type LayerPrepend struct {
	Reference   digest.Digest     `json:"reference"`
	Anchor      digest.Digest     `json:"anchor"`
	Annotations map[string]string `json:"annotations,omitempty"`
}

func (l *LayerPrepend) patchOp() PatchOperationType {
	return PatchOpPrepend
}

type LayerRemove struct {
	Reference   digest.Digest     `json:"reference"`
	Annotations map[string]string `json:"annotations,omitempty"`
}

func (l *LayerRemove) patchOp() PatchOperationType {
	return PatchOpRemove
}

type LayerReplace struct {
	Reference   digest.Digest     `json:"reference"`
	Anchor      digest.Digest     `json:"anchor"`
	Annotations map[string]string `json:"annotations,omitempty"`
}

func (l *LayerReplace) patchOp() PatchOperationType {
	return PatchOpReplace
}

type LayerAction struct {
	ID      string          `json:"id"`
	Kind    TargetType      `json:"kind"`
	Action  layerStep       `json:"action"`
	Actions LayerSubActions `json:"actions"`
	Reason  string          `json:"reason"`
	// Annotations contains arbitrary metadata relating to the targeted content.
	Annotations map[string]string `json:"annotations,omitempty"`
}

func (l *LayerAction) UnmarshalJSON(data []byte) error {
	var layerAction *struct {
		LayerAction `json:",inline"`
		layerOp     `json:"action"`
	}

	if err := json.Unmarshal(data, &layerAction); err != nil {
		return err
	}

	l = &layerAction.LayerAction
	l.Action = layerAction.data
	return nil
}

type LayerSubActions struct {
	Actions []layerStep `json:"actions"`
}

func (l *LayerSubActions) UnmarshalJSON(data []byte) error {
	var layerOps struct {
		Actions []layerOp `json:"actions"`
	}

	if err := json.Unmarshal(data, &layerOps); err != nil {
		return err
	}

	l.Actions = make([]layerStep, len(layerOps.Actions))
	for i, v := range layerOps.Actions {
		l.Actions[i] = v.data
	}

	return nil
}

func (l *layerOp) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, l); err != nil {
		return err
	}

	switch l.Operation {
	case PatchOpAppend:
		var layer *LayerAppend
		if err := json.Unmarshal(data, layer); err != nil {
			return err
		}
		l.data = layer
	case PatchOpPrepend:
		var layer *LayerPrepend
		if err := json.Unmarshal(data, layer); err != nil {
			return err
		}
		l.data = layer
	case PatchOpRemove:
		var layer *LayerRemove
		if err := json.Unmarshal(data, layer); err != nil {
			return err
		}
		l.data = layer
	case PatchOpReplace:
		var layer *LayerReplace
		if err := json.Unmarshal(data, layer); err != nil {
			return err
		}
		l.data = layer
	}

	return nil
}
