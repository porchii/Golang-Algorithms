package tester

import (
	segmenttree "root/Segment_tree"
	segmenttreeplus "root/Segment_tree_plus"
	segmenttreeplusstupid "root/tests/Segment_tree_test/Segment_tree_plus_stupid"
	segmenttreestupid "root/tests/Segment_tree_test/Segment_tree_stupid"
)

var example []int = make([]int, 10)

func Test_seg_tree() {
	tree := segmenttree.New(example)
	not_tree := segmenttreestupid.New(example)

	not_tree.Upd(1, 1)
	tree.Upd(0, 0, len(example), 1, 1)
}

func Test_seg_tree_plus() {
	tree := segmenttreeplus.New(example)
	not_tree := segmenttreeplusstupid.New(example)

	not_tree.Upd(1, 1, 1)
	tree.Upd(0, 0, len(example), 1, 1, 1)
}
