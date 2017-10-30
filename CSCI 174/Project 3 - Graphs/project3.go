package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Vertex represents a vertex in a graph with properties to help with BFS and DFS
type Vertex struct {
	path  []int     // Holds the path
	adj   []*Vertex // Array to hold adjacent vertices
	color int       // -1 = white, 0 = grey, 1 =black
	dist  float64   // Distance to reach vertex in BFS
	id    int
	disc  int // Vertex first discovered in DFS
	fin   int // Vertex finished in DFS
}

// Edge represents a path from on vertex to another.
type Edge struct {
	x int // the id of the starting vertex
	y int // the id of the ending vertex
}

type Graph struct {
	verts  []*Vertex
	topOrd []int  // DFS topological ordering
	time   int    // Used to keep track of order vertices are visited in DFS
	tree   []Edge // Used to keep track of tree edges in DFS
	back   []Edge // Used to keep track of back edges in DFS
	forw   []Edge // Used to keep track of forward edges in DFS
	cros   []Edge // Used to keep track of cross edges in DFS
}

// Used to sort a graph by the id.
type sortGraph []*Vertex

func (v sortGraph) Len() int {
	return len(v)
}

func (v sortGraph) Swap(i, j int) {
	v[i], v[j] = v[j], v[i]
}

func (v sortGraph) Less(i, j int) bool {
	return v[i].id < v[j].id
}

// Used to sort a list of edges. First by the x(first) vertex, then by the y(second)
type sortEdge []Edge

func (e sortEdge) Len() int {
	return len(e)
}

func (e sortEdge) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

func (e sortEdge) Less(i, j int) bool {
	if e[i].x < e[j].x {
		return true
	} else if e[i].x == e[j].x {
		return e[i].y < e[j].y
	} else {
		return false
	}
}

// BFS implements Breadth First Search.
func BFS(g *Graph, s *Vertex) {
	// First set all the vertices to the default values
	for i := range g.verts {
		g.verts[i].color = -1            // White
		g.verts[i].dist = math.Inf(1.0)  // Set distance to unvisited vertices to infinity
		g.verts[i].path = make([]int, 0) // Empties path
	}
	s.color = 0 // Set starting vertex to grey
	s.dist = 0  // Set distance from the start to 0, since s is the start

	// Queue to put new vertices into
	queue := make([]*Vertex, 0)

	// Start with s in the queue
	queue = append(queue, s)

	// As long as the queue is not empty, keep going through and processing the next vertex
	for len(queue) != 0 {
		// Pop the next vertex off the queue
		u := queue[0]
		queue = queue[1:]

		// Handle all the adjacent vertices
		for i, v := range u.adj {
			// if the adjacent vertex is white (unhandled), set it to the defaults and add it to the queue
			if v.color == -1 {
				u.adj[i].color = 0
				u.adj[i].dist = u.dist + 1                       // Distant is equal to previous vertex +1
				u.adj[i].path = append(u.adj[i].path, u.path...) // Set path to the path to the previous vertex
				u.adj[i].path = append(u.adj[i].path, u.id)      // And then add the previous vertex
				queue = append(queue, u.adj[i])                  // Finally add it to the queue
			}
		}
		u.color = 1 // Change color of handled vertex to black after finished with it
		u.path = append(u.path, u.id)
	}
}

// DFS implements Depth First Search.
func DFS(g *Graph) {
	// Sets all vertices to default values
	for i := range g.verts {
		g.verts[i].color = -1            // Set color to white
		g.verts[i].path = make([]int, 0) // Empties path
	}
	g.time = 0 // Zeroes time at start of DFS
	g.back = make([]Edge, 0)
	g.cros = make([]Edge, 0)
	g.forw = make([]Edge, 0)
	g.tree = make([]Edge, 0)

	for i, v := range g.verts {
		// If color is white, visit vertex
		if v.color == -1 {
			DFSVisit(g, g.verts[i])
		}
	}
}

// DFSVisit is a helper functioin where the meat of the DFS algorithm is handled
func DFSVisit(g *Graph, u *Vertex) {
	g.time++        // Increase time every time DFSVisit is called
	u.disc = g.time // Sets vertex first discovered to current time
	u.color = 0     // Color of vertex is set to grey

	// Process all adjactent vertices
	for i, v := range u.adj {
		// If white, visit adjacent and add edge to the tree list
		if v.color == -1 {
			u.adj[i].path = append(u.adj[i].path, u.id)
			g.tree = append(g.tree, Edge{u.id, v.id})
			DFSVisit(g, u.adj[i])
		} else if v.color == 0 { // If grey, add edge to the back list
			g.back = append(g.back, Edge{u.id, v.id})
		} else { // If black,
			if v.disc > u.disc { // If u was discovered before adjacent vertex, add edge to forward list
				g.forw = append(g.forw, Edge{u.id, v.id})
			} else { // Otherwise add edge to cross list
				g.cros = append(g.cros, Edge{u.id, v.id})
			}
		}
	}

	u.color = 1                                 // Set color to black when done
	g.time++                                    // Increase time also when finishing vertex
	u.fin = g.time                              // Set finish time
	g.topOrd = append([]int{u.id}, g.topOrd...) // Add u to the front of the topological order list
}

// makeGraph creates a graph from passed in file. The file should be a list of start and end vertices.
// Each pair is on a different line.
func makeGraph(s string) Graph {
	g := Graph{}
	g.topOrd = make([]int, 0)
	f, err := os.Open(s)
	check(err)
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		v := &Vertex{}
		a := &Vertex{}

		str := scanner.Text()
		split := strings.Split(str, " ")
		id, _ := strconv.Atoi(split[0])
		adj, _ := strconv.Atoi(split[1])

		vexists := false // Vertex exists
		aexists := false // Adjacent vertex exists
		for i, u := range g.verts {
			if u.id == id {
				vexists = true
				v = g.verts[i]
			}
			if u.id == adj {
				aexists = true
				a = g.verts[i]
			}
		}

		if !aexists {
			a.id = adj
			g.verts = append(g.verts, a)
		}

		v.adj = append(v.adj, a)
		if !vexists {
			v.id = id
			g.verts = append(g.verts, v)
		}

	}
	f.Close()
	sort.Sort(sortGraph(g.verts))
	return g
}

// printBFS runs BFS and then prints out the distances and paths for each vertex
func printBFS(g *Graph) {
	BFS(g, g.verts[0])

	fmt.Println("Breadth First Search")
	fmt.Println("Vertex: Distance [Path]")

	for _, v := range g.verts {
		fmt.Printf("%d : %f %v\n", v.id, math.Trunc(v.dist), v.path) // id, dist, path
	}
}

// printDFS runs DFS and prints info
func printDFS(g *Graph) {
	DFS(g)

	fmt.Println("\nDepth First Search")
	fmt.Println("Vertex: Discover/Finish")
	for _, v := range g.verts {
		fmt.Printf("%d : %d/%d\n", v.id, v.disc, v.fin)
	}

	sort.Sort(sortEdge(g.tree))
	fmt.Print("Tree: [")
	for _, e := range g.tree {
		fmt.Printf("(%d, %d) ", e.x, e.y)
	}
	fmt.Println("]")

	sort.Sort(sortEdge(g.back))
	fmt.Print("Back: [")
	for _, e := range g.back {
		fmt.Printf("(%d, %d) ", e.x, e.y)
	}
	fmt.Println("]")

	sort.Sort(sortEdge(g.forw))
	fmt.Print("Forward: [")
	for _, e := range g.forw {
		fmt.Printf("(%d, %d) ", e.x, e.y)
	}
	fmt.Println("]")

	sort.Sort(sortEdge(g.cros))
	fmt.Print("Cross: [")
	for _, e := range g.cros {
		fmt.Printf("(%d, %d) ", e.x, e.y)
	}
	fmt.Println("]")

	fmt.Printf("Vertices in topological order: %v\n\n", g.topOrd)
}

// Check is used to stop program if any IO operations produce an error
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fmt.Println("Graph 1")
	g1 := makeGraph("graph1.dat")
	printBFS(&g1)
	printDFS(&g1)

	fmt.Println("\nGraph 2")
	g2 := makeGraph("graph2.dat")
	printBFS(&g2)
	printDFS(&g2)

	fmt.Println("\nGraph 3")
	g3 := makeGraph("graph3.dat")
	printBFS(&g3)
	printDFS(&g3)
}
