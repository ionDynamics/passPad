{{template "header"}}

<ol class="breadcrumb">
  <li class="active">Meine Tresore</li>
</ol>

<div class="row">
	<div class="col-xs-12">

		<h2>Meine Tresore</h2>

		<table class="table table-bordered table-hover table-striped">
			<tr>
				<th>Tresorname</th>
				<th>Beschreibung</th>
			</tr>
			{{range .}}
			<tr>
				<td>
					<div><a href="/v1/vault/{{.Identifier}}">{{.Title}}</a></div>
				</td>
				<td>
					<div>{{.Description}}</div>
				</td>
			</tr>
			{{end}}
		</table>
	</div>
</div>

<div class="row">
	<div class="col-xs-12 col-sm-6 col-sm-offset-3">
		<div class="well">
			<h3>Tresor hinzufügen:</h3>
			<form action="/v1/vault" method="post">
				<div class="form-group">
					<input type="text" class="form-control" name="form-title" placeholder="Tresor-Name" required>
				</div>
				<div class="form-group">
					<input type="text" class="form-control" name="form-description" placeholder="Tresor-Beschreibung">
				</div>
				<div class="form-group">
					<button class="btn btn-primary" type="submit">Neuer Tresor</button>
				</div>
			</form>
		</div>
	</div>
</div>

{{template "footer" .}}