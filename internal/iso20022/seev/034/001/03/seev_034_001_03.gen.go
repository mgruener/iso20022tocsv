// Code generated by main. DO NOT EDIT.

package seev_034_001_03

type AcceptedReason1Choice struct {
	Cd    AcknowledgementReason4Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.03 Cd"`
	Prtry GenericIdentification20    `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.03 Prtry"`
}

type AcceptedStatus1Choice struct {
	NoSpcfdRsn NoReasonCode            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.03 NoSpcfdRsn"`
	Rsn        []AcceptedStatusReason1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.03 Rsn"`
}

type AcceptedStatusReason1 struct {
	RsnCd       AcceptedReason1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.03 RsnCd"`
	AddtlRsnInf Max210Text            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.03 AddtlRsnInf,omitempty"`
}

// May be one of ADEA, LATE, NSTP, OTHR
type AcknowledgementReason4Code string

// Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type AnyBICIdentifier string

type CancelledReason3Choice struct {
	Cd    CancelledStatusReason11Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.03 Cd"`
	Prtry GenericIdentification20     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.03 Prtry"`
}

type CancelledStatus3Choice struct {
	NoSpcfdRsn NoReasonCode             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.03 NoSpcfdRsn"`
	Rsn        []CancelledStatusReason6 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.03 Rsn"`
}

// May be one of CANI, CANO, CANS, CSUB, OTHR
type CancelledStatusReason11Code string

type CancelledStatusReason6 struct {
	RsnCd       CancelledReason3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.03 RsnCd"`
	AddtlRsnInf Max210Text             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.03 AddtlRsnInf,omitempty"`
}

type CashAccountIdentification5Choice struct {
	IBAN  IBAN2007Identifier `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.03 IBAN"`
	Prtry Max34Text          `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.03 Prtry"`
}

type CorporateActionEventType7Choice struct {
	Cd    CorporateActionEventType8Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.03 Cd"`
	Prtry GenericIdentification20       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.03 Prtry"`
}

// May be one of ACTV, ATTI, BIDS, BONU, BPUT, BRUP, CAPG, CAPI, CERT, CHAN, CLSA, CONS, CONV, COOP, DECR, DETI, DFLT, DLST, DRAW, DRIP, DSCL, DTCH, DVCA, DVOP, DVSC, DVSE, EXOF, EXRI, EXTM, EXWA, CAPD, INCR, INTR, LIQU, MCAL, MRGR, ODLT, OTHR, PARI, PCAL, PDEF, PINK, PLAC, PPMT, PRED, PRII, PRIO, REDM, REDO, REMK, RHDI, RHTS, SHPR, SMAL, SOFF, SPLF, SPLR, SUSP, TEND, TREC, WRTH, WTRC, CREV, DRCA
type CorporateActionEventType8Code string

type CorporateActionGeneralInformation34 struct {
	CorpActnEvtId      Max35Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.03 CorpActnEvtId"`
	OffclCorpActnEvtId Max35Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.03 OffclCorpActnEvtId,omitempty"`
	ClssActnNb         Max35Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.03 ClssActnNb,omitempty"`
	EvtTp              CorporateActionEventType7Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.03 EvtTp"`
}

type CorporateActionInstructionStatusAdviceV03 struct {
	InstrId        DocumentIdentification9              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.03 InstrId,omitempty"`
	OthrDocId      []DocumentIdentification14           `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.03 OthrDocId,omitempty"`
	CorpActnGnlInf CorporateActionGeneralInformation34  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.03 CorpActnGnlInf"`
	InstrPrcgSts   []InstructionProcessingStatus7Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.03 InstrPrcgSts"`
	CorpActnInstr  CorporateActionOption41              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.03 CorpActnInstr,omitempty"`
	AddtlInf       CorporateActionNarrative10           `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.03 AddtlInf,omitempty"`
	SplmtryData    []SupplementaryData1                 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.03 SplmtryData,omitempty"`
}

type CorporateActionNarrative10 struct {
	AddtlTxt     []Max350Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.03 AddtlTxt,omitempty"`
	PtyCtctNrrtv []Max350Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.03 PtyCtctNrrtv,omitempty"`
}

// May be one of ABST, AMGT, BSPL, BUYA, CASE, CASH, CEXC, CONN, CONY, CTEN, EXER, LAPS, MKDW, MKUP, MNGT, MPUT, NOAC, NOQU, OFFR, OTHR, OVER, PROX, QINV, SECU, SLLE, SPLI, TAXI, PRUN
type CorporateActionOption10Code string

type CorporateActionOption13Choice struct {
	Cd    CorporateActionOption10Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.03 Cd"`
	Prtry GenericIdentification20     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.03 Prtry"`
}

type CorporateActionOption41 struct {
	OptnNb           OptionNumber1Choice              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.03 OptnNb"`
	OptnTp           CorporateActionOption13Choice    `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.03 OptnTp"`
	AcctOwnr         PartyIdentification36Choice      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.03 AcctOwnr,omitempty"`
	SfkpgAcct        Max35Text                        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.03 SfkpgAcct,omitempty"`
	CshAcct          CashAccountIdentification5Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.03 CshAcct,omitempty"`
	SfkpgPlc         SafekeepingPlaceFormat2Choice    `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.03 SfkpgPlc,omitempty"`
	FinInstrmId      SecurityIdentification14         `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.03 FinInstrmId,omitempty"`
	TtlElgblBal      SignedQuantityFormat1            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.03 TtlElgblBal,omitempty"`
	InstdBal         SignedQuantityFormat1            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.03 InstdBal,omitempty"`
	UinstdBal        SignedQuantityFormat1            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.03 UinstdBal,omitempty"`
	StsQtyOrQtyToRcv StatusOrQuantityToReceive1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.03 StsQtyOrQtyToRcv,omitempty"`
}

// Must match the pattern [A-Z]{2,2}
type CountryCode string

type Document struct {
	CorpActnInstrStsAdvc CorporateActionInstructionStatusAdviceV03 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.03 CorpActnInstrStsAdvc"`
}

type DocumentIdentification14 struct {
	Id    DocumentIdentification1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.03 Id"`
	DocNb DocumentNumber1Choice         `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.03 DocNb,omitempty"`
}

type DocumentIdentification1Choice struct {
	AcctSvcrDocId Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.03 AcctSvcrDocId"`
	AcctOwnrDocId Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.03 AcctOwnrDocId"`
}

type DocumentIdentification9 struct {
	Id Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.03 Id"`
}

type DocumentNumber1Choice struct {
	ShrtNb  Exact3NumericText                 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.03 ShrtNb"`
	LngNb   ISO20022MessageIdentificationText `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.03 LngNb"`
	PrtryNb GenericIdentification19           `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.03 PrtryNb"`
}

// Must match the pattern [0-9]{3}
type Exact3NumericText string

// Must match the pattern [a-zA-Z0-9]{4}
type Exact4AlphaNumericText string

// Must be at least 1 items long
type ExternalFinancialInstrumentIdentificationType1Code string

type FinancialInstrumentQuantity1Choice struct {
	Unit     float64 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.03 Unit"`
	FaceAmt  float64 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.03 FaceAmt"`
	AmtsdVal float64 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.03 AmtsdVal"`
}

type GenericIdentification19 struct {
	Id      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.03 Id"`
	Issr    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.03 Issr"`
	SchmeNm Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.03 SchmeNm,omitempty"`
}

type GenericIdentification20 struct {
	Id      Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.03 Id"`
	Issr    Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.03 Issr"`
	SchmeNm Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.03 SchmeNm,omitempty"`
}

type GenericIdentification21 struct {
	Tp GenericIdentification20 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.03 Tp"`
	Id Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.03 Id,omitempty"`
}

// Must match the pattern [A-Z]{2,2}[0-9]{2,2}[a-zA-Z0-9]{1,30}
type IBAN2007Identifier string

// Must match the pattern [A-Z0-9]{12,12}
type ISINIdentifier string

// Must match the pattern [a-z]{4}\.[0-9]{3}\.[0-9]{3}\.[0-9]{2}
type ISO20022MessageIdentificationText string

type IdentificationSource3Choice struct {
	Cd    ExternalFinancialInstrumentIdentificationType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.03 Cd"`
	Prtry Max35Text                                          `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.03 Prtry"`
}

type InstructionProcessingStatus7Choice struct {
	Canc     CancelledStatus3Choice      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.03 Canc"`
	Accptd   AcceptedStatus1Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.03 Accptd"`
	Rjctd    RejectedStatus1Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.03 Rjctd"`
	Pdg      PendingStatus1Choice        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.03 Pdg"`
	DfltActn NoSpecifiedReason1          `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.03 DfltActn"`
	StgInstr NoSpecifiedReason1          `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.03 StgInstr"`
	PrtrySts ProprietaryStatusAndReason1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.03 PrtrySts"`
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
	NoSpcfdRsn NoReasonCode `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.03 NoSpcfdRsn"`
}

type OptionNumber1Choice struct {
	Nb Exact3NumericText `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.03 Nb"`
	Cd OptionNumber1Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.03 Cd"`
}

// May be one of UNSO
type OptionNumber1Code string

type OriginalAndCurrentQuantities1 struct {
	FaceAmt  float64 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.03 FaceAmt"`
	AmtsdVal float64 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.03 AmtsdVal"`
}

type OtherIdentification1 struct {
	Id  Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.03 Id"`
	Sfx Max16Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.03 Sfx,omitempty"`
	Tp  IdentificationSource3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.03 Tp"`
}

type PartyIdentification36Choice struct {
	AnyBIC  AnyBICIdentifier        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.03 AnyBIC"`
	PrtryId GenericIdentification19 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.03 PrtryId"`
}

// May be one of ADEA, AUTH, DQUA, LACK, LATE, MCER, MONY, OTHR, NPAY, NSEC, PENR, VLDA, CERT
type PendingReason5Code string

type PendingReason6Choice struct {
	Cd    PendingReason5Code      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.03 Cd"`
	Prtry GenericIdentification20 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.03 Prtry"`
}

type PendingStatus1Choice struct {
	NoSpcfdRsn NoReasonCode           `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.03 NoSpcfdRsn"`
	Rsn        []PendingStatusReason1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.03 Rsn"`
}

type PendingStatusReason1 struct {
	RsnCd       PendingReason6Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.03 RsnCd"`
	AddtlRsnInf Max210Text           `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.03 AddtlRsnInf,omitempty"`
}

type ProprietaryQuantity2 struct {
	Qty     float64                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.03 Qty"`
	QtyTp   Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.03 QtyTp"`
	Issr    Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.03 Issr"`
	SchmeNm Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.03 SchmeNm,omitempty"`
}

type ProprietaryReason1 struct {
	Rsn         GenericIdentification20 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.03 Rsn,omitempty"`
	AddtlRsnInf Max210Text              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.03 AddtlRsnInf,omitempty"`
}

type ProprietaryStatusAndReason1 struct {
	PrtrySts GenericIdentification20 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.03 PrtrySts"`
	PrtryRsn []ProprietaryReason1    `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.03 PrtryRsn,omitempty"`
}

type Quantity2Choice struct {
	Qty      FinancialInstrumentQuantity1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.03 Qty"`
	PrtryQty ProprietaryQuantity2               `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.03 PrtryQty"`
}

type Quantity6Choice struct {
	Qty             FinancialInstrumentQuantity1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.03 Qty"`
	OrgnlAndCurFace OriginalAndCurrentQuantities1      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.03 OrgnlAndCurFace"`
}

type RejectedReason1Choice struct {
	Cd    RejectionReason17Code   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.03 Cd"`
	Prtry GenericIdentification20 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.03 Prtry"`
}

type RejectedStatus1Choice struct {
	NoSpcfdRsn NoReasonCode            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.03 NoSpcfdRsn"`
	Rsn        []RejectedStatusReason8 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.03 Rsn"`
}

type RejectedStatusReason8 struct {
	RsnCd       RejectedReason1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.03 RsnCd"`
	AddtlRsnInf Max210Text            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.03 AddtlRsnInf,omitempty"`
}

// May be one of ADEA, CANC, DCAN, DPRG, DQUA, DSEC, EVNM, INIR, INTV, INVA, LACK, LATE, OTHR, NMTY, OPNM, OPTY, REFT, SAFE, ULNK, CERT
type RejectionReason17Code string

// May be one of CUST, ICSD, NCSD, SHHE
type SafekeepingPlace1Code string

// May be one of SHHE, ALLP
type SafekeepingPlace2Code string

type SafekeepingPlaceFormat2Choice struct {
	Id      SafekeepingPlaceTypeAndText2             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.03 Id"`
	Ctry    CountryCode                              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.03 Ctry"`
	TpAndId SafekeepingPlaceTypeAndAnyBICIdentifier1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.03 TpAndId"`
	Prtry   GenericIdentification21                  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.03 Prtry"`
}

type SafekeepingPlaceTypeAndAnyBICIdentifier1 struct {
	SfkpgPlcTp SafekeepingPlace1Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.03 SfkpgPlcTp"`
	Id         AnyBICIdentifier      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.03 Id"`
}

type SafekeepingPlaceTypeAndText2 struct {
	SfkpgPlcTp SafekeepingPlace2Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.03 SfkpgPlcTp"`
	Id         Max35Text             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.03 Id,omitempty"`
}

type SecurityIdentification14 struct {
	ISIN   ISINIdentifier         `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.03 ISIN,omitempty"`
	OthrId []OtherIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.03 OthrId,omitempty"`
	Desc   Max140Text             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.03 Desc,omitempty"`
}

// May be one of SHOR, LONG
type ShortLong1Code string

type SignedQuantityFormat1 struct {
	ShrtLngPos ShortLong1Code  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.03 ShrtLngPos"`
	QtyChc     Quantity2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.03 QtyChc"`
}

type StatusOrQuantityToReceive1Choice struct {
	StsQty   Quantity6Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.03 StsQty"`
	QtyToRcv Quantity6Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.03 QtyToRcv"`
}

type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.03 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.03 Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
}
