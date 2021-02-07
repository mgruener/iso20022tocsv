// Code generated by main. DO NOT EDIT.

package auth_058_001_01

type ActiveCurrencyAndAmount struct {
	Value float64            `xml:",chardata"`
	Ccy   ActiveCurrencyCode `xml:"Ccy,attr"`
}

// Must match the pattern [A-Z]{3,3}
type ActiveCurrencyCode string

type AmountAndDirection102 struct {
	Amt ActiveCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:auth.058.001.01 Amt"`
	Sgn bool                    `xml:"urn:iso:std:iso:20022:tech:xsd:auth.058.001.01 Sgn"`
}

type CCPPortfolioStressTestingResultReportV01 struct {
	ScnroStrssTstRslt []ScenarioStressTestResult1 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.058.001.01 ScnroStrssTstRslt"`
	SplmtryData       []SupplementaryData1        `xml:"urn:iso:std:iso:20022:tech:xsd:auth.058.001.01 SplmtryData,omitempty"`
}

type Document struct {
	CCPPrtflStrssTstgRsltRpt CCPPortfolioStressTestingResultReportV01 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.058.001.01 CCPPrtflStrssTstgRsltRpt"`
}

type GenericIdentification165 struct {
	Id      Max256Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:auth.058.001.01 Id"`
	Desc    Max140Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:auth.058.001.01 Desc,omitempty"`
	Issr    Max35Text                     `xml:"urn:iso:std:iso:20022:tech:xsd:auth.058.001.01 Issr,omitempty"`
	SchmeNm SchemeIdentificationType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:auth.058.001.01 SchmeNm,omitempty"`
}

type GenericIdentification168 struct {
	Id      Max256Text `xml:"urn:iso:std:iso:20022:tech:xsd:auth.058.001.01 Id"`
	Desc    Max140Text `xml:"urn:iso:std:iso:20022:tech:xsd:auth.058.001.01 Desc,omitempty"`
	Issr    Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:auth.058.001.01 Issr,omitempty"`
	SchmeNm Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:auth.058.001.01 SchmeNm,omitempty"`
}

// Must be at least 1 items long
type Max140Text string

// Must be at least 1 items long
type Max256Text string

// Must be at least 1 items long
type Max350Text string

// Must be at least 1 items long
type Max35Text string

type PortfolioStressTestResult1 struct {
	PrtflId      GenericIdentification165 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.058.001.01 PrtflId"`
	StrssLoss    AmountAndDirection102    `xml:"urn:iso:std:iso:20022:tech:xsd:auth.058.001.01 StrssLoss"`
	RawStrssLoss AmountAndDirection102    `xml:"urn:iso:std:iso:20022:tech:xsd:auth.058.001.01 RawStrssLoss,omitempty"`
	Cover1Flg    bool                     `xml:"urn:iso:std:iso:20022:tech:xsd:auth.058.001.01 Cover1Flg"`
	Cover2Flg    bool                     `xml:"urn:iso:std:iso:20022:tech:xsd:auth.058.001.01 Cover2Flg"`
}

type ScenarioStressTestResult1 struct {
	Id                GenericIdentification168     `xml:"urn:iso:std:iso:20022:tech:xsd:auth.058.001.01 Id"`
	PrtflStrssTstRslt []PortfolioStressTestResult1 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.058.001.01 PrtflStrssTstRslt"`
}

// May be one of MARG, COLL, POSI, CLIM
type SchemeIdentificationType1Code string

type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.058.001.01 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.058.001.01 Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
}