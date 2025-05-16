1.byte 是 uint8 的内置别名，可以把 byte 和 uint8 视为同一种类型。
byte 是类型 uint8 的别名，用于存放 1 字节的 ASCII 字符，比如英文字符、数字等，返回的是字符的原始字节。
rune 是类型 int32 的别名，用于存放多字节字符，比如占 3 字节的中文字符，返回的是 Unicode 码点值。

2.公开的属性是首字母大写，非公开的属性首字母是小写，仅按照这个规则来定义是否公开

3.iota 定义在 const 定义组中的第 n 行，那么 iota 的值为 n - 1
   const (
       f = 2
       g = iota  // 1
       h         // 2
       i         // 3
   )
4.函数 闭包（匿名函数） 方法
    func <function_name>(<parameter list>) (<return types>) {
        // 函数体
    }
    
    // 声明函数变量
    var <closure name> func(<parameter list>) (<return types>)
    
    // 声明闭包
    var  <closure name> func(<parameter list>) (<return types>) = func(<parameter list>) (<return types>) {
    <expressions>
    ...
    }
    
    // 声明并立刻执行
    func(<parameter list>) (<return types>) {
    <expressions>
    ...
    }(<value list>)
    
    // 作为参数，并调用
    func <function name>(...,<name> func(<parameter list>) (<return types>), ...) {
    ...
    <var1>,... := <name>(<value list>)
    ...
    }
    
    // 作为返回值
    func <function name>(...) (func(<parameter list>) (<return types>)) {
    ...
    <var1>,... := <name>(<value list>)
    ...
    }
5.切片
切片触发扩容前，切片一直共用相同的数组；
切片触发扩容后，会创建新的数组，并复制这些数据；
切片本身是一个特殊的指针，go 针对切片类型添加了一些语法糖，方便使用。
