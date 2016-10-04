# Go Replacer

It is a simple library that replace strings using callback function and search between Start Text and End Text to do it.  

# CI

[![Build Status](https://travis-ci.org/prsolucoes/goreplacer.svg?branch=master)](https://travis-ci.org/prsolucoes/goreplacer)

# Advantages

1. You can replace dinamically a MAP of strings inside your string  
2. You can replace using existing functions in Go like os.Getenv and replace any environment variable in your string
3. You can replace using your own custom functions that search on database, make a HTTP request or something else before replace
4. One line of code and everything works: goreplacer.*

# Examples

Example 1:

```golang
text := "My name is @|name1|@, your name is @|name2|@ and we are in @|country|@"
data := map[string]string{
    "name1":   "Paulo Coutinho",
    "name2":   "Guest",
    "country": "Brasil",
}

result := replacer.ReplaceByMapOfString(text, "@|", "|@", data)

// result: My name is Paulo Coutinho, your name is Guest and we are in Brasil
```

Example 2:
```golang
text := "The home var is: ENV{MY_HOME}"
result := replacer.ReplaceByEnvVars(text, "ENV{", "}")

// result: The home var is: /Users/paulo
```

You can check all examples in "goreplacer_test.go"

# Support with donation
[![Support with donation](http://donation.pcoutinho.com/images/donate-button.png)](http://donation.pcoutinho.com/)

# Supported By Jetbrains IntelliJ IDEA

![alt text](https://github.com/prsolucoes/goreplacer/raw/master/extras/jetbrains/logo.png "Supported By Jetbrains IntelliJ IDEA")

# Author WebSite

> http://www.pcoutinho.com

# License

MIT
