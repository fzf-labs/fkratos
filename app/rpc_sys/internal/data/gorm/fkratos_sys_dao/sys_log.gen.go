// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package fkratos_sys_dao

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"fkratos/app/rpc_sys/internal/data/gorm/fkratos_sys_model"
)

func newSysLog(db *gorm.DB, opts ...gen.DOOption) sysLog {
	_sysLog := sysLog{}

	_sysLog.sysLogDo.UseDB(db, opts...)
	_sysLog.sysLogDo.UseModel(&fkratos_sys_model.SysLog{})

	tableName := _sysLog.sysLogDo.TableName()
	_sysLog.ALL = field.NewAsterisk(tableName)
	_sysLog.ID = field.NewString(tableName, "id")
	_sysLog.AdminID = field.NewString(tableName, "admin_id")
	_sysLog.IP = field.NewString(tableName, "ip")
	_sysLog.URI = field.NewString(tableName, "uri")
	_sysLog.Useragent = field.NewString(tableName, "useragent")
	_sysLog.Header = field.NewField(tableName, "header")
	_sysLog.Req = field.NewField(tableName, "req")
	_sysLog.Resp = field.NewField(tableName, "resp")
	_sysLog.CreatedAt = field.NewTime(tableName, "created_at")

	_sysLog.fillFieldMap()

	return _sysLog
}

type sysLog struct {
	sysLogDo sysLogDo

	ALL       field.Asterisk
	ID        field.String // 编号
	AdminID   field.String // 管理员ID
	IP        field.String // ip
	URI       field.String // 请求路径
	Useragent field.String // 浏览器标识
	Header    field.Field  // header
	Req       field.Field  // 请求数据
	Resp      field.Field  // 响应数据
	CreatedAt field.Time   // 创建时间

	fieldMap map[string]field.Expr
}

func (s sysLog) Table(newTableName string) *sysLog {
	s.sysLogDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s sysLog) As(alias string) *sysLog {
	s.sysLogDo.DO = *(s.sysLogDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *sysLog) updateTableName(table string) *sysLog {
	s.ALL = field.NewAsterisk(table)
	s.ID = field.NewString(table, "id")
	s.AdminID = field.NewString(table, "admin_id")
	s.IP = field.NewString(table, "ip")
	s.URI = field.NewString(table, "uri")
	s.Useragent = field.NewString(table, "useragent")
	s.Header = field.NewField(table, "header")
	s.Req = field.NewField(table, "req")
	s.Resp = field.NewField(table, "resp")
	s.CreatedAt = field.NewTime(table, "created_at")

	s.fillFieldMap()

	return s
}

func (s *sysLog) WithContext(ctx context.Context) *sysLogDo { return s.sysLogDo.WithContext(ctx) }

func (s sysLog) TableName() string { return s.sysLogDo.TableName() }

func (s sysLog) Alias() string { return s.sysLogDo.Alias() }

func (s sysLog) Columns(cols ...field.Expr) gen.Columns { return s.sysLogDo.Columns(cols...) }

func (s *sysLog) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *sysLog) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 9)
	s.fieldMap["id"] = s.ID
	s.fieldMap["admin_id"] = s.AdminID
	s.fieldMap["ip"] = s.IP
	s.fieldMap["uri"] = s.URI
	s.fieldMap["useragent"] = s.Useragent
	s.fieldMap["header"] = s.Header
	s.fieldMap["req"] = s.Req
	s.fieldMap["resp"] = s.Resp
	s.fieldMap["created_at"] = s.CreatedAt
}

func (s sysLog) clone(db *gorm.DB) sysLog {
	s.sysLogDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s sysLog) replaceDB(db *gorm.DB) sysLog {
	s.sysLogDo.ReplaceDB(db)
	return s
}

type sysLogDo struct{ gen.DO }

func (s sysLogDo) Debug() *sysLogDo {
	return s.withDO(s.DO.Debug())
}

func (s sysLogDo) WithContext(ctx context.Context) *sysLogDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s sysLogDo) ReadDB() *sysLogDo {
	return s.Clauses(dbresolver.Read)
}

func (s sysLogDo) WriteDB() *sysLogDo {
	return s.Clauses(dbresolver.Write)
}

func (s sysLogDo) Session(config *gorm.Session) *sysLogDo {
	return s.withDO(s.DO.Session(config))
}

func (s sysLogDo) Clauses(conds ...clause.Expression) *sysLogDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s sysLogDo) Returning(value interface{}, columns ...string) *sysLogDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s sysLogDo) Not(conds ...gen.Condition) *sysLogDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s sysLogDo) Or(conds ...gen.Condition) *sysLogDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s sysLogDo) Select(conds ...field.Expr) *sysLogDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s sysLogDo) Where(conds ...gen.Condition) *sysLogDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s sysLogDo) Order(conds ...field.Expr) *sysLogDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s sysLogDo) Distinct(cols ...field.Expr) *sysLogDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s sysLogDo) Omit(cols ...field.Expr) *sysLogDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s sysLogDo) Join(table schema.Tabler, on ...field.Expr) *sysLogDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s sysLogDo) LeftJoin(table schema.Tabler, on ...field.Expr) *sysLogDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s sysLogDo) RightJoin(table schema.Tabler, on ...field.Expr) *sysLogDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s sysLogDo) Group(cols ...field.Expr) *sysLogDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s sysLogDo) Having(conds ...gen.Condition) *sysLogDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s sysLogDo) Limit(limit int) *sysLogDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s sysLogDo) Offset(offset int) *sysLogDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s sysLogDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *sysLogDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s sysLogDo) Unscoped() *sysLogDo {
	return s.withDO(s.DO.Unscoped())
}

func (s sysLogDo) Create(values ...*fkratos_sys_model.SysLog) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s sysLogDo) CreateInBatches(values []*fkratos_sys_model.SysLog, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s sysLogDo) Save(values ...*fkratos_sys_model.SysLog) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s sysLogDo) First() (*fkratos_sys_model.SysLog, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*fkratos_sys_model.SysLog), nil
	}
}

func (s sysLogDo) Take() (*fkratos_sys_model.SysLog, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*fkratos_sys_model.SysLog), nil
	}
}

func (s sysLogDo) Last() (*fkratos_sys_model.SysLog, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*fkratos_sys_model.SysLog), nil
	}
}

func (s sysLogDo) Find() ([]*fkratos_sys_model.SysLog, error) {
	result, err := s.DO.Find()
	return result.([]*fkratos_sys_model.SysLog), err
}

func (s sysLogDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*fkratos_sys_model.SysLog, err error) {
	buf := make([]*fkratos_sys_model.SysLog, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s sysLogDo) FindInBatches(result *[]*fkratos_sys_model.SysLog, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s sysLogDo) Attrs(attrs ...field.AssignExpr) *sysLogDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s sysLogDo) Assign(attrs ...field.AssignExpr) *sysLogDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s sysLogDo) Joins(fields ...field.RelationField) *sysLogDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s sysLogDo) Preload(fields ...field.RelationField) *sysLogDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s sysLogDo) FirstOrInit() (*fkratos_sys_model.SysLog, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*fkratos_sys_model.SysLog), nil
	}
}

func (s sysLogDo) FirstOrCreate() (*fkratos_sys_model.SysLog, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*fkratos_sys_model.SysLog), nil
	}
}

func (s sysLogDo) FindByPage(offset int, limit int) (result []*fkratos_sys_model.SysLog, count int64, err error) {
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

func (s sysLogDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s sysLogDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s sysLogDo) Delete(models ...*fkratos_sys_model.SysLog) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *sysLogDo) withDO(do gen.Dao) *sysLogDo {
	s.DO = *do.(*gen.DO)
	return s
}
