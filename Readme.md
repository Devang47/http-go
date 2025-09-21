# Implementation of HTTP 1.1 on top of TCP

## References

RFC 9110
RFC 9112

## HTTP-message format

```
start-line CRLF
\*( field-line CRLF )
CRLF
[ message-body ]
```
