package db

import (
	"context"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/byvinesse/vinance-backend/entity"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func TestAccount_InsertOne(t *testing.T) {
	// Create mock DB
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// Convert to sqlx.DB
	sqlxDB := sqlx.NewDb(db, "sqlmock")
	repo := NewAccount(sqlxDB)

	// Test data
	ctx := context.Background()
	now := time.Now()

	account := &entity.Account{
		UserID:        "test-user-id",
		Name:          "Test Account",
		Balance:       1000,
		Currency:      entity.CurrencyTypeIDR,
		Type:          entity.AccountTypeCash,
		Color:         "#000",
		IsArchived:    false,
		IsExcluded:    false,
		MarkForDelete: false,
		CreatedAt:     now,
		UpdatedAt:     now,
	}

	t.Run("Success", func(t *testing.T) {
		// Setup expected query
		rows := sqlmock.NewRows([]string{"id", "user_id", "name", "balance", "currency", "type", "color", "is_archived", "is_excluded", "mark_for_delete", "created_at", "updated_at"}).
			AddRow(account.ID, account.UserID, account.Name, account.Balance, account.Currency, account.Type, account.Color, account.IsArchived, account.IsExcluded, account.MarkForDelete, account.CreatedAt, account.UpdatedAt)

		mock.ExpectQuery("INSERT INTO accounts").
			WithArgs(account.UserID, account.Name, account.Balance, account.Currency, account.Type, account.Color, account.IsArchived, account.IsExcluded, account.MarkForDelete, account.CreatedAt, account.UpdatedAt).
			WillReturnRows(rows)

		// Call the repository
		result, err := repo.InsertOne(ctx, account)

		// Assert
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, account.ID, result.ID)
		assert.Equal(t, account.UserID, result.UserID)

		// Ensure all expectations were met
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled exepctations: %s", err)
		}
	})

	t.Run("Database Erorr", func(t *testing.T) {
		// Setup expected query with error
		mock.ExpectQuery("INSERT INTO accounts").
			WithArgs(account.UserID, account.Name, account.Balance, account.Currency, account.Type, account.Color, account.IsArchived, account.IsExcluded, account.MarkForDelete, account.CreatedAt, account.UpdatedAt).
			WillReturnError(sqlmock.ErrCancelled)

		// Call the repository
		result, err := repo.InsertOne(ctx, account)

		// Assert
		assert.Error(t, err)
		assert.Nil(t, result)

		// Ensure all expectations were met
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled exepctations: %s", err)
		}
	})
}
