# 快慢指针(Fast-slow Pointers)

## 1. 概念介绍
* 快慢指针是一种常用的技巧，用于解决链表中的问题。
* 快慢指针的思想是：两个指针以不同的速度遍历链表，从而达到目的。
* 快慢指针的常见应用：
    * 判断链表是否有环
    * 寻找链表的中点
    * 寻找链表的倒数第k个节点
    * 寻找链表的交点
    * 寻找链表的入环点
    * 计算链表的环的长度

## 2. 典型例题
### 2.1 判断链表是否有环
* 题目描述：给定一个单向链表的头指针，判断链表中是否有环。
* 解题思路：快慢指针，如果有环，快指针(fast)一定会追上慢指针(slow)。
因为当 fast 从后面接近 slow 时, 有两种可能性: 
  * fast 比 slow 慢一步
  * fast 比 slow 慢兩步
  
  其他情况的都可以归纳为这两种情况, 又因为 fast 先移动, 所以:
    * fast 比 slow 慢一步: fast 走了两步, slow 走了一步, fast 与 slow 相遇
    * fast 比 slow 慢两步: fast 走了两步, slow 走了一步, fast 与 slow 相差1步, 回到第一种情况

所以, 如果链表有环, 快指针一定会追上慢指针, 二者相遇。
* 代码实现：
* Java
```java
public static boolean hasCycle(ListNode head) {
    // 快慢指针,
    ListNode slow = head;
    ListNode fast = head;
    // 如果没有环, 快指针一定会先到达链表尾部
    while (fast != null && fast.next != null) { // 
      // 快指针每次走两步，慢指针每次走一步
      fast = fast.next.next;
      slow = slow.next;
      // 如果有环, 快指针一定会追上慢指针, 二者相遇
      if (slow == fast) return true;
    }
    
    return false;
  }
```
* Go
```go
func hasCycle(head *ListNode) bool {
    // 快慢指针,
    slow := head
    fast := head
    // 如果没有环, 快指针一定会先到达链表尾部
    for fast != nil && fast.Next != nil { // 
      // 快指针每次走两步，慢指针每次走一步
      fast = fast.Next.Next
      slow = slow.Next
      // 如果有环, 快指针一定会追上慢指针, 二者相遇
      if slow == fast {
          return true
      }
    }
    
    return false
}
```

### 2.2 寻找链表的中点
* 题目描述：给定一个单向链表的头指针，找到链表的中点。
* 解题思路：快慢指针，快指针每次走两步，慢指针每次走一步，使快指针走的距离始终是慢指针的两倍, 当快指针走到链表尾部时，慢指针正好走到链表中间。
* 代码实现：
* Java
```java
public static ListNode findMiddle(ListNode head) {
    ListNode slow = head;
    ListNode fast = head;
    
    while (fast != null && fast.next != null) {
      slow = slow.next;
      fast = fast.next.next;
    }
    
    return slow;
  }
```
* Go
```go
func findMiddle(head *ListNode) *ListNode {
    slow := head
    fast := head
    
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }
    
    return slow
}
```
