package main

import (
	digest "github.com/opencontainers/go-digest"
	ocispecs "github.com/opencontainers/image-spec/specs-go/v1"
	"github.com/google/go-containerregistry/pkg/v1/remote"
)

func main() {
	_ = digest.Digest("")
	_ = ocispecs.AnnotationAuthors
	_ = remote.Descriptor{}
}
