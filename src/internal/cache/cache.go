package cache

import "time"

type Cache struct {
  storage map[string]time.Time
}

func Init() Cache {
  return Cache{
    storage: make(map[string]time.Time),
  }
}
