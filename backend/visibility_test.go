package backend

import (
	"fmt"
	"testing"
)

func dummyEntity(x, y int, id uint64) *Entity {
	e := Entity{}
	e.X = x
	e.Y = y
	e.ID = id
	return &e
}

func runVisibilityTest(entities []*Entity, myId uint64, viewDistance int, visibleIds []uint64) error {
	var source *Entity
	for _, e := range entities {
		if e.ID == myId {
			source = e
			break
		}
	}
	visible := getVisible(source, entities, viewDistance)

	extra := []uint64{}
	missing := []uint64{}
	for _, e := range visible {
		ok := false
		for _, id := range visibleIds {
			if id == e.ID {
				ok = true
			}
		}
		if !ok {
			extra = append(extra, e.ID)
		}
	}

	for _, id := range visibleIds {
		ok := false
		for _, e := range visible {
			if id == e.ID {
				ok = true
			}
		}
		if !ok {
			missing = append(missing, id)
		}
	}

	if len(extra) > 0 || len(missing) > 0 {
		return fmt.Errorf("shouldn't have seen %v, should have seen %v", extra, missing)
	}

	return nil
}

func TestVisibility(t *testing.T) {
	type testcase struct {
		name         string
		entities     []*Entity
		myId         uint64
		viewDistance int
		visibleIds   []uint64
	}

	tests := []testcase{
		{
			name: "Basic occlusion test in first octant, first column",
			entities: []*Entity{
				dummyEntity(0, 0, 0),
				dummyEntity(0, 1, 1),
				dummyEntity(0, 2, 2),
			},
			myId:         0,
			viewDistance: 5,
			visibleIds:   []uint64{1},
		},
		{
			name: "Two entities obscure one behind, at an angle",
			entities: []*Entity{
				dummyEntity(0, 0, 0),
				dummyEntity(0, 1, 1),
				dummyEntity(1, 1, 2),
				dummyEntity(1, 3, 3),
			},
			myId:         0,
			viewDistance: 5,
			visibleIds:   []uint64{1, 2},
		},
		{
			name: "Partially visible entities are visible",
			entities: []*Entity{
				dummyEntity(0, 0, 0),
				dummyEntity(0, 1, 1),
				dummyEntity(1, 2, 2),
			},
			myId:         0,
			viewDistance: 5,
			visibleIds:   []uint64{1, 2},
		},
		{
			name: "Entities at different distances can combine to occlude",
			entities: []*Entity{
				dummyEntity(0, 0, 0),
				dummyEntity(0, 2, 1),
				dummyEntity(1, 3, 2),
				dummyEntity(2, 5, 3),
			},
			myId:         0,
			viewDistance: 5,
			visibleIds:   []uint64{1, 2},
		},
		{
			name: "Can't see beyond viewDistance",
			entities: []*Entity{
				dummyEntity(0, 0, 0),
				dummyEntity(0, 6, 1),
			},
			myId:         0,
			viewDistance: 5,
			visibleIds:   []uint64{},
		},
		{
			name: "Can see in all octants",
			entities: []*Entity{
				dummyEntity(0, 0, 0),
				dummyEntity(1, 3, 1),
				dummyEntity(3, 1, 2),
				dummyEntity(3, -1, 3),
				dummyEntity(1, -3, 4),
				dummyEntity(-1, -3, 5),
				dummyEntity(-3, -1, 6),
				dummyEntity(-3, 1, 7),
				dummyEntity(-1, 3, 8),
			},
			myId:         0,
			viewDistance: 5,
			visibleIds:   []uint64{1, 2, 3, 4, 5, 6, 7, 8},
		},
		{
			name: "Basic occlusion test in all octants",
			entities: []*Entity{
				dummyEntity(0, 0, 0),
				dummyEntity(1, 3, 1),
				dummyEntity(3, 1, 2),
				dummyEntity(3, -1, 3),
				dummyEntity(1, -3, 4),
				dummyEntity(-1, -3, 5),
				dummyEntity(-3, -1, 6),
				dummyEntity(-3, 1, 7),
				dummyEntity(-1, 3, 8),

				// These block all others from being visible.
				dummyEntity(0, 1, 9),
				dummyEntity(1, 1, 10),
				dummyEntity(1, 0, 11),
				dummyEntity(1, -1, 12),
				dummyEntity(0, -1, 13),
				dummyEntity(-1, -1, 14),
				dummyEntity(-1, 0, 15),
				dummyEntity(-1, 1, 16),
			},
			myId:         0,
			viewDistance: 5,
			visibleIds:   []uint64{9, 10, 11, 12, 13, 14, 15, 16},
		},
	}

	for ix, test := range tests {
		err := runVisibilityTest(test.entities, test.myId, test.viewDistance, test.visibleIds)
		if err != nil {
			t.Errorf("Failed test %v. %v: %v", ix, test.name, err)
		}
	}
}
