// Code generated by SQLBoiler 3.5.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/strmangle"
)

// Utxo is an object representing the database table.
type Utxo struct {
	ID           int64  `boil:"id" json:"id" toml:"id" yaml:"id"`
	Txid         string `boil:"txid" json:"txid" toml:"txid" yaml:"txid"`
	Idx          int64  `boil:"idx" json:"idx" toml:"idx" yaml:"idx"`
	Amount       int64  `boil:"amount" json:"amount" toml:"amount" yaml:"amount"`
	ScriptPubkey string `boil:"script_pubkey" json:"script_pubkey" toml:"script_pubkey" yaml:"script_pubkey"`

	R *utxoR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L utxoL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var UtxoColumns = struct {
	ID           string
	Txid         string
	Idx          string
	Amount       string
	ScriptPubkey string
}{
	ID:           "id",
	Txid:         "txid",
	Idx:          "idx",
	Amount:       "amount",
	ScriptPubkey: "script_pubkey",
}

// Generated where

var UtxoWhere = struct {
	ID           whereHelperint64
	Txid         whereHelperstring
	Idx          whereHelperint64
	Amount       whereHelperint64
	ScriptPubkey whereHelperstring
}{
	ID:           whereHelperint64{field: "\"utxo\".\"id\""},
	Txid:         whereHelperstring{field: "\"utxo\".\"txid\""},
	Idx:          whereHelperint64{field: "\"utxo\".\"idx\""},
	Amount:       whereHelperint64{field: "\"utxo\".\"amount\""},
	ScriptPubkey: whereHelperstring{field: "\"utxo\".\"script_pubkey\""},
}

// UtxoRels is where relationship names are stored.
var UtxoRels = struct {
}{}

// utxoR is where relationships are stored.
type utxoR struct {
}

// NewStruct creates a new relationship struct
func (*utxoR) NewStruct() *utxoR {
	return &utxoR{}
}

// utxoL is where Load methods for each relationship are stored.
type utxoL struct{}

var (
	utxoAllColumns            = []string{"id", "txid", "idx", "amount", "script_pubkey"}
	utxoColumnsWithoutDefault = []string{"txid", "idx", "amount", "script_pubkey"}
	utxoColumnsWithDefault    = []string{"id"}
	utxoPrimaryKeyColumns     = []string{"id"}
)

type (
	// UtxoSlice is an alias for a slice of pointers to Utxo.
	// This should generally be used opposed to []Utxo.
	UtxoSlice []*Utxo
	// UtxoHook is the signature for custom Utxo hook methods
	UtxoHook func(context.Context, boil.ContextExecutor, *Utxo) error

	utxoQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	utxoType                 = reflect.TypeOf(&Utxo{})
	utxoMapping              = queries.MakeStructMapping(utxoType)
	utxoPrimaryKeyMapping, _ = queries.BindMapping(utxoType, utxoMapping, utxoPrimaryKeyColumns)
	utxoInsertCacheMut       sync.RWMutex
	utxoInsertCache          = make(map[string]insertCache)
	utxoUpdateCacheMut       sync.RWMutex
	utxoUpdateCache          = make(map[string]updateCache)
	utxoUpsertCacheMut       sync.RWMutex
	utxoUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var utxoBeforeInsertHooks []UtxoHook
var utxoBeforeUpdateHooks []UtxoHook
var utxoBeforeDeleteHooks []UtxoHook
var utxoBeforeUpsertHooks []UtxoHook

var utxoAfterInsertHooks []UtxoHook
var utxoAfterSelectHooks []UtxoHook
var utxoAfterUpdateHooks []UtxoHook
var utxoAfterDeleteHooks []UtxoHook
var utxoAfterUpsertHooks []UtxoHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Utxo) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range utxoBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Utxo) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range utxoBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Utxo) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range utxoBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Utxo) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range utxoBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Utxo) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range utxoAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Utxo) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range utxoAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Utxo) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range utxoAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Utxo) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range utxoAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Utxo) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range utxoAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddUtxoHook registers your hook function for all future operations.
func AddUtxoHook(hookPoint boil.HookPoint, utxoHook UtxoHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		utxoBeforeInsertHooks = append(utxoBeforeInsertHooks, utxoHook)
	case boil.BeforeUpdateHook:
		utxoBeforeUpdateHooks = append(utxoBeforeUpdateHooks, utxoHook)
	case boil.BeforeDeleteHook:
		utxoBeforeDeleteHooks = append(utxoBeforeDeleteHooks, utxoHook)
	case boil.BeforeUpsertHook:
		utxoBeforeUpsertHooks = append(utxoBeforeUpsertHooks, utxoHook)
	case boil.AfterInsertHook:
		utxoAfterInsertHooks = append(utxoAfterInsertHooks, utxoHook)
	case boil.AfterSelectHook:
		utxoAfterSelectHooks = append(utxoAfterSelectHooks, utxoHook)
	case boil.AfterUpdateHook:
		utxoAfterUpdateHooks = append(utxoAfterUpdateHooks, utxoHook)
	case boil.AfterDeleteHook:
		utxoAfterDeleteHooks = append(utxoAfterDeleteHooks, utxoHook)
	case boil.AfterUpsertHook:
		utxoAfterUpsertHooks = append(utxoAfterUpsertHooks, utxoHook)
	}
}

// One returns a single utxo record from the query.
func (q utxoQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Utxo, error) {
	o := &Utxo{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for utxo")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Utxo records from the query.
func (q utxoQuery) All(ctx context.Context, exec boil.ContextExecutor) (UtxoSlice, error) {
	var o []*Utxo

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Utxo slice")
	}

	if len(utxoAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Utxo records in the query.
func (q utxoQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count utxo rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q utxoQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if utxo exists")
	}

	return count > 0, nil
}

// Utxos retrieves all the records using an executor.
func Utxos(mods ...qm.QueryMod) utxoQuery {
	mods = append(mods, qm.From("\"utxo\""))
	return utxoQuery{NewQuery(mods...)}
}

// FindUtxo retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindUtxo(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*Utxo, error) {
	utxoObj := &Utxo{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"utxo\" where \"id\"=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, utxoObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from utxo")
	}

	return utxoObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Utxo) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no utxo provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(utxoColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	utxoInsertCacheMut.RLock()
	cache, cached := utxoInsertCache[key]
	utxoInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			utxoAllColumns,
			utxoColumnsWithDefault,
			utxoColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(utxoType, utxoMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(utxoType, utxoMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"utxo\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"utxo\" () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT \"%s\" FROM \"utxo\" WHERE %s", strings.Join(returnColumns, "\",\""), strmangle.WhereClause("\"", "\"", 0, utxoPrimaryKeyColumns))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into utxo")
	}

	var lastID int64
	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = int64(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == utxoMapping["ID"] {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ID,
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, identifierCols...)
	}

	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for utxo")
	}

CacheNoHooks:
	if !cached {
		utxoInsertCacheMut.Lock()
		utxoInsertCache[key] = cache
		utxoInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Utxo.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Utxo) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	utxoUpdateCacheMut.RLock()
	cache, cached := utxoUpdateCache[key]
	utxoUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			utxoAllColumns,
			utxoPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update utxo, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"utxo\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 0, wl),
			strmangle.WhereClause("\"", "\"", 0, utxoPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(utxoType, utxoMapping, append(wl, utxoPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update utxo row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for utxo")
	}

	if !cached {
		utxoUpdateCacheMut.Lock()
		utxoUpdateCache[key] = cache
		utxoUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q utxoQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for utxo")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for utxo")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o UtxoSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), utxoPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"utxo\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, utxoPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in utxo slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all utxo")
	}
	return rowsAff, nil
}

// Delete deletes a single Utxo record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Utxo) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Utxo provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), utxoPrimaryKeyMapping)
	sql := "DELETE FROM \"utxo\" WHERE \"id\"=?"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from utxo")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for utxo")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q utxoQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no utxoQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from utxo")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for utxo")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o UtxoSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(utxoBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), utxoPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"utxo\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, utxoPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from utxo slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for utxo")
	}

	if len(utxoAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Utxo) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindUtxo(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *UtxoSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := UtxoSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), utxoPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"utxo\".* FROM \"utxo\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, utxoPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in UtxoSlice")
	}

	*o = slice

	return nil
}

// UtxoExists checks if the Utxo row exists.
func UtxoExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"utxo\" where \"id\"=? limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}

	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if utxo exists")
	}

	return exists, nil
}
