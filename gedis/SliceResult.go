package gedis

type SliceResult struct {
	Result []interface{}
	Err 	error
}

func NewSliceResult(result []interface{}, err error) *SliceResult {
	return &SliceResult{Result: result, Err: err}
}

func (s *SliceResult) Unwrap() []interface{}  {
	if s.Err != nil {
		panic(s.Err)
	}
	return s.Result
}

func (s *SliceResult) UnwrapOr(v []interface{}) []interface{}  {
	if s.Err != nil {
		return v
	}
	return s.Result
}

func (s *SliceResult) Iter() *Iterator  {
	return  NewIterator(s.Result)

}