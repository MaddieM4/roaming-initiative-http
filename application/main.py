import sys
import os.path
import bottle
from bottle import Bottle, run, template, static_file

app = Bottle()

@app.route('/')
@app.route('/index.htm')
@app.route('/index.html')
def root():
    return render('index')

@app.route('/<filename:path>')
def render(filename):
    filename = filename.rstrip("/")
    tmpl_path = 'pages/'+filename+".tpl"
    if os.path.isfile(tmpl_path):
        return template(tmpl_path)
    else:
        return static_file(filename, root="root")

application = app
