package export

const htmlExportTemplate = `
<!DOCTYPE html>
<html>
<head>
    <meta charset='utf-8'>
    <meta http-equiv='X-UA-Compatible' content='IE=edge'>
    <title>{{.Title}}</title>
    <meta name='viewport' content='width=device-width, initial-scale=1'>
</head>
<body>
	<table>
	<tr>
		{{range .Header}}
			<th>{{.}}</th>
		{{end}}
	</tr>
	<tbody>
		{{range .Rows}}
			<tr>
				{{range .}}
					<td>{{.}}</td>
				{{end}}
			</tr>
		{{end}}
	</tbody>
	</table>
</body>
</html>
`

type templateInput struct {
	Rows   [][]string
	Title  string
	Header []string
}
