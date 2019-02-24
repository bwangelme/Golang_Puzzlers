package main

type Dog struct {
	name string
}

func New(name string) Dog {
	return Dog{name}
}

func (dog *Dog) SetName(name string) {
	dog.name = name
}

func (dog Dog) Name() string {
	return dog.name
}

func main() {
	// 示例1。
	// 函数返回值是临时结果，临时结果是不可取值的
	// Go 会自动进行转译，将 New("little pig").SetName("monster") 变成  &New("little pig").SetName("monster")
	// 因为临时结果无法取值，所以 &New("little pig").SetName("monster") 会失败
	// puzzlers/article15/q2/demo36.go:21:19: cannot call pointer method on New("little pig")
	// puzzlers/article15/q2/demo36.go:21:19: cannot take the address of New("little pig")
	// New("little pig").SetName("monster")

	// 示例2。
	// 自增或自减语句左边表达式的结果值必须是可寻址的，但是对于字典字面量值的索引表达式例外
	// puzzlers/article15/q2/demo36.go:30:17: cannot assign to [2]int literal[1]
	// [2]int{0, 1}[1]++
	map[string]int{"the": 0, "word": 0, "counter": 0}["word"]++
	map1 := map[string]int{"the": 0, "word": 0, "counter": 0}
	map1["word"]++
}
