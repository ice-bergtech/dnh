// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package db

import (
	"context"
	"database/sql"

	"gorm.io/gorm"

	"gorm.io/gen"

	"gorm.io/plugin/dbresolver"
)

var (
	Q          = new(Query)
	ASNInfo    *aSNInfo
	DNSEntry   *dNSEntry
	Domain     *domain
	IPAddress  *iPAddress
	Nameserver *nameserver
	Path       *path
	Registrar  *registrar
	Whois      *whois
)

func SetDefault(db *gorm.DB, opts ...gen.DOOption) {
	*Q = *Use(db, opts...)
	ASNInfo = &Q.ASNInfo
	DNSEntry = &Q.DNSEntry
	Domain = &Q.Domain
	IPAddress = &Q.IPAddress
	Nameserver = &Q.Nameserver
	Path = &Q.Path
	Registrar = &Q.Registrar
	Whois = &Q.Whois
}

func Use(db *gorm.DB, opts ...gen.DOOption) *Query {
	return &Query{
		db:         db,
		ASNInfo:    newASNInfo(db, opts...),
		DNSEntry:   newDNSEntry(db, opts...),
		Domain:     newDomain(db, opts...),
		IPAddress:  newIPAddress(db, opts...),
		Nameserver: newNameserver(db, opts...),
		Path:       newPath(db, opts...),
		Registrar:  newRegistrar(db, opts...),
		Whois:      newWhois(db, opts...),
	}
}

type Query struct {
	db *gorm.DB

	ASNInfo    aSNInfo
	DNSEntry   dNSEntry
	Domain     domain
	IPAddress  iPAddress
	Nameserver nameserver
	Path       path
	Registrar  registrar
	Whois      whois
}

func (q *Query) Available() bool { return q.db != nil }

func (q *Query) clone(db *gorm.DB) *Query {
	return &Query{
		db:         db,
		ASNInfo:    q.ASNInfo.clone(db),
		DNSEntry:   q.DNSEntry.clone(db),
		Domain:     q.Domain.clone(db),
		IPAddress:  q.IPAddress.clone(db),
		Nameserver: q.Nameserver.clone(db),
		Path:       q.Path.clone(db),
		Registrar:  q.Registrar.clone(db),
		Whois:      q.Whois.clone(db),
	}
}

func (q *Query) ReadDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Read))
}

func (q *Query) WriteDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Write))
}

func (q *Query) ReplaceDB(db *gorm.DB) *Query {
	return &Query{
		db:         db,
		ASNInfo:    q.ASNInfo.replaceDB(db),
		DNSEntry:   q.DNSEntry.replaceDB(db),
		Domain:     q.Domain.replaceDB(db),
		IPAddress:  q.IPAddress.replaceDB(db),
		Nameserver: q.Nameserver.replaceDB(db),
		Path:       q.Path.replaceDB(db),
		Registrar:  q.Registrar.replaceDB(db),
		Whois:      q.Whois.replaceDB(db),
	}
}

type queryCtx struct {
	ASNInfo    IASNInfoDo
	DNSEntry   IDNSEntryDo
	Domain     IDomainDo
	IPAddress  IIPAddressDo
	Nameserver INameserverDo
	Path       IPathDo
	Registrar  IRegistrarDo
	Whois      IWhoisDo
}

func (q *Query) WithContext(ctx context.Context) *queryCtx {
	return &queryCtx{
		ASNInfo:    q.ASNInfo.WithContext(ctx),
		DNSEntry:   q.DNSEntry.WithContext(ctx),
		Domain:     q.Domain.WithContext(ctx),
		IPAddress:  q.IPAddress.WithContext(ctx),
		Nameserver: q.Nameserver.WithContext(ctx),
		Path:       q.Path.WithContext(ctx),
		Registrar:  q.Registrar.WithContext(ctx),
		Whois:      q.Whois.WithContext(ctx),
	}
}

func (q *Query) Transaction(fc func(tx *Query) error, opts ...*sql.TxOptions) error {
	return q.db.Transaction(func(tx *gorm.DB) error { return fc(q.clone(tx)) }, opts...)
}

func (q *Query) Begin(opts ...*sql.TxOptions) *QueryTx {
	tx := q.db.Begin(opts...)
	return &QueryTx{Query: q.clone(tx), Error: tx.Error}
}

type QueryTx struct {
	*Query
	Error error
}

func (q *QueryTx) Commit() error {
	return q.db.Commit().Error
}

func (q *QueryTx) Rollback() error {
	return q.db.Rollback().Error
}

func (q *QueryTx) SavePoint(name string) error {
	return q.db.SavePoint(name).Error
}

func (q *QueryTx) RollbackTo(name string) error {
	return q.db.RollbackTo(name).Error
}
