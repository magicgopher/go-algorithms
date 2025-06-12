package binary_tree

import (
	"reflect"
	"testing"
)

// TestInsertAndSearch 测试插入节点与查找功能
func TestInsertAndSearch(t *testing.T) {
	tree := &BinaryTree{}

	// 插入一些节点
	tree.Insert(10)
	tree.Insert(5)
	tree.Insert(15)
	tree.Insert(3)

	// 查找存在的值
	if !tree.Search(10) || !tree.Search(5) || !tree.Search(3) {
		t.Errorf("应能找到插入过的值")
	}

	// 查找不存在的值
	if tree.Search(100) {
		t.Errorf("不应找到未插入的值")
	}
}

// TestTraversal 测试三种遍历顺序
func TestTraversal(t *testing.T) {
	tree := &BinaryTree{}
	// 插入构建如下结构：
	//        10
	//       /  \
	//      5    15
	//     / \
	//    3   7
	tree.Insert(10)
	tree.Insert(5)
	tree.Insert(15)
	tree.Insert(3)
	tree.Insert(7)

	// 中序遍历应为排序后的数组
	expectedInOrder := []int{3, 5, 7, 10, 15}
	inOrder := tree.InOrder()
	if !reflect.DeepEqual(inOrder, expectedInOrder) {
		t.Errorf("中序遍历错误，期望 %v，实际 %v", expectedInOrder, inOrder)
	}

	// 前序遍历
	expectedPreOrder := []int{10, 5, 3, 7, 15}
	preOrder := tree.PreOrder()
	if !reflect.DeepEqual(preOrder, expectedPreOrder) {
		t.Errorf("前序遍历错误，期望 %v，实际 %v", expectedPreOrder, preOrder)
	}

	// 后序遍历
	expectedPostOrder := []int{3, 7, 5, 15, 10}
	postOrder := tree.PostOrder()
	if !reflect.DeepEqual(postOrder, expectedPostOrder) {
		t.Errorf("后序遍历错误，期望 %v，实际 %v", expectedPostOrder, postOrder)
	}
}
