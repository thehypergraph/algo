FNV_OFFSET_BASIS = 0xCBF29CE484222325
FNV_PRIME = 0x100000001B3

_MASK64 = (1 << 64) - 1


def _to_bytes(data: str | bytes) -> bytes:
    return data.encode("utf-8") if isinstance(data, str) else data


# FNV-1 Hashing Algorithm
# https://en.wikipedia.org/wiki/Fowler%E2%80%93Noll%E2%80%93Vo_hash_function#FNV-1_hash
def fnv1(data: str | bytes) -> int:
    h = FNV_OFFSET_BASIS
    for b in _to_bytes(data):
        h = (h * FNV_PRIME) & _MASK64
        h ^= b
    return h


# FNV-1a: same as FNV-1 but xor before multiply
# https://en.wikipedia.org/wiki/Fowler%E2%80%93Noll%E2%80%93Vo_hash_function#FNV-1a_hash
def fnv1a(data: str | bytes) -> int:
    h = FNV_OFFSET_BASIS
    for b in _to_bytes(data):
        h ^= b
        h = (h * FNV_PRIME) & _MASK64
    return h
