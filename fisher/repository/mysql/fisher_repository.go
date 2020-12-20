package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/ditdittdittt/backend-sitpi/domain"
	"github.com/ditdittdittt/backend-sitpi/helper"
	"github.com/sirupsen/logrus"
)

type mysqlFisherRepository struct {
	Conn *sql.DB
}

func NewMysqlFisherRepository(Conn *sql.DB) domain.FisherRepository {
	return &mysqlFisherRepository{Conn: Conn}
}

func (m *mysqlFisherRepository) fetch(ctx context.Context, query string, args ...interface{}) (result []domain.Fisher, err error) {
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

	result = make([]domain.Fisher, 0)
	for rows.Next() {
		r := domain.Fisher{}
		err = rows.Scan(
			&r.ID,
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

func (m *mysqlFisherRepository) Fetch(ctx context.Context, cursor string, num int64) (res []domain.Fisher, nextCursor string, err error) {
	query := `SELECT * FROM fisher WHERE created_at > ? ORDER BY created_at LIMIT ? `

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

func (m *mysqlFisherRepository) GetByID(ctx context.Context, id int64) (res domain.Fisher, err error) {
	query := `SELECT * FROM fisher WHERE ID = ?`

	list, err := m.fetch(ctx, query, id)
	if err != nil {
		return domain.Fisher{}, err
	}

	if len(list) > 0 {
		res = list[0]
	} else {
		return res, domain.ErrNotFound
	}

	return
}

func (m *mysqlFisherRepository) Update(ctx context.Context, f *domain.Fisher) (err error) {
	query := `UPDATE fisher SET nik=?, address=?, created_at=?, updated_at=? WHERE ID = ?`

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx, f.Nik, f.Address, f.CreatedAt, f.UpdatedAt, f.ID)
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

func (m *mysqlFisherRepository) Store(ctx context.Context, f *domain.Fisher) (err error) {
	query := `INSERT fisher SET nik=?, address=?, created_at=?, updated_at=?`
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx, f.Nik, f.Address, f.CreatedAt, f.UpdatedAt)
	if err != nil {
		return
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		return
	}
	f.ID = lastID
	return
}

func (m *mysqlFisherRepository) Delete(ctx context.Context, id int64) (err error) {
	query := `DELETE FROM fisher WHERE ID = ?`

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
