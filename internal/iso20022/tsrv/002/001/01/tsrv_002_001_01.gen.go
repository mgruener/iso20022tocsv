// Code generated by main. DO NOT EDIT.

package tsrv_002_001_01

import (
	"bytes"
	"encoding/base64"
	"encoding/xml"
	"time"
)

type ActiveCurrencyAndAmount struct {
	Value float64            `xml:",chardata"`
	Ccy   ActiveCurrencyCode `xml:"Ccy,attr"`
}

// Must match the pattern [A-Z]{3,3}
type ActiveCurrencyCode string

// May be one of ADDR, PBOX, HOME, BIZZ, MLTO, DLVY
type AddressType2Code string

type AdvisingPartyAdditionalInformation1 struct {
	RefNb         Max35Text     `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 RefNb,omitempty"`
	BkToBnfcryInf []Max2000Text `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 BkToBnfcryInf,omitempty"`
}

type AmountAndTrigger1 struct {
	Id         Max35Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 Id"`
	AmtDtlsChc AmountOrPercentage1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 AmtDtlsChc"`
	Trggr      []Trigger1                `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 Trggr"`
}

type AmountOrPercentage1Choice struct {
	DfndAmt UndertakingAmount4 `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 DfndAmt"`
	PctgAmt Percentage1        `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 PctgAmt"`
}

// Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type AnyBICIdentifier string

type AutoExtend1Choice struct {
	Days  float64 `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 Days"`
	Mnths float64 `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 Mnths"`
	Yrs   float64 `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 Yrs"`
	Dt    ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 Dt"`
}

type AutoExtension1 struct {
	Prd            AutoExtend1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 Prd,omitempty"`
	FnlXpryDt      ISODate           `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 FnlXpryDt,omitempty"`
	NonXtnsnNtfctn []NonExtension1   `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 NonXtnsnNtfctn,omitempty"`
}

type AutomaticVariation1 struct {
	Id          Max35Text           `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 Id"`
	Tp          VariationType1Code  `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 Tp"`
	AmtAndTrggr []AmountAndTrigger1 `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 AmtAndTrggr"`
	AddtlInf    []Max2000Text       `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 AddtlInf,omitempty"`
}

type Channel1Choice struct {
	Cd    ExternalChannel1Code   `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 Cd"`
	Prtry GenericIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 Prtry"`
}

type CommunicationChannel1 struct {
	Mtd         ExternalChannel1Code `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 Mtd"`
	DlvrToPtyTp PartyType1Choice     `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 DlvrToPtyTp"`
	DlvrToNm    Max140Text           `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 DlvrToNm,omitempty"`
	DlvrToAdr   PostalAddress6       `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 DlvrToAdr,omitempty"`
}

type CommunicationMethod1Choice struct {
	Cd    ExternalChannel1Code   `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 Cd"`
	Prtry GenericIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 Prtry"`
}

type ContactDetails2 struct {
	NmPrfx   NamePrefix1Code `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 NmPrfx,omitempty"`
	Nm       Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 Nm,omitempty"`
	PhneNb   PhoneNumber     `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 PhneNb,omitempty"`
	MobNb    PhoneNumber     `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 MobNb,omitempty"`
	FaxNb    PhoneNumber     `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 FaxNb,omitempty"`
	EmailAdr Max2048Text     `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 EmailAdr,omitempty"`
	Othr     Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 Othr,omitempty"`
}

// Must match the pattern [A-Z]{2,2}
type CountryCode string

type CountrySubdivision1Choice struct {
	Cd    Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 Cd"`
	Prtry GenericIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 Prtry"`
}

type DateAndDateTimeChoice struct {
	Dt   ISODate     `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 Dt"`
	DtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 DtTm"`
}

type DateAndPlaceOfBirth struct {
	BirthDt     ISODate     `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 BirthDt"`
	PrvcOfBirth Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 PrvcOfBirth,omitempty"`
	CityOfBirth Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 CityOfBirth"`
	CtryOfBirth CountryCode `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 CtryOfBirth"`
}

type DateInformation1 struct {
	StartDt ISODate                    `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 StartDt"`
	Frqcy   ExternalDateFrequency1Code `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 Frqcy"`
	Nb      float64                    `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 Nb"`
}

type Document struct {
	UdrtkgIssncAdvc UndertakingIssuanceAdviceV01 `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 UdrtkgIssncAdvc"`
}

type Document10 struct {
	DocTp        UndertakingDocumentType2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 DocTp"`
	PresntnChanl Channel1Choice                 `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 PresntnChanl,omitempty"`
	DocFrmt      DocumentFormat1Choice          `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 DocFrmt,omitempty"`
	CpyInd       bool                           `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 CpyInd,omitempty"`
	SgndInd      bool                           `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 SgndInd,omitempty"`
}

type Document8 struct {
	Tp          PresentationDocumentFormat1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 Tp"`
	Wrdg        Max20000Text                      `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 Wrdg,omitempty"`
	ElctrncDtls []Presentation3                   `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 ElctrncDtls,omitempty"`
}

type Document9 struct {
	Tp        UndertakingDocumentType1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 Tp"`
	Id        Max35Text                      `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 Id"`
	Frmt      DocumentFormat1Choice          `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 Frmt,omitempty"`
	Nclsr     Max2MBBinary                   `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 Nclsr"`
	DgtlSgntr PartyAndSignature2             `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 DgtlSgntr,omitempty"`
}

type DocumentFormat1Choice struct {
	Cd    ExternalDocumentFormat1Code `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 Cd"`
	Prtry GenericIdentification1      `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 Prtry"`
}

type ExpiryDetails1 struct {
	XpryTerms    ExpiryTerms1  `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 XpryTerms,omitempty"`
	AddtlXpryInf []Max2000Text `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 AddtlXpryInf,omitempty"`
}

type ExpiryTerms1 struct {
	DtTm       DateAndDateTimeChoice `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 DtTm,omitempty"`
	AutoXtnsn  AutoExtension1        `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 AutoXtnsn,omitempty"`
	Cond       Max2000Text           `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 Cond,omitempty"`
	OpnEnddInd bool                  `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 OpnEnddInd,omitempty"`
}

// Must be at least 1 items long
type ExternalChannel1Code string

// Must be at least 1 items long
type ExternalDateFrequency1Code string

// Must be at least 1 items long
type ExternalDocumentFormat1Code string

// Must be at least 1 items long
type ExternalModelFormIdentification1Code string

// Must be at least 1 items long
type ExternalNarrativeType1Code string

// Must be at least 1 items long
type ExternalOrganisationIdentification1Code string

// Must be at least 1 items long
type ExternalPersonIdentification1Code string

// Must be at least 1 items long
type ExternalRelativeTo1Code string

// Must be at least 1 items long
type ExternalTypeOfParty1Code string

// Must be at least 1 items long
type ExternalUnderlyingTradeTransactionType1Code string

// Must be at least 1 items long
type ExternalUndertakingDocumentType1Code string

// Must be at least 1 items long
type ExternalUndertakingDocumentType2Code string

// Must be at least 1 items long
type ExternalUndertakingType1Code string

type FixedOrRecurrentDate1Choice struct {
	FxdDt   ISODate          `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 FxdDt"`
	RcrntDt DateInformation1 `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 RcrntDt"`
}

type GenericIdentification1 struct {
	Id      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 Id"`
	SchmeNm Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 SchmeNm,omitempty"`
	Issr    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 Issr,omitempty"`
}

type GenericOrganisationIdentification1 struct {
	Id      Max35Text                                   `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 Id"`
	SchmeNm OrganisationIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 SchmeNm,omitempty"`
	Issr    Max35Text                                   `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 Issr,omitempty"`
}

type GenericPersonIdentification1 struct {
	Id      Max35Text                             `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 Id"`
	SchmeNm PersonIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 SchmeNm,omitempty"`
	Issr    Max35Text                             `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 Issr,omitempty"`
}

type GovernanceIdentification1Choice struct {
	Cd    GovernanceIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 Cd"`
	Prtry GenericIdentification1        `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 Prtry"`
}

// May be one of ISPR, NONE, UCPR, URDG
type GovernanceIdentification1Code string

type GovernanceRules1 struct {
	RuleId   GovernanceIdentification1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 RuleId"`
	AplblLaw Location1                       `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 AplblLaw,omitempty"`
	Jursdctn []Location1                     `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 Jursdctn,omitempty"`
}

// Must match the pattern [a-z]{2,2}
type ISO2ALanguageCode string

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

// May be one of CRQL, CRQC, ISSU, ISCO, ISAD
type IssuanceType1Code string

type Location1 struct {
	Ctry        CountryCode               `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 Ctry,omitempty"`
	CtrySubDvsn CountrySubdivision1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 CtrySubDvsn,omitempty"`
	Txt         []Max2000Text             `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 Txt,omitempty"`
}

// Must be at least 1 items long
type Max140Text string

// Must be at least 1 items long
type Max16Text string

// Must be at least 1 items long
type Max20000Text string

// Must be at least 1 items long
type Max2000Text string

// Must be at least 1 items long
type Max2048Text string

// Must be at least 1 items long
type Max256Text string

type Max2MBBinary []byte

func (t *Max2MBBinary) UnmarshalText(text []byte) error {
	return (*xsdBase64Binary)(t).UnmarshalText(text)
}
func (t Max2MBBinary) MarshalText() ([]byte, error) {
	return xsdBase64Binary(t).MarshalText()
}

// Must be at least 1 items long
type Max35Text string

// Must be at least 1 items long
type Max70Text string

type ModelFormIdentification1 struct {
	Id   ModelFormIdentification1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 Id"`
	Vrsn Max35Text                      `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 Vrsn,omitempty"`
}

type ModelFormIdentification1Choice struct {
	Cd    ExternalModelFormIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 Cd"`
	Prtry GenericIdentification1               `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 Prtry"`
}

// May be one of DOCT, MIST, MISS, MADM
type NamePrefix1Code string

type Narrative1 struct {
	Tp  NarrativeType1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 Tp,omitempty"`
	Txt []Max20000Text       `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 Txt"`
}

type NarrativeType1Choice struct {
	Cd    ExternalNarrativeType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 Cd"`
	Prtry GenericIdentification1     `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 Prtry"`
}

type NonExtension1 struct {
	NtfctnPrd     float64                    `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 NtfctnPrd,omitempty"`
	NtfctnMtd     CommunicationMethod1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 NtfctnMtd,omitempty"`
	NtfctnRcptTp  PartyType1Choice           `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 NtfctnRcptTp,omitempty"`
	NtfctnRcptNm  Max140Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 NtfctnRcptNm,omitempty"`
	NtfctnRcptAdr PostalAddress6             `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 NtfctnRcptAdr,omitempty"`
}

type OrganisationIdentification8 struct {
	AnyBIC AnyBICIdentifier                     `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 AnyBIC,omitempty"`
	Othr   []GenericOrganisationIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 Othr,omitempty"`
}

type OrganisationIdentificationSchemeName1Choice struct {
	Cd    ExternalOrganisationIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 Cd"`
	Prtry Max35Text                               `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 Prtry"`
}

type Party11Choice struct {
	OrgId  OrganisationIdentification8 `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 OrgId"`
	PrvtId PersonIdentification5       `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 PrvtId"`
}

type PartyAndSignature2 struct {
	Pty   PartyIdentification43 `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 Pty"`
	Sgntr ProprietaryData3      `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 Sgntr"`
}

type PartyAndType1 struct {
	Tp  PartyType1Choice      `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 Tp"`
	Pty PartyIdentification43 `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 Pty,omitempty"`
}

type PartyIdentification43 struct {
	Nm        Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 Nm,omitempty"`
	PstlAdr   PostalAddress6  `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 PstlAdr,omitempty"`
	Id        Party11Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 Id,omitempty"`
	CtryOfRes CountryCode     `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 CtryOfRes,omitempty"`
	CtctDtls  ContactDetails2 `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 CtctDtls,omitempty"`
}

type PartyType1Choice struct {
	Cd    ExternalTypeOfParty1Code `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 Cd"`
	Prtry GenericIdentification1   `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 Prtry"`
}

type Percentage1 struct {
	Rate   float64                 `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 Rate"`
	RltvTo ExternalRelativeTo1Code `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 RltvTo"`
}

type PersonIdentification5 struct {
	DtAndPlcOfBirth DateAndPlaceOfBirth            `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 DtAndPlcOfBirth,omitempty"`
	Othr            []GenericPersonIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 Othr,omitempty"`
}

type PersonIdentificationSchemeName1Choice struct {
	Cd    ExternalPersonIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 Cd"`
	Prtry Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 Prtry"`
}

// Must match the pattern \+[0-9]{1,3}-[0-9()+\-]{1,30}
type PhoneNumber string

type PlaceOfPresentation1 struct {
	Plc  ExternalTypeOfParty1Code `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 Plc"`
	Ctry CountryCode              `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 Ctry,omitempty"`
}

type PlaceOrUnderConfirmationChoice1 struct {
	PlcOfPresntn   PlaceOfPresentation1   `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 PlcOfPresntn"`
	PresntnUdrConf PresentationParty1Code `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 PresntnUdrConf"`
}

type PostalAddress12 struct {
	TwnNm       Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 TwnNm,omitempty"`
	CtrySubDvsn Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 CtrySubDvsn,omitempty"`
	Ctry        CountryCode `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 Ctry"`
}

type PostalAddress6 struct {
	AdrTp       AddressType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 AdrTp,omitempty"`
	Dept        Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 Dept,omitempty"`
	SubDept     Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 SubDept,omitempty"`
	StrtNm      Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 StrtNm,omitempty"`
	BldgNb      Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 BldgNb,omitempty"`
	PstCd       Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 PstCd,omitempty"`
	TwnNm       Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 TwnNm,omitempty"`
	CtrySubDvsn Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 CtrySubDvsn,omitempty"`
	Ctry        CountryCode      `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 Ctry,omitempty"`
	AdrLine     []Max70Text      `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 AdrLine,omitempty"`
}

type Presentation1 struct {
	Mdm                      PresentationMedium1Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 Mdm,omitempty"`
	PlcOfPresntnOrUdrConfChc PlaceOrUnderConfirmationChoice1 `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 PlcOfPresntnOrUdrConfChc,omitempty"`
	Doc                      []Document8                     `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 Doc,omitempty"`
	AddtlInf                 []Max2000Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 AddtlInf,omitempty"`
}

type Presentation3 struct {
	Frmt  DocumentFormat1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 Frmt,omitempty"`
	Chanl Channel1Choice        `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 Chanl,omitempty"`
	Adr   Max256Text            `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 Adr,omitempty"`
}

type PresentationDocumentFormat1Choice struct {
	Cd    ExternalUndertakingDocumentType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 Cd"`
	Prtry GenericIdentification1               `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 Prtry"`
}

type PresentationMedium1Choice struct {
	Cd    PresentationMedium1Code `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 Cd"`
	Prtry GenericIdentification1  `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 Prtry"`
}

// May be one of BOTH, ELEC, PAPR
type PresentationMedium1Code string

// May be one of ETHR, EXCN, EXIS
type PresentationParty1Code string

type ProprietaryData3 struct {
	Item string `xml:",any"`
}

type Trigger1 struct {
	DtChc      FixedOrRecurrentDate1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 DtChc,omitempty"`
	DcmntryEvt []Document10                `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 DcmntryEvt,omitempty"`
}

type UnderlyingTradeTransaction1 struct {
	Tp           UnderlyingTradeTransactionType1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 Tp"`
	Id           Max35Text                             `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 Id,omitempty"`
	TxDt         ISODate                               `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 TxDt,omitempty"`
	TndrClsgDt   ISODate                               `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 TndrClsgDt,omitempty"`
	TxAmt        ActiveCurrencyAndAmount               `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 TxAmt,omitempty"`
	CtrctAmtPctg float64                               `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 CtrctAmtPctg,omitempty"`
	AddtlInf     []Max2000Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 AddtlInf,omitempty"`
}

type UnderlyingTradeTransactionType1Choice struct {
	Cd    ExternalUnderlyingTradeTransactionType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 Cd"`
	Prtry GenericIdentification1                      `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 Prtry"`
}

type Undertaking3 struct {
	Id                  Max35Text                     `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 Id"`
	Nm                  UndertakingIssuanceName1Code  `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 Nm"`
	Tp                  UndertakingType1Choice        `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 Tp,omitempty"`
	IssncTp             IssuanceType1Code             `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 IssncTp"`
	Applcnt             []PartyIdentification43       `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 Applcnt,omitempty"`
	Issr                PartyIdentification43         `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 Issr"`
	Bnfcry              []PartyIdentification43       `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 Bnfcry"`
	DtOfIssnc           ISODate                       `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 DtOfIssnc"`
	PlcOfIsse           PostalAddress12               `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 PlcOfIsse,omitempty"`
	AdvsgPty            PartyIdentification43         `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 AdvsgPty,omitempty"`
	ScndAdvsgPty        PartyIdentification43         `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 ScndAdvsgPty,omitempty"`
	UdrtkgAmt           UndertakingAmount1            `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 UdrtkgAmt"`
	XpryDtls            ExpiryDetails1                `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 XpryDtls"`
	ConfInd             bool                          `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 ConfInd,omitempty"`
	ConfPtyTp           ExternalTypeOfParty1Code      `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 ConfPtyTp,omitempty"`
	AddtlPty            []PartyAndType1               `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 AddtlPty,omitempty"`
	GovncRulesAndLaw    GovernanceRules1              `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 GovncRulesAndLaw"`
	UndrlygTx           []UnderlyingTradeTransaction1 `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 UndrlygTx,omitempty"`
	PresntnDtls         Presentation1                 `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 PresntnDtls,omitempty"`
	UdrtkgTermsAndConds []Narrative1                  `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 UdrtkgTermsAndConds"`
	MltplDmndInd        bool                          `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 MltplDmndInd,omitempty"`
	PrtlDmndInd         bool                          `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 PrtlDmndInd,omitempty"`
	ConfChrgsPyblBy     ExternalTypeOfParty1Code      `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 ConfChrgsPyblBy,omitempty"`
	TrfChrgsPyblBy      ExternalTypeOfParty1Code      `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 TrfChrgsPyblBy,omitempty"`
	AutomtcAmtVartn     []AutomaticVariation1         `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 AutomtcAmtVartn,omitempty"`
	DlvryChanl          CommunicationChannel1         `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 DlvryChanl,omitempty"`
	TrfInd              bool                          `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 TrfInd,omitempty"`
	NclsdFile           []Document9                   `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 NclsdFile,omitempty"`
	AddtlInf            []Max2000Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 AddtlInf,omitempty"`
	ReqdLclUdrtkg       Undertaking4                  `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 ReqdLclUdrtkg,omitempty"`
}

type Undertaking4 struct {
	Nm               UndertakingName1Code          `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 Nm"`
	Tp               ExternalUndertakingType1Code  `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 Tp"`
	Applcnt          []PartyIdentification43       `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 Applcnt"`
	Bnfcry           []PartyIdentification43       `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 Bnfcry"`
	DtOfIssnc        ISODate                       `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 DtOfIssnc,omitempty"`
	AdvsgPty         PartyIdentification43         `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 AdvsgPty,omitempty"`
	ScndAdvsgPty     PartyIdentification43         `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 ScndAdvsgPty,omitempty"`
	LclUdrtkgAmt     UndertakingAmount1            `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 LclUdrtkgAmt"`
	XpryDtls         ExpiryDetails1                `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 XpryDtls"`
	ConfInd          bool                          `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 ConfInd,omitempty"`
	AddtlPty         []PartyAndType1               `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 AddtlPty,omitempty"`
	GovncRulesAndLaw GovernanceRules1              `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 GovncRulesAndLaw"`
	UndrlygTx        []UnderlyingTradeTransaction1 `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 UndrlygTx,omitempty"`
	PresntnDtls      Presentation1                 `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 PresntnDtls,omitempty"`
	UdrtkgWrdg       UndertakingWording1           `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 UdrtkgWrdg"`
	MltplDmndInd     bool                          `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 MltplDmndInd,omitempty"`
	PrtlDmndInd      bool                          `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 PrtlDmndInd,omitempty"`
	ConfChrgsPyblBy  ExternalTypeOfParty1Code      `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 ConfChrgsPyblBy,omitempty"`
	TrfChrgsPyblBy   ExternalTypeOfParty1Code      `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 TrfChrgsPyblBy,omitempty"`
	AutomtcAmtVartn  []AutomaticVariation1         `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 AutomtcAmtVartn,omitempty"`
	DlvryChanl       CommunicationChannel1         `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 DlvryChanl,omitempty"`
	TrfInd           bool                          `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 TrfInd,omitempty"`
	AddtlInf         []Max2000Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 AddtlInf,omitempty"`
}

type UndertakingAdvice1 struct {
	UdrtkgIssncMsg       UndertakingIssuanceMessage          `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 UdrtkgIssncMsg"`
	FrstAdvsgPtyAddtlInf AdvisingPartyAdditionalInformation1 `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 FrstAdvsgPtyAddtlInf,omitempty"`
	ScndAdvsgPtyAddtlInf AdvisingPartyAdditionalInformation1 `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 ScndAdvsgPtyAddtlInf,omitempty"`
	ConfDtls             UndertakingConfirmation1            `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 ConfDtls,omitempty"`
	DgtlSgntr            []PartyAndSignature2                `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 DgtlSgntr,omitempty"`
}

type UndertakingAmount1 struct {
	Amt        ActiveCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 Amt"`
	PlusTlrnce float64                 `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 PlusTlrnce,omitempty"`
	AddtlInf   []Max2000Text           `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 AddtlInf,omitempty"`
}

type UndertakingAmount4 struct {
	VartnAmt ActiveCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 VartnAmt"`
	BalAmt   ActiveCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 BalAmt,omitempty"`
}

type UndertakingConfirmation1 struct {
	Cnfrmr PartyIdentification43 `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 Cnfrmr"`
	RefNb  Max35Text             `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 RefNb"`
	Dt     DateAndDateTimeChoice `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 Dt"`
	Conf   []Max2000Text         `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 Conf,omitempty"`
}

type UndertakingDocumentType1Choice struct {
	Cd    ExternalUndertakingDocumentType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 Cd"`
	Prtry GenericIdentification1               `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 Prtry"`
}

type UndertakingDocumentType2Choice struct {
	Cd    ExternalUndertakingDocumentType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 Cd"`
	Prtry GenericIdentification1               `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 Prtry"`
}

type UndertakingIssuanceAdviceV01 struct {
	AdvsgPty            PartyIdentification43 `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 AdvsgPty"`
	ScndAdvsgPty        PartyIdentification43 `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 ScndAdvsgPty,omitempty"`
	DtOfAdvc            DateAndDateTimeChoice `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 DtOfAdvc"`
	UdrtkgIssncAdvcDtls UndertakingAdvice1    `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 UdrtkgIssncAdvcDtls"`
	BkToBkInf           []Max2000Text         `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 BkToBkInf,omitempty"`
	DgtlSgntr           PartyAndSignature2    `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 DgtlSgntr,omitempty"`
}

type UndertakingIssuanceMessage struct {
	UdrtkgDtls Undertaking3       `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 UdrtkgDtls"`
	DgtlSgntr  PartyAndSignature2 `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 DgtlSgntr,omitempty"`
}

// May be one of STBY, DGAR
type UndertakingIssuanceName1Code string

// May be one of STBY, DGAR
type UndertakingName1Code string

type UndertakingType1Choice struct {
	Cd    ExternalUndertakingType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 Cd"`
	Prtry GenericIdentification1       `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 Prtry"`
}

type UndertakingWording1 struct {
	MdlForm             ModelFormIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 MdlForm,omitempty"`
	ReqdWrdgLang        ISO2ALanguageCode        `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 ReqdWrdgLang,omitempty"`
	UdrtkgTermsAndConds []Narrative1             `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 UdrtkgTermsAndConds,omitempty"`
}

// May be one of DECR, INCR
type VariationType1Code string

type xsdBase64Binary []byte

func (b *xsdBase64Binary) UnmarshalText(text []byte) (err error) {
	*b, err = base64.StdEncoding.DecodeString(string(text))
	return
}
func (b xsdBase64Binary) MarshalText() ([]byte, error) {
	var buf bytes.Buffer
	enc := base64.NewEncoder(base64.StdEncoding, &buf)
	enc.Write([]byte(b))
	enc.Close()
	return buf.Bytes(), nil
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
