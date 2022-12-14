package gedis

type Iterator struct {
	data  []interface{}
	index int
}

func NewIterator(data []interface{}) *Iterator {
	return &Iterator{data: data}
}

func (i *Iterator) HasNext() bool  {
	if i.data == nil  || len(i.data) == 0 {
		return false
	}
	return i.index < len(i.data)
}

func (i * Iterator)  Next() (ret interface{}) {
	ret = i.data[i.index]
	i.index = i.index+1
	return ret

}