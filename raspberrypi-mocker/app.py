from __future__ import print_function
import json
import os
import sys
from flask import Flask, render_template, redirect, request, url_for, make_response, jsonify
from flask_restful import Resource, Api, reqparse

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

@app.route('/api/v2/device/name/<id>/Color',methods=['GET'])
def changeColorGetV2(id):
    global color
    print('This is standard output')
    print('This is standard output', file=sys.stdout)
    print("remote addr", request.remote_addr,file=sys.stdout)
    print("requesting device: ", id)

    return jsonify(Color=color), 200

@app.route('/api/v2/device/name/<id>/Color',methods=['PUT','POST'])
def changeColorPutV2(id):
    global color

    data = request.get_json(force=True)
    print("changeColorPut, request body:",data)

    parser = reqparse.RequestParser()
    parser.add_argument('Color', required=True)
    args = parser.parse_args()
    reqColor = (args['Color'])
    color = reqColor
    print("changeColorPut, request device:",id,"oldColor:",reqColor,"newColor",color)


    returnData = "color change accepted, current: " + color

    return returnData, 201


if __name__ == "__main__":
	app.run(    debug=True, \
                host='0.0.0.0', \
                port=int(os.getenv('PORT', '5000')), threaded=True)
