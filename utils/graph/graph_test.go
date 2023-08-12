package graph

import "testing"

func TestFloydWarshall(t *testing.T) {
	a := New("AA", 1)
	b := New("BB", 1)
	c := New("CC", 1)
	d := New("DD", 1)

	a.AddChild(b, 3)
	a.AddChild(d, 5)
	b.AddChild(a, 2)
	b.AddChild(d, 4)
	c.AddChild(b, 1)
	d.AddChild(c, 2)

	nodes := []*Node{a, b, c, d}

    distances := FloydWarshall(nodes)

    if distances[a][d] != 5 {
        t.Error("Not the shortest path!")
    }

    if distances[a][b] != 3 {
        t.Error("Not the shortest path!")
    }
    
    if distances[a][c] != 7 {
        t.Error("Not the shortest path!")
    }
}


func TestDijkstra(t *testing.T) {
	a := New("AA", 0)
	b := New("BB", 13)
	c := New("CC", 2)
	d := New("DD", 20)
	e := New("EE", 3)
	f := New("FF", 0)
	g := New("GG", 0)
	h := New("HH", 22)
	i := New("II", 0)
	j := New("JJ", 21)

	a.AddChild(d, 1)
	a.AddChild(i, 1)
	a.AddChild(b, 1)

	
	b.AddChild(c, 1)
	b.AddChild(a, 1)
    
	c.AddChild(b, 1)
	c.AddChild(d, 1)

	d.AddChild(c, 1)
	d.AddChild(a, 1)
	d.AddChild(e, 1)

	e.AddChild(f, 1)
    e.AddChild(d, 1)
	
    f.AddChild(e, 1)
	f.AddChild(g, 1)
	
    g.AddChild(f, 1)
	g.AddChild(h, 1)
	
    h.AddChild(g, 1)

    i.AddChild(a, 1)
    i.AddChild(j, 1)

    j.AddChild(i, 1)

    nodes := []*Node{
		a, b, c, d, e, f, g, h, i, j,
	}

    distances := a.Dijkstra(nodes)

	if distances[b] != 1 {
		t.Errorf("Shortest path from AA to BB should be 1, not %d", distances[b])
	}
    if distances[c] != 2 {
		t.Errorf("Shortest path from AA to CC should be 2, not %d", distances[c])
	}
    if distances[d] != 1 {
		t.Errorf("Shortest path from AA to DD should be 1, not %d", distances[d])
	}
    if distances[e] != 2 {
		t.Errorf("Shortest path from AA to EE should be 2, not %d", distances[e])
	}
    if distances[h] != 5 {
		t.Errorf("Shortest path from AA to DD should be 5, not %d", distances[h])
	}
}
