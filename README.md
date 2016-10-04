# Go Replacer

It is a simple library that replace strings between Start Text and End Text.

You can use to replace environment data into your strings for example, to replace a map of string data or to replace using your own custom function by parameter.

# CI

[![Build Status](https://travis-ci.org/prsolucoes/goreplacer.svg?branch=master)](https://travis-ci.org/prsolucoes/goreplacer)

# Examples

Example 1:

```golang
text := "My name is @|name1|@, your name is @|name2|@ and we are in @|country|@"
expected := "My name is Paulo Coutinho, your name is Guest and we are in Brasil"
data := map[string]string{
    "name1":   "Paulo Coutinho",
    "name2":   "Guest",
    "country": "Brasil",
}

result := ReplaceByMapOfString(text, "@|", "|@", data)

// result: My name is Paulo Coutinho, your name is Guest and we are in Brasil
```

Example 2:
```golang
text := "The home var is: ENV{MY_HOME}"
expected := "The home var is: /tmp"
os.Setenv("MY_HOME", "/tmp")

result, _ := ReplaceFunc(text, "ENV{", "}", os.Getenv)

// result: The home var is: /tmp
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
