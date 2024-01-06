# HTTX

*small, sharp tools for HTTP.*


## Introduction

HTTX is an attempt to make a super simple, meta framework -- similar to [NextJS](https://nextjs.org/), [SvelteKit](https://kit.svelte.dev/), [SolidStart](https://start.solidjs.com) or [Remix](https://remix.run/) -- that is built on Go and Bash. HTTX is a library of [Tools](/cmd) for both serving web content as websites and testing your websites. HTTX is built on top of normal file servers like `npx serve` and uses file base routing.

## Installation

```
go install github.com/fly-high-bird/httx/cmd/...
```

## Quick Start

First consider the following:

```html
<!-- ./index.html -->
<html>
  <head>...</head>
  <body>
    Hello Luke Skywalker!
  </body>
</html>
```

This will be a great new homepage, but one problem not everyone's name is Luke Skywalker. To fix this lets move the file to `./_templates/homepage.html` and make a new file:

```bash
# ./index.sh

# Set the name property to the value of the query param with the key name
with-prop "name" $QUERY_NAME |

# render the homepage template we created
render "homepage"
```

Next change the template to use [Go's HTML Templating](https://pkg.go.dev/html/template) with a map of props passed in as the global object.

```html
<!-- ./_templates/homepage.html -->
<html>
  <head>...</head>
  <body>
    Hello {{.name}}!
  </body>
</html>
```

Now we can run our server and navigat to `http://localhost:8080?name=Mario%20Mario` and see our new message.

```
# run the httx command to run a dev server at the given path
$ httx --path . --http :8080
```

**This example can be found as the [Readme Example](/examples/readme)**

## Motivation

HTTX is built on the philosophy that moving all of our problems into more higher level abstracts does not always help us most effectively deliver content to our users. This is a philosophy heavily inspired by [HTMX](https://htmx.org) -- even down to the name ;)

Today's frameworks work more and more on moving logic into interpreted languages -- usually transpiled, interpreted languages like Typescript. While this is a fun effort and moves all components of a system into one runtime with one lingua franca, we often ignore the reality that we are making more complex problems for ourselves.

Some projects like [Bun](https://bun.sh/) are focusing on taking the parts of the app that we rely on, like http servers, file io, and memory management to more modern languages -- namely Rust. We appreciate this effort but we want to remove one more contraint, we don't want to build a toolkit for Javascript.

## How It Works

Javascript, and more specifically NodeJs, is a powerful runtime that support much of the web today -- both frontend and backend -- and you should not be rewriting your business logic to Bash. HTTX offers a solutution for taking the insights from your services -- NodeJs or otherwise -- and embedding that data into html templates.

The processing of a request revolves around the Context. This is a datastructure used as a state machine to hold information for the html template that should be rendered.

Once, the `render` -- or `redir` and `to-json` -- command is called the Context will be consumed into a Response object. This Response object will then be interpreted and formatted into an http response using Go's [net/http](https://pkg.go.dev/net/http) package.

