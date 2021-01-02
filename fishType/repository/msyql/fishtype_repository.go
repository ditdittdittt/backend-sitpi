package msyql

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/ditdittdittt/backend-sitpi/domain"
	"github.com/sirupsen/logrus"
)

type mysqlFishTypeRepository struct {
	Conn *sql.DB
}

func NewMysqlFishTypeRepository(Conn *sql.DB) domain.FishTypeRepository {
	return &mysqlFishTypeRepository{Conn: Conn}
}

func (m *mysqlFishTypeRepository) fetch(ctx context.Context, query string, args ...interface{}) (result []domain.FishType, err error) {
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

	result = make([]domain.FishType, 0)
	for rows.Next() {
		r := domain.FishType{}
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

func (m *mysqlFishTypeRepository) Fetch(ctx context.Context) (res []domain.FishType, err error) {
	query := `SELECT id, name, created_at, updated_at FROM fish_type ORDER BY name`

	res, err = m.fetch(ctx, query)
	if err != nil {
		return nil, err
	}

	return
}

func (m *mysqlFishTypeRepository) GetByID(ctx context.Context, id int64) (res domain.FishType, err error) {
	query := `SELECT id, name, created_at, updated_at FROM fish_type WHERE ID = ?`

	list, err := m.fetch(ctx, query, id)
	if err != nil {
		return domain.FishType{}, err
	}

	if len(list) > 0 {
		res = list[0]
	} else {
		return res, domain.ErrNotFound
	}

	return

}

func (m *mysqlFishTypeRepository) Update(ctx context.Context, ft *domain.FishType) (err error) {
	query := `UPDATE fish_type SET name=?, created_at=?, update_at=? WHERE ID = ?`

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx, ft.Name, ft.CreatedAt, ft.UpdatedAt, ft.ID)
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

func (m *mysqlFishTypeRepository) Store(ctx context.Context, ft *domain.FishType) (err error) {
	query := `INSERT fish_type SET name=?, created_at=?, updated_at=?`

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx, ft.Name, ft.CreatedAt, ft.UpdatedAt)
	if err != nil {
		return
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		return
	}
	ft.ID = lastID
	return
}

func (m *mysqlFishTypeRepository) Delete(ctx context.Context, id int64) (err error) {
	query := `DELETE FROM fish_type WHERE ID = ?`

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
