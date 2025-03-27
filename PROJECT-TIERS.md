# Introduction and motivation

As the Sigstore TSC looks to formalize and future-proof our stewardship of the Sigstore project, we find ourselves needing to honor lifecycle differences amongst the codebases and services that comprise the Sigstore estate. The existing flat landscape of projects has caused a certain amount of confusion amongst would-be contributors / adopters to-date, and the time has come to create a robust framework for signalling project maturity to the community as well as to create a notional “on-ramp” for future community projects. In this document we will establish that framework.

## Goals

* Create a tiered labeling system for GitHub projects to signal the level of support / community involvement that a given repo has.  
* Establish a tier/label for all existing projects  
* Establish guidelines for future projects \- both for accepting new projects and for moving existing ones between tiers

# Tiers and their definitions

## Core projects

Core projects are fundamental to the goals of Sigstore as articulated by the TSC. They encompass the basic services that power the Sigstore Public Good Instance, as well as the official CLI tool and the client libraries and repos that are currently in use by 3rd party integrators (e.g. package registries like npm and RubyGems and software vendors like GitHub and GitLab) to support implementations of Sigstore for use cases like provenance attestations, publishing attestations, runtime verification etc. 

## Community projects

Community projects are ones which are deemed to have sufficient mindshare and promise that the Sigstore TSC determines that they should be brought into the Sigstore GitHub organization and live under the governance of the overall Sigstore umbrella.

### Creating a community project

Obviously anyone with an interest in Sigstore can create their own GitHub repository and begin writing code. If such a repo originator believes that their source code / project should become a Community project, they must complete the following steps:

* Create a `ROADMAP.md` document in the root of their repo describing the intended evolution of the project.  
* Create a `CONTRIBUTORS.md` document in the root of their repo naming the GitHub users who will be responsible for maintaining the project.  
* Open an issue on the Sigstore Community Repo describing the project, its benefit, and describing why it’s an appropriate candidate for stewardship as a Community project.  
* Notify the Sigstore TSC (by opening an issue on [github.com/sigstore/tsc](http://github.com/sigstore/tsc)) that they’d like to discuss their project at an upcoming meeting. A Sigstore TSC member will add it to the meeting agenda.  
* Attend the meeting of the Sigstore TSC where their project is to be discussed  
* Sigstore TSC will then discuss the project and its roadmap (in closed session, if desired) and vote on whether to bring it into the GitHub organization as a designated Community project.  
* The originating maintainer will then transfer the GitHub repository to Sigstore, including having its GitHub configuration managed by the Sigstore TSC

## Criteria for up-leveling from Community to Core

* TSC must vote to move a project from Community to Core  
* There must be a dedicated maintainer and at least one other person acting as a contributor. Ideally these should be from two different employers.  
* The project must have had at least X contributions within the trailing Y months  
* In the case of a language client, it must pass all tests described in the sigstore-conformance repo

## Why would a project want to go from Community to Core?

Core projects enjoy consideration for targeted funding from OpenSSF, Sigstore, or other funding sources. Membership in Core is a signal to the broader OSS community that the project is stable enough for adoption or productization.

# Project lists

## Core

### Services / CLI tools

* [Cosign](https://github.com/sigstore/cosign)  
* [Fulcio](https://github.com/sigstore/fulcio)  
* [Rekor](https://github.com/sigstore/rekor)  
* [Policy Controller](https://github.com/sigstore/policy-controller/)  
* [Timestamp Authority](https://github.com/sigstore/timestamp-authority)
* [rekor-search-ui](https://github.com/sigstore/rekor-search-ui)

### Libraries / clients

* [sigstore-conformance](https://github.com/sigstore/sigstore-conformance)
* [sigstore-go](https://github.com/sigstore/sigstore-go)  
* [sigstore-js](https://github.com/sigstore/sigstore-js)  
* [sigstore-java](https://github.com/sigstore/sigstore-java)  
* [sigstore-python](https://github.com/sigstore/sigstore-python)  
* [sigstore-ruby](https://github.com/sigstore/sigstore-ruby)  
* [cosign-installer](https://github.com/sigstore/cosign-installer)  

## Community

### Services / CLI tools

* [GitSign](https://github.com/sigstore/gitsign)  
* [Rekor-monitor](https://github.com/sigstore/rekor-monitor)  
* [helm-sigstore](https://github.com/sigstore/helm-sigstore)  
* [cosign-gatekeeper-provider](https://github.com/sigstore/cosign-gatekeeper-provider)

### Libraries / clients
* [model-transparency](https://github.com/sigstore/model-transparency)  