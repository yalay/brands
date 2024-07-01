package mm

import (
	"strings"
	"sync"
)

var cache sync.Map

func Model2Brand(model string) string {
	upperModel := strings.ToUpper(model)
	if cachedBrand, ok := cache.Load(upperModel); ok {
		return cachedBrand.(string)
	}

	for brand, regs := range BrandRegexps {
		for _, reg := range regs {
			if reg.MatchString(upperModel) {
				cache.Store(upperModel, brand)
				return brand
			}
		}
	}

	cache.Store(upperModel, BrandUnknown)
	return BrandUnknown
}
