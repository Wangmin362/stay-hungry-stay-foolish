package cache

import "sync"

type Tag struct {
	WechatUrl string
}

func NewCache() *Cache {
	return &Cache{
		objs: make(map[string]Tag),
	}
}

type Cache struct {
	lock sync.RWMutex
	objs map[string]Tag
}

func (s *Cache) ObjExist(obj string) (bool, Tag) {
	s.lock.RLock()
	defer s.lock.RUnlock()

	tag, ok := s.objs[obj]
	return ok, tag
}

func (s *Cache) CacheObj(obj string, tag Tag) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.objs[obj] = tag
}

func (s *Cache) DeleteObj(obj string) {
	s.lock.Lock()
	defer s.lock.Unlock()

	delete(s.objs, obj)
}

func (s *Cache) Replace(dstObj, srcObj string, tag Tag) {
	s.lock.Lock()
	defer s.lock.Unlock()

	delete(s.objs, srcObj)
	s.objs[dstObj] = tag
}
