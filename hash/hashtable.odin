package hash

import "base:runtime"
import "core:fmt"
import "core:mem"
import "core:testing"

INITIAL_CAPACITY :: 8
MAX_LOAD_FACTOR :: 0.75

Entry :: struct {
	key:   string,
	value: any,
}

// Hash table with separate chaining, keyed by FNV-1a.
HashTable :: struct {
	buckets: [dynamic][dynamic]Entry,
	count:   int,
}

hashtable_init :: proc(allocator := context.allocator) -> HashTable {
	return HashTable{buckets = make([dynamic][dynamic]Entry, INITIAL_CAPACITY, allocator)}
}

hashtable_destroy :: proc(t: ^HashTable) {
	for bucket in t.buckets {
		for entry in bucket {
			_free_value(entry.value)
		}
		delete(bucket)
	}
	delete(t.buckets)
	t.count = 0
}

// Copy the data behind an `any` onto the heap. This matters because an
// `any` made from a variable points at that variable's memory rather than
// owning a copy, so storing it directly would alias the caller's variable.
@(private)
_clone_value :: proc(v: any, allocator := context.allocator) -> any {
	if v == nil {
		return nil
	}
	ti := type_info_of(v.id)
	data, _ := mem.alloc(ti.size, ti.align, allocator)
	mem.copy_non_overlapping(data, v.data, ti.size)
	return any{data, v.id}
}

@(private)
_free_value :: proc(v: any, allocator := context.allocator) {
	if v != nil {
		free(v.data, allocator)
	}
}

@(private)
_bucket :: proc(t: ^HashTable, key: string) -> ^[dynamic]Entry {
	i := int(fnv1a(key) % u64(len(t.buckets)))
	return &t.buckets[i]
}

@(private)
_resize :: proc(t: ^HashTable) {
	old := t.buckets
	t.buckets = make([dynamic][dynamic]Entry, len(old) * 2, context.allocator)
	for bucket in old {
		for entry in bucket {
			b := _bucket(t, entry.key)
			append(b, entry)
		}
		delete(bucket)
	}
	delete(old)
}

// Insert or overwrite the value for key. The table stores a copy of the
// value's underlying data and frees it on overwrite, remove, or destroy.
hashtable_insert :: proc(t: ^HashTable, key: string, value: any) {
	bucket := _bucket(t, key)
	for &entry in bucket {
		if entry.key == key {
			_free_value(entry.value)
			entry.value = _clone_value(value)
			return
		}
	}
	append(bucket, Entry{key = key, value = _clone_value(value)})
	t.count += 1
	if f64(t.count) / f64(len(t.buckets)) > MAX_LOAD_FACTOR {
		_resize(t)
	}
}

// Returns the value for key, or ok = false if absent.
hashtable_get :: proc(t: ^HashTable, key: string) -> (value: any, ok: bool) {
	for entry in _bucket(t, key) {
		if entry.key == key {
			return entry.value, true
		}
	}
	return nil, false
}

// Removes key. Returns false if absent (the Python version raises KeyError).
hashtable_remove :: proc(t: ^HashTable, key: string) -> bool {
	bucket := _bucket(t, key)
	for entry, i in bucket {
		if entry.key == key {
			_free_value(entry.value)
			ordered_remove(bucket, i)
			t.count -= 1
			return true
		}
	}
	return false
}

hashtable_contains :: proc(t: ^HashTable, key: string) -> bool {
	for entry in _bucket(t, key) {
		if entry.key == key {
			return true
		}
	}
	return false
}

hashtable_len :: proc(t: ^HashTable) -> int {
	return t.count
}

// ---------------------------------------------------------------------------
// Tests
// ---------------------------------------------------------------------------

@(test)
test_insert_and_get :: proc(t: ^testing.T) {
	table := hashtable_init()
	defer hashtable_destroy(&table)

	hashtable_insert(&table, "foo", 42)
	hashtable_insert(&table, "bar", "hello")

	v, ok := hashtable_get(&table, "foo")
	testing.expect(t, ok, "foo should exist")
	testing.expect_value(t, v.(int), 42)

	v, ok = hashtable_get(&table, "bar")
	testing.expect(t, ok, "bar should exist")
	testing.expect_value(t, v.(string), "hello")

	testing.expect_value(t, hashtable_len(&table), 2)
}

@(test)
test_insert_overwrites :: proc(t: ^testing.T) {
	table := hashtable_init()
	defer hashtable_destroy(&table)

	hashtable_insert(&table, "key", 1)
	hashtable_insert(&table, "key", 2)

	v, ok := hashtable_get(&table, "key")
	testing.expect(t, ok)
	testing.expect_value(t, v.(int), 2)
	testing.expect_value(t, hashtable_len(&table), 1)
}

@(test)
test_get_missing_key :: proc(t: ^testing.T) {
	table := hashtable_init()
	defer hashtable_destroy(&table)

	_, ok := hashtable_get(&table, "missing")
	testing.expect(t, !ok, "missing key should not be found")
}

@(test)
test_remove :: proc(t: ^testing.T) {
	table := hashtable_init()
	defer hashtable_destroy(&table)

	hashtable_insert(&table, "a", 1)
	hashtable_insert(&table, "b", 2)
	testing.expect_value(t, hashtable_len(&table), 2)

	testing.expect(t, hashtable_remove(&table, "a"))
	testing.expect(t, !hashtable_contains(&table, "a"))
	testing.expect_value(t, hashtable_len(&table), 1)

	// Removing again fails (the Python version raises KeyError)
	testing.expect(t, !hashtable_remove(&table, "a"))

	// "b" is untouched
	v, ok := hashtable_get(&table, "b")
	testing.expect(t, ok)
	testing.expect_value(t, v.(int), 2)
}

@(test)
test_contains :: proc(t: ^testing.T) {
	table := hashtable_init()
	defer hashtable_destroy(&table)

	hashtable_insert(&table, "present", true)
	testing.expect(t, hashtable_contains(&table, "present"))
	testing.expect(t, !hashtable_contains(&table, "absent"))
}

@(test)
test_resize :: proc(t: ^testing.T) {
	table := hashtable_init()
	defer hashtable_destroy(&table)

	// Keys are heap-allocated because the table stores key strings by
	// reference and rehashes them on resize.
	N :: 100
	keys: [N]string
	for &k, i in keys {
		k = fmt.aprintf("key%d", i)
	}
	defer for k in keys {
		delete(k)
	}

	// Insert enough entries to force several resizes (load factor 0.75, cap 8)
	for k, i in keys {
		hashtable_insert(&table, k, i)
	}

	testing.expect_value(t, hashtable_len(&table), N)
	testing.expect(t, len(table.buckets) > INITIAL_CAPACITY, "table should have grown")

	// All entries must still be retrievable after rehashing
	for k, i in keys {
		v, ok := hashtable_get(&table, k)
		testing.expectf(t, ok, "key %q lost during resize", k)
		testing.expect_value(t, v.(int), i)
	}
}
