![logo](https://user-images.githubusercontent.com/1564322/185397346-49485044-364a-4936-b561-3d6dcbb7b3da.svg)

A terminal centric workflow for Jira / Github written in Go

_Disclaimer: I'm mainly making this to use the Jira webinterface as little as possible. Use at your own risk. PRs are very welcome as long as you use a similar codestyle._

# Features

* Create/checkout git branch based on jira issue
* Create github PR based on jira issue
* Merge PRs, optionally refresh base and set issue to Done

# Limitations

* Only works for Jira Cloud
* Only one Jira instance can be configured globally (multiple git repositories work fine since `gh` is doing the heavy lifting there)

# Dependencies

* Go >= 1.17
* Git >= 2.23
* [Github CLI](https://github.com/cli/cli) (configured and ready to go)

# Installation from binary

1. Download the binary release for your platform (Linux, Mac OSX or Windows on x64)
2. Put the binary somewhere in your $PATH to easily reach it via terminal
3. Call the binary without any arguments and it will interactively generate a configfile for you
4. ???
5. Profit

# Building it yourself

1. Checkout this repository somewhere
2. Execute `go install`
