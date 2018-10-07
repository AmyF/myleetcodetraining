package kata

func BlackOrWhiteKey(keyPressCount int) string {
	// your code here
	index := (keyPressCount - 1) % 88 % 12 % 5
	switch index {
	case 0:
		return "white"
	case 1:
		return "black"
	case 2:
		return "white"
	case 3:
		return "white"
	case 4:
		return "black"
	default:
		return ""
	}
}
