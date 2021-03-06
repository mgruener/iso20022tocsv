// Code generated by main. DO NOT EDIT.

package colr_014_001_03

import (
	"bytes"
	"encoding/xml"
	"time"
)

type ActiveCurrencyAndAmount struct {
	Value float64            `xml:",chardata"`
	Ccy   ActiveCurrencyCode `xml:"Ccy,attr"`
}

// Must match the pattern [A-Z]{3,3}
type ActiveCurrencyCode string

type Agreement2 struct {
	AgrmtDtls  Max140Text                `xml:"urn:iso:std:iso:20022:tech:xsd:colr.014.001.03 AgrmtDtls"`
	AgrmtId    Max140Text                `xml:"urn:iso:std:iso:20022:tech:xsd:colr.014.001.03 AgrmtId,omitempty"`
	AgrmtDt    ISODate                   `xml:"urn:iso:std:iso:20022:tech:xsd:colr.014.001.03 AgrmtDt"`
	BaseCcy    CurrencyCode              `xml:"urn:iso:std:iso:20022:tech:xsd:colr.014.001.03 BaseCcy"`
	AgrmtFrmwk AgreementFramework1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:colr.014.001.03 AgrmtFrmwk,omitempty"`
}

type AgreementFramework1Choice struct {
	AgrmtFrmwk AgreementFramework1Code `xml:"urn:iso:std:iso:20022:tech:xsd:colr.014.001.03 AgrmtFrmwk"`
	PrtryId    GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.014.001.03 PrtryId"`
}

// May be one of FBAA, BBAA, DERV, ISDA, NONR
type AgreementFramework1Code string

// Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type AnyBICIdentifier string

// May be one of SIMP, COMP
type CalculationMethod1Code string

type CollateralAccount1 struct {
	Id Max35Text                                  `xml:"urn:iso:std:iso:20022:tech:xsd:colr.014.001.03 Id"`
	Tp CollateralAccountIdentificationType1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:colr.014.001.03 Tp,omitempty"`
	Nm Max70Text                                  `xml:"urn:iso:std:iso:20022:tech:xsd:colr.014.001.03 Nm,omitempty"`
}

type CollateralAccountIdentificationType1Choice struct {
	Tp    CollateralAccountType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:colr.014.001.03 Tp,omitempty"`
	Prtry GenericIdentification29    `xml:"urn:iso:std:iso:20022:tech:xsd:colr.014.001.03 Prtry"`
}

// May be one of HOUS, CLIE, LIPR, MGIN, DFLT
type CollateralAccountType1Code string

type CollateralBalance1 struct {
	HeldByPtyA ActiveCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:colr.014.001.03 HeldByPtyA"`
	HeldByPtyB ActiveCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:colr.014.001.03 HeldByPtyB"`
}

type CollateralPurpose1Choice struct {
	Cd    CollateralPurpose1Code  `xml:"urn:iso:std:iso:20022:tech:xsd:colr.014.001.03 Cd"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.014.001.03 Prtry"`
}

// May be one of VAMA, SINA
type CollateralPurpose1Code string

// Must match the pattern [A-Z]{2,2}
type CountryCode string

// Must match the pattern [A-Z]{3,3}
type CurrencyCode string

type DateAndDateTimeChoice struct {
	Dt   ISODate     `xml:"urn:iso:std:iso:20022:tech:xsd:colr.014.001.03 Dt"`
	DtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:colr.014.001.03 DtTm"`
}

type DatePeriodDetails struct {
	FrDt ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:colr.014.001.03 FrDt"`
	ToDt ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:colr.014.001.03 ToDt"`
}

type Document struct {
	IntrstPmtRspn InterestPaymentResponseV03 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.014.001.03 IntrstPmtRspn"`
}

// Must match the pattern [a-zA-Z0-9]{4}
type Exact4AlphaNumericText string

// May be one of BFWD, PAYM, CCPC, COMM, CRDS, CRTL, CRSP, CCIR, CRPR, EQUI, EQPT, EQUS, EXTD, EXPT, FIXI, FORX, FORW, FUTR, OPTN, LIQU, OTCD, REPO, RVPO, SLOA, SBSC, SCRP, SLEB, SHSL, SCIR, SCIE, SWPT, TBAS, TRBD, TRCP
type ExposureType5Code string

// May be one of YEAR, MNTH, QURT, MIAN, WEEK, DAIL, ADHO, INDA
type Frequency1Code string

type GenericIdentification29 struct {
	Id      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:colr.014.001.03 Id"`
	Issr    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:colr.014.001.03 Issr"`
	SchmeNm Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:colr.014.001.03 SchmeNm,omitempty"`
}

type GenericIdentification30 struct {
	Id      Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:colr.014.001.03 Id"`
	Issr    Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:colr.014.001.03 Issr"`
	SchmeNm Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:colr.014.001.03 SchmeNm,omitempty"`
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

type InterestAmount2 struct {
	AcrdIntrstAmt  ActiveCurrencyAndAmount        `xml:"urn:iso:std:iso:20022:tech:xsd:colr.014.001.03 AcrdIntrstAmt"`
	ValDt          DateAndDateTimeChoice          `xml:"urn:iso:std:iso:20022:tech:xsd:colr.014.001.03 ValDt"`
	IntrstMtd      InterestMethod1Code            `xml:"urn:iso:std:iso:20022:tech:xsd:colr.014.001.03 IntrstMtd"`
	IntrstPrd      DatePeriodDetails              `xml:"urn:iso:std:iso:20022:tech:xsd:colr.014.001.03 IntrstPrd"`
	IntrstRate     InterestRate1Choice            `xml:"urn:iso:std:iso:20022:tech:xsd:colr.014.001.03 IntrstRate,omitempty"`
	DayCntBsis     InterestComputationMethod2Code `xml:"urn:iso:std:iso:20022:tech:xsd:colr.014.001.03 DayCntBsis,omitempty"`
	ApldWhldgTax   bool                           `xml:"urn:iso:std:iso:20022:tech:xsd:colr.014.001.03 ApldWhldgTax,omitempty"`
	ClctnMtd       CalculationMethod1Code         `xml:"urn:iso:std:iso:20022:tech:xsd:colr.014.001.03 ClctnMtd,omitempty"`
	ClctnFrqcy     Frequency1Code                 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.014.001.03 ClctnFrqcy,omitempty"`
	CollPurp       CollateralPurpose1Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:colr.014.001.03 CollPurp"`
	OpngCollBal    CollateralBalance1             `xml:"urn:iso:std:iso:20022:tech:xsd:colr.014.001.03 OpngCollBal,omitempty"`
	ClsgCollBal    CollateralBalance1             `xml:"urn:iso:std:iso:20022:tech:xsd:colr.014.001.03 ClsgCollBal"`
	StdSttlmInstrs Max140Text                     `xml:"urn:iso:std:iso:20022:tech:xsd:colr.014.001.03 StdSttlmInstrs,omitempty"`
	AddtlInf       Max210Text                     `xml:"urn:iso:std:iso:20022:tech:xsd:colr.014.001.03 AddtlInf,omitempty"`
}

// May be one of A001, A002, A003, A004, A005, A006, A007, A008, A009, A010, A011, A012, A013, A014, NARR
type InterestComputationMethod2Code string

// May be one of PHYS, ROLL
type InterestMethod1Code string

type InterestPaymentResponseV03 struct {
	TxId         Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:colr.014.001.03 TxId"`
	Oblgtn       Obligation3          `xml:"urn:iso:std:iso:20022:tech:xsd:colr.014.001.03 Oblgtn"`
	Agrmt        Agreement2           `xml:"urn:iso:std:iso:20022:tech:xsd:colr.014.001.03 Agrmt"`
	IntrstDueToA InterestAmount2      `xml:"urn:iso:std:iso:20022:tech:xsd:colr.014.001.03 IntrstDueToA,omitempty"`
	IntrstDueToB InterestAmount2      `xml:"urn:iso:std:iso:20022:tech:xsd:colr.014.001.03 IntrstDueToB,omitempty"`
	IntrstRspn   InterestResponse1    `xml:"urn:iso:std:iso:20022:tech:xsd:colr.014.001.03 IntrstRspn"`
	SplmtryData  []SupplementaryData1 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.014.001.03 SplmtryData,omitempty"`
}

type InterestRate1Choice struct {
	FxdIntrstRate   float64               `xml:"urn:iso:std:iso:20022:tech:xsd:colr.014.001.03 FxdIntrstRate"`
	VarblIntrstRate VariableInterest1Rate `xml:"urn:iso:std:iso:20022:tech:xsd:colr.014.001.03 VarblIntrstRate"`
}

// May be one of VADA, DIAM
type InterestRejectionReason1Code string

type InterestResponse1 struct {
	RspnTp         Status4Code                   `xml:"urn:iso:std:iso:20022:tech:xsd:colr.014.001.03 RspnTp"`
	RjctnRsn       RejectionReason21FormatChoice `xml:"urn:iso:std:iso:20022:tech:xsd:colr.014.001.03 RjctnRsn,omitempty"`
	RjctnRsnInf    Max140Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:colr.014.001.03 RjctnRsnInf,omitempty"`
	IntrstPmtReqId Max35Text                     `xml:"urn:iso:std:iso:20022:tech:xsd:colr.014.001.03 IntrstPmtReqId"`
}

// Must be at least 1 items long
type Max140Text string

// Must be at least 1 items long
type Max16Text string

// Must be at least 1 items long
type Max210Text string

// Must be at least 1 items long
type Max350Text string

// Must be at least 1 items long
type Max35Text string

// Must be at least 1 items long
type Max70Text string

type NameAndAddress6 struct {
	Nm  Max70Text      `xml:"urn:iso:std:iso:20022:tech:xsd:colr.014.001.03 Nm"`
	Adr PostalAddress2 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.014.001.03 Adr"`
}

type Obligation3 struct {
	PtyA       PartyIdentification33Choice `xml:"urn:iso:std:iso:20022:tech:xsd:colr.014.001.03 PtyA"`
	SvcgPtyA   PartyIdentification33Choice `xml:"urn:iso:std:iso:20022:tech:xsd:colr.014.001.03 SvcgPtyA,omitempty"`
	PtyB       PartyIdentification33Choice `xml:"urn:iso:std:iso:20022:tech:xsd:colr.014.001.03 PtyB"`
	SvcgPtyB   PartyIdentification33Choice `xml:"urn:iso:std:iso:20022:tech:xsd:colr.014.001.03 SvcgPtyB,omitempty"`
	CollAcctId CollateralAccount1          `xml:"urn:iso:std:iso:20022:tech:xsd:colr.014.001.03 CollAcctId,omitempty"`
	XpsrTp     ExposureType5Code           `xml:"urn:iso:std:iso:20022:tech:xsd:colr.014.001.03 XpsrTp,omitempty"`
	ValtnDt    DateAndDateTimeChoice       `xml:"urn:iso:std:iso:20022:tech:xsd:colr.014.001.03 ValtnDt"`
}

type PartyIdentification33Choice struct {
	AnyBIC   AnyBICIdentifier        `xml:"urn:iso:std:iso:20022:tech:xsd:colr.014.001.03 AnyBIC"`
	PrtryId  GenericIdentification29 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.014.001.03 PrtryId"`
	NmAndAdr NameAndAddress6         `xml:"urn:iso:std:iso:20022:tech:xsd:colr.014.001.03 NmAndAdr"`
}

type PostalAddress2 struct {
	StrtNm      Max70Text   `xml:"urn:iso:std:iso:20022:tech:xsd:colr.014.001.03 StrtNm,omitempty"`
	PstCdId     Max16Text   `xml:"urn:iso:std:iso:20022:tech:xsd:colr.014.001.03 PstCdId"`
	TwnNm       Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:colr.014.001.03 TwnNm"`
	CtrySubDvsn Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:colr.014.001.03 CtrySubDvsn,omitempty"`
	Ctry        CountryCode `xml:"urn:iso:std:iso:20022:tech:xsd:colr.014.001.03 Ctry"`
}

type RejectionReason21FormatChoice struct {
	Cd    InterestRejectionReason1Code `xml:"urn:iso:std:iso:20022:tech:xsd:colr.014.001.03 Cd"`
	Prtry GenericIdentification30      `xml:"urn:iso:std:iso:20022:tech:xsd:colr.014.001.03 Prtry"`
}

// May be one of REJT, PACK
type Status4Code string

type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.014.001.03 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.014.001.03 Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
}

type VariableInterest1Rate struct {
	Indx       Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:colr.014.001.03 Indx"`
	BsisPtSprd float64   `xml:"urn:iso:std:iso:20022:tech:xsd:colr.014.001.03 BsisPtSprd,omitempty"`
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
