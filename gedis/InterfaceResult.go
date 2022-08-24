package gedis


type InterfaceResult struct {
	Result interface{}
	Err error
}
func NewInterfaceResult(result interface{}, err error) *InterfaceResult {
	return &InterfaceResult{Result: result, Err: err}
}
func(i *InterfaceResult) Unwrap() interface{} {
	if i.Err!=nil{
		panic(i.Err)
	}
	return i.Result
}
func(i *InterfaceResult) UnwrapOr(v interface{}) interface{} {
	if i.Err!=nil{
		return v
	}
	return i.Result
}