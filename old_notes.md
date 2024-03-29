# Notes for general purposes

1. create models
2. create serializers
3. create views
4. add urls to router/urls in the api

## Backend Design ideas

[x] calculate total price of items on the list

[] Create calendar section

[] Categories for recipes as 'veggies', 'protein/fat', or 'starches/sugar'

[] Map the section to the calendar (drag and drop?)

[] Add to list based on calendar 'plan'

- get celery/redis to execute the task?
- generate route for aisles

## Frontend Design ideas

- IMPORTANT User feedback on how to improve!
  - send email/ prompt an 'alert' bell section asking if they liked it?
  - maybe have a star rating on the feature right at the bottom and collect metrics on user/how often pressed
- navbar with the pages (purple/black dark theme?)
- Mobile have the menu hidden in hamburger stack
- desktop have it displayed normally as a banner design
- light theme toggle button
- high contrast?
- page for todo list
- form has a "Manual Item" button to show the expanded features
- make navbar small response better for phone

## Alias to help

alias dc-bounce="docker compose down && docker compose up -d"

## Setup

Don't forget to create a superuser on django

## MVP for us

synced list
recipe add to list
route prediction (myself)

## Subscription Ideas

Free = synced lists
WARNING ON CANCELLATION - make clear upon sign up and on cancelling
-> upon cancelling all data between the date of sign up will be locked 
-> greyed out lists/features
-> add reasons why quitting to collect metrics
-> include sad cat image mayhaps?
\$ Recipe import/export/add to list
    maybe add instructions import option (long term)
\$ Route prediction including frozen/refrigerated items last option and/or effecient
-> snake route
    sort list by aislemo
-> most effecient (likely fastest)
-> cold must be last+bread on top
    -> add to the list on the right side the aisle number
\$ Calorie tracking/threshold settings
\$ Meal planning 
-> do you ACTUALLY have every meal appropriately portioned/planned
-> Visual mapping for each meal sections
-> generate list based on meal sections
\$ Recipe notebook -> collection to all recipes imported (long term maybe?)
\$ Budget planning -> allows prices to be unlocked/shown for total cost estimate (maybe included in all priced tiers)

## logic

if subscribed && tier= $tier on the components themselves

on the lists then there would be a 'premium' boolean as well as tier -> if not active premium or that same tier then it would be locked -> additional maybe for same/higher tier? -> maybe just tier at all boolean no extra
