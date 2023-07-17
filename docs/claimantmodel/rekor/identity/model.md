<!--- This content generated with:
go run github.com/google/trillian/docs/claimantmodel/experimental/cmd/render@master --domain_model_file ./docs/claimantmodel/rekor/identity/model.yaml 
-->
<dl>
<dt>Claim<sup>Rekor</sup></dt>
<dd><i>${OIDCIdentity} signs ${Hash} using the key bound by ${Certificate}</i></dd>
<dt>Statement<sup>Rekor</sup></dt>
<dd>${Hash}, X.509 certificate ${Certificate} containing ${OIDCIdentity}, and signature over ${Hash}</dd>
<dt>Claimant<sup>Rekor</sup></dt>
<dd>${OIDCIDOwner}</dd>
<dt>Believer<sup>Rekor</sup></dt>
<dd>Software Installer</dd>
<dt>Verifier<sup>Rekor</sup></dt>
<dd>${OIDCIDOwner}: <i>${OIDCIdentity} signs ${Hash} using the key bound by ${Certificate}</i></dd>
<dt>Arbiter<sup>Rekor</sup></dt>
<dd>Community, identity-artifact mapping</dd>
</dl>