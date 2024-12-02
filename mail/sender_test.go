package mail

import (
	"testing"

	"simple_bank/util"

	"github.com/stretchr/testify/require"
)

func TestSendEmailWithGmail(t *testing.T) {
	// THIS TEST WILL GET SKIPPED
	if testing.Short() {
		t.Skip()
	}

	config, err := util.LoadConfig("..")
	require.NoError(t, err)

	sender := NewGmailSender(config.EmailSenderName, config.EmailSenderAddress, config.EmailSenderPassword)

	subject := "A test email"
	content := `
	<h1>Hello world</h1>
	<p>This is a test message from <i><b>Deepanshu Mehra</b></i></p>
	`
	to := []string{"mehradeepanshu2102@gmail.com"}
	attachFiles := []string{"../Dockerfile"}

	err = sender.SendEmail(subject, content, to, nil, nil, attachFiles)
	require.NoError(t, err)
}
