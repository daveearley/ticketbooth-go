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

func testQuestionOptions(t *testing.T) {
	t.Parallel()

	query := QuestionOptions()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testQuestionOptionsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &QuestionOption{}
	if err = randomize.Struct(seed, o, questionOptionDBTypes, true, questionOptionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize QuestionOption struct: %s", err)
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

	count, err := QuestionOptions().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testQuestionOptionsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &QuestionOption{}
	if err = randomize.Struct(seed, o, questionOptionDBTypes, true, questionOptionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize QuestionOption struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := QuestionOptions().DeleteAll(tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := QuestionOptions().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testQuestionOptionsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &QuestionOption{}
	if err = randomize.Struct(seed, o, questionOptionDBTypes, true, questionOptionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize QuestionOption struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := QuestionOptionSlice{o}

	if rowsAff, err := slice.DeleteAll(tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := QuestionOptions().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testQuestionOptionsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &QuestionOption{}
	if err = randomize.Struct(seed, o, questionOptionDBTypes, true, questionOptionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize QuestionOption struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := QuestionOptionExists(tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if QuestionOption exists: %s", err)
	}
	if !e {
		t.Errorf("Expected QuestionOptionExists to return true, but got false.")
	}
}

func testQuestionOptionsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &QuestionOption{}
	if err = randomize.Struct(seed, o, questionOptionDBTypes, true, questionOptionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize QuestionOption struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	questionOptionFound, err := FindQuestionOption(tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if questionOptionFound == nil {
		t.Error("want a record, got nil")
	}
}

func testQuestionOptionsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &QuestionOption{}
	if err = randomize.Struct(seed, o, questionOptionDBTypes, true, questionOptionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize QuestionOption struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = QuestionOptions().Bind(nil, tx, o); err != nil {
		t.Error(err)
	}
}

func testQuestionOptionsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &QuestionOption{}
	if err = randomize.Struct(seed, o, questionOptionDBTypes, true, questionOptionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize QuestionOption struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := QuestionOptions().One(tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testQuestionOptionsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	questionOptionOne := &QuestionOption{}
	questionOptionTwo := &QuestionOption{}
	if err = randomize.Struct(seed, questionOptionOne, questionOptionDBTypes, false, questionOptionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize QuestionOption struct: %s", err)
	}
	if err = randomize.Struct(seed, questionOptionTwo, questionOptionDBTypes, false, questionOptionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize QuestionOption struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = questionOptionOne.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = questionOptionTwo.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := QuestionOptions().All(tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testQuestionOptionsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	questionOptionOne := &QuestionOption{}
	questionOptionTwo := &QuestionOption{}
	if err = randomize.Struct(seed, questionOptionOne, questionOptionDBTypes, false, questionOptionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize QuestionOption struct: %s", err)
	}
	if err = randomize.Struct(seed, questionOptionTwo, questionOptionDBTypes, false, questionOptionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize QuestionOption struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = questionOptionOne.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = questionOptionTwo.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := QuestionOptions().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func questionOptionBeforeInsertHook(e boil.Executor, o *QuestionOption) error {
	*o = QuestionOption{}
	return nil
}

func questionOptionAfterInsertHook(e boil.Executor, o *QuestionOption) error {
	*o = QuestionOption{}
	return nil
}

func questionOptionAfterSelectHook(e boil.Executor, o *QuestionOption) error {
	*o = QuestionOption{}
	return nil
}

func questionOptionBeforeUpdateHook(e boil.Executor, o *QuestionOption) error {
	*o = QuestionOption{}
	return nil
}

func questionOptionAfterUpdateHook(e boil.Executor, o *QuestionOption) error {
	*o = QuestionOption{}
	return nil
}

func questionOptionBeforeDeleteHook(e boil.Executor, o *QuestionOption) error {
	*o = QuestionOption{}
	return nil
}

func questionOptionAfterDeleteHook(e boil.Executor, o *QuestionOption) error {
	*o = QuestionOption{}
	return nil
}

func questionOptionBeforeUpsertHook(e boil.Executor, o *QuestionOption) error {
	*o = QuestionOption{}
	return nil
}

func questionOptionAfterUpsertHook(e boil.Executor, o *QuestionOption) error {
	*o = QuestionOption{}
	return nil
}

func testQuestionOptionsHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &QuestionOption{}
	o := &QuestionOption{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, questionOptionDBTypes, false); err != nil {
		t.Errorf("Unable to randomize QuestionOption object: %s", err)
	}

	AddQuestionOptionHook(boil.BeforeInsertHook, questionOptionBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	questionOptionBeforeInsertHooks = []QuestionOptionHook{}

	AddQuestionOptionHook(boil.AfterInsertHook, questionOptionAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	questionOptionAfterInsertHooks = []QuestionOptionHook{}

	AddQuestionOptionHook(boil.AfterSelectHook, questionOptionAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	questionOptionAfterSelectHooks = []QuestionOptionHook{}

	AddQuestionOptionHook(boil.BeforeUpdateHook, questionOptionBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	questionOptionBeforeUpdateHooks = []QuestionOptionHook{}

	AddQuestionOptionHook(boil.AfterUpdateHook, questionOptionAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	questionOptionAfterUpdateHooks = []QuestionOptionHook{}

	AddQuestionOptionHook(boil.BeforeDeleteHook, questionOptionBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	questionOptionBeforeDeleteHooks = []QuestionOptionHook{}

	AddQuestionOptionHook(boil.AfterDeleteHook, questionOptionAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	questionOptionAfterDeleteHooks = []QuestionOptionHook{}

	AddQuestionOptionHook(boil.BeforeUpsertHook, questionOptionBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	questionOptionBeforeUpsertHooks = []QuestionOptionHook{}

	AddQuestionOptionHook(boil.AfterUpsertHook, questionOptionAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	questionOptionAfterUpsertHooks = []QuestionOptionHook{}
}

func testQuestionOptionsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &QuestionOption{}
	if err = randomize.Struct(seed, o, questionOptionDBTypes, true, questionOptionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize QuestionOption struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := QuestionOptions().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testQuestionOptionsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &QuestionOption{}
	if err = randomize.Struct(seed, o, questionOptionDBTypes, true); err != nil {
		t.Errorf("Unable to randomize QuestionOption struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Whitelist(questionOptionColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := QuestionOptions().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testQuestionOptionToOneQuestionUsingQuestion(t *testing.T) {

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()

	var local QuestionOption
	var foreign Question

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, questionOptionDBTypes, false, questionOptionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize QuestionOption struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, questionDBTypes, false, questionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Question struct: %s", err)
	}

	if err := foreign.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.QuestionID = foreign.ID
	if err := local.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Question().One(tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := QuestionOptionSlice{&local}
	if err = local.L.LoadQuestion(tx, false, (*[]*QuestionOption)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Question == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Question = nil
	if err = local.L.LoadQuestion(tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Question == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testQuestionOptionToOneSetOpQuestionUsingQuestion(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()

	var a QuestionOption
	var b, c Question

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, questionOptionDBTypes, false, strmangle.SetComplement(questionOptionPrimaryKeyColumns, questionOptionColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, questionDBTypes, false, strmangle.SetComplement(questionPrimaryKeyColumns, questionColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, questionDBTypes, false, strmangle.SetComplement(questionPrimaryKeyColumns, questionColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Question{&b, &c} {
		err = a.SetQuestion(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Question != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.QuestionOptions[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.QuestionID != x.ID {
			t.Error("foreign key was wrong value", a.QuestionID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.QuestionID))
		reflect.Indirect(reflect.ValueOf(&a.QuestionID)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.QuestionID != x.ID {
			t.Error("foreign key was wrong value", a.QuestionID, x.ID)
		}
	}
}

func testQuestionOptionsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &QuestionOption{}
	if err = randomize.Struct(seed, o, questionOptionDBTypes, true, questionOptionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize QuestionOption struct: %s", err)
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

func testQuestionOptionsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &QuestionOption{}
	if err = randomize.Struct(seed, o, questionOptionDBTypes, true, questionOptionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize QuestionOption struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := QuestionOptionSlice{o}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}

func testQuestionOptionsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &QuestionOption{}
	if err = randomize.Struct(seed, o, questionOptionDBTypes, true, questionOptionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize QuestionOption struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := QuestionOptions().All(tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	questionOptionDBTypes = map[string]string{`CreatedAt`: `timestamp without time zone`, `ID`: `integer`, `QuestionID`: `integer`, `Title`: `text`, `UpdatedAt`: `timestamp without time zone`}
	_                     = bytes.MinRead
)

func testQuestionOptionsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(questionOptionPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(questionOptionColumns) == len(questionOptionPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &QuestionOption{}
	if err = randomize.Struct(seed, o, questionOptionDBTypes, true, questionOptionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize QuestionOption struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := QuestionOptions().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, questionOptionDBTypes, true, questionOptionPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize QuestionOption struct: %s", err)
	}

	if rowsAff, err := o.Update(tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testQuestionOptionsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(questionOptionColumns) == len(questionOptionPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &QuestionOption{}
	if err = randomize.Struct(seed, o, questionOptionDBTypes, true, questionOptionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize QuestionOption struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := QuestionOptions().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, questionOptionDBTypes, true, questionOptionPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize QuestionOption struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(questionOptionColumns, questionOptionPrimaryKeyColumns) {
		fields = questionOptionColumns
	} else {
		fields = strmangle.SetComplement(
			questionOptionColumns,
			questionOptionPrimaryKeyColumns,
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

	slice := QuestionOptionSlice{o}
	if rowsAff, err := slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testQuestionOptionsUpsert(t *testing.T) {
	t.Parallel()

	if len(questionOptionColumns) == len(questionOptionPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := QuestionOption{}
	if err = randomize.Struct(seed, &o, questionOptionDBTypes, true); err != nil {
		t.Errorf("Unable to randomize QuestionOption struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert QuestionOption: %s", err)
	}

	count, err := QuestionOptions().Count(tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, questionOptionDBTypes, false, questionOptionPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize QuestionOption struct: %s", err)
	}

	if err = o.Upsert(tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert QuestionOption: %s", err)
	}

	count, err = QuestionOptions().Count(tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
