from flask import abort, Flask, request, Response
import json
import re
# will use flask to setup server for our example

app = Flask(__name__)

# insert URL you have provided to Benzinga
url = '/placeholder/webhook/v1'

@app.route(url, methods=['POST'])
def respond():
    req_id = request.headers.get('X-BZ-Delivery')

    if request.json is None:
        # handle error here
        print("error"+req_id)
        abort(400, req_id + ": Does not contain a body")
    # Show me the data!
    print(request.json)
    #
    return Response(status=200)


if __name__ == '__main__':
    app.run(debug=True, use_reloader=True)