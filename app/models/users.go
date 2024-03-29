package models

import (
	"log"
	"time"
)

type User struct {
	ID         int
	UUID       string
	Name       string
	Email      string
	PassWord   string
	CreatedAt  time.Time
	Phone      string
	Department string
	Position   string
	PhotoURL   string
	Todos      []Todo
	ChatGroup  ChatGroup
	ChatGroups []ChatGroup
}

type Session struct {
	ID        int
	UUID      string
	Email     string
	UserID    int
	CreatedAt time.Time
}

func (u *User) CreateUser() (err error) {
	cmd := `insert into users(
		uuid,
		name,
		email,
		password,
		created_at,
		phone,
		department,
		position) values(?, ?, ?, ?, ?, ?, ?,?)`

	_, err = Db.Exec(cmd,
		createUUID(),
		u.Name,
		u.Email,
		Encrypt(u.PassWord),
		time.Now(),
		u.Phone,
		u.Department,
		u.Position)

	if err != nil {
		log.Fatalln(err)
	}

	return err

}

func GetUser(id int) (user User, err error) {
	user = User{}
	cmd := `select id, uuid, name , email, password, created_at, phone, department, position
	from users where id =?`

	err = Db.QueryRow(cmd, id).Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.PassWord,
		&user.CreatedAt,
		&user.Phone,
		&user.Department,
		&user.Position)

	return user, err
}

func GetUsers() (users []User, err error) {
	cmd := `select id, uuid, name , email, password, created_at, phone, department, position
	from users`
	rows, err := Db.Query(cmd)
	if err != nil {
		log.Fatalln(err)
	}

	for rows.Next() {
		var user User
		err = rows.Scan(
			&user.ID,
			&user.UUID,
			&user.Name,
			&user.Email,
			&user.PassWord,
			&user.CreatedAt,
			&user.Phone,
			&user.Department,
			&user.Position)
		if err != nil {
			log.Fatalln(err)
		}
		users = append(users, user)
	}
	rows.Close()

	return users, err
}

func (u *User) UpdateUser() (err error) {
	cmd := `update users set name = ?, email = ? ,phone = ?,  department = ?, position=? where id = ?`
	_, err = Db.Exec(cmd, u.Name, u.Email, u.Phone, u.Department, u.Position, u.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func (u *User) DeleteUser() (err error) {
	cmd := `delete from users where id =?`

	_, err = Db.Exec(cmd, u.ID)
	if err != nil {
		log.Fatalln(err)
	}

	return err
}

func GetUserByEmail(email string) (user User, err error) {
	user = User{}
	cmd := `select id,uuid,name,email,password,created_at from users where email =?`

	err = Db.QueryRow(cmd, email).Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.PassWord,
		&user.CreatedAt)

	return user, err

}

func (u *User) CreateSession() (session Session, err error) {
	session = Session{}
	cmd1 := `insert into sessions
	(uuid, email, user_id, created_at)  values(?,?,?,?)`

	_, err = Db.Exec(cmd1,
		createUUID(), u.Email, u.ID, time.Now())
	if err != nil {
		log.Println(err)
	}

	cmd2 := `select id, uuid, email, user_id, created_at 
	from sessions where user_id =? and email =?`

	err = Db.QueryRow(cmd2, u.ID, u.Email).Scan(
		&session.ID,
		&session.UUID,
		&session.Email,
		&session.UserID,
		&session.CreatedAt)

	return session, err

}

func (sess *Session) CheckSession() (valid bool, err error) {
	cmd := `select id, uuid, email, user_id, created_at 
	from sessions where uuid = ?`

	err = Db.QueryRow(cmd, sess.UUID).Scan(
		&sess.ID,
		&sess.UUID,
		&sess.Email,
		&sess.UserID,
		&sess.CreatedAt)

	if err != nil {
		valid = false
		return
	}
	if sess.ID != 0 {
		valid = true
	}
	return valid, err
}

func (sess *Session) DeleteSessionByUUID() (err error) {
	cmd := `delete from sessions where uuid =?`
	_, err = Db.Exec(cmd, sess.UUID)
	if err != nil {
		log.Println(err)
	}
	return err
}

func (sess *Session) GetUserBySession() (user User, err error) {
	user = User{}
	cmd := `select id, uuid, name, email, created_at, phone, department, position
	from users where id =?`
	err = Db.QueryRow(cmd, sess.UserID).Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.CreatedAt,
		&user.Phone,
		&user.Department,
		&user.Position)
	return user, err
}
