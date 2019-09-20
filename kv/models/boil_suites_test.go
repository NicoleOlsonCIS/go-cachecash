// +build sqlboiler_test

// Code generated by SQLBoiler 3.5.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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
	t.Run("Kvstores", testKvstores)
}

func TestDelete(t *testing.T) {
	t.Run("Kvstores", testKvstoresDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("Kvstores", testKvstoresQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("Kvstores", testKvstoresSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("Kvstores", testKvstoresExists)
}

func TestFind(t *testing.T) {
	t.Run("Kvstores", testKvstoresFind)
}

func TestBind(t *testing.T) {
	t.Run("Kvstores", testKvstoresBind)
}

func TestOne(t *testing.T) {
	t.Run("Kvstores", testKvstoresOne)
}

func TestAll(t *testing.T) {
	t.Run("Kvstores", testKvstoresAll)
}

func TestCount(t *testing.T) {
	t.Run("Kvstores", testKvstoresCount)
}

func TestHooks(t *testing.T) {
	t.Run("Kvstores", testKvstoresHooks)
}

func TestInsert(t *testing.T) {
	t.Run("Kvstores", testKvstoresInsert)
	t.Run("Kvstores", testKvstoresInsertWhitelist)
}

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {}

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
func TestToManyAdd(t *testing.T) {}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {}

func TestReload(t *testing.T) {
	t.Run("Kvstores", testKvstoresReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("Kvstores", testKvstoresReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("Kvstores", testKvstoresSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("Kvstores", testKvstoresUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("Kvstores", testKvstoresSliceUpdateAll)
}
