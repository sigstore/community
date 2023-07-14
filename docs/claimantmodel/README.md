# Sigstore Claimant Model

Sigstore's claimant model includes a set of claims produced by Rekor and Fulcio logs. See [Claimant Model](https://github.com/google/trillian/blob/master/docs/claimantmodel/CoreModel.md) for more information and terminology.

Sequence diagrams generated with [Claimant Model Render Tool](https://github.com/google/trillian/tree/master/docs/claimantmodel/experimental/cmd/render). This tool is a work in progress, so there may be some errors in the generated models.

Each claim is in `model.md` in each folder. `full.md` contains the claim made in both the domains of Rekor or Fulcio and the log. The sequence diagram shows the interactions between all Actors. 

All claims are also provided in this readme.

## Fulcio: Identity-based signing

<dl>
<dt>Claim<sup>Fulcio</sup></dt>
<dd><i>${OIDCIDOwner} authorizes Fulcio to bind ${PubKey} to ${OIDCIdentity}</i></dd>
<dt>Statement<sup>Fulcio</sup></dt>
<dd>X.509 certificate containing ${PubKey} and ${OIDCIdentity}, signed by Fulcio</dd>
<dt>Claimant<sup>Fulcio</sup></dt>
<dd>Fulcio</dd>
<dt>Believer<sup>Fulcio</sup></dt>
<dd>Software Installer</dd>
<dt>Verifier<sup>Fulcio</sup></dt>
<dd>${OIDCIDOwner}: <i>${OIDCIDOwner} authorizes Fulcio to bind ${PubKey} to ${OIDCIdentity}. ${OIDCIDOwner} must actively look for Claims made on their behalf that they didn't knowingly authorize.</i></dd>
<dt>Arbiter<sup>Fulcio</sup></dt>
<dd>Community</dd>
</dl>

## Rekor: Identity-based signature

<dl>
<dt>Claim<sup>Rekor</sup></dt>
<dd><i>${OIDCIdentity} signs ${Hash} using the key bound by ${Certificate}</i></dd>
<dt>Statement<sup>Rekor</sup></dt>
<dd>${Hash}, X.509 certificate ${Certificate} containing ${OIDCIdentity}, and signature over ${Hash}</dd>
<dt>Claimant<sup>Rekor</sup></dt>
<dd>${OIDCIdentity}</dd>
<dt>Believer<sup>Rekor</sup></dt>
<dd>Software Installer</dd>
<dt>Verifier<sup>Rekor</sup></dt>
<dd>${OIDCIDOwner}: <i>${OIDCIdentity} signs ${Hash} using the key bound by ${Certificate}. ${OIDCIDOwner} must actively look for Claims made on their behalf that they didn't knowingly authorize.</i></dd>
<dt>Arbiter<sup>Rekor</sup></dt>
<dd>Community, identity-artifact mapping</dd>
</dl>

## Rekor: Key-based signature

<dl>
<dt>Claim<sup>Rekor</sup></dt>
<dd><i>${Key} signs ${Hash}, verifiable with ${PubKey} </i></dd>
<dt>Statement<sup>Rekor</sup></dt>
<dd>${Hash}, public key ${PubKey}, and signature over ${Hash}</dd>
<dt>Claimant<sup>Rekor</sup></dt>
<dd>${Key}</dd>
<dt>Believer<sup>Rekor</sup></dt>
<dd>Software Installer</dd>
<dt>Verifier<sup>Rekor</sup></dt>
<dd>${KeyOwner}: <i>${Key} signs ${Hash}, verifiable with ${PubKey}. ${KeyOwner} must actively look for Claims made on their behalf that they didn't knowingly authorize.</i></dd>
<dt>Arbiter<sup>Rekor</sup></dt>
<dd>Community, key-artifact mapping</dd>
</dl>

## Rekor: Provenance

<dl>
<dt>Claim<sup>Rekor</sup></dt>
<dd><i>${OIDCIdentity} signs ${Provenance} containing ${Subject}, using the key bound by ${Certificate}</i></dd>
<dt>Statement<sup>Rekor</sup></dt>
<dd>${Provenance} with ${Hash}, X.509 certificate ${Certificate} containing ${OIDCIdentity}, signature over ${Subject}</dd>
<dt>Claimant<sup>Rekor</sup></dt>
<dd>${OIDCIdentity}</dd>
<dt>Believer<sup>Rekor</sup></dt>
<dd>Software Installer</dd>
<dt>Verifier<sup>Rekor</sup></dt>
<dd>${OIDCIDOwner}/Artifact Builder: <i>${OIDCIdentity} signs ${Provenance} containing ${Subject}, using the key bound by ${Certificate}. ${OIDCIDOwner} or Artifact Builder must actively look for Claims made on their behalf that they didn't knowingly authorize.</i></dd>
<dt>Arbiter<sup>Rekor</sup></dt>
<dd>Community, identity-artifact mapping</dd>
</dl>

## Rekor: Timestamping

Note that this claim is a work in progress, as uploading signed timestamps is not yet supported in Rekor. See the [timestamping readme](rekor/timestamping/README.md) for more information.

<dl>
<dt>Claim<sup>Rekor</sup></dt>
<dd><i>Claim<sup>Rekor<sup>Identity</sup></sup> occurs at ${Timestamp}</i></dd>
<dt>Statement<sup>Rekor</sup></dt>
<dd>Signed ${Timestamp} over Statement<sup>Rekor<sup>Identity</sup></sup></dd>
<dt>Claimant<sup>Rekor</sup></dt>
<dd>${TSA}</dd>
<dt>Believer<sup>Rekor</sup></dt>
<dd>Software Installer</dd>
<dt>Verifier<sup>Rekor</sup></dt>
<dd>None: <i>Claim<sup>Rekor<sup>Identity</sup></sup> occurs at ${Timestamp}</i></dd>
<dt>Arbiter<sup>Rekor</sup></dt>
<dd>Community</dd>
</dl>

## Timestamp Authority

<dl>
<dt>Claim<sup>TSA</sup></dt>
<dd><i>${TimestampAuthority} claims a monotonically increasing ${Time}</i></dd>
<dt>Statement<sup>TSA</sup></dt>
<dd>Signed timestamp containing ${Time}</dd>
<dt>Claimant<sup>TSA</sup></dt>
<dd>${TimestampAuthority}</dd>
<dt>Believer<sup>TSA</sup></dt>
<dd>Software Installer, entity consuming short-lived code-signing certificate</dd>
<dt>Verifier<sup>TSA</sup></dt>
<dd>${TimestampMonotonicVerifier}: <i>${TimestampAuthority} claims a monotonically increasing ${Time}</i></dd>
<dt>Arbiter<sup>TSA</sup></dt>
<dd>Community</dd>
</dl>