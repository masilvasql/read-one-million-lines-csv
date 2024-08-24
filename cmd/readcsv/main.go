package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"

	"github.com/masilvasq/read-one-million-lines-csv/pkg/rootpath"
)

type Customer struct {
	Index            int
	CustomerID       string
	FirstName        string
	LastName         string
	Company          string
	City             string
	Country          string
	Phone1           string
	Phone2           string
	Email            string
	SubscriptionDate string
	WebSite          string
}

func main() {

	startTime := time.Now()
	fmt.Printf("Started At: %v\n", startTime.Format("02-01-2006 15:04:05"))

	rootPath := rootpath.Get()

	absoluteFilePath := fmt.Sprintf("%s/customers-10.csv", rootPath)

	file, err := os.Open(absoluteFilePath)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	csvReader := csv.NewReader(file)

	countador := 0

	for {
		rec, err := csvReader.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			panic(err)
		}

		index, err := strconv.Atoi(rec[0])

		if err != nil {
			panic(err)
		}

		customer := Customer{
			Index:            index,
			CustomerID:       rec[1],
			FirstName:        rec[2],
			LastName:         rec[3],
			Company:          rec[4],
			City:             rec[5],
			Country:          rec[6],
			Phone1:           rec[7],
			Phone2:           rec[8],
			Email:            rec[9],
			SubscriptionDate: rec[10],
			WebSite:          rec[11],
		}

		fmt.Printf("Index: %d, CustomerID: %s, FirstName: %s, LastName: %s, Company: %s, City: %s, Country: %s, Phone1: %s, Phone2: %s, Email: %s, SubscriptionDate: %s, WebSite: %s\n", customer.Index, customer.CustomerID, customer.FirstName, customer.LastName, customer.Company, customer.City, customer.Country, customer.Phone1, customer.Phone2, customer.Email, customer.SubscriptionDate, customer.WebSite)
		countador++
	}

	endTime := time.Now()
	fmt.Println("--------------------------------------------------")
	fmt.Printf("Finished At: %v\n", endTime.Sub(startTime))
	fmt.Println("--------------------------------------------------")
	fmt.Printf("Total Records: %d\n", countador)

}
