// Code generated by main. DO NOT EDIT.

package catm_007_001_03

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

// May be one of ERS2, ERS1, RPSS, ECC5, ECC1, ECC4, ECC2, ECC3, ERS3, ECP2, ECP3, ECP5
type Algorithm19Code string

// May be one of EA2C, E3DC, DKP9, UKPT, UKA2, EA9C, EA5C, DA12, DA19, DA25, N108, EA5R, EA9R, EA2R, E3DR, E36C, E36R, SD5C, UKA1, UKA3
type Algorithm24Code string

// May be one of ERSA, RSAO
type Algorithm7Code string

// May be one of MGF1
type Algorithm8Code string

type AlgorithmIdentification12 struct {
	Algo  Algorithm8Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.03 Algo"`
	Param Parameter5     `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.03 Param,omitempty"`
}

type AlgorithmIdentification18 struct {
	Algo  Algorithm8Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.03 Algo"`
	Param Parameter9     `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.03 Param,omitempty"`
}

type AlgorithmIdentification19 struct {
	Algo  Algorithm7Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.03 Algo"`
	Param Parameter10    `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.03 Param,omitempty"`
}

type AlgorithmIdentification20 struct {
	Algo  Algorithm19Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.03 Algo"`
	Param Parameter11     `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.03 Param,omitempty"`
}

type AlgorithmIdentification21 struct {
	Algo Algorithm16Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.03 Algo"`
}

type AlgorithmIdentification22 struct {
	Algo  Algorithm17Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.03 Algo"`
	Param Parameter7      `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.03 Param,omitempty"`
}

type AlgorithmIdentification29 struct {
	Algo  Algorithm24Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.03 Algo"`
	Param Parameter12     `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.03 Param,omitempty"`
}

// May be one of CNAT, LATT, OATT, OUAT, CATT
type AttributeType1Code string

// May be one of EMAL, CHLG
type AttributeType2Code string

type AuthenticatedData6 struct {
	Vrsn        float64                   `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.03 Vrsn,omitempty"`
	Rcpt        []Recipient8Choice        `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.03 Rcpt"`
	MACAlgo     AlgorithmIdentification22 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.03 MACAlgo"`
	NcpsltdCntt EncapsulatedContent3      `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.03 NcpsltdCntt"`
	MAC         Max140Binary              `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.03 MAC"`
}

// May be one of LNGT, NUL8, NULG, NULL, RAND
type BytePadding1Code string

// May be one of CRTC, CRTR, CRTK, WLSR, WLSA
type CardPaymentServiceType10Code string

type CertificateIssuer1 struct {
	RltvDstngshdNm []RelativeDistinguishedName1 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.03 RltvDstngshdNm"`
}

type CertificateManagementRequest2 struct {
	POIId            GenericIdentification176     `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.03 POIId"`
	TMId             GenericIdentification176     `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.03 TMId,omitempty"`
	CertSvc          CardPaymentServiceType10Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.03 CertSvc"`
	SctyDomn         Max70Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.03 SctyDomn,omitempty"`
	BinryCertfctnReq Max20000Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.03 BinryCertfctnReq,omitempty"`
	CertfctnReq      CertificationRequest1        `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.03 CertfctnReq,omitempty"`
	ClntCert         Max10KBinary                 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.03 ClntCert,omitempty"`
	WhtListId        PointOfInteraction6          `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.03 WhtListId,omitempty"`
}

type CertificateManagementRequestV03 struct {
	Hdr         TMSHeader1                    `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.03 Hdr"`
	CertMgmtReq CertificateManagementRequest2 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.03 CertMgmtReq"`
	SctyTrlr    ContentInformationType21      `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.03 SctyTrlr,omitempty"`
}

type CertificationRequest1 struct {
	CertReqInf CertificationRequest2 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.03 CertReqInf"`
	KeyId      Max140Text            `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.03 KeyId,omitempty"`
	KeyVrsn    Max140Text            `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.03 KeyVrsn,omitempty"`
}

type CertificationRequest2 struct {
	Vrsn           float64                      `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.03 Vrsn,omitempty"`
	SbjtNm         CertificateIssuer1           `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.03 SbjtNm,omitempty"`
	SbjtPblcKeyInf PublicRSAKey2                `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.03 SbjtPblcKeyInf"`
	Attr           []RelativeDistinguishedName2 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.03 Attr"`
}

type ContentInformationType21 struct {
	CnttTp       ContentType2Code   `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.03 CnttTp"`
	AuthntcdData AuthenticatedData6 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.03 AuthntcdData,omitempty"`
	SgndData     SignedData5        `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.03 SgndData,omitempty"`
}

// May be one of DATA, SIGN, EVLP, DGST, AUTH
type ContentType2Code string

type Document struct {
	CertMgmtReq CertificateManagementRequestV03 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.03 CertMgmtReq"`
}

type EncapsulatedContent3 struct {
	CnttTp ContentType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.03 CnttTp"`
	Cntt   Max100KBinary    `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.03 Cntt,omitempty"`
}

// May be one of TR31, TR34, I238
type EncryptionFormat2Code string

type GenericIdentification176 struct {
	Id     Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.03 Id"`
	Tp     PartyType33Code   `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.03 Tp,omitempty"`
	Issr   PartyType33Code   `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.03 Issr,omitempty"`
	Ctry   Min2Max3AlphaText `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.03 Ctry,omitempty"`
	ShrtNm Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.03 ShrtNm,omitempty"`
}

type GenericIdentification177 struct {
	Id       Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.03 Id"`
	Tp       PartyType33Code    `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.03 Tp,omitempty"`
	Issr     PartyType33Code    `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.03 Issr,omitempty"`
	Ctry     Min2Max3AlphaText  `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.03 Ctry,omitempty"`
	ShrtNm   Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.03 ShrtNm,omitempty"`
	RmotAccs NetworkParameters7 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.03 RmotAccs,omitempty"`
	Glctn    Geolocation1       `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.03 Glctn,omitempty"`
}

type GenericInformation1 struct {
	Nm  Max70Text  `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.03 Nm"`
	Val Max140Text `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.03 Val,omitempty"`
}

type Geolocation1 struct {
	GeogcCordints GeolocationGeographicCoordinates1 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.03 GeogcCordints,omitempty"`
	UTMCordints   GeolocationUTMCoordinates1        `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.03 UTMCordints,omitempty"`
}

type GeolocationGeographicCoordinates1 struct {
	Lat  Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.03 Lat"`
	Long Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.03 Long"`
}

type GeolocationUTMCoordinates1 struct {
	UTMZone    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.03 UTMZone"`
	UTMEstwrd  Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.03 UTMEstwrd"`
	UTMNrthwrd Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.03 UTMNrthwrd"`
}

type ISODateTime time.Time

func (t *ISODateTime) UnmarshalText(text []byte) error {
	return (*xsdDateTime)(t).UnmarshalText(text)
}
func (t ISODateTime) MarshalText() ([]byte, error) {
	return xsdDateTime(t).MarshalText()
}

type IssuerAndSerialNumber1 struct {
	Issr  CertificateIssuer1 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.03 Issr"`
	SrlNb Max35Binary        `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.03 SrlNb"`
}

type KEK7 struct {
	Vrsn          float64                   `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.03 Vrsn,omitempty"`
	KEKId         KEKIdentifier2            `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.03 KEKId"`
	KeyNcrptnAlgo AlgorithmIdentification29 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.03 KeyNcrptnAlgo"`
	NcrptdKey     Max500Binary              `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.03 NcrptdKey"`
}

type KEKIdentifier2 struct {
	KeyId     Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.03 KeyId"`
	KeyVrsn   Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.03 KeyVrsn"`
	SeqNb     float64         `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.03 SeqNb,omitempty"`
	DerivtnId Min5Max16Binary `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.03 DerivtnId,omitempty"`
}

type KeyTransport5 struct {
	Vrsn          float64                   `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.03 Vrsn,omitempty"`
	RcptId        Recipient5Choice          `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.03 RcptId"`
	KeyNcrptnAlgo AlgorithmIdentification19 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.03 KeyNcrptnAlgo"`
	NcrptdKey     Max5000Binary             `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.03 NcrptdKey"`
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
type Max20000Text string

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
type Max500Text string

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

type NetworkParameters7 struct {
	Adr        []NetworkParameters9 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.03 Adr"`
	UsrNm      Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.03 UsrNm,omitempty"`
	AccsCd     Max35Binary          `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.03 AccsCd,omitempty"`
	SvrCert    []Max10KBinary       `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.03 SvrCert,omitempty"`
	SvrCertIdr []Max140Binary       `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.03 SvrCertIdr,omitempty"`
	ClntCert   []Max10KBinary       `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.03 ClntCert,omitempty"`
	SctyPrfl   Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.03 SctyPrfl,omitempty"`
}

type NetworkParameters9 struct {
	NtwkTp NetworkType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.03 NtwkTp"`
	AdrVal Max500Text       `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.03 AdrVal"`
}

// May be one of IPNW, PSTN
type NetworkType1Code string

type Parameter10 struct {
	NcrptnFrmt   EncryptionFormat2Code     `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.03 NcrptnFrmt,omitempty"`
	DgstAlgo     Algorithm16Code           `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.03 DgstAlgo,omitempty"`
	MskGnrtrAlgo AlgorithmIdentification18 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.03 MskGnrtrAlgo,omitempty"`
}

type Parameter11 struct {
	DgstAlgo     Algorithm16Code           `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.03 DgstAlgo"`
	MskGnrtrAlgo AlgorithmIdentification12 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.03 MskGnrtrAlgo"`
	SaltLngth    float64                   `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.03 SaltLngth"`
	TrlrFld      float64                   `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.03 TrlrFld,omitempty"`
}

type Parameter12 struct {
	NcrptnFrmt   EncryptionFormat2Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.03 NcrptnFrmt,omitempty"`
	InitlstnVctr Max500Binary          `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.03 InitlstnVctr,omitempty"`
	BPddg        BytePadding1Code      `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.03 BPddg,omitempty"`
}

type Parameter5 struct {
	DgstAlgo Algorithm11Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.03 DgstAlgo,omitempty"`
}

type Parameter7 struct {
	InitlstnVctr Max500Binary     `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.03 InitlstnVctr,omitempty"`
	BPddg        BytePadding1Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.03 BPddg,omitempty"`
}

type Parameter9 struct {
	DgstAlgo Algorithm16Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.03 DgstAlgo,omitempty"`
}

// May be one of OPOI, MERC, ACCP, ITAG, ACQR, CISS, DLIS, MTMG, TAXH, TMGT
type PartyType33Code string

type PointOfInteraction6 struct {
	ManfctrIdr Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.03 ManfctrIdr"`
	Mdl        Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.03 Mdl"`
	SrlNb      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.03 SrlNb"`
}

type PublicRSAKey1 struct {
	Mdlus Max5000Binary `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.03 Mdlus"`
	Expnt Max5000Binary `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.03 Expnt"`
}

type PublicRSAKey2 struct {
	Algo       Algorithm7Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.03 Algo,omitempty"`
	PblcKeyVal PublicRSAKey1  `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.03 PblcKeyVal"`
}

type Recipient5Choice struct {
	IssrAndSrlNb IssuerAndSerialNumber1 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.03 IssrAndSrlNb"`
	KeyIdr       KEKIdentifier2         `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.03 KeyIdr"`
}

type Recipient8Choice struct {
	KeyTrnsprt KeyTransport5  `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.03 KeyTrnsprt"`
	KEK        KEK7           `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.03 KEK"`
	KeyIdr     KEKIdentifier2 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.03 KeyIdr"`
}

type RelativeDistinguishedName1 struct {
	AttrTp  AttributeType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.03 AttrTp"`
	AttrVal Max140Text         `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.03 AttrVal"`
}

type RelativeDistinguishedName2 struct {
	AttrTp  AttributeType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.03 AttrTp"`
	AttrVal Max140Text         `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.03 AttrVal"`
}

type SignedData5 struct {
	Vrsn        float64                     `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.03 Vrsn,omitempty"`
	DgstAlgo    []AlgorithmIdentification21 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.03 DgstAlgo,omitempty"`
	NcpsltdCntt EncapsulatedContent3        `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.03 NcpsltdCntt,omitempty"`
	Cert        []Max5000Binary             `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.03 Cert,omitempty"`
	Sgnr        []Signer4                   `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.03 Sgnr,omitempty"`
}

type Signer4 struct {
	Vrsn        float64                   `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.03 Vrsn,omitempty"`
	SgnrId      Recipient5Choice          `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.03 SgnrId,omitempty"`
	DgstAlgo    AlgorithmIdentification21 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.03 DgstAlgo"`
	SgndAttrbts []GenericInformation1     `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.03 SgndAttrbts,omitempty"`
	SgntrAlgo   AlgorithmIdentification20 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.03 SgntrAlgo"`
	Sgntr       Max3000Binary             `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.03 Sgntr"`
}

type TMSHeader1 struct {
	DwnldTrf bool                     `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.03 DwnldTrf"`
	FrmtVrsn Max6Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.03 FrmtVrsn"`
	XchgId   float64                  `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.03 XchgId"`
	CreDtTm  ISODateTime              `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.03 CreDtTm"`
	InitgPty GenericIdentification176 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.03 InitgPty"`
	RcptPty  GenericIdentification177 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.03 RcptPty,omitempty"`
	Tracblt  []Traceability8          `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.03 Tracblt,omitempty"`
}

type Traceability8 struct {
	RlayId      GenericIdentification177 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.03 RlayId"`
	PrtcolNm    Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.03 PrtcolNm,omitempty"`
	PrtcolVrsn  Max6Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.03 PrtcolVrsn,omitempty"`
	TracDtTmIn  ISODateTime              `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.03 TracDtTmIn"`
	TracDtTmOut ISODateTime              `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.03 TracDtTmOut"`
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
