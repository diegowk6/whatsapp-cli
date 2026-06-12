# Análise Estratégica da Base de Conhecimento Wake

> Análise sobre 668 artigos extraídos dos grupos WhatsApp da Wake (FAQ, Troubleshooting, How-To, Best Practices), distribuídos em 14 grupos / domínios. As referências usam o formato `<arquivo> <ID>` (ex.: `kb-agencia FAQ-008`) conforme `kb-master-index.md`.

## Sumário Executivo

A base é rica mas fragmentada: o mesmo conhecimento crítico (certificação, Salesforce, carrossel, Wake Summit, ERPs, gateways, Wake vs. VTEX) é re-respondido dezenas de vezes em grupos diferentes, com pequenas variações. Quatro frentes destacam-se como prioritárias:

1. **Consolidar 23 clusters duplicados** em artigos canônicos versionados (redução estimada de ~30% do volume sem perda de informação).
2. **Cobrir 18 gaps estruturais** — começando por tabela de preços 2026, battlecard único e catálogo de integrações.
3. **Normalizar a taxonomia de tags** (40+ regras de fusão e hierarquia em 4 níveis).
4. **Tratar 7 riscos sistêmicos** — sendo o maior a dependência do WhatsApp como sistema de registro de conhecimento.

---

## 1. Duplicatas e Candidatos a Fusão

23 clusters de artigos que tratam do mesmo assunto e podem (devem) ser consolidados em artigos canônicos. Cada cluster lista o artigo-base sugerido, os artigos que devem ser fundidos e o que cada um agrega ao tema.

### Cluster 1 — Certificação Wake / Wake Score / elegibilidade para leads
- **Base**: `kb-parcerias-wave6c HOW_TO-9` ("Como funciona o Wake Score e a Certificação Wake").
- **Fundir**:
  - `kb-parceiros-agencia FAQ-2` (tipos de certificação Nairuz)
  - `kb-parceiros-wave5a FAQ-7` (A&EIGHT — Wake Score como requisito)
  - `kb-parceiros-wave7a FAQ-10` (3W — o que é Wake Score)
  - `kb-parceiros-marketplace FAQ-6` (Duo — certificação obrigatória)
  - `kb-parceiros-tech FAQ-5` (FG — Wake Education / certificação)
  - `kb-parceiros-wave5b HOW_TO-4` (DRIVEN — certificação de devs)
  - `kb-parceiros-wave7a HOW_TO-7` (Visie — certificação de agências)
  - `kb-agencia BP-004`, `kb-agencia BP-013` (certificação como pré-requisito para carrossel)
- **O que cada um agrega**: o base define o cálculo do Score; os de wave5/6/7 trazem requisitos por perfil (dev, agência, marketplace); os BP da kb-agencia explicitam que certificação destrava recebimento de leads.
- **Título consolidado**: *"Certificação Wake e Wake Score: pré-requisitos, trilhas, elegibilidade para leads e diferencial vs. VTEX"*.

### Cluster 2 — Wake Partner Day / Dia do Parceiro
- **Base**: `kb-parceiros-martech HOW_TO-5` (After Click — funcionamento).
- **Fundir**: `kb-parceiros-wave5b FAQ-2` (BonifiQ), `kb-parceiros-wave5b BEST_PRACTICES-3` (DRIVEN — quartas quinzenais), `kb-parceiros-tech HOW_TO-3` (Omnichat — LWSA), `kb-parceiros-wave6c HOW_TO-4` (Eitri — semanal), `kb-parceiros-wave6a BEST_PRACTICES-3` (Metakosmos — formato), `kb-parceiros-wave7a HOW_TO-2` (Vtrina), `kb-coordenacao BEST_PRACTICES-1` (agenda de quartas), `kb-coordenacao BEST_PRACTICES-8` (benefícios do presencial), `kb-parceiros-martech HOW_TO-11` (Revi — credenciais/Wi-Fi).
- **Agrega**: formato semanal, regras de credenciais, ROI do presencial, periodicidade por tier.
- **Título consolidado**: *"Wake Partner Day: formato semanal, regras de presença, credenciais e materiais"*.

### Cluster 3 — Registro de lead/indicação no Salesforce
- **Base**: `kb-parcerias FAQ-005` + `kb-agencia HT-002` (parceiro-origem).
- **Fundir**: `kb-parceiros-agencia HOW_TO-1` (Nairuz), `kb-parceiros-agencia HOW_TO-5` (Série//A — community), `kb-parceiros-tech HOW_TO-4` (FG — carrossel), `kb-parceiros-wave5a HOW_TO-1` (SAVVI), `kb-parceiros-wave5a FAQ-2` (SAVVI — evento), `kb-parceiros-wave5c HOW_TO-2` (Avocado), `kb-parceiros-wave5c HOW_TO-5` (Wicomm — lead inválido), `kb-parceiros-wave6a FAQ-1` (Lemoon — SDR), `kb-parceiros-wave6c HOW_TO-1` (Eitri — eventos), `kb-parceiros-wave7a HOW_TO-1` (Vtrina — formulário), `kb-parceiros-wave7a HOW_TO-6` (Visie — atribuição), `kb-parceiros-wave7a BEST_PRACTICES-5` (3W — registrar antes de evento), `kb-agencia FAQ-001`, `kb-agencia FAQ-003`, `kb-agencia FAQ-009`.
- **Agrega**: passo-a-passo por perfil (agência, tech, marketplace), tratamento de leads inválidos, registro pré-evento, casos sem botão de cadastro.
- **Título consolidado**: *"Como registrar leads e indicações de parceiros no Salesforce (carrossel, parceiro-origem, eventos)"*.

### Cluster 4 — Carrossel de leads do Salesforce
- **Base**: `kb-agencia FAQ-002`, `FAQ-003`, `FAQ-007` + `TRB-001` a `TRB-005`.
- **Fundir**: `kb-agencia HT-001` (indicação manual sem carrossel), `kb-agencia HT-003` (retirar agências indevidas), `kb-parceiros-agencia TROUBLESHOOTING-6` (Série//A), `kb-parceiros-tech TROUBLESHOOTING-2` (FG — leads com erro), `kb-parceiros-wave5c HOW_TO-5` (Wicomm — lead inválido).
- **Agrega**: lógica de distribuição, exceções, bugs conhecidos, recurso de leads não qualificados.
- **Título consolidado**: *"Carrossel de leads Wake/Salesforce: distribuição, exceções, indicação manual e troubleshooting"*.

### Cluster 5 — Programa Storefront 2.0 / Redesign de Checkout
- **Base**: `kb-parceiros-tech FAQ-7` (Eficaz).
- **Fundir**: `kb-parceiros-wave5c FAQ-4` (Wicomm), `kb-parceiros-wave6a FAQ-3` (Lemoon), `kb-parceiros-wave6c HOW_TO-8` (O2), `kb-parceiros-wave7a FAQ-9` (Visie — kit migração), `kb-parcerias FAQ-015` (custo da migração), `kb-comercial-cross FAQ-6` (nomenclatura headless).
- **Agrega**: escopo, custos por tier, empacotamento, elegibilidade da agência para tocar a migração.
- **Título consolidado**: *"Programa Storefront 2.0 e Redesign de Checkout: escopo, custos, empacotamento, elegibilidade"*.

### Cluster 6 — Campanha de Vendas Wake (incentivo a parceiros)
- **Base**: `kb-parceiros-wave5a FAQ-5` (Dotkom).
- **Fundir**: `kb-parceiros-agencia FAQ-4` (Nairuz), `kb-parceiros-wave6a FAQ-2` (Lemoon), `kb-parceiros-wave6c FAQ-2` (O2 — vale para tech?), `kb-parceiros-wave5c FAQ-5` (Wicomm — DXP/Commerce score), `kb-parceiros-wave7a HOW_TO-10` (3W), `kb-comercial-cross HOW-TO-4` (indicar parceiros tech), `kb-parcerias BP-005` (faixa 300k GMV).
- **Agrega**: escopo (Commerce vs. DXP vs. tech), faixas de GMV, fluxo de pagamento, perfis elegíveis.
- **Título consolidado**: *"Campanha de Vendas Wake: elegibilidade, escopo, faixas GMV, pagamento"*.

### Cluster 7 — Wake Summit
- **Base**: `kb-parcerias FAQ-006`, `FAQ-007`, `FAQ-008` + `BP-007`.
- **Fundir**: `kb-parceiros-wave6a FAQ-5`, `FAQ-6` (Metakosmos — QR/check-in CPF), `kb-parceiros-wave6b FAQ-3` (Estoca — brindes), `kb-parceiros-wave6b HOW_TO-4` (cota Silver), `kb-parceiros-wave6b HOW_TO-5` (podcast Summit), `kb-parceiros-wave6b BEST_PRACTICES-1` (Serasa — stand/palco), `kb-parceiros-wave6c FAQ-1` (Eitri — app do evento), `kb-parceiros-pagamento FAQ-8` (Vindi — expositor), `kb-parcerias HT-007` (jantar pré-Summit), `kb-parcerias BP-007` (cota maior = palco), `kb-comercial-cross BEST-PRACTICES-6`, `BEST-PRACTICES-12`.
- **Agrega**: cotas, palco, brindes, check-in, podcast, jantar, uso comercial.
- **Título consolidado**: *"Wake Summit — cotas, ingressos, palco, brindes, jantar, check-in"*.

### Cluster 8 — ERPs com integração à Wake
- **Base**: `kb-parcerias FAQ-012`.
- **Fundir**: `kb-parceiros-marketplace FAQ-7` (Duo), `kb-parceiros-agencia FAQ-5` (Catus — TOTVS Moda), `kb-parceiros-agencia FAQ-7` (Catus — Winthor), `kb-parceiros-wave6a FAQ-4` (Lemoon — intermediador), `kb-parceiros-wave7a FAQ-1`, `FAQ-2` (Vtrina — Winthor/Linx/Alpha7), `kb-parcerias TRB-007` (Everest — homologação), `kb-parceiros-wave5a BEST_PRACTICES-3` (SAVVI — seleção de ERP).
- **Agrega**: nativos, intermediadores (Vtrina, Bling, Proex), ERPs em homologação, conselho consultivo de seleção.
- **Título consolidado**: *"ERPs e integração Wake: nativos, intermediadores, Winthor/TOTVS/Linx/Everest"*.

### Cluster 9 — Wake vs. VTEX / Shopify / Linx (battlecard)
- **Base**: `kb-agencia BP-008` + `BP-013` (certificação como diferencial; sem exigência dos 3 projetos).
- **Fundir**: `kb-parceiros-agencia BEST_PRACTICES-3`, `4`, `6` (Nairuz/Catus), `kb-parcerias-marketing FAQ-5` (VTEX Rise), `kb-parcerias-marketing BP-6` (posicionamento vs. VTEX Rise), `kb-parceiros-wave6c BEST_PRACTICES-5` (Social — Shopify também), `kb-parceiros-wave7a HOW_TO-8` (Visie — Shopify).
- **Agrega**: objeções de preço, dependência de agência, certificação, comparativos com VTEX Rise e Shopify.
- **Título consolidado**: *"Wake vs. VTEX / Shopify / Linx: battlecard de objeções"*.

### Cluster 10 — Planos Essential, Special, Elite
- **Base**: `kb-agencia FAQ-008`.
- **Fundir**: `kb-parcerias FAQ-009` (Elite vs Special), `kb-parceiros-wave6a FAQ-10` (Appmax — Essential vs Special), `kb-parceiros-marketplace FAQ-9` (Duo — Partner Day/benefícios), `kb-parceiros-wave6a HOW_TO-6` (Estoca — entrada Elite), `kb-parceiros-wave6a BEST_PRACTICES-7` (Appmax — entregáveis Special), `kb-parceiros-wave6b BEST_PRACTICES-4` (parceiros concorrentes).
- **Agrega**: revenue share, partner fee, entregáveis por tier, regras com parceiros concorrentes.
- **Título consolidado**: *"Planos do programa de parcerias: Essential, Special, Elite (revenue share, fee, entregáveis)"*.

### Cluster 11 — B2B + B2C no mesmo ambiente
- **Base**: `kb-parceiros-tech FAQ-4` (FG).
- **Fundir**: `kb-parceiros-agencia HOW_TO-3` (Nairuz), `kb-parceiros-marketplace FAQ-7` (Catus — Winthor B2B), `kb-parceiros-marketplace FAQ-8` (Duo — múltiplos CNPJs), `kb-parceiros-wave5a HOW_TO-4` (Omni360 — proposta B2B+B2C), `kb-parceiros-wave6c FAQ-6` (Social — dúvidas técnicas B2B).
- **Agrega**: tabela de preços, múltiplos CNPJs, ERP (Winthor), proposta comercial.
- **Título consolidado**: *"Wake B2B + B2C no mesmo ambiente: preço, múltiplos CNPJs, ERP, proposta"*.

### Cluster 12 — Onboarding/KickOff de parceiro tech
- **Base**: `kb-coordenacao HOW-TO-4`.
- **Fundir**: `kb-parceiros-tech HOW_TO-6` (Leanwork), `kb-parceiros-tech HOW_TO-2` (Omnichat), `kb-parceiros-agencia HOW_TO-2` (Nairuz), `kb-parceiros-wave5b HOW_TO-3` (DRIVEN — Partner Ops), `kb-parceiros-wave7a HOW_TO-5` (Fidelizar+), `kb-parceiros-martech HOW_TO-1`, `HOW_TO-2` (Compra Rápida), `kb-parceiros-pagamento HOW_TO-3` (Givex — NDA).
- **Agrega**: contrato → NDA → ambiente trial → kickoff → entregáveis por perfil.
- **Título consolidado**: *"Onboarding e KickOff de parceiro tech: contrato → NDA → trial → kickoff"*.

### Cluster 13 — Ambiente trial / homologação / sandbox
- **Base**: `kb-parceiros-martech HOW_TO-1` + `TROUBLESHOOTING-1`, `TROUBLESHOOTING-2` (Compra Rápida).
- **Fundir**: `kb-parceiros-wave6b TROUBLESHOOTING-5` (CRMBonus), `kb-parceiros-wave6c HOW_TO-2` (Eitri — sandbox), `kb-parceiros-wave7a FAQ-3`, `HOW_TO-4` (Fidelizar — loja teste), `kb-parceiros-martech HOW_TO-11` (Revi — credenciais).
- **Agrega**: criação de loja teste, acessos, problemas com CSS/SSL/trial expirado.
- **Título consolidado**: *"Ambiente trial/sandbox/homologação para parceiros tech: criação, acessos, troubleshooting"*.

### Cluster 14 — Suporte oficial e escalation (incluindo Black Friday)
- **Base**: `kb-parceiros-wave5a FAQ-9` (Omni360 — canais oficiais).
- **Fundir**: `kb-parceiros-martech FAQ-1` (Compra Rápida), `kb-parceiros-wave5b HOW_TO-6` (Pagaleve — SLA), `kb-parceiros-wave5c TROUBLESHOOTING-2` (Econverse), `kb-parceiros-wave5a FAQ-6` (A&EIGHT — escalation), `kb-parceiros-martech FAQ-3` (Stayfilm — BF), `kb-parceiros-wave6a BEST_PRACTICES-5` (Biso — protocolo BF), `kb-parceiros-wave6c HOW_TO-5` (Eitri — BF).
- **Agrega**: hierarquia N1→N2, SLA por tier, protocolo Black Friday/Black November.
- **Título consolidado**: *"Canais oficiais de suporte Wake e escalation (incluindo plantão Black Friday)"*.

### Cluster 15 — Webhook e API de carrinho abandonado
- **Base**: `kb-parceiros-tech FAQ-1` (Omnichat — Commerce vs Experience).
- **Fundir**: `kb-parceiros-wave7a FAQ-7` (Fidelizar — webhook), `kb-parceiros-martech HOW_TO-10` (Revi — pixel/tag), `kb-parceiros-martech FAQ-4` (Revi — API checkout), `kb-parceiros-tech FAQ-2` (Omnichat — impacto do checkout novo).
- **Agrega**: diferença Commerce vs. Experience, pixel/tag, eventos disponíveis.
- **Título consolidado**: *"Carrinho abandonado na Wake: Commerce vs. Experience, webhook, pixel, novo checkout"*.

### Cluster 16 — Take-rate Wake em pedidos Anymarket
- **Base**: `kb-parcerias-marketing FAQ-4`.
- **Fundir**: `kb-parceiros-marketplace FAQ-2` (Blueman — TR sem finalização), `kb-parcerias-marketing FAQ-3` (Unlock/Estoca/Afterclick).
- **Agrega**: regras gerais, exceções por falha de finalização, concorrência OMS.
- **Título consolidado**: *"Take rate Wake em pedidos Anymarket: regras, casos especiais e exceções"*.

### Cluster 17 — Conectores de pagamento (homologação e disponibilidade)
- **Base**: `kb-parcerias TRB-008` (Rede) + `TRB-005` (PagBank pós-Storefront).
- **Fundir**: `kb-parceiros-wave5a FAQ-8` (Omni360 — adquirentes), todo o grupo `kb-parceiros-pagamento` (PayPal, MP, Safrapay, Givex, Vindi), `kb-parcerias FAQ-013`, `FAQ-014` (Vindi repasse, Pagar.me migração), `kb-parcerias TRB-001` (conector oculto), `kb-parceiros-wave6c HOW_TO-13` (Tuna — homologação).
- **Agrega**: status por conector, regras de tokenização, repasse, migração Pagar.me.
- **Título consolidado**: *"Adquirentes e gateways no Wake Gateway: homologação, disponibilidade, conectores ocultos, migração"*.

### Cluster 18 — Indicação cruzada entre parceiros
- **Base**: `kb-comercial-cross FAQ-2` (indicar integradores).
- **Fundir**: `kb-parcerias HT-004` (indicação por CS), `kb-parceiros-wave5b FAQ-4` (DRIVEN — CRM), `kb-parcerias FAQ-016` (Vurdere pitch), `kb-parceiros-martech HOW_TO-6` (After Click), `kb-parceiros-tech HOW_TO-1` (Omnichat — co-sell).
- **Agrega**: catálogo, papel de CS, pitch unificado, fluxo de comissionamento cruzado.
- **Título consolidado**: *"Como indicar parceiros tech/agência entre si e para clientes Wake (catálogo + processo)"*.

### Cluster 19 — Selo / logo / branding (Wake ↔ parceiro)
- **Base**: `kb-parceiros-wave5c HOW_TO-3` (Avocado — selos).
- **Fundir**: `kb-parceiros-wave6c FAQ-3` (O2 — rodapé), `kb-parceiros-wave5c BEST_PRACTICES-3` (Wicomm — padronização), `kb-parceiros-tech TROUBLESHOOTING-3` (FG — Elite não aparece no site), `kb-parceiros-tech TROUBLESHOOTING-5` (Leanwork — logo pequeno), `kb-parceiros-wave6c HOW_TO-3` (Eitri — apresentação), `kb-parceiros-martech HOW_TO-4` (After Click — inclusão no site).
- **Agrega**: regras de uso, padrões visuais, listagem oficial, troubleshooting de aparecimento.
- **Título consolidado**: *"Branding Wake↔parceiro: selos, rodapé, listagem de agências, padrões visuais"*.

### Cluster 20 — Co-marketing (PR, blog, podcast, LinkedIn)
- **Base**: `kb-parceiros-martech HOW_TO-8` (Stayfilm — entregáveis).
- **Fundir**: `kb-parceiros-martech HOW_TO-13` (Vurdere — lançamento Elite), `kb-parceiros-wave5b HOW_TO-1` (BonifiQ — PR), `kb-parceiros-wave5b HOW_TO-2` (BonifiQ — estúdio podcast), `kb-parceiros-martech HOW_TO-9` (Stayfilm — podcasts), `kb-parcerias-marketing HOW-1` (PR estruturado), `kb-parceiros-wave6c BEST_PRACTICES-2` (Eitri — PR), `kb-parceiros-martech BEST_PRACTICES-5` (Vurdere — cases públicos), `kb-parceiros-wave6b HOW_TO-5` (Estoca — podcast Summit).
- **Agrega**: entregáveis por tier, plano de lançamento Elite, estúdio interno.
- **Título consolidado**: *"Co-marketing Wake↔parceiro: PR, blog, podcast, painel, LinkedIn"*.

### Cluster 21 — Leads para produtos além de Commerce (OMS, DXP, Creators)
- **Base**: `kb-agencia FAQ-012`.
- **Fundir**: `kb-parceiros-wave5c FAQ-5` (Wicomm — DXP/Commerce), `kb-parceiros-wave5b HOW_TO-5` (DRIVEN — Wake Creators), `kb-parceiros-agencia FAQ-3` (Nairuz — Creators).
- **Agrega**: regras de carrossel quando o produto não é Commerce, elegibilidade por produto.
- **Título consolidado**: *"Distribuição de leads para OMS, DXP e Wake Creators: regras de carrossel e elegibilidade"*.

### Cluster 22 — Wake Experience (CRM/DXP) — demo, integração, posicionamento
- **Base**: `kb-comercial-cross FAQ-9` + `BEST-PRACTICES-5` (upsell).
- **Fundir**: `kb-parceiros-tech FAQ-1`, `FAQ-3` (Omnichat), `kb-parceiros-tech HOW_TO-10` (Eficaz — Experience), `kb-parceiros-wave6b HOW_TO-8` (Konfidency), `kb-parceiros-wave6b HOW_TO-10` (CRMBonus — agenda vendedor), `kb-comercial-cross TROUBLESHOOTING-2`, `TROUBLESHOOTING-5` (setup grátis, tabela de disparos).
- **Agrega**: posicionamento, demo, integração com parceiros de CRM, cuidados comerciais de upsell.
- **Título consolidado**: *"Wake Experience (CRM/DXP): posicionamento, demo, integração, upsell"*.

### Cluster 23 — SDR Wake → parceiros
- **Base**: `kb-parceiros-wave6a FAQ-1` (Lemoon — SDR).
- **Fundir**: `kb-parceiros-marketplace HOW_TO-4` (Duo — qualificação SDR), `kb-agencia BP-006` (agenda SDR), `kb-parceiros-wave5a TROUBLESHOOTING-2` (Dotkom — sem resposta).
- **Agrega**: critério de qualificação, agenda, regras de inativo, follow-up.
- **Título consolidado**: *"SDR Wake → parceiros: qualificação, repasse, regras de inativo"*.

---

## 2. Gaps — Perguntas Sem Resposta

18 lacunas de conteúdo evidenciadas por perguntas recorrentes que nenhum artigo isolado cobre por completo.

### Gap 1 — Tabela de Preços Wake 2026 oficial
- **Aparece em**: `kb-parcerias-wave6c FAQ-4`, `kb-parceiros-wave7a FAQ-8`, `kb-parceiros-agencia FAQ-6` (Wake Up).
- **Por quê hoje falha**: cada artigo dá uma fatia (B2C, Wake Up, add-ons), nenhum apresenta a tabela completa.
- **Sugestão**: *"Tabela de Preços Wake 2026: B2C, Wake Up, B2B, add-ons, kit migração"* — tipo **FAQ + Best Practices**.

### Gap 2 — Battlecard único Wake vs. VTEX / Shopify / Linx / Tray / VTEX Rise
- **Aparece em**: ver Cluster 9.
- **Por quê falha**: fragmentado em 6+ artigos, sem visão lado-a-lado por dimensão (preço, certificação, dependência, headless, marketplaces).
- **Sugestão**: *"Battlecard competitivo Wake"* — tipo **Best Practices**.

### Gap 3 — Catálogo central de adquirentes/gateways compatíveis
- **Aparece em**: `kb-parcerias TRB-008` (Rede), `kb-parcerias TRB-005` (PagBank).
- **Por quê falha**: agência vendeu com Rede sem saber compatibilidade — não há fonte única do estado atual.
- **Sugestão**: *"Catálogo de adquirentes e gateways Wake (status: homologado / em homologação / indisponível por conector)"* — tipo **FAQ**.

### Gap 4 — Matriz definitiva de ERPs × Wake
- **Aparece em**: `kb-parcerias FAQ-012` (mais próximo, mas curto).
- **Por quê falha**: não diferencia nativo vs. via intermediador (Vtrina/Bling/Proex) vs. em homologação vs. sem suporte.
- **Sugestão**: *"Matriz de ERPs × Wake"* — tipo **Best Practices** com tabela.

### Gap 5 — Marketplaces no Anymarket — lista canônica atualizada
- **Aparece em**: `kb-comercial-cross FAQ-3`, `TROUBLESHOOTING-1` (rumor Dafiti), `kb-parcerias-marketing FAQ-3`.
- **Por quê falha**: rumores e desinformação (Dafiti, SHEIN, Privalia) circulam sem fonte oficial.
- **Sugestão**: *"Marketplaces no Anymarket: lista canônica + roadmap"* — tipo **FAQ**.

### Gap 6 — Definição operacional de "indicação qualificada"
- **Aparece em**: `kb-parceiros-wave6b FAQ-4`.
- **Por quê falha**: começa a definir mas não tem checklist BANT, momento do registro nem prazo de validade.
- **Sugestão**: *"Definição operacional de 'indicação qualificada': BANT, registro, validade"* — tipo **Best Practices**.

### Gap 7 — Política completa de conflito de canal
- **Aparece em**: `kb-parceiros-agencia TROUBLESHOOTING-3`, `kb-parcerias-canais FAQ-004`, `TRB-003`, `kb-coordenacao TROUBLESHOOTING-4`.
- **Por quê falha**: cada caso é resolvido pontualmente (CS ofertando serviços, parceiros concorrentes na base, reserva de CNPJ).
- **Sugestão**: *"Política de conflito de canal Wake"* — tipo **Best Practices**.

### Gap 8 — Fluxo financeiro completo (Revenue Share / Partner Fee / RV)
- **Aparece em**: `kb-agencia FAQ-005`, `kb-coordenacao FAQ-3`, `FAQ-4`.
- **Por quê falha**: cobre por pedaços, sem fluxo end-to-end (registro → fatura → NF → repasse → clawback).
- **Sugestão**: *"Fluxo financeiro de remuneração de parceiros"* — tipo **How-To**.

### Gap 9 — Roadmap público de APIs Wake
- **Aparece em**: `kb-parceiros-martech FAQ-5` e diversos TRBs (Vurdere, Revi, Omnichat, Fidelizar+).
- **Por quê falha**: cada parceiro pergunta caso a caso; não há changelog público.
- **Sugestão**: *"Roadmap de APIs públicas Wake: Commerce, Storefront, Experience, Data Wake"* — tipo **FAQ**.

### Gap 10 — Guia oficial de migração de storefront/checkout
- **Aparece em**: `kb-parcerias FAQ-015`, `kb-parceiros-wave7a FAQ-9`.
- **Por quê falha**: fragmentos com números diferentes (kit migração); sem custo por plano nem responsabilidades.
- **Sugestão**: *"Kit de migração para o novo checkout Wake: custos, prazos, papéis"* — tipo **How-To**.

### Gap 11 — Wake OMS vs. concorrentes
- **Aparece em**: `kb-parcerias-marketing FAQ-3`, `kb-parceiros-wave5c FAQ-3`.
- **Por quê falha**: trata Unlock/Estoca/Afterclick como objeções isoladas, sem comparação estruturada.
- **Sugestão**: *"Wake OMS: escopo nativo, concorrentes, quando indicar parceiro"* — tipo **Best Practices**.

### Gap 12 — Matriz "build vs. buy" (nativo Wake × add-ons de parceiros)
- **Aparece em**: `kb-parceiros-wave5c FAQ-1` (conciliação), `FAQ-2` (reviews/Trustvox), `FAQ-3` (OMS).
- **Por quê falha**: cada feature é discutida isoladamente; sem visão consolidada (reviews, OMS, antifraude, fidelidade, frete, conciliação).
- **Sugestão**: *"Matriz Wake nativo × add-ons de parceiros"* — tipo **Best Practices**.

### Gap 13 — Processo formal de Win/Loss com parceiros
- **Aparece em**: `kb-comercial-cross BEST-PRACTICES-7`, `kb-parceiros-marketplace TROUBLESHOOTING-9`, `kb-parceiros-tech BEST_PRACTICES-9`.
- **Por quê falha**: importância citada, processo não formalizado (template, cadência, owner, impacto no Wake Score).
- **Sugestão**: *"Processo Win/Loss com parceiros"* — tipo **How-To**.

### Gap 14 — Cases Wake por vertical
- **Aparece em**: `kb-comercial-cross FAQ-1`, `FAQ-4`, `FAQ-5`, `FAQ-8`, `BEST-PRACTICES-9`, `BEST-PRACTICES-15`.
- **Por quê falha**: cases mencionados em respostas pontuais; sem referência única (cosméticos, moda, autopeças, grocery, B2B, farma).
- **Sugestão**: *"Cases Wake por vertical: referência única para comercial e parcerias"* — tipo **Best Practices**.

### Gap 15 — Calendário oficial de eventos anuais
- **Aparece em**: `kb-parcerias-canais FAQ-008` + dezenas de menções dispersas.
- **Por quê falha**: cada evento (Summit, Trends, Education, Wake on The Road, Partner Day, ECBR, Black Friday) tem sua própria conversa; sem visão de calendário.
- **Sugestão**: *"Calendário oficial de eventos Wake"* — tipo **FAQ**.

### Gap 16 — Política de uso de IA / MCP / Gemini / transcrição
- **Aparece em**: `kb-parceiros-agencia BEST_PRACTICES-11`, `kb-parceiros-wave5a BEST_PRACTICES-7`, `kb-parcerias-canais FAQ-002`, `FAQ-011`, `kb-parceiros-wave6c BEST_PRACTICES-4`.
- **Por quê falha**: cada uso é discutido caso a caso, sem governança formal.
- **Sugestão**: *"Ferramentas de IA aprovadas Wake (MCP, LLM, transcrição, Gemini): uso e governança"* — tipo **Best Practices**.

### Gap 17 — Política de inadimplência e bloqueio de painel
- **Aparece em**: `kb-anymarket-wake TRB-006`, `kb-parceiros-wave5c TROUBLESHOOTING-1`, `kb-coordenacao TROUBLESHOOTING-1`, `TROUBLESHOOTING-5`, `kb-parcerias FAQ-002`.
- **Por quê falha**: tratamento ad-hoc por CS/financeiro; sem gatilhos, prazos e papéis claros.
- **Sugestão**: *"Política de inadimplência: gatilhos, prazos, bloqueio, retomada, papéis"* — tipo **Best Practices**.

### Gap 18 — Manual do parceiro no Salesforce Wake
- **Aparece em**: `kb-parcerias FAQ-004`, `FAQ-005`, `TRB-002`, `kb-parceiros-wave6c TROUBLESHOOTING-2` (Community vazio), `kb-parceiros-tech HOW_TO-5` (notificações), `kb-parceiros-agencia HOW_TO-5` (cadastro sem botão).
- **Por quê falha**: dúvidas estruturais sobre campos, tier, parceiro-origem, Community e bugs ficam dispersas.
- **Sugestão**: *"Manual do parceiro no Salesforce Wake: Lead, OPT, Tier, parceiro-origem, Community"* — tipo **How-To**.

---

## 3. Taxonomia de Tags Normalizada

### 3.1. Mapa de consolidação canônica (40+ regras)

Regras `entrada1`, `entrada2`, ... → **canônica**.

| Entradas existentes | Tag canônica |
|---|---|
| `salesforce`, `SF`, `crm`, `CRM` | **salesforce** |
| `pagamento`, `gateway`, `meio-de-pagamento`, `conector`, `payment` | **pagamento** (mantém `gateway` como subtag) |
| `agência`, `agencia`, `parceiro-agência`, `agência-parceira` | **agencia** |
| `parceiro-tech`, `tech`, `parceiros-tech`, `parceiro tech` | **parceiro-tech** |
| `programa-de-parcerias`, `programa de parcerias`, `programa-parcerias`, `programa parceiros`, `programa Special`, `programa-parceiros` | **programa-parcerias** |
| `wake-score`, `wake score`, `Wake-Score`, `Wake Score` | **wake-score** |
| `wake-summit`, `Wake Summit`, `summit`, `wake summit` | **wake-summit** |
| `partner-day`, `wake-partner-day`, `Wake Partners Day`, `partner day`, `parceiro day` | **wake-partner-day** |
| `wake-day`, `Wake Day`, `wake day`, `wake-education`, `Wake Education` | **wake-education** (com `wake-day` como subtag) |
| `certificação`, `certificacao`, `certification` | **certificacao** |
| `ECBR`, `fórum-ecbr`, `forum-ecbr`, `fórum ECBR`, `forum` | **forum-ecbr** |
| `mercado-livre`, `Mercado Livre`, `ml`, `mercadolivre` | **mercado-livre** |
| `mercado-pago`, `Mercado Pago`, `MP`, `mp` | **mercado-pago** |
| `anymarket`, `Anymarket`, `ANYMARKET` | **anymarket** |
| `migração`, `migracao`, `migration`, `kit-migração` | **migracao** |
| `onboarding`, `onboarding-agência`, `onboarding parceiro`, `onboarding-técnico`, `onboarding-tecnico` | **onboarding** |
| `b2b`, `B2B`, `BtoB` | **b2b** |
| `b2c`, `B2C` | **b2c** |
| `cs`, `CS`, `CSM`, `csm`, `customer-success` | **cs** |
| `nf`, `NF`, `nfe`, `NF-e`, `nota-fiscal`, `notafiscal` | **nota-fiscal** |
| `revenue-share`, `revenue share`, `RV` (de canal), `resultado-variável` | **revenue-share** (com `rv` distinto para pessoas) |
| `lead`, `leads`, `Lead` | **lead** |
| `indicação`, `indicacao`, `referral` | **indicacao** |
| `carrossel`, `carousel` | **carrossel** |
| `checkout`, `checkout-novo`, `redesign-checkout` | **checkout** |
| `storefront`, `Storefront`, `StoreFront`, `storefront 2.0` | **storefront** |
| `api`, `API`, `api-storefront`, `api-hero`, `api-key` | **api** (com subtags `api-storefront`, `api-hero`, `api-key`) |
| `erp`, `ERP` | **erp** |
| `tray`, `Tray`, `tray-commerce` | **tray** |
| `vtex`, `VTEX`, `vtex rise` | **vtex** (com `vtex-rise` como subtag) |
| `shopify`, `Shopify` | **shopify** |
| `linx`, `Linx` | **linx** |
| `inadimplência`, `inadimplencia`, `atraso`, `cobrança` | **inadimplencia** |
| `homologação`, `homologacao` | **homologacao** |
| `precificação`, `precificacao`, `pricing`, `tabela-de-preços` | **precificacao** |
| `marketplace`, `marketplaces`, `Marketplace` | **marketplace** |
| `black-friday`, `Black Friday`, `BF`, `black-november`, `Black November` | **black-friday** |
| `evento`, `eventos`, `evento-presencial`, `evento-experiência` | **evento** |
| `webhook`, `Webhook` | **webhook** |
| `whatsapp`, `WhatsApp`, `WhatsApp-grupo` | **whatsapp** |
| `community`, `Community`, `comunidade-salesforce` | **sf-community** |
| `concorrência`, `concorrencia`, `competição`, `competicao`, `battlecard` | **concorrencia** |
| `co-sell`, `cosell`, `co-venda`, `co venda`, `cross-account` | **co-sell** |
| `co-marketing`, `co marketing`, `comarketing` | **co-marketing** |
| `pix`, `PIX`, `Pix` | **pix** |
| `boleto`, `Boleto` | **boleto** |
| `cartão`, `cartao`, `cartão-de-crédito` | **cartao** |

### 3.2. Hierarquia recomendada (5 níveis)

```
platform/
  wake, wake-commerce, wake-experience, wake-up, wake-gateway,
  anymarket, storefront, checkout, oms, dxp, data-wake,
  wake-creators, wake-dev-toolkit

integration/
  erp/        totvs, winthor, linx, alpha7, procfit, vetor, everest, tiny
  marketplace/ mercado-livre, magalu, shein, dafiti, privalia,
               riachuelo, netshoes, epoca-cosmeticos
  gateway-payment/ pagar-me, paypal, mercado-pago, safrapay, vindi,
                   givex, appmax, pagaleve, click-to-pay, rede, pagbank, tuna
  logistics/  estoca, correios, multicd, crossdocking
  martech/    omnichat, vurdere, revi, after-click, compra-rapida, biso,
              btg360, bonifiq, crmbonus, konfidency, eitri, hero-seguros,
              serasa-pass, metakosmos, stayfilm

commercial/
  planos/      essential, special, elite, wake-up
  financeiro/  revenue-share, partner-fee, rv, comissao, clawback,
               nota-fiscal, inadimplencia, faturamento, faturamento-retroativo
  propostas/   bant, proposta, briefing, war-room, forecast,
               pipeline, commit, best-case
  canais/      agencia, parceiro-tech, parceiro-marketplace,
               parceiro-pagamento, parceiro-martech, sdr, canal,
               indicacao, lead
  concorrencia/ vtex, vtex-rise, shopify, linx, tray, mercado-shops

process/
  onboarding/  certificacao, wake-score, kickoff, trial, sandbox,
               homologacao, kit-migracao
  suporte/     ticket, escalation, sla, plantao, black-friday,
               n1, n2, status-page
  eventos/     wake-summit, wake-partner-day, wake-education, wake-day,
               wake-trends, wake-on-the-road, forum-ecbr, podcast,
               jantar, patrocinio, cota, ingresso, brindes
  governanca/  contrato, nda, jurídico, lgpd, compliance, cnpj, cnae,
               reserva-cnpj, conflito-canal, prm

topics/
  funcional/    estoque, pedido, status-pedido, cancelado,
                carrinho-abandonado, cupom, promocao, frete-gratis, sku,
                anuncios, multistore, multicd, split-pedido
  modelos/      b2b, b2c, omnichannel, headless, multistore,
                embedded-finance, white-label, marketplace-out, marketplace-in
  tecnologia/   api, webhook, sdk, script, datalayer, gtm, css, ssl,
                biometria, qrcode, ip, autenticacao, etl, middleware, deploy
  ia/           ia, llm, mcp, gemini, transcricao, automacao
```

### 3.3. 15 artigos com problemas de tagging

1. `kb-comercial-cross FAQ-7` (Clear Sale) — só tem `clear-sale`, `parceiro`; faltam **antifraude** (canônica) e **pagamento** (parent).
2. `kb-comercial-cross FAQ-9` (demo Wake Experience) — usa `experience`, `demo`; falta **wake-experience** canônico.
3. `kb-comercial-cross HOW-TO-1` (War Room) — faltam **pipeline**, **forecast**.
4. `kb-parceiros-tech TROUBLESHOOTING-1` (UncleK) — typo "uncleck" nas tags; falta **integracao-tecnica**.
5. `kb-parceiros-martech TROUBLESHOOTING-6` (Vurdere admin) — tag quebrada com aspas e espaços (`'`,`'`).
6. `kb-parceiros-marketplace FAQ-1` (Blueman — etiqueta Anymarket) — falta **marketplace** como parent.
7. `kb-parceiros-wave5b FAQ-3` (Bio Diego Santos) — só tem tag de nome próprio; faltam **marketing**, **branding**.
8. `kb-parceiros-wave6a TROUBLESHOOTING-1` (OPT TNT) — `TNT` ambíguo; precisa de **glossario** e **comunicacao-interna**.
9. `kb-parceiros-wave6b FAQ-3` (Bio Diego Santos duplicado em outro grupo) — sem cross-tag para a duplicata em wave5b.
10. `kb-parcerias FAQ-017` (PRM ambiente) — só tags genéricas; faltam **governanca**, **salesforce**.
11. `kb-parceiros-agencia FAQ-5` (Catus — TOTVS Moda) — faltam **integracao** (parent) e **vertical-moda**.
12. `kb-parceiros-wave5c FAQ-2` (Avocado — Trustvox) — só `trust`; faltam **parceiro**, **martech**, **review**.
13. `kb-comercial-cross TROUBLESHOOTING-1` (rumor Dafiti) — faltam **objecao-comercial**, **enablement**.
14. `kb-parceiros-wave6c FAQ-7` (Tuna WhatsApp) — faltam **governanca**, **canal-comunicacao**.
15. `kb-anymarket-wake TRB-005` (AWS) — faltam **incidente**, **infra-cliente**.

---

## 4. Oportunidades Estratégicas

### 4.1. Top 5 temas mais críticos (frequência × impacto de negócio)

1. **Salesforce e carrossel de leads** — 25+ artigos sobre regras, bugs, atribuição, tier, parceiro-origem e Community vazio. Bloqueia comissionamento e gera disputas. Risco direto à confiança da rede de parceiros.
2. **Wake vs. concorrentes (VTEX, Shopify, Linx, Tray, VTEX Rise)** — 10+ artigos. Objeção #1 em deals (preço, dependência de agência, exigência de 3 projetos, certificação).
3. **Programa de parcerias (planos, certificação, Wake Score, revenue share)** — 30+ artigos. Foundation do negócio de canal; toda dúvida nova de parceiro toca pelo menos um destes eixos.
4. **Integrações (ERP, marketplace, gateway, OMS)** — 50+ artigos. Deal-blocker recorrente: *"Wake integra com X?"*.
5. **Anymarket / marketplaces** — 25+ artigos, maioria de troubleshooting (Privalia, SHEIN, Mercado Livre duplicado, Magalu V2). Maior fonte de chamados técnicos do KB.

### 4.2. Conteúdo ausente de alto valor (5 artigos novos prioritários)

1. **Battlecard único Wake vs. VTEX / Shopify / Linx / Tray / VTEX Rise** — substitui 10+ fragmentos espalhados (Gap 2 / Cluster 9).
2. **Matriz "build vs. buy"**: nativo Wake × add-ons de parceiros (OMS, reviews, fidelidade, antifraude, frete, conciliação) — Gap 12.
3. **Catálogo canônico de integrações** com 3 tabelas: ERPs × Wake, Marketplaces × Anymarket, Gateways × Wake Gateway — Gaps 3, 4, 5.
4. **Manual completo do Salesforce para Parcerias** (Lead → OPT → Tier → parceiro-origem → Community → carrossel + FAQ de bugs conhecidos) — Gap 18.
5. **Playbook de Wake Score e Certificação**: pré-requisitos para receber leads, trilhas técnicas vs. comerciais, regras de regressão, penalizações — apoia Clusters 1 e 23.

### 4.3. Padrões de parceria bem-sucedidos (extraídos do corpus)

1. **Co-sell / Co-Build ritualizado com BANT estruturado** — Catus, Série//A, Nairuz, Wake & Leanwork. Agências com ritual semanal convertem mais (kb-agencia BP-006; kb-parceiros-agencia HOW_TO-2).
2. **Presença física no Partner Day (LWSA, quartas-feiras)** — DRIVEN, Eitri, Vtrina, BonifiQ, Omnichat. Parceiros com presença recorrente aceleram pipeline e ganham informação privilegiada (`kb-coordenacao BEST_PRACTICES-8`).
3. **Overlap de carteira como acelerador** — Revi e Stayfilm começam pela identificação de clientes em comum com a Wake e atingem co-sell mais rapidamente (`kb-parceiros-martech HOW_TO-6`).
4. **Co-marketing coordenado com plano de lançamento Elite** — Vurdere, Stayfilm, BonifiQ usam o plano oficial (PR + blog + podcast + LinkedIn) e capturam mais leads inbound (`kb-parceiros-martech HOW_TO-13`).
5. **Certificação técnica + comercial separadas e obrigatórias antes de receber leads** — agências certificadas têm Score maior e melhor distribuição (`kb-agencia BP-004`, `BP-013`; `kb-parceiros-wave5b HOW_TO-4`).

### 4.4. Riscos e pontos cegos sistêmicos (7)

1. **Conflito de canal recorrente** — CS Wake ofertando serviços de agência, parceiros concorrentes prospectando a base, reserva de CNPJ informal. Falta política única e processo de denúncia/escalation (Gap 7).
2. **Inadimplência sem protocolo formal** — StayFilm, Vtrina, Inbrands/Anymarket, Econverse. O protocolo de contenção (`kb-coordenacao TROUBLESHOOTING-5`) deveria virar política oficial (Gap 17).
3. **Bugs recorrentes no Salesforce afetando comissionamento** — Tier errado, parceiro-origem perdido, Community vazia, carrossel disparando errado, dashboards inconsistentes. Erosão sistêmica de confiança da rede.
4. **Concorrentes apresentando Wake errado em grupos públicos** — `kb-parcerias-marketing TRBL-5`, `kb-parcerias-marketing TRBL-4` (Serasa). Falta governança de comunicação de parceiros sobre a Wake.
5. **Roadmap de APIs e novo checkout sem changelog público** — TRBs recorrentes (Omnichat, Revi, Vurdere, Hero, Serasa Pass) por mudanças não comunicadas. Sugere changelog público para parceiros tech (Gap 9).
6. **Bugs de canal no Anymarket sem fluxo N1→N2 documentado** — Privalia/Inbrands, SHEIN, Mercado Livre duplicado, Magalu V2. Seller final não sabe quem aciona.
7. **Dependência do WhatsApp como sistema de registro** — todo este KB existe porque o conhecimento crítico vive em chats. Risco operacional alto: churn de pessoas-chave (Julie, Diego, Rafa) significa perda direta de contexto. Indica necessidade urgente de um KB oficial versionado e de tornar o pipeline atual (extração → indexação) parte do processo permanente, não um one-off.

---

*Fim do relatório. Referências cruzadas com `kb-master-index.md`. Próximo passo sugerido: priorizar os 5 artigos novos da seção 4.2 e atacar os clusters 1, 3, 4, 7, 9 (maior frequência × maior impacto).*
