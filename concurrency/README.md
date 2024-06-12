# RESUMO

Tornando as coisas mais rapidas, aprendemos sobre:

- goroutines, a unidade baisca de concorrencia em GO, que nos permite verificar mais do que um site ao mesmo tempo.

- funcoes anonimas, que usamos para iniciar cada um dos processos concorrentes que verificam os sites.

- canais, para nos ajudar a organizar e controlar a comunicacao entre diferentes processos, nos permitindo evitar um bug de condicao de corrida.

- o detector de corrida, que nos ajudou a desvendar problemas com codigo concorrente.

## Torne-o rapido

Uma formulacao da forma agil de desenvolver software, erroneamente atribuida a Kent Beck, e:

```
Faca funcionar, faca da forma certa, torne-o rapido
```

Onde 'funcionar' e fazer os testes passarem, 'forma certa' e refatorar o codigo e 'tornar rapido' e otimizar o codigo para, por exemplo, tornar sua execucao rapida. So podemos 'torna-lo rapido' quando fizermos funcionar da forma certa.

Nunca devemos tentar 'torna-lo rapido' antes das outras duas etapas terem sido feitas, porque:

```
Otimizacao prematura e a raiz de todo o mal
```
