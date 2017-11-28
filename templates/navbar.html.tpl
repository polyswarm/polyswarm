{{define "navbar"}}
<nav class="navbar navbar-inverse navbar-fixed-top">
  <div class="container">
    <div class="navbar-header">
      <button type="button" class="navbar-toggle collapsed" data-toggle="collapse" data-target="#navbar" aria-expanded="false" aria-controls="navbar">
        <span class="sr-only">Toggle navigation</span>
        <span class="icon-bar"></span>
        <span class="icon-bar"></span>
        <span class="icon-bar"></span>
      </button>
      <a class="navbar-brand" href="#">Polyswarm</a>
    </div>
    <div id="navbar" class="collapse navbar-collapse">
      <ul class="nav navbar-nav">
        <li><a href="/">Current Bounties</a></li>
        <li><a href="/bounty">New Bounty</a></li>
        <li><a href="/assertion">New Assertion</a></li>
      </ul>
    </div><!--/.nav-collapse -->
  </div>
</nav>
{{end}}
