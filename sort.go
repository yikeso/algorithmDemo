package main

import (
	"time"
	"fmt"
	"math/rand"
	"runtime"
)

func main(){
	runtime.GOMAXPROCS(runtime.NumCPU())
	l := 100000
	arr := getIntArry(l)
	//fmt.Println("排序前：",arr)
	fmt.Println("数组长度：",l)
	SelectionSort(arr)
	InsertSort(arr)
	//BubbleSort(arr)
	MergeSortEntry(arr)
	MergeSortBuEntry(arr)
	QuikSortEntry(arr)
	QuikSortEntry2(arr)
	QuikSortEntry3(arr)
}

func getIntArry(l int)[]int{
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	arr := make([]int,l)
	for i:=0; i < l; i++ {
		arr[i] = r.Intn(100*l)
	}
	return arr
}
//选择排序
func SelectionSort(a []int) {
	start := time.Now().UnixNano()
	l := len(a)
	if l < 1 {
		fmt.Println("数组不得为空")
		return
	}
	arr := make([]int,l)
	for i := 0;i < l;i++ {
		arr[i] = a[i]
	}
	//min[0]储存最小数的下标
	//min[1]储存最小数的值
	min := new([2]int)
	for i := 0;i < l-1;i++{
		min[1] = arr[i]
		min[0] = i
		for j := i;j < l;j++ {
			if arr[j] < min[1]{
				min[1] = arr[j]
				min[0] = j
			}
		}
        arr[min[0]] = arr[i]
		arr[i] = min[1]
	}
	end := time.Now().UnixNano()
	fmt.Println("选择排序耗时：",(end - start)/1e6,"ms")
}
//插入排序
func InsertSort(a []int){
	start := time.Now().UnixNano()
	l := len(a)
	if l < 1 {
		fmt.Println("数组不得为空")
		return
	}
	arr := make([]int,l)
	for i := 0;i < l;i++ {
		arr[i] = a[i]
	}
	var s int
    for i := 1;i < l;i++{
		s = arr[i]
		j := i
		for ;j > 0 && arr[j-1] > s;j--{
			arr[j] = arr[j-1]
		}
		arr[j] = s
	}
	end := time.Now().UnixNano()
	fmt.Println("插入排序耗时：",(end - start)/1e6,"ms")
	//fmt.Println("排序后：",arr)
}

//冒泡排序
func BubbleSort(a []int) {
	start := time.Now().UnixNano()
	l := len(a)
	if l < 1 {
		fmt.Println("数组不得为空")
		return
	}
	arr := make([]int, l)
	for i := 0; i < l; i++ {
		arr[i] = a[i]
	}
	var s int
	sort := true
    for i := l-1;i > 0 && sort;i-- {
		sort = false
        for j := 0;j < i;j++{
			if arr[j] > arr[j+1]{
				s = arr[j+1]
				arr[j+1] = arr[j]
				arr[j] = s
				sort = true
			}
		}
	}
	end := time.Now().UnixNano()
	fmt.Println("冒泡排序耗时：",(end - start)/1e6,"ms")
}
//归并排序入口
func MergeSortEntry(a []int){
	start := time.Now().UnixNano()
	l := len(a)
	if l < 1 {
		fmt.Println("数组不得为空")
		return
	}
	arr := make([]int, l)
	for i := 0; i < l; i++ {
		arr[i] = a[i]
	}
	MergeSort(arr,0,l)
	end := time.Now().UnixNano()
	fmt.Println("递归归并排序耗时：",(end - start)/1e6,"ms")
	//fmt.Println("排序后：",arr)
}
//递归进行归并排序
//空间换时间，内存占用极大
func MergeSort(arr []int,s int,e int){
/*	if (e - s) < 3 {
		InsertSortboundary(arr,s,e)
		return
	}*/
	if (e - s) < 2 {
		return
	}
	m := (s+e)/2
	MergeSort(arr,s,m)
	MergeSort(arr,m,e)
    mergeArray(arr,s,m,e)
}

func mergeArray(arr []int,s,m,e int){
	if arr[m] >= arr[m-1]{
		return
	}
	l := e - s
	p := make([]int,l)
	i := m
	j := s
	for k := 0;k < l;k++{
		if j >= m {
			p[k] = arr[i]
			i++
		}else if i >= e{
			p[k] = arr[j]
			j++
		}else if arr[j] < arr[i] {
			p[k] = arr[j]
			j++
		}else {
			p[k] = arr[i]
			i++
		}
	}
	for _,v := range p{
		arr[s] = v
		s++
	}
}

//对部分数组插入排序
func InsertSortboundary(arr []int,s int,e int){
	l := e -s
	if l < 2 {
		return
	}
	var t int
	for ;s < e;s++{
		t = arr[s]
		j := s
		for ;j > 0 && arr[j-1] > t;j--{
			arr[j] = arr[j-1]
		}
		arr[j] = t
	}
}

//自底向上归并排序入口
//只迭代不递归
func MergeSortBuEntry(a []int){
	start := time.Now().UnixNano()
	l := len(a)
	if l < 1 {
		fmt.Println("数组不得为空")
		return
	}
	arr := make([]int, l)
	for i := 0; i < l; i++ {
		arr[i] = a[i]
	}
	MergeSortBU(arr)
	end := time.Now().UnixNano()
	fmt.Println("迭代归并排序耗时：",(end - start)/1e6,"ms")
	//fmt.Println("排序后：",arr)
}
//自底向上迭代归并
func MergeSortBU(arr []int){
	l := len(arr)
	var e,m int
    for step := 1;step < l;step += step{
        for i := 0;i + step < l;i += step + step{
			m = i + step
			if arr[m] > arr[m-1]{
				continue
			}
			e = i + step + step
			if e > l {
				e = l
			}
			mergeArray(arr,i,m,e)
		}
	}
}

//快速排序入口
func QuikSortEntry(a []int){
	start := time.Now().UnixNano()
	l := len(a)
	if l < 1 {
		fmt.Println("数组不得为空")
		return
	}
	arr := make([]int, l)
	for i := 0; i < l; i++ {
		arr[i] = a[i]
	}
	quikSort(arr,0,l)
	end := time.Now().UnixNano()
	fmt.Println("快速排序耗时：",(end - start)/1e6,"ms")
	//fmt.Println("排序后：",arr)
}

//快速排序
func quikSort(arr []int,s,e int){
	if e - s < 14 {
		InsertSortboundary(arr,s,e)
		return
	}
	p := patition(arr,s,e)
	quikSort(arr,s,p)
	quikSort(arr,p+1,e)
}

func patition(arr []int, s int, e int) int{
	m := (s+e)/2
	c := arr[m]
	arr[m] = arr[s]
	arr[s] = c
	j := s
	var t int
	for i := s;i < e;i++{
		if arr[i] < c {
			j++
			t = arr[i]
			arr[i] = arr[j]
			arr[j] = t
		}
	}
	arr[s] = arr[j]
	arr[j] = c
	return j
}

//双路快速排序入口
func QuikSortEntry2(a []int){
	start := time.Now().UnixNano()
	l := len(a)
	if l < 1 {
		fmt.Println("数组不得为空")
		return
	}
	arr := make([]int, l)
	for i := 0; i < l; i++ {
		arr[i] = a[i]
	}
	quikSort2(arr,0,l)
	end := time.Now().UnixNano()
	fmt.Println("双路快速排序耗时：",(end - start)/1e6,"ms")
	//fmt.Println("排序后：",arr)
}


//快速排序,对于大量重复值得优化
func quikSort2(arr []int,s,e int){
	if e - s < 14 {
		InsertSortboundary(arr,s,e)
		return
	}
/*	if e - s < 2 {
		return
	}*/
	p := patition2(arr,s,e)
	quikSort2(arr,s,p)
	quikSort2(arr,p+1,e)
}

//均分等于c的值
func patition2(arr []int, s int, e int) int{
	m := (s+e)/2
	c := arr[m]
	arr[m] = arr[s]
	arr[s] = c
	j := s+1
	k := e-1
	var t int
	for {
		for j < e && arr[j] < c {
			j++
		}
		for k > s && arr[k] > c {
			k--
		}
		if j > k {
			break
		}
		t = arr[k]
		arr[k] = arr[j]
		arr[j] = t
		k--
		j++
	}
	arr[s] = arr[k]
	arr[k] = c
	return k
}

//三路快速排序入口
func QuikSortEntry3(a []int){
	start := time.Now().UnixNano()
	l := len(a)
	if l < 1 {
		fmt.Println("数组不得为空")
		return
	}
	arr := make([]int, l)
	for i := 0; i < l; i++ {
		arr[i] = a[i]
	}
	quikSort3(arr,0,l)
	end := time.Now().UnixNano()
	fmt.Println("三路快速排序耗时：",(end - start)/1e6,"ms")
	//fmt.Println("排序后：",arr)
}


//快速排序,对于大量重复值得优化
//三路快速排序
func quikSort3(arr []int,s,e int){
	if e - s < 14 {
		InsertSortboundary(arr,s,e)
		return
	}
/*	if e - s < 2 {
		return
	}*/
	lt,gt := patition3(arr,s,e)
	quikSort3(arr,s,lt)
	quikSort3(arr,gt,e)
}

func patition3(arr []int, s int, e int) (lt,gt int){
	m := (s+e)/2
	c := arr[m]
	arr[m] = arr[s]
	arr[s] = c
	lt = s
	gt = e
	var t int
	for i := s+1;i < gt;{
		if arr[i] < c {
			lt++
            t = arr[i]
			arr[i] = arr[lt]
			arr[lt] = t
			i++
		}else if arr[i] > c {
			gt--
            t = arr[i]
			arr[i] = arr[gt]
			arr[gt] = t
		}else {
			i++
		}
	}
	arr[s] = arr[lt]
	arr[lt] = c
	return
}