package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Domian, HasMX, HasSPF, SpfRecord, HasDMARC, DmarcRecord")

	for scanner.Scan() {
		domain := strings.TrimSpace(scanner.Text())
		if domain == "" {
			continue
		}

		if err := mchkx(domain); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error: could not read from input: %v\n", err)
		os.Exit(1)
	}

	os.Exit(0)
}

func mchkx(domain string) error {
	hasMX, mxErr := mx(domain)
	if mxErr != nil {
		return fmt.Errorf("could not retrieve MX records for %v: %w", domain, mxErr)
	}

	hasSPF, spfRecord, spfErr := spf(domain)
	if spfErr != nil {
		return fmt.Errorf("could not retrieve SPF records for %v: %w", domain, spfErr)
	}

	hasDMARC, dmarcRecord, dmarcErr := dmarc(domain)
	if dmarcErr != nil {
		return fmt.Errorf("could not retrieve DMARC records for %v: %w", domain, dmarcErr)
	}

	fmt.Printf("%v, %v, %v, %q, %v, %q\n",
		domain, hasMX, hasSPF, spfRecord, hasDMARC, dmarcRecord)

	return nil
}

func mx(domain string) (bool, error) {
	mxRecords, err := net.LookupMX(domain)
	return len(mxRecords) > 0, err
}

func spf(domain string) (bool, string, error) {
	txtRecords, err := net.LookupTXT(domain)
	if err != nil {
		return false, "", err
	}

	for _, record := range txtRecords {
		if strings.HasPrefix(record, "v=spf1") {
			return true, record, nil
		}
	}

	return false, "", nil
}

func dmarc(domain string) (bool, string, error) {
	dmarcRecords, err := net.LookupTXT("_dmarc." + domain)
	if err != nil {
		return false, "", err
	}

	for _, record := range dmarcRecords {
		if strings.HasPrefix(record, "v=DMARC1") {
			return true, record, nil
		}
	}

	return false, "", nil
}
