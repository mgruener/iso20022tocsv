// Code generated by main. DO NOT EDIT.

package auth_033_001_01

import (
	"bytes"
	"encoding/xml"
	"time"
)

// Must match the pattern [A-Z]{3,3}
type ActiveCurrencyCode string

// Must match the pattern [A-Z]{3,3}
type ActiveOrHistoricCurrencyCode string

// May be one of DLVR, NDLV
type AssetClassSubProductType19Code string

// May be one of WIBO, TREA, TIBO, TLBO, SWAP, STBO, PRBO, PFAN, NIBO, MAAA, MOSP, LIBO, LIBI, JIBA, ISDA, GCFR, FUSW, EUCH, EUUS, EURI, EONS, EONA, CIBO, CDOR, BUBO, BBSW
type BenchmarkCurveName2Code string

type BenchmarkCurveName5Choice struct {
	Indx BenchmarkCurveName2Code `xml:"urn:iso:std:iso:20022:tech:xsd:auth.033.001.01 Indx"`
	Nm   Max25Text               `xml:"urn:iso:std:iso:20022:tech:xsd:auth.033.001.01 Nm"`
}

type BondDerivative2 struct {
	Issr    LEIIdentifier `xml:"urn:iso:std:iso:20022:tech:xsd:auth.033.001.01 Issr"`
	MtrtyDt ISODate       `xml:"urn:iso:std:iso:20022:tech:xsd:auth.033.001.01 MtrtyDt,omitempty"`
	IssncDt ISODate       `xml:"urn:iso:std:iso:20022:tech:xsd:auth.033.001.01 IssncDt,omitempty"`
}

// May be one of EUSB, OEPB, CVTB, CRPB, CVDB, OTHR
type BondType1Code string

type CommodityDerivate2Choice struct {
	Frght CommodityDerivate5 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.033.001.01 Frght"`
	Nrgy  CommodityDerivate6 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.033.001.01 Nrgy"`
}

type CommodityDerivate4 struct {
	ClssSpcfc CommodityDerivate2Choice     `xml:"urn:iso:std:iso:20022:tech:xsd:auth.033.001.01 ClssSpcfc,omitempty"`
	NtnlCcy   ActiveOrHistoricCurrencyCode `xml:"urn:iso:std:iso:20022:tech:xsd:auth.033.001.01 NtnlCcy"`
}

type CommodityDerivate5 struct {
	Sz          Max25Text `xml:"urn:iso:std:iso:20022:tech:xsd:auth.033.001.01 Sz"`
	AvrgTmChrtr Max25Text `xml:"urn:iso:std:iso:20022:tech:xsd:auth.033.001.01 AvrgTmChrtr"`
}

type CommodityDerivate6 struct {
	SttlmLctn Max25Text `xml:"urn:iso:std:iso:20022:tech:xsd:auth.033.001.01 SttlmLctn"`
}

type ContractForDifference2 struct {
	UndrlygTp UnderlyingContractForDifferenceType3Code `xml:"urn:iso:std:iso:20022:tech:xsd:auth.033.001.01 UndrlygTp"`
	NtnlCcy1  ActiveOrHistoricCurrencyCode             `xml:"urn:iso:std:iso:20022:tech:xsd:auth.033.001.01 NtnlCcy1,omitempty"`
	NtnlCcy2  ActiveOrHistoricCurrencyCode             `xml:"urn:iso:std:iso:20022:tech:xsd:auth.033.001.01 NtnlCcy2,omitempty"`
}

// Must match the pattern [A-Z]{2,2}
type CountryCode string

// Must match the pattern [A-Z]{2,2}\-[0-9A-Z]{1,3}
type CountrySubDivisionCode string

type CreditDefaultSwapDerivative3 struct {
	UndrlygIndxId ISINOct2015Identifier   `xml:"urn:iso:std:iso:20022:tech:xsd:auth.033.001.01 UndrlygIndxId,omitempty"`
	IndxNm        Max25Text               `xml:"urn:iso:std:iso:20022:tech:xsd:auth.033.001.01 IndxNm"`
	Indx          CreditDefaultSwapIndex2 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.033.001.01 Indx"`
}

type CreditDefaultSwapDerivative4 struct {
	UndrlygNmId ISINOct2015Identifier        `xml:"urn:iso:std:iso:20022:tech:xsd:auth.033.001.01 UndrlygNmId,omitempty"`
	OblgtnId    ISINOct2015Identifier        `xml:"urn:iso:std:iso:20022:tech:xsd:auth.033.001.01 OblgtnId"`
	SnglNm      CreditDefaultSwapSingleName2 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.033.001.01 SnglNm"`
}

type CreditDefaultSwapIndex2 struct {
	Srs       float64                       `xml:"urn:iso:std:iso:20022:tech:xsd:auth.033.001.01 Srs,omitempty"`
	Vrsn      float64                       `xml:"urn:iso:std:iso:20022:tech:xsd:auth.033.001.01 Vrsn,omitempty"`
	RollMnth  []RestrictedMonthExact2Number `xml:"urn:iso:std:iso:20022:tech:xsd:auth.033.001.01 RollMnth,omitempty"`
	NxtRollDt ISODate                       `xml:"urn:iso:std:iso:20022:tech:xsd:auth.033.001.01 NxtRollDt,omitempty"`
	NtnlCcy   ActiveOrHistoricCurrencyCode  `xml:"urn:iso:std:iso:20022:tech:xsd:auth.033.001.01 NtnlCcy"`
}

type CreditDefaultSwapSingleName2 struct {
	SvrgnIssr bool                                 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.033.001.01 SvrgnIssr"`
	RefPty    DerivativePartyIdentification1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:auth.033.001.01 RefPty,omitempty"`
	NtnlCcy   ActiveOrHistoricCurrencyCode         `xml:"urn:iso:std:iso:20022:tech:xsd:auth.033.001.01 NtnlCcy"`
}

type CreditDefaultSwapsDerivative3Choice struct {
	SnglNmCdtDfltSwp      CreditDefaultSwapSingleName2 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.033.001.01 SnglNmCdtDfltSwp"`
	CdtDfltSwpIndx        CreditDefaultSwapIndex2      `xml:"urn:iso:std:iso:20022:tech:xsd:auth.033.001.01 CdtDfltSwpIndx"`
	SnglNmCdtDfltSwpDeriv CreditDefaultSwapDerivative4 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.033.001.01 SnglNmCdtDfltSwpDeriv"`
	CdtDfltSwpIndxDeriv   CreditDefaultSwapDerivative3 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.033.001.01 CdtDfltSwpIndxDeriv"`
}

type DebtInstrument5 struct {
	Tp      BondType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:auth.033.001.01 Tp"`
	IssncDt ISODate       `xml:"urn:iso:std:iso:20022:tech:xsd:auth.033.001.01 IssncDt"`
}

type Derivative2Choice struct {
	Cmmdty       CommodityDerivate4                  `xml:"urn:iso:std:iso:20022:tech:xsd:auth.033.001.01 Cmmdty"`
	IntrstRate   InterestRateDerivative5             `xml:"urn:iso:std:iso:20022:tech:xsd:auth.033.001.01 IntrstRate"`
	FX           ForeignExchangeDerivative2          `xml:"urn:iso:std:iso:20022:tech:xsd:auth.033.001.01 FX"`
	Eqty         EquityDerivative2                   `xml:"urn:iso:std:iso:20022:tech:xsd:auth.033.001.01 Eqty"`
	CtrctForDiff ContractForDifference2              `xml:"urn:iso:std:iso:20022:tech:xsd:auth.033.001.01 CtrctForDiff"`
	Cdt          CreditDefaultSwapsDerivative3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:auth.033.001.01 Cdt"`
	EmssnAllwnc  EmissionAllowanceProductType1Code   `xml:"urn:iso:std:iso:20022:tech:xsd:auth.033.001.01 EmssnAllwnc"`
}

type DerivativePartyIdentification1Choice struct {
	Ctry        CountryCode            `xml:"urn:iso:std:iso:20022:tech:xsd:auth.033.001.01 Ctry"`
	CtrySubDvsn CountrySubDivisionCode `xml:"urn:iso:std:iso:20022:tech:xsd:auth.033.001.01 CtrySubDvsn"`
	LEI         LEIIdentifier          `xml:"urn:iso:std:iso:20022:tech:xsd:auth.033.001.01 LEI"`
}

type Document struct {
	FinInstrmRptgNonEqtyTrnsprncyDataRpt FinancialInstrumentReportingNonEquityTransparencyDataReportV01 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.033.001.01 FinInstrmRptgNonEqtyTrnsprncyDataRpt"`
}

// May be one of EUAA, EUAE, ERUE, CERE, OTHR
type EmissionAllowanceProductType1Code string

// May be one of CERE, ERUE, EUAE, EUAA
type EmissionAllowanceProductType2Code string

type EquityDerivative2 struct {
	UndrlygTp EquityDerivative3Choice    `xml:"urn:iso:std:iso:20022:tech:xsd:auth.033.001.01 UndrlygTp"`
	Param     EquityReturnParameter1Code `xml:"urn:iso:std:iso:20022:tech:xsd:auth.033.001.01 Param,omitempty"`
}

type EquityDerivative3Choice struct {
	Bskt   UnderlyingEquityType3Code `xml:"urn:iso:std:iso:20022:tech:xsd:auth.033.001.01 Bskt"`
	Indx   UnderlyingEquityType4Code `xml:"urn:iso:std:iso:20022:tech:xsd:auth.033.001.01 Indx"`
	SnglNm UnderlyingEquityType5Code `xml:"urn:iso:std:iso:20022:tech:xsd:auth.033.001.01 SnglNm"`
	Othr   UnderlyingEquityType6Code `xml:"urn:iso:std:iso:20022:tech:xsd:auth.033.001.01 Othr"`
}

// May be one of PRDV, PRVA, PRVO, PRBP
type EquityReturnParameter1Code string

// May be one of CFDS, FORW, FRAS, FUTR, OPTN, OTHR, SPDB, SWAP, SWPT, FONS, PSWP, FFAS, FWOS
type FinancialInstrumentContractType1Code string

type FinancialInstrumentReportingNonEquityTransparencyDataReportV01 struct {
	RptHdr               SecuritiesMarketReportHeader1 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.033.001.01 RptHdr"`
	NonEqtyTrnsprncyData []TransparencyDataReport10    `xml:"urn:iso:std:iso:20022:tech:xsd:auth.033.001.01 NonEqtyTrnsprncyData"`
	SplmtryData          []SupplementaryData1          `xml:"urn:iso:std:iso:20022:tech:xsd:auth.033.001.01 SplmtryData,omitempty"`
}

type FloatingInterestRate8 struct {
	RefRate BenchmarkCurveName5Choice `xml:"urn:iso:std:iso:20022:tech:xsd:auth.033.001.01 RefRate"`
	Term    InterestRateContractTerm2 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.033.001.01 Term,omitempty"`
}

type ForeignExchangeDerivative2 struct {
	CtrctSubTp AssetClassSubProductType19Code `xml:"urn:iso:std:iso:20022:tech:xsd:auth.033.001.01 CtrctSubTp"`
}

// Must match the pattern [A-Z]{2,2}[A-Z0-9]{9,9}[0-9]{1,1}
type ISINOct2015Identifier string

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

type InflationIndex1Choice struct {
	ISIN ISINOct2015Identifier `xml:"urn:iso:std:iso:20022:tech:xsd:auth.033.001.01 ISIN"`
	Nm   Max25Text             `xml:"urn:iso:std:iso:20022:tech:xsd:auth.033.001.01 Nm"`
}

type InterestRateContractTerm2 struct {
	Unit RateBasis1Code `xml:"urn:iso:std:iso:20022:tech:xsd:auth.033.001.01 Unit"`
	Val  float64        `xml:"urn:iso:std:iso:20022:tech:xsd:auth.033.001.01 Val"`
}

type InterestRateDerivative2Choice struct {
	SwpRltd SwapType1Code                   `xml:"urn:iso:std:iso:20022:tech:xsd:auth.033.001.01 SwpRltd"`
	Othr    UnderlyingInterestRateType3Code `xml:"urn:iso:std:iso:20022:tech:xsd:auth.033.001.01 Othr"`
}

type InterestRateDerivative5 struct {
	UndrlygTp         InterestRateDerivative2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:auth.033.001.01 UndrlygTp"`
	UndrlygBd         BondDerivative2               `xml:"urn:iso:std:iso:20022:tech:xsd:auth.033.001.01 UndrlygBd,omitempty"`
	SwptnNtnlCcy      ActiveCurrencyCode            `xml:"urn:iso:std:iso:20022:tech:xsd:auth.033.001.01 SwptnNtnlCcy,omitempty"`
	UndrlygSwpMtrtyDt ISODate                       `xml:"urn:iso:std:iso:20022:tech:xsd:auth.033.001.01 UndrlygSwpMtrtyDt,omitempty"`
	InfltnIndx        InflationIndex1Choice         `xml:"urn:iso:std:iso:20022:tech:xsd:auth.033.001.01 InfltnIndx,omitempty"`
	IntrstRateRef     FloatingInterestRate8         `xml:"urn:iso:std:iso:20022:tech:xsd:auth.033.001.01 IntrstRateRef"`
}

// Must match the pattern [A-Z0-9]{18,18}[0-9]{2,2}
type LEIIdentifier string

// Must match the pattern [A-Z0-9]{4,4}
type MICIdentifier string

// Must be at least 1 items long
type Max25Text string

// Must be at least 1 items long
type Max350Text string

// Must be at least 1 items long
type Max35Text string

// Must be at least 1 items long
type Max50Text string

// May be one of SFPS, SDRV, DERV, EMAL, BOND, ETCS, ETNS
type NonEquityInstrumentReportingClassification1Code string

type Period2 struct {
	FrDt ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:auth.033.001.01 FrDt"`
	ToDt ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:auth.033.001.01 ToDt"`
}

type Period4Choice struct {
	Dt       ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:auth.033.001.01 Dt"`
	FrDt     ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:auth.033.001.01 FrDt"`
	ToDt     ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:auth.033.001.01 ToDt"`
	FrDtToDt Period2 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.033.001.01 FrDtToDt"`
}

// May be one of EMAL, INTR, EQUI, COMM, CRDT, CURR
type ProductType5Code string

// May be one of DAYS, MNTH, WEEK, YEAR
type RateBasis1Code string

// Must match the pattern [0-9]{2,2}
type RestrictedMonthExact2Number float64

type SecuritiesMarketReportHeader1 struct {
	RptgNtty     TradingVenueIdentification1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:auth.033.001.01 RptgNtty"`
	RptgPrd      Period4Choice                     `xml:"urn:iso:std:iso:20022:tech:xsd:auth.033.001.01 RptgPrd"`
	SubmissnDtTm ISODateTime                       `xml:"urn:iso:std:iso:20022:tech:xsd:auth.033.001.01 SubmissnDtTm,omitempty"`
}

type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.033.001.01 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.033.001.01 Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
}

// May be one of OSSC, XFSC, XFMC, XXSC, XXMC, IFMC, FFSC, FFMC, IFSC, OSMC
type SwapType1Code string

// May be one of APPA, CTPS
type TradingVenue2Code string

type TradingVenueIdentification1Choice struct {
	MktIdCd          MICIdentifier               `xml:"urn:iso:std:iso:20022:tech:xsd:auth.033.001.01 MktIdCd"`
	NtlCmptntAuthrty CountryCode                 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.033.001.01 NtlCmptntAuthrty"`
	Othr             TradingVenueIdentification2 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.033.001.01 Othr"`
}

type TradingVenueIdentification2 struct {
	Id Max50Text         `xml:"urn:iso:std:iso:20022:tech:xsd:auth.033.001.01 Id"`
	Tp TradingVenue2Code `xml:"urn:iso:std:iso:20022:tech:xsd:auth.033.001.01 Tp"`
}

type TransparencyDataReport10 struct {
	TechRcrdId            Max35Text                                       `xml:"urn:iso:std:iso:20022:tech:xsd:auth.033.001.01 TechRcrdId,omitempty"`
	Id                    ISINOct2015Identifier                           `xml:"urn:iso:std:iso:20022:tech:xsd:auth.033.001.01 Id"`
	FullNm                Max350Text                                      `xml:"urn:iso:std:iso:20022:tech:xsd:auth.033.001.01 FullNm,omitempty"`
	TradgVn               MICIdentifier                                   `xml:"urn:iso:std:iso:20022:tech:xsd:auth.033.001.01 TradgVn,omitempty"`
	RptgDt                ISODate                                         `xml:"urn:iso:std:iso:20022:tech:xsd:auth.033.001.01 RptgDt,omitempty"`
	MtrtyDt               ISODate                                         `xml:"urn:iso:std:iso:20022:tech:xsd:auth.033.001.01 MtrtyDt,omitempty"`
	FinInstrmClssfctn     NonEquityInstrumentReportingClassification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:auth.033.001.01 FinInstrmClssfctn"`
	UndrlygInstrmAsstClss ProductType5Code                                `xml:"urn:iso:std:iso:20022:tech:xsd:auth.033.001.01 UndrlygInstrmAsstClss,omitempty"`
	DerivCtrctTp          FinancialInstrumentContractType1Code            `xml:"urn:iso:std:iso:20022:tech:xsd:auth.033.001.01 DerivCtrctTp,omitempty"`
	Bd                    DebtInstrument5                                 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.033.001.01 Bd,omitempty"`
	EmssnAllwncTp         EmissionAllowanceProductType2Code               `xml:"urn:iso:std:iso:20022:tech:xsd:auth.033.001.01 EmssnAllwncTp,omitempty"`
	Deriv                 Derivative2Choice                               `xml:"urn:iso:std:iso:20022:tech:xsd:auth.033.001.01 Deriv,omitempty"`
}

// May be one of BOND, COMM, CURR, EMAL, EQUI, FTEQ, OPEQ, OTHR
type UnderlyingContractForDifferenceType3Code string

// May be one of BSKT
type UnderlyingEquityType3Code string

// May be one of STIX, DIVI, OTHR, VOLI
type UnderlyingEquityType4Code string

// May be one of OTHR, ETFS, SHRS, DVSE
type UnderlyingEquityType5Code string

// May be one of BSKT, DIVI, ETFS, OTHR, SHRS, DVSE, STIX, VOLI
type UnderlyingEquityType6Code string

// May be one of BOND, BNDF, INTR, IFUT
type UnderlyingInterestRateType3Code string

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