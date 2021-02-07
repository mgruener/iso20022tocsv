// Code generated by main. DO NOT EDIT.

package caaa_015_001_05

import (
	"bytes"
	"encoding/base64"
	"encoding/xml"
	"time"
)

type AcceptorRejection2 struct {
	RjctRsn  RejectReason1Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.015.001.05 RjctRsn"`
	AddtlInf Max500Text        `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.015.001.05 AddtlInf,omitempty"`
	MsgInErr Max100KBinary     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.015.001.05 MsgInErr,omitempty"`
}

type AcceptorRejectionV05 struct {
	Hdr  Header26           `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.015.001.05 Hdr"`
	Rjct AcceptorRejection2 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.015.001.05 Rjct"`
}

type Document struct {
	AccptrRjctn AcceptorRejectionV05 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.015.001.05 AccptrRjctn"`
}

type GenericIdentification53 struct {
	Id     Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.015.001.05 Id"`
	Tp     PartyType3Code    `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.015.001.05 Tp,omitempty"`
	Issr   PartyType4Code    `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.015.001.05 Issr,omitempty"`
	Ctry   Min2Max3AlphaText `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.015.001.05 Ctry,omitempty"`
	ShrtNm Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.015.001.05 ShrtNm,omitempty"`
}

type GenericIdentification76 struct {
	Id     Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.015.001.05 Id"`
	Tp     PartyType3Code    `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.015.001.05 Tp"`
	Issr   PartyType4Code    `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.015.001.05 Issr,omitempty"`
	Ctry   Min2Max3AlphaText `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.015.001.05 Ctry,omitempty"`
	ShrtNm Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.015.001.05 ShrtNm,omitempty"`
}

type GenericIdentification94 struct {
	Id       Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.015.001.05 Id"`
	Tp       PartyType3Code     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.015.001.05 Tp,omitempty"`
	Issr     PartyType4Code     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.015.001.05 Issr,omitempty"`
	Ctry     Min2Max3AlphaText  `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.015.001.05 Ctry,omitempty"`
	ShrtNm   Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.015.001.05 ShrtNm,omitempty"`
	RmotAccs NetworkParameters5 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.015.001.05 RmotAccs,omitempty"`
}

type Header26 struct {
	MsgFctn    MessageFunction9Code    `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.015.001.05 MsgFctn"`
	PrtcolVrsn Max6Text                `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.015.001.05 PrtcolVrsn"`
	XchgId     float64                 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.015.001.05 XchgId,omitempty"`
	CreDtTm    ISODateTime             `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.015.001.05 CreDtTm"`
	InitgPty   GenericIdentification53 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.015.001.05 InitgPty,omitempty"`
	RcptPty    GenericIdentification94 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.015.001.05 RcptPty,omitempty"`
	Tracblt    []Traceability5         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.015.001.05 Tracblt,omitempty"`
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

type Max10KBinary []byte

func (t *Max10KBinary) UnmarshalText(text []byte) error {
	return (*xsdBase64Binary)(t).UnmarshalText(text)
}
func (t Max10KBinary) MarshalText() ([]byte, error) {
	return xsdBase64Binary(t).MarshalText()
}

type Max140Binary []byte

func (t *Max140Binary) UnmarshalText(text []byte) error {
	return (*xsdBase64Binary)(t).UnmarshalText(text)
}
func (t Max140Binary) MarshalText() ([]byte, error) {
	return xsdBase64Binary(t).MarshalText()
}

type Max35Binary []byte

func (t *Max35Binary) UnmarshalText(text []byte) error {
	return (*xsdBase64Binary)(t).UnmarshalText(text)
}
func (t Max35Binary) MarshalText() ([]byte, error) {
	return xsdBase64Binary(t).MarshalText()
}

// Must be at least 1 items long
type Max35Text string

// Must be at least 1 items long
type Max500Text string

// Must be at least 1 items long
type Max6Text string

// Must be at least 1 items long
type Max70Text string

// May be one of RJCQ, RJCP
type MessageFunction9Code string

// Must match the pattern [a-zA-Z]{2,3}
type Min2Max3AlphaText string

type NetworkParameters4 struct {
	NtwkTp NetworkType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.015.001.05 NtwkTp"`
	AdrVal Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.015.001.05 AdrVal"`
}

type NetworkParameters5 struct {
	Adr        []NetworkParameters4 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.015.001.05 Adr"`
	UsrNm      Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.015.001.05 UsrNm,omitempty"`
	AccsCd     Max35Binary          `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.015.001.05 AccsCd,omitempty"`
	SvrCert    []Max10KBinary       `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.015.001.05 SvrCert,omitempty"`
	SvrCertIdr []Max140Binary       `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.015.001.05 SvrCertIdr,omitempty"`
	ClntCert   []Max10KBinary       `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.015.001.05 ClntCert,omitempty"`
	SctyPrfl   Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.015.001.05 SctyPrfl,omitempty"`
}

// May be one of IPNW, PSTN
type NetworkType1Code string

// May be one of OPOI, MERC, ACCP, ITAG, ACQR, CISS, DLIS
type PartyType3Code string

// May be one of MERC, ACCP, ITAG, ACQR, CISS, TAXH
type PartyType4Code string

// May be one of UNPR, IMSG, PARS, SECU, INTP, RCPP, DPMG, VERS, MSGT
type RejectReason1Code string

type Traceability5 struct {
	RlayId      GenericIdentification76 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.015.001.05 RlayId"`
	PrtcolNm    Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.015.001.05 PrtcolNm,omitempty"`
	PrtcolVrsn  Max6Text                `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.015.001.05 PrtcolVrsn,omitempty"`
	TracDtTmIn  ISODateTime             `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.015.001.05 TracDtTmIn"`
	TracDtTmOut ISODateTime             `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.015.001.05 TracDtTmOut"`
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