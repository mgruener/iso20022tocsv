// Code generated by main. DO NOT EDIT.

package semt_001_001_03

import (
	"bytes"
	"encoding/xml"
	"time"
)

type AdditionalReference3 struct {
	Ref     Max35Text                  `xml:"urn:swift:xsd:semt.001.001.03 Ref"`
	RefIssr PartyIdentification2Choice `xml:"urn:swift:xsd:semt.001.001.03 RefIssr,omitempty"`
	MsgNm   Max35Text                  `xml:"urn:swift:xsd:semt.001.001.03 MsgNm,omitempty"`
}

// May be one of ADDR, PBOX, HOME, BIZZ, MLTO, DLVY
type AddressType2Code string

// Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type AnyBICIdentifier string

// Must match the pattern [A-Z]{2,2}
type CountryCode string

type Document struct {
	SctiesMsgRjctn SecuritiesMessageRejectionV03 `xml:"urn:swift:xsd:semt.001.001.03 SctiesMsgRjctn"`
}

type GenericIdentification1 struct {
	Id      Max35Text `xml:"urn:swift:xsd:semt.001.001.03 Id"`
	SchmeNm Max35Text `xml:"urn:swift:xsd:semt.001.001.03 SchmeNm,omitempty"`
	Issr    Max35Text `xml:"urn:swift:xsd:semt.001.001.03 Issr,omitempty"`
}

type ISODateTime time.Time

func (t *ISODateTime) UnmarshalText(text []byte) error {
	return (*xsdDateTime)(t).UnmarshalText(text)
}
func (t ISODateTime) MarshalText() ([]byte, error) {
	return xsdDateTime(t).MarshalText()
}

type LinkedMessage1Choice struct {
	PrvsRef AdditionalReference3 `xml:"urn:swift:xsd:semt.001.001.03 PrvsRef"`
	OthrRef AdditionalReference3 `xml:"urn:swift:xsd:semt.001.001.03 OthrRef"`
	RltdRef AdditionalReference3 `xml:"urn:swift:xsd:semt.001.001.03 RltdRef"`
}

// Must be at least 1 items long
type Max140Text string

// Must be at least 1 items long
type Max16Text string

// Must be at least 1 items long
type Max350Text string

// Must be at least 1 items long
type Max35Text string

// Must be at least 1 items long
type Max70Text string

type MessageIdentification1 struct {
	Id      Max35Text   `xml:"urn:swift:xsd:semt.001.001.03 Id"`
	CreDtTm ISODateTime `xml:"urn:swift:xsd:semt.001.001.03 CreDtTm"`
}

// May be one of REFE, NALO
type MessageRejectedReason1Code string

type NameAndAddress5 struct {
	Nm  Max350Text     `xml:"urn:swift:xsd:semt.001.001.03 Nm"`
	Adr PostalAddress1 `xml:"urn:swift:xsd:semt.001.001.03 Adr,omitempty"`
}

type PartyIdentification2Choice struct {
	BICOrBEI AnyBICIdentifier       `xml:"urn:swift:xsd:semt.001.001.03 BICOrBEI"`
	PrtryId  GenericIdentification1 `xml:"urn:swift:xsd:semt.001.001.03 PrtryId"`
	NmAndAdr NameAndAddress5        `xml:"urn:swift:xsd:semt.001.001.03 NmAndAdr"`
}

type PostalAddress1 struct {
	AdrTp       AddressType2Code `xml:"urn:swift:xsd:semt.001.001.03 AdrTp,omitempty"`
	AdrLine     []Max70Text      `xml:"urn:swift:xsd:semt.001.001.03 AdrLine,omitempty"`
	StrtNm      Max70Text        `xml:"urn:swift:xsd:semt.001.001.03 StrtNm,omitempty"`
	BldgNb      Max16Text        `xml:"urn:swift:xsd:semt.001.001.03 BldgNb,omitempty"`
	PstCd       Max16Text        `xml:"urn:swift:xsd:semt.001.001.03 PstCd,omitempty"`
	TwnNm       Max35Text        `xml:"urn:swift:xsd:semt.001.001.03 TwnNm,omitempty"`
	CtrySubDvsn Max35Text        `xml:"urn:swift:xsd:semt.001.001.03 CtrySubDvsn,omitempty"`
	Ctry        CountryCode      `xml:"urn:swift:xsd:semt.001.001.03 Ctry"`
}

type RejectionReason23 struct {
	Rsn      MessageRejectedReason1Code `xml:"urn:swift:xsd:semt.001.001.03 Rsn"`
	AddtlInf Max140Text                 `xml:"urn:swift:xsd:semt.001.001.03 AddtlInf,omitempty"`
	LkdMsg   LinkedMessage1Choice       `xml:"urn:swift:xsd:semt.001.001.03 LkdMsg,omitempty"`
}

type SecuritiesMessageRejectionV03 struct {
	MsgId   MessageIdentification1 `xml:"urn:swift:xsd:semt.001.001.03 MsgId"`
	RltdRef AdditionalReference3   `xml:"urn:swift:xsd:semt.001.001.03 RltdRef"`
	Rsn     RejectionReason23      `xml:"urn:swift:xsd:semt.001.001.03 Rsn"`
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