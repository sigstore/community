# pulumi-github-sync

This projects uses Pulumi and the Pulumi's GitHub provider to create/update Users, Teams, Collaborators and Repositories in GitHub.

## How to use

1. Open a PR to add an Org Member, team, collaborator or repository in the directory `github-data`

Note:
  - `github-data/bots.yaml` contains the bot users for the Mattermost Organization
  - `github-data/users.yaml` contains all the Mattermost Org Members
  - `github-data/core-commiters.yaml` contains all Mattermost Core Commiters Org Members, those are not Mattermost employees
  - `github-data/repositories.yaml` contains all public/private repositories for mattermost org (not included archived and forks).
   - `github-data/teams.yaml` contains all teams.

2. Review the CI to validate if the actions is the one that are expected.
3. After the MR is reviewd and merged the post pipeline will run to run the actions.


## Add a new user

To add a new user to be part of the GitHub Org you should edit the file `github-data/users.yaml`
and add a new entry following the example

```yaml
- username: a-new-user
  role: member
  teams:
    - Core Commiters
```

- You can check the teams available in `github-data/teams.yaml`, if the user does not need to be in a team just remove the section.
- The role should be `member` in some special cases some users will be `admin`

If we need to add a `Community Core Member` add the similar but in the file `github-data/core-commiters.yaml`
If we need to add a new Bot user add the similar but in the file `github-data/bots.yaml`

## Add a new team

To add a new team you should edit the file `github-data/users.yaml`
and add a new entry following the example

```yaml
- name: Cloud
  privacy: closed
  description: Cloud Team
```

## Add a new repository or update existing ones

To add a new repository you should edit the file `github-data/repositories.yaml``

example:

```yaml
- name: my-new-repo
  owner: my-org
  homepageUrl: https://example.com
  description: My description
  archived: false
  allowAutoMerge: true
  allowMergeCommit: false
  allowRebaseMerge: false
  allowSquashMerge: true
  deleteBranchOnMerge: true
  hasDownloads: true
  hasIssues: true
  hasProjects: false
  hasWiki: false
  visibility: public/private
  vulnerabilityAlerts: true
  topics:
    - tag1
    - tag2
  collaborators:
    - username: extgernal-collab
      permission: triage
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
  template:
    owner: my-orge
    repository: template
```
