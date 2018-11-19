package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/Azure/azure-storage-blob-go/azblob"
	"net/url"
)

type AzureBlobDataSource struct {
	accountName string
	accountKey  string
}

func (ads AzureBlobDataSource) get() pagesArray {

	credential, err := azblob.NewSharedKeyCredential(ads.accountName, ads.accountKey)

	if err != nil {
		panic(err)
	}

	p := azblob.NewPipeline(credential, azblob.PipelineOptions{})

	URL, _ := url.Parse(fmt.Sprintf("https://%s.blob.core.windows.net/%s", ads.accountName, "priceswatcher"))

	containerURL := azblob.NewContainerURL(*URL, p)

	ctx := context.Background()
	_, _ = containerURL.Create(ctx, azblob.Metadata{}, azblob.PublicAccessNone)

	blobURL := containerURL.NewBlockBlobURL("dataSource.json")

	downloadResponse, err := blobURL.Download(ctx, 0, azblob.CountToEnd, azblob.BlobAccessConditions{}, false)
	stream := downloadResponse.Body(azblob.RetryReaderOptions{MaxRetryRequests: 20})

	downloadedData := &bytes.Buffer{}
	_, _ = downloadedData.ReadFrom(stream)

	pages := pagesArray{}
	err = json.Unmarshal(downloadedData.Bytes(), &pages)

	if err != nil {
		panic(err)
	}

	return pages
}
