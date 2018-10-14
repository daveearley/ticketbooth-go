// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/randomize"
	"github.com/volatiletech/sqlboiler/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testCustomers(t *testing.T) {
	t.Parallel()

	query := Customers()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testCustomersDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Customer{}
	if err = randomize.Struct(seed, o, customerDBTypes, true, customerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Customer struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Customers().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testCustomersQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Customer{}
	if err = randomize.Struct(seed, o, customerDBTypes, true, customerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Customer struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Customers().DeleteAll(tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Customers().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testCustomersSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Customer{}
	if err = randomize.Struct(seed, o, customerDBTypes, true, customerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Customer struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := CustomerSlice{o}

	if rowsAff, err := slice.DeleteAll(tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Customers().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testCustomersExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Customer{}
	if err = randomize.Struct(seed, o, customerDBTypes, true, customerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Customer struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := CustomerExists(tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Customer exists: %s", err)
	}
	if !e {
		t.Errorf("Expected CustomerExists to return true, but got false.")
	}
}

func testCustomersFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Customer{}
	if err = randomize.Struct(seed, o, customerDBTypes, true, customerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Customer struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	customerFound, err := FindCustomer(tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if customerFound == nil {
		t.Error("want a record, got nil")
	}
}

func testCustomersBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Customer{}
	if err = randomize.Struct(seed, o, customerDBTypes, true, customerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Customer struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Customers().Bind(nil, tx, o); err != nil {
		t.Error(err)
	}
}

func testCustomersOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Customer{}
	if err = randomize.Struct(seed, o, customerDBTypes, true, customerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Customer struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Customers().One(tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testCustomersAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	customerOne := &Customer{}
	customerTwo := &Customer{}
	if err = randomize.Struct(seed, customerOne, customerDBTypes, false, customerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Customer struct: %s", err)
	}
	if err = randomize.Struct(seed, customerTwo, customerDBTypes, false, customerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Customer struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = customerOne.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = customerTwo.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Customers().All(tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testCustomersCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	customerOne := &Customer{}
	customerTwo := &Customer{}
	if err = randomize.Struct(seed, customerOne, customerDBTypes, false, customerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Customer struct: %s", err)
	}
	if err = randomize.Struct(seed, customerTwo, customerDBTypes, false, customerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Customer struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = customerOne.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = customerTwo.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Customers().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func customerBeforeInsertHook(e boil.Executor, o *Customer) error {
	*o = Customer{}
	return nil
}

func customerAfterInsertHook(e boil.Executor, o *Customer) error {
	*o = Customer{}
	return nil
}

func customerAfterSelectHook(e boil.Executor, o *Customer) error {
	*o = Customer{}
	return nil
}

func customerBeforeUpdateHook(e boil.Executor, o *Customer) error {
	*o = Customer{}
	return nil
}

func customerAfterUpdateHook(e boil.Executor, o *Customer) error {
	*o = Customer{}
	return nil
}

func customerBeforeDeleteHook(e boil.Executor, o *Customer) error {
	*o = Customer{}
	return nil
}

func customerAfterDeleteHook(e boil.Executor, o *Customer) error {
	*o = Customer{}
	return nil
}

func customerBeforeUpsertHook(e boil.Executor, o *Customer) error {
	*o = Customer{}
	return nil
}

func customerAfterUpsertHook(e boil.Executor, o *Customer) error {
	*o = Customer{}
	return nil
}

func testCustomersHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &Customer{}
	o := &Customer{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, customerDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Customer object: %s", err)
	}

	AddCustomerHook(boil.BeforeInsertHook, customerBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	customerBeforeInsertHooks = []CustomerHook{}

	AddCustomerHook(boil.AfterInsertHook, customerAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	customerAfterInsertHooks = []CustomerHook{}

	AddCustomerHook(boil.AfterSelectHook, customerAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	customerAfterSelectHooks = []CustomerHook{}

	AddCustomerHook(boil.BeforeUpdateHook, customerBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	customerBeforeUpdateHooks = []CustomerHook{}

	AddCustomerHook(boil.AfterUpdateHook, customerAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	customerAfterUpdateHooks = []CustomerHook{}

	AddCustomerHook(boil.BeforeDeleteHook, customerBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	customerBeforeDeleteHooks = []CustomerHook{}

	AddCustomerHook(boil.AfterDeleteHook, customerAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	customerAfterDeleteHooks = []CustomerHook{}

	AddCustomerHook(boil.BeforeUpsertHook, customerBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	customerBeforeUpsertHooks = []CustomerHook{}

	AddCustomerHook(boil.AfterUpsertHook, customerAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	customerAfterUpsertHooks = []CustomerHook{}
}

func testCustomersInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Customer{}
	if err = randomize.Struct(seed, o, customerDBTypes, true, customerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Customer struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Customers().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testCustomersInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Customer{}
	if err = randomize.Struct(seed, o, customerDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Customer struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Whitelist(customerColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Customers().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testCustomerToManyAttendees(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()

	var a Customer
	var b, c Attendee

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, customerDBTypes, true, customerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Customer struct: %s", err)
	}

	if err := a.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, attendeeDBTypes, false, attendeeColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, attendeeDBTypes, false, attendeeColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.CustomerID = a.ID
	c.CustomerID = a.ID

	if err = b.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	attendee, err := a.Attendees().All(tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range attendee {
		if v.CustomerID == b.CustomerID {
			bFound = true
		}
		if v.CustomerID == c.CustomerID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := CustomerSlice{&a}
	if err = a.L.LoadAttendees(tx, false, (*[]*Customer)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Attendees); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.Attendees = nil
	if err = a.L.LoadAttendees(tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Attendees); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", attendee)
	}
}

func testCustomerToManyTransactions(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()

	var a Customer
	var b, c Transaction

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, customerDBTypes, true, customerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Customer struct: %s", err)
	}

	if err := a.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, transactionDBTypes, false, transactionColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, transactionDBTypes, false, transactionColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.CustomerID = a.ID
	c.CustomerID = a.ID

	if err = b.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	transaction, err := a.Transactions().All(tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range transaction {
		if v.CustomerID == b.CustomerID {
			bFound = true
		}
		if v.CustomerID == c.CustomerID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := CustomerSlice{&a}
	if err = a.L.LoadTransactions(tx, false, (*[]*Customer)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Transactions); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.Transactions = nil
	if err = a.L.LoadTransactions(tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Transactions); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", transaction)
	}
}

func testCustomerToManyAddOpAttendees(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()

	var a Customer
	var b, c, d, e Attendee

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, customerDBTypes, false, strmangle.SetComplement(customerPrimaryKeyColumns, customerColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Attendee{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, attendeeDBTypes, false, strmangle.SetComplement(attendeePrimaryKeyColumns, attendeeColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*Attendee{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddAttendees(tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.CustomerID {
			t.Error("foreign key was wrong value", a.ID, first.CustomerID)
		}
		if a.ID != second.CustomerID {
			t.Error("foreign key was wrong value", a.ID, second.CustomerID)
		}

		if first.R.Customer != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Customer != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.Attendees[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.Attendees[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.Attendees().Count(tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testCustomerToManyAddOpTransactions(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()

	var a Customer
	var b, c, d, e Transaction

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, customerDBTypes, false, strmangle.SetComplement(customerPrimaryKeyColumns, customerColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Transaction{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, transactionDBTypes, false, strmangle.SetComplement(transactionPrimaryKeyColumns, transactionColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*Transaction{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddTransactions(tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.CustomerID {
			t.Error("foreign key was wrong value", a.ID, first.CustomerID)
		}
		if a.ID != second.CustomerID {
			t.Error("foreign key was wrong value", a.ID, second.CustomerID)
		}

		if first.R.Customer != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Customer != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.Transactions[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.Transactions[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.Transactions().Count(tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testCustomerToOneTransactionUsingTransaction(t *testing.T) {

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()

	var local Customer
	var foreign Transaction

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, customerDBTypes, false, customerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Customer struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, transactionDBTypes, false, transactionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Transaction struct: %s", err)
	}

	if err := foreign.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.TransactionID = foreign.ID
	if err := local.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Transaction().One(tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := CustomerSlice{&local}
	if err = local.L.LoadTransaction(tx, false, (*[]*Customer)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Transaction == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Transaction = nil
	if err = local.L.LoadTransaction(tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Transaction == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testCustomerToOneEventUsingEvent(t *testing.T) {

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()

	var local Customer
	var foreign Event

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, customerDBTypes, false, customerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Customer struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, eventDBTypes, false, eventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Event struct: %s", err)
	}

	if err := foreign.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.EventID = foreign.ID
	if err := local.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Event().One(tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := CustomerSlice{&local}
	if err = local.L.LoadEvent(tx, false, (*[]*Customer)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Event == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Event = nil
	if err = local.L.LoadEvent(tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Event == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testCustomerToOneSetOpTransactionUsingTransaction(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()

	var a Customer
	var b, c Transaction

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, customerDBTypes, false, strmangle.SetComplement(customerPrimaryKeyColumns, customerColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, transactionDBTypes, false, strmangle.SetComplement(transactionPrimaryKeyColumns, transactionColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, transactionDBTypes, false, strmangle.SetComplement(transactionPrimaryKeyColumns, transactionColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Transaction{&b, &c} {
		err = a.SetTransaction(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Transaction != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.Customers[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.TransactionID != x.ID {
			t.Error("foreign key was wrong value", a.TransactionID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.TransactionID))
		reflect.Indirect(reflect.ValueOf(&a.TransactionID)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.TransactionID != x.ID {
			t.Error("foreign key was wrong value", a.TransactionID, x.ID)
		}
	}
}
func testCustomerToOneSetOpEventUsingEvent(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()

	var a Customer
	var b, c Event

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, customerDBTypes, false, strmangle.SetComplement(customerPrimaryKeyColumns, customerColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, eventDBTypes, false, strmangle.SetComplement(eventPrimaryKeyColumns, eventColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, eventDBTypes, false, strmangle.SetComplement(eventPrimaryKeyColumns, eventColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Event{&b, &c} {
		err = a.SetEvent(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Event != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.Customers[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.EventID != x.ID {
			t.Error("foreign key was wrong value", a.EventID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.EventID))
		reflect.Indirect(reflect.ValueOf(&a.EventID)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.EventID != x.ID {
			t.Error("foreign key was wrong value", a.EventID, x.ID)
		}
	}
}

func testCustomersReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Customer{}
	if err = randomize.Struct(seed, o, customerDBTypes, true, customerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Customer struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testCustomersReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Customer{}
	if err = randomize.Struct(seed, o, customerDBTypes, true, customerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Customer struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := CustomerSlice{o}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}

func testCustomersSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Customer{}
	if err = randomize.Struct(seed, o, customerDBTypes, true, customerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Customer struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Customers().All(tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	customerDBTypes = map[string]string{`CreatedAt`: `timestamp without time zone`, `DeletedAt`: `timestamp without time zone`, `Email`: `character varying`, `EventID`: `integer`, `FirstName`: `character varying`, `ID`: `integer`, `LastName`: `character varying`, `Password`: `character varying`, `Status`: `enum.customer_status('ACTIVE')`, `TransactionID`: `integer`, `UpdatedAt`: `timestamp without time zone`}
	_               = bytes.MinRead
)

func testCustomersUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(customerPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(customerColumns) == len(customerPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Customer{}
	if err = randomize.Struct(seed, o, customerDBTypes, true, customerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Customer struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Customers().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, customerDBTypes, true, customerPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Customer struct: %s", err)
	}

	if rowsAff, err := o.Update(tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testCustomersSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(customerColumns) == len(customerPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Customer{}
	if err = randomize.Struct(seed, o, customerDBTypes, true, customerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Customer struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Customers().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, customerDBTypes, true, customerPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Customer struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(customerColumns, customerPrimaryKeyColumns) {
		fields = customerColumns
	} else {
		fields = strmangle.SetComplement(
			customerColumns,
			customerPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := CustomerSlice{o}
	if rowsAff, err := slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testCustomersUpsert(t *testing.T) {
	t.Parallel()

	if len(customerColumns) == len(customerPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Customer{}
	if err = randomize.Struct(seed, &o, customerDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Customer struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Customer: %s", err)
	}

	count, err := Customers().Count(tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, customerDBTypes, false, customerPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Customer struct: %s", err)
	}

	if err = o.Upsert(tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Customer: %s", err)
	}

	count, err = Customers().Count(tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}