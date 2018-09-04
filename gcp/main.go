package main

import (
    "fmt"
    "context"
    _ "flag"
    "log"

    "github.com/google/go-cloud/gcp"
    "github.com/google/go-cloud/blob"
    "github.com/google/go-cloud/blob/gcsblob"
)

func main() {
    ctx := context.Background()
    fmt.Println(ctx)

    bucketName := "go-cloud-bucket231"
    var (
        b *blob.Bucket
        err error
    )
    b, err = setupGCP(ctx, bucketName)

    if err != nil {
        log.Fatalf("Failed to setup bucket: %s", err)
    }

    fmt.Println(b, err)
}

func setupGCP(ctx context.Context, bucket string) (*blob.Bucket, error) {
    // DefaultCredentials assumes a user has logged in with gcloud.
    // See here for more information:
    // https://cloud.google.com/docs/authentication/getting-started
    creds, err := gcp.DefaultCredentials(ctx)
    if err != nil {
        return nil, err
    }
    c, err := gcp.NewHTTPClient(gcp.DefaultTransport(), gcp.CredentialsTokenSource(creds))
    if err != nil {
        return nil, err
    }
    // The bucket name must be globally unique.
    return gcsblob.OpenBucket(ctx, bucket, c)
}

