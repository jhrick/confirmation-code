package cache

import "time"

type Cache struct {
  storage     map[string]time.Time
  DataExpTime time.Duration
}

func Init(dataExpTime time.Duration) Cache {
  return Cache{
    storage: make(map[string]time.Time),
    DataExpTime: dataExpTime,
  }
}
