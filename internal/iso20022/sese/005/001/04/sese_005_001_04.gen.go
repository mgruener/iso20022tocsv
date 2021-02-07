// Code generated by main. DO NOT EDIT.

package sese_005_001_04

import (
	"bytes"
	"encoding/xml"
	"time"
)

type Account7 struct {
	Id       AccountIdentification1     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 Id"`
	AcctSvcr PartyIdentification2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 AcctSvcr,omitempty"`
}

type AccountIdentification1 struct {
	Prtry SimpleIdentificationInformation `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 Prtry"`
}

type ActiveCurrencyAnd13DecimalAmount struct {
	Value float64            `xml:",chardata"`
	Ccy   ActiveCurrencyCode `xml:"Ccy,attr"`
}

type ActiveCurrencyAndAmount struct {
	Value float64            `xml:",chardata"`
	Ccy   ActiveCurrencyCode `xml:"Ccy,attr"`
}

// Must match the pattern [A-Z]{3,3}
type ActiveCurrencyCode string

type ActiveOrHistoricCurrencyAnd13DecimalAmount struct {
	Value float64                      `xml:",chardata"`
	Ccy   ActiveOrHistoricCurrencyCode `xml:"Ccy,attr"`
}

type ActiveOrHistoricCurrencyAndAmount struct {
	Value float64                      `xml:",chardata"`
	Ccy   ActiveOrHistoricCurrencyCode `xml:"Ccy,attr"`
}

// Must match the pattern [A-Z]{3,3}
type ActiveOrHistoricCurrencyCode string

type AdditionalReference2 struct {
	Ref     Max35Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 Ref"`
	RefIssr PartyIdentification1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 RefIssr,omitempty"`
	MsgNm   Max35Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 MsgNm,omitempty"`
}

// May be one of ADDR, PBOX, HOME, BIZZ, MLTO, DLVY
type AddressType2Code string

type AlternateSecurityIdentification1 struct {
	Id         Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 Id"`
	DmstIdSrc  CountryCode `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 DmstIdSrc"`
	PrtryIdSrc Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 PrtryIdSrc"`
}

// Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type AnyBICIdentifier string

type BICIdentification1 struct {
	BIC BICIdentifier `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 BIC"`
}

// Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type BICIdentifier string

// May be one of NCER, ELEC, PHYS
type BeneficiaryCertificationCompletion1Code string

// Must be at least 1 items long
type BloombergIdentifier string

// May be one of SLDP, SLRP, DLPR
type BusinessFlowType1Code string

type Charge20 struct {
	Tp             ChargeType12Code                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 Tp"`
	XtndedTp       Extended350Code                   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 XtndedTp"`
	Amt            ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 Amt"`
	ChrgBsis       TaxationBasis2Code                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 ChrgBsis,omitempty"`
	XtndedChrgBsis Extended350Code                   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 XtndedChrgBsis,omitempty"`
	ChrgBr         ChargeBearer1Code                 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 ChrgBr,omitempty"`
	RcptId         PartyIdentification2Choice        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 RcptId,omitempty"`
}

// May be one of OUR, BEN, SHA
type ChargeBearer1Code string

// May be one of BEND, DISC, FEND, POST, REGF, SHIP, SPCN, TRAN
type ChargeType12Code string

type Commission12 struct {
	Tp             CommissionType7Code              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 Tp"`
	XtndedTp       Extended350Code                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 XtndedTp"`
	Bsis           TaxationBasis4Code               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 Bsis,omitempty"`
	XtndedBsis     Extended350Code                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 XtndedBsis,omitempty"`
	Amt            ActiveCurrencyAnd13DecimalAmount `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 Amt"`
	RcptId         PartyIdentification2Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 RcptId,omitempty"`
	ComrclAgrmtRef Max35Text                        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 ComrclAgrmtRef,omitempty"`
}

// May be one of FEND, BEND
type CommissionType7Code string

// Must be at least 1 items long
type ConsolidatedTapeAssociationIdentifier string

type ContactIdentification2 struct {
	NmPrfx   NamePrefix1Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 NmPrfx,omitempty"`
	GvnNm    Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 GvnNm,omitempty"`
	Nm       Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 Nm"`
	PhneNb   PhoneNumber     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 PhneNb,omitempty"`
	MobNb    PhoneNumber     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 MobNb,omitempty"`
	FaxNb    PhoneNumber     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 FaxNb,omitempty"`
	EmailAdr Max256Text      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 EmailAdr,omitempty"`
}

type CopyInformation2 struct {
	CpyInd    bool               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 CpyInd"`
	OrgnlRcvr BICIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 OrgnlRcvr,omitempty"`
}

// Must match the pattern [A-Z]{2,2}
type CountryCode string

type DateAndDateTimeChoice struct {
	Dt   ISODate     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 Dt"`
	DtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 DtTm"`
}

type DateFormat1Choice struct {
	Dt   ISODate             `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 Dt"`
	Cd   SettlementDate1Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 Cd"`
	DtTm ISODateTime         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 DtTm"`
}

type DeliverInformation9 struct {
	ReqdSttlmDt    ISODate                      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 ReqdSttlmDt,omitempty"`
	SttlmAmt       ActiveCurrencyAndAmount      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 SttlmAmt,omitempty"`
	StmpDty        StampDutyType2Code           `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 StmpDty,omitempty"`
	NetAmt         ActiveCurrencyAndAmount      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 NetAmt,omitempty"`
	SttlmPtiesDtls DeliveringPartiesAndAccount8 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 SttlmPtiesDtls,omitempty"`
	ChrgDtls       []Charge20                   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 ChrgDtls,omitempty"`
	ComssnDtls     []Commission12               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 ComssnDtls,omitempty"`
	TaxDtls        []Tax15                      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 TaxDtls,omitempty"`
	PhysTrf        PhysicalTransferType1Code    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 PhysTrf,omitempty"`
	PhysTrfDtls    DeliveryParameters4          `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 PhysTrfDtls,omitempty"`
}

type DeliveringPartiesAndAccount8 struct {
	DlvrrDtls        InvestmentAccount24            `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 DlvrrDtls,omitempty"`
	DlvrrsCtdnDtls   PartyIdentificationAndAccount5 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 DlvrrsCtdnDtls,omitempty"`
	DlvrrsIntrmyDtls PartyIdentificationAndAccount5 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 DlvrrsIntrmyDtls,omitempty"`
	DlvrgAgtDtls     PartyIdentificationAndAccount4 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 DlvrgAgtDtls"`
	SctiesSttlmSys   Max35Text                      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 SctiesSttlmSys,omitempty"`
	PlcOfSttlmDtls   PartyIdentification21          `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 PlcOfSttlmDtls,omitempty"`
}

type DeliveryParameters4 struct {
	RegdAdrInd bool                   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 RegdAdrInd"`
	NmAndAdr   NameAndAddress4        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 NmAndAdr,omitempty"`
	CtctPrsn   ContactIdentification2 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 CtctPrsn,omitempty"`
}

// May be one of DIST, ACCU
type DistributionPolicy1Code string

type Document struct {
	TrfInInstr TransferInInstructionV04 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 TrfInInstr"`
}

// Must be at least 1 items long
type EuroclearClearstreamIdentifier string

// Must be at least 1 items long
type Extended350Code string

type Extension1 struct {
	PlcAndNm Max350Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 PlcAndNm"`
	Txt      Max350Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 Txt"`
}

type FinancialInstrument13 struct {
	Id          SecurityIdentification3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 Id"`
	Nm          Max350Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 Nm,omitempty"`
	SplmtryId   Max35Text                     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 SplmtryId,omitempty"`
	ClssTp      Max35Text                     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 ClssTp,omitempty"`
	SctiesForm  FormOfSecurity1Code           `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 SctiesForm,omitempty"`
	DstrbtnPlcy DistributionPolicy1Code       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 DstrbtnPlcy,omitempty"`
}

type FinancialInstrumentQuantity1 struct {
	Unit float64 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 Unit"`
}

// May be one of BEAR, REGD
type FormOfSecurity1Code string

type GenericIdentification1 struct {
	Id      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 Id"`
	SchmeNm Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 SchmeNm,omitempty"`
	Issr    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 Issr,omitempty"`
}

type GenericIdentification27 struct {
	Id      Max4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 Id"`
	SchmeNm Max4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 SchmeNm,omitempty"`
	Issr    Max4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 Issr"`
}

// May be one of INVP, SWIP, PLAR
type HoldingsPlanType1Code string

// Must match the pattern [A-Z0-9]{12,12}
type ISINIdentifier string

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

// May be one of CASH, DRIP
type IncomePreference1Code string

type Intermediary10 struct {
	Id         PartyIdentification2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 Id"`
	Acct       Account7                   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 Acct,omitempty"`
	Role       InvestmentFundRole2Code    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 Role,omitempty"`
	XtndedRole Extended350Code            `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 XtndedRole,omitempty"`
	CtctPrsn   ContactIdentification2     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 CtctPrsn,omitempty"`
}

type Intermediary11 struct {
	Id         PartyIdentification2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 Id"`
	Acct       Account7                   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 Acct,omitempty"`
	Role       InvestmentFundRole2Code    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 Role,omitempty"`
	XtndedRole Extended350Code            `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 XtndedRole,omitempty"`
}

type InvestmentAccount22 struct {
	OwnrId               []PartyIdentification2Choice            `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 OwnrId,omitempty"`
	AcctId               AccountIdentification1                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 AcctId"`
	AcctNm               Max35Text                               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 AcctNm,omitempty"`
	AcctDsgnt            Max35Text                               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 AcctDsgnt,omitempty"`
	IntrmyInf            []Intermediary11                        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 IntrmyInf,omitempty"`
	SctiesForm           FormOfSecurity1Code                     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 SctiesForm,omitempty"`
	DmtrlsdInd           bool                                    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 DmtrlsdInd,omitempty"`
	IncmPref             IncomePreference1Code                   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 IncmPref,omitempty"`
	BnfcryCertfctnCmpltn BeneficiaryCertificationCompletion1Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 BnfcryCertfctnCmpltn,omitempty"`
	SfkpgPlc             PartyIdentification2Choice              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 SfkpgPlc,omitempty"`
	AcctSvcr             PartyIdentification2Choice              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 AcctSvcr,omitempty"`
}

type InvestmentAccount24 struct {
	OwnrId               []PartyIdentification2Choice            `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 OwnrId,omitempty"`
	AcctId               AccountIdentification1                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 AcctId,omitempty"`
	AcctNm               Max35Text                               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 AcctNm,omitempty"`
	AcctDsgnt            Max35Text                               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 AcctDsgnt,omitempty"`
	IntrmyInf            []Intermediary10                        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 IntrmyInf,omitempty"`
	SctiesForm           FormOfSecurity1Code                     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 SctiesForm,omitempty"`
	DmtrlsdInd           bool                                    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 DmtrlsdInd,omitempty"`
	IncmPref             IncomePreference1Code                   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 IncmPref,omitempty"`
	BnfcryCertfctnCmpltn BeneficiaryCertificationCompletion1Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 BnfcryCertfctnCmpltn,omitempty"`
	AcctSvcr             PartyIdentification2Choice              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 AcctSvcr,omitempty"`
	SubAcctDtls          SubAccount1                             `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 SubAcctDtls,omitempty"`
}

// May be one of FMCO, REGI, TRAG, INTR, DIST, CONC, UCL1, UCL2, TRAN
type InvestmentFundRole2Code string

type LongPostalAddress1Choice struct {
	Ustrd Max140Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 Ustrd"`
	Strd  StructuredLongPostalAddress1 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 Strd"`
}

// Must be at least 1 items long
type Max140Text string

// Must be at least 1 items long
type Max16Text string

// Must be at least 1 items long
type Max256Text string

// Must be at least 1 items long
type Max350Text string

// Must be at least 1 items long
type Max35Text string

// Must match the pattern [a-zA-Z0-9]{1,4}
type Max4AlphaNumericText string

// Must be at least 1 items long
type Max70Text string

type MessageIdentification1 struct {
	Id      Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 Id"`
	CreDtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 CreDtTm"`
}

type NameAndAddress2 struct {
	Nm  Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 Nm"`
	Adr LongPostalAddress1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 Adr,omitempty"`
}

type NameAndAddress4 struct {
	Nm  Max350Text     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 Nm,omitempty"`
	Adr PostalAddress1 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 Adr"`
}

type NameAndAddress5 struct {
	Nm  Max350Text     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 Nm"`
	Adr PostalAddress1 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 Adr,omitempty"`
}

// May be one of DOCT, MIST, MISS, MADM
type NamePrefix1Code string

type PartyIdentification1Choice struct {
	BICOrBEI AnyBICIdentifier       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 BICOrBEI"`
	PrtryId  GenericIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 PrtryId"`
	NmAndAdr NameAndAddress2        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 NmAndAdr"`
}

type PartyIdentification21 struct {
	PtyId    PartyIdentification2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 PtyId"`
	PrcgRef  Max35Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 PrcgRef,omitempty"`
	PrcgDt   DateAndDateTimeChoice      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 PrcgDt,omitempty"`
	CtctPrsn ContactIdentification2     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 CtctPrsn,omitempty"`
}

type PartyIdentification2Choice struct {
	BICOrBEI AnyBICIdentifier       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 BICOrBEI"`
	PrtryId  GenericIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 PrtryId"`
	NmAndAdr NameAndAddress5        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 NmAndAdr"`
}

type PartyIdentificationAndAccount4 struct {
	PtyId       PartyIdentification2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 PtyId"`
	AcctId      AccountIdentification1     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 AcctId,omitempty"`
	PrcgRef     Max35Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 PrcgRef,omitempty"`
	PrcgDt      DateAndDateTimeChoice      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 PrcgDt,omitempty"`
	SubAcctDtls SubAccount1                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 SubAcctDtls,omitempty"`
	CtctPrsn    ContactIdentification2     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 CtctPrsn,omitempty"`
}

type PartyIdentificationAndAccount5 struct {
	PtyId   PartyIdentification2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 PtyId"`
	AcctId  AccountIdentification1     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 AcctId,omitempty"`
	PrcgRef Max35Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 PrcgRef,omitempty"`
	PrcgDt  DateAndDateTimeChoice      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 PrcgDt,omitempty"`
}

type PartyIdentificationAndAccount93 struct {
	PtyId      PartyIdentification2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 PtyId,omitempty"`
	AcctId     AccountIdentification1     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 AcctId,omitempty"`
	PlcOfSttlm PartyIdentification2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 PlcOfSttlm"`
}

// Must match the pattern \+[0-9]{1,3}-[0-9()+\-]{1,30}
type PhoneNumber string

// May be one of DEMT, PHYS
type PhysicalTransferType1Code string

type PostalAddress1 struct {
	AdrTp       AddressType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 AdrTp,omitempty"`
	AdrLine     []Max70Text      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 AdrLine,omitempty"`
	StrtNm      Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 StrtNm,omitempty"`
	BldgNb      Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 BldgNb,omitempty"`
	PstCd       Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 PstCd,omitempty"`
	TwnNm       Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 TwnNm,omitempty"`
	CtrySubDvsn Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 CtrySubDvsn,omitempty"`
	Ctry        CountryCode      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 Ctry"`
}

// Must be at least 1 items long
type RICIdentifier string

type SecurityIdentification3Choice struct {
	ISIN        ISINIdentifier                        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 ISIN"`
	SEDOL       string                                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 SEDOL"`
	CUSIP       string                                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 CUSIP"`
	RIC         RICIdentifier                         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 RIC"`
	TckrSymb    TickerIdentifier                      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 TckrSymb"`
	Blmbrg      BloombergIdentifier                   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 Blmbrg"`
	CTA         ConsolidatedTapeAssociationIdentifier `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 CTA"`
	QUICK       string                                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 QUICK"`
	Wrtppr      string                                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 Wrtppr"`
	Dtch        string                                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 Dtch"`
	Vlrn        string                                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 Vlrn"`
	SCVM        string                                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 SCVM"`
	Belgn       string                                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 Belgn"`
	Cmon        EuroclearClearstreamIdentifier        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 Cmon"`
	OthrPrtryId AlternateSecurityIdentification1      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 OthrPrtryId"`
}

// May be one of ASAP, ENDC, WHIF
type SettlementDate1Code string

type SimpleIdentificationInformation struct {
	Id Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 Id"`
}

// May be one of ASTD, SDRN
type StampDutyType2Code string

type StructuredLongPostalAddress1 struct {
	BldgNm     Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 BldgNm,omitempty"`
	StrtNm     Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 StrtNm,omitempty"`
	StrtBldgId Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 StrtBldgId,omitempty"`
	Flr        Max16Text   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 Flr,omitempty"`
	TwnNm      Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 TwnNm"`
	DstrctNm   Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 DstrctNm,omitempty"`
	RgnId      Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 RgnId,omitempty"`
	Stat       Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 Stat,omitempty"`
	CtyId      Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 CtyId,omitempty"`
	Ctry       CountryCode `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 Ctry"`
	PstCdId    Max16Text   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 PstCdId"`
	POB        Max16Text   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 POB,omitempty"`
}

type SubAccount1 struct {
	Id    AccountIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 Id"`
	Nm    Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 Nm,omitempty"`
	Chrtc Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 Chrtc,omitempty"`
}

type Tax15 struct {
	Tp             TaxType13Code                              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 Tp"`
	XtndedTp       Extended350Code                            `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 XtndedTp"`
	Amt            ActiveOrHistoricCurrencyAnd13DecimalAmount `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 Amt"`
	Bsis           TaxationBasis2Code                         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 Bsis,omitempty"`
	XtndedBsis     Extended350Code                            `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 XtndedBsis,omitempty"`
	RcptId         PartyIdentification2Choice                 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 RcptId,omitempty"`
	XmptnInd       bool                                       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 XmptnInd"`
	XmptnRsn       TaxExemptReason1Code                       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 XmptnRsn,omitempty"`
	XtndedXmptnRsn Extended350Code                            `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 XtndedXmptnRsn,omitempty"`
}

// May be one of NONE, MASA, MISA, SISA, IISA, CUYP, PRYP, ASTR, EMPY, EMCY, EPRY, ECYE, NFPI, NFQP, DECP, IRAC, IRAR, KEOG, PFSP, 401K, SIRA, 403B, 457X, RIRA, RIAN, RCRF, RCIP, EIFP, EIOP
type TaxExemptReason1Code string

// May be one of COAX, EUTR, LOCL, NATI, LEVY, PROV, STAM, STAT, STEX, CTAX, TRAX, TRAN, VATA, WITH
type TaxType13Code string

// May be one of FLAT, PERU
type TaxationBasis2Code string

// May be one of FLAT, PERU, GRAM, NEAM
type TaxationBasis4Code string

// Must be at least 1 items long
type TickerIdentifier string

type Transfer21 struct {
	TrfRef         Max35Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 TrfRef"`
	ClntRef        Max35Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 ClntRef,omitempty"`
	CtrPtyRef      AdditionalReference2            `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 CtrPtyRef,omitempty"`
	BizFlowTp      BusinessFlowType1Code           `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 BizFlowTp,omitempty"`
	TrfRsn         TransferReason1                 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 TrfRsn,omitempty"`
	TrfDt          DateFormat1Choice               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 TrfDt,omitempty"`
	ReqdSttlmDt    ISODate                         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 ReqdSttlmDt,omitempty"`
	HldgsPlanTp    []HoldingsPlanType1Code         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 HldgsPlanTp,omitempty"`
	FinInstrmDtls  FinancialInstrument13           `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 FinInstrmDtls"`
	TtlUnitsNb     FinancialInstrumentQuantity1    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 TtlUnitsNb"`
	OwnAcctTrfInd  bool                            `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 OwnAcctTrfInd,omitempty"`
	NonStdSttlmInf Max350Text                      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 NonStdSttlmInf,omitempty"`
	RcvgAgtDtls    PartyIdentificationAndAccount93 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 RcvgAgtDtls,omitempty"`
	DlvrgAgtDtls   PartyIdentificationAndAccount93 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 DlvrgAgtDtls,omitempty"`
}

type TransferInInstructionV04 struct {
	MsgId     MessageIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 MsgId"`
	PoolRef   AdditionalReference2   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 PoolRef,omitempty"`
	PrvsRef   AdditionalReference2   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 PrvsRef,omitempty"`
	RltdRef   AdditionalReference2   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 RltdRef,omitempty"`
	MstrRef   Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 MstrRef,omitempty"`
	TrfDtls   []Transfer21           `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 TrfDtls"`
	AcctDtls  InvestmentAccount22    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 AcctDtls"`
	SttlmDtls DeliverInformation9    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 SttlmDtls,omitempty"`
	CpyDtls   CopyInformation2       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 CpyDtls,omitempty"`
	Xtnsn     []Extension1           `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 Xtnsn,omitempty"`
}

type TransferReason1 struct {
	Cd    TransferReason1Code     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 Cd"`
	Prtry GenericIdentification27 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 Prtry"`
}

// May be one of TRAU, TRAC, TRAT, TRAO, TRAI, TRAG, TPLD, TTDT, TRPE, TRAF, TRAN
type TransferReason1Code string

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
