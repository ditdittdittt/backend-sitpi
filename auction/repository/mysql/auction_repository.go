package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/ditdittdittt/backend-sitpi/domain"
	"github.com/sirupsen/logrus"
)

type mysqlAcutionRepository struct {
	Conn *sql.DB
}

func NewMysqlAuctionRepository(Conn *sql.DB) domain.AuctionRepository {
	return &mysqlAcutionRepository{Conn: Conn}
}

func (m *mysqlAcutionRepository) fetch(ctx context.Context, query string, args ...interface{}) (result []domain.Auction, err error) {
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

	result = make([]domain.Auction, 0)
	for rows.Next() {
		r := domain.Auction{}
		err = rows.Scan(
			&r.ID,
			&r.TpiID,
			&r.OfficerID,
			&r.CaughtFishID,
			&r.Weight,
			&r.WeightUnit,
			&r.Status,
			&r.CreatedAt,
			&r.UpdatedAt,
			&r.FisherName,
			&r.FishType,
			&r.StatusName,
		)

		if err != nil {
			logrus.Error(err)
			return nil, err
		}

		result = append(result, r)
	}

	return result, nil
}

func (m *mysqlAcutionRepository) Fetch(ctx context.Context) (res []domain.Auction, err error) {
	query := `SELECT a.id, a.tpi_id, a.officer_id, a.caught_fish_id, a.weight, a.weight_unit, a.status, a.created_at, a.updated_at, f.name, ft.name, s.status
		FROM auction AS a
		INNER JOIN caught_fish AS cf ON a.caught_fish_id=cf.id
		INNER JOIN fisher AS f ON cf.fisher_id=f.id
		INNER JOIN fish_type AS ft ON cf.fish_type_id=ft.id
		INNER JOIN auction_status AS s ON a.status=s.id
		ORDER BY a.created_at`

	res, err = m.fetch(ctx, query)
	if err != nil {
		return nil, err
	}

	return
}

func (m *mysqlAcutionRepository) GetByID(ctx context.Context, id int64) (res domain.Auction, err error) {
	query := `SELECT a.id, a.tpi_id, a.officer_id, a.caught_fish_id, a.weight, a.weight_unit, a.status, a.created_at, a.updated_at, f.name, ft.name, s.status
		FROM auction AS a
		INNER JOIN caught_fish AS cf ON a.caught_fish_id=cf.id
		INNER JOIN fisher AS f ON cf.fisher_id=f.id
		INNER JOIN fish_type AS ft ON cf.fish_type_id=ft.id
		INNER JOIN auction_status AS s ON a.status=s.id
		WHERE a.id=?`

	list, err := m.fetch(ctx, query, id)
	if err != nil {
		return domain.Auction{}, err
	}

	if len(list) > 0 {
		res = list[0]
	} else {
		return res, domain.ErrNotFound
	}

	return
}

func (m *mysqlAcutionRepository) Update(ctx context.Context, a *domain.Auction) (err error) {
	query := `UPDATE auction SET tpi_id=?, officer_id=?, caught_fish_id=?, weight=?, weight_unit=?, status=?, created_at=?, updated_at=? WHERE ID = ?`

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx, a.TpiID, a.OfficerID, a.CaughtFishID, a.Weight, a.WeightUnit, a.Status, a.CreatedAt, a.UpdatedAt, a.ID)
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

func (m *mysqlAcutionRepository) Store(ctx context.Context, a *domain.Auction) (err error) {
	query := `INSERT auction SET tpi_id=?, officer_id=?, caught_fish_id=?, weight=?, weight_unit=?, status=?, created_at=?, updated_at=?`
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx, a.TpiID, a.OfficerID, a.CaughtFishID, a.Weight, a.WeightUnit, a.Status, a.CreatedAt, a.UpdatedAt)
	if err != nil {
		return
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		return
	}
	a.ID = lastID
	return
}

func (m *mysqlAcutionRepository) Delete(ctx context.Context, id int64) (err error) {
	query := `DELETE FROM auction WHERE ID = ?`

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

func (m *mysqlAcutionRepository) UpdateStatus(ctx context.Context, id int64) (err error) {
	query := `UPDATE auction SET status=2 WHERE id=?`

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx, id)
	if err != nil {
		return
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return
	}

	if rowsAffected != 1 {
		err = fmt.Errorf("Weird  Behavior. Total Affected: %d", rowsAffected)
		return
	}

	return
}
