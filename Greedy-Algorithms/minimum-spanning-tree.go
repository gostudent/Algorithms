package main
import "fmt"

func (gr *Graph) Mst() (mst []Edge) {
	var edgeToAdd, groupID uint64
	mst = []Edge{}

	// Using union-find algorithm to detect cycles
	sort.Sort(byWeight(gr.RawEdges))
	vertexByGroup := make(map[uint64][]uint64)
	vertexGroups := make(map[uint64]uint64)
	connect := make([]uint64, 2)
	lastUsedGroup := uint64(0)
queueLoop:
	for _, e := range gr.RawEdges {
		addToExistingGroup := false
		lG, lIn := vertexGroups[e.From]
		rG, rIn := vertexGroups[e.To]
		switch {
		case lIn && rIn:
			// We have a vertex :'(
			if lG == rG {
				continue queueLoop
			}

			// We will connect two groups
			if len(vertexByGroup[lG]) < len(vertexByGroup[rG]) {
				connect[0] = rG
				connect[1] = lG
			} else {
				connect[0] = lG
				connect[1] = rG
			}

			for _, v := range vertexByGroup[connect[1]] {
				vertexByGroup[connect[0]] = append(vertexByGroup[connect[0]], v)
				vertexGroups[v] = connect[0]
			}
			delete(vertexByGroup, connect[1])
		case lIn:
			groupID = lG
			edgeToAdd = e.To
			addToExistingGroup = true
		case rIn:
			groupID = rG
			edgeToAdd = e.From
			addToExistingGroup = true
		default:
			vertexByGroup[lastUsedGroup] = []uint64{e.From, e.To}

			vertexGroups[e.From] = lastUsedGroup
			vertexGroups[e.To] = lastUsedGroup

			lastUsedGroup++
		}
		mst = append(mst, e)

		if addToExistingGroup {
			vertexGroups[edgeToAdd] = groupID
			if _, ok := vertexByGroup[groupID]; ok {
				vertexByGroup[groupID] = append(vertexByGroup[groupID], edgeToAdd)
			} else {
				vertexByGroup[groupID] = []uint64{edgeToAdd}
			}
		}
	}

	return
}

type Graph struct {
	RawEdges    []Edge
	Vertices    map[uint64]bool
	VertexEdges map[uint64]map[uint64]float64
	Undirected  bool
	NegEdges    bool
}

// Distance this structure is used to represent the distance of a vertex to
// another one in the graph taking the From vertex as immediate origin for the
// path with such distance
type Distance struct {
	From uint64
	Dist float64
}

// Edge representation of one of the edges of a directed graph,
// contains the from and to vertices, and weight for weighted graphs
type Edge struct {
	From   uint64
	To     uint64
	Weight float64
}

// byWeight Used to sort the graph edges by weight
type byWeight []Edge

func (a byWeight) Len() int           { return len(a) }
func (a byWeight) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byWeight) Less(i, j int) bool { return a[i].Weight < a[j].Weight }

// byDistance Used to sort the graph edges by distance
type byDistance []Distance

func (a byDistance) Len() int           { return len(a) }
func (a byDistance) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byDistance) Less(i, j int) bool { return a[i].Dist < a[j].Dist }

// GetUnWeightGraph Returns an unweighted graph containing the specified edges.
// Use the second boolean parameter in order to specify if the graph to be
// constructed is directed (true) or undirected (false)
func GetUnWeightGraph(edges [][]uint64, undirected bool) *Graph {
	aux := make([]Edge, len(edges))
	for i, edge := range edges {
		aux[i] = Edge{edge[0], edge[1], 0}
	}

	return GetGraph(aux, undirected)
}

// GetGraph Returns an weighted graph containing the specified edges.
// Use the second boolean parameter in order to specify if the graph to be
// constructed is directed (true) or undirected (false)
func GetGraph(edges []Edge, undirected bool) (ug *Graph) {
	var weight float64

	ug = &Graph{
		RawEdges:    edges,
		Vertices:    make(map[uint64]bool),
		VertexEdges: make(map[uint64]map[uint64]float64),
		Undirected:  undirected,
		NegEdges:    false,
	}

	for _, edge := range edges {
		weight = edge.Weight
		if weight < 0 {
			ug.NegEdges = true
		}
		ug.Vertices[edge.From] = true
		ug.Vertices[edge.To] = true
		if _, ok := ug.VertexEdges[edge.From]; ok {
			ug.VertexEdges[edge.From][edge.To] = weight
		} else {
			ug.VertexEdges[edge.From] = map[uint64]float64{edge.To: weight}
		}
		if undirected {
			if _, ok := ug.VertexEdges[edge.To]; ok {
				ug.VertexEdges[edge.To][edge.From] = weight
			} else {
				ug.VertexEdges[edge.To] = map[uint64]float64{edge.From: weight}
			}
		}
	}

	return
}

func main(){
	
}
