package repository

import (
	"chat-api/pkg/entity"

	"github.com/gocql/gocql"
)

type UserRepository interface {
	// create
	CreateUser(u entity.User) (*entity.User, error)

	// read
	GetByID(id string) (*entity.User, error)

	ListUser() ([]entity.User, error)

	// Update
	Update(id string, u entity.User) (*entity.User, error)

	// delete
	Delete(id string) error
}

type User struct {
	session *gocql.Session
}

func NewUser(session *gocql.Session) *User {
	return &User{session: session}
}

func (r *User) CreateUser(u entity.User) (*entity.User, error) {
	var query = `INSERT INTO user(user_id, name, password) VALUES(?, ?, ?)`

	if err := r.session.Query(query, u.UserID, u.Name, u.Password).Exec(); err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *User) GetByID(id string) (*entity.User, error) {
	var query = `SELECT * FROM user WHERE user_id = ?`
	var u entity.User

	if err := r.session.Query(query, id).Scan(&u.UserID, &u.Name, &u.Password); err != nil {
		if err == gocql.ErrNotFound {
			return nil, err
		}

		return nil, err
	}

	return &u, nil
}

func (r *User) ListUser() (userList []entity.User, _ error) {
	var query = `SELECT * FROM user`

	iter := r.session.Query(query).Iter()
	var user entity.User
	for iter.Scan(&user.UserID, &user.Name, &user.Password) {
		userList = append(userList, user)
	}

	if err := iter.Close(); err != nil {
		return nil, err
	}

	return userList, nil
}

func (r *User) Update(id string, u entity.User) (*entity.User, error) {
	// UPDATE user SET name = ?, password = ? WHERE user_id = ?;

	var query = `UPDATE user SET name = ?, password = ? WHERE user_id = ?`
	if err := r.session.Query(query, u.Name, u.Password, id).Exec(); err != nil {
		return nil, err
	}

	return &u, nil
}

func (r *User) Delete(id string) error {
	var query = `DELETE FROM user WHERE user_id = ?`

	if err := r.session.Query(query, id).Exec(); err != nil {
		return err
	}

	return nil
}
