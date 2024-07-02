package mm

import (
	"testing"
)

func TestXiaoMiModel2Brand(t *testing.T) {
	models := []string{
		"M1803E1A", "M1903F10G", "M2012K11I",
		"2206123SC", "22081212UG", "23127PN0CG",
		"M2006C3LC", "2404ARN45A", "23106RN0DA",
		"21091116AC", "2201116SG", "2312DRAABC",
	}

	for _, model := range models {
		brand := Model2Brand(model)
		if brand != BrandXiaoMi {
			t.Errorf("%s expect mi. but got %s", model, brand)
		}
	}
}

func TestViVoModel2Brand(t *testing.T) {
	models := []string{
		"V1821A", "V1836T", "V1938A",
		"V2031EA", "V1901T", "V1801A0",
		"V1730EA", "V1824BA", "V1936AL",
		"PA2170", "iPA2375",
	}

	for _, model := range models {
		brand := Model2Brand(model)
		if brand != BrandVivo {
			t.Errorf("%s expect vivo. but got %s", model, brand)
		}
	}
}

func TestOppoModel2Brand(t *testing.T) {
	models := []string{
		"PAFM00", "PAFT10", "PDET10",
		"PGFM10", "PCAM00", "PCKT80",
		"PJW110", "PDHM00", "PBDM00",
		"PCDM00", "PJT110", "OPD2301",
	}

	for _, model := range models {
		brand := Model2Brand(model)
		if brand != BrandOppo {
			t.Errorf("%s expect oppo. but got %s", model, brand)
		}
	}
}

func TestHuaWeiModel2Brand(t *testing.T) {
	models := []string{
		"HUAWEI MT1-T00", "NXT-AL10", "BLA-AL00",
		"NOH-AN01", "ALT-AL10", "EVA-AL00",
		"MNA-AL00", "HBP-AL00", "LEM-AL00",
		"WAS-AL00", "SEA-TL10", "ADA-AL10U",
		"RNE-AL00", "TNN-AN00", "DIG-AL00",
		"FIG-AL00", "LDN-AL00", "DRA-AL00",
		"MED-AL00", "GFY-AL00", "MON-W19",
		"BAH-AL00", "SCM-AL09", "PCE-AL40",
		"BZK-W00",
	}

	for _, model := range models {
		brand := Model2Brand(model)
		if brand != BrandHuaWei {
			t.Errorf("%s expect huawei. but got %s", model, brand)
		}
	}
}

func TestSamsungModel2Brand(t *testing.T) {
	models := []string{
		"GT-I9000", "SCH-i909", "GT-I9100G",
		"SCH-I939D", "SCH-I939I", "GT-I9502",
		"GT-I9507V", "SM-G9006V", "SM-G9810",
		"SM-S9210", "GT-N7108D", "SM-F9000",
		"SM-J3119S", "SM-W7023", "SM-T819C",
		"SM-R9650",
	}

	for _, model := range models {
		brand := Model2Brand(model)
		if brand != BrandSamsung {
			t.Errorf("%s expect samsung. but got %s", model, brand)
		}
	}
}
