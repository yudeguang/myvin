package myvin

import (
	"github.com/yudeguang/stringsx"
	"strings"
)

type country struct {
	From   string
	To     string
	NameEn string
	NameZh string
}

var countries = []country{
	{
		From:   "AA",
		To:     "AH",
		NameEn: "South Africa",
		NameZh: "南非",
	},
	{
		From:   "AJ",
		To:     "AN",
		NameEn: "Ivory Coast",
		NameZh: "科特迪瓦",
	},
	{
		From:   "BA",
		To:     "BE",
		NameEn: "Angola",
		NameZh: "安哥拉",
	},
	{
		From:   "BF",
		To:     "BK",
		NameEn: "Kenya",
		NameZh: "肯尼亚",
	},
	{
		From:   "BL",
		To:     "BR",
		NameEn: "Tanzania",
		NameZh: "坦桑尼亚",
	},
	{
		From:   "CA",
		To:     "CE",
		NameEn: "Benin",
		NameZh: "贝宁",
	},
	{
		From:   "CF",
		To:     "CK",
		NameEn: "Madagascar",
		NameZh: "马达加斯加",
	},
	{
		From:   "CL",
		To:     "CR",
		NameEn: "Tunisia",
		NameZh: "突尼斯",
	},
	{
		From:   "DA",
		To:     "DE",
		NameEn: "Egypt",
		NameZh: "埃及",
	},
	{
		From:   "DF",
		To:     "DK",
		NameEn: "Morocco",
		NameZh: "摩洛哥 ",
	},
	{
		From:   "DL",
		To:     "DR",
		NameEn: "Zambia",
		NameZh: "赞比亚",
	},
	{
		From:   "EA",
		To:     "EE",
		NameEn: "Ethiopia",
		NameZh: "埃塞俄比亚",
	},
	{
		From:   "EF",
		To:     "EK",
		NameEn: "Mozambique",
		NameZh: "莫桑比克",
	},
	{
		From:   "FA",
		To:     "FE",
		NameEn: "Ghana",
		NameZh: "加纳",
	},
	{
		From:   "FF",
		To:     "FK",
		NameEn: "Nigeria",
		NameZh: "尼日利亚",
	},
	{
		From:   "HA",
		To:     "HZ",
		NameEn: "China",
		NameZh: "中国",
	},

	{
		From:   "JA",
		To:     "J0",
		NameEn: "Japan",
		NameZh: "日本",
	},
	{
		From:   "KA",
		To:     "KE",
		NameEn: "Sri Lanka",
		NameZh: "斯里兰卡",
	},
	{
		From:   "KF",
		To:     "KK",
		NameEn: "Israel",
		NameZh: "以色列",
	},
	{
		From:   "KL",
		To:     "KR",
		NameEn: "Korea (South)",
		NameZh: "朝鲜(韩国)",
	},
	{
		From:   "KS",
		To:     "K0",
		NameEn: "Kazakhstan",
		NameZh: "哈萨克斯坦",
	},
	{
		From:   "LA",
		To:     "L0",
		NameEn: "China",
		NameZh: "中国",
	},
	{
		From:   "MA",
		To:     "ME",
		NameEn: "India",
		NameZh: "印度",
	},
	{
		From:   "MF",
		To:     "MK",
		NameEn: "Indonesia",
		NameZh: "印度尼西亚 ",
	},
	{
		From:   "ML",
		To:     "MR",
		NameEn: "Thailand",
		NameZh: "泰国",
	},
	{
		From:   "NA",
		To:     "NE",
		NameEn: "Iran",
		NameZh: "伊朗",
	},
	{
		From:   "NF",
		To:     "NK",
		NameEn: "Pakistan",
		NameZh: "巴基斯坦",
	},
	{
		From:   "NL",
		To:     "NR",
		NameEn: "Turkey",
		NameZh: "土耳其",
	},
	{
		From:   "PA",
		To:     "PE",
		NameEn: "Philippines",
		NameZh: "菲律宾",
	},
	{
		From:   "PF",
		To:     "PK",
		NameEn: "Singapore",
		NameZh: "新加坡",
	},
	{
		From:   "PL",
		To:     "PR",
		NameEn: "Malaysia",
		NameZh: "马来西亚",
	},
	{
		From:   "RA",
		To:     "RE",
		NameEn: "United Arab Emirates",
		NameZh: "阿联酋",
	},
	{
		From:   "RF",
		To:     "RK",
		NameEn: "Taiwan China",
		NameZh: "中国台湾",
	},
	{
		From:   "RL",
		To:     "RR",
		NameEn: "Vietnam",
		NameZh: "越南",
	},
	{
		From:   "RS",
		To:     "R0",
		NameEn: "Saudi Arabia",
		NameZh: "沙特阿拉伯",
	},
	{
		From:   "SA",
		To:     "SM",
		NameEn: "United Kingdom",
		NameZh: "英国",
	},
	{
		From:   "SN",
		To:     "ST",
		NameEn: "Germany",
		NameZh: "德国",
	},
	{
		From:   "SU",
		To:     "SZ",
		NameEn: "Poland",
		NameZh: "波兰",
	},
	{
		From:   "S1",
		To:     "S4",
		NameEn: "Latvia",
		NameZh: "拉脱维亚",
	},
	{
		From:   "TA",
		To:     "TH",
		NameEn: "Switzerland",
		NameZh: "瑞士",
	},
	{
		From:   "TJ",
		To:     "TP",
		NameEn: "Czech Republic",
		NameZh: "捷克",
	},
	{
		From:   "TR",
		To:     "TV",
		NameEn: "Hungary",
		NameZh: "匈牙利",
	},
	{
		From:   "TW",
		To:     "T1",
		NameEn: "Portugal",
		NameZh: "葡萄牙",
	},
	{
		From:   "UH",
		To:     "UM",
		NameEn: "Denmark",
		NameZh: "丹麦",
	},
	{
		From:   "UN",
		To:     "UT",
		NameEn: "Ireland",
		NameZh: "爱尔兰",
	},
	{
		From:   "UU",
		To:     "UZ",
		NameEn: "Romania",
		NameZh: "罗马尼亚",
	},
	{
		From:   "U5",
		To:     "U7",
		NameEn: "Slovakia",
		NameZh: "斯洛伐克",
	},
	{
		From:   "VA",
		To:     "VE",
		NameEn: "Austria",
		NameZh: "奥地利",
	},
	{
		From:   "VF",
		To:     "VR",
		NameEn: "France",
		NameZh: "法国",
	},
	{
		From:   "VS",
		To:     "VW",
		NameEn: "Spain",
		NameZh: "西班牙",
	},
	{
		From:   "VX",
		To:     "V2",
		NameEn: "Serbia",
		NameZh: "塞尔维亚",
	},
	{
		From:   "V3",
		To:     "V5",
		NameEn: "Croatia",
		NameZh: "克罗地亚",
	},
	{
		From:   "V6",
		To:     "V0",
		NameEn: "Estonia",
		NameZh: "爱沙尼亚",
	},
	{
		From:   "WA",
		To:     "W0",
		NameEn: "Germany",
		NameZh: "德国",
	},
	{
		From:   "XA",
		To:     "XE",
		NameEn: "Bulgaria",
		NameZh: "保加利亚",
	},
	{
		From:   "XF",
		To:     "XK",
		NameEn: "Greece",
		NameZh: "希腊",
	},
	{
		From:   "XL",
		To:     "XR",
		NameEn: "Netherlands",
		NameZh: "荷兰",
	},
	{
		From:   "XS",
		To:     "XW",
		NameEn: "Russia",
		NameZh: "俄罗斯",
	},
	{
		From:   "XX",
		To:     "X2",
		NameEn: "Luxembourg",
		NameZh: "卢森堡",
	},
	{
		From:   "X3",
		To:     "X0",
		NameEn: "Russia",
		NameZh: "俄罗斯",
	},
	{
		From:   "YA",
		To:     "YE",
		NameEn: "Belgium",
		NameZh: "比利时",
	},
	{
		From:   "YF",
		To:     "YK",
		NameEn: "Finland",
		NameZh: "芬兰",
	},
	{
		From:   "YL",
		To:     "YR",
		NameEn: "Malta",
		NameZh: "马耳他",
	},
	{
		From:   "YS",
		To:     "YW",
		NameEn: "Sweden",
		NameZh: "瑞典",
	},
	{
		From:   "YX",
		To:     "Y2",
		NameEn: "Norway",
		NameZh: "挪威",
	},
	{
		From:   "Y3",
		To:     "Y5",
		NameEn: "Belarus",
		NameZh: "白俄罗斯",
	},
	{
		From:   "Y6",
		To:     "Y0",
		NameEn: "Ukraine",
		NameZh: "乌克兰",
	},
	{
		From:   "ZA",
		To:     "ZR",
		NameEn: "Italy",
		NameZh: "意大利",
	},
	{
		From:   "ZX",
		To:     "Z2",
		NameEn: "Slovenia",
		NameZh: "斯洛文尼亚",
	},
	{
		From:   "Z3",
		To:     "Z5",
		NameEn: "Lithuania",
		NameZh: "立陶宛",
	},
	{
		From:   "Z6",
		To:     "Z9",
		NameEn: "Russia",
		NameZh: "俄罗斯",
	},
	{
		From:   "1A",
		To:     "10",
		NameEn: "United States",
		NameZh: "美国",
	},
	{
		From:   "2A",
		To:     "20",
		NameEn: "Canada",
		NameZh: "加拿大",
	},
	{
		From:   "3A",
		To:     "37",
		NameEn: "Mexico",
		NameZh: "墨西哥",
	},
	{
		From:   "38",
		To:     "30",
		NameEn: "Cayman Islands",
		NameZh: "开曼群岛",
	},
	{
		From:   "4A",
		To:     "40",
		NameEn: "United States",
		NameZh: "美国",
	},
	{
		From:   "5A",
		To:     "50",
		NameEn: "United States",
		NameZh: "美国",
	},
	{
		From:   "6A",
		To:     "6W",
		NameEn: "Australia",
		NameZh: "澳大利亚",
	},
	{
		From:   "7A",
		To:     "7E",
		NameEn: "New Zealand",
		NameZh: "新西兰",
	},
	{
		From:   "8A",
		To:     "8E",
		NameEn: "Argentina",
		NameZh: "阿根廷",
	},
	{
		From:   "8F",
		To:     "8K",
		NameEn: "Chile",
		NameZh: "智利",
	},
	{
		From:   "8L",
		To:     "8R",
		NameEn: "Ecuador",
		NameZh: "厄瓜多尔",
	},
	{
		From:   "8S",
		To:     "8W",
		NameEn: "Peru",
		NameZh: "秘鲁",
	},
	{
		From:   "8X",
		To:     "82",
		NameEn: "Venezuela",
		NameZh: "委内瑞拉",
	},
	{
		From:   "9A",
		To:     "9E",
		NameEn: "Brazil",
		NameZh: "巴西",
	},
	{
		From:   "9F",
		To:     "9K",
		NameEn: "Colombia",
		NameZh: "哥伦比亚",
	},
	{
		From:   "9L",
		To:     "9R",
		NameEn: "Paraguay",
		NameZh: "巴拉圭",
	},
	{
		From:   "9S",
		To:     "9W",
		NameEn: "Uruguay",
		NameZh: "乌拉圭",
	},
	{
		From:   "9X",
		To:     "92",
		NameEn: "Trinidad & Tobago",
		NameZh: "特立尼达和多巴哥",
	},
	{
		From:   "93",
		To:     "99",
		NameEn: "Brazil",
		NameZh: "巴西",
	},
}

//获取生产地区英文名
func GetRegionEn(vin string) string {
	region := strings.ToUpper(stringsx.Left(vin, 1))
	if region >= "A" && region <= "H" {
		return "Africa"
	} else if region >= "J" && region <= "R" {
		return "Asia"
	} else if region >= "S" && region <= "Z" {
		return "Europe"
	} else if region >= "1" && region <= "5" {
		return "North America"
	} else if region >= "6" && region <= "7" {
		return "Oceania"
	} else if region >= "8" && region <= "9" {
		return "South America"
	} else {
		return "Unknown Region"
	}
}

//获取生产地区中文名
func GetRegionZh(vin string) string {
	region := strings.ToUpper(stringsx.Left(vin, 1))
	if region >= "A" && region <= "H" {
		return "非洲"
	} else if region >= "J" && region <= "R" {
		return "亚洲"
	} else if region >= "S" && region <= "Z" {
		return "欧洲"
	} else if region >= "1" && region <= "5" {
		return "北美洲"
	} else if region >= "6" && region <= "7" {
		return "大洋洲"
	} else if region >= "8" && region <= "9" {
		return "南美洲"
	} else {
		return "未知区域"
	}
}

//获取生产国家英文名
func GetCountryEn(vin string) string {
	qi := indexOf(strings.ToUpper(stringsx.Left(vin, 2)))
	for _, country := range countries {
		i := indexOf(country.From)
		j := indexOf(country.To)
		if qi >= i && qi <= j {
			return country.NameEn
		}
	}
	return "Unknown country"
}

//获取生产国家中文名
func GetCountryZh(vin string) string {
	qi := indexOf(strings.ToUpper(stringsx.Left(vin, 2)))
	for _, country := range countries {
		i := indexOf(country.From)
		j := indexOf(country.To)
		if qi >= i && qi <= j {
			return country.NameZh
		}
	}
	return "未知国家"
}

const chars = "ABCDEFGHIJKLMNOPRSTUVWXYZ1234567890"

func indexOf(s string) int {
	return strings.IndexByte(chars, s[0])*len(chars) + strings.IndexByte(chars, s[1])
}
