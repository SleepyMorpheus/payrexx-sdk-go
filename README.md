# Go SDK for Payrexx REST API



## Usage
```go
import "github.com/SleepyMorpheus/payrexxsdk"

func main() {
    // Create a payrexx client instance
    var client, _ = payrexxsdk.NewClient("<INSTANCE_NAME>", "<SECRET>", payrexxsdk.APIBaseDefault)
    
    // Fetch a gateway
    gateway, err = client.GatewayRetrieve("1")
    if err != nil {
        log.Println(err)
        return
    }

    // ... use gateway
}
```

### Missing Endpoints
In case an endpoint is missing, you can use the built-in functions of the sdk to perform a request. `NewClient -> NewRequest -> Send`
