package cmdutil

import "github.com/matthiasbruns/ecwid-go/ecwid"

// AppClient is the initialized Ecwid API client, set during PersistentPreRunE.
var AppClient *ecwid.Client
