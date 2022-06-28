# LinkShortener
## An URL shortener written in GO

This was built using Docker-Compose. In order to run in your machine, just clone the repository and run:
* sudo docker-compose up --build

If you don't have docker installed in your computer, please download and install it at:
https://www.docker.com/

### Instructions

To shorten an URL, just http POST to http://localhost:8040/api/shorten with the following payload (example):

{
    "url":"https://twitter.com/"
}

The api will return the following:
{
    "id": "62bb89bc994c04314169f066",
    "shortened_url": "http://localhost:8040/7da105de659ce893"
}

Then, just throw the shortened_url value in your browser and you will get the URL redirection.