{{define "assertion"}}
<html>
  <head>
    <title>New Assertion</title>
    {{template "includes" .}}
  </head>
  <body>
    {{template "navbar" .}}
    <div class="container">
      <h1>New Assertion</h1>
      <form action="/assertion" method="POST">
        <div class="form-group">
          <label for="guid">Guid:</label>
          <input type="text" class="form-control" name="guid" />
        </div>
        <div class="form-group">
          <label for="malicious">Is Malicious:</label>
          <input type="checkbox" class="form-control" name="malicious" value="true" />
        </div>
        <div class="form-group">
          <label for="bid">Bid:</label>
          <input type="number" class="form-control" name="bid" />
        </div>
        <div class="form-group">
          <label for="metadata">Metadata:</label>
          <input type="text" class="form-control" name="metadata" />
        </div>

        <button type="submit" class="btn btn-default">Submit</button>
      </form>
    </div>
  </body>
</html>
{{end}}
