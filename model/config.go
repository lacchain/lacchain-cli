package model

type OrionConfig struct {
	Nodeurl string `toml:"nodeurl"`
    Nodeport int `toml:"nodeport"`
	Nodenetworkinterface string `toml:"nodenetworkinterface"`
	Clienturl string `toml:"clienturl"`
	Clientport int `toml:"clientport"`
	Publickeys []string `toml:"publickeys"`
	Privatekeys []string `toml:"privatekeys"`
	Passwords string `toml:"passwords"`
	Othernodes []string `toml:"othernodes"`
	Tls string `toml:"tls"`

	Tlsservercert string `toml:"tlsservercert"`
	Tlsserverchain []string `toml:"tlsserverchain"`
	Tlsserverkey string `toml:"tlsserverkey"`

	Tlsclientcert string `toml:"tlsclientcert"`
	Tlsclientchain []string `toml:"tlsclientchain"`
	Tlsclientkey string `toml:"tlsclientkey"`
 
	Tlsservertrust string `toml:"tlsservertrust"`
	Tlsclienttrust string `toml:"tlsclienttrust"`
}