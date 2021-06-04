<!-- DO NOT REMOVE - contributor_list:data:start:["gleich"]:end -->

# cihat

🥧 View the status of repo checks from an RPi sense hat LED matrix

![build](https://github.com/gleich/cihat/workflows/build/badge.svg)
![test](https://github.com/gleich/cihat/workflows/test/badge.svg)
![lint](https://github.com/gleich/cihat/workflows/lint/badge.svg)

## ❓ Explanation

cihat is a go based application for viewing the [commit checks](https://docs.github.com/en/free-pro-team@latest/github/collaborating-with-issues-and-pull-requests/about-status-checks) of GitHub repos you have contributor access to on a [Raspberry Pi Sense Hat](https://www.raspberrypi.org/products/sense-hat/). Each row is a repo and they go from top to bottom based on when a commit was made to the default branch (recent on top). Repos that have commits with no checks are skipped. Each pixel displays a check based on the status of a check. Green is for checks that pass, yellow is for checks that are in progress, and red is for checks that fail. Gets the latest checks every 2 seconds.

## 👀 Demo

[Click this to see the demo!](https://www.youtube.com/watch?v=9989GZIfGQk)

## 🚀 Setup

Please run all the following steps on your RPi with the sense hat plugged in.

### 1. Install docker and docker-compose

You can install both by running `sh ./install-docker.sh` from inside this repo.

### 2. Creating the PAT (Personal Access Token)

[Create a new personal access token](https://github.com/settings/tokens/new) and check off the `repo:status` box. Copy that token and put it in a file in `~/cihat-config/pat.txt`.

### 3. Start it up!

Run `docker-compose up -d` from inside this repo. All done!

### 4. Shutting it down

You can stop cihat by running `docker-compose down` from inside this repo and just restart your RPi to turn off the lights.

## 🙌 Contributing

Before contributing please read the [CONTRIBUTING.md file](https://github.com/gleich/cihat/blob/master/CONTRIBUTING.md).

<!-- DO NOT REMOVE - contributor_list:start -->

## 👥 Contributors

- **[@gleich](https://github.com/gleich)**

<!-- DO NOT REMOVE - contributor_list:end -->
