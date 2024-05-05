package util

func Query(payload map[string]string) string {
	queryString := ""

	for key, value := range payload {
		queryString += key + "=" + value + "&"
	}

	return queryString
}
