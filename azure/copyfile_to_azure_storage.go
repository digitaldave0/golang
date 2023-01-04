package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"os"

	"github.com/Azure/azure-storage-blob-go/azblob"
)

func copyFolderToAzure(localFolderPath, azureContainerName, storageAccountName, storageAccountKey string) {
	// Create an Azure Storage account context
	credential, err := azblob.NewSharedKeyCredential(storageAccountName, storageAccountKey)
	if err != nil {
		log.Fatal(err)
	}
	pipeline := azblob.NewPipeline(credential, azblob.PipelineOptions{})
	URL, _ := url.Parse(fmt.Sprintf("https://%s.blob.core.windows.net/%s", storageAccountName, azureContainerName))
	containerURL := azblob.NewContainerURL(*URL, pipeline)

	// Get a list of all files in the local folder
	localFiles, err := ioutil.ReadDir(localFolderPath)
	if err != nil {
		log.Fatal(err)
	}

	// Iterate through the list of local files and upload each one to Azure storage
	for _, file := range localFiles {
		// Open the local file
		f, err := os.Open(file.Name())
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		// Create a blob URL for the file
		blobURL := containerURL.NewBlockBlobURL(file.Name())

		// Upload the local file to Azure
		_, err = azblob.UploadFileToBlockBlob(context.Background(), f, blobURL, azblob.UploadToBlockBlobOptions{
			BlockSize:   4 * 1024 * 1024,
			Parallelism: 16})
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Successfully uploaded %s to Azure storage\n", file.Name())
	}
}

func main() {
	// Example usage:
	copyFolderToAzure(
		"/path/to/local/folder",
		"my-azure-container",
		"my-storage-account",
		"my-storage-account-key",
	)
}
