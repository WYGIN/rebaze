package bazer

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/google/go-containerregistry/pkg/authn"
	"github.com/google/go-containerregistry/pkg/name"
	v1 "github.com/google/go-containerregistry/pkg/v1"
	"github.com/google/go-containerregistry/pkg/v1/remote"
	"github.com/spf13/cobra"
)

var list = &cobra.Command{
	Use: "list",
	Args: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	Short:             "",
	Long:              ``,
	Example:           ``,
	ValidArgsFunction: inspect.ValidArgsFunction,
	Annotations:       map[string]string{},
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return errors.New("reference is required")
		}
	
		ref, err := parseReference(args[0])
		if err != nil {
			return err
		}
	
		desc, err := fetchRemoteDescriptor(ref)
		if err != nil {
			return err
		}
	
		manifest, err := resolveManifest(desc)
		if err != nil {
			return err
		}
	
		return printJSON(manifest)
	},
}

func parseReference(refStr string) (name.Reference, error) {
	return name.ParseReference(
		refStr,
		name.Insecure,
		name.WeakValidation,
	)
}

func fetchRemoteDescriptor(ref name.Reference) (*remote.Descriptor, error) {
	return remote.Get(
		ref,
		remote.WithAuthFromKeychain(authn.DefaultKeychain),
	)
}

func resolveManifest(desc *remote.Descriptor) (any, error) {
	if manifest, err := indexManifest(desc); err == nil {
		return manifest, nil
	}

	return imageManifest(desc)
}

func indexManifest(desc *remote.Descriptor) (*v1.IndexManifest, error) {
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

func imageManifest(desc *remote.Descriptor) (*v1.Manifest, error) {
	img, err := desc.Image()
	if err != nil {
		return nil, err
	}

	manifest, err := img.Manifest()
	if err != nil {
		return nil, err
	}

	if manifest == nil {
		return nil, errors.New("image manifest is nil")
	}

	return manifest, nil
}

func printJSON(v any) error {
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	return enc.Encode(v)
}
