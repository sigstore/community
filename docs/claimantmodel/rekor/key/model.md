<!--- This content generated with:
go run github.com/google/trillian/docs/claimantmodel/experimental/cmd/render@master --domain_model_file ./docs/claimantmodel/rekor/key/model.yaml 
-->
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
<dd>${KeyOwner}: <i>${Key} signs ${Hash}, verifiable with ${PubKey}</i></dd>
<dt>Arbiter<sup>Rekor</sup></dt>
<dd>Community, key-artifact mapping</dd>
</dl>