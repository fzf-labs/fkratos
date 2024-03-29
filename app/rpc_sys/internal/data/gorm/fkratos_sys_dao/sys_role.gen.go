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

func newSysRole(db *gorm.DB, opts ...gen.DOOption) sysRole {
	_sysRole := sysRole{}

	_sysRole.sysRoleDo.UseDB(db, opts...)
	_sysRole.sysRoleDo.UseModel(&fkratos_sys_model.SysRole{})

	tableName := _sysRole.sysRoleDo.TableName()
	_sysRole.ALL = field.NewAsterisk(tableName)
	_sysRole.ID = field.NewString(tableName, "id")
	_sysRole.Pid = field.NewString(tableName, "pid")
	_sysRole.Name = field.NewString(tableName, "name")
	_sysRole.PermissionIds = field.NewString(tableName, "permission_ids")
	_sysRole.Remark = field.NewString(tableName, "remark")
	_sysRole.Status = field.NewInt16(tableName, "status")
	_sysRole.Sort = field.NewInt64(tableName, "sort")
	_sysRole.CreatedAt = field.NewTime(tableName, "created_at")
	_sysRole.UpdatedAt = field.NewTime(tableName, "updated_at")
	_sysRole.DeletedAt = field.NewField(tableName, "deleted_at")

	_sysRole.fillFieldMap()

	return _sysRole
}

type sysRole struct {
	sysRoleDo sysRoleDo

	ALL           field.Asterisk
	ID            field.String // 编号
	Pid           field.String // 父级id
	Name          field.String // 名称
	PermissionIds field.String // 菜单权限集合
	Remark        field.String // 备注
	Status        field.Int16  // 0=禁用 1=开启
	Sort          field.Int64  // 排序值
	CreatedAt     field.Time   // 创建时间
	UpdatedAt     field.Time   // 更新时间
	DeletedAt     field.Field  // 删除时间

	fieldMap map[string]field.Expr
}

func (s sysRole) Table(newTableName string) *sysRole {
	s.sysRoleDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s sysRole) As(alias string) *sysRole {
	s.sysRoleDo.DO = *(s.sysRoleDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *sysRole) updateTableName(table string) *sysRole {
	s.ALL = field.NewAsterisk(table)
	s.ID = field.NewString(table, "id")
	s.Pid = field.NewString(table, "pid")
	s.Name = field.NewString(table, "name")
	s.PermissionIds = field.NewString(table, "permission_ids")
	s.Remark = field.NewString(table, "remark")
	s.Status = field.NewInt16(table, "status")
	s.Sort = field.NewInt64(table, "sort")
	s.CreatedAt = field.NewTime(table, "created_at")
	s.UpdatedAt = field.NewTime(table, "updated_at")
	s.DeletedAt = field.NewField(table, "deleted_at")

	s.fillFieldMap()

	return s
}

func (s *sysRole) WithContext(ctx context.Context) *sysRoleDo { return s.sysRoleDo.WithContext(ctx) }

func (s sysRole) TableName() string { return s.sysRoleDo.TableName() }

func (s sysRole) Alias() string { return s.sysRoleDo.Alias() }

func (s sysRole) Columns(cols ...field.Expr) gen.Columns { return s.sysRoleDo.Columns(cols...) }

func (s *sysRole) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *sysRole) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 10)
	s.fieldMap["id"] = s.ID
	s.fieldMap["pid"] = s.Pid
	s.fieldMap["name"] = s.Name
	s.fieldMap["permission_ids"] = s.PermissionIds
	s.fieldMap["remark"] = s.Remark
	s.fieldMap["status"] = s.Status
	s.fieldMap["sort"] = s.Sort
	s.fieldMap["created_at"] = s.CreatedAt
	s.fieldMap["updated_at"] = s.UpdatedAt
	s.fieldMap["deleted_at"] = s.DeletedAt
}

func (s sysRole) clone(db *gorm.DB) sysRole {
	s.sysRoleDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s sysRole) replaceDB(db *gorm.DB) sysRole {
	s.sysRoleDo.ReplaceDB(db)
	return s
}

type sysRoleDo struct{ gen.DO }

func (s sysRoleDo) Debug() *sysRoleDo {
	return s.withDO(s.DO.Debug())
}

func (s sysRoleDo) WithContext(ctx context.Context) *sysRoleDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s sysRoleDo) ReadDB() *sysRoleDo {
	return s.Clauses(dbresolver.Read)
}

func (s sysRoleDo) WriteDB() *sysRoleDo {
	return s.Clauses(dbresolver.Write)
}

func (s sysRoleDo) Session(config *gorm.Session) *sysRoleDo {
	return s.withDO(s.DO.Session(config))
}

func (s sysRoleDo) Clauses(conds ...clause.Expression) *sysRoleDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s sysRoleDo) Returning(value interface{}, columns ...string) *sysRoleDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s sysRoleDo) Not(conds ...gen.Condition) *sysRoleDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s sysRoleDo) Or(conds ...gen.Condition) *sysRoleDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s sysRoleDo) Select(conds ...field.Expr) *sysRoleDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s sysRoleDo) Where(conds ...gen.Condition) *sysRoleDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s sysRoleDo) Order(conds ...field.Expr) *sysRoleDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s sysRoleDo) Distinct(cols ...field.Expr) *sysRoleDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s sysRoleDo) Omit(cols ...field.Expr) *sysRoleDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s sysRoleDo) Join(table schema.Tabler, on ...field.Expr) *sysRoleDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s sysRoleDo) LeftJoin(table schema.Tabler, on ...field.Expr) *sysRoleDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s sysRoleDo) RightJoin(table schema.Tabler, on ...field.Expr) *sysRoleDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s sysRoleDo) Group(cols ...field.Expr) *sysRoleDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s sysRoleDo) Having(conds ...gen.Condition) *sysRoleDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s sysRoleDo) Limit(limit int) *sysRoleDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s sysRoleDo) Offset(offset int) *sysRoleDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s sysRoleDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *sysRoleDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s sysRoleDo) Unscoped() *sysRoleDo {
	return s.withDO(s.DO.Unscoped())
}

func (s sysRoleDo) Create(values ...*fkratos_sys_model.SysRole) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s sysRoleDo) CreateInBatches(values []*fkratos_sys_model.SysRole, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s sysRoleDo) Save(values ...*fkratos_sys_model.SysRole) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s sysRoleDo) First() (*fkratos_sys_model.SysRole, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*fkratos_sys_model.SysRole), nil
	}
}

func (s sysRoleDo) Take() (*fkratos_sys_model.SysRole, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*fkratos_sys_model.SysRole), nil
	}
}

func (s sysRoleDo) Last() (*fkratos_sys_model.SysRole, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*fkratos_sys_model.SysRole), nil
	}
}

func (s sysRoleDo) Find() ([]*fkratos_sys_model.SysRole, error) {
	result, err := s.DO.Find()
	return result.([]*fkratos_sys_model.SysRole), err
}

func (s sysRoleDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*fkratos_sys_model.SysRole, err error) {
	buf := make([]*fkratos_sys_model.SysRole, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s sysRoleDo) FindInBatches(result *[]*fkratos_sys_model.SysRole, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s sysRoleDo) Attrs(attrs ...field.AssignExpr) *sysRoleDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s sysRoleDo) Assign(attrs ...field.AssignExpr) *sysRoleDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s sysRoleDo) Joins(fields ...field.RelationField) *sysRoleDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s sysRoleDo) Preload(fields ...field.RelationField) *sysRoleDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s sysRoleDo) FirstOrInit() (*fkratos_sys_model.SysRole, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*fkratos_sys_model.SysRole), nil
	}
}

func (s sysRoleDo) FirstOrCreate() (*fkratos_sys_model.SysRole, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*fkratos_sys_model.SysRole), nil
	}
}

func (s sysRoleDo) FindByPage(offset int, limit int) (result []*fkratos_sys_model.SysRole, count int64, err error) {
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

func (s sysRoleDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s sysRoleDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s sysRoleDo) Delete(models ...*fkratos_sys_model.SysRole) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *sysRoleDo) withDO(do gen.Dao) *sysRoleDo {
	s.DO = *do.(*gen.DO)
	return s
}
