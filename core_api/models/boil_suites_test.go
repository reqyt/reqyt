// Code generated by SQLBoiler 3.6.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import "testing"

// This test suite runs each operation test in parallel.
// Example, if your database has 3 tables, the suite will run:
// table1, table2 and table3 Delete in parallel
// table1, table2 and table3 Insert in parallel, and so forth.
// It does NOT run each operation group in parallel.
// Separating the tests thusly grants avoidance of Postgres deadlocks.
func TestParent(t *testing.T) {
	t.Run("FunctionCalls", testFunctionCalls)
	t.Run("Functions", testFunctions)
	t.Run("Users", testUsers)
	t.Run("Wallets", testWallets)
}

func TestDelete(t *testing.T) {
	t.Run("FunctionCalls", testFunctionCallsDelete)
	t.Run("Functions", testFunctionsDelete)
	t.Run("Users", testUsersDelete)
	t.Run("Wallets", testWalletsDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("FunctionCalls", testFunctionCallsQueryDeleteAll)
	t.Run("Functions", testFunctionsQueryDeleteAll)
	t.Run("Users", testUsersQueryDeleteAll)
	t.Run("Wallets", testWalletsQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("FunctionCalls", testFunctionCallsSliceDeleteAll)
	t.Run("Functions", testFunctionsSliceDeleteAll)
	t.Run("Users", testUsersSliceDeleteAll)
	t.Run("Wallets", testWalletsSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("FunctionCalls", testFunctionCallsExists)
	t.Run("Functions", testFunctionsExists)
	t.Run("Users", testUsersExists)
	t.Run("Wallets", testWalletsExists)
}

func TestFind(t *testing.T) {
	t.Run("FunctionCalls", testFunctionCallsFind)
	t.Run("Functions", testFunctionsFind)
	t.Run("Users", testUsersFind)
	t.Run("Wallets", testWalletsFind)
}

func TestBind(t *testing.T) {
	t.Run("FunctionCalls", testFunctionCallsBind)
	t.Run("Functions", testFunctionsBind)
	t.Run("Users", testUsersBind)
	t.Run("Wallets", testWalletsBind)
}

func TestOne(t *testing.T) {
	t.Run("FunctionCalls", testFunctionCallsOne)
	t.Run("Functions", testFunctionsOne)
	t.Run("Users", testUsersOne)
	t.Run("Wallets", testWalletsOne)
}

func TestAll(t *testing.T) {
	t.Run("FunctionCalls", testFunctionCallsAll)
	t.Run("Functions", testFunctionsAll)
	t.Run("Users", testUsersAll)
	t.Run("Wallets", testWalletsAll)
}

func TestCount(t *testing.T) {
	t.Run("FunctionCalls", testFunctionCallsCount)
	t.Run("Functions", testFunctionsCount)
	t.Run("Users", testUsersCount)
	t.Run("Wallets", testWalletsCount)
}

func TestHooks(t *testing.T) {
	t.Run("FunctionCalls", testFunctionCallsHooks)
	t.Run("Functions", testFunctionsHooks)
	t.Run("Users", testUsersHooks)
	t.Run("Wallets", testWalletsHooks)
}

func TestInsert(t *testing.T) {
	t.Run("FunctionCalls", testFunctionCallsInsert)
	t.Run("FunctionCalls", testFunctionCallsInsertWhitelist)
	t.Run("Functions", testFunctionsInsert)
	t.Run("Functions", testFunctionsInsertWhitelist)
	t.Run("Users", testUsersInsert)
	t.Run("Users", testUsersInsertWhitelist)
	t.Run("Wallets", testWalletsInsert)
	t.Run("Wallets", testWalletsInsertWhitelist)
}

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {
	t.Run("FunctionCallToFunctionUsingWallet", testFunctionCallToOneFunctionUsingWallet)
	t.Run("FunctionCallToFunctionUsingFunction", testFunctionCallToOneFunctionUsingFunction)
}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {
	t.Run("FunctionToWalletFunctionCalls", testFunctionToManyWalletFunctionCalls)
	t.Run("FunctionToFunctionCalls", testFunctionToManyFunctionCalls)
	t.Run("UserToWallets", testUserToManyWallets)
	t.Run("WalletToUsers", testWalletToManyUsers)
}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {
	t.Run("FunctionCallToFunctionUsingWalletFunctionCalls", testFunctionCallToOneSetOpFunctionUsingWallet)
	t.Run("FunctionCallToFunctionUsingFunctionCalls", testFunctionCallToOneSetOpFunctionUsingFunction)
}

// TestToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneRemove(t *testing.T) {}

// TestOneToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneSet(t *testing.T) {}

// TestOneToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneRemove(t *testing.T) {}

// TestToManyAdd tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyAdd(t *testing.T) {
	t.Run("FunctionToWalletFunctionCalls", testFunctionToManyAddOpWalletFunctionCalls)
	t.Run("FunctionToFunctionCalls", testFunctionToManyAddOpFunctionCalls)
	t.Run("UserToWallets", testUserToManyAddOpWallets)
	t.Run("WalletToUsers", testWalletToManyAddOpUsers)
}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {
	t.Run("UserToWallets", testUserToManySetOpWallets)
	t.Run("WalletToUsers", testWalletToManySetOpUsers)
}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {
	t.Run("UserToWallets", testUserToManyRemoveOpWallets)
	t.Run("WalletToUsers", testWalletToManyRemoveOpUsers)
}

func TestReload(t *testing.T) {
	t.Run("FunctionCalls", testFunctionCallsReload)
	t.Run("Functions", testFunctionsReload)
	t.Run("Users", testUsersReload)
	t.Run("Wallets", testWalletsReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("FunctionCalls", testFunctionCallsReloadAll)
	t.Run("Functions", testFunctionsReloadAll)
	t.Run("Users", testUsersReloadAll)
	t.Run("Wallets", testWalletsReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("FunctionCalls", testFunctionCallsSelect)
	t.Run("Functions", testFunctionsSelect)
	t.Run("Users", testUsersSelect)
	t.Run("Wallets", testWalletsSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("FunctionCalls", testFunctionCallsUpdate)
	t.Run("Functions", testFunctionsUpdate)
	t.Run("Users", testUsersUpdate)
	t.Run("Wallets", testWalletsUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("FunctionCalls", testFunctionCallsSliceUpdateAll)
	t.Run("Functions", testFunctionsSliceUpdateAll)
	t.Run("Users", testUsersSliceUpdateAll)
	t.Run("Wallets", testWalletsSliceUpdateAll)
}
