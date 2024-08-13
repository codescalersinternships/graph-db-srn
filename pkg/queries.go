package graph

import (
	"slices"
)

func (graph *Graph)FilterNodeByProperty(key, value string)[]uint{
	result:= []uint{}
	for _, node := range graph.nodes{
		if ok:= node.details[key] ;ok!=value{
			continue
		}
		result = append(result, node.id)
	}
	return result
}

func (graph *Graph) QueryGetParents(id int) []uint {
	return graph.edgesParent[uint(id)]
}

func (graph *Graph) QueryGetGrandParents(id int) []uint {
	var grandparents [] uint
	for _, parent := range graph.edgesParent[uint(id)] {
		grandparents = append(grandparents, graph.edgesParent[parent]...)
	}
	
	return grandparents
}

func (graph *Graph) QueryGetSibilings(id int) []uint {
	var siblings []uint
	parents := graph.edgesParent[uint(id)]
	for _, child := range parents {
		if child == uint(id) {
			continue
		}
		siblings = append(siblings, child)
	}
	return siblings
}

func (graph *Graph) QueryGetCousins(id uint) []uint {
	var cousins []uint
	parents := graph.edgesParent[id]
	var grandparents []uint
	for _, parent := range parents {
		grandparents = append(grandparents, graph.edgesParent[parent]...)
	}
	slices.Sort(parents)
	uniCousins := make(map[uint]bool);
	for _, grandparent := range grandparents{
		for _, uncle := range graph.edgesChild[grandparent] {
			if _, found := slices.BinarySearch(parents, uncle); !found {
				for _,cousin := range graph.edgesChild[uncle] {
					uniCousins[cousin] = true
				}
			}
		}
		
	}
	for key, value := range uniCousins{
		if value {
			cousins = append(cousins, key)
		}	
	}
	return cousins
}