package main

import (
	"github.com/SpeedPHP/zeroswagger/pkg/zeroswagger"
)

func main() {
	zeroswagger.Generator("greet.api", "goc.json")
}
