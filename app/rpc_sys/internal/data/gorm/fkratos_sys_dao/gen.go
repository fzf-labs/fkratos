// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package fkratos_sys_dao

import (
	"context"
	"database/sql"

	"gorm.io/gorm"

	"gorm.io/gen"

	"gorm.io/plugin/dbresolver"
)

func Use(db *gorm.DB, opts ...gen.DOOption) *Query {
	return &Query{
		db:            db,
		SysAPI:        newSysAPI(db, opts...),
		SysAdmin:      newSysAdmin(db, opts...),
		SysDept:       newSysDept(db, opts...),
		SysDict:       newSysDict(db, opts...),
		SysJob:        newSysJob(db, opts...),
		SysLog:        newSysLog(db, opts...),
		SysPermission: newSysPermission(db, opts...),
		SysRole:       newSysRole(db, opts...),
		SysTenant:     newSysTenant(db, opts...),
	}
}

type Query struct {
	db *gorm.DB

	SysAPI        sysAPI
	SysAdmin      sysAdmin
	SysDept       sysDept
	SysDict       sysDict
	SysJob        sysJob
	SysLog        sysLog
	SysPermission sysPermission
	SysRole       sysRole
	SysTenant     sysTenant
}

func (q *Query) Available() bool { return q.db != nil }

func (q *Query) clone(db *gorm.DB) *Query {
	return &Query{
		db:            db,
		SysAPI:        q.SysAPI.clone(db),
		SysAdmin:      q.SysAdmin.clone(db),
		SysDept:       q.SysDept.clone(db),
		SysDict:       q.SysDict.clone(db),
		SysJob:        q.SysJob.clone(db),
		SysLog:        q.SysLog.clone(db),
		SysPermission: q.SysPermission.clone(db),
		SysRole:       q.SysRole.clone(db),
		SysTenant:     q.SysTenant.clone(db),
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
		db:            db,
		SysAPI:        q.SysAPI.replaceDB(db),
		SysAdmin:      q.SysAdmin.replaceDB(db),
		SysDept:       q.SysDept.replaceDB(db),
		SysDict:       q.SysDict.replaceDB(db),
		SysJob:        q.SysJob.replaceDB(db),
		SysLog:        q.SysLog.replaceDB(db),
		SysPermission: q.SysPermission.replaceDB(db),
		SysRole:       q.SysRole.replaceDB(db),
		SysTenant:     q.SysTenant.replaceDB(db),
	}
}

type queryCtx struct {
	SysAPI        *sysAPIDo
	SysAdmin      *sysAdminDo
	SysDept       *sysDeptDo
	SysDict       *sysDictDo
	SysJob        *sysJobDo
	SysLog        *sysLogDo
	SysPermission *sysPermissionDo
	SysRole       *sysRoleDo
	SysTenant     *sysTenantDo
}

func (q *Query) WithContext(ctx context.Context) *queryCtx {
	return &queryCtx{
		SysAPI:        q.SysAPI.WithContext(ctx),
		SysAdmin:      q.SysAdmin.WithContext(ctx),
		SysDept:       q.SysDept.WithContext(ctx),
		SysDict:       q.SysDict.WithContext(ctx),
		SysJob:        q.SysJob.WithContext(ctx),
		SysLog:        q.SysLog.WithContext(ctx),
		SysPermission: q.SysPermission.WithContext(ctx),
		SysRole:       q.SysRole.WithContext(ctx),
		SysTenant:     q.SysTenant.WithContext(ctx),
	}
}

func (q *Query) Transaction(fc func(tx *Query) error, opts ...*sql.TxOptions) error {
	return q.db.Transaction(func(tx *gorm.DB) error { return fc(q.clone(tx)) }, opts...)
}

func (q *Query) Begin(opts ...*sql.TxOptions) *QueryTx {
	tx := q.db.Begin(opts...)
	return &QueryTx{Query: q.clone(tx), Error: tx.Error}
}

type QueryTx struct {
	*Query
	Error error
}

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