{{define "navigation"}}
		<nav class="navbar navbar-default navbar-fixed-top">
			<div class="container">
				<div class="navbar-header">
					<button type="button" class="navbar-toggle collapsed" data-toggle="collapse" data-target="#navbar" aria-expanded="false" aria-controls="navbar">
						<span class="sr-only">Toggle navigation</span>
						<span class="icon-bar"></span>
					</button>
					<a class="navbar-brand" href="/">PassPad</a>
				</div>
				<div id="navbar" class="collapse navbar-collapse">
					<ul class="nav navbar-nav navbar-right">
						<li class="pull-right"><a href="/v1/logout"><span class="glyphicon glyphicon-log-out"></span></a></li>
					</ul>
				</div><!--/.nav-collapse -->
			</div>
	    </nav>
{{end}}