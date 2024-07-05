package brands

import "regexp"

const (
	BrandXiaoMi     = "Xiaomi"
	BrandRedMi      = "Redmi"
	BrandBlackSHARK = "SHARK"
	BrandCoolpad    = "coolpad"
	BrandGoogle     = "google"
	BrandVivo       = "vivo"
	BrandSamsung    = "Samsung"
	BrandRealme     = "realme"
	BrandOppo       = "OPPO"
	BrandOnePlus    = "OnePlus"
	BrandMeiZu      = "MEIZU"
	BrandHuaWei     = "HUAWEI"
	BrandHonor      = "HONOR"
	BrandApple      = "Apple"
	Brand360        = "360"
	BrandNubia      = "nubia"
	BrandNokia      = "nokia"
	BrandNothing    = "nothing"
	BrandLenovo     = "Lenovo"
	BrandZTE        = "ZTE"
	BrandASUS       = "ASUS"
	BrandLetv       = "Letv"
	BrandMotorola   = "Motorola"
	BrandIQOO       = "iQOO"
	BrandSmartisan  = "Smartisan"
	BrandSony       = "SONY"
	BrandUnknown    = "other"
)

// The rules are quoted from https://github.com/KHwang9883/MobileModels/blob/master/misc/naming-rules.md
var (
	modelRegexpXiaoMi1 = regexp.MustCompile(`(^|\s|/)M(\d{4})(\w{2,4})[AECGHYIPLRTW]($|\s)`)
	modelRegexpXiaoMi2 = regexp.MustCompile(`(^|\s|/)(\d{4})(\w{3,4})[CGYIPHLRT]($|\s)`)
	modelRegexpXiaoMi3 = regexp.MustCompile(`(^|\s|/)(\d{4})(\w{5})[CAGOXYIPHLRT]($|\s)`)
	modelRegexpXiaoMi4 = regexp.MustCompile(`(^|\s|/)(MI|REDMI|XIAOMI) \w+.*`)
	modelRegexpXiaoMi5 = regexp.MustCompile(`(^|\s|/)[A-Z]{3}-(A0|H0)($|\s)`) // 黑鲨

	modelRegexpViVo1 = regexp.MustCompile(`(^|\s|/)([VI])(\d{4})[A-Z]+([AT0])?($|\s)`)
	modelRegexpViVo2 = regexp.MustCompile(`(^|\s|/)(PA|IPA)(\d{4})($|\s)`)

	modelRegexpSamsung1 = regexp.MustCompile(`(^|\s|/)SM-([SGNFAMECJTXRW])(\d{3})(\w{1,2})($|\s)`)
	modelRegexpSamsung2 = regexp.MustCompile(`(^|\s|/)(GT|SCH)-[IN]\d{3,4}[A-Z]?($|\s)`)

	modelRegexpRealme = regexp.MustCompile(`(^|\s|/)(RMX|RMP|RMW)(\d{4})($|\s)`)

	modelRegexpOppo1 = regexp.MustCompile(`(^|\s|/)P([A-Z]{2})(M|T|1)(\d{2})($|\s)`)
	modelRegexpOppo2 = regexp.MustCompile(`(^|\s|/)(OPD|CPH)(\d{4})($|\s)`) // oppo pad

	modelRegexpMeiZu = regexp.MustCompile(`(^|\s|/)([MYUSL])(\d{3})([QMYUCDAH]?)($|\s)`)

	modelRegexpHuaWei1 = regexp.MustCompile(`(^|\s|/)(\w{2,4})-([A-Z]{1,2})(\d{2})([A-Z]?)($|\s)`)
	modelRegexpHuaWei2 = regexp.MustCompile(`(^|\s|/)(P|MATE|NOVA)\d{1,3}(-\w{2,3})?($|\s)`)

	modelRegexpApple1 = regexp.MustCompile(`(^|\s|/)A(\d{4})($|\s)`)
	modelRegexpApple2 = regexp.MustCompile(`(^|\s|/)(IPAD|IPHONE)(\d{1,2})`)

	modelRegexp360 = regexp.MustCompile(`(^|\s|/)(\d{4})([AM])(\d{2})($|\s)`)

	modelRegexpNubia = regexp.MustCompile(`(^|\s|/)NX(\d{3})[JH]($|\s)`)

	modelRegexpMotorola = regexp.MustCompile(`(^|\s|/)XT\d{4}-\d{1,2}($|\s)`)
)

var brandRegexps = map[string][]*regexp.Regexp{
	BrandXiaoMi:     {modelRegexpXiaoMi3, modelRegexpXiaoMi4, modelRegexpXiaoMi2, modelRegexpXiaoMi1, modelRegexpXiaoMi5},
	BrandRedMi:      {},
	BrandHuaWei:     {modelRegexpHuaWei1, modelRegexpHuaWei2},
	BrandHonor:      {},
	BrandVivo:       {modelRegexpViVo1, modelRegexpViVo2},
	BrandOppo:       {modelRegexpOppo1, modelRegexpOppo2},
	BrandApple:      {modelRegexpApple1, modelRegexpApple2},
	BrandRealme:     {modelRegexpRealme},
	BrandMeiZu:      {modelRegexpMeiZu},
	BrandOnePlus:    {},
	BrandSamsung:    {modelRegexpSamsung1, modelRegexpSamsung2},
	Brand360:        {modelRegexp360},
	BrandNubia:      {modelRegexpNubia},
	BrandMotorola:   {modelRegexpMotorola},
	BrandLenovo:     {},
	BrandZTE:        {},
	BrandASUS:       {},
	BrandLetv:       {},
	BrandIQOO:       {},
	BrandBlackSHARK: {},
	BrandCoolpad:    {},
	BrandGoogle:     {},
	BrandNokia:      {},
	BrandNothing:    {},
	BrandSmartisan:  {},
	BrandSony:       {},
}
