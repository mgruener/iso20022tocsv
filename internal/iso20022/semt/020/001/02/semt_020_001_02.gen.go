// Code generated by main. DO NOT EDIT.

package semt_020_001_02

// Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type AnyBICIdentifier string

// May be one of FREE, APMT
type DeliveryReceiptType2Code string

type Document struct {
	SctiesMsgCxlAdvc SecuritiesMessageCancellationAdviceV02 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.020.001.02 SctiesMsgCxlAdvc"`
}

// Must match the pattern [a-zA-Z0-9]{4}
type Exact4AlphaNumericText string

type GenericIdentification19 struct {
	Id      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:semt.020.001.02 Id"`
	Issr    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:semt.020.001.02 Issr"`
	SchmeNm Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:semt.020.001.02 SchmeNm,omitempty"`
}

type GenericIdentification20 struct {
	Id      Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:semt.020.001.02 Id"`
	Issr    Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:semt.020.001.02 Issr"`
	SchmeNm Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:semt.020.001.02 SchmeNm,omitempty"`
}

// Must be at least 1 items long
type Max350Text string

// Must be at least 1 items long
type Max35Text string

// Must be at least 1 items long
type Max70Text string

type PartyIdentification36Choice struct {
	AnyBIC  AnyBICIdentifier        `xml:"urn:iso:std:iso:20022:tech:xsd:semt.020.001.02 AnyBIC"`
	PrtryId GenericIdentification19 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.020.001.02 PrtryId"`
}

// May be one of DELI, RECE
type ReceiveDelivery1Code string

type References18Choice struct {
	SctiesSttlmTxConfId           SettlementTypeAndIdentification13 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.020.001.02 SctiesSttlmTxConfId"`
	IntraPosMvmntConfId           Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:semt.020.001.02 IntraPosMvmntConfId"`
	SctiesBalAcctgRptId           Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:semt.020.001.02 SctiesBalAcctgRptId"`
	SctiesBalCtdyRptId            Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:semt.020.001.02 SctiesBalCtdyRptId"`
	IntraPosMvmntPstngRptId       Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:semt.020.001.02 IntraPosMvmntPstngRptId"`
	SctiesFincgConfId             SettlementTypeAndIdentification13 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.020.001.02 SctiesFincgConfId"`
	SctiesTxPdgRptId              Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:semt.020.001.02 SctiesTxPdgRptId"`
	SctiesTxPstngRptId            Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:semt.020.001.02 SctiesTxPstngRptId"`
	SctiesSttlmTxAllgmtRptId      Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:semt.020.001.02 SctiesSttlmTxAllgmtRptId"`
	SctiesSttlmTxAllgmtNtfctnTxId SettlementTypeAndIdentification13 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.020.001.02 SctiesSttlmTxAllgmtNtfctnTxId"`
	PrtflTrfNtfctnId              Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:semt.020.001.02 PrtflTrfNtfctnId"`
	SctiesSttlmTxGnrtnNtfctnId    SettlementTypeAndIdentification13 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.020.001.02 SctiesSttlmTxGnrtnNtfctnId"`
	OthrMsgId                     Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:semt.020.001.02 OthrMsgId"`
}

type SecuritiesAccount13 struct {
	Id Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:semt.020.001.02 Id"`
	Tp GenericIdentification20 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.020.001.02 Tp,omitempty"`
	Nm Max70Text               `xml:"urn:iso:std:iso:20022:tech:xsd:semt.020.001.02 Nm,omitempty"`
}

type SecuritiesMessageCancellationAdviceV02 struct {
	Ref         References18Choice          `xml:"urn:iso:std:iso:20022:tech:xsd:semt.020.001.02 Ref"`
	AcctOwnr    PartyIdentification36Choice `xml:"urn:iso:std:iso:20022:tech:xsd:semt.020.001.02 AcctOwnr,omitempty"`
	SfkpgAcct   SecuritiesAccount13         `xml:"urn:iso:std:iso:20022:tech:xsd:semt.020.001.02 SfkpgAcct"`
	SplmtryData []SupplementaryData1        `xml:"urn:iso:std:iso:20022:tech:xsd:semt.020.001.02 SplmtryData,omitempty"`
}

type SettlementTypeAndIdentification13 struct {
	TxId          Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:semt.020.001.02 TxId"`
	SctiesMvmntTp ReceiveDelivery1Code     `xml:"urn:iso:std:iso:20022:tech:xsd:semt.020.001.02 SctiesMvmntTp"`
	Pmt           DeliveryReceiptType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:semt.020.001.02 Pmt"`
}

type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.020.001.02 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.020.001.02 Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
}
