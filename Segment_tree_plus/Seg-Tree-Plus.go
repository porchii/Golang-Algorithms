package segmenttreeplus

type Tree struct {
	T []int
	C []int
}

func build(v, l, r int, t []int, c []int, arr []int) {
	if r-l == 1 {
		t[v] = arr[l]
		return
	}

	m := (r + l) / 2
	build(2*v+1, l, m, t, c, arr)
	build(2*v+2, m, r, t, c, arr)

	t[v] = max(t[2*v+1], t[2*v+2])
}

func New(arr []int) Tree {
	tmp := make([]int, len(arr)*8)
	c := make([]int, len(arr)*8)

	for i := 0; i < len(arr)*8; i++ {
		c[i] = -1
	}

	build(0, 0, len(arr), tmp, c, arr)

	return Tree{tmp, c}
}

// Seg Tree Methods

func (tree *Tree) push(v int) {
	if tree.C[v] != -1 {
		tree.C[2*v+1] = tree.C[v]
		tree.C[2*v+2] = tree.C[v]
		tree.C[v] = -1
	}
}

func (tree *Tree) update(v int) {
	var left_v, right_v int

	if tree.C[2*v+1] != -1 {
		left_v = tree.C[2*v+1]
	} else {
		left_v = tree.T[2*v+1]
	}

	if tree.C[2*v+2] != -1 {
		right_v = tree.C[2*v+2]
	} else {
		right_v = tree.T[2*v+2]
	}

	tree.T[v] = max(left_v, right_v)
}

func (tree *Tree) Ask(v, l, r, ql, qr int) int {
	if l >= qr || r <= ql {
		return -1
	} else if l >= ql && r <= qr {
		if tree.C[v] != -1 {
			return tree.C[v]
		} else {
			return tree.T[v]
		}
	}

	m := (l + r) / 2
	tree.push(v)
	x := max(tree.Ask(2*v+1, l, m, ql, qr), tree.Ask(2*v+2, m, r, ql, qr))
	tree.update(v)
	return x
}

func (tree *Tree) Upd(v, l, r, ql, qr, val int) {
	if l >= qr || r <= ql {
		return
	} else if l >= ql && r <= qr {
		tree.C[v] = val
		return
	}

	m := (l + r) / 2
	tree.push(v)
	tree.Upd(2*v+1, l, m, ql, qr, val)
	tree.Upd(2*v+2, m, r, ql, qr, val)
	tree.update(v)
}
