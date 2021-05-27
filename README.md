# gogoauth
golang tool for google authenticator

## create google authenticator url
```C++
gogoauth  -create -name "xxx@xx.com" -issuer "abc.com" -s "secret" -png 
```

## or without the secret, a new secret will be created and print out
```C++
gogoauth  -create -name "xxx@xx.com" -issuer "abc.com" -png
```

## compute code with the given secret
```C++
gogoauth -s "secret"
```