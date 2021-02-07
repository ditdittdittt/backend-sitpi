package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/ditdittdittt/backend-sitpi/domain"
)

type mysqlTransactionRepository struct {
	Conn *sql.DB
}

func (m *mysqlTransactionRepository) getTotalBuyer(ctx context.Context, query string, args ...interface{}) (result []domain.Transaction, err error) {
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

	result = make([]domain.Transaction, 0)
	for rows.Next() {
		c := domain.Transaction{}
		err = rows.Scan(
			&c.TotalBuyer,
		)

		if err != nil {
			logrus.Error(err)
			return nil, err
		}

		result = append(result, c)
	}

	return result, nil
}

func (m *mysqlTransactionRepository) fetch(ctx context.Context, query string, args ...interface{}) (result []domain.Transaction, err error) {
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

	result = make([]domain.Transaction, 0)
	for rows.Next() {
		r := domain.Transaction{}
		err = rows.Scan(
			&r.ID,
			&r.AuctionID,
			&r.DistributionArea,
			&r.Price,
			&r.CreatedAt,
			&r.UpdatedAt,
			&r.BuyerName,
			&r.FisherName,
			&r.FishType,
			&r.Weight,
			&r.WeightUnit,
		)

		if err != nil {
			logrus.Error(err)
			return nil, err
		}

		result = append(result, r)
	}

	return result, nil
}

func (m *mysqlTransactionRepository) Fetch(ctx context.Context, from time.Time, to time.Time, buyerID int64, fishTypeID int64) (res []domain.Transaction, err error) {
	query := `SELECT t.id, t.auction_id, t.distribution_area, t.price, t.created_at, t.updated_at, b.name, f.name, ft.name, cf.weight, cf.weight_unit
		FROM transaction AS t
		INNER JOIN auction AS a ON t.auction_id=a.id
		INNER JOIN caught_fish AS cf ON a.caught_fish_id=cf.id
		INNER JOIN buyer AS b ON t.buyer_id=b.id
		INNER JOIN fisher AS f ON cf.fisher_id=f.id
		INNER JOIN fish_type AS ft ON cf.fish_type_id=ft.id
		WHERE t.created_at BETWEEN ? AND ? 
		AND t.buyer_id = IF (?=0, t.buyer_id, ?) 
		AND cf.fish_type_id = IF (?=0, cf.fish_type_id, ?)
		ORDER BY t.created_at `

	res, err = m.fetch(ctx, query, from, to, buyerID, buyerID, fishTypeID, fishTypeID)
	if err != nil {
		return nil, err
	}

	return
}

func (m *mysqlTransactionRepository) GetByID(ctx context.Context, id int64) (res domain.Transaction, err error) {
	query := `SELECT t.id, t.auction_id, t.distribution_area, t.price, t.created_at, t.updated_at, b.name, f.name, ft.name, cf.weight, cf.weight_unit
		FROM transaction AS t
		INNER JOIN auction AS a ON t.auction_id=a.id
		INNER JOIN caught_fish AS cf ON a.caught_fish_id=cf.id
		INNER JOIN buyer AS b ON t.buyer_id=b.id
		INNER JOIN fisher AS f ON cf.fisher_id=f.id
		INNER JOIN fish_type AS ft ON cf.fish_type_id=ft.id
		WHERE t.id=?
		`

	list, err := m.fetch(ctx, query, id)
	if err != nil {
		return domain.Transaction{}, err
	}

	if len(list) > 0 {
		res = list[0]
	} else {
		return res, domain.ErrNotFound
	}

	return
}

func (m *mysqlTransactionRepository) Update(ctx context.Context, t *domain.Transaction) (err error) {
	query := `UPDATE transaction SET user_id=?, tpi_id=?, auction_id=?, buyer_id=?, distribution_area=?, price=?, created_at=?, updated_at=? WHERE ID = ?`

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx, t.UserID, t.TpiID, t.AuctionID, t.BuyerID, t.DistributionArea, t.Price, t.CreatedAt, t.UpdatedAt, t.ID)
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

func (m *mysqlTransactionRepository) Store(ctx context.Context, t *domain.Transaction) (err error) {
	query := `INSERT transaction SET user_id=?, tpi_id=?, auction_id=?, buyer_id=?, distribution_area=?, price=?, created_at=?, updated_at=?`
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx, t.UserID, t.TpiID, t.AuctionID, t.BuyerID, t.DistributionArea, t.Price, t.CreatedAt, t.UpdatedAt)
	if err != nil {
		return
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		return
	}
	t.ID = lastID
	return
}

func (m *mysqlTransactionRepository) Delete(ctx context.Context, id int64) (err error) {
	query := `DELETE FROM transaction WHERE ID = ?`

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

func (m *mysqlTransactionRepository) GetTotalBuyer(ctx context.Context, from time.Time, to time.Time) (res domain.Transaction, err error) {
	query := `SELECT COUNT(DISTINCT t.buyer_id)
			FROM transaction AS t
			WHERE t.created_at BETWEEN ? AND ?`

	list, err := m.getTotalBuyer(ctx, query, from, to)
	if err != nil {
		return domain.Transaction{}, err
	}

	if len(list) > 0 {
		res = list[0]
	} else {
		return res, domain.ErrNotFound
	}

	return
}

func NewMysqlTransactionRepository(Conn *sql.DB) domain.TransactionRepository {
	return &mysqlTransactionRepository{Conn: Conn}
}
