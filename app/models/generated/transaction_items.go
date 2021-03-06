// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/strmangle"
	"github.com/volatiletech/sqlboiler/types"
)

// TransactionItem is an object representing the database table.
type TransactionItem struct {
	ID            int           `boil:"id" json:"id" toml:"id" yaml:"id"`
	Total         types.Decimal `boil:"total" json:"total" toml:"total" yaml:"total"`
	TotalTax      types.Decimal `boil:"total_tax" json:"total_tax" toml:"total_tax" yaml:"total_tax"`
	TotalDiscount types.Decimal `boil:"total_discount" json:"total_discount" toml:"total_discount" yaml:"total_discount"`
	Quantity      int           `boil:"quantity" json:"quantity" toml:"quantity" yaml:"quantity"`
	TransactionID int           `boil:"transaction_id" json:"transaction_id" toml:"transaction_id" yaml:"transaction_id"`
	TicketID      int           `boil:"ticket_id" json:"ticket_id" toml:"ticket_id" yaml:"ticket_id"`

	R *transactionItemR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L transactionItemL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var TransactionItemColumns = struct {
	ID            string
	Total         string
	TotalTax      string
	TotalDiscount string
	Quantity      string
	TransactionID string
	TicketID      string
}{
	ID:            "id",
	Total:         "total",
	TotalTax:      "total_tax",
	TotalDiscount: "total_discount",
	Quantity:      "quantity",
	TransactionID: "transaction_id",
	TicketID:      "ticket_id",
}

// TransactionItemRels is where relationship names are stored.
var TransactionItemRels = struct {
	Transaction string
	Ticket      string
}{
	Transaction: "Transaction",
	Ticket:      "Ticket",
}

// transactionItemR is where relationships are stored.
type transactionItemR struct {
	Transaction *Transaction
	Ticket      *Ticket
}

// NewStruct creates a new relationship struct
func (*transactionItemR) NewStruct() *transactionItemR {
	return &transactionItemR{}
}

// transactionItemL is where Load methods for each relationship are stored.
type transactionItemL struct{}

var (
	transactionItemColumns               = []string{"id", "total", "total_tax", "total_discount", "quantity", "transaction_id", "ticket_id"}
	transactionItemColumnsWithoutDefault = []string{"total", "total_tax", "total_discount", "quantity", "transaction_id", "ticket_id"}
	transactionItemColumnsWithDefault    = []string{"id"}
	transactionItemPrimaryKeyColumns     = []string{"id"}
)

type (
	// TransactionItemSlice is an alias for a slice of pointers to TransactionItem.
	// This should generally be used opposed to []TransactionItem.
	TransactionItemSlice []*TransactionItem
	// TransactionItemHook is the signature for custom TransactionItem hook methods
	TransactionItemHook func(boil.Executor, *TransactionItem) error

	transactionItemQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	transactionItemType                 = reflect.TypeOf(&TransactionItem{})
	transactionItemMapping              = queries.MakeStructMapping(transactionItemType)
	transactionItemPrimaryKeyMapping, _ = queries.BindMapping(transactionItemType, transactionItemMapping, transactionItemPrimaryKeyColumns)
	transactionItemInsertCacheMut       sync.RWMutex
	transactionItemInsertCache          = make(map[string]insertCache)
	transactionItemUpdateCacheMut       sync.RWMutex
	transactionItemUpdateCache          = make(map[string]updateCache)
	transactionItemUpsertCacheMut       sync.RWMutex
	transactionItemUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
)

var transactionItemBeforeInsertHooks []TransactionItemHook
var transactionItemBeforeUpdateHooks []TransactionItemHook
var transactionItemBeforeDeleteHooks []TransactionItemHook
var transactionItemBeforeUpsertHooks []TransactionItemHook

var transactionItemAfterInsertHooks []TransactionItemHook
var transactionItemAfterSelectHooks []TransactionItemHook
var transactionItemAfterUpdateHooks []TransactionItemHook
var transactionItemAfterDeleteHooks []TransactionItemHook
var transactionItemAfterUpsertHooks []TransactionItemHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *TransactionItem) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range transactionItemBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *TransactionItem) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range transactionItemBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *TransactionItem) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range transactionItemBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *TransactionItem) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range transactionItemBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *TransactionItem) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range transactionItemAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *TransactionItem) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range transactionItemAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *TransactionItem) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range transactionItemAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *TransactionItem) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range transactionItemAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *TransactionItem) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range transactionItemAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddTransactionItemHook registers your hook function for all future operations.
func AddTransactionItemHook(hookPoint boil.HookPoint, transactionItemHook TransactionItemHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		transactionItemBeforeInsertHooks = append(transactionItemBeforeInsertHooks, transactionItemHook)
	case boil.BeforeUpdateHook:
		transactionItemBeforeUpdateHooks = append(transactionItemBeforeUpdateHooks, transactionItemHook)
	case boil.BeforeDeleteHook:
		transactionItemBeforeDeleteHooks = append(transactionItemBeforeDeleteHooks, transactionItemHook)
	case boil.BeforeUpsertHook:
		transactionItemBeforeUpsertHooks = append(transactionItemBeforeUpsertHooks, transactionItemHook)
	case boil.AfterInsertHook:
		transactionItemAfterInsertHooks = append(transactionItemAfterInsertHooks, transactionItemHook)
	case boil.AfterSelectHook:
		transactionItemAfterSelectHooks = append(transactionItemAfterSelectHooks, transactionItemHook)
	case boil.AfterUpdateHook:
		transactionItemAfterUpdateHooks = append(transactionItemAfterUpdateHooks, transactionItemHook)
	case boil.AfterDeleteHook:
		transactionItemAfterDeleteHooks = append(transactionItemAfterDeleteHooks, transactionItemHook)
	case boil.AfterUpsertHook:
		transactionItemAfterUpsertHooks = append(transactionItemAfterUpsertHooks, transactionItemHook)
	}
}

// One returns a single transactionItem record from the query.
func (q transactionItemQuery) One(exec boil.Executor) (*TransactionItem, error) {
	o := &TransactionItem{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(nil, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for transaction_items")
	}

	if err := o.doAfterSelectHooks(exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all TransactionItem records from the query.
func (q transactionItemQuery) All(exec boil.Executor) (TransactionItemSlice, error) {
	var o []*TransactionItem

	err := q.Bind(nil, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to TransactionItem slice")
	}

	if len(transactionItemAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all TransactionItem records in the query.
func (q transactionItemQuery) Count(exec boil.Executor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count transaction_items rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q transactionItemQuery) Exists(exec boil.Executor) (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if transaction_items exists")
	}

	return count > 0, nil
}

// Transaction pointed to by the foreign key.
func (o *TransactionItem) Transaction(mods ...qm.QueryMod) transactionQuery {
	queryMods := []qm.QueryMod{
		qm.Where("id=?", o.TransactionID),
	}

	queryMods = append(queryMods, mods...)

	query := Transactions(queryMods...)
	queries.SetFrom(query.Query, "\"transactions\"")

	return query
}

// Ticket pointed to by the foreign key.
func (o *TransactionItem) Ticket(mods ...qm.QueryMod) ticketQuery {
	queryMods := []qm.QueryMod{
		qm.Where("id=?", o.TicketID),
	}

	queryMods = append(queryMods, mods...)

	query := Tickets(queryMods...)
	queries.SetFrom(query.Query, "\"tickets\"")

	return query
}

// LoadTransaction allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (transactionItemL) LoadTransaction(e boil.Executor, singular bool, maybeTransactionItem interface{}, mods queries.Applicator) error {
	var slice []*TransactionItem
	var object *TransactionItem

	if singular {
		object = maybeTransactionItem.(*TransactionItem)
	} else {
		slice = *maybeTransactionItem.(*[]*TransactionItem)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &transactionItemR{}
		}
		args = append(args, object.TransactionID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &transactionItemR{}
			}

			for _, a := range args {
				if a == obj.TransactionID {
					continue Outer
				}
			}

			args = append(args, obj.TransactionID)
		}
	}

	query := NewQuery(qm.From(`transactions`), qm.WhereIn(`id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.Query(e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Transaction")
	}

	var resultSlice []*Transaction
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Transaction")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for transactions")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for transactions")
	}

	if len(transactionItemAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Transaction = foreign
		if foreign.R == nil {
			foreign.R = &transactionR{}
		}
		foreign.R.TransactionItems = append(foreign.R.TransactionItems, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.TransactionID == foreign.ID {
				local.R.Transaction = foreign
				if foreign.R == nil {
					foreign.R = &transactionR{}
				}
				foreign.R.TransactionItems = append(foreign.R.TransactionItems, local)
				break
			}
		}
	}

	return nil
}

// LoadTicket allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (transactionItemL) LoadTicket(e boil.Executor, singular bool, maybeTransactionItem interface{}, mods queries.Applicator) error {
	var slice []*TransactionItem
	var object *TransactionItem

	if singular {
		object = maybeTransactionItem.(*TransactionItem)
	} else {
		slice = *maybeTransactionItem.(*[]*TransactionItem)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &transactionItemR{}
		}
		args = append(args, object.TicketID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &transactionItemR{}
			}

			for _, a := range args {
				if a == obj.TicketID {
					continue Outer
				}
			}

			args = append(args, obj.TicketID)
		}
	}

	query := NewQuery(qm.From(`tickets`), qm.WhereIn(`id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.Query(e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Ticket")
	}

	var resultSlice []*Ticket
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Ticket")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for tickets")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for tickets")
	}

	if len(transactionItemAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Ticket = foreign
		if foreign.R == nil {
			foreign.R = &ticketR{}
		}
		foreign.R.TransactionItems = append(foreign.R.TransactionItems, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.TicketID == foreign.ID {
				local.R.Ticket = foreign
				if foreign.R == nil {
					foreign.R = &ticketR{}
				}
				foreign.R.TransactionItems = append(foreign.R.TransactionItems, local)
				break
			}
		}
	}

	return nil
}

// SetTransaction of the transactionItem to the related item.
// Sets o.R.Transaction to related.
// Adds o to related.R.TransactionItems.
func (o *TransactionItem) SetTransaction(exec boil.Executor, insert bool, related *Transaction) error {
	var err error
	if insert {
		if err = related.Insert(exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"transaction_items\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"transaction_id"}),
		strmangle.WhereClause("\"", "\"", 2, transactionItemPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.TransactionID = related.ID
	if o.R == nil {
		o.R = &transactionItemR{
			Transaction: related,
		}
	} else {
		o.R.Transaction = related
	}

	if related.R == nil {
		related.R = &transactionR{
			TransactionItems: TransactionItemSlice{o},
		}
	} else {
		related.R.TransactionItems = append(related.R.TransactionItems, o)
	}

	return nil
}

// SetTicket of the transactionItem to the related item.
// Sets o.R.Ticket to related.
// Adds o to related.R.TransactionItems.
func (o *TransactionItem) SetTicket(exec boil.Executor, insert bool, related *Ticket) error {
	var err error
	if insert {
		if err = related.Insert(exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"transaction_items\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"ticket_id"}),
		strmangle.WhereClause("\"", "\"", 2, transactionItemPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.TicketID = related.ID
	if o.R == nil {
		o.R = &transactionItemR{
			Ticket: related,
		}
	} else {
		o.R.Ticket = related
	}

	if related.R == nil {
		related.R = &ticketR{
			TransactionItems: TransactionItemSlice{o},
		}
	} else {
		related.R.TransactionItems = append(related.R.TransactionItems, o)
	}

	return nil
}

// TransactionItems retrieves all the records using an executor.
func TransactionItems(mods ...qm.QueryMod) transactionItemQuery {
	mods = append(mods, qm.From("\"transaction_items\""))
	return transactionItemQuery{NewQuery(mods...)}
}

// FindTransactionItem retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindTransactionItem(exec boil.Executor, iD int, selectCols ...string) (*TransactionItem, error) {
	transactionItemObj := &TransactionItem{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"transaction_items\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(nil, exec, transactionItemObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from transaction_items")
	}

	return transactionItemObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *TransactionItem) Insert(exec boil.Executor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no transaction_items provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(transactionItemColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	transactionItemInsertCacheMut.RLock()
	cache, cached := transactionItemInsertCache[key]
	transactionItemInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			transactionItemColumns,
			transactionItemColumnsWithDefault,
			transactionItemColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(transactionItemType, transactionItemMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(transactionItemType, transactionItemMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"transaction_items\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"transaction_items\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRow(cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.Exec(cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into transaction_items")
	}

	if !cached {
		transactionItemInsertCacheMut.Lock()
		transactionItemInsertCache[key] = cache
		transactionItemInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// Update uses an executor to update the TransactionItem.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *TransactionItem) Update(exec boil.Executor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	transactionItemUpdateCacheMut.RLock()
	cache, cached := transactionItemUpdateCache[key]
	transactionItemUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			transactionItemColumns,
			transactionItemPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update transaction_items, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"transaction_items\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, transactionItemPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(transactionItemType, transactionItemMapping, append(wl, transactionItemPrimaryKeyColumns...))
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
	result, err = exec.Exec(cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update transaction_items row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for transaction_items")
	}

	if !cached {
		transactionItemUpdateCacheMut.Lock()
		transactionItemUpdateCache[key] = cache
		transactionItemUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(exec)
}

// UpdateAll updates all rows with the specified column values.
func (q transactionItemQuery) UpdateAll(exec boil.Executor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.Exec(exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for transaction_items")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for transaction_items")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o TransactionItemSlice) UpdateAll(exec boil.Executor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), transactionItemPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"transaction_items\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, transactionItemPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in transactionItem slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all transactionItem")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *TransactionItem) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no transaction_items provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(transactionItemColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	transactionItemUpsertCacheMut.RLock()
	cache, cached := transactionItemUpsertCache[key]
	transactionItemUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			transactionItemColumns,
			transactionItemColumnsWithDefault,
			transactionItemColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			transactionItemColumns,
			transactionItemPrimaryKeyColumns,
		)

		if len(update) == 0 {
			return errors.New("models: unable to upsert transaction_items, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(transactionItemPrimaryKeyColumns))
			copy(conflict, transactionItemPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"transaction_items\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(transactionItemType, transactionItemMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(transactionItemType, transactionItemMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRow(cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.Exec(cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert transaction_items")
	}

	if !cached {
		transactionItemUpsertCacheMut.Lock()
		transactionItemUpsertCache[key] = cache
		transactionItemUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// Delete deletes a single TransactionItem record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *TransactionItem) Delete(exec boil.Executor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no TransactionItem provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), transactionItemPrimaryKeyMapping)
	sql := "DELETE FROM \"transaction_items\" WHERE \"id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from transaction_items")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for transaction_items")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q transactionItemQuery) DeleteAll(exec boil.Executor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no transactionItemQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.Exec(exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from transaction_items")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for transaction_items")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o TransactionItemSlice) DeleteAll(exec boil.Executor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no TransactionItem slice provided for delete all")
	}

	if len(o) == 0 {
		return 0, nil
	}

	if len(transactionItemBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), transactionItemPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"transaction_items\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, transactionItemPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from transactionItem slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for transaction_items")
	}

	if len(transactionItemAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *TransactionItem) Reload(exec boil.Executor) error {
	ret, err := FindTransactionItem(exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *TransactionItemSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := TransactionItemSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), transactionItemPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"transaction_items\".* FROM \"transaction_items\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, transactionItemPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(nil, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in TransactionItemSlice")
	}

	*o = slice

	return nil
}

// TransactionItemExists checks if the TransactionItem row exists.
func TransactionItemExists(exec boil.Executor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"transaction_items\" where \"id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}

	row := exec.QueryRow(sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if transaction_items exists")
	}

	return exists, nil
}
