package user_pg

import (
	"context"
	"database/sql"
	"log"

	"github.com/adityatresnobudi/bank-teller-service-go/internal/entity"
	"github.com/adityatresnobudi/bank-teller-service-go/internal/repositories/user_repo"
	"github.com/adityatresnobudi/bank-teller-service-go/pkg/errs"
	"github.com/google/uuid"
)

type userPG struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) user_repo.Repository {
	return &userPG{
		db: db,
	}
}

func (u *userPG) GetAll(ctx context.Context) ([]entity.User, errs.MessageErr) {
	rows, err := u.db.QueryContext(ctx, GET_ALL_USER)

	if err != nil {
		log.Printf("db get all todos: %s\n", err.Error())
		return nil, errs.NewInternalServerError()
	}

	users := []entity.User{}

	for rows.Next() {
		user := entity.User{}

		if err = rows.Scan(
			&user.Id,
			&user.Name,
			&user.PhoneNumber,
			&user.Password,
			&user.Role,
			&user.Email,
			&user.CreatedAt,
			&user.UpdatedAt,
		); err != nil {
			log.Printf("db scan get all users: %s\n", err.Error())
			return nil, errs.NewInternalServerError()
		}

		users = append(users, user)
	}

	return users, nil
}
func (u *userPG) GetOneById(ctx context.Context, id uuid.UUID) (*entity.User, errs.MessageErr) {
	user := entity.User{}

	if err := u.db.QueryRowContext(
		ctx,
		GET_USER_BY_ID,
		id,
	).Scan(
		&user.Id,
		&user.Name,
		&user.PhoneNumber,
		&user.Password,
		&user.Role,
		&user.Email,
		&user.CreatedAt,
		&user.UpdatedAt,
	); err != nil {
		log.Printf("db scan get one user by id: %s\n", err.Error())
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("user was not found")
		}
		return nil, errs.NewInternalServerError()
	}

	return &user, nil
}
func (u *userPG) GetOneByEmail(ctx context.Context, email string) (*entity.User, errs.MessageErr) {
	user := entity.User{}

	if err := u.db.QueryRowContext(
		ctx,
		GET_USER_BY_EMAIL,
		email,
	).Scan(
		&user.Id,
		&user.Name,
		&user.PhoneNumber,
		&user.Password,
		&user.Role,
		&user.Email,
		&user.CreatedAt,
		&user.UpdatedAt,
	); err != nil {
		log.Printf("db scan get one user by email: %s\n", err.Error())
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("user was not found")
		}
		return nil, errs.NewInternalServerError()
	}

	return &user, nil
}
func (u *userPG) Create(ctx context.Context, user entity.User) (errs.MessageErr) {
	if _, err := u.db.ExecContext(
		ctx,
		INSERT_USER,
		user.Name,
		user.PhoneNumber,
		user.Password,
		user.Email,
	); err != nil {
		log.Printf("db scan create user: %s\n", err.Error())
		return errs.NewInternalServerError()
	}

	return nil
}
func (u *userPG) UpdateById(ctx context.Context, user entity.User) (*entity.User, errs.MessageErr) {
	updatedUser := entity.User{}

	if err := u.db.QueryRowContext(
		ctx,
		UPDATE_USER,
		user.Name,
		user.PhoneNumber,
		user.Password,
		user.Email,
		user.Role,
		user.Id,
	).Scan(
		&updatedUser.Id,
		&updatedUser.Name,
		&updatedUser.PhoneNumber,
		&updatedUser.Password,
		&updatedUser.Role,
		&updatedUser.Email,
		&updatedUser.CreatedAt,
		&updatedUser.UpdatedAt,
	); err != nil {
		log.Printf("db scan update user by id: %s\n", err.Error())
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("user was not found")
		}
		return nil, errs.NewInternalServerError()
	}

	return &updatedUser, nil
}
func (u *userPG) DeleteById(ctx context.Context, id uuid.UUID) errs.MessageErr {
	return nil
}
