<!--- This content generated with:
go run github.com/google/trillian/docs/claimantmodel/experimental/cmd/render@master --full_model_file ./docs/claimantmodel/rekor/identity/full.yaml 
-->
```mermaid
sequenceDiagram
    actor ${OIDCIDOwner}
    actor Log Operator
    actor Software Installer
    actor Witness
    actor Witness Quorum
    ${OIDCIDOwner}->>Log Operator: Add new ${Hash}, X.509 certificate ${Certificate} containing ${OIDCIdentity}, and signature over ${Hash}
    Log Operator->>Log Operator: Integrate ${Hash}, X.509 certificate ${Certificate} containing ${OIDCIdentity}, and signature over ${Hash}s and issue Log Checkpoint
    Log Operator->>${OIDCIDOwner}: Log Checkpoint and inclusion proof
    ${OIDCIDOwner}->>Software Installer: ${Hash}, X.509 certificate ${Certificate} containing ${OIDCIdentity}, and signature over ${Hash} with proof bundle
    Software Installer->>Software Installer: Verify bundle and install software
    loop Periodic append-only Verification
        Witness->>Log Operator: Fetch merkle data
        Witness->>Witness: Verify append-only
    end
    loop Periodic ${Hash}, X.509 certificate ${Certificate} containing ${OIDCIdentity}, and signature over ${Hash} Verification
        ${OIDCIDOwner}->>Log Operator: Get all entries
        ${OIDCIDOwner}->>${OIDCIDOwner}: Verify: ${OIDCIdentity} signs ${Hash} using the key bound by ${Certificate}
    end
```