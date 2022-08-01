package handle

import (
	"context"
	"encoding/csv"
	"esc/ascendaRoyaltyPoint/pkg/config"
	"esc/ascendaRoyaltyPoint/pkg/models"
	"log"
	"os"
	"strconv"
	"strings"
)

var Queries *models.Queries

func ReadHandbackFile() (err error) {
	config.Connect()
	Queries = models.New(config.GetDB())
	conn, sc := ConnectToSFTP()

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
		// column title = records [0]
		// transfer_date := records [1][2]
		// amount := records [2][2]
		reference_number := strings.Trim(records[3][2], "\"")
		outcome_code := strings.Trim(records[4][2], "\"")
		int_reference_number, err := strconv.ParseInt(reference_number, 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		int_outcome_code, err := strconv.ParseInt(outcome_code, 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		credit_request, err := Queries.GetCreditRequestByID(context.Background(), int_reference_number)
		if err != nil {
			log.Fatal(err)
		}
		// int_outcome_code = 0 -> approved
		// int_outcome_code = 1 -> rejected
		if int_outcome_code == 0 {
			// Update transaction status to approved
			args := models.UpdateTransactionStatusByIDParams{
				TransactionStatus: models.TransactionStatusEnumApproved,
				ReferenceNumber:   int_reference_number,
			}
			_ = Queries.UpdateTransactionStatusByID(context.Background(), args)
			log.Printf("Reference Number %v Is Successfully Approved", reference_number)
		} else {
			// Update transaction status to rejected
			args := models.UpdateTransactionStatusByIDParams{
				TransactionStatus: models.TransactionStatusEnumRejected,
				ReferenceNumber:   int_reference_number,
			}
			err = Queries.UpdateTransactionStatusByID(context.Background(), args)
			if err != nil {
				log.Fatal(err)
			}
			log.Printf("Reference Number %v Is Successfully Rejected", reference_number)

			credit_details, err := Queries.GetCreditRequestByID(context.Background(), int_reference_number)
			if err != nil {
				log.Fatal(err)
			}

			// Refunded credit used since transaction was rejected
			user_id := credit_details.UserID
			credit_used := credit_details.CreditUsed
			log.Printf("%.2f credits are refunded to USERID %v", credit_used, user_id)
			balanceargs := models.IncrBalanceParams{
				CreditBalance: credit_used,
				ID:            int64(user_id),
			}
			err = Queries.IncrBalance(context.Background(), balanceargs)
			if err != nil {
				log.Fatal(err)
			}

			//Notify user through email

		}
		_ = credit_request
	}

	log.Println("Disconnecting from SFTP server ...")
	conn.Close()
	sc.Close()

	return err

}
