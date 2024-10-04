package smileface

var validEye = map[string]bool{
	":": true,
	"8": true,
	";": true,
}

var validNose = map[string]bool{
	"-": true,
	"J": true,
	"~": true,
}

var validMouth = map[string]bool{
	")": true,
	"D": true,
}

func CountSmile(input []string) int {

	count := 0
	for _, v := range input {
		_, isValid := validEye[string(v[0])]
		if !isValid {
			continue
		}

		if len(v) == 3 {
			_, isValid = validNose[string(v[1])]
			if !isValid {
				continue
			}
			_, isValid = validMouth[string(v[2])]
			if !isValid {
				continue
			}
		} else {
			_, isValid = validMouth[string(v[1])]
			if !isValid {
				continue
			}
		}

		count++

	}

	return count
}
