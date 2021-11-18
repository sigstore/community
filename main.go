package main

import (
	"fmt"
	"log"
	"os"
	"path"

	github "github.com/pulumi/pulumi-github/sdk/v4/go/github"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"

	ghData "github.com/cpanato/pulumi-github-sync/pkg/config"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		conf := config.New(ctx, "")
		ghConfig := conf.Require("github-data-directory")

		stat, err := os.Stat(ghConfig)
		if err != nil {
			log.Fatalf("Failed to stat %s: %v\n", ghConfig, err)
		}
		p := ghData.NewParser()

		if stat.IsDir() {
			err = p.ParseDir(ghConfig)
		} else {
			err = p.ParseFile(ghConfig, path.Dir(ghConfig))
		}
		if err != nil {
			log.Fatalf("Failed to load config: %v\n", err)
		}

		// sync users
		for _, member := range p.Config.Users {
			_, err := github.NewMembership(ctx, member.Username, &github.MembershipArgs{
				Role:     pulumi.String(member.Role),
				Username: pulumi.String(member.Username),
			}, pulumi.Protect(true))
			if err != nil {
				return err
			}
		}

		for _, team := range p.Config.Teams {
			syncedTeam, err := github.NewTeam(ctx, team.Name, &github.TeamArgs{
				Name:        pulumi.String(team.Name),
				Description: pulumi.String(team.Description),
				Privacy:     pulumi.String(team.Privacy),
			}, pulumi.Protect(true))
			if err != nil {
				return err
			}

			for _, member := range p.Config.Users {
				for _, userTeam := range member.Teams {
					if userTeam == team.Name {
						_, err = github.NewTeamMembership(ctx, fmt.Sprintf("%s-%s", member.Username, team.Name), &github.TeamMembershipArgs{
							TeamId:   syncedTeam.ID(),
							Username: pulumi.String(member.Username),
							Role:     pulumi.String("member"),
						})
						if err != nil {
							return err
						}
					}
				}
			}
		}

		for _, repo := range p.Config.Repositories {
			// sync repos
			_, err := github.NewRepository(ctx, repo.Name, &github.RepositoryArgs{
				Name:                pulumi.String(repo.Name),
				Description:         pulumi.String(repo.Description),
				AllowAutoMerge:      pulumi.Bool(repo.AllowAutoMerge),
				AllowMergeCommit:    pulumi.Bool(repo.AllowMergeCommit),
				AllowRebaseMerge:    pulumi.Bool(repo.AllowRebaseMerge),
				AllowSquashMerge:    pulumi.Bool(repo.AllowSquashMerge),
				AutoInit:            pulumi.Bool(repo.AutoInit),
				DeleteBranchOnMerge: pulumi.Bool(repo.DeleteBranchOnMerge),
				HasDownloads:        pulumi.Bool(repo.HasDownloads),
				HasIssues:           pulumi.Bool(repo.HasIssues),
				HasProjects:         pulumi.Bool(repo.HasProjects),
				HasWiki:             pulumi.Bool(repo.HasWiki),
				HomepageUrl:         pulumi.String(repo.HomepageUrl),
				LicenseTemplate:     pulumi.String(repo.LicenseTemplate),
				Topics:              pulumi.ToStringArray(repo.Topics),
				VulnerabilityAlerts: pulumi.Bool(repo.VulnerabilityAlerts),
				Visibility:          pulumi.String(repo.Visibility),
			}, pulumi.Protect(true))
			if err != nil {
				return err
			}

			for _, collaborator := range repo.Collaborators {
				// sync collaborators
				_, err := github.NewRepositoryCollaborator(ctx, fmt.Sprintf("%s-%s", repo.Name, collaborator.Username), &github.RepositoryCollaboratorArgs{
					Permission: pulumi.String(collaborator.Permission),
					Repository: pulumi.String(repo.Name),
					Username:   pulumi.String(collaborator.Username),
				})
				if err != nil {
					return err
				}
			}
		}

		return nil
	})
}
