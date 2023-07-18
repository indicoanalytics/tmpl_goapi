package helpers

func Contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func MapToBytes(datamap map[string]interface{}) ([]byte, error) {
	databyte, err := Marshal(datamap)

	return databyte, err
}
