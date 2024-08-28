# challenge.zaehlerfreunde.com

## Prerequisites

 1 - Golang (v 1.22.0 or later) on the Backend

 2 -  Node.js version 8.9 or above (v10+ recommended) to run Vuejs on the Frontend


## Instructions

Am energy meter is a device that calculates the energy consumption of a certain household or establishment. Some energy meters can also send meter readings every 15 mins and based on the consumption data provided by the energy meter we can calculate the spot price.

A spot price in energy is the current price for buying or selling electricity, natural gas, or other energy resources right now, for immediate use or delivery. It's like the price you see at a gas station that tells you how much you would pay to fill up your tank at that moment. the formula to calculate the spot price is : ``` consumptionInKwh * marketPrice```.

 The goal of this challenge is to be able to calculate the spot price ( Market price and price / KWH) of meter readings in a specific time range.

#### Backend
We want to create a REST endpoint that takes a one or more meter readings and returns the energy cost.

In order to get the market prices we will use a specific REST API GET endpoint:
- https://api.awattar.de/v1/marketdata?start=[START-TIMESTAMP]&end=[END-TIMESTAMP]

- Timestamps should be in milliseconds (i.e 1724857124000).This website https://www.epochconverter.com/  can be helpful.

- the following meter readings can be used as an example: 

```
{
    readings:[

        // meter reading A
        {
            timestamp:1724760000000,
            value : 2000.0 //consumption in in kwh
        },

         // meter reading B
        {
            timestamp:1724760900000,
            value : 2500.0 //consumption in in kwh
        }
    ]
}
```

Please not that in order to retreive the consumptionInKwh in a specific time range (15 mins in this example) we need to : 

- get the time range by subtracting the timestamp (timestamp from meter reading B - timestamp from meter reading A  ) so in our case ```1724760900000 - 1724760000000 = 900000 ms = 15 mins```

- get the energy consumption during the time range (value from meter reading B - value from meter reading A ) so in our case ```2500.0 - 2000.0 = 500 Kwh```

so the spot price will be ```500 * marketPrice```
the ```marketPrice``` for the specific time range can be retrieved from https://api.awattar.de as mentioned earlier.


#### Frontend
On the Vue Web application side , we want to have :

- table with the following columns: 
From (Date), To (Date), Consumption (KWH).This table will show the added meter readings with their timestamps and the consumption.

- 1 Datepicker field to add a timestamp 
- 1 Input field to add the consumption in KWH of a reading
- One button that adds a reading to the table 
- One button that calculates the cost of the meter readings inside the table.
(this Button should be integrated with the Backend to get the cost and show it on the frontend)

