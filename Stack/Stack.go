package Stack

type Stack struct {
	List []interface{}
}

func NewStack () *Stack {
	return &Stack{}
}

func (s *Stack)StackLen() int {
	return len(s.List)
}

func (s *Stack) Pop() interface{}  {
	if len(s.List)==0{
		return nil
	}
	res:=s.List[len(s.List)-1:]
	s.List = s.List[:len(s.List)-1]
	return res[0]
}

func (s *Stack)Push(param interface{}) {
	s.List = append(s.List,param)
}