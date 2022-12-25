package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/Api_Go_Gin/controllers"
	"github.com/Api_Go_Gin/db"
	"github.com/Api_Go_Gin/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var ID int //Como queremos que o "ID" seja visto por outras classes, fizemos a declaração dele aqui para todo o código.

func SetupDasRotasDeTeste() *gin.Engine {
	gin.SetMode(gin.ReleaseMode) //Este Set melhora a visualização das respostas de nossos testes.
	rotas := gin.Default()
	return rotas
}

// Aqui definimos a criação e o delete de um aluno exemplo para que apareça e seja deletado dentro de nossas funções que atuam com nosso banco de dados.
func CriaAlunoMock() {
	aluno := models.Aluno{Nome: "Nome do Aluno Testes", CPF: "12312345689", RG: "1231237"}
	db.DB.Create(&aluno)
	ID = int(aluno.ID) //Por ter sido declarado para todo o código não há a necessidade de colocar os " : ".
}

func DeletaAlunoMock() {
	var aluno models.Aluno
	db.DB.Delete(&aluno, ID)
}

func TestVerificaStatusCodeDaSaudacaoComParametro(t *testing.T) {
	r := SetupDasRotasDeTeste()                        //Este comando cria uma nova instância no Gin que ainda não tem uma rota cadastrada, que definimos na próxima linha de código.
	r.GET("/:nome", controllers.Saudacao)              //Rota definida (models).
	req, _ := http.NewRequest("GET", "/anderson", nil) //Aqui definimos como será a nossa requisição. No caso, nos trás mais uma mensagem de erro mas, que não faremos uma verificação agora, por isso foi colocado o "_" após a requisição. Também definimos o metodo "GET", definimos também a string que vai na URL: "/anderson" e qual é o copor da requisição? tem JSON? algum dado? neste caso não passamos nada então: "nil".
	resposta := httptest.NewRecorder()                 //Esta linha de comando sever para armazenar a resposta da requisição utilizando o "httptest".Utilizando também a função "NewRecorder" que o objetivo de implementar a interface de quem vai armazenar essa resposta.
	r.ServeHTTP(resposta, req)                         //Aqui é onde de fato estamos realizando a requisição, onde passamos os argumentos já estruturados: "resposta" e "req".

	/* if resposta.Code != http.StatusOK {                //Aqui apenas verificamos se o status code(200) é igual a resposta de nossa requisição, realiando assim o teste da função "Saudação".
		t.Fatalf("Status error: Valor recebido foi %d e o esperado era %d", resposta.Code, http.StatusOK) //Importante salientar que o code recebido vem de "resposta.code" e o valor esperado vem de "http.StatusOK".
	}*/

	assert.Equal(t, http.StatusOK, resposta.Code, "Deveriam ser iguais!") // Aqui substituimos todo o "if" utilizado anteriormente pelo o "assert" que importamos com o comando:"go get github.com/stretchr/testify". O Equal é para sabermos se é igual, passamos também nossa instância de teste "t", agora passamos o valor esperado através do "http.StatusOK" e o valor recebido pelo "resposta.code" e para finalizar passamos uma mensagem de error.

}

func TestListandoTodosOsAlunosHandle(t *testing.T) {
	db.ConectaComBancoDeDados()
	CriaAlunoMock()         //Aqui criamos um aluno mock
	defer DeletaAlunoMock() //Aqui deletamos o aluno mock
	r := SetupDasRotasDeTeste()
	r.GET("/alunos", controllers.ExibeTodosAlunos)
	req, _ := http.NewRequest("GET", "/alunos", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code, "Deveria ser igual")
	/*fmt.Println(resposta.Body) //Aqui utilizamos esse comando para verificar se realmente estou acessand o banco de dados de nossa  aplicação.*/
}

func TestBuscaAlunoPorCPFHandle(t *testing.T) {
	db.ConectaComBancoDeDados()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	r := SetupDasRotasDeTeste()
	r.GET("/alunos/cpf/:cpf", controllers.BuscaAlunoPorCPF)
	req, _ := http.NewRequest("GET", "/alunos/cpf/12332132189", nil) //importante colocar um CPF válido, um CPF que conste no banco de dados do projeto.
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code, "Deveria ser igual")
}

func TestBuscaAlunoPorIdHandle(t *testing.T) {
	db.ConectaComBancoDeDados()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	r := SetupDasRotasDeTeste()
	r.GET("/alunos/:id", controllers.BuscaAlunoPorId)
	pathDaBusca := "/alunos/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("GET", pathDaBusca, nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	var AlunoMock models.Aluno
	json.Unmarshal(resposta.Body.Bytes(), &AlunoMock)
	//fmt.Println(AlunoMock.Nome)
	assert.Equal(t, "Nome do aluno teste", AlunoMock.Nome, "Os nomes deveriam ser iguais")
	assert.Equal(t, "12332132189", AlunoMock.CPF, "Os CPFs deveriam ser iguais")
	assert.Equal(t, "1231237", AlunoMock.RG, "Os RGs deveriam ser iguais ")
	assert.Equal(t, http.StatusOK, resposta.Code, "Deveria ser igual")
}

func TestDeleteAlunoHandle(t *testing.T) {
	db.ConectaComBancoDeDados()
	CriaAlunoMock()
	r := SetupDasRotasDeTeste()
	r.DELETE("/alunos/:id", controllers.DeleteAluno)
	pathDeBusca := "/alunos/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("Delete", pathDeBusca, nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code)
}

func TestEditaUmAlunoGHandle(t *testing.T) {
	db.ConectaComBancoDeDados()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	r := SetupDasRotasDeTeste()
	r.PATCH("/alunos/:id", controllers.AtualizaAluno)
	aluno := models.Aluno{Nome: "Nome do Aluno Testes", CPF: "412312345689", RG: "123123700"}
	valorjson, _ := json.Marshal(aluno)
	pathParaEditar := "/alunos/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("PATCH", pathParaEditar, bytes.NewBuffer(valorjson)) // Aqui passamos um valor para o corpo da requição mas, não podemos passar essa informação em forma de JSON direto, para isso utilizamos o "bytes.NewBuffer()" e colocamos nosso JSON "valorjson"
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	var alunoMockAtualizado models.Aluno
	json.Unmarshal(resposta.Body.Bytes(), &alunoMockAtualizado) //Aqui pegamos todo o body da resposta e colocamos no "alunoMockAtualizado", nos possibilitando fazer nossas verificações.
	assert.Equal(t, "412312345689", alunoMockAtualizado.CPF)
	assert.Equal(t, "123123700", alunoMockAtualizado.RG)
	assert.Equal(t, "Nome do Aluno Testes", alunoMockAtualizado.Nome)
}
