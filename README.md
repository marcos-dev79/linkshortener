# LinkShortener
## An URL shortener written in GO

This was built using Docker-Compose. In order to run in your machine, just clone the repository and run:
* sudo docker-compose up --build (first time, then you go up without the --build, which is much faster)

If you don't have docker installed in your computer, please download and install it at:
https://www.docker.com/

### Instructions

To shorten an URL, just http POST to http://localhost:8040/api/shorten with the following json payload (example):

```json
{
    "url":"https://www.vultr.com/docs/create-a-crud-application-with-golang-and-mongodb-on-ubuntu-20-04/?utm_source=performance-max-latam&utm_medium=paidmedia&obility_id=17096555207&utm_adgroup=&utm_campaign=&utm_term=&utm_content=&gclid=CjwKCAjwk_WVBhBZEiwAUHQCme_6kOgaeQWOKjalscslO99kCatxV5FJFdtFbqGv1127YkYBURCQ0BoCHnMQAvD_BwE"
}
```

The api will return the following:

```json
{
    "id": "2775fc20bfdacc37822f7c8e1849b5bc567abfd6",
    "shortened_url": "http://localhost:8040/2J2VhO"
}
```

Then, just throw the shortened_url value in your browser and you will get the URL redirection.

## Counter

Every time you access a shortened URL, the app will compute it and persist into the database. Just access:

http://localhost:8040/api/counter/2J2VhO ( Use your URL hash )

It will show how many times the URL was accessed.

## Delete

Send an http DELETE call to:
http://localhost:8040/api/delete/2J2VhO ( Use your URL hash )

## Documentation

I'm using godoc to generate a documentation server. Just hit 
http://localhost:6060/pkg/ and http://localhost:6060/pkg/linkshortener

After the Docker finish building the containers - If you build it first time it takes some minutes.

## Automated tests

Due to lack of time I didnt add to much tests but at least I put the mechanism in place and 1 test for avaliation purposes.