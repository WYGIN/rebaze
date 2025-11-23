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

package v1

import ocispec "github.com/opencontainers/image-spec/specs-go/v1"

// Descriptor mirrors the OCI Descriptor specification, as defined in
// opencontainers/image-spec/specs-go/v1/descriptor.go. It provides a
// consistent structure for referencing content-addressable resources.
type Descriptor ocispec.Descriptor
