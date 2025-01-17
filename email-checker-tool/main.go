package main 

import(
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
    "strings"
)


func main(){
	scanner := bufio.NewScanner(os.Stdin)
    fmt.Printf("domain,hasMX,hasSPF, spfRecord,hasDAMRC, damrcRecord\n")

	for scanner.Scan(){
		checkDomain(scanner.Text())
	}

	if err:=scanner.Err(); err != nil{
		log.Printf("Error: Could not read the input %v\n", err)
	}


}


func checkDomain (domain string){

	var hasMX, hasSPF, hasDMARC bool
	var spfRecord, dmarcRecord string

	mxRecord,err := net.LookupMX(domain)

	if err != nil{
		log.Printf("Error: %v\n", err)
	}

	if len(mxRecord) >0 {
		hasMX = true
	}

	txtRecord, err := net.LookupTXT(domain)

	if err != nil {
	log.Printf("Error: %v\n",err)
	}
	
	for _,record := range txtRecord{

		if strings.HasPrefix(record , "v=spf1"){
			hasSPF = true
			spfRecord = record
			break
		}
	}

	dmarcRecords,err := net.LookupTXT("_damrc." + domain)
	if err != nil {
		log.Printf("Error: %v\n", err)

	}

	for _,record := range dmarcRecords{
		if strings.HasPrefix(record,"v=DMARC1"){
			hasDMARC = true
			dmarcRecord = record
			break

		}
	}

	fmt.Printf("%v , %v, %v, %v, %v, %v  ", domain, hasMX, hasSPF,spfRecord,hasDMARC,dmarcRecord)
}