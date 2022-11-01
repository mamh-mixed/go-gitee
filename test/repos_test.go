package test

import (
	"fmt"
	"github.com/mamh-mixed/go-gitee/gitee"
	"testing"
)

func TestListBranches(t *testing.T) {
	branches, response, err := client.Repositories.ListBranches(ctx, "mamh-java", "jenkins-jenkins")
	for index, br := range branches {
		fmt.Println(index, *br.Name, *br.Protected, *br.ProtectionUrl)
	}
	fmt.Println(response)
	fmt.Println(err)
}

func TestGetBranch(t *testing.T) {
	branch, response, err := client.Repositories.GetBranch(ctx, "mamh-mixed", "go-gitee", "main")
	fmt.Println(branch)
	fmt.Println(response)
	fmt.Println(err)
}

func TestGetCommit(t *testing.T) {
	commit, response, err := client.Repositories.GetCommit(ctx, "mamh-mixed", "go-gitee", "8896821c53eda6698ef5c75ba5182e547e8476f1")

	fmt.Println(commit, response, err)

}

func TestListCommits(t *testing.T) {
	var opts = &gitee.CommitsListOptions{}
	for {
		commits, response, err := client.Repositories.ListCommits(ctx, "mamh-mixed", "go-gitee", opts)
		if err != nil {
			fmt.Println(err)
			return
		}
		for index, commit := range commits {
			fmt.Println(index, len(commits), *commit.Commit.Message, *commit.SHA)
		}
		if response.NextPage == 0 {
			break
		}
		opts.Page = response.NextPage
	}

}

func TestList(t *testing.T) {
	opts := &gitee.RepositoryListOptions{}
	repository, response, err := client.Repositories.List(ctx, "", opts)

	fmt.Println("Namespace", repository[0].Namespace)
	fmt.Println("Owner", repository[0].Owner)
	fmt.Println("Aassigner", repository[0].Aassigner)
	fmt.Println("Parent", repository[0].Parent)
	fmt.Println("Permission", repository[0].Permission)
	fmt.Println("Assignee", repository[0].Assignee[0])
	fmt.Println("Testers", repository[0].Testers[0])
	fmt.Println("Programs", repository[0].Programs)
	fmt.Println("Enterprise", repository[0].Enterprise)
	fmt.Println("ProjectLabels", repository[0].ProjectLabels)

	fmt.Println("response", response)
	fmt.Println("err", err)

}

func TestList1(t *testing.T) {
	opts := &gitee.RepositoryListOptions{}
	repository, response, err := client.Repositories.List(ctx, "elunez", opts)

	fmt.Println("repository", repository)

	fmt.Println("response", response)
	fmt.Println("err", err)

}
