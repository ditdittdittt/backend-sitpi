package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/ditdittdittt/backend-sitpi/domain"
	"github.com/sirupsen/logrus"
)

type mysqlWeightUnitRepository struct {
	Conn *sql.DB
}

func (m *mysqlWeightUnitRepository) fetch(ctx context.Context, query string, args ...interface{}) (result []domain.WeightUnit, err error) {
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

	result = make([]domain.WeightUnit, 0)
	for rows.Next() {
		r := domain.WeightUnit{}
		err = rows.Scan(
			&r.ID,
			&r.Unit,
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

func (m *mysqlWeightUnitRepository) Fetch(ctx context.Context) (res []domain.WeightUnit, err error) {
	query := `SELECT id, unit, created_at, updated_at FROM weight_unit`

	res, err = m.fetch(ctx, query)
	if err != nil {
		return nil, err
	}

	return
}

func (m *mysqlWeightUnitRepository) GetByID(ctx context.Context, id int64) (res domain.WeightUnit, err error) {
	query := `SELECT id, unit, created_at, updated_at FROM weight_unit WHERE ID = ?`

	list, err := m.fetch(ctx, query, id)
	if err != nil {
		return domain.WeightUnit{}, err
	}

	if len(list) > 0 {
		res = list[0]
	} else {
		return res, domain.ErrNotFound
	}

	return
}

func (m *mysqlWeightUnitRepository) Update(ctx context.Context, wu *domain.WeightUnit) (err error) {
	query := `UPDATE weight_unit SET unit=?, created_at=?, updated_at=? WHERE ID = ?`

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx, wu.Unit, wu.CreatedAt, wu.UpdatedAt, wu.ID)
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

func (m *mysqlWeightUnitRepository) Store(ctx context.Context, wu *domain.WeightUnit) (err error) {
	query := `INSERT weight_unit SET unit=?, created_at=?, updated_at=?`

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx, wu.Unit, wu.CreatedAt, wu.UpdatedAt)
	if err != nil {
		return
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		return
	}
	wu.ID = lastID
	return
}

func (m *mysqlWeightUnitRepository) Delete(ctx context.Context, id int64) (err error) {
	query := `DELETE FROM weight_unit WHERE ID = ?`

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

func NewMysqlWeightUnitRepository(Conn *sql.DB) domain.WeightUnitRepository {
	return &mysqlWeightUnitRepository{Conn: Conn}
}
