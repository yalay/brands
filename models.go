package mm

import (
	"regexp"
	"strings"
	"sync"
)

var cache sync.Map

func Model2Brand(model string) string {
	upperModel := strings.ToUpper(model)
	if cachedBrand, ok := cache.Load(upperModel); ok {
		return cachedBrand.(string)
	}

	// first matching with brand keywords
	for brand, _ := range BrandRegexps {
		upperBrand := strings.ToUpper(brand)
		if ok, _ := regexp.MatchString(`(^|\s|/)`+upperBrand, upperModel); ok {
			cache.Store(upperModel, brand)
			return brand
		}
	}

	// second regular rule matching
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
