from __future__ import print_function
import json
import os
import sys
from flask import Flask, render_template, redirect, request, url_for, make_response, jsonify
from flask_restful import Resource, Api, reqparse
from raspberry import m_RGB_LED

app = Flask(__name__)
color = "green"

@app.route('/')
def index():
    content = make_response(render_template('index.html'))
    return content


@app.route('/_ajaxAutoRefresh', methods= ['GET'])
def stuff():
    return jsonify(color=color)


@app.route('/api/v1/device/register',methods=['POST'])
def register():
    request.get_json(force=True)

    parser = reqparse.RequestParser()
    parser.add_argument('id', required=True)
    args = parser.parse_args()

    id = args['id']

    print("registering device: ", id)

    returnData = "Device registered"

    return returnData, 201

@app.route('/api/v1/device/<id>/changeColor',methods=['GET'])
def changeColorGet(id):
    global color
    print('This is standard output')
    print('This is standard output', file=sys.stdout)
    print("remote addr", request.remote_addr,file=sys.stdout)
    print("requesting device: ", id)

    return color, 200

@app.route('/api/v1/device/<id>/changeColor',methods=['PUT'])
def changeColorPut(id):
    global color

    data = request.get_json(force=True)
    print("changeColorPut, request body:",data)

    parser = reqparse.RequestParser()
    parser.add_argument('color', required=True)
    args = parser.parse_args()
    reqColor = (args['color'])
    color = reqColor
    print("changeColorPut, request device:",id,"oldColor:",reqColor,"newColor",color)

    if color == "red":
      m_RGB_LED.setColor((255,0,0))
    elif color == "green":
      m_RGB_LED.setColor((0,255,0))
    elif color == "blue":
      m_RGB_LED.setColor((0,0,255))
    else:
      # https://www.rapidtables.com/web/color/Yellow_Color.html
      m_RGB_LED.setColor((255,255,0))

    returnData = "color change accepted, current: " + color

    return returnData, 201

if __name__ == "__main__":
	app.run(    debug=True, \
                host='0.0.0.0', \
                port=int(os.getenv('PORT', '5000')), threaded=True)
