package regexp

import (
	"regexp"
	"time"
)

type regexpStrRetSliceStrCache struct {
	*cache
}

func newRegexpStrRetSliceStrCache(ttl time.Duration, isEnabled bool) *regexpStrRetSliceStrCache {
	return &regexpStrRetSliceStrCache{
		cache: newCache(
			ttl,
			isEnabled,
		),
	}
}

func (c *regexpStrRetSliceStrCache) do(r *regexp.Regexp, s string, noCacheFn func(s string) []string) []string {
	// return if cache is not enabled
	if !c.enabled() {
		return noCacheFn(s)
	}

	// generate key, check key size
	key := r.String() + s
	if len(key) > maxKeySize {
		return noCacheFn(s)
	}

	// cache hit
	if res, found := c.getStrSlice(key); found {
		return res
	}

	// cache miss, add to cache if value is not too big
	res := noCacheFn(s)
	if len(res) <= maxValueSize {
		c.add(key, res)
	}

	return res
}
