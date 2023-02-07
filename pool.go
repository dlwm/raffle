package raffle

type Pool struct {
	list      []Elem
	weightSum uint32
	// TODO 后续做池累计等
}

func (p *Pool) RepeatDraw(count int) (typList []uint32, contentList []interface{}) {
	if p.weightSum == 0 {
		return
	}

	for i := 0; i < count; i++ {
		val := r.Intn(int(p.weightSum)) + 1
		for _, e := range p.list {
			val -= int(e.weight)
			if val <= 0 {
				typList = append(typList, e.typ)
				contentList = append(contentList, e.content)
				break
			}
		}
	}

	return
}

func (p *Pool) ProbList() (probList []float32, typList []uint32, contentList []interface{}) {
	for _, elem := range p.list {
		typList = append(typList, elem.typ)
		contentList = append(contentList, elem.content)
		probList = append(probList, float32(elem.weight)/float32(p.weightSum))
	}

	return
}

//func (p *Pool) SingleDraw(count int)  (typList []uint32, contentList []interface{})

func NewPool(list []Elem) (p *Pool) {
	p = &Pool{list: list}

	for _, e := range list {
		p.weightSum += e.weight
	}

	return
}
