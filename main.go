package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"strings"

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
			syncedTeams := strings.ToLower(strings.ReplaceAll(team.Name, " ", "-"))

			teamArgs := &github.TeamArgs{
				Name:                    pulumi.String(team.Name),
				CreateDefaultMaintainer: pulumi.Bool(false),
				Description:             pulumi.String(team.Description),
				Privacy:                 pulumi.String(team.Privacy),
			}
			if team.ParentTeamID != 0 {
				teamArgs.ParentTeamId = pulumi.Int(team.ParentTeamID)
			}
			syncedTeam, err := github.NewTeam(ctx, syncedTeams, teamArgs, pulumi.Protect(true))
			if err != nil {
				return err
			}

			// sync users in teams
			for _, member := range p.Config.Users {
				for _, userTeam := range member.Teams {
					if userTeam.Name == team.Name {
						_, err = github.NewTeamMembership(ctx, fmt.Sprintf("%s-%s", member.Username, strings.ToLower(syncedTeams)), &github.TeamMembershipArgs{
							TeamId:   syncedTeam.ID(),
							Username: pulumi.String(member.Username),
							Role:     pulumi.String(userTeam.Role),
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
			repoSync := &github.RepositoryArgs{
				Name:                pulumi.String(repo.Name),
				Description:         pulumi.String(repo.Description),
				HomepageUrl:         pulumi.String(repo.HomepageUrl),
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
				LicenseTemplate:     pulumi.String(repo.LicenseTemplate),
				Topics:              pulumi.ToStringArray(repo.Topics),
				VulnerabilityAlerts: pulumi.Bool(repo.VulnerabilityAlerts),
				Visibility:          pulumi.String(repo.Visibility),
				IsTemplate:          pulumi.Bool(repo.IsTemplate),
			}

			if repo.Pages.Branch != "" {
				repoPages := &github.RepositoryPagesArgs{}

				source := &github.RepositoryPagesSourceArgs{
					Branch: pulumi.String(repo.Pages.Branch),
				}

				if repo.Pages.Path != "" {
					source.Path = pulumi.String(repo.Pages.Path)
				}

				repoPages.Source = source

				if repo.Pages.CNAME != "" {
					repoPages.Cname = pulumi.String(repo.Pages.CNAME)
				}

				repoSync.Pages = repoPages
			}

			if repo.Template.Owner != "" && repo.Template.Repository != "" {
				repoSync.Template = &github.RepositoryTemplateArgs{
					Owner:      pulumi.String(repo.Template.Owner),
					Repository: pulumi.String(repo.Template.Repository),
				}
			}

			newRepo, err := github.NewRepository(ctx, repo.Name, repoSync, pulumi.Protect(true))
			if err != nil {
				return err
			}

			for _, protection := range repo.BranchesProtection {
				_, err = github.NewBranchProtection(ctx, fmt.Sprintf("%s-%s", repo.Name, protection.Branch), &github.BranchProtectionArgs{
					RepositoryId:                  newRepo.NodeId,
					Pattern:                       pulumi.String(protection.Branch),
					EnforceAdmins:                 pulumi.Bool(protection.EnforceAdmins),
					AllowsDeletions:               pulumi.Bool(protection.AllowsDeletions),
					AllowsForcePushes:             pulumi.Bool(protection.AllowsForcePushes),
					RequiredLinearHistory:         pulumi.Bool(protection.RequiredLinearHistory),
					RequireSignedCommits:          pulumi.Bool(protection.RequireSignedCommits),
					RequireConversationResolution: pulumi.Bool(protection.RequireConversationResolution),
					RequiredStatusChecks: github.BranchProtectionRequiredStatusCheckArray{
						&github.BranchProtectionRequiredStatusCheckArgs{
							Strict:   pulumi.Bool(false),
							Contexts: pulumi.ToStringArray(protection.StatusChecks),
						},
					},
					RequiredPullRequestReviews: github.BranchProtectionRequiredPullRequestReviewArray{
						&github.BranchProtectionRequiredPullRequestReviewArgs{
							DismissStaleReviews:          pulumi.Bool(protection.DismissStaleReviews),
							RestrictDismissals:           pulumi.Bool(protection.RestrictDismissals),
							RequireCodeOwnerReviews:      pulumi.Bool(protection.RequireCodeOwnerReviews),
							RequiredApprovingReviewCount: pulumi.Int(protection.RequiredApprovingReviewCount),
						},
					},
				})
				if err != nil {
					return err
				}
			}

			for _, collaborator := range repo.Collaborators {
				// sync collaborators
				_, err := github.NewRepositoryCollaborator(ctx, fmt.Sprintf("%s-%s", repo.Name, collaborator.Username), &github.RepositoryCollaboratorArgs{
					Permission:                pulumi.String(collaborator.Permission),
					Repository:                pulumi.String(repo.Name),
					Username:                  pulumi.String(collaborator.Username),
					PermissionDiffSuppression: pulumi.Bool(false),
				})
				if err != nil {
					return err
				}
			}

			for _, team := range repo.Teams {
				// sync teams for a repo
				// format the team name to be the team slug, eg. "My Team" become "my-team"
				formatedTeam := strings.ToLower(strings.ReplaceAll(team.Name, " ", "-"))

				teamID := formatedTeam
				// used when importing existing team
				if team.ID != "" {
					teamID = team.ID
				}
				_, err := github.NewTeamRepository(ctx, fmt.Sprintf("%s-%s", repo.Name, formatedTeam), &github.TeamRepositoryArgs{
					Permission: pulumi.String(team.Permission),
					Repository: pulumi.String(repo.Name),
					TeamId:     pulumi.String(teamID),
				})
				if err != nil {
					return err
				}
			}
		}

		return nil
	})
}
