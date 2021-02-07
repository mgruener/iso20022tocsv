// Code generated by main. DO NOT EDIT.

package catp_005_001_02

import (
	"bytes"
	"encoding/base64"
	"encoding/xml"
	"time"
)

// May be one of ABAL, ASTS, CFGT, CCNT, DISC, SNDM, RPTC
type ATMCommand4Code string

type ATMCommand7 struct {
	Tp        ATMCommand4Code             `xml:"urn:iso:std:iso:20022:tech:xsd:catp.005.001.02 Tp"`
	Urgcy     TMSContactLevel2Code        `xml:"urn:iso:std:iso:20022:tech:xsd:catp.005.001.02 Urgcy"`
	DtTm      ISODateTime                 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.005.001.02 DtTm,omitempty"`
	CmdId     ATMCommandIdentification1   `xml:"urn:iso:std:iso:20022:tech:xsd:catp.005.001.02 CmdId,omitempty"`
	CmdParams ATMCommandParameters1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:catp.005.001.02 CmdParams,omitempty"`
}

type ATMCommandIdentification1 struct {
	Orgn Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:catp.005.001.02 Orgn,omitempty"`
	Ref  Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:catp.005.001.02 Ref,omitempty"`
	Prcr Max140Text `xml:"urn:iso:std:iso:20022:tech:xsd:catp.005.001.02 Prcr,omitempty"`
}

type ATMCommandParameters1Choice struct {
	ATMReqrdGblSts  ATMStatus1Code             `xml:"urn:iso:std:iso:20022:tech:xsd:catp.005.001.02 ATMReqrdGblSts"`
	XpctdMsgFctn    MessageFunction8Code       `xml:"urn:iso:std:iso:20022:tech:xsd:catp.005.001.02 XpctdMsgFctn"`
	ReqrdCfgtnParam ATMConfigurationParameter1 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.005.001.02 ReqrdCfgtnParam"`
}

type ATMConfigurationParameter1 struct {
	Tp   DataSetCategory7Code `xml:"urn:iso:std:iso:20022:tech:xsd:catp.005.001.02 Tp"`
	Vrsn Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:catp.005.001.02 Vrsn"`
}

type ATMMessageFunction2 struct {
	Fctn     MessageFunction11Code `xml:"urn:iso:std:iso:20022:tech:xsd:catp.005.001.02 Fctn"`
	ATMSvcCd Max35Text             `xml:"urn:iso:std:iso:20022:tech:xsd:catp.005.001.02 ATMSvcCd,omitempty"`
	HstSvcCd Max35Text             `xml:"urn:iso:std:iso:20022:tech:xsd:catp.005.001.02 HstSvcCd,omitempty"`
}

type ATMReject2 struct {
	RjctInitrId Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:catp.005.001.02 RjctInitrId,omitempty"`
	RjctRsn     RejectReason1Code `xml:"urn:iso:std:iso:20022:tech:xsd:catp.005.001.02 RjctRsn"`
	AddtlInf    Max500Text        `xml:"urn:iso:std:iso:20022:tech:xsd:catp.005.001.02 AddtlInf,omitempty"`
	Cmd         []ATMCommand7     `xml:"urn:iso:std:iso:20022:tech:xsd:catp.005.001.02 Cmd,omitempty"`
	MsgInErr    Max100KBinary     `xml:"urn:iso:std:iso:20022:tech:xsd:catp.005.001.02 MsgInErr,omitempty"`
}

type ATMRejectV02 struct {
	Hdr     Header33   `xml:"urn:iso:std:iso:20022:tech:xsd:catp.005.001.02 Hdr"`
	ATMRjct ATMReject2 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.005.001.02 ATMRjct"`
}

// May be one of INSV, OUTS
type ATMStatus1Code string

// May be one of ATMC, ATMP, APPR, CRAP, CPRC, OEXR, AMNT, LOCC, MNOC
type DataSetCategory7Code string

type Document struct {
	ATMRjct ATMRejectV02 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.005.001.02 ATMRjct"`
}

type GenericIdentification77 struct {
	Id     Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:catp.005.001.02 Id"`
	Tp     PartyType12Code   `xml:"urn:iso:std:iso:20022:tech:xsd:catp.005.001.02 Tp"`
	Issr   PartyType12Code   `xml:"urn:iso:std:iso:20022:tech:xsd:catp.005.001.02 Issr,omitempty"`
	Ctry   Min2Max3AlphaText `xml:"urn:iso:std:iso:20022:tech:xsd:catp.005.001.02 Ctry,omitempty"`
	ShrtNm Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:catp.005.001.02 ShrtNm,omitempty"`
}

type Header33 struct {
	MsgFctn    ATMMessageFunction2 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.005.001.02 MsgFctn"`
	PrtcolVrsn Max6Text            `xml:"urn:iso:std:iso:20022:tech:xsd:catp.005.001.02 PrtcolVrsn"`
	XchgId     Max3NumericText     `xml:"urn:iso:std:iso:20022:tech:xsd:catp.005.001.02 XchgId,omitempty"`
	CreDtTm    ISODateTime         `xml:"urn:iso:std:iso:20022:tech:xsd:catp.005.001.02 CreDtTm"`
	InitgPty   Max35Text           `xml:"urn:iso:std:iso:20022:tech:xsd:catp.005.001.02 InitgPty,omitempty"`
	RcptPty    Max35Text           `xml:"urn:iso:std:iso:20022:tech:xsd:catp.005.001.02 RcptPty,omitempty"`
	PrcStat    Max35Text           `xml:"urn:iso:std:iso:20022:tech:xsd:catp.005.001.02 PrcStat,omitempty"`
	Tracblt    []Traceability4     `xml:"urn:iso:std:iso:20022:tech:xsd:catp.005.001.02 Tracblt,omitempty"`
}

type ISODateTime time.Time

func (t *ISODateTime) UnmarshalText(text []byte) error {
	return (*xsdDateTime)(t).UnmarshalText(text)
}
func (t ISODateTime) MarshalText() ([]byte, error) {
	return xsdDateTime(t).MarshalText()
}

type Max100KBinary []byte

func (t *Max100KBinary) UnmarshalText(text []byte) error {
	return (*xsdBase64Binary)(t).UnmarshalText(text)
}
func (t Max100KBinary) MarshalText() ([]byte, error) {
	return xsdBase64Binary(t).MarshalText()
}

// Must be at least 1 items long
type Max140Text string

// Must be at least 1 items long
type Max35Text string

// Must match the pattern [0-9]{1,3}
type Max3NumericText string

// Must be at least 1 items long
type Max500Text string

// Must be at least 1 items long
type Max6Text string

// May be one of BALN, CMPA, CMPD, ACMD, DVCC, DIAQ, DIAP, GSTS, INQQ, INQP, KYAQ, KYAP, PINQ, PINP, RJAQ, RJAP, WITV, WITK, WITQ, WITP, INQC, H2AP, H2AQ, TMOP, CSEC, DSEC, SKSC, SSTS, DPSK, DPSV, DPSQ, DPSP, EXPK, EXPV, TRFQ, TRFP, RPTC
type MessageFunction11Code string

// May be one of BALN, GSTS, DSEC, INQC, KEYQ, SSTS
type MessageFunction8Code string

// Must match the pattern [a-zA-Z]{2,3}
type Min2Max3AlphaText string

// May be one of ACQR, ATMG, CISP, DLIS, HSTG, ITAG, OATM
type PartyType12Code string

// May be one of UNPR, IMSG, PARS, SECU, INTP, RCPP, DPMG, VERS, MSGT
type RejectReason1Code string

// May be one of ASAP, CRIT, DTIM, ENCS
type TMSContactLevel2Code string

type Traceability4 struct {
	RlayId      GenericIdentification77 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.005.001.02 RlayId"`
	SeqNb       Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:catp.005.001.02 SeqNb,omitempty"`
	TracDtTmIn  ISODateTime             `xml:"urn:iso:std:iso:20022:tech:xsd:catp.005.001.02 TracDtTmIn"`
	TracDtTmOut ISODateTime             `xml:"urn:iso:std:iso:20022:tech:xsd:catp.005.001.02 TracDtTmOut"`
}

type xsdBase64Binary []byte

func (b *xsdBase64Binary) UnmarshalText(text []byte) (err error) {
	*b, err = base64.StdEncoding.DecodeString(string(text))
	return
}
func (b xsdBase64Binary) MarshalText() ([]byte, error) {
	var buf bytes.Buffer
	enc := base64.NewEncoder(base64.StdEncoding, &buf)
	enc.Write([]byte(b))
	enc.Close()
	return buf.Bytes(), nil
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
