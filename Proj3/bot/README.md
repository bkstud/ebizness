# Slack bot
### To start app run:
`gradle run` <br>
App runs on port 3000 for synchronising with real slack server ngrok.io can be used. <br>
In my setup following commands were added in app panel: <br>
lists all categories: <br>
    /categories -> $address/slack/categories  <br>
lists specific category: <br>
    /list category -> $address/slack/categories/list <br>
Category is read as text form parameter.
