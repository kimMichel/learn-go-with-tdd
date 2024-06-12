# RESUMO

- Como testar um handler HTTP que teve a requisicao cancelada pelo cliente.
- Como usar o contexto para gerenciar o cancelamento.
- Como escrever uma funcao que aceita context e o usa para se cancelar usaando goroutines, select e canais.
- Seguir as diretrizes da Google a respeito de como controlar o cancelamento propagando o contexto escopado da requisicao atraves da sua pilha de chamadas (call stack)
- Como levar seu proprio spy para http.ResponseWriter se voce precisar dele.

## E quanto ao context.Value?

Alguns engenheiros tem defendido a passagem de valores atraves do context porque parece conveniente.

A conveniencia e muitas vezes a causa do codigo ruim.

O problema ccom context.Values e que ele e apenas um mapa nao tipado para que voce nao tenha nenhum tipo de seguranca e voce tem que lidar com ele nao realmente contendo seu valor. Voce tem que criar um acoplamento de chaves de mapa de um modulo para outro e se alguem muda alguma coisa comecar a quebrar.

Resumindo, se uma funcao necessita de alguns valores, coloque-os como parametros tipados em vez de tentar obte-los a partir de context.Value. Isto torna-o estaticamente verificado e documentado para que todos o vejam.

Mas...

Por outro lado, pode ser util incluir informacoes que sejam ortogonoais a uma requisicao em um contexto, como um identificador unico. Potencialmente esta informacao nao seria necessaria para todas as funcoes da sua pilha de chamadas e tornaria as suas assinaturas funcionais muito confusas.
