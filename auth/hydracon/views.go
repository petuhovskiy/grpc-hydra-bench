package hydracon

const loginTemplate = `
<html>
	<body>
		<h1>Log in</h1>
		<form action="/login" method="POST">
			<input type="hidden" name="challenge" value="%s"/>
			<input type="text" id="username" name="username" placeholder="Username" />
			<br/>
			<input type="submit" value="Log in"/>
		</form>
	</body>
</html>`
