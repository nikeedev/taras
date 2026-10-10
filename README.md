# Taras: a Gemini browser written in Go

> [!WARNING]
> This page is still under development, you can do basic "browsing", but it only fetches the websites and thats it :)


### How to use?

```console
$ go build .
$ ./taras gemini://geminiprotocol.net/
20 text/gemini
 Project Gemini

## Gemini in 100 words
...
```

Just in case you are wondering what the `20 text/gemini` means, it is the server reporting `Everything went OK,` were 20 means "OK", and then the server also specifies [MIME](https://en.wikipedia.org/wiki/MIME) code of Gemini text, aka `text/gemini`. This data is usually just for us browser developers to interpret and adjust onto the situation.


If we for example got code `31` (which means "permanent redirection" according to the [protocol specification](gemini://geminiprotocol.net/docs/protocol-specification.gmi)), we have to make sure to redirecto the the correct page so that user doesn't just get left at a blank screen.



