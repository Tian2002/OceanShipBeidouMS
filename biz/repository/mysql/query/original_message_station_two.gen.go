// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"OceanShipBeidouMS/biz/repository/mysql/model"
)

func newOriginalMessageStationTwo(db *gorm.DB, opts ...gen.DOOption) originalMessageStationTwo {
	_originalMessageStationTwo := originalMessageStationTwo{}

	_originalMessageStationTwo.originalMessageStationTwoDo.UseDB(db, opts...)
	_originalMessageStationTwo.originalMessageStationTwoDo.UseModel(&model.OriginalMessageStationTwo{})

	tableName := _originalMessageStationTwo.originalMessageStationTwoDo.TableName()
	_originalMessageStationTwo.ALL = field.NewAsterisk(tableName)
	_originalMessageStationTwo.ID = field.NewInt32(tableName, "id")
	_originalMessageStationTwo.OriginatorNumber = field.NewString(tableName, "originator_number")
	_originalMessageStationTwo.DestinationNumber = field.NewString(tableName, "destination_number")
	_originalMessageStationTwo.MessageType = field.NewBool(tableName, "message_type")
	_originalMessageStationTwo.Timestamp = field.NewTime(tableName, "timestamp")
	_originalMessageStationTwo.Longitude = field.NewFloat64(tableName, "longitude")
	_originalMessageStationTwo.Latitude = field.NewFloat64(tableName, "latitude")
	_originalMessageStationTwo.MessageContent = field.NewString(tableName, "message_content")
	_originalMessageStationTwo.CreatedAt = field.NewTime(tableName, "created_at")
	_originalMessageStationTwo.UpdatedAt = field.NewTime(tableName, "updated_at")
	_originalMessageStationTwo.Extra = field.NewString(tableName, "extra")

	_originalMessageStationTwo.fillFieldMap()

	return _originalMessageStationTwo
}

type originalMessageStationTwo struct {
	originalMessageStationTwoDo originalMessageStationTwoDo

	ALL               field.Asterisk
	ID                field.Int32
	OriginatorNumber  field.String
	DestinationNumber field.String
	MessageType       field.Bool
	Timestamp         field.Time
	Longitude         field.Float64
	Latitude          field.Float64
	MessageContent    field.String
	CreatedAt         field.Time
	UpdatedAt         field.Time
	Extra             field.String

	fieldMap map[string]field.Expr
}

func (o originalMessageStationTwo) Table(newTableName string) *originalMessageStationTwo {
	o.originalMessageStationTwoDo.UseTable(newTableName)
	return o.updateTableName(newTableName)
}

func (o originalMessageStationTwo) As(alias string) *originalMessageStationTwo {
	o.originalMessageStationTwoDo.DO = *(o.originalMessageStationTwoDo.As(alias).(*gen.DO))
	return o.updateTableName(alias)
}

func (o *originalMessageStationTwo) updateTableName(table string) *originalMessageStationTwo {
	o.ALL = field.NewAsterisk(table)
	o.ID = field.NewInt32(table, "id")
	o.OriginatorNumber = field.NewString(table, "originator_number")
	o.DestinationNumber = field.NewString(table, "destination_number")
	o.MessageType = field.NewBool(table, "message_type")
	o.Timestamp = field.NewTime(table, "timestamp")
	o.Longitude = field.NewFloat64(table, "longitude")
	o.Latitude = field.NewFloat64(table, "latitude")
	o.MessageContent = field.NewString(table, "message_content")
	o.CreatedAt = field.NewTime(table, "created_at")
	o.UpdatedAt = field.NewTime(table, "updated_at")
	o.Extra = field.NewString(table, "extra")

	o.fillFieldMap()

	return o
}

func (o *originalMessageStationTwo) WithContext(ctx context.Context) IOriginalMessageStationTwoDo {
	return o.originalMessageStationTwoDo.WithContext(ctx)
}

func (o originalMessageStationTwo) TableName() string {
	return o.originalMessageStationTwoDo.TableName()
}

func (o originalMessageStationTwo) Alias() string { return o.originalMessageStationTwoDo.Alias() }

func (o originalMessageStationTwo) Columns(cols ...field.Expr) gen.Columns {
	return o.originalMessageStationTwoDo.Columns(cols...)
}

func (o *originalMessageStationTwo) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := o.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (o *originalMessageStationTwo) fillFieldMap() {
	o.fieldMap = make(map[string]field.Expr, 11)
	o.fieldMap["id"] = o.ID
	o.fieldMap["originator_number"] = o.OriginatorNumber
	o.fieldMap["destination_number"] = o.DestinationNumber
	o.fieldMap["message_type"] = o.MessageType
	o.fieldMap["timestamp"] = o.Timestamp
	o.fieldMap["longitude"] = o.Longitude
	o.fieldMap["latitude"] = o.Latitude
	o.fieldMap["message_content"] = o.MessageContent
	o.fieldMap["created_at"] = o.CreatedAt
	o.fieldMap["updated_at"] = o.UpdatedAt
	o.fieldMap["extra"] = o.Extra
}

func (o originalMessageStationTwo) clone(db *gorm.DB) originalMessageStationTwo {
	o.originalMessageStationTwoDo.ReplaceConnPool(db.Statement.ConnPool)
	return o
}

func (o originalMessageStationTwo) replaceDB(db *gorm.DB) originalMessageStationTwo {
	o.originalMessageStationTwoDo.ReplaceDB(db)
	return o
}

type originalMessageStationTwoDo struct{ gen.DO }

type IOriginalMessageStationTwoDo interface {
	gen.SubQuery
	Debug() IOriginalMessageStationTwoDo
	WithContext(ctx context.Context) IOriginalMessageStationTwoDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IOriginalMessageStationTwoDo
	WriteDB() IOriginalMessageStationTwoDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IOriginalMessageStationTwoDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IOriginalMessageStationTwoDo
	Not(conds ...gen.Condition) IOriginalMessageStationTwoDo
	Or(conds ...gen.Condition) IOriginalMessageStationTwoDo
	Select(conds ...field.Expr) IOriginalMessageStationTwoDo
	Where(conds ...gen.Condition) IOriginalMessageStationTwoDo
	Order(conds ...field.Expr) IOriginalMessageStationTwoDo
	Distinct(cols ...field.Expr) IOriginalMessageStationTwoDo
	Omit(cols ...field.Expr) IOriginalMessageStationTwoDo
	Join(table schema.Tabler, on ...field.Expr) IOriginalMessageStationTwoDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IOriginalMessageStationTwoDo
	RightJoin(table schema.Tabler, on ...field.Expr) IOriginalMessageStationTwoDo
	Group(cols ...field.Expr) IOriginalMessageStationTwoDo
	Having(conds ...gen.Condition) IOriginalMessageStationTwoDo
	Limit(limit int) IOriginalMessageStationTwoDo
	Offset(offset int) IOriginalMessageStationTwoDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IOriginalMessageStationTwoDo
	Unscoped() IOriginalMessageStationTwoDo
	Create(values ...*model.OriginalMessageStationTwo) error
	CreateInBatches(values []*model.OriginalMessageStationTwo, batchSize int) error
	Save(values ...*model.OriginalMessageStationTwo) error
	First() (*model.OriginalMessageStationTwo, error)
	Take() (*model.OriginalMessageStationTwo, error)
	Last() (*model.OriginalMessageStationTwo, error)
	Find() ([]*model.OriginalMessageStationTwo, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.OriginalMessageStationTwo, err error)
	FindInBatches(result *[]*model.OriginalMessageStationTwo, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.OriginalMessageStationTwo) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IOriginalMessageStationTwoDo
	Assign(attrs ...field.AssignExpr) IOriginalMessageStationTwoDo
	Joins(fields ...field.RelationField) IOriginalMessageStationTwoDo
	Preload(fields ...field.RelationField) IOriginalMessageStationTwoDo
	FirstOrInit() (*model.OriginalMessageStationTwo, error)
	FirstOrCreate() (*model.OriginalMessageStationTwo, error)
	FindByPage(offset int, limit int) (result []*model.OriginalMessageStationTwo, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IOriginalMessageStationTwoDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (o originalMessageStationTwoDo) Debug() IOriginalMessageStationTwoDo {
	return o.withDO(o.DO.Debug())
}

func (o originalMessageStationTwoDo) WithContext(ctx context.Context) IOriginalMessageStationTwoDo {
	return o.withDO(o.DO.WithContext(ctx))
}

func (o originalMessageStationTwoDo) ReadDB() IOriginalMessageStationTwoDo {
	return o.Clauses(dbresolver.Read)
}

func (o originalMessageStationTwoDo) WriteDB() IOriginalMessageStationTwoDo {
	return o.Clauses(dbresolver.Write)
}

func (o originalMessageStationTwoDo) Session(config *gorm.Session) IOriginalMessageStationTwoDo {
	return o.withDO(o.DO.Session(config))
}

func (o originalMessageStationTwoDo) Clauses(conds ...clause.Expression) IOriginalMessageStationTwoDo {
	return o.withDO(o.DO.Clauses(conds...))
}

func (o originalMessageStationTwoDo) Returning(value interface{}, columns ...string) IOriginalMessageStationTwoDo {
	return o.withDO(o.DO.Returning(value, columns...))
}

func (o originalMessageStationTwoDo) Not(conds ...gen.Condition) IOriginalMessageStationTwoDo {
	return o.withDO(o.DO.Not(conds...))
}

func (o originalMessageStationTwoDo) Or(conds ...gen.Condition) IOriginalMessageStationTwoDo {
	return o.withDO(o.DO.Or(conds...))
}

func (o originalMessageStationTwoDo) Select(conds ...field.Expr) IOriginalMessageStationTwoDo {
	return o.withDO(o.DO.Select(conds...))
}

func (o originalMessageStationTwoDo) Where(conds ...gen.Condition) IOriginalMessageStationTwoDo {
	return o.withDO(o.DO.Where(conds...))
}

func (o originalMessageStationTwoDo) Order(conds ...field.Expr) IOriginalMessageStationTwoDo {
	return o.withDO(o.DO.Order(conds...))
}

func (o originalMessageStationTwoDo) Distinct(cols ...field.Expr) IOriginalMessageStationTwoDo {
	return o.withDO(o.DO.Distinct(cols...))
}

func (o originalMessageStationTwoDo) Omit(cols ...field.Expr) IOriginalMessageStationTwoDo {
	return o.withDO(o.DO.Omit(cols...))
}

func (o originalMessageStationTwoDo) Join(table schema.Tabler, on ...field.Expr) IOriginalMessageStationTwoDo {
	return o.withDO(o.DO.Join(table, on...))
}

func (o originalMessageStationTwoDo) LeftJoin(table schema.Tabler, on ...field.Expr) IOriginalMessageStationTwoDo {
	return o.withDO(o.DO.LeftJoin(table, on...))
}

func (o originalMessageStationTwoDo) RightJoin(table schema.Tabler, on ...field.Expr) IOriginalMessageStationTwoDo {
	return o.withDO(o.DO.RightJoin(table, on...))
}

func (o originalMessageStationTwoDo) Group(cols ...field.Expr) IOriginalMessageStationTwoDo {
	return o.withDO(o.DO.Group(cols...))
}

func (o originalMessageStationTwoDo) Having(conds ...gen.Condition) IOriginalMessageStationTwoDo {
	return o.withDO(o.DO.Having(conds...))
}

func (o originalMessageStationTwoDo) Limit(limit int) IOriginalMessageStationTwoDo {
	return o.withDO(o.DO.Limit(limit))
}

func (o originalMessageStationTwoDo) Offset(offset int) IOriginalMessageStationTwoDo {
	return o.withDO(o.DO.Offset(offset))
}

func (o originalMessageStationTwoDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IOriginalMessageStationTwoDo {
	return o.withDO(o.DO.Scopes(funcs...))
}

func (o originalMessageStationTwoDo) Unscoped() IOriginalMessageStationTwoDo {
	return o.withDO(o.DO.Unscoped())
}

func (o originalMessageStationTwoDo) Create(values ...*model.OriginalMessageStationTwo) error {
	if len(values) == 0 {
		return nil
	}
	return o.DO.Create(values)
}

func (o originalMessageStationTwoDo) CreateInBatches(values []*model.OriginalMessageStationTwo, batchSize int) error {
	return o.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (o originalMessageStationTwoDo) Save(values ...*model.OriginalMessageStationTwo) error {
	if len(values) == 0 {
		return nil
	}
	return o.DO.Save(values)
}

func (o originalMessageStationTwoDo) First() (*model.OriginalMessageStationTwo, error) {
	if result, err := o.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.OriginalMessageStationTwo), nil
	}
}

func (o originalMessageStationTwoDo) Take() (*model.OriginalMessageStationTwo, error) {
	if result, err := o.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.OriginalMessageStationTwo), nil
	}
}

func (o originalMessageStationTwoDo) Last() (*model.OriginalMessageStationTwo, error) {
	if result, err := o.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.OriginalMessageStationTwo), nil
	}
}

func (o originalMessageStationTwoDo) Find() ([]*model.OriginalMessageStationTwo, error) {
	result, err := o.DO.Find()
	return result.([]*model.OriginalMessageStationTwo), err
}

func (o originalMessageStationTwoDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.OriginalMessageStationTwo, err error) {
	buf := make([]*model.OriginalMessageStationTwo, 0, batchSize)
	err = o.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (o originalMessageStationTwoDo) FindInBatches(result *[]*model.OriginalMessageStationTwo, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return o.DO.FindInBatches(result, batchSize, fc)
}

func (o originalMessageStationTwoDo) Attrs(attrs ...field.AssignExpr) IOriginalMessageStationTwoDo {
	return o.withDO(o.DO.Attrs(attrs...))
}

func (o originalMessageStationTwoDo) Assign(attrs ...field.AssignExpr) IOriginalMessageStationTwoDo {
	return o.withDO(o.DO.Assign(attrs...))
}

func (o originalMessageStationTwoDo) Joins(fields ...field.RelationField) IOriginalMessageStationTwoDo {
	for _, _f := range fields {
		o = *o.withDO(o.DO.Joins(_f))
	}
	return &o
}

func (o originalMessageStationTwoDo) Preload(fields ...field.RelationField) IOriginalMessageStationTwoDo {
	for _, _f := range fields {
		o = *o.withDO(o.DO.Preload(_f))
	}
	return &o
}

func (o originalMessageStationTwoDo) FirstOrInit() (*model.OriginalMessageStationTwo, error) {
	if result, err := o.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.OriginalMessageStationTwo), nil
	}
}

func (o originalMessageStationTwoDo) FirstOrCreate() (*model.OriginalMessageStationTwo, error) {
	if result, err := o.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.OriginalMessageStationTwo), nil
	}
}

func (o originalMessageStationTwoDo) FindByPage(offset int, limit int) (result []*model.OriginalMessageStationTwo, count int64, err error) {
	result, err = o.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = o.Offset(-1).Limit(-1).Count()
	return
}

func (o originalMessageStationTwoDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = o.Count()
	if err != nil {
		return
	}

	err = o.Offset(offset).Limit(limit).Scan(result)
	return
}

func (o originalMessageStationTwoDo) Scan(result interface{}) (err error) {
	return o.DO.Scan(result)
}

func (o originalMessageStationTwoDo) Delete(models ...*model.OriginalMessageStationTwo) (result gen.ResultInfo, err error) {
	return o.DO.Delete(models)
}

func (o *originalMessageStationTwoDo) withDO(do gen.Dao) *originalMessageStationTwoDo {
	o.DO = *do.(*gen.DO)
	return o
}
