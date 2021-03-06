// Code generated by main. DO NOT EDIT.

package seev_042_001_09

import (
	"bytes"
	"encoding/xml"
	"time"
)

type AccountIdentification53 struct {
	SfkpgAcct         Max35Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 SfkpgAcct"`
	AcctOwnr          PartyIdentification127Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 AcctOwnr,omitempty"`
	SfkpgPlc          SafekeepingPlaceFormat28Choice     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 SfkpgPlc,omitempty"`
	CorpActnEvtAndBal []CorporateActionEventAndBalance17 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 CorpActnEvtAndBal,omitempty"`
}

type ActiveCurrencyAnd13DecimalAmount struct {
	Value float64            `xml:",chardata"`
	Ccy   ActiveCurrencyCode `xml:"Ccy,attr"`
}

// Must match the pattern [A-Z]{3,3}
type ActiveCurrencyCode string

type AmountPrice3 struct {
	AmtPricTp AmountPriceType1Code             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 AmtPricTp"`
	PricVal   ActiveCurrencyAnd13DecimalAmount `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 PricVal"`
}

// May be one of ACTU, DISC, PLOT, PREM
type AmountPriceType1Code string

// Must match the pattern [A-Z0-9]{4,4}[A-Z]{2,2}[A-Z0-9]{2,2}([A-Z0-9]{3,3}){0,1}
type AnyBICDec2014Identifier string

type BalanceFormat5Choice struct {
	Bal         SignedQuantityFormat7 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 Bal"`
	ElgblBal    SignedQuantityFormat6 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 ElgblBal"`
	NotElgblBal SignedQuantityFormat6 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 NotElgblBal"`
}

type CancelledReason8Choice struct {
	Cd    CancelledStatusReason6Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 Cd"`
	Prtry GenericIdentification30    `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 Prtry"`
}

type CancelledStatus12Choice struct {
	NoSpcfdRsn NoReasonCode              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 NoSpcfdRsn"`
	Rsn        []CancelledStatusReason11 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 Rsn"`
}

type CancelledStatusReason11 struct {
	RsnCd       CancelledReason8Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 RsnCd"`
	AddtlRsnInf Max210Text             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 AddtlRsnInf,omitempty"`
}

// May be one of CANI, CANO, CANS, CSUB, OTHR
type CancelledStatusReason6Code string

type CorporateActionBalance41 struct {
	TtlElgblBal      Quantity17Choice      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 TtlElgblBal"`
	UinstdBal        BalanceFormat5Choice  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 UinstdBal"`
	TtlInstdBalDtls  InstructedBalance11   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 TtlInstdBalDtls"`
	BlckdBal         SignedQuantityFormat6 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 BlckdBal,omitempty"`
	BrrwdBal         SignedQuantityFormat6 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 BrrwdBal,omitempty"`
	CollInBal        SignedQuantityFormat6 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 CollInBal,omitempty"`
	CollOutBal       SignedQuantityFormat6 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 CollOutBal,omitempty"`
	OnLnBal          SignedQuantityFormat6 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 OnLnBal,omitempty"`
	OutForRegnBal    SignedQuantityFormat6 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 OutForRegnBal,omitempty"`
	SttlmPosBal      SignedQuantityFormat6 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 SttlmPosBal,omitempty"`
	StrtPosBal       SignedQuantityFormat6 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 StrtPosBal,omitempty"`
	TradDtPosBal     SignedQuantityFormat6 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 TradDtPosBal,omitempty"`
	InTrnsShipmntBal SignedQuantityFormat6 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 InTrnsShipmntBal,omitempty"`
	RegdBal          SignedQuantityFormat6 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 RegdBal,omitempty"`
	OblgtdBal        SignedQuantityFormat6 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 OblgtdBal,omitempty"`
	PdgDlvryBal      []PendingBalance5     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 PdgDlvryBal,omitempty"`
	PdgRctBal        []PendingBalance5     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 PdgRctBal,omitempty"`
}

type CorporateActionEventAndBalance17 struct {
	GnlInf      EventInformation13       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 GnlInf"`
	UndrlygScty SecurityIdentification19 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 UndrlygScty"`
	Bal         CorporateActionBalance41 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 Bal,omitempty"`
	SplmtryData []SupplementaryData1     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 SplmtryData,omitempty"`
}

type CorporateActionEventDeadlines3 struct {
	EarlyRspnDdln  DateFormat43Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 EarlyRspnDdln,omitempty"`
	RspnDdln       DateFormat44Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 RspnDdln,omitempty"`
	MktDdln        DateFormat43Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 MktDdln,omitempty"`
	PrtctDdln      DateFormat43Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 PrtctDdln,omitempty"`
	CoverPrtctDdln DateFormat43Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 CoverPrtctDdln,omitempty"`
}

// May be one of ACTV, ATTI, BRUP, DFLT, BONU, EXRI, CAPD, CAPG, CAPI, DRCA, DVCA, CHAN, COOP, CLSA, CONS, CONV, CREV, DECR, DETI, DSCL, DVOP, DRIP, DRAW, DTCH, EXOF, REDM, MCAL, INCR, PPMT, INTR, RHDI, LIQU, EXTM, MRGR, NOOF, CERT, ODLT, OTHR, PARI, PCAL, PRED, PINK, PLAC, PDEF, PRIO, BPUT, REDO, REMK, BIDS, SPLR, RHTS, DVSC, SHPR, SMAL, SOFF, DVSE, SPLF, TREC, TEND, DLST, SUSP, EXWA, WTRC, WRTH
type CorporateActionEventType29Code string

type CorporateActionEventType85Choice struct {
	Cd    CorporateActionEventType29Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 Cd"`
	Prtry GenericIdentification30        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 Prtry"`
}

type CorporateActionInstructionStatementReportV09 struct {
	Pgntn           Pagination1               `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 Pgntn"`
	StmtGnlDtls     Statement72               `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 StmtGnlDtls"`
	AcctAndStmtDtls []AccountIdentification53 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 AcctAndStmtDtls"`
	SplmtryData     []SupplementaryData1      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 SplmtryData,omitempty"`
}

// May be one of MAND, CHOS, VOLU
type CorporateActionMandatoryVoluntary1Code string

type CorporateActionMandatoryVoluntary3Choice struct {
	Cd    CorporateActionMandatoryVoluntary1Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 Cd"`
	Prtry GenericIdentification30                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 Prtry"`
}

// May be one of ABST, BSPL, BUYA, CASE, CASH, CEXC, CONN, CONY, CTEN, EXER, LAPS, MPUT, NOAC, NOQU, OFFR, OTHR, OVER, QINV, SECU, SLLE, PRUN
type CorporateActionOption11Code string

type CorporateActionOption30Choice struct {
	Cd    CorporateActionOption11Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 Cd"`
	Prtry GenericIdentification30     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 Prtry"`
}

// May be one of MASE, SAME
type CorporateActionStatementReportingType1Code string

// May be one of MISS, ALLL, BALO, BALI
type CorporateActionStatementType2Code string

// Must match the pattern [A-Z]{2,2}
type CountryCode string

type DateAndDateTime2Choice struct {
	Dt   ISODate     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 Dt"`
	DtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 DtTm"`
}

type DateCode19Choice struct {
	Cd    DateType8Code           `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 Cd"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 Prtry"`
}

type DateCode21Choice struct {
	Cd    DateType7Code           `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 Cd"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 Prtry"`
}

type DateCodeAndTimeFormat3 struct {
	DtCd DateCode21Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 DtCd"`
	Tm   ISOTime          `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 Tm"`
}

type DateFormat43Choice struct {
	Dt   DateAndDateTime2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 Dt"`
	DtCd DateCode19Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 DtCd"`
}

type DateFormat44Choice struct {
	Dt        DateAndDateTime2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 Dt"`
	DtCdAndTm DateCodeAndTimeFormat3 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 DtCdAndTm"`
	DtCd      DateCode19Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 DtCd"`
}

type DateOrDateTimePeriod1Choice struct {
	Dt   DatePeriod2     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 Dt"`
	DtTm DateTimePeriod1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 DtTm"`
}

type DatePeriod2 struct {
	FrDt ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 FrDt"`
	ToDt ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 ToDt"`
}

type DateTimePeriod1 struct {
	FrDtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 FrDtTm"`
	ToDtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 ToDtTm"`
}

// May be one of ONGO
type DateType7Code string

// May be one of UKWN, ONGO
type DateType8Code string

type DefaultProcessingOrStandingInstruction1Choice struct {
	DfltOptnInd bool `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 DfltOptnInd"`
	StgInstrInd bool `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 StgInstrInd"`
}

// May be one of FREE, APMT
type DeliveryReceiptType2Code string

type Document struct {
	CorpActnInstrStmtRpt CorporateActionInstructionStatementReportV09 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 CorpActnInstrStmtRpt"`
}

// May be one of YEAR, ADHO, MNTH, DAIL, INDA, WEEK
type EventFrequency4Code string

type EventInformation13 struct {
	CorpActnEvtId      Max35Text                                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 CorpActnEvtId"`
	OffclCorpActnEvtId Max35Text                                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 OffclCorpActnEvtId,omitempty"`
	EvtTp              CorporateActionEventType85Choice         `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 EvtTp"`
	MndtryVlntryEvtTp  CorporateActionMandatoryVoluntary3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 MndtryVlntryEvtTp"`
	LastNtfctnId       NotificationIdentification5              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 LastNtfctnId,omitempty"`
}

// Must match the pattern [0-9]{3}
type Exact3NumericText string

// Must match the pattern [a-zA-Z0-9]{4}
type Exact4AlphaNumericText string

// Must be at least 1 items long
type ExternalFinancialInstrumentIdentificationType1Code string

type FinancialInstrumentQuantity1Choice struct {
	Unit     float64 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 Unit"`
	FaceAmt  float64 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 FaceAmt"`
	AmtsdVal float64 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 AmtsdVal"`
}

type Frequency25Choice struct {
	Cd    EventFrequency4Code     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 Cd"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 Prtry"`
}

type GenericIdentification30 struct {
	Id      Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 Id"`
	Issr    Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 Issr"`
	SchmeNm Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 SchmeNm,omitempty"`
}

type GenericIdentification36 struct {
	Id      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 Id"`
	Issr    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 Issr"`
	SchmeNm Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 SchmeNm,omitempty"`
}

type GenericIdentification78 struct {
	Tp GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 Tp"`
	Id Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 Id,omitempty"`
}

// Must match the pattern [A-Z]{2,2}[A-Z0-9]{9,9}[0-9]{1,1}
type ISINOct2015Identifier string

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

type IdentificationSource3Choice struct {
	Cd    ExternalFinancialInstrumentIdentificationType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 Cd"`
	Prtry Max35Text                                          `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 Prtry"`
}

type InstructedBalance11 struct {
	TtlInstdBal       BalanceFormat5Choice                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 TtlInstdBal"`
	TtlAccptdInstrBal SignedQuantityFormat6               `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 TtlAccptdInstrBal,omitempty"`
	TtlCancInstrBal   SignedQuantityFormat6               `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 TtlCancInstrBal,omitempty"`
	TtlPdgInstrBal    SignedQuantityFormat6               `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 TtlPdgInstrBal,omitempty"`
	TtlRjctdInstrBal  SignedQuantityFormat6               `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 TtlRjctdInstrBal,omitempty"`
	TtlPrtctInstrBal  SignedQuantityFormat6               `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 TtlPrtctInstrBal,omitempty"`
	OptnDtls          []InstructedCorporateActionOption12 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 OptnDtls,omitempty"`
}

type InstructedCorporateActionOption12 struct {
	OptnNb             Exact3NumericText                             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 OptnNb,omitempty"`
	OptnTp             CorporateActionOption30Choice                 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 OptnTp"`
	InstdBal           BalanceFormat5Choice                          `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 InstdBal"`
	DfltActn           DefaultProcessingOrStandingInstruction1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 DfltActn,omitempty"`
	OptnAccptdInstdBal SignedQuantityFormat6                         `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 OptnAccptdInstdBal,omitempty"`
	OptnCancInstrBal   SignedQuantityFormat6                         `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 OptnCancInstrBal,omitempty"`
	OptnPdgInstrBal    SignedQuantityFormat6                         `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 OptnPdgInstrBal,omitempty"`
	OptnRjctdInstrBal  SignedQuantityFormat6                         `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 OptnRjctdInstrBal,omitempty"`
	OptnPrtctInstrBal  SignedQuantityFormat6                         `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 OptnPrtctInstrBal,omitempty"`
	EvtDdlns           CorporateActionEventDeadlines3                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 EvtDdlns"`
	OptnInstrDtls      []OptionInstructionDetails3                   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 OptnInstrDtls,omitempty"`
}

type InstructionProcessingStatus32Choice struct {
	Accptd             NoSpecifiedReason1               `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 Accptd"`
	Canc               CancelledStatus12Choice          `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 Canc"`
	AccptdForFrthrPrcg NoSpecifiedReason1               `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 AccptdForFrthrPrcg"`
	Rjctd              RejectedStatus26Choice           `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 Rjctd"`
	Pdg                NoSpecifiedReason1               `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 Pdg"`
	PdgCxl             PendingCancellationStatus7Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 PdgCxl"`
	Cvrd               NoSpecifiedReason1               `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 Cvrd"`
	Ucvrd              NoSpecifiedReason1               `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 Ucvrd"`
}

// Must be at least 1 items long
type Max140Text string

// Must be at least 1 items long
type Max15Text string

// Must be at least 1 items long
type Max16Text string

// Must be at least 1 items long
type Max210Text string

// Must be at least 1 items long
type Max350Text string

// Must be at least 1 items long
type Max35Text string

// Must match the pattern [0-9]{1,3}
type Max3NumericText string

// Must match the pattern [0-9]{1,5}
type Max5NumericText string

// May be one of NORE
type NoReasonCode string

type NoSpecifiedReason1 struct {
	NoSpcfdRsn NoReasonCode `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 NoSpcfdRsn"`
}

type NotificationIdentification5 struct {
	Id      Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 Id"`
	CreDtTm DateAndDateTime2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 CreDtTm,omitempty"`
}

type OptionInstructionDetails3 struct {
	InstrId      Max15Text                           `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 InstrId"`
	InstrSeqNb   Max3NumericText                     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 InstrSeqNb,omitempty"`
	PrtctInd     ProtectTransactionType2Code         `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 PrtctInd,omitempty"`
	InstrQty     FinancialInstrumentQuantity1Choice  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 InstrQty"`
	InstrDt      ISODate                             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 InstrDt"`
	PrtctDt      ISODate                             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 PrtctDt,omitempty"`
	CoverPrtctDt ISODate                             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 CoverPrtctDt,omitempty"`
	BidPric      PriceFormat45Choice                 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 BidPric,omitempty"`
	CondlQty     FinancialInstrumentQuantity1Choice  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 CondlQty,omitempty"`
	CstmrRef     string                              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 CstmrRef,omitempty"`
	InstrNrrtv   Max350Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 InstrNrrtv,omitempty"`
	InstrSts     InstructionProcessingStatus32Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 InstrSts"`
}

type OriginalAndCurrentQuantities6 struct {
	ShrtLngPos ShortLong1Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 ShrtLngPos"`
	FaceAmt    float64        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 FaceAmt"`
	AmtsdVal   float64        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 AmtsdVal"`
}

type OtherIdentification1 struct {
	Id  Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 Id"`
	Sfx Max16Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 Sfx,omitempty"`
	Tp  IdentificationSource3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 Tp"`
}

type Pagination1 struct {
	PgNb      Max5NumericText `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 PgNb"`
	LastPgInd bool            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 LastPgInd"`
}

type PartyIdentification127Choice struct {
	AnyBIC  AnyBICDec2014Identifier `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 AnyBIC"`
	PrtryId GenericIdentification36 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 PrtryId"`
}

type PendingBalance5 struct {
	Bal    SignedQuantityFormat6               `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 Bal"`
	PdgTxs []SettlementTypeAndIdentification25 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 PdgTxs,omitempty"`
}

type PendingCancellationReason5Choice struct {
	Cd    PendingCancellationReason5Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 Cd"`
	Prtry GenericIdentification30        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 Prtry"`
}

// May be one of ADEA, DQUA, DQCS, LATE, OTHR
type PendingCancellationReason5Code string

type PendingCancellationStatus7Choice struct {
	NotSpcfdRsn NoReasonCode                       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 NotSpcfdRsn"`
	Rsn         []PendingCancellationStatusReason7 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 Rsn"`
}

type PendingCancellationStatusReason7 struct {
	RsnCd       PendingCancellationReason5Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 RsnCd"`
	AddtlRsnInf Max210Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 AddtlRsnInf,omitempty"`
}

type PercentagePrice1 struct {
	PctgPricTp PriceRateType3Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 PctgPricTp"`
	PricVal    float64            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 PricVal"`
}

type PriceFormat45Choice struct {
	PctgPric     PercentagePrice1     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 PctgPric"`
	AmtPric      AmountPrice3         `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 AmtPric"`
	NotSpcfdPric PriceValueType10Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 NotSpcfdPric"`
}

// May be one of DISC, PREM, PRCT, YIEL
type PriceRateType3Code string

// May be one of UKWN
type PriceValueType10Code string

type ProprietaryQuantity7 struct {
	ShrtLngPos ShortLong1Code         `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 ShrtLngPos,omitempty"`
	Qty        float64                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 Qty"`
	QtyTp      Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 QtyTp"`
	Issr       Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 Issr"`
	SchmeNm    Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 SchmeNm,omitempty"`
}

type ProprietaryQuantity8 struct {
	Qty     float64                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 Qty"`
	QtyTp   Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 QtyTp"`
	Issr    Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 Issr"`
	SchmeNm Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 SchmeNm,omitempty"`
}

// May be one of PROT, COVP, COVR
type ProtectTransactionType2Code string

type Quantity17Choice struct {
	QtyChc   Quantity18Choice     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 QtyChc"`
	PrtryQty ProprietaryQuantity7 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 PrtryQty"`
}

type Quantity18Choice struct {
	OrgnlAndCurFaceAmt OriginalAndCurrentQuantities6 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 OrgnlAndCurFaceAmt"`
	SgndQty            SignedQuantityFormat6         `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 SgndQty"`
}

type Quantity19Choice struct {
	Qty      FinancialInstrumentQuantity1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 Qty"`
	PrtryQty ProprietaryQuantity8               `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 PrtryQty"`
}

type RejectedReason25Choice struct {
	Cd    RejectionReason49Code   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 Cd"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 Prtry"`
}

type RejectedStatus26Choice struct {
	NoSpcfdRsn NoReasonCode             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 NoSpcfdRsn"`
	Rsn        []RejectedStatusReason24 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 Rsn"`
}

type RejectedStatusReason24 struct {
	RsnCd       RejectedReason25Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 RsnCd"`
	AddtlRsnInf Max210Text             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 AddtlRsnInf,omitempty"`
}

// May be one of ADEA, CERT, INVA, OPTY, ULNK, DSEC, LACK, LATE, NMTY, FULL, CANC, INTV, OPNM, OTHR, DQUA, REFT, SAFE, EVNM, DQCS, DQCC, DQAM, IRDQ, DQBV, DQBI, DCAN, DPRG, INIR, SHAR
type RejectionReason49Code string

// May be one of CUST, ICSD, NCSD, SHHE
type SafekeepingPlace1Code string

// May be one of SHHE, ALLP
type SafekeepingPlace2Code string

type SafekeepingPlaceFormat28Choice struct {
	Id      SafekeepingPlaceTypeAndText6           `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 Id"`
	Ctry    CountryCode                            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 Ctry"`
	TpAndId SafekeepingPlaceTypeAndIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 TpAndId"`
	Prtry   GenericIdentification78                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 Prtry"`
}

type SafekeepingPlaceTypeAndIdentification1 struct {
	SfkpgPlcTp SafekeepingPlace1Code   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 SfkpgPlcTp"`
	Id         AnyBICDec2014Identifier `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 Id"`
}

type SafekeepingPlaceTypeAndText6 struct {
	SfkpgPlcTp SafekeepingPlace2Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 SfkpgPlcTp"`
	Id         Max35Text             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 Id,omitempty"`
}

type SecurityIdentification19 struct {
	ISIN   ISINOct2015Identifier  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 ISIN,omitempty"`
	OthrId []OtherIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 OthrId,omitempty"`
	Desc   Max140Text             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 Desc,omitempty"`
}

type SettlementTypeAndIdentification25 struct {
	Pmt     DeliveryReceiptType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 Pmt"`
	TxId    Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 TxId"`
	SttlmDt DateAndDateTime2Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 SttlmDt,omitempty"`
}

// May be one of SHOR, LONG
type ShortLong1Code string

type SignedQuantityFormat6 struct {
	ShrtLngPos ShortLong1Code                     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 ShrtLngPos"`
	Qty        FinancialInstrumentQuantity1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 Qty"`
}

type SignedQuantityFormat7 struct {
	ShrtLngPos ShortLong1Code   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 ShrtLngPos"`
	QtyChc     Quantity19Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 QtyChc"`
}

type Statement72 struct {
	StmtTp        CorporateActionStatementType2Code          `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 StmtTp"`
	RptgTp        CorporateActionStatementReportingType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 RptgTp"`
	StmtId        Max35Text                                  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 StmtId"`
	InstrAggtnPrd DatePeriod2                                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 InstrAggtnPrd,omitempty"`
	RptNb         Max5NumericText                            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 RptNb,omitempty"`
	StmtDtTm      DateAndDateTime2Choice                     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 StmtDtTm"`
	Frqcy         Frequency25Choice                          `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 Frqcy"`
	UpdTp         UpdateType15Choice                         `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 UpdTp"`
	ActvtyInd     bool                                       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 ActvtyInd"`
	NtfctnDdlnPrd DateOrDateTimePeriod1Choice                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 NtfctnDdlnPrd,omitempty"`
}

// May be one of COMP, DELT
type StatementUpdateType1Code string

type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
}

type UpdateType15Choice struct {
	Cd    StatementUpdateType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 Cd"`
	Prtry GenericIdentification30  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.09 Prtry"`
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
