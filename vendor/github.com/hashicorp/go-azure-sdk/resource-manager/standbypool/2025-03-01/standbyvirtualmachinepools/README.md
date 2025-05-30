
## `github.com/hashicorp/go-azure-sdk/resource-manager/standbypool/2025-03-01/standbyvirtualmachinepools` Documentation

The `standbyvirtualmachinepools` SDK allows for interaction with Azure Resource Manager `standbypool` (API Version `2025-03-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/standbypool/2025-03-01/standbyvirtualmachinepools"
```


### Client Initialization

```go
client := standbyvirtualmachinepools.NewStandbyVirtualMachinePoolsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `StandbyVirtualMachinePoolsClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := standbyvirtualmachinepools.NewStandbyVirtualMachinePoolID("12345678-1234-9876-4563-123456789012", "example-resource-group", "standbyVirtualMachinePoolName")

payload := standbyvirtualmachinepools.StandbyVirtualMachinePoolResource{
	// ...
}


if err := client.CreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `StandbyVirtualMachinePoolsClient.Delete`

```go
ctx := context.TODO()
id := standbyvirtualmachinepools.NewStandbyVirtualMachinePoolID("12345678-1234-9876-4563-123456789012", "example-resource-group", "standbyVirtualMachinePoolName")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `StandbyVirtualMachinePoolsClient.Get`

```go
ctx := context.TODO()
id := standbyvirtualmachinepools.NewStandbyVirtualMachinePoolID("12345678-1234-9876-4563-123456789012", "example-resource-group", "standbyVirtualMachinePoolName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `StandbyVirtualMachinePoolsClient.ListByResourceGroup`

```go
ctx := context.TODO()
id := commonids.NewResourceGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group")

// alternatively `client.ListByResourceGroup(ctx, id)` can be used to do batched pagination
items, err := client.ListByResourceGroupComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `StandbyVirtualMachinePoolsClient.ListBySubscription`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.ListBySubscription(ctx, id)` can be used to do batched pagination
items, err := client.ListBySubscriptionComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `StandbyVirtualMachinePoolsClient.Update`

```go
ctx := context.TODO()
id := standbyvirtualmachinepools.NewStandbyVirtualMachinePoolID("12345678-1234-9876-4563-123456789012", "example-resource-group", "standbyVirtualMachinePoolName")

payload := standbyvirtualmachinepools.StandbyVirtualMachinePoolResourceUpdate{
	// ...
}


read, err := client.Update(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
