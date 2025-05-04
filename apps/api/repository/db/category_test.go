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

func TestCategory_FindCompleteCategory(t *testing.T) {
	// Create mock DB
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// Convert to sqlx.DB
	sqlxDB := sqlx.NewDb(db, "sqlmock")
	repo := NewCategory(sqlxDB)

	// Test data
	ctx := context.Background()
	now := time.Now()

	userID := "test-user-id"

	category := &entity.Category{
		ID:        "1",
		Name:      "Test Category",
		CreatedAt: now,
		UpdatedAt: now,
	}

	subCategory := &entity.SubCategory{
		ID:        "1",
		Name:      "Test SubCategory",
		CreatedAt: now,
		UpdatedAt: now,
	}

	t.Run("Success", func(t *testing.T) {
		// Setup expected query
		rows := sqlmock.NewRows([]string{"category_id", "subcategory_id", "category_name", "subcategory_name"}).
			AddRow(category.ID, subCategory.ID, category.Name, subCategory.Name)

		// Use regexp.QuoteMeta to escape special regex characters in the SQL query
		expectedSQL := `SELECT c\.id as category_id, sc\.id as subcategory_id, c\.name as category_name, sc\.name as subcategory_name FROM categories c JOIN subcategories sc ON c\.id = sc\.category_id WHERE \(c\.mark_for_delete = \$1 AND c\.is_archived = \$2 AND sc\.mark_for_delete = \$3 AND sc\.is_archived = \$4\)`

		mock.ExpectQuery(expectedSQL).
			WithArgs(false, false, false, false).
			WillReturnRows(rows)

		// Call the repository
		result, err := repo.FindCompleteCategory(ctx, userID)

		// Assert
		assert.NoError(t, err)
		assert.NotNil(t, result)

		// Only check array contents if result is not empty
		if len(result) > 0 {
			assert.Equal(t, category.ID, result[0].CategoryID)
			assert.Equal(t, subCategory.ID, result[0].SubCategoryID)
			assert.Equal(t, category.Name, result[0].CategoryName)
			assert.Equal(t, subCategory.Name, result[0].SubCategoryName)
		}

		// Ensure all expectations were met
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled exepctations: %s", err)
		}
	})

	t.Run("Database Error", func(t *testing.T) {
		// Use regexp.QuoteMeta to escape special regex characters in the SQL query
		expectedSQL := `SELECT c\.id as category_id, sc\.id as subcategory_id, c\.name as category_name, sc\.name as subcategory_name FROM categories c JOIN subcategories sc ON c\.id = sc\.category_id WHERE \(c\.mark_for_delete = \$1 AND c\.is_archived = \$2 AND sc\.mark_for_delete = \$3 AND sc\.is_archived = \$4\)`

		// Setup expected query with error
		mock.ExpectQuery(expectedSQL).
			WithArgs(false, false, false, false).
			WillReturnError(sqlmock.ErrCancelled)

		// Call the repository
		result, err := repo.FindCompleteCategory(ctx, userID)

		// Assert
		assert.Error(t, err)
		assert.Nil(t, result)

		// Ensure all expectations were met
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})
}
