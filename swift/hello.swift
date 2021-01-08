

// 函数，参数类似 python 中的注解
func hello(name: String) -> String {
    // let 声明一个常量
    let hello = "Hello"
    // var 声明为变量
    // var deci = 1080

    let result = "Oh, \(hello) \(name)"


    return result
}

print(hello(name:"World"))