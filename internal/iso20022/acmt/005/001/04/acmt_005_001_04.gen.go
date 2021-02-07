// Code generated by main. DO NOT EDIT.

package acmt_005_001_04

import (
	"bytes"
	"encoding/xml"
	"time"
)

type AccountManagementMessageReference3 struct {
	LkdRef      LinkedMessage3Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.04 LkdRef,omitempty"`
	StsReqTp    AccountManagementType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.04 StsReqTp"`
	AcctApplId  Max35Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.04 AcctApplId,omitempty"`
	ExstgAcctId Max35Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.04 ExstgAcctId,omitempty"`
	InvstmtAcct InvestmentAccount53        `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.04 InvstmtAcct,omitempty"`
}

// May be one of ACCO, ACCM
type AccountManagementType1Code string

type AdditionalReference6 struct {
	Ref     Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.04 Ref"`
	RefIssr PartyIdentification90Choice `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.04 RefIssr,omitempty"`
	MsgNm   Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.04 MsgNm,omitempty"`
}

// May be one of ADDR, PBOX, HOME, BIZZ, MLTO, DLVY
type AddressType2Code string

// Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type AnyBICIdentifier string

// Must match the pattern [A-Z]{2,2}
type CountryCode string

type Document struct {
	ReqForAcctMgmtStsRpt RequestForAccountManagementStatusReportV04 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.04 ReqForAcctMgmtStsRpt"`
}

// Must match the pattern [a-zA-Z0-9]{4}
type Exact4AlphaNumericText string

// May be one of MALE, FEMA
type GenderCode string

type GenericIdentification1 struct {
	Id      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.04 Id"`
	SchmeNm Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.04 SchmeNm,omitempty"`
	Issr    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.04 Issr,omitempty"`
}

type GenericIdentification47 struct {
	Id      Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.04 Id"`
	Issr    Max4AlphaNumericText   `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.04 Issr"`
	SchmeNm Max4AlphaNumericText   `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.04 SchmeNm,omitempty"`
}

type GenericIdentification81 struct {
	Id   Max35Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.04 Id"`
	IdTp OtherIdentification3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.04 IdTp"`
}

type ISODate time.Time

func (t *ISODate) UnmarshalText(text []byte) error {
	return (*xsdDate)(t).UnmarshalText(text)
}
func (t ISODate) MarshalText() ([]byte, error) {
	return xsdDate(t).MarshalText()
}

type ISODateTime time.Time

func (t *ISODateTime) UnmarshalText(text []byte) error {
	return (*xsdDateTime)(t).UnmarshalText(text)
}
func (t ISODateTime) MarshalText() ([]byte, error) {
	return xsdDateTime(t).MarshalText()
}

type IndividualPerson30 struct {
	GvnNm   Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.04 GvnNm,omitempty"`
	MddlNm  Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.04 MddlNm,omitempty"`
	Nm      Max350Text `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.04 Nm"`
	Gndr    GenderCode `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.04 Gndr,omitempty"`
	BirthDt ISODate    `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.04 BirthDt,omitempty"`
}

type IndividualPersonIdentification2Choice struct {
	IdNb   GenericIdentification81 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.04 IdNb"`
	PrsnNm IndividualPerson30      `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.04 PrsnNm"`
}

type InvestmentAccount53 struct {
	AcctId    Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.04 AcctId"`
	AcctNm    Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.04 AcctNm,omitempty"`
	AcctDsgnt Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.04 AcctDsgnt,omitempty"`
	OwnrId    OwnerIdentification2Choice  `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.04 OwnrId,omitempty"`
	AcctSvcr  PartyIdentification70Choice `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.04 AcctSvcr,omitempty"`
}

// Must match the pattern [A-Z0-9]{18,18}[0-9]{2,2}
type LEIIdentifier string

type LinkedMessage3Choice struct {
	PrvsRef AdditionalReference6 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.04 PrvsRef"`
	OthrRef AdditionalReference6 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.04 OthrRef"`
}

// Must be at least 1 items long
type Max16Text string

// Must be at least 1 items long
type Max350Text string

// Must be at least 1 items long
type Max35Text string

// Must match the pattern [a-zA-Z0-9]{1,4}
type Max4AlphaNumericText string

// Must be at least 1 items long
type Max70Text string

type MessageIdentification1 struct {
	Id      Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.04 Id"`
	CreDtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.04 CreDtTm"`
}

type NameAndAddress5 struct {
	Nm  Max350Text     `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.04 Nm"`
	Adr PostalAddress1 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.04 Adr,omitempty"`
}

type OtherIdentification3Choice struct {
	Cd    PartyIdentificationType7Code `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.04 Cd"`
	Prtry GenericIdentification47      `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.04 Prtry"`
}

type OwnerIdentification2Choice struct {
	IndvOwnrId IndividualPersonIdentification2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.04 IndvOwnrId"`
	OrgOwnrId  PartyIdentification95                 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.04 OrgOwnrId"`
}

type PartyIdentification70Choice struct {
	AnyBIC   AnyBICIdentifier       `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.04 AnyBIC"`
	PrtryId  GenericIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.04 PrtryId"`
	NmAndAdr NameAndAddress5        `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.04 NmAndAdr"`
}

type PartyIdentification90Choice struct {
	AnyBIC   AnyBICIdentifier       `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.04 AnyBIC"`
	PrtryId  GenericIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.04 PrtryId"`
	NmAndAdr NameAndAddress5        `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.04 NmAndAdr"`
}

type PartyIdentification95 struct {
	Id         PartyIdentification70Choice `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.04 Id,omitempty"`
	LglNttyIdr LEIIdentifier               `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.04 LglNttyIdr,omitempty"`
}

// May be one of ATIN, IDCD, NRIN, OTHR, PASS, POCD, SOCS, SRSA, GUNL, GTIN, ITIN, CPFA, AREG, DRLC, EMID, NINV, INCL, GIIN
type PartyIdentificationType7Code string

type PostalAddress1 struct {
	AdrTp       AddressType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.04 AdrTp,omitempty"`
	AdrLine     []Max70Text      `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.04 AdrLine,omitempty"`
	StrtNm      Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.04 StrtNm,omitempty"`
	BldgNb      Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.04 BldgNb,omitempty"`
	PstCd       Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.04 PstCd,omitempty"`
	TwnNm       Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.04 TwnNm,omitempty"`
	CtrySubDvsn Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.04 CtrySubDvsn,omitempty"`
	Ctry        CountryCode      `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.04 Ctry"`
}

type RequestForAccountManagementStatusReportV04 struct {
	MsgId   MessageIdentification1             `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.04 MsgId"`
	ReqDtls AccountManagementMessageReference3 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.04 ReqDtls"`
}

type xsdDate time.Time

func (t *xsdDate) UnmarshalText(text []byte) error {
	return _unmarshalTime(text, (*time.Time)(t), "2006-01-02")
}
func (t xsdDate) MarshalText() ([]byte, error) {
	return _marshalTime((time.Time)(t), "2006-01-02")
}
func (t xsdDate) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if (time.Time)(t).IsZero() {
		return nil
	}
	m, err := t.MarshalText()
	if err != nil {
		return err
	}
	return e.EncodeElement(m, start)
}
func (t xsdDate) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
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