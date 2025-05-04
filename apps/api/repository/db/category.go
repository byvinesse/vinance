package db

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/byvinesse/vinance-backend/entity"
	"github.com/jmoiron/sqlx"
)

type Category struct {
	db *sqlx.DB
}

func NewCategory(db *sqlx.DB) *Category {
	return &Category{
		db: db,
	}
}

func (r *Category) FindCompleteCategory(ctx context.Context, userID string) ([]entity.CategoryWithSubCategory, error) {
	queryBuilder := sq.Select("c.id as category_id, sc.id as subcategory_id, c.name as category_name, sc.name as subcategory_name").
		From("categories c").
		Join("subcategories sc ON c.id = sc.category_id").
		Where(sq.And{
			sq.Eq{"c.mark_for_delete": false},
			sq.Eq{"c.is_archived": false},
			sq.Eq{"sc.mark_for_delete": false},
			sq.Eq{"sc.is_archived": false},
		})

	query, args, _ := queryBuilder.PlaceholderFormat(sq.Dollar).ToSql()

	var res []entity.CategoryWithSubCategory
	if err := r.db.SelectContext(ctx, &res, query, args...); err != nil {
		return nil, err
	}

	return res, nil
}
