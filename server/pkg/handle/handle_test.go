package handle

import (
	"esc/ascendaRoyaltyPoint/pkg/utils"
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

// To test connection with SFTP and read a directory
func TestSFTPListFiles(t *testing.T) {
	conn, sc := ConnectToSFTP()
	sc.Mkdir("./test")

	// List files in the remote handback directory .
	theFiles, err := listFiles(*sc, "./test")
	require.NoError(t, err)

	for _, theFile := range theFiles {
		require.Contains(t, theFile.Name, "csv")
	}
	sc.RemoveDirectory("./test")
	conn.Close()
	sc.Close()
}

// Test whether SFTP has the required directory
func TestSFTPDir(t *testing.T) {

	conn, sc := ConnectToSFTP()

	_, err := sc.ReadDir("./handback")
	require.NoError(t, err)

	conn.Close()
	sc.Close()

	conn2, sc2 := ConnectToSFTP()

	_, err2 := sc2.ReadDir("./accrual")
	require.NoError(t, err2)

	conn2.Close()
	sc2.Close()
}

// Test if a file can be uploaded to SFTP server
func TestUploadFile(t *testing.T) {
	conn, sc := ConnectToSFTP()
	sc.Mkdir("./test")
	err := uploadFile(*sc, "./testdata/test.txt", "./test/test.txt")
	require.NoError(t, err)
	conn.Close()
	sc.Close()

	conn, sc = ConnectToSFTP()
	theremoteFiles, err := listFiles(*sc, "./test")
	require.NoError(t, err)

	for _, theFile := range theremoteFiles {
		require.Contains(t, theFile.Name, "test.txt")
	}

	sc.RemoveDirectory("./test")
	conn.Close()
	sc.Close()
}

// Test if a file from SFTP server can be downloaded
func TestDownloadFile(t *testing.T) {
	conn, sc := ConnectToSFTP()
	sc.Mkdir("./test")
	err := uploadFile(*sc, "./testdata/test.txt", "./test/test.txt")
	require.NoError(t, err)
	conn.Close()
	sc.Close()

	conn, sc = ConnectToSFTP()
	downloadFile(*sc, "./test/test.txt", "./testdata/test2.txt")
	require.FileExists(t, "./testdata/test2.txt")

	os.Remove("./testdata/test2.txt")
	sc.RemoveDirectory("./test")
	conn.Close()
	sc.Close()
}

// Test UploadAccrual function
func TestUploadAccrual(t *testing.T) {

	conn, sc := ConnectToSFTP()
	sc.Mkdir("./test")
	conn.Close()
	sc.Close()

	UploadAccrual("./testdata/", "./test/")

	os.Mkdir("./testresult", 0777)
	conn, sc = ConnectToSFTP()
	downloadFile(*sc, "./test/test.txt", "./testresult/test2.txt")
	require.FileExists(t, "./testresult/test2.txt")

	os.RemoveAll("./testresult/")
	sc.RemoveDirectory("./test")
	conn.Close()
	sc.Close()
}

func TestSendEmail(t *testing.T) {
	email_to := "ghostkirito84@gmail.com"
	user_name := utils.RandomString(5)
	reference_number := strconv.Itoa(int(utils.RandomInt(5, 10)))
	user_id := int32(utils.RandomInt(1, 4))
	credit_balance := utils.RandomFloat(1000)
	credit_used := utils.RandomFloat(100)
	err := sendEmail(email_to, true, user_name, reference_number, user_id, credit_balance, credit_used, "Approved")
	require.NoError(t, err)
	err = sendEmail(email_to, false, user_name, reference_number, user_id, credit_balance, credit_used, "Rejected")
	require.NoError(t, err)
}
