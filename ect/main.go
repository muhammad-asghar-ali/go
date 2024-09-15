package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Domain\tHas MX\tHas SPF\tSPF Record\tHas DMARC\tDMARC Record")

	for scanner.Scan() {
		domain := scanner.Text()
		hasMX, hasSPF, spfRecord, hasDMARC, dmarcRecord, err := checkDomain(domain)

		if err != nil {
			log.Printf("Error checking domain %s: %v\n", domain, err)
			continue
		}

		// Print the results
		fmt.Printf("%s\t%v\t%v\t%v\t%v\t%v\n", domain, hasMX, hasSPF, spfRecord, hasDMARC, dmarcRecord)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error: could not read from input: %v\n", err)
	}
}

func checkDomain(domain string) (hasMX bool, hasSPF bool, spfRecord string, hasDMARC bool, dmarcRecord string, err error) {
	mxRecords, err := net.LookupMX(domain)
	if err != nil {
		return false, false, "", false, "", fmt.Errorf("could not lookup MX records: %w", err)
	}
	hasMX = len(mxRecords) > 0

	txtRecords, err := net.LookupTXT(domain)
	if err != nil {
		return hasMX, false, "", false, "", fmt.Errorf("could not lookup TXT records: %w", err)
	}
	for _, txt := range txtRecords {
		if strings.HasPrefix(txt, "v=spf1") {
			hasSPF = true
			spfRecord = txt
			break
		}
	}

	dmarcTXTRecords, err := net.LookupTXT("_dmarc." + domain)
	if err != nil {
		return hasMX, hasSPF, spfRecord, false, "", fmt.Errorf("could not lookup DMARC records: %w", err)
	}
	for _, txt := range dmarcTXTRecords {
		if strings.HasPrefix(txt, "v=DMARC1") {
			hasDMARC = true
			dmarcRecord = txt
			break
		}
	}

	return hasMX, hasSPF, spfRecord, hasDMARC, dmarcRecord, nil
}
