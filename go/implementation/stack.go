package implementation

type Node struct {
	Next *Node
	Data interface{}
}

type Stack struct {
	Head *Node
	Size int
}

func (s *Stack) Push(data interface{}) {
	newNode := &Node{
		Next: s.Head,
		Data: data,
	}

	s.Head = newNode
	s.Size++
}

func (s *Stack) Pop() interface{} {
	if s.Head == nil {
		return nil
	}

	data := s.Head.Data
	s.Head = s.Head.Next

	s.Size--
	return data
}

func (s *Stack) Peek() interface{} {
	return s.Head.Data
}

func (s *Stack) IsEmpty() bool {
	return s.Size == 0
}

func (s *Stack) GetSize() int {
	return s.Size
}
