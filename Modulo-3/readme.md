Essa é uma aplicação de terminal feita no terceiro módulo do curso de [**GOLANG**](https://www.udemy.com/course/aprenda-golang-do-zero-desenvolva-uma-aplicacao-completa/) da Udemy.
---
A aplicação funciona da seguinte forma:

O comando abaixo retorna o endereço IP de um host passado por paramêtro
```
$ go run main.go ip --host <endereço>
```

O comando abaixo retorna o nome do servidor de um host passado por paramêtro
```
$ go run main.go server --host <endereço>
```

Os comandos podem ser executados sem os paramêtros de **--host**  **< endereço >**, nesse caso será usado o host padrão (lesse.com.br). 