// Code generated by main. DO NOT EDIT.

package caaa_015_001_06

import (
	"bytes"
	"encoding/base64"
	"encoding/xml"
	"time"
)

type AcceptorRejection2 struct {
	RjctRsn  RejectReason1Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.015.001.06 RjctRsn"`
	AddtlInf Max500Text        `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.015.001.06 AddtlInf,omitempty"`
	MsgInErr Max100KBinary     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.015.001.06 MsgInErr,omitempty"`
}

type AcceptorRejectionV06 struct {
	Hdr  Header57           `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.015.001.06 Hdr"`
	Rjct AcceptorRejection2 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.015.001.06 Rjct"`
}

type Document struct {
	AccptrRjctn AcceptorRejectionV06 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.015.001.06 AccptrRjctn"`
}

type GenericIdentification176 struct {
	Id     Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.015.001.06 Id"`
	Tp     PartyType33Code   `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.015.001.06 Tp,omitempty"`
	Issr   PartyType33Code   `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.015.001.06 Issr,omitempty"`
	Ctry   Min2Max3AlphaText `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.015.001.06 Ctry,omitempty"`
	ShrtNm Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.015.001.06 ShrtNm,omitempty"`
}

type GenericIdentification177 struct {
	Id       Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.015.001.06 Id"`
	Tp       PartyType33Code    `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.015.001.06 Tp,omitempty"`
	Issr     PartyType33Code    `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.015.001.06 Issr,omitempty"`
	Ctry     Min2Max3AlphaText  `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.015.001.06 Ctry,omitempty"`
	ShrtNm   Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.015.001.06 ShrtNm,omitempty"`
	RmotAccs NetworkParameters7 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.015.001.06 RmotAccs,omitempty"`
	Glctn    Geolocation1       `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.015.001.06 Glctn,omitempty"`
}

type Geolocation1 struct {
	GeogcCordints GeolocationGeographicCoordinates1 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.015.001.06 GeogcCordints,omitempty"`
	UTMCordints   GeolocationUTMCoordinates1        `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.015.001.06 UTMCordints,omitempty"`
}

type GeolocationGeographicCoordinates1 struct {
	Lat  Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.015.001.06 Lat"`
	Long Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.015.001.06 Long"`
}

type GeolocationUTMCoordinates1 struct {
	UTMZone    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.015.001.06 UTMZone"`
	UTMEstwrd  Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.015.001.06 UTMEstwrd"`
	UTMNrthwrd Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.015.001.06 UTMNrthwrd"`
}

type Header57 struct {
	MsgFctn    MessageFunction9Code     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.015.001.06 MsgFctn"`
	PrtcolVrsn Max6Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.015.001.06 PrtcolVrsn"`
	XchgId     float64                  `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.015.001.06 XchgId,omitempty"`
	CreDtTm    ISODateTime              `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.015.001.06 CreDtTm"`
	InitgPty   GenericIdentification176 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.015.001.06 InitgPty,omitempty"`
	RcptPty    GenericIdentification177 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.015.001.06 RcptPty,omitempty"`
	Tracblt    []Traceability8          `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.015.001.06 Tracblt,omitempty"`
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

// May be one of RJCQ, RJCP
type MessageFunction9Code string

// Must match the pattern [a-zA-Z]{2,3}
type Min2Max3AlphaText string

type NetworkParameters7 struct {
	Adr        []NetworkParameters9 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.015.001.06 Adr"`
	UsrNm      Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.015.001.06 UsrNm,omitempty"`
	AccsCd     Max35Binary          `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.015.001.06 AccsCd,omitempty"`
	SvrCert    []Max10KBinary       `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.015.001.06 SvrCert,omitempty"`
	SvrCertIdr []Max140Binary       `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.015.001.06 SvrCertIdr,omitempty"`
	ClntCert   []Max10KBinary       `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.015.001.06 ClntCert,omitempty"`
	SctyPrfl   Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.015.001.06 SctyPrfl,omitempty"`
}

type NetworkParameters9 struct {
	NtwkTp NetworkType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.015.001.06 NtwkTp"`
	AdrVal Max500Text       `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.015.001.06 AdrVal"`
}

// May be one of IPNW, PSTN
type NetworkType1Code string

// May be one of OPOI, MERC, ACCP, ITAG, ACQR, CISS, DLIS, MTMG, TAXH, TMGT
type PartyType33Code string

// May be one of UNPR, IMSG, PARS, SECU, INTP, RCPP, DPMG, VERS, MSGT
type RejectReason1Code string

type Traceability8 struct {
	RlayId      GenericIdentification177 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.015.001.06 RlayId"`
	PrtcolNm    Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.015.001.06 PrtcolNm,omitempty"`
	PrtcolVrsn  Max6Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.015.001.06 PrtcolVrsn,omitempty"`
	TracDtTmIn  ISODateTime              `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.015.001.06 TracDtTmIn"`
	TracDtTmOut ISODateTime              `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.015.001.06 TracDtTmOut"`
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
