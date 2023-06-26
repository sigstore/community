<dl>
<dt>Claim<sup>Fulcio</sup></dt>
<dd><i>${OIDCIDOwner} authorizes ${CA} to bind ${PubKey} to ${OIDCIdentity}</i></dd>
<dt>Statement<sup>Fulcio</sup></dt>
<dd>X.509 certificate containing ${PubKey} and ${OIDCIdentity}, signed by ${CA}</dd>
<dt>Claimant<sup>Fulcio</sup></dt>
<dd>${CA}</dd>
<dt>Believer<sup>Fulcio</sup></dt>
<dd>Software Installer</dd>
<dt>Verifier<sup>Fulcio</sup></dt>
<dd>${OIDCIDOwner}: <i>${OIDCIDOwner} authorizes ${CA} to bind ${PubKey} to ${OIDCIdentity}</i></dd>
<dt>Arbiter<sup>Fulcio</sup></dt>
<dd>Community</dd>
</dl>
<dl>
<dt>Claim<sup>LOG_Fulcio</sup></dt>
<dd><i><ol><li>This data structure is append-only from any previous version</li><li>This data structure is globally consistent</li><li>This data structure contains only leaves of type `X.509 certificate containing ${PubKey} and ${OIDCIdentity}, signed by ${CA}`</li></ol></i></dd>
<dt>Statement<sup>LOG_Fulcio</sup></dt>
<dd>Log Checkpoint</dd>
<dt>Claimant<sup>LOG_Fulcio</sup></dt>
<dd>Log Operator</dd>
<dt>Believer<sup>LOG_Fulcio</sup></dt>
<dd><ul><li>Software Installer</li><li>${OIDCIDOwner}</li></ul></dd>
<dt>Verifier<sup>LOG_Fulcio</sup></dt>
<dd><ul><li>Witness: <i>This data structure is append-only from any previous version</i></li><li>Witness Quorum: <i>This data structure is globally consistent</i></li><li>${OIDCIDOwner}: <i>This data structure contains only leaves of type `X.509 certificate containing ${PubKey} and ${OIDCIdentity}, signed by ${CA}`</i></li></ul></dd>
<dt>Arbiter<sup>LOG_Fulcio</sup></dt>
<dd>Community</dd>
</dl>