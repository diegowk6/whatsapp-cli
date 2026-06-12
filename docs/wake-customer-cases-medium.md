# Wake Customer Cases — Medium-confidence watchlist

**Purpose**: Customers with 2 hits in the corpus (or captured via structured `Cliente:`/`Lojista:` fields). Signal too thin for a full `CASE-NNN` writeup in `wake-customer-cases.md` but tracked for promotion once a third hit accumulates.

**Source**: `docs/wake-customer-cases-gap-list.md` — medium-confidence (2-hit) section + structured-field extras section.

**Discovery**: 2026-06. All entries verified against `store/messages.db` group text messages (`chat_jid LIKE '%@g.us'` and `(media_type IS NULL OR media_type='')`).

---

## Medium-confidence (2-hit candidates)

### LEAD-001 — Miss Victoria

| Campo | Valor |
|---|---|
| **Cliente** | Miss Victoria |
| **Setor** | moda ⚠ inferido |
| **Status** | ⚠ desconhecido — repassado para FG como CS já existente buscando agência de sustentação CRM |
| **URL** | `missvicttoria.com.br` |
| **Hits no DB** | 2 |
| **Última menção** | 2025-10-24 |
| **Grupo principal** | FG + Wake |
| **Tags** | `audience-sales`, `proposta` |

**Contexto:** Em 2025-07-07 (FG + Wake), CS Wake indicou Miss Victoria para a agência FG porque o cliente contratou CRM RD Station mas não consegue gerir, precisando de sustentação. URL reaparece em Comercial Cross WAKE em 2025-10-24 sem novo contexto.

**Por que está aqui (não em wake-customer-cases.md):** Sinal limitado a indicação de sustentação CRM; sem confirmação de fechamento ou go-live. Promover quando aparecer status de aceite pela FG ou ativação de projeto.

---

### LEAD-002 — Curavino

| Campo | Valor |
|---|---|
| **Cliente** | Curavino |
| **Setor** | vinhos / e-commerce de vinhos ⚠ inferido |
| **Status** | em qualificação — citado em listas multi-agência |
| **URL** | `curavino.com.br` |
| **Hits no DB** | 2 (+ menções em Break the Pattern e Estratégia Wake na AWS, fora dos grupos de parceria) |
| **Última menção** | 2026-02-27 |
| **Grupo principal** | Agencia Virtualiti | Wake |
| **Tags** | `audience-sales`, `lead-pipeline` |

**Contexto:** Aparece em lista compartilhada de prospects pela Virtualiti em 2026-02-27. Em Break the Pattern (grupo fora do escopo principal de parcerias), em 2026-01-28: "Curavino já paga a conta" e "Vou colocar na Curavino e tem mais um cliente pedindo". Há também citação em Estratégia Wake na AWS (2025-12-22) sobre email da Curavino.

**Por que está aqui (não em wake-customer-cases.md):** Sinal cruza vários grupos mas sem contexto operacional concreto sobre fase de venda ou go-live. Promover quando aparecer registro em grupo de parceria com agência atribuída.

---

### LEAD-003 — Casa Simões

| Campo | Valor |
|---|---|
| **Cliente** | Casa Simões |
| **Setor** | varejo ⚠ inferido |
| **Status** | em qualificação — lead aberto pela Catus |
| **URL** | `casasimoes.com.br` |
| **Hits no DB** | 2 (Wake <> Catus) + citação em WAKE | Time Agência |
| **Última menção** | 2025-08-15 |
| **Grupo principal** | Wake <> Catus |
| **Tags** | `audience-pam`, `lead-pipeline` |

**Contexto:** Em 2025-08-13 Catus enviou URL no grupo Wake <> Catus para registrar lead. Em 2025-08-15 (WAKE | Time Agência), surge no resumo semanal como "Casa Simões - Catus - Time Canais" e é citada como exemplo de lead que abre pelo Community mas não contabiliza no dash de QuickSight — issue operacional de atribuição.

**Por que está aqui (não em wake-customer-cases.md):** Sem evolução posterior nem proposta/closing reportados. Bug de atribuição no dash é o sinal mais concreto, e ele é sobre processo Salesforce, não sobre o cliente.

---

### LEAD-004 — Uniland

| Campo | Valor |
|---|---|
| **Cliente** | Uniland |
| **Setor** | ? (marketplace de manufatura) |
| **Status** | em qualificação — distribuído pelo Canais |
| **URL** | `uniland.com.br` |
| **Hits no DB** | 2 |
| **Última menção** | 2026-01-22 |
| **Grupo principal** | Wake <> 3W |
| **Tags** | `audience-sales`, `lead-pipeline` |

**Contexto:** Em 2025-08-15 mesmo lead enviado em Simples Inovação e Wake <> 3W como "Uniland - Commerce - Canais" — descrito como "marketplace para m..." (truncado). Em 2026-01-22 reaparece em listagem da FZ Commerce de oportunidades de Canais ativas.

**Por que está aqui (não em wake-customer-cases.md):** Lead distribuído por canais; sem fechamento ou perda confirmada. Promover quando agência sinalizar próximo passo.

---

### LEAD-005 — Anhanguera Ferramentas

| Campo | Valor |
|---|---|
| **Cliente** | Anhanguera Ferramentas |
| **Setor** | ferramentas / construção |
| **Status** | ⚠ usado como referência competitiva (cliente Wake existente, ainda assim re-aparece em listas) |
| **URL** | `anhangueraferramentas.com.br` |
| **Hits no DB** | 2 |
| **Última menção** | 2026-02-27 |
| **Grupo principal** | SAVVI & Wake |
| **Tags** | `audience-solutions`, `lead-pipeline` |

**Contexto:** Em 2025-09-09 (SAVVI & Wake), citada como "concorrente direto que usa Wake" de um prospect — o que indica que Anhanguera Ferramentas **já é cliente Wake**. Reaparece em 2026-02-27 em listagem da Virtualiti misturada com URLs de prospects.

**Por que está aqui (não em wake-customer-cases.md):** Provável cliente Wake existente sem case já catalogado. Promover após validar status no Salesforce e confirmar se merece CASE-NNN próprio ou só referência em battlecard de ferramentas.

---

### LEAD-006 — Imperial Ferramentas

| Campo | Valor |
|---|---|
| **Cliente** | Imperial Ferramentas |
| **Setor** | ferramentas |
| **Status** | ⚠ desconhecido — citado como exemplo de uso de HERO Seguros |
| **URL** | `imperialferramentas.com.br` |
| **Hits no DB** | 2 |
| **Última menção** | 2025-09-22 |
| **Grupo principal** | HERO & Wake |
| **Tags** | `audience-solutions`, `lead-pipeline` |

**Contexto:** Em 2025-09-12 (HERO & Wake) URL citada como "outro exemplo" no contexto de discussão de HERO Seguros, sugerindo é cliente Wake com HERO integrado. Em 2025-09-22 (WAKE | Time Agência) URL aparece sem contexto adicional.

**Por que está aqui (não em wake-customer-cases.md):** Possível cliente Wake existente com integração HERO ativa, mas sinal puxado de mensagem comparativa. Promover quando aparecer detalhe operacional (ticket, deal, renovação).

---

### LEAD-007 — Voomp Creators

| Campo | Valor |
|---|---|
| **Cliente** | Voomp Creators (Grupo Cogna) |
| **Setor** | educação ⚠ inferido |
| **Status** | em qualificação — repassado pela Ecommerce House |
| **URL** | `voompcreators.com.br` |
| **Hits no DB** | 2 (Ecommerce House) |
| **Última menção** | 2026-01-12 |
| **Grupo principal** | Ecommerce House |
| **Tags** | `audience-sales`, `lead-pipeline` |

**Contexto:** Em 2026-01-12 a Ecommerce House compartilhou URL com tag "Maior conglomerado educacional do Brasil. Grupo Cogna." em 2025-09-24 a mesma URL já havia sido compartilhada sem contexto adicional.

**Por que está aqui (não em wake-customer-cases.md):** Pertence ao Grupo Cogna (potencial conta enterprise) mas nenhum status comercial concreto. Promover se Ecommerce House abrir oportunidade formal.

---

### LEAD-008 — GoPure

| Campo | Valor |
|---|---|
| **Cliente** | GoPure |
| **Setor** | beleza / saúde |
| **Status** | em qualificação — opt de Canais multi-agência |
| **URL** | `gopure.com.br` |
| **Hits no DB** | 2 |
| **Última menção** | 2025-09-29 |
| **Grupo principal** | FZ COMMERCE & Wake |
| **Tags** | `audience-sales`, `lead-pipeline` |

**Contexto:** Em 2025-09-29, oportunidade "GoPure - Commerce - Canais" distribuída em paralelo para Simples Inovação e FZ Commerce, com mesmo template "A agencia deve apresentar Wake e proposta comercial da plataforma". URL `gopure.com.br` com 8.2K de visualização.

**Por que está aqui (não em wake-customer-cases.md):** Distribuição multi-agência típica de Canais, mas sem follow-up posterior nas threads consultadas. Promover quando uma agência reportar avanço.

---

### LEAD-009 — Vital Bio

| Campo | Valor |
|---|---|
| **Cliente** | Vital Bio (Vitalbio) |
| **Setor** | saúde / orgânicos |
| **Status** | em qualificação — opt de Canais multi-agência |
| **URL** | `vitalbio.com.br` |
| **Hits no DB** | 2 |
| **Última menção** | 2025-10-02 |
| **Grupo principal** | FZ COMMERCE & Wake |
| **Tags** | `audience-sales`, `lead-pipeline` |

**Contexto:** Em 2025-10-02, mesma mecânica: "Oportunidade: Vitalbio - Commerce - Canais" distribuída em paralelo para Simples Inovação e FZ Commerce.

**Por que está aqui (não em wake-customer-cases.md):** Mesmo padrão de distribuição de Canais sem follow-up. Promover quando uma agência reportar avanço.

---

### LEAD-010 — Biggie (abiggie)

| Campo | Valor |
|---|---|
| **Cliente** | Biggie |
| **Setor** | ? |
| **Status** | em qualificação — saindo de Shopify, ERP Tiny |
| **URL** | `abiggie.com.br` |
| **Hits no DB** | 2 |
| **Última menção** | 2025-10-09 |
| **Grupo principal** | Wake <> Ge Commerce |
| **Tags** | `audience-solutions`, `lead-pipeline` |

**Contexto:** Em 2025-10-09 mesmo lead distribuído para Ge Commerce e O2 Consultoria: "Biggie - Commerce - Canais — Cliente Shopify. ERP - Tiny" (texto truncado). Lead com sinais técnicos claros: migração de Shopify e ERP Tiny (nativo Wake).

**Por que está aqui (não em wake-customer-cases.md):** Sinal técnico forte mas sem evolução posterior. Promover quando alguma agência confirmar próximo passo. Boa candidata para CASE no battlecard Shopify→Wake.

---

### LEAD-011 — Berzerk

| Campo | Valor |
|---|---|
| **Cliente** | Berzerk |
| **Setor** | moda ⚠ inferido (produto "oversized sereia" citado) |
| **Status** | ⚠ desconhecido — possivelmente já cliente Wake na comunidade Ecommerce House |
| **URL** | `berzerk.com.br` |
| **Hits no DB** | 2 (Ecommerce House) |
| **Última menção** | 2025-10-16 |
| **Grupo principal** | Ecommerce House |
| **Tags** | `audience-sales`, `lead-pipeline` |

**Contexto:** Em 2025-10-16 (Ecommerce House) pergunta "Alguem da loja Berzeck loja aqui??" sugerindo que pode ser membro da comunidade. URL compartilhada para compra direta. Sem mensagem formal de oportunidade.

**Por que está aqui (não em wake-customer-cases.md):** Sinal de comunidade Ecommerce House, não de pipeline formal. Promover quando aparecer oportunidade registrada ou referência operacional.

---

### LEAD-012 — Armazém CEC (Armazém Casa & Construção)

| Campo | Valor |
|---|---|
| **Cliente** | Armazém Casa & Construção |
| **Setor** | construção / casa |
| **Status** | ❌ **perdido** — "escolheu outro parceiro" (2026-01-12) |
| **URL** | `armazemcec.com.br` |
| **Hits no DB** | 2 |
| **Última menção** | 2026-01-12 |
| **Grupo principal** | O2 Consultoria & Wake / DRIVEN & Wake |
| **Tags** | `audience-sales`, `lead-pipeline` |

**Contexto:** Em 2025-11-04 oportunidade "Armazém Casa & Construção - Commerce - Canais" distribuída para O2 e DRIVEN. Em 2026-01-12 (DRIVEN & Wake) status: "Armazemcec - Erik Raphael - escolheu outro parceiro" — explicitamente perdido.

**Por que está aqui (não em wake-customer-cases.md):** Lead perdido. Pode ser promovido como CASE de win-loss/lições aprendidas, mas sem detalhes do "outro parceiro".

---

### LEAD-013 — QualyAir

| Campo | Valor |
|---|---|
| **Cliente** | QualyAir (Grupo Rede HDS) |
| **Setor** | ar condicionado automotivo |
| **Status** | em qualificação — cliente solicitou reunião com Wake |
| **URL** | `qualyair.com.br` |
| **Hits no DB** | 2 (Wake <> Série//A) |
| **Última menção** | 2025-12-02 |
| **Grupo principal** | Wake <> Série//A |
| **Tags** | `audience-sales`, `lead-pipeline` |

**Contexto:** Em 2025-11-12 (Wake <> Série//A) lead detalhado: pertence ao Grupo Rede HDS (25+ anos), B2C há ~8 anos em ar condicionado automotivo, atualmente na plataforma B4X. Em 2025-12-02 cliente solicitou reunião direta com Wake via Henrique da Série//A.

**Por que está aqui (não em wake-customer-cases.md):** Contexto forte (grupo, vertical, plataforma atual) mas só duas mensagens. Promover assim que reunião com Wake for confirmada — boa candidata para case B4X→Wake.

---

### LEAD-014 — Menegatti Couro

| Campo | Valor |
|---|---|
| **Cliente** | Menegatti Couro |
| **Setor** | moda / couro |
| **Status** | em qualificação — saindo de Tray |
| **URL** | `menegatticouro.com.br` |
| **Hits no DB** | 2 |
| **Última menção** | 2025-11-24 |
| **Grupo principal** | Wake <> Nairuz / Wake <> 3W |
| **Tags** | `audience-sales`, `lead-pipeline` |

**Contexto:** Em 2025-11-24 lead "Menegatti Couro - Commerce" distribuído em paralelo para Nairuz e 3W, contexto: "Plataforma atual: Tray. Comentou que a Tray não atende de for..." (truncado).

**Por que está aqui (não em wake-customer-cases.md):** Sinal claro de churn-from-Tray mas sem evolução. Promover quando agência confirmar avanço. Boa candidata para battlecard Tray→Wake.

---

### LEAD-015 — Made in Guarda

| Campo | Valor |
|---|---|
| **Cliente** | Made in Guarda |
| **Setor** | ? |
| **Status** | em qualificação — saindo de WooCommerce |
| **URL** | `madeinguarda.com.br` |
| **Hits no DB** | 2 |
| **Última menção** | 2025-11-25 |
| **Grupo principal** | Agencia Virtualiti | Wake / Wake+Omni360+Affys |
| **Tags** | `audience-sales`, `lead-pipeline` |

**Contexto:** Em 2025-11-25 oportunidade "Made in Guarda - Canais" distribuída para Virtualiti e Omni360+Affys: "Cliente da WooCommerce. O lead está pesquisando o mercado." 11,8k de visualização.

**Por que está aqui (não em wake-customer-cases.md):** Pesquisa de mercado em fase inicial, sem decisão. Promover quando houver sinal de proposta ou agência designada.

---

### LEAD-016 — Loja Brasil Vida

| Campo | Valor |
|---|---|
| **Cliente** | Loja Brasil Vida |
| **Setor** | ? |
| **Status** | em qualificação — sob a Avocado Tech |
| **URL** | `lojabrasilvida.com.br` |
| **Hits no DB** | 2 (Wake <> Avocado Tech) |
| **Última menção** | 2026-01-26 |
| **Grupo principal** | Wake <> Avocado Tech |
| **Tags** | `audience-pam`, `lead-pipeline` |

**Contexto:** URL compartilhada duas vezes (2025-11-27 e 2026-01-26) por Avocado Tech sem contexto adicional nas mensagens cabeçalho. Sinal pertinente à carteira Avocado Tech.

**Por que está aqui (não em wake-customer-cases.md):** Sinal mínimo — duas postagens de URL apenas. Promover quando aparecer detalhe (ERP, GMV, plataforma atual).

---

### LEAD-017 — Magazine Claudia Lima

| Campo | Valor |
|---|---|
| **Cliente** | Magazine Claudia Lima |
| **Setor** | calçados / moda |
| **Status** | ⚠ provável cliente já existente — operação em produção (Black Friday ativa) |
| **URL** | `magazineclaudialima.com.br` |
| **Hits no DB** | 2 (Ecommerce House) |
| **Última menção** | 2025-11-29 |
| **Grupo principal** | Ecommerce House |
| **Tags** | `audience-sales`, `marketing-eventos` |

**Contexto:** Em 2025-11-27 e 2025-11-29 mensagens promocionais de Black Friday na URL: marcas Dakota, Bottero, Piccadilly, Umbro, Molekinha — sugere multimarca de calçados. Site ativo com produtos.

**Por que está aqui (não em wake-customer-cases.md):** Sinal apenas de marketing/comunidade, não de pipeline. Pode já ser cliente Wake — validar com Sales Ops antes de promover.

---

### LEAD-018 — Magnani Embalagens

| Campo | Valor |
|---|---|
| **Cliente** | Magnani Embalagens |
| **Setor** | embalagens |
| **Status** | ⚠ desconhecido — sinal de marketing apenas |
| **URL** | `magnaniembalagens.com.br` |
| **Hits no DB** | 2 (Ecommerce House) |
| **Última menção** | 2025-12-15 |
| **Grupo principal** | Ecommerce House |
| **Tags** | `audience-sales`, `lead-pipeline` |

**Contexto:** Em 2025-12-15 URLs compartilhadas com parâmetros de Google Ads (campanhas brand ativas) — sugere operação digital ativa. Sem registro de oportunidade ou agência designada.

**Por que está aqui (não em wake-customer-cases.md):** Apenas links promocionais. Validar status comercial antes de promover.

---

### LEAD-019 — Você Constrói

| Campo | Valor |
|---|---|
| **Cliente** | Você Constrói |
| **Setor** | construção / B2B |
| **Status** | ✓ **cliente em produção** — incidente de tracking GA4/Meta reportado |
| **URL** | `voceconstroi.com.br` |
| **Hits no DB** | 2 (Eficaz <> WAKE) |
| **Última menção** | 2026-02-18 |
| **Grupo principal** | Eficaz <> WAKE |
| **Tags** | `audience-cs`, `suporte-escalação` |

**Contexto:** Em 2025-12-15 (Eficaz <> WAKE): evento Purchase está disparando no Meta sem parâmetro de receita (value=0), destruindo ROAS. Em 2026-02-18: incidente reaberto — "Pessoal, precisamos de ajuda urgente com o cliente Voce constroi. Após um ajuste no Front feito pela Auaha em novembro, o evento Purchase passou a não contabilizar a receita".

**Por que está aqui (não em wake-customer-cases.md):** Cliente em produção (ligado a CASE-051 Esplane segundo gap list). Promover como CASE-NNN próprio ou anexar como subcustomer de Esplane. Boa candidata para troubleshooting article (TRB) sobre evento Purchase + Meta CAPI sem valor.

---

### LEAD-020 — Camys

| Campo | Valor |
|---|---|
| **Cliente** | Camys |
| **Setor** | ? |
| **Status** | ✓ **cliente Wake gravando case** — case audiovisual em produção |
| **URL** | `camys.com.br` |
| **Hits no DB** | 2 (Mastercard & Wake, Agencia Virtualiti) + 3 menções em WAKE | Time Parcerias e The Rise of Skywalker |
| **Última menção** | 2026-06-01 |
| **Grupo principal** | The Rise of Skywalker / WAKE | Time Parcerias |
| **Tags** | `audience-sales`, `marketing-eventos` |

**Contexto:** Em 2026-04-13 (WAKE | Time Parcerias) menção de implementação parcial (Mastercard? feature) que foi desativada. Em 2026-05-15 (The Rise of Skywalker): "acabamos de sair da Camys, gravamos um case super massa!!!". Em 2026-06-01: confirmação de pessoa contato da Camys.

**Por que está aqui (não em wake-customer-cases.md):** Cliente em produção com case sendo gravado — sinal forte. Promover assim que o case oficial Wake for publicado. Boa candidata para CASE com data de publicação do case marketing.

---

### LEAD-021 — Puriflora

| Campo | Valor |
|---|---|
| **Cliente** | Puriflora |
| **Setor** | ? — modelo de assinaturas/recorrência, vendas via telefone |
| **Status** | em qualificação — reunião positiva com Wake |
| **URL** | `puriflora.com.br` |
| **Hits no DB** | 2 |
| **Última menção** | 2026-01-27 |
| **Grupo principal** | Eficaz <> WAKE / Wake+Omni360+Affys |
| **Tags** | `audience-sales`, `proposta` |

**Contexto:** Em 2026-01-15 (Wake+Omni360+Affys) lead distribuído ao Canais. Em 2026-01-27 (Eficaz <> WAKE): "Wake teve muita aderência ao cenário deles, principalmente pelo modelo de assinaturas/recorrência, que hoje é a principal dor da operação. A Puriflora tem uma venda muito forte no telefone e o digital a..."

**Por que está aqui (não em wake-customer-cases.md):** Sinal qualitativo forte (aderência confirmada, dor de assinatura mapeada) mas sem proposta enviada nos hits. Promover quando proposta for enviada — caso de uso de recorrência para battlecard.

---

### LEAD-022 — Eleve Suplementos

| Campo | Valor |
|---|---|
| **Cliente** | Eleve Suplementos |
| **Setor** | suplementos |
| **Status** | em qualificação — opt Canais |
| **URL** | `elevesuplementos.com.br` |
| **Hits no DB** | 2 |
| **Última menção** | 2026-01-15 |
| **Grupo principal** | Wake+Omni360+Affys / Wake <> 3W |
| **Tags** | `audience-sales`, `lead-pipeline` |

**Contexto:** Em 2026-01-15 lead enviado em paralelo para Omni360+Affys e 3W como opt de Canais (texto template padrão "agência deve seguir a regra de atender de ponta a ponta").

**Por que está aqui (não em wake-customer-cases.md):** Distribuição template sem follow-up nos hits. Promover quando uma agência sinalizar progresso.

---

### LEAD-023 — Termomax

| Campo | Valor |
|---|---|
| **Cliente** | Termomax |
| **Setor** | industrial ⚠ inferido |
| **Status** | em qualificação — reunião de proposta desmarcada |
| **URL** | `termomax.com.br` |
| **Hits no DB** | 2 (SAVVI & Wake / O2 Consultoria) + acompanhamento em SAVVI |
| **Última menção** | 2026-03-17 |
| **Grupo principal** | SAVVI & Wake |
| **Tags** | `audience-sales`, `proposta` |

**Contexto:** Em 2026-01-28 (SAVVI & Wake / O2 Consultoria) lead enviado. Em 2026-03-06 (O2 Consultoria): "você conseguiu falar com a Ana da Termomax?". Em 2026-03-11 (SAVVI): "Termomax: Reunião de apresentação de Proposta 17/03". Em 2026-03-17: "o lead Termomax desmarcou a reunião de apresentação de proposta que estava agendada para as 15h".

**Por que está aqui (não em wake-customer-cases.md):** Sinal de proposta em curso mas com fricção (reunião desmarcada). Promover quando aparecer status final (ganho, perdido, em standby).

---

### LEAD-024 — Tanto Quanto Revestimentos

| Campo | Valor |
|---|---|
| **Cliente** | Tanto Quanto Revestimentos |
| **Setor** | revestimentos / construção |
| **Status** | em pausa — "estão com outras prioridades" (2026-03-11) |
| **URL** | `tanto.com.br` |
| **Hits no DB** | 2 |
| **Última menção** | 2026-03-11 |
| **Grupo principal** | SAVVI & Wake / O2 Consultoria & Wake |
| **Tags** | `audience-sales`, `lead-pipeline` |

**Contexto:** Em 2026-02-05 (SAVVI & Wake / O2 Consultoria) oportunidade "TANTO QUANTO REVESTIMENTOS - Canais" distribuída — cliente declarou ter agência externa Wake. Em 2026-03-11 (SAVVI) follow-up de Rogerio: "Tanto: Não tive mais retorno apos Rogerio ter dito que estão com outras prioridades".

**Por que está aqui (não em wake-customer-cases.md):** Lead esfriou. Promover só se aquecer novamente.

---

### LEAD-025 — Dominik Metalcenter

| Campo | Valor |
|---|---|
| **Cliente** | Dominik Metalcenter |
| **Setor** | industrial / metais |
| **Status** | em qualificação — encaminhado via carrossel para Série//A |
| **URL** | `dominik.com.br` |
| **Hits no DB** | 2 (Wake <> Série//A) + 1 (Wake <> Catus, com nome de executiva trocada) |
| **Última menção** | 2026-02-19 |
| **Grupo principal** | Wake <> Série//A |
| **Tags** | `audience-pam`, `lead-pipeline` |

**Contexto:** Em 2026-02-06 (Wake <> Catus) lead chegou com nome de executiva trocada — "Chegou outro de vcs mas esta com o nome da Vanessa Placoná como executiva". Em 2026-02-13 (Série//A): "Dominik: o Guto está tocando essa negociação, ainda não é cliente Wake, e direcionamos para vocês via carrossel" — sem informação no Salesforce. Em 2026-02-19: confirmação de URL.

**Por que está aqui (não em wake-customer-cases.md):** Há ambiguidade de atribuição/executivo (carrossel chegou em Catus mas direcionado para Série//A). Sinal limpo mas pequeno. Promover quando Guto avançar a negociação.

---

### LEAD-026 — Editora Brasil Seikyo

| Campo | Valor |
|---|---|
| **Cliente** | Editora Brasil Seikyo |
| **Setor** | editora / livros / clube de assinatura |
| **Status** | em qualificação — opt Canais, ERP Protheus, GMV 100k/mês |
| **URL** | `cile.brasilseikyo.com.br` |
| **Hits no DB** | 2 (O2 Consultoria, Duo) + 1 (Duo, 2026-04-27) |
| **Última menção** | 2026-04-27 |
| **Grupo principal** | Duo e Wake Parcerias / O2 Consultoria & Wake |
| **Tags** | `audience-solutions`, `lead-pipeline` |

**Contexto:** Em 2026-02-18 (O2 Consultoria, Duo) oportunidade detalhada: "Editora Brasil Seikyo — ERP: Protheus, GMV: 100k mês. Atualmente, a editora possui um e-commerce de livros e um clube de assinatura..." Em 2026-04-27 (Duo): "é da editora Brasil Seikyo né?" confirmando que ainda está sendo trabalhado.

**Por que está aqui (não em wake-customer-cases.md):** Sinal técnico forte (Protheus + assinatura) mas ainda sem proposta fechada. Promover quando avançar. Boa candidata para case de assinaturas + Protheus.

---

### LEAD-027 — Malhas GDom

| Campo | Valor |
|---|---|
| **Cliente** | Malhas GDom |
| **Setor** | malhas / têxtil |
| **Status** | em qualificação — call inicial feita |
| **URL** | `gdom.com.br` |
| **Hits no DB** | 2 |
| **Última menção** | 2026-02-26 |
| **Grupo principal** | Duo e Wake Parcerias / Wake <> Catus |
| **Tags** | `audience-sales`, `lead-pipeline` |

**Contexto:** Em 2026-02-26 mesmo lead distribuído para Duo e Catus: "Malhas GDom - Commerce - Canais. Acabei de fazer a call com o Douglas e o Diogo. Apresentei a plataforma dando ênfase nos diferenciais da Wake como enterprise, mas é legal você como agência reforçar os..." (truncado).

**Por que está aqui (não em wake-customer-cases.md):** Call inicial feita, sem follow-up posterior. Promover quando agência reportar próximo passo.

---

### LEAD-028 — Casa Pedro

| Campo | Valor |
|---|---|
| **Cliente** | Casa Pedro |
| **Setor** | ? — pedidos via WhatsApp + IA, migrando para e-commerce profissional |
| **Status** | em qualificação — projeto re-direcionado |
| **URL** | `casaspedro.com.br` |
| **Hits no DB** | 2 |
| **Última menção** | 2026-05-06 |
| **Grupo principal** | Wake <> Nairuz / Parcerias Dotkom & Wake |
| **Tags** | `audience-sales`, `lead-pipeline` |

**Contexto:** Em 2026-05-04 (Dotkom): "Wander atualizou hoje comigo o direcionamento do projeto e querem simplificar; O Wander atualizou que o projeto mudou de direção". Em 2026-05-06 (Nairuz): "Antes, a ideia era uma solução mais simples para pedidos via WhatsApp com IA. Agora, eles querem partir para um e-commerce mais profissional. O foco inicial passou a ser uma loja virtual..." (truncado).

**Por que está aqui (não em wake-customer-cases.md):** Mudança de escopo recente — promover quando definirem agência e plataforma.

---

### LEAD-029 — Crise CIA

| Campo | Valor |
|---|---|
| **Cliente** | Crise CIA |
| **Setor** | ? — operação de R$2MM ("entrou vendendo 2MM") |
| **Status** | ⚠ desconhecido — sinal de descoberta recente |
| **URL** | `crisecia.com.br` |
| **Hits no DB** | 2 |
| **Última menção** | 2026-06-02 |
| **Grupo principal** | Comercial Cross WAKE / The Rise of Skywalker |
| **Tags** | `audience-sales`, `lead-pipeline` |

**Contexto:** Em 2026-05-26 (The Rise of Skywalker): "Caraca, esse cara entrou vendendo 2MM https://www.crisecia.com.br/". Em 2026-06-02 (Comercial Cross WAKE) URL re-compartilhada sem contexto adicional.

**Por que está aqui (não em wake-customer-cases.md):** Sinal de descoberta de operação de alto faturamento, mas zero contexto de processo comercial. Promover quando aparecer agência atribuída ou abertura de lead formal.

---

### LEAD-030 — Fallco

| Campo | Valor |
|---|---|
| **Cliente** | Fallco |
| **Setor** | ? |
| **Status** | em qualificação — solicitação de pagamento de lead em discussão |
| **URL** | `fallco.com.br` |
| **Hits no DB** | 2 (FZ COMMERCE & Wake) |
| **Última menção** | 2025-08-21 |
| **Grupo principal** | FZ COMMERCE & Wake |
| **Tags** | `audience-pam`, `lead-pipeline` |

**Contexto:** Em 2025-08-21 (FZ Commerce): "vocês conseguem me pagar o lead da Fallco?". URL `fallco.com.br`. Contato: Diego (`diego@fallco.com.br`), tel 41 9687-9789. Sinal de "Migração Ecomm" no rodapé.

**Por que está aqui (não em wake-customer-cases.md):** Discussão de pagamento de lead-fee, contexto comercial limitado sobre o cliente. Promover quando agência formalizar oportunidade.

---

### LEAD-031 — TJCC (evento)

| Campo | Valor |
|---|---|
| **Cliente** | TJCC |
| **Setor** | evento / congresso ⚠ inferido |
| **Status** | ❌ **não é cliente** — evento, NÃO cliente Wake |
| **URL** | `congresso.tjcc.com.br` |
| **Hits no DB** | 2 (SoulPink Geral) + menções em AnjoRosa |
| **Última menção** | 2026-05-05 |
| **Grupo principal** | SoulPink Geral / AnjoRosa: Vendas |
| **Tags** | `audience-pam`, `marketing-eventos` |

**Contexto:** TJCC aparece em contexto de evento — "vou passar os códigos pro TJCC" (códigos de evento). Em AnjoRosa: "O evento da AZ que mudou para o TJCC" e "negociamos entregas para o TJCC". É evento, não cliente Wake.

**Por que está aqui (não em wake-customer-cases.md):** **Drop candidate** — é evento/congresso (provavelmente Congresso de Jovens / similar) onde criadores/marcas participam, não cliente da plataforma. Manter no watchlist apenas como nota de "não promover".

---

### LEAD-032 — Empório Bud

| Campo | Valor |
|---|---|
| **Cliente** | Empório Bud |
| **Setor** | ? |
| **Status** | em qualificação — indicação Konduto com contato + CNPJ |
| **URL** | `emporiobud.com.br` |
| **Hits no DB** | 2 |
| **Última menção** | 2025-09-12 |
| **Grupo principal** | Konduto & Wake / Wake <> Oásis |
| **Tags** | `audience-sales`, `lead-pipeline` |

**Contexto:** Em 2025-07-06 (Wake <> Oásis): Convite 3 do evento "Empório Bud, Fernando Felix, 11 98121-0915". Em 2025-09-12 (Konduto & Wake): "Indicação de cliente: Empório Bud, fernando@emporiobud.com.br, 11 98121-0915 - Fernando Felix, CNPJ: 11..." — indicação formal com dados completos.

**Por que está aqui (não em wake-customer-cases.md):** Indicação Konduto com pacote de dados completo — sinal forte para promoção. Promover assim que ficar claro qual agência tocou o lead.

---

### LEAD-033 — Loja Orio do Peixe

| Campo | Valor |
|---|---|
| **Cliente** | Loja Orio do Peixe |
| **Setor** | ? |
| **Status** | em qualificação — discussão de HERO Seguros para o cliente |
| **URL** | `lojaoriodopeixe.com.br` |
| **Hits no DB** | 2 |
| **Última menção** | 2025-09-19 |
| **Grupo principal** | WAKE | Time Agência / FG + Wake |
| **Tags** | `audience-solutions`, `lead-pipeline` |

**Contexto:** Em 2025-08-04 (FG + Wake): contato "jobsonluiz@lojaoriodopeixe.com.br" capturado. Em 2025-09-19 (WAKE | Time Agência): URL + "cabe hero seguros" — sugere fit para integração HERO.

**Por que está aqui (não em wake-customer-cases.md):** Sinal técnico (fit HERO) e contato capturado, mas sem oportunidade formalizada. Promover quando agência confirmar avanço.

---

## Structured-field extras

### LEAD-034 — Grupo Oásis

| Campo | Valor |
|---|---|
| **Cliente** | Grupo Oásis |
| **Setor** | ? — coordenação de performance, tech lead presentes |
| **Status** | ⚠ desconhecido — convidados de evento |
| **URL** | — |
| **Hits no DB** | 2+ (Wake <> Oásis, todos como convidados) |
| **Última menção** | 2025-07-06 |
| **Grupo principal** | Wake <> Oásis |
| **Tags** | `audience-pam`, `marketing-eventos` |

**Contexto:** Em 2025-07-06 mensagens "Convite 4" e "Convite 5" para evento citam Thamires Ferreira Carpi (`thamires@agenciaoasis.com.br`) e Lucius Faltz Lassarote da Silva (`lucius@agenciaoasis.com.br`) como "Empresa: Grupo Oásis". **Atenção**: emails são `@agenciaoasis.com.br` — pode ser que Grupo Oásis seja a própria agência Oásis, não cliente.

**Por que está aqui (não em wake-customer-cases.md):** Provável que seja o nome de holding da agência Oásis, não cliente da plataforma. **Drop candidate** após validação.

---

### LEAD-035 — Cogem (Cooperativa)

| Campo | Valor |
|---|---|
| **Cliente** | Cogem — Cooperativa de Crédito |
| **Setor** | cooperativa de crédito |
| **Status** | ⚠ provavelmente fora de escopo Wake — projeto de evento/conteúdo |
| **URL** | — |
| **Hits no DB** | 2 (Comercial Anjo Rosa) |
| **Última menção** | 2025-10-16 |
| **Grupo principal** | Comercial Anjo Rosa |
| **Tags** | `audience-sales`, `lead-pipeline` |

**Contexto:** Em 2025-10-14 (Comercial Anjo Rosa): "fiz um alinhamento com o pessoal da Cogem, estão ansiosos para amanhã, falei para entrarmos 08h50 no link, pois começa 09h... pelo Teams, com transmissão no youtube para as empresas que...". Em 2025-10-16: "Feedback Cogem + Polo Palestra".

**Por que está aqui (não em wake-customer-cases.md):** Contexto sugere evento/conteúdo com influencer (AnjoRosa), não venda de plataforma. **Drop candidate** salvo nova evidência.

---

### LEAD-036 — Adecil (seller Anymarket)

| Campo | Valor |
|---|---|
| **Cliente** | Adecil |
| **Setor** | varejo (seller Mercado Livre via Anymarket) |
| **Status** | ✓ **cliente Anymarket em produção** — cenário crítico de importação de pedidos ML |
| **URL** | — |
| **Hits no DB** | 2 (ANYMARKET & WAKE TÁTICO/OPERACIONAL) |
| **Última menção** | 2026-05-19 |
| **Grupo principal** | ANYMARKET & WAKE (TÁTICO / OPERACIONAL) |
| **Tags** | `audience-cs`, `suporte-escalação` |

**Contexto:** Em 2026-05-19 ticket aberto: "Estamos com um cenário extremamente crítico com o seller Adecil (Neto). A operação dele está com sendo muito impactada por gargalos na importação de pedidos do Mercado Livre (erros de SKUs e conflitos de..." — ticket Anymarket #21735 com solicitação de prioridade.

**Por que está aqui (não em wake-customer-cases.md):** Cliente Anymarket em produção com incidente crítico. Promover como CASE-NNN ou anexar a um case Anymarket. Boa candidata para troubleshooting article sobre SKU conflicts em importação ML.

---

### LEAD-037 — Grupo Reluz

| Campo | Valor |
|---|---|
| **Cliente** | Grupo Reluz |
| **Setor** | ? |
| **Status** | ⚠ convidado de evento |
| **URL** | — |
| **Hits no DB** | 1 (Wake <> Oásis) |
| **Última menção** | 2025-07-03 |
| **Grupo principal** | Wake <> Oásis |
| **Tags** | `audience-pam`, `marketing-eventos` |

**Contexto:** Em 2025-07-03 convite 1 para evento Wake-Oásis: Matheus Crescencio (CEO, Grupo Reluz, `matheus@gruporeluz.com.br`).

**Por que está aqui (não em wake-customer-cases.md):** Apenas convite de evento — CEO de prospect potencial. Promover quando aparecer follow-up pós-evento.

---

### LEAD-038 — Madel Home

| Campo | Valor |
|---|---|
| **Cliente** | MADEL HOME COMERCIO DE MOVEIS E DECORACOES LTDA |
| **Setor** | móveis e decoração |
| **Status** | ✓ **cliente Wake em renovação** — oportunidade "Renovação - Madel - Julho 2025" |
| **URL** | — |
| **Hits no DB** | 1 (WAKE | Time Agência) |
| **Última menção** | 2025-07-03 |
| **Grupo principal** | WAKE | Time Agência |
| **Tags** | `audience-pam`, `contrato-comercial` |

**Contexto:** Em 2025-07-03 (WAKE | Time Agência): "Nome da oportunidade: Renovação - Madel - Julho 2025. Nome da empresa: MADEL HOME COMERCIO DE MOVEIS E DECORACOES LTDA. Nome do executivo responsável: Samires Gomes (samires.souza@wake.tech)".

**Por que está aqui (não em wake-customer-cases.md):** Cliente em renovação — sinal forte de produção. Promover como CASE de cliente ativo. Apenas 1 hit no DB mas evidência clara de existência.

---

### LEAD-039 — Vou de Click

| Campo | Valor |
|---|---|
| **Cliente** | Vou de Click Ltda |
| **Setor** | ? |
| **Status** | em qualificação — Avocado Tech entrou com forecast |
| **URL** | — (contato `vanessa@voudeclick.com.br`) |
| **Hits no DB** | 2 (WAKE | Time Parcerias, WAKE | Time Agência) |
| **Última menção** | 2025-10-03 |
| **Grupo principal** | WAKE | Time Agência |
| **Tags** | `audience-pam`, `lead-pipeline` |

**Contexto:** Em 2025-09-12 (WAKE | Time Parcerias): dados completos — Vanessa Tatiane da Silva Bermudes, Vou de Click Ltda, CNPJ 52.885.806/0001-08, `vanessa@voudeclick.com.br`. Em 2025-10-03 (WAKE | Time Agência): "Avocado Tech entrou com os 3k que estão ai. Vou inserir apenas uma possível nova, Vou de Click".

**Por que está aqui (não em wake-customer-cases.md):** Lead com dados completos + agência atribuída (Avocado Tech). Promover quando aparecer status comercial.

---

### LEAD-040 — Livraria do IPP

| Campo | Valor |
|---|---|
| **Cliente** | Livraria do IPP |
| **Setor** | livraria |
| **Status** | em qualificação — opt distribuída via E-Plus |
| **URL** | — |
| **Hits no DB** | 1 (Wake <> E-Plus) |
| **Última menção** | 2025-09-18 |
| **Grupo principal** | Wake <> E-Plus |
| **Tags** | `audience-sales`, `lead-pipeline` |

**Contexto:** Em 2025-09-18 (Wake <> E-Plus): "Nome da oportunidade: Livraria do IPP - Commerce. Nome do executivo: Augusto Testa". Pedido de informações sobre o lead.

**Por que está aqui (não em wake-customer-cases.md):** Apenas 1 hit explícito. Promover quando avançar.

---

### LEAD-041 — Atacadão Maromba

| Campo | Valor |
|---|---|
| **Cliente** | Atacadão Maromba |
| **Setor** | B2B / atacado |
| **Status** | ⚠ desconhecido — mencionado em conjunto com Empório do Sertão |
| **URL** | — |
| **Hits no DB** | 2 (Coder Ivy + Wake, Wake <> Nairuz) |
| **Última menção** | 2026-05-19 |
| **Grupo principal** | Coder Ivy + Wake / Wake <> Nairuz |
| **Tags** | `audience-sales`, `lead-pipeline` |

**Contexto:** Em 2026-05-06 (Wake <> Nairuz): proposta detalhada para "Empório do Sertão (Projeto Omnichannel) - Varejo físico consolidado no Rio de Janeiro. Faturamento offline R$ 20 milhõe..." — Atacadão Maromba aparece em texto longo do mesmo conjunto de oportunidades. Em 2026-05-19 (Coder Ivy + Wake): "E Atacadão Maromba" — confirmação curta.

**Por que está aqui (não em wake-customer-cases.md):** Provavelmente parte de cluster Empório do Sertão / Atacadão Maromba — investigar se é mesma operação ou irmãs. Promover após desambiguar.

---

### LEAD-042 — Dominik Metalcenter

Ver LEAD-025 — mesma operação (Dominik). Não duplicar.

---

### LEAD-043 — Home Sport Center

| Campo | Valor |
|---|---|
| **Cliente** | Home Sport Center |
| **Setor** | esporte |
| **Status** | ✓ **cliente Anymarket em produção** — tickets #10440 e #6136 abertos |
| **URL** | — |
| **Hits no DB** | 2 (ANYMARKET & WAKE TÁTICO) |
| **Última menção** | 2026-04-13 |
| **Grupo principal** | ANYMARKET & WAKE (TÁTICO / OPERACIONAL) |
| **Tags** | `audience-cs`, `suporte-escalação` |

**Contexto:** Em 2026-02-24: "Alguma informação sobre cenário home sport center?". Em 2026-04-13: pedido de status dos tickets #10440 e #6136, ambos da Home Sport Center.

**Por que está aqui (não em wake-customer-cases.md):** Cliente Anymarket em produção com histórico de tickets. Promover como CASE.

---

### LEAD-044 — Linha e Conforto

| Campo | Valor |
|---|---|
| **Cliente** | Linha e Conforto |
| **Setor** | casa / têxtil |
| **Status** | ✓ **cliente Anymarket em produção** — problema reportado diariamente |
| **URL** | — |
| **Hits no DB** | 2 (ANYMARKET & WAKE TÁTICO) |
| **Última menção** | 2026-03-13 |
| **Grupo principal** | ANYMARKET & WAKE (TÁTICO / OPERACIONAL) |
| **Tags** | `audience-cs`, `suporte-escalação` |

**Contexto:** Em 2026-03-12: ticket Anymarket #5546 — "Loja: Linha e Conforto - podem validar?". Em 2026-03-13: "todos os dias tem reportado problema com os produtos dele dentro da Any. Existe alguma disponibilidade do lado de vocês para marcarmos uma agenda com este cliente e sanarmos..."

**Por que está aqui (não em wake-customer-cases.md):** Cliente Anymarket em produção com escalonamento contínuo. Promover como CASE.

---

### LEAD-045 — Cacto Fitwear

| Campo | Valor |
|---|---|
| **Cliente** | Cacto Fitwear |
| **Setor** | moda fitness |
| **Status** | em qualificação — via Engajar |
| **URL** | — |
| **Hits no DB** | 2 (Wake <> Engajar, WAKE | Time Parcerias) |
| **Última menção** | 2026-03-04 |
| **Grupo principal** | Wake <> Engajar / WAKE | Time Parcerias |
| **Tags** | `audience-sales`, `lead-pipeline` |

**Contexto:** Em 2025-11-18 (WAKE | Time Parcerias): "A Cacto Fitwear, cliente do segmento de moda fitness, chegou até nós através da Engajar, uma agência recém integrada ao nosso ecossistema que já demonstrou habilidade em vender a solução Wake de ponta a ponta. Com um faturamento atu..." Em 2026-03-04 (Wake <> Engajar): "Cliente: Cacto Fitwear".

**Por que está aqui (não em wake-customer-cases.md):** Sinal forte (descritivo + agência confirmada) mas só 2 hits. Promover quando aparecer status de proposta/closing.

---

### LEAD-046 — AnjoRosa Consultoria em Saúde

| Campo | Valor |
|---|---|
| **Cliente** | ANJOROSA CONSULTORIA EM SAUDE E BEM ESTAR LTDA |
| **Setor** | saúde / consultoria — influenciador |
| **Status** | ⚠ não é cliente da plataforma — empresa de influenciador (contrato com Estadão) |
| **URL** | — |
| **Hits no DB** | 1 (Comercial Anjo Rosa) |
| **Última menção** | 2026-03-14 |
| **Grupo principal** | Comercial Anjo Rosa |
| **Tags** | `audience-pam`, `contrato-comercial` |

**Contexto:** Em 2026-03-14: "Dados empresa do influenciador. Razão Social: ANJOROSA CONSULTORIA EM SAUDE E BEM ESTAR LTDA, CNPJ 47.82..." — jurídico do Estadão pediu para fazer contrato.

**Por que está aqui (não em wake-customer-cases.md):** **Drop candidate** — é PJ de influenciadora (creator economy / Wake Creators?), não cliente da plataforma de e-commerce.

---

### LEAD-047 — Central Acessórios

| Campo | Valor |
|---|---|
| **Cliente** | Central Acessórios |
| **Setor** | acessórios ⚠ inferido |
| **Status** | em qualificação — indicação Nairuz |
| **URL** | `centralcessorios.com.br` |
| **Hits no DB** | 1 (Wake <> Nairuz) |
| **Última menção** | 2026-04-23 |
| **Grupo principal** | Wake <> Nairuz |
| **Tags** | `audience-sales`, `lead-pipeline` |

**Contexto:** Em 2026-04-23 (Wake <> Nairuz): "Nome do cliente: Paloma Cesarani, Da empresa: Central Acessório, Telefone: 11 94504-5820, E-mail: paloma@centralacessorios.com.br, Site www.centralcessorios.com.br".

**Por que está aqui (não em wake-customer-cases.md):** Apenas dados de indicação, sem evolução. Promover quando aparecer status.

---

### LEAD-048 — Rabusch / GF10

| Campo | Valor |
|---|---|
| **Cliente** | Rabusch (Razão Social: GF10 INDUSTRIA E COMÉRCIO DE VESTUÁRIO LTDA) |
| **Setor** | moda / vestuário |
| **Status** | em qualificação — indicação Catus |
| **URL** | — |
| **Hits no DB** | 1 (Wake <> Catus) |
| **Última menção** | 2026-06-02 |
| **Grupo principal** | Wake <> Catus |
| **Tags** | `audience-sales`, `lead-pipeline` |

**Contexto:** Em 2026-06-02 (Wake <> Catus): "Nome da empresa: Rabusch. Razão Social: GF10 INDUSTRIA E COMÉRCIO DE VESTUÁRIO LTDA. CNPJ: 91.289.934/0001-40. Nome do contato: Telmo Leguissamo. Email: telmo.leguissamo@rabusch.com.br".

**Por que está aqui (não em wake-customer-cases.md):** Sinal único e recente (junho/2026). Consolida duas entradas do gap list (Rabusch + GF10 são mesma operação). Promover quando avançar.

---

### LEAD-049 — MPS Distribuidora Mercantil

| Campo | Valor |
|---|---|
| **Cliente** | MPS DISTRIBUIDORA MERCANTIL LTDA (oportunidade "A Dias") |
| **Setor** | B2B / distribuição |
| **Status** | em qualificação — oportunidade Trad-corp/Virtualiti |
| **URL** | — |
| **Hits no DB** | 1 (Agencia Virtualiti) |
| **Última menção** | 2025-06-30 |
| **Grupo principal** | Agencia Virtualiti | Wake |
| **Tags** | `audience-sales`, `lead-pipeline` |

**Contexto:** Em 2025-06-30 (Agencia Virtualiti): "Nome da oportunidade: A Dias-. Nome da empresa: MPS DISTRIBUIDORA MERCANTIL LTDA. Nome do executivo: Augusto Testa (augusto.testa@traycorp.com.br)" — atenção: email do executivo é Tray Corp.

**Por que está aqui (não em wake-customer-cases.md):** Sinal único e antigo (junho/2025). Executivo é de Tray Corp — investigar se é prospect compartilhado ou lead cross-platform. Promover só após confirmar contexto Wake.

---

### LEAD-050 — RENSZ Calçados

| Campo | Valor |
|---|---|
| **Cliente** | RENSZ CALCADOS |
| **Setor** | calçados |
| **Status** | em qualificação — Marcos Esperoni (executivo Wake) |
| **URL** | — (contato `juan.oliveira@rensz.com.br`) |
| **Hits no DB** | 2 (Wake <> Catus, Kobe Apps <> Wake) |
| **Última menção** | 2025-10-13 |
| **Grupo principal** | Wake <> Catus |
| **Tags** | `audience-pam`, `lead-pipeline` |

**Contexto:** Em 2025-09-02 (Kobe Apps <> Wake): email `juan.oliveira@rensz.com.br`. Em 2025-10-13 (Wake <> Catus): "Nome da oportunidade: Rensz. Nome da empresa: RENSZ CALCADOS. Nome do executivo: Marcos Esperoni (marcos.esperoni@wake.tech)".

**Por que está aqui (não em wake-customer-cases.md):** Lead com executivo Wake atribuído mas só 2 hits. Promover quando avançar.

---

### LEAD-051 — Digitron

| Campo | Valor |
|---|---|
| **Cliente** | DIGITRON |
| **Setor** | ? |
| **Status** | em qualificação — indicação Ana Ferrari, atribuição em discussão |
| **URL** | — |
| **Hits no DB** | 2 (Comercial Cross WAKE, Wake <> E-Plus) |
| **Última menção** | 2025-12-18 |
| **Grupo principal** | Wake <> E-Plus |
| **Tags** | `audience-pam`, `lead-pipeline` |

**Contexto:** Em 2025-12-15 (Wake <> E-Plus): "Recebemos a indicação do lead DIGITRON que é da Ana Ferrari, por gentileza poderiam compartilhar o Whatsapp dela? Nome da oportunidade: Digitron- Commerce". Em 2025-12-18 (Comercial Cross WAKE): "Digitron vai ficar com quem, @53253383397513?".

**Por que está aqui (não em wake-customer-cases.md):** Discussão de atribuição em andamento — promover quando atribuído.

---

### LEAD-052 — Duloren

| Campo | Valor |
|---|---|
| **Cliente** | Duloren |
| **Setor** | moda íntima |
| **Status** | ✓ **cliente em projeto** — Go-live em evolução, alinhamento Wake×Econverse |
| **URL** | — |
| **Hits no DB** | 2 (Wake <> Econverse) |
| **Última menção** | 2026-05-14 |
| **Grupo principal** | Wake <> Econverse |
| **Tags** | `audience-cs`, `integração` |

**Contexto:** Em 2026-05-14 (Wake <> Econverse): "Pavoni nosso diretor de operações me acionou para alinharmos uma agenda sobre o cliente Duloren. A ideia é revisarmos juntos o cenário atual do projeto e próximos passos para evolução do Go-live". Confirmação de Leandro como Head de Operações.

**Por que está aqui (não em wake-customer-cases.md):** Cliente em projeto ativo via Econverse. Promover como CASE quando aparecer detalhe técnico ou data de go-live.

---

## Candidatos descartados (drop)

Durante a investigação no DB, os seguintes itens do gap list foram descartados como **não-clientes**:

| Nome | Motivo do drop |
|---|---|
| Babel-Azza | Sem hits no DB com esse nome; "Babel" só aparece em mensagens promocionais da Ecommerce House sobre votação TOP 100 Gestores. |
| Empório das Toalhas | Zero hits no DB. |
| Snacks do Johnny | Aparece em grupo "Investidores #22" anunciando rodada de R$2MM — startup investida, não cliente Wake. |
| Urban Helmets | Em WAKE \| Time Parcerias (2026-03-19): "URL deles ativa até hoje no cliente Urban Helmets" — é referência de URL técnica de **parceiro de integração** dentro de outro cliente, não cliente próprio. |
| Magazord | Aparece em vaga de coordenação de parcerias divulgada — **plataforma concorrente / agência**, não cliente. |
| Saurus | Único hit sem ligação a cliente — provavelmente termo coincidente em texto sobre ACC Modas (DRIVEN/Cava). |
| People Interactive | Hits são acidentais (nome embebido em texto de Palácio das Essências). Não é cliente. |
| Própria Day / Marca Própria Day | **Evento** (Anhembi SP, set/2025), não cliente. |
| Benova | Mencionada como parceira/agência em almoço/agenda — não cliente. |
| Stayfilm Inc (referência) | Já é parceiro Wake, não cliente. |
| Rabittohs (pocmoda) | Subdomínio `.fbits.app` — sandbox/POC, não cliente em produção. |
| Loja Slívia | Subdomínio `.fbits.app` — sandbox/POC. |
| VNDA | Plataforma concorrente (mencionada como churn-source). |
| Secret Outlet | Não localizado como cliente — único hit em checklist de abordagem Omnichat, sem contexto Wake direto. |

Adicionalmente, o LEAD-031 (TJCC), LEAD-034 (Grupo Oásis) e LEAD-046 (AnjoRosa Consultoria) foram **mantidos no watchlist com nota explícita** de que provavelmente não são clientes da plataforma — para evitar re-descoberta.

---

## URLs documentadas sem identificação confiável de cliente

Pulled do gap list seção "URLs found but no customer name yet identified":

| URL | Grupo | Nota |
|---|---|---|
| `tudoentregue.com.br` | WAKE \| Time Parcerias | URL pelada, sem contexto |
| `ed3.com.br` | WAKE \| Time Agência | Check de agência ("alguém conhece essa agencia?") — não cliente |
| `controlid.com.br` | — | Partner suggestion |
| `parcelex.com.br` | — | Partner suggestion |
| `equals.com.br` | — | Partner suggestion |
| `dguests.com.br` | — | Partner suggestion |
| `unabeleza.com.br` | — | **CHURN signal** — Unikapro migrou para Nuvem. Vale CASE de churn próprio. |
| `wavecommerce.com.br` | WAKE \| Time Parcerias | Agency partner reference |

---

## Promotion criteria

Um LEAD-NNN deste arquivo é promovido a CASE-NNN em `wake-customer-cases.md` quando atender ≥2 dos critérios:

- 3+ hits de mensagem no DB
- Contexto claro de setor + status + agência atribuída
- Sinal de proposta enviada, go-live, renovação ou ticket de produção
- O agente `wake-partnership-kb-writer` consegue marcar ✓ verificado em pelo menos a afirmação primária de status (em qualificação / em produção / perdido / renovação).

**Candidatos com sinal mais forte para promoção próxima** (já com sinal de produção/decisão):
- LEAD-019 Você Constrói (incidente em produção)
- LEAD-020 Camys (case audiovisual em produção)
- LEAD-036 Adecil (ticket crítico Anymarket)
- LEAD-038 Madel Home (renovação aberta)
- LEAD-043 Home Sport Center (tickets Anymarket #10440, #6136)
- LEAD-044 Linha e Conforto (incidente recorrente Anymarket)
- LEAD-052 Duloren (go-live em curso via Econverse)

---

## Verificação

| Tipo de query | Filtro padrão | Notas |
|---|---|---|
| Enriquecimento por nome/URL | `c.jid LIKE '%@g.us' AND (m.media_type IS NULL OR m.media_type='') AND m.content LIKE '%TERM%'` | LIMIT 2-4 por termo |
| Termos consultados | 48 nomes (medium-confidence) + 26 nomes (structured-field extras) + 9 termos de cross-check (Babel, Magazord, Saurus, People Interactive, Própria Day, Marca Própria Day, Benova, Empório Toalhas) | — |
| Hits efetivos | 2 hits/cliente confirmados na maioria; alguns nomes (Adecil, Madel Home, Vou de Click, Cacto Fitwear) tinham contexto rico em ≤2 hits | — |
| Drops por cross-check | 14 nomes descartados como não-cliente (lista acima) | Documentados para evitar re-descoberta |

Filtro de mensagens texto-em-grupo aplicado em **todas** as queries de enriquecimento.
