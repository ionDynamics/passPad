{{template "header" .}}

<ol class="breadcrumb">
  <li><a href="/v1/">Meine Tresore</a></li>
  <li class="active">{{.Vault.Title}}</li>
</ol>

<div class="row">
	<div class="col-xs-12">
		<h2>Meine Einträge</h2>
	</div>
</div>

<div class="row">
	<div class="col-xs-12">
		<table class="table table-bordered table-striped table-hover">
			<tr>
				<th>Name</th>
				<th>Benutzername</th>
				<th>Passwort</th>
				<th>URL</th>
			</tr>
			{{range .Vault.Entries}}
			<tr>
				<td>
					<div class="clipboard-trigger" x-data-tbc="{{.Name}}">{{.Name}}</div>
				</td>
				<td>
					<div class="clipboard-trigger" x-data-tbc="{{.User}}">{{.User}}</div>
				</td>
				<td>
					<div class="clipboard-trigger" x-data-tbc="{{.Pass}}">*******</div>
				</td>
				<td>
					{{with .Url}}
					<div class="clipboard-trigger" x-data-tbc="{{.}}">
						<a href="{{.}}">Link</a>
					</div>
					{{else}}
					-
					{{end}}
				</td>
			</tr>
			{{end}}
		</table>
	</div>
</div>

<div class="row">
	<div class="col-xs-12 col-sm-6 col-sm-offset-3">
		<div class="well">
			<h3>Eintrag hinzufügen:</h3>
			<form action="/v1/vault/{{.Vault.Identifier}}" method="post" autocomplete="off">
				<div class="form-group">
					<input type="text" class="form-control" name="form-name" required placeholder="Name">
				</div>
				<div class="form-group">
					<input type="text" class="form-control" name="form-user" placeholder="User">
				</div>
				<div class="form-group">
					<input type="text" class="form-control" name="form-pass" value="{{64 | rand}}">
				</div>
				<div class="form-group">
					<input type="text" class="form-control" name="form-url" placeholder="URL">
				</div>
				<div class="form-group">
					<button class="btn btn-primary" type="submit">Datensatz eintragen</button>
				</div>
			</form>
		</div>
	</div>
</div>

{{template "footer" .}}