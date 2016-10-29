package color

import "io"

type colorWriter struct {
	DefaultColor Color
	Color        Color
	Writer       io.Writer
}

func (c *colorWriter)Write(data []byte) (int, error) {
	b := make([]byte, 0)
	b = append(b, []byte(c.Color)...)
	b = append(b, data...)
	b = append(b, []byte(c.DefaultColor)...)
	return c.Writer.Write(b)
}

func Write(w io.Writer, color Color) io.Writer {
	return &colorWriter{
		DefaultColor:Default,
		Color:color,
		Writer:w,
	}
}
