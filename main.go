package main

import (
    "fmt"
    "os"
    "strings"

    "github.com/c-bata/go-prompt"

    "github.com/google/go-cloud/gcp"
    "github.com/google/go-cloud/blob"
    "github.com/google/go-cloud/blob/gcsblob"

    _ "github.com/aws/aws-sdk-go/aws"
    _ "github.com/aws/aws-sdk-go/aws/credentials"
    _ "github.com/aws/aws-sdk-go/aws/session"
    _ "github.com/google/go-cloud/blob/s3blob"
)

func executer(in string) {
    if in == "exit" {
        os.Exit(0)
    } else if in == "help" {
        fmt.Println("command format:\n    cloud [service] [action]\n")
    }

    switch parse(in)[0] {
    case "gcp":
        fmt.Println("Google Cloud Platform")
        setupGCP()
    case "aws":
        fmt.Println("Amazon Web Service")
    }
}

func completer(d prompt.Document) []prompt.Suggest {
    s := []prompt.Suggest {
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

func parse(in string) []string {
    s := strings.Split(in, " ")
    return s
}

func main() {
    p := prompt.New(
        executer,
        completer,
        prompt.OptionPrefix(">>> "),
        prompt.OptionTitle("cloud console"),
    )
    p.Run()
}

// func setupGCP(ctx context.Context, bucket string) (*blob.Bucket, error) {
func setupGCP(ctx context.Context, bucket string) string {
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
    // return gcsblob.OpenBucket(ctx, bucket, c)
    return "Success!"
}

