# Zardo's Captcha Solver

This is an automated non-browser based recaptcha/cloudfare solver for the famous discord bot poketwo. It uses a private service for solving the captchas and logs them in a secure place. 

## Features
- Fully Automated Captcha Solving.
- Under 2mins solving speed on server provided proxies.
- Under 50s solving speed on webshare proxies.
- Under 20s solving speed on iproyal proxies.
- Embedded with a log checker for you.
- Newbie friendly.
- Requires very less coding knowledge.
- Very affordable.

## Available Endpoints
- **/solve** - To solve the captcha.
- **/credits** - To check available Balance of a user.
- **/logs** - To get json of logs of a user.
- **/logger** - To get a web-based GUI for checking logs.

## Snippets (Will be updated in the future)
Server-Side Logs:
![image](https://github.com/sleepyzardo/ZardoCaptchaSolver/assets/88527682/5c5a601e-1993-4056-8d9a-e80ec610abc9)

Logs Checker:
![image](https://github.com/sleepyzardo/captcha_solver/assets/88527682/7ed4ce03-65ab-4eae-b0b7-d58a97e8d931)

Client-Side Snippet (Server-Provided Proxies):
![image](https://github.com/sleepyzardo/ZardoCaptchaSolver/assets/88527682/1d71ab12-af3d-4ac6-b5b1-105a75527f79)

## Server Responses
1. **500** - Internal Server error has occured. This can mean anything from my solving server failing to your api key having some issues. Further diagnostics will be done if such problem occurs.
2. **404** - The requested endpoint was not found.
3. **204** or **200** - Indicates a successful response
4. **400** - Forbidden. You already have the requested UID in the queue.

## Understanding Responses
- On any successful request with the solved being true, 1 credit will be deducted. You will get a json object with timeTaken, solved and creditsUsed, UID, and proxy if any.
- On any successful request with the solved being false, No credits will be deducted. You can request again to solve the captcha. If it still fails, Please contact an admin.
- 500 errors are very rare but not impossible, if they do occur please contact an admin. Compensation will be provided incase the error persists.
  
# Note
This is not a free api due to the cost of managing it and hence there is a very affordable price to it. If you wish to know more contact me at discord on `@sleepyzardo`.


**This is only for educational purposes, Demonstrating that it is indeed possible. I've no connections to any 3rd party that may try to misuse it.**




*For any suggestions you may have in mind. Please contact me at `@sleepyzardo`*
