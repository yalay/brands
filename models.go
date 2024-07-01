package main

import "sync"

var cache sync.Map

func Model2Brand(model string) string {
	if cachedBrand, ok := cache.Load(model); ok {
		return cachedBrand.(string)
	}

	for brand, regs := range BrandRegexps {
		for _, reg := range regs {
			if reg.MatchString(model) {
				cache.Store(model, brand)
				return brand
			}
		}
	}

	cache.Store(model, BrandUnknown)
	return BrandUnknown
}
