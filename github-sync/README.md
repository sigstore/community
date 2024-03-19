# sigstore-github-sync

This projects uses Pulumi and the Pulumi's GitHub provider to create/update Users, Teams, Collaborators and Repositories in GitHub for the Sigstore Organization.
We have a [Pulumi's Open Source account](https://www.pulumi.com/pricing/open-source-free-tier/), thanks for [Pulumi](https://www.pulumi.com/)!

## How to use

1. Open a Pull Request to add/change an Org Member, team, collaborator or repository in the directory `github-data`

Note:
  - `github-data/bots.yaml` contains the bot users for the Sigstore Organization
  - `github-data/users.yaml` contains all the Sigstore Org Members
  - `github-data/repositories.yaml` contains all public/private repositories for Sigstore Organization (not included archived and forks).
  - `github-data/teams.yaml` contains all teams.

2. Review the CI to validate if the actions is the one that are expected.
3. After the Pull Request is reviewed and merged the post pipeline will run to run the actions.


## Add a new user

To add a new user to be part of the GitHub Org you should edit the file `github-data/users.yaml`
and add a new entry following the example

```yaml
- username: a-new-user
  role: member # most of the time the role will be member, but also can be admin
  teams:
    - name: cosign-codeowners
      role: member # Must be one of `member` or `maintainer`. Defaults to `member`.
```

- You can check the teams available in `github-data/teams.yaml`, if the user does not need to be in a team just remove the section.
- The role should be `member` in some special cases some users will be `admin`

If we need to add a new Bot user add the similar but in the file `github-data/bots.yaml`

## Add a new team

To add a new team you should edit the file `github-data/teams.yaml`
and add a new entry following the example

```yaml
- name: my-new-team
  privacy: closed
  description: Optional description
```

## Add a new repository or update existing ones

To add a new repository you should edit the file `github-data/repositories.yaml``

example:

```yaml
- name: My-RepoName
  owner: sigstore
  description: description # optional
  homepageUrl: ""
  allowAutoMerge: true|false # optional
  allowMergeCommit: true|false # optional
  allowRebaseMerge: true|false # optional
  allowSquashMerge: true|false # optional
  archived: true|false # optional
  autoInit: true|false # optional
  deleteBranchOnMerge: true|false # optional
  hasDownloads: true|false # optional
  hasIssues: true|false # optional
  hasProjects: true|false # optional
  hasWiki: true|false # optional
  vulnerabilityAlerts: true|false # optional
  visibility: public|private
  licenseTemplate: ""
  topics: []
  collaborators: []
    # - username: sigstore-bot
    #   permission: push # Must be one of `pull`, `push`, `maintain`, `triage` or `admin` for organization-owned repositories.
  teams:
    - name: Core Team
      id: 4563391
      permission: push # Must be one of `pull`, `triage`, `push`, `maintain`, or `admin`. Defaults to `pull`.
```

If you need to configure the GitHub Pages you can add the following definition

```yaml
pages:
  cname: custom cname opcional
  branch: gh-pages
  path: /docs
```

- if path is not set it will default to `/`
- if you set a cname make sure you configure the DNS in aws route53 or similar

If you need to create a new repository that will use a template repository you can set the following

```yaml
isTemplate: true
template:
  owner: my-org
  repository: template
```

To add a branch protection and configure the branch add the following settings:

```yaml
branchesProtection:
  - pattern: main
    enforceAdmins: true|false # optional
    allowsDeletions: true|false # optional
    allowsForcePushes: true|false # optional
    requiredLinearHistory: true|false # optional
    dismissStaleReviews: true|false # optional
    requireSignedCommits: true|false # optional
    requiredApprovingReviewCount: 1
    requireCodeOwnerReviews: true|false # optional
    requireConversationResolution: true|false # optional
    restrictDismissals: true|false # optional
    requireBranchesUpToDate: true|false # optional
    statusChecks:
      - DCO # optional. The name of the status checks that you want to be required for a PR.
    pushRestrictions:
      - MyTeam # optional. Only people, teams, or apps allowed to push will be able to create new branches matching this rule.
    dismissalRestrictions:
      - MyTeam # optional. Specify people, teams, or apps allowed to dismiss pull request reviews.
```

### Remove Users/Collaborators, teams and repositories

All resources, when created, are protected against accidental deletion. If your Pull Request removes a user,
collaborator, team, or repository, the CI will fail, and some Administrators with Pulumi's access will need to run a command
manually to unlock the required resource.
