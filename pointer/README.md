# RESUMO

## Ponteiros

- Go copia os valroes quando sao passados para funcoes/metodos. Entao, se estiver escrevendo uma funcao que precise mudar o estado, voce precisara de um ponteiro para o valor que voce quer mudar.

- O fato de que Go pega um a copia dos valores e muito util na maior parte do tempo, mas as vezes voce nao vai querer que o seu sistema faca copia de alguma coisa. Nesse caso, voce precisa passar uma referencia. Podemos, por exemplo, ter dados muito grandes, ou coisas que voce talvez pretenda ter apenas uma instancia (como conexoes a banco de dados).

## nil

- Ponteiros podem ser nil.

- Quando uma funcao retorna um ponteiro para algo, voce precisa ter certeza de verificar se ele e nil ou isso vai gerar uma excessao em tempo de execucao, ja que o compilador nao te consegue te ajudar nesses casos.

- Util para quando voce quer descrever um valor que pode estar faltando.

## Erros

- Erros sao a forma de sinalizar falhas na execucao de uma funcao/metodo.

- Analisando nossos testes, concluimos que buscar por uma string em um erro poderia resultar em um teste nao muito confiavel. Entao, refatoramos para usar um valor significativo, que resultou em um codigo mais facil de ser testado e concluimos que tambem seria mais fcail para usuarios de nossa API.

- Nao somente verifique os erros, trate-os graciosamente
