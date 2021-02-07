// Code generated by main. DO NOT EDIT.

package seev_041_001_07

// Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type AnyBICIdentifier string

type CancelledReason9Choice struct {
	Cd    CancelledStatusReason5Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.07 Cd"`
	Prtry GenericIdentification30    `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.07 Prtry"`
}

type CancelledStatus11Choice struct {
	NoSpcfdRsn NoReasonCode              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.07 NoSpcfdRsn"`
	Rsn        []CancelledStatusReason12 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.07 Rsn"`
}

type CancelledStatusReason12 struct {
	RsnCd       CancelledReason9Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.07 RsnCd"`
	AddtlRsnInf Max210Text             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.07 AddtlRsnInf,omitempty"`
}

// May be one of CANI, OTHR
type CancelledStatusReason5Code string

type CashAccountIdentification5Choice struct {
	IBAN  IBAN2007Identifier `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.07 IBAN"`
	Prtry Max34Text          `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.07 Prtry"`
}

// May be one of ACTV, ATTI, BRUP, DFLT, BONU, EXRI, CAPD, CAPG, CAPI, DRCA, DVCA, CHAN, COOP, CLSA, CONS, CONV, CREV, DECR, DETI, DSCL, DVOP, DRIP, DRAW, DTCH, EXOF, REDM, MCAL, INCR, PPMT, INTR, PRII, RHDI, LIQU, EXTM, MRGR, NOOF, CERT, ODLT, OTHR, PARI, PCAL, PRED, PINK, PLAC, PDEF, PRIO, BPUT, REDO, REMK, BIDS, SPLR, RHTS, DVSC, SHPR, SMAL, SOFF, DVSE, SPLF, TREC, TEND, DLST, SUSP, EXWA, WTRC, WRTH
type CorporateActionEventType20Code string

type CorporateActionEventType52Choice struct {
	Cd    CorporateActionEventType20Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.07 Cd"`
	Prtry GenericIdentification30        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.07 Prtry"`
}

type CorporateActionGeneralInformation109 struct {
	CorpActnEvtId      Max35Text                        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.07 CorpActnEvtId"`
	OffclCorpActnEvtId Max35Text                        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.07 OffclCorpActnEvtId,omitempty"`
	ClssActnNb         Max35Text                        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.07 ClssActnNb,omitempty"`
	EvtTp              CorporateActionEventType52Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.07 EvtTp"`
}

type CorporateActionInstructionCancellationRequestStatusAdviceV07 struct {
	InstrCxlReqId  DocumentIdentification9                        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.07 InstrCxlReqId,omitempty"`
	OthrDocId      []DocumentIdentification33                     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.07 OthrDocId,omitempty"`
	CorpActnGnlInf CorporateActionGeneralInformation109           `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.07 CorpActnGnlInf"`
	InstrCxlReqSts []InstructionCancellationRequestStatus11Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.07 InstrCxlReqSts"`
	CorpActnInstr  CorporateActionOption116                       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.07 CorpActnInstr,omitempty"`
	AddtlInf       CorporateActionNarrative10                     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.07 AddtlInf,omitempty"`
	SplmtryData    []SupplementaryData1                           `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.07 SplmtryData,omitempty"`
}

type CorporateActionNarrative10 struct {
	AddtlTxt     []Max350Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.07 AddtlTxt,omitempty"`
	PtyCtctNrrtv []Max350Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.07 PtyCtctNrrtv,omitempty"`
}

// May be one of ABST, AMGT, BSPL, BUYA, CASE, CASH, CEXC, CONN, CONY, CTEN, EXER, LAPS, MKDW, MKUP, MNGT, MPUT, NOAC, NOQU, OFFR, OTHR, OVER, PROX, QINV, SECU, SLLE, SPLI, TAXI, PRUN
type CorporateActionOption10Code string

type CorporateActionOption116 struct {
	OptnNb      OptionNumber1Choice              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.07 OptnNb"`
	OptnTp      CorporateActionOption21Choice    `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.07 OptnTp"`
	AcctOwnr    PartyIdentification92Choice      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.07 AcctOwnr,omitempty"`
	SfkpgAcct   Max35Text                        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.07 SfkpgAcct,omitempty"`
	CshAcct     CashAccountIdentification5Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.07 CshAcct,omitempty"`
	SfkpgPlc    SafekeepingPlaceFormat8Choice    `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.07 SfkpgPlc,omitempty"`
	FinInstrmId SecurityIdentification19         `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.07 FinInstrmId,omitempty"`
	TtlElgblBal SignedQuantityFormat7            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.07 TtlElgblBal,omitempty"`
	InstdBal    SignedQuantityFormat7            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.07 InstdBal,omitempty"`
	UinstdBal   SignedQuantityFormat7            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.07 UinstdBal,omitempty"`
	StsQty      Quantity6Choice                  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.07 StsQty"`
}

type CorporateActionOption21Choice struct {
	Cd    CorporateActionOption10Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.07 Cd"`
	Prtry GenericIdentification30     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.07 Prtry"`
}

// Must match the pattern [A-Z]{2,2}
type CountryCode string

type Document struct {
	CorpActnInstrCxlReqStsAdvc CorporateActionInstructionCancellationRequestStatusAdviceV07 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.07 CorpActnInstrCxlReqStsAdvc"`
}

type DocumentIdentification33 struct {
	Id    DocumentIdentification3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.07 Id"`
	DocNb DocumentNumber5Choice         `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.07 DocNb,omitempty"`
}

type DocumentIdentification3Choice struct {
	AcctSvcrDocId Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.07 AcctSvcrDocId"`
	AcctOwnrDocId Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.07 AcctOwnrDocId"`
}

type DocumentIdentification9 struct {
	Id Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.07 Id"`
}

type DocumentNumber5Choice struct {
	ShrtNb  Exact3NumericText                 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.07 ShrtNb"`
	LngNb   ISO20022MessageIdentificationText `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.07 LngNb"`
	PrtryNb GenericIdentification36           `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.07 PrtryNb"`
}

// Must match the pattern [0-9]{3}
type Exact3NumericText string

// Must match the pattern [a-zA-Z0-9]{4}
type Exact4AlphaNumericText string

// Must be at least 1 items long
type ExternalFinancialInstrumentIdentificationType1Code string

type FinancialInstrumentQuantity1Choice struct {
	Unit     float64 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.07 Unit"`
	FaceAmt  float64 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.07 FaceAmt"`
	AmtsdVal float64 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.07 AmtsdVal"`
}

type GenericIdentification30 struct {
	Id      Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.07 Id"`
	Issr    Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.07 Issr"`
	SchmeNm Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.07 SchmeNm,omitempty"`
}

type GenericIdentification36 struct {
	Id      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.07 Id"`
	Issr    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.07 Issr"`
	SchmeNm Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.07 SchmeNm,omitempty"`
}

type GenericIdentification78 struct {
	Tp GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.07 Tp"`
	Id Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.07 Id,omitempty"`
}

// Must match the pattern [A-Z]{2,2}[0-9]{2,2}[a-zA-Z0-9]{1,30}
type IBAN2007Identifier string

// Must match the pattern [A-Z]{2,2}[A-Z0-9]{9,9}[0-9]{1,1}
type ISINOct2015Identifier string

// Must match the pattern [a-z]{4}\.[0-9]{3}\.[0-9]{3}\.[0-9]{2}
type ISO20022MessageIdentificationText string

type IdentificationSource3Choice struct {
	Cd    ExternalFinancialInstrumentIdentificationType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.07 Cd"`
	Prtry Max35Text                                          `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.07 Prtry"`
}

type InstructionCancellationRequestStatus11Choice struct {
	CxlCmpltd CancelledStatus11Choice          `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.07 CxlCmpltd"`
	Accptd    NoSpecifiedReason1               `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.07 Accptd"`
	Rjctd     RejectedStatus18Choice           `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.07 Rjctd"`
	PdgCxl    PendingCancellationStatus7Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.07 PdgCxl"`
	PrtrySts  ProprietaryStatusAndReason6      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.07 PrtrySts"`
}

// Must be at least 1 items long
type Max140Text string

// Must be at least 1 items long
type Max16Text string

// Must be at least 1 items long
type Max210Text string

// Must be at least 1 items long
type Max34Text string

// Must be at least 1 items long
type Max350Text string

// Must be at least 1 items long
type Max35Text string

// May be one of NORE
type NoReasonCode string

type NoSpecifiedReason1 struct {
	NoSpcfdRsn NoReasonCode `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.07 NoSpcfdRsn"`
}

type OptionNumber1Choice struct {
	Nb Exact3NumericText `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.07 Nb"`
	Cd OptionNumber1Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.07 Cd"`
}

// May be one of UNSO
type OptionNumber1Code string

type OriginalAndCurrentQuantities1 struct {
	FaceAmt  float64 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.07 FaceAmt"`
	AmtsdVal float64 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.07 AmtsdVal"`
}

type OtherIdentification1 struct {
	Id  Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.07 Id"`
	Sfx Max16Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.07 Sfx,omitempty"`
	Tp  IdentificationSource3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.07 Tp"`
}

type PartyIdentification92Choice struct {
	AnyBIC  AnyBICIdentifier        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.07 AnyBIC"`
	PrtryId GenericIdentification36 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.07 PrtryId"`
}

type PendingCancellationReason5Choice struct {
	Cd    PendingCancellationReason5Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.07 Cd"`
	Prtry GenericIdentification30        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.07 Prtry"`
}

// May be one of ADEA, DQUA, DQCS, LATE, OTHR
type PendingCancellationReason5Code string

type PendingCancellationStatus7Choice struct {
	NotSpcfdRsn NoReasonCode                       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.07 NotSpcfdRsn"`
	Rsn         []PendingCancellationStatusReason7 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.07 Rsn"`
}

type PendingCancellationStatusReason7 struct {
	RsnCd       PendingCancellationReason5Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.07 RsnCd"`
	AddtlRsnInf Max210Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.07 AddtlRsnInf,omitempty"`
}

type ProprietaryQuantity8 struct {
	Qty     float64                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.07 Qty"`
	QtyTp   Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.07 QtyTp"`
	Issr    Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.07 Issr"`
	SchmeNm Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.07 SchmeNm,omitempty"`
}

type ProprietaryReason4 struct {
	Rsn         GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.07 Rsn,omitempty"`
	AddtlRsnInf Max210Text              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.07 AddtlRsnInf,omitempty"`
}

type ProprietaryStatusAndReason6 struct {
	PrtrySts GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.07 PrtrySts"`
	PrtryRsn []ProprietaryReason4    `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.07 PrtryRsn,omitempty"`
}

type Quantity19Choice struct {
	Qty      FinancialInstrumentQuantity1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.07 Qty"`
	PrtryQty ProprietaryQuantity8               `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.07 PrtryQty"`
}

type Quantity6Choice struct {
	Qty             FinancialInstrumentQuantity1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.07 Qty"`
	OrgnlAndCurFace OriginalAndCurrentQuantities1      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.07 OrgnlAndCurFace"`
}

type RejectedReason14Choice struct {
	Cd    RejectionReason45Code   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.07 Cd"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.07 Prtry"`
}

type RejectedStatus18Choice struct {
	NoSpcfdRsn NoReasonCode             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.07 NoSpcfdRsn"`
	Rsn        []RejectedStatusReason18 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.07 Rsn"`
}

type RejectedStatusReason18 struct {
	RsnCd       RejectedReason14Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.07 RsnCd"`
	AddtlRsnInf Max210Text             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.07 AddtlRsnInf,omitempty"`
}

// May be one of ADEA, ULNK, OTHR, DCAN, DPRG, INIR, LATE
type RejectionReason45Code string

// May be one of CUST, ICSD, NCSD, SHHE
type SafekeepingPlace1Code string

// May be one of SHHE, ALLP
type SafekeepingPlace2Code string

type SafekeepingPlaceFormat8Choice struct {
	Id      SafekeepingPlaceTypeAndText6             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.07 Id"`
	Ctry    CountryCode                              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.07 Ctry"`
	TpAndId SafekeepingPlaceTypeAndAnyBICIdentifier1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.07 TpAndId"`
	Prtry   GenericIdentification78                  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.07 Prtry"`
}

type SafekeepingPlaceTypeAndAnyBICIdentifier1 struct {
	SfkpgPlcTp SafekeepingPlace1Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.07 SfkpgPlcTp"`
	Id         AnyBICIdentifier      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.07 Id"`
}

type SafekeepingPlaceTypeAndText6 struct {
	SfkpgPlcTp SafekeepingPlace2Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.07 SfkpgPlcTp"`
	Id         Max35Text             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.07 Id,omitempty"`
}

type SecurityIdentification19 struct {
	ISIN   ISINOct2015Identifier  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.07 ISIN,omitempty"`
	OthrId []OtherIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.07 OthrId,omitempty"`
	Desc   Max140Text             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.07 Desc,omitempty"`
}

// May be one of SHOR, LONG
type ShortLong1Code string

type SignedQuantityFormat7 struct {
	ShrtLngPos ShortLong1Code   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.07 ShrtLngPos"`
	QtyChc     Quantity19Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.07 QtyChc"`
}

type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.07 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.07 Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
}
