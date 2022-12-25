package routes

import (
	"github.com/Api_Go_Gin/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./assets")
	r.GET("/alunos", controllers.ExibeTodosAlunos)          // Aqui definimos que toda requisão GET feita para o endpoint: /alunos, irá responder com a função "ExibeTodosAlunos", nos devolvendo o código 200 e a mensagem JSON: "id":"1" e "nome":"Anderson Roberto", na URL: http://localhost:8080/alunos
	r.GET("/:nome", controllers.Saudacao)                   //Aqui definimos mais uma requisição GET, sendo específicado o endpoint "/:nome" que tem a possibilidade de ser alterado e o responsável por essa requisição GET é o controllers."Saudacao".
	r.POST("/alunos", controllers.CriaNovoAluno)            //Aqui definimos uma requisição POST, responsável por criar um novo aluno no endpoint "/alunos".
	r.GET("/alunos/:id", controllers.BuscaAlunoPorId)       //Aqui definimos uma requisição GET, específicando um novo endpoint "/alunos/:id" e tendo o controllers."BuscaAlunoPorId" como responsável por essa requisição.
	r.DELETE("/alunos/:id", controllers.DeleteAluno)        //Aqui definimos uma requisição DELETE, específicando que dentro do endpoint "/alunos/:id" e tendo o controllers."DeleteAluno" como responsável por essa requisição.
	r.PATCH("/alunos/:id", controllers.AtualizaAluno)       //Aqui definimos uma requisição UPDATE, específicando que dentro do endpoint "/alunos/:id" e tendo o controllers."AtualizaAluno" como responsável por essa requisição.
	r.GET("/alunos/cpf/:cpf", controllers.BuscaAlunoPorCPF) //Aqui definimos uma requisição GET, específicando um novo endpoint "/alunos/cpf/:cpf" e tendo o controllers."BuscaAlunoPorCPF" como responsável por essa requisição.
	r.GET("/index", controllers.ExibePaginaIndex)           //Aqui definimos uma requisição GET, específicando um novo endpoint "/index" e tendo o controllers."ExibePaginaIndex" como responsável por essa requisição.
	r.NoRoute(controllers.RotaNaoEcontrada)
	r.Run()
}
