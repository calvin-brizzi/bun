package bun

// This template is compiled as a binary as mucking with file access in
// google app engine was not on my list of priorities
var helpTemplate = `
<head>
<style>
  table, th, td {
    border: solid 1px #DDEEEE;
    border-collapse: collapse;
    font: normal 13px Arial, sans-serif;
  }

	th {
		text-align: center;
		font-weight:bold;
	}

	th, td {
		padding: 10px;
	}

	table {
		width: 100%;
	}
</style>
</head>
<body>
<h1> Bun, like bunny1, but with some stuff cut off </h1>

<p> Bun is a fast redirection service, that allows you to create custom search engines very quickly, and make everyone more productive </p>

<h2> This instance of Bun has the following options: </h2>
<table>
	<tr>
		<th>
			Name
		</th>
		<th>
			Command
		</th>
		<th>
			Help
		</th>
	</tr>
	{{ range $key, $value := . }}
	<tr>
		<td>
			{{ $value.Name }}
		</td>
		<td>
			{{ $value.Key }}
		</td>
		<td>
			{{ $value.Help }}
		</td>
	{{ end }}
</table>
</body>
`
