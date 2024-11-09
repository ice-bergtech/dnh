// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package db

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	dnh "github.com/ice-bergtech/dnh/src/internal/lib"
)

func newIPAddress(db *gorm.DB, opts ...gen.DOOption) iPAddress {
	_iPAddress := iPAddress{}

	_iPAddress.iPAddressDo.UseDB(db, opts...)
	_iPAddress.iPAddressDo.UseModel(&dnh.IPAddress{})

	tableName := _iPAddress.iPAddressDo.TableName()
	_iPAddress.ALL = field.NewAsterisk(tableName)
	_iPAddress.ID = field.NewUint(tableName, "id")
	_iPAddress.CreatedAt = field.NewTime(tableName, "created_at")
	_iPAddress.UpdatedAt = field.NewTime(tableName, "updated_at")
	_iPAddress.DeletedAt = field.NewField(tableName, "deleted_at")
	_iPAddress.IP = field.NewField(tableName, "ip")
	_iPAddress.Mask = field.NewField(tableName, "mask")
	_iPAddress.Tags = field.NewField(tableName, "tags")
	_iPAddress.Advertisers = iPAddressManyToManyAdvertisers{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Advertisers", "dnh.ASNInfo"),
	}

	_iPAddress.fillFieldMap()

	return _iPAddress
}

type iPAddress struct {
	iPAddressDo iPAddressDo

	ALL         field.Asterisk
	ID          field.Uint
	CreatedAt   field.Time
	UpdatedAt   field.Time
	DeletedAt   field.Field
	IP          field.Field
	Mask        field.Field
	Tags        field.Field
	Advertisers iPAddressManyToManyAdvertisers

	fieldMap map[string]field.Expr
}

func (i iPAddress) Table(newTableName string) *iPAddress {
	i.iPAddressDo.UseTable(newTableName)
	return i.updateTableName(newTableName)
}

func (i iPAddress) As(alias string) *iPAddress {
	i.iPAddressDo.DO = *(i.iPAddressDo.As(alias).(*gen.DO))
	return i.updateTableName(alias)
}

func (i *iPAddress) updateTableName(table string) *iPAddress {
	i.ALL = field.NewAsterisk(table)
	i.ID = field.NewUint(table, "id")
	i.CreatedAt = field.NewTime(table, "created_at")
	i.UpdatedAt = field.NewTime(table, "updated_at")
	i.DeletedAt = field.NewField(table, "deleted_at")
	i.IP = field.NewField(table, "ip")
	i.Mask = field.NewField(table, "mask")
	i.Tags = field.NewField(table, "tags")

	i.fillFieldMap()

	return i
}

func (i *iPAddress) WithContext(ctx context.Context) IIPAddressDo {
	return i.iPAddressDo.WithContext(ctx)
}

func (i iPAddress) TableName() string { return i.iPAddressDo.TableName() }

func (i iPAddress) Alias() string { return i.iPAddressDo.Alias() }

func (i iPAddress) Columns(cols ...field.Expr) gen.Columns { return i.iPAddressDo.Columns(cols...) }

func (i *iPAddress) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := i.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (i *iPAddress) fillFieldMap() {
	i.fieldMap = make(map[string]field.Expr, 8)
	i.fieldMap["id"] = i.ID
	i.fieldMap["created_at"] = i.CreatedAt
	i.fieldMap["updated_at"] = i.UpdatedAt
	i.fieldMap["deleted_at"] = i.DeletedAt
	i.fieldMap["ip"] = i.IP
	i.fieldMap["mask"] = i.Mask
	i.fieldMap["tags"] = i.Tags

}

func (i iPAddress) clone(db *gorm.DB) iPAddress {
	i.iPAddressDo.ReplaceConnPool(db.Statement.ConnPool)
	return i
}

func (i iPAddress) replaceDB(db *gorm.DB) iPAddress {
	i.iPAddressDo.ReplaceDB(db)
	return i
}

type iPAddressManyToManyAdvertisers struct {
	db *gorm.DB

	field.RelationField
}

func (a iPAddressManyToManyAdvertisers) Where(conds ...field.Expr) *iPAddressManyToManyAdvertisers {
	if len(conds) == 0 {
		return &a
	}

	exprs := make([]clause.Expression, 0, len(conds))
	for _, cond := range conds {
		exprs = append(exprs, cond.BeCond().(clause.Expression))
	}
	a.db = a.db.Clauses(clause.Where{Exprs: exprs})
	return &a
}

func (a iPAddressManyToManyAdvertisers) WithContext(ctx context.Context) *iPAddressManyToManyAdvertisers {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a iPAddressManyToManyAdvertisers) Session(session *gorm.Session) *iPAddressManyToManyAdvertisers {
	a.db = a.db.Session(session)
	return &a
}

func (a iPAddressManyToManyAdvertisers) Model(m *dnh.IPAddress) *iPAddressManyToManyAdvertisersTx {
	return &iPAddressManyToManyAdvertisersTx{a.db.Model(m).Association(a.Name())}
}

type iPAddressManyToManyAdvertisersTx struct{ tx *gorm.Association }

func (a iPAddressManyToManyAdvertisersTx) Find() (result []*dnh.ASNInfo, err error) {
	return result, a.tx.Find(&result)
}

func (a iPAddressManyToManyAdvertisersTx) Append(values ...*dnh.ASNInfo) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a iPAddressManyToManyAdvertisersTx) Replace(values ...*dnh.ASNInfo) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a iPAddressManyToManyAdvertisersTx) Delete(values ...*dnh.ASNInfo) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a iPAddressManyToManyAdvertisersTx) Clear() error {
	return a.tx.Clear()
}

func (a iPAddressManyToManyAdvertisersTx) Count() int64 {
	return a.tx.Count()
}

type iPAddressDo struct{ gen.DO }

type IIPAddressDo interface {
	gen.SubQuery
	Debug() IIPAddressDo
	WithContext(ctx context.Context) IIPAddressDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IIPAddressDo
	WriteDB() IIPAddressDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IIPAddressDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IIPAddressDo
	Not(conds ...gen.Condition) IIPAddressDo
	Or(conds ...gen.Condition) IIPAddressDo
	Select(conds ...field.Expr) IIPAddressDo
	Where(conds ...gen.Condition) IIPAddressDo
	Order(conds ...field.Expr) IIPAddressDo
	Distinct(cols ...field.Expr) IIPAddressDo
	Omit(cols ...field.Expr) IIPAddressDo
	Join(table schema.Tabler, on ...field.Expr) IIPAddressDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IIPAddressDo
	RightJoin(table schema.Tabler, on ...field.Expr) IIPAddressDo
	Group(cols ...field.Expr) IIPAddressDo
	Having(conds ...gen.Condition) IIPAddressDo
	Limit(limit int) IIPAddressDo
	Offset(offset int) IIPAddressDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IIPAddressDo
	Unscoped() IIPAddressDo
	Create(values ...*dnh.IPAddress) error
	CreateInBatches(values []*dnh.IPAddress, batchSize int) error
	Save(values ...*dnh.IPAddress) error
	First() (*dnh.IPAddress, error)
	Take() (*dnh.IPAddress, error)
	Last() (*dnh.IPAddress, error)
	Find() ([]*dnh.IPAddress, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*dnh.IPAddress, err error)
	FindInBatches(result *[]*dnh.IPAddress, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*dnh.IPAddress) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IIPAddressDo
	Assign(attrs ...field.AssignExpr) IIPAddressDo
	Joins(fields ...field.RelationField) IIPAddressDo
	Preload(fields ...field.RelationField) IIPAddressDo
	FirstOrInit() (*dnh.IPAddress, error)
	FirstOrCreate() (*dnh.IPAddress, error)
	FindByPage(offset int, limit int) (result []*dnh.IPAddress, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IIPAddressDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (i iPAddressDo) Debug() IIPAddressDo {
	return i.withDO(i.DO.Debug())
}

func (i iPAddressDo) WithContext(ctx context.Context) IIPAddressDo {
	return i.withDO(i.DO.WithContext(ctx))
}

func (i iPAddressDo) ReadDB() IIPAddressDo {
	return i.Clauses(dbresolver.Read)
}

func (i iPAddressDo) WriteDB() IIPAddressDo {
	return i.Clauses(dbresolver.Write)
}

func (i iPAddressDo) Session(config *gorm.Session) IIPAddressDo {
	return i.withDO(i.DO.Session(config))
}

func (i iPAddressDo) Clauses(conds ...clause.Expression) IIPAddressDo {
	return i.withDO(i.DO.Clauses(conds...))
}

func (i iPAddressDo) Returning(value interface{}, columns ...string) IIPAddressDo {
	return i.withDO(i.DO.Returning(value, columns...))
}

func (i iPAddressDo) Not(conds ...gen.Condition) IIPAddressDo {
	return i.withDO(i.DO.Not(conds...))
}

func (i iPAddressDo) Or(conds ...gen.Condition) IIPAddressDo {
	return i.withDO(i.DO.Or(conds...))
}

func (i iPAddressDo) Select(conds ...field.Expr) IIPAddressDo {
	return i.withDO(i.DO.Select(conds...))
}

func (i iPAddressDo) Where(conds ...gen.Condition) IIPAddressDo {
	return i.withDO(i.DO.Where(conds...))
}

func (i iPAddressDo) Order(conds ...field.Expr) IIPAddressDo {
	return i.withDO(i.DO.Order(conds...))
}

func (i iPAddressDo) Distinct(cols ...field.Expr) IIPAddressDo {
	return i.withDO(i.DO.Distinct(cols...))
}

func (i iPAddressDo) Omit(cols ...field.Expr) IIPAddressDo {
	return i.withDO(i.DO.Omit(cols...))
}

func (i iPAddressDo) Join(table schema.Tabler, on ...field.Expr) IIPAddressDo {
	return i.withDO(i.DO.Join(table, on...))
}

func (i iPAddressDo) LeftJoin(table schema.Tabler, on ...field.Expr) IIPAddressDo {
	return i.withDO(i.DO.LeftJoin(table, on...))
}

func (i iPAddressDo) RightJoin(table schema.Tabler, on ...field.Expr) IIPAddressDo {
	return i.withDO(i.DO.RightJoin(table, on...))
}

func (i iPAddressDo) Group(cols ...field.Expr) IIPAddressDo {
	return i.withDO(i.DO.Group(cols...))
}

func (i iPAddressDo) Having(conds ...gen.Condition) IIPAddressDo {
	return i.withDO(i.DO.Having(conds...))
}

func (i iPAddressDo) Limit(limit int) IIPAddressDo {
	return i.withDO(i.DO.Limit(limit))
}

func (i iPAddressDo) Offset(offset int) IIPAddressDo {
	return i.withDO(i.DO.Offset(offset))
}

func (i iPAddressDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IIPAddressDo {
	return i.withDO(i.DO.Scopes(funcs...))
}

func (i iPAddressDo) Unscoped() IIPAddressDo {
	return i.withDO(i.DO.Unscoped())
}

func (i iPAddressDo) Create(values ...*dnh.IPAddress) error {
	if len(values) == 0 {
		return nil
	}
	return i.DO.Create(values)
}

func (i iPAddressDo) CreateInBatches(values []*dnh.IPAddress, batchSize int) error {
	return i.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (i iPAddressDo) Save(values ...*dnh.IPAddress) error {
	if len(values) == 0 {
		return nil
	}
	return i.DO.Save(values)
}

func (i iPAddressDo) First() (*dnh.IPAddress, error) {
	if result, err := i.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*dnh.IPAddress), nil
	}
}

func (i iPAddressDo) Take() (*dnh.IPAddress, error) {
	if result, err := i.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*dnh.IPAddress), nil
	}
}

func (i iPAddressDo) Last() (*dnh.IPAddress, error) {
	if result, err := i.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*dnh.IPAddress), nil
	}
}

func (i iPAddressDo) Find() ([]*dnh.IPAddress, error) {
	result, err := i.DO.Find()
	return result.([]*dnh.IPAddress), err
}

func (i iPAddressDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*dnh.IPAddress, err error) {
	buf := make([]*dnh.IPAddress, 0, batchSize)
	err = i.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (i iPAddressDo) FindInBatches(result *[]*dnh.IPAddress, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return i.DO.FindInBatches(result, batchSize, fc)
}

func (i iPAddressDo) Attrs(attrs ...field.AssignExpr) IIPAddressDo {
	return i.withDO(i.DO.Attrs(attrs...))
}

func (i iPAddressDo) Assign(attrs ...field.AssignExpr) IIPAddressDo {
	return i.withDO(i.DO.Assign(attrs...))
}

func (i iPAddressDo) Joins(fields ...field.RelationField) IIPAddressDo {
	for _, _f := range fields {
		i = *i.withDO(i.DO.Joins(_f))
	}
	return &i
}

func (i iPAddressDo) Preload(fields ...field.RelationField) IIPAddressDo {
	for _, _f := range fields {
		i = *i.withDO(i.DO.Preload(_f))
	}
	return &i
}

func (i iPAddressDo) FirstOrInit() (*dnh.IPAddress, error) {
	if result, err := i.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*dnh.IPAddress), nil
	}
}

func (i iPAddressDo) FirstOrCreate() (*dnh.IPAddress, error) {
	if result, err := i.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*dnh.IPAddress), nil
	}
}

func (i iPAddressDo) FindByPage(offset int, limit int) (result []*dnh.IPAddress, count int64, err error) {
	result, err = i.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = i.Offset(-1).Limit(-1).Count()
	return
}

func (i iPAddressDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = i.Count()
	if err != nil {
		return
	}

	err = i.Offset(offset).Limit(limit).Scan(result)
	return
}

func (i iPAddressDo) Scan(result interface{}) (err error) {
	return i.DO.Scan(result)
}

func (i iPAddressDo) Delete(models ...*dnh.IPAddress) (result gen.ResultInfo, err error) {
	return i.DO.Delete(models)
}

func (i *iPAddressDo) withDO(do gen.Dao) *iPAddressDo {
	i.DO = *do.(*gen.DO)
	return i
}
