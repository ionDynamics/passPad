{{template "header" .}}
	<form class="form-register pp-login" method="post" action="{{.Action}}">
		<h2 class="form-register-heading">Registrieren</h2>
		<label for="input-user" class="sr-only">E-Mail</label>
		<input type="text" id="input-user" name="input-user" class="form-control" placeholder="E-Mail" required autofocus>
		<label for="input-password" class="sr-only">Passwort</label>
		<input type="password" id="input-password" name="input-password" class="form-control" placeholder="Passwort" required>
		<label for="input-password2" class="sr-only">Passwort wiederholen</label>
		<input type="password" id="input-password2" name="input-password2" class="form-control" placeholder="Passwort wiederholen" required>
		<input type="hidden" id="redirect-to" name="redirect-to" value="{{.Data.RedirectTo}}">
		<button class="btn btn-lg btn-primary btn-block" type="submit">Registrieren</button>
	</form>
{{template "footer" .}}