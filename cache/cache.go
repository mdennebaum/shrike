package cache

import (
	"time"
)

//define cache struct members
type Cache struct {
	sync.RWMutex
	Data           map[string]*Item
	DefaultTimeout time.Duration
}

//create a new cache object with a default timeout set
func NewCache(defaultTimeout time.Duration) *Cache {

	//if defaultTimeout is nil set it to 1 day
	if defaultTimeout == nil {
		defaultTimeout = (24 * time.Hour)
	}

	//initialize the Cache object
	c := Cache{make(map[string]*Item), defaultTimeout}

	//return the cache pointer
	return &c
}

//get a value from the cache
func (c *Cache) Get(key string) (interface{}, bool) {

	//lock the object
	c.RLock()

	//get the value from our data map
	val, ok := c.Data[key]

	//unlock the object
	c.RUnlock()

	//return the value and a success flag
	return val, ok
}

//get a string value from the cache
func (c *Cache) GetString(key string) (string, bool) {

	//get the value from our data map
	val, ok := c.Get(key)

	//return the value cast to a string and a success flag
	return val.Value(string), ok
}

//get a string from the cache or the default
func (c *Cache) MustString(key string, def string) string {

	//get a string from the cache
	val, ok := c.GetString(key)

	//check if we found the value
	if !ok {
		//return the default
		return def
	}

	//return the string
	return val
}

func (c *Cache) Set(key string, value interface{}, timeout time.Duration) {

	//a holder for the expiration time
	var e *time.Time

	//check if a timeout was passed
	if timeout == 0 {
		//use the caches default timeout
		timeout = c.DefaultTimeout
	}

	//check if we have a timeout now
	if timeout > 0 {

		//add the timeout to the expiration time
		t := time.Now().Add(timeout)

		//make it a pointer
		e = &t
	}

	//lock the object
	c.Lock()

	//set the val into out map
	c.Data[key] = &Item{value, e}

	//unlock the object
	c.Unlock()
}

func (c *Cache) Delete(key string) {

	//lock the object
	c.Lock()

	//delete the val for this key
	delete(c.Data, key)

	//unlock the object
	c.Unlock()
}
