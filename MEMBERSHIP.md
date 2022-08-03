# Community membership

**Note:** This document is a work in progress

This doc outlines the various responsibilities of contributor roles in
Sigstore Organization.
Responsibilities for most roles are scoped to these subprojects.

| Role | Responsibilities | Requirements | Defined by |
| -----| ---------------- | ------------ | -------|
| Member | Active contributor in the community | Sponsored by 2 reviewers and multiple contributions to the project | Sigstore GitHub org member|
| Triage | Triage issues/pull requests | History of review and authorship in a subproject | GitHub Team membership |
| Codeowners | Contributions acceptance approval| Highly experienced active reviewer and contributor to a subproject | GitHub Team membership |

## New contributors

[New contributors] should be welcomed to the community by existing members,
helped with PR workflow, and directed to relevant documentation and
communication channels.

## Established community members

Established community members are expected to demonstrate their adherence to the
principles in this document, familiarity with project organization, roles,
policies, procedures, conventions, etc., and technical and/or writing ability.
Role-specific expectations, responsibilities, and requirements are enumerated
below.

## Member

Members are continuously active contributors in the community. They can have
issues and PRs assigned to them, participate through GitHub teams, and
pre-submit tests are automatically run for their PRs. Members are expected to
remain active contributors to the community.

**Defined by:** Member of the Sigstore GitHub organization

### Requirements

- Enabled [two-factor authentication] on their GitHub account
- Have made multiple contributions to the project or community.  Contribution may include, but is not limited to:
    - Authoring or reviewing PRs on GitHub. At least one PR must be **merged**.
    - Filing or commenting on issues on GitHub
    - Contributing to a project, or community discussions (e.g. meetings, Slack, email discussion
      forums, Stack Overflow)
- Subscribed to https://groups.google.com/g/sigstore-dev
- Actively contributing to 1 or more projects.
- Sponsored by 2 reviewers. **Note the following requirements for sponsors**:
    - Sponsors must have close interactions with the prospective member - e.g. code/design/proposal review, coordinating
      on issues, etc.
    - Sponsors must be reviewers or approvers in at least one CODEOWNERS file.
    - Sponsors must be from multiple member companies to demonstrate integration across community.
- **[Open an issue][membership request] against the sigstore/community repo**
   - Ensure your sponsors are @mentioned on the issue
   - Complete every item on the checklist ([preview the current version of the template][membership template])
   - Make sure that the list of contributions included is representative of your work on the project.
- Have your sponsoring reviewers reply confirmation of sponsorship: `+1`

### Responsibilities and privileges

- Responsive to issues and PRs assigned to them
- Responsive to mentions of teams they are members of
- Active owner of code they have contributed (unless ownership is explicitly transferred)
  - Code is well tested
  - Tests consistently pass
  - Addresses bugs or issues discovered after code is accepted
- They can be assigned to issues and PRs, and people can ask members for reviews with a `/cc @username`.
- Tests can be run against their PRs automatically.

**Note:** members who frequently contribute code are expected to proactively
perform code reviews and work towards becoming a primary *reviewer* for the
subproject that they are active in.

## Triage

Triage role are able to review code for quality and correctness on some part of a
subproject. They are knowledgeable about both the codebase and software
engineering principles.

**Defined by:** *triage* team membership in the subproject

Reviewer status is scoped to a part of the codebase.

**Note:** Acceptance of code contributions requires at least one approver in
addition to the assigned reviewers.

### Requirements

- member for at least 3 months
- Helped to triage issues and Pull requests
- Knowledgeable about the codebase
- Sponsored by a project approver
  - With no objections from other approvers
  - Done through PR to update the CODEOWNERS file
- May either self-nominate, be nominated by an approver in this subproject.

### Responsibilities and privileges

- Tests are automatically run for PullRequests from members of the Sigstore GitHub organization
- Code reviewer status may be a precondition to accepting large code contributions
- Responsible for project quality control via [code reviews]
  - Focus on code quality and correctness, including testing and factoring
  - May also review for more holistic issues, but not a requirement
- Expected to be responsive to review requests as per [community expectations]
- Assigned PRs to review related to the project of expertise
- Assigned test bugs related to the project of expertise

## Codeowners

Codeowners are able to both review and approve code contributions.  While
code review is focused on code quality and correctness, approval is focused on
holistic acceptance of a contribution including: backwards / forwards
compatibility, adhering to API and flag conventions, subtle performance and
correctness issues, interactions with other parts of the system, etc.

**Defined by:** *Team membership* in the Codeowners team for a subproject and update the
entry in an CODEOWNERS file in a repo owned by the Sigstore project.

Codeowners status is scoped to a part of the codebase.

### Requirements

- Reviewer of the codebase for at least 3 months
- Primary reviewer for at least 10 substantial PRs to the codebase
- Reviewed or merged at least 30 PRs to the codebase
- Being a Subject Matter Expert in the area of the code that they're responsible for, i.e PKI, transparency logs, OCI, etc.
- Nominated by a subproject owner
  - With no objections from other subproject owners
  - Done through PR to update the team membership and the top-level CODEOWNERS file

### Responsibilities and privileges

- Approver status may be a precondition to accepting large code contributions
- Demonstrate sound technical judgement
- Responsible for project quality control via [code reviews]
  - Focus on holistic acceptance of contribution such as dependencies with other features, backwards / forwards
    compatibility, API and flag definitions, etc
- Expected to be responsive to review requests as per [community expectations]
- Mentor contributors and reviewers
- May approve code contributions for acceptance

## Inactive members

_Members are continuously active contributors in the community._

A core principle in maintaining a healthy community is encouraging active
participation. It is inevitable that people's focuses will change over time and
they are not expected to be actively contributing forever.

- Lack of significant contribution to the repository
- Not assisting with reviews or triaging issues
- Not replying when get mention either in GitHub or in the Sigstore Collaboration tool (today we are using Slack)

The LFX Dashboard can help to extract metrics.

[New contributors]: /CONTRIBUTING.md
[two-factor authentication]: https://help.github.com/articles/about-two-factor-authentication
[membership request]: https://github.com/sigstore/community/issues/new?assignees=&labels=area%2Fgithub-membership&template=membership.yml&title=REQUEST%3A+New+membership+for+%3Cyour-GH-handle%3E
