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
	BrandUnknown = "unknown"
)

// The rules are quoted from https://github.com/KHwang9883/MobileModels/blob/master/misc/naming-rules.md
var (
	ModelRegexpXiaoMi1 = regexp.MustCompile(`M(\d{4})(\w{3})[AECGHYIPLRTW]`)
	ModelRegexpXiaoMi2 = regexp.MustCompile(`(\d{4})(\w{3,4})[CGYIPHLRT]`)
	ModelRegexpXiaoMi3 = regexp.MustCompile(`(\d{4})(\w{5})[CAGOXYIPHLRT]`)
	ModelRegexpXiaoMi4 = regexp.MustCompile(`(MI|REDMI|XIAOMI) \w+.*`)
	ModelRegexpXiaoMi5 = regexp.MustCompile(`[A-Z]{3}-(A0|H0)`) // 黑鲨

	ModelRegexpViVo1 = regexp.MustCompile(`([VI])(\d{4})([AT])?`)
	ModelRegexpViVo2 = regexp.MustCompile(`VIVO.*`)

	ModelRegexpSamsung = regexp.MustCompile(`SM-([SGNFAMECJTXR])(\d{3})([0-9BEFGHUVAPTWNQMC]{1,2})`)

	ModelRegexpRealme = regexp.MustCompile(`(RMX|RMP|RMW)(\d{4})`)

	ModelRegexpOppo1 = regexp.MustCompile(`P([A-Z]{2})(M|T|1)(\d{2})`)
	ModelRegexpOppo2 = regexp.MustCompile(`OPPO \w+.*`)

	ModelRegexpOnePlus = regexp.MustCompile(`([A-Z]{2})(\d{4})`)

	ModelRegexpMeiZu = regexp.MustCompile(`([MYUSL])(\d{3})([QMYUCDAH]?)`)

	ModelRegexpHuaWei = regexp.MustCompile(`(\w{2,4})-([A-Z]{1,2})(\d{2})([A-Z]?)`)

	ModelRegexpApple = regexp.MustCompile(`^A(\d{4})$`)

	ModelRegexp360 = regexp.MustCompile(`^(\d{4})([AM])(\d{2})$`)

	ModelRegexpNubia = regexp.MustCompile(`^NX(\d{3})[JH]`)
)

var BrandRegexps = map[string][]*regexp.Regexp{
	BrandXiaoMi:  {ModelRegexpXiaoMi3, ModelRegexpXiaoMi4, ModelRegexpXiaoMi2, ModelRegexpXiaoMi1, ModelRegexpXiaoMi5},
	BrandHuaWei:  {ModelRegexpHuaWei},
	BrandVivo:    {ModelRegexpViVo1, ModelRegexpViVo2},
	BrandOppo:    {ModelRegexpOppo1, ModelRegexpOppo2},
	BrandApple:   {ModelRegexpApple},
	BrandRealme:  {ModelRegexpRealme},
	BrandMeiZu:   {ModelRegexpMeiZu},
	BrandOnePlus: {ModelRegexpOnePlus},
	BrandSamsung: {ModelRegexpSamsung},
	Brand360:     {ModelRegexp360},
	BrandNubia:   {ModelRegexpNubia},
}
