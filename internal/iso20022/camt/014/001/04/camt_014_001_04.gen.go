// Code generated by main. DO NOT EDIT.

package camt_014_001_04

import (
	"bytes"
	"encoding/xml"
	"time"
)

type AccountIdentification4Choice struct {
	IBAN IBAN2007Identifier            `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.04 IBAN"`
	Othr GenericAccountIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.04 Othr"`
}

type AccountSchemeName1Choice struct {
	Cd    ExternalAccountIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.04 Cd"`
	Prtry Max35Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.04 Prtry"`
}

// Must match the pattern [A-Z]{3,3}
type ActiveOrHistoricCurrencyCode string

// Must match the pattern [A-Z0-9]{4,4}[A-Z]{2,2}[A-Z0-9]{2,2}([A-Z0-9]{3,3}){0,1}
type BICFIDec2014Identifier string

type CashAccount38 struct {
	Id   AccountIdentification4Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.04 Id"`
	Tp   CashAccountType2Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.04 Tp,omitempty"`
	Ccy  ActiveOrHistoricCurrencyCode `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.04 Ccy,omitempty"`
	Nm   Max70Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.04 Nm,omitempty"`
	Prxy ProxyAccountIdentification1  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.04 Prxy,omitempty"`
}

type CashAccountType2Choice struct {
	Cd    ExternalCashAccountType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.04 Cd"`
	Prtry Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.04 Prtry"`
}

type ClearingSystemIdentification2Choice struct {
	Cd    ExternalClearingSystemIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.04 Cd"`
	Prtry Max35Text                                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.04 Prtry"`
}

type ClearingSystemMemberIdentification2 struct {
	ClrSysId ClearingSystemIdentification2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.04 ClrSysId,omitempty"`
	MmbId    Max35Text                           `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.04 MmbId"`
}

type CommunicationAddress10 struct {
	PstlAdr  LongPostalAddress1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.04 PstlAdr"`
	PhneNb   PhoneNumber              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.04 PhneNb"`
	FaxNb    PhoneNumber              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.04 FaxNb,omitempty"`
	EmailAdr Max2048Text              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.04 EmailAdr,omitempty"`
}

type ContactIdentificationAndAddress2 struct {
	Nm     Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.04 Nm,omitempty"`
	Role   PaymentRole1Choice     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.04 Role"`
	ComAdr CommunicationAddress10 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.04 ComAdr"`
}

// Must match the pattern [A-Z]{2,2}
type CountryCode string

type Document struct {
	RtrMmb ReturnMemberV04 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.04 RtrMmb"`
}

type ErrorHandling1Choice struct {
	Cd    ErrorHandling1Code   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.04 Cd"`
	Prtry Max4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.04 Prtry"`
}

// May be one of X020, X030, X050
type ErrorHandling1Code string

type ErrorHandling3 struct {
	Err  ErrorHandling1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.04 Err"`
	Desc Max140Text           `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.04 Desc,omitempty"`
}

// Must be at least 1 items long
type ExternalAccountIdentification1Code string

// Must be at least 1 items long
type ExternalCashAccountType1Code string

// Must be at least 1 items long
type ExternalClearingSystemIdentification1Code string

// Must be at least 1 items long
type ExternalEnquiryRequestType1Code string

// Must be at least 1 items long
type ExternalFinancialInstitutionIdentification1Code string

// Must be at least 1 items long
type ExternalPaymentControlRequestType1Code string

// Must be at least 1 items long
type ExternalPaymentRole1Code string

// Must be at least 1 items long
type ExternalProxyAccountType1Code string

// Must be at least 1 items long
type ExternalSystemMemberType1Code string

type FinancialIdentificationSchemeName1Choice struct {
	Cd    ExternalFinancialInstitutionIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.04 Cd"`
	Prtry Max35Text                                       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.04 Prtry"`
}

type GenericAccountIdentification1 struct {
	Id      Max34Text                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.04 Id"`
	SchmeNm AccountSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.04 SchmeNm,omitempty"`
	Issr    Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.04 Issr,omitempty"`
}

type GenericFinancialIdentification1 struct {
	Id      Max35Text                                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.04 Id"`
	SchmeNm FinancialIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.04 SchmeNm,omitempty"`
	Issr    Max35Text                                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.04 Issr,omitempty"`
}

type GenericIdentification1 struct {
	Id      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.04 Id"`
	SchmeNm Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.04 SchmeNm,omitempty"`
	Issr    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.04 Issr,omitempty"`
}

// Must match the pattern [A-Z]{2,2}[0-9]{2,2}[a-zA-Z0-9]{1,30}
type IBAN2007Identifier string

type ISODateTime time.Time

func (t *ISODateTime) UnmarshalText(text []byte) error {
	return (*xsdDateTime)(t).UnmarshalText(text)
}
func (t ISODateTime) MarshalText() ([]byte, error) {
	return xsdDateTime(t).MarshalText()
}

type LongPostalAddress1Choice struct {
	Ustrd Max140Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.04 Ustrd"`
	Strd  StructuredLongPostalAddress1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.04 Strd"`
}

// Must be at least 1 items long
type Max140Text string

// Must be at least 1 items long
type Max16Text string

// Must be at least 1 items long
type Max2048Text string

// Must be at least 1 items long
type Max34Text string

// Must be at least 1 items long
type Max350Text string

// Must be at least 1 items long
type Max35Text string

// Must match the pattern [a-zA-Z0-9]{1,4}
type Max4AlphaNumericText string

// Must be at least 1 items long
type Max70Text string

type Member5 struct {
	Nm      Max35Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.04 Nm,omitempty"`
	RtrAdr  []MemberIdentification3Choice      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.04 RtrAdr,omitempty"`
	Acct    []CashAccount38                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.04 Acct,omitempty"`
	Tp      SystemMemberType1Choice            `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.04 Tp,omitempty"`
	Sts     SystemMemberStatus1Choice          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.04 Sts,omitempty"`
	CtctRef []ContactIdentificationAndAddress2 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.04 CtctRef,omitempty"`
	ComAdr  CommunicationAddress10             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.04 ComAdr,omitempty"`
}

type MemberIdentification3Choice struct {
	BICFI       BICFIDec2014Identifier              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.04 BICFI"`
	ClrSysMmbId ClearingSystemMemberIdentification2 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.04 ClrSysMmbId"`
	Othr        GenericFinancialIdentification1     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.04 Othr"`
}

type MemberReport5 struct {
	MmbId    MemberIdentification3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.04 MmbId"`
	MmbOrErr MemberReportOrError6Choice  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.04 MmbOrErr"`
}

type MemberReportOrError5Choice struct {
	Rpt     []MemberReport5  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.04 Rpt"`
	OprlErr []ErrorHandling3 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.04 OprlErr"`
}

type MemberReportOrError6Choice struct {
	Mmb    Member5        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.04 Mmb"`
	BizErr ErrorHandling3 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.04 BizErr"`
}

// May be one of ENBL, DSBL, DLTD, JOIN
type MemberStatus1Code string

type MessageHeader7 struct {
	MsgId       Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.04 MsgId"`
	CreDtTm     ISODateTime            `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.04 CreDtTm,omitempty"`
	ReqTp       RequestType4Choice     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.04 ReqTp,omitempty"`
	OrgnlBizQry OriginalBusinessQuery1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.04 OrgnlBizQry,omitempty"`
	QryNm       Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.04 QryNm,omitempty"`
}

type OriginalBusinessQuery1 struct {
	MsgId   Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.04 MsgId"`
	MsgNmId Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.04 MsgNmId,omitempty"`
	CreDtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.04 CreDtTm,omitempty"`
}

type PaymentRole1Choice struct {
	Cd    ExternalPaymentRole1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.04 Cd"`
	Prtry Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.04 Prtry"`
}

// Must match the pattern \+[0-9]{1,3}-[0-9()+\-]{1,30}
type PhoneNumber string

type ProxyAccountIdentification1 struct {
	Tp ProxyAccountType1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.04 Tp,omitempty"`
	Id Max2048Text             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.04 Id"`
}

type ProxyAccountType1Choice struct {
	Cd    ExternalProxyAccountType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.04 Cd"`
	Prtry Max35Text                     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.04 Prtry"`
}

type RequestType4Choice struct {
	PmtCtrl ExternalPaymentControlRequestType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.04 PmtCtrl"`
	Enqry   ExternalEnquiryRequestType1Code        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.04 Enqry"`
	Prtry   GenericIdentification1                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.04 Prtry"`
}

type ReturnMemberV04 struct {
	MsgHdr      MessageHeader7             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.04 MsgHdr"`
	RptOrErr    MemberReportOrError5Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.04 RptOrErr"`
	SplmtryData []SupplementaryData1       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.04 SplmtryData,omitempty"`
}

type StructuredLongPostalAddress1 struct {
	BldgNm     Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.04 BldgNm,omitempty"`
	StrtNm     Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.04 StrtNm,omitempty"`
	StrtBldgId Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.04 StrtBldgId,omitempty"`
	Flr        Max16Text   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.04 Flr,omitempty"`
	TwnNm      Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.04 TwnNm"`
	DstrctNm   Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.04 DstrctNm,omitempty"`
	RgnId      Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.04 RgnId,omitempty"`
	Stat       Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.04 Stat,omitempty"`
	CtyId      Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.04 CtyId,omitempty"`
	Ctry       CountryCode `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.04 Ctry"`
	PstCdId    Max16Text   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.04 PstCdId"`
	POB        Max16Text   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.04 POB,omitempty"`
}

type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.04 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.04 Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
}

type SystemMemberStatus1Choice struct {
	Cd    MemberStatus1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.04 Cd"`
	Prtry Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.04 Prtry"`
}

type SystemMemberType1Choice struct {
	Cd    ExternalSystemMemberType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.04 Cd"`
	Prtry Max35Text                     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.04 Prtry"`
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
