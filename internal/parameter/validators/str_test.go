package validators

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxLength(t *testing.T) {

}

func TestMinLength(t *testing.T) {

}

func TestIPv4Validator(t *testing.T) {
	tests := []struct {
		ip       string
		expected bool
	}{
		{"fe80:0000:0000:0000:0204:61ff:fe9d:f156", false},      // full form of IPv6
		{"fe80:0:0:0:204:61ff:fe9d:f156", false},                // drop leading zeroes
		{"fe80::204:61ff:fe9d:f156", false},                     // collapse multiple zeroes to :: in the IPv6 address
		{"fe80:0000:0000:0000:0204:61ff:254.157.241.86", false}, // IPv4 dotted quad at the end
		{"fe80:0:0:0:0204:61ff:254.157.241.86", false},          // drop leading zeroes, IPv4 dotted quad at the end
		{"fe80::204:61ff:254.157.241.86", false},                // dotted quad at the end, multiple zeroes collapsed
		{"::1", false},                                          // localhost
		{"fe80::", false},                                       // link-local prefix
		{"2001::", false},                                       // global unicast prefix
		{"1127.01.0.1", false},
		{"255.0:3.255", false},
		{"275.0.3.255", false},
		{"127.010.0.1", false},
		{"027.01.0.1", false},
		{"0.0.0.0", true},
		{"255.255.255.255", true},
		{"255.0.3.255", true},
	}

	for _, v := range tests {
		err := IPv4Validator()(v.ip)
		if v.expected {
			assert.NoError(t, err)
		} else {
			assert.Error(t, err)
		}
	}
}

func TestIPv6Validator(t *testing.T) {
	tests := []struct {
		ip       string
		expected bool
	}{
		{"fe80:0000:0000:0000:0204:61ff:fe9d:f156", true},      // full form of IPv6
		{"fe80:0:0:0:204:61ff:fe9d:f156", true},                // drop leading zeroes
		{"fe80::204:61ff:fe9d:f156", true},                     // collapse multiple zeroes to :: in the IPv6 address
		{"fe80:0000:0000:0000:0204:61ff:254.157.241.86", true}, // IPv4 dotted quad at the end
		{"fe80:0:0:0:0204:61ff:254.157.241.86", true},          // drop leading zeroes, IPv4 dotted quad at the end
		{"fe80::204:61ff:254.157.241.86", true},                // dotted quad at the end, multiple zeroes collapsed
		{"::1", true},                                          // localhost
		{"fe80::", true},                                       // link-local prefix
		{"2001::", true},                                       // global unicast prefix
		{"0.0.0.0", false},
		{"255.255.255.255", false},
		{"255.0.3.255", false},
		{"127.010.0.1", false},
		{"027.01.0.1", false},
		{"1127.01.0.1", false},
		{"255.0:3.255", false},
		{"275.0.3.255", false},
	}

	for _, v := range tests {
		err := IPv6Validator()(v.ip)
		if v.expected {
			assert.NoError(t, err)
		} else {
			assert.Error(t, err)
		}
	}
}

func TestIPValidator(t *testing.T) {
	tests := []struct {
		ip       string
		expected bool
	}{
		{"fe80:0000:0000:0000:0204:61ff:fe9d:f156", true},      // full form of IPv6
		{"fe80:0:0:0:204:61ff:fe9d:f156", true},                // drop leading zeroes
		{"fe80::204:61ff:fe9d:f156", true},                     // collapse multiple zeroes to :: in the IPv6 address
		{"fe80:0000:0000:0000:0204:61ff:254.157.241.86", true}, // IPv4 dotted quad at the end
		{"fe80:0:0:0:0204:61ff:254.157.241.86", true},          // drop leading zeroes, IPv4 dotted quad at the end
		{"fe80::204:61ff:254.157.241.86", true},                // dotted quad at the end, multiple zeroes collapsed
		{"::1", true},                                          // localhost
		{"fe80::", true},                                       // link-local prefix
		{"2001::", true},                                       // global unicast prefix
		{"0.0.0.0", true},
		{"255.255.255.255", true},
		{"255.0.3.255", true},
	}

	for _, v := range tests {
		err := IPValidator()(v.ip)
		if v.expected {
			assert.NoError(t, err)
		} else {
			assert.Error(t, err)
		}
	}
}

func TestFQDNValidator(t *testing.T) {
	tests := []struct {
		fqdn     string
		expected bool
	}{
		{"test.example.com", true},
		{"example.com", true},
		{"example24.com", true},
		{"test.example24.com", true},
		{"test24.example24.com", true},
		{"test.example.com.", true},
		{"example.com.", true},
		{"example24.com.", true},
		{"test.example24.com.", true},
		{"test24.example24.com.", true},
		{"24.example24.com", true},
		{"test.24.example.com", true},
		{"test24.example24.com..", false},
		{"example", false},
		{"192.168.0.1", false},
		{"email@example.com", false},
		{"2001:cdba:0000:0000:0000:0000:3257:9652", false},
		{"2001:cdba:0:0:0:0:3257:9652", false},
		{"2001:cdba::3257:9652", false},
		{"", false},
	}

	for _, v := range tests {
		err := FQDNValidator()(v.fqdn)
		if v.expected {
			assert.NoError(t, err)
		} else {
			assert.Error(t, err)
		}
	}
}
