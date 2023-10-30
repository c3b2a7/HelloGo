package options

import (
	"crypto/elliptic"
	"crypto/x509"
	"crypto/x509/pkix"
	"io"
	"net"
	"time"
)

func WithCA(isCA bool) CertificationOption {
	return func(options *certificationOptions) {
		options.IsCA = isCA
	}
}

func WithIssuer(issuer *x509.Certificate) CertificationOption {
	return func(options *certificationOptions) {
		options.Issuer = issuer
	}
}

func WithIssuerPrivateKey(issuerPrivateKey any) CertificationOption {
	return func(options *certificationOptions) {
		options.IssuerPrivateKey = issuerPrivateKey
	}
}

func WithSubject(subject pkix.Name) CertificationOption {
	return func(options *certificationOptions) {
		options.Subject = subject
	}
}

func WithNotBefore(notBefore time.Time) CertificationOption {
	return func(options *certificationOptions) {
		options.NotBefore = notBefore
	}
}

func WithNotAfter(notAfter time.Time) CertificationOption {
	return func(options *certificationOptions) {
		options.NotAfter = notAfter
	}
}

func WithIPs(ips []net.IP) CertificationOption {
	return func(options *certificationOptions) {
		options.IPs = ips
	}
}

func WithDomains(domains []string) CertificationOption {
	return func(options *certificationOptions) {
		options.Domains = domains
	}
}

func WithAlgorithm(algorithm string) CertificationOption {
	return func(options *certificationOptions) {
		options.Algorithm = algorithm
	}
}

func WithRandom(random io.Reader) KeyOption {
	return func(options *keyOptions) {
		options.Random = random
	}
}

func WithKeySize(keySize int) KeyOption {
	return func(options *keyOptions) {
		options.KeySize = keySize
	}
}

func WithCurve(curve elliptic.Curve) KeyOption {
	return func(options *keyOptions) {
		options.Curve = curve
	}
}
