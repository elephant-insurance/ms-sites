package services

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"context"
	"fmt"
	"io"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
)

type BlobService struct {
	accountName   string
	accountKey    string
	containerName string
	client        *azblob.Client
}

var tarFiles []string
var AllContents map[string][]byte

func NewBlobService(name, key, cName string) *BlobService {
	bs := BlobService{accountName: name, accountKey: key, containerName: cName}
	return &bs
}

func (bs *BlobService) Initialize() {

	//First we create a client
	client, err := bs.InitializeConnection()
	if err != nil {

	}
	bs.client = client
	AllContents = make(map[string][]byte)
	// Then enumerate the container and store all the tar file names
	// blob listings are returned across multiple pages
	pager := client.NewListBlobsFlatPager(bs.containerName, nil)

	// continue fetching pages until no more remain
	for pager.More() {
		// advance to the next page
		page, err := pager.NextPage(context.TODO())
		if err != nil {
		}
		// print the blob names for this page
		for _, blob := range page.Segment.BlobItems {
			tarFiles = append(tarFiles, *blob.Name)
		}

	}
	// Next is to Download each tar files and unzip and save it
	for _, name := range tarFiles {
		buffer, _ := bs.DownloadFile(name)
		bs.unzipTar(buffer)
	}
	fmt.Print("Completed")

}

func (bs *BlobService) InitializeConnection() (*azblob.Client, error) {
	blobURL := fmt.Sprintf("https://%s.blob.core.windows.net/", bs.accountName)
	credential, _ := azblob.NewSharedKeyCredential(bs.accountName, bs.accountKey)
	fmt.Print(blobURL)
	client, err := azblob.NewClientWithSharedKeyCredential(blobURL, credential, nil)
	// bs.client = client
	// AllContents = make(map[string][]byte)

	return client, err

}

func (bs *BlobService) DownloadFile(blobName string) (*bytes.Buffer, error) {
	// Create a BlockBlobURL object to a blob in the container (we assume the container already exists).
	//contentLength := int64(0) // Used for progress reporting to report the total number of bytes being downloaded.

	// downloadedData := &bytes.Buffer{}
	dr, err := bs.client.DownloadStream(context.TODO(), bs.containerName, blobName, nil)
	if err != nil {
		fmt.Println("Error ")
	}
	rs := dr.Body
	// NewResponseBodyProgress wraps the GetRetryStream with progress reporting; it returns an io.ReadCloser.
	// stream := streaming.NewResponseProgress(
	// 	rs,
	// 	func(bytesTransferred int64) {
	// 		fmt.Printf("Downloaded %d of %d bytes.\n", bytesTransferred, contentLength)
	// 	},
	// )
	// defer func(stream io.ReadCloser) {
	// 	err := stream.Close()
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// }(stream) // The client must close the response body when finished with it
	fmt.Print(*dr.ContentType)

	actualBlobData, err := io.ReadAll(rs)
	// downloadedData.ReadFrom(stream)

	// fmt.Println(len(downloadedData.Bytes()))
	//AllContents[`go-ang/`+blobName] = actualBlobData
	downloadedData := bytes.NewBuffer(actualBlobData)
	return downloadedData, err
}

func (bs *BlobService) unzipTar(buff *bytes.Buffer) {

	// buf := bytes.Buffer{}
	//file, err := os.Open("go-ang.tar.gz")

	archive, err := gzip.NewReader(buff)
	if err != nil {

		fmt.Println("There is a problem with os.Open")
	}
	tr := tar.NewReader(archive)
	for {
		hdr, err := tr.Next()
		//check if dir

		if err == io.EOF {
			break // End of archive
		}

		if err != nil {
			log.Fatal(err)
		}
		switch hdr.Typeflag {
		case tar.TypeDir:
			// if err := os.Mkdir(hdr.Name, 0755); err != nil {
			// 	log.Fatalf("ExtractTarGz: Mkdir() failed: %s", err.Error())
			// }
		case tar.TypeReg:
			// outFile, err := os.Create(hdr.Name)
			if err != nil {
				log.Fatalf("ExtractTarGz: Create() failed: %s", err.Error())
			}
			// readBytes, err := tr.Read(buf.Bytes())
			// fmt.Print(readBytes)
			// if err != nil {
			// 	log.Fatalf("ExtractTarGz: Create() failed: %s", err.Error())
			// }
			buf, _ := io.ReadAll(tr)
			// if _, err := io.Copy(&buf, tr.Read(&buf)); err != nil {
			// 	log.Fatalf("ExtractTarGz: Copy() failed: %s", err.Error())
			// }
			AllContents[hdr.Name] = buf

			// outFile.Close()

		default:
			log.Fatalf(
				"ExtractTarGz: uknown type: %s in %s",
				hdr.Typeflag,
				hdr.Name)
		}

		// if hdr.Typeflag != tar.TypeDir {
		// 	//file.Name = hdr.Name
		// 	fmt.Printf("Contents of %s:\n", hdr.Name)
		// 	// tr.Read(buf.Bytes())
		// 	_, err = io.Copy(&buf, tr)
		// 	AllContents[hdr.Name] = buf.Bytes()
		// }

		// if _, err := io.Copy(os.Stdout, tr); err != nil {
		// 	log.Fatal(err)
		// }
	}
	// file.Files = contents
	// Projects = append(Projects, file)

}
