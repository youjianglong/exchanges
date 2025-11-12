package common

type Auth struct {
	ObjectID   string
	ApiKey     string
	SecretKey  string
	Passphrase string
}

func NewAuth(objectID, apiKey, secretKey, passphrase string) Auth {
	return Auth{
		ObjectID:   objectID,
		ApiKey:     apiKey,
		SecretKey:  secretKey,
		Passphrase: passphrase,
	}
}

func (a Auth) Signature(method, path, body string, isUnix bool) *Signature {
	return &Signature{
		Key:    a.SecretKey,
		Method: method,
		Path:   path,
		Body:   body,
		IsUnix: isUnix,
	}
}
