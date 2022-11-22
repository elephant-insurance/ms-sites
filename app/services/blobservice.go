package services

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/elephant-insurance/go-microservice-arch/v2/log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"github.com/elephant-insurance/go-microservice-arch/v2/msrqc"
	goCache "github.com/patrickmn/go-cache"
)

type BlobService struct {
	accountName   string
	accountKey    string
	containerName string
	client        *azblob.Client
}

var Blob *BlobService

const (
	cacheExpirationSeconds = 30
	cachePurgeSeconds      = 60
)

var cache = goCache.New(cacheExpirationSeconds*time.Second, cachePurgeSeconds*time.Second)

func NewBlobService(name, key, cName string) *BlobService {
	bs := BlobService{accountName: name, accountKey: key, containerName: cName}
	return &bs
}

func (bs *BlobService) initializeClient(c msrqc.Context) (*azblob.Client, error) {
	blobURL := fmt.Sprintf("https://%s.blob.core.windows.net/", bs.accountName)
	credential, _ := azblob.NewSharedKeyCredential(bs.accountName, bs.accountKey)
	fmt.Print(blobURL)
	client, err := azblob.NewClientWithSharedKeyCredential(blobURL, credential, nil)
	return client, err

}

func (bs *BlobService) ensureContainer(c msrqc.Context) error {

	_, err := bs.client.CreateContainer(c, bs.containerName, &azblob.CreateContainerOptions{})
	return err

}

func (bs *BlobService) DownloadFiles(c msrqc.Context, name string) (error, int) {
	lw := log.ForFunc(c)
	var downloadedData bytes.Buffer
	client, err := bs.initializeClient(c)
	if err != nil {
		return err, http.StatusInternalServerError
	}
	bs.client = client

	blobName := name + `.tar.gz`

	dr, err := bs.client.DownloadStream(c, bs.containerName, blobName, nil)
	if err != nil {
		lw.WithError(err).Error("error downloading blob stream")
		blobError := err.(*azcore.ResponseError)
		return err, blobError.StatusCode
	}
	resp := dr.Body
	lw.Debug(*dr.ContentType)

	actualBlobData, errRead := io.ReadAll(resp)
	if errRead != nil {
		lw.WithError(errRead).Error("error reading blob download response body")
		return errRead, http.StatusInternalServerError
	}

	errClose := resp.Close()
	if errClose != nil {
		lw.WithError(errClose).Error("error closing blob stream")
	}
	downloadedData = *bytes.NewBuffer(actualBlobData)
	bs.unzipTar(c, downloadedData)
	return nil, http.StatusOK
}

func (bs *BlobService) unzipTar(c msrqc.Context, buff bytes.Buffer) {
	lw := log.ForFunc(c)
	archive, err := gzip.NewReader(&buff)
	if err != nil {
		lw.WithError(err).Error("error creating new gzip reader")
	}
	tr := tar.NewReader(archive)
	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break // End of archive
		}
		if err != nil {
			lw.WithError(err).Error("invalid tar header")
		}
		switch hdr.Typeflag {
		case tar.TypeDir:
		case tar.TypeReg:
			buf, errRead := io.ReadAll(tr)
			if errRead != nil {
				lw.SetName(hdr.Name).WithError(errRead).Error("error reading a fil for tarball")
			}
			// set the cache with key as file path (hdr.Name) and []byte as value
			cache.Set(hdr.Name, buf, 0)
		}
	}
}
