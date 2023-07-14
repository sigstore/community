<!--- This content generated with:
go run github.com/google/trillian/docs/claimantmodel/experimental/cmd/render@master --full_model_file ./docs/claimantmodel/rekor/identity/full.yaml 
-->
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
<dd>${OIDCIDOwner}: <i>${OIDCIdentity} signs ${Hash} using the key bound by ${Certificate}</i></dd>
<dt>Arbiter<sup>Rekor</sup></dt>
<dd>Community, identity-artifact mapping</dd>
</dl>
<dl>
<dt>Claim<sup>LOG_Rekor</sup></dt>
<dd><i><ol><li>This data structure is append-only from any previous version</li><li>This data structure is globally consistent</li><li>This data structure contains only leaves of type `${Hash}, X.509 certificate ${Certificate} containing ${OIDCIdentity}, and signature over ${Hash}`</li></ol></i></dd>
<dt>Statement<sup>LOG_Rekor</sup></dt>
<dd>Log Checkpoint</dd>
<dt>Claimant<sup>LOG_Rekor</sup></dt>
<dd>Log Operator</dd>
<dt>Believer<sup>LOG_Rekor</sup></dt>
<dd><ul><li>Software Installer</li><li>${OIDCIDOwner}</li></ul></dd>
<dt>Verifier<sup>LOG_Rekor</sup></dt>
<dd><ul><li>Witness: <i>This data structure is append-only from any previous version</i></li><li>Witness Quorum: <i>This data structure is globally consistent</i></li><li>${OIDCIDOwner}: <i>This data structure contains only leaves of type `${Hash}, X.509 certificate ${Certificate} containing ${OIDCIdentity}, and signature over ${Hash}`</i></li></ul></dd>
<dt>Arbiter<sup>LOG_Rekor</sup></dt>
<dd>Community, identity-artifact mapping</dd>
</dl>