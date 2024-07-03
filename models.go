package mm

import (
	"regexp"
	"strings"

	"github.com/golang/groupcache/lru"
)

var cache = lru.New(10000)

func Model2Brand(model string) string {
	upperModel := strings.ToUpper(model)
	if cachedBrand, ok := cache.Get(upperModel); ok {
		return cachedBrand.(string)
	}

	// first matching with brand keywords
	for brand, _ := range BrandRegexps {
		upperBrand := strings.ToUpper(brand)
		if ok, _ := regexp.MatchString(`(^|\s|/)`+upperBrand, upperModel); ok {
			cache.Add(upperModel, brand)
			return brand
		}
	}

	// second matching with regular rule
	for brand, regs := range BrandRegexps {
		for _, reg := range regs {
			if reg.MatchString(upperModel) {
				cache.Add(upperModel, brand)
				return brand
			}
		}
	}

	cache.Add(upperModel, BrandUnknown)
	return BrandUnknown
}
