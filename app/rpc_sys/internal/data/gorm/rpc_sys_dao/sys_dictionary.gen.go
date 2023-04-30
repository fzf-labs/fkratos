// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package rpc_sys_dao

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"fkratos/app/rpc_sys/internal/data/gorm/rpc_sys_model"
)

func newSysDictionary(db *gorm.DB, opts ...gen.DOOption) sysDictionary {
	_sysDictionary := sysDictionary{}

	_sysDictionary.sysDictionaryDo.UseDB(db, opts...)
	_sysDictionary.sysDictionaryDo.UseModel(&rpc_sys_model.SysDictionary{})

	tableName := _sysDictionary.sysDictionaryDo.TableName()
	_sysDictionary.ALL = field.NewAsterisk(tableName)
	_sysDictionary.ID = field.NewString(tableName, "id")
	_sysDictionary.Pid = field.NewString(tableName, "pid")
	_sysDictionary.Name = field.NewString(tableName, "name")
	_sysDictionary.Type = field.NewInt16(tableName, "type")
	_sysDictionary.UniqueKey = field.NewString(tableName, "unique_key")
	_sysDictionary.Value = field.NewString(tableName, "value")
	_sysDictionary.Status = field.NewInt16(tableName, "status")
	_sysDictionary.Sort = field.NewFloat64(tableName, "sort")
	_sysDictionary.Remark = field.NewString(tableName, "remark")
	_sysDictionary.CreatedAt = field.NewTime(tableName, "created_at")
	_sysDictionary.UpdatedAt = field.NewTime(tableName, "updated_at")
	_sysDictionary.DeletedAt = field.NewField(tableName, "deleted_at")

	_sysDictionary.fillFieldMap()

	return _sysDictionary
}

type sysDictionary struct {
	sysDictionaryDo sysDictionaryDo

	ALL       field.Asterisk
	ID        field.String  // 编号
	Pid       field.String  // 0=配置集 !0=父级id
	Name      field.String  // 名称
	Type      field.Int16   // 1文本 2数字 3数组 4单选 5多选 6下拉 7日期 8时间 9单图 10多图 11单文件 12多文件
	UniqueKey field.String  // 唯一值
	Value     field.String  // 配置值
	Status    field.Int16   // 0=禁用 1=开启
	Sort      field.Float64 // 排序值
	Remark    field.String  // 备注
	CreatedAt field.Time    // 创建时间
	UpdatedAt field.Time    // 更新时间
	DeletedAt field.Field   // 删除时间

	fieldMap map[string]field.Expr
}

func (s sysDictionary) Table(newTableName string) *sysDictionary {
	s.sysDictionaryDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s sysDictionary) As(alias string) *sysDictionary {
	s.sysDictionaryDo.DO = *(s.sysDictionaryDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *sysDictionary) updateTableName(table string) *sysDictionary {
	s.ALL = field.NewAsterisk(table)
	s.ID = field.NewString(table, "id")
	s.Pid = field.NewString(table, "pid")
	s.Name = field.NewString(table, "name")
	s.Type = field.NewInt16(table, "type")
	s.UniqueKey = field.NewString(table, "unique_key")
	s.Value = field.NewString(table, "value")
	s.Status = field.NewInt16(table, "status")
	s.Sort = field.NewFloat64(table, "sort")
	s.Remark = field.NewString(table, "remark")
	s.CreatedAt = field.NewTime(table, "created_at")
	s.UpdatedAt = field.NewTime(table, "updated_at")
	s.DeletedAt = field.NewField(table, "deleted_at")

	s.fillFieldMap()

	return s
}

func (s *sysDictionary) WithContext(ctx context.Context) *sysDictionaryDo {
	return s.sysDictionaryDo.WithContext(ctx)
}

func (s sysDictionary) TableName() string { return s.sysDictionaryDo.TableName() }

func (s sysDictionary) Alias() string { return s.sysDictionaryDo.Alias() }

func (s *sysDictionary) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *sysDictionary) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 12)
	s.fieldMap["id"] = s.ID
	s.fieldMap["pid"] = s.Pid
	s.fieldMap["name"] = s.Name
	s.fieldMap["type"] = s.Type
	s.fieldMap["unique_key"] = s.UniqueKey
	s.fieldMap["value"] = s.Value
	s.fieldMap["status"] = s.Status
	s.fieldMap["sort"] = s.Sort
	s.fieldMap["remark"] = s.Remark
	s.fieldMap["created_at"] = s.CreatedAt
	s.fieldMap["updated_at"] = s.UpdatedAt
	s.fieldMap["deleted_at"] = s.DeletedAt
}

func (s sysDictionary) clone(db *gorm.DB) sysDictionary {
	s.sysDictionaryDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s sysDictionary) replaceDB(db *gorm.DB) sysDictionary {
	s.sysDictionaryDo.ReplaceDB(db)
	return s
}

type sysDictionaryDo struct{ gen.DO }

func (s sysDictionaryDo) Debug() *sysDictionaryDo {
	return s.withDO(s.DO.Debug())
}

func (s sysDictionaryDo) WithContext(ctx context.Context) *sysDictionaryDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s sysDictionaryDo) ReadDB() *sysDictionaryDo {
	return s.Clauses(dbresolver.Read)
}

func (s sysDictionaryDo) WriteDB() *sysDictionaryDo {
	return s.Clauses(dbresolver.Write)
}

func (s sysDictionaryDo) Session(config *gorm.Session) *sysDictionaryDo {
	return s.withDO(s.DO.Session(config))
}

func (s sysDictionaryDo) Clauses(conds ...clause.Expression) *sysDictionaryDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s sysDictionaryDo) Returning(value interface{}, columns ...string) *sysDictionaryDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s sysDictionaryDo) Not(conds ...gen.Condition) *sysDictionaryDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s sysDictionaryDo) Or(conds ...gen.Condition) *sysDictionaryDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s sysDictionaryDo) Select(conds ...field.Expr) *sysDictionaryDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s sysDictionaryDo) Where(conds ...gen.Condition) *sysDictionaryDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s sysDictionaryDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *sysDictionaryDo {
	return s.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (s sysDictionaryDo) Order(conds ...field.Expr) *sysDictionaryDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s sysDictionaryDo) Distinct(cols ...field.Expr) *sysDictionaryDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s sysDictionaryDo) Omit(cols ...field.Expr) *sysDictionaryDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s sysDictionaryDo) Join(table schema.Tabler, on ...field.Expr) *sysDictionaryDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s sysDictionaryDo) LeftJoin(table schema.Tabler, on ...field.Expr) *sysDictionaryDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s sysDictionaryDo) RightJoin(table schema.Tabler, on ...field.Expr) *sysDictionaryDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s sysDictionaryDo) Group(cols ...field.Expr) *sysDictionaryDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s sysDictionaryDo) Having(conds ...gen.Condition) *sysDictionaryDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s sysDictionaryDo) Limit(limit int) *sysDictionaryDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s sysDictionaryDo) Offset(offset int) *sysDictionaryDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s sysDictionaryDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *sysDictionaryDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s sysDictionaryDo) Unscoped() *sysDictionaryDo {
	return s.withDO(s.DO.Unscoped())
}

func (s sysDictionaryDo) Create(values ...*rpc_sys_model.SysDictionary) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s sysDictionaryDo) CreateInBatches(values []*rpc_sys_model.SysDictionary, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s sysDictionaryDo) Save(values ...*rpc_sys_model.SysDictionary) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s sysDictionaryDo) First() (*rpc_sys_model.SysDictionary, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*rpc_sys_model.SysDictionary), nil
	}
}

func (s sysDictionaryDo) Take() (*rpc_sys_model.SysDictionary, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*rpc_sys_model.SysDictionary), nil
	}
}

func (s sysDictionaryDo) Last() (*rpc_sys_model.SysDictionary, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*rpc_sys_model.SysDictionary), nil
	}
}

func (s sysDictionaryDo) Find() ([]*rpc_sys_model.SysDictionary, error) {
	result, err := s.DO.Find()
	return result.([]*rpc_sys_model.SysDictionary), err
}

func (s sysDictionaryDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*rpc_sys_model.SysDictionary, err error) {
	buf := make([]*rpc_sys_model.SysDictionary, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s sysDictionaryDo) FindInBatches(result *[]*rpc_sys_model.SysDictionary, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s sysDictionaryDo) Attrs(attrs ...field.AssignExpr) *sysDictionaryDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s sysDictionaryDo) Assign(attrs ...field.AssignExpr) *sysDictionaryDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s sysDictionaryDo) Joins(fields ...field.RelationField) *sysDictionaryDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s sysDictionaryDo) Preload(fields ...field.RelationField) *sysDictionaryDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s sysDictionaryDo) FirstOrInit() (*rpc_sys_model.SysDictionary, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*rpc_sys_model.SysDictionary), nil
	}
}

func (s sysDictionaryDo) FirstOrCreate() (*rpc_sys_model.SysDictionary, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*rpc_sys_model.SysDictionary), nil
	}
}

func (s sysDictionaryDo) FindByPage(offset int, limit int) (result []*rpc_sys_model.SysDictionary, count int64, err error) {
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

func (s sysDictionaryDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s sysDictionaryDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s sysDictionaryDo) Delete(models ...*rpc_sys_model.SysDictionary) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *sysDictionaryDo) withDO(do gen.Dao) *sysDictionaryDo {
	s.DO = *do.(*gen.DO)
	return s
}
