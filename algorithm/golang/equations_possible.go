package main

func equationsPossible(equations []string) bool {
	m := make(map[byte]byte)
	for _, expr := range(equations) {
		if expr[1] == '=' {
            s := _get(m, expr[0])
            m[_get(m, expr[0])] = s
			m[_get(m, expr[3])] = s
		}
	}
	for _, expr := range(equations) {
		if expr[1] == '!' {
			if _get(m, expr[0]) == _get(m, expr[3]) {
				return false
			}
		}
	}
	return true
}

func _get(m map[byte]byte, x byte) byte {
	_, exist := m[x]
	if !exist {
		return x;
	}
	for x != m[x] {
		x = m[x]
	}
	return x
}
