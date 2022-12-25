package controllers

import (
	"net/http"

	"github.com/Api_Go_Gin/db"
	"github.com/Api_Go_Gin/models"
	"github.com/gin-gonic/gin"
)

// Aqui definimos todos as funções que utilizaremos no projeto mas, poderiamos colocar cada função em um file específico dentro da pasta controllers, modularizando ainda mais nosso código. Neste projeto faremos todas as funções dentro deste único file.
func ExibeTodosAlunos(c *gin.Context) {
	var alunos []models.Aluno
	db.DB.Find(&alunos)
	c.JSON(200, alunos) //Aqui fazemos uso da variável "Alunos" que nos trás uma lista da struct "Aluno" criada em "models", para ser exibida pelo JSON.
}

// Saliento que todo o contexto de nossas requisições é controlada por nossa propriedade "c".
func Saudacao(c *gin.Context) {
	nome := c.Params.ByName("nome") // Está variável recupera o parametro "nome" do endpoint criado no file "routes.go" através do "c.Params.ByName("nome")".
	c.JSON(200, gin.H{
		"API diz:": "Iae! " + nome + ", tudo bem com você?",
	})
}

func CriaNovoAluno(c *gin.Context) {
	var aluno models.Aluno
	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": err.Error()})
		return
	}
	db.DB.Create(&aluno)
	c.JSON(http.StatusOK, aluno)
}

func BuscaAlunoPorId(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")
	db.DB.First(&aluno, id)
	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Aluno não encontrado!"})
		return
	}
	c.JSON(http.StatusOK, aluno)
}

func DeleteAluno(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")
	db.DB.Delete(&aluno, id)
	c.JSON(http.StatusOK, gin.H{"Data": "Aluno deletado com sucesso!"})
}

func AtualizaAluno(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")
	db.DB.First(&aluno, id)
	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.DB.Model(&aluno).UpdateColumns(aluno)
	c.JSON(http.StatusOK, aluno)
}

func BuscaAlunoPorCPF(c *gin.Context) {
	var aluno models.Aluno
	cpf := c.Param("cpf")
	db.DB.Where(&models.Aluno{CPF: cpf}).First(&aluno)
	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Aluno não encontrado!"})
		return
	}
	c.JSON(http.StatusOK, aluno)
}

func ExibePaginaIndex(c *gin.Context) {
	var alunos []models.Aluno
	db.DB.Find(&alunos)
	c.HTML(http.StatusOK, "index.html", gin.H{
		"alunos": alunos,
	})

}

func RotaNaoEcontrada(c *gin.Context) {
	c.HTML(http.StatusFound, "404.html", nil)
}
