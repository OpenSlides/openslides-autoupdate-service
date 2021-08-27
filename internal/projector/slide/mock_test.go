package slide_test

func convertData(data map[string]string) map[string][]byte {
	converted := make(map[string][]byte, len(data))
	for k, v := range data {
		converted[k] = []byte(v)
	}
	return converted
}
