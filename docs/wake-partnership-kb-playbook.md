# Wake Partnership KB Playbook

**Audience**: the `wake-partnership-kb-writer` agent and any human author producing Wake partnership KB content.

**Purpose**: domain knowledge companion. The agent definition (`~/.claude/agents/wake-partnership-kb-writer.md`) contains the behavioral rules. This file contains the facts, taxonomies, decision trees, and SQL queries the agent needs to ground its work.

**Version**: 1.0 — 2026-06

---

## 1. Authoritative source hierarchy

When facts conflict, resolve in this order:

1. **`store/messages.db`** — raw WhatsApp messages. Ground truth.
2. **`docs/kb-priority-*.md`** — canonical synthesized articles. Useful when the DB has too much noise to read directly, but always verify the underlying claim with a DB query before quoting.
3. **`docs/kb-delivery-batch-*.md`** — extracted batch articles. Navigation index. Tells you where to look, not what is true.
4. **`docs/kb-opus-analysis.md`** — strategic analysis from a prior Opus pass. Useful for context on duplicate clusters and known gaps. Not authoritative on facts.

Never cite anything below level 1 as the only source for a claim that goes into a published article. Levels 2–4 are navigation tools.

---

## 2. Canonical partner taxonomy

Every partner mentioned in any article must fall into exactly one of these categories. The categories are defined by what the partner *does for the customer*, not by what they sell.

| Category | Definition | Examples (verify in DB before citing) |
|---|---|---|
| `agência` | Implements Wake stores: setup, theme, integrations, ongoing operation. Owns the customer relationship on the implementation side. | Eficaz, Nairuz, Catus, FZ Commerce, Wicomm, Avocado Tech, Série//A, Ge Commerce, Econverse, Lemoon, DRIVEN, FG, Duo, B8One |
| `parceiro-tech` | Adds functional capability via API/integration: search, reviews, personalization, conversational, identity. Doesn't replace the agency. | Biso (search), Vurdere (reviews), Konfidency (reviews), Omnichat (conversational), Stayfilm (video), Metakosmos (AI try-on), HERO (insurance), Serasa Pass (identity) |
| `parceiro-marketplace` | Hub or connector for selling on marketplaces. **Wake's exclusive marketplace hub is Anymarket.** | Anymarket (hub), Blueman (specific marketplace) |
| `parceiro-pagamento` | Gateway, acquirer, wallet, BNPL, or anti-fraud. | Pagar.me (Wake's native acquirer + gateway), Vindi, PayPal (BCDC + Braintree), Mercado Pago, Pagaleve, Safrapay, Mastercard, Givex, Appmax, Tuna |
| `parceiro-martech` | Marketing, loyalty, retention, CRM. | CRMBonus, BonifiQ, Fidelizar+, Revi, Compra Rápida, After Click |
| `parceiro-logistica` | Shipping, fulfillment, marketplace logistics. | Loggi, Melhor Envio, Estoca |
| `integrador-erp` | iPaaS or middleware that connects an external ERP to Wake. Most ERPs require one. | Argento Digital, Integra.do (Kore Hub), EBI, Floui, Iota.Hub, Intergate, Jitterbit, LinkToU, Vtrina, Leanwork |

**Non-negotiables (verify in DB before contradicting):**

- **Native ERPs in Wake**: only **Bling** and **Tiny**. Every other ERP requires a `integrador-erp` partner.
- **Anymarket is Wake's exclusive marketplace hub.** No other generic marketplace hub is supported.
- **Pagar.me is Wake's native gateway and acquirer.** "Rede" can mean either an acquirer (supported via Pagar.me) or a gateway (not the same thing) — always clarify which one the customer is asking about.

---

## 3. Partner selection decision tree

When writing an article that helps a Solution Architect or Sales executive choose a partner, use this skeleton. Adapt to the specific question.

### 3.1 Customer needs an ERP integration

```
Cliente usa qual ERP?
├── Bling ou Tiny → integração nativa, sem parceiro
├── TOTVS Protheus → Integra.do, EBI, Intergate, Floui
├── TOTVS Winthor → Vtrina, EBI, Integra.do
├── TOTVS RM → Integra.do
├── SAP Business One → Intergate, LinkToU
├── SAP S/4HANA → Floui, Iota.Hub
├── Linx Microvix → Argento Digital, EBI
├── Omie → Argento Digital, Integra.do
├── Sankhya → Integra.do, Floui
├── Sistema legado / customizado → Jitterbit, Iota.Hub, Leanwork
└── ERP não listado → verificar com Partner Ops antes de prometer
```

Verify the specific partner recommendation against the DB before publishing — partner capabilities evolve.

### 3.2 Customer needs marketplace presence

```
Cliente quer vender em marketplace?
├── Sim → Anymarket (hub exclusivo)
│   ├── Canal específico (Mercado Livre, Magalu, Shopee) → suportado via Anymarket
│   ├── Canal nicho (Privalia, SHEIN, Riachuelo) → verificar status atual no DB
│   └── Canal não suportado → escalar com Anymarket CS antes de prometer
└── Não → não recomendar parceiro de marketplace
```

### 3.3 Customer needs payment beyond Pagar.me native

```
Cliente quer:
├── Crossborder + Apple Pay + Google Pay → PayPal Braintree
├── PayPal wallet + checkout transparente → PayPal BCDC
├── BNPL (compre agora, pague depois) → Pagaleve
├── 18x sem juros via Mercado Pago → Mercado Pago
├── Gestão de assinaturas → Vindi
├── Antifraude externo (não o de Pagar.me) → integração via Pagar.me GTW, não PSP
└── Adquirente alternativa (Rede como acquirer) → suportado via Pagar.me
```

### 3.4 Customer needs reviews / social proof

```
Cliente quer:
├── Reviews + perguntas/respostas + UGC → Vurdere (Elite)
├── Reviews de baixo ticket / SMB → Konfidency
└── Cases públicos + LinkedIn co-marketing → ambos têm programa de co-marketing
```

These trees are skeletons. The agent must verify the partner is still active and the recommendation still applies by querying the DB before citing.

---

## 4. Escalation matrix

When writing for `audience-cs` or `audience-pam`, refer to this matrix. Verify specific contact names against the DB — people change roles.

| Situation | First point | Second point | Channel |
|---|---|---|---|
| Bug em produto Wake | CS Wake (account-specific) | Suporte Wake (`suporte@wake.tech`, portal `atendimento.wake.tech`) | Email + portal |
| Problema com Anymarket | CS Wake | Anymarket support (`suporte.anymarket.com.br`) | Ticket portal |
| Dúvida operacional de parceria | PAM responsável | Partner Ops | WhatsApp grupo da parceria + Slack |
| Conflito comercial entre parceiros | PAM responsável | Coordenação Parcerias | Slack |
| Black Friday / pico operacional | CS Wake | Plantão BF (canal específico, ativado sazonalmente) | Slack + WhatsApp |
| Lead/carrossel com problema no Salesforce | Sales Ops | Coordenação Parcerias | Slack |
| Pagamento de Revenue Share | Coordenação Parcerias | Time OTC (Order-to-Cash) | Email |
| Onboarding novo parceiro | Partner Development | PAM (após assinatura) | Email + WhatsApp |

**SLA reality check**: as of the source corpus, **Wake does not publish a formal partner-facing SLA**. If an article needs to state a response time, either find a concrete commitment in the messages or mark the claim `⚠ inferido` with the actual quoted phrase from the message.

---

## 5. Canonical source files by topic

When the agent needs to start from a known-good navigation source before going to the DB:

| Topic | Canonical synthesis file | Authoritative DB groups (filter by `c.name LIKE`) |
|---|---|---|
| Partner ecosystem overview | `docs/kb-priority-parceiros-wake.md` | `WAKE \| Time Parcerias`, `Parcerias \| Canais`, `Parcerias & Marketing Wake` |
| ERP integration | `docs/kb-priority-erp-wake.md` | `Comercial Cross WAKE`, `WAKE \| Time Parcerias` |
| Support and docs | `docs/kb-priority-suporte-docs.md` | `WAKE \| Time Parcerias`, `WAKE \| Coordenação`, `ANYMARKET & WAKE` |
| Dev environment | `docs/kb-priority-ambiente-dev.md` | `Projeto Hero <> Leanwork <> Wake`, `Projeto Serasa Pass <> Leanwork <> Wake`, `Wake & Biso` |
| Agency program | `docs/kb-delivery-batch-01.md` (kb-agencia section) | `WAKE \| Time Agência`, `Parcerias \| Canais` |
| Anymarket operations | `docs/kb-delivery-batch-01.md` + `02.md` (anymarket sections) | `ANYMARKET & WAKE (TÁTICO / OPERACIONAL)` |
| PayPal | `docs/kb-delivery-batch-07.md` (pagamento section) | `WAKE & PAYPAL` |
| Mercado Pago | `docs/kb-delivery-batch-07.md` | `Mercado Pago + Wake` |
| Wake vs VTEX positioning | `docs/kb-delivery-batch-01.md` (BPR-008, BPR-013 of kb-agencia) | `WAKE \| Time Agência`, `Comercial Cross WAKE` |

---

## 6. Top Wake-related groups in the message store

For reference when crafting queries. Counts are group text messages only.

| Group name | Approx text msg count | Primary topic |
|---|---|---|
| WAKE \| Time Parcerias | 13,803 | Internal partnerships team — strategy, programs, escalation |
| WAKE \| Time Agência | 4,746 | Internal agency program — Salesforce, carrossel, tier rules |
| Comercial Cross WAKE | 3,337 | Cross-team commercial — ERP selection, deals, pricing |
| Projeto Hero <> Leanwork <> Wake | 1,962 | HERO Seguros integration project |
| ANYMARKET & WAKE (TÁTICO / OPERACIONAL) | 1,954 | Anymarket cross-support — tickets, marketplace issues |
| Parcerias & Marketing Wake | 1,684 | Co-marketing, content, events |
| Wake <> Catus | 1,378 | Catus agency operational |
| Wake <> Nairuz | 1,371 | Nairuz agency operational |
| Projeto Serasa Pass <> Leanwork <> Wake | 1,369 | Serasa Pass integration project |
| Wake + Marketplace Blueman | 1,249 | Blueman marketplace integration |
| WAKE \| Coordenação | 672 | Coordenação Parcerias — internal coordination |
| WAKE & PAYPAL | (see DB) | PayPal partner relationship |

---

## 7. Ground-truth query cookbook

Use these patterns when verifying claims. All queries assume the agent runs from the repo root.

### 7.1 Find every mention of a term in group messages

```bash
sqlite3 store/messages.db "
SELECT c.name, datetime(m.timestamp) AS ts, substr(m.content, 1, 400)
FROM messages m JOIN chats c ON c.jid = m.chat_jid
WHERE c.jid LIKE '%@g.us'
  AND (m.media_type IS NULL OR m.media_type = '')
  AND m.content LIKE '%TERM%'
ORDER BY m.timestamp DESC
LIMIT 50;
"
```

### 7.2 Find mentions in a specific group

```bash
sqlite3 store/messages.db "
SELECT datetime(m.timestamp) AS ts, substr(m.content, 1, 400)
FROM messages m JOIN chats c ON c.jid = m.chat_jid
WHERE c.name = 'WAKE | Time Agência'
  AND (m.media_type IS NULL OR m.media_type = '')
  AND m.content LIKE '%TERM%'
ORDER BY m.timestamp DESC
LIMIT 50;
"
```

### 7.3 Verify a partner is associated with an ERP

```bash
sqlite3 store/messages.db "
SELECT c.name, datetime(m.timestamp), substr(m.content, 1, 400)
FROM messages m JOIN chats c ON c.jid = m.chat_jid
WHERE c.jid LIKE '%@g.us'
  AND (m.media_type IS NULL OR m.media_type = '')
  AND m.content LIKE '%PARTNER%'
  AND m.content LIKE '%ERP%'
ORDER BY m.timestamp DESC
LIMIT 30;
"
```

### 7.4 Find pricing mentions for a product or service

```bash
sqlite3 store/messages.db "
SELECT c.name, datetime(m.timestamp), substr(m.content, 1, 400)
FROM messages m JOIN chats c ON c.jid = m.chat_jid
WHERE c.jid LIKE '%@g.us'
  AND (m.media_type IS NULL OR m.media_type = '')
  AND m.content LIKE '%PRODUCT_NAME%'
  AND (m.content LIKE '%R\$%' OR m.content LIKE '%mensalidade%' OR m.content LIKE '%setup%' OR m.content LIKE '%fee%')
ORDER BY m.timestamp DESC
LIMIT 30;
"
```

### 7.5 Find SLA / response-time mentions

```bash
sqlite3 store/messages.db "
SELECT c.name, datetime(m.timestamp), substr(m.content, 1, 400)
FROM messages m JOIN chats c ON c.jid = m.chat_jid
WHERE c.jid LIKE '%@g.us'
  AND (m.media_type IS NULL OR m.media_type = '')
  AND (m.content LIKE '%SLA%' OR m.content LIKE '%prazo de resposta%' OR m.content LIKE '%tempo de resposta%' OR m.content LIKE '%resposta em%')
ORDER BY m.timestamp DESC
LIMIT 30;
"
```

### 7.6 Find competitive (VTEX/Shopify/Linx/Tray) mentions

```bash
sqlite3 store/messages.db "
SELECT c.name, datetime(m.timestamp), substr(m.content, 1, 400)
FROM messages m JOIN chats c ON c.jid = m.chat_jid
WHERE c.jid LIKE '%@g.us'
  AND (m.media_type IS NULL OR m.media_type = '')
  AND (m.content LIKE '%VTEX%' OR m.content LIKE '%Shopify%' OR m.content LIKE '%Linx Commerce%' OR m.content LIKE '%Tray%')
ORDER BY m.timestamp DESC
LIMIT 50;
"
```

### 7.7 Schema reference

```sql
chats (jid TEXT PK, name TEXT, last_message_time TIMESTAMP)

messages (
  id TEXT, chat_jid TEXT,
  sender TEXT, content TEXT, timestamp TIMESTAMP,
  is_from_me BOOLEAN, media_type TEXT,
  filename TEXT, url TEXT, direct_path TEXT,
  mime_type TEXT, media_key BLOB, file_sha256 BLOB, file_enc_sha256 BLOB,
  file_length INTEGER, local_path TEXT, downloaded_at TIMESTAMP,
  PRIMARY KEY (id, chat_jid)
)
```

**Critical filter**: text messages have `media_type IS NULL OR media_type = ''`, **not** `media_type = 'text'`. Group chats have `chat_jid LIKE '%@g.us'`. 1:1 chats end in `@s.whatsapp.net` — skip these for partnership KB work unless explicitly told otherwise.

---

## 8. Anti-pattern catalog (domain-specific)

Things the corpus has seen claimed that are wrong, misleading, or context-dependent:

- **"Wake integra nativamente com [ERP X]"** — almost always wrong. Only Bling and Tiny. Verify before repeating.
- **"O SLA da Wake é [N horas]"** — Wake does not publish a formal partner SLA. Quote the specific commitment from the message, mark `⚠ inferido` otherwise.
- **"Rede é suportada"** — ambiguous. Rede as an acquirer is supported via Pagar.me. Rede as a standalone gateway is not the same conversation. Always disambiguate.
- **"Anymarket suporta [canal Y]"** — channel availability changes frequently (Privalia, SHEIN, Riachuelo had status changes in 2025–2026). Verify against the most recent message.
- **"O parceiro [X] é Elite"** — tier changes. Verify against recent Time Parcerias messages.
- **"Existe sandbox gratuito"** — there is no free sandbox. Any test environment requires an active contract. State this explicitly when the topic comes up.
- **"O Wake Score funciona assim"** — the scoring rules evolved across 2025–2026. Quote the most recent message that defines the rules, and timestamp it.
- **"O parceiro X tem desconto"** — pricing and discounts are confidential and partner-specific. Never publish without explicit confirmation.

---

## 9. PII and confidentiality guardrails

The agent must strip the following from any article produced:

- Personal phone numbers of Wake employees and partner contacts — replace with `[contato interno]` or role/title.
- CPF, RG, personal email addresses (gmail, hotmail, etc.) — remove.
- Specific deal values for named customers — remove or anonymize.
- Internal Slack channel mentions — keep if they describe a process, remove if they expose private team-only spaces.

JID metadata is automatically stripped by the normalization tooling; the agent should not include `JID` rows in metadata tables.

---

## 10. Version log

- 1.0 (2026-06) — initial playbook seeded from the priority articles and Opus analysis. Agent definition: `~/.claude/agents/wake-partnership-kb-writer.md`. Format spec: `docs/kb-article-spec.md`.
