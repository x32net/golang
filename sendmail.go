/* Feb 20 2016 at 10:20
There is only one way to send an e-mail message without directly communicating with an SMTP server: delegate this action to some other program.
What program to pick, is an open and wide question in itself because there are lots of programs which are able to send mails.
On the other hand, if we talk about POSIX systems such as GNU/Linux- or *BSD-based operating systems, they usually come with the binary /usr/sbin/sendmail or, sometimes, /usr/bin/sendmail which is either provided by the real Sendmail package or by another MTA which exist in abundance—ranging from full-blown MTAs such as Postfix or Exim to narrow-scoped "null-clients" such as ssmtp or nullmailer.
This program is usually called with the -t command-line option which makes it read the addresses of the message recipients from the To headers of the mail message itself.
Basically calling /usr/sbin/sendmail -t and piping the full message's text to its standard input is what PHP's mail() function does, FWIW.
While you can do this call directly using the os/exec, net/mail and net/textproto standard packages, the popular gomail package provides a simpler way to do that: its Message type provides the WriteTo() method whith is able to write to a running Sendmail instance, like this (copied from a real program of mine):
*/
const sendmail = "/usr/sbin/sendmail"

func submitMail(m *gomail.Message) (err error) {
    cmd := exec.Command(sendmail, "-t")
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr

    pw, err := cmd.StdinPipe()
    if err != nil {
        return
    }

    err = cmd.Start()
    if err != nil {
        return
    }

    var errs [3]error
    _, errs[0] = m.WriteTo(pw)
    errs[1] = pw.Close()
    errs[2] = cmd.Wait()
    for _, err = range errs {
        if err != nil {
            return
        }
    }
    return
}
/*
Actually, relying on MTA to deliver your mails has an advantage is that full-blown MTAs support mail queuing: that is, if an MTA is unable to send your message right away (say, due to a network outage etc) is will save the message in a special directory and will then periodically try delivering it again and again until that succeeds or a (usually huge, like 4-5 days) timeout expires.
*/
