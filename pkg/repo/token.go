package repo

import (
	mylog "business/pkg/log"
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

var log = mylog.Log

//Token ...
type Token struct {
	UserID string
	Login  string
	RoleID string

	jwt.StandardClaims
}

//GenerateToken ...
func (t *Token) GenerateToken(userID, login, roleID, signedString string) string {

	tk := &Token{
		UserID: userID,
		Login:  login,
		RoleID: roleID,
	}

	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)

	tokenString, _ := token.SignedString([]byte(signedString))

	return tokenString
}

//Decode ...
func (t *Token) Decode(tokenString string) (*Token, error) {

	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})
	if err != nil {
		log.Error(fmt.Sprint(err))
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		log.Error("Can't convert token's claims to standard claims")
		return nil, err
	}

	t.UserID = fmt.Sprintf("%v", claims["UserID"])
	t.Login = fmt.Sprintf("%v", claims["Login"])
	t.RoleID = fmt.Sprintf("%v", claims["RoleID"])

	//log.PrintfErrorMsg("__cl", claims)
	return t, nil
}
