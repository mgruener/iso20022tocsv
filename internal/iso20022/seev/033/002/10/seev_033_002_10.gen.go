// Code generated by main. DO NOT EDIT.

package seev_033_002_10

import (
	"bytes"
	"encoding/xml"
	"time"
)

type AccountAndBalance44 struct {
	SfkpgAcct RestrictedFINXMax35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 SfkpgAcct"`
	AcctOwnr  PartyIdentification136Choice    `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 AcctOwnr,omitempty"`
	SfkpgPlc  SafekeepingPlaceFormat32Choice  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 SfkpgPlc,omitempty"`
	Bal       CorporateActionBalanceDetails34 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 Bal,omitempty"`
}

// Must match the pattern [A-Z]{3,3}
type ActiveCurrencyCode string

// Must match the pattern [A-Z]{3,3}
type ActiveOrHistoricCurrencyCode string

type AlternatePartyIdentification9 struct {
	IdTp    IdentificationType44Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 IdTp"`
	Ctry    CountryCode                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 Ctry"`
	AltrnId RestrictedFINXMax30Text    `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 AltrnId"`
}

type AmountPrice5 struct {
	AmtPricTp AmountPriceType1Code                          `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 AmtPricTp"`
	PricVal   RestrictedFINActiveCurrencyAnd13DecimalAmount `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 PricVal"`
}

type AmountPricePerAmount3 struct {
	AmtPricTp AmountPriceType1Code                          `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 AmtPricTp"`
	PricVal   RestrictedFINActiveCurrencyAnd13DecimalAmount `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 PricVal"`
	Amt       RestrictedFINActiveCurrencyAnd13DecimalAmount `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 Amt"`
}

type AmountPricePerFinancialInstrumentQuantity7 struct {
	AmtPricTp    AmountPriceType1Code                          `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 AmtPricTp"`
	PricVal      RestrictedFINActiveCurrencyAnd13DecimalAmount `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 PricVal"`
	FinInstrmQty FinancialInstrumentQuantity15Choice           `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 FinInstrmQty"`
}

// May be one of ACTU, DISC, PLOT, PREM
type AmountPriceType1Code string

// Must match the pattern [A-Z0-9]{4,4}[A-Z]{2,2}[A-Z0-9]{2,2}([A-Z0-9]{3,3}){0,1}
type AnyBICDec2014Identifier string

type BalanceFormat7Choice struct {
	Bal         SignedQuantityFormat8 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 Bal"`
	ElgblBal    SignedQuantityFormat9 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 ElgblBal"`
	NotElgblBal SignedQuantityFormat9 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 NotElgblBal"`
}

type BeneficiaryCertificationType11Choice struct {
	Cd    BeneficiaryCertificationType5Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 Cd"`
	Prtry GenericIdentification47           `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 Prtry"`
}

// May be one of ACCI, NCOM, QIBB
type BeneficiaryCertificationType5Code string

// Must match the pattern [A-Z]{6,6}
type CFIOct2015Identifier string

type ClassificationType33Choice struct {
	ClssfctnFinInstrm CFIOct2015Identifier    `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 ClssfctnFinInstrm"`
	AltrnClssfctn     GenericIdentification86 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 AltrnClssfctn"`
}

type CorporateActionBalanceDetails34 struct {
	TtlElgblBal      Quantity22Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 TtlElgblBal,omitempty"`
	BlckdBal         BalanceFormat7Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 BlckdBal,omitempty"`
	BrrwdBal         BalanceFormat7Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 BrrwdBal,omitempty"`
	CollInBal        BalanceFormat7Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 CollInBal,omitempty"`
	CollOutBal       BalanceFormat7Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 CollOutBal,omitempty"`
	OnLnBal          BalanceFormat7Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 OnLnBal,omitempty"`
	PdgDlvryBal      []BalanceFormat7Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 PdgDlvryBal,omitempty"`
	PdgRctBal        []BalanceFormat7Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 PdgRctBal,omitempty"`
	OutForRegnBal    BalanceFormat7Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 OutForRegnBal,omitempty"`
	SttlmPosBal      BalanceFormat7Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 SttlmPosBal,omitempty"`
	StrtPosBal       BalanceFormat7Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 StrtPosBal,omitempty"`
	TradDtPosBal     BalanceFormat7Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 TradDtPosBal,omitempty"`
	InTrnsShipmntBal BalanceFormat7Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 InTrnsShipmntBal,omitempty"`
	RegdBal          BalanceFormat7Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 RegdBal,omitempty"`
}

// May be one of BERE, CERT, DEPH, GPPH, GTGP, GTPH, NAME, PHDE, REBE, TERM
type CorporateActionChangeType2Code string

type CorporateActionChangeTypeFormat7Choice struct {
	Cd    CorporateActionChangeType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 Cd"`
	Prtry GenericIdentification47        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 Prtry"`
}

type CorporateActionEventReference4 struct {
	EvtId CorporateActionEventReference4Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 EvtId"`
	LkgTp ProcessingPosition10Choice           `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 LkgTp,omitempty"`
}

type CorporateActionEventReference4Choice struct {
	LkdOffclCorpActnEvtId RestrictedFINXMax16Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 LkdOffclCorpActnEvtId"`
	LkdCorpActnId         RestrictedFINXMax16Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 LkdCorpActnId"`
}

// May be one of ACTV, ATTI, BRUP, DFLT, BONU, EXRI, CAPD, CAPG, CAPI, DRCA, DVCA, CHAN, COOP, CLSA, CONS, CONV, CREV, DECR, DETI, DSCL, DVOP, DRIP, DRAW, DTCH, EXOF, REDM, MCAL, INCR, PPMT, INTR, RHDI, LIQU, EXTM, MRGR, NOOF, CERT, ODLT, OTHR, PARI, PCAL, PRED, PINK, PLAC, PDEF, PRIO, BPUT, REDO, REMK, BIDS, SPLR, RHTS, DVSC, SHPR, SMAL, SOFF, DVSE, SPLF, TREC, TEND, DLST, SUSP, EXWA, WTRC, WRTH
type CorporateActionEventType29Code string

type CorporateActionEventType90Choice struct {
	Cd    CorporateActionEventType29Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 Cd"`
	Prtry GenericIdentification47        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 Prtry"`
}

type CorporateActionGeneralInformation148 struct {
	CorpActnEvtId      RestrictedFINXMax16Text          `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 CorpActnEvtId"`
	OffclCorpActnEvtId RestrictedFINXMax16Text          `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 OffclCorpActnEvtId,omitempty"`
	EvtTp              CorporateActionEventType90Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 EvtTp"`
	UndrlygScty        FinancialInstrumentAttributes84  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 UndrlygScty,omitempty"`
}

type CorporateActionInstruction002V10 struct {
	ChngInstrInd   bool                                 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 ChngInstrInd,omitempty"`
	CancInstrId    DocumentIdentification37             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 CancInstrId,omitempty"`
	InstrCxlReqId  DocumentIdentification37             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 InstrCxlReqId,omitempty"`
	OthrDocId      []DocumentIdentification38           `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 OthrDocId,omitempty"`
	EvtsLkg        []CorporateActionEventReference4     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 EvtsLkg,omitempty"`
	CorpActnGnlInf CorporateActionGeneralInformation148 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 CorpActnGnlInf"`
	AcctDtls       AccountAndBalance44                  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 AcctDtls"`
	BnfclOwnrDtls  []PartyIdentification234             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 BnfclOwnrDtls,omitempty"`
	CorpActnInstr  CorporateActionOption172             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 CorpActnInstr"`
	PrtctInstr     ProtectInstruction5                  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 PrtctInstr,omitempty"`
	AddtlInf       CorporateActionNarrative34           `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 AddtlInf,omitempty"`
	SplmtryData    []SupplementaryData1                 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 SplmtryData,omitempty"`
}

type CorporateActionNarrative33 struct {
	InfToCmplyWth    []RestrictedFINXMax350Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 InfToCmplyWth,omitempty"`
	DlvryDtls        []RestrictedFINXMax350Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 DlvryDtls,omitempty"`
	FXInstrsAddtlInf []RestrictedFINXMax350Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 FXInstrsAddtlInf,omitempty"`
	InstrAddtlInf    []RestrictedFINXMax350Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 InstrAddtlInf,omitempty"`
}

type CorporateActionNarrative34 struct {
	RegnDtls       []RestrictedFINXMax350Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 RegnDtls,omitempty"`
	PtyCtctNrrtv   []RestrictedFINXMax350Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 PtyCtctNrrtv,omitempty"`
	CertfctnBrkdwn []RestrictedFINXMax350Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 CertfctnBrkdwn,omitempty"`
}

// May be one of ABST, BSPL, BUYA, CASE, CASH, CERT, CEXC, CONN, CONY, CTEN, EXER, LAPS, MKDW, MKUP, MPUT, NOAC, NOQU, OFFR, OTHR, OVER, QINV, SECU, SLLE, TAXI, PRUN
type CorporateActionOption13Code string

type CorporateActionOption172 struct {
	OptnNb              OptionNumber1Choice                      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 OptnNb"`
	OptnTp              CorporateActionOption35Choice            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 OptnTp"`
	OptnFeatrs          OptionFeaturesFormat27Choice             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 OptnFeatrs,omitempty"`
	FrctnDspstn         FractionDispositionType29Choice          `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 FrctnDspstn,omitempty"`
	ChngTp              []CorporateActionChangeTypeFormat7Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 ChngTp,omitempty"`
	ElgblForCollInd     bool                                     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 ElgblForCollInd,omitempty"`
	CcyToBuy            ActiveCurrencyCode                       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 CcyToBuy,omitempty"`
	CcyToSell           ActiveCurrencyCode                       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 CcyToSell,omitempty"`
	CcyOptn             ActiveCurrencyCode                       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 CcyOptn,omitempty"`
	SctyId              SecurityIdentification20                 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 SctyId,omitempty"`
	SctiesQtyOrInstdAmt SecuritiesQuantityOrAmount5Choice        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 SctiesQtyOrInstdAmt"`
	ExctnReqdDtTm       DateAndDateTime2Choice                   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 ExctnReqdDtTm,omitempty"`
	RateAndAmtDtls      CorporateActionRate73                    `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 RateAndAmtDtls,omitempty"`
	PricDtls            CorporateActionPrice62                   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 PricDtls,omitempty"`
	ShrhldrNb           RestrictedFINXMax25Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 ShrhldrNb,omitempty"`
	AddtlInf            CorporateActionNarrative33               `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 AddtlInf,omitempty"`
}

type CorporateActionOption35Choice struct {
	Cd    CorporateActionOption13Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 Cd"`
	Prtry GenericIdentification47     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 Prtry"`
}

type CorporateActionPrice62 struct {
	IndctvOrMktPric       IndicativeOrMarketPrice9Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 IndctvOrMktPric,omitempty"`
	IssePric              PriceFormat52Choice            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 IssePric,omitempty"`
	GncCshPricRcvdPerPdct PriceFormat53Choice            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 GncCshPricRcvdPerPdct,omitempty"`
	GncCshPricPdPerPdct   PriceFormat52Choice            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 GncCshPricPdPerPdct,omitempty"`
}

type CorporateActionRate73 struct {
	PropsdRate         float64                       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 PropsdRate,omitempty"`
	OvrsbcptRate       RateAndAmountFormat43Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 OvrsbcptRate,omitempty"`
	ReqdWhldgTaxRate   []RateAndAmountFormat45Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 ReqdWhldgTaxRate,omitempty"`
	ReqdScndLvlTaxRate []RateAndAmountFormat45Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 ReqdScndLvlTaxRate,omitempty"`
}

// Must match the pattern [A-Z]{2,2}
type CountryCode string

type DateAndDateTime2Choice struct {
	Dt   ISODate     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 Dt"`
	DtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 DtTm"`
}

type Document struct {
	CorpActnInstr CorporateActionInstruction002V10 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 CorpActnInstr"`
}

type DocumentIdentification37 struct {
	Id    RestrictedFINXMax16Text    `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 Id"`
	LkgTp ProcessingPosition10Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 LkgTp,omitempty"`
}

type DocumentIdentification38 struct {
	Id    DocumentIdentification4Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 Id"`
	DocNb DocumentNumber6Choice         `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 DocNb,omitempty"`
	LkgTp ProcessingPosition10Choice    `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 LkgTp,omitempty"`
}

type DocumentIdentification4Choice struct {
	AcctSvcrDocId RestrictedFINXMax16Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 AcctSvcrDocId"`
	AcctOwnrDocId RestrictedFINXMax16Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 AcctOwnrDocId"`
}

type DocumentNumber6Choice struct {
	ShrtNb  Exact3NumericText                 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 ShrtNb"`
	LngNb   ISO20022MessageIdentificationText `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 LngNb"`
	PrtryNb GenericIdentification86           `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 PrtryNb"`
}

// Must match the pattern [0-9]{3}
type Exact3NumericText string

// Must match the pattern [a-zA-Z0-9]{4}
type Exact4AlphaNumericText string

// Must be at least 1 items long
type ExternalFinancialInstrumentIdentificationType1Code string

type FinancialInstrumentAttributes84 struct {
	FinInstrmId   SecurityIdentification20               `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 FinInstrmId,omitempty"`
	PlcOfListg    MarketIdentification4Choice            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 PlcOfListg,omitempty"`
	DayCntBsis    InterestComputationMethodFormat5Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 DayCntBsis,omitempty"`
	ClssfctnTp    ClassificationType33Choice             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 ClssfctnTp,omitempty"`
	DnmtnCcy      ActiveOrHistoricCurrencyCode           `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 DnmtnCcy,omitempty"`
	NxtCpnDt      ISODate                                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 NxtCpnDt,omitempty"`
	XpryDt        ISODate                                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 XpryDt,omitempty"`
	FltgRateFxgDt ISODate                                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 FltgRateFxgDt,omitempty"`
	MtrtyDt       ISODate                                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 MtrtyDt,omitempty"`
	IsseDt        ISODate                                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 IsseDt,omitempty"`
	NxtCllblDt    ISODate                                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 NxtCllblDt,omitempty"`
	PutblDt       ISODate                                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 PutblDt,omitempty"`
	DtdDt         ISODate                                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 DtdDt,omitempty"`
	ConvsDt       ISODate                                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 ConvsDt,omitempty"`
	PrvsFctr      float64                                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 PrvsFctr,omitempty"`
	NxtFctr       float64                                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 NxtFctr,omitempty"`
	IntrstRate    float64                                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 IntrstRate,omitempty"`
	NxtIntrstRate float64                                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 NxtIntrstRate,omitempty"`
	MinNmnlQty    FinancialInstrumentQuantity15Choice    `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 MinNmnlQty,omitempty"`
	CtrctSz       FinancialInstrumentQuantity15Choice    `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 CtrctSz,omitempty"`
}

type FinancialInstrumentQuantity15Choice struct {
	Unit     float64 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 Unit"`
	FaceAmt  float64 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 FaceAmt"`
	AmtsdVal float64 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 AmtsdVal"`
}

// May be one of BUYU, CINL, EXPI, DIST
type FractionDispositionType10Code string

type FractionDispositionType29Choice struct {
	Cd    FractionDispositionType10Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 Cd"`
	Prtry GenericIdentification47       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 Prtry"`
}

type GenericIdentification47 struct {
	Id      Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 Id"`
	Issr    Max4AlphaNumericText   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 Issr"`
	SchmeNm Max4AlphaNumericText   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 SchmeNm,omitempty"`
}

type GenericIdentification84 struct {
	Id      RestrictedFINXMax34Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 Id"`
	Issr    Max4AlphaNumericText    `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 Issr"`
	SchmeNm Max4AlphaNumericText    `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 SchmeNm,omitempty"`
}

type GenericIdentification85 struct {
	Tp GenericIdentification47 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 Tp"`
	Id RestrictedFINXMax30Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 Id,omitempty"`
}

type GenericIdentification86 struct {
	Id      RestrictedFINXMax30Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 Id"`
	Issr    Max4AlphaNumericText    `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 Issr"`
	SchmeNm Max4AlphaNumericText    `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 SchmeNm,omitempty"`
}

// Must match the pattern [A-Z]{2,2}[A-Z0-9]{9,9}[0-9]{1,1}
type ISINOct2015Identifier string

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

type IdentificationSource4Choice struct {
	Cd    ExternalFinancialInstrumentIdentificationType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 Cd"`
	Prtry RestrictedFINExact2Text                            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 Prtry"`
}

type IdentificationType44Choice struct {
	Cd    TypeOfIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 Cd"`
	Prtry GenericIdentification47   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 Prtry"`
}

type IndicativeOrMarketPrice9Choice struct {
	IndctvPric PriceFormat52Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 IndctvPric"`
	MktPric    PriceFormat52Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 MktPric"`
}

// May be one of A001, A002, A003, A004, A005, A006, A007, A008, A009, A010, A011, A012, A013, A014, NARR
type InterestComputationMethod2Code string

type InterestComputationMethodFormat5Choice struct {
	Cd    InterestComputationMethod2Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 Cd"`
	Prtry GenericIdentification47        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 Prtry"`
}

// Must match the pattern [A-Z0-9]{18,18}[0-9]{2,2}
type LEIIdentifier string

// Must match the pattern [A-Z0-9]{4,4}
type MICIdentifier string

type MarketIdentification4Choice struct {
	MktIdrCd MICIdentifier           `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 MktIdrCd"`
	Desc     RestrictedFINXMax30Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 Desc"`
}

// Must be at least 1 items long
type Max16Text string

// Must be at least 1 items long
type Max350Text string

// Must match the pattern [a-zA-Z0-9]{1,4}
type Max4AlphaNumericText string

type NameAndAddress12 struct {
	Nm RestrictedFINXMax140Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 Nm"`
}

// May be one of OPLF
type OptionFeatures12Code string

type OptionFeaturesFormat27Choice struct {
	Cd    OptionFeatures12Code    `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 Cd"`
	Prtry GenericIdentification47 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 Prtry"`
}

type OptionNumber1Choice struct {
	Nb Exact3NumericText `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 Nb"`
	Cd OptionNumber1Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 Cd"`
}

// May be one of UNSO
type OptionNumber1Code string

type OriginalAndCurrentQuantities4 struct {
	FaceAmt  float64 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 FaceAmt"`
	AmtsdVal float64 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 AmtsdVal"`
}

type OriginalAndCurrentQuantities7 struct {
	ShrtLngPos ShortLong1Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 ShrtLngPos"`
	FaceAmt    float64        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 FaceAmt"`
	AmtsdVal   float64        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 AmtsdVal"`
}

type OtherIdentification2 struct {
	Id  RestrictedFINXMax31Text     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 Id"`
	Sfx Max16Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 Sfx,omitempty"`
	Tp  IdentificationSource4Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 Tp"`
}

type PartyIdentification136Choice struct {
	AnyBIC  AnyBICDec2014Identifier `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 AnyBIC"`
	PrtryId GenericIdentification84 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 PrtryId"`
}

type PartyIdentification137Choice struct {
	AnyBIC   AnyBICDec2014Identifier `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 AnyBIC"`
	PrtryId  GenericIdentification84 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 PrtryId"`
	NmAndAdr NameAndAddress12        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 NmAndAdr"`
}

type PartyIdentification234 struct {
	OwnrId         PartyIdentification137Choice           `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 OwnrId"`
	LEIId          LEIIdentifier                          `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 LEIId,omitempty"`
	AltrnId        []AlternatePartyIdentification9        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 AltrnId,omitempty"`
	DmclCtry       CountryCode                            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 DmclCtry,omitempty"`
	NonDmclCtry    []CountryCode                          `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 NonDmclCtry,omitempty"`
	OwndSctiesQty  FinancialInstrumentQuantity15Choice    `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 OwndSctiesQty"`
	CertfctnTp     []BeneficiaryCertificationType11Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 CertfctnTp,omitempty"`
	WhldgTaxRate   RateAndAmountFormat46Choice            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 WhldgTaxRate,omitempty"`
	CertfctnBrkdwn []RestrictedFINXMax350Text             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 CertfctnBrkdwn,omitempty"`
}

type PercentagePrice1 struct {
	PctgPricTp PriceRateType3Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 PctgPricTp"`
	PricVal    float64            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 PricVal"`
}

type PriceFormat52Choice struct {
	PctgPric PercentagePrice1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 PctgPric"`
	AmtPric  AmountPrice5     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 AmtPric"`
}

type PriceFormat53Choice struct {
	PctgPric               PercentagePrice1                           `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 PctgPric"`
	AmtPric                AmountPrice5                               `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 AmtPric"`
	NotSpcfdPric           PriceValueType9Code                        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 NotSpcfdPric"`
	AmtPricPerFinInstrmQty AmountPricePerFinancialInstrumentQuantity7 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 AmtPricPerFinInstrmQty"`
	AmtPricPerAmt          AmountPricePerAmount3                      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 AmtPricPerAmt"`
	IndxPts                float64                                    `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 IndxPts"`
}

// May be one of DISC, PREM, PRCT, YIEL
type PriceRateType3Code string

// May be one of TBSP, UNSP, UKWN
type PriceValueType9Code string

type ProcessingPosition10Choice struct {
	Cd    ProcessingPosition3Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 Cd"`
	Prtry GenericIdentification47 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 Prtry"`
}

// May be one of AFTE, WITH, BEFO, INFO
type ProcessingPosition3Code string

type ProprietaryQuantity10 struct {
	ShrtLngPos ShortLong1Code         `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 ShrtLngPos,omitempty"`
	Qty        float64                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 Qty"`
	QtyTp      Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 QtyTp"`
	Issr       Max4AlphaNumericText   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 Issr"`
	SchmeNm    Max4AlphaNumericText   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 SchmeNm,omitempty"`
}

type ProprietaryQuantity9 struct {
	Qty     float64                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 Qty"`
	QtyTp   Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 QtyTp"`
	Issr    Max4AlphaNumericText   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 Issr"`
	SchmeNm Max4AlphaNumericText   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 SchmeNm,omitempty"`
}

type ProtectInstruction5 struct {
	TxTp           ProtectTransactionType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 TxTp"`
	TxId           RestrictedFINMax15Text      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 TxId,omitempty"`
	PrtctSfkpgAcct RestrictedFINMax35Text      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 PrtctSfkpgAcct,omitempty"`
	PrtctDt        ISODate                     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 PrtctDt,omitempty"`
}

// May be one of PROT, COVP, COVR
type ProtectTransactionType2Code string

// May be one of QALL
type Quantity1Code string

type Quantity21Choice struct {
	Qty      FinancialInstrumentQuantity15Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 Qty"`
	PrtryQty ProprietaryQuantity9                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 PrtryQty"`
}

type Quantity22Choice struct {
	QtyChc   Quantity23Choice      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 QtyChc"`
	PrtryQty ProprietaryQuantity10 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 PrtryQty"`
}

type Quantity23Choice struct {
	OrgnlAndCurFaceAmt OriginalAndCurrentQuantities7 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 OrgnlAndCurFaceAmt"`
	SgndQty            SignedQuantityFormat9         `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 SgndQty"`
}

type Quantity40Choice struct {
	Cd                 Quantity1Code                       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 Cd"`
	OrgnlAndCurFaceAmt OriginalAndCurrentQuantities4       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 OrgnlAndCurFaceAmt"`
	Qty                FinancialInstrumentQuantity15Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 Qty"`
}

type RateAndAmountFormat43Choice struct {
	Rate float64                                       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 Rate"`
	Amt  RestrictedFINActiveCurrencyAnd13DecimalAmount `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 Amt"`
}

type RateAndAmountFormat45Choice struct {
	Rate          float64                                       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 Rate"`
	Amt           RestrictedFINActiveCurrencyAnd13DecimalAmount `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 Amt"`
	RateTpAndRate RateTypeAndPercentageRate9                    `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 RateTpAndRate"`
}

type RateAndAmountFormat46Choice struct {
	Rate         float64                                       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 Rate"`
	NotSpcfdRate RateValueType7Code                            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 NotSpcfdRate"`
	Amt          RestrictedFINActiveCurrencyAnd13DecimalAmount `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 Amt"`
}

type RateType46Choice struct {
	Cd    WithholdingTaxRateType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 Cd"`
	Prtry GenericIdentification47     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 Prtry"`
}

type RateTypeAndPercentageRate9 struct {
	RateTp RateType46Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 RateTp"`
	Rate   float64          `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 Rate"`
}

// May be one of UKWN
type RateValueType7Code string

type RestrictedFINActiveCurrencyAnd13DecimalAmount struct {
	Value float64            `xml:",chardata"`
	Ccy   ActiveCurrencyCode `xml:"Ccy,attr"`
}

type RestrictedFINActiveCurrencyAndAmount struct {
	Value float64            `xml:",chardata"`
	Ccy   ActiveCurrencyCode `xml:"Ccy,attr"`
}

// Must match the pattern XX|TS
type RestrictedFINExact2Text string

// Must be at least 1 items long
type RestrictedFINMax15Text string

// Must match the pattern ([^/]+/)+([^/]+)|([^/]*)
type RestrictedFINMax35Text string

// Must match the pattern [0-9a-zA-Z/\-\?:\(\)\.\n\r,'\+ ]{1,140}
type RestrictedFINXMax140Text string

// Must match the pattern ([0-9a-zA-Z\-\?:\(\)\.,'\+ ]([0-9a-zA-Z\-\?:\(\)\.,'\+ ]*(/[0-9a-zA-Z\-\?:\(\)\.,'\+ ])?)*)
type RestrictedFINXMax16Text string

// Must match the pattern ([0-9a-zA-Z\-\?:\(\)\.,'\+ ]([0-9a-zA-Z\-\?:\(\)\.,'\+ ]*(/[0-9a-zA-Z\-\?:\(\)\.,'\+ ])?)*)
type RestrictedFINXMax25Text string

// Must match the pattern ([0-9a-zA-Z\-\?:\(\)\.,'\+ ]([0-9a-zA-Z\-\?:\(\)\.,'\+ ]*(/[0-9a-zA-Z\-\?:\(\)\.,'\+ ])?)*)
type RestrictedFINXMax30Text string

// Must match the pattern [0-9a-zA-Z/\-\?:\(\)\.,'\+ ]{1,31}
type RestrictedFINXMax31Text string

// Must match the pattern ([0-9a-zA-Z\-\?:\(\)\.,'\+ ]([0-9a-zA-Z\-\?:\(\)\.,'\+ ]*(/[0-9a-zA-Z\-\?:\(\)\.,'\+ ])?)*)
type RestrictedFINXMax34Text string

// Must match the pattern [0-9a-zA-Z/\-\?:\(\)\.\n\r,'\+ ]{1,350}
type RestrictedFINXMax350Text string

// Must match the pattern [0-9a-zA-Z/\-\?:\(\)\.,'\+ ]{1,35}
type RestrictedFINXMax35Text string

// May be one of CUST, ICSD, NCSD, SHHE
type SafekeepingPlace1Code string

// May be one of SHHE, ALLP
type SafekeepingPlace2Code string

type SafekeepingPlaceFormat32Choice struct {
	Id      SafekeepingPlaceTypeAndText9           `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 Id"`
	Ctry    CountryCode                            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 Ctry"`
	TpAndId SafekeepingPlaceTypeAndIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 TpAndId"`
	Prtry   GenericIdentification85                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 Prtry"`
}

type SafekeepingPlaceTypeAndIdentification1 struct {
	SfkpgPlcTp SafekeepingPlace1Code   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 SfkpgPlcTp"`
	Id         AnyBICDec2014Identifier `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 Id"`
}

type SafekeepingPlaceTypeAndText9 struct {
	SfkpgPlcTp SafekeepingPlace2Code   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 SfkpgPlcTp"`
	Id         RestrictedFINXMax30Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 Id,omitempty"`
}

type SecuritiesOption73 struct {
	CondlQty      FinancialInstrumentQuantity15Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 CondlQty,omitempty"`
	InstdQty      Quantity40Choice                    `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 InstdQty"`
	AddtlRndUpQty FinancialInstrumentQuantity15Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 AddtlRndUpQty,omitempty"`
}

type SecuritiesQuantityOrAmount5Choice struct {
	SctiesQty SecuritiesOption73                   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 SctiesQty"`
	InstdAmt  RestrictedFINActiveCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 InstdAmt"`
}

type SecurityIdentification20 struct {
	ISIN   ISINOct2015Identifier    `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 ISIN,omitempty"`
	OthrId []OtherIdentification2   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 OthrId,omitempty"`
	Desc   RestrictedFINXMax140Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 Desc,omitempty"`
}

// May be one of SHOR, LONG
type ShortLong1Code string

type SignedQuantityFormat8 struct {
	ShrtLngPos ShortLong1Code   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 ShrtLngPos"`
	QtyChc     Quantity21Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 QtyChc"`
}

type SignedQuantityFormat9 struct {
	ShrtLngPos ShortLong1Code                      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 ShrtLngPos"`
	Qty        FinancialInstrumentQuantity15Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 Qty"`
}

type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.10 Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
}

// May be one of ARNU, CCPT, CHTY, CORP, DRLC, FIIN, TXID
type TypeOfIdentification1Code string

// May be one of BWIT, FTCA, NRAT
type WithholdingTaxRateType1Code string

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
