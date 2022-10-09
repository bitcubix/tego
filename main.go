package main

import "github.com/bitcubix/tego/cli"

func main() {
	template_path, err := cli.Load()
	if err != nil {
		panic(err)
	}

	println(template_path)
}
