package baseview

import "html/template"

var tpl *template.Template
var tplErr error

func init() {
	tpl, tplErr = template.New("base").Parse(base)
	tpl.New("css").Parse(css)
	tpl.New("header").Parse(header)
	tpl.New("footer").Parse(footer)
	tpl.New("js").Parse(js)
	tpl.New("content").Parse(content)
}

type Page struct {
	Title string
	Data  map[string]interface{}
}

// GetBaseTheme provides a pointer to the template containing
func GetBaseTheme() (*template.Template, error) {
	return tpl, tplErr
}

var base = `
<html>
<head>
<title>{{.Title}}</title>
{{block "css" .}}{{end}}
</head>
<body>
{{block "header" .}}{{end}}
{{template "content" .}}
{{block "footer" .}}{{end}}
{{block "js" .}}{{end}}
</body>
</html>
`
var header = `
{{define "header"}}
<header>
<div class="logo">soFileServe</div>
<nav>
    <a href="/">Root</a>
    <a href="" id="upLevel">Up Level</a>
</nav>
</header>
{{end}}
`
var content = `
{{define "content"}}
<form method="post" action="/">
<div class="table">
    <div class="table-row">
        <div class="table-cell"><input id="select-all" type="checkbox"/></div>
        <div class="table-cell">Name</div>
        <div class="table-cell">Type</div>
        <div class="table-cell">Size</div>
        <input type="hidden" name="zip" id="zip" value=0 />
        <input class="zip-download-button" type="submit" value="Zip Selected &amp; Download"/>&nbsp;
        <!-- <a class="download-button">Download Selected</a> -->
    </div>
{{range .Data.items}}
    <div class="table-row">
        <div class="table-cell"><input class="checkbox" name="filePath" type="checkbox" value="{{.Path}}" /></div>
        <div class="table-cell label"><a href="{{.Path}}">{{.Name}}</a></div>
        <div class="table-cell label">{{.Type}}</div>
        <div class="table-cell label">{{.Size}}</div>
        <input class="download-button" type="button" data-path="{{.Path}}" value="Download"/>
    </div>
{{end}}
    <input class="zip-download-button" type="submit" value="Zip Selected &amp; Download"/>&nbsp;
</div>
</form>
{{end}}
`
var footer = `
{{define "footer"}}
<footer>Copyright &copy; <a href="https://plainmotif.co.uk">Plainmotif Ltd.</a></footer>
{{end}}
`
var js = `
{{define "js"}}
<script>
document.getElementById('select-all').addEventListener('click', selectAll);
document.getElementById('upLevel').href = getUpLevel()

zipButtons = document.getElementsByClassName('zip-download-button');
buttons = document.getElementsByClassName('download-button');

for (var i = zipButtons.length - 1; i >= 0; i--) {
    zipButtons[i].addEventListener('click', zipFlagTrue);
}

for (var i = buttons.length - 1; i >= 0; i--) {
    buttons[i].addEventListener('click', function() {
        zipFlagFalse();
        console.log(this.getAttribute("data-path"));
    });
}


function zipFlagTrue()
{
    document.getElementById('zip').value = 1
}

function zipFlagFalse()
{
    document.getElementById('zip').value = 0
}

/**
 * @todo make sure there is no 'checked' mismatch on the selectAll checkbox
 * if first is checked and then selectAll is called
 */
function selectAll()
{
    var checkboxes = document.getElementsByClassName('checkbox');
    var state = true;

    if (checkboxes[0].checked == true) {
        state = false
    }

    for (i=0; i < checkboxes.length; i++) {
        checkboxes[i].checked = state;
    }
}

function getUpLevel()
{
    var url = window.location.href;

    if (url.substr(-1) == '/') url = url.substr(0, url.length - 2);

    url = url.split('/');
    url.pop();

    return url.join('/');
}
</script>
{{end}}
`
var css = `
{{define "css"}}
<style>
div {padding: 5px;}
.table {display:table;text-align: center;}
.table-row {display: table-row;}
.table-cell {display: table-cell;}
.charts {}
.clear {clear:both;}
.my-chart {float:left;max-width: 30%;padding:0 1% 20px;}
</style>
{{end}}
`
