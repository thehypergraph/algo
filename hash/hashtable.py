from .fnv import fnv1a

_INITIAL_CAPACITY = 8
_MAX_LOAD_FACTOR = 0.75


class HashTable:
    """Hash table with separate chaining, keyed by FNV-1a."""

    def __init__(self):
        self._buckets: list[list[tuple[str, object]]] = [
            [] for _ in range(_INITIAL_CAPACITY)
        ]
        self._count = 0

    def _bucket(self, key: str) -> list[tuple[str, object]]:
        return self._buckets[fnv1a(key) % len(self._buckets)]

    def _resize(self) -> None:
        old = self._buckets
        self._buckets = [[] for _ in range(len(old) * 2)]
        for bucket in old:
            for key, value in bucket:
                self._bucket(key).append((key, value))

    def insert(self, key: str, value: object) -> None:
        bucket = self._bucket(key)
        for i, (k, _) in enumerate(bucket):
            if k == key:
                bucket[i] = (key, value)
                return
        bucket.append((key, value))
        self._count += 1
        if self._count / len(self._buckets) > _MAX_LOAD_FACTOR:
            self._resize()

    def get(self, key: str) -> object:
        for k, v in self._bucket(key):
            if k == key:
                return v
        raise KeyError(key)

    def remove(self, key: str) -> None:
        bucket = self._bucket(key)
        for i, (k, _) in enumerate(bucket):
            if k == key:
                del bucket[i]
                self._count -= 1
                return
        raise KeyError(key)

    def __len__(self) -> int:
        return self._count

    def __contains__(self, key: str) -> bool:
        return any(k == key for k, _ in self._bucket(key))

    __setitem__ = insert
    __getitem__ = get
    __delitem__ = remove
