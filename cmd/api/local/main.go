package main

import "github.com/judegiordano/sst_template/internal"

func main() {
	app := internal.Server()
	app.Listen(":3000")
}
