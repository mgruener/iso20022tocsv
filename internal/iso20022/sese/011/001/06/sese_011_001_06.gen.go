// Code generated by main. DO NOT EDIT.

package sese_011_001_06

import (
	"bytes"
	"encoding/xml"
	"time"
)

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

type AdditionalReference7 struct {
	Ref     Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 Ref"`
	RefIssr PartyIdentification97Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 RefIssr,omitempty"`
	MsgNm   Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 MsgNm,omitempty"`
}

type AdditionalReference8 struct {
	Ref     Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 Ref"`
	RefIssr PartyIdentification113 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 RefIssr,omitempty"`
	MsgNm   Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 MsgNm,omitempty"`
}

// May be one of ADDR, PBOX, HOME, BIZZ, MLTO, DLVY
type AddressType2Code string

// Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type AnyBICIdentifier string

type CancellationPendingStatus7Choice struct {
	Rsn          Max350Text             `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 Rsn"`
	DataSrcSchme GenericIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 DataSrcSchme"`
	NoSpcfdRsn   NoReasonCode           `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 NoSpcfdRsn"`
}

type CancelledStatus13Choice struct {
	NoSpcfdRsn   NoReasonCode               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 NoSpcfdRsn"`
	Rsn          CancelledStatusReason3Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 Rsn"`
	XtndedRsn    Extended350Code            `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 XtndedRsn"`
	DataSrcSchme GenericIdentification1     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 DataSrcSchme"`
}

// May be one of CNTA, CNCL, CNIN
type CancelledStatusReason3Code string

type ChargeBasis2Choice struct {
	Cd    TaxationBasis5Code      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 Cd"`
	Prtry GenericIdentification47 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 Prtry"`
}

type ChargeOrCommissionDiscount1 struct {
	Amt  ActiveCurrencyAndAmount   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 Amt,omitempty"`
	Rate float64                   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 Rate,omitempty"`
	Bsis WaivingInstruction2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 Bsis,omitempty"`
}

type ChargeType5Choice struct {
	Cd    InvestmentFundFee1Code  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 Cd"`
	Prtry GenericIdentification47 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 Prtry"`
}

// Must match the pattern [A-Z]{2,2}
type CountryCode string

type Document struct {
	TrfInstrStsRpt TransferInstructionStatusReportV06 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 TrfInstrStsRpt"`
}

// Must match the pattern [a-zA-Z0-9]{4}
type Exact4AlphaNumericText string

// Must be at least 1 items long
type Extended350Code string

type Extension1 struct {
	PlcAndNm Max350Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 PlcAndNm"`
	Txt      Max350Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 Txt"`
}

type FailedSettlementStatus2Choice struct {
	Rsn          Max350Text             `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 Rsn"`
	DataSrcSchme GenericIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 DataSrcSchme"`
	NoSpcfdRsn   NoReasonCode           `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 NoSpcfdRsn"`
}

type Fee2 struct {
	Tp           ChargeType5Choice           `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 Tp"`
	Bsis         ChargeBasis2Choice          `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 Bsis,omitempty"`
	StdAmt       ActiveCurrencyAndAmount     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 StdAmt,omitempty"`
	StdRate      float64                     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 StdRate,omitempty"`
	DscntDtls    ChargeOrCommissionDiscount1 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 DscntDtls,omitempty"`
	ApldAmt      ActiveCurrencyAndAmount     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 ApldAmt,omitempty"`
	ApldRate     float64                     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 ApldRate,omitempty"`
	NonStdSLARef Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 NonStdSLARef,omitempty"`
	RcptId       PartyIdentification113      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 RcptId,omitempty"`
	InftvInd     bool                        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 InftvInd"`
}

type GenericIdentification1 struct {
	Id      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 Id"`
	SchmeNm Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 SchmeNm,omitempty"`
	Issr    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 Issr,omitempty"`
}

type GenericIdentification36 struct {
	Id      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 Id"`
	Issr    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 Issr"`
	SchmeNm Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 SchmeNm,omitempty"`
}

type GenericIdentification47 struct {
	Id      Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 Id"`
	Issr    Max4AlphaNumericText   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 Issr"`
	SchmeNm Max4AlphaNumericText   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 SchmeNm,omitempty"`
}

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

type InRepairStatus4Choice struct {
	Rsn          Max350Text             `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 Rsn"`
	DataSrcSchme GenericIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 DataSrcSchme"`
	NoSpcfdRsn   NoReasonCode           `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 NoSpcfdRsn"`
}

// May be one of BEND, BRKF, COMM, CDPL, CDSC, CBCH, DLEV, FEND, INIT, ADDF, POST, PREM, CHAR, SHIP, SWIT, UCIC, REGF, PENA
type InvestmentFundFee1Code string

// Must match the pattern [A-Z0-9]{18,18}[0-9]{2,2}
type LEIIdentifier string

type MarketPracticeVersion1 struct {
	Nm Max35Text    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 Nm"`
	Dt ISOYearMonth `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 Dt,omitempty"`
	Nb Max35Text    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 Nb,omitempty"`
}

// Must be at least 1 items long
type Max16Text string

// Must be at least 1 items long
type Max350Text string

// Must be at least 1 items long
type Max35Text string

// Must match the pattern [a-zA-Z0-9]{1,4}
type Max4AlphaNumericText string

// Must be at least 1 items long
type Max70Text string

type MessageIdentification1 struct {
	Id      Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 Id"`
	CreDtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 CreDtTm"`
}

type NameAndAddress5 struct {
	Nm  Max350Text     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 Nm"`
	Adr PostalAddress1 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 Adr,omitempty"`
}

// May be one of NORE
type NoReasonCode string

type OtherAmount1 struct {
	Tp  OtherAmountType1Choice  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 Tp"`
	Amt ActiveCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 Amt"`
}

type OtherAmountType1Choice struct {
	Cd      OtherAmountType1Code   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 Cd"`
	PrtryCd GenericIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 PrtryCd"`
}

// May be one of PINT, SINT
type OtherAmountType1Code string

type PartyIdentification113 struct {
	Pty PartyIdentification90Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 Pty"`
	LEI LEIIdentifier               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 LEI,omitempty"`
}

type PartyIdentification90Choice struct {
	AnyBIC   AnyBICIdentifier       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 AnyBIC"`
	PrtryId  GenericIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 PrtryId"`
	NmAndAdr NameAndAddress5        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 NmAndAdr"`
}

type PartyIdentification97Choice struct {
	AnyBIC     AnyBICIdentifier       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 AnyBIC"`
	LglNttyIdr LEIIdentifier          `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 LglNttyIdr"`
	NmAndAdr   NameAndAddress5        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 NmAndAdr"`
	PrtryId    GenericIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 PrtryId"`
}

type PendingSettlementStatus3Choice struct {
	Rsn          PendingSettlementStatusReason2Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 Rsn"`
	XtndedRsn    Extended350Code                    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 XtndedRsn"`
	DataSrcSchme GenericIdentification1             `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 DataSrcSchme"`
	NoSpcfdRsn   NoReasonCode                       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 NoSpcfdRsn"`
}

// May be one of AWSH, BLOC, CAIS, CLAC, DOCC, DOCY, IAAD, LACK, LINK, PHCK, PHSE, SBLO, MINF, ACOP, IINV, CINV, AINV, WTRF, USUA, ASTA, AFST, STST, LPRO, ADRQ, ADS1, ADS2, DRJC, CYIN, CYDV, OVER, WCPA, SDUT, TAPR, XCNF, ESCA, NRCP, FVER
type PendingSettlementStatusReason2Code string

type PostalAddress1 struct {
	AdrTp       AddressType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 AdrTp,omitempty"`
	AdrLine     []Max70Text      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 AdrLine,omitempty"`
	StrtNm      Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 StrtNm,omitempty"`
	BldgNb      Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 BldgNb,omitempty"`
	PstCd       Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 PstCd,omitempty"`
	TwnNm       Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 TwnNm,omitempty"`
	CtrySubDvsn Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 CtrySubDvsn,omitempty"`
	Ctry        CountryCode      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 Ctry"`
}

// May be one of FORW, HIST
type PriceMethod1Code string

type PriceValue1 struct {
	Amt ActiveCurrencyAnd13DecimalAmount `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 Amt"`
}

type References61Choice struct {
	RltdRef []AdditionalReference8 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 RltdRef"`
	OthrRef []AdditionalReference8 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 OthrRef"`
}

type RejectedReason15Choice struct {
	Cd    TransferRejectedStatusReason2Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 Cd"`
	Prtry GenericIdentification36           `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 Prtry"`
}

type RejectionReason32 struct {
	Rsn         RejectedReason15Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 Rsn"`
	AddtlRsnInf Max350Text             `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 AddtlRsnInf,omitempty"`
}

type ReversedStatus2Choice struct {
	Rsn          Max350Text             `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 Rsn"`
	DataSrcSchme GenericIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 DataSrcSchme"`
	NoSpcfdRsn   NoReasonCode           `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 NoSpcfdRsn"`
}

type Tax31 struct {
	Tp           TaxType3Choice              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 Tp"`
	ApldAmt      ActiveCurrencyAndAmount     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 ApldAmt"`
	ApldRate     float64                     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 ApldRate,omitempty"`
	Ctry         CountryCode                 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 Ctry,omitempty"`
	RcptId       PartyIdentification113      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 RcptId,omitempty"`
	TaxClctnDtls TaxCalculationInformation10 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 TaxClctnDtls,omitempty"`
}

type TaxBasis1Choice struct {
	Cd    TaxationBasis2Code      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 Cd"`
	Prtry GenericIdentification47 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 Prtry"`
}

type TaxCalculationInformation10 struct {
	Bsis     TaxBasis1Choice         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 Bsis,omitempty"`
	TaxblAmt ActiveCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 TaxblAmt"`
}

// May be one of PROV, NATI, STAT, WITH, KAPA, NKAP, INPO, STAM, WTAX, INHT, SOSU, CTAX, GIFT, COAX, EUTR, AKT1, AKT2, ZWIS
type TaxType17Code string

type TaxType3Choice struct {
	Cd    TaxType17Code           `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 Cd"`
	Prtry GenericIdentification47 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 Prtry"`
}

type TaxableIncomePerShareCalculated2Choice struct {
	Cd    TaxableIncomePerShareCalculated2Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 Cd"`
	Prtry GenericIdentification47              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 Prtry"`
}

// May be one of TSIY, TSIN, UKWN
type TaxableIncomePerShareCalculated2Code string

// May be one of FLAT, PERU
type TaxationBasis2Code string

// May be one of FLAT, GRAM, NEAM, NAVP, PERU
type TaxationBasis5Code string

type TotalFeesAndTaxes40 struct {
	TtlOvrhdApld   ActiveCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 TtlOvrhdApld,omitempty"`
	TtlFees        ActiveCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 TtlFees,omitempty"`
	TtlTaxs        ActiveCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 TtlTaxs,omitempty"`
	ComrclAgrmtRef Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 ComrclAgrmtRef,omitempty"`
	IndvFee        []Fee2                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 IndvFee,omitempty"`
	IndvTax        []Tax31                 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 IndvTax,omitempty"`
}

type TransferInstructionStatus4 struct {
	Sts TransferStatus4Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 Sts"`
	Rsn Max350Text          `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 Rsn,omitempty"`
}

type TransferInstructionStatusReportV06 struct {
	MsgId        MessageIdentification1   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 MsgId"`
	CtrPtyRef    AdditionalReference7     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 CtrPtyRef,omitempty"`
	Ref          References61Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 Ref,omitempty"`
	StsRpt       TransferStatusAndReason5 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 StsRpt"`
	MktPrctcVrsn MarketPracticeVersion1   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 MktPrctcVrsn,omitempty"`
	Xtnsn        []Extension1             `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 Xtnsn,omitempty"`
}

// May be one of DDAT, DEPT, DSEC, SECU, ISTP, ICTN, SAFE, IAQD, BLCA, DOCC, IDNA, DLVY, LEGL, NSLA, DQUA, INUK, INID, INAC, INNA, INPM, CYPA, PTNS, FTAX, ISAT, CASH, TREF
type TransferRejectedStatusReason2Code string

type TransferStatus2Choice struct {
	Sts        TransferInstructionStatus4       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 Sts"`
	PdgSttlm   PendingSettlementStatus3Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 PdgSttlm"`
	Umtchd     TransferUnmatchedStatus3Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 Umtchd"`
	InRpr      InRepairStatus4Choice            `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 InRpr"`
	Rjctd      []RejectionReason32              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 Rjctd"`
	FaildSttlm FailedSettlementStatus2Choice    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 FaildSttlm"`
	Canc       CancelledStatus13Choice          `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 Canc"`
	Rvsd       ReversedStatus2Choice            `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 Rvsd"`
	CxlPdg     CancellationPendingStatus7Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 CxlPdg"`
}

// May be one of PACK, COSE, MACH, RECE, STNP, SETT, COMP
type TransferStatus4Code string

type TransferStatusAndReason5 struct {
	MstrRef    Max35Text                                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 MstrRef,omitempty"`
	TrfRef     Max35Text                                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 TrfRef"`
	ClntRef    AdditionalReference7                       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 ClntRef,omitempty"`
	CxlRef     Max35Text                                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 CxlRef,omitempty"`
	TrfSts     TransferStatus2Choice                      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 TrfSts"`
	TradDt     ISODate                                    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 TradDt,omitempty"`
	SndOutDt   ISODate                                    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 SndOutDt,omitempty"`
	TtlUnitsNb float64                                    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 TtlUnitsNb,omitempty"`
	AvrgPric   ActiveOrHistoricCurrencyAnd13DecimalAmount `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 AvrgPric,omitempty"`
	UnitsDtls  []Unit8                                    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 UnitsDtls,omitempty"`
	StsInitr   PartyIdentification113                     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 StsInitr,omitempty"`
}

// May be one of CMIS, CPCA, DELN, DSEC, PHYS, PODU, DEPT, DDAT, DQUA
type TransferUnmatchedReason2Code string

type TransferUnmatchedStatus3Choice struct {
	NoSpcfdRsn   NoReasonCode                 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 NoSpcfdRsn"`
	Rsn          TransferUnmatchedReason2Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 Rsn"`
	XtndedRsn    Extended350Code              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 XtndedRsn"`
	DataSrcSchme GenericIdentification1       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 DataSrcSchme"`
}

// May be one of BIDE, OFFR, NAVL, CREA, CANC, INTE, SWNG, MIDD, RINV, SWIC, DDVR, ACTU
type TypeOfPrice10Code string

type TypeOfPrice46Choice struct {
	Cd    TypeOfPrice10Code       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 Cd"`
	Prtry GenericIdentification47 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 Prtry"`
}

// May be one of GRP1, GRP2
type UKTaxGroupUnit1Code string

type Unit8 struct {
	UnitsNb      float64             `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 UnitsNb"`
	OrdrDt       ISODate             `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 OrdrDt,omitempty"`
	AcqstnDt     ISODate             `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 AcqstnDt,omitempty"`
	CertNb       []Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 CertNb,omitempty"`
	Grp1Or2Units UKTaxGroupUnit1Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 Grp1Or2Units,omitempty"`
	Ref          Max35Text           `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 Ref,omitempty"`
	PricDtls     UnitPrice23         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 PricDtls,omitempty"`
	TxOvrhd      TotalFeesAndTaxes40 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 TxOvrhd,omitempty"`
	OthrAmt      []OtherAmount1      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 OthrAmt,omitempty"`
}

type UnitPrice23 struct {
	Tp                   TypeOfPrice46Choice                    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 Tp"`
	Val                  PriceValue1                            `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 Val"`
	PricMtd              PriceMethod1Code                       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 PricMtd,omitempty"`
	AcrdIntrstNAV        ActiveOrHistoricCurrencyAndAmount      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 AcrdIntrstNAV,omitempty"`
	NbOfDaysAcrd         float64                                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 NbOfDaysAcrd,omitempty"`
	TaxblIncmPerShr      ActiveCurrencyAnd13DecimalAmount       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 TaxblIncmPerShr,omitempty"`
	TaxblIncmPerShrClctd TaxableIncomePerShareCalculated2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 TaxblIncmPerShrClctd,omitempty"`
}

// May be one of WICA, WIUN
type WaivingInstruction1Code string

type WaivingInstruction2Choice struct {
	Cd    WaivingInstruction1Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 Cd"`
	Prtry GenericIdentification47 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.06 Prtry"`
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
