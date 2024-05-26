import requests

api_url = 'http://45.140.188.54:6200/solve'
api_key = 'your api key'
uid = '123456789'

payload = {
    'apiKey': api_key,
    'uid': uid
}

try:
    response = requests.post(api_url, json=payload)
    response.raise_for_status()  # Check if the request was successful
    print('Response:', response.json())
except requests.exceptions.RequestException as error:
    print('Error:', error)
