package main

import "fmt"

type unionFind struct {
	id []int
	rank []int //以i为根的元素的层数
	count int
}

func NewUnionFind(n int) *unionFind{
	id := make([]int,n)
	rank := make([]int,n)
	for i := 0;i < n;i++ {
		id[i] = i
		rank[i] = 1
	}
	uf := &unionFind{id,rank,n}
    return uf
}

func (uf *unionFind)Find(p int)int{
	if p > -1 && p < uf.count {
		return uf.id[p]
	}else {
		return -1
	}
}

func getRoot(uf *unionFind,p int)int{
	pId := uf.id[p]
	if p != pId {
		//递归路径压缩
		uf.id[p] = getRoot(uf,pId)
	}
	/*for p != pId {
		//父节点不是根节点，则指向父节点的父节点
		//路径压缩
		uf.id[p] = uf.id[pId]
		p = pId
		pId = uf.id[p]
	}*/
	return uf.id[p]
}

/*
p,q是否相连
 */
func (uf *unionFind)IsConnext(p,q int)bool{
	return getRoot(uf,p) == getRoot(uf,q)
}
/*
合并两个元素
 */
func (uf *unionFind)UnionRootElements(p,q int){
	pId := getRoot(uf,p)
	pRank := uf.rank[p]
	qId := getRoot(uf,q)
	qRank := uf.rank[q]
	if pId == qId {
		return
	}
	if pRank > qRank {
		uf.id[qId] = uf.id[pId]
	}else if pRank < qRank{
		uf.id[pId] = uf.id[qId]
	}else {
		uf.id[pId] = uf.id[qId]
		uf.rank[qId]++
	}
}

/*
func (uf *unionFind)UnionElements(p,q int){
	pId := uf.Find(p)
	qId := uf.Find(q)
	if pId == qId {
		return
	}
	for i := 0;i < uf.count;i++ {
		if uf.Find(i) == pId {
			uf.id[i] = qId
		}
	}
}*/

func main(){
	union := NewUnionFind(10)
	union.UnionRootElements(1,2)
	union.UnionRootElements(2,3)
	fmt.Println(union.rank[2])
}