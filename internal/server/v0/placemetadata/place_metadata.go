// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package placemetadata

import (
	"context"

	pb "github.com/datacommonsorg/mixer/internal/proto"
	"github.com/datacommonsorg/mixer/internal/server/place"
	"github.com/datacommonsorg/mixer/internal/store"
	"github.com/datacommonsorg/mixer/internal/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GetPlaceMetadata implements API for Mixer.GetPlaceMetadata.
func GetPlaceMetadata(
	ctx context.Context,
	in *pb.GetPlaceMetadataRequest,
	store *store.Store,
) (*pb.GetPlaceMetadataResponse, error) {
	places := in.GetPlaces()
	if len(places) == 0 {
		return nil, status.Error(codes.InvalidArgument, "Missing required arguments: places")
	}
	if !util.CheckValidDCIDs(places) {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid places")
	}

	res, err := place.GetPlaceMetadataHelper(ctx, places, store)
	if err != nil {
		return nil, err
	}

	return &pb.GetPlaceMetadataResponse{Data: res}, nil
}
