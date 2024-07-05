package brands

import (
	"regexp"
	"strings"

	"github.com/golang/groupcache/lru"
)

var cache = lru.New(10000)
var upperBrands = make(map[string]string, len(brandRegexps))

func init() {
	for brand, _ := range brandRegexps {
		upperBrand := strings.ToUpper(brand)
		upperBrands[upperBrand] = brand
	}
}

func Model2Brand(model string) string {
	upperModel := strings.ToUpper(model)
	if cachedBrand, ok := cache.Get(upperModel); ok {
		return cachedBrand.(string)
	}

	// first matching with brand keywords
	for brand, _ := range brandRegexps {
		upperBrand := strings.ToUpper(brand)
		if ok, _ := regexp.MatchString(`(^|\s|/)`+upperBrand, upperModel); ok {
			cache.Add(upperModel, brand)
			return brand
		}
	}

	// second matching with regular rule
	for brand, regs := range brandRegexps {
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

func IsPopularBrand(brand string) (string, bool) {
	upperBrand := strings.ToUpper(brand)
	if standardBrand, ok := upperBrands[upperBrand]; ok {
		return standardBrand, true
	} else {
		return "", false
	}
}
