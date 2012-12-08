import sys
import os.path
import bottle
from bottle import Bottle, run, template, static_file, request

app = Bottle()
bottle.debug(True)

ip_constants = {
    "hyperboria": "hype",
    "clearnet4": "4",
    "clearnet6": "6",
}

def make_template_variables(tmpl_path):
    ip = dict(ip_constants)
    ip['server'] = request.environ['Server-Addr']

    return {
        'ip':ip,
    }

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
        return template(tmpl_path, make_template_variables(tmpl_path))
    else:
        return static_file(filename, root="root")

application = app
