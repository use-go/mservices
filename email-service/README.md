## Usage

### Setting config
```shell
micro config set email.spam.spamd_address 127.0.0.1:783
micro config set email.smtp.addr smtp.qq.com:25
micro config set email.smtp.username 2637309949@qq.com
micro config set email.smtp.identity 
micro config set email.smtp.password jhprqpetmlfteabe
micro config set email.smtp.host smtp.qq.com
```

### Install spamassassin

```shell
sudo yum update
sudo yum -y install spamassassin
```

```shell
vi /etc/mail/spamassassin/local.cf
```

Uncomment, or insert the following:

```
required_hits 5.0
report_safe 0
required_score 5
rewrite_header Subject [SPAM]
```

start the service
```shell
systemctl enable spamassassin
systemctl start spamassassin
```

Update the spam rules by running
```shell
sa-update
```

## Test

Classify email
```shell
curl -X POST \
  http://127.0.0.1:8080/email/Classify \
  -H 'content-type: application/json' \
  -d '{
	"email_body": "123",
    "html_body": "123",
    "text_body": "123",
    "email": "123"
}'
```

```json
{
    "request_id": "ac01de2a-776a-4a71-8a68-3b80d8f81cce",
    "code": 200,
    "is_spam": true,
    "score": 7.9,
    "details": [
        "NO_RELAYS, Informational: message was not relayed via SMTP, -0",
        "MISSING_HEADERS, Missing To: header, 1.2",
        "MISSING_MID, Missing Message-Id: header, 0.1",
        "MISSING_SUBJECT, Missing Subject: header, 1.8",
        "EMPTY_MESSAGE, Message appears to have no textual parts and no, 2.3",
        "MISSING_FROM, Missing From: header, 1",
        "NO_RECEIVED, Informational: message has no Received headers, -0",
        "MISSING_DATE, Missing Date: header, 1.4",
        "NO_HEADERS_MESSAGE, Message appears to be missing most RFC-822 headers, 0"
    ]
}
```


Send email
```shell
curl -X POST \
  http://127.0.0.1:8080/email/Send \
  -H 'cache-control: no-cache' \
  -d '{
	"to": "2637309949@qq.com",
	"subject": "中奖通知",
	"text_body": "123$"
}'
```

```json
{
    "request_id": "94b27929-a85e-4d6e-867b-c10c5a0eb306",
    "code": 200
}
```