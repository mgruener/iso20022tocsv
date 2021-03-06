// Code generated by main. DO NOT EDIT.

package setr_002_001_04

import (
	"bytes"
	"encoding/xml"
	"time"
)

type AdditionalReference8 struct {
	Ref     Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:setr.002.001.04 Ref"`
	RefIssr PartyIdentification113 `xml:"urn:iso:std:iso:20022:tech:xsd:setr.002.001.04 RefIssr,omitempty"`
	MsgNm   Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:setr.002.001.04 MsgNm,omitempty"`
}

type AdditionalReference9 struct {
	Ref     Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:setr.002.001.04 Ref"`
	RefIssr PartyIdentification113 `xml:"urn:iso:std:iso:20022:tech:xsd:setr.002.001.04 RefIssr,omitempty"`
	MsgNm   Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:setr.002.001.04 MsgNm,omitempty"`
}

// May be one of ADDR, PBOX, HOME, BIZZ, MLTO, DLVY
type AddressType2Code string

// Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type AnyBICIdentifier string

type CancellationReason32Choice struct {
	Cd    Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:setr.002.001.04 Cd"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:setr.002.001.04 Prtry"`
}

type CopyInformation4 struct {
	CpyInd    bool             `xml:"urn:iso:std:iso:20022:tech:xsd:setr.002.001.04 CpyInd"`
	OrgnlRcvr AnyBICIdentifier `xml:"urn:iso:std:iso:20022:tech:xsd:setr.002.001.04 OrgnlRcvr,omitempty"`
}

// Must match the pattern [A-Z]{2,2}
type CountryCode string

type Document struct {
	RedBlkOrdrCxlReq RedemptionBulkOrderCancellationRequestV04 `xml:"urn:iso:std:iso:20022:tech:xsd:setr.002.001.04 RedBlkOrdrCxlReq"`
}

// Must match the pattern [a-zA-Z0-9]{4}
type Exact4AlphaNumericText string

type GenericIdentification1 struct {
	Id      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:setr.002.001.04 Id"`
	SchmeNm Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:setr.002.001.04 SchmeNm,omitempty"`
	Issr    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:setr.002.001.04 Issr,omitempty"`
}

type GenericIdentification30 struct {
	Id      Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:setr.002.001.04 Id"`
	Issr    Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:setr.002.001.04 Issr"`
	SchmeNm Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:setr.002.001.04 SchmeNm,omitempty"`
}

type ISODateTime time.Time

func (t *ISODateTime) UnmarshalText(text []byte) error {
	return (*xsdDateTime)(t).UnmarshalText(text)
}
func (t ISODateTime) MarshalText() ([]byte, error) {
	return xsdDateTime(t).MarshalText()
}

type InvestmentFundOrder9 struct {
	OrdrRef Max35Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:setr.002.001.04 OrdrRef"`
	ClntRef Max35Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:setr.002.001.04 ClntRef,omitempty"`
	CxlRef  Max35Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:setr.002.001.04 CxlRef,omitempty"`
	CxlRsn  CancellationReason32Choice `xml:"urn:iso:std:iso:20022:tech:xsd:setr.002.001.04 CxlRsn,omitempty"`
}

// Must match the pattern [A-Z0-9]{18,18}[0-9]{2,2}
type LEIIdentifier string

// Must be at least 1 items long
type Max16Text string

// Must be at least 1 items long
type Max350Text string

// Must be at least 1 items long
type Max35Text string

// Must be at least 1 items long
type Max70Text string

type MessageIdentification1 struct {
	Id      Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:setr.002.001.04 Id"`
	CreDtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:setr.002.001.04 CreDtTm"`
}

type NameAndAddress5 struct {
	Nm  Max350Text     `xml:"urn:iso:std:iso:20022:tech:xsd:setr.002.001.04 Nm"`
	Adr PostalAddress1 `xml:"urn:iso:std:iso:20022:tech:xsd:setr.002.001.04 Adr,omitempty"`
}

type PartyIdentification113 struct {
	Pty PartyIdentification90Choice `xml:"urn:iso:std:iso:20022:tech:xsd:setr.002.001.04 Pty"`
	LEI LEIIdentifier               `xml:"urn:iso:std:iso:20022:tech:xsd:setr.002.001.04 LEI,omitempty"`
}

type PartyIdentification90Choice struct {
	AnyBIC   AnyBICIdentifier       `xml:"urn:iso:std:iso:20022:tech:xsd:setr.002.001.04 AnyBIC"`
	PrtryId  GenericIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:setr.002.001.04 PrtryId"`
	NmAndAdr NameAndAddress5        `xml:"urn:iso:std:iso:20022:tech:xsd:setr.002.001.04 NmAndAdr"`
}

type PostalAddress1 struct {
	AdrTp       AddressType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:setr.002.001.04 AdrTp,omitempty"`
	AdrLine     []Max70Text      `xml:"urn:iso:std:iso:20022:tech:xsd:setr.002.001.04 AdrLine,omitempty"`
	StrtNm      Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:setr.002.001.04 StrtNm,omitempty"`
	BldgNb      Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:setr.002.001.04 BldgNb,omitempty"`
	PstCd       Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:setr.002.001.04 PstCd,omitempty"`
	TwnNm       Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:setr.002.001.04 TwnNm,omitempty"`
	CtrySubDvsn Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:setr.002.001.04 CtrySubDvsn,omitempty"`
	Ctry        CountryCode      `xml:"urn:iso:std:iso:20022:tech:xsd:setr.002.001.04 Ctry"`
}

type RedemptionBulkOrderCancellationRequestV04 struct {
	MsgId    MessageIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:setr.002.001.04 MsgId"`
	PoolRef  AdditionalReference9   `xml:"urn:iso:std:iso:20022:tech:xsd:setr.002.001.04 PoolRef,omitempty"`
	PrvsRef  AdditionalReference8   `xml:"urn:iso:std:iso:20022:tech:xsd:setr.002.001.04 PrvsRef,omitempty"`
	MstrRef  Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:setr.002.001.04 MstrRef,omitempty"`
	OrdrRefs []InvestmentFundOrder9 `xml:"urn:iso:std:iso:20022:tech:xsd:setr.002.001.04 OrdrRefs"`
	CpyDtls  CopyInformation4       `xml:"urn:iso:std:iso:20022:tech:xsd:setr.002.001.04 CpyDtls,omitempty"`
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
