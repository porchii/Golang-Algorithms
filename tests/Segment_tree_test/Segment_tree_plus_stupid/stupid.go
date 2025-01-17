package segmenttreeplusstupid

type Item struct {
	arr []int
}

func New(arr []int) Item {
	return Item{arr}
}

func (item *Item) Ask(l, r int) int {
	var mx int = -1

	for i := l; i < r; i++ {
		mx = max(mx, item.arr[i])
	}
	return mx
}

func (item *Item) Upd(l, r, val int) {
	for i := l; i < r; i++ {
		item.arr[i] = val
	}
}
