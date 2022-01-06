package Queue

type Queue struct {
	List []interface{}
}

func NewQueue () *Queue {
	return &Queue{}
}

func (q *Queue)QueueLen() int {
	return len(q.List)
}

func (q *Queue) Pop() interface{}  {
	if len(q.List)==0{
		return nil
	}
	res:= q.List[0]
	q.List = q.List[1:]
	return res
}

func (q *Queue)Push(param interface{}) {
	q.List = append(q.List,param)
}