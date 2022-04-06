# assignment - 2

## Assignment task
In this assignment, you are going to develop a REST web application in Golang that provides the client with the ability to retrieve information about Corona cases occurring in different countries, as well as the number and stringency of current policies in place. For this purpose, you will interrogate existing web services and return the result in a given output format.
The REST web services you will be using for this purpose are:


Covid 19 Cases API: https://github.com/rlindskog/covid19-graphql





Corona Policy Stringency API: https://covidtracker.bsg.ox.ac.uk/about-api



The first API focuses on the provision of information about Corona cases per country as reported by the John Hopkins Institute. The second API provides you with an assessment of policy responses addressing the corona situation.
The API documentation is provided under the corresponding links, and both services vary vastly with respect to feature set and quality of documentation. Use Postman to explore the APIs, but be mindful of rate-limiting.
A general note: When you develop your services that interrogate existing services, try to find the most efficient way of retrieving the necessary information. This generally means reducing the number of requests to these services to a minimum by using the most suitable endpoint that those APIs provide. As part of the development, and for the purpose of testing, we expect you to stub the services. e.g. make sure NOT to use the API services in your tests.
The final web service should be deployed on our local OpenStack instance Skyhigh. The initial development should occur on your local machine. For the submission, you will need to provide both a URL to the deployed service as well as your code repository.
In the following, you will find the specification for the REST API exposed to the user for interrogation/testing.

## How to run application
To run the application run "go run .\src\app\cmd\main.go" from the root forlder of the project

# Endpoints
This application have 4 different endpoints that can be used. 

## /corona/v1/cases/
This endpoint is used for getting informaion about how many covid cases there is in a country. It will use data from the day before.

### -Request
```
Method: GET
URL: serverIP/corona/v1/cases/"Country name"
```
### -Response
```
{
    "data": {
        "country": {
            "name": "Norway",
            "mostRecent": {
                "date": "2022-04-04",
                "confirmed": 1410051,
                "deaths": 2518,
                "recovered": 0,
                "growthRate": 0.0009533558409549743
            }
        }
    }
}
```
## /corona/v1/policy/
This endpoint is used for finding the different policies that are in a given country at a given date. To search you have to use the aplha-3 code of the country and date writen in YYYY-MM-DD.
### -Request
```
Method: GET
URL: servIP/corona/v1/policy/"aplha-3 code for country"/"YYYY-MM-DD"
```
### -Response
```
[
    {
        "policyActions": [
            {
                "policy_type_code": "C1",
                "policy_type_display": "School closing",
                "policyvalue": "1",
                "is_general": true,
                "notes": null
            },
        ...],
        "stringencyData": {
            "date_value": "2021-01-01",
            "country_code": "NOR",
            "confirmed": 49803,
            "stringency_actual": 56.02,
            "stringency": 56.02
        }
    }
]
```
## /corona/v1/status/
This endpoint is used for getting information about the this api and the other apis that this api uses. 

### -Request

```
Method: GET
URL: serverIP/corona/v1/status/
```

### -Response
```
{
    "CovidCasesApi": 200,
    "CovidPolicyApi": 200,
    "Webhooks": 7,
    "Version": "v1",
    "Uptime": 6.5401128
}
```
## /corona/v1/notifications/
### -Request
```
Method: GET
URL: serverIP/corona/v1/notifications/"webhook_id"(optinal)
```
- if none webhook_id is in the URL all the webhooks will be returned

```
Method: POST
URL: serverIP/corona/v1/notifications/
```

```
Method: DELETE
URL: serverIP/corona/v1/notifications/"webhook_id"
```
### -Response
GET response:
```
[
    {
        "webhookID": "FJWWE3BAXjNwmMEqlYIM",
        "URL": "https://localhost:8080/client/",
        "country": "Norway",
        "calls": 3
    }
]
```

POST response:
```
NWPmLMYHbaueTBWCpX7I
```

DELETE response:
```
```