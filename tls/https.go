// openssl genrsa -out localhost.key 2048
// openssl req -new -x509 -key localhost.key -out localhost.cert -days 3650 -subj /CN=localhost
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

// Create a handler function that calls an upgrade to websocket session using our upgrader parameters
func handler(w http.ResponseWriter, r *http.Request) {
	// TODO encode a msg to: gzip, deflate, sdch.
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "Method: %s\n", r.Method)
	fmt.Fprintf(w, "Protocol: %s\n", r.Proto)
	fmt.Fprintf(w, "Host: %s\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr: %s\n", r.RemoteAddr)
	fmt.Fprintf(w, "RequestURI: %q\n", r.RequestURI)
	fmt.Fprintf(w, "URL: %#v\n", r.URL)
	fmt.Fprintf(w, "Body.ContentLength: %d (-1 means unknown)\n", r.ContentLength)
	fmt.Fprintf(w, "Close: %v (relevant for HTTP/1 only)\n", r.Close)
	fmt.Fprintf(w, "TLS: %#v\n", r.TLS)
	fmt.Fprintf(w, "\nHeaders:\n")
	r.Header.Write(w)

	io.WriteString(w, "boby: tls")
}

func main() {
	http.HandleFunc("/", handler)
	log.Printf("Listening for SECURE websocket connections on port 8000.")
	log.Fatal(http.ListenAndServeTLS(":8000", "localhost.cert", "localhost.key", nil))
}

/*
url = https://localhost:8000
*/

/* Chrome 49.0
Method: GET
Protocol: HTTP/2.0
Host: localhost:8000
RemoteAddr: [::1]:52439
RequestURI: "/"
URL: &url.URL{Scheme:"", Opaque:"", User:(*url.Userinfo)(nil), Host:"", Path:"/", RawPath:"", RawQuery:"", Fragment:""}
Body.ContentLength: 0 (-1 means unknown)
Close: false (relevant for HTTP/1 only)
TLS: &tls.ConnectionState{Version:0x303, HandshakeComplete:true, DidResume:false, CipherSuite:0xc02f, NegotiatedProtocol:"h2", NegotiatedProtocolIsMutual:true, ServerName:"localhost", PeerCertificates:[]*x509.Certificate(nil), VerifiedChains:[][]*x509.Certificate(nil), SignedCertificateTimestamps:[][]uint8(nil), OCSPResponse:[]uint8(nil), TLSUnique:[]uint8{0xc7, 0xda, 0x29, 0xa2, 0x4d, 0xad, 0x8, 0x84, 0x9e, 0x36, 0x24, 0x1e}}

Headers:
Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8
Accept-Encoding: gzip, deflate, sdch
Accept-Language: ru-RU,ru;q=0.8,en-US;q=0.6,en;q=0.4
Cache-Control: max-age=0
Upgrade-Insecure-Requests: 1
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/49.0.2623.112 Safari/537.36
boby: tls
*/

/* Safari 9.1
Method: GET
Protocol: HTTP/1.1
Host: localhost:8000
RemoteAddr: [::1]:52523
RequestURI: "/"
URL: &url.URL{Scheme:"", Opaque:"", User:(*url.Userinfo)(nil), Host:"", Path:"/", RawPath:"", RawQuery:"", Fragment:""}
Body.ContentLength: 0 (-1 means unknown)
Close: false (relevant for HTTP/1 only)
TLS: &tls.ConnectionState{Version:0x303, HandshakeComplete:true, DidResume:false, CipherSuite:0xc013, NegotiatedProtocol:"", NegotiatedProtocolIsMutual:true, ServerName:"localhost", PeerCertificates:[]*x509.Certificate(nil), VerifiedChains:[][]*x509.Certificate(nil), SignedCertificateTimestamps:[][]uint8(nil), OCSPResponse:[]uint8(nil), TLSUnique:[]uint8{0x52, 0xa3, 0xfe, 0x14, 0x85, 0xd5, 0xfe, 0x4d, 0xf5, 0xc, 0xd1, 0x7b}}

Headers:
Accept: text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8
Accept-Encoding: gzip, deflate
Accept-Language: en-us
Cache-Control: max-age=0
Connection: keep-alive
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_5) AppleWebKit/601.5.17 (KHTML, like Gecko) Version/9.1 Safari/537.86.5
boby: tls
*/
