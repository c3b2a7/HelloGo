package options

import (
	"crypto/elliptic"
	"crypto/x509"
	"crypto/x509/pkix"
	"io"
	"net"
	"time"
)

type CertificationOptionBuilder struct {
	opts []CertificationOption
}

func (b *CertificationOptionBuilder) appendOpt(opt func(*certificationOptions)) {
	b.opts = append(b.opts, opt)
}

func (b *CertificationOptionBuilder) Build() []CertificationOption {
	return b.opts
}

func (b *CertificationOptionBuilder) WithCA(isCA bool) CertificationOptionBuilder {
	b.appendOpt(func(options *certificationOptions) {
		options.IsCA = isCA
	})
	return *b
}

func (b *CertificationOptionBuilder) WithIssuer(issuer *x509.Certificate) CertificationOptionBuilder {
	b.appendOpt(func(options *certificationOptions) {
		options.Issuer = issuer
	})
	return *b
}

func (b *CertificationOptionBuilder) WithIssuerPrivateKey(issuerPrivateKey any) CertificationOptionBuilder {
	b.appendOpt(func(options *certificationOptions) {
		options.IssuerPrivateKey = issuerPrivateKey
	})
	return *b
}

func (b *CertificationOptionBuilder) WithSubject(subject pkix.Name) CertificationOptionBuilder {
	b.appendOpt(func(options *certificationOptions) {
		options.Subject = subject
	})
	return *b
}

func (b *CertificationOptionBuilder) WithNotBefore(notBefore time.Time) CertificationOptionBuilder {
	b.appendOpt(func(options *certificationOptions) {
		options.NotBefore = notBefore
	})
	return *b
}

func (b *CertificationOptionBuilder) WithNotAfter(notAfter time.Time) CertificationOptionBuilder {
	b.appendOpt(func(options *certificationOptions) {
		options.NotAfter = notAfter
	})
	return *b
}

func (b *CertificationOptionBuilder) WithIPs(ips []net.IP) CertificationOptionBuilder {
	b.appendOpt(func(options *certificationOptions) {
		options.IPs = ips
	})
	return *b
}

func (b *CertificationOptionBuilder) WithDomains(domains []string) CertificationOptionBuilder {
	b.appendOpt(func(options *certificationOptions) {
		options.Domains = domains
	})
	return *b
}

func (b *CertificationOptionBuilder) WithAlgorithm(algorithm string) CertificationOptionBuilder {
	b.appendOpt(func(options *certificationOptions) {
		options.Algorithm = algorithm
	})
	return *b
}

type KeyOptionBuilder struct {
	opts []KeyOption
}

func (b *KeyOptionBuilder) appendOpt(opt func(options *keyOptions)) {
	b.opts = append(b.opts, opt)
}

func (b *KeyOptionBuilder) Build() []KeyOption {
	return b.opts
}

func (b *KeyOptionBuilder) WithRandom(random io.Reader) KeyOptionBuilder {
	b.appendOpt(func(options *keyOptions) {
		options.Random = random
	})
	return *b
}

func (b *KeyOptionBuilder) WithKeySize(keySize int) KeyOptionBuilder {
	b.appendOpt(func(options *keyOptions) {
		options.KeySize = keySize
	})
	return *b
}

func (b *KeyOptionBuilder) WithCurve(curve elliptic.Curve) KeyOptionBuilder {
	b.appendOpt(func(options *keyOptions) {
		options.Curve = curve
	})
	return *b
}
