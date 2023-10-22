package headers

/* __________________________________________________ */

type HeaderKey string

func (hk HeaderKey) String() string {
	return string(hk)
}

type HeaderValue string

func (hv HeaderValue) String() string {
	return string(hv)
}

/* __________________________________________________ */
