# wait

A small improvement on Go's `sync.WaitGroup`.

## quickstart

Write this:

```golang
var wg wait.Group
wg.Run(func() {
  // ... do work
})
wg.Wait()
```

Instead of this:

```golang
var wg sync.WaitGroup
wg.Add(1)
go func() {
  defer wg.Done()
  // ... do work
}()
wg.Wait()
```

## cancellable group

This package also comes with a cancellable wait group:

```golang
var wg wait.GroupWithCancellation
wg.Run(func(cancel <-chan struct{}) {
  for {
      select {
      case work <-input:
          // ... do work
      case <-cancel:
          // ... done!
      }
  }
})
wg.Cancel()
wg.Wait()
```