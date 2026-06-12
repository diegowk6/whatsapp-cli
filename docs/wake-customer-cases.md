# Wake Customer Cases — Catálogo de Clientes e Casos de Uso

**Artefato**: catálogo de customer cases extraído do corpus WhatsApp Wake/parceiros (`store/messages.db`).
**Data**: 2026-06
**Fonte**: mensagens de grupo (`chat_jid LIKE '%@g.us'`, texto puro), recortes 2025-06 a 2026-06.
**Total de casos**: 59

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
