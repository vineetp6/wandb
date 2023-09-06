// package gowandb implements the go Weights & Biases SDK
package gowandb

type History map[string]interface{}

func NewSession(opts ...SessionOption) (*Session, error) {
	session := &Session{}
	for _, opt := range opts {
		opt(session)
	}
	session.start()
	return session, nil
}