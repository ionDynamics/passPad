{{template "header" .}}
{{with .Png}}
	<img src="data:image/png;base64,{{.}}">
{{end}}
	<form class="form-setup pp-login" method="post" action="{{.Action}}">
		<h2 class="form-setup-heading">Token bestätigen</h2>
		<label for="input-token" class="sr-only">Token</label>
		<input type="text" id="input-token" name="input-token" class="form-control" placeholder="Token" required autofocus>
		<input type="hidden" id="redirect-to" name="redirect-to" value="{{.Data.RedirectTo}}">
		<button class="btn btn-lg btn-primary btn-block" type="submit">Bestätigen</button>
	</form>
{{template "footer" .}}