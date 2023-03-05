# Bookmarkey API

Bookmarkey is a convenient bookmarking app that lets you save web pages and add RSS feeds. This project is the backend web server, running PocketBase. Which has been extended and used as framework with some custom business logic

## Production

You can visit the site at: https://bookmarkey.app

## Run Locally

Install [`go-task`](https://taskfile.dev/installation/), 

To setup the project locally you can do:

```bash
git@gitlab.com:bookmarkey/gui.git
cd gui
go mod tidy

# start server locally; on http://localhost:8080
task start
```

## Test Credentials

To login to the integration tests "instance" and update entries for the database do:

```bash
task start:test
# go to http://localhost:8080/_/
```

Then enter these credentials

```
username: test@bookmarkey.app
password: password11
```

## Related

Here are some related projects:

[Bookmarkey GUI](https://gitlab.com/bookmarkey/gui)
