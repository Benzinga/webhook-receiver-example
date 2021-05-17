from flask import abort, Flask, request, Response
from urllib import unquote_plus
import json
import re


app = Flask(__name__)

url = '/placeholder/webhook/v1'

# puts the payload into an empty map
def parse_request(req):
    payload = req.get_data()
    payload = unquote_plus(payload)
    payload = re.sub('payload=', '', payload)
    payload = json.loads(payload)

    return payload

@app.route(url, methods=['POST'])
def respond():
    req_id = request.headers.get('X-BZ-Delivery')

    if request.json is None:
        # handle error here
        print("error"+req_id)
        abort(400, id + ": Does not contain a body")

    parse_request(request)
    return Response(status=200)


if __name__ == '__main__':
    app.run(debug=True, use_reloader=True)
