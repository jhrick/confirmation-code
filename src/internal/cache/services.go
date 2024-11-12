package cache

import (
	"time"
)

type cacheManagerSignal int8

const (
  NotFound cacheManagerSignal = -1
  Ok       cacheManagerSignal = 0
  Expired  cacheManagerSignal = 1
)

func (c *Cache) Store(value string) {
  c.storage[value] = time.Now()
}

func (c *Cache) Check(value string) cacheManagerSignal {
  codeGenTime, ok := c.storage[value] 
  if !ok {
    return NotFound
  }

  now := time.Now()
  difference := now.Sub(codeGenTime)
  if time.Duration(difference.Seconds()) >= time.Duration(c.DataExpTime.Seconds()) {
    delete(c.storage, value)
    return Expired
  }

  return Ok
}

func (c *Cache) Remove(value string) {
  delete(c.storage, value)
}
