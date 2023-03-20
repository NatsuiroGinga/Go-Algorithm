# LeetCode 33. Search in Rotated Sorted Array
## 题目描述
There is an integer array nums sorted in ascending order (with distinct values).

Prior to being passed to your function, nums is possibly rotated at an unknown pivot index k (1 <= k < nums.length) such that the resulting array is [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]] (0-indexed). For example, [0,1,2,4,5,6,7] might be rotated at pivot index 3 and become [4,5,6,7,0,1,2].

Given the array nums after the possible rotation and an integer target, return the index of target if it is in nums, or -1 if it is not in nums.

You must write an algorithm with O(log n) runtime complexity.

 ## Example 1:
```
Input: nums = [4,5,6,7,0,1,2], target = 0
Output: 4
```
## Example 2:
```
Input: nums = [4,5,6,7,0,1,2], target = 3
Output: -1
```
## Example 3:
```
Input: nums = [1], target = 0
Output: -1
```
## 思路
这道题目是二分查找的变种，因为数组是有序的，所以可以使用二分查找，但是这道题目的数组是旋转过的，所以需要先找到原数组的起始位置，然后再进行二分查找。

原数组的起始位置可以通过二分查找的方式找到, 原数组的起始位置就是数组中最小的元素的位置。

以数组[4,5,6,7,0,1,2]为例，原数组的起始位置就是0的位置, 而0也是数组中最小的元素。

**那么如何找到起始的位置呢？**

可以通过二分查找找到数组中间的元素，如果中间元素大于数组的最后一个元素，那么说明原数组的起始位置在中间元素的右边，如果中间元素小于数组的最后元素，那么说明原数组的起始位置在中间元素的左边或者就是中间元素。

以数组[4,5,6,7,0,1,2]为例，中间元素是7，7大于数组的最后一个元素2，所以原数组的起始位置在7的右边，也就是0的位置。

再以数组[6,7,0,1,2,4,5]为例，中间元素是0，0小于数组的最后一个元素5，所以原数组的起始位置在0的左边或者就是0。这里就是0。

**如何找到目标元素的位置呢？**

我们需要先将原数组扩展, 比如数组[4,5,6,7,0,1,2]扩展为[4,5,6,7,0,1,2,4,5,6,7]，其实就是把原数组的起始位置之前的元素复制到数组的后面。

在扩展之后的数组里使用二分查找, 必须将原数组的下标转换为原数组的下标，因为旋转后的数组不是递增的。

假设数组元素个数为`n`, 数组的起始位置为`rot`, 数组的最后一位为`high`, 数组的第一位为`low`, 
那么旋转后的数组的中间元素位置是`mid = (low + high) / 2`, 真正的中间元素位置是`realMid = (mid + rot) % n`。

可以这样理解, 在[4,5,6,7,0,1,2,4,5,6,7]里, `realMid = (rot + (hi + rot) / 2) % n`, 也就是` (rot + hi/2) % n`, 实际上是`(rot + mid) % n`。

## 代码
```java
class Solution {
    public int search(int[] nums, int target) {
        int n = nums.length;
        if (n == 0) {
            return -1;
        }
        if (n == 1) {
            return nums[0] == target ? 0 : -1;
        }
        int low = 0;
        int high = n - 1;
        while (low < high) {
            int mid = (low + high) / 2;
            if (nums[mid] > nums[high]) {
                low = mid + 1;
            } else {
                high = mid;
            }
        }
        int rot = low;
        low = 0;
        high = n - 1;
        while (low <= high) {
            int mid = (low + high) / 2;
            int realMid = (mid + rot) % n;
            if (nums[realMid] == target) {
                return realMid;
            }
            if (nums[realMid] < target) {
                low = mid + 1;
            } else {
                high = mid - 1;
            }
        }
        return -1;
    }
}
```