package cacerts

import (
	"io/ioutil"
	"encoding/pem"
	"crypto/x509"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
	"sort"
	"log"
)

type CaCerts struct {
	certMap map[string]*x509.Certificate
}

func New() *CaCerts {
	return &CaCerts{
		certMap:make(map[string]*x509.Certificate),
	}
}

func (p *CaCerts) print(header string) {
	fmt.Println(header)

	skeys := make([]string, 0)
	for k, _ := range p.certMap {
		skeys = append(skeys, k)
	}
	sort.Strings(skeys)

	for _, k := range skeys {
		v := p.certMap[k]
		fmt.Println("   ",  v.NotAfter, v.Subject.Organization,
			v.Subject.OrganizationalUnit, v.Subject.Country,
			v.Subject.CommonName)
	}
}

func Compare(pone *CaCerts, ptwo *CaCerts) {

	inboth := make([]string, 0)

	// remove pool two certificates
	for k, _ := range pone.certMap {
		if _, ok := ptwo.certMap[k]; ok {
			inboth = append(inboth, k)
		}
	}

	fmt.Printf("pool one total: %d\npool two total: %d\n", len(pone.certMap), len(ptwo.certMap))

	for _, ib := range inboth {
		delete(pone.certMap, ib)
		delete(ptwo.certMap, ib)
	}
	pone.print("*** Certificates only in pool one")
	ptwo.print("*** Certificates only in pool two")

}

func (c *CaCerts) AddCert(cert *x509.Certificate) {

	if time.Now().After(cert.NotAfter) {
		return
	}

	sha := sha256.Sum256(cert.Raw)
	shab := make([]byte, sha256.Size)
	for idx, b := range sha {
		shab[idx] = b
	}

	ckey := hex.EncodeToString(shab)

	if _, ok := c.certMap[ckey]; ok {
		panic("hash collision")
	}

	c.certMap[ckey] = cert
}

func (c *CaCerts) Parse(file string) error {
	pemCerts, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalf("unable to read file %q : %s", file, err)
	}

	for len(pemCerts) > 0 {
		var block *pem.Block
		block, pemCerts = pem.Decode(pemCerts)
		if block == nil {
			break
		}
		if block.Type != "CERTIFICATE" || len(block.Headers) != 0 {
			continue
		}

		cert, err := x509.ParseCertificate(block.Bytes)
		if err != nil {
			continue
		}

		c.AddCert(cert)
	}
	return nil
}