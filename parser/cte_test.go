package parser_test

import (
	"strings"
	"testing"
	"valdemarceccon/cte_parser/parser"
	"valdemarceccon/cte_parser/parser/types"
)

const xmlTest string = `<?xml version="1.0" encoding="UTF-8"?>
<cteProc xmlns="http://www.portalfiscal.inf.br/cte" versao="1.04">
    <CTe xmlns="http://www.portalfiscal.inf.br/cte">
        <infCte versao="1.04" Id="CTe35110899999999000191570010000001011000001018">
            <ide>
                <cUF>35</cUF>
                <cCT>00000101</cCT>
                <CFOP>5353</CFOP>
                <natOp>PRESTACAO DE SERVICO</natOp>
                <forPag>0</forPag>
                <mod>57</mod>
                <serie>1</serie>
                <nCT>101</nCT>
                <dhEmi>2011-08-20T18:24:19</dhEmi>
                <tpImp>1</tpImp>
                <tpEmis>1</tpEmis>
                <cDV>8</cDV>
                <tpAmb>2</tpAmb>
                <tpCTe>0</tpCTe>
                <procEmi>0</procEmi>
                <verProc>4.0</verProc>
                <cMunEnv>3503208</cMunEnv>
                <xMunEnv>ARARAQUARA</xMunEnv>
                <UFEnv>SP</UFEnv>
                <modal>01</modal>
                <tpServ>0</tpServ>
                <cMunIni>3505708</cMunIni>
                <xMunIni>BARUERI</xMunIni>
                <UFIni>SP</UFIni>
                <cMunFim>3503208</cMunFim>
                <xMunFim>ARARAQUARA</xMunFim>
                <UFFim>SP</UFFim>
                <retira>0</retira>
                <toma03>
                    <toma>0</toma>
                </toma03>
            </ide>
            <compl>
                <xEmi>ITALO</xEmi>
                <Entrega>
                    <semData>
                        <tpPer>0</tpPer>
                    </semData>
                    <semHora>
                        <tpHor>0</tpHor>
                    </semHora>
                </Entrega>
                <origCalc>BARUERI</origCalc>
                <destCalc>ARARAQUARA</destCalc>
                <xObs>;Documento emitido por ME ou EPP optante pelo Simples Nacional;Nao gera direito a credito fiscal de ICMS</xObs>
            </compl>
            <emit>
                <CNPJ>99999999000191</CNPJ>
                <IE>181288960178</IE>
                <xNome>TESTANDO TESTANDO LTDA-ME</xNome>
                <xFant>GMP</xFant>
                <enderEmit>
                    <xLgr>RUA EXPEDICIONARIO</xLgr>
                    <nro>123</nro>
                    <xBairro>VILA YAMADA</xBairro>
                    <cMun>3503208</cMun>
                    <xMun>ARARAQUARA</xMun>
                    <CEP>14802150</CEP>
                    <UF>SP</UF>
                    <fone>1234567891</fone>
                </enderEmit>
            </emit>
            <rem>
                <CNPJ>61365284015136</CNPJ>
                <IE>206274823111</IE>
                <xNome>TESTANDO TESTANDO LTDA-ME</xNome>
                <xFant>TESTANDO TESTANDO LTDA-ME</xFant>
                <fone>1234567891</fone>
                <enderReme>
                    <xLgr>AV. PROJETADA</xLgr>
                    <nro>751</nro>
                    <xBairro>NOVA ALVORADA</xBairro>
                    <cMun>3505708</cMun>
                    <xMun>BARUERI</xMun>
                    <CEP>06463400</CEP>
                    <UF>SP</UF>
                    <cPais>1058</cPais>
                    <xPais>BRASIL</xPais>
                </enderReme>
                <infNFe>
                    <chave>35110561375284015136550140028647541589947329</chave>
                </infNFe>
            </rem>
            <dest>
                <CPF>06760213874</CPF>
                <IE>ISENTO</IE>
                <xNome>TESTANDO TESTANDO LTDA-ME</xNome>
                <fone>1633312039</fone>
                <enderDest>
                    <xLgr>AL. FERRAZ</xLgr>
                    <nro>106</nro>
                    <xBairro>FONTE</xBairro>
                    <cMun>3503208</cMun>
                    <xMun>ARARAQUARA</xMun>
                    <CEP>14802428</CEP>
                    <UF>SP</UF>
                    <cPais>1058</cPais>
                    <xPais>BRASIL</xPais>
                </enderDest>
            </dest>
            <vPrest>
                <vTPrest>2.00</vTPrest>
                <vRec>2.00</vRec>
                <Comp>
                    <xNome>FRETE PESO</xNome>
                    <vComp>2.00</vComp>
                </Comp>
            </vPrest>
            <imp>
                <ICMS>
                    <ICMS45>
                        <CST>41</CST>
                    </ICMS45>
                </ICMS>
            </imp>
            <infCTeNorm>
                <infCarga>
                    <vCarga>88.30</vCarga>
                    <proPred>CAIXA</proPred>
                    <xOutCat>DVD</xOutCat>
                    <infQ>
                        <cUnid>01</cUnid>
                        <tpMed>Kg</tpMed>
                        <qCarga>2.0000</qCarga>
                    </infQ>
                    <infQ>
                        <cUnid>03</cUnid>
                        <tpMed>CAIXA</tpMed>
                        <qCarga>1.0000</qCarga>
                    </infQ>
                </infCarga>
                <infModal versaoModal="1.04">
                    <rodo>
                        <RNTRC>06013631</RNTRC>
                        <dPrev>2011-08-14</dPrev>
                        <lota>0</lota>
                    </rodo>
                </infModal>
            </infCTeNorm>
        </infCte>
        <Signature xmlns="http://www.w3.org/2000/09/xmldsig#">
            <SignedInfo>
                <CanonicalizationMethod Algorithm="http://www.w3.org/TR/2001/REC-xml-c14n-20010315"/>
                <SignatureMethod Algorithm="http://www.w3.org/2000/09/xmldsig#rsa-sha1"/>
                <Reference URI="#CTe35110899999999000191570010000001011000001018">
                    <Transforms>
                        <Transform Algorithm="http://www.w3.org/2000/09/xmldsig#enveloped-signature"/>
                        <Transform Algorithm="http://www.w3.org/TR/2001/REC-xml-c14n-20010315"/>
                    </Transforms>
                    <DigestMethod Algorithm="http://www.w3.org/2000/09/xmldsig#sha1"/>
                    <DigestValue>cGkh9p2OT/dDA3jTw/4DuLfzdCd=</DigestValue>
                </Reference>
            </SignedInfo>
            <SignatureValue>jfGyueaAABym1GmPMn6IGKRj99fMKwu39BtaxtUruL6ssEcaKQfo8t3qbpB9rnM8sh3iaTaf0DBP5bMjotL/2wsUXtot7vWZi+hzJIMX0Gq0d179xVK85Ey1ohkvEYt3NYSO7GygpMqTJdKvKG1uKuvlnERSjlI5aEd/9V0yZ00=</SignatureValue>
            <KeyInfo>
                <X509Data>
                    <X509Certificate>MIIFtDCCBJygAwIBAgIIG1mv0XRoWy4wDQYJKoZIhvcNAQEFBQAwXzELMAkGA1UEBhMCQlIxEzARBgNVBAoMCklDUC1CcmFzaWwxIDAeBgNVBAsMF0NhaXhhIEVjb25vbWljYSBGZWRlcmFsMRkwFwYDVQQDDBBBQyBDQUlYQSBQSi0xIHYxMB4XDTExMTExMTA2MDg0NloXDTEyMTExMDA2MDg0NlowgY4xCzAJBgNVBAYTAkJSMRMwEQYDVQQKDApJQ1AtQnJhc2lsMSAwHgYDVQQLDBdDYWl4YSBFY29ub21pY2EgRmVkZXJhbDEZMBcGA1UECwwQQUMgQ0FJWEEgUEotMSBWMTEtMCsGA1UEAwwkTSBSIE0gS0FUTyBBU0FLVVJBIC0gRVBQOjY5NjIxMTg3OTE1MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQC0751AtgqX3lAPZpQTIQWWSpXuoEN+JYiJrA99Toz2oLfRPjPPINj8J7IFHQ0sT3KzhNIykVxYAL22tpX9dvnUMs1mm0C/O+ImEzGr7xqSFVd7sp9qsdrc/JWEUaRVqKUwZSd38HWBEAwPNFBa3/UN1EhXgjev+pH4oKsu1msthwIDAQABo4ICxjCCAsIwDgYDVR0PAQH/BAQDAgXgMCkGA1UdJQQiMCAGCCsGAQUFBwMCBggrBgEFBQcDBAYKKwYBBAGCNxQCAjAdBgNVHQ4EFgQU22ocgo1bv7EO0N4T6wMl9l+cCV0wHwYDVR0jBBgwFoAUF8RVhQ+bjMcBI4CdEbNH3LOT1qswgcQGA1UdEQSBvDCBuYEWbXJlZ2luYS5rYXRvQGdtYWlsLmNvbaAXBgVgTAEDB6AOBAwwMDAwMDAwMDAwMDCgPQYFYEwBAwSgNAQyMTQwOTE5NzI2OTYyMTE4NzkxNTE2ODY1MDA4MjEwMDAwMDAwMDAwMDAwMDAwU1NQU1CgGQYFYEwBAwOgEAQOMTAxNDI3ODUwMDAxOTCgLAYFYEwBAwKgIwQhTUFSQ0lBIFJFR0lOQSBNSVlVS0kgS0FUTyBBU0FLVVJBMGYGA1UdIARfMF0wWwYGYEwBAgEJMFEwTwYIKwYBBQUHAgEWQ2h0dHA6Ly9jZXJ0aWZpY2Fkb2RpZ2l0YWwuY2FpeGEuZ292LmJyL2RvY3VtZW50b3MvZHBjYWMtY2FpeGFwai5wZGYwgbwGA1UdHwSBtDCBsTAuoCygKoYoaHR0cDovL2xjci5jYWl4YS5nb3YuYnIvYWNjYWl4YXBqMXYxLmNybDAvoC2gK4YpaHR0cDovL2xjcjIuY2FpeGEuZ292LmJyL2FjY2FpeGFwajF2MS5jcmwwTqBMoEqGSGh0dHA6Ly9yZXBvc2l0b3Jpby5pY3BicmFzaWwuZ292LmJyL2xjci9DQUlYQS9BQ0NBSVhBUEovYWNjYWl4YXBqMXYxLmNybDBXBggrBgEFBQcBAQRLMEkwRwYIKwYBBQUHMAKGO2h0dHA6Ly9jZXJ0aWZpY2Fkb2RpZ2l0YWwuY2FpeGEuZ292LmJyL2FpYS9hY2NhaXhhcGoxdjEucDdiMA0GCSqGSIb3DQEBBQUAA4IBAQB7lxQ/YPm3Tl4CqTVuPmo/ipWlM7mqMUIh4if2shlXXlZwzotNLdacsr1nh9wFLETVSn6T4gllo42aSgOgaYmzhzeITKsEw4NdMq4RkCKytuxbIqAPmTD8gZEV144AK6aB8sThv8LpZ+D002LTRG9t013C3DbiGeYD3Cdu1gxRDHOW1lWula+9e5tPZ8Bbh/BDR5cSFUWBzt3ellQ4BJY+xh2a2iteM9+KAHsZChIPrHPRV/LG5HSlpxlAtkGSxZDK5EIEYJ7SVpWovV1oUnz8PvKO2nQQKflTLVSE9C+NbfjXfD6f5GtED6q4AxkHNUQBimd6HAhA0oDjO/oaDeEC</X509Certificate>
                </X509Data>
            </KeyInfo>
        </Signature>
    </CTe>
    <protCTe versao="1.04" xmlns="http://www.portalfiscal.inf.br/cte">
        <infProt>
            <tpAmb>2</tpAmb>
            <verAplic>SP_PL_CTe_104</verAplic>
            <chCTe>35110899999999000191570010000001011000001018</chCTe>
            <dhRecbto>2011-08-20T18:24:22</dhRecbto>
            <nProt>135110002038337</nProt>
            <digVal>nBkh9p2OT/dDA3jTw/4DuLfzdZM=</digVal>
            <cStat>100</cStat>
            <xMotivo>Autorizado o uso do CT-e</xMotivo>
        </infProt>
    </protCTe>
</cteProc>`

func TestParserExampleCTE(t *testing.T) {
	procCte, err := parser.Parse(strings.NewReader(xmlTest))
	if err != nil {
		t.Fatal(err)
	}
	t.Run("given 1.04 versao procCte, want 1.04 versao", func(t *testing.T) {
		want := types.TVersaoCTe("1.04")

		if procCte.Versao != want {
			t.Errorf("Estava esperando a versão %s mas recebeu %s", want, procCte.Versao)
		}
	})

	t.Run("given protCte parse correctly", func(t *testing.T) {
		want := types.TpAmb("2")

		if procCte.ProtCTe.InfProt.TpAmb != want {
			t.Errorf("Estava esperando a versão %s mas recebeu %s", want, procCte.ProtCTe.InfProt.TpAmb)
		}
	})

}
