package repositories

import (
	"context"
	"practice-go/models"
)

func (r *repository) FindByEmail(ctx context.Context, email string) (models.User, error) {
	user := models.User{}
	query := `SELECT * 
				FROM users
				WHERE users.email = ?
				LIMIT 1`

	result, err := r.db.QueryContext(ctx, query, email)

	if err != nil {
		return user, err
	}

	defer result.Close()

	if result.Next() {
		result.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.Password,
			&user.Phone,
		)
		return user, nil
	}

	return user, nil
}

func (r *repository) CreateUser(ctx context.Context, user models.User) (models.User, error) {
	query := `INSERT INTO users(name, email, password, phone)
			  	VALUES (?, ?, ?, ?)`

	result, error := r.db.ExecContext(ctx, query, user.Name, user.Email, user.Password, user.Phone)

	if error != nil {
		return user, error
	}

	id, error := result.LastInsertId()

	if error != nil {
		return user, error
	}

	user.ID = int(id)

	return user, nil
}

func (r *repository) UpdateUser(id int, user models.User) (bool, error) {
	query := `UPDATE FROM users
				WHERE id = ?
				SET name = ?, email = ?, password = ?, phone = ?`

	_, error := r.db.Exec(query, id, user.Name, user.Email, user.Password, user.Phone)

	if error != nil {
		return false, error
	}

	return true, nil
}
