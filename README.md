# Hello Full Cycle utilizando Buffalo

Informações do desafio

	Esse desafio consiste em criar uma página web com o conteúdo "Hello Full Cycle" utilizando a linguagem Golang e o framework Buffalo. 
	Quando o usuário acessar o projeto no endpoint /hello, ele deverá ver a mensagem "Hello Full Cycle". 
	Para fazer a entrega do desafio, gere uma imagem docker de sua aplicação funcionando e a disponibilize em sua conta no Docker Hub. 

	Quando alguém executar: docker run -p 3000:3000 seuuser/suaimagem, a aplicação deverá estar disponível na porta 3000. 

## Starting the Application

Buffalo ships with a command that will watch your application and automatically rebuild the Go binary and any assets for you. To do that run the "buffalo dev" command:

	$ buffalo dev
	or
	$ make run-dev

If you point your browser to [http://127.0.0.1:3000/hello](http://127.0.0.1:3000/hello) you should see a "Welcome to Buffalo!" page.

Segue o [link](https://hub.docker.com/repository/docker/samuelsantos/hello_fullcycle) da imagem.

