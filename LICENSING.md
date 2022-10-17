# License Policy

This document serves as the overarching policy for licensing across all sigstore projects.

## Policy Statement

All sigstore projects MUST be licensed under the Apache License, Version 2.0 (Apache-2.0). This is non-negotiable, given that sigstore is a project within the OpenSSF and the [charter of the OpenSSF](https://charter.openssf.org) requires that all software projects be Apache-2.0 licensed.

## Enforcement

Each project within the `sigstore` GitHub organization MUST have:
- a LICENSE file in the root of the repository that contains the appropriate Apache-2.0 license text
- automation that checks all the dependencies used by the project(s) contained in that repository use licenses that are compliant with the terms of the Apache-2.0 license.

An example tool that can be used to achieve this goal is the [Dependency Review GitHub Action](https://github.com/marketplace/actions/dependency-review). A reusable workflow for leveraging this action can be found in this repository [.github/workflows/reusable-dependency-review.yml](https://github.com/sigstore/community/blob/main/.github/workflows/reusable-dependency-review.yml)
