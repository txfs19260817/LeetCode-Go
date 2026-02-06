import requests
import json

# JSON data
data = [
    {
        "request": {
            "url": "/v1/charges",
            "headers": {
                "X-Stripe-Client-User-Agent": "{\"lang\": \"ruby\", \"publisher\": \"stripe\", \"uname\": \"Linux version 3.13.0-57-generic (buildd@brownie) (gcc version 4.8.2 (Ubuntu 4.8.2-19ubuntu1) ) #95-Ubuntu SMP Fri Jun 19 09:28:15 UTC 2015\", \"hostname\": \"caron\", \"lang_version\": \"2.1.2 p95 (2014-05-08)\", \"engine\": \"ruby\", \"platform\": \"x86_64-linux\", \"bindings_version\": \"1.23.0\"}",
                "Host": "api.stripe.com",
                "Accept-Encoding": "gzip, deflate",
                "Authorization": "Bearer sk_test_xxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
                "Content-Type": "application/x-www-form-urlencoded",
                "Accept": "*/*; q=0.5, application/xml",
                "User-Agent": "Stripe/v1 RubyBindings/1.23.0"
            },
            "body": "amount=123&currency=usd&card[number]=4242424242424242&card[exp_month]=12&card[exp_year]=2020",
            "method": "POST"
        },
        "response": {
            "code": 200
        }
    },
    # Add other requests here as necessary
]

# API Base URL
base_url = "https://api.stripe.com"

# Function to issue a request and parse the response
def make_request(request_data):
    url = base_url + request_data["url"]
    headers = request_data["headers"]
    method = request_data["method"]
    body = request_data["body"]

    # Parse body into a dictionary for `data` parameter in requests
    body_dict = dict(x.split('=') for x in body.split('&'))

    # Issue the request
    if method == "POST":
        response = requests.post(url, headers=headers, data=body_dict)
    elif method == "GET":
        response = requests.get(url, headers=headers, params=body_dict)
    else:
        raise ValueError("Unsupported method")

    # Parse the response
    try:
        response_json = response.json()
        print("Response JSON:", json.dumps(response_json, indent=2))
    except json.JSONDecodeError:
        print("Failed to parse response as JSON:", response.text)

    # Additional handling based on response status code
    if response.status_code == 200:
        print("Request succeeded:", response_json)
    else:
        print("Request failed with code", response.status_code)
        if "error" in response_json:
            print("Error message:", response_json["error"]["message"])

# Process each request in the data
for item in data:
    make_request(item["request"])
