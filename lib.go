package mylib

type MyStruct struct {
	X int64
}

type MyInterface interface {
	DoSomething(m MyStruct)
}

func ExportedAPI(x any) {
	if y, ok := x.(MyInterface); ok {
		y.DoSomething(MyStruct{X: 1})
	}
}
