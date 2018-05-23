package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/sirupsen/logrus"
)

// Session variable related to the current AWS account
var (
	Session *session.Session
)

// LoadProfile ensures we have a session
func LoadProfile() {
	Session = session.Must(session.NewSession())
	if _, err := Session.Config.Credentials.Get(); err != nil {
		logrus.Fatalln("Loading credentials, ensure credentials setup under ~/.aws/credentials")
	}

	// defaults to us-west-2
	if *Session.Config.Region == "" {
		Session.Config = Session.Config.WithRegion("us-west-2")
	}
}
