package repository

import (
	"fmt"

	"github.com/dynastymasra/avalon/domain/repository"

	"github.com/jinzhu/gorm"
)

func translateQuery(db *gorm.DB, query *repository.Query) *gorm.DB {
	for _, filter := range query.Filters {
		switch filter.Condition {
		case repository.FilterEqual:
			q := fmt.Sprintf("%s = ?", filter.Property)
			db = db.Where(q, filter.Value)
		case repository.FilterGreaterThan:
			q := fmt.Sprintf("%s > ?", filter.Property)
			db = db.Where(q, filter.Value)
		case repository.FilterGreaterThanEqual:
			q := fmt.Sprintf("%s >= ?", filter.Property)
			db = db.Where(q, filter.Value)
		case repository.FilterLessThan:
			q := fmt.Sprintf("%s < ?", filter.Property)
			db = db.Where(q, filter.Value)
		case repository.FilterLessThanEqual:
			q := fmt.Sprintf("%s <= ?", filter.Property)
			db = db.Where(q, filter.Value)
		case repository.FilterJsonb:
			q := fmt.Sprintf("%s @> ?", filter.Property)
			db = db.Where(q, filter.Value)
		default:
			q := fmt.Sprintf("%s = ?", filter.Property)
			db = db.Where(q, filter.Value)
		}
	}

	for _, order := range query.Orderings {
		o := fmt.Sprintf("%s %s", order.Property, order.Direction)
		db = db.Order(o)
	}

	if query.Offset > 0 {
		db = db.Offset(query.Offset)
	}

	if query.Limit > 0 {
		db = db.Limit(query.Limit)
	}

	return db
}
