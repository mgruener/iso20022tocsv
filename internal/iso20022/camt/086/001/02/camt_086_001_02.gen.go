// Code generated by main. DO NOT EDIT.

package camt_086_001_02

import (
	"bytes"
	"encoding/xml"
	"time"
)

type AccountIdentification4Choice struct {
	IBAN IBAN2007Identifier            `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 IBAN"`
	Othr GenericAccountIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 Othr"`
}

// May be one of INTM, SMRY
type AccountLevel1Code string

// May be one of INTM, SMRY, DETL
type AccountLevel2Code string

type AccountSchemeName1Choice struct {
	Cd    ExternalAccountIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 Cd"`
	Prtry Max35Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 Prtry"`
}

type AccountTax1 struct {
	ClctnMtd   BillingTaxCalculationMethod1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 ClctnMtd"`
	Rgn        Max40Text                        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 Rgn,omitempty"`
	NonResCtry ResidenceLocation1Choice         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 NonResCtry,omitempty"`
}

type ActiveOrHistoricCurrencyAndAmount struct {
	Value float64                      `xml:",chardata"`
	Ccy   ActiveOrHistoricCurrencyCode `xml:"Ccy,attr"`
}

// Must match the pattern [A-Z]{3,3}
type ActiveOrHistoricCurrencyCode string

// May be one of ADDR, PBOX, HOME, BIZZ, MLTO, DLVY
type AddressType2Code string

type AmountAndDirection34 struct {
	Amt ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 Amt"`
	Sgn bool                              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 Sgn"`
}

// Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type AnyBICIdentifier string

// Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type BICFIIdentifier string

type BalanceAdjustment1 struct {
	Tp                BalanceAdjustmentType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 Tp"`
	Desc              Max105Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 Desc"`
	BalAmt            AmountAndDirection34       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 BalAmt"`
	AvrgAmt           AmountAndDirection34       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 AvrgAmt,omitempty"`
	ErrDt             ISODate                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 ErrDt,omitempty"`
	PstngDt           ISODate                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 PstngDt"`
	Days              float64                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 Days,omitempty"`
	EarngsAdjstmntAmt AmountAndDirection34       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 EarngsAdjstmntAmt,omitempty"`
}

// May be one of LDGR, FLOT, CLLD
type BalanceAdjustmentType1Code string

type BankServicesBillingStatementV02 struct {
	RptHdr      ReportHeader3     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 RptHdr"`
	BllgStmtGrp []StatementGroup2 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 BllgStmtGrp"`
}

type BankTransactionCodeStructure4 struct {
	Domn  BankTransactionCodeStructure5            `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 Domn,omitempty"`
	Prtry ProprietaryBankTransactionCodeStructure1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 Prtry,omitempty"`
}

type BankTransactionCodeStructure5 struct {
	Cd   ExternalBankTransactionDomain1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 Cd"`
	Fmly BankTransactionCodeStructure6      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 Fmly"`
}

type BankTransactionCodeStructure6 struct {
	Cd        ExternalBankTransactionFamily1Code    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 Cd"`
	SubFmlyCd ExternalBankTransactionSubFamily1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 SubFmlyCd"`
}

type BillingBalance1 struct {
	Tp    BillingBalanceType1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 Tp"`
	Val   AmountAndDirection34      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 Val"`
	CcyTp BillingCurrencyType1Code  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 CcyTp,omitempty"`
}

type BillingBalanceType1Choice struct {
	Cd    ExternalBillingBalanceType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 Cd"`
	Prtry Max35Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 Prtry"`
}

// May be one of UPRC, STAM, BCHG, DPRC, FCHG, LPRC, MCHG, MXRD, TIR1, TIR2, TIR3, TIR4, TIR5, TIR6, TIR7, TIR8, TIR9, TPRC, ZPRC, BBSE
type BillingChargeMethod1Code string

type BillingCompensation1 struct {
	Tp    BillingCompensationType1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 Tp"`
	Val   AmountAndDirection34           `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 Val"`
	CcyTp BillingCurrencyType2Code       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 CcyTp,omitempty"`
}

type BillingCompensationType1Choice struct {
	Cd    ExternalBillingCompensationType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 Cd"`
	Prtry Max35Text                            `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 Prtry"`
}

// May be one of ACCT, STLM, PRCG
type BillingCurrencyType1Code string

// May be one of ACCT, STLM, PRCG, HOST
type BillingCurrencyType2Code string

type BillingMethod1 struct {
	SvcChrgHstAmt AmountAndDirection34   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 SvcChrgHstAmt"`
	SvcTax        BillingServicesAmount1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 SvcTax"`
	TtlChrg       BillingServicesAmount2 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 TtlChrg"`
	TaxId         []BillingServicesTax1  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 TaxId"`
}

type BillingMethod1Choice struct {
	MtdA BillingMethod1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 MtdA"`
	MtdB BillingMethod2 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 MtdB"`
	MtdD BillingMethod3 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 MtdD"`
}

type BillingMethod2 struct {
	SvcChrgHstAmt AmountAndDirection34   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 SvcChrgHstAmt"`
	SvcTax        BillingServicesAmount1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 SvcTax"`
	TaxId         []BillingServicesTax1  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 TaxId"`
}

type BillingMethod3 struct {
	SvcTaxPricAmt AmountAndDirection34  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 SvcTaxPricAmt"`
	TaxId         []BillingServicesTax2 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 TaxId"`
}

type BillingMethod4 struct {
	SvcDtl   []BillingServiceParameters2 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 SvcDtl"`
	TaxClctn TaxCalculation1             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 TaxClctn"`
}

type BillingPrice1 struct {
	Ccy      ActiveOrHistoricCurrencyCode `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 Ccy,omitempty"`
	UnitPric AmountAndDirection34         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 UnitPric,omitempty"`
	Mtd      BillingChargeMethod1Code     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 Mtd,omitempty"`
	Rule     Max20Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 Rule,omitempty"`
}

type BillingRate1 struct {
	Id        BillingRateIdentification1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 Id"`
	Val       float64                          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 Val"`
	DaysInPrd float64                          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 DaysInPrd,omitempty"`
	DaysInYr  float64                          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 DaysInYr,omitempty"`
}

type BillingRateIdentification1Choice struct {
	Cd    ExternalBillingRateIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 Cd"`
	Prtry Max35Text                              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 Prtry"`
}

type BillingService2 struct {
	SvcDtl            BillingServiceParameters3 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 SvcDtl"`
	Pric              BillingPrice1             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 Pric,omitempty"`
	PmtMtd            ServicePaymentMethod1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 PmtMtd"`
	OrgnlChrgPric     AmountAndDirection34      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 OrgnlChrgPric"`
	OrgnlChrgSttlmAmt AmountAndDirection34      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 OrgnlChrgSttlmAmt,omitempty"`
	BalReqrdAcctAmt   AmountAndDirection34      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 BalReqrdAcctAmt,omitempty"`
	TaxDsgnt          ServiceTaxDesignation1    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 TaxDsgnt"`
	TaxClctn          BillingMethod1Choice      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 TaxClctn,omitempty"`
}

type BillingServiceAdjustment1 struct {
	Tp           ServiceAdjustmentType1Code       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 Tp"`
	Desc         Max140Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 Desc"`
	Amt          AmountAndDirection34             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 Amt"`
	BalReqrdAmt  AmountAndDirection34             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 BalReqrdAmt,omitempty"`
	ErrDt        ISODate                          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 ErrDt,omitempty"`
	AdjstmntId   Max35Text                        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 AdjstmntId,omitempty"`
	SubSvc       BillingSubServiceIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 SubSvc,omitempty"`
	PricChng     AmountAndDirection34             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 PricChng,omitempty"`
	OrgnlPric    AmountAndDirection34             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 OrgnlPric,omitempty"`
	NewPric      AmountAndDirection34             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 NewPric,omitempty"`
	VolChng      float64                          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 VolChng,omitempty"`
	OrgnlVol     float64                          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 OrgnlVol,omitempty"`
	NewVol       float64                          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 NewVol,omitempty"`
	OrgnlChrgAmt AmountAndDirection34             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 OrgnlChrgAmt,omitempty"`
	NewChrgAmt   AmountAndDirection34             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 NewChrgAmt,omitempty"`
}

type BillingServiceCommonIdentification1 struct {
	Issr Max6Text `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 Issr"`
	Id   Max8Text `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 Id"`
}

type BillingServiceIdentification2 struct {
	Id     Max35Text                        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 Id"`
	SubSvc BillingSubServiceIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 SubSvc,omitempty"`
	Desc   Max70Text                        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 Desc"`
}

type BillingServiceIdentification3 struct {
	Id     Max35Text                           `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 Id"`
	SubSvc BillingSubServiceIdentification1    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 SubSvc,omitempty"`
	Desc   Max70Text                           `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 Desc"`
	CmonCd BillingServiceCommonIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 CmonCd,omitempty"`
	BkTxCd BankTransactionCodeStructure4       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 BkTxCd,omitempty"`
	SvcTp  Max12Text                           `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 SvcTp,omitempty"`
}

type BillingServiceParameters2 struct {
	BkSvc      BillingServiceIdentification2 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 BkSvc"`
	Vol        float64                       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 Vol,omitempty"`
	UnitPric   AmountAndDirection34          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 UnitPric,omitempty"`
	SvcChrgAmt AmountAndDirection34          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 SvcChrgAmt"`
}

type BillingServiceParameters3 struct {
	BkSvc BillingServiceIdentification3 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 BkSvc"`
	Vol   float64                       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 Vol,omitempty"`
}

type BillingServicesAmount1 struct {
	HstAmt   AmountAndDirection34 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 HstAmt"`
	PricgAmt AmountAndDirection34 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 PricgAmt,omitempty"`
}

type BillingServicesAmount2 struct {
	HstAmt   AmountAndDirection34 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 HstAmt"`
	SttlmAmt AmountAndDirection34 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 SttlmAmt,omitempty"`
	PricgAmt AmountAndDirection34 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 PricgAmt,omitempty"`
}

type BillingServicesAmount3 struct {
	SrcAmt AmountAndDirection34 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 SrcAmt"`
	HstAmt AmountAndDirection34 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 HstAmt"`
}

type BillingServicesTax1 struct {
	Nb       Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 Nb"`
	Desc     Max40Text            `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 Desc,omitempty"`
	Rate     float64              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 Rate"`
	HstAmt   AmountAndDirection34 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 HstAmt"`
	PricgAmt AmountAndDirection34 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 PricgAmt,omitempty"`
}

type BillingServicesTax2 struct {
	Nb       Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 Nb"`
	Desc     Max40Text            `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 Desc,omitempty"`
	Rate     float64              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 Rate"`
	PricgAmt AmountAndDirection34 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 PricgAmt"`
}

type BillingServicesTax3 struct {
	Nb        Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 Nb"`
	Desc      Max40Text            `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 Desc,omitempty"`
	Rate      float64              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 Rate"`
	TtlTaxAmt AmountAndDirection34 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 TtlTaxAmt"`
}

type BillingStatement2 struct {
	StmtId      Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 StmtId"`
	FrToDt      DatePeriod1                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 FrToDt"`
	CreDtTm     ISODateTime                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 CreDtTm"`
	Sts         BillingStatementStatus1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 Sts"`
	AcctChrtcs  CashAccountCharacteristics2 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 AcctChrtcs"`
	RateData    []BillingRate1              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 RateData,omitempty"`
	CcyXchg     []CurrencyExchange6         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 CcyXchg,omitempty"`
	Bal         []BillingBalance1           `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 Bal,omitempty"`
	Compstn     []BillingCompensation1      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 Compstn,omitempty"`
	Svc         []BillingService2           `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 Svc,omitempty"`
	TaxRgn      []BillingTaxRegion1         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 TaxRgn,omitempty"`
	BalAdjstmnt []BalanceAdjustment1        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 BalAdjstmnt,omitempty"`
	SvcAdjstmnt []BillingServiceAdjustment1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 SvcAdjstmnt,omitempty"`
}

// May be one of ORGN, RPLC, TEST
type BillingStatementStatus1Code string

type BillingSubServiceIdentification1 struct {
	Issr BillingSubServiceQualifier1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 Issr"`
	Id   Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 Id"`
}

type BillingSubServiceQualifier1Choice struct {
	Cd    BillingSubServiceQualifier1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 Cd"`
	Prtry Max35Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 Prtry"`
}

// May be one of LBOX, STOR, BILA, SEQN, MACT
type BillingSubServiceQualifier1Code string

// May be one of NTAX, MTDA, MTDB, MTDC, MTDD, UDFD
type BillingTaxCalculationMethod1Code string

type BillingTaxIdentification1 struct {
	VATRegnNb Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 VATRegnNb,omitempty"`
	TaxRegnNb Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 TaxRegnNb,omitempty"`
	TaxCtct   ContactDetails3 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 TaxCtct,omitempty"`
}

type BillingTaxRegion1 struct {
	RgnNb       Max40Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 RgnNb"`
	RgnNm       Max40Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 RgnNm"`
	CstmrTaxId  Max40Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 CstmrTaxId"`
	PtDt        ISODate                   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 PtDt,omitempty"`
	SndgFI      BillingTaxIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 SndgFI,omitempty"`
	InvcNb      Max40Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 InvcNb,omitempty"`
	MtdC        BillingMethod4            `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 MtdC,omitempty"`
	SttlmAmt    AmountAndDirection34      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 SttlmAmt"`
	TaxDueToRgn AmountAndDirection34      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 TaxDueToRgn"`
}

type BranchAndFinancialInstitutionIdentification5 struct {
	FinInstnId FinancialInstitutionIdentification8 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 FinInstnId"`
	BrnchId    BranchData2                         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 BrnchId,omitempty"`
}

type BranchData2 struct {
	Id      Max35Text      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 Id,omitempty"`
	Nm      Max140Text     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 Nm,omitempty"`
	PstlAdr PostalAddress6 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 PstlAdr,omitempty"`
}

type CashAccount24 struct {
	Id  AccountIdentification4Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 Id"`
	Tp  CashAccountType2Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 Tp,omitempty"`
	Ccy ActiveOrHistoricCurrencyCode `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 Ccy,omitempty"`
	Nm  Max70Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 Nm,omitempty"`
}

type CashAccountCharacteristics2 struct {
	AcctLvl      AccountLevel2Code                            `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 AcctLvl"`
	CshAcct      CashAccount24                                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 CshAcct"`
	AcctSvcr     BranchAndFinancialInstitutionIdentification5 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 AcctSvcr,omitempty"`
	PrntAcct     ParentCashAccount2                           `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 PrntAcct,omitempty"`
	CompstnMtd   CompensationMethod1Code                      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 CompstnMtd"`
	DbtAcct      AccountIdentification4Choice                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 DbtAcct,omitempty"`
	DelydDbtDt   ISODate                                      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 DelydDbtDt,omitempty"`
	SttlmAdvc    Max105Text                                   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 SttlmAdvc,omitempty"`
	AcctBalCcyCd ActiveOrHistoricCurrencyCode                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 AcctBalCcyCd"`
	SttlmCcyCd   ActiveOrHistoricCurrencyCode                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 SttlmCcyCd,omitempty"`
	HstCcyCd     ActiveOrHistoricCurrencyCode                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 HstCcyCd,omitempty"`
	Tax          AccountTax1                                  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 Tax,omitempty"`
	AcctSvcrCtct ContactDetails3                              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 AcctSvcrCtct"`
}

type CashAccountType2Choice struct {
	Cd    ExternalCashAccountType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 Cd"`
	Prtry Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 Prtry"`
}

type ClearingSystemIdentification2Choice struct {
	Cd    ExternalClearingSystemIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 Cd"`
	Prtry Max35Text                                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 Prtry"`
}

type ClearingSystemMemberIdentification2 struct {
	ClrSysId ClearingSystemIdentification2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 ClrSysId,omitempty"`
	MmbId    Max35Text                           `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 MmbId"`
}

// May be one of NOCP, DBTD, INVD, DDBT
type CompensationMethod1Code string

type ContactDetails3 struct {
	NmPrfx    NamePrefix1Code             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 NmPrfx,omitempty"`
	Nm        Max140Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 Nm,omitempty"`
	PhneNb    PhoneNumber                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 PhneNb,omitempty"`
	MobNb     PhoneNumber                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 MobNb,omitempty"`
	FaxNb     PhoneNumber                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 FaxNb,omitempty"`
	EmailAdr  Max2048Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 EmailAdr,omitempty"`
	Othr      []OtherContact1             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 Othr,omitempty"`
	PrefrdMtd PreferredContactMethod1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 PrefrdMtd,omitempty"`
}

// Must match the pattern [A-Z]{2,2}
type CountryCode string

type CurrencyExchange6 struct {
	SrcCcy   ActiveOrHistoricCurrencyCode `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 SrcCcy"`
	TrgtCcy  ActiveOrHistoricCurrencyCode `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 TrgtCcy"`
	XchgRate float64                      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 XchgRate"`
	Desc     Max40Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 Desc,omitempty"`
	UnitCcy  ActiveOrHistoricCurrencyCode `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 UnitCcy,omitempty"`
	Cmnts    Max70Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 Cmnts,omitempty"`
	QtnDt    ISODateTime                  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 QtnDt,omitempty"`
}

type DatePeriod1 struct {
	FrDt ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 FrDt,omitempty"`
	ToDt ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 ToDt"`
}

type Document struct {
	BkSvcsBllgStmt BankServicesBillingStatementV02 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 BkSvcsBllgStmt"`
}

// Must be at least 1 items long
type ExternalAccountIdentification1Code string

// Must be at least 1 items long
type ExternalBankTransactionDomain1Code string

// Must be at least 1 items long
type ExternalBankTransactionFamily1Code string

// Must be at least 1 items long
type ExternalBankTransactionSubFamily1Code string

// Must be at least 1 items long
type ExternalBillingBalanceType1Code string

// Must be at least 1 items long
type ExternalBillingCompensationType1Code string

// Must be at least 1 items long
type ExternalBillingRateIdentification1Code string

// Must be at least 1 items long
type ExternalCashAccountType1Code string

// Must be at least 1 items long
type ExternalClearingSystemIdentification1Code string

// Must be at least 1 items long
type ExternalFinancialInstitutionIdentification1Code string

// Must be at least 1 items long
type ExternalOrganisationIdentification1Code string

type FinancialIdentificationSchemeName1Choice struct {
	Cd    ExternalFinancialInstitutionIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 Cd"`
	Prtry Max35Text                                       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 Prtry"`
}

type FinancialInstitutionIdentification8 struct {
	BICFI       BICFIIdentifier                     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 BICFI,omitempty"`
	ClrSysMmbId ClearingSystemMemberIdentification2 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 ClrSysMmbId,omitempty"`
	Nm          Max140Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 Nm,omitempty"`
	PstlAdr     PostalAddress6                      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 PstlAdr,omitempty"`
	Othr        GenericFinancialIdentification1     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 Othr,omitempty"`
}

type FinancialInstitutionIdentification9 struct {
	BICFI       BICFIIdentifier                     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 BICFI,omitempty"`
	ClrSysMmbId ClearingSystemMemberIdentification2 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 ClrSysMmbId,omitempty"`
	Othr        GenericFinancialIdentification1     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 Othr,omitempty"`
}

type GenericAccountIdentification1 struct {
	Id      Max34Text                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 Id"`
	SchmeNm AccountSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 SchmeNm,omitempty"`
	Issr    Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 Issr,omitempty"`
}

type GenericFinancialIdentification1 struct {
	Id      Max35Text                                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 Id"`
	SchmeNm FinancialIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 SchmeNm,omitempty"`
	Issr    Max35Text                                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 Issr,omitempty"`
}

type GenericOrganisationIdentification1 struct {
	Id      Max35Text                                   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 Id"`
	SchmeNm OrganisationIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 SchmeNm,omitempty"`
	Issr    Max35Text                                   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 Issr,omitempty"`
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

// Must be at least 1 items long
type Max105Text string

// Must be at least 1 items long
type Max10Text string

// Must be at least 1 items long
type Max128Text string

// Must be at least 1 items long
type Max12Text string

// Must be at least 1 items long
type Max140Text string

// Must be at least 1 items long
type Max16Text string

// Must be at least 1 items long
type Max2048Text string

// Must be at least 1 items long
type Max20Text string

// Must be at least 1 items long
type Max34Text string

// Must be at least 1 items long
type Max35Text string

// Must be at least 1 items long
type Max40Text string

// Must be at least 1 items long
type Max4Text string

// Must match the pattern [0-9]{1,5}
type Max5NumericText string

// Must be at least 1 items long
type Max6Text string

// Must be at least 1 items long
type Max70Text string

// Must be at least 1 items long
type Max8Text string

// May be one of DOCT, MIST, MISS, MADM
type NamePrefix1Code string

type OrganisationIdentification8 struct {
	AnyBIC AnyBICIdentifier                     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 AnyBIC,omitempty"`
	Othr   []GenericOrganisationIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 Othr,omitempty"`
}

type OrganisationIdentificationSchemeName1Choice struct {
	Cd    ExternalOrganisationIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 Cd"`
	Prtry Max35Text                               `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 Prtry"`
}

type OtherContact1 struct {
	ChanlTp Max4Text   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 ChanlTp"`
	Id      Max128Text `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 Id,omitempty"`
}

type Pagination struct {
	PgNb      Max5NumericText `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 PgNb"`
	LastPgInd bool            `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 LastPgInd"`
}

type ParentCashAccount2 struct {
	Lvl  AccountLevel1Code                            `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 Lvl,omitempty"`
	Id   CashAccount24                                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 Id"`
	Svcr BranchAndFinancialInstitutionIdentification5 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 Svcr,omitempty"`
}

type Party13Choice struct {
	OrgId OrganisationIdentification8         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 OrgId"`
	FIId  FinancialInstitutionIdentification9 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 FIId"`
}

type PartyIdentification58 struct {
	Nm        Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 Nm"`
	LglNm     Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 LglNm,omitempty"`
	PstlAdr   PostalAddress11 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 PstlAdr,omitempty"`
	Id        Party13Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 Id"`
	CtryOfRes CountryCode     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 CtryOfRes,omitempty"`
	CtctDtls  ContactDetails3 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 CtctDtls,omitempty"`
}

// Must match the pattern \+[0-9]{1,3}-[0-9()+\-]{1,30}
type PhoneNumber string

type PostalAddress11 struct {
	AdrTp       AddressType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 AdrTp,omitempty"`
	Dept        Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 Dept,omitempty"`
	SubDept     Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 SubDept,omitempty"`
	StrtNm      Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 StrtNm,omitempty"`
	BldgNb      Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 BldgNb,omitempty"`
	PstCd       Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 PstCd,omitempty"`
	TwnNm       Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 TwnNm,omitempty"`
	CtrySubDvsn Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 CtrySubDvsn,omitempty"`
	Ctry        CountryCode      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 Ctry,omitempty"`
	AdrLine     []Max70Text      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 AdrLine,omitempty"`
	Flr         Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 Flr,omitempty"`
	PstBx       Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 PstBx,omitempty"`
	BldgNm      Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 BldgNm,omitempty"`
	Room        Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 Room,omitempty"`
}

type PostalAddress6 struct {
	AdrTp       AddressType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 AdrTp,omitempty"`
	Dept        Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 Dept,omitempty"`
	SubDept     Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 SubDept,omitempty"`
	StrtNm      Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 StrtNm,omitempty"`
	BldgNb      Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 BldgNb,omitempty"`
	PstCd       Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 PstCd,omitempty"`
	TwnNm       Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 TwnNm,omitempty"`
	CtrySubDvsn Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 CtrySubDvsn,omitempty"`
	Ctry        CountryCode      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 Ctry,omitempty"`
	AdrLine     []Max70Text      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 AdrLine,omitempty"`
}

// May be one of LETT, MAIL, PHON, FAXX, CELL
type PreferredContactMethod1Code string

type ProprietaryBankTransactionCodeStructure1 struct {
	Cd   Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 Cd"`
	Issr Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 Issr,omitempty"`
}

type ReportHeader3 struct {
	RptId    Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 RptId"`
	MsgPgntn Pagination `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 MsgPgntn,omitempty"`
}

type ResidenceLocation1Choice struct {
	Ctry CountryCode `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 Ctry"`
	Area Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 Area"`
}

// May be one of COMP, NCMP
type ServiceAdjustmentType1Code string

// May be one of BCMP, FLAT, PVCH, INVS, WVED, FREE
type ServicePaymentMethod1Code string

type ServiceTaxDesignation1 struct {
	Cd     ServiceTaxDesignation1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 Cd"`
	Rgn    Max35Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 Rgn,omitempty"`
	TaxRsn []TaxReason1               `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 TaxRsn,omitempty"`
}

// May be one of XMPT, ZERO, TAXE
type ServiceTaxDesignation1Code string

type StatementGroup2 struct {
	GrpId        Max35Text             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 GrpId"`
	Sndr         PartyIdentification58 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 Sndr"`
	SndrIndvCtct []ContactDetails3     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 SndrIndvCtct,omitempty"`
	Rcvr         PartyIdentification58 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 Rcvr"`
	RcvrIndvCtct []ContactDetails3     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 RcvrIndvCtct,omitempty"`
	BllgStmt     []BillingStatement2   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 BllgStmt"`
}

type TaxCalculation1 struct {
	HstCcy                ActiveOrHistoricCurrencyCode `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 HstCcy"`
	TaxblSvcChrgConvs     []BillingServicesAmount3     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 TaxblSvcChrgConvs"`
	TtlTaxblSvcChrgHstAmt AmountAndDirection34         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 TtlTaxblSvcChrgHstAmt"`
	TaxId                 []BillingServicesTax3        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 TaxId"`
	TtlTax                AmountAndDirection34         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 TtlTax"`
}

type TaxReason1 struct {
	Cd     Max10Text  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 Cd"`
	Expltn Max105Text `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.02 Expltn"`
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
