/*
For the purposes of this question, a balanced tree is defined to be a tree
Implement a function to check if a binary tree is balanced.
such that the heights of the two subtrees of any node
never differ by more than one

https://leetcode.com/problems/balanced-binary-tree/
*/
package main

import (
	"fmt"
)

type TreeNode struct {
	Value int
	Left *TreeNode
	Right *TreeNode
}
