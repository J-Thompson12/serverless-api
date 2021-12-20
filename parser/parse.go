package parser

import "strings"

type IPData struct {
	Name      string `json:"Name:"`
	City      string `json:"City"`
	StateProv string `json:"StateProv"`
	Email     string `json:"Email"`
}

func CreateMapData(data string) map[string]string {
	result := make(map[string]string)
	lines := strings.Split(data, "\n")
	for _, line := range lines {
		if !strings.Contains(line, "#") && strings.Contains(line, ":") {
			lineArray := strings.Split(line, ":")
			key := strings.TrimSpace(lineArray[0])
			value := strings.TrimSpace(lineArray[1])
			result[key] = value
		}
	}

	return result
}

func BuildIPData(data map[string]string) IPData {
	if isDomain(data) {
		return IPData{
			Name:      data["Domain Name"],
			City:      data["Registrant City"],
			StateProv: data["Admin State/Province"],
			Email:     data["Admin Email"],
		}
	}
	return IPData{
		Name:      data["OrgName"],
		City:      data["City"],
		StateProv: data["StateProv"],
		Email:     data["OrgAbuseEmail"],
	}
}

func isDomain(data map[string]string) bool {
	if _, ok := data["Domain Name"]; ok {
		return true
	}
	return false
}
