// Code generated by main. DO NOT EDIT.

package seev_040_001_02

type AccountIdentification17 struct {
	SfkpgAcct Max35Text                     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.02 SfkpgAcct"`
	AcctOwnr  PartyIdentification36Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.02 AcctOwnr,omitempty"`
	SfkpgPlc  SafekeepingPlaceFormat2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.02 SfkpgPlc,omitempty"`
}

// Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type AnyBICIdentifier string

type CorporateActionEventType3Choice struct {
	Cd    CorporateActionEventType6Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.02 Cd"`
	Prtry GenericIdentification20       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.02 Prtry"`
}

// May be one of ACTV, ATTI, BIDS, BONU, BPUT, BRUP, CAPG, CAPI, CERT, CHAN, CLSA, CONS, CONV, COOP, DECR, DETI, DFLT, DLST, DRAW, DRIP, DSCL, DTCH, DVCA, DVOP, DVSC, DVSE, EXOF, EXRI, EXTM, EXWA, CAPD, INCR, INTR, LIQU, MCAL, MRGR, ODLT, OTHR, PARI, PCAL, PDEF, PINK, PLAC, PPMT, PRED, PRII, PRIO, REDM, REDO, REMK, RHDI, RHTS, SHPR, SMAL, SOFF, SPLF, SPLR, SUSP, TEND, TREC, WRTH, WTRC, CREV
type CorporateActionEventType6Code string

type CorporateActionGeneralInformation26 struct {
	CorpActnEvtId      Max35Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.02 CorpActnEvtId"`
	OffclCorpActnEvtId Max35Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.02 OffclCorpActnEvtId,omitempty"`
	EvtTp              CorporateActionEventType3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.02 EvtTp"`
	UndrlygSctyId      SecurityIdentification14        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.02 UndrlygSctyId,omitempty"`
}

type CorporateActionInstructionCancellationRequestV02 struct {
	ChngInstrInd   bool                                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.02 ChngInstrInd,omitempty"`
	InstrId        DocumentIdentification15            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.02 InstrId"`
	CorpActnGnlInf CorporateActionGeneralInformation26 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.02 CorpActnGnlInf"`
	AcctDtls       AccountIdentification17             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.02 AcctDtls"`
	CorpActnInstr  CorporateActionOption21             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.02 CorpActnInstr"`
	SplmtryData    []SupplementaryData1                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.02 SplmtryData,omitempty"`
}

type CorporateActionOption21 struct {
	OptnNb          OptionNumber1Choice                  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.02 OptnNb"`
	OptnTp          CorporateActionOption4Choice         `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.02 OptnTp"`
	InstdOrQtyToRcv InstructedOrQuantityToReceive1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.02 InstdOrQtyToRcv"`
}

type CorporateActionOption4Choice struct {
	Cd    CorporateActionOption4Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.02 Cd"`
	Prtry GenericIdentification20    `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.02 Prtry"`
}

// May be one of ABST, AMGT, BSPL, BUYA, CASE, CASH, CERT, CEXC, CONN, CONY, CTEN, EXER, LAPS, MKDW, MKUP, MNGT, MPUT, NOAC, NOQU, OFFR, OTHR, OVER, PROX, QINV, SECU, SLLE, SPLI, TAXI
type CorporateActionOption4Code string

// Must match the pattern [A-Z]{2,2}
type CountryCode string

type Document struct {
	CorpActnInstrCxlReq CorporateActionInstructionCancellationRequestV02 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.02 CorpActnInstrCxlReq"`
}

type DocumentIdentification15 struct {
	Id    Max35Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.02 Id"`
	LkgTp ProcessingPosition1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.02 LkgTp,omitempty"`
}

// Must match the pattern [0-9]{3}
type Exact3NumericText string

// Must match the pattern [a-zA-Z0-9]{4}
type Exact4AlphaNumericText string

// Must be at least 1 items long
type ExternalFinancialInstrumentIdentificationType1Code string

type FinancialInstrumentQuantity1Choice struct {
	Unit     float64 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.02 Unit"`
	FaceAmt  float64 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.02 FaceAmt"`
	AmtsdVal float64 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.02 AmtsdVal"`
}

type GenericIdentification19 struct {
	Id      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.02 Id"`
	Issr    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.02 Issr"`
	SchmeNm Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.02 SchmeNm,omitempty"`
}

type GenericIdentification20 struct {
	Id      Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.02 Id"`
	Issr    Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.02 Issr"`
	SchmeNm Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.02 SchmeNm,omitempty"`
}

type GenericIdentification21 struct {
	Tp GenericIdentification20 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.02 Tp"`
	Id Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.02 Id,omitempty"`
}

// Must match the pattern [A-Z0-9]{12,12}
type ISINIdentifier string

type IdentificationSource3Choice struct {
	Cd    ExternalFinancialInstrumentIdentificationType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.02 Cd"`
	Prtry Max35Text                                          `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.02 Prtry"`
}

type InstructedOrQuantityToReceive1Choice struct {
	InstdQty Quantity5Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.02 InstdQty"`
	QtyToRcv Quantity5Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.02 QtyToRcv"`
}

// Must be at least 1 items long
type Max140Text string

// Must be at least 1 items long
type Max16Text string

// Must be at least 1 items long
type Max350Text string

// Must be at least 1 items long
type Max35Text string

type OptionNumber1Choice struct {
	Nb Exact3NumericText `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.02 Nb"`
	Cd OptionNumber1Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.02 Cd"`
}

// May be one of UNSO
type OptionNumber1Code string

type OriginalAndCurrentQuantities1 struct {
	FaceAmt  float64 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.02 FaceAmt"`
	AmtsdVal float64 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.02 AmtsdVal"`
}

type OtherIdentification1 struct {
	Id  Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.02 Id"`
	Sfx Max16Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.02 Sfx,omitempty"`
	Tp  IdentificationSource3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.02 Tp"`
}

type PartyIdentification36Choice struct {
	AnyBIC  AnyBICIdentifier        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.02 AnyBIC"`
	PrtryId GenericIdentification19 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.02 PrtryId"`
}

type ProcessingPosition1Choice struct {
	Cd    ProcessingPosition3Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.02 Cd"`
	Prtry GenericIdentification20 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.02 Prtry"`
}

// May be one of AFTE, WITH, BEFO, INFO
type ProcessingPosition3Code string

// May be one of QALL
type Quantity1Code string

type Quantity5Choice struct {
	Cd                 Quantity1Code                      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.02 Cd"`
	OrgnlAndCurFaceAmt OriginalAndCurrentQuantities1      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.02 OrgnlAndCurFaceAmt"`
	Qty                FinancialInstrumentQuantity1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.02 Qty"`
}

// May be one of CUST, ICSD, NCSD, SHHE
type SafekeepingPlace1Code string

// May be one of SHHE, ALLP
type SafekeepingPlace2Code string

type SafekeepingPlaceFormat2Choice struct {
	Id      SafekeepingPlaceTypeAndText2             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.02 Id"`
	Ctry    CountryCode                              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.02 Ctry"`
	TpAndId SafekeepingPlaceTypeAndAnyBICIdentifier1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.02 TpAndId"`
	Prtry   GenericIdentification21                  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.02 Prtry"`
}

type SafekeepingPlaceTypeAndAnyBICIdentifier1 struct {
	SfkpgPlcTp SafekeepingPlace1Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.02 SfkpgPlcTp"`
	Id         AnyBICIdentifier      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.02 Id"`
}

type SafekeepingPlaceTypeAndText2 struct {
	SfkpgPlcTp SafekeepingPlace2Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.02 SfkpgPlcTp"`
	Id         Max35Text             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.02 Id,omitempty"`
}

type SecurityIdentification14 struct {
	ISIN   ISINIdentifier         `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.02 ISIN,omitempty"`
	OthrId []OtherIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.02 OthrId,omitempty"`
	Desc   Max140Text             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.02 Desc,omitempty"`
}

type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.02 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.02 Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
}
