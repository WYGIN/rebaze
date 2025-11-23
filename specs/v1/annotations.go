// Copyright 2024 The Bazer Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may not obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package v1 defines annotations used in OCI images for patch metadata.
// These annotations provide a standardized way to link an image to its
// associated patch image, enabling tools to discover and apply patches.
package v1

const (
	// AnnotationPatchImageDigest is the annotation key for the digest of the
	// image's patch image. This allows for content-addressable lookup of the patch.
	// The value of this annotation should be the digest of the patch image manifest.
	AnnotationPatchImageDigest = "io.bazer.image.patch/digest"

	// AnnotationPatchImageName is the annotation key for the image reference
	// (e.g., "example.com/repo/app:patch-v1") of the image's patch image.
	// This provides a human-readable and location-addressable way to find the patch.
	AnnotationPatchImageName = "io.bazer.image.patch/name"
)

const MediaTypePatchImageV1 = "application/vnd.bazer.image.patch.v1+json"
