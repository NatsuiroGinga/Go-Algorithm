# LeetCode 116. Populating Next Right Pointers in Each Node
https://leetcode.com/problems/populating-next-right-pointers-in-each-node/description/?envType=study-plan&id=algorithm-i
## 题目描述

You are given a perfect binary tree where all leaves are on the same level, and every parent has two children. The binary tree has the following definition:

struct Node {
int val;
Node *left;
Node *right;
Node *next;
}
Populate each next pointer to point to its next right node. If there is no next right node, the next pointer should be set to NULL.

Initially, all next pointers are set to NULL.

## 解题思路 1
1. DFS
2. dfs函数的两个参数是当前节点和Next节点, main函数中调用参数为root和null
3. 如果当前节点为空，直接返回
4. 如果当前节点不为空，将当前节点的next指针指向Next节点
5. 当前节点的左子树的next指针指向当前节点的右子树
6. 如果当前节点的next指针的左子树不为空, 当前节点的右子树的next指针指向当前节点的next指针的左子树
7. 否则，当前节点的右子树的next指针指向null

## 算法 1
```go
func connect(root *Node) *Node {
    dfs(root, nil)
    return root
}

func dfs(cur, next *Node) {
	if cur == nil { // 如果当前节点为空，直接返回
		return
	}

	cur.Next = next // 如果当前节点不为空，将当前节点的next指针指向Next节点

	dfs(cur.Left, cur.Right) // 当前节点的左子树的next指针指向当前节点的右子树

	if cur.Next != nil { // 如果当前节点的next指针的左子树不为空
		// 当前节点的右子树的next指针指向当前节点的next指针的左子树
		dfs(cur.Right, cur.Next.Left)
	} else {
		// 否则，当前节点的右子树的next指针指向null
		dfs(cur.Right, nil)
	}
}
```
## 思路 2
1. BFS, 从右到左, 一层一层的遍历 
2. 用一个队列存储每一层的节点, 从右到左将每一层的节点入队
3. 每次循环初始化一个RightNode指针，指向null
4. 不断弹出队列中的节点，将当前节点的next指针指向RightNode指针, 并将RightNode指针指向当前节点
5. 如果当前节点不是最后一层, 也就是说节点有左右子树, 则先将当前节点的右子树入队, 再将当前节点的左子树入队

## 算法 2
```go
func connect(root *Node) *Node {
    bfs(root)
    return root
}

func bfs(root *Node) *Node {
	if root == nil { // 如果根节点为空，直接返回
		return nil
	}
	queue := make([]*Node, 0) // 用于存储每一层的节点
	queue = append(queue, root) // 将根节点入队

	for len(queue) > 0 { // 当队列不为空时，循环
		var rightNode *Node // 初始化一个RightNode指针，指向null

		for i := len(queue) - 1; i >= 0; i-- { // 从右到左遍历每一层的节点
			cur := queue[0] // 弹出队列中的节点
			queue = queue[1:] // 将弹出的节点从队列中删除
			cur.Next = rightNode // 将当前节点的next指针指向RightNode指针
			rightNode = cur // 将RightNode指针指向当前节点

			if cur.Right != nil { // 如果当前节点不是最后一层, 也就是说节点有左右子树, 则先将当前节点的右子树入队, 再将当前节点的左子树入队
				queue = append(queue, cur.Right)
				queue = append(queue, cur.Left)
			}
		}
	}

	return root
}

```