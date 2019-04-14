package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	fmt.Println("TCP Server starts listening, on port 8080...")
	for {
		conn, err := listener.Accept() // listen to connections, so uses FOR loop to loop throught every connections.
		if err != nil {
			panic(err)
		}
		go handler(conn) // handler handles connection.
	}
}

func handler(conn net.Conn) {
	defer conn.Close()
	firstLine := true
	h := make(map[string]string)
	scan := bufio.NewScanner(conn)
	for scan.Scan() {
		ln := scan.Text()
		bs := strings.Fields(ln)
		fmt.Println(ln)
		if firstLine {
			h["method"], h["uri"], h["protocol"] = bs[0], bs[1], bs[2]
			firstLine = false
			continue
		}
		if ln == "" {
			mux(conn, h)
			break
		}
		h[bs[0]] = strings.Join(bs[1:], " ")
	}

	// print out the headers
	for i, v := range h {
		fmt.Println(i, "\t\t", v)
	}
	fmt.Println()
}

func mux(conn net.Conn, h map[string]string) {
	switch h["method"] {
	case "GET":
		switch h["uri"] {
		case "/":
			root(conn)
		case "/about":
			about(conn)
		}
	case "POST":
	case "UPDATE":
	case "DELETE":
	}
}

func root(conn net.Conn) {
	body := `<!doctype html><html>
	<head><title>root</title></head>
	<body><h3>hello world<br>
	<a href="/about">go to about</a>
	</body>
	</html>`
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func about(conn net.Conn) {
	body := `<!doctype html><html>
	<head><title>about</title></head>
	<body><h1>about page<br>
	<h3><a href="/">go to root</a>
	</body>
	</html>`
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}
