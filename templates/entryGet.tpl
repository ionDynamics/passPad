{{template "header" .}}
<table>
	{{range .Entries}}
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
		{{with .Url}}<td>
			<div class="clipboard-trigger" x-data-tbc="{{.}}">
				<a href="{{.}}">Link</a>
			</div>
		</td>{{end}}
	</tr>
	{{end}}
</table>
<form action="/v1/vault/{{.Identifier}}" method="post" autocomplete="off">
	<input type="text" name="form-name" required placeholder="Name">
	<input type="text" name="form-user" placeholder="User">
	<input type="text" name="form-pass" value="{{64 | rand}}">
	<input type="text" name="form-url" placeholder="URL">
	<button type="submit">Datensatz eintragen</button>
</form>
{{template "footer" .}}