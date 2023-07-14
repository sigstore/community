<!--- This content generated with:
go run github.com/google/trillian/docs/claimantmodel/experimental/cmd/render@master --full_model_file ./docs/claimantmodel/fulcio/identity/full.yaml 
-->
```mermaid
sequenceDiagram
    actor Fulcio
    actor ${OIDCIDOwner}
    actor Log Operator
    actor Software Installer
    actor Witness
    actor Witness Quorum
    Fulcio->>Log Operator: Add new X.509 certificate containing ${PubKey} and ${OIDCIdentity}, signed by Fulcio
    Log Operator->>Log Operator: Integrate X.509 certificate containing ${PubKey} and ${OIDCIdentity}, signed by Fulcios and issue Log Checkpoint
    Log Operator->>Fulcio: Log Checkpoint and inclusion proof
    Fulcio->>Software Installer: X.509 certificate containing ${PubKey} and ${OIDCIdentity}, signed by Fulcio with proof bundle
    Software Installer->>Software Installer: Verify bundle and install software
    loop Periodic append-only Verification
        Witness->>Log Operator: Fetch merkle data
        Witness->>Witness: Verify append-only
    end
    loop Periodic X.509 certificate containing ${PubKey} and ${OIDCIdentity}, signed by Fulcio Verification
        ${OIDCIDOwner}->>Log Operator: Get all entries
        ${OIDCIDOwner}->>${OIDCIDOwner}: Verify: ${OIDCIDOwner} authorizes Fulcio to bind ${PubKey} to ${OIDCIdentity}
    end
```