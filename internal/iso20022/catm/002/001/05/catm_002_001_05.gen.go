// Code generated by main. DO NOT EDIT.

package catm_002_001_05

import (
	"bytes"
	"encoding/base64"
	"encoding/xml"
	"time"
)

// May be one of HS25, HS38, HS51, HS01
type Algorithm11Code string

// May be one of MACC, MCCS, CMA1, MCC1, CMA9, CMA5
type Algorithm12Code string

// May be one of EA2C, E3DC, DKP9, UKPT, UKA1, EA9C, EA5C
type Algorithm13Code string

// May be one of ERS2, ERS1, RPSS
type Algorithm14Code string

// May be one of EA2C, E3DC, EA9C, EA5C
type Algorithm15Code string

// May be one of ERSA, RSAO
type Algorithm7Code string

// May be one of MGF1
type Algorithm8Code string

type AlgorithmIdentification11 struct {
	Algo  Algorithm7Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 Algo"`
	Param Parameter4     `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 Param,omitempty"`
}

type AlgorithmIdentification12 struct {
	Algo  Algorithm8Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 Algo"`
	Param Parameter5     `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 Param,omitempty"`
}

type AlgorithmIdentification13 struct {
	Algo  Algorithm13Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 Algo"`
	Param Parameter6      `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 Param,omitempty"`
}

type AlgorithmIdentification14 struct {
	Algo  Algorithm15Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 Algo"`
	Param Parameter6      `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 Param,omitempty"`
}

type AlgorithmIdentification15 struct {
	Algo  Algorithm12Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 Algo"`
	Param Parameter7      `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 Param,omitempty"`
}

type AlgorithmIdentification16 struct {
	Algo Algorithm11Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 Algo"`
}

type AlgorithmIdentification17 struct {
	Algo  Algorithm14Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 Algo"`
	Param Parameter8      `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 Param,omitempty"`
}

// May be one of CNAT, LATT, OATT, OUAT, CATT
type AttributeType1Code string

type AuthenticatedData4 struct {
	Vrsn        float64                   `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 Vrsn,omitempty"`
	Rcpt        []Recipient4Choice        `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 Rcpt"`
	MACAlgo     AlgorithmIdentification15 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 MACAlgo"`
	NcpsltdCntt EncapsulatedContent3      `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 NcpsltdCntt"`
	MAC         Max140Binary              `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 MAC"`
}

// May be one of LNGT, NUL8, NULG, NULL, RAND
type BytePadding1Code string

type CertificateIssuer1 struct {
	RltvDstngshdNm []RelativeDistinguishedName1 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 RltvDstngshdNm"`
}

type ContentInformationType12 struct {
	CnttTp       ContentType2Code   `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 CnttTp"`
	EnvlpdData   EnvelopedData4     `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 EnvlpdData,omitempty"`
	AuthntcdData AuthenticatedData4 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 AuthntcdData,omitempty"`
	SgndData     SignedData4        `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 SgndData,omitempty"`
	DgstdData    DigestedData4      `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 DgstdData,omitempty"`
}

// May be one of DATA, SIGN, EVLP, DGST, AUTH
type ContentType2Code string

// May be one of AQPR, APPR, TXCP, AKCP, DLGT, MGTP, MRPR, SCPR, SWPK, STRP, TRPR, VDPR, PARA, TMSP, CRTF
type DataSetCategory9Code string

type DataSetIdentification6 struct {
	Nm      Max256Text           `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 Nm,omitempty"`
	Tp      DataSetCategory9Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 Tp"`
	Vrsn    Max256Text           `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 Vrsn,omitempty"`
	CreDtTm ISODateTime          `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 CreDtTm,omitempty"`
}

type DigestedData4 struct {
	Vrsn        float64                   `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 Vrsn,omitempty"`
	DgstAlgo    AlgorithmIdentification16 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 DgstAlgo"`
	NcpsltdCntt EncapsulatedContent3      `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 NcpsltdCntt"`
	Dgst        Max140Binary              `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 Dgst"`
}

type Document struct {
	MgmtPlanRplcmnt ManagementPlanReplacementV05 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 MgmtPlanRplcmnt"`
}

type EncapsulatedContent3 struct {
	CnttTp ContentType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 CnttTp"`
	Cntt   Max100KBinary    `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 Cntt,omitempty"`
}

type EncryptedContent3 struct {
	CnttTp         ContentType2Code          `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 CnttTp"`
	CnttNcrptnAlgo AlgorithmIdentification14 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 CnttNcrptnAlgo"`
	NcrptdData     Max100KBinary             `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 NcrptdData"`
}

// May be one of TR31, TR34
type EncryptionFormat1Code string

type EnvelopedData4 struct {
	Vrsn       float64            `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 Vrsn,omitempty"`
	Rcpt       []Recipient4Choice `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 Rcpt"`
	NcrptdCntt EncryptedContent3  `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 NcrptdCntt,omitempty"`
}

type ErrorAction2 struct {
	ActnRslt  []TerminalManagementActionResult1Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 ActnRslt"`
	ActnToPrc TerminalManagementErrorAction2Code    `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 ActnToPrc"`
}

type GenericIdentification71 struct {
	Id     Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 Id"`
	Tp     PartyType5Code    `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 Tp,omitempty"`
	Issr   PartyType6Code    `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 Issr,omitempty"`
	Ctry   Min2Max3AlphaText `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 Ctry,omitempty"`
	ShrtNm Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 ShrtNm,omitempty"`
}

type GenericIdentification92 struct {
	Id       Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 Id"`
	Tp       PartyType5Code     `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 Tp,omitempty"`
	Issr     PartyType6Code     `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 Issr,omitempty"`
	Ctry     Min2Max3AlphaText  `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 Ctry,omitempty"`
	ShrtNm   Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 ShrtNm,omitempty"`
	RmotAccs NetworkParameters5 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 RmotAccs,omitempty"`
}

type Header27 struct {
	DwnldTrf bool                    `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 DwnldTrf"`
	FrmtVrsn Max6Text                `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 FrmtVrsn"`
	XchgId   float64                 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 XchgId"`
	CreDtTm  ISODateTime             `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 CreDtTm"`
	InitgPty GenericIdentification71 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 InitgPty"`
	RcptPty  GenericIdentification92 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 RcptPty,omitempty"`
}

type ISODateTime time.Time

func (t *ISODateTime) UnmarshalText(text []byte) error {
	return (*xsdDateTime)(t).UnmarshalText(text)
}
func (t ISODateTime) MarshalText() ([]byte, error) {
	return xsdDateTime(t).MarshalText()
}

type IssuerAndSerialNumber1 struct {
	Issr  CertificateIssuer1 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 Issr"`
	SrlNb Max35Binary        `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 SrlNb"`
}

type KEK4 struct {
	Vrsn          float64                   `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 Vrsn,omitempty"`
	KEKId         KEKIdentifier2            `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 KEKId"`
	KeyNcrptnAlgo AlgorithmIdentification13 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 KeyNcrptnAlgo"`
	NcrptdKey     Max500Binary              `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 NcrptdKey"`
}

type KEKIdentifier2 struct {
	KeyId     Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 KeyId"`
	KeyVrsn   Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 KeyVrsn"`
	SeqNb     float64         `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 SeqNb,omitempty"`
	DerivtnId Min5Max16Binary `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 DerivtnId,omitempty"`
}

type KeyTransport4 struct {
	Vrsn          float64                   `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 Vrsn,omitempty"`
	RcptId        Recipient5Choice          `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 RcptId"`
	KeyNcrptnAlgo AlgorithmIdentification11 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 KeyNcrptnAlgo"`
	NcrptdKey     Max5000Binary             `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 NcrptdKey"`
}

type ManagementPlan5 struct {
	POIId       GenericIdentification71     `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 POIId,omitempty"`
	TermnlMgrId GenericIdentification71     `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 TermnlMgrId"`
	DataSet     TerminalManagementDataSet18 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 DataSet"`
}

type ManagementPlanContent5 struct {
	TMChllng        Max140Binary   `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 TMChllng,omitempty"`
	KeyNcphrmntCert []Max10KBinary `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 KeyNcphrmntCert,omitempty"`
	Actn            []TMSAction5   `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 Actn"`
}

type ManagementPlanReplacementV05 struct {
	Hdr      Header27                 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 Hdr"`
	MgmtPlan ManagementPlan5          `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 MgmtPlan"`
	SctyTrlr ContentInformationType12 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 SctyTrlr,omitempty"`
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

// Must be at least 1 items long
type Max140Text string

// Must be at least 1 items long
type Max256Text string

type Max3000Binary []byte

func (t *Max3000Binary) UnmarshalText(text []byte) error {
	return (*xsdBase64Binary)(t).UnmarshalText(text)
}
func (t Max3000Binary) MarshalText() ([]byte, error) {
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

type Max5000Binary []byte

func (t *Max5000Binary) UnmarshalText(text []byte) error {
	return (*xsdBase64Binary)(t).UnmarshalText(text)
}
func (t Max5000Binary) MarshalText() ([]byte, error) {
	return xsdBase64Binary(t).MarshalText()
}

type Max500Binary []byte

func (t *Max500Binary) UnmarshalText(text []byte) error {
	return (*xsdBase64Binary)(t).UnmarshalText(text)
}
func (t Max500Binary) MarshalText() ([]byte, error) {
	return xsdBase64Binary(t).MarshalText()
}

// Must be at least 1 items long
type Max6Text string

// Must be at least 1 items long
type Max70Text string

// Must match the pattern [0-9]{1,9}
type Max9NumericText string

// Must match the pattern [a-zA-Z]{2,3}
type Min2Max3AlphaText string

type Min5Max16Binary []byte

func (t *Min5Max16Binary) UnmarshalText(text []byte) error {
	return (*xsdBase64Binary)(t).UnmarshalText(text)
}
func (t Min5Max16Binary) MarshalText() ([]byte, error) {
	return xsdBase64Binary(t).MarshalText()
}

type NetworkParameters4 struct {
	NtwkTp NetworkType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 NtwkTp"`
	AdrVal Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 AdrVal"`
}

type NetworkParameters5 struct {
	Adr        []NetworkParameters4 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 Adr"`
	UsrNm      Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 UsrNm,omitempty"`
	AccsCd     Max35Binary          `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 AccsCd,omitempty"`
	SvrCert    []Max10KBinary       `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 SvrCert,omitempty"`
	SvrCertIdr []Max140Binary       `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 SvrCertIdr,omitempty"`
	ClntCert   []Max10KBinary       `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 ClntCert,omitempty"`
	SctyPrfl   Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 SctyPrfl,omitempty"`
}

// May be one of IPNW, PSTN
type NetworkType1Code string

type Parameter4 struct {
	NcrptnFrmt   EncryptionFormat1Code     `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 NcrptnFrmt,omitempty"`
	DgstAlgo     Algorithm11Code           `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 DgstAlgo,omitempty"`
	MskGnrtrAlgo AlgorithmIdentification12 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 MskGnrtrAlgo,omitempty"`
}

type Parameter5 struct {
	DgstAlgo Algorithm11Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 DgstAlgo,omitempty"`
}

type Parameter6 struct {
	NcrptnFrmt   EncryptionFormat1Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 NcrptnFrmt,omitempty"`
	InitlstnVctr Max500Binary          `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 InitlstnVctr,omitempty"`
	BPddg        BytePadding1Code      `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 BPddg,omitempty"`
}

type Parameter7 struct {
	InitlstnVctr Max500Binary     `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 InitlstnVctr,omitempty"`
	BPddg        BytePadding1Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 BPddg,omitempty"`
}

type Parameter8 struct {
	DgstAlgo     Algorithm11Code           `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 DgstAlgo"`
	MskGnrtrAlgo AlgorithmIdentification12 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 MskGnrtrAlgo"`
	SaltLngth    float64                   `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 SaltLngth"`
	TrlrFld      float64                   `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 TrlrFld,omitempty"`
}

// May be one of OPOI, ACCP, MERC, ACQR, ITAG, MTMG, TMGT
type PartyType5Code string

// May be one of ACCP, MERC, ACQR, ITAG, MTMG, TMGT
type PartyType6Code string

type ProcessRetry2 struct {
	Dely  Max9NumericText `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 Dely"`
	MaxNb float64         `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 MaxNb,omitempty"`
}

type ProcessTiming3 struct {
	WtgTm   Max9NumericText `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 WtgTm,omitempty"`
	StartTm ISODateTime     `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 StartTm,omitempty"`
	EndTm   ISODateTime     `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 EndTm,omitempty"`
	Prd     Max9NumericText `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 Prd,omitempty"`
	MaxNb   float64         `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 MaxNb,omitempty"`
}

type Recipient4Choice struct {
	KeyTrnsprt KeyTransport4  `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 KeyTrnsprt"`
	KEK        KEK4           `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 KEK"`
	KeyIdr     KEKIdentifier2 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 KeyIdr"`
}

type Recipient5Choice struct {
	IssrAndSrlNb IssuerAndSerialNumber1 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 IssrAndSrlNb"`
	KeyIdr       KEKIdentifier2         `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 KeyIdr"`
}

type RelativeDistinguishedName1 struct {
	AttrTp  AttributeType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 AttrTp"`
	AttrVal Max140Text         `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 AttrVal"`
}

type SignedData4 struct {
	Vrsn        float64                     `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 Vrsn,omitempty"`
	DgstAlgo    []AlgorithmIdentification16 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 DgstAlgo"`
	NcpsltdCntt EncapsulatedContent3        `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 NcpsltdCntt"`
	Cert        []Max5000Binary             `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 Cert,omitempty"`
	Sgnr        []Signer3                   `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 Sgnr"`
}

type Signer3 struct {
	Vrsn      float64                   `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 Vrsn,omitempty"`
	SgnrId    Recipient5Choice          `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 SgnrId,omitempty"`
	DgstAlgo  AlgorithmIdentification16 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 DgstAlgo"`
	SgntrAlgo AlgorithmIdentification17 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 SgntrAlgo"`
	Sgntr     Max3000Binary             `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 Sgntr"`
}

type TMSAction5 struct {
	Tp               TerminalManagementAction2Code              `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 Tp"`
	RmotAccs         NetworkParameters5                         `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 RmotAccs,omitempty"`
	TermnlMgrId      GenericIdentification71                    `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 TermnlMgrId,omitempty"`
	TMSPrtcol        Max35Text                                  `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 TMSPrtcol,omitempty"`
	TMSPrtcolVrsn    Max35Text                                  `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 TMSPrtcolVrsn,omitempty"`
	DataSetId        DataSetIdentification6                     `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 DataSetId,omitempty"`
	CmpntTp          []DataSetCategory9Code                     `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 CmpntTp,omitempty"`
	DlgtnScpId       Max35Text                                  `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 DlgtnScpId,omitempty"`
	DlgtnScpDef      Max3000Binary                              `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 DlgtnScpDef,omitempty"`
	DlgtnProof       Max5000Binary                              `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 DlgtnProof,omitempty"`
	PrtctdDlgtnProof ContentInformationType12                   `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 PrtctdDlgtnProof,omitempty"`
	Trggr            TerminalManagementActionTrigger1Code       `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 Trggr"`
	AddtlPrc         []TerminalManagementAdditionalProcess1Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 AddtlPrc,omitempty"`
	ReTry            ProcessRetry2                              `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 ReTry,omitempty"`
	TmCond           ProcessTiming3                             `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 TmCond,omitempty"`
	TMChllng         Max140Binary                               `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 TMChllng,omitempty"`
	KeyNcphrmntCert  []Max10KBinary                             `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 KeyNcphrmntCert,omitempty"`
	ErrActn          []ErrorAction2                             `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 ErrActn,omitempty"`
	AddtlInf         []Max3000Binary                            `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 AddtlInf,omitempty"`
}

// May be one of ACTV, DCTV, DELT, DWNL, INST, RSTR, UPLD, UPDT
type TerminalManagementAction2Code string

// May be one of ACCD, CNTE, FMTE, INVC, LENE, OVER, MISS, NSUP, SIGE, SUCC, SYNE, TIMO, UKDT, UKRF
type TerminalManagementActionResult1Code string

// May be one of DATE, HOST, MANU, SALE
type TerminalManagementActionTrigger1Code string

// May be one of MANC, RCNC, RSRT
type TerminalManagementAdditionalProcess1Code string

type TerminalManagementDataSet18 struct {
	Id      DataSetIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 Id"`
	SeqCntr Max9NumericText        `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 SeqCntr,omitempty"`
	Cntt    ManagementPlanContent5 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 Cntt,omitempty"`
}

// May be one of SDSR, STOP
type TerminalManagementErrorAction2Code string

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
