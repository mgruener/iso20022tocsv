// Code generated by main. DO NOT EDIT.

package tsmt_009_001_05

import (
	"bytes"
	"encoding/xml"
	"time"
)

type AccountIdentification4Choice struct {
	IBAN IBAN2007Identifier            `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 IBAN"`
	Othr GenericAccountIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 Othr"`
}

type AccountSchemeName1Choice struct {
	Cd    ExternalAccountIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 Cd"`
	Prtry Max35Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 Prtry"`
}

type ActiveCurrencyAndAmount struct {
	Value float64            `xml:",chardata"`
	Ccy   ActiveCurrencyCode `xml:"Ccy,attr"`
}

// Must match the pattern [A-Z]{3,3}
type ActiveCurrencyCode string

// Must match the pattern [A-Z]{3,3}
type ActiveOrHistoricCurrencyCode string

type Adjustment7 struct {
	Tp        AdjustmentType1Choice     `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 Tp"`
	AmtOrPctg AmountOrPercentage2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 AmtOrPctg"`
	Drctn     AdjustmentDirection1Code  `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 Drctn"`
}

// May be one of ADDD, SUBS
type AdjustmentDirection1Code string

type AdjustmentType1Choice struct {
	Tp             AdjustmentType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 Tp"`
	OthrAdjstmntTp Max35Text           `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 OthrAdjstmntTp"`
}

// May be one of REBA, DISC, CREN, SURC
type AdjustmentType2Code string

type AirportDescription1 struct {
	Twn      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 Twn"`
	AirprtNm Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 AirprtNm,omitempty"`
}

type AirportName1Choice struct {
	AirprtCd       Max6Text            `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 AirprtCd"`
	OthrAirprtDesc AirportDescription1 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 OthrAirprtDesc"`
}

type AmountOrPercentage2Choice struct {
	Amt  ActiveCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 Amt"`
	Pctg float64                 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 Pctg"`
}

// May be one of BUYE, SELL, BUBA, SEBA
type AssuredType1Code string

type BICIdentification1 struct {
	BIC BICIdentifier `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 BIC"`
}

// Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type BICIdentifier string

type BPOApplicableRules1Choice struct {
	URBPOVrsn        float64   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 URBPOVrsn"`
	OthrRulesAndVrsn Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 OthrRulesAndVrsn"`
}

// May be one of BUYB, OBLB, RECB, SELB
type BankRole1Code string

type Baseline5 struct {
	SubmitrBaselnId      DocumentIdentification1  `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 SubmitrBaselnId"`
	SvcCd                TradeFinanceService2Code `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 SvcCd"`
	PurchsOrdrRef        DocumentIdentification7  `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 PurchsOrdrRef"`
	Buyr                 PartyIdentification26    `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 Buyr"`
	Sellr                PartyIdentification26    `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 Sellr"`
	BuyrBk               BICIdentification1       `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 BuyrBk"`
	SellrBk              BICIdentification1       `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 SellrBk"`
	BuyrSdSubmitgBk      []BICIdentification1     `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 BuyrSdSubmitgBk,omitempty"`
	SellrSdSubmitgBk     []BICIdentification1     `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 SellrSdSubmitgBk,omitempty"`
	BllTo                PartyIdentification26    `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 BllTo,omitempty"`
	ShipTo               PartyIdentification26    `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 ShipTo,omitempty"`
	Consgn               PartyIdentification26    `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 Consgn,omitempty"`
	Goods                LineItem13               `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 Goods"`
	PmtTerms             []PaymentTerms5          `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 PmtTerms"`
	SttlmTerms           SettlementTerms3         `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 SttlmTerms,omitempty"`
	PmtOblgtn            []PaymentObligation2     `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 PmtOblgtn,omitempty"`
	LatstMtchDt          ISODate                  `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 LatstMtchDt,omitempty"`
	ComrclDataSetReqrd   RequiredSubmission2      `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 ComrclDataSetReqrd"`
	TrnsprtDataSetReqrd  RequiredSubmission2      `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 TrnsprtDataSetReqrd,omitempty"`
	InsrncDataSetReqrd   RequiredSubmission3      `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 InsrncDataSetReqrd,omitempty"`
	CertDataSetReqrd     []RequiredSubmission4    `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 CertDataSetReqrd,omitempty"`
	OthrCertDataSetReqrd []RequiredSubmission6    `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 OthrCertDataSetReqrd,omitempty"`
	InttToPayXpctd       bool                     `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 InttToPayXpctd"`
}

type BaselineAmendmentRequestV05 struct {
	ReqId           MessageIdentification1          `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 ReqId"`
	TxId            SimpleIdentificationInformation `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 TxId"`
	SubmitrTxRef    SimpleIdentificationInformation `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 SubmitrTxRef,omitempty"`
	Baseln          Baseline5                       `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 Baseln"`
	BuyrCtctPrsn    []ContactIdentification1        `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 BuyrCtctPrsn,omitempty"`
	SellrCtctPrsn   []ContactIdentification1        `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 SellrCtctPrsn,omitempty"`
	BuyrBkCtctPrsn  []ContactIdentification1        `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 BuyrBkCtctPrsn,omitempty"`
	SellrBkCtctPrsn []ContactIdentification1        `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 SellrBkCtctPrsn,omitempty"`
	OthrBkCtctPrsn  []ContactIdentification3        `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 OthrBkCtctPrsn,omitempty"`
}

type CashAccount24 struct {
	Id  AccountIdentification4Choice `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 Id"`
	Tp  CashAccountType2Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 Tp,omitempty"`
	Ccy ActiveOrHistoricCurrencyCode `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 Ccy,omitempty"`
	Nm  Max70Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 Nm,omitempty"`
}

type CashAccountType2Choice struct {
	Cd    ExternalCashAccountType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 Cd"`
	Prtry Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 Prtry"`
}

type Charge24 struct {
	Tp    FreightCharges1Code `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 Tp"`
	Chrgs []ChargesDetails3   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 Chrgs,omitempty"`
}

// May be one of SIGN, STDE, STOR, PACK, PICK, DNGR, SECU, INSU, COLF, CHOR, CHDE, AIRF, TRPT
type ChargeType8Code string

type Charges5 struct {
	ChrgsPyer BankRole1Code     `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 ChrgsPyer"`
	ChrgsPyee BankRole1Code     `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 ChrgsPyee"`
	Amt       CurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 Amt,omitempty"`
	Pctg      float64           `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 Pctg,omitempty"`
	Tp        Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 Tp,omitempty"`
}

type ChargesDetails3 struct {
	Tp        ChargesType1Choice        `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 Tp"`
	AmtOrPctg AmountOrPercentage2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 AmtOrPctg"`
}

type ChargesType1Choice struct {
	Tp          ChargeType8Code `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 Tp"`
	OthrChrgsTp Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 OthrChrgsTp"`
}

type ContactIdentification1 struct {
	Nm       Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 Nm"`
	NmPrfx   NamePrefix1Code `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 NmPrfx,omitempty"`
	GvnNm    Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 GvnNm,omitempty"`
	Role     Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 Role,omitempty"`
	PhneNb   PhoneNumber     `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 PhneNb,omitempty"`
	FaxNb    PhoneNumber     `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 FaxNb,omitempty"`
	EmailAdr Max256Text      `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 EmailAdr,omitempty"`
}

type ContactIdentification3 struct {
	BIC      BICIdentifier   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 BIC"`
	Nm       Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 Nm"`
	NmPrfx   NamePrefix1Code `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 NmPrfx,omitempty"`
	GvnNm    Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 GvnNm,omitempty"`
	Role     Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 Role,omitempty"`
	PhneNb   PhoneNumber     `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 PhneNb,omitempty"`
	FaxNb    PhoneNumber     `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 FaxNb,omitempty"`
	EmailAdr Max256Text      `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 EmailAdr,omitempty"`
}

// Must match the pattern [A-Z]{2,2}
type CountryCode string

type CountrySubdivision1Choice struct {
	Cd    Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 Cd"`
	Prtry GenericIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 Prtry"`
}

type CurrencyAndAmount struct {
	Value float64      `xml:",chardata"`
	Ccy   CurrencyCode `xml:"Ccy,attr"`
}

// Must match the pattern [A-Z]{3,3}
type CurrencyCode string

type Document struct {
	BaselnAmdmntReq BaselineAmendmentRequestV05 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 BaselnAmdmntReq"`
}

type DocumentIdentification1 struct {
	Id      Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 Id"`
	Vrsn    float64            `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 Vrsn"`
	Submitr BICIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 Submitr"`
}

type DocumentIdentification7 struct {
	Id       Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 Id"`
	DtOfIsse ISODate   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 DtOfIsse"`
}

// Must match the pattern [a-zA-Z0-9]{4}
type Exact4AlphaNumericText string

// Must be at least 1 items long
type ExternalAccountIdentification1Code string

// Must be at least 1 items long
type ExternalCashAccountType1Code string

// Must be at least 1 items long
type ExternalIncoterms1Code string

type FinancialInstitutionIdentification4Choice struct {
	BIC      BICIdentifier   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 BIC"`
	NmAndAdr NameAndAddress6 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 NmAndAdr"`
}

// May be one of CLCT, PRPD
type FreightCharges1Code string

type GenericAccountIdentification1 struct {
	Id      Max34Text                `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 Id"`
	SchmeNm AccountSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 SchmeNm,omitempty"`
	Issr    Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 Issr,omitempty"`
}

type GenericIdentification1 struct {
	Id      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 Id"`
	SchmeNm Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 SchmeNm,omitempty"`
	Issr    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 Issr,omitempty"`
}

type GenericIdentification13 struct {
	Id      Max4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 Id"`
	SchmeNm Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 SchmeNm,omitempty"`
	Issr    Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 Issr"`
}

type GenericIdentification4 struct {
	Id   Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 Id"`
	IdTp Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 IdTp"`
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

type Incoterms4 struct {
	IncotrmsCd Incoterms4Choice `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 IncotrmsCd"`
	Lctn       Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 Lctn,omitempty"`
}

type Incoterms4Choice struct {
	Cd    ExternalIncoterms1Code  `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 Cd"`
	Prtry GenericIdentification13 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 Prtry"`
}

// May be one of ICCA, ICCB, ICCC, ICAI, IWCC, ISCC, IREC, ICLC, ISMC, CMCC, IRCE
type InsuranceClauses1Code string

type LineItem13 struct {
	GoodsAndOrSvcsDesc Max70Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 GoodsAndOrSvcsDesc,omitempty"`
	PrtlShipmnt        bool                      `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 PrtlShipmnt"`
	TrnsShipmnt        bool                      `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 TrnsShipmnt,omitempty"`
	ShipmntDtRg        ShipmentDateRange1        `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 ShipmntDtRg,omitempty"`
	LineItmDtls        []LineItemDetails13       `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 LineItmDtls"`
	LineItmsTtlAmt     CurrencyAndAmount         `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 LineItmsTtlAmt"`
	RtgSummry          TransportMeans5           `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 RtgSummry,omitempty"`
	Incotrms           Incoterms4                `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 Incotrms,omitempty"`
	Adjstmnt           []Adjustment7             `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 Adjstmnt,omitempty"`
	FrghtChrgs         Charge24                  `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 FrghtChrgs,omitempty"`
	Tax                []Tax23                   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 Tax,omitempty"`
	TtlNetAmt          CurrencyAndAmount         `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 TtlNetAmt"`
	BuyrDfndInf        []UserDefinedInformation1 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 BuyrDfndInf,omitempty"`
	SellrDfndInf       []UserDefinedInformation1 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 SellrDfndInf,omitempty"`
}

type LineItemDetails13 struct {
	LineItmId    Max70Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 LineItmId"`
	Qty          Quantity9                       `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 Qty"`
	QtyTlrnce    PercentageTolerance1            `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 QtyTlrnce,omitempty"`
	UnitPric     UnitPrice18                     `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 UnitPric,omitempty"`
	PricTlrnce   PercentageTolerance1            `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 PricTlrnce,omitempty"`
	PdctNm       Max70Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 PdctNm,omitempty"`
	PdctIdr      []ProductIdentifier2Choice      `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 PdctIdr,omitempty"`
	PdctChrtcs   []ProductCharacteristics1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 PdctChrtcs,omitempty"`
	PdctCtgy     []ProductCategory1Choice        `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 PdctCtgy,omitempty"`
	PdctOrgn     []CountryCode                   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 PdctOrgn,omitempty"`
	ShipmntSchdl ShipmentSchedule2Choice         `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 ShipmntSchdl,omitempty"`
	RtgSummry    TransportMeans5                 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 RtgSummry,omitempty"`
	Adjstmnt     []Adjustment7                   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 Adjstmnt,omitempty"`
	FrghtChrgs   Charge24                        `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 FrghtChrgs,omitempty"`
	Tax          []Tax23                         `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 Tax,omitempty"`
	TtlAmt       CurrencyAndAmount               `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 TtlAmt"`
	Incotrms     Incoterms4                      `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 Incotrms,omitempty"`
}

type Location2 struct {
	Ctry        CountryCode               `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 Ctry,omitempty"`
	CtrySubDvsn CountrySubdivision1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 CtrySubDvsn,omitempty"`
	Txt         Max35Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 Txt,omitempty"`
}

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
type Max35Text string

// Must match the pattern [a-zA-Z0-9]{1,4}
type Max4AlphaNumericText string

// Must be at least 1 items long
type Max6Text string

// Must be at least 1 items long
type Max70Text string

type MessageIdentification1 struct {
	Id      Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 Id"`
	CreDtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 CreDtTm"`
}

type MultimodalTransport3 struct {
	TakngInChrg  Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 TakngInChrg"`
	PlcOfFnlDstn Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 PlcOfFnlDstn"`
}

type NameAndAddress6 struct {
	Nm  Max70Text      `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 Nm"`
	Adr PostalAddress2 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 Adr"`
}

// May be one of DOCT, MIST, MISS, MADM
type NamePrefix1Code string

type PartyIdentification26 struct {
	Nm      Max70Text              `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 Nm"`
	PrtryId GenericIdentification4 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 PrtryId,omitempty"`
	PstlAdr PostalAddress5         `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 PstlAdr"`
}

type PartyIdentification27 struct {
	Nm      Max70Text              `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 Nm"`
	PrtryId GenericIdentification4 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 PrtryId,omitempty"`
	Ctry    CountryCode            `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 Ctry"`
}

type PaymentCodeOrOther1Choice struct {
	PmtCd        PaymentPeriod3 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 PmtCd"`
	PmtDueDt     ISODate        `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 PmtDueDt"`
	OthrPmtTerms Max140Text     `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 OthrPmtTerms"`
}

type PaymentCodeOrOther2Choice struct {
	PmtCd        PaymentPeriod4 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 PmtCd"`
	PmtDueDt     ISODate        `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 PmtDueDt"`
	OthrPmtTerms Max140Text     `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 OthrPmtTerms"`
}

type PaymentObligation2 struct {
	OblgrBk       BICIdentification1        `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 OblgrBk"`
	RcptBk        BICIdentification1        `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 RcptBk"`
	PmtOblgtnAmt  AmountOrPercentage2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 PmtOblgtnAmt"`
	Chrgs         []Charges5                `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 Chrgs,omitempty"`
	XpryDt        ISODate                   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 XpryDt"`
	AplblRules    BPOApplicableRules1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 AplblRules,omitempty"`
	AplblLaw      CountryCode               `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 AplblLaw,omitempty"`
	PlcOfJursdctn Location2                 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 PlcOfJursdctn,omitempty"`
	PmtTerms      []PaymentTerms4           `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 PmtTerms,omitempty"`
	SttlmTerms    SettlementTerms3          `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 SttlmTerms,omitempty"`
}

type PaymentPeriod3 struct {
	Cd       PaymentTime3Code `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 Cd"`
	NbOfDays float64          `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 NbOfDays,omitempty"`
}

type PaymentPeriod4 struct {
	Cd       PaymentTime4Code `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 Cd"`
	NbOfDays float64          `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 NbOfDays,omitempty"`
}

type PaymentTerms4 struct {
	PmtTerms  PaymentCodeOrOther1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 PmtTerms"`
	AmtOrPctg AmountOrPercentage2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 AmtOrPctg"`
}

type PaymentTerms5 struct {
	PmtTerms  PaymentCodeOrOther2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 PmtTerms"`
	AmtOrPctg AmountOrPercentage2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 AmtOrPctg"`
}

// May be one of EMTD, EMTR, EPBE, EPRD, PRMD, PRMR, EPIN, EPAM, EPPO, EPRR, EPSD, CASH, IREC
type PaymentTime3Code string

// May be one of IREC, CASH, EPSD, EPRR, EPPO, EPIN, PRMR, PRMD, EPRD, EPBE, EMTR, EMTD
type PaymentTime4Code string

type PercentageTolerance1 struct {
	PlusPct float64 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 PlusPct"`
	MnsPct  float64 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 MnsPct"`
}

// Must match the pattern \+[0-9]{1,3}-[0-9()+\-]{1,30}
type PhoneNumber string

type PostalAddress2 struct {
	StrtNm      Max70Text   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 StrtNm,omitempty"`
	PstCdId     Max16Text   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 PstCdId"`
	TwnNm       Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 TwnNm"`
	CtrySubDvsn Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 CtrySubDvsn,omitempty"`
	Ctry        CountryCode `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 Ctry"`
}

type PostalAddress5 struct {
	StrtNm      Max70Text   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 StrtNm,omitempty"`
	PstCdId     Max16Text   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 PstCdId,omitempty"`
	TwnNm       Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 TwnNm,omitempty"`
	CtrySubDvsn Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 CtrySubDvsn,omitempty"`
	Ctry        CountryCode `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 Ctry"`
}

type ProductCategory1 struct {
	Tp   ProductCategory1Code `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 Tp"`
	Ctgy Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 Ctgy"`
}

type ProductCategory1Choice struct {
	StrdPdctCtgy ProductCategory1       `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 StrdPdctCtgy"`
	OthrPdctCtgy GenericIdentification4 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 OthrPdctCtgy"`
}

// May be one of HRTR, QOTA, PRGP, LOBU, GNDR
type ProductCategory1Code string

type ProductCharacteristics1 struct {
	Tp     ProductCharacteristics1Code `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 Tp"`
	Chrtcs Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 Chrtcs"`
}

type ProductCharacteristics1Choice struct {
	StrdPdctChrtcs ProductCharacteristics1 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 StrdPdctChrtcs"`
	OthrPdctChrtcs GenericIdentification4  `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 OthrPdctChrtcs"`
}

// May be one of BISP, CHNR, CLOR, EDSP, ENNR, OPTN, ORCR, PCTV, SISP, SIZE, SZRG, SPRM, STOR, VINR
type ProductCharacteristics1Code string

type ProductIdentifier2 struct {
	Tp  ProductIdentifier2Code `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 Tp"`
	Idr Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 Idr"`
}

type ProductIdentifier2Choice struct {
	StrdPdctIdr ProductIdentifier2     `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 StrdPdctIdr"`
	OthrPdctIdr GenericIdentification4 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 OthrPdctIdr"`
}

// May be one of BINR, COMD, EANC, HRTR, MANI, MODL, PART, QOTA, STYL, SUPI, UPCC
type ProductIdentifier2Code string

type Quantity9 struct {
	UnitOfMeasr UnitOfMeasure3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 UnitOfMeasr"`
	Val         float64              `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 Val"`
	Fctr        Max15NumericText     `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 Fctr,omitempty"`
}

type RequiredSubmission2 struct {
	Submitr []BICIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 Submitr"`
}

type RequiredSubmission3 struct {
	Submitr      []BICIdentification1    `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 Submitr"`
	MtchIssr     PartyIdentification27   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 MtchIssr,omitempty"`
	MtchIsseDt   bool                    `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 MtchIsseDt"`
	MtchTrnsprt  bool                    `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 MtchTrnsprt"`
	MtchAmt      bool                    `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 MtchAmt"`
	ClausesReqrd []InsuranceClauses1Code `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 ClausesReqrd,omitempty"`
	MtchAssrdPty AssuredType1Code        `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 MtchAssrdPty,omitempty"`
}

type RequiredSubmission4 struct {
	Submitr           []BICIdentification1      `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 Submitr"`
	CertTp            TradeCertificateType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 CertTp"`
	MtchIssr          PartyIdentification27     `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 MtchIssr,omitempty"`
	MtchIsseDt        bool                      `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 MtchIsseDt"`
	MtchInspctnDt     bool                      `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 MtchInspctnDt"`
	AuthrsdInspctrInd bool                      `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 AuthrsdInspctrInd"`
	MtchConsgn        bool                      `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 MtchConsgn"`
	MtchManfctr       PartyIdentification27     `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 MtchManfctr,omitempty"`
	LineItmId         []Max70Text               `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 LineItmId,omitempty"`
}

type RequiredSubmission6 struct {
	Submitr    []BICIdentification1   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 Submitr"`
	CertTp     Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 CertTp"`
	CertTpDesc Max140Text             `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 CertTpDesc"`
}

type SettlementTerms3 struct {
	CdtrAgt  FinancialInstitutionIdentification4Choice `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 CdtrAgt,omitempty"`
	CdtrAcct CashAccount24                             `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 CdtrAcct"`
}

type ShipmentDateRange1 struct {
	EarlstShipmntDt ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 EarlstShipmntDt,omitempty"`
	LatstShipmntDt  ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 LatstShipmntDt,omitempty"`
}

type ShipmentDateRange2 struct {
	SubQtyVal       float64 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 SubQtyVal"`
	EarlstShipmntDt ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 EarlstShipmntDt,omitempty"`
	LatstShipmntDt  ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 LatstShipmntDt,omitempty"`
}

type ShipmentSchedule2Choice struct {
	ShipmntDtRg     ShipmentDateRange1   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 ShipmntDtRg"`
	ShipmntSubSchdl []ShipmentDateRange2 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 ShipmntSubSchdl"`
}

type SimpleIdentificationInformation struct {
	Id Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 Id"`
}

type SingleTransport7 struct {
	TrnsprtByAir  []TransportByAir5  `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 TrnsprtByAir,omitempty"`
	TrnsprtBySea  []TransportBySea6  `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 TrnsprtBySea,omitempty"`
	TrnsprtByRoad []TransportByRoad5 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 TrnsprtByRoad,omitempty"`
	TrnsprtByRail []TransportByRail5 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 TrnsprtByRail,omitempty"`
}

type Tax23 struct {
	Tp        TaxType2Choice            `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 Tp"`
	AmtOrPctg AmountOrPercentage2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 AmtOrPctg"`
}

type TaxType2Choice struct {
	Tp        TaxType9Code `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 Tp"`
	OthrTaxTp Max35Text    `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 OthrTaxTp"`
}

// May be one of PROV, NATI, STAT, WITH, STAM, COAX, VATA, CUST
type TaxType9Code string

// May be one of ANLY, QUAL, QUAN, WEIG, ORIG, HEAL, PHYT
type TradeCertificateType1Code string

// May be one of LEV1, LEV2, LEV3
type TradeFinanceService2Code string

type TransportByAir5 struct {
	DprtureAirprt []AirportName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 DprtureAirprt,omitempty"`
	DstnAirprt    []AirportName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 DstnAirprt"`
	AirCrrierNm   Max70Text            `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 AirCrrierNm,omitempty"`
	AirCrrierCtry CountryCode          `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 AirCrrierCtry,omitempty"`
	CrrierAgtNm   Max70Text            `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 CrrierAgtNm,omitempty"`
	CrrierAgtCtry CountryCode          `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 CrrierAgtCtry,omitempty"`
}

type TransportByRail5 struct {
	PlcOfRct       []Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 PlcOfRct,omitempty"`
	PlcOfDlvry     []Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 PlcOfDlvry"`
	RailCrrierNm   Max70Text   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 RailCrrierNm,omitempty"`
	RailCrrierCtry CountryCode `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 RailCrrierCtry,omitempty"`
	CrrierAgtNm    Max70Text   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 CrrierAgtNm,omitempty"`
	CrrierAgtCtry  CountryCode `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 CrrierAgtCtry,omitempty"`
}

type TransportByRoad5 struct {
	PlcOfRct       []Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 PlcOfRct,omitempty"`
	PlcOfDlvry     []Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 PlcOfDlvry"`
	RoadCrrierNm   Max70Text   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 RoadCrrierNm,omitempty"`
	RoadCrrierCtry CountryCode `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 RoadCrrierCtry,omitempty"`
	CrrierAgtNm    Max70Text   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 CrrierAgtNm,omitempty"`
	CrrierAgtCtry  CountryCode `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 CrrierAgtCtry,omitempty"`
}

type TransportBySea6 struct {
	PortOfLoadng  []Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 PortOfLoadng,omitempty"`
	PortOfDschrge []Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 PortOfDschrge"`
	VsslNm        Max70Text   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 VsslNm,omitempty"`
	SeaCrrierNm   Max70Text   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 SeaCrrierNm,omitempty"`
	SeaCrrierCtry CountryCode `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 SeaCrrierCtry,omitempty"`
	CrrierAgtNm   Max70Text   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 CrrierAgtNm,omitempty"`
	CrrierAgtCtry CountryCode `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 CrrierAgtCtry,omitempty"`
}

type TransportMeans5 struct {
	IndvTrnsprt   SingleTransport7     `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 IndvTrnsprt"`
	MltmdlTrnsprt MultimodalTransport3 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 MltmdlTrnsprt,omitempty"`
}

type UnitOfMeasure3Choice struct {
	UnitOfMeasrCd   UnitOfMeasure4Code `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 UnitOfMeasrCd"`
	OthrUnitOfMeasr Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 OthrUnitOfMeasr"`
}

// May be one of KGM, EA, LTN, MTR, INH, LY, GLI, GRM, CMT, MTK, FOT, 1A, INK, FTK, MIK, ONZ, PTI, PT, QTI, QT, GLL, MMT, KTM, YDK, MMK, CMK, KMK, MMQ, CLT, LTR, LBR, STN, BLL, BX, BO, CT, CH, CR, INQ, MTQ, OZI, OZA, BG, BL, TNE
type UnitOfMeasure4Code string

type UnitPrice18 struct {
	UnitPric UnitOfMeasure3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 UnitPric"`
	Amt      CurrencyAndAmount    `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 Amt"`
	Fctr     Max15NumericText     `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 Fctr,omitempty"`
}

type UserDefinedInformation1 struct {
	Labl Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 Labl"`
	Inf  Max140Text `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 Inf"`
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
