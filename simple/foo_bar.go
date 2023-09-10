package simple

type Foo struct {}
type Bar struct {}

func NewFoo() *Foo {
	return &Foo{}
}

func NewBar() *Bar {
	return &Bar{}
}

type FooBar struct {
	*Foo
	*Bar
}

// tampa membuat contruktornya <-
// func NewFooBar(foo *Foo, bar *Bar) *FooBar {
// 	return &FooBar{
// 		Foo: foo,
// 		Bar: bar,
// 	}
// }