# Validator

## installation
```bash
$ go get github.com/Miraddo/validator
```

## How to use!

```go
package main 

import(
    "github.com/miraddo/validator"
)


func main(){
    if validator.IsImageUrl("https://www.google.com/images/branding/googlelogo/1x/googlelogo_color_272x92dp.png"){
        println("return true")
        return 
    }

    println("return false")
}

```

### List of Validators
```
IsLowercase
IsUppercase
IsNumber
IsLetter
IsLetterNumber
IsLetterNumberLine
HasInjectCharacters
IsHTMLTag

IsEmptyString
IsEmptyArray
IsEmptyMap
IsEmptyInterface
IsEmptyNumber
IsEmptyFloat

IsEqual
IsNotEqual
IsEqualType

IsTimeEmpty
IsTimeAfter
IsTimeBefore
IsTimeEqual
IsTimeAfterNow
IsTimeBeforeNow
IsTimeEqualNow
IsTimeAfterNowUTC
IsTimeBeforeNowUTC
IsTimeEqualNowUTC


IsEmail
IsURL
IsURLStatusOK


IsValidIP
IsValidIPv4
IsValidIPv6
IsValidTCPPort
IsValidUDPPort
IsValidDomain

CheckStatusOK
CheckStatusNotFound
CheckStatusUnauthorized
CheckStatusForbidden
CheckStatusBadRequest
CheckStatusInternalServerError
CheckStatusServiceUnavailable
CheckStatusGatewayTimeout
CheckStatusHTTPVersionNotSupported

```
