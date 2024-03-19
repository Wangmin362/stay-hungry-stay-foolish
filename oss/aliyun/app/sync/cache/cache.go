package cache

import "sync"

type empty struct{}

func NewCache() *Cache {
	return &Cache{
		objs: make(map[string]empty),
	}
}

type Cache struct {
	lock sync.RWMutex
	objs map[string]empty
}

func (s *Cache) ObjExist(obj string) bool {
	s.lock.RLock()
	defer s.lock.RUnlock()

	_, ok := s.objs[obj]
	return ok
}

func (s *Cache) CacheObj(obj string) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.objs[obj] = empty{}
}

func (s *Cache) DeleteObj(obj string) {
	s.lock.Lock()
	defer s.lock.Unlock()

	delete(s.objs, obj)
}

func (s *Cache) Replace(dstObj, srcObj string) {
	s.lock.Lock()
	defer s.lock.Unlock()

	delete(s.objs, srcObj)
	s.objs[dstObj] = empty{}
}
