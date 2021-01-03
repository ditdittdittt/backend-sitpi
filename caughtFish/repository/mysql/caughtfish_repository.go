package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/ditdittdittt/backend-sitpi/domain"
	"github.com/sirupsen/logrus"
)

type mysqlCaughtFishRepository struct {
	Conn *sql.DB
}

func NewMysqlCaughtFishRepository(Conn *sql.DB) domain.CaughtFishRepository {
	return &mysqlCaughtFishRepository{Conn: Conn}
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
			&c.TpiID,
			&c.OfficerID,
			&c.FisherID,
			&c.FishTypeID,
			&c.WeightUnitID,
			&c.FishingGearID,
			&c.Weight,
			&c.FishingArea,
			&c.CreatedAt,
			&c.UpdatedAt,
			&c.WeightUnit,
			&c.FishingGear,
			&c.FisherName,
			&c.FisherNik,
			&c.FishType,
		)

		if err != nil {
			logrus.Error(err)
			return nil, err
		}

		result = append(result, c)
	}

	return result, nil
}

func (m *mysqlCaughtFishRepository) Fetch(ctx context.Context) (res []domain.CaughtFish, err error) {
	query := `SELECT cf.id, cf.tpi_id, cf.officer_id, cf.fisher_id, cf.fish_type_id, cf.weight_unit_id, cf.fishing_gear_id, cf.weight, cf.fishing_area, cf.created_at, cf.updated_at, wu.unit, fg.name, f.name, f.nik, ft.name
		FROM caught_fish AS cf
		INNER JOIN weight_unit AS wu ON cf.weight_unit_id=wu.id
		INNER JOIN fishing_gear AS fg ON cf.fishing_gear_id=fg.id
		INNER JOIN fisher AS f ON cf.fisher_id=f.id
		INNER JOIN fish_type AS ft ON cf.fish_type_id=ft.id
		ORDER BY cf.created_at`

	res, err = m.fetch(ctx, query)
	if err != nil {
		return nil, err
	}

	return
}

func (m *mysqlCaughtFishRepository) GetByID(ctx context.Context, id int64) (res domain.CaughtFish, err error) {
	query := `SELECT cf.id, cf.tpi_id, cf.officer_id, cf.fisher_id, cf.fish_type_id, cf.weight_unit_id, cf.fishing_gear_id, cf.weight, cf.fishing_area, cf.created_at, cf.updated_at, wu.unit, fg.name, f.name, f.nik, ft.name
		FROM caught_fish AS cf
		INNER JOIN weight_unit AS wu ON cf.weight_unit_id=wu.id
		INNER JOIN fishing_gear AS fg ON cf.fishing_gear_id=fg.id
		INNER JOIN fisher AS f ON cf.fisher_id=f.id
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
	query := `UPDATE caught_fish SET tpi_id=?, officer_id=?, fisher_id=?, fish_type_id=?, weight_unit_id=?, fishing_gear_id=?, weight=?, fishing_area=?, created_at=?, updated_at=? WHERE ID = ?`

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx, c.TpiID, c.OfficerID, c.FisherID, c.FishTypeID, c.WeightUnitID, c.FishingGearID, c.Weight, c.FishingArea, c.CreatedAt, c.UpdatedAt, c.ID)
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
	query := `INSERT caught_fish SET tpi_id=?, officer_id=?, fisher_id=?, fish_type_id=?, weight_unit_id=?, fishing_gear_id=?, weight=?, fishing_area=?, created_at=?, updated_at=?`
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx, c.TpiID, c.OfficerID, c.FisherID, c.FishTypeID, c.WeightUnitID, c.FishingGearID, c.Weight, c.FishingArea, c.CreatedAt, c.UpdatedAt)
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
