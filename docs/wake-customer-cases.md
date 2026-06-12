# Wake Customer Cases — Catálogo de Clientes e Casos de Uso

**Artefato**: catálogo de customer cases extraído do corpus WhatsApp Wake/parceiros (`store/messages.db`).
**Data**: 2026-06
**Fonte**: mensagens de grupo (`chat_jid LIKE '%@g.us'`, texto puro), recortes 2025-06 a 2026-06.
**Total de casos**: 102

---

## Confidencialidade

Este catálogo é de **uso interno** das equipes de Vendas, CS, Solutions e PAM da Wake. Vários casos referenciam:

- Valores comerciais específicos (MRR/setup) capturados em mensagens internas dos grupos Comercial Cross WAKE, Mavericks e Tropa de Elite 3.
- Situações sensíveis de churn, migração, conflitos comerciais e relacionamentos com agências.
- Status de negociação ainda em andamento.

**Antes de qualquer redistribuição externa** (apresentação a prospect, post em LinkedIn, case público), confirme o status atual com o account owner / CS da conta e remova dados comerciais. Para uso interno em sales pitches, sempre prefira citar marca + verticais + ganhos operacionais — não os valores de contrato.

JIDs, telefones pessoais e CPFs foram removidos. Mensagens citadas referenciam apenas o nome do grupo e a data.

---

## Como ler este catálogo

Cada caso traz:

- **Status**: `✓ ativo`, `⚠ em negociação`, `❌ churn`, `✓ migrado` (de outra plataforma para Wake), `⚠ status desconhecido`.
- **Plataforma anterior / ERP / Marketplace / Parceiros envolvidos**: do que foi confirmado em mensagem; campos faltantes ficam como `—` (não informado no corpus).
- **Situação / Solução / Resultado**: situação extraída do corpus; quando uma claim depender de extrapolação, vem marcada `⚠ inferido`.
- **Insights**: usos práticos para o time de vendas/CS.

Casos sem evidência suficiente foram omitidos. Empresas que aparecem apenas em listas de eventos (Vulcabras, Lacoste, Vivara, etc., do mailing do evento Frenet + Wake 2026-04-01) **não** estão aqui — presença em lista de convite ≠ relação de cliente.

---

## Índice

1. **Anchor customers** (referências top usadas em vendas): CASE-001 a CASE-008
2. **Por segmento — Moda & Beleza**: CASE-009 a CASE-022
3. **Por segmento — Eletrônicos, Casa, Construção**: CASE-023 a CASE-032
4. **Por segmento — Autopeças, Indústria, Outros**: CASE-033 a CASE-040
5. **Por cenário — Churn e migrações de saída**: CASE-041 a CASE-046
6. **Por cenário — Prospects em negociação ativa (jun/2026)**: CASE-047 a CASE-050
7. **B2B, Afiliados e Verticais Especializadas (Editoras)**: CASE-051 a CASE-059
8. **Lote 2 — Casos descobertos no sweep 06/2026**: CASE-060 a CASE-072

---

## 1. Anchor Customers — Referências de Vendas

Estes são os casos citados explicitamente como **cases de venda** ou **clientes-âncora enterprise** em conversas internas. São os primeiros nomes a usar em pitches.

---

### CASE-001 — Aramis

| Campo | Valor |
|---|---|
| **Cliente** | Aramis (moda masculina premium) |
| **Setor** | moda masculina |
| **Status** | ✓ ativo (assinado 2026-06-02) |
| **Plataforma anterior** | VTEX |
| **ERP** | — |
| **Marketplace setup** | — |
| **Parceiros envolvidos** | Visie, Série//A, Oásis (agências concorrentes na negociação) |
| **Data da última menção** | 2026-06-03 |

**Situação:** Aramis é case-âncora da safra 2026 — moda masculina premium, vinda da VTEX. Negociação envolveu rodada de apresentações com 4 agências (Visie, Série//A, e outras). O líder tech da Aramis na VTEX certificou-se em Wake antes do fechamento e teria afirmado: *"A Wake é uma Ferrari perto da VTEX"* (referência à performance/arquitetura — citada em Tropa de Elite 3, 2026-05-22).

**Solução:** Wake Commerce + Wake OMS. Contrato fechado com **R$ 154k MRR + R$ 300k setup** (Comercial Cross WAKE, 2026-06-02). Ambiente Aramis e Corello criados em conjunto em 2026-05-19.

**Resultado:** Maior contrato da safra. Aramis é tratado como deal estratégico ("presente de Natal esse ano vai ser só da Aramis" — Comercial Cross WAKE, 2026-06-02). Ainda em onboarding técnico — EBI participando da volumetria de pedidos de franquias (2026-06-01).

**Insights / lessons:**
- Caso clássico de migração VTEX → Wake com diferencial de **performance**.
- Vale ressaltar o argumento da Série//A: diagnosticaram problemas reais da Aramis na VTEX antes da apresentação comercial.
- Para enterprise, o **endorsement do tech leader** antes do fechamento foi crítico.

---

### CASE-002 — Inbrands

| Campo | Valor |
|---|---|
| **Cliente** | Inbrands (grupo multimarca: Ellus, Richards, VR, Tommy Hilfiger BR, etc.) |
| **Setor** | moda — grupo multimarca |
| **Status** | ✓ ativo (cliente enterprise) |
| **Plataforma anterior** | — (cliente base Wake) |
| **ERP** | — |
| **Marketplace setup** | Anymarket + Privalia, TikTok Shop, Mercado Livre, Netshoes (em planejamento — 5 marcas, múltiplos CDs) |
| **Parceiros envolvidos** | CRMBonus, Conecta.la (em tratativas para marketplace) |
| **Data da última menção** | 2026-06-03 |

**Situação:** Cliente **enterprise** da Wake — tratado com cuidado especial pelo time CX Enterprise. Em 2026-04, abriu chamado #13460 sobre integração marketplace impactando **~10% do GMV (~R$ 1M)**. Em 2026-04-06, problema de cache no CRMBonus acumulando promoções indevidamente. Em 2026-04-28, status: "precisamos de alguém pra Inbrands agora" (escalação interna). Em 2026-06-03, ainda sem Head de E-commerce contratado (três candidatos rejeitados — Grupo Soma, Veste, outro).

**Solução:** Plataforma Wake Commerce + Anymarket para múltiplos canais. Múltiplos CDs precisam ser suportados. Marketplace é prioridade declarada.

**Resultado:** ⚠ inferido — operação ativa mas com fricções recorrentes envolvendo Anymarket. GMV impactado em chamados de marketplace tem sido tema constante.

**Insights / lessons:**
- Case usado em pitch: citado como uma das 3 referências em material institucional (Wake <> Catus, 2026-04-28: "Pura Vida, Soneda e Inbrands").
- Para enterprise com múltiplas marcas: **modelo de 1 conta única no marketplace** (Anymarket) ainda é a única arquitetura possível tecnicamente em 2025 (ANYMARKET & WAKE, 2025-07-23).
- Vacância de Head de E-commerce no cliente é risco de churn — manter relacionamento próximo do CXO enquanto não preenchem a vaga.

---

### CASE-003 — Karcher (Karcher Center Altex)

| Campo | Valor |
|---|---|
| **Cliente** | Karcher (linha de lavadoras / equipamentos) |
| **Setor** | eletroportáteis / equipamentos |
| **Status** | ✓ ativo |
| **Plataforma anterior** | — |
| **ERP** | Tiny |
| **Marketplace setup** | — |
| **Parceiros envolvidos** | Proex (agência), Leanwork (integração HERO Seguros), HERO Seguros (Garantia Estendida + Parcela Garantida) |
| **Data da última menção** | 2026-06-03 |

**Situação:** Karcher é o **primeiro cliente da integração HERO Seguros** dentro da Wake. Em 2026-03-16, publicado o produto **Parcela Garantida** da HERO. Em 2026-03-30, apresentada a entrega de **Garantia Estendida** funcionando em ambiente de testes. Em 2026-05-29 confirmado: prazo de garantia para Karcher é **12 meses** (formalizado por e-mail). Em 2026-06-03, testes prosseguindo.

**Solução:** Wake Commerce + ERP Tiny (nativo) + HERO Seguros (Garantia Estendida + Parcela Garantida via Leanwork como integradora). Domínio: `karcher-center-altex.com.br`.

**Resultado:** Karcher é case de **referência** para venda de seguros em Wake. Funcionou em ambiente de testes em 2026-03-30: *"Funcionou tudo certinho. Fácil de usar, fácil de explicar e fácil de vender..."* (WAKE | Time Parcerias).

**Insights / lessons:**
- Use Karcher para vender **seguro/garantia estendida** como upsell em verticais de eletro/equipamentos.
- ERP Tiny + Wake é integração **nativa** — vantagem comercial para SMB/mid-market.
- Cadastro de produto no painel é feito pela **agência (Proex) ou pela Wake**, não pelo integrador — ponto de atenção em projetos com múltiplos parceiros.

---

### CASE-004 — Shoulder

| Campo | Valor |
|---|---|
| **Cliente** | Shoulder (moda feminina premium) |
| **Setor** | moda feminina |
| **Status** | ✓ ativo (com case público de migração) |
| **Plataforma anterior** | — (migração — case interno) |
| **ERP** | — |
| **Marketplace setup** | Anymarket |
| **Parceiros envolvidos** | Figo (integração), Givex (gift card), DLK (agência) |
| **Data da última menção** | 2026-06-02 |

**Situação:** Shoulder é tratado como **case de migração público** da Wake. Em 2026-04-28: *"fiquei bem impressionado com o case que eles divulgaram da Shoulder que fala bastante especificamente sobre o processo de migração e os ganhos"* (Ecommerce House). Em 2026-05-15, RH organizou live com toda a LWSA para apresentar o case Shoulder, mediado pelo Pavoni — Comercial, Onboarding, Parcerias, CS apresentando cada etapa. Em 2026-06-02, fase de **MVP** com Figo encerrada após evento da Shoulder em 2026-06-01.

**Solução:** Wake Commerce + integração Figo (MVP) + Givex (gift card). Atende Shoulder + multimarcas do grupo.

**Resultado:** Case público de migração com **ganhos documentados** (PDF e live internos). Citado como prova de processo de migração assistida pela Wake.

**Insights / lessons:**
- Use o case Shoulder em apresentações de **migração** — há narrativa pronta sobre Comercial → Onboarding → Go-Live → CS.
- Combinação com Figo gera diferencial em projetos de moda multimarca.
- *Não confundir* "Shoulder" com "shoulder bag" em buscas no corpus — há muito ruído.

---

### CASE-005 — Oriba

| Campo | Valor |
|---|---|
| **Cliente** | Oriba (moda masculina) |
| **Setor** | moda masculina |
| **Status** | ✓ migrado (Go-Live em 2026-01-13) — depois `❌ churn` declarado em 2026-04-10 |
| **Plataforma anterior** | Shopify |
| **ERP** | — |
| **Marketplace setup** | — |
| **Parceiros envolvidos** | — |
| **Data da última menção** | 2026-06-03 |

**Situação:** Oriba migrou da Shopify para Wake e entrou ao ar em 2026-01-13 (Wake & Biso, 2026-01-29). Em 2026-03-12, confirmado: *"Oriba era Shopify"* (Comercial Cross WAKE). Case da Oriba foi publicado em PDF em 2026-05-04 destacando feature de **Kits com desconto progressivo** e resultados de SEO da migração (2026-03-31). Em 2026-04-10, mencionado como churn em Tropa de Elite 3. Em 2026-06-03, citado em retrospectiva de churns: *"se tivermos problemas com Lwsa como tivemos com Shoulder, Oriba e etc..."*

**Solução:** Wake Commerce + feature de Kits com desconto progressivo.

**Resultado:** Case com **ganhos de SEO** documentados após migração. Mas o cliente saiu posteriormente — ⚠ inferido: a saída pode estar relacionada a problemas com o grupo LWSA, não necessariamente com a plataforma em si.

**Insights / lessons:**
- Case de **migração de Shopify** com prova de ganho de SEO — útil em pitch contra Shopify.
- Cuidado: o cliente posteriormente saiu. Para uso em pitch público, validar se o case PDF ainda pode ser citado dado o churn.
- Feature **Kits com desconto progressivo** é diferencial em pitch para moda.

---

### CASE-006 — Pura Vida

| Campo | Valor |
|---|---|
| **Cliente** | Pura Vida |
| **Setor** | ⚠ status desconhecido (provável moda/beleza/cosméticos) |
| **Status** | ⚠ em risco — usado em pitch como case mas em 2026-05-11 sinalizado risco de perda |
| **Plataforma anterior** | — |
| **ERP** | — |
| **Marketplace setup** | — |
| **Parceiros envolvidos** | — |
| **Data da última menção** | 2026-05-11 |

**Situação:** Citado como uma das 3 referências em material institucional enviado por Catus (2026-04-28): *"apresentação institucional e 3 cases que conversam com as necessidades deles (Pura Vida, Soneda e Inbrands)"*. Em 2026-05-11: *"a gente já vai se ferrar com pura vida"* — sinal de churn iminente, mencionado junto com a defesa da conta Ybera.

**Solução:** ⚠ inferido — Wake Commerce.

**Resultado:** ⚠ status desconhecido — usado em pitch mas em risco.

**Insights / lessons:**
- Validar com account owner antes de usar como case ativo.
- Material institucional pode estar desatualizado se cliente sair.

---

### CASE-007 — Soneda

| Campo | Valor |
|---|---|
| **Cliente** | Soneda |
| **Setor** | ⚠ status desconhecido (provável varejo) |
| **Status** | ✓ ativo |
| **Plataforma anterior** | — |
| **ERP** | integrado via **Floui** |
| **Marketplace setup** | Anymarket — TikTok Shop |
| **Parceiros envolvidos** | Floui (integrador ERP) |
| **Data da última menção** | 2026-06-01 |

**Situação:** Cliente Wake com integração ERP via Floui. Em 2026-05-28, Floui reporta demanda aberta com Wake sobre integração de descontos entre Wake e Anymarket para o canal **TikTok Shop**. Gerente de e-commerce do Soneda (Vivian) sinalizou não estar recebendo retornos diretos. Em 2026-06-01, Soneda voltou a não retornar a Floui.

**Solução:** Wake Commerce + Floui (ERP) + Anymarket (marketplaces incluindo TikTok Shop).

**Resultado:** Operação rodando; relacionamento entre Wake CS e cliente apresentou ruídos em 2026-05/06 — ⚠ inferido como fricção operacional, não churn.

**Insights / lessons:**
- Usado como case em material institucional (Wake <> Catus, 2026-04-28).
- Bom exemplo de stack **Wake + Floui (ERP) + Anymarket (canais)** para empresa multi-canal.
- Demonstra que descontos cross Wake↔Anymarket↔canal específico (TikTok) ainda exigem demanda customizada.

---

### CASE-008 — Tommy (Tommy Hilfiger BR)

| Campo | Valor |
|---|---|
| **Cliente** | Tommy (provável Tommy Hilfiger BR, parte do Grupo Inbrands) |
| **Setor** | moda masculina premium |
| **Status** | ✓ ativo (assinado 2026-05-27) |
| **Plataforma anterior** | — |
| **ERP** | — |
| **Marketplace setup** | — |
| **Parceiros envolvidos** | Sellbie (Go-Live previsto 2026-07-20) |
| **Data da última menção** | 2026-05-28 |

**Situação:** Contrato Tommy assinado em 2026-05-27 — Commerce + OMS. Em 2026-05-28, Sellbie reporta previsão de Go-Live para 2026-07-20.

**Solução:** Wake Commerce + Wake OMS. **R$ 29,4k MRR + R$ 179,7k setup** (Comercial Cross WAKE, 2026-05-27).

**Resultado:** Em onboarding técnico.

**Insights / lessons:**
- Possível **expansão dentro do Grupo Inbrands** (se confirmado o vínculo) — pode reforçar o case Inbrands como guarda-chuva.
- Tickets de setup acima de R$ 100k indicam projeto enterprise.

---

## 2. Por Segmento — Moda & Beleza

### CASE-009 — Corello

| Campo | Valor |
|---|---|
| **Cliente** | Corello (calçados femininos) |
| **Setor** | calçados / acessórios moda |
| **Status** | ✓ ativo (assinado 2026-05-29) |
| **Plataforma anterior** | — |
| **ERP** | — |
| **Parceiros envolvidos** | — (ambiente criado em conjunto com Aramis) |
| **Data da última menção** | 2026-06-02 |

**Situação:** Corello assinado em 2026-05-29 (Commerce + OMS, **R$ 50,6k MRR + R$ 73k setup**). Ambiente Corello + Aramis criados juntos em 2026-05-19 — duas marcas em paralelo.

**Solução:** Wake Commerce + OMS.

**Insights / lessons:** Stack idêntica à da Aramis — útil em pitches de moda/calçados premium. Ambientes paralelos em setup indicam capacidade de onboarding paralelo.

---

### CASE-010 — Cless

| Campo | Valor |
|---|---|
| **Cliente** | Cless (cosméticos capilares) |
| **Setor** | cosméticos / beleza |
| **Status** | ✓ ativo (assinado 2025-10-30, Go-Live 2026-02-16) |
| **Plataforma anterior** | — |
| **ERP** | — |
| **Parceiros envolvidos** | — |
| **Data da última menção** | 2026-01-29 |

**Situação:** Cless assinada em 2025-10-30 (**R$ 20k MRR + R$ 9k setup**). Go-Live confirmado para 2026-02-16 na lista de próximos Go-Lives do parceiro Biso (Wake & Biso, 2026-01-29).

**Solução:** Wake Commerce.

**Insights / lessons:** Case de **cosméticos** com onboarding em ~4 meses (assinatura → go-live). Útil para pitch em vertical beleza.

---

### CASE-011 — Adaptogen

| Campo | Valor |
|---|---|
| **Cliente** | Adaptogen |
| **Setor** | ⚠ inferido — suplementos / wellness |
| **Status** | ✓ ativo (assinado 2026-03-27) |
| **Plataforma anterior** | — |
| **Parceiros envolvidos** | — |
| **Data da última menção** | 2026-03-27 |

**Situação:** Assinado 2026-03-27. **R$ 25,5k MRR**.

**Solução:** Wake Commerce.

**Insights / lessons:** Ticket sólido em vertical de **wellness/suplementos**; vale validar a vertical com account owner.

---

### CASE-012 — LP Beauty

| Campo | Valor |
|---|---|
| **Cliente** | LP Beauty |
| **Setor** | beleza/cosméticos |
| **Status** | ❌ churn (rollback para Bagy) |
| **Plataforma anterior** | Wake |
| **Plataforma atual** | Bagy (rollback) |
| **Parceiros envolvidos** | Biso (tentativa de implantação que não chegou a contrato definitivo) |
| **Data da última menção** | 2026-04-30 |

**Situação:** Em 2026-04-29: *"esta usando, mas o cs de voces informou que eles darao / deram churn na wake, parece que deram roll back para bagy"* (Wake & Biso). Em 2026-04-13: *"Lp beuaty q deu churn podia ser desconvidada do evento de quinta"*. Em 2026-04-14, ainda tentando ativar Apple Pay via PayPal — sinal de operação parcialmente ativa em transição.

**Solução / Resultado:** Churn confirmado. Rollback para Bagy (plataforma SMB do grupo LWSA).

**Insights / lessons:**
- Único caso documentado de **churn de Wake → Bagy** no corpus.
- Útil contexto para diferenciar Wake (mid-enterprise) de Bagy (SMB) dentro do próprio grupo LWSA — clientes podem fazer movimento de "downgrade" interno.
- Cuidado: a Bagy foi anunciada em 2026-04-12 como **descontinuada como tech** (Tropa de Elite 3). Esse rollback pode ser revertido.

---

### CASE-013 — Mamô

| Campo | Valor |
|---|---|
| **Cliente** | Mamô |
| **Setor** | ⚠ moda/lifestyle (verificar) |
| **Status** | ✓ ativo (assinado 2026-04-30) |
| **Parceiros envolvidos** | — |

**Situação:** Assinado 2026-04-30. Commerce + OMS, **R$ 35,6k MRR + R$ 60k setup**.

**Insights / lessons:** Ticket Commerce+OMS combinado — case útil para pitch de oferta integrada.

---

### CASE-014 — Letty

| Campo | Valor |
|---|---|
| **Cliente** | Letty |
| **Status** | ✓ ativo (assinado 2026-04-30) |

**Situação:** Commerce, **R$ 22k MRR + R$ 15k setup**.

---

### CASE-015 — Lizie

| Campo | Valor |
|---|---|
| **Cliente** | Lizie |
| **Status** | ✓ ativo (assinado 2026-03-25) |

**Situação:** Commerce, **R$ 33,9k MRR + R$ 60k setup**.

---

### CASE-016 — Cacay

| Campo | Valor |
|---|---|
| **Cliente** | Cacay |
| **Status** | ✓ ativo (assinado 2026-05-25) |
| **Marketplace setup** | Anymarket + Dafiti (migração) |
| **Parceiros envolvidos** | — |
| **Data da última menção** | 2026-05-19 |

**Situação:** Experience assinado 2026-05-25 (**R$ 2,9k MRR**). Em 2026-05-19, em migração para Dafiti com erro no DE/PARA de frete (ticket Anymarket #22241).

**Insights / lessons:** Case operacional ativo com Dafiti — útil para mostrar suporte cross Wake/Anymarket na fase de go-live em marketplace.

---

### CASE-017 — Baw

| Campo | Valor |
|---|---|
| **Cliente** | Baw |
| **Status** | ✓ ativo (assinado 2026-05-28) |

**Situação:** Wake Experience, **R$ 9,6k MRR + R$ 4k setup**.

---

### CASE-018 — Sumirê

| Campo | Valor |
|---|---|
| **Cliente** | Sumirê |
| **Status** | ✓ ativo (assinado 2026-04-09) |

**Situação:** Wake Experience, **R$ 3k MRR**.

---

### CASE-019 — Wamp Biquinis

| Campo | Valor |
|---|---|
| **Cliente** | Wamp Biquinis |
| **Setor** | moda praia |
| **Status** | ✓ ativo (assinado 2026-05-05) |

**Situação:** Commerce, **R$ 3,8k MRR**.

---

### CASE-020 — Ello Moda Fitness e Praia

| Campo | Valor |
|---|---|
| **Cliente** | Ello Moda Fitness e Praia |
| **Setor** | moda fitness/praia |
| **Status** | ⚠ em negociação |
| **Plataforma anterior** | (sem e-commerce — 1 loja no Brás, atacado+varejo) |
| **Parceiros envolvidos** | Nairuz, 3W, Omni360+Affys (agências em rodada) |
| **Data da última menção** | 2026-05-08 |

**Situação:** Marca de moda fitness/praia nascida em dezembro 2025. Distribuída a múltiplas agências em 2026-05-08. Foco: colocar site no ar com segurança.

**Insights / lessons:** Caso típico de **prospect SMB** de moda emergente — segmento fitness/praia tem ciclo curto e necessidade de time-to-market rápido.

---

### CASE-021 — ACC Modas

| Campo | Valor |
|---|---|
| **Cliente** | ACC Modas |
| **Setor** | moda/varejo feminino |
| **Status** | ⚠ em negociação |
| **Plataforma atual** | (foco em estruturar e-commerce próprio — 11 lojas físicas, ~R$ 200k/mês por loja) |
| **Parceiros envolvidos** | Driven, Cava, N1 (agências em rodada) |
| **Data da última menção** | 2026-05-14 |

**Situação:** Fabricante de moda feminina B2C. Opera com 11 lojas físicas. Distribuído em 2026-05-14 a múltiplas agências.

**Insights / lessons:** **Varejo multi-loja** sem e-commerce próprio — típico prospect de integração POS + online (cliente precisa OMS).

---

### CASE-022 — Lenscope

| Campo | Valor |
|---|---|
| **Cliente** | Lenscope (lentes corretivas online) |
| **Setor** | ótica / D2C digital |
| **Status** | ⚠ em negociação |
| **Plataforma anterior** | — (empresa 100% digital, sem lojas físicas) |
| **Parceiros envolvidos** | B8One, Quick Digital, Duo (agências em rodada) |
| **Data da última menção** | 2026-05-29 |

**Situação:** Empresa 100% digital de Campinas, faturamento R$ 200k/mês. Distribuída em 2026-05-29.

**Insights / lessons:** **D2C 100% digital** em vertical regulada (ótica) — exemplo de prospect onde a Wake compete por escala/performance.

---

## 3. Por Segmento — Eletrônicos, Casa, Construção

### CASE-023 — Casa Almeida

| Campo | Valor |
|---|---|
| **Cliente** | Casa Almeida (Buddemeyer) |
| **Setor** | cama/mesa/banho (varejo) |
| **Status** | ✓ ativo (Go-Live 2026-02-02) |
| **Marketplace setup** | Anymarket + suporte assistido (~R$ 1k–1,5k extra) |
| **Parceiros envolvidos** | Anymarket |
| **Data da última menção** | 2026-05-11 |

**Situação:** Go-Live em 2026-02-02 (Wake & Biso). Em 2026-03-05, em discussão suporte assistido pela Anymarket para clientes com volume — citado custo de R$ 1k/R$ 1,5k/mês. Em 2026-05-11, falou no painel do evento Wake On the Road em Curitiba (junto com Balaroti e Pavoni).

**Insights / lessons:** Casa Almeida usa Anymarket. Boa **referência regional** (Sul) para verticais de casa/decor.

---

### CASE-024 — Balaroti

| Campo | Valor |
|---|---|
| **Cliente** | Balaroti |
| **Setor** | material de construção |
| **Status** | ✓ ativo |
| **Plataforma anterior** | — |
| **Parceiros envolvidos** | Serasa Pass (ativação concluída em 2026-05-07), B8One (front-end histórico — saiu do programa) |
| **Data da última menção** | 2026-05-20 |

**Situação:** Em 2026-03-16, em fase de instalação do plugin Serasa Pass. Em 2026-05-07, *"Temos a Balaroti fechada"* (Serasa & Wake). Em 2026-05-20, ajuste de front para ativar Serasa demandou contato com B8One (agência antiga que saiu do programa).

**Insights / lessons:**
- Case de **integração com Serasa Pass** para material de construção (ticket alto, financiamento relevante).
- Falou no painel do evento Wake On the Road em Curitiba (2026-05-11).
- Demonstra que **ex-parceiros (agências fora do programa)** ainda podem ser necessários em projetos legados — manter relacionamento de transição.

---

### CASE-025 — Pavoni (homonímia — verificar)

| Campo | Valor |
|---|---|
| **Cliente** | Pavoni |
| **Setor** | ⚠ status desconhecido — vertical não confirmada |
| **Status** | ✓ ativo |
| **Data da última menção** | 2026-05-11 |

**Situação:** Falou no painel do evento Wake On the Road em Curitiba (2026-05-11) junto com Balaroti e Casa Almeida.

⚠ Cuidado: "Pavoni" também é o sobrenome do executivo Wake responsável por migrações — verificar com account owner se há cliente "Pavoni" como retailer ou se a referência no painel é apenas ao executivo.

---

### CASE-026 — Verdes Mares

| Campo | Valor |
|---|---|
| **Cliente** | Verdes Mares (material de construção — interior da Bahia) |
| **Setor** | material de construção / construtoras |
| **Status** | ⚠ em negociação |
| **Plataforma anterior** | — (sem operação digital — 1 loja física + 1 CD; nova loja prevista para julho/2026) |
| **Parceiros envolvidos** | Wicomm, Eficaz, Avocado Tech, Catus (agências em rodada) |
| **Data da última menção** | 2026-05-21 |

**Situação:** Material de construção do interior da BA. Foco B2B (construtoras, arquitetos, revendedores). Distribuído em 2026-05-21.

**Insights / lessons:** Prospect B2B regional — exemplo de venda para projetos com **componente B2B forte** (não puramente B2C).

---

### CASE-027 — LEDVANCE (ex-Osram)

| Campo | Valor |
|---|---|
| **Cliente** | LEDVANCE (oriundo da Osram) |
| **Setor** | iluminação / eletroeletrônicos (nova frente) |
| **Status** | ⚠ em negociação |
| **Plataforma anterior** | — (primeira operação digital da nova linha B2C) |
| **ERP** | SAP ECC 6.0 (upgrade para S/4HANA planejado para 2º semestre 2027) |
| **OMS** | não possui |
| **Parceiros envolvidos** | Cava Digital, Avocado Tech, Catus (agências em rodada) |
| **Data da última menção** | 2026-06-03 |

**Situação:** Grupo de 120 anos. Faturamento referencial ~R$ 30-35M/ano em euros (iluminação). Nova frente B2C em eletroeletrônicos. Distribuído em 2026-06-03.

**Insights / lessons:**
- Cliente **enterprise SAP** — case típico para integradores SAP (Floui, Iota.Hub conforme playbook §3.1).
- Início da operação digital — projeto greenfield com timeline alinhado ao upgrade S/4 (planejar arquitetura para 2027).
- Oportunidade de pitch de Wake + OMS combinado.

---

### CASE-028 — Hiper Equipamentos

| Campo | Valor |
|---|---|
| **Cliente** | Hiper Equipamentos |
| **Setor** | equipamentos |
| **Status** | ✓ ativo (em onboarding marketplace) |
| **Marketplace setup** | Anymarket — múltiplos sellers, multi-CD com crossdocking |
| **Data da última menção** | 2026-04-09 |

**Situação:** Em 2025-09, classificado como crítico — *"Seller quis ser o último a integrar e batemos ontem no impasse de gestão de crossdocking"*. Em 2026-02-10, problemas com XML desde 10/02 (ticket 2311). Em 2026-03-11, *"cliente crítico para nós"*. Em 2026-04-09, problemas seguindo.

**Insights / lessons:** Case operacional difícil — **multi-CD com crossdocking** em Anymarket é cenário de complexidade alta. Usar como referência de capacidade de atender complexidade operacional.

---

### CASE-029 — Empório do Celular

| Campo | Valor |
|---|---|
| **Cliente** | Empório do Celular |
| **Setor** | eletrônicos / acessórios celular |
| **Status** | ⚠ em negociação |
| **Plataforma anterior** | Tray |
| **ERP** | Tiny |
| **GMV** | R$ 15k/mês site + R$ 150k/mês em marketplace |
| **Parceiros envolvidos** | O2 Consultoria, Driven, Savvi (agências em rodada) |
| **Data da última menção** | 2026-04-09 |

**Situação:** *"Cliente Tray, deseja migrar porque sente que a Tray é limitada, e deseja na próxima plataforma liberdade de customização"* (O2 Consultoria, 2026-04-09). 12 anos de mercado. Vendem Huawei e marcas próprias.

**Insights / lessons:**
- Migração **Tray → Wake** clássica: motivação é **customização**.
- ERP Tiny + Wake = integração nativa (vantagem comercial).
- GMV digital pequeno (R$ 15k site) — projeto de reativação do site (estavam apenas em marketplace).

---

### CASE-030 — Wallford Brasil

| Campo | Valor |
|---|---|
| **Cliente** | Wallford Brasil (calçados sociais/casuais — multi-marcas em planejamento) |
| **Setor** | calçados |
| **Status** | ⚠ em negociação (sem retorno desde jan/2026 — risco churn de lead) |
| **Plataforma de referência** | VTEX (cliente cita politica comercial VTEX como modelo desejado) |
| **ERP** | Tiny |
| **GMV** | R$ 150k/mês |
| **Parceiros envolvidos** | Avocado Tech (agência), O2 Consultoria (canal) |
| **Data da última menção** | 2026-01-26 |

**Situação:** Lead trazido por O2 Consultoria em 2025-12-10. Plano: 5 marcas com fronts diferentes, catálogo único. Em 2025-12-22, dúvidas técnicas (arquitetura multi-marca/multi-front, análises de performance centralizadas). Em 2026-01-26, sem retorno — cliente disse estar interagindo com outra agência.

**Insights / lessons:**
- Pergunta-modelo de **arquitetura multi-marca** (catálogo único, front por marca) — a VTEX é benchmark com "politica comercial".
- Lead que esfria sem resposta é padrão de **conflito de canal** — discutir com Partner Ops.

---

### CASE-031 — Norte Refrigeração

| Campo | Valor |
|---|---|
| **Cliente** | Norte Refrigeração |
| **Setor** | refrigeração comercial / equipamentos |
| **Status** | ⚠ status desconhecido |
| **Setup particular** | "Não é seller, é multiCD, mas opera igual um" (Comercial Cross WAKE, 2026-03-31) |
| **Data da última menção** | 2026-03-31 |

**Situação:** Cliente com particularidade operacional: multi-CD que opera como seller (não é marketplace, mas usa estrutura similar).

**Insights / lessons:** Case interessante para conversas de **multi-CD** quando o cliente não é marketplace mas tem complexidade similar.

---

### CASE-032 — Empório do Sertão / "Empório"

| Campo | Valor |
|---|---|
| **Cliente** | Empório (do Sertão) — RJ |
| **Setor** | varejo físico (5 lojas no RJ) |
| **Status** | ⚠ em negociação |
| **Faturamento offline** | R$ 20M |
| **Dor** | experiência anterior com e-commerce frustrada; operação atual manual via WhatsApp; ruptura de estoque |
| **Parceiros envolvidos** | Dotkom, Eficaz, Catus, Nairuz (agências em rodada) |
| **Data da última menção** | 2026-05-06 |

**Situação:** Projeto **omnichannel** — quer centralizar e usar estoque das 5 lojas para vender online em escala.

**Insights / lessons:**
- Case de **omnichannel** clássico — cliente físico com pain de e-commerce anterior.
- Distribuição multi-agência indica deal "share o lead, melhor agência leva".

---

## 4. Por Segmento — Autopeças, Indústria, Outros

### CASE-033 — Off Rush

| Campo | Valor |
|---|---|
| **Cliente** | Off Rush |
| **Setor** | autopeças / motos |
| **Status** | ⚠ em negociação |
| **Faturamento** | R$ 400k/mês (1% via site) |
| **Plataforma anterior** | (forte em marketplaces, site representa 1% do faturamento) |
| **Parceiros envolvidos** | Duo, Catus, Avocado Tech (agências em rodada) |
| **Data da última menção** | 2026-05-05 |

**Situação:** Distribuído em 2026-05-05. Vertical autopeças/motos. Atende Vale do Paraíba.

**Insights / lessons:** **Autopeças** com forte canal marketplace e site irrelevante — oportunidade clássica de **virar o jogo** com plataforma performante + integração marketplace via Anymarket.

---

### CASE-034 — Carblue

| Campo | Valor |
|---|---|
| **Cliente** | Carblue (Campinas) |
| **Setor** | ⚠ provável autopeças (verificar) |
| **Status** | ✓ ativo (cliente base) |
| **Parceiros envolvidos** | indicação de Carblue → Catus / Nairuz para sustentação |
| **Data da última menção** | 2026-05-19 |

**Situação:** Cliente da base Wake que pediu agência regional (Campinas) para sustentação e ajustes. Indicação dada a Catus e Nairuz em 2026-05-18.

**Insights / lessons:** **Sustentação regional** é demanda comum em base — manter mapa de agências por região.

---

### CASE-035 — UncleK

| Campo | Valor |
|---|---|
| **Cliente** | UncleK |
| **Setor** | ⚠ status desconhecido |
| **Status** | ✓ ativo |
| **Parceiros envolvidos** | Floui (ERP), Naper (integração) |
| **Data da última menção** | 2026-06-01 |

**Situação:** Em 2026-05-08, *"Naper já fez integração conosco e UncleK está usando"*. Em 2026-05-14, demanda relacionada a UncleK precisava ser endereçada por CS. Em 2026-06-01, Floui reporta que UncleK não retorna mais.

**Insights / lessons:** Cliente em stack **Wake + Floui + Naper**. Sinal de falta de engajamento — risco em monitoring.

---

### CASE-036 — Soubeef

| Campo | Valor |
|---|---|
| **Cliente** | Soubeef |
| **Setor** | ⚠ provável food/grocery (carne) |
| **Status** | ✓ ativo (assinado) |

**Situação:** Listado entre os assinados (Comercial Cross WAKE). Valores não detalhados nas extrações.

---

### CASE-037 — Yamaha Motor

| Campo | Valor |
|---|---|
| **Cliente** | Yamaha Motor Brasil |
| **Setor** | veículos / motos |
| **Status** | ✓ ativo (Go-Live 2026-02-26) |
| **Parceiros envolvidos** | Biso (search) |
| **Data da última menção** | 2026-05-29 |

**Situação:** Listado entre os próximos Go-Lives em 2026-01-29 com Go-Live previsto para 2026-02-26. Em 2026-05-12, Biso aguardando confirmação de e-mails dos representantes da Yamaha para avançar tratativa de search.

**Insights / lessons:** Case **enterprise indústria automotiva** — relevante para venda em verticais industriais com canal de venda direta + concessionária.

---

### CASE-038 — Estoca (parceiro logístico — citado como exemplo de operação)

⚠ **Não é customer case** — Estoca é parceiro de fulfillment listado em playbook §2 (`parceiro-logistica`). Citado aqui apenas para esclarecer que o nome aparece no DB no contexto de venda de fulfillment para outros clientes (ex.: Baly, Ybera). Omitido do catálogo.

---

### CASE-039 — Jacto

| Campo | Valor |
|---|---|
| **Cliente** | Jacto |
| **Setor** | máquinas agrícolas / indústria |
| **Status** | ✓ ativo (assinado 2025-10-01) |

**Situação:** Commerce, **R$ 28,1k MRR + R$ 130k setup** (Comercial Cross WAKE, 2025-10-01).

**Insights / lessons:** Case **enterprise indústria agrícola** — ticket de setup alto (R$ 130k) sinaliza projeto de complexidade técnica relevante.

---

### CASE-040 — Garage

| Campo | Valor |
|---|---|
| **Cliente** | Garage |
| **Setor** | ⚠ status desconhecido |
| **Status** | ✓ ativo (assinado 2026-04-02) |

**Situação:** Commerce + OMS, **R$ 42,2k MRR + R$ 80k setup**.

---

## 5. Por Cenário — Churn e Migrações de Saída

### CASE-041 — Mpozenato (churn massivo)

| Campo | Valor |
|---|---|
| **Cliente** | Mpozenato |
| **Setor** | ⚠ status desconhecido |
| **Status** | ❌ churn — perda de **R$ 1,8M** |
| **Setup particular** | "quase na totalidade em Marketplace. E contrato era baixo, valor fixo por pedido" (Tropa de Elite 3, 2026-05-27) |
| **Data da última menção** | 2026-05-27 |

**Insights / lessons:** Modelo de contrato **valor fixo por pedido** em cliente marketplace-heavy é estrutura de receita frágil — quando o cliente sai, perda é grande mas margem por pedido era baixa. Caso revisitado para discussão de modelo comercial.

---

### CASE-042 — E-plus (agência / também relacionamento de churn)

⚠ **Atenção**: E-plus é **agência**, não cliente final. Citada como exemplo de **churn de agência** ("E-plus deu churn conosco" — WAKE | Time Parcerias, 2026-04-29) e o CEO Tiago foi para a VTEX como presidente. Histórico de fricção: *"Só tivemos dor de cabeça com e-Plus + Deco"*. Não é customer case — mantida aqui só como nota de contexto.

---

### CASE-043 — Ricardo Almeida (churn pós Go-Live curto)

| Campo | Valor |
|---|---|
| **Cliente** | Ricardo Almeida (moda masculina premium) |
| **Setor** | moda masculina premium |
| **Status** | ❌ churn (solicitado em 2025-10-29) |
| **Assinatura** | 2025-08-01 — Experience (R$ 5,8k MRR + R$ 20,3k setup) |
| **Plataforma de saída** | Shopify (decisão TI + franquia TOTVS Ibirapuera / 2B Digital) |
| **Data da última menção** | 2026-06-03 |

**Situação:** Assinou Wake Experience em 2025-08-01. Solicitou churn em 2025-10-29 — 3 meses depois. Foram para Shopify por conta da franquia da TOTVS Ibirapuera (2B Digital), e houve "rolo com nosso CRM". Em 2026-04-12, contexto detalhado: *"São TOTVS Moda. TI decidiu sozinho essa migração"*.

**Insights / lessons:**
- Cenário de **churn por decisão técnica isolada** (TI decidiu sem alinhamento comercial).
- Risco de **acoplamento ERP↔Plataforma** quando a franqueada TOTVS pressiona por stack próprio.
- Possível recuperação futura: em 2026-06-03 ainda mencionado *"DLK e Ricardo Almeida não abriram"* — pode haver janela.

---

### CASE-044 — DLK Modas (migração para Shopify)

| Campo | Valor |
|---|---|
| **Cliente** | DLK Modas |
| **Setor** | moda feminina |
| **Status** | ❌ churn — migrando para Shopify (confirmado 2026-05-25) |
| **Parceiros envolvidos** | Eitri (mantém app Wake até final), Vindi (cliente quer manter Vindi na Shopify) |
| **Data da última menção** | 2026-06-03 |

**Situação:** Em 2026-05-11, problemas com recaptcha no app travando vendas. Em 2026-05-25 confirmado: *"DLK está indo para Shopify mesmo, falei aqui com Monisi por conta da Vindi, querem manter a Vindi la"*. Em 2026-06-03: *"DLK e Ricardo Almeida não abriram"*.

**Insights / lessons:**
- Case de **churn para Shopify** com manutenção do parceiro de pagamento (Vindi). Vindi é stack-agnóstico no relacionamento com cliente final.
- DLK também é referenciada no corpus como agência de moda (Wake & Eitri) — verificar dupla identidade.

---

### CASE-045 — Mimeria (churn possível)

| Campo | Valor |
|---|---|
| **Cliente** | Mimeria |
| **Status** | ❌ churn iminente (2026-02-06: "Possível Churn") |
| **Parceiros envolvidos** | CRMBonus |
| **Data da última menção** | 2026-02-06 |

**Situação:** *"Atenção: Possível Churn Mimeria. Cliente está solicitando contato de outros parceiros..."* (Wake + CRMBonus | Parcerias, 2026-02-06).

**Insights / lessons:** Sinal precoce: cliente pedindo contato de **parceiros concorrentes** (de CRMBonus) é red flag de churn.

---

### CASE-046 — Cadiveu (saída em curso)

| Campo | Valor |
|---|---|
| **Cliente** | Cadiveu (cosmética capilar) |
| **Status** | ⚠ saída em curso (2025-07-22: "Fiquei sabendo que estão saindo, mas ainda é Wake") |
| **Data da última menção** | 2026-02-10 |

**Insights / lessons:** Vertical **cosmética capilar premium** — cliente comparável a Cless (que ainda é base ativa).

---

## 6. Prospects em Negociação Ativa — Junho/2026

### CASE-047 — Ybera

| Campo | Valor |
|---|---|
| **Cliente** | Ybera (cosmética/produtos para cabelo) |
| **Setor** | cosméticos capilares |
| **Status** | ⚠ em negociação concorrida com VTEX |
| **Plataforma anterior** | — |
| **Parceiros envolvidos** | Uble (queria construir como primeiro case Wake) |
| **Data da última menção** | 2026-06-03 |

**Situação:** Em 2026-05-11, sinal de alerta: *"a gente não pode perder Ybera de jeito nenhum"*. Em 2026-05-29, Uble afirma: *"A Ybera tem potencial para ser o nosso primeiro case Wake"* (Uble & Wake). Em 2026-06-01, *"comentaram q estão com uma POC com Vtex"*. Em 2026-06-03: *"Cara, Ybera a VTEX tá batendo lá"* + *"Logo temos um detrator declarado lá dentro"* (Vinícius, diretor de tech, ex-founder de empresa comprada pela Ybera, agora em Vitória). Wake também combinando POC de 6 dias com possível continuidade paga.

**Insights / lessons:**
- Deal **competitivo direto com VTEX** em jun/2026.
- Stakeholder técnico interno (Vinícius) é **detrator** — situação política a navegar.
- Considerado deal estratégico ("não pode perder de jeito nenhum") — escalado a alta liderança.

---

### CASE-048 — Palácio das Essências

| Campo | Valor |
|---|---|
| **Cliente** | Palácio das Essências |
| **Setor** | essências / cosméticos / perfumaria |
| **Status** | ⚠ em negociação |
| **Faturamento online** | R$ 60-80k/mês (varia conforme fonte) |
| **Setup particular** | 6 lojas físicas; linha B2B/atacado; marca automotiva paralela |
| **Parceiros envolvidos** | Série//A (passagem de bastão para agências em 2026-05-08) |
| **Data da última menção** | 2026-05-08 |

**Insights / lessons:** Cliente **multi-segmento** (B2C + B2B + automotivo) — necessidade de arquitetura multi-marca/multi-loja.

---

### CASE-049 — Verdes Mares + outros prospects em rodada multi-agência

Casos abertos em 2026-05 distribuídos entre múltiplas agências da carteira. Veja **CASE-026** (Verdes Mares), **CASE-020** (Ello Moda), **CASE-021** (ACC Modas), **CASE-022** (Lenscope), **CASE-027** (LEDVANCE), **CASE-032** (Empório do Sertão), **CASE-033** (Off Rush). Padrão: oportunidade é distribuída a 3-5 agências da carteira; melhor proposta leva.

**Insights / lessons:**
- Modelo **carrossel de agências** é o padrão para SMB/mid-market em 2026.
- Para Sales: avisar PAM logo no primeiro contato para evitar conflito.
- Para PAM: monitorar a velocidade de resposta das agências — Wake distribuiu, mas o cliente só fecha com 1.

---

### CASE-050 — Protwins (referência cruzada)

| Campo | Valor |
|---|---|
| **Cliente** | Protwins (`protwins.com.br`) |
| **Setor** | ⚠ status desconhecido |
| **Status** | ⚠ em qualificação |
| **Faturamento** | R$ 90k/mês (< R$ 1M/ano) |
| **Plataforma anterior** | Wbuy |
| **Parceiros envolvidos** | Yuri (ProEx) - indicação |
| **Data da última menção** | 2026-05-19 |

**Situação:** *"Protwins. Está hoje na Wbuy. Estou indicando migrar para Wake, marquei 1a conversa com o Yuri da ProEx. Quer mandar alguém junto ou conta muito pequena?"* (Coder Ivy + Wake, 2026-05-19). Em paralelo está sendo indicado VTEX com agência Quick.

**Insights / lessons:**
- Migração **Wbuy → Wake** é cenário recorrente para SMB; competição é VTEX.
- Conta de R$ 90k/mês está no limite SMB/mid-market — decidir se vale tempo do time comercial.

---

## Resumo executivo

### Por status

| Status | Quantidade | Casos |
|---|---|---|
| ✓ ativo (com contrato assinado e/ou Go-Live confirmado) | 30 | Aramis, Inbrands, Karcher, Shoulder, Soneda, Tommy, Corello, Cless, Adaptogen, Mamô, Letty, Lizie, Cacay, Baw, Sumirê, Wamp Biquinis, Casa Almeida, Balaroti, Hiper Equipamentos, UncleK, Yamaha, Jacto, Garage, Esplane, CVC, Yora, Destro, Gold, A Fábrica, Unitoys, Editora Fórum |
| ⚠ B2B em onboarding (B2C ativo) | 1 | Buddemeyer |
| ⚠ em negociação | 12 | Ello Moda, ACC Modas, Lenscope, Verdes Mares, LEDVANCE, Empório do Sertão, Off Rush, Wallford, Empório do Celular, Ybera, Palácio das Essências, Protwins |
| ✓ migrado | 2 | Oriba (Shopify → Wake), depois churn |
| ❌ churn | 6 | LP Beauty, Mpozenato, Ricardo Almeida, DLK Modas, Mimeria (iminente), Cadiveu (em curso) |
| ⚠ status desconhecido | 5 | Pura Vida, Norte Refrigeração, Pavoni (a verificar), Soubeef, Carblue |
| not customer case (omitidos) | 3 | Estoca (parceiro), E-plus (agência), Movimento (genérico) |

### Top 5 por frequência de menção no corpus

1. **Shoulder** (311 hits — em parte ruído com "shoulder bag") — anchor case migração
2. **Karcher** (127 hits) — case HERO Seguros
3. **Inbrands** (96 hits) — enterprise CRMBonus + Anymarket
4. **Riachuelo** (72 hits — canal marketplace, **não cliente Wake**)
5. **Estoca** (71 hits — parceiro de fulfillment, **não cliente Wake**)
6. **Aramis** (67 hits) — anchor case migração VTEX

### Distribuição por segmento

- **Moda & Beleza**: 17 casos (Aramis, Corello, Shoulder, Inbrands, Cless, LP Beauty, Mamô, Letty, Lizie, Cacay, Baw, Sumirê, Wamp Biquinis, Ello Moda, ACC Modas, Oriba, DLK Modas, Ricardo Almeida)
- **Eletrônicos / Casa / Construção**: 11 casos
- **Autopeças / Indústria**: 8 casos
- **Outros / Status desconhecido**: 5+
- **B2B (atacado / franquias / matriz↔revenda)**: 8 casos (Esplane, CVC, Yora, Destro, Buddemeyer B2B, Gold, A Fábrica, Unitoys) — ver §7
- **Editoras / Publishers**: 1 caso (Editora Fórum) — ver §7

---

## B2B, Afiliados e Verticais Especializadas

Estes casos cobrem duas jornadas adicionais e uma vertical nova que estavam ausentes do catálogo até a versão anterior:

1. **B2B (atacado / matriz↔franquia / indústria↔revendedor)** — a Wake suporta a mesma loja operando em duas frentes ao mesmo tempo: URL pública B2C + área logada B2B com listas de preço, catálogo, banners e regras de pagamento próprias. A frente B2B é citada como **"estrutura bem completa"** com crédito, regionalização, mix de produto e listas de preço diferentes (CoreBiz & Wake, 2026-01-30: *"Hoje temos uma estrutura de B2B bem completa... Temos Buddemeyer e CVC como cases legais de benchmark"*). Os cases-âncora oficiais de B2B citados internamente são **CVC, GOLD, Destro e A Fábrica** (Comercial Cross WAKE, 2026-03-05: *"Quais outros cases que temos de B2B, além de: CVC, GOLD, Destro e A Fabrica?"*). A lista completa de clientes B2B Wake foi consolidada em duas mensagens — Comercial Cross WAKE 2025-12-17 e Wake <> Nairuz 2026-03-20, ambas com texto quase idêntico — e é a base deste capítulo.

2. **Programas de afiliados / influencer marketing** — não é um produto Wake separado; é um padrão de uso da plataforma. A referência interna canônica está em **Comercial Cross WAKE, 2025-12-17**: *"Ybera / Shoulder para funcionários / BAW com influenciadores"*. Tecnicamente, programas de afiliados são montados via **StoreFront API** (mutations `CheckoutClone` para clonar carrinho mantendo o `partner/afiliado` + `CheckoutAddMetadata` para UTM — Uble & Wake, 2026-01-19) e via **módulo de campanha/cupom** do painel Wake (Monitfy & Wake, 2026-05-18: *"módulo de campanha/cupom integrado na plataforma... envolve variáveis para estratégia do cliente com os influ/afiliado"*). Os três clientes ativos com programa de afiliados já estão no catálogo (CASE-004 Shoulder, CASE-017 Baw, CASE-047 Ybera) — esta seção apenas adiciona referências cruzadas.

3. **Editoras (publishers)** — vertical nova. Editora Fórum aparece como cliente Wake em projeto de **checkout Headless + Pagbank** (Wake <> Benova, 2025-08-26).

---

### CASE-051 — Esplane (B2B da Você Constrói)

| Campo | Valor |
|---|---|
| **Cliente** | Esplane |
| **Setor** | material de construção / atacado B2B |
| **Tipo de operação** | B2B — braço atacado da Você Constrói |
| **Status** | ✓ ativo (citado em listas de cases B2B Wake) |
| **URL** | `esplane.com.br` ✓ confirmado em mensagens (Comercial Cross WAKE 2026-03-05; Wake <> Nairuz 2026-03-20) |
| **Plataforma anterior** | — |
| **ERP** | — |
| **Marketplace setup** | — |
| **Parceiros envolvidos** | — |
| **Data da última menção** | 2026-03-20 |

**Situação:** A Esplane é o **braço B2B da Você Constrói**, conforme descrito em duas mensagens canônicas listando os cases B2B Wake (Comercial Cross WAKE, 2025-12-17 e Wake <> Nairuz, 2026-03-20). A Você Constrói (B2C) entrou em produção pela Anymarket em 2025-08-22 (*"Você Constrói já liberado!"* — ANYMARKET & WAKE, 2025-08-22). Esplane também aparece de forma isolada como item em uma lista de URLs (Comercial Cross WAKE, 2026-03-05 18:47), confirmando que é uma loja Wake separada da loja-mãe B2C.

**Solução:** Wake Commerce em frente **B2B dedicada** (loja própria, não área logada da Você Constrói). Atacado para profissionais de construção / lojistas — pareado com a operação B2C (Você Constrói) também em Wake.

**Resultado:** Ativo. Sem dados de volumetria ou ticket no corpus.

**Insights / lessons:**
- Exemplo de **arquitetura B2B + B2C como duas lojas Wake separadas** (não área logada de uma única loja). Útil em pitch quando o cliente quer **isolamento total** entre frente B2C e atacado.
- Reforça padrão "marca-mãe B2C entra primeiro, B2B vem depois" — Você Constrói liberada em 08/2025, Esplane consolidada como case nas listas de 12/2025 e 03/2026.

---

### CASE-052 — CVC Corp (franquias)

| Campo | Valor |
|---|---|
| **Cliente** | CVC Corp |
| **Setor** | turismo / franquias |
| **Tipo de operação** | B2B — franquias comprando produtos da matriz |
| **Status** | ✓ ativo (case-âncora de B2B + benchmark interno) |
| **URL** | `plataformacvccorp.com.br` ✓ confirmado em mensagens (Comercial Cross WAKE 2026-03-05; Wake <> Nairuz 2026-03-20) |
| **Plataforma anterior** | — |
| **ERP** | — |
| **Marketplace setup** | — |
| **Parceiros envolvidos** | — |
| **Data da última menção** | 2026-03-20 |

**Situação:** CVC Corp opera plataforma B2B em Wake para suas **franquias comprarem produtos da matriz** (Comercial Cross WAKE, 2025-12-17; Wake <> Nairuz, 2026-03-20). Citado como **um dos 4 cases B2B canônicos da Wake** junto com GOLD, Destro e A Fábrica (Comercial Cross WAKE, 2026-03-05: *"Quais outros cases que temos de B2B, além de: CVC, GOLD, Destro e A Fabrica?"*). Citado também como **benchmark de estrutura B2B "bem completa"** ao lado da Buddemeyer (CoreBiz & Wake, 2026-01-30: *"Temos Buddemeyer e CVC como cases legais de benchmark"* — contexto: estrutura de crédito, regionalização, mix de produto e listas de preço diferentes).

**Solução:** Wake Commerce em modelo **matriz → rede de franquias** com lista de preço própria, mix de produto controlado e regras comerciais específicas por franquia.

**Resultado:** Case de referência B2B Wake.

**Insights / lessons:**
- **Único case B2B do corpus com modelo matriz↔franquia** documentado — usar como prova quando o prospect for rede franqueada querendo centralizar compras das unidades.
- A estrutura de **listas de preço por franquia + crédito + regionalização** é a combinação técnica que o time comercial cita como diferencial competitivo (mesmo combo em pitch Buddemeyer + CVC para CoreBiz em 2026-01-30).
- Não confundir com **CVC** em outros contextos do corpus (Corporate Venture Capital — predominante em grupos de startups). A CVC Corp aqui é a operadora de turismo.

---

### CASE-053 — Yora (atacado kids)

| Campo | Valor |
|---|---|
| **Cliente** | Yora |
| **Setor** | brinquedos / produtos kids — atacado |
| **Tipo de operação** | B2B — atacado de produtos kids |
| **Status** | ✓ ativo (citado em lista de cases B2B Wake) |
| **URL** | `yoraatacado.com.br` ✓ confirmado em mensagens (Wake <> Nairuz 2026-03-20) |
| **Plataforma anterior** | — |
| **ERP** | — |
| **Marketplace setup** | — |
| **Parceiros envolvidos** | — |
| **Data da última menção** | 2026-03-20 |

**Situação:** Yora aparece apenas nas duas listas canônicas de cases B2B Wake (Comercial Cross WAKE 2025-12-17; Wake <> Nairuz 2026-03-20) como *"atacado de produtos kids"*. Sem outras menções substantivas no corpus — case está vivo o suficiente para entrar nas listas internas, mas tem pouca atividade discutida.

**Solução:** Wake Commerce em frente **atacado B2B puro** para vertical infantil.

**Resultado:** ⚠ inferido — ativo mas com baixa visibilidade nas conversas operacionais.

**Insights / lessons:**
- Exemplo de **atacado B2B puro** (sem frente B2C correspondente no corpus) — modelo distinto do Esplane (Esplane é B2B sob marca B2C maior).
- Útil para pitch a fabricantes/distribuidores de **brinquedos** que querem digitalizar o canal atacado sem entrar em B2C imediatamente.
- Antes de usar publicamente, validar com account owner — única menção substantiva é em listas internas, sem conversas operacionais detalhadas.

---

### CASE-054 — Destro Macro (atacadista PR)

| Campo | Valor |
|---|---|
| **Cliente** | Destro (Destro Macro) |
| **Setor** | atacado / distribuição (PR) |
| **Tipo de operação** | B2B — atacadista do Paraná |
| **Status** | ✓ ativo (case-âncora B2B) |
| **URL** | `destromacro.com.br` ✓ confirmado em mensagens (Comercial Cross WAKE 2026-03-05; 2025-09-03; Wake <> Nairuz 2026-03-20) |
| **Plataforma anterior** | — |
| **ERP** | — |
| **Marketplace setup** | — |
| **Parceiros envolvidos** | Stone (envolvida em projeto com cliente Destro — fluxo de pagamento) |
| **Data da última menção** | 2026-03-20 |

**Situação:** Destro Macro é **atacadista do Paraná** e um dos **4 cases B2B canônicos** citados internamente (Comercial Cross WAKE, 2026-03-05: *"Quais outros cases que temos de B2B, além de: CVC, GOLD, Destro e A Fabrica?"*). URL também citada em listagem de clientes Wake do segmento alimentar/varejo regional (Comercial Cross WAKE, 2025-09-03 — junto com `mercadinhosaudavel.com.br`, `lojanovamix.com.br`, `saovito.com`, `ifrutus.com.br`, `lojazaeli.com.br`). Em 2025-07-14, projeto Wake & Stone identificou *"uma parte do fluxo que pode não estar corretamente mapeada"* envolvendo o cliente Destro — caso técnico aberto entre Wake e Stone.

**Solução:** Wake Commerce em frente **atacado regional**. Integração com Stone como parte do stack de pagamento.

**Resultado:** Case ativo de referência B2B. Caso técnico Wake↔Stone em 07/2025 foi sintoma de **fluxo de pedido cross-sistema ainda não totalmente mapeado** — alinhamento N1↔N2 necessário em projetos B2B com gateway externo.

**Insights / lessons:**
- Bom **case regional Sul** para pitch a atacadistas de outras regiões.
- Combinação **atacadista regional + Stone** é a única do corpus — projetos similares devem prever **alinhamento técnico antecipado** entre Wake e o adquirente.

---

### CASE-055 — Buddemeyer (B2B em onboarding)

| Campo | Valor |
|---|---|
| **Cliente** | Buddemeyer (cama, mesa e banho premium) |
| **Setor** | cama/mesa/banho |
| **Tipo de operação** | B2B (em onboarding) + B2C (já ativo) |
| **Status** | ⚠ B2B em onboarding (B2C contratado e ativo). Frente B2B "ainda em projeto" em 2026-03-05 |
| **URL** | ⚠ fornecido pelo usuário, não confirmado em mensagens |
| **Plataforma anterior** | — |
| **ERP** | — |
| **Marketplace setup** | Anymarket — suporte assistido (~R$ 1k–1,5k/mês — discutido como modelo conjunto com Casa Almeida em 2026-03-05) |
| **Parceiros envolvidos** | Anymarket (CS), Omnichat (em prospecção via abertura de portas para Kelly Mateus / Alvacir Schulze / César Grossl — checklist de 2026-05-21) |
| **Data da última menção** | 2026-05-21 |

**Situação:** Buddemeyer assinou Wake Commerce em **2025-06-30** (Mavericks: *"Buddemeyer tá dentro!! Commerce (48,1k mrr + 21k setup)"*). A frente **B2B ainda estava em projeto** em 2026-03-05 (*"Buddemeyer B2B ainda em projeto"* — Comercial Cross WAKE) e citada em lista canônica como *"Buddemeyer em onboarding"* no recorte B2B (Comercial Cross WAKE 2025-12-17; Wake <> Nairuz 2026-03-20). Em 2026-03-05, time de CS (Math Almeida) discutia se a Anymarket podia oferecer suporte assistido (~R$ 1k/1,5k) à Buddemeyer — modelo discutido em paralelo ao da Casa Almeida (CASE-023). Citado como **benchmark de estrutura B2B "bem completa"** ao lado da CVC (CoreBiz & Wake, 2026-01-30). Em 2026-05-21, Omnichat planejava abordagem para abrir portas com gerentes/coordenadores da Buddemeyer.

**Solução:** Wake Commerce (B2C) já contratado + frente **B2B em construção** (lista de preço, regionalização, mix). Anymarket para marketplaces. Suporte assistido Anymarket em discussão.

**Resultado:** Frente B2C operacional; frente B2B em onboarding (status "em projeto" em 03/2026, sem evidência de go-live B2B publicado no corpus até 06/2026).

**Insights / lessons:**
- Junto com CVC, é o **case-benchmark interno** que o comercial usa para defender capacidade B2B da Wake em pitches enterprise (citado em pitch CoreBiz 2026-01-30).
- Ticket cheio R$ 48,1k MRR + R$ 21k setup — operação Enterprise (presença na lista de "Enterprise - Não entra" de 2025-11-03 confirma o segmento, mas refere-se a uma rodada específica de canais que não recebeu essa conta).
- B2B com Wake é projeto de prazo: contrato em 06/2025, B2B ainda em onboarding 9 meses depois (03/2026) — calibrar expectativas com prospects.

---

### CASE-056 — Gold Food Service

| Campo | Valor |
|---|---|
| **Cliente** | Gold (Gold Food Service) |
| **Setor** | food service / distribuição B2B |
| **Tipo de operação** | B2B — distribuição food service |
| **Status** | ✓ ativo (case-âncora B2B) |
| **URL** | `goldfoodservice.com.br` ✓ confirmado em mensagens (Comercial Cross WAKE 2026-03-05; Wake <> Nairuz 2026-03-20) |
| **Plataforma anterior** | — |
| **ERP** | — |
| **Marketplace setup** | — |
| **Parceiros envolvidos** | — |
| **Data da última menção** | 2026-03-20 |

**Situação:** Gold Food Service é **um dos 4 cases B2B canônicos da Wake** (Comercial Cross WAKE, 2026-03-05: *"Quais outros cases que temos de B2B, além de: CVC, GOLD, Destro e A Fabrica?"*). Aparece apenas nas listas canônicas — sem conversas operacionais detalhadas no corpus.

**Solução:** Wake Commerce em frente **B2B food service**.

**Resultado:** ⚠ inferido ativo — sem volumetria ou ticket no corpus.

**Insights / lessons:**
- **Único case B2B do corpus na vertical food service** — usar em pitch para distribuidores/atacadistas de alimentos/bebidas.
- Combinado com Soubeef (CASE-036) cobre dois ângulos diferentes em food: distribuição B2B (Gold) vs. produto cárneo de consumo (Soubeef).

---

### CASE-057 — A Fábrica

| Campo | Valor |
|---|---|
| **Cliente** | A Fábrica |
| **Setor** | ⚠ status desconhecido (URLs sugerem indústria/fabricante de bens duráveis, não confirmado em mensagem) |
| **Tipo de operação** | B2B (case-âncora) — possível indústria → revenda |
| **Status** | ✓ ativo (case-âncora B2B) |
| **URL** | `afabrica.com` + `checkout.afabrica.com` ✓ confirmado em mensagens (Comercial Cross WAKE 2026-03-05 18:50; Wake <> Nairuz 2026-03-20) |
| **Plataforma anterior** | — |
| **ERP** | — |
| **Marketplace setup** | — |
| **Parceiros envolvidos** | — |
| **Data da última menção** | 2026-03-20 |

**Situação:** A Fábrica é **um dos 4 cases B2B canônicos da Wake** (Comercial Cross WAKE, 2026-03-05: *"Quais outros cases que temos de B2B, além de: CVC, GOLD, Destro e A Fabrica?"*). A presença de **subdomínio dedicado de checkout** (`checkout.afabrica.com`) é confirmação técnica de arquitetura B2B com checkout customizado.

**Solução:** Wake Commerce em frente B2B com **checkout em subdomínio dedicado** — padrão de arquitetura que indica fluxo de fechamento B2B distinto do B2C (lista de preço, condições, autenticação).

**Resultado:** Ativo. Sem outras conversas operacionais substantivas no corpus.

**Insights / lessons:**
- **Único case B2B do corpus que documenta uso explícito de `checkout.<dominio>`** — usar como referência técnica quando prospect perguntar sobre arquitetura de checkout B2B customizada.
- Confirma que a Wake suporta **subdomínio de checkout próprio** — relevante em pitches onde o cliente quer separar fluxo de fechamento B2B do site institucional/catálogo.

---

### CASE-058 — Unitoys

| Campo | Valor |
|---|---|
| **Cliente** | Unitoys |
| **Setor** | brinquedos |
| **Tipo de operação** | B2B — citado na lista canônica B2B; recorte interno também o trata como referência em brinquedos |
| **Status** | ✓ ativo (assinado 2026-01-21) |
| **URL** | `unitoys.com.br` ✓ confirmado em mensagens (Wake <> Nairuz 2026-03-20) |
| **Plataforma anterior** | — |
| **ERP** | — |
| **Marketplace setup** | — |
| **Parceiros envolvidos** | Série//A (visitou cliente em 2026-03-02) |
| **Data da última menção** | 2026-03-20 |

**Situação:** Unitoys assinou em **2026-01-21** — *"Unitoys assinado ✍🏻 Commerce (5k mrr + 24k setup)"* (Comercial Cross WAKE). Em 2026-03-02, Série//A visitou o cliente. Em 2026-03-17, citada como referência ao perguntarem *"além da unitoys, quem mais de brinquedos?"* (Comercial Cross WAKE) — confirma que é o caso de referência interno na vertical brinquedos. Aparece na lista B2B canônica de 2026-03-20 (Wake <> Nairuz).

**Solução:** Wake Commerce. Ticket SMB (**R$ 5k MRR + R$ 24k setup**).

**Resultado:** Ativo. Assinatura recente, agência Série//A no relacionamento.

**Insights / lessons:**
- **Único case de referência em brinquedos** no corpus interno — usar em pitches a fabricantes/distribuidores do segmento.
- Ticket R$ 5k MRR sinaliza SMB — combinar com Yora (CASE-053) para cobrir os dois recortes do vertical (brinquedos varejo Wake + atacado kids).
- A presença na lista B2B junto com Yora sugere que o recorte interno "B2B" inclui também marcas com forte canal atacado, não apenas operações puramente atacadistas.

---

### CASE-059 — Editora Fórum (publishers)

| Campo | Valor |
|---|---|
| **Cliente** | Editora Fórum |
| **Setor** | editorial / publishers / conteúdo jurídico-empresarial |
| **Tipo de operação** | B2B + B2C — venda de publicações/assinaturas |
| **Status** | ✓ ativo (em projeto técnico em 2025-08; status final 06/2026 não confirmado no corpus) |
| **URL** | `negocios.forumconhecimento.com.br` ⚠ fornecido pelo usuário, não confirmado em mensagens |
| **Plataforma anterior** | — |
| **ERP** | — |
| **Pagamento** | PagBank (no novo checkout Headless) |
| **Parceiros envolvidos** | Benova (agência de implementação) |
| **Data da última menção** | 2025-08-26 |

**Situação:** Editora Fórum é cliente Wake em projeto de **Pagbank no novo checkout Headless** conduzido pela agência **Benova** (Wake <> Benova, 2025-08-26): *"Integração Pagbank no novo checkout Headless está 100%? (Cliente Editora Fórum)"* — resposta: *"Ainda não está 100% mas já estão trabalhando nos ajustes. Tem uma entrega (Pix) prevista para início de Setembro e outra (CC) mais para o fim do mês"*. Em 2025-08-26, conversa explícita confirma a relação Editora Fórum ↔ Pagbank no projeto.

**Solução:** Wake Commerce com **checkout Headless** + integração **PagBank** (Pix + Cartão de Crédito). Implementação conduzida pela Benova.

**Resultado:** ⚠ Em 2025-08-26, integração Pagbank tinha entregas previstas para 09/2025 (Pix) e fim de 09/2025 (CC). Status atual após essas datas não consta no corpus — validar com account owner.

**Insights / lessons:**
- **Única editora/publisher** identificada no corpus como cliente Wake — vertical greenfield para a base de cases.
- Combinação **checkout Headless + PagBank** é arquitetura relevante para pitch a publishers que vendem assinaturas / produtos digitais com necessidade de fluxo de pagamento próprio.
- Caso técnico em 2025-08 abriu thread sobre maturidade da integração Pagbank no checkout Headless — útil em discussões de homologação de novos gateways.

---

## 7. Lote 2 — Casos descobertos no sweep 06/2026 (CASE-060 a CASE-072)

Este bloco consolida os casos identificados no gap-sweep de junho/2026 (`docs/wake-customer-cases-gap-list.md`). Cobertura: clientes ativos, leads em rotação multi-agência, migrações em andamento, churn e go-lives recentes que não estavam no recorte original de 59 casos.

---

### CASE-060 — Gregory (moda + Metakosmos provador virtual)

| Campo | Valor |
|---|---|
| **Cliente** | Gregory |
| **Setor** | moda feminina |
| **Tipo de operação** | B2C — varejo de moda |
| **Status** | ✓ ativo (em produção, com case público de provador virtual Metakosmos) |
| **URL** | `gregory.com.br` ✓ confirmado em mensagens (Comercial Cross WAKE 2026-05-18; Wake <> Nairuz 2026-03-20) |
| **Plataforma anterior** | — |
| **ERP** | — |
| **Marketplace setup** | Anymarket (painel + acessos sendo configurados em 03/2026) |
| **Parceiros envolvidos** | Metakosmos (provador virtual em produção), Floui (sustentação técnica — chamado de subida de produtos com SLA contratual de 15 dias), Anymarket |
| **Tags** | `audience-sales`, `integração` |
| **Data da última menção** | 2026-05-28 |

**Situação:** Gregory aparece como **case de provador virtual Metakosmos em produção** (Wake & Metakosmos, 2026-05-14: *"Gregory (provador virtual) — gregory.com.br/produto/blazer-xadrez-terracota-165222"*). Em 2026-03-11, o feature já era citado como referência para outros prospects: *"Também daria pra colocar o provador igual da Gregory, só não sei como funcionam os custos dos testes"* (Wake & Metakosmos). A conta tem **agenda direta com o CEO** sendo trabalhada — em 2026-03-11 a CS Anymarket reportou *"Daqui a pouquinho terei uma agenda com o CEO da Gregory e a Anymarket é um dos temas que está na pauta"*. Sustentação técnica está com a **Floui** (chamado de subida de produtos discutido em 2026-05-26 — Floui defendeu SLA contratual de 15 dias). Houve report de 2026-05-22 (Tropa de Elite 3): *"Gregory ta uma merda não tem nada configurado"* — sinal de fricção operacional que precisava de retomada com Cintia (contato interno) — seguido de feedback positivo em 2026-05-20 (Wake+Omni360+Affys): *"o Gregory deu um feedback positivo"*.

**Solução:** Wake Commerce + **Metakosmos provador virtual (mK Fashion+)** em produção + Anymarket para marketplaces + Floui como agência de sustentação. Contatos comerciais capturados em Tropa de Elite 3, 2026-04-15: Gregory Duca Giovanelli (Gerente Operações), Luiza Duca Giovanelli Vergara (Gerente Administrativa), Andrea Duca (Diretora) — confidenciais.

**Resultado:** Provador virtual operacional e usado como case Metakosmos. Revenue share Metakosmos por Gregory **nunca foi cobrada** até 02/2026 (WAKE | Coordenação, 2026-02-20: *"Metakosmos por exemplo nunca cobramos por Gregory. (Foco para Março)"*) — sinal de gap operacional no fluxo de cobrança de revenue share de parceiros tech.

**Insights / lessons:**
- **Único case Metakosmos em produção** explicitamente citado como referência no corpus — usar em pitch a marcas de moda que perguntam sobre provador virtual / IA visual.
- Wake **deixou de cobrar revenue share Metakosmos por Gregory** durante todo o ciclo — sinal para Coordenação Parcerias revisar fluxo de billing de revenue share de parceiros tech (não só agência).
- Sustentação **Floui defendendo SLA contratual de 15 dias** para subida de produtos — referência útil de baseline de SLA agência em projetos Wake.

---

### CASE-061 — Ave Rara / Averara (moda, Go-Live 05/2026)

| Campo | Valor |
|---|---|
| **Cliente** | Ave Rara (razão social: AVERARA CONFECCOES EIRELI) |
| **Setor** | moda |
| **Tipo de operação** | B2C — varejo de moda |
| **Status** | ✓ ativo — **Go-Live 2026-05-18** |
| **URL** | `averarabrand.com.br` ✓ confirmado (The Rise of Skywalker 2026-05-18: *"Loja no ar!!! https://www.averarabrand.com.br/"*); ambiente de teste `www3.averarabrand.com.br` |
| **Plataforma anterior** | — |
| **ERP** | — |
| **Marketplace setup** | — |
| **Parceiros envolvidos** | Simples Inovação (primeiro case da agência — *"Uhuuul, primero case da Simples conosco 🚀"*, 2026-05-19), Stone (TEF + Prateleira Infinita / POS) |
| **Tags** | `audience-sales`, `proposta` |
| **Data da última menção** | 2026-06-03 |

**Situação:** Ave Rara assinou Wake Commerce em **fevereiro/2026** (The Rise of Skywalker, 2026-02-12: *"Fechamos Averara"*; 2026-02-26: *"AveRara: 55k MRR + 68.9 Setup"* — forecast de Mavericks/TRoS). Contrato teve atraso de subida para faturamento — 2026-03-02: *"o contrato de averara não subiu ainda pra eles lá pra faturamento. Ta assinado isso já?"*. **Go-Live em 2026-05-18** com ação Family & Friends ativada (cupom interno divulgado em WAKE | Time Parcerias 2026-05-15). Primeiro case operado pela agência **Simples Inovação**. Em 06/2026, **negociação ativa de Prateleira Infinita + integração POS Stone (Connect)** em curso ([work] Wake & Stone, 2026-06-02/03): cliente já opera TEF com Stone e foi orientado a contratar Prateleira Infinita, mas há fricção porque *"nem todas as maquininhas que ele já possui são compatíveis com esse produto"*. Pergunta aberta em Comercial Cross WAKE 2026-06-03: *"quem negociou o prateleira infinita para o cliente AVERARA?"*.

**Solução:** Wake Commerce (R$ 55k MRR + R$ 68,9k setup, conforme forecast 2026-02-26) + Simples Inovação como agência + Stone TEF + projeto Prateleira Infinita em negociação.

**Resultado:** Go-Live operacional desde 2026-05-18. Pendência aberta: alinhamento de compatibilidade de maquininhas Stone para a operação Prateleira Infinita.

**Insights / lessons:**
- **Primeiro case da agência Simples Inovação** — marco para o programa de canais.
- Caso operacional para o playbook de **Prateleira Infinita + POS**: a venda da Prateleira Infinita assumiu que o parque de maquininhas atual do cliente seria compatível, o que não foi verdade. **Discovery técnica antes do contrato** deve validar compatibilidade do modelo de POS antes de fechar.
- Forecast de TRoS captura ticket Commerce relevante (R$ 55k MRR) — referência interna de pricing moda em 2026.

---

### CASE-062 — OffPremium (outlet — Anymarket + Marketplace Blueman)

| Campo | Valor |
|---|---|
| **Cliente** | OffPremium |
| **Setor** | varejo / outlet de moda |
| **Tipo de operação** | B2C — outlet |
| **Status** | ✓ ativo (em produção, com integração Marketplace Blueman em desenvolvimento) |
| **URL** | `offpremium.com.br` ✓ confirmado (Wake + Marketplace Blueman 2025-10-29, link de produto teste) |
| **Plataforma anterior** | — |
| **ERP** | — |
| **Marketplace setup** | Anymarket + **Marketplace Blueman** (em desenvolvimento — perspectiva otimista 03/2026, realista 04-05/2026, pessimista 06/2026) |
| **Parceiros envolvidos** | Anymarket, Marketplace Blueman |
| **Tags** | `audience-cs`, `integração` |
| **Data da última menção** | 2026-01-28 |

**Situação:** OffPremium é cliente Wake em produção e foi piloto da integração **Marketplace Blueman** via Anymarket. Em 2025-10-29, pedidos teste foram processados via `offpremium.com.br/raquete-slim-basica---blueman---produto-teste-403271-1606992/p`. Em 2025-11-03, a integração tinha gap conhecido: *"esse aqui é OffPremium, mas o ideal era fazer um novo depois que subi uma melhoria, pq lá não tem repasse certo? Ai o pagamento é direto para vcs, e o número título não é o marketPlacePedidoId"* (Wake + Marketplace Blueman). Em 2026-01-28, prazos foram realinhados com a Anymarket: *"Perspectiva otimista: Março / Perspectiva realista: entre Abril e Maio / Perspectiva pessimista: Junho"* — *"esse é um desenvolvimento complexo, integração de dados do pagamento, requer um estudo e desenvolvimento maior"*. Cliente listado junto com **Babado Top** como par pendente de finalização da integração Mercado Livre via Marketplace Blueman.

**Solução:** Wake Commerce + Anymarket + Marketplace Blueman (canal Mercado Livre via hub Blueman). Integração de pagamento + número de título do pedido sendo customizados para fluxo Blueman.

**Resultado:** Em 06/2026 ainda em desenvolvimento (status mais recente: 2026-01-28, sem update posterior no corpus). Validar com account owner se Go-Live final aconteceu na janela pessimista (06/2026).

**Insights / lessons:**
- Caso operacional do **piloto Marketplace Blueman + Anymarket + Wake** — útil em pitch a outlets que querem operar via hub não-padrão.
- **Integração de pagamento marketplace tem gap conhecido** quando o repasse vai direto ao lojista e o `marketPlacePedidoId` não é o número de título. Sinalizar em discovery de qualquer marca outlet que use esse modelo.
- Junto com **Babado Top** forma o par de pilotos Blueman do corpus.

---

### CASE-063 — Eletro Mais (MT) (lead Duo — saindo de OpenCart)

| Campo | Valor |
|---|---|
| **Cliente** | Eletro Mais |
| **Setor** | varejo eletro (MT) |
| **Tipo de operação** | B2C — varejo eletro regional |
| **Status** | ⚠ lead em rotação multi-agência (último contato 2026-01-12, sem resposta) |
| **URL** | `eletromaismt.com.br` ✓ confirmado (Duo e Wake Parcerias 2025-12-10) |
| **Plataforma anterior** | OpenCart (6 meses no ar, sem venda) |
| **ERP** | Santri Sistemas |
| **Marketplace setup** | — (não vende em marketplaces; só loja física) |
| **Parceiros envolvidos** | Duo (origem do lead), DRIVEN (rotação) — contato Cristiano |
| **Tags** | `audience-sales`, `lead-pipeline` |
| **Data da última menção** | 2026-01-12 |

**Situação:** Eletro Mais entrou como **Oportunidade Canais** (Duo e Wake Parcerias, 2025-11-26 / 2025-12-10 — mensagem padronizada repetida em DRIVEN & Wake e Wake & A&EIGHT B8One): *"Plataforma atual: Open Cart / 1k visitas / mês / Estão com esse e-commerce no ar à 6 meses, porém, sem venda / Canais de venda: Loja física (não vendem nos marketplaces) / ERP: Santri sistemas / Urgência: O quanto antes"*. Em 2026-01-12, DRIVEN reportou: *"Eletromaismt - Cristiano - Marcado para quarta-feira as 14hs"* — agenda marcada mas sem retorno posterior no corpus.

**Solução:** Proposta Wake Commerce para substituir OpenCart, com integração Santri Sistemas (ERP regional — verificar com Solutions se precisa integrador-erp dedicado, já que Santri **não está na taxonomia canônica** de ERPs Wake).

**Resultado:** ⚠ Status final desconhecido após 2026-01-12. Sinal de **lead com urgência alta mas sem fechamento** — vale follow-up.

**Insights / lessons:**
- Caso clássico de **lojista com e-commerce parado** (6 meses no ar sem venda) — Wake compete com a percepção de fracasso da plataforma anterior, não com a plataforma em si.
- **Santri Sistemas** é ERP regional fora da taxonomia canônica (`docs/wake-partnership-kb-playbook.md` §2) — precisa ser tratado como ERP customizado / via integrador genérico (Jitterbit, Leanwork, Iota.Hub).
- Acompanha o lead **Cimafer** (CASE-064) no mesmo formato Oportunidade Canais Duo de 12/2025.

---

### CASE-064 — Cimafer (construção/ferramentas — lead Duo, GMV R$ 1M/mês claim)

| Campo | Valor |
|---|---|
| **Cliente** | Cimafer |
| **Setor** | construção / ferramentas |
| **Tipo de operação** | B2C + loja física (omnichannel) |
| **Status** | ⚠ lead em rotação multi-agência (último contato 2025-12-10) |
| **URL** | `cimafer.com.br` ✓ confirmado (Duo e Wake Parcerias 2025-12-03/10) |
| **Plataforma anterior** | **Shopify** (claim do cliente: *"não sente segurança nas vendas online"*) |
| **ERP** | Boreo |
| **Marketplace setup** | — |
| **Parceiros envolvidos** | Duo (origem do lead), DRIVEN, Wake & A&EIGHT (B8One) — contato Geórgia Pires |
| **Tags** | `audience-sales`, `lead-pipeline` |
| **Data da última menção** | 2026-01-12 |

**Situação:** Cimafer entrou como **Oportunidade Canais** (Duo e Wake Parcerias, 2025-12-03 / 2025-12-10): *"https://www.cimafer.com.br/ / 13K de visualização / GMV: R$ 1.000.000,00 mensal (esse valor seria total, E-commerce + loja física) / ERP: Boreo / Cliente Shopify, ela gostaria de trocar de plataforma porque não sente segurança nas vendas online"*. Em 2025-12-05, A&EIGHT (B8One) confirmou tentativa de chamar Cimafer. Em 2026-01-12, DRIVEN reportou: *"Cimager [sic] - Gerogia Pires - Não atende o celular e nao responde"* — lead esfriou.

**Solução:** Proposta Wake Commerce em substituição ao Shopify, com integração ERP Boreo (fora da taxonomia canônica — provavelmente requer integrador-erp).

**Resultado:** ⚠ Lead esfriou em 01/2026 (não atende celular, não responde). Não há sinal de fechamento.

**Insights / lessons:**
- **NB**: o gap-list inicial (`docs/wake-customer-cases-gap-list.md`) classificou Cimafer como migração saindo de **OpenCart** — a evidência em mensagem mostra que a plataforma atual era **Shopify** (Duo e Wake Parcerias, 2025-12-03). Confiar na mensagem.
- Claim **GMV R$ 1M/mês** combina e-commerce + loja física — útil como referência de TAM em vertical construção, mas validar o split online/offline antes de citar em pitch.
- **Boreo** é ERP regional fora da taxonomia canônica — caso para Solutions validar integrador genérico antes de prometer integração.
- Padrão: leads multi-agência em rotação que **esfriam por falta de resposta do cliente** são comuns em ticket médio/alto regional — sugere processo de **qualificação dupla** (Wake + agência) antes de distribuir.

---

### CASE-065 — Brofitwear (moda fitness — primeiro Go-Live Quick Digital, saída VTEX)

| Campo | Valor |
|---|---|
| **Cliente** | Brofitwear |
| **Setor** | moda fitness |
| **Tipo de operação** | B2C — varejo de moda fitness |
| **Status** | ✓ ativo — **Go-Live 2026-05-19** (primeiro Go-Live da Quick Digital) |
| **URL** | `brofitwear.com.br` ✓ confirmado (Quick Digital & Wake Parcerias 2026-05-19; Wake <> Catus 2025-11-11) |
| **Plataforma anterior** | **VTEX** (saída em 2025-11) |
| **ERP** | — |
| **Marketplace setup** | — |
| **Parceiros envolvidos** | Catus (lead inicial em 11/2025 — reativado), **Quick Digital** (agência que fez o Go-Live) |
| **Tags** | `audience-sales`, `competitivo` |
| **Data da última menção** | 2026-05-19 |

**Situação:** Brofitwear entrou em 2025-11-11 como reativação de lead via Catus (*"Pode reativar a Brofitwear.com.br como lead por favor? A gente esta tirando eles da Vtex"* — Wake <> Catus). Indicação interna pela Vanessa (WAKE | Time Agência, 2025-11-11). **Assinou Wake Commerce em 2025-10-30** (Mavericks + Comercial Cross WAKE: *"Brofitwear assinado ✍🏻 Commerce (6,1k mrr + 12k setup)"*). **Go-Live em 2026-05-19** anunciado em Quick Digital & Wake Parcerias: *"https://www.brofitwear.com.br/ Primeiro Go Live da Quick 🚀"*.

**Solução:** Wake Commerce (R$ 6,1k MRR + R$ 12k setup) + Quick Digital como agência de implementação.

**Resultado:** Go-Live operacional desde 2026-05-19. **Caso-âncora do programa da Quick Digital** — primeiro projeto entregue pela agência.

**Insights / lessons:**
- **Migração de saída VTEX → Wake** confirmada em vertical moda fitness — útil em battlecard VTEX.
- **Caso de marco para a Quick Digital** — usar como prova social ao discutir capacidade de novas agências entregarem Go-Lives em prazo (assinatura 10/2025 → Go-Live 05/2026, ciclo de ~7 meses).
- Ticket Commerce R$ 6,1k MRR + R$ 12k setup é referência de SMB moda fitness 2025-Q4.

---

### CASE-066 — Casafer (construção/varejo — SAP B1 + WooCommerce migration via Eficaz)

| Campo | Valor |
|---|---|
| **Cliente** | Casafer |
| **Setor** | construção / varejo de materiais |
| **Tipo de operação** | B2C + B2B (telefone, WhatsApp, balcão, representantes, LinkedIn) |
| **Status** | ⚠ lead em qualificação (último contato 2025-06-30) |
| **URL** | `casafer.com.br` ✓ confirmado (Eficaz <> WAKE, 2025-06-30) |
| **Plataforma anterior** | **WordPress / WooCommerce** (catálogo, sem vendas efetivas) |
| **ERP** | **SAP Business One** |
| **Marketplace setup** | — |
| **Parceiros envolvidos** | Eficaz (origem do lead), O2 Consultoria (rotação) |
| **Tags** | `audience-solutions`, `lead-pipeline` |
| **Data da última menção** | 2025-06-30 |

**Situação:** Lead vindo via SDR em 2025-06-30, distribuído para **Eficaz** e **O2 Consultoria** com o mesmo briefing (Eficaz <> WAKE / O2 Consultoria & Wake): *"Casafer - Commerce - Canais / https://casafer.com.br / Plataforma atual: WordPress / WooCommerce / Canais de vendas: Telefone, WhatsApp, balcão loja, representantes, LinkedIn / ERP: SAP Business One / Comentou que hoje o site está como um catálogo. Efetivamente não tem feito vendas por lá. O que tem incomodado é o desempenho e as dificuldades de integração com o ERP. Faturamento online, está praticamente zerado. Faturamento offline, aprox. de [valor confidencial]"*. Contato: Allan Guimarães (allan@casafer.com.br) — confidencial.

**Solução:** Proposta Wake Commerce + integrador-erp para SAP Business One (Intergate ou LinkToU, conforme taxonomia `docs/wake-partnership-kb-playbook.md` §3.1).

**Resultado:** ⚠ Sem atualização após 2025-06-30. Status final desconhecido.

**Insights / lessons:**
- Padrão **SAP B1 + WooCommerce com integração ERP falhando** é replay do mesmo problema visto em cases enterprise — útil em pitch para outros lojistas SAP B1 frustrados com plataforma open-source.
- Discussão sobre **e-commerce como catálogo (não como canal de venda)** é argumento de venda forte para Wake Commerce — vale documentar em battlecard "WooCommerce/SAP".
- **Distribuição duplicada para Eficaz e O2** no mesmo dia — caso para discutir governança de carrossel quando o briefing é o mesmo (`audience-pam`).

---

### CASE-067 — Singular Medicamentos (farma — rotação multi-agência)

| Campo | Valor |
|---|---|
| **Cliente** | Singular Medicamentos |
| **Setor** | farmácia (varejo farma) |
| **Tipo de operação** | B2C — varejo farma |
| **Status** | ✓ ativo (em produção, com painel Wake operacional desde 11/2025) |
| **URL** | `singularmedicamentos.com.br` ✓ confirmado (e-mails em ANYMARKET & WAKE, 2025-11-06) |
| **Plataforma anterior** | — |
| **ERP** | — |
| **Marketplace setup** | Anymarket (painel ativo) |
| **Parceiros envolvidos** | Nairuz, N1, Cava Digital (rotação multi-agência mencionada no gap-list), LeanWork (App em discovery em 10/2025), BonifiQ (interesse mútuo capturado em 10/2025) |
| **Tags** | `audience-cs`, `suporte-escalação` |
| **Data da última menção** | 2026-06-02 |

**Situação:** Singular Medicamentos é cliente Wake **em produção** (painel "SINGULAR MEDICAMENTOS - WAKE", OI 259064779 — ANYMARKET & WAKE, 2025-11-06: pedido de reenvio de login/senha para e-mails @singularmedicamentos.com.br). Em 09/2025 estava sendo qualificado para a corrida de premiação Vindi (Comercial Cross WAKE 2025-09-25: *"Eu estou vendo a Singular medicamentos e não é para setembro"*). Em 10/2025 a LeanWork pediu update de bloqueios em conjunto com **DLK Modas, Luiza Barcelos, YBERA** — sinal de que Singular faz parte do mesmo conjunto de contas indicadas para o **App LeanWork** (Wake & Leanwork, 2025-10-06). Em 2025-10-02, **BonifiQ** reportou interesse mútuo: *"a Singular Medicamentos demonstrou bastante interesse em conhecer vocês"* (Wake & BonifiQ). Em 2026-04-28 foi citada como cliente Caaqui e em 2026-05-05 (Comercial Cross WAKE) pediu-se update de tempo de implantação + agência responsável. Em **2026-06-02** (Duo e Wake Parcerias) abriu thread técnica: *"precisa fazer uma leitura de receita de oculos também, a Singular Medicamentos pede receita/dados do médico, mas é algo mais manual hoje. Temos alguma outra alternativa na plataforma?"* — sinal de **requisito regulatório farma** (validação de receita médica para venda online) que ainda é manual.

**Solução:** Wake Commerce + Anymarket. Rotação multi-agência (Nairuz / N1 / Cava Digital / Caaqui) — confirmar account owner atual com Coordenação Parcerias.

**Resultado:** Ativo. Discovery aberto para **automação de validação de receita médica** (regulatório farma) em 06/2026.

**Insights / lessons:**
- **Único cliente farma em produção** capturado no recorte com requisito explícito de **validação de receita** — sinal para Solutions/Produto de que há gap funcional para a vertical farma (receita médica + dados do médico ainda manual em 06/2026).
- Singular convidada para o **evento farma 09/2025** (Wake/Econverse Farma - Vurdere/Leal CRM Bonus 2025-09-23) com **4 convidados** — alto engajamento na vertical farmacêutica relativa ao recorte.
- Sinal de **rotação multi-agência sem account owner único claro** (Nairuz/N1/Cava/Caaqui/Econverse) — risco de conflito de canal; recomendar consolidação com PAM.

---

### CASE-068 — OG Brindes (B2B brindes — busca por agência de sustentação)

| Campo | Valor |
|---|---|
| **Cliente** | OG Brindes |
| **Setor** | brindes corporativos / B2B |
| **Tipo de operação** | B2B — distribuição de brindes corporativos |
| **Status** | ⚠ ativo, em busca de agência de sustentação (último update 07/2025) |
| **URL** | `ogbrindes.com.br` ✓ confirmado (Eficaz <> WAKE, 2025-07-08) |
| **Plataforma anterior** | — |
| **ERP** | — |
| **Marketplace setup** | — |
| **Parceiros envolvidos** | FG, Série//A, Econverse, Eficaz (indicações de sustentação), Dotkom (CS Dotkom pedindo agenda Wake + cliente em 07/2025) |
| **Tags** | `audience-pam`, `ecossistema-parceiros` |
| **Data da última menção** | 2025-07-24 |

**Situação:** OG Brindes pediu **indicação de agências para sustentação** em 2025-06-24 (FG + Wake / Wake <> Série//A / Wake <> Econverse — mensagem replicada): *"Cliente OG Brindes pediu indicação de agências para sustentação"*. Contato comercial: José Carlos Grzybowski, Diretor Geral (jose.carlos@ogbrindes.com.br) — confidencial. Em 2025-07-17, **Dotkom** pediu agenda de call entre Wake + cliente para falar de OG Brindes (Parcerias Dotkom & Wake), agenda marcada para 28/07.

**Solução:** Wake Commerce com agência de sustentação a ser definida (rotação FG / Série//A / Econverse / Eficaz / Dotkom).

**Resultado:** ⚠ Sem update após 2025-07-24. Status atual desconhecido.

**Insights / lessons:**
- Padrão **cliente Wake pedindo indicação de agência de sustentação** sinaliza que estava operando sem agência fixa — útil para discussões de **governance de account ownership pós Go-Live**.
- **Múltiplas agências chamadas em paralelo** para o mesmo cliente sem rotação formal — caso para Coordenação Parcerias revisar fluxo de indicação de sustentação.
- Vertical **B2B brindes corporativos** ainda subdocumentada no catálogo — OG Brindes é o primeiro caso explícito do recorte.

---

### CASE-069 — Biggie / abiggie (lead desqualificado — Shopify→Wake recusado)

| Campo | Valor |
|---|---|
| **Cliente** | Biggie (abiggie) |
| **Setor** | ? (não identificado no corpus — Shopify SMB) |
| **Tipo de operação** | B2C |
| **Status** | ❌ **lead perdido / desqualificado** (cliente não quer migrar, só quer frontend em Shopify) |
| **URL** | `abiggie.com.br` ✓ confirmado (Wake <> Ge Commerce 2025-10-09) |
| **Plataforma anterior** | **Shopify** (continua) |
| **ERP** | Tiny |
| **Marketplace setup** | — |
| **Parceiros envolvidos** | Ge Commerce / O2 Consultoria (rotação inicial), **Catus** (agência prejudicada na compensação por lead desqualificado) |
| **Tags** | `audience-pam`, `lead-pipeline` |
| **Data da última menção** | 2025-10-28 |

**Situação:** Biggie entrou como **Oportunidade Canais** em 2025-10-09 (Wake <> Ge Commerce / O2 Consultoria & Wake): *"Cliente Shopify / ERP - Tiny / Faturamento: Ainda não esta faturando porque esta no processo de criação do site / A ideia é revisar o layout, estrutura das páginas e também o SEO"*. Em 2025-10-28 (WAKE | Time Agência), a Catus foi citada como prejudicada: *"precisamos compensá-los em dois leads, que direcionamos e são desqualificados. Biggie - (Perdido/Desqualificado) Cliente não quer migrar, só quer alguém para fazer frontend em shopify. Trocamotor - (Perdido/Desqualificado) Cliente queria apenas desenvolver o front na solução atual que ele tem hoje"*.

**Solução:** Não aplicável — lead recusou migração.

**Resultado:** ❌ **Lead desqualificado**. Agência **Catus** sinalizou para compensação por lead falso (junto com Trocamotor — mesmo padrão).

**Insights / lessons:**
- Caso de **discovery falho na qualificação de canais**: o lead foi distribuído como Oportunidade Canais sem confirmar a intenção de migração da plataforma. **ERP Tiny + Shopify nativo é integração funcional** — não há gancho funcional para forçar troca para Wake.
- Padrão "**cliente quer só refazer frontend**" em Shopify é sinal claro de **desqualificação antes de distribuir** — caso para playbook de SDR/PAM.
- **Compensação de agência por lead falso** (Catus, Biggie + Trocamotor) é regra precedente — referência para discussão de SLA de qualidade de lead com canal.

---

### CASE-070 — Unikapro / Unabeleza (churn — migrou para Nuvemshop)

| Campo | Valor |
|---|---|
| **Cliente** | Unikapro (renomeada/refundada como **Unabeleza**) |
| **Setor** | beleza |
| **Tipo de operação** | B2C — varejo de beleza |
| **Status** | ❌ **churn** — migrou para Nuvemshop em 03/2026 |
| **URL** | `unabeleza.com.br` (nova marca, fora da Wake) |
| **Plataforma anterior** | Wake (saiu) |
| **ERP** | — |
| **Marketplace setup** | — |
| **Parceiros envolvidos** | Econverse (citou ausência da Unabeleza em planilha de comissionamento em 01/2026) |
| **Tags** | `audience-sales`, `competitivo` |
| **Data da última menção** | 2026-03-27 |

**Situação:** Em 2026-03-27 (WAKE | Time Parcerias), confirmação direta: *"Unikapro, to vendo que virou https://www.unabeleza.com.br/ mas já migrou pra Nuvem"*. A Unabeleza apareceu também em 2026-01-15 (Wake <> Econverse) em discussão de % aplicado em opps fechadas: *"Na planilha não vi os nomes de Kikkaboo - Soneda - Unabeleza"* — sinal de que o relacionamento comercial já estava se desfazendo no 1ºT/2026.

**Solução:** Não aplicável — saída para Nuvemshop.

**Resultado:** ❌ Churn confirmado. Cliente operando em **Nuvemshop** sob nova marca (Unabeleza) em 03/2026.

**Insights / lessons:**
- **Churn para Nuvemshop** é padrão menos comum no corpus (predominante: VTEX, Shopify, Bagy) — vale registrar para battlecard Nuvemshop.
- Padrão **rebranding + migração** (Unikapro → Unabeleza + Nuvemshop) sinaliza que a saída foi planejada com tempo, não fricção operacional aguda — investigar com Sales Ops se houve sinal antecipado de NPS/CSAT que pudesse ter previsto.
- Ausência da Unabeleza em **planilha de comissionamento Econverse** em 01/2026 foi sinal antecedente de 2 meses — caso para Coordenação Parcerias usar ausências de comissionamento como **leading indicator de churn**.

---

### CASE-071 — Empório Bud (lead via Konduto — vinhos/bebidas?)

| Campo | Valor |
|---|---|
| **Cliente** | Empório Bud |
| **Setor** | ? (não identificado claramente — nome sugere bebidas/cervejas, mas não confirmado) |
| **Tipo de operação** | B2C |
| **Status** | ⚠ lead em qualificação (último update 09/2025) |
| **URL** | `emporiobud.com.br` ✓ confirmado (Konduto & Wake, 2025-09-12) |
| **Plataforma anterior** | — |
| **ERP** | — |
| **Marketplace setup** | — |
| **Parceiros envolvidos** | **Konduto** (indicação inicial em evento Konduto), Oásis (convite de 07/2025 mencionando contato) |
| **Tags** | `audience-pam`, `lead-pipeline` |
| **Data da última menção** | 2025-09-12 |

**Situação:** Empório Bud foi **indicado pela Konduto** em 2025-09-12 (Konduto & Wake): *"Eles pediram minha ajuda para fazer a ponte com vocês em nosso evento. Seguem os dados para entrarem em contato: https://emporiobud.com.br / fernando@emporiobud.com.br / 11 98121-0915 - Fernando Felix / CNPJ: 11.638.545/0001-44"*. Contato anterior em 2025-07-06 (Wake <> Oásis) capturava o mesmo Fernando Felix como "Convite 3" em uma lista de prospects. Em 2025-09-10, WAKE | Time Parcerias pediu: *"Gente, Emporio bud. Preciso de parceiro logístico, pagar.me e conduto"* — sinal de que o discovery estava progredindo (precisava de parceiro logístico + Pagar.me + Konduto). Em 2025-07-01, WAKE | Time Agência reportou: *"to caçando a tal da emporio bud ainda kkkkkkkkkk"* — sinal de dificuldade para encontrar account owner.

**Solução:** Wake Commerce + **Pagar.me** + **Konduto** (antifraude) + parceiro logístico (a definir).

**Resultado:** ⚠ Sem update após 2025-09-12. CNPJ capturado mas status do contrato desconhecido.

**Insights / lessons:**
- **Konduto como originador de lead** via seu próprio evento (Konduto & Wake) — padrão de **partner-sourced lead** útil para pitch ao programa Konduto.
- Stack discutida em 2025-09-10 (**Pagar.me + Konduto + logística**) é stack canônica Wake para SMB — referência de "kit completo" mínimo.
- Caso de **demora para identificar account owner interno** ("to caçando a tal da emporio bud ainda kkkkkkkkkk") é sintoma de fluxo de distribuição confuso quando o lead vem fora do carrossel normal.

---

### CASE-072 — Você Constrói (B2C construção — parent brand do Esplane)

| Campo | Valor |
|---|---|
| **Cliente** | Você Constrói |
| **Setor** | material de construção / varejo |
| **Tipo de operação** | B2C — varejo de material de construção (marca-mãe da Esplane CASE-051) |
| **Status** | ✓ ativo (Anymarket liberado desde 2025-08-22) |
| **URL** | `voceconstroi.com.br` ✓ confirmado (Eficaz <> WAKE 2025-12-15 / 2026-02-18) |
| **Plataforma anterior** | — |
| **ERP** | — |
| **Marketplace setup** | **Anymarket** (liberado em 2025-08-22) |
| **Parceiros envolvidos** | Eficaz (sustentação técnica), Auaha (frontend), Anymarket |
| **Tags** | `audience-cs`, `integração` |
| **Data da última menção** | 2026-02-18 |

**Situação:** Você Constrói é a **frente B2C** que opera junto com a **Esplane** (CASE-051, B2B) no mesmo grupo. Anymarket foi liberada em 2025-08-22 (ANYMARKET & WAKE: *"Você Constrói já liberado!"* + *"estou na sala com a Você Constrói já"*). A frente B2C tem **chamado técnico crítico ativo** em sustentação via Eficaz: em 2025-12-15 (Eficaz <> WAKE) abriu-se ticket 258709 sobre *"O evento de Purchase está disparando no Meta, mas não está enviando o parâmetro de receita (value) junto. O Meta registra a venda, mas com valor R$ 0,00, o que destrói as métricas de ROAS e a otimização das campanhas"*. Problema voltou após ajuste no front da Auaha em 11/2025 (*"o evento Purchase passou a não contabilizar a receita"*) e em 2026-02-18 ainda estava aberto: *"precisamos de ajuda urgente com o cliente Voce constroi"* — chamado escalado.

**Solução:** Wake Commerce + Anymarket + Eficaz (sustentação) + Auaha (front). Esplane (CASE-051) opera como loja Wake separada para o canal B2B.

**Resultado:** Ativo em produção, mas com **fricção persistente no evento Purchase do Meta** desde 11/2025 — degrada ROAS/atribuição de campanha. Caso técnico ainda aberto em 02/2026.

**Insights / lessons:**
- Padrão **B2C parent + B2B child como duas lojas Wake separadas** confirmado (Você Constrói + Esplane) — referência canônica para pitch de arquitetura B2B segregada.
- Caso operacional do **bug recorrente do evento Purchase no Meta com receita zerada** — sinal de gap conhecido na camada de tagueamento headless/Wake quando o front é customizado por terceiros (Auaha). Útil em discovery de qualquer cliente que vá rodar Wake + agência de front externa + Meta Ads.
- Caso de **escalada urgente** que cruza Wake + Eficaz + Auaha — referência para playbook de **war room** quando o bug atravessa 3 fornecedores.

---

### Programa de afiliados — referências cruzadas

A Wake não tem "produto de afiliados" separado. Programas de afiliados/influencer marketing são montados sobre **dois recursos da plataforma**:

1. **Painel Wake — módulo de campanha/cupom**: *"módulo de campanha/cupom integrado na plataforma. Ele é bem completo, envolve variáveis para estratégia do cliente com os influ/afiliado"* (Monitfy & Wake, 2026-05-18). Limitação conhecida: a API correspondente é mais restrita que o painel — operações pelo painel são mais ricas que via integração programática.
2. **StoreFront API — `partner/afiliado` no carrinho**: *"Estamos usando a StoreFront API da Wake. Utilizamos duas mutations para gerar a url do novo carrinho. CheckoutClone (para clonar um carrinho prévio, incluindo o `partner/afiliado`) e CheckoutAddMetadata para colocar as nossas UTM no novo carrinho"* (Uble & Wake, 2026-01-19). Padrão técnico documentado para preservar identificação do afiliado/parceiro ao longo da sessão de checkout.

A referência interna canônica que enumera os três cases de afiliados Wake é **Comercial Cross WAKE, 2025-12-17**:

> *"Ybera / Shoulder para funcionários / BAW com influenciadores"*

| Cliente | Modelo de afiliados | Caso completo |
|---|---|---|
| **Shoulder** | Programa de afiliados **para funcionários** (rede de vendedoras + colaboradoras como afiliadas) | **CASE-004** — Shoulder |
| **BAW** | Programa de afiliados com **influenciadores** | **CASE-017** — Baw |
| **Ybera** | Cliente em **negociação concorrida com VTEX** que **usa programa de afiliados** como parte do diferencial operacional | **CASE-047** — Ybera |

**Insights / lessons:**
- A Wake **não vende** "produto de afiliados" — vende plataforma + StoreFront API. O programa é construído pela marca (com ou sem agência), explorando cupom + atribuição de `partner/afiliado` via StoreFront.
- Os três casos cobrem **três modelos distintos** de programa de afiliados: funcionários internos (Shoulder), influenciadores externos (BAW), e blend operacional dentro de relacionamento direto-com-consumidor (Ybera). Útil em pitch a marca que pergunta *"vocês suportam programa de afiliados?"* — resposta: "sim, e temos três modelos diferentes em produção".
- Limitação atual a sinalizar em discoveries: **gestão de promoções via API é limitada** (Monitfy & Wake, 2026-05-18) — para programas que exigem orquestração programática complexa, validar escopo técnico antes de fechar.
- Em conversas comerciais externas, há marcas adjacentes citando a mesma necessidade — *"a ideia de ter uma área de afiliados organizada"* (Ello Moda, Nova Oportunidade 2026-05-08, usando DLK como benchmark). Sinaliza que **programa de afiliados estruturado é argumento de venda crescente** na vertical moda em 2026.

---

### Verificação — B2B, Afiliados e Verticais Especializadas

| Claim | Status | Query (LIKE) | Evidence |
|---|---|---|---|
| Lista canônica de cases B2B Wake é "CVC / GOLD / Destro / A Fabrica" | ✓ verificado | `%B2B%` + `%CVC%` + `%Destro%` | Comercial Cross WAKE, 2026-03-05 17:42 |
| Esplane é B2B da Você Constrói | ✓ verificado | `%Esplane%` (8 hits) | Comercial Cross WAKE 2025-12-17 + Wake <> Nairuz 2026-03-20 |
| CVC opera plataforma para franquias comprarem da matriz | ✓ verificado | `%CVC%franquia%` + `%CVC%matriz%` | Comercial Cross WAKE 2025-12-17; Wake <> Nairuz 2026-03-20 |
| CVC + Buddemeyer são benchmark interno de B2B "completo" (crédito, regionalização, listas de preço) | ✓ verificado | `%CVC%benchmark%`; `%Buddemeyer%CVC%` | CoreBiz & Wake, 2026-01-30 17:21 |
| Yora é atacado de produtos kids | ✓ verificado | `%Yora%` | Comercial Cross WAKE 2025-12-17; Wake <> Nairuz 2026-03-20 |
| Destro é atacadista do PR | ✓ verificado | `%Destro%PR%`; `%Destro Macro%` | Wake <> Nairuz, 2026-03-20 |
| Buddemeyer assinou Commerce R$ 48,1k MRR + R$ 21k setup em 2025-06-30 | ✓ verificado | `%Buddemeyer%MRR%` | Mavericks + Comercial Cross WAKE, 2025-06-30 19:06 |
| Buddemeyer B2B "ainda em projeto" em 2026-03 | ✓ verificado | `%Buddemeyer%B2B%` | Comercial Cross WAKE, 2026-03-05 18:44 |
| Gold Food Service é case B2B Wake | ✓ verificado | `%Gold Food%`; `%goldfoodservice%` | Comercial Cross WAKE 2026-03-05 18:47 |
| A Fábrica usa `checkout.afabrica.com` (subdomínio dedicado) | ✓ verificado | `%afabrica%`; `%checkout.afabrica%` | Comercial Cross WAKE 2026-03-05 18:50 |
| Unitoys assinou R$ 5k MRR + R$ 24k setup em 2026-01-21 | ✓ verificado | `%Unitoys%assinad%` | Comercial Cross WAKE, 2026-01-21 13:18 |
| Unitoys é referência interna para vertical brinquedos | ✓ verificado | `%unitoys%brinquedo%` | Comercial Cross WAKE, 2026-03-17 20:54 |
| Editora Fórum é cliente Wake em projeto Pagbank + checkout Headless via Benova | ✓ verificado | `%Editora Fórum%` (5 hits) | Wake <> Benova, 2025-08-26 |
| Wake suporta programa de afiliados via StoreFront API (CheckoutClone + CheckoutAddMetadata) | ✓ verificado | `%afiliado%CheckoutClone%`; `%partner/afiliado%` | Uble & Wake, 2026-01-19 13:10 |
| Wake tem módulo de campanha/cupom no painel para afiliados, mas API é limitada | ✓ verificado | `%campanha%cupom%`; `%influ%afiliado%` | Monitfy & Wake, 2026-05-18 16:58 e 20:46 |
| Os 3 cases Wake de afiliados são Ybera / Shoulder funcionários / BAW influenciadores | ✓ verificado | `%Ybera%Shoulder%BAW%`; `%Shoulder para funcionários%` | Comercial Cross WAKE, 2025-12-17 19:16 |
| Você Constrói está usando Anymarket desde 2025-08 | ✓ verificado | `%Você Constrói%liberado%` | ANYMARKET & WAKE, 2025-08-22 15:34 |
| URL `forumconhecimento.com.br` é da Editora Fórum | ⚠ fornecido pelo usuário, não confirmado em mensagens | `%forumconhecimento%` (0 hits) | n/a — URL não citada em mensagens do corpus |
| URL `negocios.forumconhecimento.com.br` é da Editora Fórum | ⚠ fornecido pelo usuário, não confirmado em mensagens | `%negocios.forumconhecimento%` (0 hits) | n/a |
| URL da Buddemeyer | ⚠ fornecido pelo usuário, não confirmado em mensagens | n/a | URL não citada nas mensagens analisadas |

---

### Verificação — Lote 2 (CASE-060 a CASE-072)

| Claim | Status | Query (LIKE) | Evidence |
|---|---|---|---|
| Gregory roda provador virtual Metakosmos em produção | ✓ verificado | `%Gregory%` (67 hits); `%Gregory%provador%` | Wake & Metakosmos, 2026-05-14 12:35; 2026-03-11 19:38 |
| Gregory tem sustentação técnica via Floui com SLA contratual de 15 dias | ✓ verificado | `%Gregory%Floui%`; `%SLA%15 dias%` | Floui & Wake, 2026-05-26 12:51 e 17:27 |
| Wake nunca cobrou revenue share Metakosmos por Gregory até 02/2026 | ✓ verificado | `%Metakosmos%nunca cobramos%` | WAKE \| Coordenação, 2026-02-20 15:16 |
| Ave Rara teve Go-Live em 2026-05-18 com averarabrand.com.br | ✓ verificado | `%averarabrand%` (14 hits) | The Rise of Skywalker, 2026-05-18 13:55 |
| Ave Rara foi o primeiro case da agência Simples Inovação | ✓ verificado | `%Averara%primero case%`; `%Simples%Averara%` | Simples Inovação & Wake Parcerias, 2026-05-19 16:11 |
| Ave Rara fechou Commerce R$ 55k MRR + R$ 68,9k setup em 02/2026 | ✓ verificado | `%AveRara%55k%` | The Rise of Skywalker, 2026-02-26 21:50 |
| Ave Rara em negociação ativa de Prateleira Infinita + POS Stone em 06/2026 | ✓ verificado | `%AVERARA%prateleira%`; `%Averara%POS%` | Comercial Cross WAKE, 2026-06-03 14:29; [work] Wake & Stone, 2026-06-02/03 |
| OffPremium é piloto Marketplace Blueman + Anymarket | ✓ verificado | `%OffPremium%` (23 hits); `%offpremium%Blueman%` | Wake + Marketplace Blueman, 2025-10-29 a 2026-01-28 |
| OffPremium tem gap conhecido no número de título do pedido (não é marketPlacePedidoId) | ✓ verificado | `%OffPremium%marketPlacePedidoId%` | Wake + Marketplace Blueman, 2025-11-03 22:50 |
| Eletro Mais MT estava saindo de OpenCart com ERP Santri Sistemas | ✓ verificado | `%eletromaismt%`; `%Open Cart%` + `%Santri%` | Duo e Wake Parcerias, 2025-12-10 18:43 |
| Cimafer saía de Shopify (não OpenCart como sugeria o gap-list) com ERP Boreo | ✓ verificado | `%Cimafer%`; `%Cimafer%Shopify%`; `%Cimafer%Boreo%` | Duo e Wake Parcerias, 2025-12-03 17:21; 2025-12-10 18:43 |
| Cimafer claim GMV R$ 1M/mês (e-commerce + loja física combinado) | ✓ verificado | `%Cimafer%1.000.000%` | Duo e Wake Parcerias, 2025-12-03 17:21 |
| Cimafer esfriou em 01/2026 (não atende celular) | ✓ verificado | `%Cimager%não atende%` | DRIVEN & Wake, 2026-01-12 19:45 |
| Brofitwear fez Go-Live em 2026-05-19 como primeiro Go-Live da Quick Digital | ✓ verificado | `%Brofitwear%Go Live%`; `%Quick%Brofitwear%` | Quick Digital & Wake Parcerias, 2026-05-19 14:03 |
| Brofitwear saía de VTEX, indicação via Catus | ✓ verificado | `%Brofitwear%Vtex%`; `%Brofitwear%Catus%` | Wake <> Catus, 2025-11-11 12:15 e 12:21 |
| Brofitwear assinou Commerce R$ 6,1k MRR + R$ 12k setup em 2025-10-30 | ✓ verificado | `%Brofitwear%assinad%` | Mavericks + Comercial Cross WAKE, 2025-10-30 18:15/33 |
| Casafer saía de WooCommerce com SAP B1 via Eficaz/O2 | ✓ verificado | `%Casafer%WooCommerce%`; `%Casafer%SAP%` | Eficaz <> WAKE / O2 Consultoria & Wake, 2025-06-30 13:15-21 |
| Singular Medicamentos é cliente Wake em produção com painel ativo | ✓ verificado | `%Singular%Medicamentos%` + `%painel%` | ANYMARKET & WAKE, 2025-11-06 13:26 |
| Singular Medicamentos tem requisito manual de validação de receita médica | ✓ verificado | `%Singular%receita%` | Duo e Wake Parcerias, 2026-06-02 12:43 |
| Singular Medicamentos foi indicada para App LeanWork junto com DLK/Luiza Barcelos/YBERA | ✓ verificado | `%Singular%LeanWork%` | Wake & Leanwork, 2025-10-06 14:37 |
| OG Brindes pediu indicação de agência de sustentação em 06/2025 | ✓ verificado | `%OG Brindes%sustentação%`; `%ogbrindes%agência%` | FG + Wake / Wake <> Série//A / Wake <> Econverse, 2025-06-24 21:35-36 |
| OG Brindes contato comercial é José Carlos Grzybowski (Diretor Geral) | ✓ verificado em DM | `%Grzybowski%` | FG + Wake, 2025-06-24 21:36 (contato confidencial — não publicar) |
| Biggie/abiggie foi desqualificado (cliente Shopify+Tiny não quer migrar) | ✓ verificado | `%Biggie%Desqualificado%`; `%abiggie%Shopify%` | WAKE \| Time Agência, 2025-10-28 14:07; Wake <> Ge Commerce, 2025-10-09 19:03 |
| Catus foi prejudicada e pediu compensação por leads Biggie + Trocamotor desqualificados | ✓ verificado | `%Catus%compensar%Biggie%` | WAKE \| Time Agência, 2025-10-28 14:07 |
| Unikapro virou Unabeleza e migrou para Nuvemshop em 03/2026 | ✓ verificado | `%Unikapro%Unabeleza%`; `%unabeleza%Nuvem%` | WAKE \| Time Parcerias, 2026-03-27 14:08 |
| Empório Bud foi indicação Konduto via evento Konduto em 09/2025 | ✓ verificado | `%emporiobud%Konduto%`; `%Emporio bud%` | Konduto & Wake, 2025-09-12 13:06; WAKE \| Time Parcerias, 2025-09-10 17:42 |
| Você Constrói teve Anymarket liberado em 2025-08-22 | ✓ verificado | `%Você Constrói%liberado%` | ANYMARKET & WAKE, 2025-08-22 15:34 |
| Você Constrói tem bug recorrente no evento Purchase do Meta (receita zerada) — ticket 258709 aberto em 12/2025 e ainda escalado em 02/2026 | ✓ verificado | `%voceconstroi%Purchase%`; `%Voce constroi%ROAS%` | Eficaz <> WAKE, 2025-12-15 16:43 e 2026-02-18 14:15 |
| Você Constrói teve ajuste de front feito pela Auaha em 11/2025 que originou o bug | ✓ verificado | `%Voce constroi%Auaha%` | Eficaz <> WAKE, 2026-02-18 14:15 |

**Não escritos** (avaliados durante este sweep mas sem evidência suficiente ou contexto fora de escopo):

| Candidato | Razão | Evidence |
|---|---|---|
| Belluno | **Não é cliente** — é parceiro de pagamento (Belluno Pag, `bellunopag.com.br`), avaliado como gateway alternativo / "estepe caso Appmax não aceite" (WAKE \| Time Parcerias, 2026-02-09 19:42-44; Belluno + Wake, 2026-03-26). Grupo dedicado "Belluno + Wake" reflete relacionamento de parceria, não case de customer. | WAKE \| Time Parcerias 2025-07-18 a 2026-06-03; Belluno + Wake 2026-03-26 |

---

## 9. Leads e prospects de Canais — segunda leva (CASE-073–CASE-102)

Esta seção cobre prospects/leads de Canais identificados no sweep de gap analysis (2026-06), tier high-confidence (3+ hits no DB). A maioria são oportunidades distribuídas no modelo carrossel (CASE-049) para 3–5 agências — leads em fase de qualificação ou negociação. Casos com Go-Live ou contrato assinado são marcados explicitamente. Casos sem desfecho confirmado ficam como `⚠ em negociação`, `❌ perdido` ou `⚠ status desconhecido`.

**Fonte**: `docs/wake-customer-cases-gap-list.md` (high-confidence tier, candidatos não cobertos pelos "notable findings" CASE-060–CASE-072).

---

### CASE-073 — Cerealista Express

| Campo | Valor |
|---|---|
| **Cliente** | Cerealista Express (`cerealistaexpress.com.br`) |
| **Setor** | alimentos / cereais e mercearia online |
| **Status** | ⚠ em negociação |
| **Plataforma anterior** | Magento ✓ verificado |
| **Visualizações** | 25,5K/mês |
| **Parceiros envolvidos** | Avocado Tech, Ge Commerce, O2 Consultoria (rotação de canais) |
| **Data da última menção** | 2026-02-10 |

**Situação:** Lead de Canais distribuído em 2025-11-27 a 3 agências. *"Felipe é Sócio, e atualmente é cliente da Magento. Ele explicou que possui um e-commerce e que a plataforma responsável pelo site já não o satisfaz mais... investiu um valor alto... paga uma mensalidade elevada"* (Wake <> Avocado Tech, 2025-11-27). Status em 2026-01-26: agências aguardando retorno do cliente. Em 2026-02-10, ReclameAqui pergunta *"este cliente roda com vocês?"* — confirmando que ainda não havia migrado.

**Insights:**
- **Migração Magento → Wake** com dor de custo de manutenção alto e plataforma "engessada".
- Rotação de 3 agências (Avocado / Ge / O2) deixou o lead esfriar — sinal de falha de follow-up.

---

### CASE-074 — Eletrônica Santana (Grupo)

| Campo | Valor |
|---|---|
| **Cliente** | Grupo Eletrônica Santana — Eletrônica Santana + ES TECH + Dealer Shop (B2B) (`eletronicasantana.com.br`) |
| **Setor** | varejo eletro / B2C + B2B |
| **Status** | ✓ ativo |
| **Parceiros envolvidos** | Ecommerce House (agência), Biso (busca) |
| **Data da última menção** | 2026-04-08 |

**Situação:** Cliente Wake confirmado em produção. *"O confirmado é Eletrônica Santana, que é nosso cliente"* (Wake & Biso, 2025-09-16). Em 2026-02-03, Ecommerce House anuncia *"Colocamos no ar o novo modelo de PDP do site inteiro da Eletrônica Santana, com foco em melhorar a experiência do cliente, aumentar a conversão e corrigir alguns bugs"*. Em 2026-04-08, vaga aberta para Analista de E-commerce do grupo cuidando das 3 marcas — sinaliza escala da operação.

**Solução:** Wake Commerce com **3 lojas no grupo** (B2C + B2B Dealer Shop) operadas por Ecommerce House. Biso é o parceiro de busca.

**Insights:**
- Case de **grupo multi-marca varejo eletro com braço B2B** (Dealer Shop) — útil para discoveries com varejistas eletro que querem juntar canais.
- Caso de **operação contínua via Ecommerce House** — agência ativa publicando releases de PDP.

---

### CASE-075 — Inoxlon

| Campo | Valor |
|---|---|
| **Cliente** | Inoxlon (`inoxlon.com.br`) |
| **Setor** | industrial / utensílios em aço inox |
| **Status** | ⚠ em negociação (sem contato em 2026-02) |
| **Plataforma anterior** | Climba ✓ verificado |
| **ERP** | Bling |
| **Faturamento online** | R$ 200k/mês (GMV) |
| **Parceiros envolvidos** | Engajar, A&EIGHT (B8One), Cava Digital (carrossel) |
| **Data da última menção** | 2026-02-23 |

**Situação:** Lead SDR distribuído em 2026-02-12: *"INOXLON- Canais. ERP: Bling. GMV: 200k mês no online. Possui 1 loja física. Chegou na Wake através das campanhas e anúncios do Instagram... cliente da plataforma Climba"* (Wake & A&EIGHT B8One, 2026-02-12). Em 2026-02-19, email de contato deu como inexistente. Em 2026-02-23: *"INOXLON não te respondeu?"* — lead esfriou.

**Insights:**
- Lead de **inbound (Instagram ads Wake)** que esfriou por dados de contato incorretos.
- Migração **Climba → Wake** com Bling nativo — qualificação SMB padrão.

---

### CASE-076 — AZ Acessórios (Grupo Cambuci)

| Campo | Valor |
|---|---|
| **Cliente** | AZ Acessórios (`azacessorios.com.br`) |
| **Setor** | moda / acessórios — Grupo Cambuci |
| **Status** | ✓ ativo |
| **Parceiros envolvidos** | Eficaz (agência) |
| **Data da última menção** | 2026-03-16 |

**Situação:** Go-Live anunciado em 2025-06-23: *"Saiiuuu pessoal site da AZ acessórios esta no AR, case gigantesco do grupo Cambuci"* (Eficaz <> WAKE, 2025-06-23). Em 2026-03-04, time discute migração para o novo storefront: *"AZ Acessórios, está no novo storefront?"*. Em 2026-03-16, time usa AZ Acessórios como referência em discovery de outro prospect: *"mostrei o site da AZ ACESSÓRIOS"*. Em 2026-03-04, time da AZ visitou o stand Wake (provável evento ECBR/Wake Summit).

**Solução:** Wake Commerce operado pela Eficaz. **Cliente do Grupo Cambuci**.

**Insights:**
- Case de **referência viva** usado em discoveries de outros prospects (Eficaz é forte em autopeças/acessórios).
- Conexão Grupo Cambuci ⚠ inferido — relação societária com outros clientes Wake não confirmada no corpus.

---

### CASE-077 — Dallanese Parafusos

| Campo | Valor |
|---|---|
| **Cliente** | Dallanese Parafusos (`dallanese.com.br`) |
| **Setor** | industrial / ferramentas e parafusos |
| **Status** | ❌ sem retorno (lead frio) |
| **ERP** | Sankhya |
| **Faturamento online** | R$ 50k/mês |
| **Parceiros envolvidos** | A&EIGHT (B8One), Duo, Ge Commerce (carrossel) |
| **Data da última menção** | 2025-07-21 |

**Situação:** Lead SDR distribuído em 2025-06-25: *"Dallanese Parafusos - Commerce - Canais. Necessidade: Migração de e-commerce. ERP: Sankhya. Faturamento online: 50.000. Vendem nos marketplaces / ML. 1 loja física"* (Duo e Wake Parcerias, 2025-06-25). Em 2025-07-15, B8One reporta: *"Sem sucesso em contato com a Dallanese"*. Em 2025-07-21, Duo perguntando *"Dallanese tem algum retorno????"* — sem resposta após.

**Insights:**
- **ERP Sankhya** precisa integrador (ver playbook §3.1 — Integra.do ou Floui).
- Padrão **lead distribuído + sem follow-up do cliente** — esfriou em 4 semanas.

---

### CASE-078 — Balletto

| Campo | Valor |
|---|---|
| **Cliente** | Balletto (`balletto.com.br`) |
| **Setor** | moda / calçados |
| **Status** | ⚠ em negociação |
| **Plataforma anterior** | VTEX ✓ verificado |
| **ERP** | Sisplan |
| **Parceiros envolvidos** | O2 Consultoria, 3W, Eficaz (carrossel) |
| **Data da última menção** | 2025-07-22 |

**Situação:** Lead vindo de **indicação DXP→Commerce** distribuído em 2025-07-11: *"Balletto - Commerce - Canais. Oportunidade vinda de indicação de DXP para Commerce... Migração de VTEX"* (Wake <> 3W, 2025-07-11). Em 2025-07-22, O2 reporta: *"Fiz a apresentação da proposta para a Balletto... Estão considerando a operação do Digital com a Loja Física adicional. Contudo não ficou claro pra mim sobre a integração com o ERP deles Sisplan"* (O2 Consultoria & Wake, 2025-07-22).

**Insights:**
- **Migração VTEX → Wake** indicada via DXP (produto da própria Wake) — bom canal interno.
- **ERP Sisplan**: não-nativo, precisa integrador.

---

### CASE-079 — AR15 Sport

| Campo | Valor |
|---|---|
| **Cliente** | AR 15 Sport (`ar15sport.com.br`) |
| **Setor** | esporte / airsoft e equipamentos de lazer |
| **Status** | ⚠ em negociação (lead frio, sem retorno) |
| **Plataforma anterior** | Tray ✓ verificado |
| **Parceiros envolvidos** | E-Plus, Lemoon (carrossel) |
| **Data da última menção** | 2025-09-29 |

**Situação:** Lead SDR distribuído em 2025-07-22: *"AR15 Sport - Commerce - Canais. Plataforma atual: Tray. Segmento: Air soft e equipamentos de lazer... Necessidade: Migração de solução"* (Wake <> E-Plus, 2025-07-22). Em 2025-09-19 a 2025-09-29, WAKE | Time Agência reporta dificuldade: *"sigo tentando (e sem nem respostas) da AR15"* (2025-09-23).

**Insights:**
- Vertical **nicho (airsoft)** — útil em pitch a marcas de esporte/outdoor.
- Tomador de decisão **difícil de engajar** — padrão SMB autodidata.

---

### CASE-080 — Campinas Diesel

| Campo | Valor |
|---|---|
| **Cliente** | Campinas Diesel (`campinasdiesel.com.br`) |
| **Setor** | automotivo / peças e diesel |
| **Status** | ⚠ em negociação |
| **Parceiros envolvidos** | Catus, Eficaz, Proex (carrossel) |
| **Data da última menção** | 2025-08-19 |

**Situação:** Lead SDR distribuído em 2025-08-12 via Msam: *"Campinas Diesel... Mensagem original do lead: site com vendas, integracao mercado livre... Hoje trabalham com loja física, e querem um e-commerce B2C e que tenha integr[ação ML]"* (Eficaz <> WAKE, 2025-08-12). Em 2025-08-15, Catus reporta: *"Campinas Diesel: Show, se precisar de algo, conta comigo!"*. Em 2025-08-19, Eficaz lista Campinas Diesel em status report consolidado — em negociação ativa.

**Insights:**
- Caso **automotivo + integração Mercado Livre** — pitch padrão para autopeças (combinar com Anymarket).

---

### CASE-081 — 7 Oliveiras

| Campo | Valor |
|---|---|
| **Cliente** | 7 Oliveiras (`7oliveiras.com.br`) |
| **Setor** | indústria / alimentos |
| **Status** | ⚠ em rotação carteira (2026-01) |
| **Plataforma anterior** | Visual e-commerce ✓ verificado |
| **ERP** | ZADA |
| **Faturamento online** | R$ 20k site + R$ 160k marketplaces |
| **Parceiros envolvidos** | Simples Inovação, Proex, FZ Commerce (carrossel) |
| **Data da última menção** | 2026-01-22 |

**Situação:** Lead SDR distribuído em 2025-08-18: *"https://www.7oliveiras.com.br/. trabalham com a plataforma Visual e-commerce... a parte dos marketplaces é muito importante p/ eles, pois hoje vende mais por esse canal. Faturamento site: Em torno de 20k. Faturamento marketplaces: 160k. ERP: ZADA"* (Simples Inovação & Wake Parcerias, 2025-08-18).

**Insights:**
- **GMV majoritariamente em marketplaces** (R$ 160k vs R$ 20k site) — pitch deve enfatizar Anymarket como hub central.
- **ERP ZADA**: não-nativo, validar integrador antes da proposta.

---

### CASE-082 — Conceito Rouparia

| Campo | Valor |
|---|---|
| **Cliente** | Conceito Rouparia (`conceitorouparia.com.br`) |
| **Setor** | moda / multi-marca |
| **Status** | ⚠ em rotação carteira (2026-01) |
| **Plataforma anterior** | Loja Integrada ✓ verificado |
| **ERP** | Datasystem (manter) — integrador 4middleware (via API) |
| **Faturamento online** | < R$ 300k |
| **Parceiros envolvidos** | Simples Inovação, FZ Commerce (carrossel) |
| **Data da última menção** | 2026-01-22 |

**Situação:** Lead SDR distribuído em 2025-09-03: *"Plataforma atual: Loja Integrada. Necessidade: Migração de commerce. usavam a Neomode para integrar, pois o ERP é o Datasystem e a idéia é manter ele. Porém, encerraram o contrato com a Neomode e gostariam de mudar a plataforma de e-commerce. gostariam de usar como integrador o 4middleware. Validei com parceria tech, e seria possível via API (cliente ciente). São 5 lojas físicas de segmentos diferentes"* (Simples Inovação & Wake Parcerias, 2025-09-03).

**Insights:**
- **Multi-segmento físico (5 lojas)** com necessidade de manter ERP Datasystem — caso de iPaaS via 4middleware.
- Padrão **iPaaS via API com integrador independente** quando ERP não tem conector nativo — validado por Wake Parceria Tech.

---

### CASE-083 — Carinion Mix

| Campo | Valor |
|---|---|
| **Cliente** | Carinion Mix (`carinonmix.com.br`) |
| **Setor** | multi-categoria |
| **Status** | ❌ perdido para VNDA |
| **Plataforma anterior** | HostGator |
| **Visualizações** | 1.2K/mês |
| **Parceiros envolvidos** | E-Plus, Eficaz, Dotkom (carrossel) |
| **Data da última menção** | 2025-09-19 |

**Situação:** Lead SDR distribuído em 2025-09-15: *"1.2K de visitas por mês no site. O cliente possui uma loja virtual, mas a plataforma atual apresenta limitações significativas (Loja HostGator)"* (Eficaz <> WAKE, 2025-09-15). Em 2025-09-19, Eficaz reporta perda: *"Carinion mix, avisa a Eficaz que fechou com VNDA devido eles terem diversas precificações acessíveis, sendo a maior R$600,00 mensais. Achou nossa plataforma com investimento alto, ou seja, sem perfil"* (Eficaz <> WAKE, 2025-09-19).

**Insights:**
- **Perdido para VNDA por preço** (R$ 600/mês teto cliente vs proposta Wake) — SMB low-end sem perfil Wake.
- VNDA aparece como **competidor recorrente em SMB low-end** (também citado em CASE-085 Zinnia como churn-source).

---

### CASE-084 — Believe Professional + Tazin Cosméticos (lote beleza SMB)

| Campo | Valor |
|---|---|
| **Cliente A** | Believe Professional (`believepro.com.br`) — beleza profissional |
| **Cliente B** | Tazin Cosméticos (`tazincosmeticos.com.br`) — cosméticos |
| **Status** | ⚠ ambos em negociação (rotação carteira 2026-01) |
| **Plataforma anterior** | Believe: Loja Integrada; Tazin: Nuvem Shop |
| **ERP / custo atual** | Believe: ERP Bling; Tazin: paga R$ 7.500/mês para agência |
| **Parceiros envolvidos** | Simples Inovação, FZ Commerce, Lemoon, Proex (carrossel) |
| **Data última menção** | 2026-01-22 |

**Believe Professional** (lead SDR 2025-09-22): *"Plataforma atual: Loja Integrada. Estão em busca de uma nova solução, pois a LI esta deixando a desejar. Vendem somente na loja física + e-commerce. ERP: Bling"* (Simples Inovação & Wake Parcerias, 2025-09-22).

**Tazin Cosméticos** (lead SDR 2025-09-24): *"Plataforma atual: Nuvem Shop. Reclamou bastante dizendo que o e-commerce não está vendendo, sendo que os gastos atuais estão altos. Comentou que hoje paga R$ 7.500 p/ uma agência"* (Wake <> Lemoon, 2025-09-24).

**Insights:**
- Padrão **Loja Integrada / Nuvem Shop → Wake** repetido em vários leads — pitch SMB beleza tem caminho conhecido.
- Custo R$ 7.500/mês com agência (Tazin) é **benchmark de SMB-mid** para proposta.

---

### CASE-085 — Regina Salomão

| Campo | Valor |
|---|---|
| **Cliente** | Regina Salomão (`reginasalomao.com.br`) |
| **Setor** | moda feminina |
| **Status** | ⚠ projeto adiado para 2026 (proposta com desconto vencida em 10/2025) |
| **Plataforma anterior** | Tray ✓ verificado |
| **ERP** | MV Tecnologia |
| **Faturamento online** | R$ 150k–200k/mês |
| **Parceiros envolvidos** | Avocado Tech, Simples Inovação, FZ Commerce (carrossel) |
| **Data da última menção** | 2026-01-06 |

**Situação:** Lead SDR distribuído em 2025-10-02. Em 2025-11-27, Simples Inovação reporta: *"Regina Salomão, a proposta deles com desconto e etc, venceu no mês 10... além disso, nossa precificação de 2026 pode sofrer reajuste"* — projeto passou para 2026. Em 2026-01-06, follow-up sem resposta.

**Insights:**
- **Projeto adiado com proposta desconto vencida** — risco de perda por reajuste 2026.
- **ERP MV Tecnologia**: não-nativo, validar integrador.

---

### CASE-086 — Grupo Loja do Detalhamento (Loja do Detalhamento + SGCB Brasil + IGL Coatings)

| Campo | Valor |
|---|---|
| **Cliente** | Grupo de 3 marcas detailing: `lojadodetalhamento.com.br`, `sgcbbrasil.com.br`, `loja.iglcoatings.com.br` |
| **Setor** | automotivo / detailing (estética automotiva) |
| **Status** | ⚠ em negociação multi-loja |
| **Plataforma anterior** | Tray ✓ verificado |
| **ERP** | Arth Magnus (precisa integrador) |
| **Parceiros envolvidos** | Simples Inovação, Proex, FZ Commerce (carrossel) |
| **Data da última menção** | 2025-11-11 |

**Situação:** Lead SDR distribuído em 2025-10-29 como **grupo de 3 lojas**: *"Loja do Detalhamento - Commerce - Canais. Plataforma atual: Tray. Necessidade: A idéia é migrar as 3 lojas e manter o mesmo estoque p/ todas"* (Simples Inovação & Wake Parcerias, 2025-10-29). Em 2025-10-31, discovery: *"hoje mais cedo conversei com o natã que é o mkt da Loja do Detalhamento, SGCB e IGL Coatings... eles usam o ERP Arth Magnus (nesse caso iremos precisar de um integrador) e têm um ponto importante: as três lojas vendem separadas, mas tudo sai do mesmo CD. Ele quer controlar estoque e pedidos num sistema único"* (Simples Inovação & Wake Parcerias, 2025-10-31).

**Insights:**
- Caso **multi-loja com estoque unificado** — pitch deve mostrar Wake OMS como diferencial.
- **ERP Arth Magnus**: não-nativo, integrador-ERP obrigatório.
- **Vertical detailing automotivo** é nicho relevante e recorrente.

---

### CASE-087 — Alfredo Meyer

| Campo | Valor |
|---|---|
| **Cliente** | Alfredo Meyer (`alfredomeyer.com.br`) |
| **Setor** | moda |
| **Status** | ⚠ em negociação |
| **Plataforma anterior** | WordPress ✓ verificado |
| **Visualizações** | 6.1K/mês |
| **Parceiros envolvidos** | Cava Digital, Proex, FZ Commerce (carrossel) |
| **Data da última menção** | 2025-12-05 |

**Situação:** Lead SDR distribuído em 2025-11-25: *"Plataforma atual: Wordpress. operação tem um volume relevante de pedidos e ticket médio alto, e estão estruturando o plano de crescimento p/ 2026. Lista de prioridades: Migração segura... Loja responsiva... Integrações com marketplaces, ERP, WhatsApp e RD Station..."* (Cava Digital & Wake Parcerias, 2025-11-25).

**Insights:**
- **Migração WordPress → Wake** com requisitos detalhados (RD Station, WhatsApp, ERP) — perfil de cliente mais maduro.
- **Plano 2026 ativo** — boa oportunidade de fechar com tabela 2026.

---

### CASE-088 — Festa Color

| Campo | Valor |
|---|---|
| **Cliente** | Festa Color (`festacolor.com.br`) |
| **Setor** | festas e decoração / confeitaria |
| **Status** | ❌ desqualificado |
| **Plataforma anterior** | nenhuma (primeira experiência online) |
| **ERP** | Sankhya (não vai vincular ao e-commerce) |
| **Faturamento loja física** | R$ 1,5M/mês |
| **Parceiros envolvidos** | Omni360, Affys, 3W, Eficaz (carrossel) |
| **Data da última menção** | 2025-12-05 |

**Situação:** Lead SDR distribuído em 2025-12-01. Cliente citou **Shop2gether e Puravida como referência**. Em 2025-12-05, desqualificado: *"Festa Color - Desqualificado (Cliente informou que entra em contato quando precisar)"* (WAKE | Time Agência, 2025-12-05).

**Insights:**
- **R$ 1,5M loja física** é bom perfil mas **maturidade digital zero** + sem intenção imediata = lead frio.
- Cliente citou **Shop2gether e Puravida como referência** — sinaliza pull-marketing dos cases públicos Wake.

---

### CASE-089 — Ivan Joias

| Campo | Valor |
|---|---|
| **Cliente** | Ivan Joias (`ivanjoias.com.br`) |
| **Setor** | joalheria (Salvador/BA) |
| **Status** | ⚠ em negociação |
| **Plataforma anterior** | Tray ✓ verificado |
| **ERP** | 4Control |
| **GMV** | R$ 150k/mês |
| **Lojas físicas** | 6 (shoppings em Salvador) |
| **Parceiros envolvidos** | FZ Commerce, Proex, Cava Digital (carrossel) |
| **Data da última menção** | 2026-01-22 |

**Situação:** Lead SDR distribuído em 2025-12-01: *"Atualmente cliente Tray. Eles se interessaram pela Wake, além de fazer parte do mesmo grupo que a Tray, eles analisaram o site da Sholuder antes de entrar em contato conosco, informando que o dono da empresa gostou da ideia de cashback e agenda vendedor. ERP: 4Control. GMV médio mensal: R$ 150 mil. Possuem 6 lojas físicas em shopping em Salvador (Bahia)"* (FZ Commerce, 2025-12-01).

**Insights:**
- **Prospect por benchmark Shoulder** — features-gatilho: cashback + agenda vendedor — confirma valor das referências públicas Wake.
- **Migração Tray → Wake dentro da mesma LWSA** — cliente procurando upgrade dentro da casa.

---

### CASE-090 — Sustent

| Campo | Valor |
|---|---|
| **Cliente** | Sustent Intermediação de Negócios (`sustent.com.br`) |
| **Setor** | beleza / marketplace B2B2C |
| **Status** | ⚠ em qualificação (marketplace próprio) |
| **ERP** | Omie + Bling |
| **Parceiros envolvidos** | Omni360, Affys, Cava Digital, FZ Commerce (carrossel) |
| **Data da última menção** | 2025-12-16 |

**Situação:** Lead SDR distribuído em 2025-12-16: *"O lead está procurando uma solução SaaS para criar um marketplace B2B2C, voltado para o segmento de beleza e cuidados pessoais. Ele quer ser um marketplace próprio, 100% online, sem loja física... Sobre custos, ele entendeu que o valor apresentado é para criação de e-commerce, mas para operar como marketplace será necessário incluir os valores de um parceiro da Wake"* (Wake+Omni360+Affys, 2025-12-16).

**Insights:**
- **Cliente quer ser marketplace próprio** (B2B2C) — Wake não tem produto marketplace próprio; depende de parceiro custom. Validar escopo com Solutions antes de prometer.
- Sinaliza necessidade de **playbook de marketplace próprio** distinto de vendas em marketplaces (Anymarket).

---

### CASE-091 — Shop2gether (Grupo Icomm)

| Campo | Valor |
|---|---|
| **Cliente** | Shop2gether (`shop2gether.com.br`) — Grupo Icomm |
| **Setor** | moda / varejo feminino multi-marca |
| **Status** | ✓ ativo (case público publicado) |
| **Plataforma anterior** | VTEX ✓ verificado |
| **Parceiros envolvidos** | Oásis (agência), DRIVEN (tentativa inicial, transferida a Oásis), Leanwork (app) |
| **Data da última menção** | 2026-05-28 |

**Situação:** **Fechamento em 2025-07-30**: *"Fecharam shop2gether"* (Mavericks, 2025-07-30); *"acabaram de fechar shop2gether"* (WAKE | Time Parcerias, 2025-07-30). Case oficial Wake publicado em 2025-08-18: `wake.tech/cases/como-a-wake-ajudou-o-grupo-icomm-shop2gether-e-oqvestir-com-sites-de-alta-performance/` (FZ Commerce, 2025-08-18). Em 2025-08-14: *"a DRIVEN começou fazer shop2gether, não conseguiu fazer, deu problema no projeto e foi ai que passou pra Oásis"* — transferência de agência. Em 2026-05-14, ticket de suporte sobre clonagem de cartão. Em 2026-05-28: *"Shop2gether tava normal ontem, mas vamos avisar o time"*.

**Solução:** Wake Commerce com checkout próprio operado pela Oásis. Cliente do **Grupo Icomm** (junto com OqVestir).

**Insights:**
- **Migração VTEX → Wake enterprise** com case público.
- **Transferência de agência durante implementação** (DRIVEN → Oásis): exemplo prático de quando trocar agência.
- **Citado como referência** em CASE-021 ACC Modas, CASE-088 Festa Color — funciona como pull-marketing.
- Combinar com **OqVestir** (mesmo grupo) — case multi-marca enterprise.

---

### CASE-092 — Zinnia Leather

| Campo | Valor |
|---|---|
| **Cliente** | Zinnia Leather (`zinnialeather.com.br`) |
| **Setor** | moda / acessórios de couro alto ticket |
| **Status** | ⚠ em negociação |
| **Plataforma anterior** | VNDA ✓ verificado |
| **Visualizações** | 1.4K/mês |
| **Parceiros envolvidos** | Engajar, A&EIGHT (B8One), Cava Digital (carrossel) |
| **Data da última menção** | 2026-02-02 |

**Situação:** Lead SDR distribuído em 2026-02-02: *"Zinnia- Canais. São clientes da VNDA, vendem peças de alto ticket. Temos necessidade de otimizar a reco[mendação]..."* (Wake <> Engajar, 2026-02-02). Lead em tabela 2026.

**Insights:**
- **Migração VNDA → Wake** com necessidade de melhorar recomendação — pitch de feature mix & match / Biso.
- Alto ticket + baixo tráfego = perfil **premium boutique**.

---

### CASE-093 — Cittius

| Campo | Valor |
|---|---|
| **Cliente** | Cittius (`cittius.com.br`) |
| **Setor** | brinquedos colecionáveis / blocos de montar licenciados |
| **Status** | ⚠ em negociação ativa (Wake UP B2B — R$ 4.900 MRR proposto, _quente_) |
| **Plataforma anterior** | nenhuma (lançamento em 03/2026 na ABRIN) |
| **Parceiros envolvidos** | 3W, Omni360, Affys, FZ Commerce (carrossel) |
| **Data da última menção** | 2026-06-02 |

**Situação:** Lead SDR distribuído em 2026-02-09: *"Eles estão estruturando uma operação de venda de produtos licenciados de blocos de montar colecionáveis... lançar no início de março, na ABRIN... Chegou até nós através de indicação de uma pessoa que trabalha na Olist"* (Wake <> 3W, 2026-02-09). Em 2026-05-25/26, negociação ativa com desconto: *"Cittius, pediram para tentarmos algum desconto na mensalidade B2B de 4900,00. É possível?"* (WAKE | Time Agência, 2026-05-26). Em 2026-06-02, listado como _quente_: *"Cittius - B2B - R$4.900,00 MRR (pode reduzir um pouco esse MRR, em negociação...) _quente_"*.

**Insights:**
- **Cliente novo B2B (Wake UP) com ticket R$ 4.900 MRR** — referência de precificação Wake UP B2B 2026 (junto com CASE-094 HVSUL).
- **Lead via Olist** (indicação de pessoa do competidor) — canal interessante.

---

### CASE-094 — HV Sul (HVSUL — Grupo Anhanguera)

| Campo | Valor |
|---|---|
| **Cliente** | HV Sul (`hvsul.com.br`) — citado junto com Grupo Anhanguera |
| **Setor** | ⚠ status desconhecido (vínculo Anhanguera mencionado) |
| **Status** | ⚠ em negociação ativa _quente_ (B2B R$ 4.900 MRR, competindo com VTEX que ofertou 6 meses) |
| **Parceiros envolvidos** | Catus, Avocado Tech, Duo (carrossel) |
| **Data da última menção** | 2026-06-02 |

**Situação:** Lead apareceu em 2026-04-14. Em 2026-05-28, negociação acirrada: *"Estou tentando negociar valores na HVSUL, eles receberam ontem uma proposta da VTEX bem agressiva e estou alinhando com o Gabriel de como agregar valor para o Diogo, comparado com Wake"* (Wake <> Avocado Tech, 2026-05-28). Em 2026-06-02, em lista de quentes: *"HVSUL - B2B - R$4.9000 MRR (pode reduzir, em negociação, VTEX deu 6 meses) _quente_"* (WAKE | Time Agência, 2026-06-02). **Conexão com Anhanguera**: *"coloco o head de e-com da anhanguera para elogiar vcs"* (Wake <> Catus, 2026-05-28).

**Insights:**
- **Caso competitivo Wake vs VTEX** ativo (2026-06) — VTEX ofertou **6 meses grátis** como blocker. Battlecard Wake vs VTEX deve referenciar este caso.
- **Wake UP B2B R$ 4.900 MRR** confirmado como tabela 2026 — junto com Cittius.

---

### CASE-095 — Inigual Decor

| Campo | Valor |
|---|---|
| **Cliente** | Inigual Decor (`inigual.com.br`) — contato Rodrigo |
| **Setor** | casa e decoração |
| **Status** | ❌ sem perfil de migração (2026-05) |
| **Plataforma anterior** | Nuvem Shop ✓ verificado |
| **ERP** | Bling |
| **Faturamento online** | R$ 20k site + 90% em marketplaces |
| **Parceiros envolvidos** | Dotkom, FZ Commerce, Eficaz (carrossel) |
| **Data da última menção** | 2026-05-08 |

**Situação:** Lead SDR distribuído em 2026-04-22. Em 2026-05-08, Eficaz conclui: *"Rodrigo Inigual: trocou recentemente de tema na Nuvem e vende R$20k no ecommerce. Sem perfil de migrar para Wake hoje"* (Eficaz <> WAKE, 2026-05-08).

**Insights:**
- **Cliente sem perfil Wake** (R$ 20k/mês site + investimento recente em tema Nuvem) — exemplo de qualificação que filtra bem.
- 90% GMV em marketplaces = pitch alternativo seria Anymarket, não Commerce.

---

### CASE-096 — Fuel (UseFuel)

| Campo | Valor |
|---|---|
| **Cliente** | Fuel (`usefuel.com.br`) |
| **Setor** | óculos de sol / eyewear |
| **Status** | ✓ ativo (Metakosmos 3D em produção; sem Go-Live confirmado em 2025-08) |
| **Parceiros envolvidos** | Metakosmos (visualizador 3D), Serasa, Duo |
| **Data da última menção** | 2026-06-02 |

**Situação:** Em 2025-08-22, CRMBonus lista Fuel como cliente sem Go-Live: *"Fuel - Ainda sem go live"* (Wake + CRMBonus | Parcerias, 2025-08-22). Em 2026-04-27: *"a Rafa, de vendas, vai te procurar para falar sobre https://www.usefuel.com.br/"* (Serasa & Wake, 2026-04-27). Em 2026-05-14, **Metakosmos 3D em produção**: *"Fuel (Visualizador 3D). https://www.usefuel.com.br/produto/oculos-de-sol-storia-152540"* (Wake & Metakosmos, 2026-05-14).

**Solução:** Wake Commerce + **Metakosmos visualizador 3D** ativo em PDP (URL específica de óculos verificada).

**Insights:**
- Case **Metakosmos em produção** (junto com Gregory — CASE-060 do batch paralelo).
- Vertical **óculos/eyewear** é greenfield no catálogo.

---

### CASE-097 — HDM Beauty (`eclerbeauty.com.br`)

| Campo | Valor |
|---|---|
| **Cliente** | HDM Beauty (`eclerbeauty.com.br` é o site da marca luxo do grupo) |
| **Setor** | beleza / cosméticos (B2C + B2B luxo) |
| **Status** | ⚠ em negociação |
| **Plataforma anterior** | Shopify ✓ verificado (origem Tray → Shopify → migrar para Wake) |
| **Pagamento atual** | Mercado Pago (Shopify) — com bloqueio recorrente |
| **Parceiros envolvidos** | SAVVI, DRIVEN, O2 Consultoria (carrossel) |
| **Data da última menção** | 2026-04-29 |

**Situação:** Lead SDR distribuído em 2026-04-29: *"HDM BEAUTY - CANAIS... conheceu através do Gemini, recomendaram a Wake. Contou que estava relatando pro Gemini os problemas na Shopify, e ele recomendou a ADOBE Commerce e Wake Commerce... iniciaram na Tray... migrar para a Shopify... número de acessos é maior, mas conversão menor. Gostava da Vindi... Shopify usa Mercado Pago e tem bloqueado pagamento de clientes, aciona fraude mesmo sendo cliente real. Tem duas frentes de negócio, uma que trabalha com várias marcas e outra que é mais luxo. A loja luxo já estava na Shopify. Essa segunda marca vai ser somente B2B... (https://eclerbeauty.com.br/) > estão negociando com fabricas para passar a vender para revendedores"* (SAVVI & Wake, 2026-04-29).

**Insights:**
- **Cliente chegou via Gemini IA** — primeira evidência no corpus de lead atribuído a recomendação de assistente IA (Gemini citou Wake junto com Adobe Commerce). Sinal de presença Wake em LLMs.
- **Migração Tray → Shopify → Wake** — segunda migração em 2 anos. Dor: conversão baixa pós-Shopify + bloqueio Mercado Pago.
- **Duas frentes**: B2C multi-marca + B2B luxo (eclerbeauty) — pitch Wake B2B + B2C numa mesma instância.

---

### CASE-098 — Pisom & Co (Pisomeco)

| Campo | Valor |
|---|---|
| **Cliente** | Pisom & Co (`pisomeco.com.br`) |
| **Setor** | construção / pisos |
| **Status** | ⚠ em negociação (sem retorno em 2026-05) |
| **ERP** | recém-migrado a ERP novo na loja física |
| **Parceiros envolvidos** | SAVVI, DRIVEN, Quick Digital (carrossel) |
| **Data da última menção** | 2026-05-14 |

**Situação:** Lead SDR distribuído em 2026-04-29: *"Está em uma pesquisa de unificação de estoque da loja física com o e-commerce. Estoques todos separados, até em lojas separadas e está inviável manter dois estoques. Tem visto cada vez o OMS e a unificação. Acabaram de migrar para um ERP na loja f[ísica]"* (SAVVI & Wake, 2026-04-29). Em 2026-05-12/14: *"Tentei contato com a Pisom algumas vezes mas continuo sem retorno"*.

**Insights:**
- Caso clássico de **unificação OMS multi-loja física** — pitch Wake OMS.
- Recém-migrou ERP físico = timing ruim para nova migração ⚠.

---

### CASE-099 — Nano4you Brasil

| Campo | Valor |
|---|---|
| **Cliente** | Nano4you Brasil (`nano4you.com.br`) |
| **Setor** | ? (Curitiba — chegou após evento Wake 2026-05-21) |
| **Status** | ⚠ em negociação (decisão depende de diretoria) |
| **Plataforma anterior** | Tray (desde março/2025) |
| **Parceiros envolvidos** | Nairuz, N1, Dotkom (carrossel) |
| **Data da última menção** | 2026-05-28 |

**Situação:** Lead SDR distribuído em 2026-05-28: *"Empresa: Nano4you Brasil. Plataforma Atual: Tray (desde março/2025). Origem: Contato iniciado após o evento do dia 21 em Curitiba. Temperatura do Cliente: Morna para Quente. Inicialmente, o time demonstrou uma postura cautelosa, mencionando que a migração não era '100% prioridade'. No entanto, conforme a tecnologia da Wake foi apresentada — principalmente as soluções de múltiplos CDs e regionalização —, o interesse aumentou drasticamente"* (Wake + N1, 2026-05-28). **Decisão depende de diretoria**.

**Insights:**
- **Lead vindo de evento Wake (Curitiba 2026-05-21)** — ROI de eventos comprovado.
- Features-gatilho: **múltiplos CDs + regionalização** — diferencial competitivo Wake confirmado em discoveries.
- Time interno técnico empolgado mas **decisor é diretoria** — agência precisa preparar pitch custo×valor robusto.

---

### CASE-100 — Unipeças MA

| Campo | Valor |
|---|---|
| **Cliente** | Unipeças (`unipecasma.com.br`) — Imperatriz/MA |
| **Setor** | autopeças agrícolas (peças e implementos para tratores) |
| **Status** | ⚠ em negociação |
| **Plataforma anterior** | nenhuma (>30 anos só físico + televendas) |
| **Faturamento** | R$ 10M–30M anuais (Neoway) |
| **Marcas distribuídas** | New Holland, Bosch, Tramontina |
| **Parceiros envolvidos** | Visie, SAVVI, DRIVEN (carrossel) |
| **Data da última menção** | 2026-06-03 |

**Situação:** Lead SDR distribuído em 2026-05-29: *"Empresa: Unipeças. Localização: Imperatriz - Maranhão. Tempo de Mercado: Mais de 30 anos de atuação exclusiva no varejo físico e televendas. Segmento: Venda de peças e implementos agrícolas, com foco em tratores. Marcas Distribuídas/Revendidas: New Holland, Bosch, Tramontina. Faturamento Estimado (Neoway): De R$ 10.000.000,01 a R$ 30.000.000,00 anuais. A Unipeças está iniciando sua jornada digital. O principal gatilho para o projeto é a mudança no comportamento do mercado: os fornecedores e fábricas começaram a vender diretamente ao consumidor final. Objetivo: Criação de um e-commerce focado inicialmente em B2C... mas com capacidade para também atender ao público B2B"* (Visie & Wake, 2026-05-29). Em 2026-06-03, sem retorno: *"Chamei a Aline da Unipeças para uma reunião para apresentarmos a Visie, mas não tive retorno"*.

**Insights:**
- Cenário **fabricante D2C disruption → revendedor precisa entrar no digital** — argumento recorrente para B2B autopeças.
- **Greenfield digital (sem plataforma anterior)** + perfil B2C transicional para B2B = combo Wake Commerce + Wake B2B futuro.

---

### CASE-101 — Gold Prime Joalheria

| Campo | Valor |
|---|---|
| **Cliente** | Gold Prime Joalheria (`goldprimejoias.com.br`) |
| **Setor** | joalheria (Itaim Bibi/SP) |
| **Status** | ❌ perdido por falta de retorno |
| **Plataforma anterior** | Tray ✓ verificado |
| **Parceiros envolvidos** | Ge Commerce, SAVVI, DRIVEN (carrossel) |
| **Data da última menção** | 2026-03-30 |

**Situação:** Lead SDR distribuído em 2026-02-13: *"Plataforma atual: Tray. Estão em busca de um e-commerce estilo Orit (https://www.orit.com.br/quero-comprar). Hoje trabalham com compra e venda de joias e estão no Itaim Bibi. A loja online hoje não vende, e os clientes chegam até eles por indicação"* (DRIVEN & Wake, 2026-02-13). Em 2026-03-30: *"Gold Prime Joalheria - perdido por falta de retorno"* (SAVVI & Wake, 2026-03-30).

**Insights:**
- Lead com **Orit como referência de site** — Orit é case ativo Wake usado como benchmark visual em joias (apesar de não ter case dedicado no catálogo, é citado como referência viva).
- Padrão de **perda por falta de retorno em 6 semanas** — risco SMB típico.

---

### CASE-102 — Tag de leads consolidados FZ Commerce (rotação carteira 2026-01)

| Campo | Valor |
|---|---|
| **Origem** | Mensagem consolidada em FZ COMMERCE & Wake, 2026-01-22 |
| **Status** | ⚠ multi-prospect — todos em rotação |
| **Data** | 2026-01-22 |

**Situação:** Em 2026-01-22, FZ Commerce compartilhou lista consolidada de oportunidades em rotação na carteira: *"Uniland - Commerce - Canais / 7 Oliveiras - Commerce - Canais / Surfs UP Club - Commerce - Canais / Ekilibre Amazônia - Commerce - Canais / ABRAFER - Commerce - Canais / Conceito Rouparia - Commerce - Canais / Believe Professional - Commerce - Canais / Tazin Cosméticos - Commerce - Canais / GoPure - Commerce - Canais / Koube - Commerce - Canais / Vitalbio - Commerce - Canais / Regina Salomão - Commerce - Canais / Loja do Detalhamento - Commerce - Canais / Alfredo Meyer - Commerce - Canais / Ivan Joias- Canais / Sustent Intermediação de Negócios- Canais"* (FZ COMMERCE & Wake, 2026-01-22).

**Casos detalhados neste catálogo:**
- 7 Oliveiras → CASE-081
- Conceito Rouparia → CASE-082
- Believe Professional + Tazin → CASE-084
- Regina Salomão → CASE-085
- Loja do Detalhamento + SGCB + IGL → CASE-086
- Alfredo Meyer → CASE-087
- Ivan Joias → CASE-089
- Sustent → CASE-090

**Casos médio-confiança (não escritos):**
- Uniland, Ekilibre Amazônia, ABRAFER, GoPure, Koube, Vitalbio, Surfs UP Club — tracked como watchlist.

**Insights:**
- Lista FZ Commerce 2026-01-22 é a **fotografia mais consolidada** do pipeline canais SMB no início de 2026.
- Muitos leads compartilhados em **4+ agências em paralelo** — sinal de saturação do modelo carrossel.

---

### Verificação — Segunda leva CASE-073 a CASE-102

| Claim | Status | Query (LIKE) | Evidence |
|---|---|---|---|
| Cerealista Express é cliente Magento querendo migrar; rotação em 3 agências | ✓ verificado | `%CEREALISTA EXPRESS%`; `%cerealistaexpress%` | Wake <> Avocado Tech 2025-11-27, 2025-12-04; O2 Consultoria & Wake 2026-01-26 |
| Eletrônica Santana é cliente Wake confirmado (Grupo com ES TECH + Dealer Shop B2B) via Ecommerce House + Biso | ✓ verificado | `%Eletrônica Santana%nosso cliente%`; `%Grupo Eletrônica Santana%` | Wake & Biso 2025-09-16; Ecommerce House 2026-02-03 e 2026-04-08 |
| Inoxlon: Climba → Wake, ERP Bling, GMV R$ 200k, lead esfriou por contato inexistente | ✓ verificado | `%INOXLON%Canais%`; `%Inoxlon%inoxlon@inoxlon%` | Wake & A&EIGHT 2026-02-12; Wake <> Engajar 2026-02-19 |
| AZ Acessórios entrou no ar em 2025-06-23 pela Eficaz; case do Grupo Cambuci | ✓ verificado | `%azacessorios%Cambuci%`; `%AZ acessórios%AR%` | Eficaz <> WAKE 2025-06-23; 2026-03-04 (novo storefront); 2026-03-16 (referência em discovery) |
| Dallanese Parafusos é lead frio (ERP Sankhya, GMV R$ 50k) sem retorno após 4 semanas | ✓ verificado | `%Dallanese%Sankhya%`; `%Dallanese%retorno%` | Duo e Wake Parcerias 2025-06-25; A&EIGHT 2025-07-15; Duo 2025-07-21 |
| Balletto é migração VTEX (ERP Sisplan) em proposta 07/2025 via O2 | ✓ verificado | `%Balletto%VTEX%`; `%Balletto%Sisplan%` | Wake <> 3W 2025-07-11; O2 Consultoria & Wake 2025-07-22 |
| AR15 Sport é migração Tray (airsoft) com tomador difícil de engajar | ✓ verificado | `%AR15 Sport%Tray%`; `%AR15%airsoft%` | Wake <> E-Plus 2025-07-22; WAKE \| Time Agência 2025-09-19 a 2025-09-29 |
| Campinas Diesel é lead SDR via Msam distribuído a Catus + Eficaz + Proex em 2025-08-12 | ✓ verificado | `%Campinas Diesel%Msam%`; `%Campinas Diesel%mercado livre%` | Eficaz <> WAKE 2025-08-12; Wake <> Catus 2025-08-15 |
| 7 Oliveiras: Visual e-commerce → Wake, ERP ZADA, 80% GMV em marketplaces | ✓ verificado | `%7oliveiras%ZADA%`; `%7 Oliveiras%marketplaces%` | Simples Inovação & Wake Parcerias 2025-08-18 |
| Conceito Rouparia: Loja Integrada → Wake, manter ERP Datasystem via 4middleware (API validado por Wake Parceria Tech) | ✓ verificado | `%conceitorouparia%Datasystem%`; `%4middleware%API%` | Simples Inovação & Wake Parcerias 2025-09-03 |
| Carinion Mix perdido para VNDA por preço (R$ 600/mês teto cliente) | ✓ verificado | `%Carinion%VNDA%`; `%carinonmix%HostGator%` | Eficaz <> WAKE 2025-09-15 e 2025-09-19 |
| Believe Professional: Loja Integrada → Wake, ERP Bling | ✓ verificado | `%believepro%Loja Integrada%` | Simples Inovação & Wake Parcerias 2025-09-22 |
| Tazin Cosméticos: Nuvem Shop → Wake, paga R$ 7.500/mês para agência atual | ✓ verificado | `%tazincosmeticos%Nuvem Shop%`; `%Tazin%7.500%` | Wake <> Lemoon 2025-09-24 |
| Regina Salomão: Tray → Wake, ERP MV Tecnologia, GMV R$ 150-200k, projeto adiado para 2026 | ✓ verificado | `%Regina Salomão%Tray%`; `%Regina Salomão%mês 10%` | Wake <> Avocado Tech 2025-10-02; Simples Inovação & Wake Parcerias 2025-11-27 |
| Loja do Detalhamento + SGCB + IGL Coatings: grupo de 3 lojas detailing automotivo, Tray → Wake com OMS unificado, ERP Arth Magnus precisa integrador | ✓ verificado | `%Loja do Detalhamento%`; `%SGCB%IGL%`; `%Arth Magnus%` | Simples Inovação & Wake Parcerias 2025-10-29 e 2025-10-31 |
| Alfredo Meyer: WordPress → Wake, requer RD Station + WhatsApp + ERP + marketplaces | ✓ verificado | `%alfredomeyer%Wordpress%`; `%Alfredo Meyer%RD Station%` | Cava Digital & Wake Parcerias 2025-11-25; FZ Commerce 2025-12-05 |
| Festa Color foi desqualificado em 2025-12-05; citou Shop2gether + Puravida como referência | ✓ verificado | `%Festa Color%Desqualificado%`; `%Festa Color%shop2gether%` | WAKE \| Time Agência 2025-12-05; Eficaz <> WAKE 2025-12-01 |
| Ivan Joias: Tray (mesma LWSA) → Wake, ERP 4Control, GMV R$ 150k, 6 lojas Salvador, decisão por benchmark Shoulder (cashback + agenda vendedor) | ✓ verificado | `%Ivan Joias%4Control%`; `%Ivan Joias%Sholuder%` | FZ Commerce 2025-12-01; Cava Digital 2026-01-19 |
| Sustent quer ser marketplace B2B2C próprio em beleza (ERP Omie+Bling, equipe interna iPaaS) | ✓ verificado | `%sustent.com.br%marketplace%`; `%Sustent%B2B2C%` | Wake+Omni360+Affys 2025-12-16 |
| Shop2gether (Grupo Icomm) fechou em 2025-07-30, caso público publicado em 2025-08-18, migração DRIVEN→Oásis | ✓ verificado | `%shop2gether%fech%`; `%icomm-shop2gether%cases%` | Mavericks 2025-07-30; FZ Commerce 2025-08-18; WAKE \| Time Agência 2025-08-14 |
| Shop2gether ainda sem Go-Live em 2025-08-22 segundo CRMBonus | ✓ verificado | `%Icomm%Shop2gether%go live%` | Wake + CRMBonus \| Parcerias 2025-08-22 |
| Zinnia Leather: VNDA → Wake, acessórios couro alto ticket, 1.4K visualizações/mês | ✓ verificado | `%Zinnia%VNDA%`; `%zinnialeather%` | Wake <> Engajar 2026-02-02; Wake & A&EIGHT 2026-02-02 |
| Cittius é prospect novo B2B (Wake UP) por R$ 4.900 MRR em negociação ativa, lead via Olist | ✓ verificado | `%Cittius%B2B%4.900%`; `%Cittius%Olist%` | Wake <> 3W 2026-02-09; WAKE \| Time Agência 2026-05-26 e 2026-06-02 |
| HVSUL é prospect B2B R$ 4.900 MRR _quente_ competindo com VTEX (que ofertou 6 meses); conexão Grupo Anhanguera | ✓ verificado | `%HVSUL%VTEX%`; `%HVSUL%4.9%`; `%anhanguera%head%e-com%` | WAKE \| Time Agência 2026-06-02; Wake <> Catus 2026-05-28 e 2026-06-02 |
| Inigual Decor (Nuvem Shop, R$ 20k/mês, 90% em mkp) desqualificado em 2026-05-08 | ✓ verificado | `%Rodrigo Inigual%`; `%inigual.com.br%` | FZ Commerce 2026-04-22; Eficaz <> WAKE 2026-05-08 |
| Fuel (UseFuel) tem Metakosmos visualizador 3D em produção em PDP (URL específica verificada) | ✓ verificado | `%usefuel%Metakosmos%`; `%Fuel%Visualizador 3D%` | Wake & Metakosmos 2026-05-14; Wake + CRMBonus 2025-08-22 |
| HDM Beauty (eclerbeauty.com.br): chegou via Gemini IA, Tray → Shopify → Wake, B2C multi-marca + B2B luxo | ✓ verificado | `%HDM BEAUTY%Gemini%`; `%eclerbeauty%B2B%` | SAVVI & Wake 2026-04-29; DRIVEN & Wake 2026-04-29 |
| Pisom & Co (pisomeco.com.br) é lead OMS multi-loja física construção/pisos sem retorno | ✓ verificado | `%Pisom%OMS%`; `%pisomeco%` | SAVVI & Wake 2026-04-29; Quick Digital 2026-05-12; SAVVI 2026-05-14 |
| Nano4you Brasil é Tray → Wake, vindo de evento Wake Curitiba 2026-05-21, gatilhos múltiplos CDs + regionalização, decisor é diretoria | ✓ verificado | `%Nano4you%Curitiba%`; `%Nano4you%múltiplos CDs%` | Wake + N1 / Wake <> Nairuz / Parcerias Dotkom 2026-05-28 |
| Unipeças MA é autopeças agrícolas em Imperatriz/MA, greenfield digital, faturamento R$ 10-30M anuais (Neoway) | ✓ verificado | `%Unipeças%Imperatriz%`; `%unipecasma%New Holland%` | Visie & Wake 2026-05-29 e 2026-06-03 |
| Gold Prime Joalheria (Itaim Bibi, Tray → Wake): perdido por falta de retorno em 2026-03-30; usou Orit como benchmark | ✓ verificado | `%Gold Prime%perdido%`; `%Gold Prime%Orit%` | DRIVEN & Wake 2026-02-13; SAVVI & Wake 2026-03-30 |
| Lista consolidada FZ Commerce 2026-01-22 lista 16 oportunidades em rotação carteira | ✓ verificado | `%Uniland%7 Oliveiras%Sustent%` | FZ COMMERCE & Wake 2026-01-22 19:23 |
| Skip: Orit é referência visual de joias citada como benchmark, **não cliente Wake direta com evidência suficiente** | ❌ sem fonte direta | `%orit.com.br%` (4 hits, todas como referência em descrição de outros leads) | Comercial Cross WAKE 2025-09-17; DRIVEN/Ge Commerce/SAVVI 2026-02-13 |
| Skip: SoulPink é projeto interno/founder (grupo "SoulPink - DONOS"), **não cliente Wake** | ❌ não-aplicável | `%SoulPink%`; `%eusoulpink%` | SoulPink - DONOS / SoulPink Geral 2025-05 a 2026-06 (decisão de domínio, planejamento) |
| Conexão Grupo Cambuci ↔ outros clientes Wake (ex: Penalty/Olympikus) | ⚠ inferido | n/a — Cambuci nomeado mas vínculo não confirmado em corpus | Eficaz <> WAKE 2025-06-23 (única menção Cambuci) |

---

## Verificação

| Claim | Status | Query (LIKE) | Evidence |
|---|---|---|---|
| Aramis assinou Commerce+OMS por R$ 154k MRR + R$ 300k setup | ✓ verificado | `%Aramis%MRR%` + `%assinad%` | Comercial Cross WAKE, 2026-06-02 |
| Karcher é cliente Tiny + Wake + HERO Seguros (Garantia Estendida + Parcela Garantida) | ✓ verificado | `%Karcher%` (127 hits); `%Karcher%HERO%`; `%Karcher%Tiny%` | Projeto Hero <> Leanwork <> Wake, múltiplas datas 2026-03 a 2026-06 |
| Inbrands é enterprise com problema marketplace impactando 10% do GMV (~R$ 1M) | ✓ verificado | `%Inbrands%GMV%`; `%Inbrands%enterprise%` | ANYMARKET & WAKE, 2026-04-16/17 |
| Shoulder tem case público de migração mediado pelo Pavoni | ✓ verificado | `%Shoulder%case%`; `%Shoulder%migração%` | WAKE \| Time Agência, 2026-05-15; Ecommerce House, 2026-04-28 |
| Oriba migrou de Shopify e teve Go-Live em 2026-01-13 | ✓ verificado | `%Oriba%Shopify%`; `%Oriba%Go Live%` | Comercial Cross WAKE 2026-03-12; Wake & Biso 2026-01-29 |
| LP Beauty deu churn e fez rollback para Bagy | ✓ verificado | `%LP Beauty%churn%`; `%LP Beauty%Bagy%` | Wake & Biso, 2026-04-29 |
| Ricardo Almeida assinou e pediu churn em 3 meses | ✓ verificado | `%Ricardo Almeida%churn%`; `%Ricardo Almeida%assinad%` | Mavericks 2025-08-01 (assinatura) + 2025-10-29 (churn) |
| Ybera tem POC com VTEX em jun/2026 | ✓ verificado | `%Ybera%`; `%Ybera%VTEX%` | The Rise of Skywalker, 2026-06-01/03 |
| Casa Almeida participou de painel Wake On the Road (Curitiba) | ✓ verificado | `%Casa Almeida%`; `%Wake On the Road%` | Wake & BonifiQ, 2026-05-11 |
| Soneda usa Wake + Floui (ERP) + Anymarket (TikTok Shop) | ✓ verificado | `%Soneda%Floui%`; `%Soneda%tiktok%` | Floui & Wake, 2026-05-28 |
| Empório do Celular é cliente Tray + Tiny querendo migrar para Wake | ✓ verificado | `%Empório do Celular%` | O2 Consultoria & Wake, 2026-04-09 |
| Mpozenato deu churn de R$ 1,8M, contrato fixo por pedido | ✓ verificado | `%Mpozenato%` | The Rise of Skywalker, 2026-05-27 |
| Cacay assinou e migrou para Dafiti via Anymarket com erro de frete | ✓ verificado | `%Cacay%`; `%Cacay%Dafiti%` | Comercial Cross WAKE, 2026-05-25; ANYMARKET & WAKE 2026-05-19 |
| Wallford Brasil é Tiny + R$ 150k/mês, projeto multi-marca | ✓ verificado | `%Wallford%` (9 hits) | Wake <> Avocado Tech, 2025-12-22 |
| DLK migrou para Shopify mas mantém Vindi | ✓ verificado | `%DLK%Shopify%` | The Rise of Skywalker, 2026-05-25 |
| Pavoni no painel pode ser executivo Wake e não cliente — não confirmado | ⚠ inferido | `%Pavoni%` (predomina como sobrenome do executivo) | WAKE & BonifiQ 2026-05-11 — citado como speaker junto com Balaroti |
| Riachuelo é cliente Wake | ❌ sem fonte (Riachuelo é canal marketplace de Blueman, não cliente Wake direta; aliás, descontinuou marketplace em 30/01/2026) | `%Riachuelo%` | Wake + Marketplace Blueman, 2026-01-28 |
| MP Brinquedos é cliente Wake | ❌ sem fonte (única menção: lista de convidados de evento) | `%MP Brinq%` | Frenet + Wake, 2026-04-01 — só convite de evento |
| Movimento é cliente Wake | ❌ sem fonte (termo predomina em uso genérico de "movimento") | `%Movimento%` (295 hits, ruído) | n/a |
| Estoca é cliente Wake | ❌ sem fonte (Estoca é parceiro de fulfillment, listado em playbook §2 como `parceiro-logistica`) | `%Estoca%` | múltiplas conversas em grupo "Estoca & Wake - Vendas" como parceiro |
| Selloutfarma é cliente Wake | ❌ sem fonte (0 hits no DB) | `%Selloutfarma%`, `%Selloutpharma%` | nenhuma menção |
| Veste S.A. (Dudalina) é cliente Wake | ⚠ inferido (cliente Grupo Veste mencionado em negociação 2025-06-30; Dudalina especificamente não em negociação recente) | `%Dudalina%`; `%Veste S%` | Comercial Cross WAKE 2025-06-30 |
| Polishop é cliente Wake | ❌ sem fonte (4 hits, todos contextos externos — podcast, lista de convidados) | `%Polishop%` | n/a |
| Pampulha Enxovais é cliente Wake | ⚠ inferido (lead em fluxo de canais, ERP Softcom, R$ 150k/mês, status final não confirmado) | `%Pampulha%` | Parcerias Dotkom & Wake, 2026-01-29 |
| Movimento (mencionado como cliente Wake) | ❌ sem fonte distinguível do ruído |  | n/a |

---

## Fontes (grupos consultados)

- WAKE \| Time Parcerias
- WAKE \| Time Agência
- WAKE \| Coordenação
- WAKE \| Time Tech
- Comercial Cross WAKE
- ANYMARKET & WAKE (TÁTICO / OPERACIONAL)
- Wake & Biso
- Wake & Omnichat
- Wake & BonifiQ
- Wake & Leanwork
- Wake & Octadesk
- Wake & A&EIGHT (B8One)
- Wake & Mundo Verde
- Wake <> Catus
- Wake <> Nairuz
- Wake <> Avocado Tech
- Wake <> Oásis
- Wake <> Série//A
- Wake <> 3W
- Wake + Marketplace Blueman
- Wake + N1
- Wake + CRMBonus | Parcerias
- Wake + Appmax
- WAKE & PAYPAL
- Mercado Pago + Wake
- Projeto Hero <> Leanwork <> Wake
- Projeto Mundo Verde - INTERNO
- Projeto Serasa Pass <> Leanwork <> Wake
- Sellbie & Wake
- Floui & Wake
- Eficaz <> WAKE
- Wicomm <> Wake
- Agencia Cava Digital & Wake Parcerias
- Quick Digital & Wake Parcerias
- Duo e Wake Parcerias
- DRIVEN & Wake
- FZ COMMERCE & Wake
- Parcerias Dotkom & Wake
- O2 Consultoria & Wake
- SAVVI & Wake
- Eitri & Wake
- Vindi + Wake
- Uble & Wake
- Konfidency & Wake
- Konduto & Wake
- Frenet + Wake
- Wake+Omni360+Affys
- Tropa de Elite 3
- The Rise of Skywalker
- Mavericks
- Shoulder + Figo + Wake
- Figo & Wake
- Coder Ivy + Wake, partnership success
- Four2one & Wake
- Serasa & Wake
- Topsort + Wake
- Abbiamo & Wake
- EBI & Wake (DEV)
- Visie & Wake
- Wake <> Benova (Editora Fórum)
- Wake & Stone — [work] (caso Destro)
- Wake + CRMBonus | Parcerias
- Monitfy & Wake (afiliados / API)
- Uble & Wake (StoreFront API + afiliados)
- CoreBiz & Wake (benchmark B2B Buddemeyer + CVC)
- Wake & Omnichat (checklist Buddemeyer)
- Anymarket + Wake (comercial)

---

## Queries SQL utilizadas

```sql
-- 1. Mensagens com 'Nova Oportunidade Wake' (descoberta de prospects)
SELECT c.name, datetime(m.timestamp), substr(m.content,1,400)
FROM messages m JOIN chats c ON c.jid=m.chat_jid
WHERE c.jid LIKE '%@g.us'
  AND (m.media_type IS NULL OR m.media_type='')
  AND m.content LIKE '%Nova Oportunidade Wake%'
ORDER BY m.timestamp DESC;

-- 2. Closings com MRR
SELECT c.name, datetime(m.timestamp), substr(m.content,1,400)
FROM messages m JOIN chats c ON c.jid=m.chat_jid
WHERE c.jid LIKE '%@g.us'
  AND (m.media_type IS NULL OR m.media_type='')
  AND m.content LIKE '%assinad%'
  AND (m.content LIKE '%MRR%' OR m.content LIKE '%mrr%');

-- 3. Churn
SELECT c.name, datetime(m.timestamp), substr(m.content,1,500)
FROM messages m JOIN chats c ON c.jid=m.chat_jid
WHERE c.jid LIKE '%@g.us'
  AND (m.media_type IS NULL OR m.media_type='')
  AND (m.content LIKE '%churn%' OR m.content LIKE '%saiu da Wake%');

-- 4. Go-Live lists
SELECT c.name, datetime(m.timestamp), substr(m.content,1,800)
FROM messages m JOIN chats c ON c.jid=m.chat_jid
WHERE c.jid LIKE '%@g.us'
  AND (m.media_type IS NULL OR m.media_type='')
  AND m.content LIKE '%Go Live%';

-- 5. Por cliente (substitua TERMO)
SELECT c.name, datetime(m.timestamp), substr(m.content,1,500)
FROM messages m JOIN chats c ON c.jid=m.chat_jid
WHERE c.jid LIKE '%@g.us'
  AND (m.media_type IS NULL OR m.media_type='')
  AND m.content LIKE '%TERMO%'
ORDER BY m.timestamp DESC;
```

Total de queries executadas para construir este catálogo: **~28** combinações distintas de termos, mais ~12 buscas de validação cruzada. **Extensão B2B/Afiliados/Editoras (CASE-051–059)**: 11 queries adicionais (Esplane, Você Constrói, CVC franquias/matriz/B2B, Yora, Destro + Wake/Stone, Buddemeyer, Gold Food, A Fabrica + checkout, Unitoys, Editora Fórum, afiliados + Ybera/Shoulder/BAW, afiliados gerais).

---

## Tags

`audience-sales`, `audience-solutions`, `audience-cs`, `proposta`, `customer-case`
