package functions

func FiltringCheck(temp string) (bool, int) {
	if temp[0] == 13 || temp[0] == 10 { // when the client want to send emty name
		return false, 1
	}
	for i := 0; i < len(temp); i++ {
		if temp[i]==' ' {
			return false, 3
		}
		if temp[i] < 32 || temp[i] > 126 {
			if temp[i] == 13 || temp[i] == 10 {
				continue // ignore '\n' and '\r\n'
			}
			return false, 2
		}
	}
	return true, 0
}
