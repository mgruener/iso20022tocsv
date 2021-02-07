// Code generated by main. DO NOT EDIT.

package sese_007_001_07

import (
	"bytes"
	"encoding/xml"
	"time"
)

type Account19 struct {
	Id    Max35Text                   `xml:"urn:swift:xsd:sese.007.001.07 Id,omitempty"`
	Dsgnt Max35Text                   `xml:"urn:swift:xsd:sese.007.001.07 Dsgnt,omitempty"`
	Svcr  PartyIdentification70Choice `xml:"urn:swift:xsd:sese.007.001.07 Svcr,omitempty"`
}

type Account20 struct {
	Id       Max35Text                   `xml:"urn:swift:xsd:sese.007.001.07 Id,omitempty"`
	AcctSvcr PartyIdentification70Choice `xml:"urn:swift:xsd:sese.007.001.07 AcctSvcr"`
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

type AdditionalReference6 struct {
	Ref     Max35Text                   `xml:"urn:swift:xsd:sese.007.001.07 Ref"`
	RefIssr PartyIdentification90Choice `xml:"urn:swift:xsd:sese.007.001.07 RefIssr,omitempty"`
	MsgNm   Max35Text                   `xml:"urn:swift:xsd:sese.007.001.07 MsgNm,omitempty"`
}

type AdditionalReference7 struct {
	Ref     Max35Text                   `xml:"urn:swift:xsd:sese.007.001.07 Ref"`
	RefIssr PartyIdentification97Choice `xml:"urn:swift:xsd:sese.007.001.07 RefIssr,omitempty"`
	MsgNm   Max35Text                   `xml:"urn:swift:xsd:sese.007.001.07 MsgNm,omitempty"`
}

// May be one of ADDR, PBOX, HOME, BIZZ, MLTO, DLVY
type AddressType2Code string

type AlternateSecurityIdentification7 struct {
	Id    Max35Text                   `xml:"urn:swift:xsd:sese.007.001.07 Id"`
	IdSrc IdentificationSource1Choice `xml:"urn:swift:xsd:sese.007.001.07 IdSrc"`
}

// Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type AnyBICIdentifier string

// May be one of NCER, ELEC, PHYS
type BeneficiaryCertificationCompletion1Code string

// Must be at least 1 items long
type BloombergIdentifier string

// May be one of SLDP, SLRP, DLPR
type BusinessFlowType1Code string

type Charge29 struct {
	Tp       ChargeType4Choice                 `xml:"urn:swift:xsd:sese.007.001.07 Tp"`
	Amt      ActiveOrHistoricCurrencyAndAmount `xml:"urn:swift:xsd:sese.007.001.07 Amt"`
	ChrgBsis ChargeBasisType1Choice            `xml:"urn:swift:xsd:sese.007.001.07 ChrgBsis,omitempty"`
	ChrgBr   ChargeBearer1Code                 `xml:"urn:swift:xsd:sese.007.001.07 ChrgBr,omitempty"`
	RcptId   PartyIdentification70Choice       `xml:"urn:swift:xsd:sese.007.001.07 RcptId,omitempty"`
}

type ChargeBasisType1Choice struct {
	Cd    TaxationBasis2Code      `xml:"urn:swift:xsd:sese.007.001.07 Cd"`
	Prtry GenericIdentification47 `xml:"urn:swift:xsd:sese.007.001.07 Prtry"`
}

// May be one of OUR, BEN, SHA
type ChargeBearer1Code string

type ChargePaymentMethod1Choice struct {
	Cd    ChargePaymentMethod1Code `xml:"urn:swift:xsd:sese.007.001.07 Cd"`
	Prtry GenericIdentification47  `xml:"urn:swift:xsd:sese.007.001.07 Prtry"`
}

// May be one of CASH, UNIT
type ChargePaymentMethod1Code string

// May be one of BEND, DISC, FEND, POST, REGF, SHIP, SPCN, TRAN
type ChargeType12Code string

type ChargeType4Choice struct {
	Cd    ChargeType12Code        `xml:"urn:swift:xsd:sese.007.001.07 Cd"`
	Prtry GenericIdentification47 `xml:"urn:swift:xsd:sese.007.001.07 Prtry"`
}

type Commission23 struct {
	Tp             CommissionType3Choice            `xml:"urn:swift:xsd:sese.007.001.07 Tp"`
	Bsis           CommissionBasis1Choice           `xml:"urn:swift:xsd:sese.007.001.07 Bsis,omitempty"`
	Amt            ActiveCurrencyAnd13DecimalAmount `xml:"urn:swift:xsd:sese.007.001.07 Amt"`
	RcptId         PartyIdentification70Choice      `xml:"urn:swift:xsd:sese.007.001.07 RcptId,omitempty"`
	ComrclAgrmtRef Max35Text                        `xml:"urn:swift:xsd:sese.007.001.07 ComrclAgrmtRef,omitempty"`
	WvgDtls        CommissionWaiver4                `xml:"urn:swift:xsd:sese.007.001.07 WvgDtls,omitempty"`
}

type CommissionBasis1Choice struct {
	Cd    TaxationBasis4Code      `xml:"urn:swift:xsd:sese.007.001.07 Cd"`
	Prtry GenericIdentification47 `xml:"urn:swift:xsd:sese.007.001.07 Prtry"`
}

type CommissionType3Choice struct {
	Cd    CommissionType7Code     `xml:"urn:swift:xsd:sese.007.001.07 Cd"`
	Prtry GenericIdentification47 `xml:"urn:swift:xsd:sese.007.001.07 Prtry"`
}

// May be one of FEND, BEND
type CommissionType7Code string

type CommissionWaiver4 struct {
	InstrBsis WaivingInstruction1Choice `xml:"urn:swift:xsd:sese.007.001.07 InstrBsis"`
	WvdRate   float64                   `xml:"urn:swift:xsd:sese.007.001.07 WvdRate"`
}

// Must be at least 1 items long
type ConsolidatedTapeAssociationIdentifier string

type ContactIdentification2 struct {
	NmPrfx   NamePrefix1Code `xml:"urn:swift:xsd:sese.007.001.07 NmPrfx,omitempty"`
	GvnNm    Max35Text       `xml:"urn:swift:xsd:sese.007.001.07 GvnNm,omitempty"`
	Nm       Max35Text       `xml:"urn:swift:xsd:sese.007.001.07 Nm"`
	PhneNb   PhoneNumber     `xml:"urn:swift:xsd:sese.007.001.07 PhneNb,omitempty"`
	MobNb    PhoneNumber     `xml:"urn:swift:xsd:sese.007.001.07 MobNb,omitempty"`
	FaxNb    PhoneNumber     `xml:"urn:swift:xsd:sese.007.001.07 FaxNb,omitempty"`
	EmailAdr Max256Text      `xml:"urn:swift:xsd:sese.007.001.07 EmailAdr,omitempty"`
}

type CopyInformation4 struct {
	CpyInd    bool             `xml:"urn:swift:xsd:sese.007.001.07 CpyInd"`
	OrgnlRcvr AnyBICIdentifier `xml:"urn:swift:xsd:sese.007.001.07 OrgnlRcvr,omitempty"`
}

// Must match the pattern [A-Z]{2,2}
type CountryCode string

type DateAndDateTimeChoice struct {
	Dt   ISODate     `xml:"urn:swift:xsd:sese.007.001.07 Dt"`
	DtTm ISODateTime `xml:"urn:swift:xsd:sese.007.001.07 DtTm"`
}

type DeliverInformation17 struct {
	Trfr           PartyIdentification70Choice   `xml:"urn:swift:xsd:sese.007.001.07 Trfr,omitempty"`
	TrfrRegdAcct   Account19                     `xml:"urn:swift:xsd:sese.007.001.07 TrfrRegdAcct,omitempty"`
	IntrmyInf      []Intermediary34              `xml:"urn:swift:xsd:sese.007.001.07 IntrmyInf,omitempty"`
	ReqdSttlmDt    ISODate                       `xml:"urn:swift:xsd:sese.007.001.07 ReqdSttlmDt,omitempty"`
	FctvSttlmDt    DateAndDateTimeChoice         `xml:"urn:swift:xsd:sese.007.001.07 FctvSttlmDt,omitempty"`
	SttlmAmt       ActiveCurrencyAndAmount       `xml:"urn:swift:xsd:sese.007.001.07 SttlmAmt,omitempty"`
	StmpDty        StampDutyType2Code            `xml:"urn:swift:xsd:sese.007.001.07 StmpDty,omitempty"`
	NetAmt         ActiveCurrencyAndAmount       `xml:"urn:swift:xsd:sese.007.001.07 NetAmt,omitempty"`
	ChrgDtls       []Charge29                    `xml:"urn:swift:xsd:sese.007.001.07 ChrgDtls,omitempty"`
	ComssnDtls     []Commission23                `xml:"urn:swift:xsd:sese.007.001.07 ComssnDtls,omitempty"`
	TaxDtls        []Tax28                       `xml:"urn:swift:xsd:sese.007.001.07 TaxDtls,omitempty"`
	FXDtls         []ForeignExchangeTerms26      `xml:"urn:swift:xsd:sese.007.001.07 FXDtls,omitempty"`
	SttlmPtiesDtls DeliveringPartiesAndAccount13 `xml:"urn:swift:xsd:sese.007.001.07 SttlmPtiesDtls,omitempty"`
	PhysTrf        PhysicalTransferType1Code     `xml:"urn:swift:xsd:sese.007.001.07 PhysTrf,omitempty"`
	PhysTrfDtls    DeliveryParameters4           `xml:"urn:swift:xsd:sese.007.001.07 PhysTrfDtls,omitempty"`
	ClntRef        AdditionalReference7          `xml:"urn:swift:xsd:sese.007.001.07 ClntRef,omitempty"`
}

type DeliveringPartiesAndAccount13 struct {
	DlvrrDtls         InvestmentAccount55              `xml:"urn:swift:xsd:sese.007.001.07 DlvrrDtls,omitempty"`
	DlvrrsCtdnDtls    PartyIdentificationAndAccount124 `xml:"urn:swift:xsd:sese.007.001.07 DlvrrsCtdnDtls,omitempty"`
	DlvrrsIntrmy1Dtls PartyIdentificationAndAccount124 `xml:"urn:swift:xsd:sese.007.001.07 DlvrrsIntrmy1Dtls,omitempty"`
	DlvrrsIntrmy2Dtls PartyIdentificationAndAccount124 `xml:"urn:swift:xsd:sese.007.001.07 DlvrrsIntrmy2Dtls,omitempty"`
	DlvrgAgtDtls      PartyIdentificationAndAccount124 `xml:"urn:swift:xsd:sese.007.001.07 DlvrgAgtDtls"`
	SctiesSttlmSys    Max35Text                        `xml:"urn:swift:xsd:sese.007.001.07 SctiesSttlmSys,omitempty"`
	PlcOfSttlmDtls    PartyIdentification97            `xml:"urn:swift:xsd:sese.007.001.07 PlcOfSttlmDtls,omitempty"`
}

type DeliveryParameters4 struct {
	RegdAdrInd bool                   `xml:"urn:swift:xsd:sese.007.001.07 RegdAdrInd"`
	NmAndAdr   NameAndAddress4        `xml:"urn:swift:xsd:sese.007.001.07 NmAndAdr,omitempty"`
	CtctPrsn   ContactIdentification2 `xml:"urn:swift:xsd:sese.007.001.07 CtctPrsn,omitempty"`
}

// May be one of DIST, ACCU
type DistributionPolicy1Code string

type Document struct {
	TrfInConf TransferInConfirmationV07 `xml:"urn:swift:xsd:sese.007.001.07 TrfInConf"`
}

// Must be at least 1 items long
type EuroclearClearstreamIdentifier string

// Must match the pattern [a-zA-Z0-9]{4}
type Exact4AlphaNumericText string

type ExemptionReason1Choice struct {
	Cd    TaxExemptReason1Code    `xml:"urn:swift:xsd:sese.007.001.07 Cd"`
	Prtry GenericIdentification47 `xml:"urn:swift:xsd:sese.007.001.07 Prtry"`
}

type Extension1 struct {
	PlcAndNm Max350Text `xml:"urn:swift:xsd:sese.007.001.07 PlcAndNm"`
	Txt      Max350Text `xml:"urn:swift:xsd:sese.007.001.07 Txt"`
}

type FinancialInstrument49 struct {
	Id          SecurityIdentification23Choice `xml:"urn:swift:xsd:sese.007.001.07 Id"`
	Nm          Max350Text                     `xml:"urn:swift:xsd:sese.007.001.07 Nm,omitempty"`
	ShrtNm      Max35Text                      `xml:"urn:swift:xsd:sese.007.001.07 ShrtNm,omitempty"`
	SplmtryId   Max35Text                      `xml:"urn:swift:xsd:sese.007.001.07 SplmtryId,omitempty"`
	ClssTp      Max35Text                      `xml:"urn:swift:xsd:sese.007.001.07 ClssTp,omitempty"`
	SctiesForm  FormOfSecurity1Code            `xml:"urn:swift:xsd:sese.007.001.07 SctiesForm,omitempty"`
	DstrbtnPlcy DistributionPolicy1Code        `xml:"urn:swift:xsd:sese.007.001.07 DstrbtnPlcy,omitempty"`
}

type FinancialInstrumentQuantity1 struct {
	Unit float64 `xml:"urn:swift:xsd:sese.007.001.07 Unit"`
}

type ForeignExchangeTerms26 struct {
	ToAmt    ActiveCurrencyAnd13DecimalAmount `xml:"urn:swift:xsd:sese.007.001.07 ToAmt,omitempty"`
	FrAmt    ActiveCurrencyAndAmount          `xml:"urn:swift:xsd:sese.007.001.07 FrAmt,omitempty"`
	UnitCcy  ActiveOrHistoricCurrencyCode     `xml:"urn:swift:xsd:sese.007.001.07 UnitCcy"`
	QtdCcy   ActiveOrHistoricCurrencyCode     `xml:"urn:swift:xsd:sese.007.001.07 QtdCcy"`
	XchgRate float64                          `xml:"urn:swift:xsd:sese.007.001.07 XchgRate"`
	QtnDt    ISODateTime                      `xml:"urn:swift:xsd:sese.007.001.07 QtnDt,omitempty"`
	QtgInstn PartyIdentification70Choice      `xml:"urn:swift:xsd:sese.007.001.07 QtgInstn,omitempty"`
}

// May be one of BEAR, REGD
type FormOfSecurity1Code string

type GenericIdentification1 struct {
	Id      Max35Text `xml:"urn:swift:xsd:sese.007.001.07 Id"`
	SchmeNm Max35Text `xml:"urn:swift:xsd:sese.007.001.07 SchmeNm,omitempty"`
	Issr    Max35Text `xml:"urn:swift:xsd:sese.007.001.07 Issr,omitempty"`
}

type GenericIdentification27 struct {
	Id      Max4AlphaNumericText `xml:"urn:swift:xsd:sese.007.001.07 Id"`
	SchmeNm Max4AlphaNumericText `xml:"urn:swift:xsd:sese.007.001.07 SchmeNm,omitempty"`
	Issr    Max4AlphaNumericText `xml:"urn:swift:xsd:sese.007.001.07 Issr"`
}

type GenericIdentification30 struct {
	Id      Exact4AlphaNumericText `xml:"urn:swift:xsd:sese.007.001.07 Id"`
	Issr    Max35Text              `xml:"urn:swift:xsd:sese.007.001.07 Issr"`
	SchmeNm Max35Text              `xml:"urn:swift:xsd:sese.007.001.07 SchmeNm,omitempty"`
}

type GenericIdentification47 struct {
	Id      Exact4AlphaNumericText `xml:"urn:swift:xsd:sese.007.001.07 Id"`
	Issr    Max4AlphaNumericText   `xml:"urn:swift:xsd:sese.007.001.07 Issr"`
	SchmeNm Max4AlphaNumericText   `xml:"urn:swift:xsd:sese.007.001.07 SchmeNm,omitempty"`
}

type GenericIdentification78 struct {
	Tp GenericIdentification30 `xml:"urn:swift:xsd:sese.007.001.07 Tp"`
	Id Max35Text               `xml:"urn:swift:xsd:sese.007.001.07 Id,omitempty"`
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

type ISOYearMonth time.Time

func (t *ISOYearMonth) UnmarshalText(text []byte) error {
	return (*xsdGYearMonth)(t).UnmarshalText(text)
}
func (t ISOYearMonth) MarshalText() ([]byte, error) {
	return xsdGYearMonth(t).MarshalText()
}

type IdentificationSource1Choice struct {
	Dmst  CountryCode `xml:"urn:swift:xsd:sese.007.001.07 Dmst"`
	Prtry Max35Text   `xml:"urn:swift:xsd:sese.007.001.07 Prtry"`
}

// May be one of CASH, SECU
type IncomePreference2Code string

type Intermediary34 struct {
	Id       PartyIdentification70Choice `xml:"urn:swift:xsd:sese.007.001.07 Id"`
	Acct     Account20                   `xml:"urn:swift:xsd:sese.007.001.07 Acct,omitempty"`
	Role     Role4Choice                 `xml:"urn:swift:xsd:sese.007.001.07 Role,omitempty"`
	CtctPrsn ContactIdentification2      `xml:"urn:swift:xsd:sese.007.001.07 CtctPrsn,omitempty"`
}

type InvestmentAccount55 struct {
	OwnrId               []PartyIdentification70Choice           `xml:"urn:swift:xsd:sese.007.001.07 OwnrId,omitempty"`
	AcctId               Max35Text                               `xml:"urn:swift:xsd:sese.007.001.07 AcctId,omitempty"`
	AcctNm               Max35Text                               `xml:"urn:swift:xsd:sese.007.001.07 AcctNm,omitempty"`
	AcctDsgnt            Max35Text                               `xml:"urn:swift:xsd:sese.007.001.07 AcctDsgnt,omitempty"`
	SctiesForm           FormOfSecurity1Code                     `xml:"urn:swift:xsd:sese.007.001.07 SctiesForm,omitempty"`
	DmtrlsdInd           bool                                    `xml:"urn:swift:xsd:sese.007.001.07 DmtrlsdInd,omitempty"`
	IncmPref             IncomePreference2Code                   `xml:"urn:swift:xsd:sese.007.001.07 IncmPref,omitempty"`
	BnfcryCertfctnCmpltn BeneficiaryCertificationCompletion1Code `xml:"urn:swift:xsd:sese.007.001.07 BnfcryCertfctnCmpltn,omitempty"`
	AcctSvcr             PartyIdentification70Choice             `xml:"urn:swift:xsd:sese.007.001.07 AcctSvcr,omitempty"`
	SubAcctDtls          SubAccount5                             `xml:"urn:swift:xsd:sese.007.001.07 SubAcctDtls,omitempty"`
}

type InvestmentAccount56 struct {
	OwnrId               []PartyIdentification70Choice           `xml:"urn:swift:xsd:sese.007.001.07 OwnrId,omitempty"`
	AcctId               Max35Text                               `xml:"urn:swift:xsd:sese.007.001.07 AcctId"`
	AcctNm               Max35Text                               `xml:"urn:swift:xsd:sese.007.001.07 AcctNm,omitempty"`
	AcctDsgnt            Max35Text                               `xml:"urn:swift:xsd:sese.007.001.07 AcctDsgnt,omitempty"`
	IntrmyInf            []Intermediary34                        `xml:"urn:swift:xsd:sese.007.001.07 IntrmyInf,omitempty"`
	SctiesForm           FormOfSecurity1Code                     `xml:"urn:swift:xsd:sese.007.001.07 SctiesForm,omitempty"`
	DmtrlsdInd           bool                                    `xml:"urn:swift:xsd:sese.007.001.07 DmtrlsdInd,omitempty"`
	IncmPref             IncomePreference2Code                   `xml:"urn:swift:xsd:sese.007.001.07 IncmPref,omitempty"`
	BnfcryCertfctnCmpltn BeneficiaryCertificationCompletion1Code `xml:"urn:swift:xsd:sese.007.001.07 BnfcryCertfctnCmpltn,omitempty"`
	SfkpgPlc             SafekeepingPlaceFormat8Choice           `xml:"urn:swift:xsd:sese.007.001.07 SfkpgPlc,omitempty"`
	AcctSvcr             PartyIdentification70Choice             `xml:"urn:swift:xsd:sese.007.001.07 AcctSvcr,omitempty"`
	SubAcctDtls          SubAccount5                             `xml:"urn:swift:xsd:sese.007.001.07 SubAcctDtls,omitempty"`
	SttlmPtiesDtls       ReceivingPartiesAndAccount14            `xml:"urn:swift:xsd:sese.007.001.07 SttlmPtiesDtls,omitempty"`
}

// May be one of FMCO, REGI, TRAG, INTR, DIST, CONC, UCL1, UCL2, TRAN
type InvestmentFundRole2Code string

// Must match the pattern [A-Z0-9]{18,18}[0-9]{2,2}
type LEIIdentifier string

type MarketPracticeVersion1 struct {
	Nm Max35Text    `xml:"urn:swift:xsd:sese.007.001.07 Nm"`
	Dt ISOYearMonth `xml:"urn:swift:xsd:sese.007.001.07 Dt,omitempty"`
	Nb Max35Text    `xml:"urn:swift:xsd:sese.007.001.07 Nb,omitempty"`
}

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
	Id      Max35Text   `xml:"urn:swift:xsd:sese.007.001.07 Id"`
	CreDtTm ISODateTime `xml:"urn:swift:xsd:sese.007.001.07 CreDtTm"`
}

type NameAndAddress4 struct {
	Nm  Max350Text     `xml:"urn:swift:xsd:sese.007.001.07 Nm,omitempty"`
	Adr PostalAddress1 `xml:"urn:swift:xsd:sese.007.001.07 Adr"`
}

type NameAndAddress5 struct {
	Nm  Max350Text     `xml:"urn:swift:xsd:sese.007.001.07 Nm"`
	Adr PostalAddress1 `xml:"urn:swift:xsd:sese.007.001.07 Adr,omitempty"`
}

// May be one of DOCT, MIST, MISS, MADM
type NamePrefix1Code string

type PartyIdentification70Choice struct {
	AnyBIC   AnyBICIdentifier       `xml:"urn:swift:xsd:sese.007.001.07 AnyBIC"`
	PrtryId  GenericIdentification1 `xml:"urn:swift:xsd:sese.007.001.07 PrtryId"`
	NmAndAdr NameAndAddress5        `xml:"urn:swift:xsd:sese.007.001.07 NmAndAdr"`
}

type PartyIdentification90Choice struct {
	AnyBIC   AnyBICIdentifier       `xml:"urn:swift:xsd:sese.007.001.07 AnyBIC"`
	PrtryId  GenericIdentification1 `xml:"urn:swift:xsd:sese.007.001.07 PrtryId"`
	NmAndAdr NameAndAddress5        `xml:"urn:swift:xsd:sese.007.001.07 NmAndAdr"`
}

type PartyIdentification97 struct {
	PtyId    PartyIdentification70Choice `xml:"urn:swift:xsd:sese.007.001.07 PtyId"`
	PrcgRef  Max35Text                   `xml:"urn:swift:xsd:sese.007.001.07 PrcgRef,omitempty"`
	PrcgDt   DateAndDateTimeChoice       `xml:"urn:swift:xsd:sese.007.001.07 PrcgDt,omitempty"`
	CtctPrsn ContactIdentification2      `xml:"urn:swift:xsd:sese.007.001.07 CtctPrsn,omitempty"`
}

type PartyIdentification97Choice struct {
	AnyBIC     AnyBICIdentifier       `xml:"urn:swift:xsd:sese.007.001.07 AnyBIC"`
	LglNttyIdr LEIIdentifier          `xml:"urn:swift:xsd:sese.007.001.07 LglNttyIdr"`
	NmAndAdr   NameAndAddress5        `xml:"urn:swift:xsd:sese.007.001.07 NmAndAdr"`
	PrtryId    GenericIdentification1 `xml:"urn:swift:xsd:sese.007.001.07 PrtryId"`
}

type PartyIdentificationAndAccount123 struct {
	PtyId       PartyIdentification70Choice `xml:"urn:swift:xsd:sese.007.001.07 PtyId"`
	AcctId      Max35Text                   `xml:"urn:swift:xsd:sese.007.001.07 AcctId,omitempty"`
	PrcgRef     Max35Text                   `xml:"urn:swift:xsd:sese.007.001.07 PrcgRef,omitempty"`
	PrcgDt      DateAndDateTimeChoice       `xml:"urn:swift:xsd:sese.007.001.07 PrcgDt,omitempty"`
	SubAcctDtls SubAccount5                 `xml:"urn:swift:xsd:sese.007.001.07 SubAcctDtls,omitempty"`
	CtctPrsn    ContactIdentification2      `xml:"urn:swift:xsd:sese.007.001.07 CtctPrsn,omitempty"`
}

type PartyIdentificationAndAccount124 struct {
	PtyId   PartyIdentification70Choice `xml:"urn:swift:xsd:sese.007.001.07 PtyId"`
	AcctId  Max35Text                   `xml:"urn:swift:xsd:sese.007.001.07 AcctId,omitempty"`
	PrcgRef Max35Text                   `xml:"urn:swift:xsd:sese.007.001.07 PrcgRef,omitempty"`
	PrcgDt  DateAndDateTimeChoice       `xml:"urn:swift:xsd:sese.007.001.07 PrcgDt,omitempty"`
}

type PartyIdentificationAndAccount125 struct {
	PtyId      PartyIdentification70Choice `xml:"urn:swift:xsd:sese.007.001.07 PtyId,omitempty"`
	AcctId     Max35Text                   `xml:"urn:swift:xsd:sese.007.001.07 AcctId,omitempty"`
	PlcOfSttlm PartyIdentification70Choice `xml:"urn:swift:xsd:sese.007.001.07 PlcOfSttlm"`
}

// Must match the pattern \+[0-9]{1,3}-[0-9()+\-]{1,30}
type PhoneNumber string

// May be one of DEMT, PHYS
type PhysicalTransferType1Code string

type PostalAddress1 struct {
	AdrTp       AddressType2Code `xml:"urn:swift:xsd:sese.007.001.07 AdrTp,omitempty"`
	AdrLine     []Max70Text      `xml:"urn:swift:xsd:sese.007.001.07 AdrLine,omitempty"`
	StrtNm      Max70Text        `xml:"urn:swift:xsd:sese.007.001.07 StrtNm,omitempty"`
	BldgNb      Max16Text        `xml:"urn:swift:xsd:sese.007.001.07 BldgNb,omitempty"`
	PstCd       Max16Text        `xml:"urn:swift:xsd:sese.007.001.07 PstCd,omitempty"`
	TwnNm       Max35Text        `xml:"urn:swift:xsd:sese.007.001.07 TwnNm,omitempty"`
	CtrySubDvsn Max35Text        `xml:"urn:swift:xsd:sese.007.001.07 CtrySubDvsn,omitempty"`
	Ctry        CountryCode      `xml:"urn:swift:xsd:sese.007.001.07 Ctry"`
}

// May be one of FORW, HIST
type PriceMethod1Code string

type PriceValue1 struct {
	Amt ActiveCurrencyAnd13DecimalAmount `xml:"urn:swift:xsd:sese.007.001.07 Amt"`
}

// Must be at least 1 items long
type RICIdentifier string

type ReceivingPartiesAndAccount14 struct {
	RcvrsCtdnDtls    PartyIdentificationAndAccount124 `xml:"urn:swift:xsd:sese.007.001.07 RcvrsCtdnDtls,omitempty"`
	RcvrsIntrmy1Dtls PartyIdentificationAndAccount124 `xml:"urn:swift:xsd:sese.007.001.07 RcvrsIntrmy1Dtls,omitempty"`
	RcvrsIntrmy2Dtls PartyIdentificationAndAccount124 `xml:"urn:swift:xsd:sese.007.001.07 RcvrsIntrmy2Dtls,omitempty"`
	RcvgAgtDtls      PartyIdentificationAndAccount123 `xml:"urn:swift:xsd:sese.007.001.07 RcvgAgtDtls"`
	SctiesSttlmSys   Max35Text                        `xml:"urn:swift:xsd:sese.007.001.07 SctiesSttlmSys,omitempty"`
	PlcOfSttlmDtls   PartyIdentification97            `xml:"urn:swift:xsd:sese.007.001.07 PlcOfSttlmDtls,omitempty"`
}

type Role4Choice struct {
	Cd    InvestmentFundRole2Code `xml:"urn:swift:xsd:sese.007.001.07 Cd"`
	Prtry GenericIdentification47 `xml:"urn:swift:xsd:sese.007.001.07 Prtry"`
}

// May be one of CUST, ICSD, NCSD, SHHE
type SafekeepingPlace1Code string

// May be one of SHHE, ALLP
type SafekeepingPlace2Code string

type SafekeepingPlaceFormat8Choice struct {
	Id      SafekeepingPlaceTypeAndText6             `xml:"urn:swift:xsd:sese.007.001.07 Id"`
	Ctry    CountryCode                              `xml:"urn:swift:xsd:sese.007.001.07 Ctry"`
	TpAndId SafekeepingPlaceTypeAndAnyBICIdentifier1 `xml:"urn:swift:xsd:sese.007.001.07 TpAndId"`
	Prtry   GenericIdentification78                  `xml:"urn:swift:xsd:sese.007.001.07 Prtry"`
}

type SafekeepingPlaceTypeAndAnyBICIdentifier1 struct {
	SfkpgPlcTp SafekeepingPlace1Code `xml:"urn:swift:xsd:sese.007.001.07 SfkpgPlcTp"`
	Id         AnyBICIdentifier      `xml:"urn:swift:xsd:sese.007.001.07 Id"`
}

type SafekeepingPlaceTypeAndText6 struct {
	SfkpgPlcTp SafekeepingPlace2Code `xml:"urn:swift:xsd:sese.007.001.07 SfkpgPlcTp"`
	Id         Max35Text             `xml:"urn:swift:xsd:sese.007.001.07 Id,omitempty"`
}

type SecurityIdentification23Choice struct {
	ISIN        ISINIdentifier                        `xml:"urn:swift:xsd:sese.007.001.07 ISIN"`
	SEDOL       string                                `xml:"urn:swift:xsd:sese.007.001.07 SEDOL"`
	CUSIP       string                                `xml:"urn:swift:xsd:sese.007.001.07 CUSIP"`
	RIC         RICIdentifier                         `xml:"urn:swift:xsd:sese.007.001.07 RIC"`
	TckrSymb    TickerIdentifier                      `xml:"urn:swift:xsd:sese.007.001.07 TckrSymb"`
	Blmbrg      BloombergIdentifier                   `xml:"urn:swift:xsd:sese.007.001.07 Blmbrg"`
	CTA         ConsolidatedTapeAssociationIdentifier `xml:"urn:swift:xsd:sese.007.001.07 CTA"`
	QUICK       string                                `xml:"urn:swift:xsd:sese.007.001.07 QUICK"`
	Wrtppr      string                                `xml:"urn:swift:xsd:sese.007.001.07 Wrtppr"`
	Dtch        string                                `xml:"urn:swift:xsd:sese.007.001.07 Dtch"`
	Vlrn        string                                `xml:"urn:swift:xsd:sese.007.001.07 Vlrn"`
	SCVM        string                                `xml:"urn:swift:xsd:sese.007.001.07 SCVM"`
	Belgn       string                                `xml:"urn:swift:xsd:sese.007.001.07 Belgn"`
	Cmon        EuroclearClearstreamIdentifier        `xml:"urn:swift:xsd:sese.007.001.07 Cmon"`
	OthrPrtryId AlternateSecurityIdentification7      `xml:"urn:swift:xsd:sese.007.001.07 OthrPrtryId"`
}

// May be one of ASTD, SDRN
type StampDutyType2Code string

type SubAccount5 struct {
	Id    Max35Text `xml:"urn:swift:xsd:sese.007.001.07 Id"`
	Nm    Max35Text `xml:"urn:swift:xsd:sese.007.001.07 Nm,omitempty"`
	Chrtc Max35Text `xml:"urn:swift:xsd:sese.007.001.07 Chrtc,omitempty"`
}

type Tax28 struct {
	Tp           TaxType1Choice                             `xml:"urn:swift:xsd:sese.007.001.07 Tp"`
	Amt          ActiveOrHistoricCurrencyAnd13DecimalAmount `xml:"urn:swift:xsd:sese.007.001.07 Amt"`
	Bsis         TaxBasis1Choice                            `xml:"urn:swift:xsd:sese.007.001.07 Bsis,omitempty"`
	RcptId       PartyIdentification70Choice                `xml:"urn:swift:xsd:sese.007.001.07 RcptId,omitempty"`
	XmptnInd     bool                                       `xml:"urn:swift:xsd:sese.007.001.07 XmptnInd"`
	XmptnRsn     ExemptionReason1Choice                     `xml:"urn:swift:xsd:sese.007.001.07 XmptnRsn,omitempty"`
	TaxClctnDtls TaxCalculationInformation8                 `xml:"urn:swift:xsd:sese.007.001.07 TaxClctnDtls,omitempty"`
}

type TaxBasis1Choice struct {
	Cd    TaxationBasis2Code      `xml:"urn:swift:xsd:sese.007.001.07 Cd"`
	Prtry GenericIdentification47 `xml:"urn:swift:xsd:sese.007.001.07 Prtry"`
}

type TaxCalculationInformation8 struct {
	Bsis     TaxBasis1Choice                  `xml:"urn:swift:xsd:sese.007.001.07 Bsis,omitempty"`
	TaxblAmt ActiveCurrencyAnd13DecimalAmount `xml:"urn:swift:xsd:sese.007.001.07 TaxblAmt"`
}

// May be one of NONE, MASA, MISA, SISA, IISA, CUYP, PRYP, ASTR, EMPY, EMCY, EPRY, ECYE, NFPI, NFQP, DECP, IRAC, IRAR, KEOG, PFSP, 401K, SIRA, 403B, 457X, RIRA, RIAN, RCRF, RCIP, EIFP, EIOP
type TaxExemptReason1Code string

// May be one of COAX, CTAX, EUTR, LEVY, LOCL, NATI, PROV, STAM, STAT, STEX, TRAN, TRAX, VATA, WITH, NKAP, KAPA
type TaxType16Code string

type TaxType1Choice struct {
	Cd    TaxType16Code           `xml:"urn:swift:xsd:sese.007.001.07 Cd"`
	Prtry GenericIdentification47 `xml:"urn:swift:xsd:sese.007.001.07 Prtry"`
}

// May be one of FLAT, PERU
type TaxationBasis2Code string

// May be one of FLAT, PERU, GRAM, NEAM
type TaxationBasis4Code string

// Must be at least 1 items long
type TickerIdentifier string

type Transfer33 struct {
	TrfConfRef     Max35Text                                  `xml:"urn:swift:xsd:sese.007.001.07 TrfConfRef"`
	TrfRef         Max35Text                                  `xml:"urn:swift:xsd:sese.007.001.07 TrfRef"`
	ClntRef        AdditionalReference7                       `xml:"urn:swift:xsd:sese.007.001.07 ClntRef,omitempty"`
	CtrPtyRef      AdditionalReference7                       `xml:"urn:swift:xsd:sese.007.001.07 CtrPtyRef,omitempty"`
	BizFlowTp      BusinessFlowType1Code                      `xml:"urn:swift:xsd:sese.007.001.07 BizFlowTp,omitempty"`
	FctvTrfDt      DateAndDateTimeChoice                      `xml:"urn:swift:xsd:sese.007.001.07 FctvTrfDt,omitempty"`
	ReqdSttlmDt    ISODate                                    `xml:"urn:swift:xsd:sese.007.001.07 ReqdSttlmDt,omitempty"`
	FctvSttlmDt    DateAndDateTimeChoice                      `xml:"urn:swift:xsd:sese.007.001.07 FctvSttlmDt,omitempty"`
	TradDt         DateAndDateTimeChoice                      `xml:"urn:swift:xsd:sese.007.001.07 TradDt,omitempty"`
	TrfOrdrDtForm  ISODate                                    `xml:"urn:swift:xsd:sese.007.001.07 TrfOrdrDtForm,omitempty"`
	TrfRsn         TransferReason1                            `xml:"urn:swift:xsd:sese.007.001.07 TrfRsn,omitempty"`
	HldgsPlanTp    []HoldingsPlanType1Code                    `xml:"urn:swift:xsd:sese.007.001.07 HldgsPlanTp,omitempty"`
	FinInstrmDtls  FinancialInstrument49                      `xml:"urn:swift:xsd:sese.007.001.07 FinInstrmDtls"`
	TtlUnitsNb     FinancialInstrumentQuantity1               `xml:"urn:swift:xsd:sese.007.001.07 TtlUnitsNb"`
	UnitsDtls      []Unit6                                    `xml:"urn:swift:xsd:sese.007.001.07 UnitsDtls,omitempty"`
	AvrgPric       ActiveOrHistoricCurrencyAnd13DecimalAmount `xml:"urn:swift:xsd:sese.007.001.07 AvrgPric,omitempty"`
	NewAvrgPric    ActiveOrHistoricCurrencyAnd13DecimalAmount `xml:"urn:swift:xsd:sese.007.001.07 NewAvrgPric,omitempty"`
	AvrgDt         ISODate                                    `xml:"urn:swift:xsd:sese.007.001.07 AvrgDt,omitempty"`
	NewAvrgDt      ISODate                                    `xml:"urn:swift:xsd:sese.007.001.07 NewAvrgDt,omitempty"`
	TrfCcy         ActiveOrHistoricCurrencyCode               `xml:"urn:swift:xsd:sese.007.001.07 TrfCcy,omitempty"`
	OwnAcctTrfInd  bool                                       `xml:"urn:swift:xsd:sese.007.001.07 OwnAcctTrfInd,omitempty"`
	NonStdSttlmInf Max350Text                                 `xml:"urn:swift:xsd:sese.007.001.07 NonStdSttlmInf,omitempty"`
	RcvgAgtDtls    PartyIdentificationAndAccount125           `xml:"urn:swift:xsd:sese.007.001.07 RcvgAgtDtls,omitempty"`
	DlvrgAgtDtls   PartyIdentificationAndAccount125           `xml:"urn:swift:xsd:sese.007.001.07 DlvrgAgtDtls,omitempty"`
	TrfExpnssPmtTp ChargePaymentMethod1Choice                 `xml:"urn:swift:xsd:sese.007.001.07 TrfExpnssPmtTp,omitempty"`
}

type TransferInConfirmationV07 struct {
	MsgId        MessageIdentification1  `xml:"urn:swift:xsd:sese.007.001.07 MsgId"`
	PoolRef      AdditionalReference6    `xml:"urn:swift:xsd:sese.007.001.07 PoolRef,omitempty"`
	PrvsRef      AdditionalReference6    `xml:"urn:swift:xsd:sese.007.001.07 PrvsRef,omitempty"`
	RltdRef      AdditionalReference6    `xml:"urn:swift:xsd:sese.007.001.07 RltdRef,omitempty"`
	Fctn         TransferInFunction2Code `xml:"urn:swift:xsd:sese.007.001.07 Fctn,omitempty"`
	MstrRef      Max35Text               `xml:"urn:swift:xsd:sese.007.001.07 MstrRef,omitempty"`
	TrfDtls      []Transfer33            `xml:"urn:swift:xsd:sese.007.001.07 TrfDtls"`
	AcctDtls     InvestmentAccount56     `xml:"urn:swift:xsd:sese.007.001.07 AcctDtls"`
	SttlmDtls    DeliverInformation17    `xml:"urn:swift:xsd:sese.007.001.07 SttlmDtls,omitempty"`
	MktPrctcVrsn MarketPracticeVersion1  `xml:"urn:swift:xsd:sese.007.001.07 MktPrctcVrsn,omitempty"`
	CpyDtls      CopyInformation4        `xml:"urn:swift:xsd:sese.007.001.07 CpyDtls,omitempty"`
	Xtnsn        []Extension1            `xml:"urn:swift:xsd:sese.007.001.07 Xtnsn,omitempty"`
}

// May be one of CONF, ADVI
type TransferInFunction2Code string

type TransferReason1 struct {
	Cd    TransferReason1Code     `xml:"urn:swift:xsd:sese.007.001.07 Cd"`
	Prtry GenericIdentification27 `xml:"urn:swift:xsd:sese.007.001.07 Prtry"`
}

// May be one of TRAU, TRAC, TRAT, TRAO, TRAI, TRAG, TPLD, TTDT, TRPE, TRAF, TRAN
type TransferReason1Code string

// May be one of BIDE, OFFR, NAVL, CREA, CANC, INTE, SWNG, MIDD, RINV, SWIC
type TypeOfPrice12Code string

type TypeOfPrice31Choice struct {
	Cd    TypeOfPrice12Code       `xml:"urn:swift:xsd:sese.007.001.07 Cd"`
	Prtry GenericIdentification47 `xml:"urn:swift:xsd:sese.007.001.07 Prtry"`
}

// May be one of GRP1, GRP2
type UKTaxGroupUnitCode string

type Unit6 struct {
	UnitsNb      FinancialInstrumentQuantity1 `xml:"urn:swift:xsd:sese.007.001.07 UnitsNb"`
	AcqstnDt     ISODate                      `xml:"urn:swift:xsd:sese.007.001.07 AcqstnDt,omitempty"`
	CertNb       []Max35Text                  `xml:"urn:swift:xsd:sese.007.001.07 CertNb,omitempty"`
	Grp1Or2Units UKTaxGroupUnitCode           `xml:"urn:swift:xsd:sese.007.001.07 Grp1Or2Units,omitempty"`
	PricDtls     UnitPrice21                  `xml:"urn:swift:xsd:sese.007.001.07 PricDtls,omitempty"`
}

type UnitPrice21 struct {
	Tp              TypeOfPrice31Choice               `xml:"urn:swift:xsd:sese.007.001.07 Tp"`
	Val             PriceValue1                       `xml:"urn:swift:xsd:sese.007.001.07 Val"`
	PricMtd         PriceMethod1Code                  `xml:"urn:swift:xsd:sese.007.001.07 PricMtd,omitempty"`
	AcrdIntrstNAV   ActiveOrHistoricCurrencyAndAmount `xml:"urn:swift:xsd:sese.007.001.07 AcrdIntrstNAV,omitempty"`
	NbOfDaysAcrd    float64                           `xml:"urn:swift:xsd:sese.007.001.07 NbOfDaysAcrd,omitempty"`
	TaxblIncmPerShr ActiveCurrencyAnd13DecimalAmount  `xml:"urn:swift:xsd:sese.007.001.07 TaxblIncmPerShr,omitempty"`
}

type WaivingInstruction1Choice struct {
	Cd    WaivingInstruction1Code `xml:"urn:swift:xsd:sese.007.001.07 Cd"`
	Prtry GenericIdentification47 `xml:"urn:swift:xsd:sese.007.001.07 Prtry"`
}

// May be one of WICA, WIUN
type WaivingInstruction1Code string

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

type xsdGYearMonth time.Time

func (t *xsdGYearMonth) UnmarshalText(text []byte) error {
	return _unmarshalTime(text, (*time.Time)(t), "2006-01")
}
func (t xsdGYearMonth) MarshalText() ([]byte, error) {
	return _marshalTime((time.Time)(t), "2006-01")
}
func (t xsdGYearMonth) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if (time.Time)(t).IsZero() {
		return nil
	}
	m, err := t.MarshalText()
	if err != nil {
		return err
	}
	return e.EncodeElement(m, start)
}
func (t xsdGYearMonth) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if (time.Time)(t).IsZero() {
		return xml.Attr{}, nil
	}
	m, err := t.MarshalText()
	return xml.Attr{Name: name, Value: string(m)}, err
}
