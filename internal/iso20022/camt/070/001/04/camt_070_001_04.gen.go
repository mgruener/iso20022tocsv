// Code generated by main. DO NOT EDIT.

package camt_070_001_04

import (
	"bytes"
	"encoding/xml"
	"time"
)

type AccountIdentification4Choice struct {
	IBAN IBAN2007Identifier            `xml:"urn:iso:std:iso:20022:tech:xsd:camt.070.001.04 IBAN"`
	Othr GenericAccountIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.070.001.04 Othr"`
}

type AccountSchemeName1Choice struct {
	Cd    ExternalAccountIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.070.001.04 Cd"`
	Prtry Max35Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.070.001.04 Prtry"`
}

type ActiveCurrencyAndAmount struct {
	Value float64            `xml:",chardata"`
	Ccy   ActiveCurrencyCode `xml:"Ccy,attr"`
}

// Must match the pattern [A-Z]{3,3}
type ActiveCurrencyCode string

// Must match the pattern [A-Z]{3,3}
type ActiveOrHistoricCurrencyCode string

// May be one of ADDR, PBOX, HOME, BIZZ, MLTO, DLVY
type AddressType2Code string

type AddressType3Choice struct {
	Cd    AddressType2Code        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.070.001.04 Cd"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.070.001.04 Prtry"`
}

type Amount2Choice struct {
	AmtWthtCcy float64                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.070.001.04 AmtWthtCcy"`
	AmtWthCcy  ActiveCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:camt.070.001.04 AmtWthCcy"`
}

// Must match the pattern [A-Z0-9]{4,4}[A-Z]{2,2}[A-Z0-9]{2,2}([A-Z0-9]{3,3}){0,1}
type BICFIDec2014Identifier string

type BranchAndFinancialInstitutionIdentification6 struct {
	FinInstnId FinancialInstitutionIdentification18 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.070.001.04 FinInstnId"`
	BrnchId    BranchData3                          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.070.001.04 BrnchId,omitempty"`
}

type BranchData3 struct {
	Id      Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.070.001.04 Id,omitempty"`
	LEI     LEIIdentifier   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.070.001.04 LEI,omitempty"`
	Nm      Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.070.001.04 Nm,omitempty"`
	PstlAdr PostalAddress24 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.070.001.04 PstlAdr,omitempty"`
}

type CashAccount38 struct {
	Id   AccountIdentification4Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.070.001.04 Id"`
	Tp   CashAccountType2Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.070.001.04 Tp,omitempty"`
	Ccy  ActiveOrHistoricCurrencyCode `xml:"urn:iso:std:iso:20022:tech:xsd:camt.070.001.04 Ccy,omitempty"`
	Nm   Max70Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.070.001.04 Nm,omitempty"`
	Prxy ProxyAccountIdentification1  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.070.001.04 Prxy,omitempty"`
}

type CashAccountType2Choice struct {
	Cd    ExternalCashAccountType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.070.001.04 Cd"`
	Prtry Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.070.001.04 Prtry"`
}

type ClearingSystemIdentification2Choice struct {
	Cd    ExternalClearingSystemIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.070.001.04 Cd"`
	Prtry Max35Text                                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.070.001.04 Prtry"`
}

type ClearingSystemMemberIdentification2 struct {
	ClrSysId ClearingSystemIdentification2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.070.001.04 ClrSysId,omitempty"`
	MmbId    Max35Text                           `xml:"urn:iso:std:iso:20022:tech:xsd:camt.070.001.04 MmbId"`
}

// Must match the pattern [A-Z]{2,2}
type CountryCode string

// May be one of CRDT, DBIT
type CreditDebitCode string

type DatePeriodDetails1 struct {
	FrDt ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:camt.070.001.04 FrDt"`
	ToDt ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:camt.070.001.04 ToDt,omitempty"`
}

type Document struct {
	RtrStgOrdr ReturnStandingOrderV04 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.070.001.04 RtrStgOrdr"`
}

type ErrorHandling3Choice struct {
	Cd    ExternalSystemErrorHandling1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.070.001.04 Cd"`
	Prtry Max35Text                        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.070.001.04 Prtry"`
}

type ErrorHandling5 struct {
	Err  ErrorHandling3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.070.001.04 Err"`
	Desc Max140Text           `xml:"urn:iso:std:iso:20022:tech:xsd:camt.070.001.04 Desc,omitempty"`
}

type EventType1Choice struct {
	Cd    ExternalSystemEventType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.070.001.04 Cd"`
	Prtry Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.070.001.04 Prtry"`
}

// Must match the pattern [a-zA-Z0-9]{4}
type Exact4AlphaNumericText string

type ExecutionType1Choice struct {
	Tm  ISOTime          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.070.001.04 Tm"`
	Evt EventType1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.070.001.04 Evt"`
}

// Must be at least 1 items long
type ExternalAccountIdentification1Code string

// Must be at least 1 items long
type ExternalCashAccountType1Code string

// Must be at least 1 items long
type ExternalClearingSystemIdentification1Code string

// Must be at least 1 items long
type ExternalFinancialInstitutionIdentification1Code string

// Must be at least 1 items long
type ExternalProxyAccountType1Code string

// Must be at least 1 items long
type ExternalSystemErrorHandling1Code string

// Must be at least 1 items long
type ExternalSystemEventType1Code string

type FinancialIdentificationSchemeName1Choice struct {
	Cd    ExternalFinancialInstitutionIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.070.001.04 Cd"`
	Prtry Max35Text                                       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.070.001.04 Prtry"`
}

type FinancialInstitutionIdentification18 struct {
	BICFI       BICFIDec2014Identifier              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.070.001.04 BICFI,omitempty"`
	ClrSysMmbId ClearingSystemMemberIdentification2 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.070.001.04 ClrSysMmbId,omitempty"`
	LEI         LEIIdentifier                       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.070.001.04 LEI,omitempty"`
	Nm          Max140Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.070.001.04 Nm,omitempty"`
	PstlAdr     PostalAddress24                     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.070.001.04 PstlAdr,omitempty"`
	Othr        GenericFinancialIdentification1     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.070.001.04 Othr,omitempty"`
}

// May be one of YEAR, MNTH, QURT, MIAN, WEEK, DAIL, ADHO, INDA, OVNG
type Frequency2Code string

type GenericAccountIdentification1 struct {
	Id      Max34Text                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.070.001.04 Id"`
	SchmeNm AccountSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.070.001.04 SchmeNm,omitempty"`
	Issr    Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.070.001.04 Issr,omitempty"`
}

type GenericFinancialIdentification1 struct {
	Id      Max35Text                                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.070.001.04 Id"`
	SchmeNm FinancialIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.070.001.04 SchmeNm,omitempty"`
	Issr    Max35Text                                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.070.001.04 Issr,omitempty"`
}

type GenericIdentification1 struct {
	Id      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:camt.070.001.04 Id"`
	SchmeNm Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:camt.070.001.04 SchmeNm,omitempty"`
	Issr    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:camt.070.001.04 Issr,omitempty"`
}

type GenericIdentification30 struct {
	Id      Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:camt.070.001.04 Id"`
	Issr    Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.070.001.04 Issr"`
	SchmeNm Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.070.001.04 SchmeNm,omitempty"`
}

// Must match the pattern [A-Z]{2,2}[0-9]{2,2}[a-zA-Z0-9]{1,30}
type IBAN2007Identifier string

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

type ISOTime time.Time

func (t *ISOTime) UnmarshalText(text []byte) error {
	return (*xsdTime)(t).UnmarshalText(text)
}
func (t ISOTime) MarshalText() ([]byte, error) {
	return xsdTime(t).MarshalText()
}

// Must match the pattern [A-Z0-9]{18,18}[0-9]{2,2}
type LEIIdentifier string

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

// Must be at least 1 items long
type Max70Text string

type MessageHeader6 struct {
	MsgId       Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.070.001.04 MsgId"`
	CreDtTm     ISODateTime            `xml:"urn:iso:std:iso:20022:tech:xsd:camt.070.001.04 CreDtTm,omitempty"`
	OrgnlBizQry OriginalBusinessQuery1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.070.001.04 OrgnlBizQry,omitempty"`
	QryNm       Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.070.001.04 QryNm,omitempty"`
	ReqTp       RequestType3Choice     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.070.001.04 ReqTp,omitempty"`
}

type OriginalBusinessQuery1 struct {
	MsgId   Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.070.001.04 MsgId"`
	MsgNmId Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.070.001.04 MsgNmId,omitempty"`
	CreDtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:camt.070.001.04 CreDtTm,omitempty"`
}

type PostalAddress24 struct {
	AdrTp       AddressType3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.070.001.04 AdrTp,omitempty"`
	Dept        Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.070.001.04 Dept,omitempty"`
	SubDept     Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.070.001.04 SubDept,omitempty"`
	StrtNm      Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.070.001.04 StrtNm,omitempty"`
	BldgNb      Max16Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.070.001.04 BldgNb,omitempty"`
	BldgNm      Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.070.001.04 BldgNm,omitempty"`
	Flr         Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.070.001.04 Flr,omitempty"`
	PstBx       Max16Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.070.001.04 PstBx,omitempty"`
	Room        Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.070.001.04 Room,omitempty"`
	PstCd       Max16Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.070.001.04 PstCd,omitempty"`
	TwnNm       Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.070.001.04 TwnNm,omitempty"`
	TwnLctnNm   Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.070.001.04 TwnLctnNm,omitempty"`
	DstrctNm    Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.070.001.04 DstrctNm,omitempty"`
	CtrySubDvsn Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.070.001.04 CtrySubDvsn,omitempty"`
	Ctry        CountryCode        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.070.001.04 Ctry,omitempty"`
	AdrLine     []Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.070.001.04 AdrLine,omitempty"`
}

type ProxyAccountIdentification1 struct {
	Tp ProxyAccountType1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.070.001.04 Tp,omitempty"`
	Id Max2048Text             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.070.001.04 Id"`
}

type ProxyAccountType1Choice struct {
	Cd    ExternalProxyAccountType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.070.001.04 Cd"`
	Prtry Max35Text                     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.070.001.04 Prtry"`
}

type RequestType3Choice struct {
	Cd    StandingOrderQueryType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.070.001.04 Cd"`
	Prtry GenericIdentification1      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.070.001.04 Prtry"`
}

type ReturnStandingOrderV04 struct {
	MsgHdr      MessageHeader6              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.070.001.04 MsgHdr"`
	RptOrErr    StandingOrderOrError5Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.070.001.04 RptOrErr"`
	SplmtryData []SupplementaryData1        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.070.001.04 SplmtryData,omitempty"`
}

type StandingOrder6 struct {
	Amt             Amount2Choice                                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.070.001.04 Amt"`
	CdtDbtInd       CreditDebitCode                              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.070.001.04 CdtDbtInd"`
	Ccy             ActiveCurrencyCode                           `xml:"urn:iso:std:iso:20022:tech:xsd:camt.070.001.04 Ccy,omitempty"`
	Tp              StandingOrderType1Choice                     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.070.001.04 Tp,omitempty"`
	AssoctdPoolAcct AccountIdentification4Choice                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.070.001.04 AssoctdPoolAcct,omitempty"`
	Ref             Max35Text                                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.070.001.04 Ref,omitempty"`
	Frqcy           Frequency2Code                               `xml:"urn:iso:std:iso:20022:tech:xsd:camt.070.001.04 Frqcy,omitempty"`
	VldtyPrd        DatePeriodDetails1                           `xml:"urn:iso:std:iso:20022:tech:xsd:camt.070.001.04 VldtyPrd,omitempty"`
	SysMmb          BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.070.001.04 SysMmb,omitempty"`
	RspnsblPty      BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.070.001.04 RspnsblPty,omitempty"`
	LkSetId         Max35Text                                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.070.001.04 LkSetId,omitempty"`
	LkSetOrdrId     Max35Text                                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.070.001.04 LkSetOrdrId,omitempty"`
	LkSetOrdrSeq    float64                                      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.070.001.04 LkSetOrdrSeq,omitempty"`
	ExctnTp         ExecutionType1Choice                         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.070.001.04 ExctnTp,omitempty"`
	Cdtr            BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.070.001.04 Cdtr,omitempty"`
	CdtrAcct        CashAccount38                                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.070.001.04 CdtrAcct,omitempty"`
	Dbtr            BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.070.001.04 Dbtr,omitempty"`
	DbtrAcct        CashAccount38                                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.070.001.04 DbtrAcct,omitempty"`
	TtlsPerStgOrdr  StandingOrderTotalAmount1                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.070.001.04 TtlsPerStgOrdr,omitempty"`
	ZeroSweepInd    bool                                         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.070.001.04 ZeroSweepInd,omitempty"`
}

type StandingOrderIdentification4 struct {
	Id       Max35Text                                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.070.001.04 Id,omitempty"`
	Acct     CashAccount38                                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.070.001.04 Acct"`
	AcctOwnr BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.070.001.04 AcctOwnr,omitempty"`
}

type StandingOrderOrError5Choice struct {
	Rpt     []StandingOrderReport1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.070.001.04 Rpt"`
	OprlErr []ErrorHandling5       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.070.001.04 OprlErr"`
}

type StandingOrderOrError6Choice struct {
	StgOrdr StandingOrder6   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.070.001.04 StgOrdr"`
	BizErr  []ErrorHandling5 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.070.001.04 BizErr"`
}

// May be one of SLST, SDTL, TAPS, SLSL, SWLS
type StandingOrderQueryType1Code string

type StandingOrderReport1 struct {
	StgOrdrId    StandingOrderIdentification4 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.070.001.04 StgOrdrId"`
	StgOrdrOrErr StandingOrderOrError6Choice  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.070.001.04 StgOrdrOrErr"`
}

type StandingOrderTotalAmount1 struct {
	SetPrdfndOrdr TotalAmountAndCurrency1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.070.001.04 SetPrdfndOrdr"`
	PdgPrdfndOrdr TotalAmountAndCurrency1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.070.001.04 PdgPrdfndOrdr"`
	SetStgOrdr    TotalAmountAndCurrency1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.070.001.04 SetStgOrdr"`
	PdgStgOrdr    TotalAmountAndCurrency1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.070.001.04 PdgStgOrdr"`
}

type StandingOrderType1Choice struct {
	Cd    StandingOrderType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.070.001.04 Cd"`
	Prtry GenericIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.070.001.04 Prtry"`
}

// May be one of USTO, PSTO
type StandingOrderType1Code string

type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.070.001.04 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.070.001.04 Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
}

type TotalAmountAndCurrency1 struct {
	TtlAmt    float64            `xml:"urn:iso:std:iso:20022:tech:xsd:camt.070.001.04 TtlAmt"`
	CdtDbtInd CreditDebitCode    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.070.001.04 CdtDbtInd,omitempty"`
	Ccy       ActiveCurrencyCode `xml:"urn:iso:std:iso:20022:tech:xsd:camt.070.001.04 Ccy,omitempty"`
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

type xsdTime time.Time

func (t *xsdTime) UnmarshalText(text []byte) error {
	return _unmarshalTime(text, (*time.Time)(t), "15:04:05.999999999")
}
func (t xsdTime) MarshalText() ([]byte, error) {
	return _marshalTime((time.Time)(t), "15:04:05.999999999")
}
func (t xsdTime) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if (time.Time)(t).IsZero() {
		return nil
	}
	m, err := t.MarshalText()
	if err != nil {
		return err
	}
	return e.EncodeElement(m, start)
}
func (t xsdTime) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if (time.Time)(t).IsZero() {
		return xml.Attr{}, nil
	}
	m, err := t.MarshalText()
	return xml.Attr{Name: name, Value: string(m)}, err
}
