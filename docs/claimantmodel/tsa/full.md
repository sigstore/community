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
<dl>
<dt>Claim<sup>LOG_TSA</sup></dt>
<dd><i><ol><li>This data structure is append-only from any previous version</li><li>This data structure is globally consistent</li><li>This data structure contains only leaves of type `Signed timestamp containing ${Time}`</li></ol></i></dd>
<dt>Statement<sup>LOG_TSA</sup></dt>
<dd>Log Checkpoint</dd>
<dt>Claimant<sup>LOG_TSA</sup></dt>
<dd>Log Operator</dd>
<dt>Believer<sup>LOG_TSA</sup></dt>
<dd><ul><li>Software Installer, entity consuming short-lived code-signing certificate</li><li>${TimestampMonotonicVerifier}</li></ul></dd>
<dt>Verifier<sup>LOG_TSA</sup></dt>
<dd><ul><li>Witness: <i>This data structure is append-only from any previous version</i></li><li>Witness Quorum: <i>This data structure is globally consistent</i></li><li>${TimestampMonotonicVerifier}: <i>This data structure contains only leaves of type `Signed timestamp containing ${Time}`</i></li></ul></dd>
<dt>Arbiter<sup>LOG_TSA</sup></dt>
<dd>Community</dd>
</dl>