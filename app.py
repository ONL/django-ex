# -*- coding: utf-8 -*-
import os
from datetime import datetime
from flask import Flask, request, flash, url_for, redirect, \
     render_template, abort, send_from_directory
from flask.ext.login import login_required, fresh_login_required, login_user, logout_user, LoginManager, current_user
from kartenapp import userclass, afrika

app = Flask(__name__)
app.config.from_pyfile('flaskapp.cfg')

# Login Manager Initialisierung
login_manager = LoginManager()
login_manager.login_view = "to_login"
login_manager.refresh_view = "to_login"
login_manager.init_app(app)

# Daten für current_user wiederherstellen
@login_manager.user_loader
def load_user(userid):
    appUser = userclass.User()
    appUser.get_user(userid)
    return appUser

# Login
@app.route('/login/', methods=['GET', 'POST'])
def to_login():
    if 'GET' == request.method:
        return render_template('login.html', next = request.args.get('next') or '')
    else:
        if request.form['pw']:
            appUser = userclass.User()
            appUser.try_login(request.form['pw'])
            login_user(appUser, True)
            return redirect(request.form['next'] or url_for("index"))
        else:
            return redirect(url_for('to_login'))

# Logout
@app.route("/logout/")
@login_required
def to_logout():
    logout_user()
    return redirect(url_for('index'))

@app.route('/')
def index():
    return render_template('main.html', content='index.html', grouphead='indexhead.html')

@app.route('/about/')
def about():
    return render_template('main.html', content='about.html', grouphead='indexhead.html')

@app.route('/quellen/')
def quellen():
    return render_template('main.html', content='quellen.html', grouphead='indexhead.html')

@app.route('/<path:resource>')
def serveStaticResource(resource):
    return send_from_directory('static/', resource)

@app.route("/afrika-klima/", methods=['GET', 'POST'])
def afrikaKlima():
    return afrika.klima()

@app.route("/afrika-vegetation/", methods=['GET', 'POST'])
def afrikaVegetation():
    return afrika.vegetation()
    
@app.route("/afrika-klima-los/", methods=['GET'])
@login_required
def afrikaKlimaLos():
    return afrika.klima()

@app.route("/afrika-vegetation-los/", methods=['GET'])
@login_required
def afrikaVegetationLos():
    return afrika.vegetation()

@app.route("/afrika/", methods=['POST'])
def afrikaInfo():
    return afrika.info()

if __name__ == '__main__':
    app.run()
