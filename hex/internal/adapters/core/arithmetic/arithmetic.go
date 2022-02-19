package arithmetic

/*
This is the arithmetic adapter. it supplies the Adapter datatype wich follows the specifications of the arithmetic interface
*/

type Adapter struct {
}

func NewAdapter() *Adapter {
	return &Adapter{} // This adapter struct literal returns the adapter with all functions attached
}

func (arith Adapter) Addition(a int32, b int32) (int32, error) {
	return a + b, nil
}

func (arith Adapter) Subtraction(a int32, b int32) (int32, error) {
	return a - b, nil
}

func (arith Adapter) Multiplication(a int32, b int32) (int32, error) {
	return a * b, nil
}

func (arith Adapter) Division(a int32, b int32) (int32, error) {
	return a / b, nil
}
