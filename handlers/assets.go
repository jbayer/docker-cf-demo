package handlers

import "text/template"

var styledTemplate = template.Must(template.New("experiment").Parse(`
<html>
<head>
<style>

body {
    font-family: "helveticaneue-light";
    font-size: 16px;
    color: #333;
    position: absolute;
    margin:0;
    width:100%;
    height:100%;
}

.hello {
  position:absolute;
  top:0;
  height:120px;
  left:0;
  right:0;
  text-align:center;
  font-size:60px;
  font-weight:bold;
  line-height:120px;
  color: rgb(0,139,185);
}

.goodbye {
  position:absolute;
  top:0;
  height:120px;
  left:0;
  right:0;
  text-align:center;
  font-size:60px;
  font-weight:bold;
  line-height:120px;
  background-color: #f00;
  color: white;
}

.my-index {
  position:absolute;
  top:120px;
  height:60px;
  font-size:30px;
  line-height:30px;
  left:0;
  right:0;
  text-align:center;
  color: rgb(0,139,185);
}

.my-index-goodbye {
  position:absolute;
  top:120px;
  height:60px;
  font-size:30px;
  line-height:30px;
  left:0;
  right:0;
  text-align:center;
  background-color: #f00;
  color: white;
}

.index {
  position:absolute;
  top:180px;
  height:120px;
  left:0;
  right:0;
  color: #fff;
  font-size: 30px;
  line-height: 120px;
  background-color: rgb(36,184,235);
  text-align:center;
}

.index-goodbye {
  position:absolute;
  top:180px;
  height:120px;
  left:0;
  right:0;
  color: #fff;
  font-size: 30px;
  line-height: 120px;
  background-color: rgb(235, 13, 5);
  text-align:center;
}

.mid-color {
  position:absolute;
  top:300px;
  height:120px;
  left:0;
  right:0;
  color: #fff;
  font-size: 30px;
  line-height: 120px;
  background-color: rgb(36,184,235);
  text-align: center;
}

.mid-color-goodbye {
  position:absolute;
  top:300px;
  height:120px;
  left:0;
  right:0;
  color: #fff;
  font-size: 30px;
  line-height: 120px;
  background-color: rgb(206, 26, 4);
  text-align: center;
}

.bottom-color {
  position:absolute;
  top:420px;
  bottom:0;
  left:0;
  right:0;
  text-align:center;
  background-color: rgb(0,139,185);
  color: white;
}

.bottom-color-goodbye {
  position:absolute;
  top:420px;
  bottom:0;
  left:0;
  right:0;
  text-align:center;
  background-color: rgb(185, 4, 9);
  color: white;
}

dt {
  color:#777;
}

.envs {
  margin:10px
}

</style>
</head>
<body class="{{.Class}}">
{{.Body}}
</body>
</html>
`))

type Body struct {
	Body  string
	Class string
}
