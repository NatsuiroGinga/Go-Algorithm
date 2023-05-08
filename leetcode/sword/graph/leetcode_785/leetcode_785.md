# LeetCode 785. 判断二分图

## 题目描述

给定一个无向图graph，当这个图为二分图时返回true。

如果我们能将一个图的节点集合分割成两个独立的子集A和B，并使图中的每一条边的两个节点一个来自A集合，一个来自B集合，我们就将这个图称为二分图。

graph将会以邻接表方式给出，graph[i]表示图中与节点i相连的所有节点。每个节点都是一个在0到graph.length-1之间的整数。这图中没有自环和平行边： graph[i] 中不存在i，并且graph[i]中没有重复的值。

示例 1:
```

输入: [[1,3], [0,2], [1,3], [0,2]]

输出: true

解释:

无向图如下:

0----1

|    |

|    |

3----2

我们可以将节点分成两组: {0, 2} 和 {1, 3}。

```

示例 2:
```

输入: [[1,2,3], [0,2], [0,1,3], [0,2]]

输出: false

解释:

无向图如下:

0----1

| \  |

|  \ |

3----2

我们不能将节点分割成两个独立的子集。

```

注意:

- graph 的长度范围为 [1, 100]。
- graph[i] 中的元素的范围为 [0, graph.length - 1]。
- graph[i] 不会包含 i 或者有重复的值。
- 图是无向的: 如果j 在 graph[i]里边, 那么 i 也会在 graph[j]里边。

## 解题思路

深度优先搜索, 染色法
1. 从任意一个节点开始, 将其染成红色, 并将其相邻的节点染成蓝色, 以此类推, 如果相邻的节点已经染色, 则判断其颜色是否与当前节点相同, 如果相同, 则返回false, 否则返回true
2. 如果当前节点的所有相邻节点都已经染色, 则返回true
3. 如果当前节点的所有相邻节点都已经染色, 且颜色不同, 则返回false

## 代码

**Java**

```java
class Solution {
    public boolean isBipartite(int[][] graph) {
        int[] colors = new int[graph.length];
        for (int i = 0; i < graph.length; i++) {
            if (colors[i] == 0 && !dfs(graph, colors, i, 1)) {
                return false;
            }
        }
        return true;
    }

    private boolean dfs(int[][] graph, int[] colors, int index, int color) {
        if (colors[index] != 0) {
            return colors[index] == color;
        }
        colors[index] = color;
        for (int i = 0; i < graph[index].length; i++) {
            if (!dfs(graph, colors, graph[index][i], -color)) {
                return false;
            }
        }
        return true;
    }
}
```

**Python3**

```python
class Solution:
    def isBipartite(self, graph: List[List[int]]) -> bool:
        colors = [0] * len(graph)
        for i in range(len(graph)):
            if colors[i] == 0 and not self.dfs(graph, colors, i, 1):
                return False
        return True

    def dfs(self, graph, colors, index, color):
        if colors[index] != 0:
            return colors[index] == color
        colors[index] = color
        for i in range(len(graph[index])):
            if not self.dfs(graph, colors, graph[index][i], -color):
                return False
        return True
```