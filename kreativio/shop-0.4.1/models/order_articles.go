// Code generated by SQLBoiler 4.2.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/v4/types"
	"github.com/volatiletech/strmangle"
)

// OrderArticle is an object representing the database table.
type OrderArticle struct {
	OrderID   int           `boil:"order_id" json:"order_id" toml:"order_id" yaml:"order_id"`
	ArticleID int           `boil:"article_id" json:"article_id" toml:"article_id" yaml:"article_id"`
	Amount    int           `boil:"amount" json:"amount" toml:"amount" yaml:"amount"`
	ID        int           `boil:"id" json:"id" toml:"id" yaml:"id"`
	Title     string        `boil:"title" json:"title" toml:"title" yaml:"title"`
	Price     types.Decimal `boil:"price" json:"price" toml:"price" yaml:"price"`
	Details   null.JSON     `boil:"details" json:"details,omitempty" toml:"details" yaml:"details,omitempty"`

	R *orderArticleR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L orderArticleL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var OrderArticleColumns = struct {
	OrderID   string
	ArticleID string
	Amount    string
	ID        string
	Title     string
	Price     string
	Details   string
}{
	OrderID:   "order_id",
	ArticleID: "article_id",
	Amount:    "amount",
	ID:        "id",
	Title:     "title",
	Price:     "price",
	Details:   "details",
}

// Generated where

type whereHelpernull_JSON struct{ field string }

func (w whereHelpernull_JSON) EQ(x null.JSON) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_JSON) NEQ(x null.JSON) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_JSON) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_JSON) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }
func (w whereHelpernull_JSON) LT(x null.JSON) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_JSON) LTE(x null.JSON) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_JSON) GT(x null.JSON) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_JSON) GTE(x null.JSON) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

var OrderArticleWhere = struct {
	OrderID   whereHelperint
	ArticleID whereHelperint
	Amount    whereHelperint
	ID        whereHelperint
	Title     whereHelperstring
	Price     whereHelpertypes_Decimal
	Details   whereHelpernull_JSON
}{
	OrderID:   whereHelperint{field: "\"shop\".\"order_articles\".\"order_id\""},
	ArticleID: whereHelperint{field: "\"shop\".\"order_articles\".\"article_id\""},
	Amount:    whereHelperint{field: "\"shop\".\"order_articles\".\"amount\""},
	ID:        whereHelperint{field: "\"shop\".\"order_articles\".\"id\""},
	Title:     whereHelperstring{field: "\"shop\".\"order_articles\".\"title\""},
	Price:     whereHelpertypes_Decimal{field: "\"shop\".\"order_articles\".\"price\""},
	Details:   whereHelpernull_JSON{field: "\"shop\".\"order_articles\".\"details\""},
}

// OrderArticleRels is where relationship names are stored.
var OrderArticleRels = struct {
	Order string
}{
	Order: "Order",
}

// orderArticleR is where relationships are stored.
type orderArticleR struct {
	Order *Order `boil:"Order" json:"Order" toml:"Order" yaml:"Order"`
}

// NewStruct creates a new relationship struct
func (*orderArticleR) NewStruct() *orderArticleR {
	return &orderArticleR{}
}

// orderArticleL is where Load methods for each relationship are stored.
type orderArticleL struct{}

var (
	orderArticleAllColumns            = []string{"order_id", "article_id", "amount", "id", "title", "price", "details"}
	orderArticleColumnsWithoutDefault = []string{"order_id", "article_id", "amount", "title", "price", "details"}
	orderArticleColumnsWithDefault    = []string{"id"}
	orderArticlePrimaryKeyColumns     = []string{"id"}
)

type (
	// OrderArticleSlice is an alias for a slice of pointers to OrderArticle.
	// This should generally be used opposed to []OrderArticle.
	OrderArticleSlice []*OrderArticle
	// OrderArticleHook is the signature for custom OrderArticle hook methods
	OrderArticleHook func(context.Context, boil.ContextExecutor, *OrderArticle) error

	orderArticleQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	orderArticleType                 = reflect.TypeOf(&OrderArticle{})
	orderArticleMapping              = queries.MakeStructMapping(orderArticleType)
	orderArticlePrimaryKeyMapping, _ = queries.BindMapping(orderArticleType, orderArticleMapping, orderArticlePrimaryKeyColumns)
	orderArticleInsertCacheMut       sync.RWMutex
	orderArticleInsertCache          = make(map[string]insertCache)
	orderArticleUpdateCacheMut       sync.RWMutex
	orderArticleUpdateCache          = make(map[string]updateCache)
	orderArticleUpsertCacheMut       sync.RWMutex
	orderArticleUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var orderArticleBeforeInsertHooks []OrderArticleHook
var orderArticleBeforeUpdateHooks []OrderArticleHook
var orderArticleBeforeDeleteHooks []OrderArticleHook
var orderArticleBeforeUpsertHooks []OrderArticleHook

var orderArticleAfterInsertHooks []OrderArticleHook
var orderArticleAfterSelectHooks []OrderArticleHook
var orderArticleAfterUpdateHooks []OrderArticleHook
var orderArticleAfterDeleteHooks []OrderArticleHook
var orderArticleAfterUpsertHooks []OrderArticleHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *OrderArticle) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range orderArticleBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *OrderArticle) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range orderArticleBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *OrderArticle) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range orderArticleBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *OrderArticle) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range orderArticleBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *OrderArticle) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range orderArticleAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *OrderArticle) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range orderArticleAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *OrderArticle) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range orderArticleAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *OrderArticle) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range orderArticleAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *OrderArticle) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range orderArticleAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddOrderArticleHook registers your hook function for all future operations.
func AddOrderArticleHook(hookPoint boil.HookPoint, orderArticleHook OrderArticleHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		orderArticleBeforeInsertHooks = append(orderArticleBeforeInsertHooks, orderArticleHook)
	case boil.BeforeUpdateHook:
		orderArticleBeforeUpdateHooks = append(orderArticleBeforeUpdateHooks, orderArticleHook)
	case boil.BeforeDeleteHook:
		orderArticleBeforeDeleteHooks = append(orderArticleBeforeDeleteHooks, orderArticleHook)
	case boil.BeforeUpsertHook:
		orderArticleBeforeUpsertHooks = append(orderArticleBeforeUpsertHooks, orderArticleHook)
	case boil.AfterInsertHook:
		orderArticleAfterInsertHooks = append(orderArticleAfterInsertHooks, orderArticleHook)
	case boil.AfterSelectHook:
		orderArticleAfterSelectHooks = append(orderArticleAfterSelectHooks, orderArticleHook)
	case boil.AfterUpdateHook:
		orderArticleAfterUpdateHooks = append(orderArticleAfterUpdateHooks, orderArticleHook)
	case boil.AfterDeleteHook:
		orderArticleAfterDeleteHooks = append(orderArticleAfterDeleteHooks, orderArticleHook)
	case boil.AfterUpsertHook:
		orderArticleAfterUpsertHooks = append(orderArticleAfterUpsertHooks, orderArticleHook)
	}
}

// One returns a single orderArticle record from the query.
func (q orderArticleQuery) One(ctx context.Context, exec boil.ContextExecutor) (*OrderArticle, error) {
	o := &OrderArticle{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for order_articles")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all OrderArticle records from the query.
func (q orderArticleQuery) All(ctx context.Context, exec boil.ContextExecutor) (OrderArticleSlice, error) {
	var o []*OrderArticle

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to OrderArticle slice")
	}

	if len(orderArticleAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all OrderArticle records in the query.
func (q orderArticleQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count order_articles rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q orderArticleQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if order_articles exists")
	}

	return count > 0, nil
}

// Order pointed to by the foreign key.
func (o *OrderArticle) Order(mods ...qm.QueryMod) orderQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.OrderID),
	}

	queryMods = append(queryMods, mods...)

	query := Orders(queryMods...)
	queries.SetFrom(query.Query, "\"shop\".\"orders\"")

	return query
}

// LoadOrder allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (orderArticleL) LoadOrder(ctx context.Context, e boil.ContextExecutor, singular bool, maybeOrderArticle interface{}, mods queries.Applicator) error {
	var slice []*OrderArticle
	var object *OrderArticle

	if singular {
		object = maybeOrderArticle.(*OrderArticle)
	} else {
		slice = *maybeOrderArticle.(*[]*OrderArticle)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &orderArticleR{}
		}
		args = append(args, object.OrderID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &orderArticleR{}
			}

			for _, a := range args {
				if a == obj.OrderID {
					continue Outer
				}
			}

			args = append(args, obj.OrderID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`shop.orders`),
		qm.WhereIn(`shop.orders.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Order")
	}

	var resultSlice []*Order
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Order")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for orders")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for orders")
	}

	if len(orderArticleAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Order = foreign
		if foreign.R == nil {
			foreign.R = &orderR{}
		}
		foreign.R.OrderArticles = append(foreign.R.OrderArticles, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.OrderID == foreign.ID {
				local.R.Order = foreign
				if foreign.R == nil {
					foreign.R = &orderR{}
				}
				foreign.R.OrderArticles = append(foreign.R.OrderArticles, local)
				break
			}
		}
	}

	return nil
}

// SetOrder of the orderArticle to the related item.
// Sets o.R.Order to related.
// Adds o to related.R.OrderArticles.
func (o *OrderArticle) SetOrder(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Order) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"shop\".\"order_articles\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"order_id"}),
		strmangle.WhereClause("\"", "\"", 2, orderArticlePrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.OrderID = related.ID
	if o.R == nil {
		o.R = &orderArticleR{
			Order: related,
		}
	} else {
		o.R.Order = related
	}

	if related.R == nil {
		related.R = &orderR{
			OrderArticles: OrderArticleSlice{o},
		}
	} else {
		related.R.OrderArticles = append(related.R.OrderArticles, o)
	}

	return nil
}

// OrderArticles retrieves all the records using an executor.
func OrderArticles(mods ...qm.QueryMod) orderArticleQuery {
	mods = append(mods, qm.From("\"shop\".\"order_articles\""))
	return orderArticleQuery{NewQuery(mods...)}
}

// FindOrderArticle retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindOrderArticle(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*OrderArticle, error) {
	orderArticleObj := &OrderArticle{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"shop\".\"order_articles\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, orderArticleObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from order_articles")
	}

	return orderArticleObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *OrderArticle) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no order_articles provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(orderArticleColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	orderArticleInsertCacheMut.RLock()
	cache, cached := orderArticleInsertCache[key]
	orderArticleInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			orderArticleAllColumns,
			orderArticleColumnsWithDefault,
			orderArticleColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(orderArticleType, orderArticleMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(orderArticleType, orderArticleMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"shop\".\"order_articles\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"shop\".\"order_articles\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into order_articles")
	}

	if !cached {
		orderArticleInsertCacheMut.Lock()
		orderArticleInsertCache[key] = cache
		orderArticleInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the OrderArticle.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *OrderArticle) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	orderArticleUpdateCacheMut.RLock()
	cache, cached := orderArticleUpdateCache[key]
	orderArticleUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			orderArticleAllColumns,
			orderArticlePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update order_articles, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"shop\".\"order_articles\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, orderArticlePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(orderArticleType, orderArticleMapping, append(wl, orderArticlePrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update order_articles row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for order_articles")
	}

	if !cached {
		orderArticleUpdateCacheMut.Lock()
		orderArticleUpdateCache[key] = cache
		orderArticleUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q orderArticleQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for order_articles")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for order_articles")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o OrderArticleSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), orderArticlePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"shop\".\"order_articles\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, orderArticlePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in orderArticle slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all orderArticle")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *OrderArticle) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no order_articles provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(orderArticleColumnsWithDefault, o)

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

	orderArticleUpsertCacheMut.RLock()
	cache, cached := orderArticleUpsertCache[key]
	orderArticleUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			orderArticleAllColumns,
			orderArticleColumnsWithDefault,
			orderArticleColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			orderArticleAllColumns,
			orderArticlePrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert order_articles, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(orderArticlePrimaryKeyColumns))
			copy(conflict, orderArticlePrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"shop\".\"order_articles\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(orderArticleType, orderArticleMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(orderArticleType, orderArticleMapping, ret)
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

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert order_articles")
	}

	if !cached {
		orderArticleUpsertCacheMut.Lock()
		orderArticleUpsertCache[key] = cache
		orderArticleUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single OrderArticle record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *OrderArticle) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no OrderArticle provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), orderArticlePrimaryKeyMapping)
	sql := "DELETE FROM \"shop\".\"order_articles\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from order_articles")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for order_articles")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q orderArticleQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no orderArticleQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from order_articles")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for order_articles")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o OrderArticleSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(orderArticleBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), orderArticlePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"shop\".\"order_articles\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, orderArticlePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from orderArticle slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for order_articles")
	}

	if len(orderArticleAfterDeleteHooks) != 0 {
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
func (o *OrderArticle) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindOrderArticle(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *OrderArticleSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := OrderArticleSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), orderArticlePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"shop\".\"order_articles\".* FROM \"shop\".\"order_articles\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, orderArticlePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in OrderArticleSlice")
	}

	*o = slice

	return nil
}

// OrderArticleExists checks if the OrderArticle row exists.
func OrderArticleExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"shop\".\"order_articles\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if order_articles exists")
	}

	return exists, nil
}