package main

import (
	"fmt"
	"time"
	"math/rand"
)
//最大堆结构
type maxHeap struct{
	count int//挂载数据个数
	data []int //堆数据容器
	cap int //数组容量
}
/*
获取堆中元素个数
 */
func (mh *maxHeap)GetCount()int  {
	return mh.count
}
/*
向堆中插入一个元素
 */
func (mh *maxHeap)Insert(n int){
	mh.count++
	if mh.count > mh.cap{
		mh.data = append(mh.data,n)
		mh.cap = len(mh.data)-1
	}else {
		mh.data[mh.count] = n
	}
	shiftUp(mh.count,mh.data)
}

/*
从堆中弹出顶部元素
为堆中最大值
 */
func (mh *maxHeap)Pop()(n int){
	if mh.count < 1{
		return
	}
	n = mh.data[1]
	mh.data[1] = mh.data[mh.count]
	mh.count--
	shiftDown(1,mh.count,mh.data)
	return
}

/*
将下标元素移动到合适位置保持最大堆得定义
 */
func shiftUp(index int,data []int){
    if index < 2 {
		return
	}
	parentIndex := index/2
	var t int
	for index >1 && data[index] > data[parentIndex]{
		t = data[index]
		data[index] = data[parentIndex]
		data[parentIndex] = t
		index = parentIndex
		parentIndex = index/2
	}
}

/*
将根元素移动到合适位置保持最大堆得定义
 */
func shiftDown(index,capacity int,data []int){
	t := data[index]
	for 2*index <= capacity {
		r := maxChild(index,data)
		if r[1] > t {
			data[index] = r[1]
			data[r[0]] = t
			index = r[0]
		}else {
			return
		}
	}
}
/*
的到两个孩子中大的那一个孩子
 */
func maxChild(index int,data []int) []int{
	//0储存孩子的下标，1储存孩子的值
	r := make([]int,2)
	n := index*2
    r[0] = n
	r[1] = data[r[0]]
	n++
	if n < len(data) && data[n] > r[1] {
		r[1] = data[n]
		r[0] = n
	}
	return r
}

//根据指定容量大小返回一个最大堆结构
func NewMaxHeapByCap(size int)*maxHeap{
	return &maxHeap{0,make([]int,size+1),size}
}

//根据传入的数组返回一个最大堆结构
func NewMaxHeapByArray(arr []int)*maxHeap{
	c := len(arr)
	l := c + 1
    mhArr := make([]int,l)
	for i,v := range arr{
		mhArr[i+1] = v
	}
	for i := c / 2;i > 0;i--{
		shiftDown(i,c,mhArr)
	}
	return &maxHeap{c,mhArr,l}
}
/*
堆排序
 */
func HeapSort(arr []int){
	mh := NewMaxHeapByArray(arr)
	for i := range arr{
		arr[i] = mh.Pop()
	}
}

func main(){
	//mh := NewMaxHeap(100)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	l := 100000
	arr := make([]int,l)
	for i:=0; i < l; i++ {
		arr[i] = r.Intn(10000000)
	}
    //fmt.Println("数组:",arr)
	start := time.Now().UnixNano()
	HeapSort(arr)
	end := time.Now().UnixNano()
	fmt.Println("堆排序耗时：",(end - start)/1e6,"ms")
	//fmt.Println("排序后：",arr)
}

