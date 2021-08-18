package models;

import (
	"errors"
	"fmt"
)

type User struct {
	ID int 
	FirstName string
	LastName string
}

var (
	users []*User
	nextID = 1
)

func GetUsers() []*User {
	return users;
}

/*
AddUser will add a new user.
*/
func AddUser(u User) (User, error) {
	if u.ID != 0 {
		return User{}, errors.New("New users must not include ID or it must be equal to 0")
	}
	u.ID = nextID
	nextID++
	users = append(users, &u)
	return u, nil
}

/*
GetUserByID function returns the user info wrt to the id given as parameter.
*/
func GetUserByID(id int) (User, error) {
	for _, u := range users {
		if u.ID == id {
			return *u, nil
		}
	}
	return User{}, fmt.Errorf("User with ID %v was not found", id)
}

/*
UpdateUser functino will update the information of the user.
*/
func UpdateUser(u User) (User, error) {
	for i, candidate := range users {
		if candidate.ID == u.ID {
			users[i] = &u
			return u, nil
		}
	}
	return User{}, fmt.Errorf("User with ID %v was not found", u.ID)
}

/*
RemoveUserByID function will remove the user by ID given in parameter.
*/
func RemoveUserByID(id int) error {
	for i, u := range users {
		if u.ID == id {
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("User with ID %v not found", id)
}

func remove(slice []*User, s int)  []*User {
    return append(slice[:s], slice[s+1:]...)
}