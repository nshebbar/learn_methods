package main

import "fmt"

type myInt int

func (mi myInt) isEven() bool {
	return mi%2 == 0
}

func (mi *myInt) Double() {
	*mi = *mi * 2
}


type Foo struct {
	A int
	B string
}

func (f Foo) fieldCount() int {
	return 2
}

func (f Foo) String () string {
	return fmt.Sprintf("%d fields: A: %d and B: %s", f.fieldCount(), f.A, f.B)
}

func (f *Foo) Double() {
	f.A = f.A * 2
}

type Bar struct {
	C bool
	Foo
}

func (b Bar) String() string {
	return fmt.Sprintf("%s and C: %v", b.Foo.String, b.C)
}

func (b Bar) fieldCount() int {
	return 3
}

func main() {
	f := Foo {
		A: 10,
		B: "hello",
	}
	bb := Bar{
		C: true,
		Foo: f,
	}

	fmt.Println("f:", f.String())
	f.Double()
	fmt.Println("f:", f.String())

	fmt.Println("bb:", bb.String())
	bb.Double()
	fmt.Println("bb:", bb.String())

	m := myInt(10)
	fmt.Println("m:", m)
	fmt.Println("m:", m.isEven())
	m.Double()
	fmt.Println("New m:", m)
}