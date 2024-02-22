package middlewares

import (
    "golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
    hashedPwd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

    if err != nil {
	return "", err
    }

    return string(hashedPwd), nil
}

func CheckPassword(password, hashedPwd string) (bool, error) {
    
    err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(password))

    if err == nil {
	return true, nil
    
    } else if err == bcrypt.ErrMismatchedHashAndPassword {
	return false, nil
    
    } else {
	return false, err
    }
	
}
