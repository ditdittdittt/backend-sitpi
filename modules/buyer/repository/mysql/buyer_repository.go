package mysql

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/sirupsen/logrus"

	"github.com/ditdittdittt/backend-sitpi/domain"
)

type mysqlBuyerRepository struct {
	Conn *sql.DB
}

func (m *mysqlBuyerRepository) fetch(ctx context.Context, query string, args ...interface{}) (result []domain.Buyer, err error) {
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

	result = make([]domain.Buyer, 0)
	for rows.Next() {
		r := domain.Buyer{}
		err = rows.Scan(
			&r.ID,
			&r.UserID,
			&r.Name,
			&r.Nik,
			&r.Address,
			&r.Status,
			&r.PhoneNumber,
			&r.CreatedAt,
			&r.UpdatedAt,
			&r.UserName,
		)

		if err != nil {
			logrus.Error(err)
			return nil, err
		}

		result = append(result, r)
	}

	return result, nil
}

func (m *mysqlBuyerRepository) inquiry(ctx context.Context, query string, args ...interface{}) (result []domain.Buyer, err error) {

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

	result = make([]domain.Buyer, 0)
	for rows.Next() {
		r := domain.Buyer{}
		err = rows.Scan(
			&r.ID,
			&r.Name,
			&r.Nik,
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

func (m *mysqlBuyerRepository) Fetch(ctx context.Context) (res []domain.Buyer, err error) {
	query := `SELECT b.id, b.user_id, b.name, b.nik, b.address, b.status, b.phone_number, b.created_at, b.updated_at, u.name
			FROM buyer AS b 
			INNER JOIN users AS u ON b.user_id=u.id
			ORDER BY created_at`

	res, err = m.fetch(ctx, query)
	if err != nil {
		return nil, err
	}

	return
}

func (m *mysqlBuyerRepository) GetByID(ctx context.Context, id int64) (res domain.Buyer, err error) {
	query := `SELECT b.id, b.user_id, b.name, b.nik, b.address, b.status, b.phone_number, b.created_at, b.updated_at, u.name
			FROM buyer AS b 
			INNER JOIN users AS u ON b.user_id=u.id
			WHERE b.id = ?`

	list, err := m.fetch(ctx, query, id)
	if err != nil {
		return domain.Buyer{}, err
	}

	if len(list) > 0 {
		res = list[0]
	} else {
		return res, domain.ErrNotFound
	}

	return
}

func (m *mysqlBuyerRepository) Update(ctx context.Context, b *domain.Buyer) (err error) {
	query := `UPDATE buyer SET user_id=?, nik=?, name=?, address=?, status=?, phone_number=?, created_at=?, updated_at=? WHERE ID = ?`

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx, b.UserID, b.Nik, b.Name, b.Address, b.Status, b.PhoneNumber, b.CreatedAt, b.UpdatedAt, b.ID)
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

func (m *mysqlBuyerRepository) Store(ctx context.Context, b *domain.Buyer) (err error) {
	query := `INSERT buyer SET user_id=?, nik=?, name=?, address=?, status=?, phone_number=?, created_at=?, updated_at=?`
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx, b.UserID, b.Nik, b.Name, b.Address, b.Status, b.PhoneNumber, b.CreatedAt, b.UpdatedAt)
	if err != nil {
		return
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		return
	}
	b.ID = lastID
	return
}

func (m *mysqlBuyerRepository) Delete(ctx context.Context, id int64) (err error) {
	query := `DELETE FROM buyer WHERE ID = ?`

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx, id)
	if err != nil {
		return
	}

	rowsAfected, err := res.RowsAffected()
	if err != nil {
		return
	}

	if rowsAfected != 1 {
		err = fmt.Errorf("Weird  Behavior. Total Affected: %d", rowsAfected)
		return
	}

	return
}

func (m *mysqlBuyerRepository) Inquiry(ctx context.Context) (res []domain.Buyer, err error) {
	query := `SELECT id, name, nik, created_at, updated_at FROM buyer`

	res, err = m.inquiry(ctx, query)
	if err != nil {
		return nil, err
	}

	return
}

func NewMysqlBuyerRepository(Conn *sql.DB) domain.BuyerRepository {
	return &mysqlBuyerRepository{Conn: Conn}
}
