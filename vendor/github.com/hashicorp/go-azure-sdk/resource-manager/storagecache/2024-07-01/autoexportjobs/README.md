
## `github.com/hashicorp/go-azure-sdk/resource-manager/storagecache/2024-07-01/autoexportjobs` Documentation

The `autoexportjobs` SDK allows for interaction with Azure Resource Manager `storagecache` (API Version `2024-07-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/storagecache/2024-07-01/autoexportjobs"
```


### Client Initialization

```go
client := autoexportjobs.NewAutoExportJobsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `AutoExportJobsClient.Delete`

```go
ctx := context.TODO()
id := autoexportjobs.NewAutoExportJobID("12345678-1234-9876-4563-123456789012", "example-resource-group", "amlFilesystemName", "autoExportJobName")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `AutoExportJobsClient.Get`

```go
ctx := context.TODO()
id := autoexportjobs.NewAutoExportJobID("12345678-1234-9876-4563-123456789012", "example-resource-group", "amlFilesystemName", "autoExportJobName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
