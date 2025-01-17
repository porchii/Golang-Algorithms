package segmenttreestupid

type Item struct {
	arr []int
}

func New(arr []int) Item {
	return Item{arr}
}

func (item *Item) Ask(l, r int) int {
	var sum int

	for i := l; i < r; i++ {
		sum += item.arr[i]
	}

	return sum
}

func (item *Item) Upd(pos, val int) {
	item.arr[pos-1] = val
}
