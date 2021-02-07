// Code generated by main. DO NOT EDIT.

package auth_029_001_02

import (
	"bytes"
	"encoding/xml"
	"time"
)

// May be one of ADDR, PBOX, HOME, BIZZ, MLTO, DLVY
type AddressType2Code string

// Must match the pattern [A-Z0-9]{4,4}[A-Z]{2,2}[A-Z0-9]{2,2}([A-Z0-9]{3,3}){0,1}
type AnyBICDec2014Identifier string

// May be one of ANYM
type AnyMIC1Code string

// Must match the pattern [A-Z]{6,6}
type CFIOct2015Identifier string

type CorporateSectorCriteria3 struct {
	FISctr  []FinancialPartySectorType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:auth.029.001.02 FISctr,omitempty"`
	NFISctr []NonFinancialPartySector1Code  `xml:"urn:iso:std:iso:20022:tech:xsd:auth.029.001.02 NFISctr,omitempty"`
	NotRptd NotReported1Code                `xml:"urn:iso:std:iso:20022:tech:xsd:auth.029.001.02 NotRptd,omitempty"`
}

// Must match the pattern [A-Z]{2,2}
type CountryCode string

type DateOrBlankQuery2Choice struct {
	Rg      DatePeriod1      `xml:"urn:iso:std:iso:20022:tech:xsd:auth.029.001.02 Rg"`
	NotRptd NotReported1Code `xml:"urn:iso:std:iso:20022:tech:xsd:auth.029.001.02 NotRptd"`
}

type DatePeriod1 struct {
	FrDt ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:auth.029.001.02 FrDt,omitempty"`
	ToDt ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:auth.029.001.02 ToDt"`
}

type DateTimePeriod1 struct {
	FrDtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:auth.029.001.02 FrDtTm"`
	ToDtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:auth.029.001.02 ToDtTm"`
}

type DerivativesTradeReportQueryV02 struct {
	RqstngAuthrty PartyIdentification121Choice `xml:"urn:iso:std:iso:20022:tech:xsd:auth.029.001.02 RqstngAuthrty"`
	TradQryData   TradeReportQuery10Choice     `xml:"urn:iso:std:iso:20022:tech:xsd:auth.029.001.02 TradQryData"`
	SplmtryData   []SupplementaryData1         `xml:"urn:iso:std:iso:20022:tech:xsd:auth.029.001.02 SplmtryData,omitempty"`
}

type Document struct {
	DerivsTradRptQry DerivativesTradeReportQueryV02 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.029.001.02 DerivsTradRptQry"`
}

// May be one of AIFD, ASSU, CDTI, INUN, INVF, ORPI, REIN, UCIT, OTHR
type FinancialPartySectorType1Code string

// May be one of DAIL, WEEK, MNTH, ADHO
type Frequency14Code string

type GenericIdentification1 struct {
	Id      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:auth.029.001.02 Id"`
	SchmeNm Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:auth.029.001.02 SchmeNm,omitempty"`
	Issr    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:auth.029.001.02 Issr,omitempty"`
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

// Must match the pattern [A-Z0-9]{18,18}[0-9]{2,2}
type LEIIdentifier string

// Must match the pattern [A-Z0-9]{4,4}
type MICIdentifier string

// Must be at least 1 items long
type Max1000Text string

// Must be at least 1 items long
type Max16Text string

// Must be at least 1 items long
type Max25Text string

// Must be at least 1 items long
type Max350Text string

// Must be at least 1 items long
type Max35Text string

// Must be at least 1 items long
type Max50Text string

// Must be at least 1 items long
type Max52Text string

// Must be at least 1 items long
type Max70Text string

type NameAndAddress5 struct {
	Nm  Max350Text     `xml:"urn:iso:std:iso:20022:tech:xsd:auth.029.001.02 Nm"`
	Adr PostalAddress1 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.029.001.02 Adr,omitempty"`
}

// May be one of WTER, MING, MAFG, SPLY, CSTR, AGRI, ACAF, EDUC, AEAR, FINA, HHSW, INCO, WRRM, OTSA, PSTA, PADS, RESA, TRAS, ASSA, AHAE, AEOB
type NonFinancialPartySector1Code string

// May be one of NTAV
type NotAvailable1Code string

// May be one of NORP
type NotReported1Code string

// May be one of ANDD, ORRR
type Operation3Code string

type PartyIdentification121Choice struct {
	AnyBIC     AnyBICDec2014Identifier `xml:"urn:iso:std:iso:20022:tech:xsd:auth.029.001.02 AnyBIC"`
	LglNttyIdr LEIIdentifier           `xml:"urn:iso:std:iso:20022:tech:xsd:auth.029.001.02 LglNttyIdr"`
	NmAndAdr   NameAndAddress5         `xml:"urn:iso:std:iso:20022:tech:xsd:auth.029.001.02 NmAndAdr"`
	PrtryId    GenericIdentification1  `xml:"urn:iso:std:iso:20022:tech:xsd:auth.029.001.02 PrtryId"`
}

// May be one of OTHR, NFIN, FIIN, CCPS
type PartyNatureType1Code string

type PostalAddress1 struct {
	AdrTp       AddressType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:auth.029.001.02 AdrTp,omitempty"`
	AdrLine     []Max70Text      `xml:"urn:iso:std:iso:20022:tech:xsd:auth.029.001.02 AdrLine,omitempty"`
	StrtNm      Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:auth.029.001.02 StrtNm,omitempty"`
	BldgNb      Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:auth.029.001.02 BldgNb,omitempty"`
	PstCd       Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:auth.029.001.02 PstCd,omitempty"`
	TwnNm       Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:auth.029.001.02 TwnNm,omitempty"`
	CtrySubDvsn Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:auth.029.001.02 CtrySubDvsn,omitempty"`
	Ctry        CountryCode      `xml:"urn:iso:std:iso:20022:tech:xsd:auth.029.001.02 Ctry"`
}

type ProductClassificationCriteria1 struct {
	ClssfctnFinInstrm []CFIOct2015Identifier `xml:"urn:iso:std:iso:20022:tech:xsd:auth.029.001.02 ClssfctnFinInstrm,omitempty"`
	UnqPdctIdr        []Max52Text            `xml:"urn:iso:std:iso:20022:tech:xsd:auth.029.001.02 UnqPdctIdr,omitempty"`
}

// May be one of CRDT, CURR, EQUI, INTR, COMM, OTHR
type ProductType4Code string

type SecuritiesTradeVenueCriteria1Choice struct {
	MIC    []MICIdentifier `xml:"urn:iso:std:iso:20022:tech:xsd:auth.029.001.02 MIC"`
	AnyMIC AnyMIC1Code     `xml:"urn:iso:std:iso:20022:tech:xsd:auth.029.001.02 AnyMIC"`
}

type SecurityIdentification20Choice struct {
	ISIN ISINOct2015Identifier `xml:"urn:iso:std:iso:20022:tech:xsd:auth.029.001.02 ISIN"`
	Nm   Max25Text             `xml:"urn:iso:std:iso:20022:tech:xsd:auth.029.001.02 Nm"`
}

type SecurityIdentificationQuery3Choice struct {
	ISIN            []ISINOct2015Identifier          `xml:"urn:iso:std:iso:20022:tech:xsd:auth.029.001.02 ISIN,omitempty"`
	AltrntvInstrmId []Max52Text                      `xml:"urn:iso:std:iso:20022:tech:xsd:auth.029.001.02 AltrntvInstrmId,omitempty"`
	NotAvlbl        NotAvailable1Code                `xml:"urn:iso:std:iso:20022:tech:xsd:auth.029.001.02 NotAvlbl,omitempty"`
	UnqPdctIdr      []Max52Text                      `xml:"urn:iso:std:iso:20022:tech:xsd:auth.029.001.02 UnqPdctIdr,omitempty"`
	Indx            []SecurityIdentification20Choice `xml:"urn:iso:std:iso:20022:tech:xsd:auth.029.001.02 Indx,omitempty"`
	NotRptd         NotReported1Code                 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.029.001.02 NotRptd,omitempty"`
}

type SecurityIdentificationQueryCriteria1 struct {
	ISIN            []ISINOct2015Identifier `xml:"urn:iso:std:iso:20022:tech:xsd:auth.029.001.02 ISIN,omitempty"`
	AltrntvInstrmId []Max52Text             `xml:"urn:iso:std:iso:20022:tech:xsd:auth.029.001.02 AltrntvInstrmId,omitempty"`
}

type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.029.001.02 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.029.001.02 Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
}

type TradeAdditionalQueryCriteria3 struct {
	ActnTp       []TransactionOperationType3Code     `xml:"urn:iso:std:iso:20022:tech:xsd:auth.029.001.02 ActnTp,omitempty"`
	ExctnVn      SecuritiesTradeVenueCriteria1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:auth.029.001.02 ExctnVn,omitempty"`
	NtrOfCtrPty  PartyNatureType1Code                `xml:"urn:iso:std:iso:20022:tech:xsd:auth.029.001.02 NtrOfCtrPty,omitempty"`
	CorpSctr     CorporateSectorCriteria3            `xml:"urn:iso:std:iso:20022:tech:xsd:auth.029.001.02 CorpSctr,omitempty"`
	AsstClss     []ProductType4Code                  `xml:"urn:iso:std:iso:20022:tech:xsd:auth.029.001.02 AsstClss,omitempty"`
	PdctClssfctn ProductClassificationCriteria1      `xml:"urn:iso:std:iso:20022:tech:xsd:auth.029.001.02 PdctClssfctn,omitempty"`
}

type TradeDateTimeQueryCriteria2 struct {
	RptgDtTm  DateTimePeriod1         `xml:"urn:iso:std:iso:20022:tech:xsd:auth.029.001.02 RptgDtTm,omitempty"`
	ExctnDtTm DateTimePeriod1         `xml:"urn:iso:std:iso:20022:tech:xsd:auth.029.001.02 ExctnDtTm,omitempty"`
	MtrtyDt   DateOrBlankQuery2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:auth.029.001.02 MtrtyDt,omitempty"`
	TermntnDt DateOrBlankQuery2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:auth.029.001.02 TermntnDt,omitempty"`
}

type TradePartyIdentificationQuery8 struct {
	LEI     []LEIIdentifier           `xml:"urn:iso:std:iso:20022:tech:xsd:auth.029.001.02 LEI,omitempty"`
	AnyBIC  []AnyBICDec2014Identifier `xml:"urn:iso:std:iso:20022:tech:xsd:auth.029.001.02 AnyBIC,omitempty"`
	ClntId  []Max50Text               `xml:"urn:iso:std:iso:20022:tech:xsd:auth.029.001.02 ClntId,omitempty"`
	NotRptd NotReported1Code          `xml:"urn:iso:std:iso:20022:tech:xsd:auth.029.001.02 NotRptd,omitempty"`
}

type TradePartyQueryCriteria3 struct {
	Oprtr      Operation3Code                 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.029.001.02 Oprtr"`
	RptgCtrPty TradePartyIdentificationQuery8 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.029.001.02 RptgCtrPty,omitempty"`
	OthrCtrPty TradePartyIdentificationQuery8 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.029.001.02 OthrCtrPty,omitempty"`
	Bnfcry     TradePartyIdentificationQuery8 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.029.001.02 Bnfcry,omitempty"`
	SubmitgAgt TradePartyIdentificationQuery8 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.029.001.02 SubmitgAgt,omitempty"`
	Brkr       TradePartyIdentificationQuery8 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.029.001.02 Brkr,omitempty"`
	CCP        TradePartyIdentificationQuery8 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.029.001.02 CCP,omitempty"`
}

type TradeQueryCriteria4 struct {
	TradLifeCyclHstry bool                                      `xml:"urn:iso:std:iso:20022:tech:xsd:auth.029.001.02 TradLifeCyclHstry"`
	OutsdngTradInd    bool                                      `xml:"urn:iso:std:iso:20022:tech:xsd:auth.029.001.02 OutsdngTradInd"`
	TradPtyCrit       TradePartyQueryCriteria3                  `xml:"urn:iso:std:iso:20022:tech:xsd:auth.029.001.02 TradPtyCrit,omitempty"`
	FinInstrmCrit     TradeSecurityIdentificationQueryCriteria2 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.029.001.02 FinInstrmCrit,omitempty"`
	TmCrit            TradeDateTimeQueryCriteria2               `xml:"urn:iso:std:iso:20022:tech:xsd:auth.029.001.02 TmCrit,omitempty"`
	OthrCrit          TradeAdditionalQueryCriteria3             `xml:"urn:iso:std:iso:20022:tech:xsd:auth.029.001.02 OthrCrit,omitempty"`
}

type TradeQueryExecutionFrequency3 struct {
	FrqcyTp   Frequency14Code `xml:"urn:iso:std:iso:20022:tech:xsd:auth.029.001.02 FrqcyTp"`
	DlvryDay  []WeekDay3Code  `xml:"urn:iso:std:iso:20022:tech:xsd:auth.029.001.02 DlvryDay,omitempty"`
	DayOfMnth []float64       `xml:"urn:iso:std:iso:20022:tech:xsd:auth.029.001.02 DayOfMnth,omitempty"`
}

type TradeRecurrentQuery5 struct {
	QryTp    Max1000Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:auth.029.001.02 QryTp"`
	Frqcy    TradeQueryExecutionFrequency3 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.029.001.02 Frqcy"`
	VldUntil ISODate                       `xml:"urn:iso:std:iso:20022:tech:xsd:auth.029.001.02 VldUntil"`
}

type TradeReportQuery10Choice struct {
	AdHocQry TradeQueryCriteria4  `xml:"urn:iso:std:iso:20022:tech:xsd:auth.029.001.02 AdHocQry"`
	RcrntQry TradeRecurrentQuery5 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.029.001.02 RcrntQry"`
}

type TradeSecurityIdentificationQueryCriteria2 struct {
	Oprtr           Operation3Code                         `xml:"urn:iso:std:iso:20022:tech:xsd:auth.029.001.02 Oprtr"`
	Id              []SecurityIdentificationQueryCriteria1 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.029.001.02 Id,omitempty"`
	UndrlygInstrmId []SecurityIdentificationQuery3Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:auth.029.001.02 UndrlygInstrmId,omitempty"`
}

// May be one of CORR, ETRM, EROR, NEWT, POSC, VALU, COMP, MODI, OTHR
type TransactionOperationType3Code string

// May be one of ALLD, XBHL, IBHL, FRID, MOND, SATD, SUND, THUD, TUED, WEDD, WDAY, WEND
type WeekDay3Code string

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
