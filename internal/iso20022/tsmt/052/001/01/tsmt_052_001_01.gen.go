// Code generated by main. DO NOT EDIT.

package tsmt_052_001_01

import (
	"bytes"
	"encoding/xml"
	"time"
)

// May be one of SBTW, RSTW, RSBS, ARDM, ARCS, ARES, WAIT, UPDT, SBDS, ARBA, ARRO, CINR
type Action2Code string

type BICIdentification1 struct {
	BIC BICIdentifier `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.052.001.01 BIC"`
}

// Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type BICIdentifier string

// May be one of PROP, CLSD, PMTC, ESTD, ACTV, COMP, AMRQ, RARQ, CLRQ, SCRQ, SERQ, DARQ
type BaselineStatus3Code string

type Document struct {
	RoleAndBaselnRjctnNtfctn RoleAndBaselineRejectionNotificationV01 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.052.001.01 RoleAndBaselnRjctnNtfctn"`
}

type DocumentIdentification3 struct {
	Id   Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.052.001.01 Id"`
	Vrsn float64   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.052.001.01 Vrsn"`
}

type DocumentIdentification5 struct {
	Id     Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.052.001.01 Id"`
	IdIssr BICIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.052.001.01 IdIssr"`
}

type ISODateTime time.Time

func (t *ISODateTime) UnmarshalText(text []byte) error {
	return (*xsdDateTime)(t).UnmarshalText(text)
}
func (t ISODateTime) MarshalText() ([]byte, error) {
	return xsdDateTime(t).MarshalText()
}

// Must be at least 1 items long
type Max140Text string

// Must be at least 1 items long
type Max35Text string

type MessageIdentification1 struct {
	Id      Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.052.001.01 Id"`
	CreDtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.052.001.01 CreDtTm"`
}

type PendingActivity2 struct {
	Tp   Action2Code `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.052.001.01 Tp"`
	Desc Max140Text  `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.052.001.01 Desc,omitempty"`
}

type Reason2 struct {
	Desc Max140Text `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.052.001.01 Desc"`
}

type RoleAndBaselineRejectionNotificationV01 struct {
	NtfctnId          MessageIdentification1          `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.052.001.01 NtfctnId"`
	TxId              SimpleIdentificationInformation `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.052.001.01 TxId"`
	EstblishdBaselnId DocumentIdentification3         `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.052.001.01 EstblishdBaselnId,omitempty"`
	TxSts             TransactionStatus4              `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.052.001.01 TxSts"`
	UsrTxRef          []DocumentIdentification5       `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.052.001.01 UsrTxRef,omitempty"`
	Initr             BICIdentification1              `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.052.001.01 Initr"`
	RjctnRsn          Reason2                         `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.052.001.01 RjctnRsn,omitempty"`
	ReqForActn        PendingActivity2                `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.052.001.01 ReqForActn,omitempty"`
}

type SimpleIdentificationInformation struct {
	Id Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.052.001.01 Id"`
}

type TransactionStatus4 struct {
	Sts BaselineStatus3Code `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.052.001.01 Sts"`
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
