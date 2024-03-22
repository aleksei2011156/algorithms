package main

func ValidParentheses(s string) bool {
	not_c := string(s[0]) != ")" || string(s[0]) != "]" || string(s[0]) != "}"
	not_o := string(s[len(s)-1]) != "(" || string(s[len(s)-1]) != "[" || string(s[len(s)-1]) != "{"
	if len(s)%2 != 0 || !not_c || !not_o {
		return false
	}

	_open := map[string]string{
		"(": ")",
		"[": "]",
		"{": "}",
	}
	_close := map[string]string{
		"}": "{",
		")": "(",
		"]": "[",
	}

	new_val := ""

	for _, sign := range s {
		if _, ok := _open[string(sign)]; ok {
			new_val += string(sign)
		} else {
			if new_val != "" && _close[string(sign)] == string(new_val[len(new_val)-1]) {
				new_val = new_val[:len(new_val)-1]
			} else {
				return false
			}
		}
	}

	if new_val != "" {
		return false
	}
	return true
}
