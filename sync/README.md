# Resumo

Falamos sobre algumas coisas do pacote sync:

- Mutex nos permite adicionar travas aos nossos dados
- WaitGroup e uma maneira de esperar as goroutines terminarem suas tarefas

## Quando usar travas em vez de channels e goroutines?

Resumindo:

- Use channels quando for passar a propriedade de um dado
- Use mutexes para gerenciar estados

## go vet

Nao se esqueca de usar go vet nos seus scripts de build porque ele pode te alertar a respeito de bugs mais sutis no seu codigo antes que eles atinjam seus pobres usuarios.

## Nao use codigos embutidos apenas porque e conveniente

- Pense a respeito do efeito que embutir codigos tem na sua API publica
- Voce realmente quer expor esses metodos e ter pessoas acoplando codigo porprio delas a ele?
- Mutexes podem se tornar um desastre de maneiras muito imprevisiveis e estranhas. Imagine um codigo inesperado destravando um mutex quando nao deveria? Isso causaria erros muito estranhos que seriam muito dificies de encontrar.
