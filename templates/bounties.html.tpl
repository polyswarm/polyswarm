<html>
<head>
    <title>Active Bounties</title>
    <style>
        table, th, td {
            border: 1px solid black;
            border-collapse: collapse;
        }
    </style>
</head>
<body>
    <h1>Active Bounties</h1>
    <table>
        <tr>
            <th>Originator</th>
            <th>Bounty Fee</th>
            <th>Bounty Amount</th>
            <th>Artifact Hash</th>
            <th>Artifact URI</th>
            <th>Block Deadline</th>
            <th>Guid</th>
        </tr>
        {{range .}}
        <tr>
            <td>{{printf "%s" .Originator}}</td>
            <td>{{printf "%s" .BountyFee}}</td>
            <td>{{printf "%s" .BountyAmount}}</td>
            <td>{{printf "%s" .ArtifactHash}}</td>
            <td>{{printf "%s" .ArtifactURI}}</td>
            <td>{{printf "%s" .BlockDeadline}}</td>
            <td>{{printf "%s" .Guid}}</td>
        </tr>
        {{end}}
    </table>
<a href="/bounty">Create a New Bounty</a>
<a href="/assertion">Create a New Assertion</a>
</body>
</html>

