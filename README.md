# 语法

# 一、常量
iota，特殊常量，可以认为是一个可以被编译器修改的常量。
iota 在 const关键字出现时将被重置为 0(const 内部的第一行之前)，const 中每新增一行常量声明将使 iota 计数一次**(iota 可理解为 const 语句块中的行索引)。**
iota 可以被用作枚举值：
```go
const （
	a=inta
）
```
# 二、GO运算符
和c基本一样
&取地址
*指针

## 三、条件控制语句
### 3.1 if
```go
func main() {
   /* 定义局部变量 */
   var a int = 10
 
   /* 使用 if 语句判断布尔表达式 */
   if a < 20 {
       /* 如果条件为 true 则执行以下语句 */
       fmt.Printf("a 小于 20\n" )
   }
   fmt.Printf("a 的值为 : %d\n", a)
}
```
### 3.2 switch
```go
 /* 定义局部变量 */
   var grade string = "B"
   var marks int = 90

   switch marks {
      case 90: grade = "A"
      case 80: grade = "B"
      case 50,60,70 : grade = "C"
      default: grade = "D"  
   }

   switch {
      case grade == "A" :
         fmt.Printf("优秀!\n" )    
      case grade == "B", grade == "C" :
         fmt.Printf("良好\n" )      
      case grade == "D" :
         fmt.Printf("及格\n" )      
      case grade == "F":
         fmt.Printf("不及格\n" )
      default:
         fmt.Printf("差\n" );
   }
   fmt.Printf("你的等级是 %s\n", grade );     
```
### 3.3 select
select 是 Go 中的一个控制结构，类似于用于通信的 switch 语句。每个 case 必须是一个通信操作，要么是发送要么是接收。
select 随机执行一个可运行的 case**。如果没有 case 可运行，它将阻塞**，直到有 case 可运行。一个默认的子句应该总是可运行的。
```go
select {
    case communication clause  :
       statement(s);      
    case communication clause  :
       statement(s);
    /* 你可以定义任意数量的 case */
    default : /* 可选 */
       statement(s);
}

```

- 每个 case 都必须是一个通信
- 所有 channel 表达式都会被求值
- 所有被发送的表达式都会被求值
- 如果任意某个通信可以进行，它就执行，其他被忽略。
- 如果有多个 case 都可以运行，Select 会随机公平地选出一个执行。其他不会执行。否则：
   1. 如果有 default 子句，则执行该语句。
   1. 如果没有 default 子句，select 将阻塞，直到某个通信可以运行；Go 不会重新对 channel 或值进行求值。
# 四、函数定义
```go
func function_name( [parameter list] ) [return_types] {
   函数体
}
```

- func：函数由 func 开始声明
- function_name：函数名称，函数名和参数列表一起构成了函数签名。
- parameter list：参数列表，参数就像一个占位符，当函数被调用时，你可以将值传递给参数，这个值被称为实际参数。参数列表指定的是参数类型、顺序、及参数个数。参数是可选的，也就是说函数也可以不包含参数。
- return_types：返回类型，函数返回一列值。return_types 是该列值的数据类型。有些功能不需要返回值，这种情况下 return_types 不是必须的。
- 函数体：函数定义的代码集合。



```go
/* 函数返回两个数的最大值 */
func max(num1, num2 int) int {
   /* 声明局部变量 */
   var result int

   if (num1 > num2) {
      result = num1
   } else {
      result = num2
   }
   return result
}
```
## 五、数组
```go
var variable_name [SIZE] variable_type
```
初始化数组中 {} 中的元素个数不能大于 [] 中的数字。
如果忽略 [] 中的数字不设置数组大小，Go 语言会根据元素的个数来设置数组的大小：
```go
 var balance = [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
```
## 六、指针
```go
var ip *int
var fp *float32
```
eg
```go
package main

import "fmt"

func main() {
   var a int= 20   /* 声明实际变量 */
   var ip *int        /* 声明指针变量 */

   ip = &a  /* 指针变量的存储地址 */

   fmt.Printf("a 变量的地址是: %x\n", &a  )

   /* 指针变量的存储地址 */
   fmt.Printf("ip 变量储存的指针地址: %x\n", ip )

   /* 使用指针访问值 */
   fmt.Printf("*ip 变量的值: %d\n", *ip )
}

```
空指针判断
```go
if(ptr !=nil) 
if(ptr == nil)
```
指向指针的指针
```go
var ptr **int;
```
# 七、结构体
```go
type struct_variable_type struct {
   member definition
   member definition
   ...
   member definition
}
```
```go
package main

import "fmt"

type Books struct {
   title string
   author string
   subject string
   book_id int
}


func main() {

    // 创建一个新的结构体
    fmt.Println(Books{"Go 语言", "www.runoob.com", "Go 语言教程", 6495407})

    // 也可以使用 key => value 格式
    fmt.Println(Books{title: "Go 语言", author: "www.runoob.com", subject: "Go 语言教程", book_id: 6495407})

    // 忽略的字段为 0 或 空
   fmt.Println(Books{title: "Go 语言", author: "www.runoob.com"})
}
```
# 八、切片
Go 数组的长度不可改变，在特定场景中这样的集合就不太适用，Go中提供了一种灵活，功能强悍的内置类型切片("动态数组"),与数组相比切片的长度是不固定的，可以追加元素，在追加时可能使切片的容量增大。

声明一个未指定大小的数组来定义切片
```go
var identifier []type
```
切片不需要说明长度。
或使用make()函数来创建切片:
```go
var slice1 []type = make([]type, len)

也可以简写为

slice1 := make([]type, len)
```
也可以指定容量，其中capacity为可选参数。

```go
make([]T, length, capacity)
```
这里 len 是数组的长度并且也是切片的初始长度。




```go
s :=[] int {1,2,3 } 
直接初始化切片，[]表示是切片类型，{1,2,3}初始化值依次是1,2,3.其cap=len=3

s := arr[:] 
初始化切片s,是数组arr的引用

s := arr[startIndex:endIndex] 
将arr中从下标startIndex到endIndex-1 下的元素创建为一个新的切片

s := arr[startIndex:] 
默认 endIndex 时将表示一直到arr的最后一个元素

s := arr[:endIndex] 
默认 startIndex 时将表示从arr的第一个元素开始

s1 := s[startIndex:endIndex] 
通过切片s初始化切片s1

s :=make([]int,len,cap) 
```
## append() 和 copy() 函数
如果想增加切片的容量，我们必须创建一个新的更大的切片并把原分片的内容都拷贝过来。
下面的代码描述了从拷贝切片的 copy 方法和向切片追加新元素的 append 方法。
ss := []string{"武汉","啊啊啊啊"}
s1 = append(s1, ss...)
... 表示拆开


# 九、range
Go 语言中 range 关键字用于 for 循环中迭代数组(array)、切片(slice)、通道(channel)或集合(map)的元素。在数组和切片中它返回元素的索引和索引对应的值，在集合中返回 key-value 对。

```go
package main
import "fmt"
func main() {
    //这是我们使用range去求一个slice的和。使用数组跟这个很类似
    nums := []int{2, 3, 4}
    sum := 0
    for _, num := range nums {
        sum += num
    }
    fmt.Println("sum:", sum)
    //在数组上使用range将传入index和值两个变量。上面那个例子我们不需要使用该元素的序号，所以我们使用空白符"_"省略了。有时侯我们确实需要知道它的索引。
    for i, num := range nums {
        if num == 3 {
            fmt.Println("index:", i)
        }
    }
    //range也可以用在map的键值对上。
    kvs := map[string]string{"a": "apple", "b": "banana"}
    for k, v := range kvs {
        fmt.Printf("%s -> %s\n", k, v)
    }
    //range也可以用来枚举Unicode字符串。第一个参数是字符的索引，第二个是字符（Unicode的值）本身。
    for i, c := range "go" {
        fmt.Println(i, c)
    }
}
```
# 十、Map
```go
/* 声明变量，默认 map 是 nil */
var map_variable map[key_data_type]value_data_type

/* 使用 make 函数 */
map_variable := make(map[key_data_type]value_data_type)
```
如果不初始化 map，那么就会创建一个 nil map。nil map 不能用来存放键值对

delete() 函数用于删除集合的元素, 参数为 map 和其对应的 key。实例如下：
delete(countryCapitalMap, "France") 
入参 map,key
# 十一、类型转化
类型转换用于将一种数据类型的变量转换为另外一种类型的变量
type_name(expression)
```go
func main() {
   var sum int = 17
   var count int = 5
   var mean float32
   
   mean = float32(sum)/float32(count)
   fmt.Printf("mean 的值为: %f\n",mean)
}
```
# 十二、接口
Go 语言提供了另外一种数据类型即接口，它把所有的具有共性的方法定义在一起，任何其他类型只要实现了这些方法就是实现了这个接口。

## init函数
每一个源文件都可以包含一个init函数在main函数执行前，被Go运行框架调用时


如何一个文件同时包含全局变量定义，init函数 和main函数，则执行的流程**全局变量定义>init函数>main函数**
**
## 匿名函数
### 1 定义匿名函数的时候直接调用，这样匿名函数只能调用一次  
![image.png](https://cdn.nlark.com/yuque/0/2020/png/1497025/1605798448818-900025e2-f0b4-42ba-adb9-4121e903407d.png#align=left&display=inline&height=129&margin=%5Bobject%20Object%5D&name=image.png&originHeight=258&originWidth=679&size=145118&status=done&style=none&width=339.5)
### 2 将匿名函数赋给一个变量（函数变量），再通过变量来调用匿名函数


![image.png](https://cdn.nlark.com/yuque/0/2020/png/1497025/1605798468718-3d9bbe0f-c9d5-4cf4-a2d6-72f1ea043509.png#align=left&display=inline&height=124&margin=%5Bobject%20Object%5D&name=image.png&originHeight=247&originWidth=640&size=147819&status=done&style=none&width=320)
## 3 如果将匿名函数赋给一个全局变量，那么这个匿名函数，就成为一个全局匿名函数，可以在程序有效。


# 十三、大小写区别
如果一个结构体是大写开头，那么这个结构体就可以被其他包使用
如果一个结构体的属性是大写开头，那么这个结构体就是可以被其他包引用


结构体字段=属性=field
# 十四、结构体类型使用细节
指针，slice和map的零值都是nil，即还没有分配空间
使用这样的字段，需要先make


- 结构体是值类型，不是引用类型

两个结构体  取等号是值拷贝的过程


```go
声明方法：
方式1
p2 := Person{"mary",20}
fmt.Println(p2)
方式2
var person Persion = Person{}

方式3
var p3 *Person = new(Person)
因为p3是一个指针，因此标准的给字段赋值的方式
(*p3).Name = "smith" 也可以这样写 p3.Name = "smith"

//原因：go的设计者为了程序员使用方便，底层会对p3.Name = "smith"
进行处理
会给p3加上 取值运算 (*p3).Name = "smith"

方式4 {}
var person *Person = &Person{}

因为person是一个指针，因此标准的访问字段的方法

(*person).Name = "scott"
person.Name = "scott~~"


```
第三第四种方式返回的是 **结构体指针**
**结构体指针访问字段的标准方式应该是：(*结构体指针).字段名，比如（*person）.Name = "tom"**
**但go做了一个简化，也支持结构体指针.字段名 **
**go编译器底层 对person.Name 做了转化 (*person).Name**
**
## 结构体内存
结构体的所有字段在内存中是连续的
可以在每个字段上写上一个tag 
该tag可以通过反射机制获取，常见的使用场景就是序列化和反序列化
用 json.Marshal()


# 十五、方法声明和调用
```go
type A struct{
	Num int
}
func (a A)test(){
    fmt.Println(a.Num)
}
表面A结构体有一方法，方法名为test
(a A)体现是和A类型绑定的


一个具体的例子
给Person结构体jisuan 该方法可以接受一个数n，计算从1+..+n的结果
func (p Person) jisuan (n int){
    res :=0
    for i :=1;i <=n ; i++{
    	res += i
    }
    fmt.PrintLn(p.Name,"计算的结果是="，res)
}


```
# 十六、方法的调用和传参机制原来（重要！）
方法的调用和传参机制和函数基本一样，不一样的地方是方法调用时，会将调用方法的变量，当做实参也传递给方法。

对于普通函数，接受者为值类型时，不能将指针类型的数据直接传递，反之亦然
func test01(p Person){
fmt.Println(p.Name)
}


func test02(p *Person){
fmt.Println(p.Name)
}
func main(){
p := Person{"tom"}
test01(p)
test02(&p) 
}
# 十七、工厂模式解决结构体构造问题
如果model包的结构体变量首字母大写，引入后，直接使用，没有问题


如果model包的结构体变量首字母小写，引入后是不能直接使用，所以可以用工程模式解决


```go
//定义一个结构体
type student struct{
	Name string
    Score float64
}

func NewStudent(n String,s float64) *student{
    return &student{
        Name : n,
        Score :s
    }
}



在main.go中

func main(){
    var stu = model.NewStudent("free",88.8)
    fmt.PrintLb(*stu)
    fmt.Printlb("name=",stu.Name,"score=",stu.Score)
}
```


java中的private 属性的 set和get方法
```go
//定义一个结构体
type student struct{
	Name string
    score float64
}

func NewStudent(n String,s float64) *student{
    return &student{
        Name : n,
        Score :s
    }
}
func (s *student)GetScore() float64{
	return s.score
}


main.go中
func main(){
    var stu = model.NewStudent("tom~",98.8)
    
    fmt.Println(*stu)
    fmt.Println("name=",stu.Name,"score=",stu.GeScore())
}
```
# 十八、继承如何实现？
struct嵌套struct，用匿名结构体
```go
type Goods struct{
	Name string
    Price int
}
type Book struct{
	Goods //这里就是嵌套匿名结构体Goods
    Writer String
}
```
当我们直接通过一个东西来访问字段或方法，其执行流程如下
b.Name
b.A.Name
编译器会先看b对应的类型有没有Name，如果有，则直接调用B类型的Name字段


如果没有就去看B中嵌入的匿名结构体A有没有什么Name字段，如果有就调用，如果没有就继续查找，如果什么都找不到就报错


当结构体和匿名结构体有 相同的字段或者方法时，编译器会就近访问，如希望访问匿名结构体的字段和方法，可以通过匿名结构体来区分


如果当结构体嵌入两个或多个匿名结构体，如两个匿名结构体有相同的字段和方法（同时结构体本身没有同名的字段和方法），访问时，就必须明确指定匿名结构体名字，否则就便以报错。


**如果一个struct嵌套了一个有名的结构体，这种模式就是组合**
如果是组合关系，那么在访问组合的结构体的字段或方法时，必选带上结构体的名字


嵌套匿名结构体后，也可以在创建结构体变量（实例）时，直接指定各个匿名结构体字段的值
```go
type TV2 struct{
	*Goods
    *Brand
}

tv2 := TV2{
    &Goods{
        Price : 5000.99
        Name : "电视剧002"
    },
    &Brand{
        Name : "夏普"，
        Address :"北京"
    }
}



或者是

tv4 := TV2{&Goods{"电视机003"，7000},&Brand{"创维"，"河南"}}



```
如果一个结构体有一个int类型的匿名字段，就不能第二个
如果需要多个int的字段，则必须给int字段指定名字


# 十九、接口
```go
type Usb interface{
    Start()
    Stop()
}
type Phone struct{
}

//让Phone实现 Usb接口的方法
func (p Phone)Start(){
    fmt.Println("手机开始工作...")
}
func(p Phone)Stop(){
    fmt.Println("手机停止工作...")
}

实现Usb接口（所谓实现Usb接口，就是指实现了Usb接口声明所有方法）


编写一个Working 方法，接收一个Usb接口类型变量
func (c Computer)Working(usb Usb){
	//通过usb接口变量来diaoongStart和Stop方法
    usb.Start()
    usb.Stop()
    //usb变量会根据传入的实参，来判断到底是Phone还是Camera
}


```


```go
type Usb interface{
    Say()
}
type Stu struct {
}
func (this *Stu) Say(){
    fmt.Println("Say()")
}

func main(){
    var stu Stu = Stu{}
    var u Usb = &stu
    u.Say()
    fmt.Println("here",u)
}
```
# 二十、类型断言
```go
type Point struct{
	x int
    y int
}

func main(){
    var a interface{}
    var point Point =Point{1,2}
    a = point
    
}
```
类型断言，由于接口是一般类型，不知道具体类型，如果要转成具体类型，就需要使用类型断言
```go
var x inteeface{}
var b2 float32 = 1.1
x = b2  //空接口，可以接收任意类型
y := x.(float32)

对进行类型断言时，如果类型不匹配，就会报panic，因此进行类型断言时，要确保原来的空接口只想的就是断言的类型

```
# 二十一、文件操作
判断文件是否存在
os.Stat():
返回错误为nil 说明文件 或文件夹存在
如果返回错误类型使用os.isNotExist()判断为true，说明文件或文件夹不存在


# 二十二、json
data,err := json.Marshal(&。。。)


# 二十三、测试框架
go语言自带了一个轻量级测试框架testing 和自带的go test 命令来执行单元测试


1. 测试用例文件名必须以 _test.go 结尾  比如 cal_test.go  cal 不是固定的
1. 测试用例函数必须以Test开头，一般来说就是Test加被测试的函数名
1. TestAddUpper(t *testing T) 的形参类型必须是  *testing.
1. 一个测试用例文件中，可以有呕多个测试用例函数，比如TestAddUpper TestSub 
1. 运行测试用例指令 go test   go test -v





# 二十四、goroutine和channel


GO主线程：一个Go线程上，可以起多个协程，携程可以理解为轻量级线程


## GO协程的特点

1. 有独立的栈空间
1. 共享程序堆空间
1. 调度由用户控制
1. 轻量级线程





![image.png](https://cdn.nlark.com/yuque/0/2020/png/1497025/1606290248219-400b0b59-36ba-4ac1-9c11-2452dcde4805.png#align=left&display=inline&height=293&margin=%5Bobject%20Object%5D&name=image.png&originHeight=586&originWidth=705&size=281476&status=done&style=none&width=352.5)
快速入门小结

1. 主线程是一个物理线程，直接作用在cpu上的。是重量级的，非常耗费cpu资源
1. 协程从主线程开启的，是轻量级的线程，是逻辑态的。对资源消耗相对小
1. Goloang的协程机制是重要的特点，可以轻松的开启上万个协程。



## goroutine的调度模型
### MPG模型基本介绍






# ！channel 

1. channel 本质上就是一个数据结构-队列
1. 数据是先进先出
1. 线程安全，多goroutine访问时，不需要加锁，就是说channel本身就是线程安全的
1. channel有类型的，一个string的channel只能存放string类型数据





说明：

1. channel是引用类型
1. channel必须初始化才能写入数据，即make后才能使用
1. 管道是有类型，intChan只能写入 整数int





## 从管道中拿结构体 
需要先做断言判断
比如
```go
newCat := <-allChan
fmt.Printf("newCat = %T","newCat=%v\n",newCat,newCat)

a := newCat.(Cat)
先断言判断了 然后再去获取属性

```
## channel的遍历与关闭
使用内置函数close可以关闭channel，当channel关闭之后，就不能再向channel写数据了，但是可以继续读数据
```go
intChan := make(make int,3)
int
```
在遍历时，如果channel没有关闭，会出现deadlock的错误
在遍历时，如果channel已经关闭，就会正常遍历数组，便利完之后，就会退出遍历




## channel阻塞
如果编译器（运行），发现一个管道只有写，而没有读，则该管道会阻塞。


默认情况下 管道是双向的
var chan1 chan int 


但是若是 var **chan2 chan<**- int 则是只写
若是 var **chan3 <-chan** int 只写
## select可以解决从那个管道取数据的阻塞问题
# 二十六、反射
```go
reflect.TypeOf(变量名)，获取变量的类型，返回reflect.type类型
reflect.Value
```
变量、interface{}、reflect.Value



