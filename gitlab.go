package main

import (
	"encoding/base64"
	"fmt"
	"log"

	"github.com/xanzy/go-gitlab"
)

func main() {
	git, err := gitlab.NewClient("MTeJxyxBBEaUt9hpuj_f", gitlab.WithBaseURL("http://192.168.1.45:9080/api/v4"))
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	users, _, err := git.Users.ListUsers(&gitlab.ListUsersOptions{})
	for _, u := range users {
		fmt.Printf("u.ID=%d\n", u.ID)
		fmt.Printf("u.Identities=%v\n", u.Identities)

	}
	fmt.Println(len(users))
	projects, response, err := git.Projects.ListProjects(nil)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("%v\n", response)
	for _, p := range projects {
		fmt.Printf("%v\n", p)
	}
	tree, r, err := git.Repositories.ListTree(2, nil)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("总共有%d个项 \n", r.TotalItems)
	for i, t := range tree {
		fmt.Printf("%d,%s", i, t.Path)

		if t.Type == "tree" {
			fmt.Printf("\t%s\n", "文件夹")
			continue
		}
		fmt.Printf("\t%s", "文件")
		file, _, err := git.RepositoryFiles.GetFile(2, t.Path, &gitlab.GetFileOptions{Ref: sPtr("master")})
		if err != nil {
			fmt.Println(err.Error())
		}
		bytes, _ := base64.StdEncoding.DecodeString(file.Content)
		fmt.Printf("\t%s....\n", string(bytes)[0:10])
	}
}
