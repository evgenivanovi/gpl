package io

/* __________________________________________________ */

type OnErrorReader struct {
	err error
}

func (er *OnErrorReader) Read([]byte) (int, error) {
	return 0, er.err
}

func NewOnErrorReader(err error) *OnErrorReader {
	return &OnErrorReader{
		err: err,
	}
}

/* __________________________________________________ */
