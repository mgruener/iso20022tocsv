// Code generated by main. DO NOT EDIT.

package seev_037_001_06

import (
	"bytes"
	"encoding/xml"
	"time"
)

type Account8Choice struct {
	CshAcct   CashAccountIdentification5Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 CshAcct"`
	ChrgsAcct CashAccountIdentification5Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 ChrgsAcct"`
	TaxAcct   CashAccountIdentification5Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 TaxAcct"`
}

type AccountAndBalance4 struct {
	SfkpgAcct Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 SfkpgAcct"`
	ConfdBal  BalanceFormat1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 ConfdBal"`
}

type ActiveCurrencyAndAmount struct {
	Value float64            `xml:",chardata"`
	Ccy   ActiveCurrencyCode `xml:"Ccy,attr"`
}

// Must match the pattern [A-Z]{3,3}
type ActiveCurrencyCode string

// May be one of CLAI, TAXR, ACLA, ATXF, CNTR, CONS, NAMC, NPLE, SCHM
type AdditionalBusinessProcess4Code string

type AdditionalBusinessProcessFormat7Choice struct {
	Cd    AdditionalBusinessProcess4Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 Cd"`
	Prtry GenericIdentification20        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 Prtry"`
}

// May be one of ADDR, PBOX, HOME, BIZZ, MLTO, DLVY
type AddressType2Code string

// Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type AnyBICIdentifier string

type BalanceFormat1Choice struct {
	Bal         SignedQuantityFormat1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 Bal"`
	ElgblBal    SignedQuantityFormat2 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 ElgblBal"`
	NotElgblBal SignedQuantityFormat2 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 NotElgblBal"`
}

type CashAccountIdentification5Choice struct {
	IBAN  IBAN2007Identifier `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 IBAN"`
	Prtry Max34Text          `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 Prtry"`
}

type CashOption19 struct {
	CdtDbtInd    CreditDebitCode         `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 CdtDbtInd"`
	Acct         Account8Choice          `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 Acct,omitempty"`
	PstngDt      DateAndDateTimeChoice   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 PstngDt"`
	OrgnlPstngDt DateAndDateTimeChoice   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 OrgnlPstngDt,omitempty"`
	ValDt        DateAndDateTimeChoice   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 ValDt,omitempty"`
	PstngAmt     ActiveCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 PstngAmt"`
}

type CorporateAction13 struct {
	DtDtls  CorporateActionDate30                  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 DtDtls,omitempty"`
	EvtStag CorporateActionEventStageFormat6Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 EvtStag,omitempty"`
	LtryTp  LotteryTypeFormat1Choice               `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 LtryTp,omitempty"`
}

type CorporateActionDate30 struct {
	RcrdDt   DateFormat19Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 RcrdDt,omitempty"`
	ExDvddDt DateFormat19Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 ExDvddDt,omitempty"`
}

type CorporateActionEventReference1 struct {
	EvtId CorporateActionEventReference1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 EvtId"`
	LkgTp ProcessingPosition1Choice            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 LkgTp,omitempty"`
}

type CorporateActionEventReference1Choice struct {
	LkdOffclCorpActnEvtId Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 LkdOffclCorpActnEvtId"`
	LkdCorpActnId         Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 LkdCorpActnId"`
}

// May be one of FULL, PART, RESC
type CorporateActionEventStage4Code string

type CorporateActionEventStageFormat6Choice struct {
	Cd    CorporateActionEventStage4Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 Cd"`
	Prtry GenericIdentification20        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 Prtry"`
}

// May be one of ACTV, ATTI, BRUP, DFLT, BONU, EXRI, CAPD, CAPG, CAPI, DRCA, DVCA, CHAN, CLSA, COOP, CONS, CONV, CREV, DECR, DETI, DSCL, DVOP, DRIP, DRAW, DTCH, EXOF, REDM, MCAL, INCR, PPMT, INTR, PRII, RHDI, EXTM, MRGR, LIQU, CERT, ODLT, OTHR, PARI, PCAL, PRED, PINK, PRIO, BPUT, REDO, PDEF, PLAC, REMK, BIDS, SPLR, RHTS, DVSC, MTNG, SHPR, SMAL, SOFF, DVSE, SPLF, TREC, TEND, DLST, SUSP, EXWA, WTRC, WRTH, NOOF, ACCU
type CorporateActionEventType11Code string

type CorporateActionEventType12Choice struct {
	Cd    CorporateActionEventType11Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 Cd"`
	Prtry GenericIdentification20        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 Prtry"`
}

type CorporateActionGeneralInformation79 struct {
	CorpActnEvtId          Max35Text                                           `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 CorpActnEvtId"`
	OffclCorpActnEvtId     Max35Text                                           `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 OffclCorpActnEvtId,omitempty"`
	ClssActnNb             Max35Text                                           `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 ClssActnNb,omitempty"`
	EvtTp                  CorporateActionEventType12Choice                    `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 EvtTp"`
	FinInstrmId            SecurityIdentification14                            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 FinInstrmId"`
	AddtlBizPrcInd         AdditionalBusinessProcessFormat7Choice              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 AddtlBizPrcInd,omitempty"`
	IntrmdtSctiesDstrbtnTp IntermediateSecuritiesDistributionTypeFormat5Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 IntrmdtSctiesDstrbtnTp,omitempty"`
	FrctnlQty              FinancialInstrumentQuantity1Choice                  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 FrctnlQty,omitempty"`
}

type CorporateActionMovementReversalAdviceV06 struct {
	MvmntConfId      DocumentIdentification15            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 MvmntConfId"`
	OthrDocId        []DocumentIdentification13          `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 OthrDocId,omitempty"`
	EvtsLkg          []CorporateActionEventReference1    `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 EvtsLkg,omitempty"`
	RvslRsn          CorporateActionReversalReason1      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 RvslRsn,omitempty"`
	CorpActnGnlInf   CorporateActionGeneralInformation79 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 CorpActnGnlInf"`
	AcctDtls         AccountAndBalance4                  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 AcctDtls"`
	CorpActnDtls     CorporateAction13                   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 CorpActnDtls,omitempty"`
	CorpActnConfDtls CorporateActionOption39             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 CorpActnConfDtls"`
	AddtlInf         CorporateActionNarrative4           `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 AddtlInf,omitempty"`
	IssrAgt          []PartyIdentification40Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 IssrAgt,omitempty"`
	PngAgt           []PartyIdentification40Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 PngAgt,omitempty"`
	SubPngAgt        []PartyIdentification40Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 SubPngAgt,omitempty"`
	SplmtryData      []SupplementaryData1                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 SplmtryData,omitempty"`
}

type CorporateActionNarrative4 struct {
	DclrtnDtls    []Max350Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 DclrtnDtls,omitempty"`
	AddtlTxt      []Max350Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 AddtlTxt,omitempty"`
	NrrtvVrsn     []Max350Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 NrrtvVrsn,omitempty"`
	RegnDtls      []Max350Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 RegnDtls,omitempty"`
	InfConds      []Max350Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 InfConds,omitempty"`
	InfToCmplyWth []Max350Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 InfToCmplyWth,omitempty"`
	PtyCtctNrrtv  []Max350Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 PtyCtctNrrtv,omitempty"`
	TaxtnConds    []Max350Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 TaxtnConds,omitempty"`
	BsktOrIndxInf []Max350Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 BsktOrIndxInf,omitempty"`
}

type CorporateActionOption11Choice struct {
	Cd    CorporateActionOption8Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 Cd"`
	Prtry GenericIdentification20    `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 Prtry"`
}

type CorporateActionOption39 struct {
	OptnNb          OptionNumber1Choice           `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 OptnNb"`
	OptnTp          CorporateActionOption11Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 OptnTp"`
	SctiesMvmntDtls []SecuritiesOption27          `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 SctiesMvmntDtls,omitempty"`
	CshMvmntDtls    []CashOption19                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 CshMvmntDtls,omitempty"`
}

// May be one of ABST, AMGT, BSPL, BUYA, CASE, CASH, CEXC, CONN, CONY, CTEN, EXER, LAPS, MKDW, MKUP, MNGT, MPUT, NOAC, NOQU, OFFR, OTHR, OVER, PROX, QINV, SECU, SLLE, SPLI, PRUN
type CorporateActionOption8Code string

type CorporateActionReversalReason1 struct {
	Rsn         CorporateActionReversalReason1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 Rsn"`
	AddtlRsnInf Max256Text                           `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 AddtlRsnInf,omitempty"`
}

type CorporateActionReversalReason1Choice struct {
	Cd    CorporateActionReversalReason1Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 Cd"`
	Prtry GenericIdentification20            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 Prtry"`
}

// May be one of DCBD, IVAD, IRED, IPRI, UPAY, IETR, FNRC, POCS, IPCU
type CorporateActionReversalReason1Code string

// Must match the pattern [A-Z]{2,2}
type CountryCode string

// May be one of CRDT, DBIT
type CreditDebitCode string

type DateAndDateTimeChoice struct {
	Dt   ISODate     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 Dt"`
	DtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 DtTm"`
}

type DateCode11Choice struct {
	Cd    DateType8Code           `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 Cd"`
	Prtry GenericIdentification20 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 Prtry"`
}

type DateFormat19Choice struct {
	Dt   DateAndDateTimeChoice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 Dt"`
	DtCd DateCode11Choice      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 DtCd"`
}

// May be one of UKWN, ONGO
type DateType8Code string

type Document struct {
	CorpActnMvmntRvslAdvc CorporateActionMovementReversalAdviceV06 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 CorpActnMvmntRvslAdvc"`
}

type DocumentIdentification13 struct {
	Id    DocumentIdentification1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 Id"`
	DocNb DocumentNumber1Choice         `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 DocNb,omitempty"`
	LkgTp ProcessingPosition1Choice     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 LkgTp,omitempty"`
}

type DocumentIdentification15 struct {
	Id    Max35Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 Id"`
	LkgTp ProcessingPosition1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 LkgTp,omitempty"`
}

type DocumentIdentification1Choice struct {
	AcctSvcrDocId Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 AcctSvcrDocId"`
	AcctOwnrDocId Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 AcctOwnrDocId"`
}

type DocumentNumber1Choice struct {
	ShrtNb  Exact3NumericText                 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 ShrtNb"`
	LngNb   ISO20022MessageIdentificationText `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 LngNb"`
	PrtryNb GenericIdentification19           `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 PrtryNb"`
}

// Must match the pattern [0-9]{3}
type Exact3NumericText string

// Must match the pattern [a-zA-Z0-9]{4}
type Exact4AlphaNumericText string

// Must be at least 1 items long
type ExternalFinancialInstrumentIdentificationType1Code string

type FinancialInstrumentQuantity1Choice struct {
	Unit     float64 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 Unit"`
	FaceAmt  float64 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 FaceAmt"`
	AmtsdVal float64 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 AmtsdVal"`
}

type GenericIdentification19 struct {
	Id      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 Id"`
	Issr    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 Issr"`
	SchmeNm Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 SchmeNm,omitempty"`
}

type GenericIdentification20 struct {
	Id      Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 Id"`
	Issr    Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 Issr"`
	SchmeNm Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 SchmeNm,omitempty"`
}

// Must match the pattern [A-Z]{2,2}[0-9]{2,2}[a-zA-Z0-9]{1,30}
type IBAN2007Identifier string

// Must match the pattern [A-Z0-9]{12,12}
type ISINIdentifier string

// Must match the pattern [a-z]{4}\.[0-9]{3}\.[0-9]{3}\.[0-9]{2}
type ISO20022MessageIdentificationText string

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
	Cd    ExternalFinancialInstrumentIdentificationType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 Cd"`
	Prtry Max35Text                                          `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 Prtry"`
}

type IntermediateSecuritiesDistributionTypeFormat5Choice struct {
	Cd    IntermediateSecurityDistributionType4Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 Cd"`
	Prtry GenericIdentification20                   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 Prtry"`
}

// May be one of BIDS, DRIP, DVCA, DVOP, EXRI, PRIO, DVSC, DVSE, INTR, LIQU, SOFF, SPLF, BONU, EXOF
type IntermediateSecurityDistributionType4Code string

// May be one of ORIG, SUPP
type LotteryType1Code string

type LotteryTypeFormat1Choice struct {
	Cd    LotteryType1Code        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 Cd"`
	Prtry GenericIdentification20 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 Prtry"`
}

// Must be at least 1 items long
type Max140Text string

// Must be at least 1 items long
type Max16Text string

// Must be at least 1 items long
type Max256Text string

// Must be at least 1 items long
type Max34Text string

// Must be at least 1 items long
type Max350Text string

// Must be at least 1 items long
type Max35Text string

// Must be at least 1 items long
type Max70Text string

type NameAndAddress5 struct {
	Nm  Max350Text     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 Nm"`
	Adr PostalAddress1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 Adr,omitempty"`
}

type OptionNumber1Choice struct {
	Nb Exact3NumericText `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 Nb"`
	Cd OptionNumber1Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 Cd"`
}

// May be one of UNSO
type OptionNumber1Code string

type OriginalAndCurrentQuantities1 struct {
	FaceAmt  float64 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 FaceAmt"`
	AmtsdVal float64 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 AmtsdVal"`
}

type OtherIdentification1 struct {
	Id  Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 Id"`
	Sfx Max16Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 Sfx,omitempty"`
	Tp  IdentificationSource3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 Tp"`
}

type PartyIdentification40Choice struct {
	AnyBIC   AnyBICIdentifier        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 AnyBIC"`
	PrtryId  GenericIdentification19 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 PrtryId"`
	NmAndAdr NameAndAddress5         `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 NmAndAdr"`
}

type PostalAddress1 struct {
	AdrTp       AddressType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 AdrTp,omitempty"`
	AdrLine     []Max70Text      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 AdrLine,omitempty"`
	StrtNm      Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 StrtNm,omitempty"`
	BldgNb      Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 BldgNb,omitempty"`
	PstCd       Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 PstCd,omitempty"`
	TwnNm       Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 TwnNm,omitempty"`
	CtrySubDvsn Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 CtrySubDvsn,omitempty"`
	Ctry        CountryCode      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 Ctry"`
}

type ProcessingPosition1Choice struct {
	Cd    ProcessingPosition3Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 Cd"`
	Prtry GenericIdentification20 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 Prtry"`
}

// May be one of AFTE, WITH, BEFO, INFO
type ProcessingPosition3Code string

type ProprietaryQuantity2 struct {
	Qty     float64                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 Qty"`
	QtyTp   Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 QtyTp"`
	Issr    Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 Issr"`
	SchmeNm Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 SchmeNm,omitempty"`
}

type Quantity2Choice struct {
	Qty      FinancialInstrumentQuantity1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 Qty"`
	PrtryQty ProprietaryQuantity2               `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 PrtryQty"`
}

type Quantity6Choice struct {
	Qty             FinancialInstrumentQuantity1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 Qty"`
	OrgnlAndCurFace OriginalAndCurrentQuantities1      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 OrgnlAndCurFace"`
}

type SecuritiesOption27 struct {
	FinInstrmId  SecurityIdentification14 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 FinInstrmId"`
	CdtDbtInd    CreditDebitCode          `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 CdtDbtInd"`
	PstngQty     Quantity6Choice          `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 PstngQty"`
	PstngDt      DateAndDateTimeChoice    `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 PstngDt"`
	OrgnlPstngDt DateAndDateTimeChoice    `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 OrgnlPstngDt,omitempty"`
}

type SecurityIdentification14 struct {
	ISIN   ISINIdentifier         `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 ISIN,omitempty"`
	OthrId []OtherIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 OthrId,omitempty"`
	Desc   Max140Text             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 Desc,omitempty"`
}

// May be one of SHOR, LONG
type ShortLong1Code string

type SignedQuantityFormat1 struct {
	ShrtLngPos ShortLong1Code  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 ShrtLngPos"`
	QtyChc     Quantity2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 QtyChc"`
}

type SignedQuantityFormat2 struct {
	ShrtLngPos ShortLong1Code                     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 ShrtLngPos"`
	Qty        FinancialInstrumentQuantity1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 Qty"`
}

type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.06 Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
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
