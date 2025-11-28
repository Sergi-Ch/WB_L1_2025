package main

import "fmt"

type Sender interface { // Ожидаемый интерфейс
	Send()
}

type JsonSender struct { // Изначальный класс
}

func (jSender JsonSender) SendJson() {
	fmt.Println("sending Json")
}

type JsonToSenderAdapt struct { // Сам адаптер
	JsonSender JsonSender
}

func (JsTSA JsonToSenderAdapt) Send() { // функция для имплементации нужного интерфейса
	JsTSA.JsonSender.SendJson()
}

func TestClient(send Sender) {
	send.Send() // тест клиент, который на вход получает интерфейс и в который мы садаптировали изначальную функцию.
}

func main() {
	jsonSender := JsonSender{}
	jsonAdapter := JsonToSenderAdapt{JsonSender: jsonSender}
	TestClient(jsonAdapter)
}
