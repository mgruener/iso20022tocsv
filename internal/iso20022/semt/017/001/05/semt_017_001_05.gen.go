// Code generated by main. DO NOT EDIT.

package semt_017_001_05

import (
	"bytes"
	"encoding/xml"
	"time"
)

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

// May be one of ADDR, PBOX, HOME, BIZZ, MLTO, DLVY
type AddressType2Code string

type AmountAndDirection3 struct {
	Amt    ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 Amt"`
	CdtDbt CreditDebitCode                   `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 CdtDbt"`
}

type AmountAndDirection4 struct {
	Amt       ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 Amt"`
	CdtDbtInd CreditDebitCode                   `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 CdtDbtInd,omitempty"`
}

// Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type AnyBICIdentifier string

type BalanceQuantity5Choice struct {
	Qty   FinancialInstrumentQuantity1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 Qty"`
	Prtry GenericIdentification22            `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 Prtry"`
}

type BeneficialOwnership1Choice struct {
	Ind   bool                    `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 Ind"`
	Prtry GenericIdentification20 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 Prtry"`
}

type BlockTrade1Choice struct {
	Cd    BlockTrade1Code         `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 Cd"`
	Prtry GenericIdentification20 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 Prtry"`
}

// May be one of BLPA, BLCH
type BlockTrade1Code string

type CashSettlementSystem1Choice struct {
	Cd    CashSettlementSystem2Code `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 Cd"`
	Prtry GenericIdentification20   `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 Prtry"`
}

// May be one of GROS, NETS
type CashSettlementSystem2Code string

type CentralCounterPartyEligibility1Choice struct {
	Ind   bool                    `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 Ind"`
	Prtry GenericIdentification20 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 Prtry"`
}

type ClosingBalance1 struct {
	ShrtLngInd ShortLong1Code        `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 ShrtLngInd"`
	ClsgBal    ClosingBalance1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 ClsgBal"`
}

type ClosingBalance1Choice struct {
	Fnl    BalanceQuantity5Choice `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 Fnl"`
	Intrmy BalanceQuantity5Choice `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 Intrmy"`
}

// May be one of ACCU, ACTV, ATTI, BRUP, DFLT, BONU, EXRI, CAPD, CAPG, DRCA, DVCA, CAPI, CHAN, CLSA, COOP, CONS, CONV, CREV, DECR, DETI, DSCL, DVOP, DRIP, DRAW, DTCH, EXOF, REDM, MCAL, INCR, INFO, PPMT, INTR, PRII, RHDI, LIQU, EXTM, MRGR, NOOF, CERT, ODLT, OTHR, PARI, PCAL, PRED, PINK, PLAC, PDEF, PRIO, BPUT, REDO, REMK, BIDS, SPLR, RHTS, DVSC, MTNG, SHPR, SMAL, SOFF, SPLF, DVSE, TREC, TEND, DLST, SUSP, EXWA, WTRC, WRTH
type CorporateActionEventType14Code string

type CorporateActionEventType15Choice struct {
	Cd    CorporateActionEventType14Code `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 Cd"`
	Prtry GenericIdentification20        `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 Prtry"`
}

// Must match the pattern [A-Z]{2,2}
type CountryCode string

// May be one of CRDT, DBIT
type CreditDebitCode string

type DateAndDateTimeChoice struct {
	Dt   ISODate     `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 Dt"`
	DtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 DtTm"`
}

type DateTimePeriodDetails struct {
	FrDtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 FrDtTm"`
	ToDtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 ToDtTm"`
}

// May be one of VARI
type DateType3Code string

// May be one of FREE, APMT
type DeliveryReceiptType2Code string

type Document struct {
	SctiesTxPstngRpt SecuritiesTransactionPostingReportV05 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 SctiesTxPstngRpt"`
}

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

type FinancialInstrumentDetails17 struct {
	FinInstrmId SecurityIdentification14      `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 FinInstrmId"`
	PricDtls    PriceInformation6             `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 PricDtls,omitempty"`
	SfkpgPlc    SafekeepingPlaceFormat3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 SfkpgPlc,omitempty"`
	OpngBal     OpeningBalance1               `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 OpngBal,omitempty"`
	ClsgBal     ClosingBalance1               `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 ClsgBal,omitempty"`
	Tx          []Transaction36               `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 Tx"`
}

type FinancialInstrumentQuantity1Choice struct {
	Unit     float64 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 Unit"`
	FaceAmt  float64 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 FaceAmt"`
	AmtsdVal float64 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 AmtsdVal"`
}

type Frequency4Choice struct {
	Cd    EventFrequency4Code     `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 Cd"`
	Prtry GenericIdentification20 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 Prtry"`
}

type GenericIdentification19 struct {
	Id      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 Id"`
	Issr    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 Issr"`
	SchmeNm Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 SchmeNm,omitempty"`
}

type GenericIdentification20 struct {
	Id      Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 Id"`
	Issr    Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 Issr"`
	SchmeNm Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 SchmeNm,omitempty"`
}

type GenericIdentification21 struct {
	Tp GenericIdentification20 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 Tp"`
	Id Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 Id,omitempty"`
}

type GenericIdentification22 struct {
	Id      Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 Id"`
	Issr    Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 Issr"`
	SchmeNm Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 SchmeNm,omitempty"`
	Bal     float64                `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 Bal"`
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

type IdentificationSource3Choice struct {
	Cd    ExternalFinancialInstrumentIdentificationType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 Cd"`
	Prtry Max35Text                                          `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 Prtry"`
}

type LetterOfGuarantee1Choice struct {
	Ind   bool                    `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 Ind"`
	Prtry GenericIdentification20 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 Prtry"`
}

// Must match the pattern [A-Z0-9]{4,4}
type MICIdentifier string

type MarketClientSide1Choice struct {
	Cd    MarketClientSideCode    `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 Cd"`
	Prtry GenericIdentification20 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 Prtry"`
}

// May be one of MAKT, CLNT
type MarketClientSideCode string

type MarketIdentification1Choice struct {
	MktIdrCd MICIdentifier `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 MktIdrCd"`
	Desc     Max35Text     `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 Desc"`
}

type MarketIdentification6 struct {
	Id MarketIdentification1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 Id,omitempty"`
	Tp MarketType4Choice           `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 Tp"`
}

type MarketIdentification78 struct {
	Id MarketIdentification1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 Id,omitempty"`
	Tp MarketType3Choice           `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 Tp"`
}

// May be one of PRIM, SECM, OTCO, VARI, EXCH
type MarketType2Code string

type MarketType3Choice struct {
	Cd    MarketType2Code         `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 Cd"`
	Prtry GenericIdentification20 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 Prtry"`
}

type MarketType4Choice struct {
	Cd    MarketType4Code         `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 Cd"`
	Prtry GenericIdentification20 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 Prtry"`
}

// May be one of FUND, LMAR, THEO, VEND
type MarketType4Code string

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

type NameAndAddress5 struct {
	Nm  Max350Text     `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 Nm"`
	Adr PostalAddress1 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 Adr,omitempty"`
}

type NettingEligibility1Choice struct {
	Ind   bool                    `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 Ind"`
	Prtry GenericIdentification20 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 Prtry"`
}

type Number3Choice struct {
	Shrt Exact3NumericText `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 Shrt"`
	Lng  Exact5NumericText `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 Lng"`
}

type OpeningBalance1 struct {
	ShrtLngInd ShortLong1Code        `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 ShrtLngInd"`
	OpngBal    OpeningBalance1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 OpngBal"`
}

type OpeningBalance1Choice struct {
	Frst   BalanceQuantity5Choice `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 Frst"`
	Intrmy BalanceQuantity5Choice `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 Intrmy"`
}

type OriginalAndCurrentQuantities1 struct {
	FaceAmt  float64 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 FaceAmt"`
	AmtsdVal float64 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 AmtsdVal"`
}

type OtherIdentification1 struct {
	Id  Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 Id"`
	Sfx Max16Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 Sfx,omitempty"`
	Tp  IdentificationSource3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 Tp"`
}

// May be one of A144, NRST, RSTR
type OwnershipLegalRestrictions1Code string

type Pagination struct {
	PgNb      Max5NumericText `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 PgNb"`
	LastPgInd bool            `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 LastPgInd"`
}

type PartyIdentification36Choice struct {
	AnyBIC  AnyBICIdentifier        `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 AnyBIC"`
	PrtryId GenericIdentification19 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 PrtryId"`
}

type PartyIdentification44Choice struct {
	AnyBIC   AnyBICIdentifier `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 AnyBIC"`
	NmAndAdr NameAndAddress5  `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 NmAndAdr"`
	Ctry     CountryCode      `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 Ctry"`
}

type PartyIdentification45Choice struct {
	AnyBIC   AnyBICIdentifier        `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 AnyBIC"`
	PrtryId  GenericIdentification19 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 PrtryId"`
	NmAndAdr NameAndAddress5         `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 NmAndAdr"`
}

type PartyIdentification46 struct {
	Id     PartyIdentification44Choice `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 Id"`
	PrcgId Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 PrcgId,omitempty"`
}

type PartyIdentificationAndAccount44 struct {
	Id        PartyIdentification45Choice `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 Id"`
	SfkpgAcct SecuritiesAccount13         `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 SfkpgAcct,omitempty"`
	PrcgId    Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 PrcgId,omitempty"`
}

type Period2 struct {
	FrDt ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 FrDt"`
	ToDt ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 ToDt"`
}

type Period2Choice struct {
	FrDtTmToDtTm DateTimePeriodDetails `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 FrDtTmToDtTm"`
	FrDtToDt     Period2               `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 FrDtToDt"`
}

type PostalAddress1 struct {
	AdrTp       AddressType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 AdrTp,omitempty"`
	AdrLine     []Max70Text      `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 AdrLine,omitempty"`
	StrtNm      Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 StrtNm,omitempty"`
	BldgNb      Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 BldgNb,omitempty"`
	PstCd       Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 PstCd,omitempty"`
	TwnNm       Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 TwnNm,omitempty"`
	CtrySubDvsn Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 CtrySubDvsn,omitempty"`
	Ctry        CountryCode      `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 Ctry"`
}

type PriceInformation6 struct {
	Tp        TypeOfPrice6Choice               `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 Tp"`
	ValTp     YieldedOrValueType1Choice        `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 ValTp"`
	Val       PriceRateOrAmountOrUnknownChoice `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 Val"`
	SrcOfPric MarketIdentification6            `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 SrcOfPric,omitempty"`
	QtnDt     DateAndDateTimeChoice            `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 QtnDt,omitempty"`
}

type PriceRateOrAmountOrUnknownChoice struct {
	Rate     float64                                    `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 Rate"`
	Amt      ActiveOrHistoricCurrencyAnd13DecimalAmount `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 Amt"`
	UknwnInd bool                                       `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 UknwnInd"`
}

// May be one of DISC, PREM, PARV
type PriceValueType1Code string

type PurposeCode2Choice struct {
	Cd    SecuritiesAccountPurposeType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 Cd"`
	Prtry GenericIdentification20           `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 Prtry"`
}

type Quantity6Choice struct {
	Qty             FinancialInstrumentQuantity1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 Qty"`
	OrgnlAndCurFace OriginalAndCurrentQuantities1      `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 OrgnlAndCurFace"`
}

// May be one of DELI, RECE
type ReceiveDelivery1Code string

type Registration1Choice struct {
	Cd    Registration1Code       `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 Cd"`
	Prtry GenericIdentification20 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 Prtry"`
}

// May be one of NREG, YREG
type Registration1Code string

type RepurchaseType3Choice struct {
	Cd    RepurchaseType3Code     `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 Cd"`
	Prtry GenericIdentification20 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 Prtry"`
}

// May be one of CADJ, CALL, PAIR, RATE, ROLP, TOPU, WTHD, PADJ
type RepurchaseType3Code string

type Restriction1Choice struct {
	Cd    OwnershipLegalRestrictions1Code `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 Cd"`
	Prtry GenericIdentification20         `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 Prtry"`
}

// May be one of CUST, ICSD, NCSD, SHHE
type SafekeepingPlace1Code string

// May be one of SHHE
type SafekeepingPlace3Code string

type SafekeepingPlaceFormat3Choice struct {
	Id      SafekeepingPlaceTypeAndText3             `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 Id"`
	Ctry    CountryCode                              `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 Ctry"`
	TpAndId SafekeepingPlaceTypeAndAnyBICIdentifier1 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 TpAndId"`
	Prtry   GenericIdentification21                  `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 Prtry"`
}

type SafekeepingPlaceTypeAndAnyBICIdentifier1 struct {
	SfkpgPlcTp SafekeepingPlace1Code `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 SfkpgPlcTp"`
	Id         AnyBICIdentifier      `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 Id"`
}

type SafekeepingPlaceTypeAndText3 struct {
	SfkpgPlcTp SafekeepingPlace3Code `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 SfkpgPlcTp"`
	Id         Max35Text             `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 Id,omitempty"`
}

type SecuritiesAccount13 struct {
	Id Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 Id"`
	Tp GenericIdentification20 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 Tp,omitempty"`
	Nm Max70Text               `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 Nm,omitempty"`
}

type SecuritiesAccount14 struct {
	Id Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 Id"`
	Tp PurposeCode2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 Tp,omitempty"`
	Nm Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 Nm,omitempty"`
}

// May be one of MARG, SHOR, ABRD, CEND, DVPA, PHYS
type SecuritiesAccountPurposeType1Code string

type SecuritiesRTGS1Choice struct {
	Ind   bool                    `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 Ind"`
	Prtry GenericIdentification20 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 Prtry"`
}

type SecuritiesTransactionPostingReportV05 struct {
	Pgntn         Pagination                     `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 Pgntn"`
	StmtGnlDtls   Statement11                    `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 StmtGnlDtls"`
	AcctOwnr      PartyIdentification36Choice    `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 AcctOwnr,omitempty"`
	SfkpgAcct     SecuritiesAccount13            `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 SfkpgAcct"`
	FinInstrmDtls []FinancialInstrumentDetails17 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 FinInstrmDtls,omitempty"`
	SubAcctDtls   []SubAccountIdentification34   `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 SubAcctDtls,omitempty"`
}

type SecuritiesTransactionType10Choice struct {
	Cd    SecuritiesTransactionType8Code `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 Cd"`
	Prtry GenericIdentification20        `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 Prtry"`
}

// May be one of BSBK, BIYI, CNCB, COLI, COLO, CONV, FCTA, INSP, ISSU, MKDW, MKUP, NETT, NSYN, OWNE, OWNI, PAIR, PLAC, PORT, REAL, REDI, REDM, RELE, REPU, RODE, RVPO, SBBK, SBRE, SECB, SECL, SLRE, SUBS, SYND, TBAC, TRAD, TRPO, TRVO, TURN, CLAI, CORP, AUTO, ETFT
type SecuritiesTransactionType8Code string

type SecurityIdentification14 struct {
	ISIN   ISINIdentifier         `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 ISIN,omitempty"`
	OthrId []OtherIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 OthrId,omitempty"`
	Desc   Max140Text             `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 Desc,omitempty"`
}

type SettlementDate1Choice struct {
	Dt   DateAndDateTimeChoice     `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 Dt"`
	DtCd SettlementDateCode1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 DtCd"`
}

// May be one of WISS
type SettlementDate4Code string

type SettlementDateCode1Choice struct {
	Cd    SettlementDate4Code     `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 Cd"`
	Prtry GenericIdentification20 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 Prtry"`
}

type SettlementDetails75 struct {
	SttlmTxCond    []SettlementTransactionCondition12Choice `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 SttlmTxCond,omitempty"`
	SttlgCpcty     SettlingCapacity4Choice                  `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 SttlgCpcty,omitempty"`
	StmpDtyTaxBsis GenericIdentification20                  `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 StmpDtyTaxBsis,omitempty"`
	SctiesRTGS     SecuritiesRTGS1Choice                    `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 SctiesRTGS,omitempty"`
	Regn           Registration1Choice                      `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 Regn,omitempty"`
	BnfclOwnrsh    BeneficialOwnership1Choice               `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 BnfclOwnrsh,omitempty"`
	CshClrSys      CashSettlementSystem1Choice              `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 CshClrSys,omitempty"`
	TaxCpcty       TaxCapacityParty1Choice                  `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 TaxCpcty,omitempty"`
	RpTp           RepurchaseType3Choice                    `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 RpTp,omitempty"`
	MktClntSd      MarketClientSide1Choice                  `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 MktClntSd,omitempty"`
	BlckTrad       BlockTrade1Choice                        `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 BlckTrad,omitempty"`
	LglRstrctns    Restriction1Choice                       `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 LglRstrctns,omitempty"`
	SttlmSysMtd    SettlementSystemMethod1Choice            `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 SttlmSysMtd,omitempty"`
	NetgElgblty    NettingEligibility1Choice                `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 NetgElgblty,omitempty"`
	CCPElgblty     CentralCounterPartyEligibility1Choice    `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 CCPElgblty,omitempty"`
	LttrOfGrnt     LetterOfGuarantee1Choice                 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 LttrOfGrnt,omitempty"`
	PrtlSttlmInd   SettlementTransactionCondition5Code      `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 PrtlSttlmInd,omitempty"`
}

type SettlementOrCorporateActionEvent9Choice struct {
	SctiesTxTp    SecuritiesTransactionType10Choice `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 SctiesTxTp"`
	CorpActnEvtTp CorporateActionEventType15Choice  `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 CorpActnEvtTp"`
}

type SettlementParties13 struct {
	Dpstry PartyIdentification46           `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 Dpstry,omitempty"`
	Pty1   PartyIdentificationAndAccount44 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 Pty1,omitempty"`
	Pty2   PartyIdentificationAndAccount44 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 Pty2,omitempty"`
	Pty3   PartyIdentificationAndAccount44 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 Pty3,omitempty"`
	Pty4   PartyIdentificationAndAccount44 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 Pty4,omitempty"`
	Pty5   PartyIdentificationAndAccount44 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 Pty5,omitempty"`
}

type SettlementSystemMethod1Choice struct {
	Cd    SettlementSystemMethod1Code `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 Cd"`
	Prtry GenericIdentification20     `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 Prtry"`
}

// May be one of NSET, YSET
type SettlementSystemMethod1Code string

type SettlementTransactionCondition12Choice struct {
	Cd    SettlementTransactionCondition8Code `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 Cd"`
	Prtry GenericIdentification20             `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 Prtry"`
}

// May be one of PART, NPAR, PARC, PARQ
type SettlementTransactionCondition5Code string

// May be one of ASGN, BUTC, CLEN, DIRT, DLWM, DRAW, EXER, FRCL, KNOC, PHYS, RESI, SHOR, SPDL, SPST, EXPI, PENS, UNEX, TRIP, NOMC, TRAN, RHYP, ADEA, RPTO
type SettlementTransactionCondition8Code string

// May be one of SAGE, CUST, SPRI, RISP
type SettlingCapacity2Code string

type SettlingCapacity4Choice struct {
	Cd    SettlingCapacity2Code   `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 Cd"`
	Prtry GenericIdentification20 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 Prtry"`
}

// May be one of SHOR, LONG
type ShortLong1Code string

type Statement11 struct {
	RptNb      Number3Choice         `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 RptNb,omitempty"`
	QryRef     Max35Text             `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 QryRef,omitempty"`
	StmtId     Max35Text             `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 StmtId,omitempty"`
	StmtPrd    Period2Choice         `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 StmtPrd"`
	Frqcy      Frequency4Choice      `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 Frqcy,omitempty"`
	UpdTp      UpdateType2Choice     `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 UpdTp,omitempty"`
	StmtBsis   StatementBasis2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 StmtBsis"`
	ActvtyInd  bool                  `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 ActvtyInd"`
	SubAcctInd bool                  `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 SubAcctInd"`
}

type StatementBasis2Choice struct {
	Cd    StatementBasis2Code     `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 Cd"`
	Prtry GenericIdentification20 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 Prtry"`
}

// May be one of SETT, TRAD
type StatementBasis2Code string

// May be one of COMP, DELT
type StatementUpdateType1Code string

type SubAccountIdentification34 struct {
	AcctOwnr      PartyIdentification36Choice    `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 AcctOwnr,omitempty"`
	SfkpgAcct     SecuritiesAccount14            `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 SfkpgAcct"`
	ActvtyInd     bool                           `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 ActvtyInd"`
	FinInstrmDtls []FinancialInstrumentDetails17 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 FinInstrmDtls,omitempty"`
}

type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
}

type TaxCapacityParty1Choice struct {
	Cd    TaxLiability1Code       `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 Cd"`
	Prtry GenericIdentification20 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 Prtry"`
}

// May be one of PRIN, AGEN
type TaxLiability1Code string

type TradeDate1Choice struct {
	Dt   DateAndDateTimeChoice `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 Dt"`
	DtCd TradeDateCode1Choice  `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 DtCd"`
}

type TradeDateCode1Choice struct {
	Cd    DateType3Code           `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 Cd"`
	Prtry GenericIdentification20 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 Prtry"`
}

type Transaction36 struct {
	AcctOwnrTxId      Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 AcctOwnrTxId"`
	AcctSvcrTxId      Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 AcctSvcrTxId,omitempty"`
	MktInfrstrctrTxId Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 MktInfrstrctrTxId,omitempty"`
	PrcrTxId          Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 PrcrTxId,omitempty"`
	TradId            []Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 TradId,omitempty"`
	PoolId            Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 PoolId,omitempty"`
	CmonId            Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 CmonId,omitempty"`
	CorpActnEvtId     Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 CorpActnEvtId,omitempty"`
	TrptyAgtCollTxId  Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 TrptyAgtCollTxId,omitempty"`
	ClntTrptyCollTxId Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 ClntTrptyCollTxId,omitempty"`
	ClntCollInstrId   Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 ClntCollInstrId,omitempty"`
	TrptyCollInstrId  Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 TrptyCollInstrId,omitempty"`
	TxDtls            TransactionDetails63 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 TxDtls,omitempty"`
	SplmtryData       []SupplementaryData1 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 SplmtryData,omitempty"`
}

type TransactionActivity1Choice struct {
	Cd    TransactionActivity1Code `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 Cd"`
	Prtry GenericIdentification20  `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 Prtry"`
}

// May be one of BOLE, CLAI, COLL, CORP, SETT
type TransactionActivity1Code string

type TransactionDetails63 struct {
	TxActvty               TransactionActivity1Choice              `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 TxActvty"`
	SttlmTxOrCorpActnEvtTp SettlementOrCorporateActionEvent9Choice `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 SttlmTxOrCorpActnEvtTp,omitempty"`
	SctiesMvmntTp          ReceiveDelivery1Code                    `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 SctiesMvmntTp"`
	Pmt                    DeliveryReceiptType2Code                `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 Pmt"`
	SttlmParams            SettlementDetails75                     `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 SttlmParams,omitempty"`
	PlcOfTrad              MarketIdentification78                  `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 PlcOfTrad,omitempty"`
	SfkpgPlc               SafekeepingPlaceFormat3Choice           `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 SfkpgPlc,omitempty"`
	PlcOfClr               AnyBICIdentifier                        `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 PlcOfClr,omitempty"`
	PstngQty               Quantity6Choice                         `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 PstngQty"`
	NbOfDaysAcrd           float64                                 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 NbOfDaysAcrd,omitempty"`
	PstngAmt               AmountAndDirection3                     `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 PstngAmt,omitempty"`
	AcrdIntrstAmt          AmountAndDirection4                     `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 AcrdIntrstAmt,omitempty"`
	TradDt                 TradeDate1Choice                        `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 TradDt,omitempty"`
	FctvSttlmDt            DateAndDateTimeChoice                   `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 FctvSttlmDt"`
	SttlmDt                SettlementDate1Choice                   `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 SttlmDt,omitempty"`
	ValDt                  DateAndDateTimeChoice                   `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 ValDt,omitempty"`
	DlvrgSttlmPties        SettlementParties13                     `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 DlvrgSttlmPties,omitempty"`
	RcvgSttlmPties         SettlementParties13                     `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 RcvgSttlmPties,omitempty"`
	RvslInd                bool                                    `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 RvslInd,omitempty"`
	TxAddtlDtls            Max350Text                              `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 TxAddtlDtls,omitempty"`
}

// May be one of MRKT, INDC
type TypeOfPrice16Code string

type TypeOfPrice6Choice struct {
	Cd    TypeOfPrice16Code       `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 Cd"`
	Prtry GenericIdentification20 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 Prtry"`
}

type UpdateType2Choice struct {
	Cd    StatementUpdateType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 Cd"`
	Prtry GenericIdentification20  `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 Prtry"`
}

type YieldedOrValueType1Choice struct {
	Yldd  bool                `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 Yldd"`
	ValTp PriceValueType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.05 ValTp"`
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