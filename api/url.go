package api

import "time"

func ComposeURL(host string, urlType string, school string, date any, extension any) string {
	var dateStr string
	switch v := date.(type) {
	case time.Time:
		dateStr = v.Format("20060102")
	case string:
		dateStr = v
	default:
		dateStr = ""
	}

	var extensionStr string
	if _, ok := extension.(string); !ok {
		extensionStr = ""
	} else {
		extensionStr = extension.(string)
	}

	return "https://" + host + "/" + school + urlType + dateStr + extensionStr

}
