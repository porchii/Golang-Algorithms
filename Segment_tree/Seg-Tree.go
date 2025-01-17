package segmenttree

type Tree struct {
	T []int
}

func build(v int, l int, r int, t []int, a []int) {
	if r-l == 1 {
		t[v] = a[l]
		return
	}
	m := (l + r) / 2

	build(2*v+1, l, m, t, a)
	build(2*v+2, m, r, t, a)
	t[v] = t[2*v+1] + t[2*v+2]
}

func New(a []int) Tree {
	tmp := make([]int, len(a)*8)
	build(0, 0, len(a), tmp, a)

	return Tree{tmp}
}

// Segment Tree Methods

func (tree *Tree) Ask(v, l, r, ql, qr int) int {
	if l >= qr || r <= ql {
		return 0 // Neutral Value
	} else if l >= ql && r <= qr {
		return tree.T[v]
	}

	m := (l + r) / 2
	return tree.Ask(2*v+1, l, m, ql, qr) + tree.Ask(2*v+2, m, r, ql, qr)
}

func (tree *Tree) Upd(v, l, r, pos, val int) {
	if r-l == 1 {
		tree.T[v] = val
		return
	}

	m := (l + r) / 2
	if pos < m {
		tree.Upd(2*v+1, l, m, pos, val)
	} else {
		tree.Upd(2*v+2, m, r, pos, val)
	}
	tree.T[v] = tree.T[2*v+1] + tree.T[2*v+2]
}
