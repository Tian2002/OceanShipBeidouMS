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

func newBeiDouCard(db *gorm.DB, opts ...gen.DOOption) beiDouCard {
	_beiDouCard := beiDouCard{}

	_beiDouCard.beiDouCardDo.UseDB(db, opts...)
	_beiDouCard.beiDouCardDo.UseModel(&model.BeiDouCard{})

	tableName := _beiDouCard.beiDouCardDo.TableName()
	_beiDouCard.ALL = field.NewAsterisk(tableName)
	_beiDouCard.ID = field.NewInt32(tableName, "id")
	_beiDouCard.Number = field.NewString(tableName, "number")
	_beiDouCard.Type = field.NewBool(tableName, "type")
	_beiDouCard.Station = field.NewInt32(tableName, "station")
	_beiDouCard.ShipID = field.NewInt32(tableName, "ship_id")
	_beiDouCard.CreatedAt = field.NewTime(tableName, "created_at")
	_beiDouCard.UpdatedAt = field.NewTime(tableName, "updated_at")
	_beiDouCard.Extra = field.NewString(tableName, "extra")

	_beiDouCard.fillFieldMap()

	return _beiDouCard
}

type beiDouCard struct {
	beiDouCardDo beiDouCardDo

	ALL       field.Asterisk
	ID        field.Int32
	Number    field.String
	Type      field.Bool
	Station   field.Int32
	ShipID    field.Int32
	CreatedAt field.Time
	UpdatedAt field.Time
	Extra     field.String

	fieldMap map[string]field.Expr
}

func (b beiDouCard) Table(newTableName string) *beiDouCard {
	b.beiDouCardDo.UseTable(newTableName)
	return b.updateTableName(newTableName)
}

func (b beiDouCard) As(alias string) *beiDouCard {
	b.beiDouCardDo.DO = *(b.beiDouCardDo.As(alias).(*gen.DO))
	return b.updateTableName(alias)
}

func (b *beiDouCard) updateTableName(table string) *beiDouCard {
	b.ALL = field.NewAsterisk(table)
	b.ID = field.NewInt32(table, "id")
	b.Number = field.NewString(table, "number")
	b.Type = field.NewBool(table, "type")
	b.Station = field.NewInt32(table, "station")
	b.ShipID = field.NewInt32(table, "ship_id")
	b.CreatedAt = field.NewTime(table, "created_at")
	b.UpdatedAt = field.NewTime(table, "updated_at")
	b.Extra = field.NewString(table, "extra")

	b.fillFieldMap()

	return b
}

func (b *beiDouCard) WithContext(ctx context.Context) IBeiDouCardDo {
	return b.beiDouCardDo.WithContext(ctx)
}

func (b beiDouCard) TableName() string { return b.beiDouCardDo.TableName() }

func (b beiDouCard) Alias() string { return b.beiDouCardDo.Alias() }

func (b beiDouCard) Columns(cols ...field.Expr) gen.Columns { return b.beiDouCardDo.Columns(cols...) }

func (b *beiDouCard) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := b.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (b *beiDouCard) fillFieldMap() {
	b.fieldMap = make(map[string]field.Expr, 8)
	b.fieldMap["id"] = b.ID
	b.fieldMap["number"] = b.Number
	b.fieldMap["type"] = b.Type
	b.fieldMap["station"] = b.Station
	b.fieldMap["ship_id"] = b.ShipID
	b.fieldMap["created_at"] = b.CreatedAt
	b.fieldMap["updated_at"] = b.UpdatedAt
	b.fieldMap["extra"] = b.Extra
}

func (b beiDouCard) clone(db *gorm.DB) beiDouCard {
	b.beiDouCardDo.ReplaceConnPool(db.Statement.ConnPool)
	return b
}

func (b beiDouCard) replaceDB(db *gorm.DB) beiDouCard {
	b.beiDouCardDo.ReplaceDB(db)
	return b
}

type beiDouCardDo struct{ gen.DO }

type IBeiDouCardDo interface {
	gen.SubQuery
	Debug() IBeiDouCardDo
	WithContext(ctx context.Context) IBeiDouCardDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IBeiDouCardDo
	WriteDB() IBeiDouCardDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IBeiDouCardDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IBeiDouCardDo
	Not(conds ...gen.Condition) IBeiDouCardDo
	Or(conds ...gen.Condition) IBeiDouCardDo
	Select(conds ...field.Expr) IBeiDouCardDo
	Where(conds ...gen.Condition) IBeiDouCardDo
	Order(conds ...field.Expr) IBeiDouCardDo
	Distinct(cols ...field.Expr) IBeiDouCardDo
	Omit(cols ...field.Expr) IBeiDouCardDo
	Join(table schema.Tabler, on ...field.Expr) IBeiDouCardDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IBeiDouCardDo
	RightJoin(table schema.Tabler, on ...field.Expr) IBeiDouCardDo
	Group(cols ...field.Expr) IBeiDouCardDo
	Having(conds ...gen.Condition) IBeiDouCardDo
	Limit(limit int) IBeiDouCardDo
	Offset(offset int) IBeiDouCardDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IBeiDouCardDo
	Unscoped() IBeiDouCardDo
	Create(values ...*model.BeiDouCard) error
	CreateInBatches(values []*model.BeiDouCard, batchSize int) error
	Save(values ...*model.BeiDouCard) error
	First() (*model.BeiDouCard, error)
	Take() (*model.BeiDouCard, error)
	Last() (*model.BeiDouCard, error)
	Find() ([]*model.BeiDouCard, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.BeiDouCard, err error)
	FindInBatches(result *[]*model.BeiDouCard, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.BeiDouCard) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IBeiDouCardDo
	Assign(attrs ...field.AssignExpr) IBeiDouCardDo
	Joins(fields ...field.RelationField) IBeiDouCardDo
	Preload(fields ...field.RelationField) IBeiDouCardDo
	FirstOrInit() (*model.BeiDouCard, error)
	FirstOrCreate() (*model.BeiDouCard, error)
	FindByPage(offset int, limit int) (result []*model.BeiDouCard, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IBeiDouCardDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (b beiDouCardDo) Debug() IBeiDouCardDo {
	return b.withDO(b.DO.Debug())
}

func (b beiDouCardDo) WithContext(ctx context.Context) IBeiDouCardDo {
	return b.withDO(b.DO.WithContext(ctx))
}

func (b beiDouCardDo) ReadDB() IBeiDouCardDo {
	return b.Clauses(dbresolver.Read)
}

func (b beiDouCardDo) WriteDB() IBeiDouCardDo {
	return b.Clauses(dbresolver.Write)
}

func (b beiDouCardDo) Session(config *gorm.Session) IBeiDouCardDo {
	return b.withDO(b.DO.Session(config))
}

func (b beiDouCardDo) Clauses(conds ...clause.Expression) IBeiDouCardDo {
	return b.withDO(b.DO.Clauses(conds...))
}

func (b beiDouCardDo) Returning(value interface{}, columns ...string) IBeiDouCardDo {
	return b.withDO(b.DO.Returning(value, columns...))
}

func (b beiDouCardDo) Not(conds ...gen.Condition) IBeiDouCardDo {
	return b.withDO(b.DO.Not(conds...))
}

func (b beiDouCardDo) Or(conds ...gen.Condition) IBeiDouCardDo {
	return b.withDO(b.DO.Or(conds...))
}

func (b beiDouCardDo) Select(conds ...field.Expr) IBeiDouCardDo {
	return b.withDO(b.DO.Select(conds...))
}

func (b beiDouCardDo) Where(conds ...gen.Condition) IBeiDouCardDo {
	return b.withDO(b.DO.Where(conds...))
}

func (b beiDouCardDo) Order(conds ...field.Expr) IBeiDouCardDo {
	return b.withDO(b.DO.Order(conds...))
}

func (b beiDouCardDo) Distinct(cols ...field.Expr) IBeiDouCardDo {
	return b.withDO(b.DO.Distinct(cols...))
}

func (b beiDouCardDo) Omit(cols ...field.Expr) IBeiDouCardDo {
	return b.withDO(b.DO.Omit(cols...))
}

func (b beiDouCardDo) Join(table schema.Tabler, on ...field.Expr) IBeiDouCardDo {
	return b.withDO(b.DO.Join(table, on...))
}

func (b beiDouCardDo) LeftJoin(table schema.Tabler, on ...field.Expr) IBeiDouCardDo {
	return b.withDO(b.DO.LeftJoin(table, on...))
}

func (b beiDouCardDo) RightJoin(table schema.Tabler, on ...field.Expr) IBeiDouCardDo {
	return b.withDO(b.DO.RightJoin(table, on...))
}

func (b beiDouCardDo) Group(cols ...field.Expr) IBeiDouCardDo {
	return b.withDO(b.DO.Group(cols...))
}

func (b beiDouCardDo) Having(conds ...gen.Condition) IBeiDouCardDo {
	return b.withDO(b.DO.Having(conds...))
}

func (b beiDouCardDo) Limit(limit int) IBeiDouCardDo {
	return b.withDO(b.DO.Limit(limit))
}

func (b beiDouCardDo) Offset(offset int) IBeiDouCardDo {
	return b.withDO(b.DO.Offset(offset))
}

func (b beiDouCardDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IBeiDouCardDo {
	return b.withDO(b.DO.Scopes(funcs...))
}

func (b beiDouCardDo) Unscoped() IBeiDouCardDo {
	return b.withDO(b.DO.Unscoped())
}

func (b beiDouCardDo) Create(values ...*model.BeiDouCard) error {
	if len(values) == 0 {
		return nil
	}
	return b.DO.Create(values)
}

func (b beiDouCardDo) CreateInBatches(values []*model.BeiDouCard, batchSize int) error {
	return b.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (b beiDouCardDo) Save(values ...*model.BeiDouCard) error {
	if len(values) == 0 {
		return nil
	}
	return b.DO.Save(values)
}

func (b beiDouCardDo) First() (*model.BeiDouCard, error) {
	if result, err := b.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.BeiDouCard), nil
	}
}

func (b beiDouCardDo) Take() (*model.BeiDouCard, error) {
	if result, err := b.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.BeiDouCard), nil
	}
}

func (b beiDouCardDo) Last() (*model.BeiDouCard, error) {
	if result, err := b.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.BeiDouCard), nil
	}
}

func (b beiDouCardDo) Find() ([]*model.BeiDouCard, error) {
	result, err := b.DO.Find()
	return result.([]*model.BeiDouCard), err
}

func (b beiDouCardDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.BeiDouCard, err error) {
	buf := make([]*model.BeiDouCard, 0, batchSize)
	err = b.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (b beiDouCardDo) FindInBatches(result *[]*model.BeiDouCard, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return b.DO.FindInBatches(result, batchSize, fc)
}

func (b beiDouCardDo) Attrs(attrs ...field.AssignExpr) IBeiDouCardDo {
	return b.withDO(b.DO.Attrs(attrs...))
}

func (b beiDouCardDo) Assign(attrs ...field.AssignExpr) IBeiDouCardDo {
	return b.withDO(b.DO.Assign(attrs...))
}

func (b beiDouCardDo) Joins(fields ...field.RelationField) IBeiDouCardDo {
	for _, _f := range fields {
		b = *b.withDO(b.DO.Joins(_f))
	}
	return &b
}

func (b beiDouCardDo) Preload(fields ...field.RelationField) IBeiDouCardDo {
	for _, _f := range fields {
		b = *b.withDO(b.DO.Preload(_f))
	}
	return &b
}

func (b beiDouCardDo) FirstOrInit() (*model.BeiDouCard, error) {
	if result, err := b.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.BeiDouCard), nil
	}
}

func (b beiDouCardDo) FirstOrCreate() (*model.BeiDouCard, error) {
	if result, err := b.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.BeiDouCard), nil
	}
}

func (b beiDouCardDo) FindByPage(offset int, limit int) (result []*model.BeiDouCard, count int64, err error) {
	result, err = b.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = b.Offset(-1).Limit(-1).Count()
	return
}

func (b beiDouCardDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = b.Count()
	if err != nil {
		return
	}

	err = b.Offset(offset).Limit(limit).Scan(result)
	return
}

func (b beiDouCardDo) Scan(result interface{}) (err error) {
	return b.DO.Scan(result)
}

func (b beiDouCardDo) Delete(models ...*model.BeiDouCard) (result gen.ResultInfo, err error) {
	return b.DO.Delete(models)
}

func (b *beiDouCardDo) withDO(do gen.Dao) *beiDouCardDo {
	b.DO = *do.(*gen.DO)
	return b
}
