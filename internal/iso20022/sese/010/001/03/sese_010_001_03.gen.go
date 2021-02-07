// Code generated by main. DO NOT EDIT.

package sese_010_001_03

import (
	"bytes"
	"encoding/xml"
	"time"
)

type AdditionalReference2 struct {
	Ref     Max35Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.03 Ref"`
	RefIssr PartyIdentification1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.03 RefIssr,omitempty"`
	MsgNm   Max35Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.03 MsgNm,omitempty"`
}

type AdditionalReference3 struct {
	Ref     Max35Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.03 Ref"`
	RefIssr PartyIdentification2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.03 RefIssr,omitempty"`
	MsgNm   Max35Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.03 MsgNm,omitempty"`
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
	MstrRef  Max35Text                                    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.03 MstrRef,omitempty"`
	TrfRef   Max35Text                                    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.03 TrfRef"`
	ClntRef  Max35Text                                    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.03 ClntRef,omitempty"`
	CxlRef   Max35Text                                    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.03 CxlRef,omitempty"`
	Sts      TransferCancellationStatus2                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.03 Sts"`
	Rjctd    TransferCancellationRejectedStatus1          `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.03 Rjctd"`
	Cmplt    TransferCancellationCompleteStatusAndReason1 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.03 Cmplt"`
	Pdg      TransferCancellationPendingStatus1           `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.03 Pdg"`
	StsInitr PartyIdentification2Choice                   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.03 StsInitr,omitempty"`
}

// May be one of CANI, CANS, CSUB
type CancelledStatusReason1Code string

// Must match the pattern [A-Z]{2,2}
type CountryCode string

type Document struct {
	TrfCxlStsRpt TransferCancellationStatusReportV03 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.03 TrfCxlStsRpt"`
}

// Must be at least 1 items long
type Extended350Code string

type Extension1 struct {
	PlcAndNm Max350Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.03 PlcAndNm"`
	Txt      Max350Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.03 Txt"`
}

type GenericIdentification1 struct {
	Id      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.03 Id"`
	SchmeNm Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.03 SchmeNm,omitempty"`
	Issr    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.03 Issr,omitempty"`
}

type ISODateTime time.Time

func (t *ISODateTime) UnmarshalText(text []byte) error {
	return (*xsdDateTime)(t).UnmarshalText(text)
}
func (t ISODateTime) MarshalText() ([]byte, error) {
	return xsdDateTime(t).MarshalText()
}

type LongPostalAddress1Choice struct {
	Ustrd Max140Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.03 Ustrd"`
	Strd  StructuredLongPostalAddress1 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.03 Strd"`
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
	Id      Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.03 Id"`
	CreDtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.03 CreDtTm"`
}

type NameAndAddress2 struct {
	Nm  Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.03 Nm"`
	Adr LongPostalAddress1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.03 Adr,omitempty"`
}

type NameAndAddress5 struct {
	Nm  Max350Text     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.03 Nm"`
	Adr PostalAddress1 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.03 Adr,omitempty"`
}

type PartyIdentification1Choice struct {
	BICOrBEI AnyBICIdentifier       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.03 BICOrBEI"`
	PrtryId  GenericIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.03 PrtryId"`
	NmAndAdr NameAndAddress2        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.03 NmAndAdr"`
}

type PartyIdentification2Choice struct {
	BICOrBEI AnyBICIdentifier       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.03 BICOrBEI"`
	PrtryId  GenericIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.03 PrtryId"`
	NmAndAdr NameAndAddress5        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.03 NmAndAdr"`
}

type PostalAddress1 struct {
	AdrTp       AddressType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.03 AdrTp,omitempty"`
	AdrLine     []Max70Text      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.03 AdrLine,omitempty"`
	StrtNm      Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.03 StrtNm,omitempty"`
	BldgNb      Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.03 BldgNb,omitempty"`
	PstCd       Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.03 PstCd,omitempty"`
	TwnNm       Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.03 TwnNm,omitempty"`
	CtrySubDvsn Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.03 CtrySubDvsn,omitempty"`
	Ctry        CountryCode      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.03 Ctry"`
}

type StructuredLongPostalAddress1 struct {
	BldgNm     Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.03 BldgNm,omitempty"`
	StrtNm     Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.03 StrtNm,omitempty"`
	StrtBldgId Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.03 StrtBldgId,omitempty"`
	Flr        Max16Text   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.03 Flr,omitempty"`
	TwnNm      Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.03 TwnNm"`
	DstrctNm   Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.03 DstrctNm,omitempty"`
	RgnId      Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.03 RgnId,omitempty"`
	Stat       Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.03 Stat,omitempty"`
	CtyId      Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.03 CtyId,omitempty"`
	Ctry       CountryCode `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.03 Ctry"`
	PstCdId    Max16Text   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.03 PstCdId"`
	POB        Max16Text   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.03 POB,omitempty"`
}

type TransferCancellationCompleteStatusAndReason1 struct {
	Rsn          CancelledStatusReason1Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.03 Rsn"`
	XtndedRsn    Extended350Code            `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.03 XtndedRsn"`
	DataSrcSchme GenericIdentification1     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.03 DataSrcSchme"`
}

type TransferCancellationPendingStatus1 struct {
	Rsn Max350Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.03 Rsn,omitempty"`
}

type TransferCancellationRejectedStatus1 struct {
	Rsn          CancellationRejectedReason1Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.03 Rsn"`
	XtndedRsn    Extended350Code                 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.03 XtndedRsn"`
	DataSrcSchme []GenericIdentification1        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.03 DataSrcSchme"`
}

type TransferCancellationStatus2 struct {
	Sts CancellationStatus2Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.03 Sts"`
	Rsn Max350Text              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.03 Rsn,omitempty"`
}

type TransferCancellationStatusReportV03 struct {
	MsgId     MessageIdentification1       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.03 MsgId"`
	CtrPtyRef AdditionalReference2         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.03 CtrPtyRef,omitempty"`
	RltdRef   []AdditionalReference3       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.03 RltdRef,omitempty"`
	OthrRef   []AdditionalReference3       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.03 OthrRef,omitempty"`
	StsRpt    CancellationStatusAndReason2 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.03 StsRpt"`
	Xtnsn     []Extension1                 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.03 Xtnsn,omitempty"`
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
