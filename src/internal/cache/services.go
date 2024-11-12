package cache

import "time"

func (c *Cache) Store(value string) {
  c.storage[value] = time.Now()
}

func (c *Cache) Check(value string) (found bool) {
  found = true

  if _, ok := c.storage[value]; !ok {
    found = false
  }

  return
}

func (c *Cache) Remove(value string) {
  delete(c.storage, value)
}
