package project

import (
	"github.com/patrickmn/go-cache"
)

// default cache
var c = cache.New(cache.NoExpiration, cache.NoExpiration)

func (p *Project) setcache() {
	key := p.getprojectpath()
	c.SetDefault(key, p)
}

func getcache(key string) (p *Project) {
	x, found := c.Get(key)
	if found {
		p = x.(*Project)
	}
	return
}

func (p *Project) delcache() {
	key := p.getprojectpath()
	c.Delete(key)
}
