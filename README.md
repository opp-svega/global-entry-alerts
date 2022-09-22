# Global Entry Alerts

## Problem
It is very hard to get an appointment to complete your Global Entry Applications. This tool will monitor the Global Entry System for available appointments in the few selected locations.

## Setup
### Twilio Setup
Twilio has a very generous free trial in which you dont even need a credit card to sign up. Once you have signed up, you can fill out the following environment variables.
```
# From Twilio
export TWILIO_FROM_PHONE_NUMBER=""
export TWILIO_ACCOUNT_SID=""
export TWILIO_AUTH_TOKEN=""

# Mobile number you wish to recieve text messages to
export TWILIO_TO_PHONE_NUMBER=""
```

## How to run
Just running the application
```
cd src
go run main.go
```

Building an then running
```
cd src
go build 
```