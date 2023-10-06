// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package fkratos_user_dao

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"fkratos/app/rpc_user/internal/data/gorm/fkratos_user_model"
)

func newUserRule(db *gorm.DB, opts ...gen.DOOption) userRule {
	_userRule := userRule{}

	_userRule.userRuleDo.UseDB(db, opts...)
	_userRule.userRuleDo.UseModel(&fkratos_user_model.UserRule{})

	tableName := _userRule.userRuleDo.TableName()
	_userRule.ALL = field.NewAsterisk(tableName)
	_userRule.ID = field.NewString(tableName, "id")
	_userRule.Pid = field.NewString(tableName, "pid")
	_userRule.Type = field.NewString(tableName, "type")
	_userRule.Title = field.NewString(tableName, "title")
	_userRule.Name = field.NewString(tableName, "name")
	_userRule.Path = field.NewString(tableName, "path")
	_userRule.Icon = field.NewString(tableName, "icon")
	_userRule.MenuType = field.NewString(tableName, "menu_type")
	_userRule.URL = field.NewString(tableName, "url")
	_userRule.Component = field.NewString(tableName, "component")
	_userRule.Extend = field.NewString(tableName, "extend")
	_userRule.Remark = field.NewString(tableName, "remark")
	_userRule.Status = field.NewInt16(tableName, "status")
	_userRule.CreatedAt = field.NewTime(tableName, "created_at")
	_userRule.UpdatedAt = field.NewTime(tableName, "updated_at")
	_userRule.DeletedAt = field.NewField(tableName, "deleted_at")

	_userRule.fillFieldMap()

	return _userRule
}

type userRule struct {
	userRuleDo userRuleDo

	ALL       field.Asterisk
	ID        field.String // ID
	Pid       field.String // 上级菜单
	Type      field.String // 类型:route=路由,menu_dir=菜单目录,menu=菜单项,nav_user_menu=顶栏会员菜单下拉项,nav=顶栏菜单项,button=页面按钮
	Title     field.String // 标题
	Name      field.String // 规则名称
	Path      field.String // 路由路径
	Icon      field.String // 图标
	MenuType  field.String // 菜单类型:tab=选项卡,link=链接,iframe=Iframe
	URL       field.String // URL
	Component field.String // 组件路径
	Extend    field.String // 扩展属性:none=无,add_rules_only=只添加为路由,add_menu_only=只添加为菜单
	Remark    field.String // 备注
	Status    field.Int16  // 状态
	CreatedAt field.Time   // 创建时间
	UpdatedAt field.Time   // 更新时间
	DeletedAt field.Field  // 删除时间

	fieldMap map[string]field.Expr
}

func (u userRule) Table(newTableName string) *userRule {
	u.userRuleDo.UseTable(newTableName)
	return u.updateTableName(newTableName)
}

func (u userRule) As(alias string) *userRule {
	u.userRuleDo.DO = *(u.userRuleDo.As(alias).(*gen.DO))
	return u.updateTableName(alias)
}

func (u *userRule) updateTableName(table string) *userRule {
	u.ALL = field.NewAsterisk(table)
	u.ID = field.NewString(table, "id")
	u.Pid = field.NewString(table, "pid")
	u.Type = field.NewString(table, "type")
	u.Title = field.NewString(table, "title")
	u.Name = field.NewString(table, "name")
	u.Path = field.NewString(table, "path")
	u.Icon = field.NewString(table, "icon")
	u.MenuType = field.NewString(table, "menu_type")
	u.URL = field.NewString(table, "url")
	u.Component = field.NewString(table, "component")
	u.Extend = field.NewString(table, "extend")
	u.Remark = field.NewString(table, "remark")
	u.Status = field.NewInt16(table, "status")
	u.CreatedAt = field.NewTime(table, "created_at")
	u.UpdatedAt = field.NewTime(table, "updated_at")
	u.DeletedAt = field.NewField(table, "deleted_at")

	u.fillFieldMap()

	return u
}

func (u *userRule) WithContext(ctx context.Context) *userRuleDo { return u.userRuleDo.WithContext(ctx) }

func (u userRule) TableName() string { return u.userRuleDo.TableName() }

func (u userRule) Alias() string { return u.userRuleDo.Alias() }

func (u userRule) Columns(cols ...field.Expr) gen.Columns { return u.userRuleDo.Columns(cols...) }

func (u *userRule) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := u.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (u *userRule) fillFieldMap() {
	u.fieldMap = make(map[string]field.Expr, 16)
	u.fieldMap["id"] = u.ID
	u.fieldMap["pid"] = u.Pid
	u.fieldMap["type"] = u.Type
	u.fieldMap["title"] = u.Title
	u.fieldMap["name"] = u.Name
	u.fieldMap["path"] = u.Path
	u.fieldMap["icon"] = u.Icon
	u.fieldMap["menu_type"] = u.MenuType
	u.fieldMap["url"] = u.URL
	u.fieldMap["component"] = u.Component
	u.fieldMap["extend"] = u.Extend
	u.fieldMap["remark"] = u.Remark
	u.fieldMap["status"] = u.Status
	u.fieldMap["created_at"] = u.CreatedAt
	u.fieldMap["updated_at"] = u.UpdatedAt
	u.fieldMap["deleted_at"] = u.DeletedAt
}

func (u userRule) clone(db *gorm.DB) userRule {
	u.userRuleDo.ReplaceConnPool(db.Statement.ConnPool)
	return u
}

func (u userRule) replaceDB(db *gorm.DB) userRule {
	u.userRuleDo.ReplaceDB(db)
	return u
}

type userRuleDo struct{ gen.DO }

func (u userRuleDo) Debug() *userRuleDo {
	return u.withDO(u.DO.Debug())
}

func (u userRuleDo) WithContext(ctx context.Context) *userRuleDo {
	return u.withDO(u.DO.WithContext(ctx))
}

func (u userRuleDo) ReadDB() *userRuleDo {
	return u.Clauses(dbresolver.Read)
}

func (u userRuleDo) WriteDB() *userRuleDo {
	return u.Clauses(dbresolver.Write)
}

func (u userRuleDo) Session(config *gorm.Session) *userRuleDo {
	return u.withDO(u.DO.Session(config))
}

func (u userRuleDo) Clauses(conds ...clause.Expression) *userRuleDo {
	return u.withDO(u.DO.Clauses(conds...))
}

func (u userRuleDo) Returning(value interface{}, columns ...string) *userRuleDo {
	return u.withDO(u.DO.Returning(value, columns...))
}

func (u userRuleDo) Not(conds ...gen.Condition) *userRuleDo {
	return u.withDO(u.DO.Not(conds...))
}

func (u userRuleDo) Or(conds ...gen.Condition) *userRuleDo {
	return u.withDO(u.DO.Or(conds...))
}

func (u userRuleDo) Select(conds ...field.Expr) *userRuleDo {
	return u.withDO(u.DO.Select(conds...))
}

func (u userRuleDo) Where(conds ...gen.Condition) *userRuleDo {
	return u.withDO(u.DO.Where(conds...))
}

func (u userRuleDo) Order(conds ...field.Expr) *userRuleDo {
	return u.withDO(u.DO.Order(conds...))
}

func (u userRuleDo) Distinct(cols ...field.Expr) *userRuleDo {
	return u.withDO(u.DO.Distinct(cols...))
}

func (u userRuleDo) Omit(cols ...field.Expr) *userRuleDo {
	return u.withDO(u.DO.Omit(cols...))
}

func (u userRuleDo) Join(table schema.Tabler, on ...field.Expr) *userRuleDo {
	return u.withDO(u.DO.Join(table, on...))
}

func (u userRuleDo) LeftJoin(table schema.Tabler, on ...field.Expr) *userRuleDo {
	return u.withDO(u.DO.LeftJoin(table, on...))
}

func (u userRuleDo) RightJoin(table schema.Tabler, on ...field.Expr) *userRuleDo {
	return u.withDO(u.DO.RightJoin(table, on...))
}

func (u userRuleDo) Group(cols ...field.Expr) *userRuleDo {
	return u.withDO(u.DO.Group(cols...))
}

func (u userRuleDo) Having(conds ...gen.Condition) *userRuleDo {
	return u.withDO(u.DO.Having(conds...))
}

func (u userRuleDo) Limit(limit int) *userRuleDo {
	return u.withDO(u.DO.Limit(limit))
}

func (u userRuleDo) Offset(offset int) *userRuleDo {
	return u.withDO(u.DO.Offset(offset))
}

func (u userRuleDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *userRuleDo {
	return u.withDO(u.DO.Scopes(funcs...))
}

func (u userRuleDo) Unscoped() *userRuleDo {
	return u.withDO(u.DO.Unscoped())
}

func (u userRuleDo) Create(values ...*fkratos_user_model.UserRule) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Create(values)
}

func (u userRuleDo) CreateInBatches(values []*fkratos_user_model.UserRule, batchSize int) error {
	return u.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (u userRuleDo) Save(values ...*fkratos_user_model.UserRule) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Save(values)
}

func (u userRuleDo) First() (*fkratos_user_model.UserRule, error) {
	if result, err := u.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*fkratos_user_model.UserRule), nil
	}
}

func (u userRuleDo) Take() (*fkratos_user_model.UserRule, error) {
	if result, err := u.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*fkratos_user_model.UserRule), nil
	}
}

func (u userRuleDo) Last() (*fkratos_user_model.UserRule, error) {
	if result, err := u.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*fkratos_user_model.UserRule), nil
	}
}

func (u userRuleDo) Find() ([]*fkratos_user_model.UserRule, error) {
	result, err := u.DO.Find()
	return result.([]*fkratos_user_model.UserRule), err
}

func (u userRuleDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*fkratos_user_model.UserRule, err error) {
	buf := make([]*fkratos_user_model.UserRule, 0, batchSize)
	err = u.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (u userRuleDo) FindInBatches(result *[]*fkratos_user_model.UserRule, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return u.DO.FindInBatches(result, batchSize, fc)
}

func (u userRuleDo) Attrs(attrs ...field.AssignExpr) *userRuleDo {
	return u.withDO(u.DO.Attrs(attrs...))
}

func (u userRuleDo) Assign(attrs ...field.AssignExpr) *userRuleDo {
	return u.withDO(u.DO.Assign(attrs...))
}

func (u userRuleDo) Joins(fields ...field.RelationField) *userRuleDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Joins(_f))
	}
	return &u
}

func (u userRuleDo) Preload(fields ...field.RelationField) *userRuleDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Preload(_f))
	}
	return &u
}

func (u userRuleDo) FirstOrInit() (*fkratos_user_model.UserRule, error) {
	if result, err := u.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*fkratos_user_model.UserRule), nil
	}
}

func (u userRuleDo) FirstOrCreate() (*fkratos_user_model.UserRule, error) {
	if result, err := u.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*fkratos_user_model.UserRule), nil
	}
}

func (u userRuleDo) FindByPage(offset int, limit int) (result []*fkratos_user_model.UserRule, count int64, err error) {
	result, err = u.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = u.Offset(-1).Limit(-1).Count()
	return
}

func (u userRuleDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = u.Count()
	if err != nil {
		return
	}

	err = u.Offset(offset).Limit(limit).Scan(result)
	return
}

func (u userRuleDo) Scan(result interface{}) (err error) {
	return u.DO.Scan(result)
}

func (u userRuleDo) Delete(models ...*fkratos_user_model.UserRule) (result gen.ResultInfo, err error) {
	return u.DO.Delete(models)
}

func (u *userRuleDo) withDO(do gen.Dao) *userRuleDo {
	u.DO = *do.(*gen.DO)
	return u
}
