package modbus

func isAvailable(data []string, str string) bool {
	for _, x := range data {
		if x == str {
			return true
		}
	}
	return false
}

type Hook interface {
	// Init(...any)
	Run(...any)
}

func ValidHooks() (ret []string) {
	ret = []string{"beforeRead", "afterRead", "beforeSend", "afterSend"}
	return
}
