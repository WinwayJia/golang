package main

type MyInterface interface {
	Print()
	Hello()
	World()
	AWK()
}

func Foo(me MyInterface) {
	me.Print()
	me.Hello()
	me.World()
	me.AWK()

}

type MyStruct struct{}

func (me MyStruct) Print() {}
func (me MyStruct) Hello() {}
func (me MyStruct) World() {}
func (me MyStruct) AWK()   {}

func main() {
	var me MyStruct
	Foo(me)
}
