package myvin

import (
	"fmt"
	"strings"
)

//根据VIN8及年份代码等生成VIN码
func CreateVinFromYearCode(vin8, yearCode, factoryCode string, vinLast6 int) (vin string, err error) {
	for _, vin9 := range strings.Split(Vin9strs, ",") {
		vin := vin8 + vin9 + yearCode + factoryCode + GetSerialNumber6(vinLast6)
		if IsVinLegal(vin) {
			return vin, nil
		}
	}
	return "", fmt.Errorf("has illegal character in vin")
}

//根据VIN8及年份等生成VIN码
func CreateVinFromYear(vin8, year, factoryCode string, vinLast6 int) (vin string, err error) {
	for _, vin9 := range strings.Split(Vin9strs, ",") {
		vin := vin8 + vin9 + GetYearCodeFromYear(year) + factoryCode + GetSerialNumber6(vinLast6)
		if IsVinLegal(vin) {
			return vin, nil
		}
	}
	return "", fmt.Errorf("has illegal character in vin")
}

//根据区间生成多个VIN
func CreateVins(vin8, yearCode, factoryCode string, vinLast6Begin, vinLast6End int) (ret []string) {
	if vinLast6Begin < 1 {
		vinLast6Begin = 1
	}
	if vinLast6End > 999999 {
		vinLast6End = 999999
	}
	for vinLast6 := vinLast6Begin; vinLast6 <= vinLast6End; vinLast6++ {
		if vin, err := CreateVinFromYearCode(vin8, yearCode, factoryCode, vinLast6); err == nil {
			ret = append(ret, vin)
		}
	}
	return ret
}
