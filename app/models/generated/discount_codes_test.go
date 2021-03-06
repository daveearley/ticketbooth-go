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

func testDiscountCodes(t *testing.T) {
	t.Parallel()

	query := DiscountCodes()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testDiscountCodesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DiscountCode{}
	if err = randomize.Struct(seed, o, discountCodeDBTypes, true, discountCodeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscountCode struct: %s", err)
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

	count, err := DiscountCodes().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDiscountCodesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DiscountCode{}
	if err = randomize.Struct(seed, o, discountCodeDBTypes, true, discountCodeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscountCode struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := DiscountCodes().DeleteAll(tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := DiscountCodes().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDiscountCodesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DiscountCode{}
	if err = randomize.Struct(seed, o, discountCodeDBTypes, true, discountCodeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscountCode struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := DiscountCodeSlice{o}

	if rowsAff, err := slice.DeleteAll(tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := DiscountCodes().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDiscountCodesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DiscountCode{}
	if err = randomize.Struct(seed, o, discountCodeDBTypes, true, discountCodeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscountCode struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := DiscountCodeExists(tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if DiscountCode exists: %s", err)
	}
	if !e {
		t.Errorf("Expected DiscountCodeExists to return true, but got false.")
	}
}

func testDiscountCodesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DiscountCode{}
	if err = randomize.Struct(seed, o, discountCodeDBTypes, true, discountCodeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscountCode struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	discountCodeFound, err := FindDiscountCode(tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if discountCodeFound == nil {
		t.Error("want a record, got nil")
	}
}

func testDiscountCodesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DiscountCode{}
	if err = randomize.Struct(seed, o, discountCodeDBTypes, true, discountCodeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscountCode struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = DiscountCodes().Bind(nil, tx, o); err != nil {
		t.Error(err)
	}
}

func testDiscountCodesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DiscountCode{}
	if err = randomize.Struct(seed, o, discountCodeDBTypes, true, discountCodeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscountCode struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := DiscountCodes().One(tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testDiscountCodesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discountCodeOne := &DiscountCode{}
	discountCodeTwo := &DiscountCode{}
	if err = randomize.Struct(seed, discountCodeOne, discountCodeDBTypes, false, discountCodeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscountCode struct: %s", err)
	}
	if err = randomize.Struct(seed, discountCodeTwo, discountCodeDBTypes, false, discountCodeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscountCode struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = discountCodeOne.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = discountCodeTwo.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := DiscountCodes().All(tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testDiscountCodesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	discountCodeOne := &DiscountCode{}
	discountCodeTwo := &DiscountCode{}
	if err = randomize.Struct(seed, discountCodeOne, discountCodeDBTypes, false, discountCodeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscountCode struct: %s", err)
	}
	if err = randomize.Struct(seed, discountCodeTwo, discountCodeDBTypes, false, discountCodeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscountCode struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = discountCodeOne.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = discountCodeTwo.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := DiscountCodes().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func discountCodeBeforeInsertHook(e boil.Executor, o *DiscountCode) error {
	*o = DiscountCode{}
	return nil
}

func discountCodeAfterInsertHook(e boil.Executor, o *DiscountCode) error {
	*o = DiscountCode{}
	return nil
}

func discountCodeAfterSelectHook(e boil.Executor, o *DiscountCode) error {
	*o = DiscountCode{}
	return nil
}

func discountCodeBeforeUpdateHook(e boil.Executor, o *DiscountCode) error {
	*o = DiscountCode{}
	return nil
}

func discountCodeAfterUpdateHook(e boil.Executor, o *DiscountCode) error {
	*o = DiscountCode{}
	return nil
}

func discountCodeBeforeDeleteHook(e boil.Executor, o *DiscountCode) error {
	*o = DiscountCode{}
	return nil
}

func discountCodeAfterDeleteHook(e boil.Executor, o *DiscountCode) error {
	*o = DiscountCode{}
	return nil
}

func discountCodeBeforeUpsertHook(e boil.Executor, o *DiscountCode) error {
	*o = DiscountCode{}
	return nil
}

func discountCodeAfterUpsertHook(e boil.Executor, o *DiscountCode) error {
	*o = DiscountCode{}
	return nil
}

func testDiscountCodesHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &DiscountCode{}
	o := &DiscountCode{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, discountCodeDBTypes, false); err != nil {
		t.Errorf("Unable to randomize DiscountCode object: %s", err)
	}

	AddDiscountCodeHook(boil.BeforeInsertHook, discountCodeBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	discountCodeBeforeInsertHooks = []DiscountCodeHook{}

	AddDiscountCodeHook(boil.AfterInsertHook, discountCodeAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	discountCodeAfterInsertHooks = []DiscountCodeHook{}

	AddDiscountCodeHook(boil.AfterSelectHook, discountCodeAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	discountCodeAfterSelectHooks = []DiscountCodeHook{}

	AddDiscountCodeHook(boil.BeforeUpdateHook, discountCodeBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	discountCodeBeforeUpdateHooks = []DiscountCodeHook{}

	AddDiscountCodeHook(boil.AfterUpdateHook, discountCodeAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	discountCodeAfterUpdateHooks = []DiscountCodeHook{}

	AddDiscountCodeHook(boil.BeforeDeleteHook, discountCodeBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	discountCodeBeforeDeleteHooks = []DiscountCodeHook{}

	AddDiscountCodeHook(boil.AfterDeleteHook, discountCodeAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	discountCodeAfterDeleteHooks = []DiscountCodeHook{}

	AddDiscountCodeHook(boil.BeforeUpsertHook, discountCodeBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	discountCodeBeforeUpsertHooks = []DiscountCodeHook{}

	AddDiscountCodeHook(boil.AfterUpsertHook, discountCodeAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	discountCodeAfterUpsertHooks = []DiscountCodeHook{}
}

func testDiscountCodesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DiscountCode{}
	if err = randomize.Struct(seed, o, discountCodeDBTypes, true, discountCodeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscountCode struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := DiscountCodes().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testDiscountCodesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DiscountCode{}
	if err = randomize.Struct(seed, o, discountCodeDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DiscountCode struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Whitelist(discountCodeColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := DiscountCodes().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testDiscountCodeToManyTransactionDiscountCodes(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()

	var a DiscountCode
	var b, c TransactionDiscountCode

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, discountCodeDBTypes, true, discountCodeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscountCode struct: %s", err)
	}

	if err := a.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, transactionDiscountCodeDBTypes, false, transactionDiscountCodeColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, transactionDiscountCodeDBTypes, false, transactionDiscountCodeColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.DiscountCodeID = a.ID
	c.DiscountCodeID = a.ID

	if err = b.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	transactionDiscountCode, err := a.TransactionDiscountCodes().All(tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range transactionDiscountCode {
		if v.DiscountCodeID == b.DiscountCodeID {
			bFound = true
		}
		if v.DiscountCodeID == c.DiscountCodeID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := DiscountCodeSlice{&a}
	if err = a.L.LoadTransactionDiscountCodes(tx, false, (*[]*DiscountCode)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.TransactionDiscountCodes); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.TransactionDiscountCodes = nil
	if err = a.L.LoadTransactionDiscountCodes(tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.TransactionDiscountCodes); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", transactionDiscountCode)
	}
}

func testDiscountCodeToManyAddOpTransactionDiscountCodes(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()

	var a DiscountCode
	var b, c, d, e TransactionDiscountCode

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, discountCodeDBTypes, false, strmangle.SetComplement(discountCodePrimaryKeyColumns, discountCodeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*TransactionDiscountCode{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, transactionDiscountCodeDBTypes, false, strmangle.SetComplement(transactionDiscountCodePrimaryKeyColumns, transactionDiscountCodeColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*TransactionDiscountCode{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddTransactionDiscountCodes(tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.DiscountCodeID {
			t.Error("foreign key was wrong value", a.ID, first.DiscountCodeID)
		}
		if a.ID != second.DiscountCodeID {
			t.Error("foreign key was wrong value", a.ID, second.DiscountCodeID)
		}

		if first.R.DiscountCode != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.DiscountCode != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.TransactionDiscountCodes[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.TransactionDiscountCodes[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.TransactionDiscountCodes().Count(tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testDiscountCodesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DiscountCode{}
	if err = randomize.Struct(seed, o, discountCodeDBTypes, true, discountCodeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscountCode struct: %s", err)
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

func testDiscountCodesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DiscountCode{}
	if err = randomize.Struct(seed, o, discountCodeDBTypes, true, discountCodeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscountCode struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := DiscountCodeSlice{o}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}

func testDiscountCodesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DiscountCode{}
	if err = randomize.Struct(seed, o, discountCodeDBTypes, true, discountCodeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscountCode struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := DiscountCodes().All(tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	discountCodeDBTypes = map[string]string{`Code`: `character varying`, `CreatedAt`: `timestamp without time zone`, `DeletedAt`: `timestamp without time zone`, `Discount`: `numeric`, `ID`: `integer`, `Type`: `enum.discount_code_type('ACTIVE')`, `UpdatedAt`: `timestamp without time zone`}
	_                   = bytes.MinRead
)

func testDiscountCodesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(discountCodePrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(discountCodeColumns) == len(discountCodePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &DiscountCode{}
	if err = randomize.Struct(seed, o, discountCodeDBTypes, true, discountCodeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscountCode struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := DiscountCodes().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, discountCodeDBTypes, true, discountCodePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DiscountCode struct: %s", err)
	}

	if rowsAff, err := o.Update(tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testDiscountCodesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(discountCodeColumns) == len(discountCodePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &DiscountCode{}
	if err = randomize.Struct(seed, o, discountCodeDBTypes, true, discountCodeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscountCode struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := DiscountCodes().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, discountCodeDBTypes, true, discountCodePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DiscountCode struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(discountCodeColumns, discountCodePrimaryKeyColumns) {
		fields = discountCodeColumns
	} else {
		fields = strmangle.SetComplement(
			discountCodeColumns,
			discountCodePrimaryKeyColumns,
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

	slice := DiscountCodeSlice{o}
	if rowsAff, err := slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testDiscountCodesUpsert(t *testing.T) {
	t.Parallel()

	if len(discountCodeColumns) == len(discountCodePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := DiscountCode{}
	if err = randomize.Struct(seed, &o, discountCodeDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DiscountCode struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert DiscountCode: %s", err)
	}

	count, err := DiscountCodes().Count(tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, discountCodeDBTypes, false, discountCodePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DiscountCode struct: %s", err)
	}

	if err = o.Upsert(tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert DiscountCode: %s", err)
	}

	count, err = DiscountCodes().Count(tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
