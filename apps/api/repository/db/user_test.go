package db

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/byvinesse/vinance-backend/entity"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func TestUser_InsertOne(t *testing.T) {
	// Create mock DB
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// Convert to sqlx.DB
	sqlxDB := sqlx.NewDb(db, "sqlmock")
	repo := NewUser(sqlxDB)

	// Test data
	ctx := context.Background()
	now := time.Now()

	user := &entity.User{
		Email:        "test@example.com",
		PasswordHash: "hashed_password",
		Username:     "testuser",
		PhoneNumber:  "1234567890",
		Gender:       "male",
		DateOfBirth:  time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC),
		CreatedAt:    now,
		UpdatedAt:    now,
	}

	t.Run("Success", func(t *testing.T) {
		// Setup expected query
		rows := sqlmock.NewRows([]string{"id", "email", "password_hash", "username", "phone_number", "date_of_birth", "gender", "created_at", "updated_at"}).
			AddRow("1", user.Email, user.PasswordHash, user.Username, user.PhoneNumber, user.DateOfBirth, user.Gender, user.CreatedAt, user.UpdatedAt)

		mock.ExpectQuery("INSERT INTO users").
			WithArgs(user.Email, user.PasswordHash, user.Username, user.PhoneNumber, user.DateOfBirth, user.Gender, user.CreatedAt, user.UpdatedAt).
			WillReturnRows(rows)

		// Call the repository
		result, err := repo.InsertOne(ctx, user)

		// Assert
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, "1", result.ID)
		assert.Equal(t, user.Email, result.Email)

		// Ensure all expectations were met
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})

	t.Run("Database Error", func(t *testing.T) {
		// Setup expected query with error
		mock.ExpectQuery("INSERT INTO users").
			WithArgs(user.Email, user.PasswordHash, user.Username, user.PhoneNumber, user.DateOfBirth, user.Gender, user.CreatedAt, user.UpdatedAt).
			WillReturnError(sqlmock.ErrCancelled)

		// Call the repository
		result, err := repo.InsertOne(ctx, user)

		// Assert
		assert.Error(t, err)
		assert.Nil(t, result)

		// Ensure all expectations were met
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})
}

func TestUser_FindOneByEmail(t *testing.T) {
	// Create mock DB
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// Convert to sqlx.DB
	sqlxDB := sqlx.NewDb(db, "sqlmock")
	repo := NewUser(sqlxDB)

	// Test data
	ctx := context.Background()
	email := "test@example.com"
	now := time.Now()

	t.Run("Success", func(t *testing.T) {
		// Setup expected query
		rows := sqlmock.NewRows([]string{"id", "email", "password_hash", "username", "phone_number", "date_of_birth", "gender", "created_at", "updated_at"}).
			AddRow("1", email, "hashed_password", "testuser", "1234567890", time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC), "male", now, now)

		mock.ExpectQuery("SELECT \\* FROM users").
			WithArgs(email).
			WillReturnRows(rows)

		// Call the repository
		result, err := repo.FindOneByEmail(ctx, email)

		// Assert
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, "1", result.ID)
		assert.Equal(t, email, result.Email)

		// Ensure all expectations were met
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})

	t.Run("Not Found", func(t *testing.T) {
		// Setup expected query with no rows
		mock.ExpectQuery("SELECT \\* FROM users").
			WithArgs(email).
			WillReturnError(errors.New("no rows in result set"))

		// Call the repository
		result, err := repo.FindOneByEmail(ctx, email)

		// Assert
		assert.Error(t, err)
		assert.Nil(t, result)

		// Ensure all expectations were met
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})
}
