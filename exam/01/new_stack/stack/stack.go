package stack

// Stack - структура стека
type Stack struct {
	items []interface{}
}

// NewStack - создает новый экземпляр стека
func NewStack() *Stack {
	return &Stack{
		items: []interface{}{},
	}
}

// Push - добавляет элемент на вершину стека
func (s *Stack) Push(f interface{}) {
	s.items = append(s.items, f)
}

// Pop - удаляет и возвращает элемент с вершины стека
func (s *Stack) Pop() interface{} {
	if len(s.items) == 0 {
		return nil
	}
	lastIndex := len(s.items) - 1
	item := s.items[lastIndex]
	s.items = s.items[:lastIndex]
	return item
}
