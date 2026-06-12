# Wake KB — Article Format Specification

**Version**: 1.0  
**Date**: 2026-06

This document defines the canonical format for all Wake Knowledge Base articles. Any article that deviates from this spec must be normalized before import into the target system.

---

## 1. Article Types

Four types, lowercase, used exactly as written:

| Type | Use when |
|---|---|
| `faq` | A clear question with a clear answer |
| `troubleshooting` | A problem with symptoms, cause, and resolution |
| `how_to` | A numbered step-by-step process |
| `best_practices` | A recommendation with rationale |

---

## 2. Article ID Format

Each article has a short prefix + zero-padded 3-digit number + em dash + title.

| Type | Canonical Prefix | Example |
|---|---|---|
| `faq` | `FAQ` | `FAQ-001 — Como registrar um lead no Salesforce` |
| `troubleshooting` | `TRB` | `TRB-003 — Carrossel disparando em renovações` |
| `how_to` | `HTO` | `HTO-007 — Como fazer onboarding de nova agência` |
| `best_practices` | `BPR` | `BPR-012 — Nunca vender Wake sem agência definida` |

**Rules:**
- Numbers are zero-padded to 3 digits: `001`, `012`, `099`
- IDs are unique within a file (not globally)
- Separator between number and title is ` — ` (space, em dash, space)
- Title is sentence case (first word capitalized, rest lowercase except proper nouns)

**Do NOT use:** `FAQ-1`, `BEST_PRACTICES-5`, `BP-003`, `HT-001`, `HOW_TO-6`, `HOW-TO-4`, `TRBL-2`, `TROUBLESHOOTING-7` — these are legacy formats from the initial extraction.

---

## 3. Metadata Table

Every article must have this table immediately after the `###` heading, before the body:

```markdown
| Campo | Valor |
|---|---|
| **Grupo/Chat** | Nome do grupo de origem |
| **Data** | YYYY-MM-DD |
| **Categoria** | `faq` |
| **Tags** | `tag-1`, `tag-2`, `tag-3` |
```

**Rules:**
- `Grupo/Chat`: name of the WhatsApp group the content came from, exactly as stored
- `Data`: ISO 8601 date (`YYYY-MM-DD`) of the original conversation
- `Categoria`: one of the four types, in backticks, lowercase
- `Tags`: 3–7 tags, comma-separated, each in backticks, lowercase, hyphenated (`revenue-share`, not `Revenue Share`)
- **Omit the `JID` field** — it is internal metadata, not useful in the target system

---

## 4. Body Format by Type

### 4.1 FAQ

```markdown
### FAQ-001 — [Title]

| Campo | Valor |
|---|---|
| **Grupo/Chat** | ... |
| **Data** | YYYY-MM-DD |
| **Categoria** | `faq` |
| **Tags** | `tag-1`, `tag-2` |

**Pergunta:** [The question as it appeared or was implied in the conversation.]

**Resposta:** [The answer. Direct, factual, in the same language as the source. No padding.]
```

### 4.2 Troubleshooting

```markdown
### TRB-001 — [Title]

| Campo | Valor |
|---|---|
| **Grupo/Chat** | ... |
| **Data** | YYYY-MM-DD |
| **Categoria** | `troubleshooting` |
| **Tags** | `tag-1`, `tag-2` |

**Sintomas:** [What the user observes. Observable behavior only.]

**Causa raiz:** [Why it happens.]

**Resolução:** [Steps taken or recommended to fix. Numbered if multiple steps.]

**Workaround:** [Optional. Only if a temporary fix exists while root cause is unresolved.]
```

### 4.3 How-To

```markdown
### HTO-001 — [Title]

| Campo | Valor |
|---|---|
| **Grupo/Chat** | ... |
| **Data** | YYYY-MM-DD |
| **Categoria** | `how_to` |
| **Tags** | `tag-1`, `tag-2` |

1. [First step. Imperative verb. One action per step.]
2. [Second step.]
3. [Third step.]

**Observação:** [Optional note for edge cases or gotchas.]
```

### 4.4 Best Practices

```markdown
### BPR-001 — [Title]

| Campo | Valor |
|---|---|
| **Grupo/Chat** | ... |
| **Data** | YYYY-MM-DD |
| **Categoria** | `best_practices` |
| **Tags** | `tag-1`, `tag-2` |

**Recomendação:** [The practice. One clear sentence or short paragraph.]

**Rationale:** [Why. The business reason, incident that motivated it, or risk being avoided.]
```

---

## 5. File Structure

Each delivery file groups articles by type in this order:

```markdown
# KB — [Topic or Group Name]

## FAQ

### FAQ-001 — ...
### FAQ-002 — ...

---

## Troubleshooting

### TRB-001 — ...

---

## How-To

### HTO-001 — ...

---

## Best Practices

### BPR-001 — ...
```

**Rules:**
- Section headers (`## FAQ`, etc.) are present only if the file contains articles of that type
- Articles within each section are sorted by ID number ascending
- A `---` separator appears between each article
- No separator after the last article in a section

---

## 6. Tag Taxonomy

The previous taxonomy grew to 913 unique tags across 675 articles — unsearchable and unmaintainable. A first pass cut that to 66 across five facets; this revision tightens it further to a **hard cap of 20 tags total** across two facets. Tag bloat is the enemy; the article body is the right place for everything that doesn't fit a canonical facet.

### Principles

- **Two facets only.** Every tag belongs to exactly one of two facets: Audience or Topic. A reader filters by facet, not by guessing tag names.
- **Hard cap: 20 tags.** The entire canonical vocabulary is 20 tags. Adding a 21st requires retiring an existing one.
- **Rule of 2–4.** Every article carries **2 to 4 tags total**: 1+ audience (required) + exactly 1 topic (required). Additional audience tags only when the article genuinely serves multiple readers in the same shape.
- **No synonyms.** Every concept has exactly one canonical tag. The synonym map in `docs/.tag-synonym-map.json` enforces this on import.
- **No partner names, no product names, no competitor names as tags.** Partner identity, the specific Wake product (Wake Commerce, Wake Up, Anymarket, Pagar.me, etc.), and the named competitor (VTEX, Shopify, Linx, Tray) all live in the article body and `## Fontes`. They are never tags.
- **No long-tail tags.** Client names, ticket IDs, bug IDs, single-use descriptors, dates, and ad-hoc qualifiers go in the body. If a candidate tag does not appear in the canonical list below, it does not get added — find the right canonical tag or expand the body instead.

The total canonical vocabulary across Facets 1–2 is **15 tags** (4 audience + 11 topic), leaving 5 slots of headroom under the 20-tag cap.

---

### Facet 1 — Audience (REQUIRED, 1+)

Every article carries at least one. An article without a named audience does not get written.

| Tag | Reader profile |
|---|---|
| `audience-pam` | Partner Account Manager — rules, tiers, contracts, lead flow |
| `audience-solutions` | Solution Architect — partner/ERP selection, integration topology |
| `audience-sales` | Sales executive — battlecards, objection handlers, deal unlocks |
| `audience-cs` | Customer Success — escalation paths, N1→N2 boundaries, partner contacts |

---

### Facet 2 — Topic (REQUIRED, exactly 1)

The operational subject of the article. Choose the single most representative topic — do not stack topics. Each topic absorbs several legacy subcategories; the article body carries the specificity.

| Tag | Use when the article is about… | Absorbs |
|---|---|---|
| `onboarding` | Bringing a new partner, agency, or integration online — first contact through go-live | `cadastro-parceiro`, `kickoff`, `setup-inicial` |
| `programa-parceiros` | Program structure, tiers (Elite/Essential/Special), Wake Score, certificação, governance, PRM, benefits | `certificação`, `tiering`, `wake-score`, `prm` |
| `contrato-comercial` | Anything paper-or-money: contracts, NDAs, MSAs, exclusivity, revenue share, comissionamento, billing/NF-e, pricing, mensalidade, setup fees, discount policy | `contrato`, `revenue-share`, `faturamento`, `precificação` |
| `lead-pipeline` | Lead registration/distribution/attribution, Salesforce carrossel, opportunity stages, forecast, co-sell, conflito-canal | `lead`, `carrossel`, `pipeline`, `indicação` |
| `proposta` | Proposal, cotação, negotiation, deal scope, closing, objection handling at the deal level | `cotação`, `closing`, `negociação` |
| `suporte-escalação` | Support tickets, bug intake, SLA, status pages, atendimento, escalation paths, N1/N2 boundaries, plantão, war room, who-to-call | `suporte`, `escalação`, `plantão` |
| `homologação` | Homologation of an integration, sandbox/test environment access, credenciais, trial setup, certificação técnica de integração | `sandbox`, `certificação-técnica` |
| `integração` | Integration work: API/webhook/middleware/iPaaS topology, ERP wiring, checkout/payment gateway behavior, parcelamento, antifraude, migrating a customer onto Wake from another platform or between Wake products | `migração`, `checkout`, `api-integração` |
| `marketing-eventos` | Co-marketing, conteúdo, campanhas, lançamentos, posicionamento, Wake Summit, ECBR, Partner Day, patrocínio, treinamento, workshops, webinars, capacitação | `marketing`, `eventos`, `treinamento` |
| `ecossistema-parceiros` | Catalog views, partner selection decision trees, "which partner for use case X", category overviews, partner directory | `catálogo`, `seleção-parceiro` |
| `competitivo` | Battlecards, objection handlers, win-loss analysis against a competitor. Name the competitor (VTEX, Shopify, Linx, Tray, etc.) in the body. | `vtex`, `shopify`, `linx`, `tray`, `battlecard` |

---

### Facets 3–5 — Collapsed into body content

Earlier drafts of this spec defined three additional facets — **Partner category**, **Product**, and **Competitive** — totalling 21 tags. They are **removed**. Their information lives in the article body, the title, and the `## Fontes` section:

- **Partner category** (agência, parceiro-tech, parceiro-marketplace, parceiro-pagamento, parceiro-martech, parceiro-logistica, integrador-erp): named explicitly in the article body and title. The taxonomy in `docs/wake-partnership-kb-playbook.md` §2 is the authoritative reference; the article cites it.
- **Product** (wake-commerce, wake-experience, wake-up, wake-oms, wake-dxp, storefront, commerce-api, data-wake, anymarket, pagar-me): named in the title and body. If an article is scoped to Wake Up, the title says so.
- **Competitive** (vtex, shopify, linx, tray): the topic tag is `competitivo` and the competitor is named in the title and body.

This is a deliberate trade: searchability per category drops slightly, but the canonical vocabulary becomes small enough that every author can hold it in working memory and the synonym map stays maintainable.

---

### What NOT to tag

These all go in the **article body**, the title, and the `## Fontes` / `## Verificação` sections:

- **Partner categories** (`agência`, `parceiro-tech`, `parceiro-marketplace`, `parceiro-pagamento`, `parceiro-martech`, `parceiro-logistica`, `integrador-erp`) — name them in title and body; they were a facet, now they aren't.
- **Wake product names** (`wake-commerce`, `wake-experience`, `wake-up`, `wake-oms`, `wake-dxp`, `storefront`, `commerce-api`, `data-wake`, `anymarket`, `pagar-me`) — title and body only.
- **Competitor names** (`vtex`, `shopify`, `linx`, `tray`, `rise`) — topic is `competitivo`; the competitor is named in the title and body.
- **Partner names** (`vurdere`, `biso`, `crmbonus`, `paypal`, `vindi`, `pagaleve`, `safrapay`, `omnichat`, `metakosmos`, `stayfilm`, `serasa-pass`, `hero-seguros`, `konfidency`, `revi`, `bonifiq`, `givex`, `mastercard`, `appmax`, `tuna`, `mercado-pago`, `nairuz`, `eficaz`, `catus`, `leanwork`, `vtrina`, `integra-do`, `iota-hub`, `jitterbit`, `argento-digital`, `linktou`, `floui`, etc.) — body only.
- **Client / lojista names** (`karcher`, `inbrands`, `balaroti`, `riachuelo`, `época-cosméticos`, `netshoes`, `dafiti`, etc.) — body only.
- **ERP product names** (`bling`, `tiny`, `totvs`, `sap`, `protheus`, `winthor`, `omie`, `sankhya`, `linx-microvix`) — body only; topic is `integração`.
- **Marketplace channel names** (`mercado-livre`, `magalu`, `shopee`, `privalia`, `shein`) — body only.
- **Bug IDs, ticket numbers, dates, version numbers** — body only.
- **Generic technical descriptors** (`api`, `webhook`, `token`, `json`, `payload`, `script`, `deploy`, `mobile`, `desktop`, `url`, `ip`, `sandbox`, `middleware`, `ipaas`, `conector`) — body only; topic is `integração` or `homologação`.
- **Generic commercial descriptors** (`vendas`, `comercial`, `cliente`, `produto`, `negociação`, `oportunidade`, `conversão`, `gmv`, `kpi`, `metas`, `forecast`, `prospect`, `prospecção`, `qualificação`, `follow-up`, `objeção`, `pitch`, `demo`, `cases`) — body only; choose a Facet 2 topic.
- **Money/paper sub-descriptors** (`contrato`, `nda`, `msa`, `revenue-share`, `comissão`, `comissionamento`, `faturamento`, `nf-e`, `fiscal`, `mensalidade`, `setup-fee`, `desconto`, `precificação`) — body only; topic is `contrato-comercial`.
- **Lead/pipeline sub-descriptors** (`lead`, `indicação`, `carrossel`, `pipeline`, `bant`, `conflito-canal`, `co-sell`) — body only; topic is `lead-pipeline`.
- **Marketing/event micro-descriptors** (`linkedin`, `podcast`, `banner`, `email`, `webinar`, `site`, `redes-sociais`, `conteúdo`, `campanha`, `convite`, `stand`, `palco`, `wake-summit`, `ecbr`, `partner-day`, `treinamento`, `workshop`) — body only; topic is `marketing-eventos`.
- **Support/escalation sub-descriptors** (`suporte`, `escalação`, `plantão`, `war-room`, `n1`, `n2`, `sla`, `ticket`, `atendimento`) — body only; topic is `suporte-escalação`.
- **Workflow / process verbs** (`cadastro`, `acesso`, `acompanhamento`, `validação`, `aprovação`, `mapeamento`, `sincronização`, `configuração`, `setup`, `rotina`) — body only.
- **Wake internal team names** (`parcerias-e-canais`, `sales-ops`, `partner-ops`, `coordenação`) — body only; they appear in the resolution/escalation steps, not as tags.
- **Audience-segment qualifiers** (`b2b`, `b2c`, `enterprise`, `smb`, `pessoa-física`, `omnichannel`, `farma`, `moda`, `cosméticos`, `food`, `grocery`, `autopeças`) — body only.
- **Tier labels** (`elite`, `essential`, `special`) — body only; topic is `programa-parceiros`.
- **Adjacent programs** (`wake-score`, `wakepédia`, `wake-mcp`, `wake-creators`, `wake-trends`, `wake-day`, `community`, `prm`) — body only.

If a candidate tag is not in Facets 1–2, the answer is always: **expand the body, not the tag list**.

---

### Migration from the legacy 913-tag corpus

The existing extracted articles use the pre-facet vocabulary. They will be re-normalized via a **rewritten** synonym map at `docs/.tag-synonym-map.json` — the previous 66-tag map is obsolete and must be regenerated for this stricter 15-tag scheme. The follow-up normalization pass will:

1. Map each legacy tag to its canonical Facet 1 or Facet 2 tag (or to `__drop__` if it belongs in the body).
2. Collapse the retired Facet 2/3/4/5 tags (partner category, product, competitive) by moving their information into the article body where it isn't already present.
3. Enforce the rule-of-2-to-4 by keeping at most one topic tag and the named audience tag(s) per article.
4. Flag articles that lose their topic tag during normalization for manual review.

**Follow-up required**: regenerate `docs/.tag-synonym-map.json` against this 15-tag canonical list before running the next normalization pass.

**All new articles produced by the `wake-partnership-kb-writer` agent must conform to this taxonomy from day one.** Legacy batches are grandfathered until the normalization pass runs; new content is not.

---

## 7. Language

- Write in **Brazilian Portuguese** (pt-BR)
- Match the register of the source conversation — technical articles stay technical, commercial articles stay commercial
- Do not translate proper nouns: Salesforce, Anymarket, Wake, TOTVS, SAP, etc.
- Avoid filler phrases: "é importante ressaltar que", "vale mencionar que", "conforme discutido"

---

## 8. What to Exclude

Do not create articles from:
- Greetings, celebrations, scheduling, logistics ("bom dia!", "vou checar aqui")
- Messages with no resolution ("alguém sabe?", "vou verificar")
- Duplicate topics already covered in the same file — merge instead
- Personal information: phone numbers, CPFs, personal emails, health data
- Credentials: API keys, passwords, tokens — describe the process, not the values

---

## 9. Delivery File Naming

```
kb-delivery-batch-NN.md
```

Where `NN` is zero-padded batch number: `01`, `02`, `03`...

Each delivery file must start with:

```markdown
# KB Delivery — Batch NN
**Artigos**: [start]–[end] do registry  
**Fontes**: [list of source files]  
**Total**: [N] artigos  
**Status**: prontos para importação  
**Data**: YYYY-MM

---
```

---

## 10. Authoring Authority

All new and modified Wake partnership KB articles **must** be produced via the `wake-partnership-kb-writer` agent unless the user explicitly overrides.

The agent enforces this spec and adds three guarantees beyond format compliance:

1. **Ground truth**: every factual claim is verified against `store/messages.db` before being written. Sources cited by group name + date.
2. **Audience targeting**: every article carries at least one of `audience-pam`, `audience-solutions`, `audience-sales`, `audience-cs`.
3. **Verificação footer**: every article ends with a table listing the SQLite queries used and the evidence found per claim, marked ✓ verificado / ⚠ inferido / ❌ sem fonte.

Companion files the agent uses:
- Agent definition: `~/.claude/agents/wake-partnership-kb-writer.md`
- Domain playbook: `docs/wake-partnership-kb-playbook.md`

Invoke via: `Agent(subagent_type="wake-partnership-kb-writer", prompt="...")`.

---

## 11. Known Legacy Inconsistencies (do not replicate)

These non-canonical patterns exist in the initial 668 extracted articles. They are acceptable in already-delivered batches but must not appear in new articles:

| Legacy pattern | Canonical replacement |
|---|---|
| `BEST_PRACTICES-5`, `BEST-PRACTICES-3`, `BP-004` | `BPR-005` |
| `HOW_TO-3`, `HOW-TO-6`, `HT-002` | `HTO-003` |
| `TROUBLESHOOTING-4`, `TRBL-2`, `TRB-001` | `TRB-004` |
| `FAQ-1` (no zero-pad) | `FAQ-001` |
| JID field in metadata table | Omit entirely |
| Phone numbers in contact tables | Replace with `[contato interno]` |
| Tags without backticks | Wrap in backticks |
| Tags with spaces (`revenue share`) | Hyphenate (`revenue-share`) |
