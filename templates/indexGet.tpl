{{template "header" .}}
<table>
	{{range .}}
	<tr>
		<td>
			<div><a href="/v1/vault/{{.Identifier}}">{{.Description}}</a></div>
		</td>
	</tr>
	{{end}}
</table>
<form action="/v1/vault" method="post">
	<input type="text" name="form-description" required>
	<button type="submit">Neuer Tresor</button>
</form>
{{template "footer" .}}