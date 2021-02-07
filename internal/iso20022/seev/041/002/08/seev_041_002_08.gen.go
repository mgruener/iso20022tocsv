// Code generated by main. DO NOT EDIT.

package seev_041_002_08

// Must match the pattern [A-Z]{3,3}
type ActiveCurrencyCode string

// Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type AnyBICIdentifier string

type CancelledReason10Choice struct {
	Cd    CancelledStatusReason5Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.002.08 Cd"`
	Prtry GenericIdentification47    `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.002.08 Prtry"`
}

type CancelledStatus14Choice struct {
	NoSpcfdRsn NoReasonCode              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.002.08 NoSpcfdRsn"`
	Rsn        []CancelledStatusReason13 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.002.08 Rsn"`
}

type CancelledStatusReason13 struct {
	RsnCd       CancelledReason10Choice  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.002.08 RsnCd"`
	AddtlRsnInf RestrictedFINXMax210Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.002.08 AddtlRsnInf,omitempty"`
}

// May be one of CANI, OTHR
type CancelledStatusReason5Code string

type CashAccountIdentification6Choice struct {
	IBAN  IBAN2007Identifier       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.002.08 IBAN"`
	Prtry RestrictedFINX2Max34Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.002.08 Prtry"`
}

// May be one of ACTV, ATTI, BRUP, DFLT, BONU, EXRI, CAPD, CAPG, CAPI, DRCA, DVCA, CHAN, COOP, CLSA, CONS, CONV, CREV, DECR, DETI, DSCL, DVOP, DRIP, DRAW, DTCH, EXOF, REDM, MCAL, INCR, PPMT, INTR, PRII, RHDI, LIQU, EXTM, MRGR, NOOF, CERT, ODLT, OTHR, PARI, PCAL, PRED, PINK, PLAC, PDEF, PRIO, BPUT, REDO, REMK, BIDS, SPLR, RHTS, DVSC, SHPR, SMAL, SOFF, DVSE, SPLF, TREC, TEND, DLST, SUSP, EXWA, WTRC, WRTH
type CorporateActionEventType20Code string

type CorporateActionEventType58Choice struct {
	Cd    CorporateActionEventType20Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.002.08 Cd"`
	Prtry GenericIdentification47        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.002.08 Prtry"`
}

type CorporateActionGeneralInformation114 struct {
	CorpActnEvtId      RestrictedFINXMax16Text          `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.002.08 CorpActnEvtId"`
	OffclCorpActnEvtId RestrictedFINXMax16Text          `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.002.08 OffclCorpActnEvtId,omitempty"`
	ClssActnNb         RestrictedFINXMax16Text          `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.002.08 ClssActnNb,omitempty"`
	EvtTp              CorporateActionEventType58Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.002.08 EvtTp"`
}

type CorporateActionInstructionCancellationRequestStatusAdvice002V08 struct {
	InstrCxlReqId  DocumentIdentification17                       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.002.08 InstrCxlReqId,omitempty"`
	OthrDocId      []DocumentIdentification34                     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.002.08 OthrDocId,omitempty"`
	CorpActnGnlInf CorporateActionGeneralInformation114           `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.002.08 CorpActnGnlInf"`
	InstrCxlReqSts []InstructionCancellationRequestStatus12Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.002.08 InstrCxlReqSts"`
	CorpActnInstr  CorporateActionOption144                       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.002.08 CorpActnInstr,omitempty"`
	AddtlInf       CorporateActionNarrative19                     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.002.08 AddtlInf,omitempty"`
	SplmtryData    []SupplementaryData1                           `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.002.08 SplmtryData,omitempty"`
}

type CorporateActionNarrative19 struct {
	AddtlTxt     []RestrictedFINXMax350Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.002.08 AddtlTxt,omitempty"`
	PtyCtctNrrtv []RestrictedFINXMax350Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.002.08 PtyCtctNrrtv,omitempty"`
}

// May be one of ABST, AMGT, BSPL, BUYA, CASE, CASH, CEXC, CONN, CONY, CTEN, EXER, LAPS, MKDW, MKUP, MNGT, MPUT, NOAC, NOQU, OFFR, OTHR, OVER, PROX, QINV, SECU, SLLE, SPLI, TAXI, PRUN
type CorporateActionOption10Code string

type CorporateActionOption144 struct {
	OptnNb      OptionNumber1Choice                  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.002.08 OptnNb"`
	OptnTp      CorporateActionOption22Choice        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.002.08 OptnTp"`
	AcctOwnr    PartyIdentification103Choice         `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.002.08 AcctOwnr,omitempty"`
	SfkpgAcct   RestrictedFINXMax35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.002.08 SfkpgAcct,omitempty"`
	CshAcct     CashAccountIdentification6Choice     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.002.08 CshAcct,omitempty"`
	SfkpgPlc    SafekeepingPlaceFormat11Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.002.08 SfkpgPlc,omitempty"`
	FinInstrmId SecurityIdentification20             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.002.08 FinInstrmId,omitempty"`
	TtlElgblBal SignedQuantityFormat8                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.002.08 TtlElgblBal,omitempty"`
	InstdBal    SignedQuantityFormat8                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.002.08 InstdBal,omitempty"`
	UinstdBal   SignedQuantityFormat8                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.002.08 UinstdBal,omitempty"`
	StsQty      Quantity10Choice                     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.002.08 StsQty,omitempty"`
	StsCshAmt   RestrictedFINActiveCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.002.08 StsCshAmt,omitempty"`
}

type CorporateActionOption22Choice struct {
	Cd    CorporateActionOption10Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.002.08 Cd"`
	Prtry GenericIdentification47     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.002.08 Prtry"`
}

// Must match the pattern [A-Z]{2,2}
type CountryCode string

type Document struct {
	CorpActnInstrCxlReqStsAdvc CorporateActionInstructionCancellationRequestStatusAdvice002V08 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.002.08 CorpActnInstrCxlReqStsAdvc"`
}

type DocumentIdentification17 struct {
	Id RestrictedFINXMax16Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.002.08 Id"`
}

type DocumentIdentification34 struct {
	Id    DocumentIdentification4Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.002.08 Id"`
	DocNb DocumentNumber6Choice         `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.002.08 DocNb,omitempty"`
}

type DocumentIdentification4Choice struct {
	AcctSvcrDocId RestrictedFINXMax16Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.002.08 AcctSvcrDocId"`
	AcctOwnrDocId RestrictedFINXMax16Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.002.08 AcctOwnrDocId"`
}

type DocumentNumber6Choice struct {
	ShrtNb  Exact3NumericText                 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.002.08 ShrtNb"`
	LngNb   ISO20022MessageIdentificationText `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.002.08 LngNb"`
	PrtryNb GenericIdentification86           `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.002.08 PrtryNb"`
}

// Must match the pattern [0-9]{3}
type Exact3NumericText string

// Must match the pattern [a-zA-Z0-9]{4}
type Exact4AlphaNumericText string

// Must be at least 1 items long
type ExternalFinancialInstrumentIdentificationType1Code string

type FinancialInstrumentQuantity15Choice struct {
	Unit     float64 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.002.08 Unit"`
	FaceAmt  float64 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.002.08 FaceAmt"`
	AmtsdVal float64 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.002.08 AmtsdVal"`
}

type GenericIdentification47 struct {
	Id      Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.002.08 Id"`
	Issr    Max4AlphaNumericText   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.002.08 Issr"`
	SchmeNm Max4AlphaNumericText   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.002.08 SchmeNm,omitempty"`
}

type GenericIdentification84 struct {
	Id      RestrictedFINXMax34Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.002.08 Id"`
	Issr    Max4AlphaNumericText    `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.002.08 Issr"`
	SchmeNm Max4AlphaNumericText    `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.002.08 SchmeNm,omitempty"`
}

type GenericIdentification85 struct {
	Tp GenericIdentification47 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.002.08 Tp"`
	Id RestrictedFINXMax30Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.002.08 Id,omitempty"`
}

type GenericIdentification86 struct {
	Id      RestrictedFINXMax30Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.002.08 Id"`
	Issr    Max4AlphaNumericText    `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.002.08 Issr"`
	SchmeNm Max4AlphaNumericText    `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.002.08 SchmeNm,omitempty"`
}

// Must match the pattern [A-Z]{2,2}[0-9]{2,2}[a-zA-Z0-9]{1,30}
type IBAN2007Identifier string

// Must match the pattern [A-Z]{2,2}[A-Z0-9]{9,9}[0-9]{1,1}
type ISINOct2015Identifier string

// Must match the pattern [a-z]{4}\.[0-9]{3}\.[0-9]{3}\.[0-9]{2}
type ISO20022MessageIdentificationText string

type IdentificationSource4Choice struct {
	Cd    ExternalFinancialInstrumentIdentificationType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.002.08 Cd"`
	Prtry RestrictedFINExact2Text                            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.002.08 Prtry"`
}

type InstructionCancellationRequestStatus12Choice struct {
	CxlCmpltd CancelledStatus14Choice          `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.002.08 CxlCmpltd"`
	Accptd    NoSpecifiedReason1               `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.002.08 Accptd"`
	Rjctd     RejectedStatus20Choice           `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.002.08 Rjctd"`
	PdgCxl    PendingCancellationStatus8Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.002.08 PdgCxl"`
	PrtrySts  ProprietaryStatusAndReason7      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.002.08 PrtrySts"`
}

// Must be at least 1 items long
type Max16Text string

// Must be at least 1 items long
type Max350Text string

// Must match the pattern [a-zA-Z0-9]{1,4}
type Max4AlphaNumericText string

// May be one of NORE
type NoReasonCode string

type NoSpecifiedReason1 struct {
	NoSpcfdRsn NoReasonCode `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.002.08 NoSpcfdRsn"`
}

type OptionNumber1Choice struct {
	Nb Exact3NumericText `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.002.08 Nb"`
	Cd OptionNumber1Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.002.08 Cd"`
}

// May be one of UNSO
type OptionNumber1Code string

type OriginalAndCurrentQuantities4 struct {
	FaceAmt  float64 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.002.08 FaceAmt"`
	AmtsdVal float64 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.002.08 AmtsdVal"`
}

type OtherIdentification2 struct {
	Id  RestrictedFINXMax31Text     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.002.08 Id"`
	Sfx Max16Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.002.08 Sfx,omitempty"`
	Tp  IdentificationSource4Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.002.08 Tp"`
}

type PartyIdentification103Choice struct {
	AnyBIC  AnyBICIdentifier        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.002.08 AnyBIC"`
	PrtryId GenericIdentification84 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.002.08 PrtryId"`
}

// May be one of ADEA, DQUA, DQCS, LATE, OTHR
type PendingCancellationReason5Code string

type PendingCancellationReason6Choice struct {
	Cd    PendingCancellationReason5Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.002.08 Cd"`
	Prtry GenericIdentification47        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.002.08 Prtry"`
}

type PendingCancellationStatus8Choice struct {
	NotSpcfdRsn NoReasonCode                       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.002.08 NotSpcfdRsn"`
	Rsn         []PendingCancellationStatusReason8 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.002.08 Rsn"`
}

type PendingCancellationStatusReason8 struct {
	RsnCd       PendingCancellationReason6Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.002.08 RsnCd"`
	AddtlRsnInf RestrictedFINXMax210Text         `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.002.08 AddtlRsnInf,omitempty"`
}

type ProprietaryQuantity9 struct {
	Qty     float64                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.002.08 Qty"`
	QtyTp   Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.002.08 QtyTp"`
	Issr    Max4AlphaNumericText   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.002.08 Issr"`
	SchmeNm Max4AlphaNumericText   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.002.08 SchmeNm,omitempty"`
}

type ProprietaryReason5 struct {
	Rsn         GenericIdentification47  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.002.08 Rsn,omitempty"`
	AddtlRsnInf RestrictedFINXMax210Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.002.08 AddtlRsnInf,omitempty"`
}

type ProprietaryStatusAndReason7 struct {
	PrtrySts GenericIdentification47 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.002.08 PrtrySts"`
	PrtryRsn []ProprietaryReason5    `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.002.08 PrtryRsn,omitempty"`
}

type Quantity10Choice struct {
	Qty             FinancialInstrumentQuantity15Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.002.08 Qty"`
	OrgnlAndCurFace OriginalAndCurrentQuantities4       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.002.08 OrgnlAndCurFace"`
}

type Quantity21Choice struct {
	Qty      FinancialInstrumentQuantity15Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.002.08 Qty"`
	PrtryQty ProprietaryQuantity9                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.002.08 PrtryQty"`
}

type RejectedReason18Choice struct {
	Cd    RejectionReason45Code   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.002.08 Cd"`
	Prtry GenericIdentification47 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.002.08 Prtry"`
}

type RejectedStatus20Choice struct {
	NoSpcfdRsn NoReasonCode             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.002.08 NoSpcfdRsn"`
	Rsn        []RejectedStatusReason19 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.002.08 Rsn"`
}

type RejectedStatusReason19 struct {
	RsnCd       RejectedReason18Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.002.08 RsnCd"`
	AddtlRsnInf RestrictedFINXMax210Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.002.08 AddtlRsnInf,omitempty"`
}

// May be one of ADEA, ULNK, OTHR, DCAN, DPRG, INIR, LATE
type RejectionReason45Code string

type RestrictedFINActiveCurrencyAndAmount struct {
	Value float64            `xml:",chardata"`
	Ccy   ActiveCurrencyCode `xml:"Ccy,attr"`
}

// Must match the pattern XX|TS
type RestrictedFINExact2Text string

// Must match the pattern [0-9a-zA-Z/\-\?:\(\)\.,'\+ ]{1,34}
type RestrictedFINX2Max34Text string

// Must match the pattern [0-9a-zA-Z/\-\?:\(\)\.\n\r,'\+ ]{1,140}
type RestrictedFINXMax140Text string

// Must match the pattern ([0-9a-zA-Z\-\?:\(\)\.,'\+ ]([0-9a-zA-Z\-\?:\(\)\.,'\+ ]*(/[0-9a-zA-Z\-\?:\(\)\.,'\+ ])?)*)
type RestrictedFINXMax16Text string

// Must match the pattern [0-9a-zA-Z/\-\?:\(\)\.\n\r,'\+ ]{1,210}
type RestrictedFINXMax210Text string

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

type SafekeepingPlaceFormat11Choice struct {
	Id      SafekeepingPlaceTypeAndText9             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.002.08 Id"`
	Ctry    CountryCode                              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.002.08 Ctry"`
	TpAndId SafekeepingPlaceTypeAndAnyBICIdentifier1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.002.08 TpAndId"`
	Prtry   GenericIdentification85                  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.002.08 Prtry"`
}

type SafekeepingPlaceTypeAndAnyBICIdentifier1 struct {
	SfkpgPlcTp SafekeepingPlace1Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.002.08 SfkpgPlcTp"`
	Id         AnyBICIdentifier      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.002.08 Id"`
}

type SafekeepingPlaceTypeAndText9 struct {
	SfkpgPlcTp SafekeepingPlace2Code   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.002.08 SfkpgPlcTp"`
	Id         RestrictedFINXMax30Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.002.08 Id,omitempty"`
}

type SecurityIdentification20 struct {
	ISIN   ISINOct2015Identifier    `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.002.08 ISIN,omitempty"`
	OthrId []OtherIdentification2   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.002.08 OthrId,omitempty"`
	Desc   RestrictedFINXMax140Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.002.08 Desc,omitempty"`
}

// May be one of SHOR, LONG
type ShortLong1Code string

type SignedQuantityFormat8 struct {
	ShrtLngPos ShortLong1Code   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.002.08 ShrtLngPos"`
	QtyChc     Quantity21Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.002.08 QtyChc"`
}

type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.002.08 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.002.08 Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
}