package mm

import "regexp"

const (
	BrandXiaoMi  = "Xiaomi"
	BrandVivo    = "vivo"
	BrandSamsung = "Samsung"
	BrandRealme  = "Realme"
	BrandOppo    = "oppo"
	BrandOnePlus = "OnePlus"
	BrandMeiZu   = "MEIZU"
	BrandHuaWei  = "HUAWEI"
	BrandApple   = "Apple"
	Brand360     = "360"
	BrandNubia   = "nubia"
	BrandLenovo  = "Lenovo"
	BrandZTE     = "ZTE"
	BrandASUS    = "ASUS"
	BrandHONOR   = "HONOR"
	BrandLetv    = "Letv"
	BrandUnknown = "unknown"
)

// The rules are quoted from https://github.com/KHwang9883/MobileModels/blob/master/misc/naming-rules.md
var (
	ModelRegexpXiaoMi1 = regexp.MustCompile(`(^|\s|/)M(\d{4})(\w{2,4})[AECGHYIPLRTW]($|\s)`)
	ModelRegexpXiaoMi2 = regexp.MustCompile(`(^|\s|/)(\d{4})(\w{3,4})[CGYIPHLRT]($|\s)`)
	ModelRegexpXiaoMi3 = regexp.MustCompile(`(^|\s|/)(\d{4})(\w{5})[CAGOXYIPHLRT]($|\s)`)
	ModelRegexpXiaoMi4 = regexp.MustCompile(`(^|\s|/)(MI|REDMI|XIAOMI) \w+.*`)
	ModelRegexpXiaoMi5 = regexp.MustCompile(`(^|\s|/)[A-Z]{3}-(A0|H0)($|\s)`) // 黑鲨

	ModelRegexpViVo = regexp.MustCompile(`(^|\s|/)([VI])(\d{4})[A-Z]+([AT])?($|\s)`)

	ModelRegexpSamsung = regexp.MustCompile(`(^|\s|/)SM-([SGNFAMECJTXR])(\d{3})([0-9BEFGHUVAPTWNQMC]{1,2})($|\s)`)

	ModelRegexpRealme = regexp.MustCompile(`(^|\s|/)(RMX|RMP|RMW)(\d{4})($|\s)`)

	ModelRegexpOppo1 = regexp.MustCompile(`(^|\s|/)P([A-Z]{2})(M|T|1)(\d{2})($|\s)`)
	ModelRegexpOppo2 = regexp.MustCompile(`(^|\s|/)OPD(\d{4})($|\s)`) // oppo pad

	ModelRegexpOnePlus = regexp.MustCompile(`(^|\s|/)([A-Z]{2})(\d{4})($|\s)`)

	ModelRegexpMeiZu = regexp.MustCompile(`(^|\s|/)([MYUSL])(\d{3})([QMYUCDAH]?)($|\s)`)

	ModelRegexpHuaWei1 = regexp.MustCompile(`(^|\s|/)(\w{2,4})-([A-Z]{1,2})(\d{2})([A-Z]?)`)
	ModelRegexpHuaWei2 = regexp.MustCompile(`(^|\s|/)(P|MATE|NOVA)\d{1,3}(-\w{2,3})?($|\s)`)

	ModelRegexpApple1 = regexp.MustCompile(`(^|\s|/)A(\d{4})($|\s)`)
	ModelRegexpApple2 = regexp.MustCompile(`(^|\s|/)(IPAD|IPHONE)(\d{1,2})`)

	ModelRegexp360 = regexp.MustCompile(`(^|\s|/)(\d{4})([AM])(\d{2})($|\s)`)

	ModelRegexpNubia = regexp.MustCompile(`(^|\s|/)NX(\d{3})[JH]($|\s)`)
)

var BrandRegexps = map[string][]*regexp.Regexp{
	BrandXiaoMi:  {ModelRegexpXiaoMi3, ModelRegexpXiaoMi4, ModelRegexpXiaoMi2, ModelRegexpXiaoMi1, ModelRegexpXiaoMi5},
	BrandHuaWei:  {ModelRegexpHuaWei1, ModelRegexpHuaWei2},
	BrandVivo:    {ModelRegexpViVo},
	BrandOppo:    {ModelRegexpOppo1, ModelRegexpOppo2},
	BrandApple:   {ModelRegexpApple1, ModelRegexpApple2},
	BrandRealme:  {ModelRegexpRealme},
	BrandMeiZu:   {ModelRegexpMeiZu},
	BrandOnePlus: {ModelRegexpOnePlus},
	BrandSamsung: {ModelRegexpSamsung},
	Brand360:     {ModelRegexp360},
	BrandNubia:   {ModelRegexpNubia},
	BrandLenovo:  {},
	BrandZTE:     {},
	BrandASUS:    {},
	BrandHONOR:   {},
	BrandLetv:    {},
}
