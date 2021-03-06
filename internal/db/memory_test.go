package db

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/110y/tododemo/internal/todo"
)

func TestMemoryDBPutTODO(t *testing.T) {
	t.Parallel()

	todo1 := &todo.TODO{
		ID:    "b8231e9e-0b83-4e45-b31f-5d3ba498ba60",
		Title: "brush the gopher",
	}

	tests := map[string]struct {
		todo     *todo.TODO
		expected map[string]*todo.TODO
	}{
		"put": {
			todo:     todo1,
			expected: map[string]*todo.TODO{todo1.ID: todo1},
		},
	}

	ctx := context.Background()
	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			d := &memoryDB{db: map[string]*todo.TODO{}}
			if err := d.PutTODO(ctx, test.todo); err != nil {
				t.Fatalf("failed to put a todo: %s", err.Error())
			}

			if diff := cmp.Diff(test.expected, d.db); diff != "" {
				t.Errorf("\n(-expected, +actual)\n%s", diff)
			}
		})
	}
}
