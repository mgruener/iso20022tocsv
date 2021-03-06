// Code generated by main. DO NOT EDIT.

package pain_014_001_01

import (
	"bytes"
	"encoding/xml"
	"time"
)

type AccountIdentification4Choice struct {
	IBAN IBAN2007Identifier            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 IBAN"`
	Othr GenericAccountIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 Othr"`
}

type AccountSchemeName1Choice struct {
	Cd    ExternalAccountIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 Cd"`
	Prtry Max35Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 Prtry"`
}

type ActiveOrHistoricCurrencyAndAmount struct {
	Value float64                      `xml:",chardata"`
	Ccy   ActiveOrHistoricCurrencyCode `xml:"Ccy,attr"`
}

// Must match the pattern [A-Z]{3,3}
type ActiveOrHistoricCurrencyCode string

// May be one of ADDR, PBOX, HOME, BIZZ, MLTO, DLVY
type AddressType2Code string

type AmendmentInformationDetails7 struct {
	OrgnlMndtId      Max35Text                                    `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 OrgnlMndtId,omitempty"`
	OrgnlCdtrSchmeId PartyIdentification43                        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 OrgnlCdtrSchmeId,omitempty"`
	OrgnlCdtrAgt     BranchAndFinancialInstitutionIdentification5 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 OrgnlCdtrAgt,omitempty"`
	OrgnlDbtr        PartyIdentification43                        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 OrgnlDbtr,omitempty"`
	OrgnlDbtrAcct    CashAccount16                                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 OrgnlDbtrAcct,omitempty"`
	OrgnlDbtrAgt     BranchAndFinancialInstitutionIdentification5 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 OrgnlDbtrAgt,omitempty"`
	OrgnlFnlColltnDt ISODate                                      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 OrgnlFnlColltnDt,omitempty"`
	OrgnlFrqcy       Frequency1Code                               `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 OrgnlFrqcy,omitempty"`
}

type AmountType3Choice struct {
	InstdAmt ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 InstdAmt"`
	EqvtAmt  EquivalentAmount2                 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 EqvtAmt"`
}

// Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type AnyBICIdentifier string

// Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type BICFIIdentifier string

type BranchAndFinancialInstitutionIdentification5 struct {
	FinInstnId FinancialInstitutionIdentification8 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 FinInstnId"`
	BrnchId    BranchData2                         `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 BrnchId,omitempty"`
}

type BranchData2 struct {
	Id      Max35Text      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 Id,omitempty"`
	Nm      Max140Text     `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 Nm,omitempty"`
	PstlAdr PostalAddress6 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 PstlAdr,omitempty"`
}

type CashAccount16 struct {
	Id  AccountIdentification4Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 Id"`
	Tp  CashAccountType2             `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 Tp,omitempty"`
	Ccy ActiveOrHistoricCurrencyCode `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 Ccy,omitempty"`
	Nm  Max70Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 Nm,omitempty"`
}

type CashAccountType2 struct {
	Cd    CashAccountType4Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 Cd"`
	Prtry Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 Prtry"`
}

// May be one of CASH, CHAR, COMM, TAXE, CISH, TRAS, SACC, CACC, SVGS, ONDP, MGLD, NREX, MOMA, LOAN, SLRY, ODFT
type CashAccountType4Code string

type CategoryPurpose1Choice struct {
	Cd    ExternalCategoryPurpose1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 Cd"`
	Prtry Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 Prtry"`
}

type ChargesInformation7 struct {
	Amt ActiveOrHistoricCurrencyAndAmount            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 Amt"`
	Pty BranchAndFinancialInstitutionIdentification5 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 Pty"`
}

// May be one of RTGS, RTNS, MPNS, BOOK
type ClearingChannel2Code string

type ClearingSystemIdentification2Choice struct {
	Cd    ExternalClearingSystemIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 Cd"`
	Prtry Max35Text                                 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 Prtry"`
}

type ClearingSystemIdentification3Choice struct {
	Cd    ExternalCashClearingSystem1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 Cd"`
	Prtry Max35Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 Prtry"`
}

type ClearingSystemMemberIdentification2 struct {
	ClrSysId ClearingSystemIdentification2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 ClrSysId,omitempty"`
	MmbId    Max35Text                           `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 MmbId"`
}

type ContactDetails2 struct {
	NmPrfx   NamePrefix1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 NmPrfx,omitempty"`
	Nm       Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 Nm,omitempty"`
	PhneNb   PhoneNumber     `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 PhneNb,omitempty"`
	MobNb    PhoneNumber     `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 MobNb,omitempty"`
	FaxNb    PhoneNumber     `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 FaxNb,omitempty"`
	EmailAdr Max2048Text     `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 EmailAdr,omitempty"`
	Othr     Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 Othr,omitempty"`
}

// Must match the pattern [A-Z]{2,2}
type CountryCode string

// May be one of CRDT, DBIT
type CreditDebitCode string

type CreditorPaymentActivationRequestStatusReportV01 struct {
	GrpHdr            GroupHeader46                 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 GrpHdr"`
	OrgnlGrpInfAndSts OriginalGroupInformation25    `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 OrgnlGrpInfAndSts"`
	OrgnlPmtInfAndSts []OriginalPaymentInformation5 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 OrgnlPmtInfAndSts,omitempty"`
}

type CreditorReferenceInformation2 struct {
	Tp  CreditorReferenceType2 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 Tp,omitempty"`
	Ref Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 Ref,omitempty"`
}

type CreditorReferenceType1Choice struct {
	Cd    DocumentType3Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 Cd"`
	Prtry Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 Prtry"`
}

type CreditorReferenceType2 struct {
	CdOrPrtry CreditorReferenceType1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 CdOrPrtry"`
	Issr      Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 Issr,omitempty"`
}

type DateAndPlaceOfBirth struct {
	BirthDt     ISODate     `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 BirthDt"`
	PrvcOfBirth Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 PrvcOfBirth,omitempty"`
	CityOfBirth Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 CityOfBirth"`
	CtryOfBirth CountryCode `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 CtryOfBirth"`
}

type Document struct {
	CdtrPmtActvtnReqStsRpt CreditorPaymentActivationRequestStatusReportV01 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 CdtrPmtActvtnReqStsRpt"`
}

type DocumentAdjustment1 struct {
	Amt       ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 Amt"`
	CdtDbtInd CreditDebitCode                   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 CdtDbtInd,omitempty"`
	Rsn       Max4Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 Rsn,omitempty"`
	AddtlInf  Max140Text                        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 AddtlInf,omitempty"`
}

// May be one of RADM, RPIN, FXDR, DISP, PUOR, SCOR
type DocumentType3Code string

// May be one of MSIN, CNFA, DNFA, CINV, CREN, DEBN, HIRI, SBIN, CMCN, SOAC, DISP, BOLD, VCHR, AROI, TSUT
type DocumentType5Code string

type EquivalentAmount2 struct {
	Amt      ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 Amt"`
	CcyOfTrf ActiveOrHistoricCurrencyCode      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 CcyOfTrf"`
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

// Must be at least 1 items long
type ExternalStatusReason1Code string

type FinancialIdentificationSchemeName1Choice struct {
	Cd    ExternalFinancialInstitutionIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 Cd"`
	Prtry Max35Text                                       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 Prtry"`
}

type FinancialInstitutionIdentification8 struct {
	BICFI       BICFIIdentifier                     `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 BICFI,omitempty"`
	ClrSysMmbId ClearingSystemMemberIdentification2 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 ClrSysMmbId,omitempty"`
	Nm          Max140Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 Nm,omitempty"`
	PstlAdr     PostalAddress6                      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 PstlAdr,omitempty"`
	Othr        GenericFinancialIdentification1     `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 Othr,omitempty"`
}

// May be one of YEAR, MNTH, QURT, MIAN, WEEK, DAIL, ADHO, INDA
type Frequency1Code string

type GenericAccountIdentification1 struct {
	Id      Max34Text                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 Id"`
	SchmeNm AccountSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 SchmeNm,omitempty"`
	Issr    Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 Issr,omitempty"`
}

type GenericFinancialIdentification1 struct {
	Id      Max35Text                                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 Id"`
	SchmeNm FinancialIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 SchmeNm,omitempty"`
	Issr    Max35Text                                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 Issr,omitempty"`
}

type GenericOrganisationIdentification1 struct {
	Id      Max35Text                                   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 Id"`
	SchmeNm OrganisationIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 SchmeNm,omitempty"`
	Issr    Max35Text                                   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 Issr,omitempty"`
}

type GenericPersonIdentification1 struct {
	Id      Max35Text                             `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 Id"`
	SchmeNm PersonIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 SchmeNm,omitempty"`
	Issr    Max35Text                             `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 Issr,omitempty"`
}

type GroupHeader46 struct {
	MsgId    Max35Text                                    `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 MsgId"`
	CreDtTm  ISODateTime                                  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 CreDtTm"`
	InitgPty PartyIdentification43                        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 InitgPty"`
	DbtrAgt  BranchAndFinancialInstitutionIdentification5 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 DbtrAgt,omitempty"`
	CdtrAgt  BranchAndFinancialInstitutionIdentification5 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 CdtrAgt,omitempty"`
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
	Cd    ExternalLocalInstrument1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 Cd"`
	Prtry Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 Prtry"`
}

type MandateRelatedInformation7 struct {
	MndtId        Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 MndtId,omitempty"`
	DtOfSgntr     ISODate                      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 DtOfSgntr,omitempty"`
	AmdmntInd     bool                         `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 AmdmntInd,omitempty"`
	AmdmntInfDtls AmendmentInformationDetails7 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 AmdmntInfDtls,omitempty"`
	ElctrncSgntr  Max1025Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 ElctrncSgntr,omitempty"`
	FrstColltnDt  ISODate                      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 FrstColltnDt,omitempty"`
	FnlColltnDt   ISODate                      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 FnlColltnDt,omitempty"`
	Frqcy         Frequency1Code               `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 Frqcy,omitempty"`
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

type NumberOfTransactionsPerStatus3 struct {
	DtldNbOfTxs Max15NumericText                 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 DtldNbOfTxs"`
	DtldSts     TransactionIndividualStatus3Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 DtldSts"`
	DtldCtrlSum float64                          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 DtldCtrlSum,omitempty"`
}

type OrganisationIdentification8 struct {
	AnyBIC AnyBICIdentifier                     `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 AnyBIC,omitempty"`
	Othr   []GenericOrganisationIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 Othr,omitempty"`
}

type OrganisationIdentificationSchemeName1Choice struct {
	Cd    ExternalOrganisationIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 Cd"`
	Prtry Max35Text                               `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 Prtry"`
}

type OriginalGroupInformation25 struct {
	OrgnlMsgId    Max35Text                        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 OrgnlMsgId"`
	OrgnlMsgNmId  Max35Text                        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 OrgnlMsgNmId"`
	OrgnlCreDtTm  ISODateTime                      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 OrgnlCreDtTm,omitempty"`
	OrgnlNbOfTxs  Max15NumericText                 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 OrgnlNbOfTxs,omitempty"`
	OrgnlCtrlSum  float64                          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 OrgnlCtrlSum,omitempty"`
	GrpSts        TransactionGroupStatus3Code      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 GrpSts,omitempty"`
	StsRsnInf     []StatusReasonInformation9       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 StsRsnInf,omitempty"`
	NbOfTxsPerSts []NumberOfTransactionsPerStatus3 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 NbOfTxsPerSts,omitempty"`
}

type OriginalPaymentInformation5 struct {
	OrgnlPmtInfId Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 OrgnlPmtInfId"`
	OrgnlNbOfTxs  Max15NumericText                  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 OrgnlNbOfTxs,omitempty"`
	OrgnlCtrlSum  float64                           `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 OrgnlCtrlSum,omitempty"`
	PmtInfSts     TransactionGroupStatus3Code       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 PmtInfSts,omitempty"`
	StsRsnInf     []StatusReasonInformation9        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 StsRsnInf,omitempty"`
	NbOfTxsPerSts []NumberOfTransactionsPerStatus3  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 NbOfTxsPerSts,omitempty"`
	TxInfAndSts   []PaymentTransactionInformation34 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 TxInfAndSts,omitempty"`
}

type OriginalTransactionReference15 struct {
	IntrBkSttlmAmt ActiveOrHistoricCurrencyAndAmount            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 IntrBkSttlmAmt,omitempty"`
	Amt            AmountType3Choice                            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 Amt,omitempty"`
	IntrBkSttlmDt  ISODate                                      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 IntrBkSttlmDt,omitempty"`
	ReqdColltnDt   ISODate                                      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 ReqdColltnDt,omitempty"`
	ReqdExctnDt    ISODate                                      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 ReqdExctnDt,omitempty"`
	CdtrSchmeId    PartyIdentification43                        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 CdtrSchmeId,omitempty"`
	SttlmInf       SettlementInformation16                      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 SttlmInf,omitempty"`
	PmtTpInf       PaymentTypeInformation22                     `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 PmtTpInf,omitempty"`
	PmtMtd         PaymentMethod4Code                           `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 PmtMtd,omitempty"`
	MndtRltdInf    MandateRelatedInformation7                   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 MndtRltdInf,omitempty"`
	RmtInf         RemittanceInformation6                       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 RmtInf,omitempty"`
	UltmtDbtr      PartyIdentification43                        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 UltmtDbtr,omitempty"`
	Dbtr           PartyIdentification43                        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 Dbtr,omitempty"`
	DbtrAcct       CashAccount16                                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 DbtrAcct,omitempty"`
	DbtrAgt        BranchAndFinancialInstitutionIdentification5 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 DbtrAgt,omitempty"`
	CdtrAgt        BranchAndFinancialInstitutionIdentification5 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 CdtrAgt"`
	Cdtr           PartyIdentification43                        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 Cdtr"`
	CdtrAcct       CashAccount16                                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 CdtrAcct,omitempty"`
	UltmtCdtr      PartyIdentification43                        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 UltmtCdtr,omitempty"`
}

type Party11Choice struct {
	OrgId  OrganisationIdentification8 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 OrgId"`
	PrvtId PersonIdentification5       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 PrvtId"`
}

type PartyIdentification43 struct {
	Nm        Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 Nm,omitempty"`
	PstlAdr   PostalAddress6  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 PstlAdr,omitempty"`
	Id        Party11Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 Id,omitempty"`
	CtryOfRes CountryCode     `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 CtryOfRes,omitempty"`
	CtctDtls  ContactDetails2 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 CtctDtls,omitempty"`
}

// May be one of CHK, TRF, DD, TRA
type PaymentMethod4Code string

type PaymentTransactionInformation34 struct {
	StsId           Max35Text                        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 StsId,omitempty"`
	OrgnlInstrId    Max35Text                        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 OrgnlInstrId,omitempty"`
	OrgnlEndToEndId Max35Text                        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 OrgnlEndToEndId,omitempty"`
	TxSts           TransactionIndividualStatus3Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 TxSts,omitempty"`
	StsRsnInf       []StatusReasonInformation9       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 StsRsnInf,omitempty"`
	ChrgsInf        []ChargesInformation7            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 ChrgsInf,omitempty"`
	AccptncDtTm     ISODateTime                      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 AccptncDtTm,omitempty"`
	AcctSvcrRef     Max35Text                        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 AcctSvcrRef,omitempty"`
	ClrSysRef       Max35Text                        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 ClrSysRef,omitempty"`
	OrgnlTxRef      OriginalTransactionReference15   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 OrgnlTxRef,omitempty"`
}

type PaymentTypeInformation22 struct {
	InstrPrty Priority2Code          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 InstrPrty,omitempty"`
	ClrChanl  ClearingChannel2Code   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 ClrChanl,omitempty"`
	SvcLvl    ServiceLevel8Choice    `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 SvcLvl,omitempty"`
	LclInstrm LocalInstrument2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 LclInstrm,omitempty"`
	SeqTp     SequenceType1Code      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 SeqTp,omitempty"`
	CtgyPurp  CategoryPurpose1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 CtgyPurp,omitempty"`
}

type PersonIdentification5 struct {
	DtAndPlcOfBirth DateAndPlaceOfBirth            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 DtAndPlcOfBirth,omitempty"`
	Othr            []GenericPersonIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 Othr,omitempty"`
}

type PersonIdentificationSchemeName1Choice struct {
	Cd    ExternalPersonIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 Cd"`
	Prtry Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 Prtry"`
}

// Must match the pattern \+[0-9]{1,3}-[0-9()+\-]{1,30}
type PhoneNumber string

type PostalAddress6 struct {
	AdrTp       AddressType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 AdrTp,omitempty"`
	Dept        Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 Dept,omitempty"`
	SubDept     Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 SubDept,omitempty"`
	StrtNm      Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 StrtNm,omitempty"`
	BldgNb      Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 BldgNb,omitempty"`
	PstCd       Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 PstCd,omitempty"`
	TwnNm       Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 TwnNm,omitempty"`
	CtrySubDvsn Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 CtrySubDvsn,omitempty"`
	Ctry        CountryCode      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 Ctry,omitempty"`
	AdrLine     []Max70Text      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 AdrLine,omitempty"`
}

// May be one of HIGH, NORM
type Priority2Code string

type ReferredDocumentInformation3 struct {
	Tp     ReferredDocumentType2 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 Tp,omitempty"`
	Nb     Max35Text             `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 Nb,omitempty"`
	RltdDt ISODate               `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 RltdDt,omitempty"`
}

type ReferredDocumentType1Choice struct {
	Cd    DocumentType5Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 Cd"`
	Prtry Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 Prtry"`
}

type ReferredDocumentType2 struct {
	CdOrPrtry ReferredDocumentType1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 CdOrPrtry"`
	Issr      Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 Issr,omitempty"`
}

type RemittanceAmount1 struct {
	DuePyblAmt        ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 DuePyblAmt,omitempty"`
	DscntApldAmt      ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 DscntApldAmt,omitempty"`
	CdtNoteAmt        ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 CdtNoteAmt,omitempty"`
	TaxAmt            ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 TaxAmt,omitempty"`
	AdjstmntAmtAndRsn []DocumentAdjustment1             `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 AdjstmntAmtAndRsn,omitempty"`
	RmtdAmt           ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 RmtdAmt,omitempty"`
}

type RemittanceInformation6 struct {
	Ustrd []Max140Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 Ustrd,omitempty"`
	Strd  []StructuredRemittanceInformation8 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 Strd,omitempty"`
}

// May be one of FRST, RCUR, FNAL, OOFF
type SequenceType1Code string

type ServiceLevel8Choice struct {
	Cd    ExternalServiceLevel1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 Cd"`
	Prtry Max35Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 Prtry"`
}

type SettlementInformation16 struct {
	SttlmMtd             SettlementMethod1Code                        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 SttlmMtd"`
	SttlmAcct            CashAccount16                                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 SttlmAcct,omitempty"`
	ClrSys               ClearingSystemIdentification3Choice          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 ClrSys,omitempty"`
	InstgRmbrsmntAgt     BranchAndFinancialInstitutionIdentification5 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 InstgRmbrsmntAgt,omitempty"`
	InstgRmbrsmntAgtAcct CashAccount16                                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 InstgRmbrsmntAgtAcct,omitempty"`
	InstdRmbrsmntAgt     BranchAndFinancialInstitutionIdentification5 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 InstdRmbrsmntAgt,omitempty"`
	InstdRmbrsmntAgtAcct CashAccount16                                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 InstdRmbrsmntAgtAcct,omitempty"`
	ThrdRmbrsmntAgt      BranchAndFinancialInstitutionIdentification5 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 ThrdRmbrsmntAgt,omitempty"`
	ThrdRmbrsmntAgtAcct  CashAccount16                                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 ThrdRmbrsmntAgtAcct,omitempty"`
}

// May be one of INDA, INGA, COVE, CLRG
type SettlementMethod1Code string

type StatusReason6Choice struct {
	Cd    ExternalStatusReason1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 Cd"`
	Prtry Max35Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 Prtry"`
}

type StatusReasonInformation9 struct {
	Orgtr    PartyIdentification43 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 Orgtr,omitempty"`
	Rsn      StatusReason6Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 Rsn,omitempty"`
	AddtlInf []Max105Text          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 AddtlInf,omitempty"`
}

type StructuredRemittanceInformation8 struct {
	RfrdDocInf  []ReferredDocumentInformation3 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 RfrdDocInf,omitempty"`
	RfrdDocAmt  RemittanceAmount1              `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 RfrdDocAmt,omitempty"`
	CdtrRefInf  CreditorReferenceInformation2  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 CdtrRefInf,omitempty"`
	Invcr       PartyIdentification43          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 Invcr,omitempty"`
	Invcee      PartyIdentification43          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 Invcee,omitempty"`
	AddtlRmtInf []Max140Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 AddtlRmtInf,omitempty"`
}

// May be one of ACTC, RCVD, PART, RJCT, PDNG, ACCP, ACSP, ACSC, ACWC
type TransactionGroupStatus3Code string

// May be one of ACTC, RJCT, PDNG, ACCP, ACSP, ACSC, ACWC
type TransactionIndividualStatus3Code string

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
