package bytestrings

import (
	"bytes"
	"io"
)

func Buffer(rawString string) *bytes.Buffer {
	rawBytes := []byte(rawString)
	// 아래 3개는 모두 같은 버퍼를 생성하는 방법이다.

	// 1.
	var b = new(bytes.Buffer)
	b.Write(rawBytes)

	// 2.
	b = bytes.NewBuffer(rawBytes)

	// 3.
	b = bytes.NewBufferString(rawString)
	return b
}

func toString(r io.Reader) (string, error) {
	b, err := io.ReadAll(r)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
