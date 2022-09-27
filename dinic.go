package dinic

const INF = 1 << 20

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type Edge struct {
	to, cap, rev int
}

type Graph struct {
	n    int
	Edge [][]Edge
}

func newGraph(n int) *Graph {
	return &Graph{n: n, Edge: make([][]Edge, n)}
}

func (g *Graph) addEdge(from, to, cap int) {
	g.Edge[from] = append(g.Edge[from], Edge{to, cap, len(g.Edge[to])})
	g.Edge[to] = append(g.Edge[to], Edge{from, 0, len(g.Edge[from]) - 1})
}

func (g *Graph) bfs(s int, t int, level []int) bool {
	for i := range level {
		level[i] = -1
	}
	q := []int{s}
	level[s] = 0
	for len(q) > 0 {
		v := q[0]
		q = q[1:]
		for _, e := range g.Edge[v] {
			if e.cap > 0 && level[e.to] < 0 {
				level[e.to] = level[v] + 1
				q = append(q, e.to)
			}
		}
	}
	return level[t] >= 0
}

func (g *Graph) dfs(v int, t int, flow int, level []int, ptr []int) int {
	if flow == 0 {
		return 0
	}
	if v == t {
		return flow
	}

	for i := ptr[v]; i < len(g.Edge[v]); i++ {
		ptr[v] = i
		e := &g.Edge[v][i]
		if e.cap > 0 && level[v] < level[e.to] {
			d := g.dfs(e.to, t, min(flow, e.cap), level, ptr)
			if d > 0 {
				e.cap -= d
				g.Edge[e.to][e.rev].cap += d
				return d
			}
		}
	}
	return 0
}

func (g *Graph) maxFlow(s, t int) int {
	level := make([]int, g.n)
	ptr := make([]int, g.n)
	flow := 0
	for g.bfs(s, t, level) {
		for i := range ptr {
			ptr[i] = 0
		}
		for {
			f := g.dfs(s, t, INF, level, ptr)
			if f == 0 {
				break
			}
			flow += f
		}
	}
	return flow
}
