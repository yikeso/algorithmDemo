package main

import (
	"math/rand"
	"time"
	"fmt"
)

type testStruct struct {
	v int
}

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
	key        int
	value      interface{}
	leftNode   *BinarySerachTreeNode
	rightNode  *BinarySerachTreeNode
	root       *BinarySerachTreeNode
	count      int
	parentNode *BinarySerachTreeNode
}

/*
中序遍历
*/
func (bstn *BinarySerachTreeNode) BinarySerachTreeLDR() {
	binarySerachTreeLDR(bstn)
}

/*
中序遍历二叉搜索树
 */
func binarySerachTreeLDR(root *BinarySerachTreeNode) {
	node := root.leftNode
	if node != nil {
		binarySerachTreeLDR(node)
	}
	fmt.Println("当前key：", root.key)
	node = root.rightNode
	if node != nil {
		binarySerachTreeLDR(node)
	}
}

/*
中序遍历二叉搜索树
 */
func changeBinarySerachTreeRootNode(node, root *BinarySerachTreeNode) {
	child := node.leftNode
	if child != nil {
		changeBinarySerachTreeRootNode(child, root)
	}
	fmt.Println("当前修改root属性节点key：", node.key)
	node.root = root
	child = node.rightNode
	if child != nil {
		changeBinarySerachTreeRootNode(child, root)
	}
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

/*
判断二分搜索树中包含key
*/
func (bstn *BinarySerachTreeNode) BinarySerachTreeContainKey(k int) bool {
	return binarySerachTreeContainKey(bstn, k)
}

/*
删除二分搜索树中最大key
*/
func (bstn *BinarySerachTreeNode) BinarySerachTreeDeleteMaxnum() *BinarySerachTreeNode {
	return binarySerachTreeDeleteMaxnum(bstn)
}

func binarySerachTreeDeleteMaxnum(node *BinarySerachTreeNode) (root *BinarySerachTreeNode) {
	right := node.rightNode
	root = node
	for right != nil {
		node = right
		right = node.rightNode
	}
	//如果最小值是根节点
	if node.root.key == node.key {
		//修改所有节点的root属性为根节点的右节点
		//将根节点的右节点作为新的根节点
		left := node.leftNode
		if left == nil {
			return nil
		}
		left.parentNode = nil
		changeBinarySerachTreeRootNode(left, left)
		root = left
	} else {
		//将最小值节点 的父节点 的左节点 改为
		//最小值节点的右节点
		node.parentNode.rightNode = node.leftNode
		if node.leftNode != nil {
			node.leftNode.parentNode = node.parentNode
		}
	}
	return root
}

/*
删除二分搜索树中最小key
*/
func (bstn *BinarySerachTreeNode) BinarySerachTreeDeleteMinnum() *BinarySerachTreeNode {
	return binarySerachTreeDeleteMinnum(bstn)
}

func binarySerachTreeDeleteMinnum(node *BinarySerachTreeNode) (root *BinarySerachTreeNode) {
	left := node.leftNode
	root = node
	for left != nil {
		node = left
		left = node.leftNode
	}
	//如果最小值是根节点
	if node.root.key == node.key {
		//修改所有节点的root属性为根节点的右节点
		//将根节点的右节点作为新的根节点
		right := node.rightNode
		if right == nil {
			return nil
		}
		right.parentNode = nil
		changeBinarySerachTreeRootNode(right, right)
		root = right
	} else {
		//将最小值节点 的父节点 的左节点 改为
		//最小值节点的右节点
		node.parentNode.leftNode = node.rightNode
		if node.rightNode != nil {
			node.rightNode.parentNode = node.parentNode
		}
	}
	return root
}

/*
删除二叉搜索树中的任意节点
 */
func (bstn *BinarySerachTreeNode) BinarySerachTreeDeleteNode(key int) *BinarySerachTreeNode {
	node := serachBinarySerachTreeNodeByKey(bstn, key)
	if node == nil {
		return bstn
	}
	var min *BinarySerachTreeNode
	var child *BinarySerachTreeNode
	child = node.rightNode
	if node.root.key == key {
		if child == nil {
			child = node.leftNode
			if child == nil {
				return nil
			}
			child.parentNode = nil
			changeBinarySerachTreeRootNode(child, child)
			return child
		} else {
			min = binarySerachTreeDeleteMinnumReturnMinnumNode(child)
			changeBinarySerachTreeRootNode(bstn, min)
			min.leftNode = bstn.leftNode
			if bstn.leftNode != nil {
				bstn.leftNode.parentNode = min
			}
			if min.key != child.key {
				min.rightNode = bstn.rightNode
				if bstn.rightNode != nil {
					bstn.rightNode.parentNode = min
				}
			}
			min.parentNode = nil
			return min
		}
	} else {
		isLeft := true
		if node.parentNode.rightNode.key == node.key {
			isLeft = false
		}
		if child == nil {
			node.leftNode.parentNode = node.parentNode
			if isLeft {
				node.parentNode.leftNode = node.leftNode
			} else {
				node.parentNode.rightNode = node.leftNode
			}
		} else {
			min = binarySerachTreeDeleteMinnumReturnMinnumNode(child)
			min.parentNode = node.parentNode
			min.leftNode = node.leftNode
			node.leftNode.parentNode = min
			if min.key != child.key {
				min.rightNode = child
				child.parentNode = min
			}
			if isLeft {
				node.parentNode.leftNode = min
			} else {
				node.parentNode.rightNode = min
			}
		}
		return bstn
	}
}

func serachBinarySerachTreeNodeByKey(node *BinarySerachTreeNode, key int) (r *BinarySerachTreeNode) {
	if node == nil {
		return
	}
	if node.key == key {
		return node
	}
	if key > node.key {
		r = serachBinarySerachTreeNodeByKey(node.rightNode, key)
	} else {
		r = serachBinarySerachTreeNodeByKey(node.leftNode, key)
	}
	return r
}

/*
删除传入子树的最小节点，返回这个节点
 */
func binarySerachTreeDeleteMinnumReturnMinnumNode(node *BinarySerachTreeNode) *BinarySerachTreeNode {
	left := node.leftNode
	rootKey := node.key
	for left != nil {
		node = left
		left = node.leftNode
	}
	//将最小值节点 的父节点 的左节点 改为
	//最小值节点的右节点
	if node.key != rootKey {
		node.parentNode.leftNode = node.rightNode
		if node.rightNode != nil {
			node.rightNode.parentNode = node.parentNode
		}
	}
	return node
}

func binarySerachTreeContainKey(node *BinarySerachTreeNode, k int) bool {
	if k == node.key {
		return true
	}
	var child *BinarySerachTreeNode
	if k > node.key {
		child = node.rightNode
	} else {
		child = node.leftNode
	}
	if child == nil {
		return false
	} else {
		return binarySerachTreeContainKey(child, k)
	}
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
		node.parentNode = parent
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
	child.parentNode = parent
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
	var a, b, c, d *testStruct
	a = new(testStruct)
	b = new(testStruct)
	c = new(testStruct)
	d = new(testStruct)
	root.InsertBinarySerachTreeNodeByKV(2, a)
	root.InsertBinarySerachTreeNodeByKV(1, b)
	root.InsertBinarySerachTreeNodeByKV(3, c)
	root.InsertBinarySerachTreeNodeByKV(4, d)
	root = root.BinarySerachTreeDeleteNode(2)
	root.BinarySerachTreeLDR()
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
