package models

import (
	"net/mail"
	"strings"

	"golang.org/x/net/idna"
)

// normalizeEmailAddress converts internationalized domain names to ASCII so
// they can be validated and sent through SMTP libraries that expect punycode.
func normalizeEmailAddress(raw string) (*mail.Address, error) {
	raw = strings.TrimSpace(raw)
	if raw == "" {
		return nil, mail.ErrHeaderNotPresent
	}

	namePart := ""
	addrPart := raw

	if start := strings.LastIndex(raw, "<"); start >= 0 && strings.HasSuffix(raw, ">") {
		namePart = strings.TrimSpace(raw[:start])
		addrPart = strings.TrimSpace(raw[start+1 : len(raw)-1])
	}

	at := strings.LastIndex(addrPart, "@")
	if at <= 0 || at == len(addrPart)-1 {
		return nil, ErrInvalidFromAddress
	}

	localPart := addrPart[:at]
	domainPart := addrPart[at+1:]
	asciiDomain, err := idna.Lookup.ToASCII(domainPart)
	if err != nil {
		return nil, err
	}

	normalized := localPart + "@" + asciiDomain
	if namePart != "" {
		normalized = namePart + " <" + normalized + ">"
	}
	return mail.ParseAddress(normalized)
}

// NormalizeEmailAddressForAPI exposes the shared normalization logic outside
// the models package.
func NormalizeEmailAddressForAPI(raw string) (*mail.Address, error) {
	return normalizeEmailAddress(raw)
}
