package cache

func (cache *Cache) InsertNumber(missingNumber, multipliedMissingNumber int) {
	node := newNode(missingNumber, multipliedMissingNumber)
	cache.List.append(node)
}
