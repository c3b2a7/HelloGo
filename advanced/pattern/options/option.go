package options

import (
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"io"
	"net"
	"os"
	"time"
)

var (
	defaultKeyOptions = keyOptions{
		Random:  rand.Reader,
		KeySize: 2048,
		Curve:   elliptic.P384(),
	}

	defaultCertificationOptions = certificationOptions{
		IsCA: true,
		Subject: pkix.Name{
			Organization: []string{"Easy IsCA"},
			CommonName:   "Easy IsCA Root",
		},
		NotBefore: time.Now(),
		NotAfter:  time.Now().AddDate(20, 0, 0),

		Algorithm: lookupEnv("ALGORITHM", "RSA"),
		KeyOpts:   defaultKeyOptions,
	}
)

func lookupEnv(envName, def string) string {
	if v, ok := os.LookupEnv(envName); ok {
		return v
	}
	return def
}

type CertificationOption func(*certificationOptions)

type certificationOptions struct {
	IsCA             bool
	Issuer           *x509.Certificate
	IssuerPrivateKey any

	Subject   pkix.Name
	NotBefore time.Time
	NotAfter  time.Time
	IPs       []net.IP
	Domains   []string

	Algorithm string
	KeyOpts   keyOptions
}

func (c *certificationOptions) ToCSR() (csr x509.Certificate) {
	csr.IsCA = c.IsCA
	csr.Subject = c.Subject
	csr.NotBefore = c.NotBefore
	csr.NotAfter = c.NotAfter
	csr.IPAddresses = c.IPs
	csr.DNSNames = c.Domains
	return csr
}

func (c *certificationOptions) IsRootCA() bool {
	return (c.IsCA && !c.IsMiddleCA()) || c.IsCA
}

func (c *certificationOptions) IsMiddleCA() bool {
	return c.IsCA && c.Issuer != nil && c.IssuerPrivateKey != nil
}

type KeyOption func(*keyOptions)

type keyOptions struct {
	Random  io.Reader
	KeySize int
	Curve   elliptic.Curve
}
