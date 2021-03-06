// Code generated by main. DO NOT EDIT.

package semt_024_001_01

import (
	"bytes"
	"encoding/xml"
	"time"
)

type AccountIdentification5 struct {
	Id Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:semt.024.001.01 Id"`
	Nm Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:semt.024.001.01 Nm,omitempty"`
	Tp GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.024.001.01 Tp,omitempty"`
}

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

type AmountAndDirection30 struct {
	Amt ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:semt.024.001.01 Amt"`
	Sgn bool                              `xml:"urn:iso:std:iso:20022:tech:xsd:semt.024.001.01 Sgn,omitempty"`
}

type AmountAndDirection31 struct {
	Amt        ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:semt.024.001.01 Amt"`
	ShrtLngInd ShortLong1Code                    `xml:"urn:iso:std:iso:20022:tech:xsd:semt.024.001.01 ShrtLngInd,omitempty"`
}

type AmountAndRate2 struct {
	Amt  AmountAndDirection30 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.024.001.01 Amt,omitempty"`
	Rate float64              `xml:"urn:iso:std:iso:20022:tech:xsd:semt.024.001.01 Rate,omitempty"`
}

type BalanceDetails5 struct {
	Tp      BalanceType6Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:semt.024.001.01 Tp"`
	Urlsd   Unrealised1Code      `xml:"urn:iso:std:iso:20022:tech:xsd:semt.024.001.01 Urlsd,omitempty"`
	Amt     AmountAndDirection31 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.024.001.01 Amt"`
	DtldBal []BalanceDetails6    `xml:"urn:iso:std:iso:20022:tech:xsd:semt.024.001.01 DtldBal,omitempty"`
}

type BalanceDetails6 struct {
	Ctgy  FinancialAssetTypeCategory1Code `xml:"urn:iso:std:iso:20022:tech:xsd:semt.024.001.01 Ctgy,omitempty"`
	Tp    BalanceType7Choice              `xml:"urn:iso:std:iso:20022:tech:xsd:semt.024.001.01 Tp,omitempty"`
	Urlsd Unrealised1Code                 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.024.001.01 Urlsd,omitempty"`
	Amt   AmountAndDirection31            `xml:"urn:iso:std:iso:20022:tech:xsd:semt.024.001.01 Amt"`
}

// May be one of INVE, CASE, BORR, REVE, EXPN, IIOF, OTHR, PAYA, RECE
type BalanceType13Code string

type BalanceType6Choice struct {
	Cd    BalanceType13Code       `xml:"urn:iso:std:iso:20022:tech:xsd:semt.024.001.01 Cd"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.024.001.01 Prtry"`
}

type BalanceType7Choice struct {
	Cd    FinancialAssetBalanceType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:semt.024.001.01 Cd"`
	Prtry GenericIdentification30        `xml:"urn:iso:std:iso:20022:tech:xsd:semt.024.001.01 Prtry"`
	Acct  AccountIdentification5         `xml:"urn:iso:std:iso:20022:tech:xsd:semt.024.001.01 Acct"`
}

type DateAndDateTimeChoice struct {
	Dt   ISODate     `xml:"urn:iso:std:iso:20022:tech:xsd:semt.024.001.01 Dt"`
	DtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:semt.024.001.01 DtTm"`
}

type DatePeriodDetails struct {
	FrDt ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:semt.024.001.01 FrDt"`
	ToDt ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:semt.024.001.01 ToDt"`
}

type Document struct {
	TtlPrtflValtnRpt TotalPortfolioValuationReportV01 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.024.001.01 TtlPrtflValtnRpt"`
}

// May be one of YEAR, SEMI, QUTR, TOMN, MNTH, TWMN, TOWK, WEEK, DAIL, ADHO, INDA, OVNG, ONDE
type EventFrequency1Code string

// Must match the pattern [a-zA-Z0-9]{4}
type Exact4AlphaNumericText string

// Must be at least 1 items long
type ExternalFinancialInstrumentIdentificationType1Code string

// May be one of ACRU, OINT, SCAS, FXTR, CASH, TIPS, EQUI, CSTK, PREF, MFUN, XFUN, RGHT, WARR, BOND, CONV, CBND, GBND, OPTN, FUTR, SWAP, CUEX, FOIV, GOLD, PROP, BAAP, SYBL, CBOO, CEOD, CDEO, CLOB, CMOO, COPR, CPPE, DISC, FEAD, FEHA, FEHL, FNMA, FLNO, GNMA, TAAB, IETM, MPRP, MBON, SLMA, STIF, TSTP, TIDE, UNBW, UNBO, VRDN, ZOOO, FWBO, FRAG, REPO, XREP, TREP, RXRP, FXFD, FXSP
type FinancialAssetBalanceType1Code string

// May be one of EQTY, DEBT, ENTL, DERV, MMKT, OTHR
type FinancialAssetTypeCategory1Code string

type Frequency8Choice struct {
	Cd    EventFrequency1Code     `xml:"urn:iso:std:iso:20022:tech:xsd:semt.024.001.01 Cd"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.024.001.01 Prtry"`
}

type GenericIdentification29 struct {
	Id      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:semt.024.001.01 Id"`
	Issr    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:semt.024.001.01 Issr"`
	SchmeNm Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:semt.024.001.01 SchmeNm,omitempty"`
}

type GenericIdentification30 struct {
	Id      Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:semt.024.001.01 Id"`
	Issr    Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:semt.024.001.01 Issr"`
	SchmeNm Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:semt.024.001.01 SchmeNm,omitempty"`
}

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
	Cd    ExternalFinancialInstrumentIdentificationType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:semt.024.001.01 Cd"`
	Prtry Max35Text                                          `xml:"urn:iso:std:iso:20022:tech:xsd:semt.024.001.01 Prtry"`
}

type InvestmentFund1 struct {
	FinInstrmId     SecurityIdentification14 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.024.001.01 FinInstrmId,omitempty"`
	ClssTp          Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:semt.024.001.01 ClssTp,omitempty"`
	TtlUnitsOutsdng float64                  `xml:"urn:iso:std:iso:20022:tech:xsd:semt.024.001.01 TtlUnitsOutsdng,omitempty"`
	TxnlUnits       float64                  `xml:"urn:iso:std:iso:20022:tech:xsd:semt.024.001.01 TxnlUnits,omitempty"`
	TtlVal          AmountAndDirection30     `xml:"urn:iso:std:iso:20022:tech:xsd:semt.024.001.01 TtlVal,omitempty"`
	Pric            []PriceInformation10     `xml:"urn:iso:std:iso:20022:tech:xsd:semt.024.001.01 Pric,omitempty"`
	SplmtryData     []SupplementaryData1     `xml:"urn:iso:std:iso:20022:tech:xsd:semt.024.001.01 SplmtryData,omitempty"`
}

// Must be at least 1 items long
type Max140Text string

// Must be at least 1 items long
type Max16Text string

// Must be at least 1 items long
type Max350Text string

// Must be at least 1 items long
type Max35Text string

// Must match the pattern [0-9]{1,5}
type Max5NumericText string

type OtherIdentification1 struct {
	Id  Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:semt.024.001.01 Id"`
	Sfx Max16Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:semt.024.001.01 Sfx,omitempty"`
	Tp  IdentificationSource3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:semt.024.001.01 Tp"`
}

type Pagination struct {
	PgNb      Max5NumericText `xml:"urn:iso:std:iso:20022:tech:xsd:semt.024.001.01 PgNb"`
	LastPgInd bool            `xml:"urn:iso:std:iso:20022:tech:xsd:semt.024.001.01 LastPgInd"`
}

type PortfolioBalance1 struct {
	SummryBal []BalanceDetails5 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.024.001.01 SummryBal"`
	DtldBal   []BalanceDetails6 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.024.001.01 DtldBal"`
}

type PriceAndDirection1 struct {
	Val ActiveOrHistoricCurrencyAnd13DecimalAmount `xml:"urn:iso:std:iso:20022:tech:xsd:semt.024.001.01 Val"`
	Sgn bool                                       `xml:"urn:iso:std:iso:20022:tech:xsd:semt.024.001.01 Sgn,omitempty"`
}

type PriceInformation10 struct {
	CurPric   ActiveOrHistoricCurrencyAnd13DecimalAmount `xml:"urn:iso:std:iso:20022:tech:xsd:semt.024.001.01 CurPric"`
	Tp        TypeOfPrice27Choice                        `xml:"urn:iso:std:iso:20022:tech:xsd:semt.024.001.01 Tp"`
	PrvsPric  ActiveOrHistoricCurrencyAnd13DecimalAmount `xml:"urn:iso:std:iso:20022:tech:xsd:semt.024.001.01 PrvsPric,omitempty"`
	AmtOfChng PriceValueAndRate4                         `xml:"urn:iso:std:iso:20022:tech:xsd:semt.024.001.01 AmtOfChng,omitempty"`
}

type PriceValueAndRate4 struct {
	Val  PriceAndDirection1 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.024.001.01 Val,omitempty"`
	Rate float64            `xml:"urn:iso:std:iso:20022:tech:xsd:semt.024.001.01 Rate,omitempty"`
}

type Report4 struct {
	RptNb       Max5NumericText        `xml:"urn:iso:std:iso:20022:tech:xsd:semt.024.001.01 RptNb,omitempty"`
	QryRef      Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:semt.024.001.01 QryRef,omitempty"`
	RptId       Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:semt.024.001.01 RptId,omitempty"`
	RptDtTm     DateAndDateTimeChoice  `xml:"urn:iso:std:iso:20022:tech:xsd:semt.024.001.01 RptDtTm"`
	CreDtTm     DateAndDateTimeChoice  `xml:"urn:iso:std:iso:20022:tech:xsd:semt.024.001.01 CreDtTm,omitempty"`
	PrvsRptDtTm DateAndDateTimeChoice  `xml:"urn:iso:std:iso:20022:tech:xsd:semt.024.001.01 PrvsRptDtTm,omitempty"`
	Frqcy       Frequency8Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:semt.024.001.01 Frqcy"`
	UpdTp       UpdateType4Choice      `xml:"urn:iso:std:iso:20022:tech:xsd:semt.024.001.01 UpdTp"`
	RptBsis     StatementBasis6Choice  `xml:"urn:iso:std:iso:20022:tech:xsd:semt.024.001.01 RptBsis"`
	RptPrd      DatePeriodDetails      `xml:"urn:iso:std:iso:20022:tech:xsd:semt.024.001.01 RptPrd,omitempty"`
	RptSrc      StatementSource1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:semt.024.001.01 RptSrc,omitempty"`
	AudtdInd    bool                   `xml:"urn:iso:std:iso:20022:tech:xsd:semt.024.001.01 AudtdInd,omitempty"`
	ActvtyInd   bool                   `xml:"urn:iso:std:iso:20022:tech:xsd:semt.024.001.01 ActvtyInd,omitempty"`
}

type SecuritiesAccount21 struct {
	Acct    AccountIdentification5       `xml:"urn:iso:std:iso:20022:tech:xsd:semt.024.001.01 Acct"`
	SubAcct AccountIdentification5       `xml:"urn:iso:std:iso:20022:tech:xsd:semt.024.001.01 SubAcct,omitempty"`
	BaseCcy ActiveOrHistoricCurrencyCode `xml:"urn:iso:std:iso:20022:tech:xsd:semt.024.001.01 BaseCcy,omitempty"`
	RptgCcy ActiveOrHistoricCurrencyCode `xml:"urn:iso:std:iso:20022:tech:xsd:semt.024.001.01 RptgCcy,omitempty"`
	FXRate  float64                      `xml:"urn:iso:std:iso:20022:tech:xsd:semt.024.001.01 FXRate,omitempty"`
}

type SecurityIdentification14 struct {
	ISIN   ISINIdentifier         `xml:"urn:iso:std:iso:20022:tech:xsd:semt.024.001.01 ISIN,omitempty"`
	OthrId []OtherIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.024.001.01 OthrId,omitempty"`
	Desc   Max140Text             `xml:"urn:iso:std:iso:20022:tech:xsd:semt.024.001.01 Desc,omitempty"`
}

// May be one of SHOR, LONG
type ShortLong1Code string

// May be one of CONT, SETT, TRAD
type StatementBasis1Code string

type StatementBasis6Choice struct {
	Cd    StatementBasis1Code     `xml:"urn:iso:std:iso:20022:tech:xsd:semt.024.001.01 Cd"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.024.001.01 Prtry"`
}

type StatementSource1Choice struct {
	Cd    StatementSource1Code    `xml:"urn:iso:std:iso:20022:tech:xsd:semt.024.001.01 Cd"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.024.001.01 Prtry"`
}

// May be one of ACCT, CUST
type StatementSource1Code string

// May be one of COMP, DELT
type StatementUpdateType1Code string

type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.024.001.01 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.024.001.01 Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
}

type TotalPortfolioValuation1 struct {
	TtlPrtflVal     AmountAndDirection30 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.024.001.01 TtlPrtflVal"`
	PrvsTtlPrtflVal AmountAndDirection30 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.024.001.01 PrvsTtlPrtflVal,omitempty"`
	TtlPrtflValChng AmountAndRate2       `xml:"urn:iso:std:iso:20022:tech:xsd:semt.024.001.01 TtlPrtflValChng,omitempty"`
	TtlBookVal      AmountAndDirection30 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.024.001.01 TtlBookVal"`
	PrvsTtlBookVal  AmountAndDirection30 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.024.001.01 PrvsTtlBookVal,omitempty"`
	TtlBookValChng  AmountAndRate2       `xml:"urn:iso:std:iso:20022:tech:xsd:semt.024.001.01 TtlBookValChng,omitempty"`
	TtlRcts         AmountAndDirection30 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.024.001.01 TtlRcts,omitempty"`
	TtlDsbrsmnts    AmountAndDirection30 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.024.001.01 TtlDsbrsmnts,omitempty"`
	IncmRcvd        AmountAndDirection30 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.024.001.01 IncmRcvd,omitempty"`
	ExpnssPd        AmountAndDirection30 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.024.001.01 ExpnssPd,omitempty"`
	UrlsdGnOrLoss   AmountAndDirection31 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.024.001.01 UrlsdGnOrLoss,omitempty"`
	RealsdGnOrLoss  AmountAndDirection31 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.024.001.01 RealsdGnOrLoss,omitempty"`
	AcrdIncm        AmountAndDirection30 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.024.001.01 AcrdIncm,omitempty"`
	InvstmtFndDtls  []InvestmentFund1    `xml:"urn:iso:std:iso:20022:tech:xsd:semt.024.001.01 InvstmtFndDtls,omitempty"`
}

type TotalPortfolioValuationReportV01 struct {
	Pgntn         Pagination               `xml:"urn:iso:std:iso:20022:tech:xsd:semt.024.001.01 Pgntn"`
	RptGnlDtls    Report4                  `xml:"urn:iso:std:iso:20022:tech:xsd:semt.024.001.01 RptGnlDtls"`
	AcctDtls      SecuritiesAccount21      `xml:"urn:iso:std:iso:20022:tech:xsd:semt.024.001.01 AcctDtls"`
	TtlPrtflValtn TotalPortfolioValuation1 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.024.001.01 TtlPrtflValtn"`
	Bal           PortfolioBalance1        `xml:"urn:iso:std:iso:20022:tech:xsd:semt.024.001.01 Bal,omitempty"`
	SplmtryData   SupplementaryData1       `xml:"urn:iso:std:iso:20022:tech:xsd:semt.024.001.01 SplmtryData,omitempty"`
}

type TypeOfPrice27Choice struct {
	Cd    TypeOfPrice30Code       `xml:"urn:iso:std:iso:20022:tech:xsd:semt.024.001.01 Cd"`
	Prtry GenericIdentification29 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.024.001.01 Prtry"`
}

// May be one of BIDE, OFFR, NAVL, CREA, CANC, INTE, SWNG, MIDD, RINV, SWIC, DDVR, ACTU, NAUP, GUAR, ENAV, REDN, SUBN
type TypeOfPrice30Code string

// May be one of GAIN, LOSS
type Unrealised1Code string

type UpdateType4Choice struct {
	Cd    StatementUpdateType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:semt.024.001.01 Cd"`
	Prtry GenericIdentification30  `xml:"urn:iso:std:iso:20022:tech:xsd:semt.024.001.01 Prtry"`
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
