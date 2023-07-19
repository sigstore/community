<!--- This content generated with:
go run github.com/google/trillian/docs/claimantmodel/experimental/cmd/render@master --full_model_file ./docs/claimantmodel/tsa/full.yaml 
-->
```mermaid
sequenceDiagram
    actor ${TimestampAuthority}
    actor ${TimestampMonotonicVerifier}
    actor Log Operator
    actor Software Installer, entity consuming short-lived code-signing certificate
    actor Witness
    actor Witness Quorum
    ${TimestampAuthority}->>Log Operator: Add new Signed timestamp containing ${Time}
    Log Operator->>Log Operator: Integrate Signed timestamp containing ${Time}s and issue Log Checkpoint
    Log Operator->>${TimestampAuthority}: Log Checkpoint and inclusion proof
    ${TimestampAuthority}->>Software Installer, entity consuming short-lived code-signing certificate: Signed timestamp containing ${Time} with proof bundle
    Software Installer, entity consuming short-lived code-signing certificate->>Software Installer, entity consuming short-lived code-signing certificate: Verify bundle and install software
    loop Periodic append-only Verification
        Witness->>Log Operator: Fetch merkle data
        Witness->>Witness: Verify append-only
    end
    loop Periodic Signed timestamp containing ${Time} Verification
        ${TimestampMonotonicVerifier}->>Log Operator: Get all entries
        ${TimestampMonotonicVerifier}->>${TimestampMonotonicVerifier}: Verify: ${TimestampAuthority} claims a monotonically increasing ${Time}
    end
```