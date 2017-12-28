package secret

var actions = []string{"wink", "double blink", "close your eyes", "jump"}

// Handshake generates a list of actions based on code
func Handshake(code uint) []string {
	res := []string{}
	for _, act := range actions {
		if code&1 == 1 {
			res = append(res, act)
		}
		code >>= 1
	}
	if code&1 == 1 {
		for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
			res[i], res[j] = res[j], res[i]
		}
	}
	return res
}
