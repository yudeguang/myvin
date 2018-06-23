package myvin

import (
	"strconv"
)

//生成6位序列号
func GetSerialNumber6(n int) string {
	serial_number := strconv.Itoa(n)
	len_serial_number := len(serial_number)
	switch len_serial_number {
	case 1:
		return "00000" + serial_number
	case 2:
		return "0000" + serial_number
	case 3:
		return "000" + serial_number
	case 4:
		return "00" + serial_number
	case 5:
		return "0" + serial_number
	case 6:
		return serial_number
	default:
		return "000000"
	}
}

//生成5位序列号
func GetSerialNumber5(n int) string {
	serial_number := strconv.Itoa(n)
	len_serial_number := len(serial_number)
	switch len_serial_number {
	case 1:
		return "0000" + serial_number
	case 2:
		return "000" + serial_number
	case 3:
		return "00" + serial_number
	case 4:
		return "0" + serial_number
	case 5:
		return serial_number
	default:
		return "00000"
	}
}

//生成4位序列号
func GetSerialNumber4(n int) string {
	serial_number := strconv.Itoa(n)
	len_serial_number := len(serial_number)
	switch len_serial_number {
	case 1:
		return "000" + serial_number
	case 2:
		return "00" + serial_number
	case 3:
		return "0" + serial_number
	case 4:
		return serial_number
	default:
		return "0000"
	}
}

//生成3位序列号
func GetSerialNumber3(n int) string {
	serial_number := strconv.Itoa(n)
	len_serial_number := len(serial_number)
	switch len_serial_number {
	case 1:
		return "00" + serial_number
	case 2:
		return "0" + serial_number
	case 3:
		return serial_number
	default:
		return "000"
	}
}

//生成2位序列号
func GetSerialNumber2(n int) string {
	serial_number := strconv.Itoa(n)
	len_serial_number := len(serial_number)
	switch len_serial_number {
	case 1:
		return "0" + serial_number
	case 2:
		return serial_number
	default:
		return "00"
	}

}

//生成1位序列号
func GetSerialNumber1(n int) string {
	serial_number := strconv.Itoa(n)
	len_serial_number := len(serial_number)
	switch len_serial_number {
	case 1:
		return serial_number
	default:
		return "0"
	}

}
