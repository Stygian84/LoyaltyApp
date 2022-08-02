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
	conn, sc := ConnectToSFTP()

	// mkdir temp
	path := "temp"
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			log.Println(err)
		}
	}

	//index, memberID, member fullname, transfer date, amount, reference no, partnercode
	// left with index and partnercode
	credit_request_list, err := Queries.ListCreditRequest(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	program_dict := make(map[int32]int)

	for _, credit_request := range credit_request_list {
		program_id := credit_request.Program
		if _, ok := program_dict[program_id]; ok {
			continue
		} else {
			program_dict[program_id] = 0
		}
	}

	//key is program_id
	for key, element := range program_dict {

		_ = element
		//update file name to include date later
		file_name := strconv.FormatInt(int64(key), 10)
		csvFile, err := os.Create("./temp/" + file_name + ".csv")
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
			tempList := []string{}
			user_id := credit_request_details.UserID
			member_id := credit_request_details.MemberID
			credit_used := credit_request_details.CreditUsed
			user_details, _ := Queries.GetUserByID(context.Background(), int64(user_id))
			full_name := user_details.FullName.String
			transfer_date := time.Now().Format("2006-01-02")
			reference_number := credit_request_details.ReferenceNumber
			tempList = append(tempList, strconv.FormatInt(int64(idx), 10), member_id, full_name, transfer_date, strconv.FormatInt(int64(credit_used), 10), strconv.FormatInt(int64(reference_number), 10), file_name)
			idx += 1
			tempData = append(tempData, tempList)
		}

		for _, empRow := range tempData {
			_ = csvwriter.Write(empRow)
		}

		csvwriter.Flush()
		csvFile.Close()
	}

	// Update all status to pending, updatetransaction
	UploadAccrual()
	err = os.RemoveAll("./temp")
	sc.Close()
	conn.Close()
	log.Print("Disconnecting from sftp server ...")
	return err
}
