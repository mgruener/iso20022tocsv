// Code generated by main. DO NOT EDIT.

package cain_013_001_01

import (
	"bytes"
	"encoding/base64"
	"encoding/xml"
	"time"
)

type AcceptorRejection4 struct {
	RjctRsn  RejectReason1Code `xml:"urn:iso:std:iso:20022:tech:xsd:cain.013.001.01 RjctRsn"`
	ErrRptg  []ErrorReporting1 `xml:"urn:iso:std:iso:20022:tech:xsd:cain.013.001.01 ErrRptg,omitempty"`
	MsgInErr Max100KBinary     `xml:"urn:iso:std:iso:20022:tech:xsd:cain.013.001.01 MsgInErr,omitempty"`
}

type AcquirerRejection struct {
	Hdr  Header19           `xml:"urn:iso:std:iso:20022:tech:xsd:cain.013.001.01 Hdr"`
	Rjct AcceptorRejection4 `xml:"urn:iso:std:iso:20022:tech:xsd:cain.013.001.01 Rjct"`
}

type Document struct {
	AcqrrRjctn AcquirerRejection `xml:"urn:iso:std:iso:20022:tech:xsd:cain.013.001.01 AcqrrRjctn"`
}

type ErrorReporting1 struct {
	Tp   Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:cain.013.001.01 Tp"`
	Desc Max500Text `xml:"urn:iso:std:iso:20022:tech:xsd:cain.013.001.01 Desc"`
}

type GenericIdentification73 struct {
	Id     Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:cain.013.001.01 Id"`
	Tp     PartyType9Code    `xml:"urn:iso:std:iso:20022:tech:xsd:cain.013.001.01 Tp,omitempty"`
	Issr   PartyType9Code    `xml:"urn:iso:std:iso:20022:tech:xsd:cain.013.001.01 Issr,omitempty"`
	Ctry   Min2Max3AlphaText `xml:"urn:iso:std:iso:20022:tech:xsd:cain.013.001.01 Ctry,omitempty"`
	ShrtNm Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:cain.013.001.01 ShrtNm,omitempty"`
}

type GenericIdentification74 struct {
	Id     Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:cain.013.001.01 Id"`
	Tp     PartyType10Code   `xml:"urn:iso:std:iso:20022:tech:xsd:cain.013.001.01 Tp"`
	Issr   PartyType10Code   `xml:"urn:iso:std:iso:20022:tech:xsd:cain.013.001.01 Issr,omitempty"`
	Ctry   Min2Max3AlphaText `xml:"urn:iso:std:iso:20022:tech:xsd:cain.013.001.01 Ctry,omitempty"`
	ShrtNm Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:cain.013.001.01 ShrtNm,omitempty"`
}

type Header19 struct {
	MsgFctn        MessageFunction6Code    `xml:"urn:iso:std:iso:20022:tech:xsd:cain.013.001.01 MsgFctn"`
	OrgnlMsgFctn   MessageFunction6Code    `xml:"urn:iso:std:iso:20022:tech:xsd:cain.013.001.01 OrgnlMsgFctn,omitempty"`
	PrtcolVrsn     Max6Text                `xml:"urn:iso:std:iso:20022:tech:xsd:cain.013.001.01 PrtcolVrsn"`
	XchgId         Max3NumericText         `xml:"urn:iso:std:iso:20022:tech:xsd:cain.013.001.01 XchgId"`
	ReTrnsmssnCntr Max3NumericText         `xml:"urn:iso:std:iso:20022:tech:xsd:cain.013.001.01 ReTrnsmssnCntr,omitempty"`
	CreDtTm        ISODateTime             `xml:"urn:iso:std:iso:20022:tech:xsd:cain.013.001.01 CreDtTm"`
	InitgPty       GenericIdentification73 `xml:"urn:iso:std:iso:20022:tech:xsd:cain.013.001.01 InitgPty,omitempty"`
	RcptPty        GenericIdentification73 `xml:"urn:iso:std:iso:20022:tech:xsd:cain.013.001.01 RcptPty,omitempty"`
	Tracblt        []Traceability3         `xml:"urn:iso:std:iso:20022:tech:xsd:cain.013.001.01 Tracblt,omitempty"`
}

type ISODateTime time.Time

func (t *ISODateTime) UnmarshalText(text []byte) error {
	return (*xsdDateTime)(t).UnmarshalText(text)
}
func (t ISODateTime) MarshalText() ([]byte, error) {
	return xsdDateTime(t).MarshalText()
}

type Max100KBinary []byte

func (t *Max100KBinary) UnmarshalText(text []byte) error {
	return (*xsdBase64Binary)(t).UnmarshalText(text)
}
func (t Max100KBinary) MarshalText() ([]byte, error) {
	return xsdBase64Binary(t).MarshalText()
}

// Must be at least 1 items long
type Max35Text string

// Must match the pattern [0-9]{1,3}
type Max3NumericText string

// Must be at least 1 items long
type Max500Text string

// Must be at least 1 items long
type Max6Text string

// May be one of RCAV, RCAN, RCAQ, REJA, REVV, REVN, REVQ, RCPV, RCPN, RCPQ, REJP, AUTV, AUTN, AUTQ, AUTP, FNCV, FNCN, FNCQ, RCIV, RCIN, RCIQ, REJI, KEYV, KEYQ, MGTV, MGTQ
type MessageFunction6Code string

// Must match the pattern [a-zA-Z]{2,3}
type Min2Max3AlphaText string

// May be one of ACCP, ACQR, ATMG, CISS, DLIS, HSTG, ITAG, MERC, OATM, OPOI
type PartyType10Code string

// May be one of ACQR, ACQP, CISS, CISP, CSCH, SCHP
type PartyType9Code string

// May be one of UNPR, IMSG, PARS, SECU, INTP, RCPP, DPMG, VERS, MSGT
type RejectReason1Code string

type Traceability3 struct {
	RlayId      GenericIdentification74 `xml:"urn:iso:std:iso:20022:tech:xsd:cain.013.001.01 RlayId"`
	TracDtTmIn  ISODateTime             `xml:"urn:iso:std:iso:20022:tech:xsd:cain.013.001.01 TracDtTmIn"`
	TracDtTmOut ISODateTime             `xml:"urn:iso:std:iso:20022:tech:xsd:cain.013.001.01 TracDtTmOut"`
}

type xsdBase64Binary []byte

func (b *xsdBase64Binary) UnmarshalText(text []byte) (err error) {
	*b, err = base64.StdEncoding.DecodeString(string(text))
	return
}
func (b xsdBase64Binary) MarshalText() ([]byte, error) {
	var buf bytes.Buffer
	enc := base64.NewEncoder(base64.StdEncoding, &buf)
	enc.Write([]byte(b))
	enc.Close()
	return buf.Bytes(), nil
}

type xsdDateTime time.Time

func (t *xsdDateTime) UnmarshalText(text []byte) error {
	return _unmarshalTime(text, (*time.Time)(t), "2006-01-02T15:04:05.999999999")
}
func (t xsdDateTime) MarshalText() ([]byte, error) {
	return _marshalTime((time.Time)(t), "2006-01-02T15:04:05.999999999")
}
func (t xsdDateTime) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if (time.Time)(t).IsZero() {
		return nil
	}
	m, err := t.MarshalText()
	if err != nil {
		return err
	}
	return e.EncodeElement(m, start)
}
func (t xsdDateTime) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if (time.Time)(t).IsZero() {
		return xml.Attr{}, nil
	}
	m, err := t.MarshalText()
	return xml.Attr{Name: name, Value: string(m)}, err
}
func _unmarshalTime(text []byte, t *time.Time, format string) (err error) {
	s := string(bytes.TrimSpace(text))
	*t, err = time.Parse(format, s)
	if _, ok := err.(*time.ParseError); ok {
		*t, err = time.Parse(format+"Z07:00", s)
	}
	return err
}
func _marshalTime(t time.Time, format string) ([]byte, error) {
	return []byte(t.Format(format + "Z07:00")), nil
}