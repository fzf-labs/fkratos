// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package rpc_common_dao

import (
	"context"
	"database/sql"

	"gorm.io/gorm"

	"gorm.io/gen"

	"gorm.io/plugin/dbresolver"
)

func Use(db *gorm.DB, opts ...gen.DOOption) *Query {
	return &Query{
		db:                db,
		SensitiveCategory: newSensitiveCategory(db, opts...),
		SensitiveWord:     newSensitiveWord(db, opts...),
	}
}

type Query struct {
	db *gorm.DB

	SensitiveCategory sensitiveCategory
	SensitiveWord     sensitiveWord
}

func (q *Query) Available() bool { return q.db != nil }

func (q *Query) clone(db *gorm.DB) *Query {
	return &Query{
		db:                db,
		SensitiveCategory: q.SensitiveCategory.clone(db),
		SensitiveWord:     q.SensitiveWord.clone(db),
	}
}

func (q *Query) ReadDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Read))
}

func (q *Query) WriteDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Write))
}

func (q *Query) ReplaceDB(db *gorm.DB) *Query {
	return &Query{
		db:                db,
		SensitiveCategory: q.SensitiveCategory.replaceDB(db),
		SensitiveWord:     q.SensitiveWord.replaceDB(db),
	}
}

type queryCtx struct {
	SensitiveCategory *sensitiveCategoryDo
	SensitiveWord     *sensitiveWordDo
}

func (q *Query) WithContext(ctx context.Context) *queryCtx {
	return &queryCtx{
		SensitiveCategory: q.SensitiveCategory.WithContext(ctx),
		SensitiveWord:     q.SensitiveWord.WithContext(ctx),
	}
}

func (q *Query) Transaction(fc func(tx *Query) error, opts ...*sql.TxOptions) error {
	return q.db.Transaction(func(tx *gorm.DB) error { return fc(q.clone(tx)) }, opts...)
}

func (q *Query) Begin(opts ...*sql.TxOptions) *QueryTx {
	return &QueryTx{q.clone(q.db.Begin(opts...))}
}

type QueryTx struct{ *Query }

func (q *QueryTx) Commit() error {
	return q.db.Commit().Error
}

func (q *QueryTx) Rollback() error {
	return q.db.Rollback().Error
}

func (q *QueryTx) SavePoint(name string) error {
	return q.db.SavePoint(name).Error
}

func (q *QueryTx) RollbackTo(name string) error {
	return q.db.RollbackTo(name).Error
}