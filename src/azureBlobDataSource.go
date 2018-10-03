package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/Azure/azure-storage-blob-go/2016-05-31/azblob"
	"net/url"
)

type AzureBlobDataSource struct {
	accountName string
	accountKey string
}

func (ads AzureBlobDataSource) get() pagesArray {

	credential := azblob.NewSharedKeyCredential( ads.accountName, ads.accountKey)
	p := azblob.NewPipeline(credential, azblob.PipelineOptions{})

	URL, _ := url.Parse(
		fmt.Sprintf("https://%s.blob.core.windows.net/%s", ads.accountName, "priceswatcher"))

	containerURL := azblob.NewContainerURL(*URL, p)

	ctx := context.Background()
	_, _ = containerURL.Create(ctx, azblob.Metadata{}, azblob.PublicAccessNone)

	blobURL := containerURL.NewBlockBlobURL("dataSource.json")

	stream := azblob.NewDownloadStream(ctx, blobURL.GetBlob, azblob.DownloadStreamOptions{})
	downloadedData := &bytes.Buffer{}
	_, _ = downloadedData.ReadFrom(stream)

	pages := pagesArray{}
	err := json.Unmarshal(downloadedData.Bytes(), &pages)

	if err != nil {
		panic(err)
	}

	return pages
}
