package validators

import (
	"fmt"
	"net"
	"regexp"
)

var (
	regexFQDN = regexp.MustCompile(`^([a-zA-Z0-9][a-zA-Z0-9-]{0,62})(\.[a-zA-Z0-9][a-zA-Z0-9-]{0,62})*?(\.[a-zA-Z][a-zA-Z0-9]{0,62})\.?$`)
)

func MaxLength(maxLen int) func(string) error {
	return func(str string) error {
		if len(str) > maxLen {
			return fmt.Errorf("string must be less than or equal to %d characters", maxLen)
		}
		return nil
	}
}

func MinLength(minLen int) func(string) error {
	return func(str string) error {
		if len(str) < minLen {
			return fmt.Errorf("string must be greater than or equal to %d characters", minLen)
		}
		return nil
	}
}

func RegexValidator(regex string) func(string) error {
	return func(str string) error {
		re := regexp.MustCompile(regex)
		if !re.MatchString(str) {
			return fmt.Errorf("contents must match the regex: %s", regex)
		}
		return nil
	}
}

func IPv4Validator() func(string) error {
	return func(str string) error {
		if !isIPv4(str) {
			return fmt.Errorf("%s is not a valid IPv4 address", str)
		}
		return nil
	}
}

func IPv6Validator() func(string) error {
	return func(s string) error {
		if !isIPv6(s) {
			return fmt.Errorf("%s is not a valid IPv6 address", s)
		}
		return nil
	}
}

func IPValidator() func(string) error {
	return func(str string) error {
		if isIP(str) == 0 {
			return fmt.Errorf("%s is not a valid IP address", str)
		}
		return nil
	}
}

func FQDNValidator() func(string) error {
	return func(str string) error {
		if !regexFQDN.MatchString(str) {
			return fmt.Errorf("%s is not a valid FQDN", str)
		}
		return nil
	}
}

func isIP(input string) int32 {
	ip := net.ParseIP(input)
	if ip == nil {
		return 0
	}

	if ip.To4() != nil {
		return 4
	}

	return 6
}

func isIPv4(input string) bool {
	return 4 == isIP(input)
}

func isIPv6(input string) bool {
	return 6 == isIP(input)
}
