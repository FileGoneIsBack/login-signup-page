package database

import (
	"golang.org/x/crypto/bcrypt"
    "database/sql"
    "errors"
)

type User struct {
	ID                                      int
	Username                                string
	Password								string
	Email								    string
}

var (
	ErrDuplicateUser    = errors.New("duplicate user")
	ErrUserNotFound     = errors.New("user couldn't be found in the database")
	ErrInvalidPassword  = errors.New("invalid password")
)

func (conn *Instance) NewUser(user *User) error {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        return err
    }
    
    _, err = Container.conn.Exec("INSERT INTO users (username, password_hash, email) VALUES (?, ?, ?)", user.Username, hashedPassword, user.Email)
    if err != nil {
        return err
    }
    
    return nil
}
func (conn *Instance) AuthenticateUser(username, password string) (*User, error) {
    // Retrieve hashed password from database based on username
    var hashedPassword string
    var user User

    err := Container.conn.QueryRow("SELECT id, username, password_hash, email FROM users WHERE username = ?", username).Scan(&user.ID, &user.Username, &hashedPassword, &user.Email)
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, ErrUserNotFound
        }
        return nil, err
    }

    // Compare the provided password with the hashed password from the database
    err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
    if err != nil {
        return nil, ErrInvalidPassword
    }

    // Clear the password field for security reasons before returning user data
    user.Password = ""

    return &user, nil
}

