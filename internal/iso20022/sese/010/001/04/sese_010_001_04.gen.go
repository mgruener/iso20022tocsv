// Code generated by main. DO NOT EDIT.

package sese_010_001_04

import (
	"bytes"
	"encoding/xml"
	"time"
)

type AdditionalReference2 struct {
	Ref     Max35Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.04 Ref"`
	RefIssr PartyIdentification1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.04 RefIssr,omitempty"`
	MsgNm   Max35Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.04 MsgNm,omitempty"`
}

type AdditionalReference3 struct {
	Ref     Max35Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.04 Ref"`
	RefIssr PartyIdentification2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.04 RefIssr,omitempty"`
	MsgNm   Max35Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.04 MsgNm,omitempty"`
}

// May be one of ADDR, PBOX, HOME, BIZZ, MLTO, DLVY
type AddressType2Code string

// Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type AnyBICIdentifier string

// May be one of CUTO, COSE
type CancellationRejectedReason1Code string

// May be one of STNP, PACK
type CancellationStatus2Code string

type CancellationStatusAndReason2 struct {
	MstrRef  Max35Text                                    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.04 MstrRef,omitempty"`
	TrfRef   Max35Text                                    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.04 TrfRef"`
	ClntRef  Max35Text                                    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.04 ClntRef,omitempty"`
	CxlRef   Max35Text                                    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.04 CxlRef,omitempty"`
	Sts      TransferCancellationStatus2                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.04 Sts"`
	Rjctd    TransferCancellationRejectedStatus1          `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.04 Rjctd"`
	Cmplt    TransferCancellationCompleteStatusAndReason1 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.04 Cmplt"`
	Pdg      TransferCancellationPendingStatus1           `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.04 Pdg"`
	StsInitr PartyIdentification2Choice                   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.04 StsInitr,omitempty"`
}

// May be one of CANI, CANS, CSUB
type CancelledStatusReason1Code string

// Must match the pattern [A-Z]{2,2}
type CountryCode string

type Document struct {
	TrfCxlStsRpt TransferCancellationStatusReportV04 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.04 TrfCxlStsRpt"`
}

// Must be at least 1 items long
type Extended350Code string

type Extension1 struct {
	PlcAndNm Max350Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.04 PlcAndNm"`
	Txt      Max350Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.04 Txt"`
}

type GenericIdentification1 struct {
	Id      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.04 Id"`
	SchmeNm Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.04 SchmeNm,omitempty"`
	Issr    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.04 Issr,omitempty"`
}

type ISODateTime time.Time

func (t *ISODateTime) UnmarshalText(text []byte) error {
	return (*xsdDateTime)(t).UnmarshalText(text)
}
func (t ISODateTime) MarshalText() ([]byte, error) {
	return xsdDateTime(t).MarshalText()
}

type ISOYearMonth time.Time

func (t *ISOYearMonth) UnmarshalText(text []byte) error {
	return (*xsdGYearMonth)(t).UnmarshalText(text)
}
func (t ISOYearMonth) MarshalText() ([]byte, error) {
	return xsdGYearMonth(t).MarshalText()
}

type LongPostalAddress1Choice struct {
	Ustrd Max140Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.04 Ustrd"`
	Strd  StructuredLongPostalAddress1 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.04 Strd"`
}

type MarketPracticeVersion1 struct {
	Nm Max35Text    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.04 Nm"`
	Dt ISOYearMonth `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.04 Dt,omitempty"`
	Nb Max35Text    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.04 Nb,omitempty"`
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
	Id      Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.04 Id"`
	CreDtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.04 CreDtTm"`
}

type NameAndAddress2 struct {
	Nm  Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.04 Nm"`
	Adr LongPostalAddress1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.04 Adr,omitempty"`
}

type NameAndAddress5 struct {
	Nm  Max350Text     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.04 Nm"`
	Adr PostalAddress1 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.04 Adr,omitempty"`
}

type PartyIdentification1Choice struct {
	BICOrBEI AnyBICIdentifier       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.04 BICOrBEI"`
	PrtryId  GenericIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.04 PrtryId"`
	NmAndAdr NameAndAddress2        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.04 NmAndAdr"`
}

type PartyIdentification2Choice struct {
	BICOrBEI AnyBICIdentifier       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.04 BICOrBEI"`
	PrtryId  GenericIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.04 PrtryId"`
	NmAndAdr NameAndAddress5        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.04 NmAndAdr"`
}

type PostalAddress1 struct {
	AdrTp       AddressType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.04 AdrTp,omitempty"`
	AdrLine     []Max70Text      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.04 AdrLine,omitempty"`
	StrtNm      Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.04 StrtNm,omitempty"`
	BldgNb      Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.04 BldgNb,omitempty"`
	PstCd       Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.04 PstCd,omitempty"`
	TwnNm       Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.04 TwnNm,omitempty"`
	CtrySubDvsn Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.04 CtrySubDvsn,omitempty"`
	Ctry        CountryCode      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.04 Ctry"`
}

type References40Choice struct {
	RltdRef []AdditionalReference3 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.04 RltdRef"`
	OthrRef []AdditionalReference3 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.04 OthrRef"`
}

type StructuredLongPostalAddress1 struct {
	BldgNm     Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.04 BldgNm,omitempty"`
	StrtNm     Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.04 StrtNm,omitempty"`
	StrtBldgId Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.04 StrtBldgId,omitempty"`
	Flr        Max16Text   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.04 Flr,omitempty"`
	TwnNm      Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.04 TwnNm"`
	DstrctNm   Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.04 DstrctNm,omitempty"`
	RgnId      Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.04 RgnId,omitempty"`
	Stat       Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.04 Stat,omitempty"`
	CtyId      Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.04 CtyId,omitempty"`
	Ctry       CountryCode `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.04 Ctry"`
	PstCdId    Max16Text   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.04 PstCdId"`
	POB        Max16Text   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.04 POB,omitempty"`
}

type TransferCancellationCompleteStatusAndReason1 struct {
	Rsn          CancelledStatusReason1Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.04 Rsn"`
	XtndedRsn    Extended350Code            `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.04 XtndedRsn"`
	DataSrcSchme GenericIdentification1     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.04 DataSrcSchme"`
}

type TransferCancellationPendingStatus1 struct {
	Rsn Max350Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.04 Rsn,omitempty"`
}

type TransferCancellationRejectedStatus1 struct {
	Rsn          CancellationRejectedReason1Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.04 Rsn"`
	XtndedRsn    Extended350Code                 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.04 XtndedRsn"`
	DataSrcSchme []GenericIdentification1        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.04 DataSrcSchme"`
}

type TransferCancellationStatus2 struct {
	Sts CancellationStatus2Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.04 Sts"`
	Rsn Max350Text              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.04 Rsn,omitempty"`
}

type TransferCancellationStatusReportV04 struct {
	MsgId        MessageIdentification1       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.04 MsgId"`
	CtrPtyRef    AdditionalReference2         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.04 CtrPtyRef,omitempty"`
	Ref          References40Choice           `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.04 Ref,omitempty"`
	StsRpt       CancellationStatusAndReason2 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.04 StsRpt"`
	MktPrctcVrsn MarketPracticeVersion1       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.04 MktPrctcVrsn,omitempty"`
	Xtnsn        []Extension1                 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.04 Xtnsn,omitempty"`
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

type xsdGYearMonth time.Time

func (t *xsdGYearMonth) UnmarshalText(text []byte) error {
	return _unmarshalTime(text, (*time.Time)(t), "2006-01")
}
func (t xsdGYearMonth) MarshalText() ([]byte, error) {
	return _marshalTime((time.Time)(t), "2006-01")
}
func (t xsdGYearMonth) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if (time.Time)(t).IsZero() {
		return nil
	}
	m, err := t.MarshalText()
	if err != nil {
		return err
	}
	return e.EncodeElement(m, start)
}
func (t xsdGYearMonth) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if (time.Time)(t).IsZero() {
		return xml.Attr{}, nil
	}
	m, err := t.MarshalText()
	return xml.Attr{Name: name, Value: string(m)}, err
}
