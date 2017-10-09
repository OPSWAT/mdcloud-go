package aws

import (
	"github.com/OPSWAT/mdcloud-go/utils"
	"github.com/aws/aws-sdk-go/aws/session"
)

// Session variable related to the current AWS account
var (
	Session *session.Session
	err     error
)

// LoadProfile ensures we have a session
func LoadProfile() {
	Session, err = session.NewSession()
	if err != nil {
		utils.ExitError("error loading session", err)
	}
	// defaults to us-west-2
	if *Session.Config.Region == "" {
		Session.Config = Session.Config.WithRegion("us-west-2")
	}
}
