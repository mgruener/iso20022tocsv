// Code generated by main. DO NOT EDIT.

package pain_001_001_03

import (
	"bytes"
	"encoding/xml"
	"time"
)

type AccountIdentification4Choice struct {
	IBAN IBAN2007Identifier            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 IBAN"`
	Othr GenericAccountIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 Othr"`
}

type AccountSchemeName1Choice struct {
	Cd    ExternalAccountIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 Cd"`
	Prtry Max35Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 Prtry"`
}

type ActiveOrHistoricCurrencyAndAmount struct {
	Value float64                      `xml:",chardata"`
	Ccy   ActiveOrHistoricCurrencyCode `xml:"Ccy,attr"`
}

// Must match the pattern [A-Z]{3,3}
type ActiveOrHistoricCurrencyCode string

// May be one of ADDR, PBOX, HOME, BIZZ, MLTO, DLVY
type AddressType2Code string

type AmountType3Choice struct {
	InstdAmt ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 InstdAmt"`
	EqvtAmt  EquivalentAmount2                 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 EqvtAmt"`
}

// Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type AnyBICIdentifier string

type Authorisation1Choice struct {
	Cd    Authorisation1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 Cd"`
	Prtry Max128Text         `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 Prtry"`
}

// May be one of AUTH, FDET, FSUM, ILEV
type Authorisation1Code string

// Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type BICIdentifier string

type BranchAndFinancialInstitutionIdentification4 struct {
	FinInstnId FinancialInstitutionIdentification7 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 FinInstnId"`
	BrnchId    BranchData2                         `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 BrnchId,omitempty"`
}

type BranchData2 struct {
	Id      Max35Text      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 Id,omitempty"`
	Nm      Max140Text     `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 Nm,omitempty"`
	PstlAdr PostalAddress6 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 PstlAdr,omitempty"`
}

type CashAccount16 struct {
	Id  AccountIdentification4Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 Id"`
	Tp  CashAccountType2             `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 Tp,omitempty"`
	Ccy ActiveOrHistoricCurrencyCode `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 Ccy,omitempty"`
	Nm  Max70Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 Nm,omitempty"`
}

type CashAccountType2 struct {
	Cd    CashAccountType4Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 Cd"`
	Prtry Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 Prtry"`
}

// May be one of CASH, CHAR, COMM, TAXE, CISH, TRAS, SACC, CACC, SVGS, ONDP, MGLD, NREX, MOMA, LOAN, SLRY, ODFT
type CashAccountType4Code string

type CategoryPurpose1Choice struct {
	Cd    ExternalCategoryPurpose1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 Cd"`
	Prtry Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 Prtry"`
}

// May be one of DEBT, CRED, SHAR, SLEV
type ChargeBearerType1Code string

type Cheque6 struct {
	ChqTp       ChequeType2Code             `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 ChqTp,omitempty"`
	ChqNb       Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 ChqNb,omitempty"`
	ChqFr       NameAndAddress10            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 ChqFr,omitempty"`
	DlvryMtd    ChequeDeliveryMethod1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 DlvryMtd,omitempty"`
	DlvrTo      NameAndAddress10            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 DlvrTo,omitempty"`
	InstrPrty   Priority2Code               `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 InstrPrty,omitempty"`
	ChqMtrtyDt  ISODate                     `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 ChqMtrtyDt,omitempty"`
	FrmsCd      Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 FrmsCd,omitempty"`
	MemoFld     []Max35Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 MemoFld,omitempty"`
	RgnlClrZone Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 RgnlClrZone,omitempty"`
	PrtLctn     Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 PrtLctn,omitempty"`
}

// May be one of MLDB, MLCD, MLFA, CRDB, CRCD, CRFA, PUDB, PUCD, PUFA, RGDB, RGCD, RGFA
type ChequeDelivery1Code string

type ChequeDeliveryMethod1Choice struct {
	Cd    ChequeDelivery1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 Cd"`
	Prtry Max35Text           `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 Prtry"`
}

// May be one of CCHQ, CCCH, BCHQ, DRFT, ELDR
type ChequeType2Code string

type ClearingSystemIdentification2Choice struct {
	Cd    ExternalClearingSystemIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 Cd"`
	Prtry Max35Text                                 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 Prtry"`
}

type ClearingSystemMemberIdentification2 struct {
	ClrSysId ClearingSystemIdentification2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 ClrSysId,omitempty"`
	MmbId    Max35Text                           `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 MmbId"`
}

type ContactDetails2 struct {
	NmPrfx   NamePrefix1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 NmPrfx,omitempty"`
	Nm       Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 Nm,omitempty"`
	PhneNb   PhoneNumber     `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 PhneNb,omitempty"`
	MobNb    PhoneNumber     `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 MobNb,omitempty"`
	FaxNb    PhoneNumber     `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 FaxNb,omitempty"`
	EmailAdr Max2048Text     `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 EmailAdr,omitempty"`
	Othr     Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 Othr,omitempty"`
}

// Must match the pattern [A-Z]{2,2}
type CountryCode string

// May be one of CRDT, DBIT
type CreditDebitCode string

type CreditTransferTransactionInformation10 struct {
	PmtId           PaymentIdentification1                       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 PmtId"`
	PmtTpInf        PaymentTypeInformation19                     `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 PmtTpInf,omitempty"`
	Amt             AmountType3Choice                            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 Amt"`
	XchgRateInf     ExchangeRateInformation1                     `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 XchgRateInf,omitempty"`
	ChrgBr          ChargeBearerType1Code                        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 ChrgBr,omitempty"`
	ChqInstr        Cheque6                                      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 ChqInstr,omitempty"`
	UltmtDbtr       PartyIdentification32                        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 UltmtDbtr,omitempty"`
	IntrmyAgt1      BranchAndFinancialInstitutionIdentification4 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 IntrmyAgt1,omitempty"`
	IntrmyAgt1Acct  CashAccount16                                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 IntrmyAgt1Acct,omitempty"`
	IntrmyAgt2      BranchAndFinancialInstitutionIdentification4 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 IntrmyAgt2,omitempty"`
	IntrmyAgt2Acct  CashAccount16                                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 IntrmyAgt2Acct,omitempty"`
	IntrmyAgt3      BranchAndFinancialInstitutionIdentification4 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 IntrmyAgt3,omitempty"`
	IntrmyAgt3Acct  CashAccount16                                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 IntrmyAgt3Acct,omitempty"`
	CdtrAgt         BranchAndFinancialInstitutionIdentification4 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 CdtrAgt,omitempty"`
	CdtrAgtAcct     CashAccount16                                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 CdtrAgtAcct,omitempty"`
	Cdtr            PartyIdentification32                        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 Cdtr,omitempty"`
	CdtrAcct        CashAccount16                                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 CdtrAcct,omitempty"`
	UltmtCdtr       PartyIdentification32                        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 UltmtCdtr,omitempty"`
	InstrForCdtrAgt []InstructionForCreditorAgent1               `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 InstrForCdtrAgt,omitempty"`
	InstrForDbtrAgt Max140Text                                   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 InstrForDbtrAgt,omitempty"`
	Purp            Purpose2Choice                               `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 Purp,omitempty"`
	RgltryRptg      []RegulatoryReporting3                       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 RgltryRptg,omitempty"`
	Tax             TaxInformation3                              `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 Tax,omitempty"`
	RltdRmtInf      []RemittanceLocation2                        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 RltdRmtInf,omitempty"`
	RmtInf          RemittanceInformation5                       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 RmtInf,omitempty"`
}

type CreditorReferenceInformation2 struct {
	Tp  CreditorReferenceType2 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 Tp,omitempty"`
	Ref Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 Ref,omitempty"`
}

type CreditorReferenceType1Choice struct {
	Cd    DocumentType3Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 Cd"`
	Prtry Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 Prtry"`
}

type CreditorReferenceType2 struct {
	CdOrPrtry CreditorReferenceType1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 CdOrPrtry"`
	Issr      Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 Issr,omitempty"`
}

type CustomerCreditTransferInitiationV03 struct {
	GrpHdr GroupHeader32                    `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 GrpHdr"`
	PmtInf []PaymentInstructionInformation3 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 PmtInf"`
}

type DateAndPlaceOfBirth struct {
	BirthDt     ISODate     `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 BirthDt"`
	PrvcOfBirth Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 PrvcOfBirth,omitempty"`
	CityOfBirth Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 CityOfBirth"`
	CtryOfBirth CountryCode `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 CtryOfBirth"`
}

type DatePeriodDetails struct {
	FrDt ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 FrDt"`
	ToDt ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 ToDt"`
}

type Document struct {
	CstmrCdtTrfInitn CustomerCreditTransferInitiationV03 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 CstmrCdtTrfInitn"`
}

type DocumentAdjustment1 struct {
	Amt       ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 Amt"`
	CdtDbtInd CreditDebitCode                   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 CdtDbtInd,omitempty"`
	Rsn       Max4Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 Rsn,omitempty"`
	AddtlInf  Max140Text                        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 AddtlInf,omitempty"`
}

// May be one of RADM, RPIN, FXDR, DISP, PUOR, SCOR
type DocumentType3Code string

// May be one of MSIN, CNFA, DNFA, CINV, CREN, DEBN, HIRI, SBIN, CMCN, SOAC, DISP, BOLD, VCHR, AROI, TSUT
type DocumentType5Code string

type EquivalentAmount2 struct {
	Amt      ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 Amt"`
	CcyOfTrf ActiveOrHistoricCurrencyCode      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 CcyOfTrf"`
}

type ExchangeRateInformation1 struct {
	XchgRate float64               `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 XchgRate,omitempty"`
	RateTp   ExchangeRateType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 RateTp,omitempty"`
	CtrctId  Max35Text             `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 CtrctId,omitempty"`
}

// May be one of SPOT, SALE, AGRD
type ExchangeRateType1Code string

// Must be at least 1 items long
type ExternalAccountIdentification1Code string

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
type ExternalPurpose1Code string

// Must be at least 1 items long
type ExternalServiceLevel1Code string

type FinancialIdentificationSchemeName1Choice struct {
	Cd    ExternalFinancialInstitutionIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 Cd"`
	Prtry Max35Text                                       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 Prtry"`
}

type FinancialInstitutionIdentification7 struct {
	BIC         BICIdentifier                       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 BIC,omitempty"`
	ClrSysMmbId ClearingSystemMemberIdentification2 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 ClrSysMmbId,omitempty"`
	Nm          Max140Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 Nm,omitempty"`
	PstlAdr     PostalAddress6                      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 PstlAdr,omitempty"`
	Othr        GenericFinancialIdentification1     `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 Othr,omitempty"`
}

type GenericAccountIdentification1 struct {
	Id      Max34Text                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 Id"`
	SchmeNm AccountSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 SchmeNm,omitempty"`
	Issr    Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 Issr,omitempty"`
}

type GenericFinancialIdentification1 struct {
	Id      Max35Text                                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 Id"`
	SchmeNm FinancialIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 SchmeNm,omitempty"`
	Issr    Max35Text                                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 Issr,omitempty"`
}

type GenericOrganisationIdentification1 struct {
	Id      Max35Text                                   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 Id"`
	SchmeNm OrganisationIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 SchmeNm,omitempty"`
	Issr    Max35Text                                   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 Issr,omitempty"`
}

type GenericPersonIdentification1 struct {
	Id      Max35Text                             `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 Id"`
	SchmeNm PersonIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 SchmeNm,omitempty"`
	Issr    Max35Text                             `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 Issr,omitempty"`
}

type GroupHeader32 struct {
	MsgId    Max35Text                                    `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 MsgId"`
	CreDtTm  ISODateTime                                  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 CreDtTm"`
	Authstn  []Authorisation1Choice                       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 Authstn,omitempty"`
	NbOfTxs  Max15NumericText                             `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 NbOfTxs"`
	CtrlSum  float64                                      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 CtrlSum,omitempty"`
	InitgPty PartyIdentification32                        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 InitgPty"`
	FwdgAgt  BranchAndFinancialInstitutionIdentification4 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 FwdgAgt,omitempty"`
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

// May be one of CHQB, HOLD, PHOB, TELB
type Instruction3Code string

type InstructionForCreditorAgent1 struct {
	Cd       Instruction3Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 Cd,omitempty"`
	InstrInf Max140Text       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 InstrInf,omitempty"`
}

type LocalInstrument2Choice struct {
	Cd    ExternalLocalInstrument1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 Cd"`
	Prtry Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 Prtry"`
}

// Must be at least 1 items long
type Max10Text string

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

type NameAndAddress10 struct {
	Nm  Max140Text     `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 Nm"`
	Adr PostalAddress6 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 Adr"`
}

// May be one of DOCT, MIST, MISS, MADM
type NamePrefix1Code string

type OrganisationIdentification4 struct {
	BICOrBEI AnyBICIdentifier                     `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 BICOrBEI,omitempty"`
	Othr     []GenericOrganisationIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 Othr,omitempty"`
}

type OrganisationIdentificationSchemeName1Choice struct {
	Cd    ExternalOrganisationIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 Cd"`
	Prtry Max35Text                               `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 Prtry"`
}

type Party6Choice struct {
	OrgId  OrganisationIdentification4 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 OrgId"`
	PrvtId PersonIdentification5       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 PrvtId"`
}

type PartyIdentification32 struct {
	Nm        Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 Nm,omitempty"`
	PstlAdr   PostalAddress6  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 PstlAdr,omitempty"`
	Id        Party6Choice    `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 Id,omitempty"`
	CtryOfRes CountryCode     `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 CtryOfRes,omitempty"`
	CtctDtls  ContactDetails2 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 CtctDtls,omitempty"`
}

type PaymentIdentification1 struct {
	InstrId    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 InstrId,omitempty"`
	EndToEndId Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 EndToEndId"`
}

type PaymentInstructionInformation3 struct {
	PmtInfId        Max35Text                                    `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 PmtInfId"`
	PmtMtd          PaymentMethod3Code                           `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 PmtMtd"`
	BtchBookg       bool                                         `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 BtchBookg,omitempty"`
	NbOfTxs         Max15NumericText                             `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 NbOfTxs,omitempty"`
	CtrlSum         float64                                      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 CtrlSum,omitempty"`
	PmtTpInf        PaymentTypeInformation19                     `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 PmtTpInf,omitempty"`
	ReqdExctnDt     ISODate                                      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 ReqdExctnDt"`
	PoolgAdjstmntDt ISODate                                      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 PoolgAdjstmntDt,omitempty"`
	Dbtr            PartyIdentification32                        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 Dbtr"`
	DbtrAcct        CashAccount16                                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 DbtrAcct"`
	DbtrAgt         BranchAndFinancialInstitutionIdentification4 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 DbtrAgt"`
	DbtrAgtAcct     CashAccount16                                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 DbtrAgtAcct,omitempty"`
	UltmtDbtr       PartyIdentification32                        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 UltmtDbtr,omitempty"`
	ChrgBr          ChargeBearerType1Code                        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 ChrgBr,omitempty"`
	ChrgsAcct       CashAccount16                                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 ChrgsAcct,omitempty"`
	ChrgsAcctAgt    BranchAndFinancialInstitutionIdentification4 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 ChrgsAcctAgt,omitempty"`
	CdtTrfTxInf     []CreditTransferTransactionInformation10     `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 CdtTrfTxInf"`
}

// May be one of CHK, TRF, TRA
type PaymentMethod3Code string

type PaymentTypeInformation19 struct {
	InstrPrty Priority2Code          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 InstrPrty,omitempty"`
	SvcLvl    ServiceLevel8Choice    `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 SvcLvl,omitempty"`
	LclInstrm LocalInstrument2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 LclInstrm,omitempty"`
	CtgyPurp  CategoryPurpose1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 CtgyPurp,omitempty"`
}

type PersonIdentification5 struct {
	DtAndPlcOfBirth DateAndPlaceOfBirth            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 DtAndPlcOfBirth,omitempty"`
	Othr            []GenericPersonIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 Othr,omitempty"`
}

type PersonIdentificationSchemeName1Choice struct {
	Cd    ExternalPersonIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 Cd"`
	Prtry Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 Prtry"`
}

// Must match the pattern \+[0-9]{1,3}-[0-9()+\-]{1,30}
type PhoneNumber string

type PostalAddress6 struct {
	AdrTp       AddressType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 AdrTp,omitempty"`
	Dept        Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 Dept,omitempty"`
	SubDept     Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 SubDept,omitempty"`
	StrtNm      Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 StrtNm,omitempty"`
	BldgNb      Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 BldgNb,omitempty"`
	PstCd       Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 PstCd,omitempty"`
	TwnNm       Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 TwnNm,omitempty"`
	CtrySubDvsn Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 CtrySubDvsn,omitempty"`
	Ctry        CountryCode      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 Ctry,omitempty"`
	AdrLine     []Max70Text      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 AdrLine,omitempty"`
}

// May be one of HIGH, NORM
type Priority2Code string

type Purpose2Choice struct {
	Cd    ExternalPurpose1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 Cd"`
	Prtry Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 Prtry"`
}

type ReferredDocumentInformation3 struct {
	Tp     ReferredDocumentType2 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 Tp,omitempty"`
	Nb     Max35Text             `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 Nb,omitempty"`
	RltdDt ISODate               `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 RltdDt,omitempty"`
}

type ReferredDocumentType1Choice struct {
	Cd    DocumentType5Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 Cd"`
	Prtry Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 Prtry"`
}

type ReferredDocumentType2 struct {
	CdOrPrtry ReferredDocumentType1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 CdOrPrtry"`
	Issr      Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 Issr,omitempty"`
}

type RegulatoryAuthority2 struct {
	Nm   Max140Text  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 Nm,omitempty"`
	Ctry CountryCode `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 Ctry,omitempty"`
}

type RegulatoryReporting3 struct {
	DbtCdtRptgInd RegulatoryReportingType1Code     `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 DbtCdtRptgInd,omitempty"`
	Authrty       RegulatoryAuthority2             `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 Authrty,omitempty"`
	Dtls          []StructuredRegulatoryReporting3 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 Dtls,omitempty"`
}

// May be one of CRED, DEBT, BOTH
type RegulatoryReportingType1Code string

type RemittanceAmount1 struct {
	DuePyblAmt        ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 DuePyblAmt,omitempty"`
	DscntApldAmt      ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 DscntApldAmt,omitempty"`
	CdtNoteAmt        ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 CdtNoteAmt,omitempty"`
	TaxAmt            ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 TaxAmt,omitempty"`
	AdjstmntAmtAndRsn []DocumentAdjustment1             `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 AdjstmntAmtAndRsn,omitempty"`
	RmtdAmt           ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 RmtdAmt,omitempty"`
}

type RemittanceInformation5 struct {
	Ustrd []Max140Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 Ustrd,omitempty"`
	Strd  []StructuredRemittanceInformation7 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 Strd,omitempty"`
}

type RemittanceLocation2 struct {
	RmtId             Max35Text                     `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 RmtId,omitempty"`
	RmtLctnMtd        RemittanceLocationMethod2Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 RmtLctnMtd,omitempty"`
	RmtLctnElctrncAdr Max2048Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 RmtLctnElctrncAdr,omitempty"`
	RmtLctnPstlAdr    NameAndAddress10              `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 RmtLctnPstlAdr,omitempty"`
}

// May be one of FAXI, EDIC, URID, EMAL, POST, SMSM
type RemittanceLocationMethod2Code string

type ServiceLevel8Choice struct {
	Cd    ExternalServiceLevel1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 Cd"`
	Prtry Max35Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 Prtry"`
}

type StructuredRegulatoryReporting3 struct {
	Tp   Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 Tp,omitempty"`
	Dt   ISODate                           `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 Dt,omitempty"`
	Ctry CountryCode                       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 Ctry,omitempty"`
	Cd   Max10Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 Cd,omitempty"`
	Amt  ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 Amt,omitempty"`
	Inf  []Max35Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 Inf,omitempty"`
}

type StructuredRemittanceInformation7 struct {
	RfrdDocInf  []ReferredDocumentInformation3 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 RfrdDocInf,omitempty"`
	RfrdDocAmt  RemittanceAmount1              `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 RfrdDocAmt,omitempty"`
	CdtrRefInf  CreditorReferenceInformation2  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 CdtrRefInf,omitempty"`
	Invcr       PartyIdentification32          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 Invcr,omitempty"`
	Invcee      PartyIdentification32          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 Invcee,omitempty"`
	AddtlRmtInf []Max140Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 AddtlRmtInf,omitempty"`
}

type TaxAmount1 struct {
	Rate         float64                           `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 Rate,omitempty"`
	TaxblBaseAmt ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 TaxblBaseAmt,omitempty"`
	TtlAmt       ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 TtlAmt,omitempty"`
	Dtls         []TaxRecordDetails1               `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 Dtls,omitempty"`
}

type TaxAuthorisation1 struct {
	Titl Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 Titl,omitempty"`
	Nm   Max140Text `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 Nm,omitempty"`
}

type TaxInformation3 struct {
	Cdtr            TaxParty1                         `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 Cdtr,omitempty"`
	Dbtr            TaxParty2                         `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 Dbtr,omitempty"`
	AdmstnZn        Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 AdmstnZn,omitempty"`
	RefNb           Max140Text                        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 RefNb,omitempty"`
	Mtd             Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 Mtd,omitempty"`
	TtlTaxblBaseAmt ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 TtlTaxblBaseAmt,omitempty"`
	TtlTaxAmt       ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 TtlTaxAmt,omitempty"`
	Dt              ISODate                           `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 Dt,omitempty"`
	SeqNb           float64                           `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 SeqNb,omitempty"`
	Rcrd            []TaxRecord1                      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 Rcrd,omitempty"`
}

type TaxParty1 struct {
	TaxId  Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 TaxId,omitempty"`
	RegnId Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 RegnId,omitempty"`
	TaxTp  Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 TaxTp,omitempty"`
}

type TaxParty2 struct {
	TaxId   Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 TaxId,omitempty"`
	RegnId  Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 RegnId,omitempty"`
	TaxTp   Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 TaxTp,omitempty"`
	Authstn TaxAuthorisation1 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 Authstn,omitempty"`
}

type TaxPeriod1 struct {
	Yr     ISODate              `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 Yr,omitempty"`
	Tp     TaxRecordPeriod1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 Tp,omitempty"`
	FrToDt DatePeriodDetails    `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 FrToDt,omitempty"`
}

type TaxRecord1 struct {
	Tp       Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 Tp,omitempty"`
	Ctgy     Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 Ctgy,omitempty"`
	CtgyDtls Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 CtgyDtls,omitempty"`
	DbtrSts  Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 DbtrSts,omitempty"`
	CertId   Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 CertId,omitempty"`
	FrmsCd   Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 FrmsCd,omitempty"`
	Prd      TaxPeriod1 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 Prd,omitempty"`
	TaxAmt   TaxAmount1 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 TaxAmt,omitempty"`
	AddtlInf Max140Text `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 AddtlInf,omitempty"`
}

type TaxRecordDetails1 struct {
	Prd TaxPeriod1                        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 Prd,omitempty"`
	Amt ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 Amt"`
}

// May be one of MM01, MM02, MM03, MM04, MM05, MM06, MM07, MM08, MM09, MM10, MM11, MM12, QTR1, QTR2, QTR3, QTR4, HLF1, HLF2
type TaxRecordPeriod1Code string

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
