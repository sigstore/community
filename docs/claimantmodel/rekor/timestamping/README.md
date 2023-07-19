# Claim for timestamping in Rekor

The claim for timestamping in Rekor is a work in progress, as uploading signed timestamps is not yet supported in Rekor.

This claim has some unexpected properties:

* The claim and statement reference the claim for identity-based signatures. The claim can be expanded to
`${OIDCIdentity} signs ${Hash} using the key bound by ${Certificate}, valid with respect to ${Timestamp} whose signature is over ${Signature}`, and the statement `${Hash}, X.509 certificate containing ${OIDCIdentity}, signed ${Timestamp}, and ${Signature} over ${Hash}`.
* There is no Verifier, because no entity can verify the veracity of a timestamping claim. Anyone can generate a timestamp with a signing event. Note that there could be a Verifier for the Claimant Model for a Timestamp Authority, but that is out of scope for this claim.
* The generated flow chart may not make sense because of the lack of a Verifier.