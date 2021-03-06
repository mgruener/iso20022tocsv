// Code generated by main. DO NOT EDIT.

package seev_039_001_04

type AccountIdentification10 struct {
	IdCd SafekeepingAccountIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.04 IdCd"`
}

type AccountIdentification13Choice struct {
	ForAllAccts AccountIdentification10   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.04 ForAllAccts"`
	AcctsList   []AccountIdentification15 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.04 AcctsList"`
}

type AccountIdentification15 struct {
	SfkpgAcct Max35Text                     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.04 SfkpgAcct"`
	AcctOwnr  PartyIdentification36Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.04 AcctOwnr,omitempty"`
	SfkpgPlc  SafekeepingPlaceFormat2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.04 SfkpgPlc,omitempty"`
}

// May be one of ADDR, PBOX, HOME, BIZZ, MLTO, DLVY
type AddressType2Code string

// Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type AnyBICIdentifier string

type CorporateActionCancellation1 struct {
	CxlRsnCd CorporateActionCancellationReason1Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.04 CxlRsnCd"`
	CxlRsn   Max140Text                             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.04 CxlRsn,omitempty"`
	PrcgSts  CorporateActionProcessingStatus1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.04 PrcgSts"`
}

type CorporateActionCancellationAdviceV04 struct {
	CxlAdvcGnlInf  CorporateActionCancellation1        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.04 CxlAdvcGnlInf"`
	CorpActnGnlInf CorporateActionGeneralInformation56 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.04 CorpActnGnlInf"`
	AcctsDtls      AccountIdentification13Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.04 AcctsDtls"`
	IssrAgt        []PartyIdentification46Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.04 IssrAgt,omitempty"`
	PngAgt         []PartyIdentification46Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.04 PngAgt,omitempty"`
	SubPngAgt      []PartyIdentification46Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.04 SubPngAgt,omitempty"`
	Regar          PartyIdentification46Choice         `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.04 Regar,omitempty"`
	RsellngAgt     []PartyIdentification46Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.04 RsellngAgt,omitempty"`
	PhysSctiesAgt  PartyIdentification46Choice         `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.04 PhysSctiesAgt,omitempty"`
	DrpAgt         PartyIdentification46Choice         `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.04 DrpAgt,omitempty"`
	SlctnAgt       []PartyIdentification46Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.04 SlctnAgt,omitempty"`
	InfAgt         PartyIdentification46Choice         `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.04 InfAgt,omitempty"`
	SplmtryData    []SupplementaryData1                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.04 SplmtryData,omitempty"`
}

// May be one of WITH, PROC
type CorporateActionCancellationReason1Code string

type CorporateActionEventStatus1 struct {
	EvtCmpltnsSts EventCompletenessStatus1Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.04 EvtCmpltnsSts"`
	EvtConfSts    EventConfirmationStatus1Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.04 EvtConfSts"`
}

// May be one of ACCU, ACTV, INFO, ATTI, BRUP, DFLT, BONU, EXRI, CAPD, CAPG, CAPI, DRCA, DVCA, CHAN, CLSA, COOP, CONS, CONV, CREV, DECR, DETI, DSCL, DVOP, DRIP, DRAW, DTCH, EXOF, REDM, MCAL, INCR, PPMT, INTR, PRII, RHDI, LIQU, EXTM, MRGR, NOOF, CERT, ODLT, OTHR, PARI, PCAL, PRED, PINK, PLAC, PDEF, PRIO, REDO, REMK, BIDS, SPLR, RHTS, DVSC, SHPR, SMAL, SOFF, DVSE, SPLF, TREC, TEND, DLST, SUSP, EXWA, WTRC, WRTH, BPUT
type CorporateActionEventType12Code string

type CorporateActionEventType13Choice struct {
	Cd    CorporateActionEventType12Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.04 Cd"`
	Prtry GenericIdentification20        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.04 Prtry"`
}

type CorporateActionGeneralInformation56 struct {
	CorpActnEvtId      Max35Text                                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.04 CorpActnEvtId"`
	OffclCorpActnEvtId Max35Text                                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.04 OffclCorpActnEvtId,omitempty"`
	ClssActnNb         Max35Text                                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.04 ClssActnNb,omitempty"`
	EvtTp              CorporateActionEventType13Choice         `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.04 EvtTp"`
	MndtryVlntryEvtTp  CorporateActionMandatoryVoluntary1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.04 MndtryVlntryEvtTp"`
	FinInstrmId        SecurityIdentification14                 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.04 FinInstrmId"`
}

type CorporateActionMandatoryVoluntary1Choice struct {
	Cd    CorporateActionMandatoryVoluntary1Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.04 Cd"`
	Prtry GenericIdentification20                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.04 Prtry"`
}

// May be one of MAND, CHOS, VOLU
type CorporateActionMandatoryVoluntary1Code string

type CorporateActionProcessingStatus1Choice struct {
	EvtSts        CorporateActionEventStatus1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.04 EvtSts"`
	ForInfOnlyInd bool                        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.04 ForInfOnlyInd"`
}

// Must match the pattern [A-Z]{2,2}
type CountryCode string

type Document struct {
	CorpActnCxlAdvc CorporateActionCancellationAdviceV04 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.04 CorpActnCxlAdvc"`
}

// May be one of COMP, INCO
type EventCompletenessStatus1Code string

// May be one of CONF, UCON
type EventConfirmationStatus1Code string

// Must match the pattern [a-zA-Z0-9]{4}
type Exact4AlphaNumericText string

// Must be at least 1 items long
type ExternalFinancialInstrumentIdentificationType1Code string

type GenericIdentification19 struct {
	Id      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.04 Id"`
	Issr    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.04 Issr"`
	SchmeNm Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.04 SchmeNm,omitempty"`
}

type GenericIdentification20 struct {
	Id      Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.04 Id"`
	Issr    Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.04 Issr"`
	SchmeNm Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.04 SchmeNm,omitempty"`
}

type GenericIdentification21 struct {
	Tp GenericIdentification20 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.04 Tp"`
	Id Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.04 Id,omitempty"`
}

// Must match the pattern [A-Z0-9]{12,12}
type ISINIdentifier string

type IdentificationSource3Choice struct {
	Cd    ExternalFinancialInstrumentIdentificationType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.04 Cd"`
	Prtry Max35Text                                          `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.04 Prtry"`
}

// Must be at least 1 items long
type Max140Text string

// Must be at least 1 items long
type Max16Text string

// Must be at least 1 items long
type Max350Text string

// Must be at least 1 items long
type Max35Text string

// Must be at least 1 items long
type Max70Text string

type NameAndAddress5 struct {
	Nm  Max350Text     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.04 Nm"`
	Adr PostalAddress1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.04 Adr,omitempty"`
}

type OtherIdentification1 struct {
	Id  Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.04 Id"`
	Sfx Max16Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.04 Sfx,omitempty"`
	Tp  IdentificationSource3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.04 Tp"`
}

type PartyIdentification36Choice struct {
	AnyBIC  AnyBICIdentifier        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.04 AnyBIC"`
	PrtryId GenericIdentification19 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.04 PrtryId"`
}

type PartyIdentification46Choice struct {
	AnyBIC   AnyBICIdentifier        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.04 AnyBIC"`
	PrtryId  GenericIdentification19 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.04 PrtryId"`
	NmAndAdr NameAndAddress5         `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.04 NmAndAdr"`
}

type PostalAddress1 struct {
	AdrTp       AddressType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.04 AdrTp,omitempty"`
	AdrLine     []Max70Text      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.04 AdrLine,omitempty"`
	StrtNm      Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.04 StrtNm,omitempty"`
	BldgNb      Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.04 BldgNb,omitempty"`
	PstCd       Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.04 PstCd,omitempty"`
	TwnNm       Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.04 TwnNm,omitempty"`
	CtrySubDvsn Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.04 CtrySubDvsn,omitempty"`
	Ctry        CountryCode      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.04 Ctry"`
}

// May be one of GENR
type SafekeepingAccountIdentification1Code string

// May be one of CUST, ICSD, NCSD, SHHE
type SafekeepingPlace1Code string

// May be one of SHHE, ALLP
type SafekeepingPlace2Code string

type SafekeepingPlaceFormat2Choice struct {
	Id      SafekeepingPlaceTypeAndText2             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.04 Id"`
	Ctry    CountryCode                              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.04 Ctry"`
	TpAndId SafekeepingPlaceTypeAndAnyBICIdentifier1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.04 TpAndId"`
	Prtry   GenericIdentification21                  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.04 Prtry"`
}

type SafekeepingPlaceTypeAndAnyBICIdentifier1 struct {
	SfkpgPlcTp SafekeepingPlace1Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.04 SfkpgPlcTp"`
	Id         AnyBICIdentifier      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.04 Id"`
}

type SafekeepingPlaceTypeAndText2 struct {
	SfkpgPlcTp SafekeepingPlace2Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.04 SfkpgPlcTp"`
	Id         Max35Text             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.04 Id,omitempty"`
}

type SecurityIdentification14 struct {
	ISIN   ISINIdentifier         `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.04 ISIN,omitempty"`
	OthrId []OtherIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.04 OthrId,omitempty"`
	Desc   Max140Text             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.04 Desc,omitempty"`
}

type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.04 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.04 Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
}
