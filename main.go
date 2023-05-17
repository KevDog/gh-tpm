package main

import (
	// "fmt"
	// "log"
	"os"

	"github.com/KevDog/gh-tpm/cmd"
	// "github.com/cli/go-gh/v2"
	// "github.com/cli/go-gh/v2/pkg/api"
)

func main() {
	cmd := cmd.NewCmdRoot()
		if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
	// fmt.Println("hi world, this is the gh-tpm extension!")
	// client, err := api.DefaultRESTClient()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// response := struct{ Login string }{}
	// err = client.Get("user", &response)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Printf("running as %s\n", response.Login)
	// issueList, _, err := gh.Exec("issue", "list", "--repo", "kevdog/gh-tpm", "--limit", "5")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("issueList" + issueList.String())
	// cmd.PrintLabels()
}

// For more examples of using go-gh, see:
// https://github.com/cli/go-gh/blob/trunk/example_gh_test.go
