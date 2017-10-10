package aws

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

// Session variable related to the current AWS account
var (
	Session *session.Session
)

// LoadProfile ensures we have a session
func LoadProfile() {
	creds := credentials.NewEnvCredentials()
	if _, err := creds.Get(); err != nil {
		log.Fatalln(fmt.Sprintf("error loading session: %s", err.Error()))
	}
	Session = session.Must(session.NewSession())

	// defaults to us-west-2
	if *Session.Config.Region == "" {
		Session.Config = Session.Config.WithRegion("us-west-2")
	}
}
