# LeetCode42.  接雨水

## 题目描述

给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。

## 示例

```text

输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]

输出：6

解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。 

```

## 提示

- n == height.length
- 0 <= n <= 3 * 104
- 0 <= height[i] <= 105

## 解题思路

### 方法一：暴力法

对于数组中的每个元素，我们找出下雨后水能达到的最高位置，等于两边最大高度的最小值减去当前高度的值。

1. 初始化 ans=0
2. 从左向右扫描数组：
   1. 初始化 max_left=0 和 max_right=0
   2. 从当前元素向左扫描并更新：max_left=max(max_left,height[j])
   3. 从当前元素向右扫描并更新：max_right=max(max_right,height[j])
   4. 将 min(max_left,max_right)−height[i] 累加到 ans
3. 返回 ans

### 方法二：动态编程

在暴力法中，我们仅仅为了找到最大值每次都要向左和向右扫描一次。但是我们可以提前存储这个值。因此，可以通过动态编程解决。

1. 找到数组中从下标 i 到最左端最高的条形块高度 left_max。
2. 找到数组中从下标 i 到最右端最高的条形块高度 right_max。
3. 扫描数组 height 并更新答案：累加 min(max_left[i],max_right[i])−height[i] 到 ans 上
4. 返回 ans
5. 优化：在上面的基础上，我们可以不用额外的数组来存储 left_max 和 right_max 的值，我们可以使用两个变量来存储当前的最大值。

### 方法三：栈的应用

我们可以用栈来跟踪可能储水的最长的条形块。使用栈就可以在一次遍历内完成计算。

1. 使用栈来存储条形块的索引下标。
2. 遍历数组：
   1. 当栈非空且 height[current]>height[st.top()]
      1. 意味着栈中元素可以被弹出。弹出栈顶元素 top。
      2. 计算当前元素和栈顶元素的距离，准备进行填充操作
      3. distance=current−st.top()−1
      4. 找出界定高度
         1. bounded_height=min(height[current],height[st.top()])−height[top]
      5. 往答案中累加积水量 ans+=distance×bounded_height
   2. 将当前索引下标入栈
   3. 将 current 移动到下个位置
3. 返回 ans

### 方法四：双指针

和方法二相比，我们不从左和从右分开计算，我们想办法一次完成遍历。

1. 初始化 left 指针为 0 并且 right 指针为 size-1
2. While left<right, do:
   1. If height[left] < height[right]
      1. If height[left]≥left_max, 更新 left_max
      2. Else 累加 left_max−height[left] 到 ans
      3. left = left + 1.
   2. Else
      1. If height[right]≥right_max, 更新 right_max
      2. Else 累加 right_max−height[right] 到 ans
      3. right = right - 1.
3. 返回 ans

## 代码

### Java

```java

class Solution {
    public int trap(int[] height) {
        int ans = 0;
        int size = height.length;
        for (int i = 1; i < size - 1; i++) {
            int max_left = 0, max_right = 0;
            for (int j = i; j >= 0; j--) { //Search the left part for max bar size
                max_left = Math.max(max_left, height[j]);
            }
            for (int j = i; j < size; j++) { //Search the right part for max bar size
                max_right = Math.max(max_right, height[j]);
            }
            ans += Math.min(max_left, max_right) - height[i];
        }
        return ans;
    }
}

```

### C++

```cpp

class Solution {
public:
    int trap(vector<int>& height) {
        int ans = 0;
        int size = height.size();
        for (int i = 1; i < size - 1; i++) {
            int max_left = 0, max_right = 0;
            for (int j = i; j >= 0; j--) { //Search the left part for max bar size
                max_left = max(max_left, height[j]);
            }
            for (int j = i; j < size; j++) { //Search the right part for max bar size
                max_right = max(max_right, height[j]);
            }
            ans += min(max_left, max_right) - height[i];
        }
        return ans;
    }
};

```