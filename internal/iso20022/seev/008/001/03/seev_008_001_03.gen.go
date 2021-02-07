// Code generated by main. DO NOT EDIT.

package seev_008_001_03

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

type AlternateIdentification1 struct {
	Id    Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.03 Id"`
	IdSrc IdentificationSource1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.03 IdSrc"`
}

type AmendInformation2 struct {
	PrvsRef MessageIdentification `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.03 PrvsRef"`
}

// Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type AnyBICIdentifier string

type CommunicationAddress4 struct {
	EmailAdr Max256Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.03 EmailAdr,omitempty"`
	URLAdr   Max256Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.03 URLAdr,omitempty"`
}

// Must match the pattern [A-Z]{2,2}
type CountryCode string

type Document struct {
	MtgRsltDssmntn MeetingResultDisseminationV03 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.03 MtgRsltDssmntn"`
}

type EligiblePosition3 struct {
	AcctId    Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.03 AcctId,omitempty"`
	AcctOwnr  PartyIdentification9Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.03 AcctOwnr,omitempty"`
	HldgBal   []HoldingBalance6            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.03 HldgBal,omitempty"`
	RghtsHldr []PartyIdentification9Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.03 RghtsHldr,omitempty"`
}

// Must match the pattern [a-zA-Z0-9]{4}
type Exact4AlphaNumericText string

type GenericIdentification1 struct {
	Id      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.03 Id"`
	SchmeNm Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.03 SchmeNm,omitempty"`
	Issr    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.03 Issr,omitempty"`
}

type GenericIdentification13 struct {
	Id      Max4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.03 Id"`
	SchmeNm Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.03 SchmeNm,omitempty"`
	Issr    Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.03 Issr"`
}

type GenericIdentification5 struct {
	Issr  Max8Text               `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.03 Issr"`
	Inf   Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.03 Inf"`
	Nrrtv Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.03 Nrrtv,omitempty"`
}

type HoldingBalance6 struct {
	Bal      UnitOrFaceAmount1Choice      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.03 Bal,omitempty"`
	BalTp    SecuritiesEntryType2Code     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.03 BalTp,omitempty"`
	SfkpgPlc SafekeepingPlaceFormatChoice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.03 SfkpgPlc,omitempty"`
	Dt       ISODate                      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.03 Dt,omitempty"`
}

// Must match the pattern [A-Z0-9]{12,12}
type ISINIdentifier string

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

type IdentificationSource1Choice struct {
	Dmst  CountryCode `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.03 Dmst"`
	Prtry Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.03 Prtry"`
}

// Must be at least 1 items long
type Max140Text string

// Must be at least 1 items long
type Max16Text string

// Must be at least 1 items long
type Max256Text string

// Must be at least 1 items long
type Max350Text string

// Must be at least 1 items long
type Max35Text string

// Must match the pattern [a-zA-Z0-9]{1,4}
type Max4AlphaNumericText string

// Must be at least 1 items long
type Max70Text string

// Must be at least 1 items long
type Max8Text string

type MeetingReference4 struct {
	MtgId      Max35Text                        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.03 MtgId,omitempty"`
	IssrMtgId  Max35Text                        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.03 IssrMtgId,omitempty"`
	MtgDtAndTm ISODateTime                      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.03 MtgDtAndTm"`
	Tp         MeetingType2Code                 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.03 Tp"`
	Clssfctn   MeetingTypeClassification1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.03 Clssfctn,omitempty"`
	Lctn       []PostalAddress1                 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.03 Lctn,omitempty"`
}

type MeetingResultDisseminationV03 struct {
	Id       MessageIdentification1     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.03 Id"`
	Amdmnt   AmendInformation2          `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.03 Amdmnt,omitempty"`
	MtgRef   MeetingReference4          `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.03 MtgRef"`
	RptgPty  PartyIdentification9Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.03 RptgPty"`
	Scty     []SecurityPosition6        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.03 Scty"`
	VoteRslt []Vote5                    `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.03 VoteRslt"`
	Prtcptn  Participation3             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.03 Prtcptn,omitempty"`
	AddtlInf CommunicationAddress4      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.03 AddtlInf,omitempty"`
}

// May be one of XMET, GMET, MIXD, SPCL
type MeetingType2Code string

type MeetingTypeClassification1Choice struct {
	Cd    MeetingTypeClassification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.03 Cd"`
	Prtry GenericIdentification13        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.03 Prtry"`
}

// May be one of AMET, OMET, CLAS, ISSU, VRHI, CORT
type MeetingTypeClassification1Code string

type MessageIdentification struct {
	Id Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.03 Id"`
}

type MessageIdentification1 struct {
	Id      Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.03 Id"`
	CreDtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.03 CreDtTm"`
}

type NameAndAddress5 struct {
	Nm  Max350Text     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.03 Nm"`
	Adr PostalAddress1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.03 Adr,omitempty"`
}

type Participation3 struct {
	TtlNbOfVtngRghts     float64                 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.03 TtlNbOfVtngRghts,omitempty"`
	PctgOfVtngRghts      float64                 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.03 PctgOfVtngRghts,omitempty"`
	TtlNbOfSctiesOutsdng ActiveCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.03 TtlNbOfSctiesOutsdng,omitempty"`
	ClctnDt              ISODate                 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.03 ClctnDt,omitempty"`
}

type PartyIdentification3 struct {
	BICOrBEI AnyBICIdentifier `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.03 BICOrBEI"`
}

type PartyIdentification9Choice struct {
	BICOrBEI AnyBICIdentifier       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.03 BICOrBEI"`
	PrtryId  GenericIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.03 PrtryId"`
	NmAndAdr NameAndAddress5        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.03 NmAndAdr"`
}

type PostalAddress1 struct {
	AdrTp       AddressType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.03 AdrTp,omitempty"`
	AdrLine     []Max70Text      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.03 AdrLine,omitempty"`
	StrtNm      Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.03 StrtNm,omitempty"`
	BldgNb      Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.03 BldgNb,omitempty"`
	PstCd       Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.03 PstCd,omitempty"`
	TwnNm       Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.03 TwnNm,omitempty"`
	CtrySubDvsn Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.03 CtrySubDvsn,omitempty"`
	Ctry        CountryCode      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.03 Ctry"`
}

// May be one of CUST, ICSD, NCSD, SHHE
type SafekeepingPlace1Code string

type SafekeepingPlaceAsCodeAndPartyIdentification struct {
	PlcSfkpg SafekeepingPlace1Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.03 PlcSfkpg"`
	Nrrtv    Max35Text             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.03 Nrrtv,omitempty"`
	Pty      PartyIdentification3  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.03 Pty,omitempty"`
}

type SafekeepingPlaceFormatChoice struct {
	Id       SafekeepingPlaceAsCodeAndPartyIdentification `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.03 Id"`
	IdAsDSS  GenericIdentification5                       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.03 IdAsDSS"`
	IdAsCtry CountryCode                                  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.03 IdAsCtry"`
}

// May be one of BLOK, ELIG, PEND, PENR, NOMI, SETD, BORR, LOAN, SPOS, TRAD, COLI, COLO, UNBA, INBA, REGO
type SecuritiesEntryType2Code string

type SecurityIdentification11 struct {
	Id   SecurityIdentification11Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.03 Id"`
	Desc Max140Text                     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.03 Desc,omitempty"`
}

type SecurityIdentification11Choice struct {
	ISIN   ISINIdentifier           `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.03 ISIN"`
	OthrId AlternateIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.03 OthrId"`
}

type SecurityPosition6 struct {
	Id  SecurityIdentification11 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.03 Id"`
	Pos []EligiblePosition3      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.03 Pos,omitempty"`
}

type UnitOrFaceAmount1Choice struct {
	Unit    float64                 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.03 Unit"`
	FaceAmt ActiveCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.03 FaceAmt"`
}

type Vote5 struct {
	IssrLabl Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.03 IssrLabl"`
	Accptd   bool      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.03 Accptd"`
	For      float64   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.03 For,omitempty"`
	Agnst    float64   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.03 Agnst,omitempty"`
	Abstn    float64   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.03 Abstn,omitempty"`
	Wthhld   float64   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.03 Wthhld,omitempty"`
	NoActn   float64   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.03 NoActn,omitempty"`
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