# aftersales-backend


## Initialize your project

```
go mod init environment-config
```

## Updating dependencies

```
go mod vendor
```

## verify dependencies

```
go mod verify
```

## remove dependencies that is not used

```
go mod tidy
```

## print dependence diagram

```
go mod graph
```

## download dependencies

```
go mod download
```

# Release Template

### Feature

* [NEW] Support for Go Modules [#17](https://github.com/sillyhatxu/convenient-utils/issues/17)

---

### Bug fix

* [FIX] Truncate Latency precision in long running request [#17](https://github.com/sillyhatxu/convenient-utils/issues/17)
