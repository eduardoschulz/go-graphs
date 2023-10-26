package graphs 


type Graph struct {
  Vertices []Vertex
}

type Vertex struct {
  value int
  edges []Edge
  color string
}

type Edge struct {
  vertex *Vertex
}

func (g *Graph) addVertex(v *Vertex){
  g.Vertices = append(g.Vertices, *v)
}

/*

func (g *Graph) addMultipleVertices(v ...*Vertex){
  for _, v := range v {
    g.addVertex(v)
  }
}
*/


func creatingVertex(value int)*Vertex{
  return &Vertex{value, nil, "white"}
}

func appendEdges(v *Vertex, e *Edge){
  v.edges = append(v.edges, *e)
}
/*
func appendEdgesBothWays(v1 *Vertex, v2 *Vertex){
  appendEdges(v1, &Edge{v2})
  appendEdges(v2, &Edge{v1})
}
*/

func appendVertices(g *Graph, v *Vertex){
  g.Vertices = append(g.Vertices, *v)
}


func visit(g *Graph, v *Vertex){
  v.color = "gray" // visited 
}



