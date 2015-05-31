{{template "header" .}}
	<form class="form-signin pp-login" method="post" action="{{.Action}}">
		<h2 class="form-signin-heading">Anmelden</h2>
		<label for="input-user" class="sr-only">Benutzername</label>
		<input type="text" id="input-user" name="input-user" class="form-control" placeholder="Benutzername" required autofocus>
		<label for="input-password" class="sr-only">Passwort</label>
		<input type="password" id="input-password" name="input-password" class="form-control" placeholder="Passwort" required>
		<label for="input-token" class="sr-only">Token</label>
		<input type="text" id="input-token" name="input-token" class="form-control" placeholder="Token" autocomplete="off" required>
		<input type="hidden" id="redirect-to" name="redirect-to" value="{{.Data.RedirectTo}}">
		<button class="btn btn-lg btn-primary btn-block" type="submit">Anmelden</button>
	</form>
{{template "footer" .}}