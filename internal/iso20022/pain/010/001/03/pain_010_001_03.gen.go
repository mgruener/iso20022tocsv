// Code generated by main. DO NOT EDIT.

package pain_010_001_03

import (
	"bytes"
	"encoding/xml"
	"time"
)

type AccountIdentification4Choice struct {
	IBAN IBAN2007Identifier            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 IBAN"`
	Othr GenericAccountIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 Othr"`
}

type AccountSchemeName1Choice struct {
	Cd    ExternalAccountIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 Cd"`
	Prtry Max35Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 Prtry"`
}

type ActiveCurrencyAndAmount struct {
	Value float64            `xml:",chardata"`
	Ccy   ActiveCurrencyCode `xml:"Ccy,attr"`
}

// Must match the pattern [A-Z]{3,3}
type ActiveCurrencyCode string

type ActiveOrHistoricCurrencyAndAmount struct {
	Value float64                      `xml:",chardata"`
	Ccy   ActiveOrHistoricCurrencyCode `xml:"Ccy,attr"`
}

// Must match the pattern [A-Z]{3,3}
type ActiveOrHistoricCurrencyCode string

// May be one of ADDR, PBOX, HOME, BIZZ, MLTO, DLVY
type AddressType2Code string

// Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type AnyBICIdentifier string

type Authorisation1Choice struct {
	Cd    Authorisation1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 Cd"`
	Prtry Max128Text         `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 Prtry"`
}

// May be one of AUTH, FDET, FSUM, ILEV
type Authorisation1Code string

// Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type BICFIIdentifier string

type BranchAndFinancialInstitutionIdentification5 struct {
	FinInstnId FinancialInstitutionIdentification8 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 FinInstnId"`
	BrnchId    BranchData2                         `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 BrnchId,omitempty"`
}

type BranchData2 struct {
	Id      Max35Text      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 Id,omitempty"`
	Nm      Max140Text     `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 Nm,omitempty"`
	PstlAdr PostalAddress6 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 PstlAdr,omitempty"`
}

type CashAccount24 struct {
	Id  AccountIdentification4Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 Id"`
	Tp  CashAccountType2Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 Tp,omitempty"`
	Ccy ActiveOrHistoricCurrencyCode `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 Ccy,omitempty"`
	Nm  Max70Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 Nm,omitempty"`
}

type CashAccountType2Choice struct {
	Cd    ExternalCashAccountType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 Cd"`
	Prtry Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 Prtry"`
}

type ClearingSystemIdentification2Choice struct {
	Cd    ExternalClearingSystemIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 Cd"`
	Prtry Max35Text                                 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 Prtry"`
}

type ClearingSystemMemberIdentification2 struct {
	ClrSysId ClearingSystemIdentification2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 ClrSysId,omitempty"`
	MmbId    Max35Text                           `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 MmbId"`
}

type ContactDetails2 struct {
	NmPrfx   NamePrefix1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 NmPrfx,omitempty"`
	Nm       Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 Nm,omitempty"`
	PhneNb   PhoneNumber     `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 PhneNb,omitempty"`
	MobNb    PhoneNumber     `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 MobNb,omitempty"`
	FaxNb    PhoneNumber     `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 FaxNb,omitempty"`
	EmailAdr Max2048Text     `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 EmailAdr,omitempty"`
	Othr     Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 Othr,omitempty"`
}

// Must match the pattern [A-Z]{2,2}
type CountryCode string

type DateAndPlaceOfBirth struct {
	BirthDt     ISODate     `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 BirthDt"`
	PrvcOfBirth Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 PrvcOfBirth,omitempty"`
	CityOfBirth Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 CityOfBirth"`
	CtryOfBirth CountryCode `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 CtryOfBirth"`
}

type DatePeriodDetails1 struct {
	FrDt ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 FrDt"`
	ToDt ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 ToDt,omitempty"`
}

type Document struct {
	MndtAmdmntReq MandateAmendmentRequestV03 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 MndtAmdmntReq"`
}

// May be one of MSIN, CNFA, DNFA, CINV, CREN, DEBN, HIRI, SBIN, CMCN, SOAC, DISP, BOLD, VCHR, AROI, TSUT
type DocumentType5Code string

// Must be at least 1 items long
type ExternalAccountIdentification1Code string

// Must be at least 1 items long
type ExternalCashAccountType1Code string

// Must be at least 1 items long
type ExternalClearingSystemIdentification1Code string

// Must be at least 1 items long
type ExternalFinancialInstitutionIdentification1Code string

// Must be at least 1 items long
type ExternalLocalInstrument1Code string

// Must be at least 1 items long
type ExternalMandateReason1Code string

// Must be at least 1 items long
type ExternalOrganisationIdentification1Code string

// Must be at least 1 items long
type ExternalPersonIdentification1Code string

// Must be at least 1 items long
type ExternalServiceLevel1Code string

type FinancialIdentificationSchemeName1Choice struct {
	Cd    ExternalFinancialInstitutionIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 Cd"`
	Prtry Max35Text                                       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 Prtry"`
}

type FinancialInstitutionIdentification8 struct {
	BICFI       BICFIIdentifier                     `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 BICFI,omitempty"`
	ClrSysMmbId ClearingSystemMemberIdentification2 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 ClrSysMmbId,omitempty"`
	Nm          Max140Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 Nm,omitempty"`
	PstlAdr     PostalAddress6                      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 PstlAdr,omitempty"`
	Othr        GenericFinancialIdentification1     `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 Othr,omitempty"`
}

// May be one of YEAR, MNTH, QURT, MIAN, WEEK, DAIL, ADHO, INDA, FRTN
type Frequency6Code string

type GenericAccountIdentification1 struct {
	Id      Max34Text                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 Id"`
	SchmeNm AccountSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 SchmeNm,omitempty"`
	Issr    Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 Issr,omitempty"`
}

type GenericFinancialIdentification1 struct {
	Id      Max35Text                                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 Id"`
	SchmeNm FinancialIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 SchmeNm,omitempty"`
	Issr    Max35Text                                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 Issr,omitempty"`
}

type GenericOrganisationIdentification1 struct {
	Id      Max35Text                                   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 Id"`
	SchmeNm OrganisationIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 SchmeNm,omitempty"`
	Issr    Max35Text                                   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 Issr,omitempty"`
}

type GenericPersonIdentification1 struct {
	Id      Max35Text                             `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 Id"`
	SchmeNm PersonIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 SchmeNm,omitempty"`
	Issr    Max35Text                             `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 Issr,omitempty"`
}

type GroupHeader47 struct {
	MsgId    Max35Text                                    `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 MsgId"`
	CreDtTm  ISODateTime                                  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 CreDtTm"`
	Authstn  []Authorisation1Choice                       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 Authstn,omitempty"`
	InitgPty PartyIdentification43                        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 InitgPty,omitempty"`
	InstgAgt BranchAndFinancialInstitutionIdentification5 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 InstgAgt,omitempty"`
	InstdAgt BranchAndFinancialInstitutionIdentification5 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 InstdAgt,omitempty"`
}

// Must match the pattern [A-Z]{2,2}[0-9]{2,2}[a-zA-Z0-9]{1,30}
type IBAN2007Identifier string

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

type LocalInstrument2Choice struct {
	Cd    ExternalLocalInstrument1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 Cd"`
	Prtry Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 Prtry"`
}

type Mandate1 struct {
	MndtId      Max35Text                                    `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 MndtId"`
	MndtReqId   Max35Text                                    `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 MndtReqId,omitempty"`
	Tp          MandateTypeInformation1                      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 Tp,omitempty"`
	Ocrncs      MandateOccurrences2                          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 Ocrncs,omitempty"`
	ColltnAmt   ActiveOrHistoricCurrencyAndAmount            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 ColltnAmt,omitempty"`
	MaxAmt      ActiveOrHistoricCurrencyAndAmount            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 MaxAmt,omitempty"`
	CdtrSchmeId PartyIdentification43                        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 CdtrSchmeId,omitempty"`
	Cdtr        PartyIdentification43                        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 Cdtr"`
	CdtrAcct    CashAccount24                                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 CdtrAcct,omitempty"`
	CdtrAgt     BranchAndFinancialInstitutionIdentification5 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 CdtrAgt,omitempty"`
	UltmtCdtr   PartyIdentification43                        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 UltmtCdtr,omitempty"`
	Dbtr        PartyIdentification43                        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 Dbtr"`
	DbtrAcct    CashAccount24                                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 DbtrAcct,omitempty"`
	DbtrAgt     BranchAndFinancialInstitutionIdentification5 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 DbtrAgt"`
	UltmtDbtr   PartyIdentification43                        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 UltmtDbtr,omitempty"`
	RfrdDoc     ReferredDocumentInformation3                 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 RfrdDoc,omitempty"`
}

type Mandate3 struct {
	MndtId      Max35Text                                    `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 MndtId"`
	MndtReqId   Max35Text                                    `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 MndtReqId,omitempty"`
	Tp          MandateTypeInformation1                      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 Tp,omitempty"`
	Ocrncs      MandateOccurrences2                          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 Ocrncs,omitempty"`
	ColltnAmt   ActiveCurrencyAndAmount                      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 ColltnAmt,omitempty"`
	MaxAmt      ActiveCurrencyAndAmount                      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 MaxAmt,omitempty"`
	CdtrSchmeId PartyIdentification43                        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 CdtrSchmeId,omitempty"`
	Cdtr        PartyIdentification43                        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 Cdtr,omitempty"`
	CdtrAcct    CashAccount24                                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 CdtrAcct,omitempty"`
	CdtrAgt     BranchAndFinancialInstitutionIdentification5 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 CdtrAgt,omitempty"`
	UltmtCdtr   PartyIdentification43                        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 UltmtCdtr,omitempty"`
	Dbtr        PartyIdentification43                        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 Dbtr,omitempty"`
	DbtrAcct    CashAccount24                                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 DbtrAcct,omitempty"`
	DbtrAgt     BranchAndFinancialInstitutionIdentification5 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 DbtrAgt,omitempty"`
	UltmtDbtr   PartyIdentification43                        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 UltmtDbtr,omitempty"`
	RfrdDoc     ReferredDocumentInformation3                 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 RfrdDoc,omitempty"`
}

type MandateAmendment3 struct {
	OrgnlMsgInf OriginalMessageInformation1 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 OrgnlMsgInf,omitempty"`
	AmdmntRsn   MandateAmendmentReason1     `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 AmdmntRsn"`
	Mndt        Mandate3                    `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 Mndt"`
	OrgnlMndt   OriginalMandate2Choice      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 OrgnlMndt"`
	SplmtryData []SupplementaryData1        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 SplmtryData,omitempty"`
}

type MandateAmendmentReason1 struct {
	Orgtr    PartyIdentification43 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 Orgtr,omitempty"`
	Rsn      MandateReason1Choice  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 Rsn"`
	AddtlInf []Max105Text          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 AddtlInf,omitempty"`
}

type MandateAmendmentRequestV03 struct {
	GrpHdr            GroupHeader47        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 GrpHdr"`
	UndrlygAmdmntDtls []MandateAmendment3  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 UndrlygAmdmntDtls"`
	SplmtryData       []SupplementaryData1 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 SplmtryData,omitempty"`
}

type MandateOccurrences2 struct {
	SeqTp        SequenceType2Code  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 SeqTp"`
	Frqcy        Frequency6Code     `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 Frqcy,omitempty"`
	Drtn         DatePeriodDetails1 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 Drtn,omitempty"`
	FrstColltnDt ISODate            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 FrstColltnDt,omitempty"`
	FnlColltnDt  ISODate            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 FnlColltnDt,omitempty"`
}

type MandateReason1Choice struct {
	Cd    ExternalMandateReason1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 Cd"`
	Prtry Max35Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 Prtry"`
}

type MandateTypeInformation1 struct {
	SvcLvl    ServiceLevel8Choice    `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 SvcLvl,omitempty"`
	LclInstrm LocalInstrument2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 LclInstrm,omitempty"`
}

// Must be at least 1 items long
type Max105Text string

// Must be at least 1 items long
type Max128Text string

// Must be at least 1 items long
type Max140Text string

// Must be at least 1 items long
type Max16Text string

// Must be at least 1 items long
type Max2048Text string

// Must be at least 1 items long
type Max34Text string

// Must be at least 1 items long
type Max350Text string

// Must be at least 1 items long
type Max35Text string

// Must be at least 1 items long
type Max70Text string

// May be one of DOCT, MIST, MISS, MADM
type NamePrefix1Code string

type OrganisationIdentification8 struct {
	AnyBIC AnyBICIdentifier                     `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 AnyBIC,omitempty"`
	Othr   []GenericOrganisationIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 Othr,omitempty"`
}

type OrganisationIdentificationSchemeName1Choice struct {
	Cd    ExternalOrganisationIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 Cd"`
	Prtry Max35Text                               `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 Prtry"`
}

type OriginalMandate2Choice struct {
	OrgnlMndtId Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 OrgnlMndtId"`
	OrgnlMndt   Mandate1  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 OrgnlMndt"`
}

type OriginalMessageInformation1 struct {
	MsgId   Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 MsgId"`
	MsgNmId Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 MsgNmId"`
	CreDtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 CreDtTm,omitempty"`
}

type Party11Choice struct {
	OrgId  OrganisationIdentification8 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 OrgId"`
	PrvtId PersonIdentification5       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 PrvtId"`
}

type PartyIdentification43 struct {
	Nm        Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 Nm,omitempty"`
	PstlAdr   PostalAddress6  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 PstlAdr,omitempty"`
	Id        Party11Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 Id,omitempty"`
	CtryOfRes CountryCode     `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 CtryOfRes,omitempty"`
	CtctDtls  ContactDetails2 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 CtctDtls,omitempty"`
}

type PersonIdentification5 struct {
	DtAndPlcOfBirth DateAndPlaceOfBirth            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 DtAndPlcOfBirth,omitempty"`
	Othr            []GenericPersonIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 Othr,omitempty"`
}

type PersonIdentificationSchemeName1Choice struct {
	Cd    ExternalPersonIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 Cd"`
	Prtry Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 Prtry"`
}

// Must match the pattern \+[0-9]{1,3}-[0-9()+\-]{1,30}
type PhoneNumber string

type PostalAddress6 struct {
	AdrTp       AddressType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 AdrTp,omitempty"`
	Dept        Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 Dept,omitempty"`
	SubDept     Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 SubDept,omitempty"`
	StrtNm      Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 StrtNm,omitempty"`
	BldgNb      Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 BldgNb,omitempty"`
	PstCd       Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 PstCd,omitempty"`
	TwnNm       Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 TwnNm,omitempty"`
	CtrySubDvsn Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 CtrySubDvsn,omitempty"`
	Ctry        CountryCode      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 Ctry,omitempty"`
	AdrLine     []Max70Text      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 AdrLine,omitempty"`
}

type ReferredDocumentInformation3 struct {
	Tp     ReferredDocumentType2 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 Tp,omitempty"`
	Nb     Max35Text             `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 Nb,omitempty"`
	RltdDt ISODate               `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 RltdDt,omitempty"`
}

type ReferredDocumentType1Choice struct {
	Cd    DocumentType5Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 Cd"`
	Prtry Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 Prtry"`
}

type ReferredDocumentType2 struct {
	CdOrPrtry ReferredDocumentType1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 CdOrPrtry"`
	Issr      Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 Issr,omitempty"`
}

// May be one of RCUR, OOFF
type SequenceType2Code string

type ServiceLevel8Choice struct {
	Cd    ExternalServiceLevel1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 Cd"`
	Prtry Max35Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 Prtry"`
}

type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 Envlp"`
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
