package mysql_test

import (
	"bou.ke/monkey"
	"context"
	"github.com/ditdittdittt/backend-sitpi/buyer/repository/mysql"
	"github.com/ditdittdittt/backend-sitpi/domain"
	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"testing"
	"time"
)

func TestMysqlBuyerRepository_Fetch(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening buyerRepository stub database connection", err)
	}

	monkey.Patch(time.Now, func() time.Time {
		return time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	})

	mockBuyer := []domain.Buyer{
		{
			ID:        1,
			Nik:       "09821312",
			Name:      "buyer test",
			Address:   "buyer address",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:        2,
			Nik:       "12345678",
			Name:      "buyer test 2",
			Address:   "buyer address 2",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	rows := sqlmock.NewRows([]string{"id", "nik", "name", "address", "created_at", "updated_at"}).
		AddRow(mockBuyer[0].ID, mockBuyer[0].Nik, mockBuyer[0].Name, mockBuyer[0].Address, mockBuyer[0].CreatedAt, mockBuyer[0].UpdatedAt).
		AddRow(mockBuyer[1].ID, mockBuyer[1].Nik, mockBuyer[1].Name, mockBuyer[0].Address, mockBuyer[1].CreatedAt, mockBuyer[1].UpdatedAt)

	query := `SELECT id, name, nik, address, created_at, updated_at FROM buyer`

	mock.ExpectQuery(query).WillReturnRows(rows)
	buyerRepository := mysql.NewMysqlBuyerRepository(db)
	list, err := buyerRepository.Fetch(context.TODO())
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening buyerRepository stub database connection", err)
	}
	assert.NoError(t, err)
	assert.Len(t, list, 2)
}

func TestMysqlBuyerRepository_GetByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	monkey.Patch(time.Now, func() time.Time {
		return time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	})

	rows := sqlmock.NewRows([]string{"id", "nik", "name", "address", "created_at", "updated_at"}).
		AddRow(1, "12345678", "buyer name", "buyer address", time.Now(), time.Now())

	query := "SELECT id, name, nik, address, created_at, updated_at FROM buyer WHERE ID = \\?"

	mock.ExpectQuery(query).WillReturnRows(rows)
	buyerRepository := mysql.NewMysqlBuyerRepository(db)

	num := int64(1)
	buyer, err := buyerRepository.GetByID(context.TODO(), num)
	assert.NoError(t, err)
	assert.NotNil(t, buyer)
}

func TestMysqlBuyerRepository_Store(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	now := time.Now()
	buyer := &domain.Buyer{
		Nik:       "12345678",
		Name:      "buyer name",
		Address:   "buyer address",
		CreatedAt: now,
		UpdatedAt: now,
	}

	query := "INSERT buyer SET nik=\\?, name=\\?, address=\\?, created_at=\\?, updated_at=\\?"
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(buyer.Nik, buyer.Name, buyer.Address, buyer.CreatedAt, buyer.UpdatedAt).WillReturnResult(sqlmock.NewResult(12, 1))

	buyerRepository := mysql.NewMysqlBuyerRepository(db)
	err = buyerRepository.Store(context.TODO(), buyer)
	assert.NoError(t, err)
	assert.Equal(t, int64(12), buyer.ID)
}

func TestMysqlBuyerRepository_Update(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	now := time.Now()
	buyer := &domain.Buyer{
		ID:        12,
		Nik:       "12345678",
		Name:      "buyer name",
		Address:   "buyer address",
		CreatedAt: now,
		UpdatedAt: now,
	}

	query := "UPDATE buyer SET nik=\\?, name=\\?, address=\\?, created_at=\\?, updated_at=\\? WHERE ID = \\?"

	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(buyer.Nik, buyer.Name, buyer.Address, buyer.CreatedAt, buyer.UpdatedAt, buyer.ID).WillReturnResult(sqlmock.NewResult(12, 1))

	buyerRepository := mysql.NewMysqlBuyerRepository(db)
	err = buyerRepository.Update(context.TODO(), buyer)
	assert.NoError(t, err)
}

func TestMysqlBuyerRepository_Delete(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	query := "DELETE FROM buyer WHERE ID = \\?"

	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(12).WillReturnResult(sqlmock.NewResult(12, 1))

	buyerRepository := mysql.NewMysqlBuyerRepository(db)

	num := int64(12)
	err = buyerRepository.Delete(context.TODO(), num)
	assert.NoError(t, err)
}
