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

## Credits

## License
