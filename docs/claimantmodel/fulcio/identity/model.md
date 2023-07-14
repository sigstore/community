<!--- This content generated with:
go run github.com/google/trillian/docs/claimantmodel/experimental/cmd/render@master --domain_model_file ./docs/claimantmodel/fulcio/identity/model.yaml 
-->
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
<dd>${OIDCIDOwner}: <i>${OIDCIDOwner} authorizes Fulcio to bind ${PubKey} to ${OIDCIdentity}</i></dd>
<dt>Arbiter<sup>Fulcio</sup></dt>
<dd>Community</dd>
</dl>