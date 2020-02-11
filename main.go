package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)


type treeNode struct {
	data string
	key int
	leftChild *treeNode
	rightChild *treeNode
}

type tree struct {

	root *treeNode
	size int
	height int

}

func treeHeight(t *treeNode) int {
	if t == nil{
		return 0
	} else {
		leftH := treeHeight(t.leftChild)
		rightH := treeHeight(t.rightChild)
		return int(1+ math.Max(float64(leftH),float64(rightH)))
	}

}

func treeSize(t *treeNode) int {
	if t == nil{
		return 0
	} else {
		leftSize := treeSize(t.leftChild)
		rightSize := treeSize(t.rightChild)
		return 1 + leftSize + rightSize
	}
}

func (t *tree)put(key int,data string)  {
	if t.root == nil{
		t.root = &treeNode{
			data:       data,
			key:        key,
			leftChild:  nil,
			rightChild: nil,
		}
	} else {
		_put(key,data,t.root)
	}
}

func midOrder(n *treeNode) {

	if n != nil{
		midOrder(n.leftChild)
		fmt.Println(n.key)
		midOrder(n.rightChild)
	}

}

func _put(key int ,data string,node *treeNode) {
	if key < node.key{
		if node.leftChild != nil{
			_put(key,data,node.leftChild)
		} else {
			node.leftChild = &treeNode{
				data:       data,
				key:        key,
				leftChild:  nil,
				rightChild: nil,
			}
		}
	} else {
		if node.rightChild != nil{
			_put(key,data,node.rightChild)
		} else {
			node.rightChild = &treeNode{
				data:       data,
				key:        key,
				leftChild:  nil,
				rightChild: nil,
			}
		}
	}

}


func main() {

	t := tree{}
	t.root = &treeNode{
		data:       "baiyun",
		key:        50,
		leftChild:  nil,
		rightChild: nil,
	}
	for i:=0;i<49;i++{
		a := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(100)
		t.put(a,"dyl")
	}
	t.height = treeHeight(t.root)
	t.size = treeSize(t.root)
	fmt.Println("中序遍历")
	midOrder(t.root)
	fmt.Println("根:",t.root.key)
	fmt.Println("节点数目:",t.size)
	fmt.Println("树高:",t.height)

}


