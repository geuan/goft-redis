package gedis

type StringResult struct {
	Result  string
	Err error
}

func NewStringResult(result string, err error) *StringResult {
	return &StringResult{Result: result, Err: err}
}

func (s *StringResult) Unwrap() string  {
	if s.Err != nil {
		panic(s.Err)
	}
	return s.Result
}

func (s *StringResult) UnwrapOr(str string) string  {
	if s.Err != nil {
		return str
	}
	return s.Result
}

func (s *StringResult) UnwrapOrElse(f func() string) string  {
	if s.Err != nil {
		return f()
	}
	return s.Result
}