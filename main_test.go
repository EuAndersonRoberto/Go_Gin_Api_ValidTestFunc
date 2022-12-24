package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Api_Go_Gin/controllers"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetupDasRotasDeTeste() *gin.Engine {
	rotas := gin.Default()
	return rotas
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
