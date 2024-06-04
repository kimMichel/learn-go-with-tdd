package hello

import "fmt"

const portuguese = "portuguese"
const spanish = "spanish"
const french = "french"

const defaultHelloPrefix = "Hello "
const portugueseHelloPrefix = "Ola "
const spanishHelloPrefix = "Hola "
const frenchHelloPrefix = "Bonjour "

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	return helloPrefix(language) + name
}

func helloPrefix(language string) (prefix string) {
	switch language {
	case portuguese:
		prefix = portugueseHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = defaultHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("World", "english"))
}