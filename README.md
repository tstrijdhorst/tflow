# TFlow
A terminal centric workflow for Jira / Github written in Go

# Features

* Create/checkout git branch based on jira issue
* Create github PR based on jira issue
* Merge PRs, refresh master and optionally set issue to Done

# Limitations

* Only works for Jira Cloud
* Only one Jira instance can be configured globally (multiple git repositories work fine since `gh` is doing the heavy lifting there)

# Dependencies

* Git >= 2.23
* [Github CLI](https://github.com/cli/cli) (configured and ready to go)

# Installation from binary

1. Download the binary release for your platform (Linux,Mac OSX or Windows on x64)
2. Put the binary somewhere in your $PATH to easily reach it via terminal
3. Call the binary without any arguments and it will interactively generate a configfile for you
4. ???
5. Profit
