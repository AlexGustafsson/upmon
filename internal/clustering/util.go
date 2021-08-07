package clustering

import "encoding/gob"

func RegisterGobNames() {
	gob.RegisterName("ConfigUpdateMessage", &ConfigUpdateMessage{})
	gob.RegisterName("StatusUpdateMessage", &StatusUpdateMessage{})
}
