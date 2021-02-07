// Code generated by main. DO NOT EDIT.

package pacs_004_001_02

import (
	"bytes"
	"encoding/xml"
	"time"
)

type AccountIdentification4Choice struct {
	IBAN IBAN2007Identifier            `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 IBAN"`
	Othr GenericAccountIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 Othr"`
}

type AccountSchemeName1Choice struct {
	Cd    ExternalAccountIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 Cd"`
	Prtry Max35Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 Prtry"`
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

type AmendmentInformationDetails6 struct {
	OrgnlMndtId      Max35Text                                    `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 OrgnlMndtId,omitempty"`
	OrgnlCdtrSchmeId PartyIdentification32                        `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 OrgnlCdtrSchmeId,omitempty"`
	OrgnlCdtrAgt     BranchAndFinancialInstitutionIdentification4 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 OrgnlCdtrAgt,omitempty"`
	OrgnlCdtrAgtAcct CashAccount16                                `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 OrgnlCdtrAgtAcct,omitempty"`
	OrgnlDbtr        PartyIdentification32                        `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 OrgnlDbtr,omitempty"`
	OrgnlDbtrAcct    CashAccount16                                `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 OrgnlDbtrAcct,omitempty"`
	OrgnlDbtrAgt     BranchAndFinancialInstitutionIdentification4 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 OrgnlDbtrAgt,omitempty"`
	OrgnlDbtrAgtAcct CashAccount16                                `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 OrgnlDbtrAgtAcct,omitempty"`
	OrgnlFnlColltnDt ISODate                                      `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 OrgnlFnlColltnDt,omitempty"`
	OrgnlFrqcy       Frequency1Code                               `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 OrgnlFrqcy,omitempty"`
}

type AmountType3Choice struct {
	InstdAmt ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 InstdAmt"`
	EqvtAmt  EquivalentAmount2                 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 EqvtAmt"`
}

// Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type AnyBICIdentifier string

type Authorisation1Choice struct {
	Cd    Authorisation1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 Cd"`
	Prtry Max128Text         `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 Prtry"`
}

// May be one of AUTH, FDET, FSUM, ILEV
type Authorisation1Code string

// Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type BICIdentifier string

type BranchAndFinancialInstitutionIdentification4 struct {
	FinInstnId FinancialInstitutionIdentification7 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 FinInstnId"`
	BrnchId    BranchData2                         `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 BrnchId,omitempty"`
}

type BranchData2 struct {
	Id      Max35Text      `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 Id,omitempty"`
	Nm      Max140Text     `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 Nm,omitempty"`
	PstlAdr PostalAddress6 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 PstlAdr,omitempty"`
}

type CashAccount16 struct {
	Id  AccountIdentification4Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 Id"`
	Tp  CashAccountType2             `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 Tp,omitempty"`
	Ccy ActiveOrHistoricCurrencyCode `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 Ccy,omitempty"`
	Nm  Max70Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 Nm,omitempty"`
}

type CashAccountType2 struct {
	Cd    CashAccountType4Code `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 Cd"`
	Prtry Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 Prtry"`
}

// May be one of CASH, CHAR, COMM, TAXE, CISH, TRAS, SACC, CACC, SVGS, ONDP, MGLD, NREX, MOMA, LOAN, SLRY, ODFT
type CashAccountType4Code string

type CategoryPurpose1Choice struct {
	Cd    ExternalCategoryPurpose1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 Cd"`
	Prtry Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 Prtry"`
}

// May be one of DEBT, CRED, SHAR, SLEV
type ChargeBearerType1Code string

type ChargesInformation5 struct {
	Amt ActiveOrHistoricCurrencyAndAmount            `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 Amt"`
	Pty BranchAndFinancialInstitutionIdentification4 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 Pty"`
}

// May be one of RTGS, RTNS, MPNS, BOOK
type ClearingChannel2Code string

type ClearingSystemIdentification2Choice struct {
	Cd    ExternalClearingSystemIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 Cd"`
	Prtry Max35Text                                 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 Prtry"`
}

type ClearingSystemIdentification3Choice struct {
	Cd    ExternalCashClearingSystem1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 Cd"`
	Prtry Max35Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 Prtry"`
}

type ClearingSystemMemberIdentification2 struct {
	ClrSysId ClearingSystemIdentification2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 ClrSysId,omitempty"`
	MmbId    Max35Text                           `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 MmbId"`
}

type ContactDetails2 struct {
	NmPrfx   NamePrefix1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 NmPrfx,omitempty"`
	Nm       Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 Nm,omitempty"`
	PhneNb   PhoneNumber     `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 PhneNb,omitempty"`
	MobNb    PhoneNumber     `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 MobNb,omitempty"`
	FaxNb    PhoneNumber     `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 FaxNb,omitempty"`
	EmailAdr Max2048Text     `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 EmailAdr,omitempty"`
	Othr     Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 Othr,omitempty"`
}

// Must match the pattern [A-Z]{2,2}
type CountryCode string

// May be one of CRDT, DBIT
type CreditDebitCode string

type CreditorReferenceInformation2 struct {
	Tp  CreditorReferenceType2 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 Tp,omitempty"`
	Ref Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 Ref,omitempty"`
}

type CreditorReferenceType1Choice struct {
	Cd    DocumentType3Code `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 Cd"`
	Prtry Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 Prtry"`
}

type CreditorReferenceType2 struct {
	CdOrPrtry CreditorReferenceType1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 CdOrPrtry"`
	Issr      Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 Issr,omitempty"`
}

type DateAndPlaceOfBirth struct {
	BirthDt     ISODate     `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 BirthDt"`
	PrvcOfBirth Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 PrvcOfBirth,omitempty"`
	CityOfBirth Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 CityOfBirth"`
	CtryOfBirth CountryCode `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 CtryOfBirth"`
}

type Document struct {
	PmtRtr PaymentReturnV02 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 PmtRtr"`
}

type DocumentAdjustment1 struct {
	Amt       ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 Amt"`
	CdtDbtInd CreditDebitCode                   `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 CdtDbtInd,omitempty"`
	Rsn       Max4Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 Rsn,omitempty"`
	AddtlInf  Max140Text                        `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 AddtlInf,omitempty"`
}

// May be one of RADM, RPIN, FXDR, DISP, PUOR, SCOR
type DocumentType3Code string

// May be one of MSIN, CNFA, DNFA, CINV, CREN, DEBN, HIRI, SBIN, CMCN, SOAC, DISP, BOLD, VCHR, AROI, TSUT
type DocumentType5Code string

type EquivalentAmount2 struct {
	Amt      ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 Amt"`
	CcyOfTrf ActiveOrHistoricCurrencyCode      `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 CcyOfTrf"`
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
type ExternalReturnReason1Code string

// Must be at least 1 items long
type ExternalServiceLevel1Code string

type FinancialIdentificationSchemeName1Choice struct {
	Cd    ExternalFinancialInstitutionIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 Cd"`
	Prtry Max35Text                                       `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 Prtry"`
}

type FinancialInstitutionIdentification7 struct {
	BIC         BICIdentifier                       `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 BIC,omitempty"`
	ClrSysMmbId ClearingSystemMemberIdentification2 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 ClrSysMmbId,omitempty"`
	Nm          Max140Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 Nm,omitempty"`
	PstlAdr     PostalAddress6                      `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 PstlAdr,omitempty"`
	Othr        GenericFinancialIdentification1     `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 Othr,omitempty"`
}

// May be one of YEAR, MNTH, QURT, MIAN, WEEK, DAIL, ADHO, INDA
type Frequency1Code string

type GenericAccountIdentification1 struct {
	Id      Max34Text                `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 Id"`
	SchmeNm AccountSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 SchmeNm,omitempty"`
	Issr    Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 Issr,omitempty"`
}

type GenericFinancialIdentification1 struct {
	Id      Max35Text                                `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 Id"`
	SchmeNm FinancialIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 SchmeNm,omitempty"`
	Issr    Max35Text                                `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 Issr,omitempty"`
}

type GenericOrganisationIdentification1 struct {
	Id      Max35Text                                   `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 Id"`
	SchmeNm OrganisationIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 SchmeNm,omitempty"`
	Issr    Max35Text                                   `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 Issr,omitempty"`
}

type GenericPersonIdentification1 struct {
	Id      Max35Text                             `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 Id"`
	SchmeNm PersonIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 SchmeNm,omitempty"`
	Issr    Max35Text                             `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 Issr,omitempty"`
}

type GroupHeader38 struct {
	MsgId                 Max35Text                                    `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 MsgId"`
	CreDtTm               ISODateTime                                  `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 CreDtTm"`
	Authstn               []Authorisation1Choice                       `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 Authstn,omitempty"`
	BtchBookg             bool                                         `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 BtchBookg,omitempty"`
	NbOfTxs               Max15NumericText                             `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 NbOfTxs"`
	CtrlSum               float64                                      `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 CtrlSum,omitempty"`
	GrpRtr                bool                                         `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 GrpRtr,omitempty"`
	TtlRtrdIntrBkSttlmAmt ActiveCurrencyAndAmount                      `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 TtlRtrdIntrBkSttlmAmt,omitempty"`
	IntrBkSttlmDt         ISODate                                      `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 IntrBkSttlmDt,omitempty"`
	SttlmInf              SettlementInformation13                      `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 SttlmInf"`
	InstgAgt              BranchAndFinancialInstitutionIdentification4 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 InstgAgt,omitempty"`
	InstdAgt              BranchAndFinancialInstitutionIdentification4 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 InstdAgt,omitempty"`
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
	Cd    ExternalLocalInstrument1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 Cd"`
	Prtry Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 Prtry"`
}

type MandateRelatedInformation6 struct {
	MndtId        Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 MndtId,omitempty"`
	DtOfSgntr     ISODate                      `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 DtOfSgntr,omitempty"`
	AmdmntInd     bool                         `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 AmdmntInd,omitempty"`
	AmdmntInfDtls AmendmentInformationDetails6 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 AmdmntInfDtls,omitempty"`
	ElctrncSgntr  Max1025Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 ElctrncSgntr,omitempty"`
	FrstColltnDt  ISODate                      `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 FrstColltnDt,omitempty"`
	FnlColltnDt   ISODate                      `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 FnlColltnDt,omitempty"`
	Frqcy         Frequency1Code               `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 Frqcy,omitempty"`
}

// Must be at least 1 items long
type Max1025Text string

// Must be at least 1 items long
type Max105Text string

// Must be at least 1 items long
type Max128Text string

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
	BICOrBEI AnyBICIdentifier                     `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 BICOrBEI,omitempty"`
	Othr     []GenericOrganisationIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 Othr,omitempty"`
}

type OrganisationIdentificationSchemeName1Choice struct {
	Cd    ExternalOrganisationIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 Cd"`
	Prtry Max35Text                               `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 Prtry"`
}

type OriginalGroupInformation21 struct {
	OrgnlMsgId   Max35Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 OrgnlMsgId"`
	OrgnlMsgNmId Max35Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 OrgnlMsgNmId"`
	OrgnlCreDtTm ISODateTime                `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 OrgnlCreDtTm,omitempty"`
	RtrRsnInf    []ReturnReasonInformation9 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 RtrRsnInf,omitempty"`
}

type OriginalGroupInformation3 struct {
	OrgnlMsgId   Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 OrgnlMsgId"`
	OrgnlMsgNmId Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 OrgnlMsgNmId"`
	OrgnlCreDtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 OrgnlCreDtTm,omitempty"`
}

type OriginalTransactionReference13 struct {
	IntrBkSttlmAmt ActiveOrHistoricCurrencyAndAmount            `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 IntrBkSttlmAmt,omitempty"`
	Amt            AmountType3Choice                            `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 Amt,omitempty"`
	IntrBkSttlmDt  ISODate                                      `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 IntrBkSttlmDt,omitempty"`
	ReqdColltnDt   ISODate                                      `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 ReqdColltnDt,omitempty"`
	ReqdExctnDt    ISODate                                      `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 ReqdExctnDt,omitempty"`
	CdtrSchmeId    PartyIdentification32                        `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 CdtrSchmeId,omitempty"`
	SttlmInf       SettlementInformation13                      `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 SttlmInf,omitempty"`
	PmtTpInf       PaymentTypeInformation22                     `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 PmtTpInf,omitempty"`
	PmtMtd         PaymentMethod4Code                           `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 PmtMtd,omitempty"`
	MndtRltdInf    MandateRelatedInformation6                   `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 MndtRltdInf,omitempty"`
	RmtInf         RemittanceInformation5                       `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 RmtInf,omitempty"`
	UltmtDbtr      PartyIdentification32                        `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 UltmtDbtr,omitempty"`
	Dbtr           PartyIdentification32                        `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 Dbtr,omitempty"`
	DbtrAcct       CashAccount16                                `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 DbtrAcct,omitempty"`
	DbtrAgt        BranchAndFinancialInstitutionIdentification4 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 DbtrAgt,omitempty"`
	DbtrAgtAcct    CashAccount16                                `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 DbtrAgtAcct,omitempty"`
	CdtrAgt        BranchAndFinancialInstitutionIdentification4 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 CdtrAgt,omitempty"`
	CdtrAgtAcct    CashAccount16                                `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 CdtrAgtAcct,omitempty"`
	Cdtr           PartyIdentification32                        `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 Cdtr,omitempty"`
	CdtrAcct       CashAccount16                                `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 CdtrAcct,omitempty"`
	UltmtCdtr      PartyIdentification32                        `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 UltmtCdtr,omitempty"`
}

type Party6Choice struct {
	OrgId  OrganisationIdentification4 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 OrgId"`
	PrvtId PersonIdentification5       `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 PrvtId"`
}

type PartyIdentification32 struct {
	Nm        Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 Nm,omitempty"`
	PstlAdr   PostalAddress6  `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 PstlAdr,omitempty"`
	Id        Party6Choice    `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 Id,omitempty"`
	CtryOfRes CountryCode     `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 CtryOfRes,omitempty"`
	CtctDtls  ContactDetails2 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 CtctDtls,omitempty"`
}

// May be one of CHK, TRF, DD, TRA
type PaymentMethod4Code string

type PaymentReturnV02 struct {
	GrpHdr      GroupHeader38                     `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 GrpHdr"`
	OrgnlGrpInf OriginalGroupInformation21        `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 OrgnlGrpInf,omitempty"`
	TxInf       []PaymentTransactionInformation27 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 TxInf,omitempty"`
}

type PaymentTransactionInformation27 struct {
	RtrId               Max35Text                                    `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 RtrId,omitempty"`
	OrgnlGrpInf         OriginalGroupInformation3                    `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 OrgnlGrpInf,omitempty"`
	OrgnlInstrId        Max35Text                                    `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 OrgnlInstrId,omitempty"`
	OrgnlEndToEndId     Max35Text                                    `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 OrgnlEndToEndId,omitempty"`
	OrgnlTxId           Max35Text                                    `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 OrgnlTxId,omitempty"`
	OrgnlClrSysRef      Max35Text                                    `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 OrgnlClrSysRef,omitempty"`
	OrgnlIntrBkSttlmAmt ActiveOrHistoricCurrencyAndAmount            `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 OrgnlIntrBkSttlmAmt,omitempty"`
	RtrdIntrBkSttlmAmt  ActiveCurrencyAndAmount                      `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 RtrdIntrBkSttlmAmt"`
	IntrBkSttlmDt       ISODate                                      `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 IntrBkSttlmDt,omitempty"`
	RtrdInstdAmt        ActiveOrHistoricCurrencyAndAmount            `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 RtrdInstdAmt,omitempty"`
	XchgRate            float64                                      `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 XchgRate,omitempty"`
	CompstnAmt          ActiveOrHistoricCurrencyAndAmount            `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 CompstnAmt,omitempty"`
	ChrgBr              ChargeBearerType1Code                        `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 ChrgBr,omitempty"`
	ChrgsInf            []ChargesInformation5                        `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 ChrgsInf,omitempty"`
	InstgAgt            BranchAndFinancialInstitutionIdentification4 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 InstgAgt,omitempty"`
	InstdAgt            BranchAndFinancialInstitutionIdentification4 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 InstdAgt,omitempty"`
	RtrRsnInf           []ReturnReasonInformation9                   `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 RtrRsnInf,omitempty"`
	OrgnlTxRef          OriginalTransactionReference13               `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 OrgnlTxRef,omitempty"`
}

type PaymentTypeInformation22 struct {
	InstrPrty Priority2Code          `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 InstrPrty,omitempty"`
	ClrChanl  ClearingChannel2Code   `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 ClrChanl,omitempty"`
	SvcLvl    ServiceLevel8Choice    `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 SvcLvl,omitempty"`
	LclInstrm LocalInstrument2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 LclInstrm,omitempty"`
	SeqTp     SequenceType1Code      `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 SeqTp,omitempty"`
	CtgyPurp  CategoryPurpose1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 CtgyPurp,omitempty"`
}

type PersonIdentification5 struct {
	DtAndPlcOfBirth DateAndPlaceOfBirth            `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 DtAndPlcOfBirth,omitempty"`
	Othr            []GenericPersonIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 Othr,omitempty"`
}

type PersonIdentificationSchemeName1Choice struct {
	Cd    ExternalPersonIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 Cd"`
	Prtry Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 Prtry"`
}

// Must match the pattern \+[0-9]{1,3}-[0-9()+\-]{1,30}
type PhoneNumber string

type PostalAddress6 struct {
	AdrTp       AddressType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 AdrTp,omitempty"`
	Dept        Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 Dept,omitempty"`
	SubDept     Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 SubDept,omitempty"`
	StrtNm      Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 StrtNm,omitempty"`
	BldgNb      Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 BldgNb,omitempty"`
	PstCd       Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 PstCd,omitempty"`
	TwnNm       Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 TwnNm,omitempty"`
	CtrySubDvsn Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 CtrySubDvsn,omitempty"`
	Ctry        CountryCode      `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 Ctry,omitempty"`
	AdrLine     []Max70Text      `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 AdrLine,omitempty"`
}

// May be one of HIGH, NORM
type Priority2Code string

type ReferredDocumentInformation3 struct {
	Tp     ReferredDocumentType2 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 Tp,omitempty"`
	Nb     Max35Text             `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 Nb,omitempty"`
	RltdDt ISODate               `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 RltdDt,omitempty"`
}

type ReferredDocumentType1Choice struct {
	Cd    DocumentType5Code `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 Cd"`
	Prtry Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 Prtry"`
}

type ReferredDocumentType2 struct {
	CdOrPrtry ReferredDocumentType1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 CdOrPrtry"`
	Issr      Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 Issr,omitempty"`
}

type RemittanceAmount1 struct {
	DuePyblAmt        ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 DuePyblAmt,omitempty"`
	DscntApldAmt      ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 DscntApldAmt,omitempty"`
	CdtNoteAmt        ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 CdtNoteAmt,omitempty"`
	TaxAmt            ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 TaxAmt,omitempty"`
	AdjstmntAmtAndRsn []DocumentAdjustment1             `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 AdjstmntAmtAndRsn,omitempty"`
	RmtdAmt           ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 RmtdAmt,omitempty"`
}

type RemittanceInformation5 struct {
	Ustrd []Max140Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 Ustrd,omitempty"`
	Strd  []StructuredRemittanceInformation7 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 Strd,omitempty"`
}

type ReturnReason5Choice struct {
	Cd    ExternalReturnReason1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 Cd"`
	Prtry Max35Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 Prtry"`
}

type ReturnReasonInformation9 struct {
	Orgtr    PartyIdentification32 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 Orgtr,omitempty"`
	Rsn      ReturnReason5Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 Rsn,omitempty"`
	AddtlInf []Max105Text          `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 AddtlInf,omitempty"`
}

// May be one of FRST, RCUR, FNAL, OOFF
type SequenceType1Code string

type ServiceLevel8Choice struct {
	Cd    ExternalServiceLevel1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 Cd"`
	Prtry Max35Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 Prtry"`
}

type SettlementInformation13 struct {
	SttlmMtd             SettlementMethod1Code                        `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 SttlmMtd"`
	SttlmAcct            CashAccount16                                `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 SttlmAcct,omitempty"`
	ClrSys               ClearingSystemIdentification3Choice          `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 ClrSys,omitempty"`
	InstgRmbrsmntAgt     BranchAndFinancialInstitutionIdentification4 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 InstgRmbrsmntAgt,omitempty"`
	InstgRmbrsmntAgtAcct CashAccount16                                `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 InstgRmbrsmntAgtAcct,omitempty"`
	InstdRmbrsmntAgt     BranchAndFinancialInstitutionIdentification4 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 InstdRmbrsmntAgt,omitempty"`
	InstdRmbrsmntAgtAcct CashAccount16                                `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 InstdRmbrsmntAgtAcct,omitempty"`
	ThrdRmbrsmntAgt      BranchAndFinancialInstitutionIdentification4 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 ThrdRmbrsmntAgt,omitempty"`
	ThrdRmbrsmntAgtAcct  CashAccount16                                `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 ThrdRmbrsmntAgtAcct,omitempty"`
}

// May be one of INDA, INGA, COVE, CLRG
type SettlementMethod1Code string

type StructuredRemittanceInformation7 struct {
	RfrdDocInf  []ReferredDocumentInformation3 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 RfrdDocInf,omitempty"`
	RfrdDocAmt  RemittanceAmount1              `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 RfrdDocAmt,omitempty"`
	CdtrRefInf  CreditorReferenceInformation2  `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 CdtrRefInf,omitempty"`
	Invcr       PartyIdentification32          `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 Invcr,omitempty"`
	Invcee      PartyIdentification32          `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 Invcee,omitempty"`
	AddtlRmtInf []Max140Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 AddtlRmtInf,omitempty"`
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