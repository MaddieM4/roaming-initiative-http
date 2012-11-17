import sys
import bottle
from bottle import Bottle, run, template, static_file

app = Bottle()

@app.route('/')
def index():
    return open('root/index.html')

@app.route('/<filename:path>')
def static(filename):
    return static_file(filename, root="root")

reloader = False
if "--debug" in sys.argv:
    print "Running in debug mode"
    bottle.debug(True)
    #reloader = True

run(app, host='localhost', port=8000, reloader=reloader)
