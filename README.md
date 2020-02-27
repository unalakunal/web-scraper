# web-scraper
Simple web scraper in GoLang to find all links inside an HTML

## Usage

    NAME:
    web-scraper - finds all the links inside HTML

    USAGE:
    web-scraper [global options] command [command options] [arguments...]

    VERSION:
    0.0.2

    AUTHOR:
    Ünal Akünal <unal.akunal@gmail.com>

    COMMANDS:
        help, h  Shows a list of commands or help for one command

    GLOBAL OPTIONS:
    --url value, -u value  the URL to get data from (default: "http://google.com")
    --help, -h             show help
    --version, -v          print the version

## Concurrency

To parse multiple links in parallel, do:

`./web-scraper -u http://github.com -u http://facebook.com`