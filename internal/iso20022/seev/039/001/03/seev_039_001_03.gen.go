// Code generated by main. DO NOT EDIT.

package seev_039_001_03

type AccountIdentification10 struct {
	IdCd SafekeepingAccountIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.03 IdCd"`
}

type AccountIdentification13Choice struct {
	ForAllAccts AccountIdentification10   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.03 ForAllAccts"`
	AcctsList   []AccountIdentification15 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.03 AcctsList"`
}

type AccountIdentification15 struct {
	SfkpgAcct Max35Text                     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.03 SfkpgAcct"`
	AcctOwnr  PartyIdentification36Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.03 AcctOwnr,omitempty"`
	SfkpgPlc  SafekeepingPlaceFormat2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.03 SfkpgPlc,omitempty"`
}

// May be one of ADDR, PBOX, HOME, BIZZ, MLTO, DLVY
type AddressType2Code string

// Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type AnyBICIdentifier string

type CorporateActionCancellation1 struct {
	CxlRsnCd CorporateActionCancellationReason1Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.03 CxlRsnCd"`
	CxlRsn   Max140Text                             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.03 CxlRsn,omitempty"`
	PrcgSts  CorporateActionProcessingStatus1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.03 PrcgSts"`
}

type CorporateActionCancellationAdviceV03 struct {
	CxlAdvcGnlInf  CorporateActionCancellation1        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.03 CxlAdvcGnlInf"`
	CorpActnGnlInf CorporateActionGeneralInformation35 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.03 CorpActnGnlInf"`
	AcctsDtls      AccountIdentification13Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.03 AcctsDtls"`
	IssrAgt        []PartyIdentification46Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.03 IssrAgt,omitempty"`
	PngAgt         []PartyIdentification46Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.03 PngAgt,omitempty"`
	SubPngAgt      []PartyIdentification46Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.03 SubPngAgt,omitempty"`
	Regar          PartyIdentification46Choice         `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.03 Regar,omitempty"`
	RsellngAgt     []PartyIdentification46Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.03 RsellngAgt,omitempty"`
	PhysSctiesAgt  PartyIdentification46Choice         `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.03 PhysSctiesAgt,omitempty"`
	DrpAgt         PartyIdentification46Choice         `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.03 DrpAgt,omitempty"`
	SlctnAgt       []PartyIdentification46Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.03 SlctnAgt,omitempty"`
	InfAgt         PartyIdentification46Choice         `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.03 InfAgt,omitempty"`
	SplmtryData    []SupplementaryData1                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.03 SplmtryData,omitempty"`
}

// May be one of WITH, PROC
type CorporateActionCancellationReason1Code string

type CorporateActionEventStatus1 struct {
	EvtCmpltnsSts EventCompletenessStatus1Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.03 EvtCmpltnsSts"`
	EvtConfSts    EventConfirmationStatus1Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.03 EvtConfSts"`
}

type CorporateActionEventType7Choice struct {
	Cd    CorporateActionEventType8Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.03 Cd"`
	Prtry GenericIdentification20       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.03 Prtry"`
}

// May be one of ACTV, ATTI, BIDS, BONU, BPUT, BRUP, CAPG, CAPI, CERT, CHAN, CLSA, CONS, CONV, COOP, DECR, DETI, DFLT, DLST, DRAW, DRIP, DSCL, DTCH, DVCA, DVOP, DVSC, DVSE, EXOF, EXRI, EXTM, EXWA, CAPD, INCR, INTR, LIQU, MCAL, MRGR, ODLT, OTHR, PARI, PCAL, PDEF, PINK, PLAC, PPMT, PRED, PRII, PRIO, REDM, REDO, REMK, RHDI, RHTS, SHPR, SMAL, SOFF, SPLF, SPLR, SUSP, TEND, TREC, WRTH, WTRC, CREV, DRCA
type CorporateActionEventType8Code string

type CorporateActionGeneralInformation35 struct {
	CorpActnEvtId      Max35Text                                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.03 CorpActnEvtId"`
	OffclCorpActnEvtId Max35Text                                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.03 OffclCorpActnEvtId,omitempty"`
	ClssActnNb         Max35Text                                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.03 ClssActnNb,omitempty"`
	EvtTp              CorporateActionEventType7Choice          `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.03 EvtTp"`
	MndtryVlntryEvtTp  CorporateActionMandatoryVoluntary1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.03 MndtryVlntryEvtTp"`
	FinInstrmId        SecurityIdentification14                 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.03 FinInstrmId"`
}

type CorporateActionMandatoryVoluntary1Choice struct {
	Cd    CorporateActionMandatoryVoluntary1Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.03 Cd"`
	Prtry GenericIdentification20                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.03 Prtry"`
}

// May be one of MAND, CHOS, VOLU
type CorporateActionMandatoryVoluntary1Code string

type CorporateActionProcessingStatus1Choice struct {
	EvtSts        CorporateActionEventStatus1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.03 EvtSts"`
	ForInfOnlyInd bool                        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.03 ForInfOnlyInd"`
}

// Must match the pattern [A-Z]{2,2}
type CountryCode string

type Document struct {
	CorpActnCxlAdvc CorporateActionCancellationAdviceV03 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.03 CorpActnCxlAdvc"`
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
	Id      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.03 Id"`
	Issr    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.03 Issr"`
	SchmeNm Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.03 SchmeNm,omitempty"`
}

type GenericIdentification20 struct {
	Id      Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.03 Id"`
	Issr    Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.03 Issr"`
	SchmeNm Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.03 SchmeNm,omitempty"`
}

type GenericIdentification21 struct {
	Tp GenericIdentification20 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.03 Tp"`
	Id Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.03 Id,omitempty"`
}

// Must match the pattern [A-Z0-9]{12,12}
type ISINIdentifier string

type IdentificationSource3Choice struct {
	Cd    ExternalFinancialInstrumentIdentificationType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.03 Cd"`
	Prtry Max35Text                                          `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.03 Prtry"`
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
	Nm  Max350Text     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.03 Nm"`
	Adr PostalAddress1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.03 Adr,omitempty"`
}

type OtherIdentification1 struct {
	Id  Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.03 Id"`
	Sfx Max16Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.03 Sfx,omitempty"`
	Tp  IdentificationSource3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.03 Tp"`
}

type PartyIdentification36Choice struct {
	AnyBIC  AnyBICIdentifier        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.03 AnyBIC"`
	PrtryId GenericIdentification19 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.03 PrtryId"`
}

type PartyIdentification46Choice struct {
	AnyBIC   AnyBICIdentifier        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.03 AnyBIC"`
	PrtryId  GenericIdentification19 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.03 PrtryId"`
	NmAndAdr NameAndAddress5         `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.03 NmAndAdr"`
}

type PostalAddress1 struct {
	AdrTp       AddressType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.03 AdrTp,omitempty"`
	AdrLine     []Max70Text      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.03 AdrLine,omitempty"`
	StrtNm      Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.03 StrtNm,omitempty"`
	BldgNb      Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.03 BldgNb,omitempty"`
	PstCd       Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.03 PstCd,omitempty"`
	TwnNm       Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.03 TwnNm,omitempty"`
	CtrySubDvsn Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.03 CtrySubDvsn,omitempty"`
	Ctry        CountryCode      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.03 Ctry"`
}

// May be one of GENR
type SafekeepingAccountIdentification1Code string

// May be one of CUST, ICSD, NCSD, SHHE
type SafekeepingPlace1Code string

// May be one of SHHE, ALLP
type SafekeepingPlace2Code string

type SafekeepingPlaceFormat2Choice struct {
	Id      SafekeepingPlaceTypeAndText2             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.03 Id"`
	Ctry    CountryCode                              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.03 Ctry"`
	TpAndId SafekeepingPlaceTypeAndAnyBICIdentifier1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.03 TpAndId"`
	Prtry   GenericIdentification21                  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.03 Prtry"`
}

type SafekeepingPlaceTypeAndAnyBICIdentifier1 struct {
	SfkpgPlcTp SafekeepingPlace1Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.03 SfkpgPlcTp"`
	Id         AnyBICIdentifier      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.03 Id"`
}

type SafekeepingPlaceTypeAndText2 struct {
	SfkpgPlcTp SafekeepingPlace2Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.03 SfkpgPlcTp"`
	Id         Max35Text             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.03 Id,omitempty"`
}

type SecurityIdentification14 struct {
	ISIN   ISINIdentifier         `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.03 ISIN,omitempty"`
	OthrId []OtherIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.03 OthrId,omitempty"`
	Desc   Max140Text             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.03 Desc,omitempty"`
}

type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.03 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.03 Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
}
