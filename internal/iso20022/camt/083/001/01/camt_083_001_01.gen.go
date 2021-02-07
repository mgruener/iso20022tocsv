// Code generated by main. DO NOT EDIT.

package camt_083_001_01

import (
	"bytes"
	"encoding/xml"
	"time"
)

type AccountIdentification4Choice struct {
	IBAN IBAN2007Identifier            `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 IBAN"`
	Othr GenericAccountIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 Othr"`
}

type AccountSchemeName1Choice struct {
	Cd    ExternalAccountIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 Cd"`
	Prtry Max35Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 Prtry"`
}

type AcknowledgedAcceptedStatus24Choice struct {
	NoSpcfdRsn NoReasonCode              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 NoSpcfdRsn"`
	Rsn        []AcknowledgementReason12 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 Rsn"`
}

type AcknowledgementReason12 struct {
	Cd          AcknowledgementReason15Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 Cd"`
	AddtlRsnInf Max210Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 AddtlRsnInf,omitempty"`
}

type AcknowledgementReason15Choice struct {
	Cd    AcknowledgementReason3Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 Cd"`
	Prtry GenericIdentification30    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 Prtry"`
}

// May be one of ADEA, SMPG, OTHR
type AcknowledgementReason3Code string

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
	Cd    AddressType2Code        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 Cd"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 Prtry"`
}

type Amount2Choice struct {
	AmtWthtCcy float64                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 AmtWthtCcy"`
	AmtWthCcy  ActiveCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 AmtWthCcy"`
}

type AmountAndDirection5 struct {
	Amt    ActiveCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 Amt"`
	CdtDbt CreditDebitCode         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 CdtDbt,omitempty"`
}

type AmountAndQuantityBreakdown1 struct {
	LotNb       GenericIdentification37            `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 LotNb,omitempty"`
	LotAmt      AmountAndDirection5                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 LotAmt,omitempty"`
	LotQty      FinancialInstrumentQuantity1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 LotQty,omitempty"`
	CshSubBalTp GenericIdentification30            `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 CshSubBalTp,omitempty"`
}

// Must match the pattern [A-Z0-9]{4,4}[A-Z]{2,2}[A-Z0-9]{2,2}([A-Z0-9]{3,3}){0,1}
type AnyBICDec2014Identifier string

// Must match the pattern [A-Z0-9]{4,4}[A-Z]{2,2}[A-Z0-9]{2,2}([A-Z0-9]{3,3}){0,1}
type BICFIDec2014Identifier string

type BranchAndFinancialInstitutionIdentification6 struct {
	FinInstnId FinancialInstitutionIdentification18 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 FinInstnId"`
	BrnchId    BranchData3                          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 BrnchId,omitempty"`
}

type BranchData3 struct {
	Id      Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 Id,omitempty"`
	LEI     LEIIdentifier   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 LEI,omitempty"`
	Nm      Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 Nm,omitempty"`
	PstlAdr PostalAddress24 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 PstlAdr,omitempty"`
}

type CancellationReason10 struct {
	Cd          CancellationReason21Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 Cd"`
	AddtlRsnInf Max210Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 AddtlRsnInf,omitempty"`
}

type CancellationReason21Choice struct {
	Cd    CancelledStatusReason5Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 Cd"`
	Prtry GenericIdentification30    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 Prtry"`
}

type CancellationStatus15Choice struct {
	NoSpcfdRsn NoReasonCode           `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 NoSpcfdRsn"`
	Rsn        []CancellationReason10 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 Rsn"`
}

// May be one of CANI, OTHR
type CancelledStatusReason5Code string

type CashAccount38 struct {
	Id   AccountIdentification4Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 Id"`
	Tp   CashAccountType2Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 Tp,omitempty"`
	Ccy  ActiveOrHistoricCurrencyCode `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 Ccy,omitempty"`
	Nm   Max70Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 Nm,omitempty"`
	Prxy ProxyAccountIdentification1  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 Prxy,omitempty"`
}

type CashAccountType2Choice struct {
	Cd    ExternalCashAccountType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 Cd"`
	Prtry Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 Prtry"`
}

type CashBalanceType3Choice struct {
	Cd    ExternalBalanceType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 Cd"`
	Prtry GenericIdentification30  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 Prtry"`
}

type CashSubBalanceTypeAndQuantityBreakdown3 struct {
	Tp        CashBalanceType3Choice        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 Tp"`
	QtyBrkdwn []AmountAndQuantityBreakdown1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 QtyBrkdwn,omitempty"`
}

type ClearingSystemIdentification2Choice struct {
	Cd    ExternalClearingSystemIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 Cd"`
	Prtry Max35Text                                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 Prtry"`
}

type ClearingSystemMemberIdentification2 struct {
	ClrSysId ClearingSystemIdentification2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 ClrSysId,omitempty"`
	MmbId    Max35Text                           `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 MmbId"`
}

// May be one of CODU, COPY, DUPL
type CopyDuplicate1Code string

// Must match the pattern [A-Z]{2,2}
type CountryCode string

// May be one of CRDT, DBIT
type CreditDebitCode string

type DateAndDateTime2Choice struct {
	Dt   ISODate     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 Dt"`
	DtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 DtTm"`
}

type DateTimePeriod1 struct {
	FrDtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 FrDtTm"`
	ToDtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 ToDtTm"`
}

type DeniedReason11 struct {
	Cd          DeniedReason16Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 Cd"`
	AddtlRsnInf Max210Text           `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 AddtlRsnInf,omitempty"`
}

type DeniedReason16Choice struct {
	Cd    DeniedReason4Code       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 Cd"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 Prtry"`
}

// May be one of ADEA, DCAN, DPRG, DREP, DSET, LATE, OTHR, CDRG, CDCY, CDRE
type DeniedReason4Code string

type DeniedStatus16Choice struct {
	NoSpcfdRsn NoReasonCode     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 NoSpcfdRsn"`
	Rsn        []DeniedReason11 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 Rsn"`
}

type Document struct {
	IntraBalMvmntCxlRpt IntraBalanceMovementCancellationReportV01 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 IntraBalMvmntCxlRpt"`
}

type DocumentIdentification51 struct {
	Id       Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 Id"`
	CreDtTm  DateAndDateTime2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 CreDtTm,omitempty"`
	CpyDplct CopyDuplicate1Code     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 CpyDplct,omitempty"`
	MsgOrgtr PartyIdentification136 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 MsgOrgtr,omitempty"`
	MsgRcpt  PartyIdentification136 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 MsgRcpt,omitempty"`
}

type ErrorHandling3Choice struct {
	Cd    ExternalSystemErrorHandling1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 Cd"`
	Prtry Max35Text                        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 Prtry"`
}

type ErrorHandling5 struct {
	Err  ErrorHandling3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 Err"`
	Desc Max140Text           `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 Desc,omitempty"`
}

// May be one of YEAR, ADHO, MNTH, DAIL, INDA, WEEK, SEMI, QUTR, TOMN, TOWK, TWMN, OVNG, ONDE
type EventFrequency7Code string

// Must match the pattern [0-9]{3}
type Exact3NumericText string

// Must match the pattern [a-zA-Z0-9]{4}
type Exact4AlphaNumericText string

// Must match the pattern [0-9]{4}
type Exact4NumericText string

// Must match the pattern [0-9]{5}
type Exact5NumericText string

// Must be at least 1 items long
type ExternalAccountIdentification1Code string

// Must be at least 1 items long
type ExternalBalanceType1Code string

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

type FinancialIdentificationSchemeName1Choice struct {
	Cd    ExternalFinancialInstitutionIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 Cd"`
	Prtry Max35Text                                       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 Prtry"`
}

type FinancialInstitutionIdentification18 struct {
	BICFI       BICFIDec2014Identifier              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 BICFI,omitempty"`
	ClrSysMmbId ClearingSystemMemberIdentification2 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 ClrSysMmbId,omitempty"`
	LEI         LEIIdentifier                       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 LEI,omitempty"`
	Nm          Max140Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 Nm,omitempty"`
	PstlAdr     PostalAddress24                     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 PstlAdr,omitempty"`
	Othr        GenericFinancialIdentification1     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 Othr,omitempty"`
}

type FinancialInstrumentQuantity1Choice struct {
	Unit     float64 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 Unit"`
	FaceAmt  float64 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 FaceAmt"`
	AmtsdVal float64 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 AmtsdVal"`
}

type Frequency22Choice struct {
	Cd    EventFrequency7Code     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 Cd"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 Prtry"`
}

type GenericAccountIdentification1 struct {
	Id      Max34Text                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 Id"`
	SchmeNm AccountSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 SchmeNm,omitempty"`
	Issr    Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 Issr,omitempty"`
}

type GenericFinancialIdentification1 struct {
	Id      Max35Text                                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 Id"`
	SchmeNm FinancialIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 SchmeNm,omitempty"`
	Issr    Max35Text                                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 Issr,omitempty"`
}

type GenericIdentification30 struct {
	Id      Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 Id"`
	Issr    Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 Issr"`
	SchmeNm Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 SchmeNm,omitempty"`
}

type GenericIdentification36 struct {
	Id      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 Id"`
	Issr    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 Issr"`
	SchmeNm Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 SchmeNm,omitempty"`
}

type GenericIdentification37 struct {
	Id   Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 Id"`
	Issr Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 Issr,omitempty"`
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

type IntraBalance5 struct {
	SttlmAmt           Amount2Choice                           `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 SttlmAmt"`
	SttlmDt            DateAndDateTime2Choice                  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 SttlmDt"`
	BalFr              CashSubBalanceTypeAndQuantityBreakdown3 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 BalFr"`
	BalTo              CashSubBalanceTypeAndQuantityBreakdown3 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 BalTo"`
	CshSubBalId        GenericIdentification37                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 CshSubBalId,omitempty"`
	Prty               PriorityNumeric4Choice                  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 Prty,omitempty"`
	InstrPrcgAddtlDtls Max350Text                              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 InstrPrcgAddtlDtls,omitempty"`
}

type IntraBalanceCancellation5 struct {
	CshAcct     CashAccount38                                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 CshAcct,omitempty"`
	CshAcctOwnr SystemPartyIdentification8                   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 CshAcctOwnr,omitempty"`
	CshAcctSvcr BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 CshAcctSvcr,omitempty"`
	PrcgSts     ProcessingStatus69Choice                     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 PrcgSts,omitempty"`
	Cxl         []IntraBalanceCancellation6                  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 Cxl"`
}

type IntraBalanceCancellation6 struct {
	CshAcct         CashAccount38                                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 CshAcct,omitempty"`
	CshAcctOwnr     SystemPartyIdentification8                   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 CshAcctOwnr,omitempty"`
	CshAcctSvcr     BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 CshAcctSvcr,omitempty"`
	PrcgSts         ProcessingStatus69Choice                     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 PrcgSts,omitempty"`
	ReqRef          Max35Text                                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 ReqRef"`
	StsDt           ISODateTime                                  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 StsDt,omitempty"`
	TxId            References14                                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 TxId,omitempty"`
	UndrlygIntraBal IntraBalance5                                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 UndrlygIntraBal,omitempty"`
}

type IntraBalanceMovementCancellationReportV01 struct {
	Id          DocumentIdentification51              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 Id,omitempty"`
	Pgntn       Pagination1                           `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 Pgntn"`
	RptGnlDtls  IntraBalanceReport5                   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 RptGnlDtls"`
	RptOrErr    IntraBalanceOrOperationalError9Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 RptOrErr,omitempty"`
	SplmtryData []SupplementaryData1                  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 SplmtryData,omitempty"`
}

type IntraBalanceOrOperationalError9Choice struct {
	Cxls    []IntraBalanceCancellation5 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 Cxls"`
	OprlErr []ErrorHandling5            `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 OprlErr"`
}

type IntraBalanceReport5 struct {
	RptNb     Number3Choice             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 RptNb,omitempty"`
	QryRef    Max35Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 QryRef,omitempty"`
	RptId     Max35Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 RptId,omitempty"`
	RptDtTm   DateAndDateTime2Choice    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 RptDtTm,omitempty"`
	RptPrd    Period7Choice             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 RptPrd,omitempty"`
	QryTp     MovementResponseType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 QryTp,omitempty"`
	Frqcy     Frequency22Choice         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 Frqcy,omitempty"`
	UpdTp     UpdateType15Choice        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 UpdTp"`
	ActvtyInd bool                      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 ActvtyInd"`
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
type Max210Text string

// Must be at least 1 items long
type Max34Text string

// Must be at least 1 items long
type Max350Text string

// Must be at least 1 items long
type Max35Text string

// Must match the pattern [0-9]{1,5}
type Max5NumericText string

// Must be at least 1 items long
type Max70Text string

// May be one of FULL, STTS
type MovementResponseType1Code string

type NameAndAddress5 struct {
	Nm  Max350Text     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 Nm"`
	Adr PostalAddress1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 Adr,omitempty"`
}

// May be one of NORE
type NoReasonCode string

type Number3Choice struct {
	Shrt Exact3NumericText `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 Shrt"`
	Lng  Exact5NumericText `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 Lng"`
}

type Pagination1 struct {
	PgNb      Max5NumericText `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 PgNb"`
	LastPgInd bool            `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 LastPgInd"`
}

type PartyIdentification120Choice struct {
	AnyBIC   AnyBICDec2014Identifier `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 AnyBIC"`
	PrtryId  GenericIdentification36 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 PrtryId"`
	NmAndAdr NameAndAddress5         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 NmAndAdr"`
}

type PartyIdentification136 struct {
	Id  PartyIdentification120Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 Id"`
	LEI LEIIdentifier                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 LEI,omitempty"`
}

type PendingReason17 struct {
	Cd          PendingReason30Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 Cd"`
	AddtlRsnInf Max210Text            `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 AddtlRsnInf,omitempty"`
}

type PendingReason30Choice struct {
	Cd    PendingReason9Code      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 Cd"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 Prtry"`
}

// May be one of ADEA, CONF, OTHR, CDRG, CDCY, CDRE, CDAC, INBC
type PendingReason9Code string

type PendingStatus39Choice struct {
	NoSpcfdRsn NoReasonCode      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 NoSpcfdRsn"`
	Rsn        []PendingReason17 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 Rsn"`
}

type Period2 struct {
	FrDt ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 FrDt"`
	ToDt ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 ToDt"`
}

type Period7Choice struct {
	FrDtTmToDtTm DateTimePeriod1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 FrDtTmToDtTm"`
	FrDtToDt     Period2         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 FrDtToDt"`
}

type PostalAddress1 struct {
	AdrTp       AddressType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 AdrTp,omitempty"`
	AdrLine     []Max70Text      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 AdrLine,omitempty"`
	StrtNm      Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 StrtNm,omitempty"`
	BldgNb      Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 BldgNb,omitempty"`
	PstCd       Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 PstCd,omitempty"`
	TwnNm       Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 TwnNm,omitempty"`
	CtrySubDvsn Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 CtrySubDvsn,omitempty"`
	Ctry        CountryCode      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 Ctry"`
}

type PostalAddress24 struct {
	AdrTp       AddressType3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 AdrTp,omitempty"`
	Dept        Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 Dept,omitempty"`
	SubDept     Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 SubDept,omitempty"`
	StrtNm      Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 StrtNm,omitempty"`
	BldgNb      Max16Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 BldgNb,omitempty"`
	BldgNm      Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 BldgNm,omitempty"`
	Flr         Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 Flr,omitempty"`
	PstBx       Max16Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 PstBx,omitempty"`
	Room        Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 Room,omitempty"`
	PstCd       Max16Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 PstCd,omitempty"`
	TwnNm       Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 TwnNm,omitempty"`
	TwnLctnNm   Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 TwnLctnNm,omitempty"`
	DstrctNm    Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 DstrctNm,omitempty"`
	CtrySubDvsn Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 CtrySubDvsn,omitempty"`
	Ctry        CountryCode        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 Ctry,omitempty"`
	AdrLine     []Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 AdrLine,omitempty"`
}

type PriorityNumeric4Choice struct {
	Nmrc  Exact4NumericText       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 Nmrc"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 Prtry"`
}

type ProcessingStatus69Choice struct {
	PdgCxl     PendingStatus39Choice              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 PdgCxl"`
	Rjctd      RejectionOrRepairStatus39Choice    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 Rjctd"`
	Rpr        RejectionOrRepairStatus39Choice    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 Rpr"`
	AckdAccptd AcknowledgedAcceptedStatus24Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 AckdAccptd"`
	Prtry      ProprietaryStatusAndReason6        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 Prtry"`
	Dnd        DeniedStatus16Choice               `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 Dnd"`
	Canc       CancellationStatus15Choice         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 Canc"`
}

type ProprietaryReason4 struct {
	Rsn         GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 Rsn,omitempty"`
	AddtlRsnInf Max210Text              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 AddtlRsnInf,omitempty"`
}

type ProprietaryStatusAndReason6 struct {
	PrtrySts GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 PrtrySts"`
	PrtryRsn []ProprietaryReason4    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 PrtryRsn,omitempty"`
}

type ProxyAccountIdentification1 struct {
	Tp ProxyAccountType1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 Tp,omitempty"`
	Id Max2048Text             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 Id"`
}

type ProxyAccountType1Choice struct {
	Cd    ExternalProxyAccountType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 Cd"`
	Prtry Max35Text                     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 Prtry"`
}

type References14 struct {
	AcctOwnrTxId      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 AcctOwnrTxId,omitempty"`
	AcctSvcrTxId      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 AcctSvcrTxId,omitempty"`
	MktInfrstrctrTxId Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 MktInfrstrctrTxId,omitempty"`
	PrcrTxId          Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 PrcrTxId,omitempty"`
	PoolId            Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 PoolId,omitempty"`
}

type RejectionAndRepairReason33Choice struct {
	Cd    RejectionReason34Code   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 Cd"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 Prtry"`
}

type RejectionOrRepairReason33 struct {
	Cd          RejectionAndRepairReason33Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 Cd"`
	AddtlRsnInf Max210Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 AddtlRsnInf,omitempty"`
}

type RejectionOrRepairStatus39Choice struct {
	NoSpcfdRsn NoReasonCode                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 NoSpcfdRsn"`
	Rsn        []RejectionOrRepairReason33 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 Rsn"`
}

// May be one of ADEA, LATE, CASH, NRGM, NRGN, OTHR, REFE
type RejectionReason34Code string

// May be one of COMP, DELT
type StatementUpdateType1Code string

type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
}

type SystemPartyIdentification8 struct {
	Id           PartyIdentification136 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 Id"`
	RspnsblPtyId PartyIdentification136 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 RspnsblPtyId,omitempty"`
}

type UpdateType15Choice struct {
	Cd    StatementUpdateType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 Cd"`
	Prtry GenericIdentification30  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.083.001.01 Prtry"`
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
