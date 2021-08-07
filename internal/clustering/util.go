package clustering

import "encoding/gob"

// RegisterGobNames regiters all the necessary interfaces used by the clustering package
func RegisterGobNames() {
	gob.RegisterName("ConfigUpdateMessage", &ConfigUpdateMessage{})
	gob.RegisterName("StatusUpdateMessage", &StatusUpdateMessage{})
}
