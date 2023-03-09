# 数据结构
## 字符串 string
1. 访问字符串中的值 
    * 通过下标访问
   ```go
    s1 := "hello world"
    first := s[0]
   ```

    * 通过切片访问
    ```go
    s2 := []byte(s1)
    first := s2[0]
   ```
   
    * 通过for-range循环访问
    ```go
    for i, v := range s1 {
        fmt.Println(i, v)
    }
    ```
   
2.  查询字符是否属于特定字符集
   ```go
       // 判断字符串中是否包含a、b、c中的任意一个字符
       strings.ContainsAny(s1, "abc")
       // 判断字符串中是否包含abc子串
       strings.Contains(s1, "abc")
       // 判断字符串中是否包含a字符
       strings.ContainsRune(s1, 'a')
   ```

3. 比较两个字符串
   * ==
    ```go
    if s1 == s2 {
        fmt.Println("s1 == s2")
    }
    ```
    * strings.Compare, 1大于，0相等，-1小于
     ```go
    if strings.Compare(s1, s2) == 0 {
        fmt.Println("s1 == s2")
    }
    ```
   * strings.EqualFold, 忽略大小写
    ```go
    if strings.EqualFold(s1, s2) {
        fmt.Println("s1 == s2")
    }
    ```
4. 字符串拼接
    * 加号
     ```go
     s3 := s1 + s2
     ```
    * strings.Join
     ```go
     s3 := strings.Join([]string{s1, s2}, "")
     ```
    * 高效拼接字符串
        ```go
        var buffer bytes.Buffer
        buffer.WriteString(s1)
        buffer.WriteString(s2)
        s3 := buffer.String()
        ```
        或者
        ```go
        var builder strings.Builder
        buffer.WriteString(s1)
        buffer.WriteString(s2)
        s3 := buffer.String()
        ```
## slice模拟stack
1. 创建栈
    ```go
    stack := make([]int, 0)
    ```
2. 入栈
    ```go
    stack = append(stack, 1)
    ```
3. 出栈
    ```go
    if len(stack) > 0 {
        stack = stack[:len(stack)-1]
    }
    ```
4. 判断栈是否为空
    ```go
    if len(stack) == 0 {
        fmt.Println("stack is empty")
    }
    ```
## slice模拟Queue
1. 创建队列
    ```go
    queue := make([]int, 0)
    ```
2. 入队
    ```go
    queue = append(queue, 1)
    ```
3. 出队
    ```go
    if len(queue) > 0 {
        queue = queue[1:]
    }
    ```
4. 判断队列是否为空
    ```go
    if len(queue) == 0 {
        fmt.Println("queue is empty")
    }
    ```
## slice模拟Set
1. 创建集合
    ```go
    set := make(map[byte]struct{})
    ```
2. 添加元素
    ```go
    set['a'] = struct{}{}
    ```
3. 删除元素
    ```go
    delete(set, 'a')
    ```
4. 判断元素是否存在
    ```go
    if _, ok := set['a']; ok {
        fmt.Println("a is in set")
    }
    ```
