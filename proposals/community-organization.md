# Light-touch project management for `sigstore` on Github

`sigstore` and its associated projects are developed and hosted primarily on Github, and, as such, make use of Github's tools for project management: particularly, issues and pull requests.

As a nascent cross-organizational effort, `sigstore` does not currently enforce much, if any, process around work items -- indeed, such structure may introduce unnecessary overhead at a time when it's not really needed. Of course, as the project scales and more people get involved, such organization will rapidly prove needed. With the project going public and the inauguration of community meetings providing a central place for contributors to discuss things, now seems like a good time to propose a plan for light-touch developer processes compatible with a loosely-organized team. This document will attempt to outline that plan, in addition to proposing the use of some select automation tools to further aid productivity.

The stated proposals assume that the project will continue to live primarily on Github, using Github issues to outline work items and pull requests to deliberate on and approve completed work.

## Kanban board(s)

The sole direct process change of this proposal will be the creation and regular use of a basic kanban board to track work.

At the very least, it should contain a "Triage" column, where new issues and PRs automatically land (with the help of some automation, detailed below). This would allow core developers to observe, in a single glance, any new items that may need their attention. Once looked at or otherwise addressed, the issue or PR should be moved to another column on the board, perhaps an "In Progress" column. (The exact structure of the rest of the board can vary depending on developer needs.) If desired, the Triage column could be assessed during community meetings.

### How many boards?

Using one board would centralize all new issues and PRs from all `sigstore` member projects in one place. An alternative to this structure would be separate boards for each sub-focus or sub-project; for example, all new `sigstore/rekor` issues could land on a Rekor board, while `sigstore/fulcio` issues could land on a Fulcio board. This proposal does not specifically dictate one approach or another.

## Project automation tools

Working in tandem with the proposed project boards would be up to several project automation tools from [`actions-automation`](https://github.com/actions-automation) (disclosure: written and maintained by the author of this proposal).

These tools are designed to be light-touch, enriching Github's issues and pull requests to be more useful without any additional feature cruft on top. They are also all written as straightforward Github Actions, meaning they do not rely on any external infrastructure to work (only Github's own Actions infrastructure). They have been deployed on several other projects on Github (including [Enarx](https://github.com/enarx) and [Keylime](https://github.com/keylime)) with positive results.

Setup would require:

- the creation of a dedicated "bot" account given "Contributor" (_not_ write) access to participating `sigstore` repos;
- the provisioning of a restricted-scope (`public_repo`, `org:read`, and `org:write`) personal access token from that bot account, available via Github's "secrets" architecture.

There are three actions to consider; how many of them get used is ultimately up to the discretion of the community, though the creation of the kanban board hinges on the `triage` action.

### [`triage`](https://github.com/actions-automation/triage)

A straightforward action that adds all newly-opened (or reopened) issues and pull requests on a given repo to a specified org-wide board, such as the kanban board proposed above.

### [`pull-request-responsibility`](https://github.com/actions-automation/pull-request-responsibility)

A set of several actions that push PRs towards completion faster -- namely:

- Automatically request a random set of PR reviewers from a specified Github team and suggested reviewers, accounting for reviewer availability.
- Automatically adjust assignees on a pull request such that the current assignees are the ones that must take action to move the pull request forward.

A full list of actions and what they do is available [here](https://github.com/actions-automation/pull-request-responsibility#supported-actions).

This action in particular is very helpful in ensuring PRs don't get "lost" or otherwise abandoned by reviewers.

### [`manage-your-labels`](https://github.com/actions-automation/manage-your-labels)

An action that allows for robust label mirroring and management across an entire organization's repositories. This effectively stores the set of labels available to organization repos in a config file -- which has the side benefit of allowing changes to labels via PR.
