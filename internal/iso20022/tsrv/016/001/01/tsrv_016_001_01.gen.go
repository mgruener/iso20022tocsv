// Code generated by main. DO NOT EDIT.

package tsrv_016_001_01

import (
	"bytes"
	"encoding/xml"
	"time"
)

type ActiveCurrencyAndAmount struct {
	Value float64            `xml:",chardata"`
	Ccy   ActiveCurrencyCode `xml:"Ccy,attr"`
}

// Must match the pattern [A-Z]{3,3}
type ActiveCurrencyCode string

// May be one of ADDR, PBOX, HOME, BIZZ, MLTO, DLVY
type AddressType2Code string

// Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type AnyBICIdentifier string

type ContactDetails2 struct {
	NmPrfx   NamePrefix1Code `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.016.001.01 NmPrfx,omitempty"`
	Nm       Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.016.001.01 Nm,omitempty"`
	PhneNb   PhoneNumber     `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.016.001.01 PhneNb,omitempty"`
	MobNb    PhoneNumber     `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.016.001.01 MobNb,omitempty"`
	FaxNb    PhoneNumber     `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.016.001.01 FaxNb,omitempty"`
	EmailAdr Max2048Text     `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.016.001.01 EmailAdr,omitempty"`
	Othr     Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.016.001.01 Othr,omitempty"`
}

// Must match the pattern [A-Z]{2,2}
type CountryCode string

type DateAndPlaceOfBirth struct {
	BirthDt     ISODate     `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.016.001.01 BirthDt"`
	PrvcOfBirth Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.016.001.01 PrvcOfBirth,omitempty"`
	CityOfBirth Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.016.001.01 CityOfBirth"`
	CtryOfBirth CountryCode `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.016.001.01 CtryOfBirth"`
}

type Demand2 struct {
	Id           Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.016.001.01 Id"`
	SubmissnDtTm ISODateTime             `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.016.001.01 SubmissnDtTm"`
	Amt          ActiveCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.016.001.01 Amt"`
	AddtlInf     []Max2000Text           `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.016.001.01 AddtlInf,omitempty"`
}

type DemandRefusal1 struct {
	UdrtkgId          Undertaking9   `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.016.001.01 UdrtkgId"`
	AdvsgPtyRefNb     Max35Text      `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.016.001.01 AdvsgPtyRefNb,omitempty"`
	ScndAdvsgPtyRefNb Max35Text      `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.016.001.01 ScndAdvsgPtyRefNb,omitempty"`
	CnfrmrRefNb       Max35Text      `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.016.001.01 CnfrmrRefNb,omitempty"`
	DmndDtls          Demand2        `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.016.001.01 DmndDtls"`
	Sts               Refused7Text   `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.016.001.01 Sts"`
	Dscrpncy          []Discrepancy1 `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.016.001.01 Dscrpncy,omitempty"`
	DspstnOfDocs      []Max2000Text  `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.016.001.01 DspstnOfDocs,omitempty"`
	AddtlInf          []Max2000Text  `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.016.001.01 AddtlInf,omitempty"`
}

type DemandRefusalNotificationV01 struct {
	DmndRfslNtfctnDtls []DemandRefusal1   `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.016.001.01 DmndRfslNtfctnDtls,omitempty"`
	DgtlSgntr          PartyAndSignature2 `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.016.001.01 DgtlSgntr,omitempty"`
}

type Discrepancy1 struct {
	Id    Max35Text    `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.016.001.01 Id"`
	Nrrtv Max20000Text `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.016.001.01 Nrrtv"`
}

type Document struct {
	DmndRfslNtfctn DemandRefusalNotificationV01 `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.016.001.01 DmndRfslNtfctn"`
}

// Must be at least 1 items long
type ExternalOrganisationIdentification1Code string

// Must be at least 1 items long
type ExternalPersonIdentification1Code string

type GenericOrganisationIdentification1 struct {
	Id      Max35Text                                   `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.016.001.01 Id"`
	SchmeNm OrganisationIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.016.001.01 SchmeNm,omitempty"`
	Issr    Max35Text                                   `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.016.001.01 Issr,omitempty"`
}

type GenericPersonIdentification1 struct {
	Id      Max35Text                             `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.016.001.01 Id"`
	SchmeNm PersonIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.016.001.01 SchmeNm,omitempty"`
	Issr    Max35Text                             `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.016.001.01 Issr,omitempty"`
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

// Must be at least 1 items long
type Max140Text string

// Must be at least 1 items long
type Max16Text string

// Must be at least 1 items long
type Max20000Text string

// Must be at least 1 items long
type Max2000Text string

// Must be at least 1 items long
type Max2048Text string

// Must be at least 1 items long
type Max35Text string

// Must be at least 1 items long
type Max70Text string

// May be one of DOCT, MIST, MISS, MADM
type NamePrefix1Code string

type OrganisationIdentification8 struct {
	AnyBIC AnyBICIdentifier                     `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.016.001.01 AnyBIC,omitempty"`
	Othr   []GenericOrganisationIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.016.001.01 Othr,omitempty"`
}

type OrganisationIdentificationSchemeName1Choice struct {
	Cd    ExternalOrganisationIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.016.001.01 Cd"`
	Prtry Max35Text                               `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.016.001.01 Prtry"`
}

type Party11Choice struct {
	OrgId  OrganisationIdentification8 `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.016.001.01 OrgId"`
	PrvtId PersonIdentification5       `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.016.001.01 PrvtId"`
}

type PartyAndSignature2 struct {
	Pty   PartyIdentification43 `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.016.001.01 Pty"`
	Sgntr ProprietaryData3      `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.016.001.01 Sgntr"`
}

type PartyIdentification43 struct {
	Nm        Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.016.001.01 Nm,omitempty"`
	PstlAdr   PostalAddress6  `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.016.001.01 PstlAdr,omitempty"`
	Id        Party11Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.016.001.01 Id,omitempty"`
	CtryOfRes CountryCode     `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.016.001.01 CtryOfRes,omitempty"`
	CtctDtls  ContactDetails2 `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.016.001.01 CtctDtls,omitempty"`
}

type PersonIdentification5 struct {
	DtAndPlcOfBirth DateAndPlaceOfBirth            `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.016.001.01 DtAndPlcOfBirth,omitempty"`
	Othr            []GenericPersonIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.016.001.01 Othr,omitempty"`
}

type PersonIdentificationSchemeName1Choice struct {
	Cd    ExternalPersonIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.016.001.01 Cd"`
	Prtry Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.016.001.01 Prtry"`
}

// Must match the pattern \+[0-9]{1,3}-[0-9()+\-]{1,30}
type PhoneNumber string

type PostalAddress6 struct {
	AdrTp       AddressType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.016.001.01 AdrTp,omitempty"`
	Dept        Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.016.001.01 Dept,omitempty"`
	SubDept     Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.016.001.01 SubDept,omitempty"`
	StrtNm      Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.016.001.01 StrtNm,omitempty"`
	BldgNb      Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.016.001.01 BldgNb,omitempty"`
	PstCd       Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.016.001.01 PstCd,omitempty"`
	TwnNm       Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.016.001.01 TwnNm,omitempty"`
	CtrySubDvsn Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.016.001.01 CtrySubDvsn,omitempty"`
	Ctry        CountryCode      `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.016.001.01 Ctry,omitempty"`
	AdrLine     []Max70Text      `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.016.001.01 AdrLine,omitempty"`
}

type ProprietaryData3 struct {
	Item string `xml:",any"`
}

// Must match the pattern REFUSED
type Refused7Text string

type Undertaking9 struct {
	Id           Max35Text             `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.016.001.01 Id"`
	Issr         PartyIdentification43 `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.016.001.01 Issr"`
	ApplcntRefNb Max35Text             `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.016.001.01 ApplcntRefNb,omitempty"`
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
