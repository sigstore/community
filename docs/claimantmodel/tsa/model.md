<!--- This content generated with:
go run github.com/google/trillian/docs/claimantmodel/experimental/cmd/render@master --domain_model_file ./docs/claimantmodel/tsa/model.yaml 
-->
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