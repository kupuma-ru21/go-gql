package model

import (
	"fmt"
	"io"
	"net/url"
)

type MyURL struct {
	url.URL
}

// MarshalGQL implements the graphql.Marshaler interface
func (u MyURL) MarshalGQL(w io.Writer) {
	io.WriteString(w, fmt.Sprintf(`"%s"`, u.URL.String()))
}

// UnmarshalGQL implements the graphql.Unmarshaler interface
func (u *MyURL) UnmarshalGQL(v interface{}) error {
	switch v := v.(type) {
	case string:
		if result, err := url.Parse(v); err != nil {
			return err
		} else {
			u = &MyURL{*result}
		}
		return nil
	case []byte:
		result := &url.URL{}
		if err := result.UnmarshalBinary(v); err != nil {
			return err
		}
		u = &MyURL{*result}
		return nil
	default:
		return fmt.Errorf("%T is not a url.URL", v)
	}
}
