package dinic

import (
	"bufio"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"testing"
)

func readIntFromFile(filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()
	n, _ := strconv.Atoi(line)
	return n
}

func readGraphFromFile(filepath string) *Graph {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	sc.Scan()

	nm := strings.Split(sc.Text(), " ")
	n, _ := strconv.Atoi(nm[0])
	m, _ := strconv.Atoi(nm[1])

	g := newGraph(n)
	for i := 0; i < m; i++ {
		sc.Scan()
		edge := strings.Split(sc.Text(), " ")

		u, _ := strconv.Atoi(edge[0])
		v, _ := strconv.Atoi(edge[1])
		c, _ := strconv.Atoi(edge[2])

		g.addEdge(u, v, c)
	}
	return g
}

func TestDinic(t *testing.T) {
	files, err := ioutil.ReadDir("data/input")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		Graph := readGraphFromFile("data/input/" + file.Name())
		assert.Equal(t, Graph.maxFlow(0, Graph.n-1), readIntFromFile("data/output/"+file.Name()), "Graph in file "+file.Name()+": incorrect answer.")
	}
}
