// Code generated by main. DO NOT EDIT.

package reda_068_001_01

import (
	"bytes"
	"encoding/base64"
	"encoding/xml"
	"time"
)

// May be one of ADDR, PBOX, HOME, BIZZ, MLTO, DLVY
type AddressType2Code string

type AddressType3Choice struct {
	Cd    AddressType2Code        `xml:"urn:iso:std:iso:20022:tech:xsd:reda.068.001.01 Cd"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.068.001.01 Prtry"`
}

// Must match the pattern [A-Z0-9]{4,4}[A-Z]{2,2}[A-Z0-9]{2,2}([A-Z0-9]{3,3}){0,1}
type AnyBICDec2014Identifier string

type Contact4 struct {
	NmPrfx    NamePrefix2Code             `xml:"urn:iso:std:iso:20022:tech:xsd:reda.068.001.01 NmPrfx,omitempty"`
	Nm        Max140Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:reda.068.001.01 Nm,omitempty"`
	PhneNb    PhoneNumber                 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.068.001.01 PhneNb,omitempty"`
	MobNb     PhoneNumber                 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.068.001.01 MobNb,omitempty"`
	FaxNb     PhoneNumber                 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.068.001.01 FaxNb,omitempty"`
	EmailAdr  Max2048Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.068.001.01 EmailAdr,omitempty"`
	EmailPurp Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:reda.068.001.01 EmailPurp,omitempty"`
	JobTitl   Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:reda.068.001.01 JobTitl,omitempty"`
	Rspnsblty Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:reda.068.001.01 Rspnsblty,omitempty"`
	Dept      Max70Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:reda.068.001.01 Dept,omitempty"`
	Othr      []OtherContact1             `xml:"urn:iso:std:iso:20022:tech:xsd:reda.068.001.01 Othr,omitempty"`
	PrefrdMtd PreferredContactMethod1Code `xml:"urn:iso:std:iso:20022:tech:xsd:reda.068.001.01 PrefrdMtd,omitempty"`
}

// Must match the pattern [A-Z]{2,2}
type CountryCode string

type CreditorEnrolment3 struct {
	Enrlmnt      CreditorServiceEnrolment1      `xml:"urn:iso:std:iso:20022:tech:xsd:reda.068.001.01 Enrlmnt"`
	CdtrTradgNm  Max140Text                     `xml:"urn:iso:std:iso:20022:tech:xsd:reda.068.001.01 CdtrTradgNm,omitempty"`
	Cdtr         RTPPartyIdentification1        `xml:"urn:iso:std:iso:20022:tech:xsd:reda.068.001.01 Cdtr"`
	UltmtCdtr    RTPPartyIdentification1        `xml:"urn:iso:std:iso:20022:tech:xsd:reda.068.001.01 UltmtCdtr,omitempty"`
	MrchntCtgyCd MerchantCategoryCodeIdentifier `xml:"urn:iso:std:iso:20022:tech:xsd:reda.068.001.01 MrchntCtgyCd"`
	CdtrLogo     Max10KBinary                   `xml:"urn:iso:std:iso:20022:tech:xsd:reda.068.001.01 CdtrLogo,omitempty"`
}

type CreditorEnrolmentCancellation2 struct {
	OrgnlBizInstr OriginalBusinessInstruction1         `xml:"urn:iso:std:iso:20022:tech:xsd:reda.068.001.01 OrgnlBizInstr,omitempty"`
	CxlRsn        CreditorEnrolmentCancellationReason2 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.068.001.01 CxlRsn,omitempty"`
	OrgnlEnrlmnt  OriginalEnrolment2Choice             `xml:"urn:iso:std:iso:20022:tech:xsd:reda.068.001.01 OrgnlEnrlmnt"`
	SplmtryData   []SupplementaryData1                 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.068.001.01 SplmtryData,omitempty"`
}

type CreditorEnrolmentCancellationReason1Choice struct {
	Cd    ExternalCreditorEnrolmentCancellationReason1Code `xml:"urn:iso:std:iso:20022:tech:xsd:reda.068.001.01 Cd"`
	Prtry Max35Text                                        `xml:"urn:iso:std:iso:20022:tech:xsd:reda.068.001.01 Prtry"`
}

type CreditorEnrolmentCancellationReason2 struct {
	Orgtr    RTPPartyIdentification1                    `xml:"urn:iso:std:iso:20022:tech:xsd:reda.068.001.01 Orgtr,omitempty"`
	Rsn      CreditorEnrolmentCancellationReason1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:reda.068.001.01 Rsn"`
	AddtlInf []Max105Text                               `xml:"urn:iso:std:iso:20022:tech:xsd:reda.068.001.01 AddtlInf,omitempty"`
}

type CreditorServiceEnrolment1 struct {
	EnrlmntStartDt  DateAndDateTime2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:reda.068.001.01 EnrlmntStartDt,omitempty"`
	EnrlmntEndDt    DateAndDateTime2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:reda.068.001.01 EnrlmntEndDt,omitempty"`
	Vsblty          Visibilty1             `xml:"urn:iso:std:iso:20022:tech:xsd:reda.068.001.01 Vsblty,omitempty"`
	SvcActvtnAllwd  bool                   `xml:"urn:iso:std:iso:20022:tech:xsd:reda.068.001.01 SvcActvtnAllwd"`
	SvcDescLk       Max2048Text            `xml:"urn:iso:std:iso:20022:tech:xsd:reda.068.001.01 SvcDescLk,omitempty"`
	CdtrSvcActvtnLk Max2048Text            `xml:"urn:iso:std:iso:20022:tech:xsd:reda.068.001.01 CdtrSvcActvtnLk,omitempty"`
}

type DateAndDateTime2Choice struct {
	Dt   ISODate     `xml:"urn:iso:std:iso:20022:tech:xsd:reda.068.001.01 Dt"`
	DtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:reda.068.001.01 DtTm"`
}

type DateAndPlaceOfBirth1 struct {
	BirthDt     ISODate     `xml:"urn:iso:std:iso:20022:tech:xsd:reda.068.001.01 BirthDt"`
	PrvcOfBirth Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:reda.068.001.01 PrvcOfBirth,omitempty"`
	CityOfBirth Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:reda.068.001.01 CityOfBirth"`
	CtryOfBirth CountryCode `xml:"urn:iso:std:iso:20022:tech:xsd:reda.068.001.01 CtryOfBirth"`
}

type Document struct {
	ReqToPayCdtrEnrlmntCxlReq RequestToPayCreditorEnrolmentCancellationRequestV01 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.068.001.01 ReqToPayCdtrEnrlmntCxlReq"`
}

type EnrolmentHeader2 struct {
	MsgId    Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:reda.068.001.01 MsgId"`
	CreDtTm  ISODateTime             `xml:"urn:iso:std:iso:20022:tech:xsd:reda.068.001.01 CreDtTm"`
	MsgOrgtr RTPPartyIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.068.001.01 MsgOrgtr,omitempty"`
	MsgRcpt  RTPPartyIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.068.001.01 MsgRcpt,omitempty"`
	InitgPty RTPPartyIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.068.001.01 InitgPty"`
}

// Must match the pattern [a-zA-Z0-9]{4}
type Exact4AlphaNumericText string

// Must be at least 1 items long
type ExternalCreditorEnrolmentCancellationReason1Code string

// Must be at least 1 items long
type ExternalOrganisationIdentification1Code string

// Must be at least 1 items long
type ExternalPersonIdentification1Code string

type GenericIdentification30 struct {
	Id      Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:reda.068.001.01 Id"`
	Issr    Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:reda.068.001.01 Issr"`
	SchmeNm Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:reda.068.001.01 SchmeNm,omitempty"`
}

type GenericOrganisationIdentification1 struct {
	Id      Max35Text                                   `xml:"urn:iso:std:iso:20022:tech:xsd:reda.068.001.01 Id"`
	SchmeNm OrganisationIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:reda.068.001.01 SchmeNm,omitempty"`
	Issr    Max35Text                                   `xml:"urn:iso:std:iso:20022:tech:xsd:reda.068.001.01 Issr,omitempty"`
}

type GenericPersonIdentification1 struct {
	Id      Max35Text                             `xml:"urn:iso:std:iso:20022:tech:xsd:reda.068.001.01 Id"`
	SchmeNm PersonIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:reda.068.001.01 SchmeNm,omitempty"`
	Issr    Max35Text                             `xml:"urn:iso:std:iso:20022:tech:xsd:reda.068.001.01 Issr,omitempty"`
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

// Must match the pattern [A-Z0-9]{18,18}[0-9]{2,2}
type LEIIdentifier string

// Must be at least 1 items long
type Max105Text string

type Max10KBinary []byte

func (t *Max10KBinary) UnmarshalText(text []byte) error {
	return (*xsdBase64Binary)(t).UnmarshalText(text)
}
func (t Max10KBinary) MarshalText() ([]byte, error) {
	return xsdBase64Binary(t).MarshalText()
}

// Must be at least 1 items long
type Max128Text string

// Must be at least 1 items long
type Max140Text string

// Must be at least 1 items long
type Max16Text string

// Must be at least 1 items long
type Max2048Text string

// Must be at least 1 items long
type Max256Text string

// Must be at least 1 items long
type Max350Text string

// Must be at least 1 items long
type Max35Text string

// Must be at least 1 items long
type Max4Text string

// Must be at least 1 items long
type Max70Text string

// Must match the pattern [0-9]{4,4}
type MerchantCategoryCodeIdentifier string

// May be one of DOCT, MADM, MISS, MIST, MIKS
type NamePrefix2Code string

type OrganisationIdentification37 struct {
	AnyBIC   AnyBICDec2014Identifier              `xml:"urn:iso:std:iso:20022:tech:xsd:reda.068.001.01 AnyBIC,omitempty"`
	LEI      LEIIdentifier                        `xml:"urn:iso:std:iso:20022:tech:xsd:reda.068.001.01 LEI,omitempty"`
	EmailAdr Max256Text                           `xml:"urn:iso:std:iso:20022:tech:xsd:reda.068.001.01 EmailAdr,omitempty"`
	Othr     []GenericOrganisationIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.068.001.01 Othr,omitempty"`
}

type OrganisationIdentificationSchemeName1Choice struct {
	Cd    ExternalOrganisationIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:reda.068.001.01 Cd"`
	Prtry Max35Text                               `xml:"urn:iso:std:iso:20022:tech:xsd:reda.068.001.01 Prtry"`
}

type OriginalBusinessInstruction1 struct {
	MsgId   Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:reda.068.001.01 MsgId"`
	MsgNmId Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:reda.068.001.01 MsgNmId,omitempty"`
	CreDtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:reda.068.001.01 CreDtTm,omitempty"`
}

type OriginalEnrolment2Choice struct {
	OrgnlCdtrId      Party49Choice      `xml:"urn:iso:std:iso:20022:tech:xsd:reda.068.001.01 OrgnlCdtrId"`
	OrgnlEnrlmntData CreditorEnrolment3 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.068.001.01 OrgnlEnrlmntData"`
}

type OtherContact1 struct {
	ChanlTp Max4Text   `xml:"urn:iso:std:iso:20022:tech:xsd:reda.068.001.01 ChanlTp"`
	Id      Max128Text `xml:"urn:iso:std:iso:20022:tech:xsd:reda.068.001.01 Id,omitempty"`
}

type Party49Choice struct {
	OrgId  OrganisationIdentification37 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.068.001.01 OrgId"`
	PrvtId PersonIdentification17       `xml:"urn:iso:std:iso:20022:tech:xsd:reda.068.001.01 PrvtId"`
}

type PersonIdentification17 struct {
	DtAndPlcOfBirth DateAndPlaceOfBirth1           `xml:"urn:iso:std:iso:20022:tech:xsd:reda.068.001.01 DtAndPlcOfBirth,omitempty"`
	EmailAdr        Max256Text                     `xml:"urn:iso:std:iso:20022:tech:xsd:reda.068.001.01 EmailAdr,omitempty"`
	Othr            []GenericPersonIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.068.001.01 Othr,omitempty"`
}

type PersonIdentificationSchemeName1Choice struct {
	Cd    ExternalPersonIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:reda.068.001.01 Cd"`
	Prtry Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:reda.068.001.01 Prtry"`
}

// Must match the pattern \+[0-9]{1,3}-[0-9()+\-]{1,30}
type PhoneNumber string

type PostalAddress24 struct {
	AdrTp       AddressType3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:reda.068.001.01 AdrTp,omitempty"`
	Dept        Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:reda.068.001.01 Dept,omitempty"`
	SubDept     Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:reda.068.001.01 SubDept,omitempty"`
	StrtNm      Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:reda.068.001.01 StrtNm,omitempty"`
	BldgNb      Max16Text          `xml:"urn:iso:std:iso:20022:tech:xsd:reda.068.001.01 BldgNb,omitempty"`
	BldgNm      Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:reda.068.001.01 BldgNm,omitempty"`
	Flr         Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:reda.068.001.01 Flr,omitempty"`
	PstBx       Max16Text          `xml:"urn:iso:std:iso:20022:tech:xsd:reda.068.001.01 PstBx,omitempty"`
	Room        Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:reda.068.001.01 Room,omitempty"`
	PstCd       Max16Text          `xml:"urn:iso:std:iso:20022:tech:xsd:reda.068.001.01 PstCd,omitempty"`
	TwnNm       Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:reda.068.001.01 TwnNm,omitempty"`
	TwnLctnNm   Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:reda.068.001.01 TwnLctnNm,omitempty"`
	DstrctNm    Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:reda.068.001.01 DstrctNm,omitempty"`
	CtrySubDvsn Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:reda.068.001.01 CtrySubDvsn,omitempty"`
	Ctry        CountryCode        `xml:"urn:iso:std:iso:20022:tech:xsd:reda.068.001.01 Ctry,omitempty"`
	AdrLine     []Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:reda.068.001.01 AdrLine,omitempty"`
}

// May be one of LETT, MAIL, PHON, FAXX, CELL
type PreferredContactMethod1Code string

type RTPPartyIdentification1 struct {
	Nm        Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:reda.068.001.01 Nm,omitempty"`
	PstlAdr   PostalAddress24 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.068.001.01 PstlAdr,omitempty"`
	Id        Party49Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:reda.068.001.01 Id,omitempty"`
	CtryOfRes CountryCode     `xml:"urn:iso:std:iso:20022:tech:xsd:reda.068.001.01 CtryOfRes,omitempty"`
	CtctDtls  Contact4        `xml:"urn:iso:std:iso:20022:tech:xsd:reda.068.001.01 CtctDtls,omitempty"`
}

type RequestToPayCreditorEnrolmentCancellationRequestV01 struct {
	Hdr         EnrolmentHeader2                 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.068.001.01 Hdr"`
	CxlData     []CreditorEnrolmentCancellation2 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.068.001.01 CxlData"`
	SplmtryData []SupplementaryData1             `xml:"urn:iso:std:iso:20022:tech:xsd:reda.068.001.01 SplmtryData,omitempty"`
}

type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.068.001.01 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.068.001.01 Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
}

type Visibilty1 struct {
	StartDt   DateAndDateTime2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:reda.068.001.01 StartDt,omitempty"`
	EndDt     DateAndDateTime2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:reda.068.001.01 EndDt,omitempty"`
	LtdVsblty bool                   `xml:"urn:iso:std:iso:20022:tech:xsd:reda.068.001.01 LtdVsblty,omitempty"`
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
