# logw
Logw is a build-time log wrapper over glog and AppEngine logging. It is not a
logging framework intended to be extensible and flexible, but rather to allow
Go libraries to log easily in multiple environments for a few fixed choices of
logging mechanisms.

## Usage
For AppEngine, use the request context:
```
ctx := appengine.NewContext(r)
```

Then call any logging or assert function:
```
logw.Printf(ctx, "Message")
logw.Warningf(ctx, "Message: %v", 3)

logw.Assert(ctx, p > 5, "Value too small")
```

## License

Logw is released under the [MIT License](http://opensource.org/licenses/MIT).
