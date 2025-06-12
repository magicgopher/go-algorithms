package binary_tree

import "fmt"

// 常用数据结构
// 二叉树

// Node 表示二叉树中的一个节点
type Node struct {
	Value int   // 节点的值
	Left  *Node // 左子节点指针
	Right *Node // 右子节点指针
}

// BinaryTree 表示一个二叉查找树（Binary Search Tree）
type BinaryTree struct {
	Root *Node // 树的根节点
}

// Insert 向二叉树中插入一个新值
func (t *BinaryTree) Insert(value int) {
	t.Root = insertNode(t.Root, value)
}

// insertNode 是递归插入函数，返回插入后的子树根节点
func insertNode(node *Node, value int) *Node {
	if node == nil {
		// 如果当前位置为空，创建一个新节点
		return &Node{Value: value}
	}
	if value < node.Value {
		// 小于当前节点值，插入左子树
		node.Left = insertNode(node.Left, value)
	} else if value > node.Value {
		// 大于当前节点值，插入右子树
		node.Right = insertNode(node.Right, value)
	}
	// 如果值相等，不插入（忽略重复值）
	return node
}

// Search 在二叉树中查找指定的值，返回是否存在
func (t *BinaryTree) Search(value int) bool {
	return searchNode(t.Root, value)
}

// searchNode 是递归查找函数
func searchNode(node *Node, value int) bool {
	if node == nil {
		return false // 到达空节点，未找到
	}
	if value == node.Value {
		return true // 找到值
	} else if value < node.Value {
		return searchNode(node.Left, value) // 在左子树查找
	} else {
		return searchNode(node.Right, value) // 在右子树查找
	}
}

// PreOrder 前序遍历（根-左-右），将结果写入切片
func (t *BinaryTree) PreOrder() []int {
	var result []int
	preOrderTraversal(t.Root, &result)
	return result
}

// preOrderTraversal 是递归的前序遍历函数
func preOrderTraversal(node *Node, result *[]int) {
	if node == nil {
		return
	}
	*result = append(*result, node.Value) // 先访问根
	preOrderTraversal(node.Left, result)  // 再左子树
	preOrderTraversal(node.Right, result) // 再右子树
}

// InOrder 中序遍历（左-根-右），将结果写入切片
func (t *BinaryTree) InOrder() []int {
	var result []int
	inOrderTraversal(t.Root, &result)
	return result
}

func inOrderTraversal(node *Node, result *[]int) {
	if node == nil {
		return
	}
	inOrderTraversal(node.Left, result)
	*result = append(*result, node.Value)
	inOrderTraversal(node.Right, result)
}

// PostOrder 后序遍历（左-右-根），将结果写入切片
func (t *BinaryTree) PostOrder() []int {
	var result []int
	postOrderTraversal(t.Root, &result)
	return result
}

func postOrderTraversal(node *Node, result *[]int) {
	if node == nil {
		return
	}
	postOrderTraversal(node.Left, result)
	postOrderTraversal(node.Right, result)
	*result = append(*result, node.Value)
}

// Print 用于调试打印整棵树（中序输出）
func (t *BinaryTree) Print() {
	printInOrder(t.Root)
	fmt.Println()
}

func printInOrder(node *Node) {
	if node == nil {
		return
	}
	printInOrder(node.Left)
	fmt.Print(node.Value, " ")
	printInOrder(node.Right)
}
