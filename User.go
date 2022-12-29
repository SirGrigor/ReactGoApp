package main

import (
	"fmt"
	"github.com/google/uuid"
)

type User struct {
	Id        int
	FirstName string
	LastName  string
	IdCode    string
	Email     string
}

func CreateNewUser(id int, firstName, lastName, idCode, email string) User {
	return User{
		Id:        id,
		FirstName: firstName,
		LastName:  lastName,
		IdCode:    idCode,
		Email:     email,
	}
}

func GenerateUsers(num int) []User {
	users := make([]User, num)
	for i := 0; i < num; i++ {
		firstName := fmt.Sprintf("User %d", i+1)
		lastName := fmt.Sprintf("User %d", i+1)
		random, err := uuid.NewRandom()
		if err != nil {
			return nil
		}
		idCode := fmt.Sprintf("%d", random.String())
		email := fmt.Sprintf("%s.%s@supermail.com", firstName, lastName)
		users[i] = CreateNewUser(i, firstName, lastName, idCode, email)
	}
	return users
}
