package mysql

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/sirupsen/logrus"

	"github.com/ditdittdittt/backend-sitpi/domain"
)

type mysqlFishingAreaRepository struct {
	Conn *sql.DB
}

func (m *mysqlFishingAreaRepository) fetch(ctx context.Context, query string, args ...interface{}) (result []domain.FishingArea, err error) {
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

	result = make([]domain.FishingArea, 0)
	for rows.Next() {
		r := domain.FishingArea{}
		err = rows.Scan(
			&r.ID,
			&r.DistrictID,
			&r.Name,
			&r.SouthLatitudeDegree,
			&r.SouthLatitudeMinute,
			&r.SouthLatitudeSecond,
			&r.EastLongitudeDegree,
			&r.EastLongitudeMinute,
			&r.EastLongitudeSecond,
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

func (m *mysqlFishingAreaRepository) Fetch(ctx context.Context) (res []domain.FishingArea, err error) {
	query := `SELECT id, district_id, name, south_latitude_degree, south_latitude_minute, south_latitude_second, east_longitude_degree, east_longitude_minute, east_longitude_second, created_at, updated_at
		FROM fishing_area`

	res, err = m.fetch(ctx, query)
	if err != nil {
		return nil, err
	}

	return
}

func (m *mysqlFishingAreaRepository) GetByID(ctx context.Context, id int64) (res domain.FishingArea, err error) {
	query := `SELECT id, district_id, name, south_latitude_degree, south_latitude_minute, south_latitude_second, east_longitude_degree, east_longitude_minute, east_longitude_second, created_at, updated_at
		FROM fishing_area 
		WHERE id = ?`

	list, err := m.fetch(ctx, query, id)
	if err != nil {
		return domain.FishingArea{}, err
	}

	if len(list) > 0 {
		res = list[0]
	} else {
		return res, domain.ErrNotFound
	}

	return
}

func (m *mysqlFishingAreaRepository) Store(ctx context.Context, fishingArea *domain.FishingArea) (err error) {
	query := `INSERT fishing_area SET name=?, district_id=?, south_latitude_degree=?, south_latitude_minute=?, south_latitude_second=?, east_longitude_degree=?, east_longitude_minute=?, east_longitude_second=?, created_at=?, updated_at=?`

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx, fishingArea.Name, fishingArea.DistrictID, fishingArea.SouthLatitudeDegree, fishingArea.SouthLatitudeMinute, fishingArea.SouthLatitudeSecond, fishingArea.EastLongitudeDegree, fishingArea.EastLongitudeMinute, fishingArea.EastLongitudeSecond, fishingArea.CreatedAt, fishingArea.UpdatedAt)
	if err != nil {
		return
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		return
	}
	fishingArea.ID = lastID
	return

}

func (m *mysqlFishingAreaRepository) Update(ctx context.Context, fishingArea *domain.FishingArea) (err error) {
	query := `UPDATE fishing_area SET name=?, district_id=?, south_latitude_degree=?, south_latitude_minute=?, south_latitude_second=?, east_longitude_degree=?, east_longitude_minute=?, east_longitude_second=?, created_at=?, updated_at=? WHERE ID = ?`

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx, fishingArea.Name, fishingArea.DistrictID, fishingArea.SouthLatitudeDegree, fishingArea.SouthLatitudeMinute, fishingArea.SouthLatitudeSecond, fishingArea.EastLongitudeDegree, fishingArea.EastLongitudeMinute, fishingArea.EastLongitudeSecond, fishingArea.CreatedAt, fishingArea.UpdatedAt, fishingArea.ID)
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

func (m *mysqlFishingAreaRepository) Delete(ctx context.Context, id int64) (err error) {
	query := `DELETE FROM fishing_area WHERE ID = ?`

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

func NewFishingAreRepository(Conn *sql.DB) domain.FishingAreaRepository {
	return &mysqlFishingAreaRepository{Conn: Conn}
}
