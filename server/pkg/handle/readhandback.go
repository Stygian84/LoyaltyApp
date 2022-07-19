package handle

import (
	"context"
	"encoding/csv"
	"esc/ascendaRoyaltyPoint/pkg/config"
	"esc/ascendaRoyaltyPoint/pkg/models"
	"log"
	"os"
	"strconv"
)

var Queries *models.Queries

func ReadHandbackFile() (err error) {
	config.Connect()
	Queries = models.New(config.GetDB())
	conn, sc := ConnectToSFTP()

	defer conn.Close()
	defer sc.Close()

	// List files in the remote handback directory .
	theFiles, err := listFiles(*sc, "./handback")
	if err != nil {
		log.Fatalf("failed to list files in ./handback: %v", err)
	}

	log.Printf("Found Files in ./handback Files")
	// Output each file name and size in bytes
	log.Printf("%19s %12s %s", "MOD TIME", "SIZE", "NAME")
	for _, theFile := range theFiles {
		log.Printf("%19s %12s %s", theFile.ModTime, theFile.Size, theFile.Name)
	}

	// Show content
	for _, theFile := range theFiles {

		remoteFile := "./handback/" + theFile.Name

		srcFile, err := sc.OpenFile(remoteFile, (os.O_RDONLY))

		if err != nil {
			log.Printf("Unable to open file: %v", err)
		}
		csvReader := csv.NewReader(srcFile)
		records, err := csvReader.ReadAll()
		if err != nil {
			log.Fatal(err)
		}
		//column title = records [0]
		// transfer_date := records [1][2]
		// amount := records [2][2]
		reference_number := records[3][2]
		outcome_code := records[4][2]
		int_reference_number, err := strconv.ParseInt(reference_number, 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		int_outcome_code, err := strconv.ParseInt(outcome_code, 10, 64)
		if err != nil {
			log.Fatal(err)
		}

		log.Print(int_outcome_code)

		credit_request, err := Queries.GetCreditRequestByID(context.Background(), int_reference_number)
		if int_outcome_code == 0 {
			// Update transaction status to approved
			_ = Queries.UpdateTransactionStatus(context.Background(), "approved")
			log.Print("Is Successfully Approved")
		} else {
			// Update transaction status to rejected
			_ = Queries.UpdateTransactionStatus(context.Background(), "rejected")
			log.Print("Is Successfully Rejected")

		}

		//
		_ = err
		_ = credit_request
	}

	return err

}
