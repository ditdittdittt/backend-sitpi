package mysql

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/sirupsen/logrus"

	"github.com/ditdittdittt/backend-sitpi/domain"
)

type mysqlFishingGearRepository struct {
	Conn *sql.DB
}

func (m *mysqlFishingGearRepository) fetch(ctx context.Context, query string, args ...interface{}) (result []domain.FishingGear, err error) {
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

	result = make([]domain.FishingGear, 0)
	for rows.Next() {
		r := domain.FishingGear{}
		err = rows.Scan(
			&r.ID,
			&r.Name,
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

func (m *mysqlFishingGearRepository) Fetch(ctx context.Context) (res []domain.FishingGear, err error) {
	query := `SELECT id, name, created_at, updated_at FROM fishing_gear ORDER BY name`

	res, err = m.fetch(ctx, query)
	if err != nil {
		return nil, err
	}

	return
}

func (m *mysqlFishingGearRepository) GetByID(ctx context.Context, id int64) (res domain.FishingGear, err error) {
	query := `SELECT id, name, created_at, updated_at FROM fishing_gear WHERE ID = ?`

	list, err := m.fetch(ctx, query, id)
	if err != nil {
		return domain.FishingGear{}, err
	}

	if len(list) > 0 {
		res = list[0]
	} else {
		return res, domain.ErrNotFound
	}

	return
}

func (m *mysqlFishingGearRepository) Update(ctx context.Context, fg *domain.FishingGear) (err error) {
	query := `UPDATE fishing_gear SET name=?, created_at=?, updated_at=? WHERE ID = ?`

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx, fg.Name, fg.CreatedAt, fg.UpdatedAt, fg.ID)
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

func (m *mysqlFishingGearRepository) Store(ctx context.Context, fg *domain.FishingGear) (err error) {
	query := `INSERT fishing_gear SET name=?, created_at=?, updated_at=?`

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx, fg.Name, fg.CreatedAt, fg.UpdatedAt)
	if err != nil {
		return
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		return
	}
	fg.ID = lastID
	return
}

func (m *mysqlFishingGearRepository) Delete(ctx context.Context, id int64) (err error) {
	query := `DELETE FROM fishing_gear WHERE ID = ?`

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

func NewMysqlFishingGearRepository(Conn *sql.DB) domain.FishingGearRepository {
	return &mysqlFishingGearRepository{Conn: Conn}
}
