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

package golden

import (
	"context"
	"path"
	"runtime"
	"testing"

	pb "github.com/datacommonsorg/mixer/internal/proto"
	"github.com/datacommonsorg/mixer/internal/util"
	"github.com/datacommonsorg/mixer/test"
	"github.com/google/go-cmp/cmp"
	"google.golang.org/protobuf/testing/protocmp"
)

func TestBulkPropertyValuesOut(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	_, filename, _, _ := runtime.Caller(0)
	goldenPath := path.Join(path.Dir(filename), "bulk_property_values_out")

	testSuite := func(mixer pb.MixerClient, recon pb.ReconClient, latencyTest bool) {
		for _, c := range []struct {
			goldenFile string
			property   string
			entities   []string
			limit      int32
			token      string
		}{
			{
				"name.json",
				"name",
				[]string{"country/USA", "geoId/06", "dummy"},
				0,
				"",
			},
			{
				"containedIn.json",
				"containedInPlace",
				[]string{"country/USA", "geoId/06053", "geoId/06"},
				0,
				"",
			},
			{
				"geoOverlaps1.json",
				"geoOverlaps",
				[]string{"country/USA", "geoId/06053", "geoId/06"},
				5,
				"",
			},
			{
				"geoOverlaps2.json",
				"geoOverlaps",
				[]string{"country/USA", "geoId/06053", "geoId/06"},
				5,
				"H4sIAAAAAAAA/+LS4eJOzi/NKymq1A8NduTiTk/N9y9LLcpJLCgWYhBi4mAUYuJgEmLiYBZi4mDh0gMr8EzRNzAzMDVGV83CwSjBiqJek4sDpp6A0QAAAAD//wEAAP//omBbzIkAAAA=",
			},
		} {
			req := &pb.BulkPropertyValuesRequest{
				Property:  c.property,
				Entities:  c.entities,
				Direction: util.DirectionOut,
				Limit:     c.limit,
				NextToken: c.token,
			}
			resp, err := mixer.BulkPropertyValues(ctx, req)
			if err != nil {
				t.Errorf("could not run mixer.BulkPropertyValues/out: %s", err)
				continue
			}
			if latencyTest {
				continue
			}
			if test.GenerateGolden {
				test.UpdateProtoGolden(resp, goldenPath, c.goldenFile)
				continue
			}
			var expected pb.BulkPropertyValuesResponse
			if err := test.ReadJSON(goldenPath, c.goldenFile, &expected); err != nil {
				t.Errorf("Can not Unmarshal golden file %s: %v", c.goldenFile, err)
				continue
			}
			if diff := cmp.Diff(resp, &expected, protocmp.Transform()); diff != "" {
				t.Errorf("Golden got diff: %v", diff)
				continue
			}
		}
	}
	if err := test.TestDriver(
		"BulkPropertyValues/out", &test.TestOption{}, testSuite); err != nil {
		t.Errorf("TestDriver() = %s", err)
	}
}
