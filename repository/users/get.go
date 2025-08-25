package users

import (
	"apigo/entity"
	"database/sql"
	"fmt"
)

type UsersGetRepository struct {
	connection *sql.DB
}

func NewUserGet(connection *sql.DB) UsersGetRepository {
	return UsersGetRepository{
		connection: connection,
	}
}

func (us *UsersGetRepository) GetUser() ([]entity.Users, error) {
	query := "SELECT * FROM users"
	row, err := us.connection.Query(query)

	if err != nil {
		return nil, err
	}
	userList := []entity.Users{}
	var usersObject entity.Users

	defer row.Close()

	for row.Next() {
		err = row.Scan(
			&usersObject.ID,
			&usersObject.UserName,
			&usersObject.Email,
			&usersObject.Password,
			&usersObject.Role,
		)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		userList = append(userList, usersObject)
	}

	return userList, nil
}
