import os
from datetime import datetime
from flask import Flask, request, flash, url_for, redirect, \
     render_template, abort, send_from_directory
from flask.ext.login import login_required, login_user, logout_user, LoginManager, current_user
     
def klima():
    if 'GET' == request.method:
        return render_template('main.html', grouphead='afrika.html', content='afrika-map.html', mapoverlay='afrika-klima.html')
    else:
        if '1' == request.form['afgnr']:
            if 'posted' == request.form['status']:
                score = 0
                if 'subtrop' == request.form['cat1']:
                    score += 1
                if 'passat' == request.form['cat2']:
                    score += 1
                if 'wechsel' == request.form['cat3']:
                    score += 1
                if 'aequatorial' == request.form['cat4']:
                    score += 1
                return render_template('afrika-klima-afg1.html', score = score, cat1 = request.form['cat1'], cat2 = request.form['cat2'], cat3 = request.form['cat3'], cat4 = request.form['cat4'])
            else:
                return render_template('afrika-klima-afg1.html', score = 5, cat1 = 'none', cat2 = 'none', cat3 = 'none', cat4 = 'none')
        else:
            return render_template('main.html', grouphead='afrika.html', content='afrika-map.html', mapoverlay='afrika-klima.html')

def vegetation():
    if 'GET' == request.method:
        return render_template('main.html', grouphead='afrika.html', content='afrika-map.html', mapoverlay='afrika-vegetation.html')
    else:
        if '1' == request.form['afgnr']:
            if 'posted' == request.form['status']:
                score = 0
                if 'hartlaub' == request.form['cat1']:
                    score += 1
                if 'wuste' == request.form['cat2']:
                    score += 1
                if 'dornenstrauch' == request.form['cat3']:
                    score += 1
                if 'trockensavanne' == request.form['cat4']:
                    score += 1
                if 'feuchtsavanne' == request.form['cat5']:
                    score += 1
                if 'tropR' == request.form['cat6']:
                    score += 1
                return render_template('afrika-vegetation-afg1.html', score = score, cat1 = request.form['cat1'], cat2 = request.form['cat2'], cat3 = request.form['cat3'], cat4 = request.form['cat4'], cat5 = request.form['cat5'], cat6 = request.form['cat6'])
            else:
                return render_template('afrika-vegetation-afg1.html', score = 7, cat1 = 'none', cat2 = 'none', cat3 = 'none', cat4 = 'none', cat5 = 'none', cat6 = 'none')
        else:
            return render_template('main.html', grouphead='afrika.html', content='afrika-map.html', mapoverlay='afrika-vegetation.html')

def info():
    if 'GET' == request.method:
        searchentry = request.args.get('q', '')
        return None
    else:
        # POST
        searchentry = request.form['q']
        return render_template('info_'+searchentry+'.html')
