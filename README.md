# Tareekh
[![Build Status](https://travis-ci.org/peanut-labs/tareekh.svg?branch=master)](https://travis-ci.org/peanut-labs/tareekh)
[![GoDoc](https://godoc.org/github.com/peanut-labs/tareekh?status.svg)](https://godoc.org/github.com/peanut-labs/tareekh)

Tareekh is a GoLang library to make it a little bit easy to work with dates. This was written to avoid reimplementing common functionality working with `time.Time`. 

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

##### BeginningOfMonth
```go
begin := tareekh.BeginningOfMonth() //first of current month
```

##### EndOfMonth
```go
end := tareekh.EndOfMonth() //last day of current month
```

##### FromDayOfMonth
```go
dayx := tareekh.FromDayOfMonth(1) //specified day of current month
```


## Status

In development, if there's other useful stuff feel free to fork and issue a pull request.
