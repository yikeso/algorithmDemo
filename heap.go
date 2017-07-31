package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Compares interface {
	ComparesTo(c interface{}) int
}

//最大堆结构
type maxHeap struct {
	count int   //挂载数据个数
	data  []int //堆数据容器
	cap   int   //数组容量
}

/*
获取堆中元素个数
*/
func (mh *maxHeap) GetCount() int {
	return mh.count
}

/*
向堆中插入一个元素
*/
func (mh *maxHeap) Insert(n int) {
	mh.count++
	if mh.count > mh.cap {
		mh.data = append(mh.data, n)
		mh.cap = len(mh.data) - 1
	} else {
		mh.data[mh.count] = n
	}
	shiftUp(mh.count, mh.data)
}

/*
从堆中弹出顶部元素
为堆中最大值
*/
func (mh *maxHeap) Pop() (n int) {
	if mh.count < 1 {
		return
	}
	n = mh.data[1]
	mh.data[1] = mh.data[mh.count]
	mh.count--
	shiftDown(1, mh.count, mh.data)
	return
}

/*
将下标元素移动到合适位置保持最大堆得定义
*/
func shiftUp(index int, data []int) {
	if index < 2 {
		return
	}
	parentIndex := index / 2
	var t int
	for index > 1 && data[index] > data[parentIndex] {
		t = data[index]
		data[index] = data[parentIndex]
		data[parentIndex] = t
		index = parentIndex
		parentIndex = index / 2
	}
}

/*
将元素移动到合适位置保持最大堆得定义
根元素为下标为1的值
*/
func shiftDown(index, capacity int, data []int) {
	t := data[index]
	for 2*index <= capacity {
		r := maxChild(index, data)
		if r[1] > t {
			data[index] = r[1]
			data[r[0]] = t
			index = r[0]
		} else {
			return
		}
	}
}

/*
的到两个孩子中大的那一个孩子
堆的根节点为数组下标为1的值
*/
func maxChild(index int, data []int) []int {
	//0储存孩子的下标，1储存孩子的值
	r := make([]int, 2)
	n := index * 2
	r[0] = n
	r[1] = data[r[0]]
	n++
	if n < len(data) && data[n] > r[1] {
		r[1] = data[n]
		r[0] = n
	}
	return r
}

/*
将元素移动到合适位置保持最大堆得定义
根元素为下标为0的值
*/
func shiftDown2(index, capacity int, data []int) {
	t := data[index]
	for 2*index+1 < capacity {
		r := maxChild2(index, capacity, data)
		if r[1] > t {
			data[index] = r[1]
			data[r[0]] = t
			index = r[0]
		} else {
			return
		}
	}
}

/*
的到两个孩子中大的那一个孩子
堆的根节点为数组下标为0的值
*/
func maxChild2(index, capacity int, data []int) []int {
	//0储存孩子的下标，1储存孩子的值
	r := make([]int, 2)
	n := index*2 + 1
	r[0] = n
	r[1] = data[r[0]]
	n++
	if n < capacity && data[n] > r[1] {
		r[1] = data[n]
		r[0] = n
	}
	return r
}

//根据指定容量大小返回一个最大堆结构
func NewMaxHeapByCap(size int) *maxHeap {
	return &maxHeap{0, make([]int, size+1), size}
}

//根据传入的数组返回一个最大堆结构
func NewMaxHeapByArray(arr []int) *maxHeap {
	c := len(arr)
	l := c + 1
	mhArr := make([]int, l)
	for i, v := range arr {
		mhArr[i+1] = v
	}
	for i := c / 2; i > 0; i-- {
		shiftDown(i, c, mhArr)
	}
	return &maxHeap{c, mhArr, l}
}

/*
堆排序
*/
func HeapSort(arr []int) {
	mh := NewMaxHeapByArray(arr)
	for i := range arr {
		arr[i] = mh.Pop()
	}
}

/*
堆排序优化，降低空间复杂度
*/
func HeapSort2(arr []int) {
	l := len(arr)
	shiftDownArray(arr, l)
	l--
	var t int
	for ; l > 0; l-- {
		t = arr[l]
		arr[l] = arr[0]
		arr[0] = t
		shiftDown2(0, l, arr)
	}
}

/*
从指定的下标开始进行shiftDown操作
*/
func shiftDownArray(arr []int, l int) {
	for i := (l - 1) / 2; i >= 0; i-- {
		shiftDown2(i, l, arr)
	}
}

type testItem struct {
	Data int
}

func (t *testItem) ComparesTo(c testItem) int {
	return t.Data - c.Data
}

//最大堆结构
type maxIndexHeap struct {
	count  int //挂载数据个数
	indexs []int
	data   []Compares //堆数据容器
	cap    int        //数组容量
}

/*
根据数组创建最大索引堆
*/
func NewMaxIndexHeapByArray(arr []Compares) *maxIndexHeap {
	l := len(arr)
	data := make([]Compares, l)
	for i := range arr {
		data[i] = arr[i]
	}
	indexs := make([]int, l)
	for i := range indexs {
		indexs[i] = i
	}
	for i := (l - 1) / 2; i >= 0; i-- {
		shiftDownIndex(data, indexs, i, l)
	}
	return &maxIndexHeap{l, indexs, data, l}
}

/*
将元素移动到合适位置保持最大堆得定义
根元素为下标为0的值
i为进行shiftDown操作的索引数组的下标
l为数组的长度
*/
func shiftDownIndex(data []Compares, indexs []int, i, l int) {
	t := indexs[i]
	var r []int
	for i*2+1 < l {
		r = maxChildIndex(data, indexs, i, l)
		if data[r[1]].ComparesTo(data[t]) > 0 {
			indexs[i] = r[1]
			indexs[r[0]] = t
			i = r[0]
		}
	}
}

/*
的到两个孩子中大的那一个孩子的索引数组的值
堆的根节点为数组下标为0的值
*/
func maxChildIndex(data []Compares, indexs []int, i, l int) []int {
	//0储存孩子的下标，1储存孩子的值
	r := make([]int, 2)
	n := i*2 + 1
	r[0] = n
	r[1] = indexs[r[0]]
	n++
	if n < l && data[indexs[n]].ComparesTo(data[r[1]]) > 0 {
		r[1] = indexs[n]
		r[0] = n
	}
	return r
}
func ChangeItem(ind int, item Compares) {

}

/*
将下标元素移动到合适位置保持最大堆得定义
*/
func shiftUpIndex(data []Compares, indexs []int, i int) {
	if i < 1 {
		return
	}
	parentIndex := (i - 1) / 2
	var t int
	for i > 0 && data[indexs[i]].ComparesTo(data[indexs[parentIndex]]) > 0 {
		t = indexs[i]
		indexs[i] = indexs[parentIndex]
		indexs[parentIndex] = t
		i = parentIndex
		parentIndex = (i - 1) / 2
	}
}

func main() {
	//mh := NewMaxHeap(100)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	l := 100000
	arr := make([]int, l)
	for i := 0; i < l; i++ {
		arr[i] = r.Intn(l * 100)
	}
	//fmt.Println("数组:",arr)
	start := time.Now().UnixNano()
	HeapSort2(arr)
	end := time.Now().UnixNano()
	fmt.Println("堆排序耗时：", (end-start)/1e6, "ms")
	//fmt.Println("排序后：",arr)
}
