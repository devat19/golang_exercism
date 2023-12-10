package airportrobot

// Write your code here.
// This exercise does not have tests for each individual task.
// Try to solve all the tasks first before running the tests.
// type Greeter interface {
// 	LanguageName() string
// 	Greet(name string) string
// }

// // func (c *Greeter) SomeOtherMethod() {
// // 	// The type can have additional methods not mentioned in the interface.
// // }

// type Italian struct {
// }

// func SayHello(name string, germanGreeter Greeter) {

// }

type Greeter interface {
	LanguageName() string
	Greet(name string) string
}

func SayHello(name string, greeter Greeter) string {
	return "I can speak " + greeter.LanguageName() + ": " + greeter.Greet(name)
}

type Italian struct{}

func (i Italian) LanguageName() string {
	return "Italian"
}
func (i Italian) Greet(name string) string {
	return "Ciao " + name + "!"
}

type Portuguese struct{}

func (p Portuguese) LanguageName() string {
	return "Portuguese"
}
func (p Portuguese) Greet(name string) string {
	return "Olá " + name + "!"
}
