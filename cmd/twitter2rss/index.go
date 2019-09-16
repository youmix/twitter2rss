package main

import "github.com/n0madic/twitter2rss"

func init() {
	twitter2rss.Index = `<!DOCTYPE html>
	<html>
	<head>
		<title>Twitter to RSS proxy</title>
		<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
		<meta name="keywords" content="Twitter, RSS, Atom, feed, reader, agregator, convert to, convert, to">
		<link rel="stylesheet" href="//maxcdn.bootstrapcdn.com/bootstrap/3.4.1/css/bootstrap.min.css">
		<link rel="shortcut icon" href="//abs.twimg.com/favicons/twitter.ico" type="image/x-icon">
	</head>
	<body style="padding: 20px;">
		<div class="container">
			<div class="jumbotron vertical-center">
				<div class="container">
					<h1>Twitter to RSS
						<small>proxy</small>
					</h1><br />
					<p>Enter Twitter name and get full RSS feed include images!</p>
					<form id="tform" class="form-horizontal" role="form" action="/" method="GET">
						<div class="form-group">
							<div class="input-group input-group-lg">
								<span class="input-group-addon">@</span>
								<input type="text" id="name" class="form-control search-query" placeholder="Twitter name" required>
									<span class="input-group-btn">
										<input class="btn btn-primary" type="submit" value="Get RSS">
									</span>
							</div>
							<div class="panel panel-default" style="margin-top: 20px;">
								<div class="panel-body">
									<div class="form-group" style="margin-bottom: 0;">
										<label for="pages" class="col-sm-3 control-label">Number of pages (max 10):</label>
										<div class="col-sm-2">
											<input name="pages" id="pages" class="form-control" placeholder="1">
										</div>
										<input style="margin-top: 10px;" name="exclude_replies" id="exclude_replies" type="checkbox"> Exclude Replies
									</div>
								</div>
							</div>
						</div>
					</form>
				</div>
			</div> <!-- jumbotron -->
				<script>
					var Form = document.getElementById('tform');
					Form.onsubmit = function(event) {
						document.getElementById('tform').action = '/' + document.getElementById('name').value;
						var pages_input = document.getElementById('pages');
						if (pages_input.value < 2) pages_input.disabled = true;
					};
				</script>
			<footer class="navbar-fixed-bottom">
				<div style="text-align: center;"><p><a href="https://github.com/n0madic/twitter2rss">GitHub</a> &copy; Nomadic 2014-2019</p></div>
			</footer>
		</div>
	</body>
	</html>`
}
