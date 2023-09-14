// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package fkratos_device_dao

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"fkratos/app/rpc_device/internal/data/gorm/fkratos_device_model"
)

func newDevice(db *gorm.DB, opts ...gen.DOOption) device {
	_device := device{}

	_device.deviceDo.UseDB(db, opts...)
	_device.deviceDo.UseModel(&fkratos_device_model.Device{})

	tableName := _device.deviceDo.TableName()
	_device.ALL = field.NewAsterisk(tableName)
	_device.ID = field.NewString(tableName, "id")
	_device.SNID = field.NewString(tableName, "SNId")
	_device.DeviceBrand = field.NewString(tableName, "deviceBrand")
	_device.ModelDevice = field.NewString(tableName, "modelDevice")
	_device.Status = field.NewInt16(tableName, "status")
	_device.CreatedAt = field.NewTime(tableName, "created_at")
	_device.UpdatedAt = field.NewTime(tableName, "updated_at")
	_device.DeletedAt = field.NewField(tableName, "deleted_at")

	_device.fillFieldMap()

	return _device
}

type device struct {
	deviceDo deviceDo

	ALL         field.Asterisk
	ID          field.String // 记录ID
	SNID        field.String // 设备SN唯一标识码
	DeviceBrand field.String // 设备品牌
	ModelDevice field.String // 设备型号
	Status      field.Int16  // 状态
	CreatedAt   field.Time   // 创建时间
	UpdatedAt   field.Time   // 更新时间
	DeletedAt   field.Field  // 删除时间

	fieldMap map[string]field.Expr
}

func (d device) Table(newTableName string) *device {
	d.deviceDo.UseTable(newTableName)
	return d.updateTableName(newTableName)
}

func (d device) As(alias string) *device {
	d.deviceDo.DO = *(d.deviceDo.As(alias).(*gen.DO))
	return d.updateTableName(alias)
}

func (d *device) updateTableName(table string) *device {
	d.ALL = field.NewAsterisk(table)
	d.ID = field.NewString(table, "id")
	d.SNID = field.NewString(table, "SNId")
	d.DeviceBrand = field.NewString(table, "deviceBrand")
	d.ModelDevice = field.NewString(table, "modelDevice")
	d.Status = field.NewInt16(table, "status")
	d.CreatedAt = field.NewTime(table, "created_at")
	d.UpdatedAt = field.NewTime(table, "updated_at")
	d.DeletedAt = field.NewField(table, "deleted_at")

	d.fillFieldMap()

	return d
}

func (d *device) WithContext(ctx context.Context) *deviceDo { return d.deviceDo.WithContext(ctx) }

func (d device) TableName() string { return d.deviceDo.TableName() }

func (d device) Alias() string { return d.deviceDo.Alias() }

func (d device) Columns(cols ...field.Expr) gen.Columns { return d.deviceDo.Columns(cols...) }

func (d *device) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := d.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (d *device) fillFieldMap() {
	d.fieldMap = make(map[string]field.Expr, 8)
	d.fieldMap["id"] = d.ID
	d.fieldMap["SNId"] = d.SNID
	d.fieldMap["deviceBrand"] = d.DeviceBrand
	d.fieldMap["modelDevice"] = d.ModelDevice
	d.fieldMap["status"] = d.Status
	d.fieldMap["created_at"] = d.CreatedAt
	d.fieldMap["updated_at"] = d.UpdatedAt
	d.fieldMap["deleted_at"] = d.DeletedAt
}

func (d device) clone(db *gorm.DB) device {
	d.deviceDo.ReplaceConnPool(db.Statement.ConnPool)
	return d
}

func (d device) replaceDB(db *gorm.DB) device {
	d.deviceDo.ReplaceDB(db)
	return d
}

type deviceDo struct{ gen.DO }

func (d deviceDo) Debug() *deviceDo {
	return d.withDO(d.DO.Debug())
}

func (d deviceDo) WithContext(ctx context.Context) *deviceDo {
	return d.withDO(d.DO.WithContext(ctx))
}

func (d deviceDo) ReadDB() *deviceDo {
	return d.Clauses(dbresolver.Read)
}

func (d deviceDo) WriteDB() *deviceDo {
	return d.Clauses(dbresolver.Write)
}

func (d deviceDo) Session(config *gorm.Session) *deviceDo {
	return d.withDO(d.DO.Session(config))
}

func (d deviceDo) Clauses(conds ...clause.Expression) *deviceDo {
	return d.withDO(d.DO.Clauses(conds...))
}

func (d deviceDo) Returning(value interface{}, columns ...string) *deviceDo {
	return d.withDO(d.DO.Returning(value, columns...))
}

func (d deviceDo) Not(conds ...gen.Condition) *deviceDo {
	return d.withDO(d.DO.Not(conds...))
}

func (d deviceDo) Or(conds ...gen.Condition) *deviceDo {
	return d.withDO(d.DO.Or(conds...))
}

func (d deviceDo) Select(conds ...field.Expr) *deviceDo {
	return d.withDO(d.DO.Select(conds...))
}

func (d deviceDo) Where(conds ...gen.Condition) *deviceDo {
	return d.withDO(d.DO.Where(conds...))
}

func (d deviceDo) Order(conds ...field.Expr) *deviceDo {
	return d.withDO(d.DO.Order(conds...))
}

func (d deviceDo) Distinct(cols ...field.Expr) *deviceDo {
	return d.withDO(d.DO.Distinct(cols...))
}

func (d deviceDo) Omit(cols ...field.Expr) *deviceDo {
	return d.withDO(d.DO.Omit(cols...))
}

func (d deviceDo) Join(table schema.Tabler, on ...field.Expr) *deviceDo {
	return d.withDO(d.DO.Join(table, on...))
}

func (d deviceDo) LeftJoin(table schema.Tabler, on ...field.Expr) *deviceDo {
	return d.withDO(d.DO.LeftJoin(table, on...))
}

func (d deviceDo) RightJoin(table schema.Tabler, on ...field.Expr) *deviceDo {
	return d.withDO(d.DO.RightJoin(table, on...))
}

func (d deviceDo) Group(cols ...field.Expr) *deviceDo {
	return d.withDO(d.DO.Group(cols...))
}

func (d deviceDo) Having(conds ...gen.Condition) *deviceDo {
	return d.withDO(d.DO.Having(conds...))
}

func (d deviceDo) Limit(limit int) *deviceDo {
	return d.withDO(d.DO.Limit(limit))
}

func (d deviceDo) Offset(offset int) *deviceDo {
	return d.withDO(d.DO.Offset(offset))
}

func (d deviceDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *deviceDo {
	return d.withDO(d.DO.Scopes(funcs...))
}

func (d deviceDo) Unscoped() *deviceDo {
	return d.withDO(d.DO.Unscoped())
}

func (d deviceDo) Create(values ...*fkratos_device_model.Device) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Create(values)
}

func (d deviceDo) CreateInBatches(values []*fkratos_device_model.Device, batchSize int) error {
	return d.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (d deviceDo) Save(values ...*fkratos_device_model.Device) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Save(values)
}

func (d deviceDo) First() (*fkratos_device_model.Device, error) {
	if result, err := d.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*fkratos_device_model.Device), nil
	}
}

func (d deviceDo) Take() (*fkratos_device_model.Device, error) {
	if result, err := d.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*fkratos_device_model.Device), nil
	}
}

func (d deviceDo) Last() (*fkratos_device_model.Device, error) {
	if result, err := d.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*fkratos_device_model.Device), nil
	}
}

func (d deviceDo) Find() ([]*fkratos_device_model.Device, error) {
	result, err := d.DO.Find()
	return result.([]*fkratos_device_model.Device), err
}

func (d deviceDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*fkratos_device_model.Device, err error) {
	buf := make([]*fkratos_device_model.Device, 0, batchSize)
	err = d.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (d deviceDo) FindInBatches(result *[]*fkratos_device_model.Device, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return d.DO.FindInBatches(result, batchSize, fc)
}

func (d deviceDo) Attrs(attrs ...field.AssignExpr) *deviceDo {
	return d.withDO(d.DO.Attrs(attrs...))
}

func (d deviceDo) Assign(attrs ...field.AssignExpr) *deviceDo {
	return d.withDO(d.DO.Assign(attrs...))
}

func (d deviceDo) Joins(fields ...field.RelationField) *deviceDo {
	for _, _f := range fields {
		d = *d.withDO(d.DO.Joins(_f))
	}
	return &d
}

func (d deviceDo) Preload(fields ...field.RelationField) *deviceDo {
	for _, _f := range fields {
		d = *d.withDO(d.DO.Preload(_f))
	}
	return &d
}

func (d deviceDo) FirstOrInit() (*fkratos_device_model.Device, error) {
	if result, err := d.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*fkratos_device_model.Device), nil
	}
}

func (d deviceDo) FirstOrCreate() (*fkratos_device_model.Device, error) {
	if result, err := d.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*fkratos_device_model.Device), nil
	}
}

func (d deviceDo) FindByPage(offset int, limit int) (result []*fkratos_device_model.Device, count int64, err error) {
	result, err = d.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = d.Offset(-1).Limit(-1).Count()
	return
}

func (d deviceDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = d.Count()
	if err != nil {
		return
	}

	err = d.Offset(offset).Limit(limit).Scan(result)
	return
}

func (d deviceDo) Scan(result interface{}) (err error) {
	return d.DO.Scan(result)
}

func (d deviceDo) Delete(models ...*fkratos_device_model.Device) (result gen.ResultInfo, err error) {
	return d.DO.Delete(models)
}

func (d *deviceDo) withDO(do gen.Dao) *deviceDo {
	d.DO = *do.(*gen.DO)
	return d
}
