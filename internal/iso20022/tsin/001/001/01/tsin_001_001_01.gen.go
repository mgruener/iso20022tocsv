// Code generated by main. DO NOT EDIT.

package tsin_001_001_01

import (
	"bytes"
	"encoding/xml"
	"time"
)

type AccountIdentification3Choice struct {
	IBAN      IBANIdentifier                   `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 IBAN"`
	BBAN      BBANIdentifier                   `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 BBAN"`
	UPIC      UPICIdentifier                   `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 UPIC"`
	PrtryAcct SimpleIdentificationInformation2 `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 PrtryAcct"`
}

type ActiveCurrencyAndAmount struct {
	Value float64            `xml:",chardata"`
	Ccy   ActiveCurrencyCode `xml:"Ccy,attr"`
}

// Must match the pattern [A-Z]{3,3}
type ActiveCurrencyCode string

type AdditionalInformation1 struct {
	InfTp  InformationType1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 InfTp"`
	InfVal Max350Text             `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 InfVal"`
}

// May be one of ADDR, PBOX, HOME, BIZZ, MLTO, DLVY
type AddressType2Code string

type Adjustment5 struct {
	Drctn AdjustmentDirection1Code `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 Drctn"`
	Amt   ActiveCurrencyAndAmount  `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 Amt"`
}

// May be one of ADDD, SUBS
type AdjustmentDirection1Code string

type AgreementClauses1 struct {
	Desc   Max256Text `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 Desc,omitempty"`
	DocURL Max350Text `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 DocURL,omitempty"`
}

// Must match the pattern AT[0-9]{5,5}
type AustrianBankleitzahlIdentifier string

// Must match the pattern [a-zA-Z0-9]{1,30}
type BBANIdentifier string

// Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type BEIIdentifier string

// Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type BICIdentifier string

// Must match the pattern CP[0-9]{4,4}
type CHIPSParticipantIdentifier string

// Must match the pattern CH[0-9]{6,6}
type CHIPSUniversalIdentifier string

// Must match the pattern CA[0-9]{9,9}
type CanadianPaymentsARNIdentifier string

type CashAccount7 struct {
	Id  AccountIdentification3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 Id"`
	Tp  CashAccountType2             `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 Tp,omitempty"`
	Ccy CurrencyCode                 `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 Ccy,omitempty"`
	Nm  Max70Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 Nm,omitempty"`
}

type CashAccountType2 struct {
	Cd    CashAccountType4Code `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 Cd"`
	Prtry Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 Prtry"`
}

// May be one of CASH, CHAR, COMM, TAXE, CISH, TRAS, SACC, CACC, SVGS, ONDP, MGLD, NREX, MOMA, LOAN, SLRY, ODFT
type CashAccountType4Code string

type ClearingSystemMemberIdentification2Choice struct {
	USCHU       CHIPSUniversalIdentifier                       `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 USCHU"`
	NZNCC       NewZealandNCCIdentifier                        `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 NZNCC"`
	IENSC       IrishNSCIdentifier                             `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 IENSC"`
	GBSC        UKDomesticSortCodeIdentifier                   `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 GBSC"`
	USCH        CHIPSParticipantIdentifier                     `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 USCH"`
	CHBC        SwissBCIdentifier                              `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 CHBC"`
	USFW        FedwireRoutingNumberIdentifier                 `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 USFW"`
	PTNCC       PortugueseNCCIdentifier                        `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 PTNCC"`
	RUCB        RussianCentralBankIdentificationCodeIdentifier `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 RUCB"`
	ITNCC       ItalianDomesticIdentifier                      `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 ITNCC"`
	ATBLZ       AustrianBankleitzahlIdentifier                 `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 ATBLZ"`
	CACPA       CanadianPaymentsARNIdentifier                  `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 CACPA"`
	CHSIC       SwissSICIdentifier                             `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 CHSIC"`
	DEBLZ       GermanBankleitzahlIdentifier                   `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 DEBLZ"`
	ESNCC       SpanishDomesticInterbankingIdentifier          `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 ESNCC"`
	ZANCC       SouthAfricanNCCIdentifier                      `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 ZANCC"`
	HKNCC       HongKongBankIdentifier                         `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 HKNCC"`
	AUBSBx      ExtensiveBranchNetworkIdentifier               `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 AUBSBx"`
	AUBSBs      SmallNetworkIdentifier                         `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 AUBSBs"`
	INIFSC      IndianFinancialSystemCodeIdentifier            `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 INIFSC"`
	GRHEBIC     HellenicBankIdentificationCodeIdentifier       `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 GRHEBIC"`
	PLKNR       PolishNationalClearingCodeIdentifier           `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 PLKNR"`
	OthrClrCdId Max35Text                                      `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 OthrClrCdId"`
}

type ContactIdentification1 struct {
	Nm       Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 Nm"`
	NmPrfx   NamePrefix1Code `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 NmPrfx,omitempty"`
	GvnNm    Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 GvnNm,omitempty"`
	Role     Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 Role,omitempty"`
	PhneNb   PhoneNumber     `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 PhneNb,omitempty"`
	FaxNb    PhoneNumber     `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 FaxNb,omitempty"`
	EmailAdr Max256Text      `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 EmailAdr,omitempty"`
}

// Must match the pattern [A-Z]{2,2}
type CountryCode string

// Must match the pattern [A-Z]{3,3}
type CurrencyCode string

type DateAndPlaceOfBirth struct {
	BirthDt     ISODate     `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 BirthDt"`
	PrvcOfBirth Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 PrvcOfBirth,omitempty"`
	CityOfBirth Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 CityOfBirth"`
	CtryOfBirth CountryCode `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 CtryOfBirth"`
}

type Document struct {
	InvcFincgReq InvoiceFinancingRequestV01 `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 InvcFincgReq"`
}

type DocumentGeneralInformation1 struct {
	DocTp         DocumentType4Code `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 DocTp"`
	DocNb         Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 DocNb"`
	SndrRcvrSeqId Max140Text        `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 SndrRcvrSeqId,omitempty"`
	IsseDt        ISODate           `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 IsseDt"`
	URL           Max256Text        `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 URL,omitempty"`
}

// May be one of MSIN, CNFA, DNFA, CINV, CREN, DEBN, HIRI, SBIN, CMCN, SOAC, DISP
type DocumentType2Code string

// May be one of CINV
type DocumentType4Code string

// Must match the pattern [0-9]{9,9}
type DunsIdentifier string

// Must match the pattern [0-9]{13,13}
type EANGLNIdentifier string

// Must match the pattern AU[0-9]{6,6}
type ExtensiveBranchNetworkIdentifier string

// Must match the pattern FW[0-9]{9,9}
type FedwireRoutingNumberIdentifier string

type FinancialInstitutionIdentification6 struct {
	ClrSysMmbId ClearingSystemMemberIdentification2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 ClrSysMmbId,omitempty"`
	PrtryId     GenericIdentification4                    `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 PrtryId,omitempty"`
	BIC         BICIdentifier                             `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 BIC,omitempty"`
}

type FinancingRateOrAmountChoice struct {
	Amt  ActiveCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 Amt"`
	Rate float64                 `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 Rate"`
}

type GenericIdentification3 struct {
	Id   Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 Id"`
	Issr Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 Issr,omitempty"`
}

type GenericIdentification4 struct {
	Id   Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 Id"`
	IdTp Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 IdTp"`
}

// Must match the pattern BL[0-9]{8,8}
type GermanBankleitzahlIdentifier string

// Must match the pattern GR[0-9]{7,7}
type HellenicBankIdentificationCodeIdentifier string

// Must match the pattern HK[0-9]{3,3}
type HongKongBankIdentifier string

// Must match the pattern [a-zA-Z]{2,2}[0-9]{2,2}[a-zA-Z0-9]{1,30}
type IBANIdentifier string

// Must match the pattern [A-Z]{2,2}[B-DF-HJ-NP-TV-XZ0-9]{7,7}[0-9]{1,1}
type IBEIIdentifier string

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

// Must match the pattern IN[a-zA-Z0-9]{11,11}
type IndianFinancialSystemCodeIdentifier string

type InformationType1Choice struct {
	Cd    InformationType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 Cd"`
	Prtry Max140Text           `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 Prtry"`
}

// May be one of INST, RELY
type InformationType1Code string

type Instalment1 struct {
	SeqId    Max70Text               `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 SeqId"`
	PmtDueDt ISODate                 `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 PmtDueDt"`
	Amt      ActiveCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 Amt"`
}

type InvoiceFinancingRequestV01 struct {
	ReqGrpInf  RequestGroupInformation1     `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 ReqGrpInf"`
	InvcReqInf []InvoiceRequestInformation1 `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 InvcReqInf"`
}

type InvoiceRequestInformation1 struct {
	InvcGnlInf    DocumentGeneralInformation1                         `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 InvcGnlInf"`
	InvcTtlsInf   InvoiceTotals1                                      `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 InvcTtlsInf"`
	CdtDbtNoteAmt ActiveCurrencyAndAmount                             `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 CdtDbtNoteAmt,omitempty"`
	InstlmtInf    []Instalment1                                       `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 InstlmtInf,omitempty"`
	ReqdAmt       FinancingRateOrAmountChoice                         `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 ReqdAmt,omitempty"`
	Spplr         PartyAndAccountIdentificationAndContactInformation1 `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 Spplr"`
	Buyr          PartyIdentificationAndContactInformation1           `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 Buyr"`
	InvcPmtInf    PaymentInformation15                                `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 InvcPmtInf"`
	RfrdDoc       []ReferredDocumentInformation2                      `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 RfrdDoc,omitempty"`
}

type InvoiceTotals1 struct {
	TtlTaxblAmt ActiveCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 TtlTaxblAmt"`
	TtlTaxAmt   ActiveCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 TtlTaxAmt"`
	Adjstmnt    Adjustment5             `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 Adjstmnt,omitempty"`
	TtlInvcAmt  ActiveCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 TtlInvcAmt"`
	PmtDueDt    ISODate                 `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 PmtDueDt"`
}

// Must match the pattern IE[0-9]{6,6}
type IrishNSCIdentifier string

// Must match the pattern IT[0-9]{10,10}
type ItalianDomesticIdentifier string

// Must be at least 1 items long
type Max128Text string

// Must be at least 1 items long
type Max140Text string

// Must match the pattern [0-9]{1,15}
type Max15NumericText string

// Must be at least 1 items long
type Max16Text string

// Must be at least 1 items long
type Max256Text string

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

// Must match the pattern NZ[0-9]{6,6}
type NewZealandNCCIdentifier string

type OrganisationIdentification2 struct {
	BIC     BICIdentifier            `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 BIC,omitempty"`
	IBEI    IBEIIdentifier           `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 IBEI,omitempty"`
	BEI     BEIIdentifier            `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 BEI,omitempty"`
	EANGLN  EANGLNIdentifier         `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 EANGLN,omitempty"`
	USCHU   CHIPSUniversalIdentifier `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 USCHU,omitempty"`
	DUNS    DunsIdentifier           `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 DUNS,omitempty"`
	BkPtyId Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 BkPtyId,omitempty"`
	TaxIdNb Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 TaxIdNb,omitempty"`
	PrtryId GenericIdentification3   `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 PrtryId,omitempty"`
}

type Party2Choice struct {
	OrgId  OrganisationIdentification2 `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 OrgId"`
	PrvtId []PersonIdentification3     `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 PrvtId"`
}

type PartyAndAccountIdentificationAndContactInformation1 struct {
	PtyId   PartyIdentification8   `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 PtyId"`
	AcctId  CashAccount7           `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 AcctId,omitempty"`
	CtctInf ContactIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 CtctInf,omitempty"`
}

type PartyIdentification25 struct {
	Nm      Max70Text              `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 Nm"`
	PrtryId GenericIdentification4 `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 PrtryId,omitempty"`
	BEI     BEIIdentifier          `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 BEI,omitempty"`
}

type PartyIdentification8 struct {
	Nm        Max70Text      `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 Nm,omitempty"`
	PstlAdr   PostalAddress1 `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 PstlAdr,omitempty"`
	Id        Party2Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 Id,omitempty"`
	CtryOfRes CountryCode    `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 CtryOfRes,omitempty"`
}

type PartyIdentificationAndAccount6 struct {
	PtyId     PartyIdentification25 `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 PtyId"`
	CdtAcct   CashAccount7          `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 CdtAcct,omitempty"`
	FincgAcct CashAccount7          `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 FincgAcct,omitempty"`
}

type PartyIdentificationAndContactInformation1 struct {
	PtyId   PartyIdentification8   `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 PtyId"`
	CtctInf ContactIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 CtctInf,omitempty"`
}

type PaymentInformation15 struct {
	PmtMtd  PaymentMethod4Code `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 PmtMtd"`
	PmtAcct CashAccount7       `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 PmtAcct,omitempty"`
}

// May be one of CHK, TRF, DD, TRA
type PaymentMethod4Code string

type PersonIdentification3 struct {
	DrvrsLicNb      Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 DrvrsLicNb"`
	CstmrNb         Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 CstmrNb"`
	SclSctyNb       Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 SclSctyNb"`
	AlnRegnNb       Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 AlnRegnNb"`
	PsptNb          Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 PsptNb"`
	TaxIdNb         Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 TaxIdNb"`
	IdntyCardNb     Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 IdntyCardNb"`
	MplyrIdNb       Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 MplyrIdNb"`
	DtAndPlcOfBirth DateAndPlaceOfBirth    `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 DtAndPlcOfBirth"`
	OthrId          GenericIdentification4 `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 OthrId"`
	Issr            Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 Issr,omitempty"`
}

// Must match the pattern \+[0-9]{1,3}-[0-9()+\-]{1,30}
type PhoneNumber string

// Must match the pattern PL[0-9]{8,8}
type PolishNationalClearingCodeIdentifier string

// Must match the pattern PT[0-9]{8,8}
type PortugueseNCCIdentifier string

type PostalAddress1 struct {
	AdrTp       AddressType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 AdrTp,omitempty"`
	AdrLine     []Max70Text      `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 AdrLine,omitempty"`
	StrtNm      Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 StrtNm,omitempty"`
	BldgNb      Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 BldgNb,omitempty"`
	PstCd       Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 PstCd,omitempty"`
	TwnNm       Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 TwnNm,omitempty"`
	CtrySubDvsn Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 CtrySubDvsn,omitempty"`
	Ctry        CountryCode      `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 Ctry"`
}

type ReferredDocumentInformation2 struct {
	Tp     ReferredDocumentType1   `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 Tp,omitempty"`
	DocNb  Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 DocNb,omitempty"`
	RltdDt ISODate                 `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 RltdDt,omitempty"`
	DocAmt ActiveCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 DocAmt,omitempty"`
}

type ReferredDocumentType1 struct {
	Cd    DocumentType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 Cd"`
	Prtry Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 Prtry"`
	Issr  Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 Issr,omitempty"`
}

type RequestGroupInformation1 struct {
	GrpId         Max35Text                           `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 GrpId"`
	CreDtTm       ISODateTime                         `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 CreDtTm"`
	Authstn       []Max128Text                        `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 Authstn,omitempty"`
	NbOfInvcReqs  Max15NumericText                    `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 NbOfInvcReqs,omitempty"`
	TtlBlkInvcAmt ActiveCurrencyAndAmount             `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 TtlBlkInvcAmt,omitempty"`
	Ccy           CurrencyCode                        `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 Ccy"`
	FincgAgrmt    Max350Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 FincgAgrmt,omitempty"`
	FincgRqstr    PartyIdentificationAndAccount6      `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 FincgRqstr"`
	IntrmyAgt     FinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 IntrmyAgt,omitempty"`
	FrstAgt       FinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 FrstAgt"`
	AgrmtClauses  []AgreementClauses1                 `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 AgrmtClauses,omitempty"`
	AddtlInf      []AdditionalInformation1            `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 AddtlInf,omitempty"`
}

// Must match the pattern RU[0-9]{9,9}
type RussianCentralBankIdentificationCodeIdentifier string

type SimpleIdentificationInformation2 struct {
	Id Max34Text `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 Id"`
}

// Must match the pattern AU[0-9]{6,6}
type SmallNetworkIdentifier string

// Must match the pattern ZA[0-9]{6,6}
type SouthAfricanNCCIdentifier string

// Must match the pattern ES[0-9]{8,9}
type SpanishDomesticInterbankingIdentifier string

// Must match the pattern SW[0-9]{3,5}
type SwissBCIdentifier string

// Must match the pattern SW[0-9]{6,6}
type SwissSICIdentifier string

// Must match the pattern SC[0-9]{6,6}
type UKDomesticSortCodeIdentifier string

// Must match the pattern [0-9]{8,17}
type UPICIdentifier string

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