# HTTX

*small, sharp tools for HTTP.*


## Introduction

HTTX is an attempt to make a super simple, meta framework -- similar to [NextJS](https://nextjs.org/), [SvelteKit](https://kit.svelte.dev/), [SolidStart](https://start.solidjs.com) or [Remix](https://remix.run/) -- that is built on Go and Bash. HTTX is a library of [Tools](/todo) for both serving web content as websites and testing your websites. HTTX is built on top of normal file servers like `npx serve` and uses file base routing.

## Motivation

HTTX is built on the philosophy that moving all of our problems into more higher level abstracts does not always help us most effectively deliver content to our users. This is a philosophy heavily inspired by [HTMX](https://htmx.org) -- even down to the name ;)

Today's frameworks work more and more on moving logic into interpreted languages -- usually transpiled, interpreted languages like Typescript. While this is a fun effort and moves all components of a system into one runtime with one lingua franca, we often ignore the reality that we are making more complex problems for ourselves.

Some projects like [Bun](https://bun.sh/) are focusing on taking the parts of the app that we rely on, like http servers, file io, and memory management to more modern languages -- namely Rust. We appreciate this effort but we want to remove one more contraint, we don't want to build a toolkit for Javascript.

## How It Works

Javascript, and more specifically NodeJs, is a powerful runtime that support much of the web today -- both frontend and backend -- and you should not be rewriting your business logic to Bash. HTTX offers a solutution for taking the insights from your services -- NodeJs or otherwise -- and embedding that data into html templates.

The processing of a request revolves around the Context. This is a datastructure used as a state machine to hold information for the html template that should be rendered.

Once, the `render` -- or `redir` and `to-json` -- command is called the Context will be consumed into a Response object. This Response object will then be interpreted and formatted into an http response using Go's [net/http](https://pkg.go.dev/net/http) package.

## Installation

```
go install github.com/fly-high-bird/httx@latest
```
