<html>
<head>
    <title>New Bounty</title>
</head>
<body>
    <h1>New Bounty</h1>
    <form enctype="multipart/form-data" action="/bounty" method="POST">
        Artifact: <input type="file" name="artifact" /><br />
        Fee: <input type="number" name="fee" /><br />
        Bounty Amount: <input type="number" name="bountyAmount" /><br />
        Deadline (Blocks from now): <input type="number" name="blocksFromNow" /><br />
        <input type="submit" /><br />
        <input type="reset" /><br />
    </form>
</body>
</html>

