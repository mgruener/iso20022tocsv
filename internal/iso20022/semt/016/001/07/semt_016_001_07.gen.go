// Code generated by main. DO NOT EDIT.

package semt_016_001_07

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

type ActiveOrHistoricCurrencyAnd13DecimalAmount struct {
	Value float64                      `xml:",chardata"`
	Ccy   ActiveOrHistoricCurrencyCode `xml:"Ccy,attr"`
}

type ActiveOrHistoricCurrencyAndAmount struct {
	Value float64                      `xml:",chardata"`
	Ccy   ActiveOrHistoricCurrencyCode `xml:"Ccy,attr"`
}

// Must match the pattern [A-Z]{3,3}
type ActiveOrHistoricCurrencyCode string

type AmountAndDirection44 struct {
	Amt                 ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 Amt"`
	CdtDbtInd           CreditDebitCode                   `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 CdtDbtInd,omitempty"`
	OrgnlCcyAndOrdrdAmt ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 OrgnlCcyAndOrdrdAmt,omitempty"`
	FXDtls              ForeignExchangeTerms23            `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 FXDtls,omitempty"`
}

// Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type AnyBICIdentifier string

// Must match the pattern [A-Z]{6,6}
type CFIOct2015Identifier string

type ClassificationType32Choice struct {
	ClssfctnFinInstrm CFIOct2015Identifier    `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 ClssfctnFinInstrm"`
	AltrnClssfctn     GenericIdentification36 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 AltrnClssfctn"`
}

// May be one of ACTV, ATTI, BRUP, DFLT, BONU, EXRI, CAPD, CAPG, CAPI, DRCA, DVCA, CHAN, COOP, CLSA, CONS, CONV, CREV, DECR, DETI, DSCL, DVOP, DRIP, DRAW, DTCH, EXOF, REDM, MCAL, INCR, PPMT, INTR, PRII, RHDI, LIQU, EXTM, MRGR, NOOF, CERT, ODLT, OTHR, PARI, PCAL, PRED, PINK, PLAC, PDEF, PRIO, BPUT, REDO, REMK, BIDS, SPLR, RHTS, DVSC, SHPR, SMAL, SOFF, DVSE, SPLF, TREC, TEND, DLST, SUSP, EXWA, WTRC, WRTH, ACCU, MTNG, INFO, TNDP
type CorporateActionEventType28Code string

type CorporateActionEventType73Choice struct {
	Cd    CorporateActionEventType28Code `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 Cd"`
	Prtry GenericIdentification30        `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 Prtry"`
}

// Must match the pattern [A-Z]{2,2}
type CountryCode string

// May be one of CRDT, DBIT
type CreditDebitCode string

type DateAndDateTime2Choice struct {
	Dt   ISODate     `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 Dt"`
	DtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 DtTm"`
}

type DateTimePeriodDetails2 struct {
	FrDtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 FrDtTm"`
	ToDtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 ToDtTm"`
}

type Document struct {
	IntraPosMvmntPstngRpt IntraPositionMovementPostingReportV07 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 IntraPosMvmntPstngRpt"`
}

// May be one of YEAR, MNTH, QUTR, SEMI, WEEK
type EventFrequency3Code string

// May be one of YEAR, ADHO, MNTH, DAIL, INDA, WEEK
type EventFrequency4Code string

// Must match the pattern [0-9]{3}
type Exact3NumericText string

// Must match the pattern [a-zA-Z0-9]{4}
type Exact4AlphaNumericText string

// Must match the pattern [0-9]{5}
type Exact5NumericText string

// Must be at least 1 items long
type ExternalFinancialInstrumentIdentificationType1Code string

type FinancialInstrumentAttributes92 struct {
	PlcOfListg             MarketIdentification3Choice            `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 PlcOfListg,omitempty"`
	DayCntBsis             InterestComputationMethodFormat4Choice `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 DayCntBsis,omitempty"`
	RegnForm               FormOfSecurity6Choice                  `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 RegnForm,omitempty"`
	PmtFrqcy               Frequency23Choice                      `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 PmtFrqcy,omitempty"`
	PmtSts                 SecuritiesPaymentStatus5Choice         `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 PmtSts,omitempty"`
	VarblRateChngFrqcy     Frequency23Choice                      `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 VarblRateChngFrqcy,omitempty"`
	ClssfctnTp             ClassificationType32Choice             `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 ClssfctnTp,omitempty"`
	OptnStyle              OptionStyle8Choice                     `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 OptnStyle,omitempty"`
	OptnTp                 OptionType6Choice                      `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 OptnTp,omitempty"`
	DnmtnCcy               ActiveOrHistoricCurrencyCode           `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 DnmtnCcy,omitempty"`
	CpnDt                  ISODate                                `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 CpnDt,omitempty"`
	XpryDt                 ISODate                                `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 XpryDt,omitempty"`
	FltgRateFxgDt          ISODate                                `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 FltgRateFxgDt,omitempty"`
	MtrtyDt                ISODate                                `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 MtrtyDt,omitempty"`
	IsseDt                 ISODate                                `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 IsseDt,omitempty"`
	NxtCllblDt             ISODate                                `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 NxtCllblDt,omitempty"`
	PutblDt                ISODate                                `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 PutblDt,omitempty"`
	DtdDt                  ISODate                                `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 DtdDt,omitempty"`
	FrstPmtDt              ISODate                                `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 FrstPmtDt,omitempty"`
	PrvsFctr               float64                                `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 PrvsFctr,omitempty"`
	CurFctr                float64                                `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 CurFctr,omitempty"`
	NxtFctr                float64                                `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 NxtFctr,omitempty"`
	IntrstRate             float64                                `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 IntrstRate,omitempty"`
	YldToMtrtyRate         float64                                `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 YldToMtrtyRate,omitempty"`
	NxtIntrstRate          float64                                `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 NxtIntrstRate,omitempty"`
	IndxRateBsis           float64                                `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 IndxRateBsis,omitempty"`
	CpnAttchdNb            Number22Choice                         `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 CpnAttchdNb,omitempty"`
	PoolNb                 GenericIdentification37                `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 PoolNb,omitempty"`
	QtyBrkdwn              []QuantityBreakdown31                  `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 QtyBrkdwn,omitempty"`
	VarblRateInd           bool                                   `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 VarblRateInd,omitempty"`
	CllblInd               bool                                   `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 CllblInd,omitempty"`
	PutblInd               bool                                   `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 PutblInd,omitempty"`
	MktOrIndctvPric        PriceType4Choice                       `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 MktOrIndctvPric,omitempty"`
	ExrcPric               Price7                                 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 ExrcPric,omitempty"`
	SbcptPric              Price7                                 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 SbcptPric,omitempty"`
	ConvsPric              Price7                                 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 ConvsPric,omitempty"`
	StrkPric               Price7                                 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 StrkPric,omitempty"`
	MinNmnlQty             FinancialInstrumentQuantity1Choice     `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 MinNmnlQty,omitempty"`
	CtrctSz                FinancialInstrumentQuantity1Choice     `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 CtrctSz,omitempty"`
	UndrlygFinInstrmId     []SecurityIdentification19             `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 UndrlygFinInstrmId,omitempty"`
	FinInstrmAttrAddtlDtls Max350Text                             `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 FinInstrmAttrAddtlDtls,omitempty"`
}

type FinancialInstrumentDetails29 struct {
	FinInstrmId      SecurityIdentification19        `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 FinInstrmId"`
	FinInstrmAttrbts FinancialInstrumentAttributes92 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 FinInstrmAttrbts,omitempty"`
	SubBal           []IntraPositionDetails45        `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 SubBal"`
}

type FinancialInstrumentQuantity1Choice struct {
	Unit     float64 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 Unit"`
	FaceAmt  float64 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 FaceAmt"`
	AmtsdVal float64 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 AmtsdVal"`
}

type ForeignExchangeTerms23 struct {
	UnitCcy  ActiveCurrencyCode      `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 UnitCcy"`
	QtdCcy   ActiveCurrencyCode      `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 QtdCcy"`
	XchgRate float64                 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 XchgRate"`
	RsltgAmt ActiveCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 RsltgAmt"`
}

// May be one of BEAR, REGD
type FormOfSecurity1Code string

type FormOfSecurity6Choice struct {
	Cd    FormOfSecurity1Code     `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 Cd"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 Prtry"`
}

type Frequency23Choice struct {
	Cd    EventFrequency3Code     `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 Cd"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 Prtry"`
}

type Frequency25Choice struct {
	Cd    EventFrequency4Code     `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 Cd"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 Prtry"`
}

type GenericIdentification1 struct {
	Id      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 Id"`
	SchmeNm Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 SchmeNm,omitempty"`
	Issr    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 Issr,omitempty"`
}

type GenericIdentification30 struct {
	Id      Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 Id"`
	Issr    Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 Issr"`
	SchmeNm Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 SchmeNm,omitempty"`
}

type GenericIdentification36 struct {
	Id      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 Id"`
	Issr    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 Issr"`
	SchmeNm Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 SchmeNm,omitempty"`
}

type GenericIdentification37 struct {
	Id   Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 Id"`
	Issr Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 Issr,omitempty"`
}

type GenericIdentification78 struct {
	Tp GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 Tp"`
	Id Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 Id,omitempty"`
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

type IdentificationSource3Choice struct {
	Cd    ExternalFinancialInstrumentIdentificationType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 Cd"`
	Prtry Max35Text                                          `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 Prtry"`
}

// May be one of A001, A002, A003, A004, A005, A006, A007, A008, A009, A010, A011, A012, A013, A014, NARR
type InterestComputationMethod2Code string

type InterestComputationMethodFormat4Choice struct {
	Cd    InterestComputationMethod2Code `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 Cd"`
	Prtry GenericIdentification30        `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 Prtry"`
}

type IntraPositionDetails45 struct {
	SfkpgPlc      SafekeepingPlaceFormat10Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 SfkpgPlc,omitempty"`
	BalFr         SecuritiesBalanceType6Choice     `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 BalFr"`
	IntraPosMvmnt []IntraPositionMovementDetails15 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 IntraPosMvmnt"`
}

type IntraPositionMovementDetails15 struct {
	Id                 References42Choice                 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 Id,omitempty"`
	SttldQty           FinancialInstrumentQuantity1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 SttldQty"`
	PrevslySttldQty    FinancialInstrumentQuantity1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 PrevslySttldQty,omitempty"`
	RmngToBeSttldQty   FinancialInstrumentQuantity1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 RmngToBeSttldQty,omitempty"`
	SctiesSubBalId     GenericIdentification37            `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 SctiesSubBalId,omitempty"`
	BalTo              SecuritiesBalanceType6Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 BalTo"`
	SttlmDt            DateAndDateTime2Choice             `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 SttlmDt"`
	AvlblDt            DateAndDateTime2Choice             `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 AvlblDt,omitempty"`
	AckdStsTmStmp      ISODateTime                        `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 AckdStsTmStmp,omitempty"`
	CorpActnEvtTp      CorporateActionEventType73Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 CorpActnEvtTp,omitempty"`
	CollMntrAmt        AmountAndDirection44               `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 CollMntrAmt,omitempty"`
	InstrPrcgAddtlDtls Max350Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 InstrPrcgAddtlDtls,omitempty"`
	SplmtryData        []SupplementaryData1               `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 SplmtryData,omitempty"`
}

type IntraPositionMovementPostingReportV07 struct {
	Pgntn       Pagination1                    `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 Pgntn"`
	StmtGnlDtls Statement60                    `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 StmtGnlDtls"`
	AcctOwnr    PartyIdentification92Choice    `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 AcctOwnr,omitempty"`
	SfkpgAcct   SecuritiesAccount19            `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 SfkpgAcct"`
	FinInstrm   []FinancialInstrumentDetails29 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 FinInstrm,omitempty"`
}

// Must match the pattern [A-Z0-9]{4,4}
type MICIdentifier string

type MarketIdentification3Choice struct {
	MktIdrCd MICIdentifier `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 MktIdrCd"`
	Desc     Max35Text     `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 Desc"`
}

// Must be at least 1 items long
type Max140Text string

// Must be at least 1 items long
type Max16Text string

// Must be at least 1 items long
type Max350Text string

// Must be at least 1 items long
type Max35Text string

// Must match the pattern [0-9]{1,5}
type Max5NumericText string

// Must be at least 1 items long
type Max70Text string

type Number22Choice struct {
	Shrt Exact3NumericText      `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 Shrt"`
	Lng  GenericIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 Lng"`
}

type Number3Choice struct {
	Shrt Exact3NumericText `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 Shrt"`
	Lng  Exact5NumericText `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 Lng"`
}

// May be one of AMER, EURO
type OptionStyle2Code string

type OptionStyle8Choice struct {
	Cd    OptionStyle2Code        `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 Cd"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 Prtry"`
}

// May be one of CALL, PUTO
type OptionType1Code string

type OptionType6Choice struct {
	Cd    OptionType1Code         `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 Cd"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 Prtry"`
}

type OtherIdentification1 struct {
	Id  Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 Id"`
	Sfx Max16Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 Sfx,omitempty"`
	Tp  IdentificationSource3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 Tp"`
}

type Pagination1 struct {
	PgNb      Max5NumericText `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 PgNb"`
	LastPgInd bool            `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 LastPgInd"`
}

type PartyIdentification92Choice struct {
	AnyBIC  AnyBICIdentifier        `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 AnyBIC"`
	PrtryId GenericIdentification36 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 PrtryId"`
}

type Period2 struct {
	FrDt ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 FrDt"`
	ToDt ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 ToDt"`
}

type Period5Choice struct {
	FrDtTmToDtTm DateTimePeriodDetails2 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 FrDtTmToDtTm"`
	FrDtToDt     Period2                `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 FrDtToDt"`
}

type Price7 struct {
	Tp  YieldedOrValueType1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 Tp"`
	Val PriceRateOrAmount3Choice  `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 Val"`
}

type PriceRateOrAmount3Choice struct {
	Rate float64                                    `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 Rate"`
	Amt  ActiveOrHistoricCurrencyAnd13DecimalAmount `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 Amt"`
}

type PriceType4Choice struct {
	Mkt    Price7 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 Mkt"`
	Indctv Price7 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 Indctv"`
}

// May be one of DISC, PREM, PARV
type PriceValueType1Code string

type QuantityBreakdown31 struct {
	LotNb  GenericIdentification37            `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 LotNb"`
	LotQty FinancialInstrumentQuantity1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 LotQty,omitempty"`
}

type References42Choice struct {
	AcctOwnrTxId      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 AcctOwnrTxId"`
	AcctSvcrTxId      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 AcctSvcrTxId"`
	PoolId            Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 PoolId"`
	MktInfrstrctrTxId Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 MktInfrstrctrTxId"`
	PrcrTxId          Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 PrcrTxId"`
}

// May be one of CUST, ICSD, NCSD, SHHE
type SafekeepingPlace1Code string

// May be one of SHHE
type SafekeepingPlace3Code string

type SafekeepingPlaceFormat10Choice struct {
	Id      SafekeepingPlaceTypeAndText8             `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 Id"`
	Ctry    CountryCode                              `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 Ctry"`
	TpAndId SafekeepingPlaceTypeAndAnyBICIdentifier1 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 TpAndId"`
	Prtry   GenericIdentification78                  `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 Prtry"`
}

type SafekeepingPlaceTypeAndAnyBICIdentifier1 struct {
	SfkpgPlcTp SafekeepingPlace1Code `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 SfkpgPlcTp"`
	Id         AnyBICIdentifier      `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 Id"`
}

type SafekeepingPlaceTypeAndText8 struct {
	SfkpgPlcTp SafekeepingPlace3Code `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 SfkpgPlcTp"`
	Id         Max35Text             `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 Id,omitempty"`
}

type SecuritiesAccount19 struct {
	Id Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 Id"`
	Tp GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 Tp,omitempty"`
	Nm Max70Text               `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 Nm,omitempty"`
}

// May be one of BLOK, AWAS, AVAI, BLCA, BLOT, BLOV, BORR, COLI, COLO, COLA, LOAN, MARG, PECA, PEDA, PLED, REGO, RSTR, OTHR, TRAN, DRAW, CLEN, DIRT, NOMI, SPOS, UNRG, ISSU, QUAS, LODE
type SecuritiesBalanceType11Code string

type SecuritiesBalanceType6Choice struct {
	Cd    SecuritiesBalanceType11Code `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 Cd"`
	Prtry GenericIdentification30     `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 Prtry"`
}

// May be one of FULL, NILL, PART
type SecuritiesPaymentStatus1Code string

type SecuritiesPaymentStatus5Choice struct {
	Cd    SecuritiesPaymentStatus1Code `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 Cd"`
	Prtry GenericIdentification30      `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 Prtry"`
}

type SecurityIdentification19 struct {
	ISIN   ISINOct2015Identifier  `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 ISIN,omitempty"`
	OthrId []OtherIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 OthrId,omitempty"`
	Desc   Max140Text             `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 Desc,omitempty"`
}

type Statement60 struct {
	RptNb     Number3Choice      `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 RptNb,omitempty"`
	QryRef    Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 QryRef,omitempty"`
	StmtId    Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 StmtId,omitempty"`
	StmtPrd   Period5Choice      `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 StmtPrd"`
	Frqcy     Frequency25Choice  `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 Frqcy,omitempty"`
	UpdTp     UpdateType15Choice `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 UpdTp,omitempty"`
	ActvtyInd bool               `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 ActvtyInd"`
}

// May be one of COMP, DELT
type StatementUpdateType1Code string

type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
}

type UpdateType15Choice struct {
	Cd    StatementUpdateType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 Cd"`
	Prtry GenericIdentification30  `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 Prtry"`
}

type YieldedOrValueType1Choice struct {
	Yldd  bool                `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 Yldd"`
	ValTp PriceValueType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.07 ValTp"`
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
