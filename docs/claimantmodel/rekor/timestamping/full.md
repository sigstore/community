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
<dd>None</dd>
<dt>Arbiter<sup>Rekor</sup></dt>
<dd>Community</dd>
</dl>
<dl>
<dt>Claim<sup>LOG_Rekor</sup></dt>
<dd><i><ol><li>This data structure is append-only from any previous version</li><li>This data structure is globally consistent</li><li>This data structure contains only leaves of type `Signed ${Timestamp} over Statement<sup>Rekor<sup>Identity</sup></sup>`</li></ol></i></dd>
<dt>Statement<sup>LOG_Rekor</sup></dt>
<dd>Log Checkpoint</dd>
<dt>Claimant<sup>LOG_Rekor</sup></dt>
<dd>Log Operator</dd>
<dt>Believer<sup>LOG_Rekor</sup></dt>
<dd><ul><li>Software Installer</li></ul></dd>
<dt>Verifier<sup>LOG_Rekor</sup></dt>
<dd><ul><li>Witness: <i>This data structure is append-only from any previous version</i></li><li>Witness Quorum: <i>This data structure is globally consistent</i></li></ul></dd>
<dt>Arbiter<sup>LOG_Rekor</sup></dt>
<dd>Community</dd>
</dl>