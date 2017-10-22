# go-csvencoder

Go util for CSV

## Installing Go

```
wget https://storage.googleapis.com/golang/go1.7.1.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.7.1.linux-amd64.tar.gz
echo "export PATH=$PATH:/usr/local/go/bin" >> ~/.bashrc && source ~/.bashrc
mkdir ~/go && echo "export GOPATH=~/go" >> ~/.bashrc && source ~/.bashrc

```

## How to use the golib

```

import (
	"fmt"
	"io"
	"log"

	csv "github.com/anaisghelfi/go-csvencoder/csv"
)


// Browse your array of Interfaces
// w is the writer (io.Write / io.Pipe)
	for _, obj := range objects {
        enc := csv.NewEncoder(w, 3, ",")
		if err := enc.Encode(v); err != nil {
			return err
		}
	}

```
