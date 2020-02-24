package main

import (
	"container/heap"
	"fmt"
)

type Element struct {
	Value    interface{}
	Priority int64
	index    int
}

type MinHeap []*Element

func NewMinHeap() *MinHeap {
	mh := &MinHeap{}
	heap.Init(mh)
	return mh
}

func (mh MinHeap) Len() int { return len(mh) }

func (mh MinHeap) Less(i, j int) bool {
	return mh[i].Priority < mh[j].Priority
}

func (mh MinHeap) Swap(i, j int) {
	mh[i], mh[j] = mh[j], mh[i]
	mh[i].index = i
	mh[j].index = j
}

func (mh *MinHeap) Push(x interface{}) {
	n := len(*mh)
	item := x.(*Element)
	item.index = n
	*mh = append(*mh, item)
}

func (mh *MinHeap) Pop() interface{} {
	old := *mh
	n := len(old)
	item := old[n-1]
	item.index = -1
	*mh = old[0 : n-1]
	return item
}

func (mh *MinHeap) PushEl(el *Element) {
	heap.Push(mh, el)
}

func (mh *MinHeap) PopEl() *Element {
	el := heap.Pop(mh)
	return el.(*Element)
}

func (mh *MinHeap) PeekEl() *Element {
	items := *mh
	return items[0]
}

func (mh *MinHeap) UpdateEl(el *Element, priority int64) {
	heap.Remove(mh, el.index)
	el.Priority = priority
	heap.Push(mh, el)
}

func (mh *MinHeap) RemoveEl(el *Element) {
	heap.Remove(mh, el.index)
}

func main() {
	mh := NewMinHeap()

	mh.PushEl(&Element{
		Value:    "2.2.2.2",
		Priority: 2,
	})
	mh.PushEl(&Element{
		Value:    "4.4.4.4",
		Priority: 4,
	})
	mh.PushEl(&Element{
		Value:    "6.6.6.6",
		Priority: 6,
	})
	mh.PushEl(&Element{
		Value:    "3.3.3.3",
		Priority: 3,
	})
	mh.PushEl(&Element{
		Value:    "5.5.5.5",
		Priority: 5,
	})

	fmt.Println(mh.PeekEl())
	fmt.Println(mh.PopEl())
	fmt.Println(mh.PeekEl())
	fmt.Println(mh.PopEl())
	fmt.Println(mh.PeekEl())
	fmt.Println(mh.PopEl())
	fmt.Println(mh.PeekEl())
}
