# BlogoMatic

This is the monorepo of the proejct OpenAI Wordpress.
The goal of the project is to automatically publish
posts on Wordpress using OpenAI to generate these posts.
This project will also generate the short description,
title and other details according to the post it generated.
The AI will also find the ideas of the posts itself.

## Configuration

- Configure the prompts, the feed URLs and etc. in
[`go-backend/config.yml`](go-backend/config.yml)
- Configure the secrets such as OpenAI token in
[`go-backend/.env`](go-backend/.env)

**IMPORTANT NOTE:** If you are having issues saying that
authorization header is missing, try setting Wordpress
environment type to local. I did not check the side-effects
of this yet, so please check it if you want to use this in
production. [Check here for more information](https://developer.wordpress.org/apis/wp-config-php/#wp-environment-type)

## Usage

1. Clone this repository and open a shell
1. Configure the project
1. Install the dependencies by changing the directory
into `go-backend` (`cd go-backend`) and running `go mod tidy`.
1. Run the go-backend using `go run .`
1. Now, check your Wordpress website to check if the post is
published.

## To-Do List

- [ ] Select the best topic using OpenAI or select random
post from feeds to write an article about instead of choosing
the first one.
- [ ] Make it user-friendly to be used with a Web UI.
  - [ ] Maybe a cloud version that will allow users to
  use the software without installing.
- [ ] Add support for other blog softwares, maybe other LLMs.
