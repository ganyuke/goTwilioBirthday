# goTwlioBirthday

Automatically text your friends when it's their birthday using Twilio. Supports randomly-selected prewritten phrases to always keep your birthday texts fresh!

**Requires a paid Twilio plan to send to numbers which have not whitelisted your Twilio number yet.**

## Usage
**Enviroment variables**
* TWILIO_ACCOUNT_SID - the SID for your Twilio
* TWILIO_AUTH_TOKEN - the auth token for your Twilio
* SENDER_NUMBER - the number of your Twilio account
* BIRTHDAY_LIST_PATH - the path to `birthday_list.json`
* MESSAGE_TEMPLATE_PATH - the path to `message_gen.json`
* APP_ENV - If set to "production", stops check for .env file.

Copy `message_gen.json` and `birthday_list.json` to your desired location and edit them to your liking. Afterward, run the program. Every day at 10:30 UTC, the program will check the birthday list to see if there are any birthdays today.