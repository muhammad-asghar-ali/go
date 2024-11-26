# Domain Mail Configuration Checker

A simple Go program to check DNS records for a domain's mail configuration. It verifies if the domain has MX, SPF, and DMARC records and outputs the results in a CSV format.

## Run

To run the program:
`go run main.go`

## Features

- Checks if the domain has an **MX record** (mail exchange).
- Checks for the presence of an **SPF record** (Sender Policy Framework).
- Checks for the presence of a **DMARC record** (Domain-based Message Authentication, Reporting, and Conformance).
- Outputs the results in a CSV format.
- Exits immediately if an error occurs during processing.

## Prerequisites

- Go 1.18 or later.
- Access to the internet to perform DNS lookups.

# Domain Mail Configuration Checker - Sample Output

The following is an example output of the Domain Mail Configuration Checker program. The results are displayed in CSV format:
`exmaple.com, true, true, "v=spf1 include:\_spf.mx.cloudflare.net ~all", true, "v=DMARC1; p=reject; pct=100; rua=mailto:f196fc49e9b6435498f27c54ac0e578f@dmarc-reports.cloudflare.net,mailto:cloudflare@dmarc.area1reports.com; ruf=mailto:cloudflare@dmarc.area1reports.com"`

### Explanation of Columns:

- **Domain**: The domain being checked.
- **HasMX**: Indicates whether the domain has an MX (Mail Exchange) record.
- **HasSPF**: Indicates whether the domain has an SPF (Sender Policy Framework) record.
- **SPFRecord**: The actual SPF record if it exists.
- **HasDMARC**: Indicates whether the domain has a DMARC (Domain-based Message Authentication, Reporting, and Conformance) record.
- **DMARCRecord**: The actual DMARC record if it exists.

### Error Handling:

- If an error occurs during the check (e.g., DNS resolution failure), the program will display the error message for the domain and stop further processing.
