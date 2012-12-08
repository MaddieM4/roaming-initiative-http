import sys
import os.path
import bottle
from bottle import Bottle, run, template, static_file, request

app = Bottle()
#bottle.debug(True)

ip_constants = {
    "clearnet4": "173.255.210.202",
    "clearnet6": "2600:3c01::f03c:91ff:feae:1082",
    "hyperboria":"fcd5:7d07:2146:f18f:f937:d46e:77c9:80e7",
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
