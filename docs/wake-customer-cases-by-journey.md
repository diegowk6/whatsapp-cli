# Wake Customer Cases — Por Jornada

**Purpose**: navigation index for sales and CS — find customers by *journey pattern* rather than vertical. Pattern recognition for objection handling, migration playbooks, churn prevention, and competitive battlecards.

**Source**: derived from `docs/wake-customer-cases.md` (102 cases, CASE-001 a CASE-102) e `docs/wake-customer-cases-medium.md` (52 leads em watchlist, LEAD-001 a LEAD-052). Cada entrada aqui referencia o ID canônico do arquivo-mestre — para detalhes (status, ticket, parceiros, contatos), abrir o caso completo lá.

**Confidencialidade**: mesmo regime do master — uso interno (Vendas / CS / Solutions / PAM). Não publicar valores comerciais nem nomes de stakeholders externos.

**Como ler**: um mesmo caso aparece em múltiplas jornadas. Status: ✓ ativo · ⚠ em negociação / risco · ❌ churn ou perdido.

---

## Journeys at a glance

| # | Jornada | Casos | Anchor reference | Insight de cabeçalho |
|---|---|---|---|---|
| 1 | Migração **VTEX → Wake** (ganhos) | 7 | CASE-004 Shoulder (público), CASE-001 Aramis | Argumento performance + TCO + atendimento; certificação Wake gratuita destrava o tech leader |
| 2 | Migração **Shopify → Wake** (ganhos) | 4 | CASE-005 Oriba (caso de SEO), CASE-097 HDM Beauty | Dor recorrente: bloqueio do Mercado Pago + conversão baixa |
| 3 | Migração **Tray → Wake** | 6 | CASE-029 Empório do Celular, CASE-089 Ivan Joias | Motivação canônica: customização e features (cashback, agenda vendedor) |
| 4 | Migração **Loja Integrada / Nuvem Shop → Wake** (SMB) | 4 | CASE-084 Believe + Tazin | Pitch SMB beleza tem caminho conhecido |
| 5 | Migração **Wbuy / OpenCart / Magento / VNDA / WordPress / WooCommerce / Climba / Visual / HostGator → Wake** | 11 | CASE-050 Protwins, CASE-073 Cerealista, CASE-092 Zinnia | Padrão "plataforma legada engessada"; competidor é frequentemente VTEX no upgrade |
| 6 | Migração de **plataforma desconhecida** (greenfield ou primeira op digital) | 8 | CASE-031 Norte Refrigeração, CASE-100 Unipeças MA, CASE-027 LEDVANCE | Cliente físico/B2B entrando no digital pela primeira vez |
| 7 | Wake → **Shopify** (churn) | 2 | CASE-044 DLK (manteve Vindi), CASE-043 Ricardo Almeida | Decisão TI / franquia ERP pressiona stack |
| 8 | Wake → **Bagy** (churn) | 1 | CASE-012 LP Beauty | Único caso no corpus; rollback interno LWSA |
| 9 | Wake → **Nuvemshop** (churn) | 1 | CASE-070 Unikapro / Unabeleza | Primeiro churn Nuvemshop no corpus |
| 10 | **Failed Go-Lives** (churn sub-6 meses) | 3 | CASE-043 Ricardo Almeida (3m), CASE-005 Oriba (3m) | Rushed implementation; decisão técnica isolada |
| 11 | **B2B journey** (atacado / franquias / matriz↔revenda / B2B luxo) | 14 | CASE-052 CVC, CASE-054 Destro, CASE-056 Gold, CASE-057 A Fábrica (anchors 4) | Mesma loja c/ área logada OU duas lojas separadas (B2C parent + B2B child) |
| 12 | **Programa de afiliados / influencers** | 3 | CASE-004 Shoulder, CASE-017 BAW, CASE-047 Ybera | Dois padrões técnicos: módulo de cupom OU StoreFront API |
| 13 | **Marketplace-led** (site < 10% da receita) | 6 | CASE-007 Soneda, CASE-033 Off Rush, CASE-081 7 Oliveiras | Anymarket-anchored; pitch Wake é virar o jogo no site |
| 14 | **Multi-CD / crossdocking / OMS unificado** | 5 | CASE-023 Casa Almeida, CASE-028 Hiper Equipamentos, CASE-086 Loja do Detalhamento | OMS Wake como diferencial |
| 15 | **Multi-marca / grupo** | 6 | CASE-002 Inbrands, CASE-091 Shop2gether/Grupo Icomm | Anymarket 1-conta + multi-CD são gargalos |
| 16 | **Primeira operação digital** (no prior e-commerce) | 5 | CASE-031 Norte, CASE-100 Unipeças, CASE-027 LEDVANCE | Físico expandindo para digital |
| 17 | **Cases públicos Wake** (uso livre em pitches) | 3 | CASE-004 Shoulder, CASE-005 Oriba, CASE-091 Shop2gether | wake.tech/cases publicados |
| 18 | **AI-acquired leads** (Gemini / LLMs) | 1 | CASE-097 HDM Beauty | Primeira evidência de Wake recomendada por LLM |
| 19 | **Battles competitivas ativas** (jun/2026) | 3 | CASE-094 HVSUL (VTEX 6 meses grátis), CASE-047 Ybera | Battlecard VTEX em validação |
| 20 | **B2B Wake UP (R$ 4,9k MRR)** | 2 | CASE-093 Cittius, CASE-094 HVSUL | Tabela 2026 confirmada |
| 21 | **Stack Metakosmos** (provador / 3D em PDP) | 2 | CASE-060 Gregory, CASE-096 Fuel | Únicos cases em produção |
| 22 | **Stack HERO Seguros** (garantia estendida + parcela garantida) | 1 | CASE-003 Karcher | Primeiro cliente HERO |
| 23 | **Stack Stone TEF + Prateleira Infinita / POS** | 2 | CASE-054 Destro, CASE-061 Ave Rara | Cuidado: validar parque de maquininhas no discovery |
| 24 | **Vertical regulada — farma** | 1 | CASE-067 Singular Medicamentos | Gap: validação de receita ainda manual |
| 25 | **Vertical regulada — joalheria** | 3 | CASE-089 Ivan Joias, CASE-101 Gold Prime, Orit (ref. viva) | Benchmark visual: orit.com.br |
| 26 | **Vertical editorial / publishers** | 1 | CASE-059 Editora Fórum | Greenfield no catálogo |
| 27 | **Leads desqualificados / perdidos por preço** | 5 | CASE-069 Biggie, CASE-083 Carinion (VNDA), CASE-088 Festa Color | Aprendizados de qualificação |
| 28 | **Leads que esfriam por falta de retorno** | 8+ | CASE-063, CASE-064, CASE-077, CASE-079, CASE-101 | Padrão SMB/regional sem follow-up |
| 29 | **Operações com bug recorrente em ROAS / Meta** | 1 | CASE-072 Você Constrói | Evento Purchase com receita zerada em headless+front terceirizado |
| 30 | **Operações multi-segmento** (B2C + B2B + automotivo na mesma marca) | 3 | CASE-048 Palácio das Essências, CASE-097 HDM Beauty | Arquitetura multi-vertical na mesma instância |

---

## 1. Migrações FROM competitors INTO Wake (wins)

### 1.1 Migração **VTEX → Wake**

**Pattern**: Cliente VTEX sente o custo do TCO (Trinio + Anymarket TR + agência), descobre certificação Wake gratuita, e o tech leader endorsa antes do fechamento. Wake ganha por **performance + atendimento + custo**.

**Cases**:
- **CASE-001 Aramis** — ✓ assinado 2026-06-02. Maior contrato da safra. Tech leader VTEX se certificou em Wake antes do fechamento; afirmou *"A Wake é uma Ferrari perto da VTEX"* (Tropa de Elite 3, 2026-05-22).
- **CASE-004 Shoulder** — ✓ case público de migração (wake.tech/cases). Narrativa de migração assistida pronta para uso em pitch.
- **CASE-030 Wallford Brasil** — ⚠ em negociação. Cliente cita política comercial VTEX como benchmark de arquitetura multi-marca.
- **CASE-047 Ybera** — ⚠ live battle. POC VTEX em curso, detrator interno (Vinícius, ex-founder).
- **CASE-065 Brofitwear** — ✓ Go-Live 2026-05-19. Primeiro Go-Live da Quick Digital. Migração concluída em ciclo de ~7 meses.
- **CASE-078 Balletto** — ⚠ em negociação via O2 Consultoria; ERP Sisplan não-nativo.
- **CASE-091 Shop2gether (Grupo Icomm)** — ✓ case público. Fechamento 2025-07-30, migração agência DRIVEN → Oásis durante implementação.
- **CASE-094 HVSUL** — ⚠ live battle. VTEX ofertou 6 meses grátis.

**Argumentos que ganharam**:
- Performance / arquitetura (Aramis: *"Ferrari perto da VTEX"*)
- TCO real (Trinio + AnyMarket TR) vs. Wake + Anymarket nativo
- Atendimento direto (CASE-004 Shoulder mediado por equipe interna Comercial→Onboarding→CS→Parcerias)
- Certificação Wake gratuita para tech leaders (Aramis)
- Indicação via DXP/Wake como produto interno (CASE-078 Balletto)

**Objeções recorrentes**:
- "Detrator técnico interno" — CASE-047 Ybera (CTO ex-founder)
- VTEX agressiva em desconto — CASE-094 HVSUL (6 meses grátis)
- Cliente quer manter política comercial VTEX inteira — CASE-030 Wallford

---

### 1.2 Migração **Shopify → Wake**

**Pattern**: Dor canônica é **bloqueio do Mercado Pago + conversão pós-migração baixa**. Cliente vem com discurso "minha plataforma não tá vendendo".

**Cases**:
- **CASE-005 Oriba** — ✓ Go-Live 2026-01-13 com ganho de SEO documentado (PDF público publicado 2026-05-04). ❌ posterior churn (relacionado ao grupo LWSA, não à plataforma).
- **CASE-064 Cimafer** — ⚠ esfriou. Saía de Shopify por *"não sente segurança nas vendas online"*. ERP Boreo não-nativo.
- **CASE-069 Biggie** — ❌ desqualificado. Cliente Shopify+Tiny só queria refazer frontend, sem intenção de migrar. ERP Tiny já tem integração nativa Shopify.
- **CASE-097 HDM Beauty (eclerbeauty.com.br)** — ⚠ em negociação. Tray → Shopify → quer Wake. Dor: bloqueio Mercado Pago + conversão baixa pós-Shopify. Lead via Gemini IA.

**Argumentos que funcionam**: ganho de SEO (Oriba), feature "Kits com desconto progressivo" (Oriba), saída do Mercado Pago para Vindi/Pagar.me.

**Cuidado de qualificação**: cliente Shopify+Tiny que só quer "refazer o front" é lead falso (CASE-069 Biggie + Trocamotor → Catus pediu compensação).

---

### 1.3 Migração **Tray → Wake**

**Pattern**: Tray frustra por **limitação de customização**. Cliente busca cashback, agenda vendedor, mix & match — features-gatilho do case Shoulder.

**Cases**:
- **CASE-029 Empório do Celular** — ⚠ em negociação. *"Sente que Tray é limitada, deseja liberdade de customização"*. ERP Tiny.
- **CASE-079 AR15 Sport** — ⚠ lead frio. Vertical airsoft.
- **CASE-085 Regina Salomão** — ⚠ projeto adiado para 2026. ERP MV Tecnologia.
- **CASE-086 Grupo Loja do Detalhamento (3 marcas detailing)** — ⚠ migração c/ OMS unificado. ERP Arth Magnus.
- **CASE-089 Ivan Joias** — ⚠ em negociação. Migração **Tray → Wake dentro da mesma LWSA**, decisão por benchmark Shoulder (cashback + agenda vendedor). ERP 4Control.
- **CASE-099 Nano4you Brasil** — ⚠ em negociação. Lead vindo de evento Wake Curitiba 2026-05-21. Features-gatilho: múltiplos CDs + regionalização.
- **CASE-101 Gold Prime Joalheria** — ❌ perdido por falta de retorno. Usou Orit como benchmark.

**Argumento principal**: capacidade de customização + cashback / agenda vendedor / regionalização.

---

### 1.4 Migração **Loja Integrada / Nuvem Shop → Wake** (SMB upgrade)

**Pattern**: SMB beleza/moda que sente o teto da plataforma "lojinha" e quer profissionalizar.

**Cases**:
- **CASE-082 Conceito Rouparia** — ⚠ Loja Integrada → Wake, mantém ERP Datasystem via 4middleware (API).
- **CASE-084 Believe Professional + Tazin Cosméticos** — ⚠ ambos em negociação. Believe sai de Loja Integrada (ERP Bling); Tazin sai de Nuvem Shop pagando R$ 7,5k/mês para agência atual.
- **CASE-095 Inigual Decor** — ❌ desqualificado em 2026-05-08. R$ 20k/mês site + 90% em marketplaces; sem perfil Wake.

---

### 1.5 Migração de **outras plataformas legadas → Wake**

**Pattern**: "Plataforma engessada / não vende / custo alto" — argumento de venda comum independente da plataforma de origem.

**Cases por plataforma de origem**:
- **Wbuy**: CASE-050 Protwins (⚠ qualificação, R$ 90k/mês, competição com VTEX/Quick).
- **OpenCart**: CASE-063 Eletro Mais MT (⚠ lead frio; 6 meses no ar sem venda).
- **Magento**: CASE-073 Cerealista Express (⚠ em rotação; mensalidade alta, plataforma engessada).
- **WordPress / WooCommerce**: CASE-066 Casafer (⚠ qualificação, SAP B1, e-commerce só catálogo); CASE-087 Alfredo Meyer (⚠ migração c/ requisitos RD Station + WhatsApp + ERP + marketplaces).
- **VNDA**: CASE-092 Zinnia Leather (⚠ couro alto ticket, melhorar recomendação).
- **Climba**: CASE-075 Inoxlon (⚠ lead frio por contato inexistente; ERP Bling; GMV R$ 200k).
- **Visual e-commerce**: CASE-081 7 Oliveiras (⚠ rotação; ERP ZADA; 80% GMV em marketplaces).
- **HostGator**: CASE-083 Carinion Mix (❌ perdido para VNDA por preço).
- **Tray → Shopify → Wake** (dupla migração): CASE-097 HDM Beauty (⚠ em negociação).

**Insight**: clientes que **saíram de plataformas SMB low-end** (HostGator, OpenCart) frequentemente perdem para outra SMB low-end (VNDA), não para a Wake — sinal de **desqualificação antes de distribuir**.

---

### 1.6 **Greenfield digital** (sem plataforma anterior)

**Pattern**: Cliente físico forte / B2B forte / fabricante entrando no e-commerce pela primeira vez.

**Cases**:
- **CASE-020 Ello Moda Fitness e Praia** — ⚠ moda fitness emergente (1 loja Brás atacado+varejo); foco time-to-market.
- **CASE-021 ACC Modas** — ⚠ 11 lojas físicas (~R$ 200k/loja/mês), estruturando e-commerce próprio.
- **CASE-026 Verdes Mares** — ⚠ material de construção interior BA; 1 loja + 1 CD; nova loja prevista jul/2026; foco B2B.
- **CASE-027 LEDVANCE (ex-Osram)** — ⚠ greenfield B2C; ERP SAP ECC 6.0 (S/4HANA planejado 2027).
- **CASE-031 Norte Refrigeração** — ⚠ multi-CD que opera como seller (sem ser marketplace).
- **CASE-032 Empório do Sertão** — ⚠ omnichannel; 5 lojas RJ; R$ 20M offline; e-commerce anterior frustrado.
- **CASE-088 Festa Color** — ❌ desqualificado. R$ 1,5M loja física + maturidade digital zero.
- **CASE-100 Unipeças MA** — ⚠ greenfield B2C → B2B. 30+ anos de varejo físico + televendas; R$ 10-30M/ano; New Holland / Bosch / Tramontina.

**Insight**: greenfield + perfil B2B forte = projeto **Wake Commerce + Wake B2B futuro** (ver §3 abaixo). Sem perfil B2B explícito, validar maturidade digital antes de distribuir (CASE-088 é o anti-padrão).

---

## 2. Wake LOSSES (churn / migrações de saída)

### 2.1 Wake → **Shopify**

**Pattern**: Decisão TI isolada (CASE-043) OU vontade de manter parceiro de pagamento (CASE-044).

**Cases**:
- **CASE-043 Ricardo Almeida** — ❌ churn em 3 meses pós Go-Live. *"TOTVS Moda. TI decidiu sozinho essa migração"* — pressão da franquia TOTVS Ibirapuera (2B Digital).
- **CASE-044 DLK Modas** — ❌ confirmado 2026-05-25. *"Estão indo para Shopify mesmo, falei aqui com Monisi por conta da Vindi, querem manter a Vindi lá"* — Vindi é stack-agnóstico no relacionamento com cliente final.

**Razões comuns**:
- Acoplamento ERP↔Plataforma (franquia TOTVS pressionando)
- Quer manter parceiro de subscription/pagamento (Vindi)
- Decisão técnica isolada sem alinhamento comercial

---

### 2.2 Wake → **Bagy**

**Cases**:
- **CASE-012 LP Beauty** — ❌ rollback para Bagy. Único caso documentado de Wake → Bagy no corpus.

**Razão**: cliente faz "downgrade" interno dentro do grupo LWSA (Wake mid-enterprise → Bagy SMB). **Cuidado**: Bagy foi anunciada como descontinuada como tech em 2026-04-12 (Tropa de Elite 3) — rollback pode ser revertido.

---

### 2.3 Wake → **Nuvemshop**

**Cases**:
- **CASE-070 Unikapro / Unabeleza** — ❌ confirmado 2026-03-27. Migração com rebranding (Unikapro → Unabeleza + Nuvemshop). Ausência da Unabeleza em **planilha de comissionamento Econverse** em jan/2026 foi **leading indicator** 2 meses antes do churn formal.

**Razão**: primeiro churn Nuvemshop do corpus. Padrão **rebranding + migração** sinaliza saída planejada, não fricção operacional aguda.

---

### 2.4 Wake → **VNDA / outras SMB** (não-existente direto, mas perdas para)

- **CASE-083 Carinion Mix** — perdido para VNDA *antes* de virar cliente Wake (não é churn; é lead perdido por preço).
- **CASE-046 Cadiveu** — ⚠ saída em curso (2025-07-22, sem destino confirmado).
- **CASE-045 Mimeria** — ❌ churn iminente; pediu contato de **parceiros concorrentes de CRMBonus** (red flag clássico).

---

## 3. Failed Go-Lives (sub-6-month churns)

**Pattern**: Rushed implementation, decisão técnica isolada, modelo de contrato frágil.

**Cases**:
- **CASE-005 Oriba** — Go-Live 2026-01-13 → churn dentro do 1ºT/2026 (≤3 meses). Contexto LWSA Shoulder + Oriba.
- **CASE-043 Ricardo Almeida** — assinado 08/2025 → churn solicitado 10/2025 (3 meses). TI decidiu sozinho.
- **CASE-041 Mpozenato** — ❌ perda de **R$ 1,8M**. Contrato de **valor fixo por pedido** em cliente marketplace-heavy — receita frágil quando cliente sai.

**Lessons**:
- Validar decisão multi-stakeholder (não só TI) antes do fechamento
- Evitar contratos de **valor fixo por pedido** para clientes marketplace-heavy
- Atenção redobrada quando ERP/franquia tem stack próprio (TOTVS Moda)

---

## 4. B2B Journey

**Arquitetura Wake**: a mesma loja pode operar B2C público + B2B logado, OU pode ser **duas lojas Wake separadas** (parent B2C + child B2B). A frente B2B suporta **listas de preço por cliente, regionalização, crédito, mix de produto, checkout dedicado** (CoreBiz & Wake, 2026-01-30: *"Hoje temos uma estrutura de B2B bem completa"*).

**Os 4 cases-âncora canônicos** (Comercial Cross WAKE, 2026-03-05: *"CVC, GOLD, Destro e A Fabrica"*):

| Variante B2B | Caso | Destaque arquitetural |
|---|---|---|
| Franquia (matriz → rede) | **CASE-052 CVC Corp** | Listas de preço por franquia + crédito + regionalização |
| Atacadista regional | **CASE-054 Destro Macro** (PR) | Stack inclui Stone como adquirente |
| Distribuição food service | **CASE-056 Gold Food Service** | Único case food service do corpus |
| Indústria → revenda c/ checkout dedicado | **CASE-057 A Fábrica** | `checkout.afabrica.com` em subdomínio próprio |

**Variantes adicionais**:
- **Atacado puro (kids)** — CASE-053 Yora (`yoraatacado.com.br`)
- **B2C parent + B2B child como duas lojas separadas** — CASE-072 Você Constrói (B2C) + CASE-051 Esplane (B2B)
- **B2B em onboarding** sobre B2C já ativo — CASE-055 Buddemeyer (B2C contratado 06/2025; B2B "ainda em projeto" 03/2026; benchmark interno c/ CVC)
- **Brinquedos varejo + atacado** — CASE-058 Unitoys (referência interna em brinquedos)
- **B2B brindes corporativos** — CASE-068 OG Brindes (em busca de agência de sustentação)
- **B2B luxo dentro de cliente B2C multi-marca** — CASE-097 HDM Beauty (eclerbeauty.com.br para revendedores)
- **Varejo eletro + Dealer Shop B2B** — CASE-074 Eletrônica Santana (grupo c/ B2C + B2B + ES TECH)
- **Wake UP B2B (R$ 4,9k MRR — tabela 2026)** — CASE-093 Cittius (blocos colecionáveis, lançamento ABRIN); CASE-094 HVSUL (live battle vs VTEX)

**Insights**:
- Padrão **B2C parent + B2B child como lojas separadas** = isolamento total entre frentes (Você Constrói + Esplane).
- **Wake B2B com lista de preço por franquia/cliente + crédito + regionalização** é o pacote técnico citado em pitch enterprise (Buddemeyer + CVC para CoreBiz).
- B2B com Wake é projeto de prazo: Buddemeyer contratado 06/2025, B2B ainda em onboarding 9 meses depois — **calibrar expectativas**.

---

## 5. Programa de Afiliados / Influencers

**A Wake não tem produto de afiliados separado.** Programas são montados sobre dois recursos da plataforma:

1. **Painel Wake — módulo de campanha/cupom** (Monitfy & Wake, 2026-05-18). Pegada operacional simples; API correspondente é mais limitada que o painel.
2. **StoreFront API com `partner/afiliado` no carrinho** (Uble & Wake, 2026-01-19). Mutations `CheckoutClone` + `CheckoutAddMetadata` preservam afiliado e UTM através do checkout.

**Os 3 cases canônicos** (Comercial Cross WAKE, 2025-12-17: *"Ybera / Shoulder para funcionários / BAW com influenciadores"*):

| Modelo | Caso | Detalhe |
|---|---|---|
| Funcionários como afiliados | **CASE-004 Shoulder** | Rede de vendedoras + colaboradoras |
| Influenciadores externos | **CASE-017 Baw** | Programa de influencer marketing |
| Operacional dentro de relacionamento D2C | **CASE-047 Ybera** | Em live battle com VTEX; afiliados parte do diferencial |

**Sinal adjacente**: CASE-020 Ello Moda pediu *"área de afiliados organizada"* usando DLK como benchmark — programa de afiliados é **argumento de venda crescente** em moda 2026.

**Limitação a sinalizar em discovery**: gestão programática via API é mais restrita que via painel — para orquestração complexa, validar escopo técnico antes de fechar.

---

## 6. Marketplace-led customers (site < 10% da receita)

**Pattern**: site representa pequena fração; venda concentrada em Anymarket + canais. Pitch Wake = "vamos virar o jogo no site".

**Cases**:
- **CASE-007 Soneda** — ✓ ativo. Wake + Floui (ERP) + Anymarket (incluindo TikTok Shop).
- **CASE-016 Cacay** — ✓ ativo. Migração para Dafiti via Anymarket (erro DE/PARA frete documentado).
- **CASE-028 Hiper Equipamentos** — ✓ multi-CD com crossdocking via Anymarket.
- **CASE-033 Off Rush** — ⚠ R$ 400k/mês total, **1% via site**. Vertical autopeças/motos.
- **CASE-062 OffPremium** — ✓ piloto Marketplace Blueman + Anymarket; gap conhecido no `marketPlacePedidoId`.
- **CASE-081 7 Oliveiras** — ⚠ R$ 20k site vs R$ 160k marketplaces. Pitch deve enfatizar Anymarket como hub central.
- **CASE-095 Inigual Decor** — ❌ desqualificado. 90% em marketplaces + R$ 20k site = sem perfil Wake (qualificação filtrou bem).

---

## 7. Multi-CD / crossdocking / OMS unificado

**Pattern**: complexidade logística é o gancho de Wake OMS como diferencial.

**Cases**:
- **CASE-023 Casa Almeida (Buddemeyer)** — ✓ Go-Live 02/2026. Suporte assistido Anymarket (~R$ 1k–1,5k extra). Referência regional Sul.
- **CASE-028 Hiper Equipamentos** — ✓ multi-CD com crossdocking ativo via Anymarket.
- **CASE-031 Norte Refrigeração** — ⚠ "Não é seller, é multi-CD mas opera igual um".
- **CASE-086 Grupo Loja do Detalhamento (3 marcas)** — ⚠ 3 lojas separadas, 1 CD único, controle unificado de estoque e pedidos.
- **CASE-098 Pisom & Co** — ⚠ unificação OMS multi-loja física construção/pisos.

**Argumentos**: Wake OMS como camada de unificação. Cuidado: **Hiper Equipamentos** é exemplo de operação difícil (multi-CD + crossdocking + Anymarket = ticket técnico recorrente).

---

## 8. Multi-marca / grupo

**Pattern**: holdings ou grupos com 3+ marcas operadas em paralelo.

**Cases**:
- **CASE-002 Inbrands** — ✓ enterprise. Ellus, Richards, VR, Tommy Hilfiger BR; 5 marcas + multi-CD planejado. Anymarket 1-conta é gargalo arquitetural conhecido.
- **CASE-008 Tommy (Tommy Hilfiger BR)** — ✓ assinado 2026-05-27. Possível expansão dentro do Grupo Inbrands.
- **CASE-023 Casa Almeida (Buddemeyer)** — ✓ multi-marca cama/mesa/banho.
- **CASE-074 Eletrônica Santana** — ✓ 3 marcas (Eletrônica Santana + ES TECH + Dealer Shop B2B).
- **CASE-076 AZ Acessórios (Grupo Cambuci)** — ✓ Go-Live 2025-06-23 via Eficaz.
- **CASE-091 Shop2gether (Grupo Icomm)** — ✓ case público; combinar com OqVestir (mesmo grupo).

**Sinal adjacente**: CASE-030 Wallford Brasil planeja 5 marcas com fronts diferentes + catálogo único (em negociação; benchmark VTEX).

---

## 9. Cases públicos Wake (uso livre em pitches)

- **CASE-004 Shoulder** — `wake.tech/cases` — migração assistida (live interna LWSA 2026-05-15)
- **CASE-005 Oriba** — PDF público 2026-05-04 (Kits desconto progressivo + ganhos SEO). ⚠ Cliente posteriormente saiu — validar uso em pitch público.
- **CASE-091 Shop2gether (Grupo Icomm)** — `wake.tech/cases/como-a-wake-ajudou-o-grupo-icomm-shop2gether-e-oqvestir-com-sites-de-alta-performance/`

**Citados como referência em discoveries de outros prospects**:
- Shoulder → CASE-088 Festa Color, CASE-089 Ivan Joias, CASE-021 ACC Modas
- Pura Vida → CASE-088 Festa Color (junto com Shop2gether)
- Orit (referência viva, sem case dedicado) → CASE-089 Ivan Joias, CASE-101 Gold Prime

---

## 10. AI-acquired leads (novo signal 2026)

- **CASE-097 HDM Beauty** — primeira evidência no corpus de lead vindo via **recomendação Gemini IA** (cliente relatou problemas Shopify, Gemini recomendou Adobe Commerce + Wake Commerce).

**Insight**: sinal de presença Wake em LLMs. Monitorar fluxo entrante por origem "IA" para validar tendência.

---

## 11. Live competitive battles (jun/2026)

**Use estas batalhas ativas para refinar o battlecard.**

- **CASE-094 HVSUL** — ⚠ live. VTEX ofertou **6 meses grátis**. Wake UP B2B R$ 4,9k MRR em negociação. Conexão Grupo Anhanguera (head de e-com pode endossar).
- **CASE-047 Ybera** — ⚠ live. POC com VTEX em curso. Detrator interno (Vinícius, diretor de tech, ex-founder). Considerado estratégico ("não pode perder de jeito nenhum").
- **CASE-050 Protwins** — ⚠ migração Wbuy. VTEX/Quick indicada em paralelo. Conta SMB no limite (R$ 90k/mês).

**Cenários adjacentes não-VTEX**:
- **CASE-083 Carinion Mix** — ❌ perdido para VNDA por preço (R$ 600/mês teto).

---

## 12. Stacks integradas (parceiros + clientes)

### 12.1 Stack **Metakosmos** (provador virtual / 3D em PDP)

**Únicos cases em produção** (referência canônica Wake & Metakosmos, 2026-05-14):
- **CASE-060 Gregory** — ✓ provador virtual em produção (`gregory.com.br`). **Revenue share Metakosmos por Gregory nunca foi cobrada** até 02/2026 (gap operacional de billing).
- **CASE-096 Fuel (UseFuel)** — ✓ visualizador 3D em PDP em produção (`usefuel.com.br/produto/oculos-de-sol-storia-152540`).

### 12.2 Stack **HERO Seguros** (Garantia Estendida + Parcela Garantida)

- **CASE-003 Karcher** — ✓ primeiro cliente da integração HERO Seguros. Prazo de garantia 12 meses confirmado por e-mail em 2026-05-29. Integração via Leanwork.

### 12.3 Stack **Stone TEF + Prateleira Infinita / POS**

- **CASE-054 Destro Macro** — projeto Wake↔Stone teve gap de mapeamento de fluxo em 07/2025.
- **CASE-061 Ave Rara** — em negociação ativa de Prateleira Infinita 06/2026. **Lição**: nem todas as maquininhas do parque atual eram compatíveis — validar compatibilidade POS no discovery antes do fechamento.

### 12.4 Stack **Pagar.me + Konduto + parceiro logístico** (kit canônico SMB)

- **CASE-071 Empório Bud** — discovery 09/2025 explicitou stack Pagar.me + Konduto + logística como mínimo viável.

### 12.5 Stack **PagBank em checkout Headless**

- **CASE-059 Editora Fórum** — implementação via Benova (Pix + CC); status pós-09/2025 não confirmado no corpus.

### 12.6 Stack **Anymarket + Marketplace Blueman**

- **CASE-062 OffPremium** — piloto. Gap conhecido: pagamento direto + número de título ≠ `marketPlacePedidoId`.

### 12.7 Stack **CRMBonus**

- **CASE-002 Inbrands** — ✓ enterprise. Cache acumulando promoções indevidamente em 2026-04-06.
- **CASE-045 Mimeria** — ❌ churn iminente (pediu contato de parceiros concorrentes de CRMBonus — red flag).
- **CASE-096 Fuel** — citado por CRMBonus em 2025-08-22 como "ainda sem Go-Live" (Go-Live posterior).

### 12.8 Stack **Serasa Pass**

- **CASE-024 Balaroti** — ✓ ativação Serasa Pass concluída 2026-05-07. Caso de financiamento em material de construção (ticket alto).
- **CASE-096 Fuel** — em discovery com Serasa em 2026-04-27.

### 12.9 Stack **Figo + Givex**

- **CASE-004 Shoulder** — Figo (MVP encerrado pós-evento 2026-06-01) + Givex (gift card).

### 12.10 Stack **Eitri** (app mobile)

- **CASE-044 DLK Modas** — Eitri mantém app Wake mesmo após cliente migrar para Shopify.

---

## 13. Verticais reguladas

### 13.1 Farma

- **CASE-067 Singular Medicamentos** — ✓ ativo em produção; painel Wake operacional desde 11/2025. **Gap funcional aberto em 06/2026**: validação de receita médica + dados do médico ainda manual. Sinal para Solutions/Produto.

### 13.2 Joalheria

- **CASE-089 Ivan Joias** — ⚠ Tray → Wake, 6 lojas Salvador, decisão por benchmark Shoulder.
- **CASE-101 Gold Prime Joalheria** — ❌ perdido por falta de retorno. Usou **Orit** como benchmark de site.
- **Orit** (`orit.com.br`) — referência viva citada em discoveries (sem case dedicado).

### 13.3 Editoras / Publishers

- **CASE-059 Editora Fórum** — ✓ em projeto técnico 08/2025 (checkout Headless + PagBank via Benova). **Único publisher do corpus** — vertical greenfield.

---

## 14. Padrões de qualificação e perdas

### 14.1 Leads desqualificados (não-perfil Wake)

- **CASE-069 Biggie** — cliente Shopify+Tiny só queria refazer frontend. Catus pediu compensação por lead falso.
- **CASE-083 Carinion Mix** — perdido para VNDA por preço (R$ 600/mês teto cliente).
- **CASE-088 Festa Color** — desqualificado. R$ 1,5M loja física mas maturidade digital zero, sem intenção imediata.
- **CASE-095 Inigual Decor** — desqualificado. R$ 20k site + 90% em marketplaces; sem perfil.

**Sinais de desqualificação**:
- Faturamento online < R$ 50k/mês com >80% em marketplaces (Anymarket-only deveria ser o pitch)
- Cliente quer só "refazer o front" da plataforma atual
- Teto de orçamento mensal R$ 600-1.500 (SMB low-end)
- Maturidade digital zero + sem urgência declarada

### 14.2 Leads que esfriam por falta de retorno

Padrão SMB/regional clássico: lead distribuído ao carrossel, agências tentam contato, cliente não responde.

- CASE-063 Eletro Mais MT (agenda marcada, sem retorno)
- CASE-064 Cimafer (não atende celular, não responde)
- CASE-066 Casafer (SAP B1 + WooCommerce; sem atualização)
- CASE-077 Dallanese Parafusos (sem sucesso em contato em 4 semanas)
- CASE-079 AR15 Sport (sem respostas, vertical airsoft)
- CASE-101 Gold Prime Joalheria (perdido por falta de retorno em 6 semanas)
- CASE-085 Regina Salomão (proposta de desconto venceu; adiado para 2026)

**Insight**: leads multi-agência em rotação carrossel **esfriam quando o cliente não responde** — sintoma de saturação do modelo carrossel (CASE-049 + CASE-102). Recomendação: **qualificação dupla** (Wake SDR + agência) antes de distribuir.

---

## 15. Operações com bugs / fricções recorrentes (escalation patterns)

### 15.1 Bug evento **Purchase do Meta** com receita zerada (headless + front terceirizado)

- **CASE-072 Você Constrói** — chamado 258709 aberto 12/2025, ainda escalado em 02/2026. Front feito por Auaha causou regressão. Padrão recorrente em Wake + agência externa de front + Meta Ads.

### 15.2 Cache acumulando promoções indevidamente

- **CASE-002 Inbrands** — CRMBonus, 2026-04-06.

### 15.3 Multi-CD + crossdocking + Anymarket

- **CASE-028 Hiper Equipamentos** — XML com problemas desde 02/2026; ticket recorrente.

### 15.4 DE/PARA de frete em migração marketplace

- **CASE-016 Cacay** — ticket Anymarket #22241 em migração para Dafiti.

### 15.5 SLA agência em sustentação

- **CASE-060 Gregory** — Floui defendeu SLA contratual de **15 dias** para subida de produtos como baseline.

---

## 16. Padrões por partner / agência (referência cruzada)

### 16.1 Agências com case de marco

- **Simples Inovação** → primeiro case CASE-061 Ave Rara (Go-Live 2026-05-18)
- **Quick Digital** → primeiro Go-Live CASE-065 Brofitwear (2026-05-19)
- **Sellbie** → CASE-008 Tommy (Go-Live previsto 2026-07-20)

### 16.2 Agências fora do programa que ainda atendem base

- **B8One** — saiu do programa; ainda necessária para ajustes legados em CASE-024 Balaroti
- **DRIVEN** — transferiu CASE-091 Shop2gether para Oásis durante implementação por dificuldade

### 16.3 Sustentação regional (Campinas)

- **CASE-034 Carblue** — cliente da base pediu agência regional; indicado a Catus + Nairuz

---

## 17. Watchlist — Leads em monitoramento (`wake-customer-cases-medium.md`)

A segunda lista (52 leads, LEAD-001 a LEAD-052) cobre **prospects 2-hit** ainda sem desfecho confirmado. Padrões de jornada que predominam ali:

- **Greenfield digital** + perfil B2B regional (LEAD-005 Anhanguera Ferramentas, LEAD-006 Imperial, LEAD-012 Armazém CEC, LEAD-049 MPS Distribuidora)
- **Migração de plataforma SMB** (LEAD-008 GoPure, LEAD-009 Vital Bio, LEAD-021 Puriflora — todos Loja Integrada / Tray)
- **Eventos** como originadores de lead (LEAD-031 TJCC, LEAD-032 Empório Bud / Konduto)
- **Setor saúde / wellness adjacente** (LEAD-022 Eleve Suplementos, LEAD-046 AnjoRosa)

**Promotion criteria** (ver §"Promotion criteria" no master): para virar CASE-NNN, o lead precisa de 3+ hits substantivos e status atual identificável.

---

## Como usar este arquivo

1. **Em pitch de prospect**: vá direto à jornada mais próxima (ex.: "cliente Shopify de moda" → §1.2 + §1.3 + §11 se houver competição VTEX).
2. **Em battlecard**: §11 (live battles) e §2 (churn outbound) trazem o material competitivo mais fresco.
3. **Em discovery técnica**: §12 (stacks integradas) lista compatibilidades já testadas com cada parceiro e os gaps conhecidos.
4. **Em qualificação de lead**: §14 lista os anti-padrões — se o lead bate em 2+ sinais de desqualificação, segurar antes de distribuir ao carrossel.
5. **Para detalhes de qualquer caso**: abrir `docs/wake-customer-cases.md` no CASE-NNN correspondente — todo dado verificado vive lá.

**Atualização**: este arquivo deriva do master. Quando o master ganhar novos CASE-NNN, revisitar as seções abaixo para inserir referência cruzada (não duplicar conteúdo).

---

## Tags

`audience-sales`, `audience-cs`, `audience-pam`, `audience-solutions`, `customer-case`, `navigation-index`
