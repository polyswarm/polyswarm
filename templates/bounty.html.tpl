{{define "bounty"}}
<html>
  <head>
    <title>New Bounty</title>
    {{template "includes" .}}
  </head>
  <body>
    {{template "navbar" .}}
    <div class="container">
      <h1>New Bounty</h1>
      <form enctype="multipart/form-data" action="/bounty" method="POST">
        <div class="form-group">
          <label for="artifact">Artifact:</label>
          <input type="file" class="form-control" name="artifact" />
        </div>
        <div class="form-group">
          <label for="fee">Fee:</label>
          <input type="number" class="form-control" name="fee" />
        </div>
        <div class="form-group">
          <label for="bountyAmount">Bounty Amount:</label>
          <input type="number" class="form-control" name="bountyAmount" />
        </div>
        <div class="form-group">
          <label for="blocksFromNow">Deadline (Blocks from now):</label>
          <input type="number" class="form-control" name="blocksFromNow" />
        </div>

        <button type="submit" class="btn btn-default"> Submit</button>
      </form>
    </div>
  </body>
</html>
{{end}}
