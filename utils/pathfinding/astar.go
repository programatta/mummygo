package pathfinding

import (
	"fmt"
)

//Defining possible graph elements
const (
	UNKNOWN int = iota - 1
	LAND
	WALL
	START
	STOP
)

//MapData to be applied to pathfinding.
type MapData [][]int

// //NewMapData Return a new MapData by value given some dimensions
// func NewMapData(rows, cols int) *MapData {
// 	result := make(MapData, rows)
// 	for i := 0; i < rows; i++ {
// 		result[i] = make([]int, cols)
// 	}
// 	return &result
// }

//Node A node is just a set of X, Y coordinates with a parent node and a
//heuristic value H
type Node struct {
	X, Y   int //Using int for efficiency
	parent *Node
	H      int //Heuristic (aproximate distance)
	cost   int //Path cost for this node
}

//NewNode Create a new node
func NewNode(x, y int) *Node {
	node := &Node{
		X:      x,
		Y:      y,
		parent: nil,
		H:      0,
		cost:   0,
	}
	return node
}

//String Return string representation of the node
func (n *Node) String() string {
	return fmt.Sprintf("<Node X:%d Y:%d addr:%d>", n.X, n.Y, &n)
}

//Graph Start, end nodes and a slice of nodes
type Graph struct {
	start, stop *Node
	nodes       []*Node
	data        *MapData
}

//NewGraph Return a Graph from a map of coordinates (those that are passible)
func NewGraph(mapData *MapData) *Graph {
	var start, stop *Node
	var nodes []*Node
	for i, row := range *mapData {
		for j, _type := range row {
			if _type == START || _type == STOP {
				node := NewNode(i, j)
				nodes = append(nodes, node)
				if _type == START {
					start = node
				}
				if _type == STOP {
					stop = node
				}
			}
		}
	}
	g := &Graph{
		nodes: nodes,
		start: start,
		stop:  stop,
		data:  mapData,
	}
	return g
}

//Node Get *Node based on X, Y coordinates.
func (g *Graph) Node(x, y int) *Node {
	//Check if node is not already in the graph and append that node
	for _, n := range g.nodes {
		if n.X == x && n.Y == y {
			return n
		}
	}
	mapData := *g.data
	if mapData[x][y] == LAND || mapData[x][y] == STOP {
		//Create a new node and add it to the graph
		n := NewNode(x, y)
		g.nodes = append(g.nodes, n)
		return n
	}
	return nil
}

//Get the nodes near some node
func (g *Graph) adjacentNodes(node *Node) []*Node {
	var result []*Node
	mapData := *g.data
	rows := len(mapData)
	cols := len(mapData[0])

	//If the coordinates are passable then create a new node and add it
	if node.X <= rows && node.Y+1 < cols {
		if newNode := g.Node(node.X, node.Y+1); newNode != nil {
			result = append(result, newNode)
		}
	}
	if node.X <= rows && node.Y-1 >= 0 {
		newNode := g.Node(node.X, node.Y-1)
		if newNode != nil {
			result = append(result, newNode)
		}
	}
	if node.Y <= cols && node.X+1 < rows {
		newNode := g.Node(node.X+1, node.Y)
		if newNode != nil {
			result = append(result, newNode)
		}
	}
	if node.Y <= cols && node.X-1 >= 0 {
		newNode := g.Node(node.X-1, node.Y)
		if newNode != nil {
			result = append(result, newNode)
		}
	}
	return result
}

// ---- A start ----

//Astar A* search algorithm. See http://en.wikipedia.org/wiki/A*_search_algorithm
func Astar(graph *Graph) []*Node {
	var path, openSet, closedSet []*Node

	openSet = append(openSet, graph.start)
	for len(openSet) != 0 {
		//Get the node with the min H
		current := minH(openSet)
		if current.parent != nil {
			current.cost = current.parent.cost + 1
		}
		if current == graph.stop {
			return retracePath(current)
		}
		openSet = removeNode(openSet, current)
		closedSet = append(closedSet, current)
		for _, tile := range graph.adjacentNodes(current) {
			if tile != nil && graph.stop != nil && !hasNode(closedSet, tile) {
				tile.H = Heuristic(graph, tile) + current.cost
				if !hasNode(openSet, tile) {
					openSet = append(openSet, tile)
				}
				tile.parent = current
			}
		}
	}
	return path
}

//Return the node with the minimum H
func minH(nodes []*Node) *Node {
	if len(nodes) == 0 {
		return nil
	}
	resultNode := nodes[0]
	minH := resultNode.H
	for _, node := range nodes {
		if node.H < minH {
			minH = node.H
			resultNode = node
		}
	}
	return resultNode
}

//Heuristic In our particular case: Manhatan distance
func Heuristic(graph *Graph, tile *Node) int {
	return abs(graph.stop.X-tile.X) + abs(graph.stop.Y-tile.Y)
}

func retracePath(currentNode *Node) []*Node {
	var path []*Node
	path = append(path, currentNode)
	for currentNode.parent != nil {
		path = append(path, currentNode.parent)
		currentNode = currentNode.parent
	}
	//Reverse path
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}
	return path
}

func removeNode(nodes []*Node, node *Node) []*Node {
	ith := -1
	for i, n := range nodes {
		if n == node {
			ith = i
			break
		}
	}
	if ith != -1 {
		copy(nodes[ith:], nodes[ith+1:])
		nodes = nodes[:len(nodes)-1]
	}
	return nodes
}

func hasNode(nodes []*Node, node *Node) bool {
	for _, n := range nodes {
		if n == node {
			return true
		}
	}
	return false
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
