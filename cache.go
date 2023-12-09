package cache

import "time"

type Second int

type Cache struct {
	values map[string]Item
}

func NewCache() *Cache {
	return &Cache{
		values: make(map[string]Item),
	}
}

func (c *Cache) Set(key string, value interface{}, lifetime Second) {
	c.values[key] = NewItem(key, value, lifetime)
}

func (c *Cache) Get(key string) interface{} {
	return c.values[key].value
}

func (c *Cache) Delete(key string) {
	delete(c.values, key)
}

type Item struct {
	key       string
	value     any
	lifetime  Second
	createdAt time.Time
}

func NewItem(key string, value any, lifetime Second) Item {
	return Item{
		key:       key,
		value:     value,
		lifetime:  lifetime,
		createdAt: time.Now(),
	}
}
