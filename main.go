package main

import (
	"github.com/Api_Go_Gin/db"
	"github.com/Api_Go_Gin/routes"
)

func main() {
	//Aqui utlizamos o main.go para alimentar com informações a variável "Alunos" (vinda de models), que representa a lista da struct "Aluno" (Struct criada também em models).
	db.ConectaComBancoDeDados() //Aqui realizamos a conexão com o banco de dados.
	/*models.Alunos = []models.Aluno{
		{Nome: "Anderson Roberto", CPF: "123.312.654-11", RG: "2.222-222"},
		{Nome: "Ana Helka", CPF: "123.312.777-22", RG: "2.111-333"},
	}*/
	routes.HandleRequests()
}
