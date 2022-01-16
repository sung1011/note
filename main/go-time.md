# time

## code

- [time](src/go/package/time_test.go)

## TimeZone

### 默认时区

```go
func initLocal() {
	// consult $TZ to find the time zone to use.
	// no $TZ means use the system default /etc/localtime.
	// $TZ="" means use UTC.
	// $TZ="foo" means use /usr/share/zoneinfo/foo.

	tz, ok := syscall.Getenv("TZ")
	switch {
	case !ok:
		z, err := loadLocation("localtime", []string{"/etc"})
		if err == nil {
			localLoc = *z
			localLoc.name = "Local"
			return
		}
	case tz != "" && tz != "UTC":
		if z, err := loadLocation(tz, zoneSources); err == nil {
			localLoc = *z
			return
		}
	}

	// Fall back to UTC.
	localLoc.name = "UTC"
}

```

### 获取

```go
now.Location()
```

### 改变

```go
// func LoadLocation(name string) (*Location, error)

loc, _ := time.LoadLocation("Etc/GMT+4")
t := now.In(loc)
```