<!--- This content generated with:
go run github.com/google/trillian/docs/claimantmodel/experimental/cmd/render@master --full_model_file ./docs/claimantmodel/rekor/key/full.yaml 
-->
```mermaid
sequenceDiagram
    actor ${KeyOwner}
    actor Log Operator
    actor Software Installer
    actor Witness
    actor Witness Quorum
    ${KeyOwner}->>Log Operator: Add new ${Hash}, public key ${PubKey}, and signature over ${Hash}
    Log Operator->>Log Operator: Integrate ${Hash}, public key ${PubKey}, and signature over ${Hash}s and issue Log Checkpoint
    Log Operator->>${KeyOwner}: Log Checkpoint and inclusion proof
    ${KeyOwner}->>Software Installer: ${Hash}, public key ${PubKey}, and signature over ${Hash} with proof bundle
    Software Installer->>Software Installer: Verify bundle and install software
    loop Periodic append-only Verification
        Witness->>Log Operator: Fetch merkle data
        Witness->>Witness: Verify append-only
    end
    loop Periodic ${Hash}, public key ${PubKey}, and signature over ${Hash} Verification
        ${KeyOwner}->>Log Operator: Get all entries
        ${KeyOwner}->>${KeyOwner}: Verify: ${Key} signs ${Hash}, verifiable with ${PubKey}
    end
```