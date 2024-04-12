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

func newShip2mi(db *gorm.DB, opts ...gen.DOOption) ship2mi {
	_ship2mi := ship2mi{}

	_ship2mi.ship2miDo.UseDB(db, opts...)
	_ship2mi.ship2miDo.UseModel(&model.Ship2mi{})

	tableName := _ship2mi.ship2miDo.TableName()
	_ship2mi.ALL = field.NewAsterisk(tableName)
	_ship2mi.ID = field.NewInt32(tableName, "id")
	_ship2mi.OriginatorShipID = field.NewInt32(tableName, "originator_ship_id")
	_ship2mi.DestinationShipID = field.NewInt32(tableName, "destination_ship_id")
	_ship2mi.Timestamp = field.NewTime(tableName, "timestamp")
	_ship2mi.MessageContent = field.NewString(tableName, "message_content")
	_ship2mi.CreatedAt = field.NewTime(tableName, "created_at")
	_ship2mi.UpdatedAt = field.NewTime(tableName, "updated_at")
	_ship2mi.Extra = field.NewString(tableName, "extra")

	_ship2mi.fillFieldMap()

	return _ship2mi
}

type ship2mi struct {
	ship2miDo ship2miDo

	ALL               field.Asterisk
	ID                field.Int32
	OriginatorShipID  field.Int32
	DestinationShipID field.Int32
	Timestamp         field.Time
	MessageContent    field.String
	CreatedAt         field.Time
	UpdatedAt         field.Time
	Extra             field.String

	fieldMap map[string]field.Expr
}

func (s ship2mi) Table(newTableName string) *ship2mi {
	s.ship2miDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s ship2mi) As(alias string) *ship2mi {
	s.ship2miDo.DO = *(s.ship2miDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *ship2mi) updateTableName(table string) *ship2mi {
	s.ALL = field.NewAsterisk(table)
	s.ID = field.NewInt32(table, "id")
	s.OriginatorShipID = field.NewInt32(table, "originator_ship_id")
	s.DestinationShipID = field.NewInt32(table, "destination_ship_id")
	s.Timestamp = field.NewTime(table, "timestamp")
	s.MessageContent = field.NewString(table, "message_content")
	s.CreatedAt = field.NewTime(table, "created_at")
	s.UpdatedAt = field.NewTime(table, "updated_at")
	s.Extra = field.NewString(table, "extra")

	s.fillFieldMap()

	return s
}

func (s *ship2mi) WithContext(ctx context.Context) IShip2miDo { return s.ship2miDo.WithContext(ctx) }

func (s ship2mi) TableName() string { return s.ship2miDo.TableName() }

func (s ship2mi) Alias() string { return s.ship2miDo.Alias() }

func (s ship2mi) Columns(cols ...field.Expr) gen.Columns { return s.ship2miDo.Columns(cols...) }

func (s *ship2mi) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *ship2mi) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 8)
	s.fieldMap["id"] = s.ID
	s.fieldMap["originator_ship_id"] = s.OriginatorShipID
	s.fieldMap["destination_ship_id"] = s.DestinationShipID
	s.fieldMap["timestamp"] = s.Timestamp
	s.fieldMap["message_content"] = s.MessageContent
	s.fieldMap["created_at"] = s.CreatedAt
	s.fieldMap["updated_at"] = s.UpdatedAt
	s.fieldMap["extra"] = s.Extra
}

func (s ship2mi) clone(db *gorm.DB) ship2mi {
	s.ship2miDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s ship2mi) replaceDB(db *gorm.DB) ship2mi {
	s.ship2miDo.ReplaceDB(db)
	return s
}

type ship2miDo struct{ gen.DO }

type IShip2miDo interface {
	gen.SubQuery
	Debug() IShip2miDo
	WithContext(ctx context.Context) IShip2miDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IShip2miDo
	WriteDB() IShip2miDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IShip2miDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IShip2miDo
	Not(conds ...gen.Condition) IShip2miDo
	Or(conds ...gen.Condition) IShip2miDo
	Select(conds ...field.Expr) IShip2miDo
	Where(conds ...gen.Condition) IShip2miDo
	Order(conds ...field.Expr) IShip2miDo
	Distinct(cols ...field.Expr) IShip2miDo
	Omit(cols ...field.Expr) IShip2miDo
	Join(table schema.Tabler, on ...field.Expr) IShip2miDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IShip2miDo
	RightJoin(table schema.Tabler, on ...field.Expr) IShip2miDo
	Group(cols ...field.Expr) IShip2miDo
	Having(conds ...gen.Condition) IShip2miDo
	Limit(limit int) IShip2miDo
	Offset(offset int) IShip2miDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IShip2miDo
	Unscoped() IShip2miDo
	Create(values ...*model.Ship2mi) error
	CreateInBatches(values []*model.Ship2mi, batchSize int) error
	Save(values ...*model.Ship2mi) error
	First() (*model.Ship2mi, error)
	Take() (*model.Ship2mi, error)
	Last() (*model.Ship2mi, error)
	Find() ([]*model.Ship2mi, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Ship2mi, err error)
	FindInBatches(result *[]*model.Ship2mi, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Ship2mi) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IShip2miDo
	Assign(attrs ...field.AssignExpr) IShip2miDo
	Joins(fields ...field.RelationField) IShip2miDo
	Preload(fields ...field.RelationField) IShip2miDo
	FirstOrInit() (*model.Ship2mi, error)
	FirstOrCreate() (*model.Ship2mi, error)
	FindByPage(offset int, limit int) (result []*model.Ship2mi, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IShip2miDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s ship2miDo) Debug() IShip2miDo {
	return s.withDO(s.DO.Debug())
}

func (s ship2miDo) WithContext(ctx context.Context) IShip2miDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s ship2miDo) ReadDB() IShip2miDo {
	return s.Clauses(dbresolver.Read)
}

func (s ship2miDo) WriteDB() IShip2miDo {
	return s.Clauses(dbresolver.Write)
}

func (s ship2miDo) Session(config *gorm.Session) IShip2miDo {
	return s.withDO(s.DO.Session(config))
}

func (s ship2miDo) Clauses(conds ...clause.Expression) IShip2miDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s ship2miDo) Returning(value interface{}, columns ...string) IShip2miDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s ship2miDo) Not(conds ...gen.Condition) IShip2miDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s ship2miDo) Or(conds ...gen.Condition) IShip2miDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s ship2miDo) Select(conds ...field.Expr) IShip2miDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s ship2miDo) Where(conds ...gen.Condition) IShip2miDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s ship2miDo) Order(conds ...field.Expr) IShip2miDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s ship2miDo) Distinct(cols ...field.Expr) IShip2miDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s ship2miDo) Omit(cols ...field.Expr) IShip2miDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s ship2miDo) Join(table schema.Tabler, on ...field.Expr) IShip2miDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s ship2miDo) LeftJoin(table schema.Tabler, on ...field.Expr) IShip2miDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s ship2miDo) RightJoin(table schema.Tabler, on ...field.Expr) IShip2miDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s ship2miDo) Group(cols ...field.Expr) IShip2miDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s ship2miDo) Having(conds ...gen.Condition) IShip2miDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s ship2miDo) Limit(limit int) IShip2miDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s ship2miDo) Offset(offset int) IShip2miDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s ship2miDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IShip2miDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s ship2miDo) Unscoped() IShip2miDo {
	return s.withDO(s.DO.Unscoped())
}

func (s ship2miDo) Create(values ...*model.Ship2mi) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s ship2miDo) CreateInBatches(values []*model.Ship2mi, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s ship2miDo) Save(values ...*model.Ship2mi) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s ship2miDo) First() (*model.Ship2mi, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Ship2mi), nil
	}
}

func (s ship2miDo) Take() (*model.Ship2mi, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Ship2mi), nil
	}
}

func (s ship2miDo) Last() (*model.Ship2mi, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Ship2mi), nil
	}
}

func (s ship2miDo) Find() ([]*model.Ship2mi, error) {
	result, err := s.DO.Find()
	return result.([]*model.Ship2mi), err
}

func (s ship2miDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Ship2mi, err error) {
	buf := make([]*model.Ship2mi, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s ship2miDo) FindInBatches(result *[]*model.Ship2mi, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s ship2miDo) Attrs(attrs ...field.AssignExpr) IShip2miDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s ship2miDo) Assign(attrs ...field.AssignExpr) IShip2miDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s ship2miDo) Joins(fields ...field.RelationField) IShip2miDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s ship2miDo) Preload(fields ...field.RelationField) IShip2miDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s ship2miDo) FirstOrInit() (*model.Ship2mi, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Ship2mi), nil
	}
}

func (s ship2miDo) FirstOrCreate() (*model.Ship2mi, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Ship2mi), nil
	}
}

func (s ship2miDo) FindByPage(offset int, limit int) (result []*model.Ship2mi, count int64, err error) {
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

func (s ship2miDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s ship2miDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s ship2miDo) Delete(models ...*model.Ship2mi) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *ship2miDo) withDO(do gen.Dao) *ship2miDo {
	s.DO = *do.(*gen.DO)
	return s
}
