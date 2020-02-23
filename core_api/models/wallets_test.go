// Code generated by SQLBoiler 3.6.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
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

func testWallets(t *testing.T) {
	t.Parallel()

	query := Wallets()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testWalletsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Wallet{}
	if err = randomize.Struct(seed, o, walletDBTypes, true, walletColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Wallet struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Wallets().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testWalletsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Wallet{}
	if err = randomize.Struct(seed, o, walletDBTypes, true, walletColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Wallet struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Wallets().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Wallets().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testWalletsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Wallet{}
	if err = randomize.Struct(seed, o, walletDBTypes, true, walletColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Wallet struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := WalletSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Wallets().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testWalletsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Wallet{}
	if err = randomize.Struct(seed, o, walletDBTypes, true, walletColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Wallet struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := WalletExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Wallet exists: %s", err)
	}
	if !e {
		t.Errorf("Expected WalletExists to return true, but got false.")
	}
}

func testWalletsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Wallet{}
	if err = randomize.Struct(seed, o, walletDBTypes, true, walletColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Wallet struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	walletFound, err := FindWallet(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if walletFound == nil {
		t.Error("want a record, got nil")
	}
}

func testWalletsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Wallet{}
	if err = randomize.Struct(seed, o, walletDBTypes, true, walletColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Wallet struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Wallets().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testWalletsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Wallet{}
	if err = randomize.Struct(seed, o, walletDBTypes, true, walletColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Wallet struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Wallets().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testWalletsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	walletOne := &Wallet{}
	walletTwo := &Wallet{}
	if err = randomize.Struct(seed, walletOne, walletDBTypes, false, walletColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Wallet struct: %s", err)
	}
	if err = randomize.Struct(seed, walletTwo, walletDBTypes, false, walletColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Wallet struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = walletOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = walletTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Wallets().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testWalletsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	walletOne := &Wallet{}
	walletTwo := &Wallet{}
	if err = randomize.Struct(seed, walletOne, walletDBTypes, false, walletColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Wallet struct: %s", err)
	}
	if err = randomize.Struct(seed, walletTwo, walletDBTypes, false, walletColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Wallet struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = walletOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = walletTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Wallets().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func walletBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Wallet) error {
	*o = Wallet{}
	return nil
}

func walletAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Wallet) error {
	*o = Wallet{}
	return nil
}

func walletAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Wallet) error {
	*o = Wallet{}
	return nil
}

func walletBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Wallet) error {
	*o = Wallet{}
	return nil
}

func walletAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Wallet) error {
	*o = Wallet{}
	return nil
}

func walletBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Wallet) error {
	*o = Wallet{}
	return nil
}

func walletAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Wallet) error {
	*o = Wallet{}
	return nil
}

func walletBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Wallet) error {
	*o = Wallet{}
	return nil
}

func walletAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Wallet) error {
	*o = Wallet{}
	return nil
}

func testWalletsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Wallet{}
	o := &Wallet{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, walletDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Wallet object: %s", err)
	}

	AddWalletHook(boil.BeforeInsertHook, walletBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	walletBeforeInsertHooks = []WalletHook{}

	AddWalletHook(boil.AfterInsertHook, walletAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	walletAfterInsertHooks = []WalletHook{}

	AddWalletHook(boil.AfterSelectHook, walletAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	walletAfterSelectHooks = []WalletHook{}

	AddWalletHook(boil.BeforeUpdateHook, walletBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	walletBeforeUpdateHooks = []WalletHook{}

	AddWalletHook(boil.AfterUpdateHook, walletAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	walletAfterUpdateHooks = []WalletHook{}

	AddWalletHook(boil.BeforeDeleteHook, walletBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	walletBeforeDeleteHooks = []WalletHook{}

	AddWalletHook(boil.AfterDeleteHook, walletAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	walletAfterDeleteHooks = []WalletHook{}

	AddWalletHook(boil.BeforeUpsertHook, walletBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	walletBeforeUpsertHooks = []WalletHook{}

	AddWalletHook(boil.AfterUpsertHook, walletAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	walletAfterUpsertHooks = []WalletHook{}
}

func testWalletsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Wallet{}
	if err = randomize.Struct(seed, o, walletDBTypes, true, walletColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Wallet struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Wallets().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testWalletsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Wallet{}
	if err = randomize.Struct(seed, o, walletDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Wallet struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(walletColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Wallets().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testWalletToManyUsers(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Wallet
	var b, c User

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, walletDBTypes, true, walletColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Wallet struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, userDBTypes, false, userColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, userDBTypes, false, userColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	_, err = tx.Exec("insert into \"rel_user_wallet\" (\"wallet_id\", \"user_id\") values (?, ?)", a.ID, b.ID)
	if err != nil {
		t.Fatal(err)
	}
	_, err = tx.Exec("insert into \"rel_user_wallet\" (\"wallet_id\", \"user_id\") values (?, ?)", a.ID, c.ID)
	if err != nil {
		t.Fatal(err)
	}

	check, err := a.Users().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.ID == b.ID {
			bFound = true
		}
		if v.ID == c.ID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := WalletSlice{&a}
	if err = a.L.LoadUsers(ctx, tx, false, (*[]*Wallet)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Users); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.Users = nil
	if err = a.L.LoadUsers(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Users); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testWalletToManyAddOpUsers(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Wallet
	var b, c, d, e User

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, walletDBTypes, false, strmangle.SetComplement(walletPrimaryKeyColumns, walletColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*User{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*User{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddUsers(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if first.R.Wallets[0] != &a {
			t.Error("relationship was not added properly to the slice")
		}
		if second.R.Wallets[0] != &a {
			t.Error("relationship was not added properly to the slice")
		}

		if a.R.Users[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.Users[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.Users().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testWalletToManySetOpUsers(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Wallet
	var b, c, d, e User

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, walletDBTypes, false, strmangle.SetComplement(walletPrimaryKeyColumns, walletColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*User{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err = a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.SetUsers(ctx, tx, false, &b, &c)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.Users().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	err = a.SetUsers(ctx, tx, true, &d, &e)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.Users().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	// The following checks cannot be implemented since we have no handle
	// to these when we call Set(). Leaving them here as wishful thinking
	// and to let people know there's dragons.
	//
	// if len(b.R.Wallets) != 0 {
	// 	t.Error("relationship was not removed properly from the slice")
	// }
	// if len(c.R.Wallets) != 0 {
	// 	t.Error("relationship was not removed properly from the slice")
	// }
	if d.R.Wallets[0] != &a {
		t.Error("relationship was not added properly to the slice")
	}
	if e.R.Wallets[0] != &a {
		t.Error("relationship was not added properly to the slice")
	}

	if a.R.Users[0] != &d {
		t.Error("relationship struct slice not set to correct value")
	}
	if a.R.Users[1] != &e {
		t.Error("relationship struct slice not set to correct value")
	}
}

func testWalletToManyRemoveOpUsers(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Wallet
	var b, c, d, e User

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, walletDBTypes, false, strmangle.SetComplement(walletPrimaryKeyColumns, walletColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*User{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.AddUsers(ctx, tx, true, foreigners...)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.Users().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 4 {
		t.Error("count was wrong:", count)
	}

	err = a.RemoveUsers(ctx, tx, foreigners[:2]...)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.Users().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if len(b.R.Wallets) != 0 {
		t.Error("relationship was not removed properly from the slice")
	}
	if len(c.R.Wallets) != 0 {
		t.Error("relationship was not removed properly from the slice")
	}
	if d.R.Wallets[0] != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}
	if e.R.Wallets[0] != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}

	if len(a.R.Users) != 2 {
		t.Error("should have preserved two relationships")
	}

	// Removal doesn't do a stable deletion for performance so we have to flip the order
	if a.R.Users[1] != &d {
		t.Error("relationship to d should have been preserved")
	}
	if a.R.Users[0] != &e {
		t.Error("relationship to e should have been preserved")
	}
}

func testWalletsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Wallet{}
	if err = randomize.Struct(seed, o, walletDBTypes, true, walletColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Wallet struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testWalletsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Wallet{}
	if err = randomize.Struct(seed, o, walletDBTypes, true, walletColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Wallet struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := WalletSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testWalletsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Wallet{}
	if err = randomize.Struct(seed, o, walletDBTypes, true, walletColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Wallet struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Wallets().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	walletDBTypes = map[string]string{`ID`: `INTEGER`, `CreatedAt`: `DATETIME`, `UpdatedAt`: `DATETIME`, `DeletedAt`: `DATETIME`, `Key`: `VARCHAR(255)`, `Credits`: `BIGINT`}
	_             = bytes.MinRead
)

func testWalletsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(walletPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(walletAllColumns) == len(walletPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Wallet{}
	if err = randomize.Struct(seed, o, walletDBTypes, true, walletColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Wallet struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Wallets().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, walletDBTypes, true, walletPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Wallet struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testWalletsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(walletAllColumns) == len(walletPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Wallet{}
	if err = randomize.Struct(seed, o, walletDBTypes, true, walletColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Wallet struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Wallets().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, walletDBTypes, true, walletPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Wallet struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(walletAllColumns, walletPrimaryKeyColumns) {
		fields = walletAllColumns
	} else {
		fields = strmangle.SetComplement(
			walletAllColumns,
			walletPrimaryKeyColumns,
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

	slice := WalletSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}