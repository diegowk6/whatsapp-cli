# KB — ERPs Integrados com a Wake Commerce

> Artigo canônico · Extraído de grupos internos WhatsApp · Atualizado: 2026-06-05
>
> **Tags**: `erp`, `integrador-erp`, `middleware`, `ipaas`, `totvs`, `sap`, `bling`, `tiny`, `audience-solutions`, `audience-sales`

---

## Resumo

A Wake Commerce possui uma **API aberta (REST/GraphQL)**. Não existem conectores ERP nativos embutidos na plataforma para a grande maioria dos ERPs do mercado. Toda integração entre um ERP e a Wake requer um **parceiro integrador/middleware** (iPaaS ou consultoria especializada) que construa ou opere o conector.

A única exceção documentada é o **ERP Tiny**, cujos IDs de grupo de pagamento precisam ser mapeados diretamente pelo cliente junto ao Tiny após a criação dos grupos no painel Wake — confirmando que há integração em uso (caso Karcher), mas sem conector oficial gerenciado pela Wake.

O ecossistema de parceiros integradores é robusto: 10 parceiros certificados cobrem os principais ERPs do mercado brasileiro. A escolha do parceiro correto depende do ERP do cliente, da complexidade do projeto (B2B, multiCD, marketplace) e do histórico do integrador com aquele ERP específico.

---

## Matriz de ERPs

### 1. ERPs com integração em uso (via middleware/parceiro)

Estes ERPs aparecem confirmados em conversas técnicas e comerciais com clientes reais da Wake:

| ERP | Parceiros Confirmados | Status | Last verified | Observações |
|---|---|---|---|---|
| **TOTVS Protheus** | EBI, Integra.do, Intergate, Floui | ✓ verificado | 2026-06-02 | Múltiplos parceiros especializados; lista reconfirmada em mensagem de Quick Digital & Wake Parcerias |
| **TOTVS Winthor** | EBI, Integra.do, Vtrina | ✓ verificado | 2026-06-02 | Vtrina confirmou integração explicitamente |
| **TOTVS RM** | Intergate (especialista) | ⚠ inferido | 2026-06-02 | Intergate é o mais especializado. **Integra.do lista TOTVS RM no site mas o closer da Integra.do confirmou que a integração não está pronta** (mensagem WAKE \| Time Parcerias, 2026-06-02 — Arnaldo da Savvi). Não indicar Integra.do para RM até nova confirmação |
| **SAP Business One** | Intergate, EBI, LinkToU, Iota.Hub | ✓ verificado | 2026-06-02 | Intergate tem salesAPI proprietária para SAP B1 |
| **SAP S/4HANA** | Floui | ✓ verificado | 2026-06-02 | Floui é referência para S/4HANA (low-code) |
| **SAP (geral)** | Integra.do, Iota.Hub, EBI, LinkToU, Jitterbit | ✓ verificado | 2026-06-02 | Múltiplos parceiros |
| **Linx Microvix** | EBI, Vtrina, Argento Digital | ✓ verificado | 2026-06-02 | Linx foi adquirida pela TOTVS |
| **Linx (geral)** | LinkToU, Vtrina, Argento Digital | ✓ verificado | 2026-06-02 | Mesma observação acima |
| **Omie** | Integra.do, Argento Digital | ✓ verificado | 2026-06-02 | |
| **Sankhya** | Integra.do, Floui | ✓ verificado | 2026-06-02 | |
| **Bling** | Nativo Wake; Integra.do e Vtrina para cenários complexos | ✓ verificado | 2026-06-02 | Bling é ERP nativo Wake. Vtrina cobre cenário multiCD com Bling + Correios Log+ |
| **Tiny** | Nativo Wake; Integra.do para cenários complexos | ✓ verificado | 2026-04-09 | Tiny é ERP nativo Wake (sem necessidade de integrador para cenários padrão). Adoção confirmada em vários leads recentes: Karcher (mar/2026), Empório do Celular (abr/2026), Wallford Brasil (dez/2025), Avocado Tech Wallford (dez/2025). Casos com middleware (ex: Karcher + Leanwork) exigem remapeamento manual de IDs de grupo de pagamento no Tiny pelo cliente |
| **Senior** | Argento Digital | ⚠ inferido | — | Sem mensagem recente confirmando — validar com Partner Ops antes de indicar |
| **Abacos** | Argento Digital | ✓ verificado | 2026-06-02 | Listado explicitamente entre os ERPs cobertos pela Argento na mensagem do ecossistema (WAKE \| Time Parcerias, 2026-06-02) |
| **Alpha7** | Vtrina | ⚠ inferido | — | "Pode ser viabilizado" — confirmar com Vtrina antes de indicar |
| **Procfit** | Vtrina | ⚠ inferido | — | "Pode ser viabilizado" — confirmar com Vtrina antes de indicar |

### 2. ERPs com integração via parceiro iPaaS global

Parceiros iPaaS (Jitterbit, Iota.Hub, Floui, LinkToU) possuem conectores pré-construídos e suportam virtualmente qualquer ERP com API disponível. Para ERPs não listados acima, consultar estes parceiros diretamente.

### 3. ERPs com status incerto (necessita validação)

| ERP | Status | Last verified | Ação recomendada |
|---|---|---|---|
| **Everest** | ⚠ validação pendente via Partner Ops | 2025-10-10 | Mencionado como "ERP EVEREST" em convenção de parceiro (Parcerias \| Canais, 2025-10-10) e como dúvida de homologação (WAKE \| Time Parcerias, 2025-07-16: "ERP chamado de Everest e a princípio ele estaria integrado (homologado) com a Wake. Sabe se realmente esta?" — pergunta sem resposta registrada). Nenhuma confirmação de integração ativa encontrada. **Não prometer integração** — escalar para Partner Ops antes de indicar |
| **Vetor** | ❌ sem fonte (remover do pipeline de indicações) | — | Nenhuma mensagem no corpus confirma a existência de "ERP Vetor" como produto homologável. Única menção próxima é "Vetor Editora" (cliente B2C, Eficaz <> WAKE, 2026-05-29) — referência de marca, não ERP. Tratar como reivindicação não substanciada até que Partner Ops apresente evidência |

### 3a. Categoria corrigida — Proex não é ERP

**Correção factual** (validada 2026-06-02): "Proex" / "ProEX Digital" é uma **agência parceira** da Wake (selo Special / Elite em diferentes momentos), não um ERP. Aparece consistentemente em listas de agências (WAKE \| Time Agência, 2025-08-12, 2025-12-19; Comercial Cross WAKE, 2025-08-18) e como agência que indicou clientes para a Wake (Coder Ivy + Wake, 2026-05-19). A classificação anterior deste artigo como "ERP que requer intermediador" estava incorreta e foi removida da matriz. Para a entrada da Proex no catálogo de agências, consultar `docs/kb-priority-parceiros-wake.md`.

### 4. ERPs sem integração conhecida

Nenhum ERP foi explicitamente listado como "não suportado" ou "bloqueado" nas fontes consultadas. A Wake, por ter API aberta, não bloqueia nenhum ERP tecnicamente — qualquer ERP que exponha uma API pode ser integrado via parceiro.

---

## Parceiros de Integração — Especialidades

### Argento Digital
- **Especialidades:** Linx (Microvix), Senior, Omie, Abacos, WMS
- **Perfil:** Middleware/hub; forte em ERPs de varejo e WMS
- **Indicado para:** Clientes com Senior, Linx ou Abacos

### Integra.do (Kore Hub)
- **Especialidades:** TOTVS (Protheus, Winthor, RM), SAP, Omie, Sankhya, Bling, Tiny
- **Perfil:** Parceiro de ampla cobertura — maior diversidade de ERPs homologados
- **Atenção:** TOTVS RM listado no site, mas closer da Integra.do confirmou (via Savvi, 2026-06-02) que **a integração não está pronta**. Não indicar Integra.do para RM até nova confirmação

### Iota.Hub
- **Especialidades:** SAP, sistemas legados, jornadas B2B/B2C
- **Perfil:** iPaaS focado em SAP e complexidade B2B
- **Indicado para:** Clientes com SAP (qualquer versão) e operações B2B complexas

### EBI (EBInterchange)
- **Especialidades:** TOTVS Protheus, TOTVS Winthor, Linx Microvix, SAP
- **Perfil:** 15+ anos de experiência em integrações de e-commerce com ERPs
- **Indicado para:** Clientes com TOTVS ou Linx que precisam de parceiro experiente

### LinkToU
- **Especialidades:** SAP, TOTVS, Linx
- **Perfil:** Conectores pré-construídos com mapeamento visual
- **Indicado para:** Projetos que precisam de velocidade de implantação com SAP/TOTVS

### Intergate Consultoria
- **Especialidades:** SAP Business One, TOTVS Protheus (especialistas), TOTVS RM
- **Perfil:** Consultoria especializada com salesAPI proprietária para SAP B1; é a parceira mais especializada para TOTVS RM
- **Indicado para:** Clientes com SAP B1 ou TOTVS RM que exigem profundidade técnica

### Floui
- **Especialidades:** SAP S/4HANA, TOTVS Protheus, Sankhya
- **Perfil:** Plataforma low-code de integração; forte em SAP S/4HANA
- **Indicado para:** Clientes com SAP S/4HANA (Floui é a principal referência neste ERP)

### Jitterbit
- **Especialidades:** iPaaS global com conectores pré-construídos e IA
- **Perfil:** Líder global em iPaaS; suporta cloud e on-premise; cobre qualquer ERP com API
- **Indicado para:** Empresas enterprise com ERPs internacionais ou complexidade alta

### Vtrina
- **Especialidades:** Hub de marketplace + middleware para integrações complexas de ERP
- **ERPs confirmados:** Winthor (TOTVS), Linx (geral), Bling; Alpha7 e Procfit podem ser viabilizados; Vetor necessita avaliação
- **Perfil:** Forte em operações com múltiplos marketplaces + ERP; integração padrão R$ 14.300 / ~80 dias; integração complexa R$ 14.700
- **Indicado para:** Clientes que precisam de hub de marketplace integrado ao ERP

### Leanwork Group
- **Especialidades:** Software house — conectores customizados para e-commerce digital
- **Perfil:** Atua como laboratório interno da Wake para novas integrações; gerencia o onboarding de novos parceiros tecnológicos
- **Indicado para:** Projetos que não têm parceiro disponível no mercado; integrações sob demanda

---

## Como Escolher o Parceiro de Integração Certo

### 1. Identifique o ERP do cliente
O ERP atual (ou novo) do cliente é o principal critério de seleção. Consulte a matriz acima para ver quais parceiros têm aquele ERP homologado.

### 2. Avalie a complexidade da operação
- **Operação simples (B2C, 1 loja, sem marketplace):** Qualquer parceiro da lista que cubra o ERP
- **Operação com múltiplos marketplaces + ERP:** Vtrina ou Jitterbit
- **Operação B2B complexa ou SAP enterprise:** Iota.Hub, Floui, Intergate
- **ERP não listado ou legado:** Jitterbit (iPaaS global) ou Leanwork (sob demanda)

### 3. Considere o histórico do integrador
Para ERPs críticos, priorize o parceiro mais especializado:
- **TOTVS RM:** Intergate é a opção recomendada (Integra.do lista no site mas **confirmou em 2026-06-02 que a integração não está pronta**)
- **SAP S/4HANA:** Floui
- **SAP B1:** Intergate
- **TOTVS Protheus/Winthor:** EBI (15+ anos), Integra.do, Intergate ou Floui
- **Linx Microvix:** EBI ou Vtrina
- **ERPs SMB (Bling, Tiny, Omie):** Integra.do

### 4. Use o processo oficial para indicações
Ao indicar um integrador para um cliente, utilize o **formulário oficial de indicação de parceiros**:
`https://docs.google.com/forms/d/e/1FAIpQLSfDs4kbyGpLUZTX18Y1Y0VTH_akMpX5Y_mHbLvB9V0toO9SMw/viewform`

Os Partner Ops fazem a ponte formal com o integrador. Indicações feitas por outros meios não contam para campanhas de premiação.

### 5. Lembre-se: a Wake tem API aberta
Não existe bloqueio técnico para nenhum ERP. Se o ERP do cliente não estiver na lista acima, o caminho é: (a) verificar se o parceiro iPaaS cobre aquele ERP, ou (b) acionar a Leanwork para avaliação de desenvolvimento sob demanda.

---

## Perguntas Frequentes sobre ERP

**A Wake tem integração nativa com algum ERP?**
Sim, dois: **Bling e Tiny** são ERPs nativos da Wake (cenários padrão dispensam integrador). Todos os demais ERPs passam por parceiro integrador ou middleware desenvolvido pelo cliente/agência. (Verificado em `wake-partnership-kb-playbook.md` §2 e em casos confirmados — vide matriz.)

**Quando o cliente fala em "agência" para integração de ERP, é a mesma coisa que "integrador"?**
Não necessariamente. O time Wake usa às vezes o termo "agência" quando quer dizer "integrador". Para integrações de ERP, o perfil correto é um parceiro integrador/consultoria técnica especializada — não uma agência de e-commerce. Confirmar com o cliente o que eles precisam (frontend da loja vs. integração técnica de ERP).

**"Proex" é um ERP?**
Não. ProEX Digital é uma **agência parceira** da Wake (selo Special/Elite em diferentes ciclos), não um ERP. A confusão veio de uma entrada anterior deste artigo que foi corrigida em 2026-06-05. Para integração de ERP, consultar a matriz acima; para o perfil da Proex como agência, consultar `kb-priority-parceiros-wake.md`.

**O ERP Tiny já foi integrado com a Wake em projetos reais?**
Sim, em vários. Como Tiny é nativo da Wake, cenários padrão não exigem integrador. Casos confirmados em leads/projetos recentes: Karcher (caso Hero, mar/2026), Empório do Celular (abr/2026), Wallford Brasil (dez/2025), e o cliente da Avocado Tech para Wallford (dez/2025). No caso Karcher, por usar middleware Leanwork para seguro, os IDs de grupos de pagamento criados no painel Wake precisam ser remapeados no Tiny pelo cliente (a agência não tem credenciais do ERP).

**O ERP Everest funciona com a Wake?**
**Validação pendente.** A única evidência no corpus é uma pergunta sem resposta em WAKE \| Time Parcerias (2025-07-16) e uma menção a um "ERP EVEREST" em convenção de parceiro (Parcerias \| Canais, 2025-10-10). Não há mensagem confirmando integração ativa. Escalar para Partner Ops antes de prometer ao cliente.

**Existe um ERP chamado "Vetor" integrado com a Wake?**
Não há evidência no corpus. A menção a "Vetor Editora" (Eficaz <> WAKE, 2026-05-29) refere-se a um cliente B2C usado como referência de marca, não a um ERP. Tratar reivindicações de "ERP Vetor" como não substanciadas até Partner Ops apresentar prova.

**A Vtrina tem integração com o Winthor?**
Sim, confirmado. A Vtrina também integra com Linx; Alpha7 e Procfit podem ser viabilizados; Vetor necessita avaliação prévia.

**Qual é a diferença entre um parceiro iPaaS e uma consultoria de integração?**
- **iPaaS (Jitterbit, Iota.Hub, Floui, LinkToU):** Plataforma de integração como serviço. Tem conectores pré-construídos, interface visual, suporte a múltiplos ERPs. Tende a ser mais rápido para ERPs já homologados.
- **Consultoria de integração (EBI, Intergate, Argento Digital):** Desenvolve e opera a integração customizada. Pode ter mais profundidade técnica para ERPs específicos.

**O ERP do cliente está no scope de pricing de projetos de agência?**
Sim. Conforme registrado em briefings de leads, o ERP integrado vs. novo é um fator de precificação em projetos de implantação. ERPs mais complexos (TOTVS, SAP) encarecem o projeto.

**A Wake pode fazer B2B e B2C com o mesmo ERP integrado?**
Sim. A Wake suporta B2B e B2C no mesmo ambiente, com uma única integração de ERP cobrindo ambos — tabelas de preço, produtos, banners e logos distintos, tudo gerenciado em um único admin.

---

## Casos Conhecidos e Lições Aprendidas

### Caso Karcher — ERP Tiny com Wake + Hero Seguros (Leanwork)
- **ERP:** Tiny
- **Contexto:** Loja Karcher implementou o seguro Hero (Parcela Garantida) via Wake + middleware Leanwork
- **Lição sobre ERP:** Ao configurar novos grupos de pagamento no checkout Wake (ex: "GTW Cartão com Seguro"), os IDs precisam ser mapeados no Tiny pelo cliente. Ao trocar de conector de teste para conector real (Vindi), um novo ID de grupo é gerado — o remapeamento no Tiny precisa ser refeito
- **Processo:** A solicitação de mapeamento de IDs é feita pelo próprio cliente ao ERP, pois a agência não tem credenciais de acesso ao ERP do cliente

### Recomendação TOTVS RM — Intergate é a opção
- A Intergate é a parceira mais especializada para TOTVS RM
- **2026-06-02 (WAKE \| Time Parcerias)**: Arnaldo da Savvi reportou que "apesar da Integra.do ter a Totvs RM no site, o closer deles me disse que eles não têm a integração pronta". Não indicar Integra.do para RM até nova confirmação
- Para clientes RM, priorizar Intergate e validar capacidade no momento da indicação

### Correção: Proex não é ERP
- Entradas anteriores deste artigo classificavam "Proex" como ERP que requer intermediador. Está incorreto.
- ProEX Digital é agência parceira (Special/Elite), confirmada em listas oficiais do time de agências (WAKE \| Time Agência, 2025-08-12, 2025-12-19; Comercial Cross WAKE, 2025-08-18)
- A linha foi removida da matriz de ERPs em 2026-06-05

### Vtrina + Winthor — Confirmação técnica
- FAQ registrado em 2025: "A Vtrina tem integração com Winthor (ERP da TOTVS)? Sim."
- A Vtrina suporta cenário multiCD com Bling + Correios Log+ (com qualificação prévia)
- Preços de referência da Vtrina: integração padrão R$ 14.300 / ~80 dias; integração complexa R$ 14.700

### Indicação de Vtrina para cliente com Winthor (Comercial Cross)
- Em conversas do time comercial, a Vtrina foi apontada explicitamente como opção válida para clientes que precisam de integração com Winthor (TOTVS)
- O processo correto de indicação é via formulário oficial para os Partner Ops formalizarem

### Leanwork como laboratório de novas integrações
- A Leanwork atua como tech lab interno da Wake para desenvolvimento de novos conectores sob demanda
- Gerencia o onboarding de novos parceiros tecnológicos para a plataforma
- Exemplos de projetos executados: integração Wake + Hero Seguros (seguros no checkout) e Wake + Serasa Pass (identidade/biometria) — ambos sem relação com ERP, mas confirmam o modelo de middleware gerenciado pela Leanwork

---

## Fontes

| Arquivo | Conteúdo relevante |
|---|---|
| `kb-parcerias.md` | FAQ-012: matriz completa de parceiros integradores de ERP; TRB-007: ERP Everest (status incerto) |
| `kb-parceiros-tech.md` | HOW_TO-6: processo de onboarding de novos parceiros tech via Leanwork |
| `kb-anymarket-wake.md` | Integração NF-e/XML via ERP: endpoint `v2/orders/{id}/nfe`; campo `NeedInvoicedXml` |
| `kb-parceiros-wave5a.md` | BEST_PRACTICES-3: matriz de parceiros com nota sobre TOTVS RM; BEST_PRACTICES-5: ERP em briefings de leads |
| `kb-parceiros-wave5c.md` | FAQ-6 (Econverse): ERP integrado como fator de precificação de projetos |
| `kb-parceiros-wave6a.md` | FAQ-4 (Lemoon): ERP Proex requer intermediador; distinção agência vs. integrador |
| `kb-parceiros-wave7a.md` | FAQ-1 a FAQ-4 (Vtrina): confirmação de ERPs suportados e preços de referência |
| `kb-comercial-cross.md` | FAQ-2: Vtrina confirmada para Winthor; formulário oficial de indicação de integradores |
| `wake-partnership-faq.md` | Tabela consolidada de parceiros integradores de ERP com especialidades |
| `kb-projeto-hero-leanwork.md` | HOW_TO-5: mapeamento de IDs de grupo de pagamento no ERP Tiny (caso Karcher) |

> **Nota sobre `kb-parceiros-wave6a.md` (FAQ-4 / Lemoon)**: a referência original alegava que "ERP Proex requer intermediador". Essa fonte está incorreta — Proex é agência, não ERP. Substituída pela verificação em DB (2026-06-02).

---

## Verificação

Verificação realizada em 2026-06-05 contra `store/messages.db` para resolver lacunas apontadas em auditoria anterior (Everest, Vetor, Proex, TOTVS RM, Tiny).

| Claim | Status | Query / método | Evidência |
|---|---|---|---|
| Bling e Tiny são nativos da Wake (não exigem integrador em cenário padrão) | ✓ verificado | Playbook §2 + DB: `LIKE '%TINY%'` AND `LIKE '%integração%'` (mensagem Duo e Wake Parcerias, 2026-06-02: "sem precisar de integração, seria BLING e TINY?") | Duo e Wake Parcerias, 2026-06-02 |
| Tiny adotado em múltiplos clientes além de Karcher | ✓ verificado | `LIKE '%Tiny%' AND m.timestamp >= '2025-06-01'` (40 hits) | Empório do Celular (O2/Driven/Savvi, 2026-04-09); Wallford Brasil (O2/Catus/Avocado Tech, 2025-12-10); Avocado Tech Wallford (2025-12-22); Karcher (Projeto Hero <> Leanwork <> Wake, 2026-03-13) |
| TOTVS RM — Integra.do lista no site mas integração não está pronta | ✓ verificado | `LIKE '%RM%' AND (LIKE '%TOTVS%' OR LIKE '%Integra.do%')` (30 hits) | WAKE \| Time Parcerias, 2026-06-02 ("apesar da Integra.do ter a Totvs RM no site, o closer deles me disse que eles não têm a integração pronta" — Arnaldo da Savvi); SAVVI & Wake, 2026-06-02 |
| TOTVS RM — Intergate é parceira recomendada | ⚠ inferido | Sem mensagem recente confirmando integração ativa Intergate <> RM; afirmação herdada do artigo anterior e mantida como especialização. Validar com Partner Ops antes de indicar | Não confirmado em DB pós-2025 |
| Ecossistema atual de integradores ERP (lista completa) | ✓ verificado | `LIKE '%integração%' AND LIKE '%ERP%'` | Mensagem espelhada em 4+ grupos (WAKE \| Time Parcerias, SAVVI & Wake, Quick Digital, Duo e Wake) em 2026-06-02 listando: Argento Digital, Integra.do, Iota, EBI, LinkToU, Intergate, Floui, Jitterbit, Vtrina, Leanwork |
| Everest — integração com Wake | ⚠ validação pendente via Partner Ops | `LIKE '%Everest%'` (20 hits) | Única evidência relevante: WAKE \| Time Parcerias, 2025-07-16 ("ERP chamado de Everest e a princípio ele estaria integrado (homologado) com a Wake. Sabe se realmente esta?" — pergunta sem resposta); Parcerias \| Canais, 2025-10-10 ("convenção de vendas de um dos maiores parceiros aqui no ERP EVEREST"). Nenhuma confirmação de integração ativa |
| Vetor — existência como ERP integrável | ❌ sem fonte | `LIKE '%Vetor%'` (20 hits) | Apenas menção a "Vetor Editora" (Eficaz <> WAKE, 2026-05-29) como cliente referência B2C e a "logo vetorizado" (uso genérico). Nenhuma menção a "ERP Vetor" no corpus. Reivindicação original (Vtrina diz "desconhecido, necessita avaliação") não confirmada em DB |
| Proex é ERP que requer intermediador | ❌ sem fonte (reivindicação refutada) | `LIKE '%Proex%'` (30+ hits) | Proex aparece consistentemente como agência: Comercial Cross WAKE, 2025-08-18 (lista de agências); WAKE \| Time Agência, 2025-08-12 e 2025-12-19 (selo Special/Elite); Coder Ivy + Wake, 2026-05-19 (Yuri da Proex indicando cliente para Wake). Nenhuma referência a "ERP Proex". Linha removida da matriz |
| Vtrina + Winthor | ✓ verificado | Reconfirmado na mensagem do ecossistema de 2026-06-02 + FAQ-2 de `kb-comercial-cross.md` | WAKE \| Time Parcerias, 2026-06-02 |
| Bling — Vtrina cobre multiCD | ⚠ inferido | Sem mensagem direta no escopo da query; herdado de `kb-parceiros-wave7a.md` (FAQ Vtrina) | Validar com Vtrina |
| Argento cobre Linx/Senior/Omie/Abacos | ✓ verificado | Mensagem do ecossistema de 2026-06-02 cita literalmente esses 4 ERPs | WAKE \| Time Parcerias, 2026-06-02 |

**Audiências**: `audience-solutions` (primária), `audience-sales` (secundária — battlecard de seleção de integrador).

**Mapeamentos por status (matriz §1):**
- ✓ verificado: 12
- ⚠ inferido: 3 (Senior, Alpha7, Procfit)
- ❌ sem fonte: 0 (Proex e Vetor não constam mais da matriz §1; Everest movido para §3 com status pendente; Vetor movido para §3 com status sem fonte)

**Próximos passos para Partner Ops:**
1. Confirmar se existe roadmap de homologação para ERP Everest e atualizar este artigo
2. Confirmar se Senior continua coberto por Argento (sem mensagem recente)
3. Solicitar à Vtrina lista oficial atualizada cobrindo Alpha7, Procfit e demais ERPs em pipeline
4. Confirmar se Intergate mantém integração pronta com TOTVS RM (a única opção restante para esse ERP após exclusão de Integra.do)
