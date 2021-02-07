package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/ditdittdittt/backend-sitpi/domain"
)

type mysqlCaughtFishRepository struct {
	Conn *sql.DB
}

func (m *mysqlCaughtFishRepository) fetch(ctx context.Context, query string, args ...interface{}) (result []domain.CaughtFish, err error) {
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

	result = make([]domain.CaughtFish, 0)
	for rows.Next() {
		c := domain.CaughtFish{}
		err = rows.Scan(
			&c.ID,
			&c.TripDay,
			&c.Weight,
			&c.WeightUnit,
			&c.FisherNik,
			&c.FisherName,
			&c.FishingGear,
			&c.FishingArea,
			&c.FishType,
			&c.CreatedAt,
			&c.UpdatedAt,
		)

		if err != nil {
			logrus.Error(err)
			return nil, err
		}

		result = append(result, c)
	}

	return result, nil
}

func (m *mysqlCaughtFishRepository) getTotalProduction(ctx context.Context, query string, args ...interface{}) (result []domain.CaughtFish, err error) {
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

	result = make([]domain.CaughtFish, 0)
	for rows.Next() {
		c := domain.CaughtFish{}
		err = rows.Scan(
			&c.TotalProduction,
		)

		if err != nil {
			logrus.Error(err)
			return nil, err
		}

		result = append(result, c)
	}

	return result, nil
}

func (m *mysqlCaughtFishRepository) getTotalFisher(ctx context.Context, query string, args ...interface{}) (result []domain.CaughtFish, err error) {
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

	result = make([]domain.CaughtFish, 0)
	for rows.Next() {
		c := domain.CaughtFish{}
		err = rows.Scan(
			&c.TotalFisher,
		)

		if err != nil {
			logrus.Error(err)
			return nil, err
		}

		result = append(result, c)
	}

	return result, nil
}

func (m *mysqlCaughtFishRepository) Fetch(ctx context.Context, from time.Time, to time.Time, fisherID int64, fishTypeID int64) (res []domain.CaughtFish, err error) {
	query := `SELECT cf.id, cf.trip_day, cf.weight, cf.weight_unit, f.nik, f.name, fg.name, fa.name, ft.name, cf.created_at, cf.updated_at
		FROM caught_fish AS cf
		INNER JOIN fisher AS f ON cf.fisher_id=f.id
		INNER JOIN fishing_gear AS fg ON cf.fishing_gear_id=fg.id
		INNER JOIN fishing_area AS fa ON cf.fishing_area_id=fa.id
		INNER JOIN fish_type AS ft ON cf.fish_type_id=ft.id
		WHERE cf.created_at BETWEEN ? AND ? 
		AND cf.fisher_id = IF (?=0, cf.fisher_id, ?) 
		AND cf.fish_type_id = IF (?=0, cf.fish_type_id, ?)
		ORDER BY cf.created_at`

	res, err = m.fetch(ctx, query, from, to, fisherID, fisherID, fishTypeID, fishTypeID)
	if err != nil {
		return nil, err
	}

	return
}

func (m *mysqlCaughtFishRepository) GetByID(ctx context.Context, id int64) (res domain.CaughtFish, err error) {
	query := `SELECT cf.id, cf.trip_day, cf.weight, cf.weight_unit, f.nik, f.name, fg.name, fa.name, ft.name, cf.created_at, cf.updated_at
		FROM caught_fish AS cf
		INNER JOIN fisher AS f ON cf.fisher_id=f.id
		INNER JOIN fishing_gear AS fg ON cf.fishing_gear_id=fg.id
		INNER JOIN fishing_area AS fa ON cf.fishing_area_id=fa.id
		INNER JOIN fish_type AS ft ON cf.fish_type_id=ft.id
		WHERE cf.id = ?`

	list, err := m.fetch(ctx, query, id)
	if err != nil {
		return domain.CaughtFish{}, err
	}

	if len(list) > 0 {
		res = list[0]
	} else {
		return res, domain.ErrNotFound
	}

	return
}

func (m *mysqlCaughtFishRepository) Update(ctx context.Context, c *domain.CaughtFish) (err error) {
	query := `UPDATE caught_fish SET user_id=?, tpi_id=?, fisher_id=?, fish_type_id=?, weight_unit_id=?, fishing_gear_id=?, fishing_area_id=?, weight=?, trip_day=?, created_at=?, updated_at=? WHERE ID = ?`

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx, c.UserID, c.TpiID, c.FisherID, c.FishTypeID, c.FishingGearID, c.FishingAreaID, c.Weight, c.TripDay, c.CreatedAt, c.UpdatedAt, c.ID)
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

func (m *mysqlCaughtFishRepository) Store(ctx context.Context, c *domain.CaughtFish) (lastID int64, err error) {
	query := `INSERT caught_fish SET user_id=?, tpi_id=?, fisher_id=?, fish_type_id=?, fishing_gear_id=?, fishing_area_id=?, weight=?, weight_unit=?, trip_day=?, created_at=?, updated_at=?`
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx, c.UserID, c.TpiID, c.FisherID, c.FishTypeID, c.FishingGearID, c.FishingAreaID, c.Weight, c.WeightUnit, c.TripDay, c.CreatedAt, c.UpdatedAt)
	if err != nil {
		return
	}

	lastID, err = res.LastInsertId()
	if err != nil {
		return
	}
	c.ID = lastID
	return
}

func (m *mysqlCaughtFishRepository) Delete(ctx context.Context, id int64) (err error) {
	query := `DELETE FROM caught_fish WHERE ID = ?`

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

func (m *mysqlCaughtFishRepository) GetTotalProduction(ctx context.Context, from time.Time, to time.Time) (res domain.CaughtFish, err error) {
	query := `SELECT COALESCE(	
				SUM(
    			CASE
 				WHEN cf.weight_unit_id = 1 THEN cf.weight * 1000
 				WHEN cf.weight_unit_id = 2 THEN cf.weight * 100
    			WHEN cf.weight_unit_id = 3 THEN cf.weight * 1
 				END), 0) AS total
			FROM caught_fish AS cf
			INNER JOIN weight_unit AS wu ON cf.weight_unit_id=wu.id
			WHERE cf.created_at BETWEEN ? AND ?`

	list, err := m.getTotalProduction(ctx, query, from, to)
	if err != nil {
		return domain.CaughtFish{}, err
	}

	if len(list) > 0 {
		res = list[0]
	} else {
		return res, domain.ErrNotFound
	}

	return
}

func (m *mysqlCaughtFishRepository) GetTotalFisher(ctx context.Context, from time.Time, to time.Time) (res domain.CaughtFish, err error) {
	query := `SELECT COUNT(DISTINCT cf.fisher_id)
			FROM caught_fish AS cf
			WHERE cf.created_at BETWEEN ? AND ?`

	list, err := m.getTotalFisher(ctx, query, from, to)
	if err != nil {
		return domain.CaughtFish{}, err
	}

	if len(list) > 0 {
		res = list[0]
	} else {
		return res, domain.ErrNotFound
	}

	return
}

func NewMysqlCaughtFishRepository(Conn *sql.DB) domain.CaughtFishRepository {
	return &mysqlCaughtFishRepository{Conn: Conn}
}
