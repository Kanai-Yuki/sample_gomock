package sample

// Sample ...
type Sample interface {
	Method(a, b int) int
}

// ABC ...
type ABC struct {
}

// Method ...
func (abc *ABC) Method(a, b int) int {
	return a + b
}
