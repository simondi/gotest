package main

import "fmt"

type Mutatable struct {
    a int
    b int
}

func (m Mutatable) StayTheSame() {
    m.a = 5
    m.b = 7
}

func (m *Mutatable) Mutate() {
    m.a = 5
    m.b = 7
}

func (m Mutatable) MakeChange(i int)  Mutatable {
    m.a = 5*i
    m.b = 7*i
	return m
}


func main() {
    m := &Mutatable{0, 0}
    fmt.Println(m)
    m.StayTheSame()
    fmt.Println(m)
    m.Mutate()
    fmt.Println(m)
	n := m.MakeChange(2)
    fmt.Println(n)
}