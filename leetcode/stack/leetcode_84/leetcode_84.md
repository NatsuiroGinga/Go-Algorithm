# LeetCode 84. Largest Rectangle in Histogram

## 题目

Given an array of integers `heights` representing the histogram's bar height where the width of each bar is `1`, return *the area of the largest rectangle in the histogram*.

## 翻译

给定一个整数数组`heights`，表示柱状图的高度，每个柱子的宽度为`1`，返回柱状图中最大矩形的面积。

## 解析

这道题目的解法是单调栈，单调栈的思路是，维护一个单调递减的栈，当遇到比栈顶元素小的元素时，就计算栈顶元素的面积，然后弹出栈顶元素，直到栈顶元素大于等于当前元素，然后将当前元素入栈。

**选用单调栈的原因**

这道题目的解法是单调栈，这里介绍一下为什么选用单调栈。
因为计算柱状图中最大矩形的面积，需要计算每个柱子的左边界和右边界，然后计算面积，这里的左边界和右边界都是指比当前柱子小的柱子，所以可以使用单调栈来计算左边界和右边界。

**计算右边界**

计算右边界的时候，需要从左往右遍历，遍历到的柱子如果比栈顶元素小，当前柱子就是栈顶元素的右边界，然后弹出栈顶元素，直到栈顶元素小于等于当前元素，然后将当前元素入栈。

**计算右边界**

因为使用的是递减栈，所以栈中的元素是单调递减的，所以栈中的元素的左边界就是栈顶元素的下一个元素，如果栈为空，说明当前元素是最小的元素，左边界就是-1。

**计算面积**

计算面积的时候，需要计算当前柱子的高度和左边界和右边界的距离，然后计算面积，最后取最大值。

## 代码

### Java

```java

class Solution {
    public int largestRectangleArea(int[] heights) {
        int n = heights.length;
        int[] left = new int[n];
        int[] right = new int[n];
        Deque<Integer> stack = new LinkedList<>();
        for (int i = 0; i < n; i++) {
            while (!stack.isEmpty() && heights[stack.peek()] >= heights[i]) {
                stack.pop();
            }
            left[i] = stack.isEmpty() ? -1 : stack.peek();
            stack.push(i);
        }
        stack.clear();
        for (int i = n - 1; i >= 0; i--) {
            while (!stack.isEmpty() && heights[stack.peek()] >= heights[i]) {
                stack.pop();
            }
            right[i] = stack.isEmpty() ? n : stack.peek();
            stack.push(i);
        }
        int ans = 0;
        for (int i = 0; i < n; i++) {
            ans = Math.max(ans, (right[i] - left[i] - 1) * heights[i]);
        }
        return ans;
    }
}

```

### C++

```cpp

class Solution {
public:
    int largestRectangleArea(vector<int>& heights) {
        int n = heights.size();
        vector<int> left(n), right(n, n);
        stack<int> mono_stack;
        for (int i = 0; i < n; ++i) {
            while (!mono_stack.empty() && heights[mono_stack.top()] >= heights[i]) {
                right[mono_stack.top()] = i;
                mono_stack.pop();
            }
            left[i] = (mono_stack.empty() ? -1 : mono_stack.top());
            mono_stack.push(i);
        }
        int ans = 0;
        for (int i = 0; i < n; ++i) {
            ans = max(ans, (right[i] - left[i] - 1) * heights[i]);
        }
        return ans;
    }
};

```