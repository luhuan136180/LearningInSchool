package main

type caculate interface {
	gettarea()
	getRound()
}
type Triangle struct{
	p float64
	side1 float64
	side2 float64
	side3 float64
}

func (t *Triangle) getares() (tarea float64){
	return t.p*(t.p - t.side1)*(t.p-t.side2)*(t.p-t.side3)
}
func (t *Triangle) getRound() (p float64) {
	t.p=t.side1+t.side2+t.side3
	return t.p
}
type Circle struct {
	r float64
}
func main() {

}

