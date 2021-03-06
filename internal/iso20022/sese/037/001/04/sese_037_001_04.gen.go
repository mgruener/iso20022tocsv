// Code generated by main. DO NOT EDIT.

package sese_037_001_04

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

// May be one of ADDR, PBOX, HOME, BIZZ, MLTO, DLVY
type AddressType2Code string

type AlternatePartyIdentification7 struct {
	IdTp    IdentificationType42Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 IdTp"`
	Ctry    CountryCode                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 Ctry"`
	AltrnId Max35Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 AltrnId"`
}

type AmountAndDirection44 struct {
	Amt                 ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 Amt"`
	CdtDbtInd           CreditDebitCode                   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 CdtDbtInd,omitempty"`
	OrgnlCcyAndOrdrdAmt ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 OrgnlCcyAndOrdrdAmt,omitempty"`
	FXDtls              ForeignExchangeTerms23            `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 FXDtls,omitempty"`
}

type AmountAndDirection52 struct {
	Amt       ActiveCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 Amt"`
	CdtDbtInd CreditDebitCode         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 CdtDbtInd"`
}

// Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type AnyBICIdentifier string

// Must match the pattern [A-Z]{6,6}
type CFIOct2015Identifier string

type ClassificationType32Choice struct {
	ClssfctnFinInstrm CFIOct2015Identifier    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 ClssfctnFinInstrm"`
	AltrnClssfctn     GenericIdentification36 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 AltrnClssfctn"`
}

// Must match the pattern [A-Z]{2,2}
type CountryCode string

// May be one of CRDT, DBIT
type CreditDebitCode string

type DateAndDateTimeChoice struct {
	Dt   ISODate     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 Dt"`
	DtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 DtTm"`
}

// May be one of VARI
type DateType3Code string

// May be one of FREE, APMT
type DeliveryReceiptType2Code string

type Document struct {
	PrtflTrfNtfctn PortfolioTransferNotificationV04 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 PrtflTrfNtfctn"`
}

// May be one of YEAR, MNTH, QUTR, SEMI, WEEK
type EventFrequency3Code string

// Must match the pattern [0-9]{3}
type Exact3NumericText string

// Must match the pattern [a-zA-Z0-9]{4}
type Exact4AlphaNumericText string

// Must match the pattern [0-9]{5}
type Exact5NumericText string

// Must be at least 1 items long
type ExternalFinancialInstrumentIdentificationType1Code string

type FinancialInstrumentAttributes64 struct {
	PlcOfListg             MarketIdentification3Choice            `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 PlcOfListg,omitempty"`
	DayCntBsis             InterestComputationMethodFormat4Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 DayCntBsis,omitempty"`
	RegnForm               FormOfSecurity6Choice                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 RegnForm,omitempty"`
	PmtFrqcy               Frequency23Choice                      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 PmtFrqcy,omitempty"`
	PmtSts                 SecuritiesPaymentStatus5Choice         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 PmtSts,omitempty"`
	VarblRateChngFrqcy     Frequency23Choice                      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 VarblRateChngFrqcy,omitempty"`
	ClssfctnTp             ClassificationType32Choice             `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 ClssfctnTp,omitempty"`
	OptnStyle              OptionStyle8Choice                     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 OptnStyle,omitempty"`
	OptnTp                 OptionType6Choice                      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 OptnTp,omitempty"`
	DnmtnCcy               ActiveOrHistoricCurrencyCode           `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 DnmtnCcy,omitempty"`
	CpnDt                  ISODate                                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 CpnDt,omitempty"`
	XpryDt                 ISODate                                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 XpryDt,omitempty"`
	FltgRateFxgDt          ISODate                                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 FltgRateFxgDt,omitempty"`
	MtrtyDt                ISODate                                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 MtrtyDt,omitempty"`
	IsseDt                 ISODate                                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 IsseDt,omitempty"`
	NxtCllblDt             ISODate                                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 NxtCllblDt,omitempty"`
	PutblDt                ISODate                                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 PutblDt,omitempty"`
	DtdDt                  ISODate                                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 DtdDt,omitempty"`
	FrstPmtDt              ISODate                                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 FrstPmtDt,omitempty"`
	PrvsFctr               float64                                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 PrvsFctr,omitempty"`
	CurFctr                float64                                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 CurFctr,omitempty"`
	NxtFctr                float64                                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 NxtFctr,omitempty"`
	IntrstRate             float64                                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 IntrstRate,omitempty"`
	YldToMtrtyRate         float64                                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 YldToMtrtyRate,omitempty"`
	NxtIntrstRate          float64                                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 NxtIntrstRate,omitempty"`
	IndxRateBsis           float64                                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 IndxRateBsis,omitempty"`
	CpnAttchdNb            Number22Choice                         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 CpnAttchdNb,omitempty"`
	PoolNb                 GenericIdentification37                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 PoolNb,omitempty"`
	VarblRateInd           bool                                   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 VarblRateInd,omitempty"`
	CllblInd               bool                                   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 CllblInd,omitempty"`
	PutblInd               bool                                   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 PutblInd,omitempty"`
	MktOrIndctvPric        PriceType1Choice                       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 MktOrIndctvPric,omitempty"`
	ExrcPric               Price2                                 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 ExrcPric,omitempty"`
	SbcptPric              Price2                                 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 SbcptPric,omitempty"`
	ConvsPric              Price2                                 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 ConvsPric,omitempty"`
	StrkPric               Price2                                 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 StrkPric,omitempty"`
	MinNmnlQty             FinancialInstrumentQuantity1Choice     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 MinNmnlQty,omitempty"`
	CtrctSz                FinancialInstrumentQuantity1Choice     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 CtrctSz,omitempty"`
	UndrlygFinInstrmId     []SecurityIdentification19             `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 UndrlygFinInstrmId,omitempty"`
	FinInstrmAttrAddtlDtls Max350Text                             `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 FinInstrmAttrAddtlDtls,omitempty"`
}

type FinancialInstrumentQuantity1Choice struct {
	Unit     float64 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 Unit"`
	FaceAmt  float64 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 FaceAmt"`
	AmtsdVal float64 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 AmtsdVal"`
}

type ForeignExchangeTerms23 struct {
	UnitCcy  ActiveCurrencyCode      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 UnitCcy"`
	QtdCcy   ActiveCurrencyCode      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 QtdCcy"`
	XchgRate float64                 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 XchgRate"`
	RsltgAmt ActiveCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 RsltgAmt"`
}

// May be one of BEAR, REGD
type FormOfSecurity1Code string

type FormOfSecurity6Choice struct {
	Cd    FormOfSecurity1Code     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 Cd"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 Prtry"`
}

type Frequency23Choice struct {
	Cd    EventFrequency3Code     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 Cd"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 Prtry"`
}

type GenericIdentification1 struct {
	Id      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 Id"`
	SchmeNm Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 SchmeNm,omitempty"`
	Issr    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 Issr,omitempty"`
}

type GenericIdentification30 struct {
	Id      Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 Id"`
	Issr    Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 Issr"`
	SchmeNm Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 SchmeNm,omitempty"`
}

type GenericIdentification36 struct {
	Id      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 Id"`
	Issr    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 Issr"`
	SchmeNm Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 SchmeNm,omitempty"`
}

type GenericIdentification37 struct {
	Id   Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 Id"`
	Issr Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 Issr,omitempty"`
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
	Cd    ExternalFinancialInstrumentIdentificationType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 Cd"`
	Prtry Max35Text                                          `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 Prtry"`
}

type IdentificationType42Choice struct {
	Cd    TypeOfIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 Cd"`
	Prtry GenericIdentification30   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 Prtry"`
}

// May be one of A001, A002, A003, A004, A005, A006, A007, A008, A009, A010, A011, A012, A013, A014, NARR
type InterestComputationMethod2Code string

type InterestComputationMethodFormat4Choice struct {
	Cd    InterestComputationMethod2Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 Cd"`
	Prtry GenericIdentification30        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 Prtry"`
}

// Must match the pattern [A-Z0-9]{18,18}[0-9]{2,2}
type LEIIdentifier string

// Must match the pattern [A-Z0-9]{4,4}
type MICIdentifier string

type MarketIdentification3Choice struct {
	MktIdrCd MICIdentifier `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 MktIdrCd"`
	Desc     Max35Text     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 Desc"`
}

// Must be at least 1 items long
type Max140Text string

// Must be at least 1 items long
type Max16Text string

// Must be at least 1 items long
type Max210Text string

// Must be at least 1 items long
type Max350Text string

// Must be at least 1 items long
type Max35Text string

// Must match the pattern [0-9]{1,5}
type Max5NumericText string

// Must be at least 1 items long
type Max70Text string

type NameAndAddress5 struct {
	Nm  Max350Text     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 Nm"`
	Adr PostalAddress1 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 Adr,omitempty"`
}

type Number22Choice struct {
	Shrt Exact3NumericText      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 Shrt"`
	Lng  GenericIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 Lng"`
}

type Number3Choice struct {
	Shrt Exact3NumericText `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 Shrt"`
	Lng  Exact5NumericText `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 Lng"`
}

// May be one of AMER, EURO
type OptionStyle2Code string

type OptionStyle8Choice struct {
	Cd    OptionStyle2Code        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 Cd"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 Prtry"`
}

// May be one of CALL, PUTO
type OptionType1Code string

type OptionType6Choice struct {
	Cd    OptionType1Code         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 Cd"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 Prtry"`
}

type OtherAmounts29 struct {
	AcrdIntrstAmt  AmountAndDirection44 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 AcrdIntrstAmt,omitempty"`
	ChrgsFees      AmountAndDirection44 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 ChrgsFees,omitempty"`
	CtryNtlFdrlTax AmountAndDirection44 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 CtryNtlFdrlTax,omitempty"`
	PmtLevyTax     AmountAndDirection44 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 PmtLevyTax,omitempty"`
	LclTax         AmountAndDirection44 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 LclTax,omitempty"`
	Othr           AmountAndDirection44 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 Othr,omitempty"`
	RgltryAmt      AmountAndDirection44 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 RgltryAmt,omitempty"`
	ShppgAmt       AmountAndDirection44 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 ShppgAmt,omitempty"`
	StmpDty        AmountAndDirection44 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 StmpDty,omitempty"`
	StockXchgTax   AmountAndDirection44 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 StockXchgTax,omitempty"`
	TrfTax         AmountAndDirection44 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 TrfTax,omitempty"`
	TxTax          AmountAndDirection44 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 TxTax,omitempty"`
	ValAddedTax    AmountAndDirection44 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 ValAddedTax,omitempty"`
	WhldgTax       AmountAndDirection44 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 WhldgTax,omitempty"`
	CsmptnTax      AmountAndDirection44 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 CsmptnTax,omitempty"`
	AcrdCptlstnAmt AmountAndDirection44 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 AcrdCptlstnAmt,omitempty"`
}

type OtherIdentification1 struct {
	Id  Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 Id"`
	Sfx Max16Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 Sfx,omitempty"`
	Tp  IdentificationSource3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 Tp"`
}

type OtherParties26 struct {
	Invstr    PartyIdentification99  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 Invstr,omitempty"`
	StockXchg PartyIdentification100 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 StockXchg,omitempty"`
	TradRgltr PartyIdentification100 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 TradRgltr,omitempty"`
}

// May be one of A144, NRST, RSTR
type OwnershipLegalRestrictions1Code string

type Pagination struct {
	PgNb      Max5NumericText `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 PgNb"`
	LastPgInd bool            `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 LastPgInd"`
}

type PartyIdentification100 struct {
	Id  PartyIdentification71Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 Id"`
	LEI LEIIdentifier               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 LEI,omitempty"`
}

type PartyIdentification44Choice struct {
	AnyBIC   AnyBICIdentifier `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 AnyBIC"`
	NmAndAdr NameAndAddress5  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 NmAndAdr"`
	Ctry     CountryCode      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 Ctry"`
}

type PartyIdentification71Choice struct {
	AnyBIC   AnyBICIdentifier        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 AnyBIC"`
	PrtryId  GenericIdentification36 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 PrtryId"`
	NmAndAdr NameAndAddress5         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 NmAndAdr"`
}

type PartyIdentification75 struct {
	Id       PartyIdentification44Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 Id"`
	LEI      LEIIdentifier                 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 LEI,omitempty"`
	AltrnId  AlternatePartyIdentification7 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 AltrnId,omitempty"`
	PrcgDt   DateAndDateTimeChoice         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 PrcgDt,omitempty"`
	PrcgId   Max35Text                     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 PrcgId,omitempty"`
	AddtlInf PartyTextInformation1         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 AddtlInf,omitempty"`
}

type PartyIdentification92Choice struct {
	AnyBIC  AnyBICIdentifier        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 AnyBIC"`
	PrtryId GenericIdentification36 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 PrtryId"`
}

type PartyIdentification93Choice struct {
	AnyBIC   AnyBICIdentifier        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 AnyBIC"`
	PrtryId  GenericIdentification36 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 PrtryId"`
	NmAndAdr NameAndAddress5         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 NmAndAdr"`
	Ctry     CountryCode             `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 Ctry"`
}

type PartyIdentification98 struct {
	Id  PartyIdentification92Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 Id"`
	LEI LEIIdentifier               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 LEI,omitempty"`
}

type PartyIdentification99 struct {
	Id  PartyIdentification93Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 Id"`
	LEI LEIIdentifier               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 LEI,omitempty"`
}

type PartyIdentificationAndAccount106 struct {
	Id        PartyIdentification71Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 Id"`
	LEI       LEIIdentifier                 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 LEI,omitempty"`
	AltrnId   AlternatePartyIdentification7 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 AltrnId,omitempty"`
	SfkpgAcct SecuritiesAccount24           `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 SfkpgAcct,omitempty"`
	PrcgDt    DateAndDateTimeChoice         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 PrcgDt,omitempty"`
	PrcgId    Max35Text                     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 PrcgId,omitempty"`
	AddtlInf  PartyTextInformation1         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 AddtlInf,omitempty"`
}

type PartyTextInformation1 struct {
	DclrtnDtls  Max350Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 DclrtnDtls,omitempty"`
	PtyCtctDtls Max140Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 PtyCtctDtls,omitempty"`
	RegnDtls    Max350Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 RegnDtls,omitempty"`
}

type PortfolioTransferNotificationV04 struct {
	Pgntn         Pagination                 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 Pgntn"`
	StmtGnlDtls   Statement46                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 StmtGnlDtls"`
	AcctOwnr      PartyIdentification98      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 AcctOwnr,omitempty"`
	SfkpgAcct     SecuritiesAccount19        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 SfkpgAcct"`
	TrfNtfctnDtls []SecuritiesTradeDetails48 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 TrfNtfctnDtls,omitempty"`
}

type PostalAddress1 struct {
	AdrTp       AddressType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 AdrTp,omitempty"`
	AdrLine     []Max70Text      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 AdrLine,omitempty"`
	StrtNm      Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 StrtNm,omitempty"`
	BldgNb      Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 BldgNb,omitempty"`
	PstCd       Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 PstCd,omitempty"`
	TwnNm       Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 TwnNm,omitempty"`
	CtrySubDvsn Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 CtrySubDvsn,omitempty"`
	Ctry        CountryCode      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 Ctry"`
}

type Price2 struct {
	Tp  YieldedOrValueType1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 Tp"`
	Val PriceRateOrAmountChoice   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 Val"`
}

type PriceRateOrAmountChoice struct {
	Rate float64                                    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 Rate"`
	Amt  ActiveOrHistoricCurrencyAnd13DecimalAmount `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 Amt"`
}

type PriceType1Choice struct {
	Mkt    Price2 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 Mkt"`
	Indctv Price2 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 Indctv"`
}

// May be one of DISC, PREM, PARV
type PriceValueType1Code string

type Quantity11 struct {
	SttlmQty  FinancialInstrumentQuantity1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 SttlmQty"`
	DnmtnChc  Max210Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 DnmtnChc,omitempty"`
	CertNb    []SecuritiesCertificate4           `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 CertNb,omitempty"`
	QtyBrkdwn []QuantityBreakdown30              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 QtyBrkdwn,omitempty"`
}

type QuantityBreakdown30 struct {
	LotNb    GenericIdentification37            `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 LotNb,omitempty"`
	LotQty   FinancialInstrumentQuantity1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 LotQty,omitempty"`
	LotDtTm  DateAndDateTimeChoice              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 LotDtTm,omitempty"`
	LotPric  Price2                             `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 LotPric,omitempty"`
	TpOfPric TypeOfPrice29Choice                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 TpOfPric,omitempty"`
}

// May be one of DELI, RECE
type ReceiveDelivery1Code string

// May be one of NREG, YREG
type Registration1Code string

type Registration9Choice struct {
	Cd    Registration1Code       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 Cd"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 Prtry"`
}

// May be one of STEX, REGU
type Reporting1Code string

type Reporting7Choice struct {
	Cd    Reporting1Code          `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 Cd"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 Prtry"`
}

type Restriction5Choice struct {
	Cd    OwnershipLegalRestrictions1Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 Cd"`
	Prtry GenericIdentification30         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 Prtry"`
}

type SecuritiesAccount19 struct {
	Id Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 Id"`
	Tp GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 Tp,omitempty"`
	Nm Max70Text               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 Nm,omitempty"`
}

type SecuritiesAccount24 struct {
	Id Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 Id"`
	Tp GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 Tp,omitempty"`
	Nm Max70Text               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 Nm,omitempty"`
}

type SecuritiesCertificate4 struct {
	Nb      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 Nb"`
	Issr    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 Issr,omitempty"`
	SchmeNm Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 SchmeNm,omitempty"`
}

// May be one of FULL, NILL, PART
type SecuritiesPaymentStatus1Code string

type SecuritiesPaymentStatus5Choice struct {
	Cd    SecuritiesPaymentStatus1Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 Cd"`
	Prtry GenericIdentification30      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 Prtry"`
}

type SecuritiesRTGS4Choice struct {
	Ind   bool                    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 Ind"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 Prtry"`
}

type SecuritiesTradeDetails48 struct {
	NtfctnSndrTxId   Max35Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 NtfctnSndrTxId,omitempty"`
	NtfctnRcvrTxId   Max35Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 NtfctnRcvrTxId,omitempty"`
	CmonId           Max35Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 CmonId,omitempty"`
	SctiesMvmntTp    ReceiveDelivery1Code            `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 SctiesMvmntTp"`
	Pmt              DeliveryReceiptType2Code        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 Pmt"`
	TradDt           TradeDate5Choice                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 TradDt,omitempty"`
	SttlmDt          SettlementDate9Choice           `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 SttlmDt"`
	NbOfDaysAcrd     float64                         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 NbOfDaysAcrd,omitempty"`
	FinInstrmId      SecurityIdentification19        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 FinInstrmId"`
	FinInstrmAttrbts FinancialInstrumentAttributes64 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 FinInstrmAttrbts,omitempty"`
	Rptg             []Reporting7Choice              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 Rptg,omitempty"`
	QtyDtls          Quantity11                      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 QtyDtls"`
	SttlmParams      SettlementDetails100            `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 SttlmParams,omitempty"`
	DlvrgSttlmPties  SettlementParties36             `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 DlvrgSttlmPties,omitempty"`
	RcvgSttlmPties   SettlementParties36             `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 RcvgSttlmPties,omitempty"`
	SttlmAmt         AmountAndDirection52            `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 SttlmAmt,omitempty"`
	OthrAmts         OtherAmounts29                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 OthrAmts,omitempty"`
	OthrBizPties     OtherParties26                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 OthrBizPties,omitempty"`
	SplmtryData      []SupplementaryData1            `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 SplmtryData,omitempty"`
}

type SecurityIdentification19 struct {
	ISIN   ISINOct2015Identifier  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 ISIN,omitempty"`
	OthrId []OtherIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 OthrId,omitempty"`
	Desc   Max140Text             `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 Desc,omitempty"`
}

// May be one of WISS
type SettlementDate4Code string

type SettlementDate9Choice struct {
	Dt   DateAndDateTimeChoice     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 Dt"`
	DtCd SettlementDateCode7Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 DtCd"`
}

type SettlementDateCode7Choice struct {
	Cd    SettlementDate4Code     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 Cd"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 Prtry"`
}

type SettlementDetails100 struct {
	SttlmTxCond    []SettlementTransactionCondition19Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 SttlmTxCond,omitempty"`
	Regn           Registration9Choice                      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 Regn,omitempty"`
	LglRstrctns    Restriction5Choice                       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 LglRstrctns,omitempty"`
	SctiesRTGS     SecuritiesRTGS4Choice                    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 SctiesRTGS,omitempty"`
	SttlmSysMtd    SettlementSystemMethod4Choice            `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 SttlmSysMtd,omitempty"`
	TaxCpcty       TaxCapacityParty4Choice                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 TaxCpcty,omitempty"`
	StmpDtyTaxBsis GenericIdentification30                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 StmpDtyTaxBsis,omitempty"`
}

type SettlementParties36 struct {
	Dpstry PartyIdentification75            `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 Dpstry,omitempty"`
	Pty1   PartyIdentificationAndAccount106 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 Pty1,omitempty"`
	Pty2   PartyIdentificationAndAccount106 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 Pty2,omitempty"`
	Pty3   PartyIdentificationAndAccount106 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 Pty3,omitempty"`
	Pty4   PartyIdentificationAndAccount106 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 Pty4,omitempty"`
	Pty5   PartyIdentificationAndAccount106 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 Pty5,omitempty"`
}

// May be one of NSET, YSET
type SettlementSystemMethod1Code string

type SettlementSystemMethod4Choice struct {
	Cd    SettlementSystemMethod1Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 Cd"`
	Prtry GenericIdentification30     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 Prtry"`
}

type SettlementTransactionCondition19Choice struct {
	Cd    SettlementTransactionCondition3Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 Cd"`
	Prtry GenericIdentification30             `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 Prtry"`
}

// May be one of ASGN, CLEN, DIRT, DLWM, DRAW, EXER, FRCL, KNOC, PHYS, RESI, SPDL, SPST, UNEX
type SettlementTransactionCondition3Code string

type Statement46 struct {
	CtrPtyPrtflTrfNtfctnRef Max35Text             `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 CtrPtyPrtflTrfNtfctnRef,omitempty"`
	RptNb                   Number3Choice         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 RptNb,omitempty"`
	StmtId                  Max35Text             `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 StmtId,omitempty"`
	StmtDtTm                DateAndDateTimeChoice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 StmtDtTm"`
	UpdTp                   UpdateType15Choice    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 UpdTp,omitempty"`
	ActvtyInd               bool                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 ActvtyInd"`
}

// May be one of COMP, DELT
type StatementUpdateType1Code string

type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
}

type TaxCapacityParty4Choice struct {
	Cd    TaxLiability1Code       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 Cd"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 Prtry"`
}

// May be one of PRIN, AGEN
type TaxLiability1Code string

type TradeDate5Choice struct {
	Dt   DateAndDateTimeChoice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 Dt"`
	DtCd TradeDateCode3Choice  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 DtCd"`
}

type TradeDateCode3Choice struct {
	Cd    DateType3Code           `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 Cd"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 Prtry"`
}

// May be one of ARNU, CCPT, CHTY, CORP, DRLC, FIIN, TXID
type TypeOfIdentification1Code string

// May be one of AVER
type TypeOfPrice14Code string

type TypeOfPrice29Choice struct {
	Cd    TypeOfPrice14Code       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 Cd"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 Prtry"`
}

type UpdateType15Choice struct {
	Cd    StatementUpdateType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 Cd"`
	Prtry GenericIdentification30  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 Prtry"`
}

type YieldedOrValueType1Choice struct {
	Yldd  bool                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 Yldd"`
	ValTp PriceValueType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.04 ValTp"`
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
