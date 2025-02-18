# Community membership

This document outlines the various responsibilities of contributor roles in
the Sigstore organization.
Responsibilities for roles are scoped to projects, aka repositories, within the Sigstore organization.

| Role | Responsibilities | Requirements | Defined by |
| -----| ---------------- | ------------ | -------|
| Reviewer | Review pull requests, triage issues | History of review and authorship in a project | GitHub "Write" Team membership |
| Codeowner | Approve and merge contributions, set project direction | Highly experienced subject matter expert, active reviewer and contributor to a project | GitHub "Maintain" Team membership |

## New contributors

New contributors are welcomed to the community by existing members,
and directed to [relevant documentation](https://docs.sigstore.dev)
and [communication channels](https://join.slack.com/t/sigstore/shared_invite/zt-1z7jzpemb-xEKSUtpgDFXpIEMwMYZQKQ).

## Established community members

Established community members are expected to demonstrate their adherence to the
principles in this document, demonstrating familiarity with project organization,
roles, policies, procedures, conventions, and technical and/or writing ability.
Role-specific expectations, responsibilities, and requirements are provided
below.

### Community member requirements

The requirements outlined below apply to Reviewers and Codeowners.

- Enabled [two-factor authentication](https://docs.github.com/en/authentication/securing-your-account-with-two-factor-authentication-2fa/about-two-factor-authentication) on their GitHub account
- Have made multiple contributions to the project or community.  Contributions should include:
    - Authoring or reviewing PRs on GitHub. Multiple PRs must be **merged** to become a Reviewer or Codeowner
    - Filing or commenting on issues on GitHub
    - Contributing to a project, or community discussions (e.g. meetings, Slack, email discussion
      forums, Stack Overflow)
- Subscribed to https://groups.google.com/g/sigstore-dev

### Community member responsibilities

The responsibilities outlined below apply to Reviewers and Codeowners.

- Responsive to issues and PRs assigned to them
- Responsive to mentions of teams they are members of
- Active owner of code they have contributed (unless ownership is explicitly transferred)
  - Code is well tested
  - Tests consistently pass
  - Addresses bugs or issues discovered after code is accepted

**Note:** Members who frequently contribute code are expected to proactively
perform code reviews and work towards becoming a *Reviewer* for the
project that they are active in.

## Reviewer

The Reviewer role is able to review code for quality and correctness on the majority of a
project. They are knowledgeable about both the codebase and software
engineering principles. A reviewer can approve, but not merge, a PR.

**Defined by:** Team membership with *Write* permissions for the project. The project
will also have a push restriction branch protection rule to allow only Codeowners
to push to branches (merge PRs).

Reviewer status may be scoped to a part of the codebase. Ideally reviewers are familiar
with the entire codebase. Reviewers should use judgment when reviewing code they are not
familiar with, and delegate reviews to other reviewers or codeowners that are familiar.

**Note:** Acceptance of code contributions requires at least one Codeowner in
addition to the assigned reviewers.

### Requirements

- Active community member for at least 3 months
- Active in issue triage and pull request reviews
- Knowledgeable about the codebase
- Sponsored by a project Codeowner
  - With no objections from other Codeowners
  - Done through PR in the Community repository, either:
    - Updating the Write Team membership
    - Granting the user the `push` permission
- May either self-nominate or be nominated by a Codeowner in the project

### Responsibilities and privileges

- Responsible for project quality control through code reviews
  - *Must* focus on code quality and correctness, including testing and factoring
  - *Must* be familiar with code content. If not, *must* delegate review to other reviewers
  - *May* also review for more holistic issues, but not a requirement
- Expected to be responsive to review requests as per community expectations
- Assigned PRs to review in the project
- Assigned issues to investigate in the project

## Codeowners

Codeowners are able to both approve and merge code contributions. While
code review is focused on code quality and correctness, approval for merge is focused on
holistic acceptance of a contribution including: backward compatibility,
adhering to API and flag conventions, security and threat modeling, subtle performance and
correctness issues, interactions with other parts of the system, etc.

**Defined by:** Team membership with *Maintain* permission for the project. The project will
also have a push restriction branch protection rule to allow only these Codeowners
to push to branches (merge PRs).

Codeowners status is for the entire project.

### Requirements

- Reviewer of the codebase for at least 3 months
- Primary reviewer for at least 10 substantial PRs to the codebase
- Reviewed a significant number of PRs in the codebase. For a large codebase, at least 30 PRs
- Be a Subject Matter Expert in the project that they're responsible for, e.g. PKI, transparency logs, OCI, etc.
- Sponsored by another Codeowner
  - With no objections from other Codeowners
  - Done through PR in the Community repository, updating the Maintain or Codeowners Team membership
- May either self-nominate or be nominated by a Codeowner in the project

### Responsibilities and privileges

- Demonstrate sound technical judgment
- Responsible for project quality control via code reviews
  - *Must* focus on holistic acceptance of contributions such as dependencies with other features, backward
    compatibility, API and flag definitions, security risks, etc.
  - *Must* be familiar with code content. If not, *must* delegate review to other codeowners
  - *Must* verify code changes are well-tested
- Expected to be responsive to review requests as per community expectations
- Mentor contributors and reviewers
- Set direction for project with input from the community

## Inactive members

_Members are continuously active contributors in the community._

A core principle in maintaining a healthy community is encouraging active
participation. It is inevitable that people's focuses will change over time and
they are not expected to be actively contributing forever.

Signs of inactivity include:

- Lack of significant contributions to the repository
- Not assisting with reviews or triaging issues
- Not replying when mentioned either in GitHub or Slack 

The LFX Dashboard and GitHub per-repo contribution insights can be used to determine activity.

Inactive members may be removed by active Codeowners for a given project or TSC members.
