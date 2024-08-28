# challenge.zaehlerfreunde.com

## Prerequisites

 1 - Golang (v 1.22.0 or later) on the Backend

 2 -  Node.js version 8.9 or above (v10+ recommended) to run Vuejs on the Frontend


## Instructions

We want to be able to calculate the energy cost ( Market price and price / KWH) of meter readings in a specific time range.

#### Backend
We want to create a REST endpoint that takes a one or more meter readings and returns the energy cost.

In order to get the market prices we will use a specific REST API GET endpoint:
- https://api.awattar.de/v1/marketdata?start=[START-TIMESTAMP]&end=[END-TIMESTAMP]

- Timestamps should be in milliseconds (i.e 1724857124000).This website https://www.epochconverter.com/  can be helpful.

- the following meter readings can be used as an example: 

```
{
    readings:[
        {
            timeStamp:1724770724000,
            value : 2000.0 //consumption in in kwh
        }
    ]
}
```


#### Frontend
On the Vue Web application side , we want to have :

- table with the following columns: 
From (Date), To (Date), Consumption (KWH).This table will show the added meter readings with their timestamps and the consumption.

- 1 Datepicker field to add a timestamp 
- 1 Input field to add the consumption in KWH of a reading
- One button that adds a reading to the table 
- One button that calculates the cost of the meter readings inside the table.
(this Button should be integrated with the Backend to get the cost and show it on the frontend)

