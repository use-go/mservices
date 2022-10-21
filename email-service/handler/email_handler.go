package handler

import (
	"bytes"
	"comm/auth"
	"comm/errors"
	"comm/logger"
	"comm/mark"
	"context"
	"fmt"
	"net/smtp"
	"proto/email"
	"regexp"
	"time"

	"github.com/Teamwork/spamc"
	el "github.com/jordan-wright/email"
	"gopkg.in/gomail.v2"
)

// validEmail defined todo
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

// Send defined todo
func (h *Handler) Send(ctx context.Context, request *email.SendRequest, response *email.SendResponse) error {
	var err error
	var timemark mark.TimeMark
	defer timemark.Init(ctx, "Send")()

	acc, ok := auth.FromContext(ctx)
	if ok {
		logger.Infof(ctx, "%v Do Send", acc.Name)
	}
	if !validEmail(request.To) {
		return errors.BadRequest("email", "Invalid to address")
	}
	if len(request.Subject) == 0 {
		return errors.BadRequest("email", "Missing subject")
	}
	if len(request.TextBody) == 0 && len(request.HtmlBody) == 0 {
		return errors.BadRequest("email", "Missing email body")
	}

	spamReq := &email.ClassifyRequest{
		TextBody: request.TextBody,
		HtmlBody: request.HtmlBody,
		To:       request.To,
		From:     request.From,
		Subject:  request.Subject,
	}
	rsp := email.ClassifyResponse{}
	err = h.Classify(ctx, spamReq, &rsp)
	timemark.Mark("Classify")
	if err != nil || rsp.IsSpam {
		logger.Errorf(ctx, "Error validating email %s %v", err, rsp)
		return errors.InternalServerError("email", "Error validating email")
	}

	if err := h.sendEmail(ctx, request); err != nil {
		logger.Errorf(ctx, "Error sending email: %v\n", err)
		return errors.InternalServerError("email", "Error sending email")
	}
	timemark.Mark("sendEmail")

	return nil
}

// Classify defined todo
func (h *Handler) Classify(ctx context.Context, request *email.ClassifyRequest, response *email.ClassifyResponse) error {
	var err error
	var timemark mark.TimeMark
	defer timemark.Init(ctx, "Classify")()

	acc, ok := auth.FromContext(ctx)
	if ok {
		logger.Infof(ctx, "%v Do Classify", acc.Name)
	}
	if len(request.EmailBody) == 0 && len(request.TextBody) == 0 && len(request.HtmlBody) == 0 {
		return errors.BadRequest("email", "spam.Classify %v", "Missing one of email_body, html_body, text_body")
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
			logger.Errorf(ctx, "Error classifying email %v", err)
			return errors.InternalServerError("email", "spam.Classify %v", err)
		}
	}
	rc, err := h.Spamc.Report(ctx, &bf, spamc.Header{}.Set("Content-Length", fmt.Sprintf("%d", bf.Len())))
	if err != nil {
		logger.Errorf(ctx, "Error checking spamd %s", err)
		return errors.InternalServerError("email", "spam.Classify %v", err)
	}
	response.IsSpam = rc.IsSpam
	response.Score = rc.Score
	response.Details = []string{}
	timemark.Mark("Report")

	for _, v := range rc.Report.Table {
		response.Details = append(response.Details, fmt.Sprintf("%s, %s, %v", v.Rule, v.Description, v.Points))
	}
	return nil
}

// sendEmail defined todo
func (h *Handler) sendEmail(ctx context.Context, req *email.SendRequest) error {
	fromName := h.Smtp.UserName
	if len(req.From) > 0 {
		fromName = req.From
	}
	ele := el.Email{
		Subject: req.Subject,
		From:    fromName,
		To:      []string{req.To},
		Text:    []byte(req.TextBody),
		HTML:    []byte(req.HtmlBody),
		ReplyTo: []string{req.ReplyTo},
	}
	err := ele.Send(h.Smtp.Addr, smtp.PlainAuth(h.Smtp.Identity, h.Smtp.UserName, h.Smtp.Password, h.Smtp.Host))
	if err != nil {
		return err
	}
	return nil
}
