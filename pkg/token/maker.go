package token

import "time"

//Maker is an interface for managing token

type Maker interface {
	//	CreateToken creates tokens

	CreateToken(username string, duration time.Duration) (string, error)
	//VerifyToken verifies if the token is valid or not
	VerifyToken(token string) (*Payload, error)
}
