package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/ditdittdittt/backend-sitpi/domain"
	"github.com/ditdittdittt/backend-sitpi/helper"
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
			&r.FisherID,
			&r.OfficerID,
			&r.FishTypeID,
			&r.Weight,
			&r.WeightUnit,
			&r.FishingGear,
			&r.FishingArea,
			&r.Price,
			&r.Status,
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

func (m *mysqlAcutionRepository) Fetch(ctx context.Context, cursor string, num int64) (res []domain.Auction, nextCursor string, err error) {
	query := `SELECT * FROM auction WHERE created_at > ? ORDER BY created_at LIMIT ? `

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

func (m *mysqlAcutionRepository) GetByID(ctx context.Context, id int64) (res domain.Auction, err error) {
	query := `SELECT * FROM auction WHERE ID = ?`

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
	query := `UPDATE auction SET tpi_id=?, officer_id=?, fisher_id=?, fish_type_id=?, weight=?, weight_unit=?, fishing_gear=?, fishing_area=?, price=?, status=?, created_at=?, updated_at=? WHERE ID = ?`

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx, a.TpiID, a.OfficerID, a.FisherID, a.FishTypeID, a.Weight, a.WeightUnit, a.FishingGear, a.FishingArea, a.Price, a.Status, a.CreatedAt, a.UpdatedAt, a.ID)
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
	query := `INSERT auction SET tpi_id=?, officer_id=?, fisher_id=?, fish_type_id=?, weight=?, weight_unit=?, fishing_gear=?, fishing_area=?, price=?, status=?, created_at=?, updated_at=?`
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx, a.TpiID, a.OfficerID, a.FisherID, a.FishTypeID, a.Weight, a.WeightUnit, a.FishingGear, a.FishingArea, a.Price, a.Status, a.CreatedAt, a.UpdatedAt)
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
