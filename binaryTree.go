package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
二分查找，返回目标值再数组中的下标
*/
func BinarySerach(arr []int, target int) int {
	//含头不含尾，左区间
	l := 0
	//右区间
	r := len(arr)
	mid := l + (r-l)/2
	for l < r {
		if target == arr[mid] {
			return mid
		}
		if arr[mid] > target {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return -1
}

/*
二分查找，返回目标值在数组中的下标
如果值重复，则返回最大下标值。
如果没有该值，返回大于该值的最小值下标
*/
func FloorBinarySerach(arr []int, target int) (i int) {
	//含头不含尾，左区间
	l := 0
	//右区间
	r := len(arr)
	mid := l + (r-l)/2
	for l < r {
		if arr[mid] == target {
			r = mid
			i = mid
		} else if arr[mid] >= target {
			r = mid
		} else {
			l = mid + 1
			i = mid
		}
		mid = l + (r-l)/2
	}
	return
}

/*
二分查找，返回目标值在数组中的下标
如果值重复，则返回最小下标值。
如果没有该值，返回小于该值的最大值下标
*/
func CeilBinarySerach(arr []int, target int) (i int) {
	//含头不含尾，左区间
	l := 0
	//右区间
	r := len(arr)
	mid := l + (r-l)/2
	for l < r {
		if arr[mid] == target {
			l = mid + 1
			i = mid
		} else if arr[mid] > target {
			r = mid
			i = mid
		} else {
			l = mid + 1
		}
		mid = l + (r-l)/2
	}
	return
}

type BinarySerachTreeNode struct {
	key       int
	value     interface{}
	leftNode  *BinarySerachTreeNode
	rightNode *BinarySerachTreeNode
	root      *BinarySerachTreeNode
	count     int
}

/*
返回节点个数
*/
func (bstn *BinarySerachTreeNode) GetBSTSize() int {
	return bstn.count
}

/*
二分搜索树是否为空
*/
func (bstn *BinarySerachTreeNode) BSTSIsEmpty() bool {
	return bstn.count == 0
}

/*
向二分搜索树种插入kv
*/
func (bstn *BinarySerachTreeNode) InsertBinarySerachTreeNodeByKV(k int, v interface{}) {
	insertBinarySerachTreeNodeByKV(bstn, k, v)
}

/*
向二分搜索树中查找value
*/
func (bstn *BinarySerachTreeNode) SerachBinarySerachTreeValueByKey(k int) interface{} {
	return serachBinarySerachTreeByValue(bstn, k)
}

func serachBinarySerachTreeByValue(node *BinarySerachTreeNode, k int) interface{} {
	if k == node.key {
		return node.value
	}
	var child *BinarySerachTreeNode
	if k > node.key {
		child = node.rightNode
	} else {
		child = node.leftNode
	}
	if child == nil {
		return nil
	} else {
		return serachBinarySerachTreeByValue(child, k)
	}
}

/*
向二分搜索树种插入一个节点
*/
func insertBinarySerachTreeNodeByKV(root *BinarySerachTreeNode, k int, v interface{}) {
	if root.GetBSTSize() == 0 {
		root.key = k
		root.value = v
		root.count++
		root.root = root
		return
	}
	node := newBinarySerachTreeNodeByKV(k, v)
	node.root = root
	//迭代插入二叉搜索树
	var parent *BinarySerachTreeNode
	for root != nil {
		if root.key == node.key {
			root.value = node.value
			return
		}
		parent = root
		if node.key > root.key {
			root = root.rightNode
		} else {
			root = root.leftNode
		}
	}
	node.root.count++
	if node.key > parent.key {
		parent.rightNode = node
	} else {
		parent.leftNode = node
	}
	//递归插入二叉搜索树
	//appendChileNode(root,node)
}

func appendChileNode(parent, child *BinarySerachTreeNode) {
	if parent.key == child.key {
		parent.value = child.value
		return
	}
	if parent.key < child.key {
		r := parent.rightNode
		if r == nil {
			child.root.count++
			parent.rightNode = child
		} else {
			appendChileNode(r, child)
		}
	} else {
		l := parent.leftNode
		if l == nil {
			child.root.count++
			parent.leftNode = child
		} else {
			appendChileNode(l, child)
		}
	}
}

func NewBinarySerachTreeNode() *BinarySerachTreeNode {
	return new(BinarySerachTreeNode)
}

func newBinarySerachTreeNodeByKV(k int, v interface{}) *BinarySerachTreeNode {
	return &BinarySerachTreeNode{key: k, value: v}
}

/*
返回节点的value
*/
func (bstn *BinarySerachTreeNode) GetBSTNValue() interface{} {
	return bstn.value
}

func main() {
	root := NewBinarySerachTreeNode()
	root.InsertBinarySerachTreeNodeByKV(1, 10)
	root.InsertBinarySerachTreeNodeByKV(2, 20)
	root.InsertBinarySerachTreeNodeByKV(3, 30)
	fmt.Println("key为3：", root.SerachBinarySerachTreeValueByKey(3))
	fmt.Println("size为：", root.GetBSTSize())
	root.InsertBinarySerachTreeNodeByKV(3, 40)
	fmt.Println("key为3：", root.SerachBinarySerachTreeValueByKey(3))
	fmt.Println("size为：", root.GetBSTSize())
	/*	l := 10
		arr := GetIntArry(l)
		quikSort3(arr,0,l)
		fmt.Println("arr：",arr)
		target := 5
		i := FloorBinarySerach(arr,target)
		fmt.Println("最小下标值：",i)
		i = CeilBinarySerach(arr,target)
		fmt.Println("最大下标值：",i)*/
}

func GetIntArry(l int) []int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	arr := make([]int, l)
	for i := 0; i < l; i++ {
		arr[i] = r.Intn(10)
	}
	return arr
}

//快速排序,对于大量重复值得优化
//三路快速排序
func quikSort3(arr []int, s, e int) {
	if e-s < 2 {
		return
	}
	lt, gt := patition3(arr, s, e)
	quikSort3(arr, s, lt)
	quikSort3(arr, gt, e)
}

func patition3(arr []int, s int, e int) (lt, gt int) {
	m := (s + e) / 2
	c := arr[m]
	arr[m] = arr[s]
	arr[s] = c
	lt = s
	gt = e
	var t int
	for i := s + 1; i < gt; {
		if arr[i] < c {
			lt++
			t = arr[i]
			arr[i] = arr[lt]
			arr[lt] = t
			i++
		} else if arr[i] > c {
			gt--
			t = arr[i]
			arr[i] = arr[gt]
			arr[gt] = t
		} else {
			i++
		}
	}
	arr[s] = arr[lt]
	arr[lt] = c
	return
}
