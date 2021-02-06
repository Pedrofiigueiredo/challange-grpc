package domain

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

// User é a entidade usuario
type User struct {
	Base
	Name     string
	Email    string
	Password string
	Token    string
}

// Prepare adiciona dados padrao do usuario (id, datas, senha criptografada)
func (user *User) Prepare() error {
	// Criprografar a senha do usuario
	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	// Adicionar dados padrao ao usuario
	user.ID = uuid.NewV4().String()
	user.CreatedAt = time.Now()
	user.Password = string(password)
	user.Token = uuid.NewV4().String()

	// Validacao dos dados
	err = user.validate()

	if err != nil {
		return err
	}

	return nil
}

func (user *User) validate() error {
	return nil
}
