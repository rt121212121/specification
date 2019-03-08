
# Vote Action

Vote Action -  A vote is created by the Contract in response to a valid Referendum (Issuer) or Initiative (User) Action.

The following breaks down the construction of a Vote Action. The action is constructed by building a single string from each of the elements in order.

<table class="waffle">
	<tr style='height:19px;'>
		<th style="width:6%" class="s0">Field</th>
		<th style="width:9%" class="s1">Label</th>
		<th style="width:9%" class="s1">Name</th>
		<th style="width:2%" class="s1">Bytes</th>
		<th style="width:29%" class="s1">Example Values</th>
		<th style="width:26%" class="s1">Comments</th>
		<th style="width:5%" class="s1">Data Type</th>
		<th style="width:14%" class="s2">Amendment Restrictions</th>
	</tr>
	<tr>
		<td class="s5" rowspan="100">Metadata (OP_RETURN Payload)</td>
		<td class="g6">Header[]</td>
		<td class="g6">Header Array</td>
		<td class="g6">-</td>
		<td class="g6">-</td>
		<td class="g6">Common header data for all actions</td>
		<td class="g6">Header</td>
		<td class="g7"></td>
	</tr>

	<tr>
		<td class="g10">Timestamp</td>
		<td class="g10">Timestamp</td>
		<td class="g10">8</td>
		<td class="g10" style="word-break:break-all">1551767413250187179</td>
		<td class="g10">Timestamp in nanoseconds of when the smart contract created the action.</td>
		<td class="g10">timestamp</td>
		<td class="g11">Cannot be changed by issuer, operator. Smart contract controls.</td>
	</tr>

</table>