package cmdlines

import "os"

func ExpandArgs(cache map[string]string, args []string) []string {
	res := make([]string, 0, 32)
	for _, arg := range args {
		res = append(res, SplitShellVars(cache, arg)...)
	}
	return res
}

// SplitShellVars do not determinate some special shell variable names, like $$/$./$1/$?/$#
func SplitShellVars(cache map[string]string, s string) []string {
	res := make([]string, 0, 16)
	var last string
	for i := 0; i < len(s); {
		j := i
		for j < len(s) && s[j] != '$' {
			j++
		}
		if j == len(s) {
			last = s[i:j]
			break
		}
		shellVarName, n := GetValidShellVarName(s[j+1:])
		if n > 0 {
			value := GetShellVarValue(cache, shellVarName)
			if j > i {
				res = append(res, s[i:j])
			}
			res = append(res, value)
			i = j + n + 1
			continue
		}
		j = j + 1 // this '$' is a normal '$', not a shell var
		if s[j] == '$' {
			j = j + 1 // "$$" is normal
		}
		for j < len(s) && s[j] != '$' {
			j++
		}
		res = append(res, s[i:j])
		i = j
	}
	if last != "" {
		res = append(res, last)
	}
	return res
}

func GetShellVarValue(cache map[string]string, v string) string {
	if vv, ok := cache[v]; ok {
		return vv
	}
	vv, ok := os.LookupEnv(v)
	if ok {
		cache[v] = vv
	}
	return vv
}

func GetValidShellVarName(s string) (string, int) {
	if len(s) == 0 {
		return "", 0
	}
	i := 0
	leftBrace := false
	if s[i] == '{' {
		leftBrace = true
		i++
	}
	for i < len(s) && (s[i] == '_' || (s[i] >= '0' && s[i] <= '9') || (s[i] >= 'a' && s[i] <= 'z') || (s[i] >= 'A' && s[i] <= 'Z')) {
		i++
	}
	if leftBrace {
		if i == len(s) || s[i] != '}' {
			return "", 0
		}
		return s[1:i], i + 1
	}
	return s[:i], i
}
