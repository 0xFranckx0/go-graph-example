package socnet

import (
	"bufio"
	"gonum.org/v1/gonum/graph/path"
	"gonum.org/v1/gonum/graph/simple"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

var (
	g *simple.UndirectedGraph
)

// The init has in charge to build the graph with dataset file
// I chose a undirected and unweighted graph to address the problem
// because this kind of graph matches the requirement for the
// Dijkstra algorithm and  finding neighbors.
func init() {

	g = simple.NewUndirectedGraph()

	// In order to build the graph we need to add Nodes and edges
	var datasetFile string
	// First we try to setup the default location of the file with the env
	// variable DATASETFILE
	if datasetFile = os.Getenv("DATASETFILE"); datasetFile == "" {
		datasetFile = os.Getenv("GOPATH") + "/src/github.com/go-graph-example/dataset"
	}
	file, err := os.Open(datasetFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// The buf map is just used to avoid inserting duplicates because it
	// triggers a panic. And it's  more efficient  instead of testing
	// if the node belongs to the graph for each insertion.
	buf := make(map[int]bool)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), " ")
		ix0, err := strconv.Atoi(s[0])
		if err != nil {
			log.Fatal("conversion error")
		}
		if !buf[ix0] {
			buf[ix0] = true
			g.AddNode(simple.Node(ix0))
		}
		ix1, err := strconv.Atoi(s[1])
		if err != nil {
			log.Fatal("conversion error")
		}
		if !buf[ix1] {
			buf[ix1] = true
			g.AddNode(simple.Node(ix1))
		}

		g.SetEdge(g.NewEdge(simple.Node(ix0), simple.Node(ix1)))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

// The GetDistance function takes two strings as argument and returns
// the distance as an int and the path as an int slice.
func GetDistance(from, to string) (int, []int) {
	// we first need to convert the From and To strings into graph.Node types
	i, err := strconv.Atoi(from)
	if err != nil {
		log.Print("conversion error")
		return -1, nil
	}
	f := simple.Node(i)

	i, err = strconv.Atoi(to)
	if err != nil {
		log.Print("conversion error")
		return -1, nil
	}
	t := simple.Node(i)

	// Once the conversion is done we are running the Dijkstra algorithm
	// to find the shortest path between from and to.
	var (
		pt       path.Shortest
		panicked bool
	)

	func() {
		defer func() {
			panicked = recover() != nil
		}()
		pt = path.DijkstraFrom(f, g)
	}()

	// If all went well we can fill out the slice which will contain the path
	// and set the distance.
	distance := 0
	var nodes []int
	if !panicked {
		p, weight := pt.To(t)
		for _, x := range p {
			nodes = append(nodes, int(x.ID()))
		}
		if int(weight) > 0 {
			distance = int(weight) + 1
		}
	} else {
		// we return distance -1 to specifies that something wrong occured while
		// parsing the path.
		return -1, nodes
	}
	// we return distance 0 if no path exists, that will trigger a 404 Not
	// Found by the API.
	return distance, nodes
}

// GetCommonFriends function takes two strings as arguments that represent the
// two IDs we want to know common friends. It returns a slice of common
// friends.
func GetCommonFriends(from, to string) []int {

	// we first need to convert the From and To strings into graph.Node types
	i, err := strconv.Atoi(from)
	if err != nil {
		log.Print("conversion error")
		return nil
	}
	f := simple.Node(i)

	i, err = strconv.Atoi(to)
	if err != nil {
		log.Print("conversion error")
		return nil
	}
	t := simple.Node(i)

	// We need a map to store friends and determine which ones are common.
	// All map entries with a value of 1 are common.
	dict := make(map[int]int)

	// In the first iteration we fill out the map with all IDs in the From
	// list. All values are set up to 0.
	list := g.From(f)
	for _, x := range list {
		dict[int(x.ID())] = 0
	}

	// In the second iteration we just need to take about IDs that are already
	// in the map, and increment their value to 1.
	list = g.From(t)
	for _, x := range list {
		if val, ok := dict[int(x.ID())]; ok {
			dict[int(x.ID())] = val + 1
		}
	}

	// Then we just need to fill out the slice of int that we will return, with
	// all keys of the map with a value equal to 1.
	var friends []int
	for k, v := range dict {
		if v == 1 {
			friends = append(friends, k)
		}
	}
	// For convenience we return a sorted slice.
	if friends != nil {
		sort.Ints(friends)
	} else {
		friends = make([]int, 0)
	}

	return friends
}
