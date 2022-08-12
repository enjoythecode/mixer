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

package propertyvalues

import (
	"container/heap"

	pb "github.com/datacommonsorg/mixer/internal/proto"
	"google.golang.org/protobuf/proto"
)

// Cache builder generates page size of 500.
// This limit applies to per type of a property, not the total response
var defaultLimit = 500

func buildDefaultCursorGroups(
	entities []string,
	properties []string,
	propType map[string]map[string][]string,
	n int,
) []*pb.CursorGroup {
	result := []*pb.CursorGroup{}
	for _, e := range entities {
		for _, p := range properties {
			for _, t := range propType[e][p] {
				cg := &pb.CursorGroup{
					Keys: []string{e, p, t},
				}
				for i := 0; i < n; i++ {
					cg.Cursors = append(cg.Cursors, &pb.Cursor{ImportGroup: int32(i)})
				}
				result = append(result, cg)
			}
		}
	}
	return result
}

var unmarshalFunc = func(jsonRaw []byte) (interface{}, error) {
	var p pb.PagedEntities
	err := proto.Unmarshal(jsonRaw, &p)
	return &p, err
}

// MergeTypedEntities merges entities by type into ordered flat list of entities.
func MergeTypedEntities(data map[string][]*pb.EntityInfo) []*pb.EntityInfo {
	res := []*pb.EntityInfo{}
	h := &entityHeap{}
	heap.Init(h)
	for t, items := range data {
		heap.Push(h, &heapElem{typ: t, pos: 0, data: items[0]})
	}
	for h.Len() > 0 {
		elem := heap.Pop(h).(*heapElem)
		e, t, pos := elem.data, elem.typ, elem.pos
		if e.GetTypes() == nil && t != "" {
			e.Types = []string{t}
		}
		res = append(res, e)
		if pos < int32(len(data[t])-1) {
			heap.Push(
				h,
				&heapElem{
					typ:  t,
					pos:  pos + 1,
					data: data[t][pos+1],
				})
		}
	}
	return res
}
