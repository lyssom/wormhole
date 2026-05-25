package index

import (
	"testing"
)

func TestHNSWAddAndSearch(t *testing.T) {
	dim := 4
	h := NewHNSW(dim, 16, 200)

	// Add some vectors
	vecs := [][]float32{
		{1.0, 0.0, 0.0, 0.0},   // ID 0
		{0.9, 0.1, 0.0, 0.0},   // ID 1 - close to 0
		{0.0, 1.0, 0.0, 0.0},   // ID 2 - far from 0
		{0.0, 0.0, 1.0, 0.0},   // ID 3 - far from 0
		{0.0, 0.0, 0.0, 1.0},   // ID 4 - far from 0
	}

	for i, vec := range vecs {
		if err := h.Add(i, vec); err != nil {
			t.Fatalf("Add(%d) failed: %v", i, err)
		}
	}

	if h.Size() != 5 {
		t.Errorf("Size: got %d, want 5", h.Size())
	}

	// Search for nearest to [1, 0, 0, 0] - should be ID 0, then 1
	results, err := h.Search([]float32{1.0, 0.0, 0.0, 0.0}, 2, 10)
	if err != nil {
		t.Fatalf("Search failed: %v", err)
	}

	if len(results) != 2 {
		t.Fatalf("expected 2 results, got %d", len(results))
	}

	// First result should be ID 0 (exact match)
	if results[0].ID != 0 {
		t.Errorf("first result: got ID %d, want 0", results[0].ID)
	}

	// Second result should be ID 1 (closest to first)
	if results[1].ID != 1 {
		t.Errorf("second result: got ID %d, want 1", results[1].ID)
	}
}

func TestHNSWSearchEmpty(t *testing.T) {
	h := NewHNSW(4, 16, 200)

	// Search on empty index
	results, err := h.Search([]float32{1.0, 0.0, 0.0, 0.0}, 2, 10)
	if err != nil {
		t.Fatalf("Search failed: %v", err)
	}

	if results != nil {
		t.Errorf("expected nil results for empty index, got %v", results)
	}
}

func TestHNSWDimensionMismatch(t *testing.T) {
	h := NewHNSW(4, 16, 200)

	err := h.Add(0, []float32{1.0, 0.0})
	if err == nil {
		t.Errorf("expected error for dimension mismatch")
	}

	_, err = h.Search([]float32{1.0, 0.0}, 2, 10)
	if err == nil {
		t.Errorf("expected error for dimension mismatch in search")
	}
}

func TestHNSWDistance(t *testing.T) {
	h := NewHNSW(3, 16, 200)

	// Distance from [1, 0, 0] to [0, 1, 0] should be sqrt(2)
	v1 := []float32{1.0, 0.0, 0.0}
	v2 := []float32{0.0, 1.0, 0.0}
	d := h.distance(v1, v2)
	expected := float32(1.41421356) // sqrt(2)
	if d < expected-0.01 || d > expected+0.01 {
		t.Errorf("distance: got %f, want ~%f", d, expected)
	}
}

func TestHNSWSerializeDeserialize(t *testing.T) {
	dim := 4
	h := NewHNSW(dim, 16, 200)

	// Add some vectors
	vecs := [][]float32{
		{1.0, 0.0, 0.0, 0.0},   // ID 0
		{0.9, 0.1, 0.0, 0.0},   // ID 1 - close to 0
		{0.0, 1.0, 0.0, 0.0},   // ID 2
	}
	for i, vec := range vecs {
		if err := h.Add(i, vec); err != nil {
			t.Fatalf("Add(%d) failed: %v", i, err)
		}
	}

	// Serialize
	data, err := h.Serialize()
	if err != nil {
		t.Fatalf("Serialize failed: %v", err)
	}

	// Deserialize
	h2, err := DeserializeHNSW(data)
	if err != nil {
		t.Fatalf("Deserialize failed: %v", err)
	}

	// Verify
	if h2.Size() != h.Size() {
		t.Errorf("Size after deserialize: got %d, want %d", h2.Size(), h.Size())
	}

	if h2.dim != dim {
		t.Errorf("dim after deserialize: got %d, want %d", h2.dim, dim)
	}

	// Search should return same results
	results1, _ := h.Search([]float32{1.0, 0.0, 0.0, 0.0}, 2, 10)
	results2, _ := h2.Search([]float32{1.0, 0.0, 0.0, 0.0}, 2, 10)

	if len(results1) != len(results2) {
		t.Errorf("result count: got %d, want %d", len(results2), len(results1))
	}
	for i := range results1 {
		if results1[i].ID != results2[i].ID {
			t.Errorf("result[%d].ID: got %d, want %d", i, results2[i].ID, results1[i].ID)
		}
	}
}
