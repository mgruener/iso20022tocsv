// Code generated by main. DO NOT EDIT.

package semt_014_002_05

import (
	"bytes"
	"encoding/xml"
	"time"
)

type AcknowledgedAcceptedStatus25Choice struct {
	NoSpcfdRsn NoReasonCode              `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.002.05 NoSpcfdRsn"`
	Rsn        []AcknowledgementReason13 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.002.05 Rsn"`
}

type AcknowledgementReason13 struct {
	Cd          AcknowledgementReason16Choice `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.002.05 Cd"`
	AddtlRsnInf RestrictedFINXMax210Text      `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.002.05 AddtlRsnInf,omitempty"`
}

type AcknowledgementReason16Choice struct {
	Cd    AcknowledgementReason5Code `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.002.05 Cd"`
	Prtry GenericIdentification47    `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.002.05 Prtry"`
}

// May be one of ADEA, SMPG, OTHR, CDCY, CDRG, CDRE, NSTP, RQWV, LATE
type AcknowledgementReason5Code string

// Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type AnyBICIdentifier string

type CancellationReason14 struct {
	Cd          CancellationReason24Choice `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.002.05 Cd"`
	AddtlRsnInf RestrictedFINXMax210Text   `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.002.05 AddtlRsnInf,omitempty"`
}

type CancellationReason24Choice struct {
	Cd    CancelledStatusReason13Code `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.002.05 Cd"`
	Prtry GenericIdentification47     `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.002.05 Prtry"`
}

type CancellationStatus17Choice struct {
	NoSpcfdRsn NoReasonCode           `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.002.05 NoSpcfdRsn"`
	Rsn        []CancellationReason14 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.002.05 Rsn"`
}

// May be one of CANI, CANS, CSUB, CXLR, CANT, CANZ, CORP, SCEX, OTHR, CTHP
type CancelledStatusReason13Code string

type DateAndDateTimeChoice struct {
	Dt   ISODate     `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.002.05 Dt"`
	DtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.002.05 DtTm"`
}

type Document struct {
	IntraPosMvmntStsAdvc IntraPositionMovementStatusAdvice002V05 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.002.05 IntraPosMvmntStsAdvc"`
}

// Must match the pattern [a-zA-Z0-9]{4}
type Exact4AlphaNumericText string

// Must be at least 1 items long
type ExternalFinancialInstrumentIdentificationType1Code string

type FailingReason10Choice struct {
	Cd    FailingReason3Code      `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.002.05 Cd"`
	Prtry GenericIdentification47 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.002.05 Prtry"`
}

// May be one of AWMO, BYIY, CLAT, ADEA, CANR, CAIS, OBJT, AWSH, PHSE, STCD, DOCY, MLAT, DOCC, BLOC, CHAS, NEWI, CLAC, MUNO, GLOB, PREA, PART, NOFX, CMON, YCOL, COLL, DEPO, FLIM, INCA, LINK, LACK, LALO, MONY, NCON, REFS, SDUT, BATC, CYCL, SBLO, CPEC, MINO, IAAD, OTHR, PHCK, BENO, BOTH, CLHT, DENO, DISA, DKNY, FROZ, LAAW, LATE, LIQU, PRCY, REGT, SETS, CERT, PRSY, INBC
type FailingReason3Code string

type FailingReason9 struct {
	Cd          FailingReason10Choice    `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.002.05 Cd"`
	AddtlRsnInf RestrictedFINXMax210Text `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.002.05 AddtlRsnInf,omitempty"`
}

type FailingStatus11Choice struct {
	NoSpcfdRsn NoReasonCode     `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.002.05 NoSpcfdRsn"`
	Rsn        []FailingReason9 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.002.05 Rsn"`
}

type FinancialInstrumentQuantity1Choice struct {
	Unit     float64 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.002.05 Unit"`
	FaceAmt  float64 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.002.05 FaceAmt"`
	AmtsdVal float64 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.002.05 AmtsdVal"`
}

type GenericIdentification39 struct {
	Id   RestrictedFINMax30Text `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.002.05 Id"`
	Issr RestrictedFINMax8Text  `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.002.05 Issr,omitempty"`
}

type GenericIdentification47 struct {
	Id      Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.002.05 Id"`
	Issr    Max4AlphaNumericText   `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.002.05 Issr"`
	SchmeNm Max4AlphaNumericText   `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.002.05 SchmeNm,omitempty"`
}

type GenericIdentification84 struct {
	Id      RestrictedFINXMax34Text `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.002.05 Id"`
	Issr    Max4AlphaNumericText    `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.002.05 Issr"`
	SchmeNm Max4AlphaNumericText    `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.002.05 SchmeNm,omitempty"`
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

type IdentificationSource4Choice struct {
	Cd    ExternalFinancialInstrumentIdentificationType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.002.05 Cd"`
	Prtry RestrictedFINExact2Text                            `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.002.05 Prtry"`
}

type IntraPositionDetails42 struct {
	PoolId        RestrictedFINXMax16Text            `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.002.05 PoolId,omitempty"`
	AcctOwnr      PartyIdentification103Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.002.05 AcctOwnr,omitempty"`
	SfkpgAcct     SecuritiesAccount30                `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.002.05 SfkpgAcct"`
	FinInstrmId   SecurityIdentification20           `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.002.05 FinInstrmId"`
	SttlmQty      FinancialInstrumentQuantity1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.002.05 SttlmQty"`
	LotNb         GenericIdentification39            `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.002.05 LotNb,omitempty"`
	SttlmDt       DateAndDateTimeChoice              `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.002.05 SttlmDt"`
	AckdStsTmStmp ISODateTime                        `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.002.05 AckdStsTmStmp,omitempty"`
	BalFr         SecuritiesBalanceType11Choice      `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.002.05 BalFr,omitempty"`
	BalTo         SecuritiesBalanceType11Choice      `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.002.05 BalTo,omitempty"`
}

type IntraPositionMovementStatusAdvice002V05 struct {
	TxId        TransactionIdentifications34         `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.002.05 TxId"`
	PrcgSts     IntraPositionProcessingStatus6Choice `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.002.05 PrcgSts,omitempty"`
	SttlmSts    SettlementStatus20Choice             `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.002.05 SttlmSts,omitempty"`
	TxDtls      IntraPositionDetails42               `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.002.05 TxDtls,omitempty"`
	SplmtryData []SupplementaryData1                 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.002.05 SplmtryData,omitempty"`
}

type IntraPositionProcessingStatus6Choice struct {
	Rjctd      RejectionOrRepairStatus33Choice    `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.002.05 Rjctd"`
	Rpr        RejectionOrRepairStatus33Choice    `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.002.05 Rpr"`
	Canc       CancellationStatus17Choice         `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.002.05 Canc"`
	AckdAccptd AcknowledgedAcceptedStatus25Choice `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.002.05 AckdAccptd"`
	Prtry      ProprietaryStatusAndReason7        `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.002.05 Prtry"`
}

// Must be at least 1 items long
type Max16Text string

// Must be at least 1 items long
type Max350Text string

// Must match the pattern [a-zA-Z0-9]{1,4}
type Max4AlphaNumericText string

// Must be at least 1 items long
type Max70Text string

// May be one of NORE
type NoReasonCode string

type OtherIdentification2 struct {
	Id  RestrictedFINXMax31Text     `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.002.05 Id"`
	Sfx Max16Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.002.05 Sfx,omitempty"`
	Tp  IdentificationSource4Choice `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.002.05 Tp"`
}

type PartyIdentification103Choice struct {
	AnyBIC  AnyBICIdentifier        `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.002.05 AnyBIC"`
	PrtryId GenericIdentification84 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.002.05 PrtryId"`
}

// May be one of AWMO, ADEA, CAIS, REFU, AWSH, PHSE, TAMM, DOCY, DOCC, BLOC, CHAS, NEWI, CLAC, MUNO, GLOB, PREA, PART, NMAS, NOFX, CMON, YCOL, COLL, DEPO, FLIM, INCA, LINK, FUTU, LACK, LALO, MONY, NCON, REFS, SDUT, BATC, CYCL, SBLO, CPEC, MINO, IAAD, OTHR, PHCK, BENO, BOTH, CLHT, DENO, DISA, DKNY, FROZ, LAAW, LATE, LIQU, PRCY, REGT, SETS, CERT, PRSY, INBC
type PendingReason10Code string

type PendingReason19 struct {
	Cd          PendingReason36Choice    `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.002.05 Cd"`
	AddtlRsnInf RestrictedFINXMax210Text `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.002.05 AddtlRsnInf,omitempty"`
}

type PendingReason36Choice struct {
	Cd    PendingReason10Code     `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.002.05 Cd"`
	Prtry GenericIdentification47 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.002.05 Prtry"`
}

type PendingStatus45Choice struct {
	NoSpcfdRsn NoReasonCode      `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.002.05 NoSpcfdRsn"`
	Rsn        []PendingReason19 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.002.05 Rsn"`
}

type ProprietaryReason5 struct {
	Rsn         GenericIdentification47  `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.002.05 Rsn,omitempty"`
	AddtlRsnInf RestrictedFINXMax210Text `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.002.05 AddtlRsnInf,omitempty"`
}

type ProprietaryStatusAndReason7 struct {
	PrtrySts GenericIdentification47 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.002.05 PrtrySts"`
	PrtryRsn []ProprietaryReason5    `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.002.05 PrtryRsn,omitempty"`
}

type RejectionAndRepairReason27Choice struct {
	Cd    RejectionReason29Code   `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.002.05 Cd"`
	Prtry GenericIdentification47 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.002.05 Prtry"`
}

type RejectionOrRepairReason27 struct {
	Cd          []RejectionAndRepairReason27Choice `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.002.05 Cd,omitempty"`
	AddtlRsnInf RestrictedFINXMax210Text           `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.002.05 AddtlRsnInf,omitempty"`
}

type RejectionOrRepairStatus33Choice struct {
	NoSpcfdRsn NoReasonCode                `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.002.05 NoSpcfdRsn"`
	Rsn        []RejectionOrRepairReason27 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.002.05 Rsn"`
}

// May be one of SAFE, ADEA, LATE, CAEV, DDAT, REFE, OTHR, DQUA, DSEC, INVB, INVL, INVN, MINO, MUNO, VALR
type RejectionReason29Code string

// Must match the pattern XX|TS
type RestrictedFINExact2Text string

// Must match the pattern ([^/]+/)+([^/]+)|([^/]*)
type RestrictedFINMax30Text string

// Must match the pattern ([^/]+/)+([^/]+)|([^/]*)
type RestrictedFINMax8Text string

// Must match the pattern [0-9a-zA-Z/\-\?:\(\)\.\n\r,'\+ ]{1,140}
type RestrictedFINXMax140Text string

// Must match the pattern ([0-9a-zA-Z\-\?:\(\)\.,'\+ ]([0-9a-zA-Z\-\?:\(\)\.,'\+ ]*(/[0-9a-zA-Z\-\?:\(\)\.,'\+ ])?)*)
type RestrictedFINXMax16Text string

// Must match the pattern [0-9a-zA-Z/\-\?:\(\)\.\n\r,'\+ ]{1,210}
type RestrictedFINXMax210Text string

// Must match the pattern [0-9a-zA-Z/\-\?:\(\)\.,'\+ ]{1,31}
type RestrictedFINXMax31Text string

// Must match the pattern ([0-9a-zA-Z\-\?:\(\)\.,'\+ ]([0-9a-zA-Z\-\?:\(\)\.,'\+ ]*(/[0-9a-zA-Z\-\?:\(\)\.,'\+ ])?)*)
type RestrictedFINXMax34Text string

// Must match the pattern [0-9a-zA-Z/\-\?:\(\)\.,'\+ ]{1,35}
type RestrictedFINXMax35Text string

type SecuritiesAccount30 struct {
	Id RestrictedFINXMax35Text `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.002.05 Id"`
	Tp GenericIdentification47 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.002.05 Tp,omitempty"`
	Nm Max70Text               `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.002.05 Nm,omitempty"`
}

type SecuritiesBalanceType11Choice struct {
	Cd    SecuritiesBalanceType13Code `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.002.05 Cd"`
	Prtry GenericIdentification47     `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.002.05 Prtry"`
}

// May be one of BLOK, AWAS, AVAI, NOMI, PLED, REGO, RSTR, OTHR, SPOS, UNRG, ISSU, QUAS, COLA
type SecuritiesBalanceType13Code string

type SecurityIdentification20 struct {
	ISIN   ISINOct2015Identifier    `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.002.05 ISIN,omitempty"`
	OthrId []OtherIdentification2   `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.002.05 OthrId,omitempty"`
	Desc   RestrictedFINXMax140Text `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.002.05 Desc,omitempty"`
}

type SettlementStatus20Choice struct {
	Pdg   PendingStatus45Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.002.05 Pdg"`
	Flng  FailingStatus11Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.002.05 Flng"`
	Prtry ProprietaryStatusAndReason7 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.002.05 Prtry"`
}

type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.002.05 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.002.05 Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
}

type TransactionIdentifications34 struct {
	AcctOwnrTxId      RestrictedFINXMax16Text `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.002.05 AcctOwnrTxId"`
	AcctSvcrTxId      RestrictedFINXMax16Text `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.002.05 AcctSvcrTxId,omitempty"`
	MktInfrstrctrTxId RestrictedFINXMax16Text `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.002.05 MktInfrstrctrTxId,omitempty"`
	PrcrTxId          RestrictedFINXMax16Text `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.002.05 PrcrTxId,omitempty"`
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
