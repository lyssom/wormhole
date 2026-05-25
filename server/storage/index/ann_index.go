package index

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math"
	"math/rand"
	"sync"
)

// HNSWIndex implements Hierarchical Navigable Small World (HNSW) algorithm
// for approximate nearest neighbor search on vector data.
type HNSWIndex struct {
	dim          int   // vector dimension
	maxLayers    int   // maximum number of layers (M × log(n))
	efConstruction int  // ef parameter during construction
	mu            sync.RWMutex
	layers        []map[int][]int // layer -> node_id -> [neighbor_ids]
	vectors       [][]float32      // node_id -> vector
	entryPoint    int              // entry point for search
	entryLayer    int              // layer of entry point
	nodeCount     int              // total number of nodes
}

// SearchResult represents a search result with distance.
type SearchResult struct {
	ID       int
	Distance float32
}

// NewHNSW creates a new HNSW index with the given parameters.
// dim: vector dimension
// maxLayers: maximum number of layers (typically M × log(n))
// efConstruction: ef parameter during construction (higher = better quality, slower)
func NewHNSW(dim, maxLayers, efConstruction int) *HNSWIndex {
	return &HNSWIndex{
		dim:            dim,
		maxLayers:      maxLayers,
		efConstruction: efConstruction,
		layers:         make([]map[int][]int, maxLayers),
		vectors:        make([][]float32, 0),
		entryPoint:    -1,
		entryLayer:    0,
	}
}

// Add adds a vector to the index with the given ID.
func (h *HNSWIndex) Add(id int, vector []float32) error {
	h.mu.Lock()
	defer h.mu.Unlock()

	if len(vector) != h.dim {
		return fmt.Errorf("Add: vector dimension %d != index dimension %d", len(vector), h.dim)
	}

	// Expand vectors slice if needed
	if id >= len(h.vectors) {
		// Fill gaps with nil
		for i := len(h.vectors); i <= id; i++ {
			h.vectors = append(h.vectors, nil)
		}
	}
	h.vectors[id] = vector

	// Determine number of layers for this node (random, exponential distribution)
	numLayers := h.randomLayer()
	if numLayers >= h.maxLayers {
		numLayers = h.maxLayers - 1
	}

	// Search for nearest neighbors in each layer
	for layer := 0; layer <= numLayers; layer++ {
		if h.layers[layer] == nil {
			h.layers[layer] = make(map[int][]int)
		}

		// Find nearest existing nodes in this layer
		neighbors := h.searchLayer(vector, layer, h.efConstruction)

		// Connect to neighbors
		for _, neighborID := range neighbors {
			h.layers[layer][neighborID] = append(h.layers[layer][neighborID], id)
			h.layers[layer][id] = append(h.layers[layer][id], neighborID)
		}
	}

	// Update entry point if this is the first node or it has higher layer
	if h.entryPoint == -1 || numLayers > h.entryLayer {
		h.entryPoint = id
		h.entryLayer = numLayers
	}

	h.nodeCount++
	return nil
}

// randomLayer generates a random layer number using exponential distribution.
// Higher layers are less likely.
func (h *HNSWIndex) randomLayer() int {
	r := rand.Float64()
	l := int(-math.Log(r))
	if l == 0 {
		l = 1
	}
	return l
}

// searchLayer searches for k nearest neighbors in a specific layer.
func (h *HNSWIndex) searchLayer(query []float32, layer, ef int) []int {
	if h.entryPoint == -1 {
		return nil
	}

	// Initialize priority queue (as slice with manual ordering)
	type nodeDist struct {
		id       int
		distance float32
	}

	// Start from entry point
	current := h.entryPoint
	visited := map[int]bool{current: true}
	candidates := []nodeDist{{current, h.distance(h.vectors[current], query)}}
	queue := []nodeDist{{current, h.distance(h.vectors[current], query)}}

	for len(queue) > 0 {
		// Find node with maximum distance in queue
		var maxIdx int
		var maxDist float32 = -1
		for i, nd := range queue {
			if nd.distance > maxDist {
				maxDist = nd.distance
				maxIdx = i
			}
		}

		// Remove from queue
		nd := queue[maxIdx]
		queue = append(queue[:maxIdx], queue[maxIdx+1:]...)

		// Check if we can stop
		if len(candidates) >= ef && nd.distance > candidates[0].distance {
			break
		}

		// Get neighbors in this layer
		neighbors := h.layers[layer][nd.id]

		for _, neighborID := range neighbors {
			if visited[neighborID] {
				continue
			}
			visited[neighborID] = true

			neighborDist := h.distance(h.vectors[neighborID], query)

			// Add to candidates if within ef
			if len(candidates) < ef || neighborDist < candidates[0].distance {
				queue = append(queue, nodeDist{neighborID, neighborDist})

				// Maintain candidates as min-heap (by distance)
				if neighborDist < candidates[0].distance {
					candidates = append([]nodeDist{{neighborID, neighborDist}}, candidates...)
					if len(candidates) > ef {
						candidates = candidates[:ef]
					}
				} else {
					candidates = append(candidates, nodeDist{neighborID, neighborDist})
				}
			}
		}
	}

	// Sort candidates by distance
	for i := 0; i < len(candidates)-1; i++ {
		for j := i + 1; j < len(candidates); j++ {
			if candidates[j].distance < candidates[i].distance {
				candidates[i], candidates[j] = candidates[j], candidates[i]
			}
		}
	}

	result := make([]int, 0, ef)
	for i := 0; i < ef && i < len(candidates); i++ {
		result = append(result, candidates[i].id)
	}
	return result
}

// Search searches for k nearest neighbors to the query vector.
// ef: search parameter (higher = better accuracy, slower)
func (h *HNSWIndex) Search(query []float32, k, ef int) ([]SearchResult, error) {
	h.mu.RLock()
	defer h.mu.RUnlock()

	if len(query) != h.dim {
		return nil, fmt.Errorf("Search: query dimension %d != index dimension %d", len(query), h.dim)
	}

	if h.entryPoint == -1 {
		return nil, nil
	}

	// Search from top layer down to layer 0
	current := h.entryPoint
	currentDist := h.distance(h.vectors[current], query)

	for layer := h.entryLayer; layer > 0; layer-- {
		changed := true
		for changed {
			changed = false
			neighbors := h.layers[layer][current]
			for _, neighborID := range neighbors {
				d := h.distance(h.vectors[neighborID], query)
				if d < currentDist {
					currentDist = d
					current = neighborID
					changed = true
				}
			}
		}
	}

	// Layer 0 search with ef
	candidates := h.searchLayer(query, 0, ef)

	results := make([]SearchResult, 0, k)
	for _, id := range candidates {
		if h.vectors[id] == nil {
			continue
		}
		results = append(results, SearchResult{
			ID:       id,
			Distance: h.distance(h.vectors[id], query),
		})
	}

	// Sort by distance and take top k
	for i := 0; i < len(results)-1; i++ {
		for j := i + 1; j < len(results); j++ {
			if results[j].Distance < results[i].Distance {
				results[i], results[j] = results[j], results[i]
			}
		}
	}

	if len(results) > k {
		results = results[:k]
	}

	return results, nil
}

// distance computes Euclidean distance between two vectors.
func (h *HNSWIndex) distance(v1, v2 []float32) float32 {
	var sum float32
	for i := 0; i < h.dim; i++ {
		d := v1[i] - v2[i]
		sum += d * d
	}
	return float32(math.Sqrt(float64(sum)))
}

// Size returns the number of vectors in the index.
func (h *HNSWIndex) Size() int {
	h.mu.RLock()
	defer h.mu.RUnlock()
	return h.nodeCount
}

const hnswMagic = "HNSW"
const hnswVersion = 1

// Serialize serializes the HNSW index to binary format.
func (h *HNSWIndex) Serialize() ([]byte, error) {
	h.mu.RLock()
	defer h.mu.RUnlock()

	// Calculate size
	size := 4 + 4 + 4 + 4 + 4 + 4 + 4 + 4 // header
	for layerNum, layer := range h.layers {
		size += 4 // layer number
		size += 4 // number of nodes
		for nodeID, neighbors := range layer {
			size += 4                         // node_id
			size += 4                         // neighbor count
			size += 4 * len(neighbors)        // neighbors
			_ = nodeID                        // avoid unused
			_ = layerNum                      // avoid unused
		}
	}
	size += 4 // sentinel (0xFFFFFFFF) to mark end of layers
	size += 4 // nodeCount
	for id, vec := range h.vectors {
		if vec != nil {
			size += 4 + 4 + 4*len(vec) // id + dim + data
			_ = id
		}
	}

	buf := make([]byte, 0, size)

	// Write header
	buf = append(buf, hnswMagic...)
	buf = binary.LittleEndian.AppendUint32(buf, hnswVersion)
	buf = binary.LittleEndian.AppendUint32(buf, uint32(h.dim))
	buf = binary.LittleEndian.AppendUint32(buf, uint32(h.maxLayers))
	buf = binary.LittleEndian.AppendUint32(buf, uint32(h.efConstruction))
	buf = binary.LittleEndian.AppendUint32(buf, uint32(h.entryPoint))
	buf = binary.LittleEndian.AppendUint32(buf, uint32(h.entryLayer))
	buf = binary.LittleEndian.AppendUint32(buf, uint32(h.nodeCount))

	// Write layers
	for layerNum, layer := range h.layers {
		if len(layer) == 0 {
			continue
		}
		buf = binary.LittleEndian.AppendUint32(buf, uint32(layerNum))
		buf = binary.LittleEndian.AppendUint32(buf, uint32(len(layer)))
		for nodeID, neighbors := range layer {
			buf = binary.LittleEndian.AppendUint32(buf, uint32(nodeID))
			buf = binary.LittleEndian.AppendUint32(buf, uint32(len(neighbors)))
			for _, neighbor := range neighbors {
				buf = binary.LittleEndian.AppendUint32(buf, uint32(neighbor))
			}
		}
	}

	// Write sentinel to mark end of layers
	buf = binary.LittleEndian.AppendUint32(buf, 0xFFFFFFFF)

	// Write vectors
	vecCount := 0
	for _, vec := range h.vectors {
		if vec != nil {
			vecCount++
		}
	}
	buf = binary.LittleEndian.AppendUint32(buf, uint32(vecCount))
	for id, vec := range h.vectors {
		if vec == nil {
			continue
		}
		buf = binary.LittleEndian.AppendUint32(buf, uint32(id))
		buf = binary.LittleEndian.AppendUint32(buf, uint32(len(vec)))
		for _, f := range vec {
			buf = binary.LittleEndian.AppendUint32(buf, math.Float32bits(f))
		}
	}

	return buf, nil
}

// DeserializeHNSW deserializes an HNSW index from binary data.
func DeserializeHNSW(data []byte) (*HNSWIndex, error) {
	r := bytes.NewReader(data)

	// Read magic
	magic := make([]byte, 4)
	if _, err := r.Read(magic); err != nil {
		return nil, err
	}
	if string(magic) != hnswMagic {
		return nil, fmt.Errorf("invalid magic: %q", string(magic))
	}

	// Read version
	var version uint32
	if err := binary.Read(r, binary.LittleEndian, &version); err != nil {
		return nil, err
	}
	if version != hnswVersion {
		return nil, fmt.Errorf("unsupported version: %d", version)
	}

	// Read header
	var dim, maxLayers, efConstruction, entryPoint, entryLayer, nodeCount uint32
	if err := binary.Read(r, binary.LittleEndian, &dim); err != nil {
		return nil, err
	}
	if err := binary.Read(r, binary.LittleEndian, &maxLayers); err != nil {
		return nil, err
	}
	if err := binary.Read(r, binary.LittleEndian, &efConstruction); err != nil {
		return nil, err
	}
	if err := binary.Read(r, binary.LittleEndian, &entryPoint); err != nil {
		return nil, err
	}
	if err := binary.Read(r, binary.LittleEndian, &entryLayer); err != nil {
		return nil, err
	}
	if err := binary.Read(r, binary.LittleEndian, &nodeCount); err != nil {
		return nil, err
	}

	h := &HNSWIndex{
		dim:            int(dim),
		maxLayers:      int(maxLayers),
		efConstruction: int(efConstruction),
		entryPoint:     int(entryPoint),
		entryLayer:     int(entryLayer),
		nodeCount:      int(nodeCount),
		layers:         make([]map[int][]int, maxLayers),
		vectors:        make([][]float32, 0),
	}

	// Read layers until we hit the sentinel (0xFFFFFFFF)
	for {
		var layerNum uint32
		if err := binary.Read(r, binary.LittleEndian, &layerNum); err != nil {
			return nil, err
		}
		if layerNum == 0xFFFFFFFF {
			// Sentinel found, next is vector count
			break
		}

		if int(layerNum) >= h.maxLayers {
			return nil, fmt.Errorf("invalid layer number: %d >= maxLayers %d", layerNum, h.maxLayers)
		}

		var nodeCountInLayer uint32
		if err := binary.Read(r, binary.LittleEndian, &nodeCountInLayer); err != nil {
			return nil, err
		}

		if h.layers[layerNum] == nil {
			h.layers[layerNum] = make(map[int][]int)
		}

		for i := uint32(0); i < nodeCountInLayer; i++ {
			var nodeID, neighborCount uint32
			if err := binary.Read(r, binary.LittleEndian, &nodeID); err != nil {
				return nil, err
			}
			if err := binary.Read(r, binary.LittleEndian, &neighborCount); err != nil {
				return nil, err
			}
			neighbors := make([]int, neighborCount)
			for j := uint32(0); j < neighborCount; j++ {
				var neighbor uint32
				if err := binary.Read(r, binary.LittleEndian, &neighbor); err != nil {
					return nil, err
				}
				neighbors[j] = int(neighbor)
			}
			h.layers[layerNum][int(nodeID)] = neighbors
		}
	}

	// Read vector count
	var vecCount uint32
	if err := binary.Read(r, binary.LittleEndian, &vecCount); err != nil {
		return nil, err
	}

	// Read vectors
	for i := uint32(0); i < vecCount; i++ {
		var id, vecDim uint32
		if err := binary.Read(r, binary.LittleEndian, &id); err != nil {
			return nil, err
		}
		if err := binary.Read(r, binary.LittleEndian, &vecDim); err != nil {
			return nil, err
		}
		vec := make([]float32, vecDim)
		for j := uint32(0); j < vecDim; j++ {
			var bits uint32
			if err := binary.Read(r, binary.LittleEndian, &bits); err != nil {
				return nil, err
			}
			vec[j] = math.Float32frombits(bits)
		}
		// Expand vectors slice if needed
		for len(h.vectors) <= int(id) {
			h.vectors = append(h.vectors, nil)
		}
		h.vectors[id] = vec
	}

	return h, nil
}
