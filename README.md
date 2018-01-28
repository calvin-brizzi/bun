# Bun

Bun is a redirection service, created because [bunny1](http://www.bunny1.org/) hasn't been updated in 5 years, so I rewrote the functionality I wanted in Go.

## Overview

Bun lets you define custom redirect rules that make searching various services easy. For example, when I type `d lunch menu` it redirects to a Google Drive search
for 'lunch menu'.

### Why not just use my browser's custom search functionality?

Because that's harder to share accross browsers. Bun's main use is for creating custom searches that everyone in your organization can use, and setting everyone's search
engine to point at Bun. For example, we have our internal company wikis, various monitoring services etc all configurable in Bun, making everyone's life easier.

## Running Bun

You can either create your own main package that calls bun's init function, or (recommended) deploy on Google App Engine with the provided app.yaml (guide [here](https://cloud.google.com/appengine/docs/go/) the standard environment has served us fine so far.

Note: The app won't run out of the box, you'll either need to create an initCustom() function in custom.go or uncomment that line in bun.go (working on getting this a bit nicer. But for v0.0.1 it'll do)

## Using Bun

In chrome, you need to switch your default search engine:  
Settings > manage search engines > ADD (next to other search engines)

Then insert this info:  
```
Name: Bun
Keyword: bun (or anything you want)
URL: <bunURL>/search?q=%s
```

## Defining new commands

See the godoc for Command and simply add them to custom.go

### What about things that I don't want the outside world to have access to?

If you run Bun with `-witelist_ips true` and set `private` to true in the Command, Bun will use the ipWhitelist in `filter.go` to determine who can connect and redirect others to google for that query only.