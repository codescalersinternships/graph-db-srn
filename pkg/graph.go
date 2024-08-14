package graph

type node struct {
	id      uint
	details map[string]string
}

type Graph struct {
	nodes       map[uint]node
	edgesChild  map[uint]([]uint)
	edgesParent map[uint]([]uint)
}

// parent to []child

func NewGraph() (graph *Graph) {
	return &Graph{}
}

func newNode(id uint, details map[string]string) *node {
	return &node{id: id, details: details}
}

func (g *Graph) AddNode(details map[string]string) {
	len := len(g.nodes)
	g.nodes[uint(len)+1] = *newNode(uint(len)+1, details)
}

func (g *Graph) AddEdge(from uint, to uint, relation string) {
	g.edgesChild[from] = append(g.edgesChild[from], to)
	g.edgesChild[to] = append(g.edgesChild[to], from)
}
func (g *Graph) GetNodeByID(id uint) node {
	return g.nodes[id]
}

// query_grandparents
// query_siblings
// query_cousins
// filter_vertices
