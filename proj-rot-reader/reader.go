package main


import (
	"fmt"
	"io"
	"strings"
)


const Message = "Let's go, Bear #17!"


// A `Rot13Reader` wrapping over another reader, modifying the stream
// by:
//   - 'A-Z': rotating 5 letters rightward
//   - 'a-z': rotating 7 letters rightward
//   - '0-9': rotating 3 numbers rightward
//   - others: stay unchanged
type RotReader struct {
	raw io.Reader
}

// Initialization function.
func NewRotReader(raw io.Reader) RotReader {
	return RotReader{raw}
}

// Implement the `io.Reader` interface.
func (r RotReader) Read(b []byte) (int, error) {
	// Read from raw.
	braw := make([]byte, cap(b))
	n, err := r.raw.Read(braw)
	if err != nil {
		return n, err
	}
	// Do rotation on `braw[:n]`.
	for i, c := range braw[:n] {
		switch {
		case byte('A') <= c && c <= byte('Z'):
			b[i] = byte('A') + (c-byte('A')+5) % 26
		case byte('a') <= c && c <= byte('z'):
			b[i] = byte('a') + (c-byte('a')+7) % 26
		case byte('0') <= c && c <= byte('9'):
			b[i] = byte('0') + (c-byte('0')+3) % 10
		default:
			b[i] = c
		}
	}
	return n, nil
}


// Testing.
func main() {
	s := strings.NewReader(Message)
	r := NewRotReader(s)
	fmt.Printf("Initialized a RotReader on message \"%s\".\n", Message)

	fmt.Print("Ciphered text: \"")
	b := make([]byte, 4)    // Read 4 bytes each time.
	for {
		n, err := r.Read(b)
		if err != nil {
			break
		}
		fmt.Printf("%s", b[:n])
	}
	fmt.Println("\".")
}
