// DBとの操作を行う

package repository

import (
	"database/sql"
	"github.com/yrk9/cardboard_task_app/model"
	"time"
)

type UserRepository interface {
	CreateUser(user *model.User) error
	FindByEmail(email string) (*model.User, error)
}

type userRepository struct {
	db *sql.DB
}

// コンストラクタ
func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

// ユーザアカウントの作成メソッド

func (r userRepository) CreateUser(user *model.User) error {
	result, err := r.db.Exec(
		`INSERT INTO users (email, password_hash, created_at, updated_at) VALUES (?, ?, ?, ?)`,
		user.Email, 
		user.PasswordHash, 
		time.Now(), 
		time.Now(),
	)
	
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	user.ID = id
	return nil
}

// Emailからユーザを探すメソッド

func (r userRepository) FindByEmail(email string) (*model.User, error) {
	user := &model.User{}

	row := r.db.QueryRow(`SELECT id, email, password_hash, created_at, updated_at FROM users WHERE email = ?`, email)
	err := row.Scan(
		&user.ID,
		&user.Email,
		&user.PasswordHash,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return user, nil
}