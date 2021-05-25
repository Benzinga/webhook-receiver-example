from flask import Flask, request, Response
import json
import re
# will use flask to setup server for our example



app = Flask(__name__)
# Handler of the endpoint provided to Benzinga
@app.route("/", methods=['POST'])
def respond():
    # Header with ID of the event
    event_id = request.headers.get('X-BZ-Delivery')
    # Request Body will always have a JSON, if it doesn't, it's an invalid request
    if request.json is None:
        app.logger.error(f"Error in event ID: {event_id}")
        # Report back the error to us with event ID
        return app.response_class(response=f"ID: {event_id}, does not contain a valid Body", status=400)
    # Show me the data!
    app.logger.info(request.json)
    # send a 200 response back to Benzinga otherwise message will resend
    return app.response_class(status=200)
if __name__ == '__main__':
    app.run(debug=True)