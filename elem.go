package raffle

type Elem struct {
	typ     uint32
	weight  uint32
	content interface{}
}

func NewElems(typ, weight uint32, num int, content interface{}) (elemList []Elem) {
	for i := 0; i < num; i++ {
		elemList = append(elemList, Elem{
			typ:     typ,
			weight:  weight,
			content: content,
		})
	}

	return
}
