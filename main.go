package main

import (
	digest "github.com/opencontainers/go-digest"
	ocispecs "github.com/opencontainers/image-spec/specs-go/v1"
)

func main() {
	_ = digest.Digest("")
	_ = ocispecs.AnnotationAuthors
}
