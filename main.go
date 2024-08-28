package main

func main() {
	server := newApiServer(":8001")
	server.Run()
}
