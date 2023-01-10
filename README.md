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
**Flags**
* env - If set to "dev", uses .env file in project root.
* text - If set to false, prints messages to logs instead of sending.
* time - Time in UTC for when the program will check and send messages for birthdays each day.

Copy `message_gen.json` and `birthday_list.json` to your desired location and edit them to your liking. Note that all enviroment variables listed above are required to function. Afterward, run the program. By default, the program checks for birthdays and sends messages at 00:00 UTC.