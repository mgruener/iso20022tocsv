// Code generated by main. DO NOT EDIT.

package colr_007_001_03

import (
	"bytes"
	"encoding/xml"
	"time"
)

type AccountIdentification26 struct {
	Prtry SimpleIdentificationInformation4 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 Prtry"`
}

type AccountIdentification4Choice struct {
	IBAN IBAN2007Identifier            `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 IBAN"`
	Othr GenericAccountIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 Othr"`
}

type AccountSchemeName1Choice struct {
	Cd    ExternalAccountIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 Cd"`
	Prtry Max35Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 Prtry"`
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

// Must match the pattern [A-Z]{3,3}
type ActiveOrHistoricCurrencyCode string

// May be one of ADDR, PBOX, HOME, BIZZ, MLTO, DLVY
type AddressType2Code string

type Agreement2 struct {
	AgrmtDtls  Max140Text                `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 AgrmtDtls"`
	AgrmtId    Max140Text                `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 AgrmtId,omitempty"`
	AgrmtDt    ISODate                   `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 AgrmtDt"`
	BaseCcy    CurrencyCode              `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 BaseCcy"`
	AgrmtFrmwk AgreementFramework1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 AgrmtFrmwk,omitempty"`
}

type AgreementFramework1Choice struct {
	AgrmtFrmwk AgreementFramework1Code `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 AgrmtFrmwk"`
	PrtryId    GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 PrtryId"`
}

// May be one of FBAA, BBAA, DERV, ISDA, NONR
type AgreementFramework1Code string

type AlternatePartyIdentification5 struct {
	IdTp    IdentificationType40Choice `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 IdTp"`
	Ctry    CountryCode                `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 Ctry"`
	AltrnId Max35Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 AltrnId"`
}

// Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type AnyBICIdentifier string

type CashCollateral2 struct {
	CollId    Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 CollId,omitempty"`
	CshAcctId AccountIdentification4Choice `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 CshAcctId,omitempty"`
	RtrXcss   bool                         `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 RtrXcss,omitempty"`
	DpstAmt   ActiveCurrencyAndAmount      `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 DpstAmt,omitempty"`
	DpstTp    DepositType1Code             `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 DpstTp,omitempty"`
	MtrtyDt   ISODate                      `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 MtrtyDt,omitempty"`
	ValDt     ISODate                      `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 ValDt,omitempty"`
	XchgRate  float64                      `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 XchgRate,omitempty"`
	CollVal   ActiveCurrencyAndAmount      `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 CollVal"`
	Hrcut     float64                      `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 Hrcut,omitempty"`
}

type CashCollateral3 struct {
	CollId    Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 CollId,omitempty"`
	CshAcctId AccountIdentification4Choice `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 CshAcctId,omitempty"`
	DpstAmt   ActiveCurrencyAndAmount      `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 DpstAmt,omitempty"`
	DpstTp    DepositType1Code             `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 DpstTp,omitempty"`
	MtrtyDt   ISODate                      `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 MtrtyDt,omitempty"`
	ValDt     ISODate                      `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 ValDt,omitempty"`
	XchgRate  float64                      `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 XchgRate,omitempty"`
	CollVal   ActiveCurrencyAndAmount      `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 CollVal"`
	Hrcut     float64                      `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 Hrcut,omitempty"`
}

type Collateral7 struct {
	MrgnCallReqId   Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 MrgnCallReqId"`
	MrgnCallRspnId  Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 MrgnCallRspnId,omitempty"`
	StdSttlmInstrs  Max140Text              `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 StdSttlmInstrs,omitempty"`
	CollPrpslRspnId Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 CollPrpslRspnId,omitempty"`
	SctiesColl      []SecuritiesCollateral3 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 SctiesColl,omitempty"`
	CshColl         []CashCollateral2       `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 CshColl,omitempty"`
	OthrColl        []OtherCollateral2      `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 OthrColl,omitempty"`
}

type Collateral8 struct {
	MrgnCallReqId   Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 MrgnCallReqId"`
	MrgnCallRspnId  Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 MrgnCallRspnId,omitempty"`
	StdSttlmInstrs  Max140Text              `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 StdSttlmInstrs,omitempty"`
	CollPrpslRspnId Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 CollPrpslRspnId,omitempty"`
	SctiesColl      []SecuritiesCollateral3 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 SctiesColl,omitempty"`
	CshColl         []CashCollateral3       `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 CshColl,omitempty"`
	OthrColl        []OtherCollateral2      `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 OthrColl,omitempty"`
}

type CollateralAccount1 struct {
	Id Max35Text                                  `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 Id"`
	Tp CollateralAccountIdentificationType1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 Tp,omitempty"`
	Nm Max70Text                                  `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 Nm,omitempty"`
}

type CollateralAccountIdentificationType1Choice struct {
	Tp    CollateralAccountType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 Tp,omitempty"`
	Prtry GenericIdentification29    `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 Prtry"`
}

// May be one of HOUS, CLIE, LIPR, MGIN, DFLT
type CollateralAccountType1Code string

type CollateralMovement3Choice struct {
	CollMvmntDrctn CollateralMovement6 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 CollMvmntDrctn"`
	Rtr            Collateral7         `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 Rtr"`
}

type CollateralMovement5 struct {
	AgrdAmt    ActiveCurrencyAndAmount     `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 AgrdAmt"`
	MvmntDrctn []CollateralMovement3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 MvmntDrctn,omitempty"`
}

type CollateralMovement6 struct {
	Dlvr Collateral8 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 Dlvr"`
	Rtr  Collateral7 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 Rtr,omitempty"`
}

type CollateralOwnership1 struct {
	Prtry  bool                        `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 Prtry"`
	ClntNm PartyIdentification33Choice `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 ClntNm,omitempty"`
}

type CollateralProposal3Choice struct {
	CollPrpslDtls   CollateralProposal4 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 CollPrpslDtls"`
	SgrtdIndpdntAmt CollateralMovement5 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 SgrtdIndpdntAmt"`
}

type CollateralProposal4 struct {
	VartnMrgn       CollateralMovement5 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 VartnMrgn"`
	SgrtdIndpdntAmt CollateralMovement5 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 SgrtdIndpdntAmt,omitempty"`
}

type CollateralProposalV03 struct {
	TxId        Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 TxId"`
	Oblgtn      Obligation3          `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 Oblgtn"`
	Agrmt       Agreement2           `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 Agrmt,omitempty"`
	TpAndDtls   Proposal3            `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 TpAndDtls"`
	SplmtryData []SupplementaryData1 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 SplmtryData,omitempty"`
}

type ContactIdentification2 struct {
	NmPrfx   NamePrefix1Code `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 NmPrfx,omitempty"`
	GvnNm    Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 GvnNm,omitempty"`
	Nm       Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 Nm"`
	PhneNb   PhoneNumber     `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 PhneNb,omitempty"`
	MobNb    PhoneNumber     `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 MobNb,omitempty"`
	FaxNb    PhoneNumber     `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 FaxNb,omitempty"`
	EmailAdr Max256Text      `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 EmailAdr,omitempty"`
}

// Must match the pattern [A-Z]{2,2}
type CountryCode string

// Must match the pattern [A-Z]{3,3}
type CurrencyCode string

type DateAndDateTimeChoice struct {
	Dt   ISODate     `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 Dt"`
	DtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 DtTm"`
}

type DateCode9Choice struct {
	Cd    DateType2Code           `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 Cd"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 Prtry"`
}

type DateFormat14Choice struct {
	Dt   ISODate         `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 Dt"`
	DtCd DateCode9Choice `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 DtCd"`
}

// May be one of OPEN
type DateType2Code string

type DeliveringPartiesAndAccount10 struct {
	Dpstry PartyIdentification34Choice      `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 Dpstry"`
	Pty1   PartyIdentificationAndAccount102 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 Pty1"`
	Pty2   PartyIdentificationAndAccount77  `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 Pty2,omitempty"`
}

// May be one of FITE, CALL
type DepositType1Code string

type Document struct {
	CollPrpsl CollateralProposalV03 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 CollPrpsl"`
}

// Must match the pattern [a-zA-Z0-9]{4}
type Exact4AlphaNumericText string

// May be one of BFWD, PAYM, CCPC, COMM, CRDS, CRTL, CRSP, CCIR, CRPR, EQUI, EQPT, EQUS, EXTD, EXPT, FIXI, FORX, FORW, FUTR, OPTN, LIQU, OTCD, REPO, RVPO, SLOA, SBSC, SCRP, SLEB, SHSL, SCIR, SCIE, SWPT, TBAS, TRBD, TRCP
type ExposureType5Code string

// Must be at least 1 items long
type ExternalAccountIdentification1Code string

// Must be at least 1 items long
type ExternalFinancialInstrumentIdentificationType1Code string

type FinancialInstrumentQuantity1Choice struct {
	Unit     float64 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 Unit"`
	FaceAmt  float64 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 FaceAmt"`
	AmtsdVal float64 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 AmtsdVal"`
}

type GenericAccountIdentification1 struct {
	Id      Max34Text                `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 Id"`
	SchmeNm AccountSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 SchmeNm,omitempty"`
	Issr    Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 Issr,omitempty"`
}

type GenericIdentification29 struct {
	Id      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 Id"`
	Issr    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 Issr"`
	SchmeNm Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 SchmeNm,omitempty"`
}

type GenericIdentification30 struct {
	Id      Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 Id"`
	Issr    Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 Issr"`
	SchmeNm Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 SchmeNm,omitempty"`
}

type GenericIdentification40 struct {
	Id      Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 Id"`
	Issr    Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 Issr"`
	SchmeNm Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 SchmeNm,omitempty"`
}

type GenericIdentification58 struct {
	Id Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 Id,omitempty"`
	Tp GenericIdentification40 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 Tp"`
}

// Must match the pattern [A-Z]{2,2}[0-9]{2,2}[a-zA-Z0-9]{1,30}
type IBAN2007Identifier string

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

type IdentificationSource3Choice struct {
	Cd    ExternalFinancialInstrumentIdentificationType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 Cd"`
	Prtry Max35Text                                          `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 Prtry"`
}

type IdentificationType40Choice struct {
	Cd    TypeOfIdentification2Code `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 Cd"`
	Prtry GenericIdentification29   `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 Prtry"`
}

// Must be at least 1 items long
type Max140Text string

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

type NameAndAddress13 struct {
	Nm  Max350Text     `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 Nm"`
	Adr PostalAddress8 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 Adr,omitempty"`
}

type NameAndAddress5 struct {
	Nm  Max350Text     `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 Nm"`
	Adr PostalAddress1 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 Adr,omitempty"`
}

type NameAndAddress6 struct {
	Nm  Max70Text      `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 Nm"`
	Adr PostalAddress2 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 Adr"`
}

// May be one of DOCT, MIST, MISS, MADM
type NamePrefix1Code string

type Obligation3 struct {
	PtyA       PartyIdentification33Choice `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 PtyA"`
	SvcgPtyA   PartyIdentification33Choice `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 SvcgPtyA,omitempty"`
	PtyB       PartyIdentification33Choice `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 PtyB"`
	SvcgPtyB   PartyIdentification33Choice `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 SvcgPtyB,omitempty"`
	CollAcctId CollateralAccount1          `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 CollAcctId,omitempty"`
	XpsrTp     ExposureType5Code           `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 XpsrTp,omitempty"`
	ValtnDt    DateAndDateTimeChoice       `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 ValtnDt"`
}

type OtherCollateral2 struct {
	CollId       Max35Text                     `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 CollId,omitempty"`
	LttrOfCdtId  Max35Text                     `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 LttrOfCdtId,omitempty"`
	LttrOfCdtAmt ActiveCurrencyAndAmount       `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 LttrOfCdtAmt,omitempty"`
	GrntAmt      ActiveCurrencyAndAmount       `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 GrntAmt,omitempty"`
	OthrTpOfColl OtherTypeOfCollateral2        `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 OthrTpOfColl,omitempty"`
	IsseDt       DateFormat14Choice            `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 IsseDt,omitempty"`
	XpryDt       DateFormat14Choice            `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 XpryDt,omitempty"`
	LtdCvrgInd   bool                          `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 LtdCvrgInd,omitempty"`
	Issr         PartyIdentification33Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 Issr,omitempty"`
	ValDt        ISODate                       `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 ValDt,omitempty"`
	XchgRate     float64                       `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 XchgRate,omitempty"`
	MktVal       ActiveCurrencyAndAmount       `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 MktVal,omitempty"`
	Hrcut        float64                       `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 Hrcut,omitempty"`
	CollVal      ActiveCurrencyAndAmount       `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 CollVal"`
	SfkpgPlc     SafekeepingPlaceFormat7Choice `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 SfkpgPlc,omitempty"`
	SfkpgAcct    SecuritiesAccount19           `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 SfkpgAcct,omitempty"`
}

type OtherIdentification1 struct {
	Id  Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 Id"`
	Sfx Max16Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 Sfx,omitempty"`
	Tp  IdentificationSource3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 Tp"`
}

type OtherTypeOfCollateral2 struct {
	Desc Max140Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 Desc"`
	Qty  FinancialInstrumentQuantity1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 Qty,omitempty"`
}

type PartyIdentification32Choice struct {
	BIC      AnyBICIdentifier        `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 BIC"`
	PrtryId  GenericIdentification29 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 PrtryId"`
	NmAndAdr NameAndAddress13        `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 NmAndAdr"`
}

type PartyIdentification33Choice struct {
	AnyBIC   AnyBICIdentifier        `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 AnyBIC"`
	PrtryId  GenericIdentification29 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 PrtryId"`
	NmAndAdr NameAndAddress6         `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 NmAndAdr"`
}

type PartyIdentification34Choice struct {
	BIC      AnyBICIdentifier `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 BIC"`
	NmAndAdr NameAndAddress5  `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 NmAndAdr"`
	Ctry     CountryCode      `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 Ctry"`
}

type PartyIdentificationAndAccount102 struct {
	PtyId    PartyIdentification33Choice `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 PtyId"`
	AcctId   Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 AcctId,omitempty"`
	PrcgId   Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 PrcgId,omitempty"`
	PrcgDt   DateAndDateTimeChoice       `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 PrcgDt,omitempty"`
	SubAcct  SubAccount4                 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 SubAcct,omitempty"`
	CtctPrsn ContactIdentification2      `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 CtctPrsn,omitempty"`
}

type PartyIdentificationAndAccount77 struct {
	Id        PartyIdentification32Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 Id"`
	AltrnId   AlternatePartyIdentification5 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 AltrnId,omitempty"`
	SfkpgAcct Max35Text                     `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 SfkpgAcct,omitempty"`
	PrcgId    Max35Text                     `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 PrcgId,omitempty"`
	AddtlInf  PartyTextInformation1         `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 AddtlInf,omitempty"`
}

type PartyTextInformation1 struct {
	DclrtnDtls  Max350Text `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 DclrtnDtls,omitempty"`
	PtyCtctDtls Max140Text `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 PtyCtctDtls,omitempty"`
	RegnDtls    Max350Text `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 RegnDtls,omitempty"`
}

// Must match the pattern \+[0-9]{1,3}-[0-9()+\-]{1,30}
type PhoneNumber string

type PostalAddress1 struct {
	AdrTp       AddressType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 AdrTp,omitempty"`
	AdrLine     []Max70Text      `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 AdrLine,omitempty"`
	StrtNm      Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 StrtNm,omitempty"`
	BldgNb      Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 BldgNb,omitempty"`
	PstCd       Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 PstCd,omitempty"`
	TwnNm       Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 TwnNm,omitempty"`
	CtrySubDvsn Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 CtrySubDvsn,omitempty"`
	Ctry        CountryCode      `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 Ctry"`
}

type PostalAddress2 struct {
	StrtNm      Max70Text   `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 StrtNm,omitempty"`
	PstCdId     Max16Text   `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 PstCdId"`
	TwnNm       Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 TwnNm"`
	CtrySubDvsn Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 CtrySubDvsn,omitempty"`
	Ctry        CountryCode `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 Ctry"`
}

type PostalAddress8 struct {
	AdrTp       AddressType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 AdrTp,omitempty"`
	AdrLine     []Max70Text      `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 AdrLine,omitempty"`
	StrtNm      Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 StrtNm,omitempty"`
	BldgNb      Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 BldgNb,omitempty"`
	PstCd       Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 PstCd,omitempty"`
	TwnNm       Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 TwnNm,omitempty"`
	CtrySubDvsn Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 CtrySubDvsn,omitempty"`
	Ctry        CountryCode      `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 Ctry"`
}

type Price2 struct {
	Tp  YieldedOrValueType1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 Tp"`
	Val PriceRateOrAmountChoice   `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 Val"`
}

type PriceRateOrAmountChoice struct {
	Rate float64                                    `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 Rate"`
	Amt  ActiveOrHistoricCurrencyAnd13DecimalAmount `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 Amt"`
}

// May be one of DISC, PREM, PARV
type PriceValueType1Code string

type Proposal3 struct {
	CollPrpslTp ProposalType1Code         `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 CollPrpslTp"`
	CollPrpsl   CollateralProposal3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 CollPrpsl"`
}

// May be one of INIT, COUN
type ProposalType1Code string

type ReceivingPartiesAndAccount10 struct {
	Dpstry PartyIdentification34Choice      `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 Dpstry"`
	Pty1   PartyIdentificationAndAccount102 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 Pty1"`
	Pty2   PartyIdentificationAndAccount77  `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 Pty2,omitempty"`
}

// May be one of CUST, ICSD, NCSD, SHHE
type SafekeepingPlace1Code string

// May be one of SHHE
type SafekeepingPlace3Code string

type SafekeepingPlaceFormat7Choice struct {
	Id      SafekeepingPlaceTypeAndText1             `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 Id"`
	Ctry    CountryCode                              `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 Ctry"`
	TpAndId SafekeepingPlaceTypeAndAnyBICIdentifier1 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 TpAndId"`
	Prtry   GenericIdentification58                  `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 Prtry"`
}

type SafekeepingPlaceTypeAndAnyBICIdentifier1 struct {
	SfkpgPlcTp SafekeepingPlace1Code `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 SfkpgPlcTp"`
	Id         AnyBICIdentifier      `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 Id"`
}

type SafekeepingPlaceTypeAndText1 struct {
	SfkpgPlcTp SafekeepingPlace3Code `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 SfkpgPlcTp"`
	Id         Max35Text             `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 Id,omitempty"`
}

type SecuritiesAccount19 struct {
	Id Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 Id"`
	Tp GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 Tp,omitempty"`
	Nm Max70Text               `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 Nm,omitempty"`
}

type SecuritiesCollateral3 struct {
	CollId      Max35Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 CollId,omitempty"`
	SctyId      SecurityIdentification14           `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 SctyId"`
	MtrtyDt     DateAndDateTimeChoice              `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 MtrtyDt,omitempty"`
	LtdCvrgInd  bool                               `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 LtdCvrgInd,omitempty"`
	Qty         FinancialInstrumentQuantity1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 Qty"`
	Pric        Price2                             `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 Pric,omitempty"`
	MktVal      ActiveCurrencyAndAmount            `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 MktVal,omitempty"`
	Hrcut       float64                            `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 Hrcut,omitempty"`
	CollVal     ActiveCurrencyAndAmount            `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 CollVal,omitempty"`
	ValDt       ISODate                            `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 ValDt,omitempty"`
	SfkpgAcct   SecuritiesAccount19                `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 SfkpgAcct,omitempty"`
	SfkpgPlc    SafekeepingPlaceFormat7Choice      `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 SfkpgPlc"`
	SttlmParams SettlementDetails88                `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 SttlmParams,omitempty"`
}

type SecurityIdentification14 struct {
	ISIN   ISINIdentifier         `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 ISIN,omitempty"`
	OthrId []OtherIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 OthrId,omitempty"`
	Desc   Max140Text             `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 Desc,omitempty"`
}

type SettlementDetails88 struct {
	TradDt     ISODateTime              `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 TradDt"`
	SttlmPties SettlementParties3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 SttlmPties,omitempty"`
	CollOwnrsh CollateralOwnership1     `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 CollOwnrsh"`
}

type SettlementParties3Choice struct {
	DlvrgSttlmPties DeliveringPartiesAndAccount10 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 DlvrgSttlmPties,omitempty"`
	RcvgSttlmPties  ReceivingPartiesAndAccount10  `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 RcvgSttlmPties,omitempty"`
}

type SimpleIdentificationInformation4 struct {
	Id Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 Id"`
}

type SubAccount4 struct {
	Id    AccountIdentification26 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 Id"`
	Nm    Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 Nm,omitempty"`
	Chrtc Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 Chrtc,omitempty"`
}

type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
}

// May be one of ARNU, CHTY, CORP, FIIN, TXID
type TypeOfIdentification2Code string

type YieldedOrValueType1Choice struct {
	Yldd  bool                `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 Yldd"`
	ValTp PriceValueType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.03 ValTp"`
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
