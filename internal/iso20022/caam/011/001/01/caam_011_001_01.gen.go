// Code generated by main. DO NOT EDIT.

package caam_011_001_01

import (
	"bytes"
	"encoding/base64"
	"encoding/xml"
	"time"
)

type ATMConfigurationParameter1 struct {
	Tp   DataSetCategory7Code `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 Tp"`
	Vrsn Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 Vrsn"`
}

type ATMContext20 struct {
	SsnRef Max35Text    `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 SsnRef,omitempty"`
	Svc    ATMService24 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 Svc"`
}

type ATMCustomer6 struct {
	Prfl         ATMCustomerProfile4              `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 Prfl,omitempty"`
	SelctdLang   string                           `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 SelctdLang,omitempty"`
	AuthntcnRslt []TransactionVerificationResult5 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 AuthntcnRslt"`
}

// May be one of CRDF, OREQ, PREQ
type ATMCustomerProfile1Code string

type ATMCustomerProfile4 struct {
	RtrvlMd ATMCustomerProfile1Code `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 RtrvlMd"`
	PrflRef Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 PrflRef,omitempty"`
	CstmrId Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 CstmrId,omitempty"`
}

type ATMEnvironment16 struct {
	Acqrr    Acquirer7               `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 Acqrr,omitempty"`
	ATMMgrId Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 ATMMgrId,omitempty"`
	HstgNtty TerminalHosting1        `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 HstgNtty,omitempty"`
	ATM      AutomatedTellerMachine9 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 ATM"`
	Cstmr    ATMCustomer6            `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 Cstmr,omitempty"`
	Card     PaymentCard23           `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 Card,omitempty"`
}

type ATMEquipment1 struct {
	Manfctr    Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 Manfctr,omitempty"`
	Mdl        Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 Mdl,omitempty"`
	SrlNb      Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 SrlNb,omitempty"`
	ApplPrvdr  Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 ApplPrvdr,omitempty"`
	ApplNm     Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 ApplNm,omitempty"`
	ApplVrsn   Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 ApplVrsn,omitempty"`
	ApprvlNb   Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 ApprvlNb,omitempty"`
	CfgtnParam []ATMConfigurationParameter1 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 CfgtnParam,omitempty"`
}

type ATMExceptionAdvice1 struct {
	Envt  ATMEnvironment16 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 Envt,omitempty"`
	Cntxt ATMContext20     `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 Cntxt,omitempty"`
	Tx    ATMTransaction27 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 Tx"`
}

type ATMExceptionAdviceV01 struct {
	Hdr                Header32                 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 Hdr"`
	PrtctdATMXcptnAdvc ContentInformationType10 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 PrtctdATMXcptnAdvc,omitempty"`
	ATMXcptnAdvc       ATMExceptionAdvice1      `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 ATMXcptnAdvc,omitempty"`
	SctyTrlr           ContentInformationType15 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 SctyTrlr,omitempty"`
}

// May be one of CARD, COIN, CMDT, CPNS, NOTE, STMP, UDTM
type ATMMediaType1Code string

// May be one of CARD, COIN, CMDT, CPNS, NOTE, STMP, UDTM, CHCK
type ATMMediaType2Code string

type ATMMessageFunction2 struct {
	Fctn     MessageFunction11Code `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 Fctn"`
	ATMSvcCd Max35Text             `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 ATMSvcCd,omitempty"`
	HstSvcCd Max35Text             `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 HstSvcCd,omitempty"`
}

type ATMService24 struct {
	SvcRef     Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 SvcRef,omitempty"`
	ATMSvcCd   Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 ATMSvcCd,omitempty"`
	HstSvcCd   Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 HstSvcCd,omitempty"`
	SvcTp      ATMServiceType10Code `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 SvcTp"`
	SvcVarntId []Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 SvcVarntId,omitempty"`
}

// May be one of TRFC, TRFI, TRFP, ASTS, BLCQ, CDVF, CHSN, CMPF, DCCS, XRTD, XRTW, MCHG, DPSN, PINC, PINR, PINU, PATH, PRFL, EMVS, STDR, SPRV, DPSV
type ATMServiceType10Code string

type ATMTransaction27 struct {
	TxId          TransactionIdentifier1 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 TxId,omitempty"`
	RcncltnId     Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 RcncltnId,omitempty"`
	Xcptn         []FailureReason8Code   `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 Xcptn"`
	XcptnDtl      []Max70Text            `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 XcptnDtl,omitempty"`
	ElctrncPrsBal CurrencyAndAmount      `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 ElctrncPrsBal,omitempty"`
}

type Acquirer7 struct {
	AcqrgInstn Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 AcqrgInstn,omitempty"`
	Brnch      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 Brnch,omitempty"`
}

// Must match the pattern [A-Z]{3,3}
type ActiveCurrencyCode string

// May be one of HS25, HS38, HS51, HS01
type Algorithm11Code string

// May be one of MACC, MCCS, CMA1, MCC1, CMA9, CMA5
type Algorithm12Code string

// May be one of EA2C, E3DC, DKP9, UKPT, UKA1, EA9C, EA5C
type Algorithm13Code string

// May be one of EA2C, E3DC, EA9C, EA5C
type Algorithm15Code string

// May be one of ERSA, RSAO
type Algorithm7Code string

// May be one of MGF1
type Algorithm8Code string

type AlgorithmIdentification11 struct {
	Algo  Algorithm7Code `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 Algo"`
	Param Parameter4     `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 Param,omitempty"`
}

type AlgorithmIdentification12 struct {
	Algo  Algorithm8Code `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 Algo"`
	Param Parameter5     `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 Param,omitempty"`
}

type AlgorithmIdentification13 struct {
	Algo  Algorithm13Code `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 Algo"`
	Param Parameter6      `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 Param,omitempty"`
}

type AlgorithmIdentification14 struct {
	Algo  Algorithm15Code `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 Algo"`
	Param Parameter6      `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 Param,omitempty"`
}

type AlgorithmIdentification15 struct {
	Algo  Algorithm12Code `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 Algo"`
	Param Parameter7      `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 Param,omitempty"`
}

// May be one of CNAT, LATT, OATT, OUAT, CATT
type AttributeType1Code string

type AuthenticatedData4 struct {
	Vrsn        float64                   `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 Vrsn,omitempty"`
	Rcpt        []Recipient4Choice        `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 Rcpt"`
	MACAlgo     AlgorithmIdentification15 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 MACAlgo"`
	NcpsltdCntt EncapsulatedContent3      `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 NcpsltdCntt"`
	MAC         Max140Binary              `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 MAC"`
}

// May be one of ICCD, AGNT, MERC, ACQR, ISSR, TRML
type AuthenticationEntity2Code string

// May be one of TOKA, BIOM, MOBL, OTHR, FPIN, NPIN, PSWD, SCRT, SCNL
type AuthenticationMethod7Code string

type AutomatedTellerMachine9 struct {
	Id       Max35Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 Id"`
	AddtlId  Max35Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 AddtlId,omitempty"`
	SeqNb    Max35Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 SeqNb,omitempty"`
	BaseCcy  ActiveCurrencyCode              `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 BaseCcy"`
	Lctn     PostalAddress17                 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 Lctn,omitempty"`
	LctnCtgy TransactionEnvironment2Code     `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 LctnCtgy,omitempty"`
	Cpblties PointOfInteractionCapabilities7 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 Cpblties,omitempty"`
	Eqpmnt   ATMEquipment1                   `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 Eqpmnt,omitempty"`
}

// May be one of LNGT, NUL8, NULG, NULL, RAND
type BytePadding1Code string

// May be one of TAGC, PHYS, BRCD, MGST, CICC, DFLE, CTLS, ECTL
type CardDataReading1Code string

// May be one of ECTL, CICC, MGST, CTLS
type CardDataReading4Code string

// May be one of FFLB, SFLB, NFLB
type CardFallback1Code string

// May be one of NPIN, FCPN, FEPN, FDSG, FBIO, FBIG, PKIS, PCOD
type CardholderVerificationCapability3Code string

type CertificateIssuer1 struct {
	RltvDstngshdNm []RelativeDistinguishedName1 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 RltvDstngshdNm"`
}

type ContentInformationType10 struct {
	CnttTp     ContentType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 CnttTp"`
	EnvlpdData EnvelopedData4   `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 EnvlpdData"`
}

type ContentInformationType15 struct {
	CnttTp       ContentType2Code   `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 CnttTp"`
	AuthntcdData AuthenticatedData4 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 AuthntcdData"`
}

// May be one of DATA, SIGN, EVLP, DGST, AUTH
type ContentType2Code string

// Must match the pattern [A-Z]{2,2}
type CountryCode string

type CurrencyAndAmount struct {
	Value float64      `xml:",chardata"`
	Ccy   CurrencyCode `xml:"Ccy,attr"`
}

// Must match the pattern [A-Z]{3,3}
type CurrencyCode string

// May be one of ATMC, ATMP, APPR, CRAP, CPRC, OEXR, AMNT, LOCC, MNOC
type DataSetCategory7Code string

type DisplayCapabilities5 struct {
	Dstn      []UserInterface5Code `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 Dstn"`
	AvlblFrmt []OutputFormat1Code  `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 AvlblFrmt,omitempty"`
	NbOfLines float64              `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 NbOfLines,omitempty"`
	LineWidth float64              `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 LineWidth,omitempty"`
	AvlblLang []string             `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 AvlblLang,omitempty"`
}

type Document struct {
	ATMXcptnAdvc ATMExceptionAdviceV01 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 ATMXcptnAdvc"`
}

type EncapsulatedContent3 struct {
	CnttTp ContentType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 CnttTp"`
	Cntt   Max100KBinary    `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 Cntt,omitempty"`
}

type EncryptedContent3 struct {
	CnttTp         ContentType2Code          `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 CnttTp"`
	CnttNcrptnAlgo AlgorithmIdentification14 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 CnttNcrptnAlgo"`
	NcrptdData     Max100KBinary             `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 NcrptdData"`
}

// May be one of TR31, TR34
type EncryptionFormat1Code string

type EnvelopedData4 struct {
	Vrsn       float64            `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 Vrsn,omitempty"`
	Rcpt       []Recipient4Choice `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 Rcpt"`
	NcrptdCntt EncryptedContent3  `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 NcrptdCntt,omitempty"`
}

// Must match the pattern [a-zA-Z0-9]{3}
type Exact3AlphaNumericText string

// May be one of CDRT, CDCP, CUCL, CDFG, MALF, SECU, SFRD, UCPT
type FailureReason8Code string

type GenericIdentification77 struct {
	Id     Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 Id"`
	Tp     PartyType12Code   `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 Tp"`
	Issr   PartyType12Code   `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 Issr,omitempty"`
	Ctry   Min2Max3AlphaText `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 Ctry,omitempty"`
	ShrtNm Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 ShrtNm,omitempty"`
}

type GeographicCoordinates1 struct {
	Lat  Max16Text `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 Lat"`
	Long Max16Text `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 Long"`
}

type GeographicLocation1Choice struct {
	GeogcCordints GeographicCoordinates1 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 GeogcCordints"`
	UTMCordints   UTMCoordinates1        `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 UTMCordints"`
}

type Header32 struct {
	MsgFctn        ATMMessageFunction2 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 MsgFctn"`
	PrtcolVrsn     Max6Text            `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 PrtcolVrsn"`
	XchgId         Max3NumericText     `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 XchgId"`
	ReTrnsmssnCntr float64             `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 ReTrnsmssnCntr,omitempty"`
	CreDtTm        ISODateTime         `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 CreDtTm"`
	InitgPty       Max35Text           `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 InitgPty"`
	RcptPty        Max35Text           `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 RcptPty,omitempty"`
	PrcStat        Max35Text           `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 PrcStat,omitempty"`
	Tracblt        []Traceability4     `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 Tracblt,omitempty"`
}

type ISODateTime time.Time

func (t *ISODateTime) UnmarshalText(text []byte) error {
	return (*xsdDateTime)(t).UnmarshalText(text)
}
func (t ISODateTime) MarshalText() ([]byte, error) {
	return xsdDateTime(t).MarshalText()
}

type IssuerAndSerialNumber1 struct {
	Issr  CertificateIssuer1 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 Issr"`
	SrlNb Max35Binary        `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 SrlNb"`
}

type KEK4 struct {
	Vrsn          float64                   `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 Vrsn,omitempty"`
	KEKId         KEKIdentifier2            `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 KEKId"`
	KeyNcrptnAlgo AlgorithmIdentification13 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 KeyNcrptnAlgo"`
	NcrptdKey     Max500Binary              `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 NcrptdKey"`
}

type KEKIdentifier2 struct {
	KeyId     Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 KeyId"`
	KeyVrsn   Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 KeyVrsn"`
	SeqNb     float64         `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 SeqNb,omitempty"`
	DerivtnId Min5Max16Binary `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 DerivtnId,omitempty"`
}

type KeyTransport4 struct {
	Vrsn          float64                   `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 Vrsn,omitempty"`
	RcptId        Recipient5Choice          `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 RcptId"`
	KeyNcrptnAlgo AlgorithmIdentification11 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 KeyNcrptnAlgo"`
	NcrptdKey     Max5000Binary             `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 NcrptdKey"`
}

type Max100KBinary []byte

func (t *Max100KBinary) UnmarshalText(text []byte) error {
	return (*xsdBase64Binary)(t).UnmarshalText(text)
}
func (t Max100KBinary) MarshalText() ([]byte, error) {
	return xsdBase64Binary(t).MarshalText()
}

// Must be at least 1 items long
type Max104Text string

// Must be at least 1 items long
type Max10Text string

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
type Max16Text string

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
type Max37Text string

// Must match the pattern [0-9]{1,3}
type Max3NumericText string

// Must be at least 1 items long
type Max3Text string

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

// Must be at least 1 items long
type Max76Text string

// May be one of BALN, CMPA, CMPD, ACMD, DVCC, DIAQ, DIAP, GSTS, INQQ, INQP, KYAQ, KYAP, PINQ, PINP, RJAQ, RJAP, WITV, WITK, WITQ, WITP, INQC, H2AP, H2AQ, TMOP, CSEC, DSEC, SKSC, SSTS, DPSK, DPSV, DPSQ, DPSP, EXPK, EXPV, TRFQ, TRFP, RPTC
type MessageFunction11Code string

// Must match the pattern [a-zA-Z]{2,3}
type Min2Max3AlphaText string

// Must match the pattern [0-9]{2,3}
type Min2Max3NumericText string

type Min5Max16Binary []byte

func (t *Min5Max16Binary) UnmarshalText(text []byte) error {
	return (*xsdBase64Binary)(t).UnmarshalText(text)
}
func (t Min5Max16Binary) MarshalText() ([]byte, error) {
	return xsdBase64Binary(t).MarshalText()
}

// Must match the pattern [0-9]{8,28}
type Min8Max28NumericText string

// May be one of MREF, TEXT, HTML
type OutputFormat1Code string

type Parameter4 struct {
	NcrptnFrmt   EncryptionFormat1Code     `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 NcrptnFrmt,omitempty"`
	DgstAlgo     Algorithm11Code           `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 DgstAlgo,omitempty"`
	MskGnrtrAlgo AlgorithmIdentification12 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 MskGnrtrAlgo,omitempty"`
}

type Parameter5 struct {
	DgstAlgo Algorithm11Code `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 DgstAlgo,omitempty"`
}

type Parameter6 struct {
	NcrptnFrmt   EncryptionFormat1Code `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 NcrptnFrmt,omitempty"`
	InitlstnVctr Max500Binary          `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 InitlstnVctr,omitempty"`
	BPddg        BytePadding1Code      `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 BPddg,omitempty"`
}

type Parameter7 struct {
	InitlstnVctr Max500Binary     `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 InitlstnVctr,omitempty"`
	BPddg        BytePadding1Code `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 BPddg,omitempty"`
}

// May be one of ACQR, ATMG, CISP, DLIS, HSTG, ITAG, OATM
type PartyType12Code string

type PaymentCard23 struct {
	CardDataNtryMd CardDataReading1Code     `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 CardDataNtryMd"`
	FllbckInd      CardFallback1Code        `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 FllbckInd,omitempty"`
	PrtctdCardData ContentInformationType10 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 PrtctdCardData,omitempty"`
	PlainCardData  PlainCardData19          `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 PlainCardData,omitempty"`
	CardCtryCd     Max3Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 CardCtryCd,omitempty"`
	CardCcyCd      Exact3AlphaNumericText   `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 CardCcyCd,omitempty"`
	ElctrncPrsBal  CurrencyAndAmount        `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 ElctrncPrsBal,omitempty"`
}

type PlainCardData19 struct {
	PAN       Min8Max28NumericText `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 PAN,omitempty"`
	CardSeqNb Min2Max3NumericText  `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 CardSeqNb,omitempty"`
	FctvDt    Max10Text            `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 FctvDt,omitempty"`
	XpryDt    Max10Text            `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 XpryDt,omitempty"`
	Trck1     Max76Text            `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 Trck1,omitempty"`
	Trck2     Max37Text            `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 Trck2,omitempty"`
	Trck3     Max104Text           `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 Trck3,omitempty"`
}

type PointOfInteractionCapabilities7 struct {
	CardRdData       []CardDataReading4Code                  `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 CardRdData,omitempty"`
	CardWrtData      []CardDataReading4Code                  `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 CardWrtData,omitempty"`
	Authntcn         []CardholderVerificationCapability3Code `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 Authntcn,omitempty"`
	PINLngthCpblties float64                                 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 PINLngthCpblties,omitempty"`
	ApprvlCdLngth    float64                                 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 ApprvlCdLngth,omitempty"`
	MxScrptLngth     float64                                 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 MxScrptLngth,omitempty"`
	CardCaptrCpbl    bool                                    `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 CardCaptrCpbl,omitempty"`
	WdrwlMdia        []ATMMediaType1Code                     `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 WdrwlMdia,omitempty"`
	DpstdMdia        []ATMMediaType2Code                     `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 DpstdMdia,omitempty"`
	MsgCpblties      []DisplayCapabilities5                  `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 MsgCpblties,omitempty"`
}

type PostalAddress17 struct {
	AdrLine     []Max70Text               `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 AdrLine,omitempty"`
	StrtNm      Max70Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 StrtNm,omitempty"`
	BldgNb      Max16Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 BldgNb,omitempty"`
	PstCd       Max16Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 PstCd,omitempty"`
	TwnNm       Max35Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 TwnNm"`
	CtrySubDvsn []Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 CtrySubDvsn,omitempty"`
	Ctry        CountryCode               `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 Ctry"`
	GLctn       GeographicLocation1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 GLctn,omitempty"`
}

type Recipient4Choice struct {
	KeyTrnsprt KeyTransport4  `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 KeyTrnsprt"`
	KEK        KEK4           `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 KEK"`
	KeyIdr     KEKIdentifier2 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 KeyIdr"`
}

type Recipient5Choice struct {
	IssrAndSrlNb IssuerAndSerialNumber1 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 IssrAndSrlNb"`
	KeyIdr       KEKIdentifier2         `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 KeyIdr"`
}

type RelativeDistinguishedName1 struct {
	AttrTp  AttributeType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 AttrTp"`
	AttrVal Max140Text         `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 AttrVal"`
}

type TerminalHosting1 struct {
	Ctgy TransactionEnvironment3Code `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 Ctgy,omitempty"`
	Id   Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 Id,omitempty"`
}

type Traceability4 struct {
	RlayId      GenericIdentification77 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 RlayId"`
	SeqNb       Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 SeqNb,omitempty"`
	TracDtTmIn  ISODateTime             `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 TracDtTmIn"`
	TracDtTmOut ISODateTime             `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 TracDtTmOut"`
}

// May be one of PRIV, PUBL
type TransactionEnvironment2Code string

// May be one of BRCH, MERC, OTHR
type TransactionEnvironment3Code string

type TransactionIdentifier1 struct {
	TxDtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 TxDtTm"`
	TxRef  Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 TxRef"`
}

type TransactionVerificationResult5 struct {
	Mtd         AuthenticationMethod7Code `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 Mtd"`
	VrfctnNtty  AuthenticationEntity2Code `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 VrfctnNtty,omitempty"`
	Rslt        Verification1Code         `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 Rslt,omitempty"`
	AddtlRslt   Max500Text                `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 AddtlRslt,omitempty"`
	AuthntcnTkn Max140Binary              `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 AuthntcnTkn,omitempty"`
}

type UTMCoordinates1 struct {
	UTMZone    Max16Text `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 UTMZone"`
	UTMEstwrd  float64   `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 UTMEstwrd"`
	UTMNrthwrd float64   `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 UTMNrthwrd"`
}

// May be one of CDSP, CRCP, CRDO
type UserInterface5Code string

// May be one of FAIL, MISS, NOVF, PART, SUCC, ERRR
type Verification1Code string

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
