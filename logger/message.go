package logger

import "github.com/sirupsen/logrus"

type Message struct {
	Logger *Logger

	Message      string
	Err          error
	customFields logrus.Fields
}

func (m Message) Info() {
	if m.Logger == nil {
		logrus.WithFields(m.fields()).Info(m.Message)
		return
	}

	m.Logger.WithFields(m.fields()).Info(m.Message)
}

func (m Message) Debug() {
	if m.Logger == nil {
		logrus.WithFields(m.fields()).Debug(m.Message)
		return
	}

	m.Logger.WithFields(m.fields()).Debug(m.Message)
}

func (m Message) Warn() {
	if m.Logger == nil {
		logrus.WithFields(m.fields()).Warn(m.Message)
		return
	}

	m.Logger.WithFields(m.fields()).Warn(m.Message)
}

func (m Message) Fatal() {
	if m.Logger == nil {
		logrus.WithFields(m.fields()).Fatal(m.Message)
		return
	}

	m.Logger.WithFields(m.fields()).Fatal(m.Message)
}

func (m Message) Panic() {
	if m.Logger == nil {
		logrus.WithFields(m.fields()).Panic(m.Message)
		return
	}

	m.Logger.WithFields(m.fields()).Panic(m.Message)
}

func (m Message) Error() {
	if m.Logger == nil {
		logrus.WithFields(m.fields()).Error(m.Message)
		return
	}

	m.Logger.WithFields(m.fields()).Error(m.Message)
}

func (m Message) AddField(key string, value interface{}) Message {
	if m.customFields == nil {
		m.customFields = logrus.Fields{}
	}

	m.customFields[key] = value

	return m
}

func (m Message) fields() logrus.Fields {
	f := logrus.Fields{}

	if m.customFields != nil {
		f = m.customFields
	}

	if m.Err != nil {
		f["error"] = m.Err
	}

	return f
}
