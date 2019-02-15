package main

type computer interface {
	_type() string
	_score() string
}

type mac struct {
	myType string
	score string
}

func (m *mac) _type()string  {
	return m.myType
}
func (m *mac) _score()string{
	return m.score
}


func main() {
	var test computer
	test = new (mac)
	test._score()
	test._type()


}
