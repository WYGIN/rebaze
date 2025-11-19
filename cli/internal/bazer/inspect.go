package bazer

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/google/go-containerregistry/pkg/authn"
	"github.com/google/go-containerregistry/pkg/name"
	"github.com/google/go-containerregistry/pkg/v1"
	"github.com/google/go-containerregistry/pkg/v1/remote"
	"github.com/spf13/cobra"
)

var inspect = &cobra.Command{
	Use:               "inspect <reference>",
	Short:             "inspect manifest information for a container reference",
	Example:           "rebaze inspect ubuntu:latest",
	RunE:              runInspect,
}

func runInspect(cmd *cobra.Command, args []string) error {
	if err := validateArgs(args); err != nil {
		return err
	}

	reference, err := parseReference(args[0])
	if err != nil {
		return err
	}

	descriptor, err := fetchDescriptor(reference)
	if err != nil {
		return err
	}

	manifest, err := loadManifest(descriptor)
	if err != nil {
		return err
	}

	return writeJSON(manifest)
}

func validateArgs(args []string) error {
	if len(args) < 1 {
		return errors.New("reference argument is required")
	}
	return nil
}

func parseReference(ref string) (name.Reference, error) {
	return name.ParseReference(
		ref,
		name.Insecure,
		name.WeakValidation,
	)
}

func fetchDescriptor(ref name.Reference) (*remote.Descriptor, error) {
	return remote.Get(
		ref,
		remote.WithAuthFromKeychain(authn.DefaultKeychain),
	)
}

func loadManifest(desc *remote.Descriptor) (any, error) {
	indexManifest, err := tryLoadIndexManifest(desc)
	if err == nil {
		return indexManifest, nil
	}

	return loadImageManifest(desc)
}

func tryLoadIndexManifest(desc *remote.Descriptor) (*v1.IndexManifest, error) {
	index, err := desc.ImageIndex()
	if err != nil {
		return nil, err
	}

	manifest, err := index.IndexManifest()
	if err != nil {
		return nil, err
	}

	if manifest == nil {
		return nil, errors.New("index manifest is nil")
	}

	return manifest, nil
}

func loadImageManifest(desc *remote.Descriptor) (*v1.Manifest, error) {
	image, err := desc.Image()
	if err != nil {
		return nil, err
	}

	manifest, err := image.Manifest()
	if err != nil {
		return nil, err
	}

	if manifest == nil {
		return nil, errors.New("image manifest is nil")
	}

	return manifest, nil
}

func writeJSON(data any) error {
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(data)
}
