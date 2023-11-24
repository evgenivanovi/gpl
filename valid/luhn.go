package valid

// Luhn is an implementation of Luhn algorithm.
func Luhn(value string) error {

	if len(value) == 0 {
		return nil
	}

	var sum int
	var alter bool

	for ind := len(value) - 1; ind >= 0; ind-- {
		d := int(value[ind] - '0')
		if alter {
			d = d * 2
		}
		sum += d / 10
		sum += d % 10
		alter = !alter
	}

	if sum%10 != 0 {
		return ErrInvalidChecksum
	}

	return nil

}
