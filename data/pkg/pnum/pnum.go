package pnum

import (
	"strconv"
	"strings"
)

func ToInt(persianNum string) (int, error) {
	if persianNum == "" {
		return 0, nil
	}

	var numBuilder strings.Builder

	segments := strings.Split(persianNum, "")

	for _, s := range segments {
		switch s {
		case "۰":
			numBuilder.WriteString("0")
		case "۱":
			numBuilder.WriteString("1")
		case "۲":
			numBuilder.WriteString("2")
		case "۳":
			numBuilder.WriteString("3")
		case "۴":
			numBuilder.WriteString("4")
		case "۵":
			numBuilder.WriteString("5")
		case "۶":
			numBuilder.WriteString("6")
		case "۷":
			numBuilder.WriteString("7")
		case "۸":
			numBuilder.WriteString("8")
		case "۹":
			numBuilder.WriteString("9")
		default:
			continue
		}
	}

	num, err := strconv.Atoi(numBuilder.String())
	if err != nil {
		return 0, err
	}

	return num, nil
}
