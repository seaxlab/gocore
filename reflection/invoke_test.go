package reflection

import "testing"

// spy 2020/1/26
type YourT1 struct{}

func (y YourT1) MethodBar() {
	//do something
	logger.Info("MethodBar")
}

type YourT2 struct{}

func (y YourT2) MethodFoo(i int, oo string) {
	//do something
	logger.Info("MethodFoo")
	logger.Info("i=", i, ",oo=", oo)
}
func TestInvoke(t *testing.T) {

	Invoke(YourT2{}, "MethodFoo", 10, "abc")
	Invoke(YourT1{}, "MethodBar")
}
