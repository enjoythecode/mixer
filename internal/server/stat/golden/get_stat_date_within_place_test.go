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
	"github.com/datacommonsorg/mixer/test"
	"github.com/google/go-cmp/cmp"
	"google.golang.org/protobuf/testing/protocmp"
)

func TestGetStatDateWithinPlace(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	_, filename, _, _ := runtime.Caller(0)
	goldenPath := path.Join(
		path.Dir(filename), "get_stat_date_within_place")

	testSuite := func(mixer pb.MixerClient, recon pb.ReconClient, latencyTest bool) {
		for _, c := range []struct {
			ancestorPlace  string
			childPlaceType string
			statVars       []string
			goldenFile     string
		}{
			{
				"geoId/06",
				"County",
				[]string{"Count_Person", "Median_Age_Person"},
				"CA_County.json",
			},
			{
				"country/USA",
				"State",
				[]string{"Count_Person", "Count_Person_Female"},
				"USA_State.json",
			},
			{
				"country/USA",
				"State",
				[]string{"Count_Person_FoodInsecure", "Mean_MealCost_Person_FoodSecure"},
				"memdb.json",
			},
		} {
			resp, err := mixer.GetStatDateWithinPlace(ctx, &pb.GetStatDateWithinPlaceRequest{
				AncestorPlace:  c.ancestorPlace,
				ChildPlaceType: c.childPlaceType,
				StatVars:       c.statVars,
			})
			if err != nil {
				t.Errorf("could not GetStatDateWithinPlace: %s", err)
				continue
			}

			if latencyTest {
				continue
			}

			if test.GenerateGolden {
				test.UpdateGolden(resp, goldenPath, c.goldenFile)
				continue
			}
			var expected pb.GetStatDateWithinPlaceResponse
			if err := test.ReadJSON(goldenPath, c.goldenFile, &expected); err != nil {
				t.Errorf("Can not Unmarshal golden file")
				continue
			}

			if diff := cmp.Diff(resp, &expected, protocmp.Transform()); diff != "" {
				t.Errorf("payload got diff: %v", diff)
				continue
			}
		}
	}

	if err := test.TestDriver(
		"GetStatDateWithinPlace",
		&test.TestOption{UseMemdb: true},
		testSuite,
	); err != nil {
		t.Errorf("TestDriver() = %s", err)
	}
}
