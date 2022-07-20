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

var UploadAccrual = func() {

	theFiles, err := ioutil.ReadDir("./temp/")
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
		remoteFile := "./accrual/" + fmt.Sprint(theFile.Name())
		localFile := "./temp/" + fmt.Sprint(theFile.Name())

		err = uploadFile(*sc, localFile, remoteFile)
		if err != nil {
			log.Fatalf("could not upload file: %v", err)
		}

		conn.Close()
		sc.Close()

	}
	os.RemoveAll("/temp/")
	log.Printf("All files are uploaded successfully")

}

var DownloadHandback = func() {

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

	// Download remote file to local file.
	// Downloaded from handback folder to files/handback folder
	for _, theFile := range theFiles {

		remoteFile := "./handback/" + theFile.Name
		localFile := "./files/handback/" + theFile.Name

		err = downloadFile(*sc, remoteFile, localFile)
		if err != nil {
			log.Fatalf("Could not download file %s; %v", theFile.Name, err)
		}
	}

	log.Printf("All files are downloaded successfully")

	conn.Close()
	sc.Close()

	return
}

type remoteFiles struct {
	Name    string
	Size    string
	ModTime string
}

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

// Upload file to sftp server
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
	_ = bytes // to avoid declared and not used error
	//log.Printf("%d bytes copied", bytes)
	dstFile.Close()
	srcFile.Close()
	return nil
}

// Download file from sftp server
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
	_ = bytes // to avoid declared and not used error
	//log.Printf("%d bytes copied to %v", bytes, dstFile)

	return nil
}

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

	//log.Printf("Connecting to %s ...\n", host)

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

	//defer conn.Close()

	// Create new SFTP client
	sc, err := sftp.NewClient(conn)
	if err != nil {
		log.Fatalf("Unable to start SFTP subsystem: %v", err)
	}
	//defer sc.Close()

	return conn, sc
}
