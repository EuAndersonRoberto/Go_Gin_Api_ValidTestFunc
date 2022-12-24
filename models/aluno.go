package models

import (
	"gopkg.in/validator.v2" //Biblioiteca de validações.
	"gorm.io/gorm"
)

// Aqui estamos utilizando o "Validator.v2" com seus parametros para realizar as validações conforme mostrado em sua documentação em: https://pkg.go.dev/gopkg.in/validator.v2#section-readme.
type Aluno struct {
	gorm.Model        // Este comando nos trás uma struct do GORM que já nos apresenta: id, create, update e delete.
	Nome       string `json:"nome" validate:"nonzero"`                //Conforme documentação este validade nos apresenta: (Para int = 0, para string = "", para ponteiros = nulo, etc.)
	CPF        string `json:"cpf" validate:"len=11, regexp=^[0-9]*$"` //Aqui estamos utilizando o "len" da documentação validator que nos retorna a validação de númeos de intem ex: len=11.
	RG         string `json:"rg" validate:"len=7, regexp=^[0-9]*$"`   //Aqui estamos utilizando o "len" da documentação validator que nos retorna a validação de númeos de intem ex: len=9.
}

func ValidaDadosDeAluno(aluno *Aluno) error {
	if err := validator.Validate(aluno); err != nil {
		return err
	}
	return nil
}
