# Tareekh

Tareekh is a GoLang library to make it a little bit easy to work with dates. This is highly specific to PL related stuff, so there's a lot of assumptions made about default date format etc. (YYYY-MM-DD).

This was originally written to help with writing Alexa custom skills.

## Time zone
By default the Local timezone is used. But if you want to specify a particular timezone, just do:

```go
tareekh.TimeZone = "America/New_York" //EST
```

## Examples

##### Yesterday

```go
yesterday := tareekh.Yesterday()
fmt.Println(yesterday) //yesterdays time object
```

##### FromDateString

```go
dt, err := tareekh.FromDateString("2016-11-22")
if err != nil {
  //handle error
}
fmt.Println(t)
```

##### ToShortDate

```go
t, _ := tareekh.FromDateString("2016-11-22")
shortDate := tareekh.ToShortDate(t)
fmt.Println(shortDate) //should print 2016-11-22
```


##### IsDateInFuture
```go
future := time.Now().AddDate(0, 0, 1)
isit := tareekh.IsDateInFuture(future)
fmt.Println(isit) //should be true
```

## Status

In development, if there's other useful stuff feel free to fork and issue a pull request.
