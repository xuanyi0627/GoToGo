package main

import (
	"fmt"
	"errors"
)

func variableName()  {

	fmt.Println("------------------------------- variableName")

	//定义一个名称为“variableName”，类型为"type"的变量
	//var variableName type
	var a int

	fmt.Println(a)

	//定义三个类型都是“type”的变量
	var b, c, d string

	fmt.Println(b, c, d)

	//初始化“variableName”的变量为“value”值，类型是“type”
	//var variableName type = value
	var e int = 2

	fmt.Println(e)

	/*
    定义三个类型都是"type"的变量,并且分别初始化为相应的值
    vname1为v1，vname2为v2，vname3为v3
	*/
	//var vname1, vname2, vname3 type= v1, v2, v3
	var o, p, q int = 1, 2, 3

	fmt.Println(o, p, q)

	/*
    定义三个变量，它们分别初始化为相应的值
    vname1为v1，vname2为v2，vname3为v3
    然后Go会根据其相应值的类型来帮你初始化它们
	*/
	//var vname1, vname2, vname3 = v1, v2, v3
	var r, s, t = 1, 3, 5

	fmt.Println(r, s, t)

	/*
    定义三个变量，它们分别初始化为相应的值
    vname1为v1，vname2为v2，vname3为v3
    编译器会根据初始化的值自动推导出相应的类型
	*/
	//vname1, vname2, vname3 := v1, v2, v3
	u, v, w := 1, 3, 5

	fmt.Println(u, v, w)

	//_（下划线）是个特殊的变量名，任何赋予它的值都会被丢弃。在这个例子中，我们将值35赋予m，并同时丢弃34
	_, m := 34, 35

	fmt.Println(m)

	//:=这个符号直接取代了var和type,这种形式叫做简短声明。不过它有一个限制，那就是它只能用在函数内部；在函数外部使用则会无法编译通过，所以一般用var方式来定义全局变量。
}

func constantName()  {

	fmt.Println("------------------------------- constantName")

	//const constantName = value
	//如果需要，也可以明确指定常量的类型：
	const PI float32 = 3.1415926

	fmt.Println(PI)
}

var isActive bool  // 全局变量声明
var enabled, disabled = true, false  // 忽略类型的声明

func baseType() {

	fmt.Println("------------------------------- baseType")

	var available bool  // 一般声明
	available = true    // 赋值操作

	valid := false      // 简短声明

	fmt.Println(available, valid)
}

var frenchHello string  // 声明变量为字符串的一般方法
var emptyString string = ""  // 声明了一个字符串变量，初始化为空字符串

func stringType()  {

	fmt.Println("------------------------------- stringType")

	no, yes, maybe := "no", "yes", "maybe"  // 简短声明，同时声明多个变量
	japaneseHello := "Konichiwa"  // 同上
	frenchHello = "Bonjour"  // 常规赋值

	fmt.Println(no, yes, maybe, japaneseHello, frenchHello)

	s := "hello"
	c := []byte(s)
	c[0] = 'c'
	s2 := string(c)

	fmt.Println(s2)

	m :=  " world"
	a := s + m

	fmt.Println(a)

	s = "C" + s[1:]

	fmt.Println(s)

	mu := `hello 
           world
          `

	fmt.Println(mu)
}

func errorType()  {
	fmt.Println("------------------------------- errorType")

	err := errors.New("emit macho dwarf: elf header corrupted")
	if err != nil {
		fmt.Println(err)
	}
}

const (
	x = iota // x == 0
	y = iota // y == 1
	z = iota // z == 2
	w        // 常量声明省略值时，默认和之前一个值的字面相同。这里隐式地说w = iota，因此w == 3。其实上面y和z可同样不用"= iota"
)

const  v = iota // 每遇到一个const关键字，iota就会重置，此时v == 0
const (
	h, i, j = iota, iota, iota  //h=0,i=0,j=0 iota在同一行值相同
)

const (
	a       = iota //a=0
	b       = "B"
	c       = iota             //c=2
	d, e, f = iota, iota, iota //d=3,e=3,f=3
	g       = iota             //g = 4
)

func arrayType()  {

	fmt.Println("------------------------------- arrayType")

	var arr [10]int
	arr[0] = 42
	arr[1] = 13

	fmt.Printf("The first element is %d\n", arr[0])  // 获取数据，返回42
	fmt.Printf("The last element is %d\n", arr[9]) //返回未赋值的最后一个元素，默认返回0

	a := [3]int{1, 2, 3} // 声明了一个长度为3的int数组

	b := [10]int{1, 2, 3} // 声明了一个长度为10的int数组，其中前三个元素初始化为1、2、3，其它默认为0

	c := [...]int{4, 5, 6} // 可以省略长度而采用`...`的方式，Go会自动根据元素个数来计算长度

	fmt.Printf("a %d\n", a)
	fmt.Printf("b %d\n", b)
	fmt.Printf("c %d\n", c)

	// 声明了一个二维数组，该数组以两个数组作为元素，其中每个数组中又有4个int类型的元素
	doubleArray := [2][4]int{[4]int{1, 2, 3, 4}, [4]int{5, 6, 7, 8}}

	// 上面的声明可以简化，直接忽略内部的类型
	easyArray := [2][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}}

	fmt.Printf("doubleArray %d\n", doubleArray)
	fmt.Printf("easyArray %d\n", easyArray)
}

func sliceType()  {

	fmt.Println("------------------------------- sliceType")

	var fslice []int
	slice := []byte{'a', 'b', 'c', 'd'}

	fmt.Printf("fslice %d\n", fslice)
	fmt.Printf("slice %c\n", slice)

	// 声明一个含有10个元素元素类型为byte的数组
	var ar = [10]byte {'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}

	fmt.Printf("ar %d\n", ar)

	var a, b []byte

	// a指向数组的第3个元素开始，并到第五个元素结束，
	a = ar[2:5]
	//现在a含有的元素: ar[2]、ar[3]和ar[4]

	// b是数组ar的另一个slice
	b = ar[3:5]
	// b的元素是：ar[3]和ar[4]


	fmt.Printf("a %c\n", a)
	fmt.Printf("b %c\n", b)


	// 声明一个数组
	var array = [10]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}
	// 声明两个slice
	var aSlice, bSlice []byte

	// 演示一些简便操作
	aSlice = array[:3] // 等价于aSlice = array[0:3] aSlice包含元素: a,b,c

	fmt.Printf("aSlice %c\n", aSlice)

	aSlice = array[5:] // 等价于aSlice = array[5:10] aSlice包含元素: f,g,h,i,j

	fmt.Printf("aSlice %c\n", aSlice)

	aSlice = array[:]  // 等价于aSlice = array[0:10] 这样aSlice包含了全部的元素

	fmt.Printf("aSlice %c\n", aSlice)

	// 从slice中获取slice
	aSlice = array[3:7]  // aSlice包含元素: d,e,f,g，len=4，cap=7

	fmt.Printf("aSlice %c\n", aSlice)

	bSlice = aSlice[1:3] // bSlice 包含aSlice[1], aSlice[2] 也就是含有: e,f

	fmt.Printf("bSlice %c\n", bSlice)

	bSlice = aSlice[:3]  // bSlice 包含 aSlice[0], aSlice[1], aSlice[2] 也就是含有: d,e,f

	fmt.Printf("bSlice %c\n", bSlice)

	bSlice = aSlice[0:5] // 对slice的slice可以在cap范围内扩展，此时bSlice包含：d,e,f,g,h

	fmt.Printf("bSlice %c\n", bSlice)

	bSlice = aSlice[:]   // bSlice包含所有aSlice的元素: d,e,f,g

	fmt.Printf("bSlice %c\n", bSlice)
}

func mapType()  {

	fmt.Println("------------------------------- mapType")

	//var number map[string]int

	numbers := make(map[string]int)

	numbers["one"] = 1  //赋值
	numbers["ten"] = 10 //赋值
	numbers["three"] = 3

	fmt.Println("第三个数字是: ", numbers["three"]) // 读取数据

	// 初始化一个字典
	rating := map[string]float32{"C":5, "Go":4.5, "Python":4.5, "C++":2 }

	// map有两个返回值，第二个返回值，如果不存在key，那么ok为false，如果存在ok为true
	csharpRating, ok := rating["C#"]

	if ok {
		fmt.Println("C# is in the map and its rating is ", csharpRating)
	} else {
		fmt.Println("We have no rating associated with C# in the map")
	}

	delete(rating, "C")  // 删除key为C的元素

	fmt.Println(rating)

	m := make(map[string]string)
	m["Hello"] = "Bonjour"
	m1 := m
	m1["Hello"] = "Salut"  // 现在m["hello"]的值已经是Salut了
}

func main() {

	variableName()
	constantName()
	baseType()
	stringType()
	errorType()

	fmt.Println(a, b, c, d, e, f, g, h, i, j, x, y, z, w, v)

	arrayType()
	sliceType()
	mapType()
}
