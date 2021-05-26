package graphiql

import (
	"fmt"
	"net/http"
)

type Handler struct {
	Url string
}

func (handler *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var page = []byte(fmt.Sprintf(`
		<!DOCTYPE html>
		<html>
			<head>
				<style>
					html,
					body {
						height: 100%%;
						margin: 0;
						overflow: hidden;
						width: 100%%;
					}
					#graphiql {
						height: 100vh;
					}
				</style>

				<link
					rel="stylesheet"
					href="https://cdn.jsdelivr.net/npm/graphiql-with-extensions@0.14.3/graphiqlWithExtensions.css"
					integrity="sha384-GBqwox+q8UtVEyBLBKloN5QDlBDsQnuoSUfMeJH1ZtDiCrrk103D7Bg/WjIvl4ya"
					crossorigin="anonymous"
				/>
				<script
					src="https://cdn.jsdelivr.net/npm/whatwg-fetch@2.0.3/fetch.min.js"
					integrity="sha384-dcF7KoWRaRpjcNbVPUFgatYgAijf8DqW6NWuqLdfB5Sb4Cdbb8iHX7bHsl9YhpKa"
					crossorigin="anonymous"
				></script>
				<script
					src="https://cdn.jsdelivr.net/npm/react@16.8.6/umd/react.production.min.js"
					integrity="sha384-qn+ML/QkkJxqn4LLs1zjaKxlTg2Bl/6yU/xBTJAgxkmNGc6kMZyeskAG0a7eJBR1"
					crossorigin="anonymous"
				></script>
				<script
					src="https://cdn.jsdelivr.net/npm/react-dom@16.8.6/umd/react-dom.production.min.js"
					integrity="sha384-85IMG5rvmoDsmMeWK/qUU4kwnYXVpC+o9hoHMLi4bpNR+gMEiPLrvkZCgsr7WWgV"
					crossorigin="anonymous"
				></script>
				<script
					src="https://cdn.jsdelivr.net/npm/graphiql-with-extensions@0.14.3/graphiqlWithExtensions.min.js"
					integrity="sha384-TqI6gT2PjmSrnEOTvGHLad1U4Vm5VoyzMmcKK0C/PLCWTnwPyXhCJY6NYhC/tp19"
					crossorigin="anonymous"
				></script>
			</head>
			<body>
				<div id="graphiql"></div>
				<script>
					var fetchURL = "%s";
					function graphQLFetcher(graphQLParams) {
						var headers = {
							Accept: 'application/json',
							'Content-Type': 'application/json',
						};

						return fetch(fetchURL, {
							method: 'post',
							headers: headers,
							body: JSON.stringify(graphQLParams),
						})
							.then(function(response) {
								return response.text();
							})
							.then(function(responseBody) {
								try {
									return JSON.parse(responseBody);
								} catch (error) {
									return responseBody;
								}
							});
					}
					ReactDOM.render(
						React.createElement(GraphiQLWithExtensions.GraphiQLWithExtensions, {
							fetcher: graphQLFetcher,
						}),
						document.getElementById('graphiql'),
					);
				</script>
			</body>
		</html>
	`, handler.Url))

	w.Write(page)
}
