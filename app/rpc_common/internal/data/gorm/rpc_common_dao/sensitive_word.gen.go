// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package rpc_common_dao

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"fkratos/app/rpc_common/internal/data/gorm/rpc_common_model"
)

func newSensitiveWord(db *gorm.DB, opts ...gen.DOOption) sensitiveWord {
	_sensitiveWord := sensitiveWord{}

	_sensitiveWord.sensitiveWordDo.UseDB(db, opts...)
	_sensitiveWord.sensitiveWordDo.UseModel(&rpc_common_model.SensitiveWord{})

	tableName := _sensitiveWord.sensitiveWordDo.TableName()
	_sensitiveWord.ALL = field.NewAsterisk(tableName)
	_sensitiveWord.ID = field.NewString(tableName, "id")
	_sensitiveWord.CategoryID = field.NewString(tableName, "category_id")
	_sensitiveWord.Word = field.NewString(tableName, "word")
	_sensitiveWord.CreatedAt = field.NewTime(tableName, "created_at")
	_sensitiveWord.UpdatedAt = field.NewTime(tableName, "updated_at")
	_sensitiveWord.DeletedAt = field.NewField(tableName, "deleted_at")

	_sensitiveWord.fillFieldMap()

	return _sensitiveWord
}

type sensitiveWord struct {
	sensitiveWordDo sensitiveWordDo

	ALL        field.Asterisk
	ID         field.String
	CategoryID field.String // 分类ID
	Word       field.String // 敏感词
	CreatedAt  field.Time   // 创建时间
	UpdatedAt  field.Time   // 更新时间
	DeletedAt  field.Field  // 删除时间

	fieldMap map[string]field.Expr
}

func (s sensitiveWord) Table(newTableName string) *sensitiveWord {
	s.sensitiveWordDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s sensitiveWord) As(alias string) *sensitiveWord {
	s.sensitiveWordDo.DO = *(s.sensitiveWordDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *sensitiveWord) updateTableName(table string) *sensitiveWord {
	s.ALL = field.NewAsterisk(table)
	s.ID = field.NewString(table, "id")
	s.CategoryID = field.NewString(table, "category_id")
	s.Word = field.NewString(table, "word")
	s.CreatedAt = field.NewTime(table, "created_at")
	s.UpdatedAt = field.NewTime(table, "updated_at")
	s.DeletedAt = field.NewField(table, "deleted_at")

	s.fillFieldMap()

	return s
}

func (s *sensitiveWord) WithContext(ctx context.Context) *sensitiveWordDo {
	return s.sensitiveWordDo.WithContext(ctx)
}

func (s sensitiveWord) TableName() string { return s.sensitiveWordDo.TableName() }

func (s sensitiveWord) Alias() string { return s.sensitiveWordDo.Alias() }

func (s *sensitiveWord) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *sensitiveWord) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 6)
	s.fieldMap["id"] = s.ID
	s.fieldMap["category_id"] = s.CategoryID
	s.fieldMap["word"] = s.Word
	s.fieldMap["created_at"] = s.CreatedAt
	s.fieldMap["updated_at"] = s.UpdatedAt
	s.fieldMap["deleted_at"] = s.DeletedAt
}

func (s sensitiveWord) clone(db *gorm.DB) sensitiveWord {
	s.sensitiveWordDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s sensitiveWord) replaceDB(db *gorm.DB) sensitiveWord {
	s.sensitiveWordDo.ReplaceDB(db)
	return s
}

type sensitiveWordDo struct{ gen.DO }

func (s sensitiveWordDo) Debug() *sensitiveWordDo {
	return s.withDO(s.DO.Debug())
}

func (s sensitiveWordDo) WithContext(ctx context.Context) *sensitiveWordDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s sensitiveWordDo) ReadDB() *sensitiveWordDo {
	return s.Clauses(dbresolver.Read)
}

func (s sensitiveWordDo) WriteDB() *sensitiveWordDo {
	return s.Clauses(dbresolver.Write)
}

func (s sensitiveWordDo) Session(config *gorm.Session) *sensitiveWordDo {
	return s.withDO(s.DO.Session(config))
}

func (s sensitiveWordDo) Clauses(conds ...clause.Expression) *sensitiveWordDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s sensitiveWordDo) Returning(value interface{}, columns ...string) *sensitiveWordDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s sensitiveWordDo) Not(conds ...gen.Condition) *sensitiveWordDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s sensitiveWordDo) Or(conds ...gen.Condition) *sensitiveWordDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s sensitiveWordDo) Select(conds ...field.Expr) *sensitiveWordDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s sensitiveWordDo) Where(conds ...gen.Condition) *sensitiveWordDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s sensitiveWordDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *sensitiveWordDo {
	return s.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (s sensitiveWordDo) Order(conds ...field.Expr) *sensitiveWordDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s sensitiveWordDo) Distinct(cols ...field.Expr) *sensitiveWordDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s sensitiveWordDo) Omit(cols ...field.Expr) *sensitiveWordDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s sensitiveWordDo) Join(table schema.Tabler, on ...field.Expr) *sensitiveWordDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s sensitiveWordDo) LeftJoin(table schema.Tabler, on ...field.Expr) *sensitiveWordDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s sensitiveWordDo) RightJoin(table schema.Tabler, on ...field.Expr) *sensitiveWordDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s sensitiveWordDo) Group(cols ...field.Expr) *sensitiveWordDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s sensitiveWordDo) Having(conds ...gen.Condition) *sensitiveWordDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s sensitiveWordDo) Limit(limit int) *sensitiveWordDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s sensitiveWordDo) Offset(offset int) *sensitiveWordDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s sensitiveWordDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *sensitiveWordDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s sensitiveWordDo) Unscoped() *sensitiveWordDo {
	return s.withDO(s.DO.Unscoped())
}

func (s sensitiveWordDo) Create(values ...*rpc_common_model.SensitiveWord) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s sensitiveWordDo) CreateInBatches(values []*rpc_common_model.SensitiveWord, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s sensitiveWordDo) Save(values ...*rpc_common_model.SensitiveWord) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s sensitiveWordDo) First() (*rpc_common_model.SensitiveWord, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*rpc_common_model.SensitiveWord), nil
	}
}

func (s sensitiveWordDo) Take() (*rpc_common_model.SensitiveWord, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*rpc_common_model.SensitiveWord), nil
	}
}

func (s sensitiveWordDo) Last() (*rpc_common_model.SensitiveWord, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*rpc_common_model.SensitiveWord), nil
	}
}

func (s sensitiveWordDo) Find() ([]*rpc_common_model.SensitiveWord, error) {
	result, err := s.DO.Find()
	return result.([]*rpc_common_model.SensitiveWord), err
}

func (s sensitiveWordDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*rpc_common_model.SensitiveWord, err error) {
	buf := make([]*rpc_common_model.SensitiveWord, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s sensitiveWordDo) FindInBatches(result *[]*rpc_common_model.SensitiveWord, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s sensitiveWordDo) Attrs(attrs ...field.AssignExpr) *sensitiveWordDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s sensitiveWordDo) Assign(attrs ...field.AssignExpr) *sensitiveWordDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s sensitiveWordDo) Joins(fields ...field.RelationField) *sensitiveWordDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s sensitiveWordDo) Preload(fields ...field.RelationField) *sensitiveWordDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s sensitiveWordDo) FirstOrInit() (*rpc_common_model.SensitiveWord, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*rpc_common_model.SensitiveWord), nil
	}
}

func (s sensitiveWordDo) FirstOrCreate() (*rpc_common_model.SensitiveWord, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*rpc_common_model.SensitiveWord), nil
	}
}

func (s sensitiveWordDo) FindByPage(offset int, limit int) (result []*rpc_common_model.SensitiveWord, count int64, err error) {
	result, err = s.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = s.Offset(-1).Limit(-1).Count()
	return
}

func (s sensitiveWordDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s sensitiveWordDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s sensitiveWordDo) Delete(models ...*rpc_common_model.SensitiveWord) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *sensitiveWordDo) withDO(do gen.Dao) *sensitiveWordDo {
	s.DO = *do.(*gen.DO)
	return s
}
