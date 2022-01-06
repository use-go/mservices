package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"time"

	"comm/auth"
	"comm/errors"
	"comm/logger"
	"proto/email"

	"github.com/Teamwork/spamc"
	"gopkg.in/gomail.v2"
)

func validEmail(email string) bool {
	if len(email) == 0 {
		return false
	}
	m, err := regexp.MatchString("^\\S+@\\S+$", email)
	if err != nil {
		return false
	}
	return m
}

func (h *Handler) Send(ctx context.Context, request *email.SendRequest, response *email.SendResponse) error {
	acc, ok := auth.FromContext(ctx)
	if ok {
		logger.Infof(ctx, "%v Do Send", acc.Name)
	}
	if len(request.From) == 0 {
		return errors.BadRequest("Missing from name")
	}
	if !validEmail(request.To) {
		return errors.BadRequest("Invalid to address")
	}
	if len(request.Subject) == 0 {
		return errors.BadRequest("Missing subject")
	}
	if len(request.TextBody) == 0 && len(request.HtmlBody) == 0 {
		return errors.BadRequest("Missing email body")
	}

	spamReq := &email.ClassifyRequest{
		TextBody: request.TextBody,
		HtmlBody: request.HtmlBody,
		To:       request.To,
		From:     request.From,
		Subject:  request.Subject,
	}
	rsp := email.ClassifyResponse{}
	err := h.Classify(ctx, spamReq, &rsp)
	if err != nil || rsp.IsSpam {
		logger.Errorf(ctx, "Error validating email %s %v", err, rsp)
		return errors.InternalServerError("Error validating email")
	}

	if err := h.sendEmail(ctx, request); err != nil {
		logger.Errorf(ctx, "Error sending email: %v\n", err)
		return errors.InternalServerError("Error sending email")
	}

	return nil
}

func (h *Handler) Classify(ctx context.Context, request *email.ClassifyRequest, response *email.ClassifyResponse) error {
	acc, ok := auth.FromContext(ctx)
	if ok {
		logger.Infof(ctx, "%v Do Classify", acc.Name)
	}
	if len(request.EmailBody) == 0 && len(request.TextBody) == 0 && len(request.HtmlBody) == 0 {
		return errors.BadRequest("spam.Classify", "Missing one of email_body, html_body, text_body")
	}
	bf := bytes.Buffer{}

	if len(request.EmailBody) > 0 {
		bf.WriteString(request.EmailBody)
	} else {
		m := gomail.NewMessage()

		if len(request.To) > 0 {
			m.SetHeader("To", request.To)
		}
		if len(request.From) > 0 {
			m.SetHeader("From", request.From)
		}
		if len(request.Subject) > 0 {
			m.SetHeader("Subject", request.Subject)
		}
		m.SetHeader("Date", time.Now().Format(time.RFC1123Z))
		if len(request.TextBody) > 0 {
			m.SetBody("text/plain", request.TextBody)
		}
		if len(request.HtmlBody) > 0 {
			m.SetBody("text/html", request.HtmlBody)
		}
		if _, err := m.WriteTo(&bf); err != nil {
			logger.Errorf(ctx, "Error classifying email %s", err)
			return errors.InternalServerError("spam.Classify", "Error classifying email")
		}

	}
	rc, err := h.Spamc.Report(ctx, &bf, spamc.Header{}.Set("Content-Length", fmt.Sprintf("%d", bf.Len())))
	if err != nil {
		logger.Errorf(ctx, "Error checking spamd %s", err)
		return errors.InternalServerError("spam.Classify", "Error classifying email")
	}
	response.IsSpam = rc.IsSpam
	response.Score = rc.Score

	response.Details = []string{}
	for _, v := range rc.Report.Table {
		response.Details = append(response.Details, fmt.Sprintf("%s, %s, %v", v.Rule, v.Description, v.Points))
	}
	return nil
}

func (h *Handler) sendEmail(ctx context.Context, req *email.SendRequest) error {
	content := []interface{}{}
	replyTo := h.Config.EmailFrom
	if len(req.ReplyTo) > 0 {
		replyTo = req.ReplyTo
	}

	if len(req.TextBody) > 0 {
		content = append(content, map[string]string{
			"type":  "text/plain",
			"value": req.TextBody,
		})
	}

	if len(req.HtmlBody) > 0 {
		content = append(content, map[string]string{
			"type":  "text/html",
			"value": req.HtmlBody,
		})
	}

	reqMap := map[string]interface{}{
		"from": map[string]string{
			"email": h.Config.EmailFrom,
			"name":  req.From,
		},
		"reply_to": map[string]string{
			"email": replyTo,
		},
		"subject": req.Subject,
		"content": content,
		"personalizations": []interface{}{
			map[string]interface{}{
				"to": []map[string]string{
					{
						"email": req.To,
					},
				},
			},
		},
	}
	if len(h.Config.PoolName) > 0 {
		reqMap["ip_pool_name"] = h.Config.PoolName
	}

	reqBody, _ := json.Marshal(reqMap)

	httpReq, err := http.NewRequest("POST", "https://api.sendgrid.com/v3/mail/send", bytes.NewBuffer(reqBody))
	if err != nil {
		return err
	}

	httpReq.Header.Set("Authorization", "Bearer "+h.Config.Key)
	httpReq.Header.Set("Content-Type", "application/json")

	rsp, err := new(http.Client).Do(httpReq)
	if err != nil {
		return fmt.Errorf("could not send email, error: %v", err)
	}
	defer rsp.Body.Close()

	if rsp.StatusCode < 200 || rsp.StatusCode > 299 {
		bytes, err := ioutil.ReadAll(rsp.Body)
		if err != nil {
			return err
		}
		return fmt.Errorf("could not send email, error: %v", string(bytes))
	}

	return nil
}
