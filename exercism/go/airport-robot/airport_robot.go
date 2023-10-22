package airportrobot

import "fmt"

type Greeter interface {
	LanguageName() string
	Greet(name string) string
}

func SayHello(name string, g Greeter) string {
	return fmt.Sprintf("I can speak %s: %s!", g.LanguageName(), g.Greet(name))
}

type Italian struct{}

func (it Italian) LanguageName() string {
	return "Italian"
}

func (it Italian) Greet(name string) string {
	return fmt.Sprintf("Ciao %s", name)
}

type Portuguese struct{}

func (pt Portuguese) LanguageName() string {
	return "Portuguese"
}

func (pt Portuguese) Greet(name string) string {
	return fmt.Sprintf("Olá %s", name)
}
