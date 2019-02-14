package Global

import (
	"../../Helper/Session"
)

var Sessions *Session.Manager

func init() {
	Sessions, _ = Session.NewManager("memory", "gosessionid", 3600)
}
