package main

import "fmt"

type Human struct {
	name string
	sex  string
}
type SuperMan struct {
	Human // 类继承了Human类的方法
	level int
}

func (this *SuperMan) Eat() {
	fmt.Println("SuperMan.Eat()")
}

func (this *SuperMan) Fly() {
	fmt.Println("SuperMan.Fly()")
}
func (this *Human) Eat() {
	fmt.Println("Human.Eat()...")
}
func (this *Human) Walk() {
	fmt.Println("Human.Walk()...")
}
func main() {
	h := Human{"zhang3", "female"}
	h.Eat()
	h.Walk()
	// 定义一个子类对象
	// s := SuperMan{Human{"li4", "female"}, 88}
	var s SuperMan
	s.name = "li4"
	s.sex = "male"
	s.level = 88

	s.Walk() // 父类的方法
	s.Eat()  // 子类的方法（重写）
	s.Fly()  // 子类的方法
}
