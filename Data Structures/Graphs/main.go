package main

import "fmt"

// Graph structure
type Graph struct {
	vertices []*Vertex
}
// Vertex structure
type Vertex struct {
	key int
	adjacent []*Vertex
}

// AddVertex adds a Vertex to the Graph
func (g *Graph) AddVertex(k int)  {
	if contains(g.vertices, k) {
		err := fmt.Errorf("vertex %v not added because it is an existing key", k)
		fmt.Println(err.Error())
	} else {
		g.vertices = append(g.vertices, &Vertex{key: k})
	}
}


// AddEdge adds an edge to the graph
func (g *Graph) AddEdge(from,to int)  {
	// get vertex
	fromVertex := g.getVertex(from)
	toVertex := g.getVertex(to)

	// check error
	if fromVertex == nil || toVertex == nil {
		err := fmt.Errorf("cannot add edge for vertex that doest not exist")
		fmt.Println(err.Error())
	} else if contains(fromVertex.adjacent, to) {
		err := fmt.Errorf("existing edge")
		fmt.Println(err.Error())
	} else {
		// add edge
		fromVertex.adjacent = append(fromVertex.adjacent, toVertex)
	}
}

// getVertex returns a pointer to the Vertex with a key integer
func (g *Graph) getVertex(k int) *Vertex {
	for i,vertex := range g.vertices {
		if vertex.key == k {
			return g.vertices[i]
		}
	}
	return nil
}
// contains
func contains(s []*Vertex, k int) bool {
	for _, v := range s {
		if k == v.key {
			return true
		}
	}
	return false
}

// Print will print the adjacent list for each vertex of the graph
func (g *Graph) Print()  {
	for _, vertex := range g.vertices {
		fmt.Printf("\nVertex %v : ", vertex.key)
		for _, v := range vertex.adjacent {
			fmt.Printf("%v ", v.key)
		}
	}
}

func main() {
	test := &Graph{}

	for i := 0; i < 5; i++ {
		test.AddVertex(i)
	}

	test.AddEdge(1, 2)
	test.AddEdge(1, 2)

	test.Print()
}