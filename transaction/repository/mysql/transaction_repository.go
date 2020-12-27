package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/ditdittdittt/backend-sitpi/domain"
	"github.com/sirupsen/logrus"
)

type mysqlTransactionRepository struct {
	Conn *sql.DB
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
			&r.TpiID,
			&r.AuctionID,
			&r.OfficerID,
			&r.BuyerID,
			&r.Price,
			&r.DistributionArea,
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

func (m *mysqlTransactionRepository) Fetch(ctx context.Context) (res []domain.Transaction, err error) {
	query := `SELECT id, tpi_id, officer_id, auction_id, buyer_id, price, distribution_area, created_at, updated_at FROM transaction ORDER BY created_at `

	res, err = m.fetch(ctx, query)
	if err != nil {
		return nil, err
	}

	return
}

func (m *mysqlTransactionRepository) GetByID(ctx context.Context, id int64) (res domain.Transaction, err error) {
	query := `SELECT * FROM transaction WHERE ID = ?`

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
	query := `UPDATE transaction SET tpi_id=?, auction_id=?, officer_id=?, buyer_id=?, distribution_area=?, price=?, created_at=?, updated_at=? WHERE ID = ?`

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx, t.TpiID, t.AuctionID, t.OfficerID, t.BuyerID, t.DistributionArea, t.Price, t.CreatedAt, t.UpdatedAt, t.ID)
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
	query := `INSERT transaction SET tpi_id=?, auction_id=?, officer_id=?, buyer_id=?, distribution_area=?, price=?, created_at=?, updated_at=?`
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx, t.TpiID, t.AuctionID, t.OfficerID, t.BuyerID, t.DistributionArea, t.Price, t.CreatedAt, t.UpdatedAt)
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

func NewMysqlTransactionRepository(Conn *sql.DB) domain.TransactionRepository {
	return &mysqlTransactionRepository{Conn: Conn}
}
