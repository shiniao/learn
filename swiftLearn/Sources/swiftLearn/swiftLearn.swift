

// 函数，参数类似 python 中的注解
// 调用函数的时候，必须包含参数
// hello(name: "chaojie")
func hello(name: String) -> String {
    // let 声明一个常量
    let hello = "Hello"
    // var 声明为变量
    var deci = 1080
    // swift不支持隐式转换，必须使用显式转换
    var testS = "length: " + String(deci)
    // 可选的，表示 None 或者某个值
    var optionalTest: String? = "Hello"


    // ---------------- String ---------------------- //

    // 使用反斜线，在 str 中包含变量，类似 python 中 f-string
    let result = "Oh, \(hello) \(name)"
    // 多行str："""
    let multiStr = """
    Test multiStr \(result)
    """

    // 数组，类似 python
    // 空数组
    var emptyList = [String]()
    var empty2List = []
    var testList = ["apple", "banana", "dog"]
    // 数组添加新元素
    testList.append("cat")



    // 字典
    // 空字典
    var emptyDict = [String: float]()
    var empty2Dict = [:]
    // 注意，字典也是中括号
    vat testDict = ["test1":1, "test2":2]


    // 控制语句
    var rangeList = [1,2,3,4,5]
    // for ... in ...
    for r in rangeList {
        // if 语句的条件必须是bool 表达式
        // 比如 if r {...} 就是错的
        if r > 3 {
            print("more")
        }else {
            print("less")
        }
    }

    // for...in...迭代字典
    for (_, number) in testDict{
        print(number)
    }

    // for ...in range 类似 python 中 range
    for i in 0..<4{
        print(i)
    }

    // while
    while condition {}

    // switch case

}

print(hello(name:"World"))


// --------------------类和对象-----------------------//
class MyClass{
    // 类变量
    var name: String
    var sideLength: Double = 100

    // 初始化
    init(name: String) {
        self.name = name
    }

    // getter and setter
    var perimeter: Double {
        get {
            return 3.0 * sideLength
        }
        set {
            sideLength = newValue / 3.0
        }
    }

    func classFunc(parameters) -> String {
        return "test class \(name)"
    }
}

// 继承类
class MyNextClass: MyClass{
    
    init(name: String) {
        super.init(name: name)
    }

    // 覆盖父类函数
    override func classFunc(parameters) -> String {
         return "my next class"
    }
    
}

// -------------------枚举------------------//
enum Rank {

    case ace = 1
    case two, three, four, five, six, seven

    // 枚举中也可以有方法
    func simpleDescription() -> String {
        // 注意这里是 switch self
        switch self {
        // 自定义值的返回
        case .ace:
            return "ace"
        default:
        // 默认返回原生值，原生值默认从 0 开始
        return String(self.rawValue)
        }
    }
}

//------------------struct---------------//
// struct 和 class 类似，但是struct 的调用总是通过复制，
// 而 class 的调用是引用
struct Card {

    var rank: Rank

    func simpleDescription() -> String {
        return "The \(rank.simpleDescription())"
    }
}

// 调用 
// 枚举的调用使用. 来代表某个元素
Card(rank: .three)

//------------------------协议和扩展-----------------------//
// 协议 protocol 类似 interface
// class、enum、struct 都可以实现 protocol

//------------------------Error-------------------------//
// 自定义 error：实现 Error接口即可
enum MyError: Error{
    case outOfPaper
    case noToner
    case onFire
}

// 抛出 error：throw，类似 python 中的raise
// handle error
// 使用 do...catch
// 在可能抛出异常的代码前使用 try
do {
    let printerResponse = try send(job: 1040, toPrinter: "Bi Sheng")
    print(printerResponse)
} catch {
    print(error)
}

