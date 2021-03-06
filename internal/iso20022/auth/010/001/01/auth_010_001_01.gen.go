// Code generated by main. DO NOT EDIT.

package auth_010_001_01

import (
	"bytes"
	"encoding/xml"
	"time"
)

// Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type AnyBICIdentifier string

type Document struct {
	RgltryTxRptStsV01 RegulatoryTransactionReportStatusV01 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.010.001.01 RgltryTxRptStsV01"`
}

type DocumentIdentification8 struct {
	Id      Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:auth.010.001.01 Id"`
	CreDtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:auth.010.001.01 CreDtTm,omitempty"`
}

// Must be at least 1 items long
type Extended350Code string

type GenericIdentification1 struct {
	Id      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:auth.010.001.01 Id"`
	SchmeNm Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:auth.010.001.01 SchmeNm,omitempty"`
	Issr    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:auth.010.001.01 Issr,omitempty"`
}

type ISODateTime time.Time

func (t *ISODateTime) UnmarshalText(text []byte) error {
	return (*xsdDateTime)(t).UnmarshalText(text)
}
func (t ISODateTime) MarshalText() ([]byte, error) {
	return xsdDateTime(t).MarshalText()
}

// Must be at least 1 items long
type Max35Text string

// Must be at least 1 items long
type Max70Text string

type PartyIdentification23Choice struct {
	BICOrBEI AnyBICIdentifier       `xml:"urn:iso:std:iso:20022:tech:xsd:auth.010.001.01 BICOrBEI"`
	PrtryId  GenericIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.010.001.01 PrtryId"`
}

type RegulatoryTransactionReportStatusV01 struct {
	Id        DocumentIdentification8            `xml:"urn:iso:std:iso:20022:tech:xsd:auth.010.001.01 Id"`
	RptgInstn PartyIdentification23Choice        `xml:"urn:iso:std:iso:20022:tech:xsd:auth.010.001.01 RptgInstn"`
	RptSts    ReportStatusAndReason1             `xml:"urn:iso:std:iso:20022:tech:xsd:auth.010.001.01 RptSts"`
	IndvTxSts []TradeTransactionStatusAndReason1 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.010.001.01 IndvTxSts"`
}

type RejectedStatusReason9Choice struct {
	Rsn          RejectedStatusReason9Code `xml:"urn:iso:std:iso:20022:tech:xsd:auth.010.001.01 Rsn"`
	XtndedRsn    Extended350Code           `xml:"urn:iso:std:iso:20022:tech:xsd:auth.010.001.01 XtndedRsn"`
	DataSrcSchme GenericIdentification1    `xml:"urn:iso:std:iso:20022:tech:xsd:auth.010.001.01 DataSrcSchme"`
}

// May be one of DSEC, IDNA, ORRF, NSLA, DQUA, NCRR, PLCE, DTRD
type RejectedStatusReason9Code string

type ReportStatusAndReason1 struct {
	RltdRef Max35Text                     `xml:"urn:iso:std:iso:20022:tech:xsd:auth.010.001.01 RltdRef"`
	Sts     Status2Code                   `xml:"urn:iso:std:iso:20022:tech:xsd:auth.010.001.01 Sts"`
	Rjctd   []RejectedStatusReason9Choice `xml:"urn:iso:std:iso:20022:tech:xsd:auth.010.001.01 Rjctd"`
}

// May be one of COMP, PDNG
type Status2Code string

type TradeTransactionStatusAndReason1 struct {
	RltdRef Max35Text                     `xml:"urn:iso:std:iso:20022:tech:xsd:auth.010.001.01 RltdRef"`
	TradRef Max70Text                     `xml:"urn:iso:std:iso:20022:tech:xsd:auth.010.001.01 TradRef"`
	Sts     Status2Code                   `xml:"urn:iso:std:iso:20022:tech:xsd:auth.010.001.01 Sts"`
	Rjctd   []RejectedStatusReason9Choice `xml:"urn:iso:std:iso:20022:tech:xsd:auth.010.001.01 Rjctd"`
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
