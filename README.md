# SnakeSync

Simple File Backup Tool

## Getting Started

### Requirements

* Go >= 1.15

## CLI Usage

### Scan
* How to use the program via scan

```
go run ./cmd -scan 
```

### With Flag

* How to run the program

```
go run ./cmd -path=<Path> 
```

- *path*  
Flag: -path  
Type: string  


## API Usage
* How to import database

```go

package main

import (
  "errors"
  snakesync "github.com/sadihakan/snake-sync"
)

func main() {

  ss := snakesync.New()

  //Create Watcher
  ss.NewWatcher()

  //Add file path
  ss.AddFilePath(path)

  done := make(chan bool)

  if ss.Error != nil {
    panic(ss.Error)
  }

  ss.Chase(done)

  <-done
}
```

Contributors names and contact info

ex. [@sadihakan](https://github.com/sadihakan/)

## License

This project is licensed under the sadihakan License - see the LICENSE.md file for details


# snake-sync
