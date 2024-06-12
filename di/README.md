# RESUMO

Gracas aos nossos testes, refatoramos o codigo para que pudessemos controlar para onde os dados eram escritos injetando uma dependencia que nos permitiu:

- Testar nosso codigo: se voce nao consegue testar uma funcao de forma simples, geralmente e porque dependencias estao acopladas em uma funcao ou estado global. Se voce tem um pool de conexao global da base de dados, por exxemplo, e provavel que seja dificil testar e vai ser lento para ser executado. A injecao de dependencia te motiva a injetar em uma dependencia da base de dados (atraves de uma interface), para que voce possa criar um mock com algo que voce possa controlar nos seus testes.

- Separar nossas preocupacoes, desacoplando onde os dados vao de como gera-los. Se voce ja achou que um metodo/funcao tem responsabilidades demais (gerando dados e escrevendo na base de dados? Lidando com requisicoes HTTP e aaplicando logica a nivel de dominio?), a injecao de dependencia provavelmente sera a ferramenta que voce precisa.

- Permitir que nosso codigo seja reutilizado em contextos diferentes: o primeiro contexto "novo" do nosso codigo que pode ser usado dentro dos testes. No entando, se alguem quiser testar algo novo com nossa funcao, a pessoa pode injetar suas proprias dependencias.

## Mas e o mock? Ouvi falar que precisa disso para trabalhar com injecao de dependencia e que tambem e do demonho

Voce mocka para substituir coisas reais que voce injeta com uma versao falsa que voce pode controlar e examinar nos seus testes. No entando, no nosso caso a biblioteca padrao ja tinha algo pronto para usarmos.
