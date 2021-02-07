// Code generated by main. DO NOT EDIT.

package catm_006_001_03

import (
	"bytes"
	"encoding/base64"
	"encoding/xml"
	"time"
)

// May be one of HS25, HS38, HS51, HS01
type Algorithm11Code string

// May be one of HS25, HS38, HS51, HS01, SH31, SH32, SH33, SH35, SHK1, SHK2
type Algorithm16Code string

// May be one of MACC, MCCS, CMA1, MCC1, CMA9, CMA5, CMA2, CM31, CM32, CM33, MCS3, CCA1, CCA2, CCA3
type Algorithm17Code string

// May be one of EA2C, E3DC, DKP9, UKPT, UKA1, EA9C, EA5C, DA12, DA19, DA25, N108, EA5R, EA9R, EA2R, E3DR, E36C, E36R, SD5C
type Algorithm18Code string

// May be one of ERS2, ERS1, RPSS, ECC5, ECC1, ECC4, ECC2, ECC3, ERS3, ECP2, ECP3, ECP5
type Algorithm19Code string

// May be one of ERSA, RSAO
type Algorithm7Code string

// May be one of MGF1
type Algorithm8Code string

type AlgorithmIdentification12 struct {
	Algo  Algorithm8Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.006.001.03 Algo"`
	Param Parameter5     `xml:"urn:iso:std:iso:20022:tech:xsd:catm.006.001.03 Param,omitempty"`
}

type AlgorithmIdentification18 struct {
	Algo  Algorithm8Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.006.001.03 Algo"`
	Param Parameter9     `xml:"urn:iso:std:iso:20022:tech:xsd:catm.006.001.03 Param,omitempty"`
}

type AlgorithmIdentification19 struct {
	Algo  Algorithm7Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.006.001.03 Algo"`
	Param Parameter10    `xml:"urn:iso:std:iso:20022:tech:xsd:catm.006.001.03 Param,omitempty"`
}

type AlgorithmIdentification20 struct {
	Algo  Algorithm19Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.006.001.03 Algo"`
	Param Parameter11     `xml:"urn:iso:std:iso:20022:tech:xsd:catm.006.001.03 Param,omitempty"`
}

type AlgorithmIdentification21 struct {
	Algo Algorithm16Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.006.001.03 Algo"`
}

type AlgorithmIdentification22 struct {
	Algo  Algorithm17Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.006.001.03 Algo"`
	Param Parameter7      `xml:"urn:iso:std:iso:20022:tech:xsd:catm.006.001.03 Param,omitempty"`
}

type AlgorithmIdentification23 struct {
	Algo  Algorithm18Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.006.001.03 Algo"`
	Param Parameter12     `xml:"urn:iso:std:iso:20022:tech:xsd:catm.006.001.03 Param,omitempty"`
}

type AlgorithmIdentification24 struct {
	Algo  Algorithm18Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.006.001.03 Algo"`
	Param Parameter12     `xml:"urn:iso:std:iso:20022:tech:xsd:catm.006.001.03 Param,omitempty"`
}

// May be one of CNAT, LATT, OATT, OUAT, CATT
type AttributeType1Code string

type AuthenticatedData5 struct {
	Vrsn        float64                   `xml:"urn:iso:std:iso:20022:tech:xsd:catm.006.001.03 Vrsn,omitempty"`
	Rcpt        []Recipient6Choice        `xml:"urn:iso:std:iso:20022:tech:xsd:catm.006.001.03 Rcpt"`
	MACAlgo     AlgorithmIdentification22 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.006.001.03 MACAlgo"`
	NcpsltdCntt EncapsulatedContent3      `xml:"urn:iso:std:iso:20022:tech:xsd:catm.006.001.03 NcpsltdCntt"`
	MAC         Max140Binary              `xml:"urn:iso:std:iso:20022:tech:xsd:catm.006.001.03 MAC"`
}

// May be one of LNGT, NUL8, NULG, NULL, RAND
type BytePadding1Code string

type CertificateIssuer1 struct {
	RltvDstngshdNm []RelativeDistinguishedName1 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.006.001.03 RltvDstngshdNm"`
}

type ContentInformationType18 struct {
	CnttTp       ContentType2Code   `xml:"urn:iso:std:iso:20022:tech:xsd:catm.006.001.03 CnttTp"`
	AuthntcdData AuthenticatedData5 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.006.001.03 AuthntcdData,omitempty"`
	SgndData     SignedData5        `xml:"urn:iso:std:iso:20022:tech:xsd:catm.006.001.03 SgndData,omitempty"`
}

type ContentInformationType19 struct {
	CnttTp       ContentType2Code   `xml:"urn:iso:std:iso:20022:tech:xsd:catm.006.001.03 CnttTp"`
	EnvlpdData   EnvelopedData5     `xml:"urn:iso:std:iso:20022:tech:xsd:catm.006.001.03 EnvlpdData,omitempty"`
	AuthntcdData AuthenticatedData5 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.006.001.03 AuthntcdData,omitempty"`
	SgndData     SignedData5        `xml:"urn:iso:std:iso:20022:tech:xsd:catm.006.001.03 SgndData,omitempty"`
	DgstdData    DigestedData5      `xml:"urn:iso:std:iso:20022:tech:xsd:catm.006.001.03 DgstdData,omitempty"`
}

// May be one of DATA, SIGN, EVLP, DGST, AUTH
type ContentType2Code string

// May be one of ACQP, APPR, APSB, KDWL, KMGT, RPRT, SWPK, TMSP, MRPR, TRPR, CRTF
type DataSetCategory11Code string

type DigestedData5 struct {
	Vrsn        float64                   `xml:"urn:iso:std:iso:20022:tech:xsd:catm.006.001.03 Vrsn,omitempty"`
	DgstAlgo    AlgorithmIdentification21 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.006.001.03 DgstAlgo"`
	NcpsltdCntt EncapsulatedContent3      `xml:"urn:iso:std:iso:20022:tech:xsd:catm.006.001.03 NcpsltdCntt"`
	Dgst        Max140Binary              `xml:"urn:iso:std:iso:20022:tech:xsd:catm.006.001.03 Dgst"`
}

type Document struct {
	MntncDlgtnRspn MaintenanceDelegationResponseV03 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.006.001.03 MntncDlgtnRspn"`
}

type EncapsulatedContent3 struct {
	CnttTp ContentType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.006.001.03 CnttTp"`
	Cntt   Max100KBinary    `xml:"urn:iso:std:iso:20022:tech:xsd:catm.006.001.03 Cntt,omitempty"`
}

type EncryptedContent4 struct {
	CnttTp         ContentType2Code          `xml:"urn:iso:std:iso:20022:tech:xsd:catm.006.001.03 CnttTp"`
	CnttNcrptnAlgo AlgorithmIdentification24 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.006.001.03 CnttNcrptnAlgo,omitempty"`
	NcrptdData     Max100KBinary             `xml:"urn:iso:std:iso:20022:tech:xsd:catm.006.001.03 NcrptdData"`
}

// May be one of TR31, TR34, I238
type EncryptionFormat2Code string

type EnvelopedData5 struct {
	Vrsn       float64                `xml:"urn:iso:std:iso:20022:tech:xsd:catm.006.001.03 Vrsn,omitempty"`
	OrgtrInf   OriginatorInformation1 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.006.001.03 OrgtrInf,omitempty"`
	Rcpt       []Recipient6Choice     `xml:"urn:iso:std:iso:20022:tech:xsd:catm.006.001.03 Rcpt"`
	NcrptdCntt EncryptedContent4      `xml:"urn:iso:std:iso:20022:tech:xsd:catm.006.001.03 NcrptdCntt,omitempty"`
}

type GenericIdentification72 struct {
	Id     Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:catm.006.001.03 Id"`
	Issr   PartyType6Code    `xml:"urn:iso:std:iso:20022:tech:xsd:catm.006.001.03 Issr,omitempty"`
	Ctry   Min2Max3AlphaText `xml:"urn:iso:std:iso:20022:tech:xsd:catm.006.001.03 Ctry,omitempty"`
	ShrtNm Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:catm.006.001.03 ShrtNm,omitempty"`
}

type GenericIdentification93 struct {
	Id       Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:catm.006.001.03 Id"`
	Issr     PartyType6Code     `xml:"urn:iso:std:iso:20022:tech:xsd:catm.006.001.03 Issr,omitempty"`
	Ctry     Min2Max3AlphaText  `xml:"urn:iso:std:iso:20022:tech:xsd:catm.006.001.03 Ctry,omitempty"`
	ShrtNm   Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:catm.006.001.03 ShrtNm,omitempty"`
	RmotAccs NetworkParameters5 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.006.001.03 RmotAccs,omitempty"`
}

type GenericInformation1 struct {
	Nm  Max70Text  `xml:"urn:iso:std:iso:20022:tech:xsd:catm.006.001.03 Nm"`
	Val Max140Text `xml:"urn:iso:std:iso:20022:tech:xsd:catm.006.001.03 Val,omitempty"`
}

type Header29 struct {
	PrtcolVrsn Max6Text                `xml:"urn:iso:std:iso:20022:tech:xsd:catm.006.001.03 PrtcolVrsn"`
	XchgId     float64                 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.006.001.03 XchgId,omitempty"`
	CreDtTm    ISODateTime             `xml:"urn:iso:std:iso:20022:tech:xsd:catm.006.001.03 CreDtTm"`
	InitgPty   GenericIdentification72 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.006.001.03 InitgPty"`
	RcptPty    GenericIdentification93 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.006.001.03 RcptPty,omitempty"`
}

type ISODateTime time.Time

func (t *ISODateTime) UnmarshalText(text []byte) error {
	return (*xsdDateTime)(t).UnmarshalText(text)
}
func (t ISODateTime) MarshalText() ([]byte, error) {
	return xsdDateTime(t).MarshalText()
}

type IssuerAndSerialNumber1 struct {
	Issr  CertificateIssuer1 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.006.001.03 Issr"`
	SrlNb Max35Binary        `xml:"urn:iso:std:iso:20022:tech:xsd:catm.006.001.03 SrlNb"`
}

type KEK5 struct {
	Vrsn          float64                   `xml:"urn:iso:std:iso:20022:tech:xsd:catm.006.001.03 Vrsn,omitempty"`
	KEKId         KEKIdentifier2            `xml:"urn:iso:std:iso:20022:tech:xsd:catm.006.001.03 KEKId"`
	KeyNcrptnAlgo AlgorithmIdentification23 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.006.001.03 KeyNcrptnAlgo"`
	NcrptdKey     Max500Binary              `xml:"urn:iso:std:iso:20022:tech:xsd:catm.006.001.03 NcrptdKey"`
}

type KEKIdentifier2 struct {
	KeyId     Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:catm.006.001.03 KeyId"`
	KeyVrsn   Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:catm.006.001.03 KeyVrsn"`
	SeqNb     float64         `xml:"urn:iso:std:iso:20022:tech:xsd:catm.006.001.03 SeqNb,omitempty"`
	DerivtnId Min5Max16Binary `xml:"urn:iso:std:iso:20022:tech:xsd:catm.006.001.03 DerivtnId,omitempty"`
}

type KeyTransport5 struct {
	Vrsn          float64                   `xml:"urn:iso:std:iso:20022:tech:xsd:catm.006.001.03 Vrsn,omitempty"`
	RcptId        Recipient5Choice          `xml:"urn:iso:std:iso:20022:tech:xsd:catm.006.001.03 RcptId"`
	KeyNcrptnAlgo AlgorithmIdentification19 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.006.001.03 KeyNcrptnAlgo"`
	NcrptdKey     Max5000Binary             `xml:"urn:iso:std:iso:20022:tech:xsd:catm.006.001.03 NcrptdKey"`
}

type MaintenanceDelegation6 struct {
	MntncSvc         []DataSetCategory11Code                 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.006.001.03 MntncSvc"`
	Rspn             Response2Code                           `xml:"urn:iso:std:iso:20022:tech:xsd:catm.006.001.03 Rspn"`
	RspnRsn          Max35Text                               `xml:"urn:iso:std:iso:20022:tech:xsd:catm.006.001.03 RspnRsn,omitempty"`
	DlgtnTp          TerminalManagementAction3Code           `xml:"urn:iso:std:iso:20022:tech:xsd:catm.006.001.03 DlgtnTp"`
	POISubset        []Max35Text                             `xml:"urn:iso:std:iso:20022:tech:xsd:catm.006.001.03 POISubset,omitempty"`
	DlgtnScpId       Max35Text                               `xml:"urn:iso:std:iso:20022:tech:xsd:catm.006.001.03 DlgtnScpId,omitempty"`
	DlgtnScpDef      Max3000Binary                           `xml:"urn:iso:std:iso:20022:tech:xsd:catm.006.001.03 DlgtnScpDef,omitempty"`
	DlgtnProof       Max5000Binary                           `xml:"urn:iso:std:iso:20022:tech:xsd:catm.006.001.03 DlgtnProof,omitempty"`
	PrtctdDlgtnProof ContentInformationType19                `xml:"urn:iso:std:iso:20022:tech:xsd:catm.006.001.03 PrtctdDlgtnProof,omitempty"`
	POIIdAssoctn     []MaintenanceIdentificationAssociation1 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.006.001.03 POIIdAssoctn,omitempty"`
}

type MaintenanceDelegationResponse3 struct {
	TMId      GenericIdentification72  `xml:"urn:iso:std:iso:20022:tech:xsd:catm.006.001.03 TMId"`
	MstrTMId  GenericIdentification72  `xml:"urn:iso:std:iso:20022:tech:xsd:catm.006.001.03 MstrTMId,omitempty"`
	DlgtnRspn []MaintenanceDelegation6 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.006.001.03 DlgtnRspn"`
}

type MaintenanceDelegationResponseV03 struct {
	Hdr            Header29                       `xml:"urn:iso:std:iso:20022:tech:xsd:catm.006.001.03 Hdr"`
	MntncDlgtnRspn MaintenanceDelegationResponse3 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.006.001.03 MntncDlgtnRspn"`
	SctyTrlr       ContentInformationType18       `xml:"urn:iso:std:iso:20022:tech:xsd:catm.006.001.03 SctyTrlr,omitempty"`
}

type MaintenanceIdentificationAssociation1 struct {
	MstrTMId Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:catm.006.001.03 MstrTMId"`
	TMId     Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:catm.006.001.03 TMId"`
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
	NtwkTp NetworkType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.006.001.03 NtwkTp"`
	AdrVal Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:catm.006.001.03 AdrVal"`
}

type NetworkParameters5 struct {
	Adr        []NetworkParameters4 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.006.001.03 Adr"`
	UsrNm      Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:catm.006.001.03 UsrNm,omitempty"`
	AccsCd     Max35Binary          `xml:"urn:iso:std:iso:20022:tech:xsd:catm.006.001.03 AccsCd,omitempty"`
	SvrCert    []Max10KBinary       `xml:"urn:iso:std:iso:20022:tech:xsd:catm.006.001.03 SvrCert,omitempty"`
	SvrCertIdr []Max140Binary       `xml:"urn:iso:std:iso:20022:tech:xsd:catm.006.001.03 SvrCertIdr,omitempty"`
	ClntCert   []Max10KBinary       `xml:"urn:iso:std:iso:20022:tech:xsd:catm.006.001.03 ClntCert,omitempty"`
	SctyPrfl   Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:catm.006.001.03 SctyPrfl,omitempty"`
}

// May be one of IPNW, PSTN
type NetworkType1Code string

type OriginatorInformation1 struct {
	Cert []Max5000Binary `xml:"urn:iso:std:iso:20022:tech:xsd:catm.006.001.03 Cert,omitempty"`
}

type Parameter10 struct {
	NcrptnFrmt   EncryptionFormat2Code     `xml:"urn:iso:std:iso:20022:tech:xsd:catm.006.001.03 NcrptnFrmt,omitempty"`
	DgstAlgo     Algorithm16Code           `xml:"urn:iso:std:iso:20022:tech:xsd:catm.006.001.03 DgstAlgo,omitempty"`
	MskGnrtrAlgo AlgorithmIdentification18 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.006.001.03 MskGnrtrAlgo,omitempty"`
}

type Parameter11 struct {
	DgstAlgo     Algorithm16Code           `xml:"urn:iso:std:iso:20022:tech:xsd:catm.006.001.03 DgstAlgo"`
	MskGnrtrAlgo AlgorithmIdentification12 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.006.001.03 MskGnrtrAlgo"`
	SaltLngth    float64                   `xml:"urn:iso:std:iso:20022:tech:xsd:catm.006.001.03 SaltLngth"`
	TrlrFld      float64                   `xml:"urn:iso:std:iso:20022:tech:xsd:catm.006.001.03 TrlrFld,omitempty"`
}

type Parameter12 struct {
	NcrptnFrmt   EncryptionFormat2Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.006.001.03 NcrptnFrmt,omitempty"`
	InitlstnVctr Max500Binary          `xml:"urn:iso:std:iso:20022:tech:xsd:catm.006.001.03 InitlstnVctr,omitempty"`
	BPddg        BytePadding1Code      `xml:"urn:iso:std:iso:20022:tech:xsd:catm.006.001.03 BPddg,omitempty"`
}

type Parameter5 struct {
	DgstAlgo Algorithm11Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.006.001.03 DgstAlgo,omitempty"`
}

type Parameter7 struct {
	InitlstnVctr Max500Binary     `xml:"urn:iso:std:iso:20022:tech:xsd:catm.006.001.03 InitlstnVctr,omitempty"`
	BPddg        BytePadding1Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.006.001.03 BPddg,omitempty"`
}

type Parameter9 struct {
	DgstAlgo Algorithm16Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.006.001.03 DgstAlgo,omitempty"`
}

// May be one of ACCP, MERC, ACQR, ITAG, MTMG, TMGT
type PartyType6Code string

type Recipient5Choice struct {
	IssrAndSrlNb IssuerAndSerialNumber1 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.006.001.03 IssrAndSrlNb"`
	KeyIdr       KEKIdentifier2         `xml:"urn:iso:std:iso:20022:tech:xsd:catm.006.001.03 KeyIdr"`
}

type Recipient6Choice struct {
	KeyTrnsprt KeyTransport5  `xml:"urn:iso:std:iso:20022:tech:xsd:catm.006.001.03 KeyTrnsprt"`
	KEK        KEK5           `xml:"urn:iso:std:iso:20022:tech:xsd:catm.006.001.03 KEK"`
	KeyIdr     KEKIdentifier2 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.006.001.03 KeyIdr"`
}

type RelativeDistinguishedName1 struct {
	AttrTp  AttributeType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.006.001.03 AttrTp"`
	AttrVal Max140Text         `xml:"urn:iso:std:iso:20022:tech:xsd:catm.006.001.03 AttrVal"`
}

// May be one of APPR, DECL
type Response2Code string

type SignedData5 struct {
	Vrsn        float64                     `xml:"urn:iso:std:iso:20022:tech:xsd:catm.006.001.03 Vrsn,omitempty"`
	DgstAlgo    []AlgorithmIdentification21 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.006.001.03 DgstAlgo,omitempty"`
	NcpsltdCntt EncapsulatedContent3        `xml:"urn:iso:std:iso:20022:tech:xsd:catm.006.001.03 NcpsltdCntt,omitempty"`
	Cert        []Max5000Binary             `xml:"urn:iso:std:iso:20022:tech:xsd:catm.006.001.03 Cert,omitempty"`
	Sgnr        []Signer4                   `xml:"urn:iso:std:iso:20022:tech:xsd:catm.006.001.03 Sgnr,omitempty"`
}

type Signer4 struct {
	Vrsn        float64                   `xml:"urn:iso:std:iso:20022:tech:xsd:catm.006.001.03 Vrsn,omitempty"`
	SgnrId      Recipient5Choice          `xml:"urn:iso:std:iso:20022:tech:xsd:catm.006.001.03 SgnrId,omitempty"`
	DgstAlgo    AlgorithmIdentification21 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.006.001.03 DgstAlgo"`
	SgndAttrbts []GenericInformation1     `xml:"urn:iso:std:iso:20022:tech:xsd:catm.006.001.03 SgndAttrbts,omitempty"`
	SgntrAlgo   AlgorithmIdentification20 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.006.001.03 SgntrAlgo"`
	Sgntr       Max3000Binary             `xml:"urn:iso:std:iso:20022:tech:xsd:catm.006.001.03 Sgntr"`
}

// May be one of CREA, DELT, UPDT
type TerminalManagementAction3Code string

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