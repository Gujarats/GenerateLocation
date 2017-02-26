# Generate Random Location.
This Library is creating random location from given location. the result will be in decimal degree.
for example if you have a base location like 48.8588377,2.2775176. You can generate new location using this library.

# How To Use It
Call ``SetupLocation(Latitude,Longitude)`` first. and then you can get the random location using
`RandomLatLong(yourInputInSecondsRange)` and `RandomLatLongMinute(InputYourMinuteRange)`.

# How does it work
First it will convert the decimal degree to degre like `48° 51' 23.8104''` and then adding the random value given to seconds and minute depends on what method you are using. After that It will convert back to decimal and you will get new latitude and longitude value.

# Contribution
I'm open to any improvement on the code.
