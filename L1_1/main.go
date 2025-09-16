package main

import "fmt"

type Human struct {
	age  int
	name string
}

func (h *Human) greeting() {
	fmt.Printf("Hi, my name is %s", h.name)
}

type Action struct {
	Human
	topic string
}

func (a *Action) showTopic() {
	fmt.Printf("my topic is %s", a.topic)
}
func main() {
	Bob := Action{Human: Human{age: 15, name: "bob"}, topic: "sleep"}
	Bob.greeting()

}
