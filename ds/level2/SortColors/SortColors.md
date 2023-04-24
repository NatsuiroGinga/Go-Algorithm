# LeetCode 75. Sort Colors
链接: https://leetcode-cn.com/problems/sort-colors/
## 题目描述
给定一个包含 n 个颜色为红色、白色或蓝色对象的数组编号，对它们进行就地排序，以便相同颜色的对象相邻，颜色按红色、白色和蓝色的顺序排列。

用整数 0、1 和 2 分别表示红色、白色和蓝色。

## 思路
1. 使用三个指针，`red`、`white`、`blue`，分别指向红色、白色、蓝色。
2. 初始化时，`red`指向第一个元素，`white`指向第一个元素，`blue`指向最后一个元素。
3. 移动`white`指针，当`white`指向的元素为红色时，交换`red`和`white`指向的元素，`red`和`white`指针都向后移动一位；当`white`指向的元素为白色时，`white`指针向后移动一位；当`white`指向的元素为蓝色时，交换`white`和`blue`指向的元素，`blue`指针向前移动一位。
4. 当`white`指针大于`blue`指针时，结束循环。

## 代码
```cpp
class Solution {
public:
    void sortColors(vector<int>& nums) {
        auto red = nums.begin(), white = nums.begin(), blue = nums.end() - 1;
        while (white <= blue) {
            if (*white == 1) white++;
            else if (*white == 0) swap(*white++, *red++);
            else swap(*white, *blue--);
        }    
   }
};
```