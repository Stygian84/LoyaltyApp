package handle

import (
	"context"
	"encoding/csv"
	"errors"
	"esc/ascendaRoyaltyPoint/pkg/config"
	"esc/ascendaRoyaltyPoint/pkg/models"
	"log"
	"os"
	"strconv"
	"time"
)

func SendAccrual() (err error) {
	config.Connect()
	Queries = models.New(config.GetDB())

	// mkdir temp
	path := "temp"
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			log.Println(err)
		}
	}

	credit_request_list, err := Queries.ListCreditRequestByStatus(context.Background(), models.TransactionStatusEnumCreated)
	if err != nil {
		log.Fatal(err)
	}

	// Dictionary key is partnercode, value is credit req reference number
	program_dict := make(map[string][]int)

	for _, credit_request := range credit_request_list {

		// Look for partnercode for each program id
		program_id := credit_request.Program
		program_details, err := Queries.GetLoyaltyByID(context.Background(), int64(program_id))
		partner_code := program_details.PartnerCode

		credit_request_reference_number := credit_request.ReferenceNumber
		if err != nil {
			log.Fatal(err)
		}
		program_dict[partner_code] = append(program_dict[partner_code], int(credit_request_reference_number))
	}

	// Key = partnercode , element = reference number
	// 1st loop : Iterate through each partnercode, to get relevant reference number
	// 2nd loop : Iterate through each reference number and append the index + content to a csv file
	for key, element := range program_dict {
		_ = element

		// Create a csv file whose file name is based on a partnercode
		partner_code := key
		csvFile, err := os.Create("./temp/" + partner_code + "_" + time.Now().Format("2006-01-02") + ".csv")
		if err != nil {
			log.Fatal(err)
		}

		// Create the first row of the csv file
		csvwriter := csv.NewWriter(csvFile)
		tempData := [][]string{
			{"Index", "Member ID", "Member FullName", "Transfer Date", "Amount", "Reference Number", "PartnerCode"},
		}

		idx := 1
		for _, reference_number := range element {
			credit_request_details, err := Queries.GetCreditRequestByID(context.Background(), int64(reference_number))
			if err != nil {
				log.Print(err)
			}

			rowList := []string{}

			user_id := credit_request_details.UserID
			member_id := credit_request_details.MemberID
			credit_used := credit_request_details.CreditUsed
			user_details, _ := Queries.GetUserByID(context.Background(), int64(user_id))
			full_name := user_details.FullName
			transfer_date := time.Now().Format("2006-01-02")
			reference_number := credit_request_details.ReferenceNumber

			rowList = append(rowList, strconv.FormatInt(int64(idx), 10), member_id, full_name, transfer_date, strconv.FormatInt(int64(credit_used), 10), strconv.FormatInt(int64(reference_number), 10), partner_code)
			idx += 1
			tempData = append(tempData, rowList)

			// Update transaction status to Pending
			args := models.UpdateTransactionStatusByIDParams{
				TransactionStatus: models.TransactionStatusEnumPending,
				ReferenceNumber:   int64(reference_number),
			}
			_ = Queries.UpdateTransactionStatusByID(context.Background(), args)

		}

		for _, tempRow := range tempData {
			_ = csvwriter.Write(tempRow)
		}

		csvwriter.Flush()
		csvFile.Close()
	}

	// Upload csv files to sftp server
	UploadAccrual("./temp/", "./accrual")

	// For demo purposes the temp folder does not get deleted

	// Delete temp folder containing newly created csv file for transfer
	// err = os.RemoveAll("./temp")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	log.Print("Disconnecting from sftp server ...")
	return err
}
