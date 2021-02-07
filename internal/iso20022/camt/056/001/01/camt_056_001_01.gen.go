// Code generated by main. DO NOT EDIT.

package camt_056_001_01

import (
	"bytes"
	"encoding/xml"
	"time"
)

type AccountIdentification4Choice struct {
	IBAN IBAN2007Identifier            `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 IBAN"`
	Othr GenericAccountIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 Othr"`
}

type AccountSchemeName1Choice struct {
	Cd    ExternalAccountIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 Cd"`
	Prtry Max35Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 Prtry"`
}

type ActiveOrHistoricCurrencyAndAmount struct {
	Value float64                      `xml:",chardata"`
	Ccy   ActiveOrHistoricCurrencyCode `xml:"Ccy,attr"`
}

// Must match the pattern [A-Z]{3,3}
type ActiveOrHistoricCurrencyCode string

// May be one of ADDR, PBOX, HOME, BIZZ, MLTO, DLVY
type AddressType2Code string

type AmendmentInformationDetails6 struct {
	OrgnlMndtId      Max35Text                                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 OrgnlMndtId,omitempty"`
	OrgnlCdtrSchmeId PartyIdentification32                        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 OrgnlCdtrSchmeId,omitempty"`
	OrgnlCdtrAgt     BranchAndFinancialInstitutionIdentification4 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 OrgnlCdtrAgt,omitempty"`
	OrgnlCdtrAgtAcct CashAccount16                                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 OrgnlCdtrAgtAcct,omitempty"`
	OrgnlDbtr        PartyIdentification32                        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 OrgnlDbtr,omitempty"`
	OrgnlDbtrAcct    CashAccount16                                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 OrgnlDbtrAcct,omitempty"`
	OrgnlDbtrAgt     BranchAndFinancialInstitutionIdentification4 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 OrgnlDbtrAgt,omitempty"`
	OrgnlDbtrAgtAcct CashAccount16                                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 OrgnlDbtrAgtAcct,omitempty"`
	OrgnlFnlColltnDt ISODate                                      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 OrgnlFnlColltnDt,omitempty"`
	OrgnlFrqcy       Frequency1Code                               `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 OrgnlFrqcy,omitempty"`
}

type AmountType3Choice struct {
	InstdAmt ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 InstdAmt"`
	EqvtAmt  EquivalentAmount2                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 EqvtAmt"`
}

// Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type AnyBICIdentifier string

// Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type BICIdentifier string

type BranchAndFinancialInstitutionIdentification4 struct {
	FinInstnId FinancialInstitutionIdentification7 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 FinInstnId"`
	BrnchId    BranchData2                         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 BrnchId,omitempty"`
}

type BranchData2 struct {
	Id      Max35Text      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 Id,omitempty"`
	Nm      Max140Text     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 Nm,omitempty"`
	PstlAdr PostalAddress6 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 PstlAdr,omitempty"`
}

type CancellationReason2Choice struct {
	Cd    CancellationReason4Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 Cd"`
	Prtry Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 Prtry"`
}

// May be one of CUST, DUPL, AGNT, CURR, UPAY, CUTA
type CancellationReason4Code string

type CancellationReasonInformation3 struct {
	Orgtr    PartyIdentification32     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 Orgtr,omitempty"`
	Rsn      CancellationReason2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 Rsn,omitempty"`
	AddtlInf []Max105Text              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 AddtlInf,omitempty"`
}

type Case2 struct {
	Id             Max35Text    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 Id"`
	Cretr          Party7Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 Cretr"`
	ReopCaseIndctn bool         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 ReopCaseIndctn,omitempty"`
}

type CaseAssignment2 struct {
	Id      Max35Text    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 Id"`
	Assgnr  Party7Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 Assgnr"`
	Assgne  Party7Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 Assgne"`
	CreDtTm ISODateTime  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 CreDtTm"`
}

type CashAccount16 struct {
	Id  AccountIdentification4Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 Id"`
	Tp  CashAccountType2             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 Tp,omitempty"`
	Ccy ActiveOrHistoricCurrencyCode `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 Ccy,omitempty"`
	Nm  Max70Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 Nm,omitempty"`
}

type CashAccountType2 struct {
	Cd    CashAccountType4Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 Cd"`
	Prtry Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 Prtry"`
}

// May be one of CASH, CHAR, COMM, TAXE, CISH, TRAS, SACC, CACC, SVGS, ONDP, MGLD, NREX, MOMA, LOAN, SLRY, ODFT
type CashAccountType4Code string

type CategoryPurpose1Choice struct {
	Cd    ExternalCategoryPurpose1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 Cd"`
	Prtry Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 Prtry"`
}

// May be one of RTGS, RTNS, MPNS, BOOK
type ClearingChannel2Code string

type ClearingSystemIdentification2Choice struct {
	Cd    ExternalClearingSystemIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 Cd"`
	Prtry Max35Text                                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 Prtry"`
}

type ClearingSystemIdentification3Choice struct {
	Cd    ExternalCashClearingSystem1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 Cd"`
	Prtry Max35Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 Prtry"`
}

type ClearingSystemMemberIdentification2 struct {
	ClrSysId ClearingSystemIdentification2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 ClrSysId,omitempty"`
	MmbId    Max35Text                           `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 MmbId"`
}

type ContactDetails2 struct {
	NmPrfx   NamePrefix1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 NmPrfx,omitempty"`
	Nm       Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 Nm,omitempty"`
	PhneNb   PhoneNumber     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 PhneNb,omitempty"`
	MobNb    PhoneNumber     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 MobNb,omitempty"`
	FaxNb    PhoneNumber     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 FaxNb,omitempty"`
	EmailAdr Max2048Text     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 EmailAdr,omitempty"`
	Othr     Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 Othr,omitempty"`
}

type ControlData1 struct {
	NbOfTxs Max15NumericText `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 NbOfTxs"`
	CtrlSum float64          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 CtrlSum,omitempty"`
}

// Must match the pattern [A-Z]{2,2}
type CountryCode string

// May be one of CRDT, DBIT
type CreditDebitCode string

type CreditorReferenceInformation2 struct {
	Tp  CreditorReferenceType2 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 Tp,omitempty"`
	Ref Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 Ref,omitempty"`
}

type CreditorReferenceType1Choice struct {
	Cd    DocumentType3Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 Cd"`
	Prtry Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 Prtry"`
}

type CreditorReferenceType2 struct {
	CdOrPrtry CreditorReferenceType1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 CdOrPrtry"`
	Issr      Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 Issr,omitempty"`
}

type DateAndPlaceOfBirth struct {
	BirthDt     ISODate     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 BirthDt"`
	PrvcOfBirth Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 PrvcOfBirth,omitempty"`
	CityOfBirth Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 CityOfBirth"`
	CtryOfBirth CountryCode `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 CtryOfBirth"`
}

type Document struct {
	FIToFIPmtCxlReq FIToFIPaymentCancellationRequestV01 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 FIToFIPmtCxlReq"`
}

type DocumentAdjustment1 struct {
	Amt       ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 Amt"`
	CdtDbtInd CreditDebitCode                   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 CdtDbtInd,omitempty"`
	Rsn       Max4Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 Rsn,omitempty"`
	AddtlInf  Max140Text                        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 AddtlInf,omitempty"`
}

// May be one of RADM, RPIN, FXDR, DISP, PUOR, SCOR
type DocumentType3Code string

// May be one of MSIN, CNFA, DNFA, CINV, CREN, DEBN, HIRI, SBIN, CMCN, SOAC, DISP, BOLD, VCHR, AROI, TSUT
type DocumentType5Code string

type EquivalentAmount2 struct {
	Amt      ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 Amt"`
	CcyOfTrf ActiveOrHistoricCurrencyCode      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 CcyOfTrf"`
}

// Must be at least 1 items long
type ExternalAccountIdentification1Code string

// Must be at least 1 items long
type ExternalCashClearingSystem1Code string

// Must be at least 1 items long
type ExternalCategoryPurpose1Code string

// Must be at least 1 items long
type ExternalClearingSystemIdentification1Code string

// Must be at least 1 items long
type ExternalFinancialInstitutionIdentification1Code string

// Must be at least 1 items long
type ExternalLocalInstrument1Code string

// Must be at least 1 items long
type ExternalOrganisationIdentification1Code string

// Must be at least 1 items long
type ExternalPersonIdentification1Code string

// Must be at least 1 items long
type ExternalServiceLevel1Code string

type FIToFIPaymentCancellationRequestV01 struct {
	Assgnmt  CaseAssignment2          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 Assgnmt"`
	Case     Case2                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 Case,omitempty"`
	CtrlData ControlData1             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 CtrlData,omitempty"`
	Undrlyg  []UnderlyingTransaction2 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 Undrlyg"`
}

type FinancialIdentificationSchemeName1Choice struct {
	Cd    ExternalFinancialInstitutionIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 Cd"`
	Prtry Max35Text                                       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 Prtry"`
}

type FinancialInstitutionIdentification7 struct {
	BIC         BICIdentifier                       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 BIC,omitempty"`
	ClrSysMmbId ClearingSystemMemberIdentification2 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 ClrSysMmbId,omitempty"`
	Nm          Max140Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 Nm,omitempty"`
	PstlAdr     PostalAddress6                      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 PstlAdr,omitempty"`
	Othr        GenericFinancialIdentification1     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 Othr,omitempty"`
}

// May be one of YEAR, MNTH, QURT, MIAN, WEEK, DAIL, ADHO, INDA
type Frequency1Code string

type GenericAccountIdentification1 struct {
	Id      Max34Text                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 Id"`
	SchmeNm AccountSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 SchmeNm,omitempty"`
	Issr    Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 Issr,omitempty"`
}

type GenericFinancialIdentification1 struct {
	Id      Max35Text                                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 Id"`
	SchmeNm FinancialIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 SchmeNm,omitempty"`
	Issr    Max35Text                                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 Issr,omitempty"`
}

type GenericOrganisationIdentification1 struct {
	Id      Max35Text                                   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 Id"`
	SchmeNm OrganisationIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 SchmeNm,omitempty"`
	Issr    Max35Text                                   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 Issr,omitempty"`
}

type GenericPersonIdentification1 struct {
	Id      Max35Text                             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 Id"`
	SchmeNm PersonIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 SchmeNm,omitempty"`
	Issr    Max35Text                             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 Issr,omitempty"`
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
	Cd    ExternalLocalInstrument1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 Cd"`
	Prtry Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 Prtry"`
}

type MandateRelatedInformation6 struct {
	MndtId        Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 MndtId,omitempty"`
	DtOfSgntr     ISODate                      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 DtOfSgntr,omitempty"`
	AmdmntInd     bool                         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 AmdmntInd,omitempty"`
	AmdmntInfDtls AmendmentInformationDetails6 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 AmdmntInfDtls,omitempty"`
	ElctrncSgntr  Max1025Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 ElctrncSgntr,omitempty"`
	FrstColltnDt  ISODate                      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 FrstColltnDt,omitempty"`
	FnlColltnDt   ISODate                      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 FnlColltnDt,omitempty"`
	Frqcy         Frequency1Code               `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 Frqcy,omitempty"`
}

// Must be at least 1 items long
type Max1025Text string

// Must be at least 1 items long
type Max105Text string

// Must be at least 1 items long
type Max140Text string

// Must match the pattern [0-9]{1,15}
type Max15NumericText string

// Must be at least 1 items long
type Max16Text string

// Must be at least 1 items long
type Max2048Text string

// Must be at least 1 items long
type Max34Text string

// Must be at least 1 items long
type Max35Text string

// Must be at least 1 items long
type Max4Text string

// Must be at least 1 items long
type Max70Text string

// May be one of DOCT, MIST, MISS, MADM
type NamePrefix1Code string

type OrganisationIdentification4 struct {
	BICOrBEI AnyBICIdentifier                     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 BICOrBEI,omitempty"`
	Othr     []GenericOrganisationIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 Othr,omitempty"`
}

type OrganisationIdentificationSchemeName1Choice struct {
	Cd    ExternalOrganisationIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 Cd"`
	Prtry Max35Text                               `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 Prtry"`
}

type OriginalGroupInformation23 struct {
	GrpCxlId     Max35Text                        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 GrpCxlId,omitempty"`
	Case         Case2                            `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 Case,omitempty"`
	OrgnlMsgId   Max35Text                        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 OrgnlMsgId"`
	OrgnlMsgNmId Max35Text                        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 OrgnlMsgNmId"`
	OrgnlCreDtTm ISODateTime                      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 OrgnlCreDtTm,omitempty"`
	NbOfTxs      Max15NumericText                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 NbOfTxs,omitempty"`
	CtrlSum      float64                          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 CtrlSum,omitempty"`
	GrpCxl       bool                             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 GrpCxl,omitempty"`
	CxlRsnInf    []CancellationReasonInformation3 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 CxlRsnInf,omitempty"`
}

type OriginalGroupInformation3 struct {
	OrgnlMsgId   Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 OrgnlMsgId"`
	OrgnlMsgNmId Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 OrgnlMsgNmId"`
	OrgnlCreDtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 OrgnlCreDtTm,omitempty"`
}

type OriginalTransactionReference13 struct {
	IntrBkSttlmAmt ActiveOrHistoricCurrencyAndAmount            `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 IntrBkSttlmAmt,omitempty"`
	Amt            AmountType3Choice                            `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 Amt,omitempty"`
	IntrBkSttlmDt  ISODate                                      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 IntrBkSttlmDt,omitempty"`
	ReqdColltnDt   ISODate                                      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 ReqdColltnDt,omitempty"`
	ReqdExctnDt    ISODate                                      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 ReqdExctnDt,omitempty"`
	CdtrSchmeId    PartyIdentification32                        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 CdtrSchmeId,omitempty"`
	SttlmInf       SettlementInformation13                      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 SttlmInf,omitempty"`
	PmtTpInf       PaymentTypeInformation22                     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 PmtTpInf,omitempty"`
	PmtMtd         PaymentMethod4Code                           `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 PmtMtd,omitempty"`
	MndtRltdInf    MandateRelatedInformation6                   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 MndtRltdInf,omitempty"`
	RmtInf         RemittanceInformation5                       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 RmtInf,omitempty"`
	UltmtDbtr      PartyIdentification32                        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 UltmtDbtr,omitempty"`
	Dbtr           PartyIdentification32                        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 Dbtr,omitempty"`
	DbtrAcct       CashAccount16                                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 DbtrAcct,omitempty"`
	DbtrAgt        BranchAndFinancialInstitutionIdentification4 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 DbtrAgt,omitempty"`
	DbtrAgtAcct    CashAccount16                                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 DbtrAgtAcct,omitempty"`
	CdtrAgt        BranchAndFinancialInstitutionIdentification4 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 CdtrAgt,omitempty"`
	CdtrAgtAcct    CashAccount16                                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 CdtrAgtAcct,omitempty"`
	Cdtr           PartyIdentification32                        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 Cdtr,omitempty"`
	CdtrAcct       CashAccount16                                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 CdtrAcct,omitempty"`
	UltmtCdtr      PartyIdentification32                        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 UltmtCdtr,omitempty"`
}

type Party6Choice struct {
	OrgId  OrganisationIdentification4 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 OrgId"`
	PrvtId PersonIdentification5       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 PrvtId"`
}

type Party7Choice struct {
	Pty PartyIdentification32                        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 Pty"`
	Agt BranchAndFinancialInstitutionIdentification4 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 Agt"`
}

type PartyIdentification32 struct {
	Nm        Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 Nm,omitempty"`
	PstlAdr   PostalAddress6  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 PstlAdr,omitempty"`
	Id        Party6Choice    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 Id,omitempty"`
	CtryOfRes CountryCode     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 CtryOfRes,omitempty"`
	CtctDtls  ContactDetails2 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 CtctDtls,omitempty"`
}

// May be one of CHK, TRF, DD, TRA
type PaymentMethod4Code string

type PaymentTransactionInformation31 struct {
	CxlId               Max35Text                                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 CxlId,omitempty"`
	Case                Case2                                        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 Case,omitempty"`
	OrgnlGrpInf         OriginalGroupInformation3                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 OrgnlGrpInf,omitempty"`
	OrgnlInstrId        Max35Text                                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 OrgnlInstrId,omitempty"`
	OrgnlEndToEndId     Max35Text                                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 OrgnlEndToEndId,omitempty"`
	OrgnlTxId           Max35Text                                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 OrgnlTxId,omitempty"`
	OrgnlClrSysRef      Max35Text                                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 OrgnlClrSysRef,omitempty"`
	OrgnlIntrBkSttlmAmt ActiveOrHistoricCurrencyAndAmount            `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 OrgnlIntrBkSttlmAmt,omitempty"`
	OrgnlIntrBkSttlmDt  ISODate                                      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 OrgnlIntrBkSttlmDt,omitempty"`
	Assgnr              BranchAndFinancialInstitutionIdentification4 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 Assgnr,omitempty"`
	Assgne              BranchAndFinancialInstitutionIdentification4 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 Assgne,omitempty"`
	CxlRsnInf           []CancellationReasonInformation3             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 CxlRsnInf,omitempty"`
	OrgnlTxRef          OriginalTransactionReference13               `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 OrgnlTxRef,omitempty"`
}

type PaymentTypeInformation22 struct {
	InstrPrty Priority2Code          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 InstrPrty,omitempty"`
	ClrChanl  ClearingChannel2Code   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 ClrChanl,omitempty"`
	SvcLvl    ServiceLevel8Choice    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 SvcLvl,omitempty"`
	LclInstrm LocalInstrument2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 LclInstrm,omitempty"`
	SeqTp     SequenceType1Code      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 SeqTp,omitempty"`
	CtgyPurp  CategoryPurpose1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 CtgyPurp,omitempty"`
}

type PersonIdentification5 struct {
	DtAndPlcOfBirth DateAndPlaceOfBirth            `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 DtAndPlcOfBirth,omitempty"`
	Othr            []GenericPersonIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 Othr,omitempty"`
}

type PersonIdentificationSchemeName1Choice struct {
	Cd    ExternalPersonIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 Cd"`
	Prtry Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 Prtry"`
}

// Must match the pattern \+[0-9]{1,3}-[0-9()+\-]{1,30}
type PhoneNumber string

type PostalAddress6 struct {
	AdrTp       AddressType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 AdrTp,omitempty"`
	Dept        Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 Dept,omitempty"`
	SubDept     Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 SubDept,omitempty"`
	StrtNm      Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 StrtNm,omitempty"`
	BldgNb      Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 BldgNb,omitempty"`
	PstCd       Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 PstCd,omitempty"`
	TwnNm       Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 TwnNm,omitempty"`
	CtrySubDvsn Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 CtrySubDvsn,omitempty"`
	Ctry        CountryCode      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 Ctry,omitempty"`
	AdrLine     []Max70Text      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 AdrLine,omitempty"`
}

// May be one of HIGH, NORM
type Priority2Code string

type ReferredDocumentInformation3 struct {
	Tp     ReferredDocumentType2 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 Tp,omitempty"`
	Nb     Max35Text             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 Nb,omitempty"`
	RltdDt ISODate               `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 RltdDt,omitempty"`
}

type ReferredDocumentType1Choice struct {
	Cd    DocumentType5Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 Cd"`
	Prtry Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 Prtry"`
}

type ReferredDocumentType2 struct {
	CdOrPrtry ReferredDocumentType1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 CdOrPrtry"`
	Issr      Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 Issr,omitempty"`
}

type RemittanceAmount1 struct {
	DuePyblAmt        ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 DuePyblAmt,omitempty"`
	DscntApldAmt      ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 DscntApldAmt,omitempty"`
	CdtNoteAmt        ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 CdtNoteAmt,omitempty"`
	TaxAmt            ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 TaxAmt,omitempty"`
	AdjstmntAmtAndRsn []DocumentAdjustment1             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 AdjstmntAmtAndRsn,omitempty"`
	RmtdAmt           ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 RmtdAmt,omitempty"`
}

type RemittanceInformation5 struct {
	Ustrd []Max140Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 Ustrd,omitempty"`
	Strd  []StructuredRemittanceInformation7 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 Strd,omitempty"`
}

// May be one of FRST, RCUR, FNAL, OOFF
type SequenceType1Code string

type ServiceLevel8Choice struct {
	Cd    ExternalServiceLevel1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 Cd"`
	Prtry Max35Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 Prtry"`
}

type SettlementInformation13 struct {
	SttlmMtd             SettlementMethod1Code                        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 SttlmMtd"`
	SttlmAcct            CashAccount16                                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 SttlmAcct,omitempty"`
	ClrSys               ClearingSystemIdentification3Choice          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 ClrSys,omitempty"`
	InstgRmbrsmntAgt     BranchAndFinancialInstitutionIdentification4 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 InstgRmbrsmntAgt,omitempty"`
	InstgRmbrsmntAgtAcct CashAccount16                                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 InstgRmbrsmntAgtAcct,omitempty"`
	InstdRmbrsmntAgt     BranchAndFinancialInstitutionIdentification4 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 InstdRmbrsmntAgt,omitempty"`
	InstdRmbrsmntAgtAcct CashAccount16                                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 InstdRmbrsmntAgtAcct,omitempty"`
	ThrdRmbrsmntAgt      BranchAndFinancialInstitutionIdentification4 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 ThrdRmbrsmntAgt,omitempty"`
	ThrdRmbrsmntAgtAcct  CashAccount16                                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 ThrdRmbrsmntAgtAcct,omitempty"`
}

// May be one of INDA, INGA, COVE, CLRG
type SettlementMethod1Code string

type StructuredRemittanceInformation7 struct {
	RfrdDocInf  []ReferredDocumentInformation3 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 RfrdDocInf,omitempty"`
	RfrdDocAmt  RemittanceAmount1              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 RfrdDocAmt,omitempty"`
	CdtrRefInf  CreditorReferenceInformation2  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 CdtrRefInf,omitempty"`
	Invcr       PartyIdentification32          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 Invcr,omitempty"`
	Invcee      PartyIdentification32          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 Invcee,omitempty"`
	AddtlRmtInf []Max140Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 AddtlRmtInf,omitempty"`
}

type UnderlyingTransaction2 struct {
	OrgnlGrpInfAndCxl OriginalGroupInformation23        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 OrgnlGrpInfAndCxl,omitempty"`
	TxInf             []PaymentTransactionInformation31 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 TxInf,omitempty"`
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