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
	c.CheckExpired()
	c.values[key] = NewItem(key, value, lifetime)
}

func (c *Cache) Get(key string) interface{} {
	c.CheckExpired()
	return c.values[key].value
}

func (c *Cache) Delete(key string) {
	delete(c.values, key)
}

// CheckExpired Want to rewrite the method call to goroutines in the next version
func (c *Cache) CheckExpired() {
	currentTime := time.Now()

	for key, item := range c.values {
		expirationTime := item.createdAt.Add(time.Duration(item.lifetime) * time.Second)

		if currentTime.After(expirationTime) {
			delete(c.values, key)
		}
	}
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
