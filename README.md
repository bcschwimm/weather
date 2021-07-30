# Weather Command Line Tool | GoLang example

## Requires OpenWeather API Key & Environment Variable named "WEATHER_KEY"

## Examples: 

> weather -c 90210

> weather -c "los angeles"

#### weather metrics will be printed in fahrenheit, the terminal and we can call those using a zip code or city + state combo. 

## TODO:
1. [x] Complete custom types / methods
2. [x] Write Test
3. [x] Complile for Linux
4. Update: Handle Flag Arguments like "Los Angeles, CA". three spaces are messing with parsing the argument so we need to combine.
5. Pull in Forecast / Rain Y/N field.
6. APIs use of max/min is not the max 24h temp. Pull in another API for this