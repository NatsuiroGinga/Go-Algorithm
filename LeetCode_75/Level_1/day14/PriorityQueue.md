# Go语言实现堆(优先队列)
## 1. 堆的定义
堆是一种特殊的树，它的每个节点都有一个权值，权值越大的节点越靠近根节点。堆的定义如下：
- 堆是一棵完全二叉树
- 堆中每个节点的值都必须大于等于（或小于等于）其子树中每个节点的值
## 2. 堆的实现
### 2.1 堆的结构体
```go
type Heap struct {
    data []int
    size int
}
```
### 2.2 堆的初始化
```go
func NewHeap() *Heap {
    return &Heap{
        data: make([]int, 0),
        size: 0,
    }
}
```
### 2.3 堆的插入
```go
func (h *Heap) Insert(val int) {
    h.data = append(h.data, val)
    h.size++
    h.up(h.size - 1)
}
```
### 2.4 堆的上浮
```go
func (h *Heap) up(i int) {
    for {
        if i == 0 {
            break
        }
        p := (i - 1) / 2
        if h.data[p] >= h.data[i] {
            break
        }
        h.data[p], h.data[i] = h.data[i], h.data[p]
        i = p
    }
}
```
### 2.5 堆的删除
```go
func (h *Heap) Remove() int {
    if h.size == 0 {
        return -1
    }
    val := h.data[0]
    h.data[0] = h.data[h.size-1]
    h.data = h.data[:h.size-1]
    h.size--
    h.down(0)
    return val
}
```
### 2.6 堆的下沉
```go
func (h *Heap) down(i int) {
    for {
        a := 2*i + 1
        if a >= h.size {
            break
        }
        b := a + 1
        max := a
        if b < h.size && h.data[b] > h.data[a] {
            max = b
        }
        if h.data[i] >= h.data[max] {
            break
        }
        h.data[i], h.data[max] = h.data[max], h.data[i]
        i = max
    }
}
```
### 2.7 堆的打印
```go
func (h *Heap) Print() {
    fmt.Println(h.data)
}
```
### 2.8 堆的测试
```go
func TestHeap(t *testing.T) {
    h := NewHeap()
    h.Insert(1)
    h.Insert(2)
    h.Insert(3)
    h.Insert(4)
    h.Insert(5)
    h.Insert(6)
    h.Insert(7)
    h.Insert(8)
    h.Insert(9)
    h.Print()
    h.Remove()
    h.Print()
    h.Remove
    h.Print()
}
```
## 3. 堆的应用
### 3.1 堆排序
堆排序的思路是先将数组构建成一个大顶堆，然后将堆顶元素与最后一个元素交换，然后将剩下的元素重新构建成一个大顶堆，如此反复，直到剩下一个元素为止。
```go
func HeapSort(arr []int) {
    h := NewHeap()
    for _, v := range arr {
        h.Insert(v)
    }
    for i := len(arr) - 1; i >= 0; i-- {
        arr[i] = h.Remove()
    }
}
```
### 3.2 前K个高频元素
给定一个非空的整数数组，返回其中出现频率前 k 高的元素。
```go
func TopKFrequent(nums []int, k int) []int {
    m := make(map[int]int)
    for _, v := range nums {
        m[v]++
    }
    h := NewHeap()
    for k, v := range m {
        h.Insert(&Item{
            key:   k,
            value: v,
        })
    }
    res := make([]int, 0)
    for i := 0; i < k; i++ {
        res = append(res, h.Remove().(*Item).key)
    }
    return res
}
```
## 4. 使用container/heap包实现堆
```go
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
    *h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}

func TestHeap2(t *testing.T) {
    h := &IntHeap{2, 1, 5}
    heap.Init(h)
    heap.Push(h, 3)
    for h.Len() > 0 {
        fmt.Printf("%d ", heap.Pop(h))
    }
}
```
**打印结果**
```
1 2 3 5
```