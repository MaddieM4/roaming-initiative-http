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
    tmpl_path = 'pages/'+filename+".tpl"
    if os.path.isfile(tmpl_path):
        return template(tmpl_path)
    else:
        return static_file(filename, root="root")

reloader = False
if "--debug" in sys.argv:
    print "Running in debug mode"
    bottle.debug(True)
    #reloader = True

run(app,
    host='localhost',
    port=8000,
    #server="gevent", # TODO: Figure out correct version for gevent
    reloader=reloader
)
