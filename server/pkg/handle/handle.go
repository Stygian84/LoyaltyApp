package handle

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"

	"golang.org/x/crypto/ssh"

	"github.com/pkg/sftp"
)

const (
	sftpUser = "sutd_2022_c4g2"
	sftpPass = "Sutd_2022_c4g2Sutd_2022_c4g2"
	sftpHost = "66.220.9.51"
	sftpPort = "22"
)

// localtempDir should be "./temp/" while remoteDir should be "./accrual/"
func UploadAccrual(localtempDir string, remoteDir string) {
	theFiles, err := ioutil.ReadDir(localtempDir)
	if err != nil {
		log.Fatal(err)
	}

	//List all files in local accrual directory
	log.Printf("%19s %12s %s", "MOD TIME", "SIZE", "NAME")
	for _, theFile := range theFiles {
		log.Printf("%19s %12d %s", theFile.ModTime().Format("2006-01-02 15:04:05"), theFile.Size(), theFile.Name())
	}

	for _, theFile := range theFiles {
		conn, sc := ConnectToSFTP()

		// Upload local file to remote file
		remoteFile := remoteDir + fmt.Sprint(theFile.Name())
		localFile := localtempDir + fmt.Sprint(theFile.Name())

		err = uploadFile(*sc, localFile, remoteFile)
		if err != nil {
			log.Fatalf("could not upload file: %v", err)
		}

		conn.Close()
		sc.Close()

	}

	log.Printf("All files are uploaded successfully")

}

type remoteFiles struct {
	Name    string
	Size    string
	ModTime string
}

// To list all files in the remoteDir
func listFiles(sc sftp.Client, remoteDir string) (theFiles []remoteFiles, err error) {

	files, err := sc.ReadDir(remoteDir)
	if err != nil {
		return theFiles, fmt.Errorf("Unable to list remote dir: %v", err)
	}

	for _, f := range files {
		var name, modTime, size string

		name = f.Name()
		modTime = f.ModTime().Format("2006-01-02 15:04:05") // Don't change this
		size = fmt.Sprintf("%12d", f.Size())

		if f.IsDir() {
			name = name + "/"
			modTime = ""
			size = "PRE"
		}

		theFiles = append(theFiles, remoteFiles{
			Name:    name,
			Size:    size,
			ModTime: modTime,
		})
	}

	return theFiles, nil
}

// Upload a file to sftp server
func uploadFile(sc sftp.Client, localFile, remoteFile string) (err error) {
	log.Printf("Uploading [%s] to [%s] ...", localFile, remoteFile)

	srcFile, err := os.Open(localFile)
	if err != nil {
		return fmt.Errorf("Unable to open local file: %v", err)
	}

	// Make remote directories recursion
	parent := filepath.Dir(remoteFile)
	path := string(filepath.Separator)
	dirs := strings.Split(parent, path)
	for _, dir := range dirs {
		path = filepath.Join(path, dir)
		sc.Mkdir(path)
	}

	dstFile, err := sc.OpenFile(remoteFile, (os.O_WRONLY | os.O_CREATE | os.O_TRUNC))
	if err != nil {
		return fmt.Errorf("Unable to open remote file: %v", err)
	}

	bytes, err := io.Copy(dstFile, srcFile)
	if err != nil {
		return fmt.Errorf("Unable to upload local file: %v", err)
	}
	_ = bytes
	dstFile.Close()
	srcFile.Close()
	return nil
}

// Download a file from sftp server
func downloadFile(sc sftp.Client, remoteFile, localFile string) (err error) {

	log.Printf("Downloading [%s] to [%s] ...\n", remoteFile, localFile)

	srcFile, err := sc.OpenFile(remoteFile, (os.O_RDONLY))
	if err != nil {
		return fmt.Errorf("unable to open remote file: %v", err)
	}
	defer srcFile.Close()

	dstFile, err := os.Create(localFile)
	if err != nil {
		return fmt.Errorf("unable to open local file: %v", err)
	}
	defer dstFile.Close()

	bytes, err := io.Copy(dstFile, srcFile)
	if err != nil {
		return fmt.Errorf("unable to download remote file: %v", err)
	}
	_ = bytes

	return nil
}

// To connect to a SFTP server
func ConnectToSFTP() (*ssh.Client, *sftp.Client) {
	// Create a url
	rawurl := fmt.Sprintf("sftp://%v:%v@%v", sftpUser, sftpPass, sftpHost)

	// Parse the URL
	parsedUrl, err := url.Parse(rawurl)
	if err != nil {
		log.Fatalf("Failed to parse SFTP To Go URL: %s", err)
	}

	// Get user name and pass
	user := parsedUrl.User.Username()
	pass, _ := parsedUrl.User.Password()

	// Parse Host and Port
	host := parsedUrl.Host

	var auths []ssh.AuthMethod

	// Use password authentication if provided
	if pass != "" {
		auths = append(auths, ssh.Password(pass))
	}

	// Initialize client configuration
	config := ssh.ClientConfig{
		User:            user,
		Auth:            auths,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),

		Timeout: 30 * time.Second,
	}

	addr := fmt.Sprintf("%s:%s", host, sftpPort)

	// Connect to server
	conn, err := ssh.Dial("tcp", addr, &config)
	if err != nil {
		log.Fatalf("Failed to connect to host [%s]: %v", addr, err)
	}

	// Create new SFTP client
	sc, err := sftp.NewClient(conn)
	if err != nil {
		log.Fatalf("Unable to start SFTP subsystem: %v", err)
	}

	return conn, sc
}
