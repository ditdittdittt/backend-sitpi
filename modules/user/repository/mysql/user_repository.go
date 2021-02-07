package mysql

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/sirupsen/logrus"

	"github.com/ditdittdittt/backend-sitpi/domain"
)

type mysqlUserRepository struct {
	Conn *sql.DB
}

func (m *mysqlUserRepository) fetch(ctx context.Context, query string, args ...interface{}) (result []domain.User, err error) {
	rows, err := m.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	defer func() {
		errRow := rows.Close()
		if errRow != nil {
			logrus.Error(errRow)
		}
	}()

	result = make([]domain.User, 0)
	for rows.Next() {
		r := domain.User{}
		err = rows.Scan(
			&r.ID,
			&r.RoleID,
			&r.StatusID,
			&r.TpiID,
			&r.Nik,
			&r.Name,
			&r.Address,
			&r.Username,
			&r.Password,
			&r.CreatedAt,
			&r.UpdatedAt,
		)

		if err != nil {
			logrus.Error(err)
			return nil, err
		}

		result = append(result, r)
	}

	return result, nil
}

func (m *mysqlUserRepository) Register(ctx context.Context, u *domain.User) (jwtToken string, err error) {
	panic("implement me")
}

func (m *mysqlUserRepository) Login(ctx context.Context, u *domain.User) {
	panic("implement me")
}

func (m *mysqlUserRepository) GetByID(ctx context.Context, id int64) (res domain.User, err error) {
	panic("implement me")
}

func (m *mysqlUserRepository) GetByUsername(ctx context.Context, username string) (res domain.User, err error) {
	query := `SELECT u.id, u.role_id, u.status_id, u.tpi_id, u.nik, u.name, u.address, u.username, u.password, u.created_at, u.updated_at
		FROM users AS u
		WHERE u.username = ?`

	list, err := m.fetch(ctx, query, username)
	if err != nil {
		return domain.User{}, err
	}

	if len(list) > 0 {
		res = list[0]
	} else {
		return res, domain.ErrNotFound
	}

	return
}

func (m *mysqlUserRepository) ChangePassword(ctx context.Context, newPassword string, id int64) (err error) {
	query := `UPDATE users SET password=? WHERE id=?`

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx, newPassword, id)
	if err != nil {
		return
	}
	affect, err := res.RowsAffected()
	if err != nil {
		return
	}
	if affect != 1 {
		err = fmt.Errorf("Weird  Behavior. Total Affected: %d", affect)
		return
	}

	return
}

func NewMysqlUserRepository(Conn *sql.DB) domain.UserRepository {
	return &mysqlUserRepository{Conn: Conn}
}
