package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/ditdittdittt/backend-sitpi/domain"
	"github.com/ditdittdittt/backend-sitpi/helper"
	"github.com/sirupsen/logrus"
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
			&r.Name,
			&r.Nik,
			&r.Address,
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

func (m *mysqlBuyerRepository) Fetch(ctx context.Context, cursor string, num int64) (res []domain.Buyer, nextCursor string, err error) {
	query := `SELECT * FROM buyer WHERE created_at > ? ORDER BY created_at LIMIT ? `

	decodedCursor, err := helper.DecodeCursor(cursor)
	if err != nil && cursor != "" {
		return nil, "", domain.ErrBadParamInput
	}

	res, err = m.fetch(ctx, query, decodedCursor, num)
	if err != nil {
		return nil, "", err
	}

	if len(res) == int(num) {
		nextCursor = helper.EncodeCursor(res[len(res)-1].CreatedAt)
	}

	return
}

func (m *mysqlBuyerRepository) GetByID(ctx context.Context, id int64) (res domain.Buyer, err error) {
	query := `SELECT * FROM buyer WHERE ID = ?`

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
	query := `UPDATE buyer SET nik=?, name=?, address=?, created_at=?, updated_at=? WHERE ID = ?`

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx, b.Nik, b.Name, b.Address, b.CreatedAt, b.UpdatedAt, b.ID)
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
	query := `INSERT buyer SET nik=?, name=?, address=?, created_at=?, updated_at=?`
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx, b.Nik, b.Name, b.Address, b.CreatedAt, b.UpdatedAt)
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

func NewMysqlBuyerRepository(Conn *sql.DB) domain.BuyerRepository {
	return &mysqlBuyerRepository{Conn: Conn}
}
