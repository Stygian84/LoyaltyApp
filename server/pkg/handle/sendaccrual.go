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

	//index, memberID, member fullname, transfer date, amount, reference no, partnercode

	credit_request_list, err := Queries.ListCreditRequestByStatus(context.Background(), models.TransactionStatusEnumCreated)
	if err != nil {
		log.Fatal(err)
	}
	// key should be partnercode , value = program id
	program_dict := make(map[int32]int)

	for _, credit_request := range credit_request_list {
		program_id := credit_request.Program
		if err != nil {
			log.Fatal(err)
		}

		if _, ok := program_dict[program_id]; ok {
			continue
		} else {
			program_dict[program_id] = 0
		}
	}
	// program_details,err := Queries.GetLoyaltyByID(context.Background(),int64(program_id))
	//key is program_code
	for key, element := range program_dict {
		_ = element
		//update file name to include date later
		file_name := strconv.FormatInt(int64(key), 10)
		csvFile, err := os.Create("./temp/" + file_name + "_" + time.Now().Format("2006-01-02") + ".csv")
		if err != nil {
			log.Fatal(err)
		}

		csvwriter := csv.NewWriter(csvFile)
		tempData := [][]string{
			{"Index", "Member ID", "Member FullName", "Transfer Date", "Amount", "Reference Number", "PartnerCode"},
		}

		credit_request_ls, err := Queries.GetCreditRequestByProg(context.Background(), key)
		if err != nil {
			log.Fatal(err)
		}

		idx := 1
		for _, credit_request_details := range credit_request_ls {
			rowList := []string{}

			user_id := credit_request_details.UserID
			member_id := credit_request_details.MemberID
			credit_used := credit_request_details.CreditUsed
			user_details, _ := Queries.GetUserByID(context.Background(), int64(user_id))
			full_name := user_details.FullName.String
			transfer_date := time.Now().Format("2006-01-02")
			reference_number := credit_request_details.ReferenceNumber
			// program_id := credit_request_details.Program
			// //search program by program id sql
			// partner_code := progra

			rowList = append(rowList, strconv.FormatInt(int64(idx), 10), member_id, full_name, transfer_date, strconv.FormatInt(int64(credit_used), 10), strconv.FormatInt(int64(reference_number), 10), file_name)
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

	// Update all status to pending, updatetransaction
	UploadAccrual()

	// Delete temp folder containing newly created csv file for transfer
	err = os.RemoveAll("./temp")
	if err != nil {
		log.Fatal(err)
	}
	log.Print("Disconnecting from sftp server ...")
	return err
}
