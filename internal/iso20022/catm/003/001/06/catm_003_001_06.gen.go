// Code generated by main. DO NOT EDIT.

package catm_003_001_06

import (
	"bytes"
	"encoding/base64"
	"encoding/xml"
	"time"
)

type AcceptorConfiguration6 struct {
	TermnlMgrId GenericIdentification71       `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 TermnlMgrId"`
	DataSet     []TerminalManagementDataSet20 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 DataSet"`
}

type AcceptorConfigurationContent6 struct {
	RplcCfgtn         bool                               `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 RplcCfgtn,omitempty"`
	TMSPrtcolParams   []TMSProtocolParameters2           `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 TMSPrtcolParams,omitempty"`
	AcqrrPrtcolParams []AcquirerProtocolParameters10     `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 AcqrrPrtcolParams,omitempty"`
	MrchntParams      []MerchantConfigurationParameters3 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 MrchntParams,omitempty"`
	TermnlParams      []PaymentTerminalParameters4       `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 TermnlParams,omitempty"`
	ApplParams        []ApplicationParameters6           `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 ApplParams,omitempty"`
	HstComParams      []HostCommunicationParameter4      `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 HstComParams,omitempty"`
	SctyParams        []SecurityParameters6              `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 SctyParams,omitempty"`
}

type AcceptorConfigurationUpdateV06 struct {
	Hdr         Header27                 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 Hdr"`
	AccptrCfgtn AcceptorConfiguration6   `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 AccptrCfgtn"`
	SctyTrlr    ContentInformationType12 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 SctyTrlr,omitempty"`
}

type AcquirerHostConfiguration4 struct {
	HstId    Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 HstId"`
	MsgToSnd []MessageFunction12Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 MsgToSnd,omitempty"`
}

type AcquirerProtocolParameters10 struct {
	ActnTp         TerminalManagementAction3Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 ActnTp"`
	AcqrrId        []GenericIdentification53     `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 AcqrrId"`
	Vrsn           Max256Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 Vrsn"`
	ApplId         []Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 ApplId,omitempty"`
	Hst            []AcquirerHostConfiguration4  `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 Hst,omitempty"`
	OnLineTx       AcquirerProtocolParameters8   `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 OnLineTx,omitempty"`
	OffLineTx      AcquirerProtocolParameters8   `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 OffLineTx,omitempty"`
	RcncltnXchg    ExchangeConfiguration6        `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 RcncltnXchg,omitempty"`
	RcncltnByAcqrr bool                          `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 RcncltnByAcqrr,omitempty"`
	TtlsPerCcy     bool                          `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 TtlsPerCcy,omitempty"`
	SpltTtls       bool                          `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 SpltTtls,omitempty"`
	RcncltnErr     bool                          `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 RcncltnErr,omitempty"`
	CardDataVrfctn bool                          `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 CardDataVrfctn,omitempty"`
	NtfyOffLineCxl bool                          `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 NtfyOffLineCxl,omitempty"`
	BtchTrfCntt    []BatchTransactionType1Code   `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 BtchTrfCntt,omitempty"`
	FileTrfBtch    bool                          `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 FileTrfBtch,omitempty"`
	BtchDgtlSgntr  bool                          `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 BtchDgtlSgntr,omitempty"`
	MsgItm         []MessageItemCondition1       `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 MsgItm,omitempty"`
	PrtctCardData  bool                          `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 PrtctCardData"`
	MndtrySctyTrlr bool                          `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 MndtrySctyTrlr,omitempty"`
}

type AcquirerProtocolParameters8 struct {
	FinCaptr   FinancialCapture1Code    `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 FinCaptr"`
	BtchTrf    ExchangeConfiguration6   `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 BtchTrf,omitempty"`
	CmpltnXchg ExchangeConfiguration7   `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 CmpltnXchg,omitempty"`
	CxlXchg    CancellationProcess1Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 CxlXchg,omitempty"`
}

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
	Algo  Algorithm7Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 Algo"`
	Param Parameter4     `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 Param,omitempty"`
}

type AlgorithmIdentification12 struct {
	Algo  Algorithm8Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 Algo"`
	Param Parameter5     `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 Param,omitempty"`
}

type AlgorithmIdentification13 struct {
	Algo  Algorithm13Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 Algo"`
	Param Parameter6      `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 Param,omitempty"`
}

type AlgorithmIdentification14 struct {
	Algo  Algorithm15Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 Algo"`
	Param Parameter6      `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 Param,omitempty"`
}

type AlgorithmIdentification15 struct {
	Algo  Algorithm12Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 Algo"`
	Param Parameter7      `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 Param,omitempty"`
}

type AlgorithmIdentification16 struct {
	Algo Algorithm11Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 Algo"`
}

type AlgorithmIdentification17 struct {
	Algo  Algorithm14Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 Algo"`
	Param Parameter8      `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 Param,omitempty"`
}

type ApplicationParameters6 struct {
	ActnTp       TerminalManagementAction3Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 ActnTp"`
	ApplId       Max35Text                     `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 ApplId"`
	Vrsn         Max256Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 Vrsn,omitempty"`
	Params       []Max100KBinary               `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 Params,omitempty"`
	NcrptdParams ContentInformationType10      `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 NcrptdParams,omitempty"`
}

// May be one of CNAT, LATT, OATT, OUAT, CATT
type AttributeType1Code string

type AuthenticatedData4 struct {
	Vrsn        float64                   `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 Vrsn,omitempty"`
	Rcpt        []Recipient4Choice        `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 Rcpt"`
	MACAlgo     AlgorithmIdentification15 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 MACAlgo"`
	NcpsltdCntt EncapsulatedContent3      `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 NcpsltdCntt"`
	MAC         Max140Binary              `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 MAC"`
}

// May be one of DTCT, CNCL, FAIL, DCLN
type BatchTransactionType1Code string

// May be one of LNGT, NUL8, NULG, NULL, RAND
type BytePadding1Code string

// May be one of ADVC, NALW, REQU
type CancellationProcess1Code string

type CertificateIssuer1 struct {
	RltvDstngshdNm []RelativeDistinguishedName1 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 RltvDstngshdNm"`
}

type ClockSynchronisation1 struct {
	POITmZone Max70Text            `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 POITmZone"`
	SynctnSvr []NetworkParameters2 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 SynctnSvr,omitempty"`
}

type ContentInformationType10 struct {
	CnttTp     ContentType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 CnttTp"`
	EnvlpdData EnvelopedData4   `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 EnvlpdData"`
}

type ContentInformationType12 struct {
	CnttTp       ContentType2Code   `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 CnttTp"`
	EnvlpdData   EnvelopedData4     `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 EnvlpdData,omitempty"`
	AuthntcdData AuthenticatedData4 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 AuthntcdData,omitempty"`
	SgndData     SignedData4        `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 SgndData,omitempty"`
	DgstdData    DigestedData4      `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 DgstdData,omitempty"`
}

// May be one of DATA, SIGN, EVLP, DGST, AUTH
type ContentType2Code string

type CryptographicKey5 struct {
	Id         Max140Text                `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 Id"`
	AddtlId    Max35Binary               `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 AddtlId,omitempty"`
	Vrsn       Max256Text                `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 Vrsn"`
	Tp         CryptographicKeyType3Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 Tp"`
	Fctn       []KeyUsage1Code           `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 Fctn"`
	ActvtnDt   ISODateTime               `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 ActvtnDt,omitempty"`
	DeactvtnDt ISODateTime               `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 DeactvtnDt,omitempty"`
	KeyVal     ContentInformationType10  `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 KeyVal"`
}

// May be one of AES2, EDE3, DKP9, AES9, AES5, EDE4
type CryptographicKeyType3Code string

// May be one of AQPR, APPR, MTMG, MRPR, MTOR, SCPR, SWPK, TRPR, CRTF, TMSP
type DataSetCategory10Code string

// May be one of AQPR, APPR, TXCP, AKCP, DLGT, MGTP, MRPR, SCPR, SWPK, STRP, TRPR, VDPR, PARA, TMSP, CRTF
type DataSetCategory9Code string

type DataSetIdentification6 struct {
	Nm      Max256Text           `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 Nm,omitempty"`
	Tp      DataSetCategory9Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 Tp"`
	Vrsn    Max256Text           `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 Vrsn,omitempty"`
	CreDtTm ISODateTime          `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 CreDtTm,omitempty"`
}

type DigestedData4 struct {
	Vrsn        float64                   `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 Vrsn,omitempty"`
	DgstAlgo    AlgorithmIdentification16 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 DgstAlgo"`
	NcpsltdCntt EncapsulatedContent3      `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 NcpsltdCntt"`
	Dgst        Max140Binary              `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 Dgst"`
}

type Document struct {
	AccptrCfgtnUpd AcceptorConfigurationUpdateV06 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 AccptrCfgtnUpd"`
}

type EncapsulatedContent3 struct {
	CnttTp ContentType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 CnttTp"`
	Cntt   Max100KBinary    `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 Cntt,omitempty"`
}

type EncryptedContent3 struct {
	CnttTp         ContentType2Code          `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 CnttTp"`
	CnttNcrptnAlgo AlgorithmIdentification14 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 CnttNcrptnAlgo"`
	NcrptdData     Max100KBinary             `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 NcrptdData"`
}

// May be one of TR31, TR34
type EncryptionFormat1Code string

type EnvelopedData4 struct {
	Vrsn       float64            `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 Vrsn,omitempty"`
	Rcpt       []Recipient4Choice `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 Rcpt"`
	NcrptdCntt EncryptedContent3  `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 NcrptdCntt,omitempty"`
}

type ExchangeConfiguration6 struct {
	XchgPlcy []ExchangePolicy1Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 XchgPlcy"`
	MaxNb    float64               `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 MaxNb,omitempty"`
	MaxAmt   float64               `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 MaxAmt,omitempty"`
	ReTry    ProcessRetry2         `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 ReTry,omitempty"`
	TmCond   ProcessTiming4        `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 TmCond,omitempty"`
}

type ExchangeConfiguration7 struct {
	XchgPlcy  []ExchangePolicy1Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 XchgPlcy"`
	MaxNb     float64               `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 MaxNb,omitempty"`
	MaxAmt    float64               `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 MaxAmt,omitempty"`
	ReTry     ProcessRetry2         `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 ReTry,omitempty"`
	TmCond    ProcessTiming4        `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 TmCond,omitempty"`
	XchgFaild bool                  `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 XchgFaild,omitempty"`
	XchgDclnd bool                  `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 XchgDclnd,omitempty"`
}

// May be one of ONDM, IMMD, ASAP, AGRP, NBLT, TTLT, CYCL, NONE
type ExchangePolicy1Code string

// May be one of AUTH, COMP, BTCH
type FinancialCapture1Code string

type GenericIdentification53 struct {
	Id     Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 Id"`
	Tp     PartyType3Code    `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 Tp,omitempty"`
	Issr   PartyType4Code    `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 Issr,omitempty"`
	Ctry   Min2Max3AlphaText `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 Ctry,omitempty"`
	ShrtNm Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 ShrtNm,omitempty"`
}

type GenericIdentification71 struct {
	Id     Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 Id"`
	Tp     PartyType5Code    `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 Tp,omitempty"`
	Issr   PartyType6Code    `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 Issr,omitempty"`
	Ctry   Min2Max3AlphaText `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 Ctry,omitempty"`
	ShrtNm Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 ShrtNm,omitempty"`
}

type GenericIdentification92 struct {
	Id       Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 Id"`
	Tp       PartyType5Code     `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 Tp,omitempty"`
	Issr     PartyType6Code     `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 Issr,omitempty"`
	Ctry     Min2Max3AlphaText  `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 Ctry,omitempty"`
	ShrtNm   Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 ShrtNm,omitempty"`
	RmotAccs NetworkParameters5 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 RmotAccs,omitempty"`
}

type Header27 struct {
	DwnldTrf bool                    `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 DwnldTrf"`
	FrmtVrsn Max6Text                `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 FrmtVrsn"`
	XchgId   float64                 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 XchgId"`
	CreDtTm  ISODateTime             `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 CreDtTm"`
	InitgPty GenericIdentification71 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 InitgPty"`
	RcptPty  GenericIdentification92 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 RcptPty,omitempty"`
}

type HostCommunicationParameter4 struct {
	ActnTp       TerminalManagementAction3Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 ActnTp"`
	HstId        Max35Text                     `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 HstId"`
	Adr          NetworkParameters3            `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 Adr,omitempty"`
	Key          []KEKIdentifier5              `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 Key,omitempty"`
	NtwkSvcPrvdr NetworkParameters5            `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 NtwkSvcPrvdr,omitempty"`
}

type ISODateTime time.Time

func (t *ISODateTime) UnmarshalText(text []byte) error {
	return (*xsdDateTime)(t).UnmarshalText(text)
}
func (t ISODateTime) MarshalText() ([]byte, error) {
	return xsdDateTime(t).MarshalText()
}

type ISOTime time.Time

func (t *ISOTime) UnmarshalText(text []byte) error {
	return (*xsdTime)(t).UnmarshalText(text)
}
func (t ISOTime) MarshalText() ([]byte, error) {
	return xsdTime(t).MarshalText()
}

type IssuerAndSerialNumber1 struct {
	Issr  CertificateIssuer1 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 Issr"`
	SrlNb Max35Binary        `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 SrlNb"`
}

type KEK4 struct {
	Vrsn          float64                   `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 Vrsn,omitempty"`
	KEKId         KEKIdentifier2            `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 KEKId"`
	KeyNcrptnAlgo AlgorithmIdentification13 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 KeyNcrptnAlgo"`
	NcrptdKey     Max500Binary              `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 NcrptdKey"`
}

type KEKIdentifier2 struct {
	KeyId     Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 KeyId"`
	KeyVrsn   Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 KeyVrsn"`
	SeqNb     float64         `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 SeqNb,omitempty"`
	DerivtnId Min5Max16Binary `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 DerivtnId,omitempty"`
}

type KEKIdentifier5 struct {
	KeyId     Max140Text                `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 KeyId"`
	KeyVrsn   Max140Text                `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 KeyVrsn"`
	SeqNb     float64                   `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 SeqNb,omitempty"`
	DerivtnId Min5Max16Binary           `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 DerivtnId,omitempty"`
	Tp        CryptographicKeyType3Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 Tp,omitempty"`
	Fctn      []KeyUsage1Code           `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 Fctn,omitempty"`
}

type KeyTransport4 struct {
	Vrsn          float64                   `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 Vrsn,omitempty"`
	RcptId        Recipient5Choice          `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 RcptId"`
	KeyNcrptnAlgo AlgorithmIdentification11 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 KeyNcrptnAlgo"`
	NcrptdKey     Max5000Binary             `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 NcrptdKey"`
}

// May be one of ENCR, DCPT, DENC, DDEC, TRNI, TRNX, MACG, MACV, SIGG, SUGV, PINE, PIND, PINV, KEYG, KEYI, KEYX, KEYD
type KeyUsage1Code string

type LocalDateTime1 struct {
	FrDtTm    ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 FrDtTm,omitempty"`
	ToDtTm    ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 ToDtTm,omitempty"`
	UTCOffset float64     `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 UTCOffset"`
}

type Max10000Binary []byte

func (t *Max10000Binary) UnmarshalText(text []byte) error {
	return (*xsdBase64Binary)(t).UnmarshalText(text)
}
func (t Max10000Binary) MarshalText() ([]byte, error) {
	return xsdBase64Binary(t).MarshalText()
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

type MerchantConfigurationParameters3 struct {
	ActnTp     TerminalManagementAction3Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 ActnTp"`
	MrchntId   Max35Text                     `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 MrchntId,omitempty"`
	Vrsn       Max256Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 Vrsn,omitempty"`
	Prxy       NetworkParameters6            `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 Prxy,omitempty"`
	OthrParams Max10000Binary                `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 OthrParams,omitempty"`
}

// May be one of FAUQ, CCAQ, CMPV, DGNP, RCLQ, CCAV, BTCH, FRVA, AUTQ, FCMV, DCCQ, RVRA, DCRR
type MessageFunction12Code string

type MessageItemCondition1 struct {
	ItmId Max140Text                `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 ItmId"`
	Cond  MessageItemCondition1Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 Cond"`
	Val   []Max140Text              `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 Val,omitempty"`
}

// May be one of MNDT, CFVL, DFLT, ALWV, IFAV, COPY, UNSP
type MessageItemCondition1Code string

// Must match the pattern [a-zA-Z]{2,3}
type Min2Max3AlphaText string

type Min5Max16Binary []byte

func (t *Min5Max16Binary) UnmarshalText(text []byte) error {
	return (*xsdBase64Binary)(t).UnmarshalText(text)
}
func (t Min5Max16Binary) MarshalText() ([]byte, error) {
	return xsdBase64Binary(t).MarshalText()
}

type NetworkParameters2 struct {
	Adr    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 Adr"`
	PortNb float64   `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 PortNb,omitempty"`
	Dely   ISOTime   `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 Dely,omitempty"`
}

type NetworkParameters3 struct {
	Adr        []NetworkParameters4 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 Adr"`
	UsrNm      Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 UsrNm,omitempty"`
	AccsCd     Max35Binary          `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 AccsCd,omitempty"`
	SvrCert    []Max3000Binary      `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 SvrCert,omitempty"`
	SvrCertIdr []Max140Binary       `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 SvrCertIdr,omitempty"`
}

type NetworkParameters4 struct {
	NtwkTp NetworkType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 NtwkTp"`
	AdrVal Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 AdrVal"`
}

type NetworkParameters5 struct {
	Adr        []NetworkParameters4 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 Adr"`
	UsrNm      Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 UsrNm,omitempty"`
	AccsCd     Max35Binary          `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 AccsCd,omitempty"`
	SvrCert    []Max10KBinary       `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 SvrCert,omitempty"`
	SvrCertIdr []Max140Binary       `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 SvrCertIdr,omitempty"`
	ClntCert   []Max10KBinary       `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 ClntCert,omitempty"`
	SctyPrfl   Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 SctyPrfl,omitempty"`
}

type NetworkParameters6 struct {
	Tp   NetworkType2Code   `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 Tp"`
	Accs NetworkParameters5 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 Accs"`
}

// May be one of IPNW, PSTN
type NetworkType1Code string

// May be one of SCK5, SCK4, HTTP
type NetworkType2Code string

type Parameter4 struct {
	NcrptnFrmt   EncryptionFormat1Code     `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 NcrptnFrmt,omitempty"`
	DgstAlgo     Algorithm11Code           `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 DgstAlgo,omitempty"`
	MskGnrtrAlgo AlgorithmIdentification12 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 MskGnrtrAlgo,omitempty"`
}

type Parameter5 struct {
	DgstAlgo Algorithm11Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 DgstAlgo,omitempty"`
}

type Parameter6 struct {
	NcrptnFrmt   EncryptionFormat1Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 NcrptnFrmt,omitempty"`
	InitlstnVctr Max500Binary          `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 InitlstnVctr,omitempty"`
	BPddg        BytePadding1Code      `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 BPddg,omitempty"`
}

type Parameter7 struct {
	InitlstnVctr Max500Binary     `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 InitlstnVctr,omitempty"`
	BPddg        BytePadding1Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 BPddg,omitempty"`
}

type Parameter8 struct {
	DgstAlgo     Algorithm11Code           `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 DgstAlgo"`
	MskGnrtrAlgo AlgorithmIdentification12 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 MskGnrtrAlgo"`
	SaltLngth    float64                   `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 SaltLngth"`
	TrlrFld      float64                   `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 TrlrFld,omitempty"`
}

// May be one of PGRP, PSYS, PSNG
type PartyType15Code string

// May be one of OPOI, MERC, ACCP, ITAG, ACQR, CISS, DLIS
type PartyType3Code string

// May be one of MERC, ACCP, ITAG, ACQR, CISS, TAXH
type PartyType4Code string

// May be one of OPOI, ACCP, MERC, ACQR, ITAG, MTMG, TMGT
type PartyType5Code string

// May be one of ACCP, MERC, ACQR, ITAG, MTMG, TMGT
type PartyType6Code string

type PaymentTerminalParameters4 struct {
	ActnTp     TerminalManagementAction3Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 ActnTp"`
	VndrId     Max35Text                     `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 VndrId,omitempty"`
	Vrsn       Max256Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 Vrsn,omitempty"`
	ClckSynctn ClockSynchronisation1         `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 ClckSynctn,omitempty"`
	TmZoneLine []Max70Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 TmZoneLine,omitempty"`
	LclDtTm    []LocalDateTime1              `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 LclDtTm,omitempty"`
	OthrParams Max10000Binary                `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 OthrParams,omitempty"`
}

type ProcessRetry2 struct {
	Dely  Max9NumericText `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 Dely"`
	MaxNb float64         `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 MaxNb,omitempty"`
}

type ProcessTiming4 struct {
	StartTm ISODateTime     `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 StartTm,omitempty"`
	EndTm   ISODateTime     `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 EndTm,omitempty"`
	Prd     Max9NumericText `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 Prd,omitempty"`
}

type Recipient4Choice struct {
	KeyTrnsprt KeyTransport4  `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 KeyTrnsprt"`
	KEK        KEK4           `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 KEK"`
	KeyIdr     KEKIdentifier2 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 KeyIdr"`
}

type Recipient5Choice struct {
	IssrAndSrlNb IssuerAndSerialNumber1 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 IssrAndSrlNb"`
	KeyIdr       KEKIdentifier2         `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 KeyIdr"`
}

type RelativeDistinguishedName1 struct {
	AttrTp  AttributeType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 AttrTp"`
	AttrVal Max140Text         `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 AttrVal"`
}

type SecurityParameters6 struct {
	ActnTp    TerminalManagementAction3Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 ActnTp"`
	Vrsn      Max256Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 Vrsn"`
	POIChllng Max140Binary                  `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 POIChllng,omitempty"`
	TMChllng  Max140Binary                  `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 TMChllng,omitempty"`
	SmmtrcKey []CryptographicKey5           `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 SmmtrcKey,omitempty"`
}

type SignedData4 struct {
	Vrsn        float64                     `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 Vrsn,omitempty"`
	DgstAlgo    []AlgorithmIdentification16 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 DgstAlgo"`
	NcpsltdCntt EncapsulatedContent3        `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 NcpsltdCntt"`
	Cert        []Max5000Binary             `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 Cert,omitempty"`
	Sgnr        []Signer3                   `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 Sgnr"`
}

type Signer3 struct {
	Vrsn      float64                   `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 Vrsn,omitempty"`
	SgnrId    Recipient5Choice          `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 SgnrId,omitempty"`
	DgstAlgo  AlgorithmIdentification16 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 DgstAlgo"`
	SgntrAlgo AlgorithmIdentification17 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 SgntrAlgo"`
	Sgntr     Max3000Binary             `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 Sgntr"`
}

type TMSProtocolParameters2 struct {
	ActnTp      TerminalManagementAction3Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 ActnTp"`
	TermnlMgrId GenericIdentification71       `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 TermnlMgrId"`
	MntncSvc    []DataSetCategory10Code       `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 MntncSvc"`
	Vrsn        Max256Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 Vrsn"`
	ApplId      []Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 ApplId,omitempty"`
	HstId       Max35Text                     `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 HstId"`
	POIId       Max35Text                     `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 POIId,omitempty"`
	InitgPtyId  Max35Text                     `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 InitgPtyId,omitempty"`
	RcptPtyId   Max35Text                     `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 RcptPtyId,omitempty"`
	FileTrf     bool                          `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 FileTrf,omitempty"`
}

// May be one of CREA, DELT, UPDT
type TerminalManagementAction3Code string

type TerminalManagementDataSet20 struct {
	Id       DataSetIdentification6        `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 Id"`
	SeqCntr  Max9NumericText               `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 SeqCntr,omitempty"`
	POIId    []GenericIdentification71     `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 POIId,omitempty"`
	CfgtnScp PartyType15Code               `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 CfgtnScp,omitempty"`
	Cntt     AcceptorConfigurationContent6 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.06 Cntt"`
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

type xsdTime time.Time

func (t *xsdTime) UnmarshalText(text []byte) error {
	return _unmarshalTime(text, (*time.Time)(t), "15:04:05.999999999")
}
func (t xsdTime) MarshalText() ([]byte, error) {
	return _marshalTime((time.Time)(t), "15:04:05.999999999")
}
func (t xsdTime) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if (time.Time)(t).IsZero() {
		return nil
	}
	m, err := t.MarshalText()
	if err != nil {
		return err
	}
	return e.EncodeElement(m, start)
}
func (t xsdTime) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if (time.Time)(t).IsZero() {
		return xml.Attr{}, nil
	}
	m, err := t.MarshalText()
	return xml.Attr{Name: name, Value: string(m)}, err
}