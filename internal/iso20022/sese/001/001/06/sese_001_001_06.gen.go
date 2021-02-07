// Code generated by main. DO NOT EDIT.

package sese_001_001_06

import (
	"bytes"
	"encoding/xml"
	"time"
)

type Account14 struct {
	Id       AccountIdentification1     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 Id"`
	AcctSvcr PartyIdentification2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 AcctSvcr,omitempty"`
}

type AccountIdentification1 struct {
	Prtry SimpleIdentificationInformation `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 Prtry"`
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
	Ref     Max35Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 Ref"`
	RefIssr PartyIdentification1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 RefIssr,omitempty"`
	MsgNm   Max35Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 MsgNm,omitempty"`
}

// May be one of ADDR, PBOX, HOME, BIZZ, MLTO, DLVY
type AddressType2Code string

type AlternateSecurityIdentification1 struct {
	Id         Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 Id"`
	DmstIdSrc  CountryCode `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 DmstIdSrc"`
	PrtryIdSrc Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 PrtryIdSrc"`
}

// Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type AnyBICIdentifier string

type BICIdentification1 struct {
	BIC BICIdentifier `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 BIC"`
}

// Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type BICIdentifier string

// May be one of NCER, ELEC, PHYS
type BeneficiaryCertificationCompletion1Code string

// Must be at least 1 items long
type BloombergIdentifier string

// May be one of SLDP, SLRP, DLPR
type BusinessFlowType1Code string

type Charge27 struct {
	Tp       ChargeType4Choice                 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 Tp"`
	Amt      ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 Amt"`
	ChrgBsis ChargeBasisType1Choice            `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 ChrgBsis,omitempty"`
	ChrgBr   ChargeBearer1Code                 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 ChrgBr,omitempty"`
	RcptId   PartyIdentification2Choice        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 RcptId,omitempty"`
}

type ChargeBasisType1Choice struct {
	Cd    TaxationBasis2Code      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 Cd"`
	Prtry GenericIdentification47 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 Prtry"`
}

// May be one of OUR, BEN, SHA
type ChargeBearer1Code string

type ChargePaymentMethod1Choice struct {
	Cd    ChargePaymentMethod1Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 Cd"`
	Prtry GenericIdentification47  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 Prtry"`
}

// May be one of CASH, UNIT
type ChargePaymentMethod1Code string

// May be one of BEND, DISC, FEND, POST, REGF, SHIP, SPCN, TRAN
type ChargeType12Code string

type ChargeType4Choice struct {
	Cd    ChargeType12Code        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 Cd"`
	Prtry GenericIdentification47 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 Prtry"`
}

type Commission22 struct {
	Tp             CommissionType3Choice            `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 Tp"`
	Bsis           CommissionBasis1Choice           `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 Bsis,omitempty"`
	Amt            ActiveCurrencyAnd13DecimalAmount `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 Amt"`
	RcptId         PartyIdentification2Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 RcptId,omitempty"`
	ComrclAgrmtRef Max35Text                        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 ComrclAgrmtRef,omitempty"`
	WvgDtls        CommissionWaiver4                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 WvgDtls,omitempty"`
}

type CommissionBasis1Choice struct {
	Cd    TaxationBasis4Code      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 Cd"`
	Prtry GenericIdentification47 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 Prtry"`
}

type CommissionType3Choice struct {
	Cd    CommissionType7Code     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 Cd"`
	Prtry GenericIdentification47 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 Prtry"`
}

// May be one of FEND, BEND
type CommissionType7Code string

type CommissionWaiver4 struct {
	InstrBsis WaivingInstruction1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 InstrBsis"`
	WvdRate   float64                   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 WvdRate"`
}

// Must be at least 1 items long
type ConsolidatedTapeAssociationIdentifier string

type ContactIdentification2 struct {
	NmPrfx   NamePrefix1Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 NmPrfx,omitempty"`
	GvnNm    Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 GvnNm,omitempty"`
	Nm       Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 Nm"`
	PhneNb   PhoneNumber     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 PhneNb,omitempty"`
	MobNb    PhoneNumber     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 MobNb,omitempty"`
	FaxNb    PhoneNumber     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 FaxNb,omitempty"`
	EmailAdr Max256Text      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 EmailAdr,omitempty"`
}

type CopyInformation2 struct {
	CpyInd    bool               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 CpyInd"`
	OrgnlRcvr BICIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 OrgnlRcvr,omitempty"`
}

// Must match the pattern [A-Z]{2,2}
type CountryCode string

// Must match the pattern [A-Z]{3,3}
type CurrencyCode string

type DateAndDateTimeChoice struct {
	Dt   ISODate     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 Dt"`
	DtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 DtTm"`
}

type DateFormat1Choice struct {
	Dt   ISODate             `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 Dt"`
	Cd   SettlementDate1Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 Cd"`
	DtTm ISODateTime         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 DtTm"`
}

type DeliveryParameters4 struct {
	RegdAdrInd bool                   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 RegdAdrInd"`
	NmAndAdr   NameAndAddress4        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 NmAndAdr,omitempty"`
	CtctPrsn   ContactIdentification2 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 CtctPrsn,omitempty"`
}

// May be one of DIST, ACCU
type DistributionPolicy1Code string

type Document struct {
	TrfOutInstr TransferOutInstructionV06 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 TrfOutInstr"`
}

// Must be at least 1 items long
type EuroclearClearstreamIdentifier string

// Must match the pattern [a-zA-Z0-9]{4}
type Exact4AlphaNumericText string

type ExemptionReason1Choice struct {
	Cd    TaxExemptReason1Code    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 Cd"`
	Prtry GenericIdentification47 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 Prtry"`
}

// Must be at least 1 items long
type Extended350Code string

type Extension1 struct {
	PlcAndNm Max350Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 PlcAndNm"`
	Txt      Max350Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 Txt"`
}

type FinancialInstrument13 struct {
	Id          SecurityIdentification3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 Id"`
	Nm          Max350Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 Nm,omitempty"`
	SplmtryId   Max35Text                     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 SplmtryId,omitempty"`
	ClssTp      Max35Text                     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 ClssTp,omitempty"`
	SctiesForm  FormOfSecurity1Code           `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 SctiesForm,omitempty"`
	DstrbtnPlcy DistributionPolicy1Code       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 DstrbtnPlcy,omitempty"`
}

type FinancialInstrumentQuantity1 struct {
	Unit float64 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 Unit"`
}

type ForeignExchangeTerms7 struct {
	ToAmt    ActiveCurrencyAnd13DecimalAmount `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 ToAmt,omitempty"`
	FrAmt    ActiveCurrencyAndAmount          `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 FrAmt,omitempty"`
	UnitCcy  ActiveOrHistoricCurrencyCode     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 UnitCcy"`
	QtdCcy   ActiveOrHistoricCurrencyCode     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 QtdCcy"`
	XchgRate float64                          `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 XchgRate"`
	QtnDt    ISODateTime                      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 QtnDt,omitempty"`
	QtgInstn PartyIdentification2Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 QtgInstn,omitempty"`
}

// May be one of BEAR, REGD
type FormOfSecurity1Code string

type GenericIdentification1 struct {
	Id      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 Id"`
	SchmeNm Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 SchmeNm,omitempty"`
	Issr    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 Issr,omitempty"`
}

type GenericIdentification27 struct {
	Id      Max4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 Id"`
	SchmeNm Max4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 SchmeNm,omitempty"`
	Issr    Max4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 Issr"`
}

type GenericIdentification47 struct {
	Id      Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 Id"`
	Issr    Max4AlphaNumericText   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 Issr"`
	SchmeNm Max4AlphaNumericText   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 SchmeNm,omitempty"`
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

// May be one of CASH, DRIP
type IncomePreference1Code string

type Intermediary25 struct {
	Id   PartyIdentification2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 Id"`
	Acct Account14                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 Acct,omitempty"`
	Role Role4Choice                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 Role,omitempty"`
}

type Intermediary26 struct {
	Id       PartyIdentification2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 Id"`
	Acct     Account14                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 Acct,omitempty"`
	Role     Role4Choice                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 Role,omitempty"`
	CtctPrsn ContactIdentification2     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 CtctPrsn,omitempty"`
}

type InvestmentAccount40 struct {
	OwnrId               []PartyIdentification2Choice            `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 OwnrId,omitempty"`
	AcctId               AccountIdentification1                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 AcctId"`
	AcctNm               Max35Text                               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 AcctNm,omitempty"`
	AcctDsgnt            Max35Text                               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 AcctDsgnt,omitempty"`
	IntrmyInf            []Intermediary25                        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 IntrmyInf,omitempty"`
	SctiesForm           FormOfSecurity1Code                     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 SctiesForm,omitempty"`
	DmtrlsdInd           bool                                    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 DmtrlsdInd,omitempty"`
	IncmPref             IncomePreference1Code                   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 IncmPref,omitempty"`
	BnfcryCertfctnCmpltn BeneficiaryCertificationCompletion1Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 BnfcryCertfctnCmpltn,omitempty"`
	SfkpgPlc             PartyIdentification2Choice              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 SfkpgPlc,omitempty"`
	AcctSvcr             PartyIdentification2Choice              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 AcctSvcr,omitempty"`
	SubAcctDtls          SubAccount1                             `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 SubAcctDtls,omitempty"`
}

type InvestmentAccount41 struct {
	OwnrId               []PartyIdentification2Choice            `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 OwnrId,omitempty"`
	AcctId               AccountIdentification1                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 AcctId,omitempty"`
	AcctNm               Max35Text                               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 AcctNm,omitempty"`
	AcctDsgnt            Max35Text                               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 AcctDsgnt,omitempty"`
	IntrmyInf            []Intermediary26                        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 IntrmyInf,omitempty"`
	SctiesForm           FormOfSecurity1Code                     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 SctiesForm,omitempty"`
	DmtrlsdInd           bool                                    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 DmtrlsdInd,omitempty"`
	IncmPref             IncomePreference1Code                   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 IncmPref,omitempty"`
	BnfcryCertfctnCmpltn BeneficiaryCertificationCompletion1Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 BnfcryCertfctnCmpltn,omitempty"`
	AcctSvcr             PartyIdentification2Choice              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 AcctSvcr,omitempty"`
	SubAcctDtls          SubAccount1                             `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 SubAcctDtls,omitempty"`
}

// May be one of FMCO, REGI, TRAG, INTR, DIST, CONC, UCL1, UCL2, TRAN
type InvestmentFundRole2Code string

type LongPostalAddress1Choice struct {
	Ustrd Max140Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 Ustrd"`
	Strd  StructuredLongPostalAddress1 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 Strd"`
}

type MarketPracticeVersion1 struct {
	Nm Max35Text    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 Nm"`
	Dt ISOYearMonth `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 Dt,omitempty"`
	Nb Max35Text    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 Nb,omitempty"`
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
	Id      Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 Id"`
	CreDtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 CreDtTm"`
}

type NameAndAddress2 struct {
	Nm  Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 Nm"`
	Adr LongPostalAddress1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 Adr,omitempty"`
}

type NameAndAddress4 struct {
	Nm  Max350Text     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 Nm,omitempty"`
	Adr PostalAddress1 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 Adr"`
}

type NameAndAddress5 struct {
	Nm  Max350Text     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 Nm"`
	Adr PostalAddress1 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 Adr,omitempty"`
}

// May be one of DOCT, MIST, MISS, MADM
type NamePrefix1Code string

type PartyIdentification1Choice struct {
	BICOrBEI AnyBICIdentifier       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 BICOrBEI"`
	PrtryId  GenericIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 PrtryId"`
	NmAndAdr NameAndAddress2        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 NmAndAdr"`
}

type PartyIdentification21 struct {
	PtyId    PartyIdentification2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 PtyId"`
	PrcgRef  Max35Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 PrcgRef,omitempty"`
	PrcgDt   DateAndDateTimeChoice      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 PrcgDt,omitempty"`
	CtctPrsn ContactIdentification2     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 CtctPrsn,omitempty"`
}

type PartyIdentification2Choice struct {
	BICOrBEI AnyBICIdentifier       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 BICOrBEI"`
	PrtryId  GenericIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 PrtryId"`
	NmAndAdr NameAndAddress5        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 NmAndAdr"`
}

type PartyIdentificationAndAccount4 struct {
	PtyId       PartyIdentification2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 PtyId"`
	AcctId      AccountIdentification1     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 AcctId,omitempty"`
	PrcgRef     Max35Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 PrcgRef,omitempty"`
	PrcgDt      DateAndDateTimeChoice      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 PrcgDt,omitempty"`
	SubAcctDtls SubAccount1                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 SubAcctDtls,omitempty"`
	CtctPrsn    ContactIdentification2     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 CtctPrsn,omitempty"`
}

type PartyIdentificationAndAccount5 struct {
	PtyId   PartyIdentification2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 PtyId"`
	AcctId  AccountIdentification1     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 AcctId,omitempty"`
	PrcgRef Max35Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 PrcgRef,omitempty"`
	PrcgDt  DateAndDateTimeChoice      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 PrcgDt,omitempty"`
}

type PartyIdentificationAndAccount93 struct {
	PtyId      PartyIdentification2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 PtyId,omitempty"`
	AcctId     AccountIdentification1     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 AcctId,omitempty"`
	PlcOfSttlm PartyIdentification2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 PlcOfSttlm"`
}

// Must match the pattern \+[0-9]{1,3}-[0-9()+\-]{1,30}
type PhoneNumber string

// May be one of DEMT, PHYS
type PhysicalTransferType1Code string

type PostalAddress1 struct {
	AdrTp       AddressType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 AdrTp,omitempty"`
	AdrLine     []Max70Text      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 AdrLine,omitempty"`
	StrtNm      Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 StrtNm,omitempty"`
	BldgNb      Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 BldgNb,omitempty"`
	PstCd       Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 PstCd,omitempty"`
	TwnNm       Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 TwnNm,omitempty"`
	CtrySubDvsn Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 CtrySubDvsn,omitempty"`
	Ctry        CountryCode      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 Ctry"`
}

// May be one of FORW, HIST
type PriceMethod1Code string

type PriceValue1 struct {
	Amt ActiveCurrencyAnd13DecimalAmount `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 Amt"`
}

type Quantity13Choice struct {
	TtlUnitsNb      FinancialInstrumentQuantity1 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 TtlUnitsNb"`
	PrtflTrfOutRate float64                      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 PrtflTrfOutRate"`
}

// Must be at least 1 items long
type RICIdentifier string

type ReceiveInformation15 struct {
	ReqdSttlmDt    ISODate                     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 ReqdSttlmDt,omitempty"`
	SttlmAmt       ActiveCurrencyAndAmount     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 SttlmAmt,omitempty"`
	StmpDty        StampDutyType2Code          `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 StmpDty,omitempty"`
	NetAmt         ActiveCurrencyAndAmount     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 NetAmt,omitempty"`
	SttlmPtiesDtls ReceivingPartiesAndAccount9 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 SttlmPtiesDtls,omitempty"`
	ChrgDtls       []Charge27                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 ChrgDtls,omitempty"`
	ComssnDtls     []Commission22              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 ComssnDtls,omitempty"`
	TaxDtls        []Tax25                     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 TaxDtls,omitempty"`
	FXDtls         []ForeignExchangeTerms7     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 FXDtls,omitempty"`
	PhysTrf        PhysicalTransferType1Code   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 PhysTrf,omitempty"`
	PhysTrfDtls    DeliveryParameters4         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 PhysTrfDtls,omitempty"`
	ClntRef        Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 ClntRef,omitempty"`
}

type ReceivingPartiesAndAccount9 struct {
	RcvrDtls        InvestmentAccount41            `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 RcvrDtls,omitempty"`
	RcvrsCtdnDtls   PartyIdentificationAndAccount5 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 RcvrsCtdnDtls,omitempty"`
	RcvrsIntrmyDtls PartyIdentificationAndAccount5 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 RcvrsIntrmyDtls,omitempty"`
	RcvgAgtDtls     PartyIdentificationAndAccount4 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 RcvgAgtDtls"`
	SctiesSttlmSys  Max35Text                      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 SctiesSttlmSys,omitempty"`
	PlcOfSttlmDtls  PartyIdentification21          `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 PlcOfSttlmDtls,omitempty"`
}

type Role4Choice struct {
	Cd    InvestmentFundRole2Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 Cd"`
	Prtry GenericIdentification47 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 Prtry"`
}

// May be one of RDUP, RDWN
type RoundingDirection2Code string

type SecurityIdentification3Choice struct {
	ISIN        ISINIdentifier                        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 ISIN"`
	SEDOL       string                                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 SEDOL"`
	CUSIP       string                                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 CUSIP"`
	RIC         RICIdentifier                         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 RIC"`
	TckrSymb    TickerIdentifier                      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 TckrSymb"`
	Blmbrg      BloombergIdentifier                   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 Blmbrg"`
	CTA         ConsolidatedTapeAssociationIdentifier `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 CTA"`
	QUICK       string                                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 QUICK"`
	Wrtppr      string                                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 Wrtppr"`
	Dtch        string                                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 Dtch"`
	Vlrn        string                                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 Vlrn"`
	SCVM        string                                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 SCVM"`
	Belgn       string                                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 Belgn"`
	Cmon        EuroclearClearstreamIdentifier        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 Cmon"`
	OthrPrtryId AlternateSecurityIdentification1      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 OthrPrtryId"`
}

// May be one of ASAP, ENDC, WHIF
type SettlementDate1Code string

type SimpleIdentificationInformation struct {
	Id Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 Id"`
}

// May be one of ASTD, SDRN
type StampDutyType2Code string

type StructuredLongPostalAddress1 struct {
	BldgNm     Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 BldgNm,omitempty"`
	StrtNm     Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 StrtNm,omitempty"`
	StrtBldgId Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 StrtBldgId,omitempty"`
	Flr        Max16Text   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 Flr,omitempty"`
	TwnNm      Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 TwnNm"`
	DstrctNm   Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 DstrctNm,omitempty"`
	RgnId      Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 RgnId,omitempty"`
	Stat       Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 Stat,omitempty"`
	CtyId      Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 CtyId,omitempty"`
	Ctry       CountryCode `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 Ctry"`
	PstCdId    Max16Text   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 PstCdId"`
	POB        Max16Text   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 POB,omitempty"`
}

type SubAccount1 struct {
	Id    AccountIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 Id"`
	Nm    Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 Nm,omitempty"`
	Chrtc Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 Chrtc,omitempty"`
}

type Tax25 struct {
	Tp           TaxType1Choice                             `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 Tp"`
	Amt          ActiveOrHistoricCurrencyAnd13DecimalAmount `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 Amt"`
	Bsis         TaxBasis1Choice                            `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 Bsis,omitempty"`
	RcptId       PartyIdentification2Choice                 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 RcptId,omitempty"`
	XmptnInd     bool                                       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 XmptnInd"`
	XmptnRsn     ExemptionReason1Choice                     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 XmptnRsn,omitempty"`
	TaxClctnDtls TaxCalculationInformation8                 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 TaxClctnDtls,omitempty"`
}

type TaxBasis1Choice struct {
	Cd    TaxationBasis2Code      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 Cd"`
	Prtry GenericIdentification47 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 Prtry"`
}

type TaxCalculationInformation8 struct {
	Bsis     TaxBasis1Choice                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 Bsis,omitempty"`
	TaxblAmt ActiveCurrencyAnd13DecimalAmount `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 TaxblAmt"`
}

// May be one of NONE, MASA, MISA, SISA, IISA, CUYP, PRYP, ASTR, EMPY, EMCY, EPRY, ECYE, NFPI, NFQP, DECP, IRAC, IRAR, KEOG, PFSP, 401K, SIRA, 403B, 457X, RIRA, RIAN, RCRF, RCIP, EIFP, EIOP
type TaxExemptReason1Code string

// May be one of COAX, CTAX, EUTR, LEVY, LOCL, NATI, PROV, STAM, STAT, STEX, TRAN, TRAX, VATA, WITH, NKAP, KAPA
type TaxType16Code string

type TaxType1Choice struct {
	Cd    TaxType16Code           `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 Cd"`
	Prtry GenericIdentification47 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 Prtry"`
}

// May be one of FLAT, PERU
type TaxationBasis2Code string

// May be one of FLAT, PERU, GRAM, NEAM
type TaxationBasis4Code string

// Must be at least 1 items long
type TickerIdentifier string

type Transfer27 struct {
	TrfRef         Max35Text                                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 TrfRef"`
	ClntRef        Max35Text                                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 ClntRef,omitempty"`
	CtrPtyRef      AdditionalReference2                       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 CtrPtyRef,omitempty"`
	BizFlowTp      BusinessFlowType1Code                      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 BizFlowTp,omitempty"`
	ReqdSttlmDt    ISODate                                    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 ReqdSttlmDt,omitempty"`
	TrfOrdrDtForm  ISODate                                    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 TrfOrdrDtForm,omitempty"`
	TrfRsn         TransferReason1                            `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 TrfRsn,omitempty"`
	HldgsPlanTp    []HoldingsPlanType1Code                    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 HldgsPlanTp,omitempty"`
	FinInstrmDtls  FinancialInstrument13                      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 FinInstrmDtls"`
	Qty            Quantity13Choice                           `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 Qty"`
	UnitsDtls      []Unit3                                    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 UnitsDtls,omitempty"`
	Rndg           RoundingDirection2Code                     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 Rndg,omitempty"`
	AvrgPric       ActiveOrHistoricCurrencyAnd13DecimalAmount `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 AvrgPric,omitempty"`
	TrfCcy         CurrencyCode                               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 TrfCcy,omitempty"`
	OwnAcctTrfInd  bool                                       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 OwnAcctTrfInd,omitempty"`
	NonStdSttlmInf Max350Text                                 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 NonStdSttlmInf,omitempty"`
	RcvgAgtDtls    PartyIdentificationAndAccount93            `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 RcvgAgtDtls,omitempty"`
	DlvrgAgtDtls   PartyIdentificationAndAccount93            `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 DlvrgAgtDtls,omitempty"`
	TrfExpnssPmtTp ChargePaymentMethod1Choice                 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 TrfExpnssPmtTp,omitempty"`
}

type TransferOutInstructionV06 struct {
	MsgId        MessageIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 MsgId"`
	PoolRef      AdditionalReference2   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 PoolRef,omitempty"`
	PrvsRef      AdditionalReference2   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 PrvsRef,omitempty"`
	RltdRef      AdditionalReference2   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 RltdRef,omitempty"`
	MstrRef      Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 MstrRef,omitempty"`
	ReqdTrfDt    DateFormat1Choice      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 ReqdTrfDt,omitempty"`
	TrfDtls      []Transfer27           `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 TrfDtls"`
	AcctDtls     InvestmentAccount40    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 AcctDtls"`
	SttlmDtls    ReceiveInformation15   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 SttlmDtls,omitempty"`
	MktPrctcVrsn MarketPracticeVersion1 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 MktPrctcVrsn,omitempty"`
	CpyDtls      CopyInformation2       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 CpyDtls,omitempty"`
	Xtnsn        []Extension1           `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 Xtnsn,omitempty"`
}

type TransferReason1 struct {
	Cd    TransferReason1Code     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 Cd"`
	Prtry GenericIdentification27 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 Prtry"`
}

// May be one of TRAU, TRAC, TRAT, TRAO, TRAI, TRAG, TPLD, TTDT, TRPE, TRAF, TRAN
type TransferReason1Code string

// May be one of BIDE, OFFR, NAVL, CREA, CANC, INTE, SWNG, MIDD, RINV, SWIC
type TypeOfPrice12Code string

// May be one of GRP1, GRP2
type UKTaxGroupUnitCode string

type Unit3 struct {
	UnitsNb      FinancialInstrumentQuantity1 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 UnitsNb"`
	AcqstnDt     ISODate                      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 AcqstnDt,omitempty"`
	CertNb       []Max35Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 CertNb,omitempty"`
	Grp1Or2Units UKTaxGroupUnitCode           `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 Grp1Or2Units,omitempty"`
	PricDtls     UnitPrice12                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 PricDtls,omitempty"`
}

type UnitPrice12 struct {
	Tp              TypeOfPrice12Code                 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 Tp"`
	XtndedTp        Extended350Code                   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 XtndedTp"`
	Val             PriceValue1                       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 Val"`
	PricMtd         PriceMethod1Code                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 PricMtd,omitempty"`
	AcrdIntrstNAV   ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 AcrdIntrstNAV,omitempty"`
	NbOfDaysAcrd    float64                           `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 NbOfDaysAcrd,omitempty"`
	TaxblIncmPerShr ActiveCurrencyAnd13DecimalAmount  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 TaxblIncmPerShr,omitempty"`
}

type WaivingInstruction1Choice struct {
	Cd    WaivingInstruction1Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 Cd"`
	Prtry GenericIdentification47 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.06 Prtry"`
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
