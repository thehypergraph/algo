package hash

import "core:testing"

FNV_OFFSET_BASIS :: u64(0xCBF29CE484222325)
FNV_PRIME :: u64(0x100000001B3)

// FNV-1 Hashing Algorithm
// https://en.wikipedia.org/wiki/Fowler%E2%80%93Noll%E2%80%93Vo_hash_function#FNV-1_hash
// Note: u64 overflow wraps around in Odin, matching Python's & _MASK64.
fnv1 :: proc(data: string) -> u64 {
	h := FNV_OFFSET_BASIS
	for b in transmute([]u8)data {
		h = h * FNV_PRIME
		h ~= u64(b)
	}
	return h
}

// FNV-1a: same as FNV-1 but xor before multiply
// https://en.wikipedia.org/wiki/Fowler%E2%80%93Noll%E2%80%93Vo_hash_function#FNV-1a_hash
fnv1a :: proc(data: string) -> u64 {
	h := FNV_OFFSET_BASIS
	for b in transmute([]u8)data {
		h ~= u64(b)
		h = h * FNV_PRIME
	}
	return h
}

// ---------------------------------------------------------------------------
// Tests
// ---------------------------------------------------------------------------

@(test)
test_fnv1 :: proc(t: ^testing.T) {
	// Values verified against the Python implementation in fnv.py
	testing.expect_value(t, fnv1(""), 0xcbf29ce484222325)
	testing.expect_value(t, fnv1("a"), 0xaf63bd4c8601b7be)
	testing.expect_value(t, fnv1("hello"), 0x7b495389bdbdd4c7)
	testing.expect_value(t, fnv1("Hello, world!"), 0x6519bd6389aaa166)
	testing.expect_value(t, fnv1("foo"), 0xd8cbc7186ba13533)
	testing.expect_value(t, fnv1("bar"), 0xd8d9a5186bad3880)
}

@(test)
test_fnv1a :: proc(t: ^testing.T) {
	// Values verified against the Python implementation in fnv.py
	testing.expect_value(t, fnv1a(""), 0xcbf29ce484222325)
	testing.expect_value(t, fnv1a("a"), 0xaf63dc4c8601ec8c)
	testing.expect_value(t, fnv1a("hello"), 0xa430d84680aabd0b)
	testing.expect_value(t, fnv1a("Hello, world!"), 0x38d1334144987bf4)
	testing.expect_value(t, fnv1a("foo"), 0xdcb27518fed9d577)
	testing.expect_value(t, fnv1a("bar"), 0x003934191339461a)
}

@(test)
test_fnv1_vs_fnv1a_differ :: proc(t: ^testing.T) {
	// Order of xor/multiply matters for any non-empty input
	testing.expect(t, fnv1("hello") != fnv1a("hello"))
}
