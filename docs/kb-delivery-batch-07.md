# KB Delivery — Batch 07
**Artigos**: 437–518 do registry  
**Fontes**: kb-parceiros-pagamento.md (completo), kb-projeto-serasa-leanwork.md (completo), kb-parceiros-wave6b.md (completo)  
**Total**: 82 artigos  
**Status**: prontos para importação  
**Data**: 2026-06

---

# KB Extract — Parceiros de Pagamento (PayPal, Mercado Pago, Givex, Safrapay, Vindi)
> Extracted: 2026-06-04

---

## Grupo: WAKE & PAYPAL
**JID:** `120363401981841957@g.us`

---

### FAQ-001 — Wallets disponíveis no conector PayPal

| Campo | Valor |
|---|---|
| **Grupo/Chat** | WAKE & PAYPAL |
| **Data** | 2025-08-11 |
| **Categoria** | `faq` |
| **Tags** | `programa-parceiros`, `audience-solutions` |

**Pergunta:** Quais wallets já estão disponíveis no conector PayPal da Wake?

**Resposta (PayPal):** A wallet do PayPal já está integrada, assim como o checkout transparente (BCDC). Apple Pay e Google Pay ainda não estão disponíveis no BCDC — ambos estão no roadmap com expectativa de entrega até final de H2 2025. Na solução de Braintree (BT), Apple Pay, Google Pay e 3DS já estão disponíveis.

**Detalhes adicionais:**
- Apple Pay e Google Pay via BCDC: previsão de entrega 08/09/2025 para início de homologação.
- Via Braintree (BT): Apple Pay, Google Pay e 3DS já funcionavam para clientes como Shoulder (que também usa crossborder).
- A Bycoders (SI parceira) era responsável pela implementação de wallets para Wake via Braintree.

---

### FAQ-002 — Diferença entre PayPal Plus (PP+) e BCDC

| Campo | Valor |
|---|---|
| **Grupo/Chat** | WAKE & PAYPAL |
| **Data** | 2025-08-06 |
| **Categoria** | `faq` |
| **Tags** | `integração`, `audience-pam`, `audience-solutions` |

**Pergunta:** É possível migrar clientes do PayPal Plus para o BCDC? Existe alguma issue pendente?

**Resposta (PayPal):** Sim, é possível migrar. Havia um plano criado junto com o time Wake dividindo os clientes por sprints: na primeira onda entram clientes menos expressivos para acompanhamento, e com sucesso confirmado, aciona-se os demais.

**Contexto estratégico:** Quanto mais lojistas migrarem do PayPal Plus para o BCDC, maior o incentivo do PayPal para investimentos na parceria com Wake. A migração não é automática — é feita loja a loja, com abordagem comercial para comunicar os benefícios.

**Processo:**
1. Time comercial Wake aborda os lojistas explicando benefícios da migração.
2. Informações sobre configuração foram incluídas tanto no artigo antigo quanto em novo artigo de configuração.
3. Qualquer comunicação de massa só é enviada após confirmação de sucesso da primeira bateria de lojas.

---

### TRB-001 — Alto abandono de carrinho / User Dropoff com PayPal

| Campo | Valor |
|---|---|
| **Grupo/Chat** | WAKE & PAYPAL |
| **Data** | 2025-07-07 |
| **Categoria** | `troubleshooting` |
| **Tags** | `integração`, `audience-cs`, `audience-solutions` |

**Sintomas:** Cliente Mimeria com 34% de User Dropoff e alto abandono de carrinho no checkout com PayPal.

**Causa raiz:** [fonte original sem informação] — pedido-teste realizado pelo time PayPal pareceu normal; causa raiz não isolada.

**Resolução:**
1. Verificar se há algo em comum nos pedidos abandonados (estado, cidade, perfil do pagador) para reproduzir o cenário.
2. Não acionar o cliente antes de investigar internamente.
3. Solicitar ao time PayPal que verifique pelo lado deles sem envolver o cliente.
4. Realizar pedido de teste próximo às condições dos pedidos abandonados (mesmo estado, cidade, método de pagamento).

---

### TRB-002 — Problema no checkout em múltiplos merchants PayPal

| Campo | Valor |
|---|---|
| **Grupo/Chat** | WAKE & PAYPAL |
| **Data** | 2025-07-10 |
| **Categoria** | `troubleshooting` |
| **Tags** | `integração`, `audience-cs` |

**Sintomas:** 3 merchants reportaram problema no checkout simultaneamente.

**Causa raiz:** [fonte original sem informação] — causa técnica específica não documentada nas conversas; sintoma foi tratado por canal de escalação.

**Resolução:** Time técnico PayPal encaminhou e-mail com dúvidas para o time Wake (Amanda). Amanda (Wake) tratou o caso diretamente após receber o e-mail. Escalação via canal técnico por e-mail foi o caminho correto.

**Canal de escalação:** Para problemas técnicos com múltiplos merchants, o time técnico do PayPal deve enviar e-mail diretamente para o responsável técnico na Wake (no período, Amanda).

---

### HTO-001 — Processo de migração PayPal Plus → BCDC

| Campo | Valor |
|---|---|
| **Grupo/Chat** | WAKE & PAYPAL |
| **Data** | 2025-08-11 |
| **Categoria** | `how_to` |
| **Tags** | `integração`, `audience-pam` |

**Fluxo de migração PayPal Plus para BCDC:**

1. **Preparação:** Time Wake cria lista de lojistas segmentada por sprints (do menos ao mais expressivo).
2. **Sprint 1:** Entrar em contato com lojistas menos expressivos para acompanhar a experiência de migração.
3. **Validação:** Confirmar que tudo está OK com a primeira bateria antes de avançar.
4. **Sprint 2+:** Acionar lojistas de maior porte e expressividade após validação positiva.
5. **Comunicação:** Time comercial Wake aborda os lojistas; PayPal não comunica diretamente.
6. **Documentação:** Artigos atualizados na central de atendimento Wake com instruções de configuração do BCDC.

**Importante:** A migração não é automática — cada lojista precisa ser acionado individualmente. O plugin do PayPal Plus deve ser ocultado para novos lojistas durante o processo.

---

### BPR-001 — Estratégia de parceria e posicionamento PayPal

| Campo | Valor |
|---|---|
| **Grupo/Chat** | WAKE & PAYPAL |
| **Data** | 2025-08-05 |
| **Categoria** | `best_practices` |
| **Tags** | `programa-parceiros`, `audience-pam`, `audience-sales` |

**Recomendação:** Posicionar a parceria Wake × PayPal com foco em: migração de lojistas para BCDC, expansão do escopo de contas que já usam Braintree (VTEX), oferta de crossborder para lojistas com ambição global, uso de wallets (Apple Pay/Google Pay) como diferencial competitivo e participação conjunta em eventos Wake Summit como patrocinador.

**Rationale:**
- **Alavanca de investimento:** Quanto mais lojistas Wake migrarem para BCDC (em detrimento do PayPal Plus), maior o ânimo do PayPal para investir na parceria.
- **Oportunidade de expansão:** Clientes que já usam PayPal via Braintree (VTEX) podem ter a parte nacional migrada para Wake — é uma oportunidade de ampliar o escopo da conta.
- **Internacionalização:** Tema crescente entre lojistas Wake — PayPal via Braintree suporta crossborder, o que pode ser argumento de venda para lojistas com ambição global.
- **Wallets como diferencial:** Apple Pay e Google Pay no roadmap do BCDC são um argumento forte para clientes com necessidade de carteiras digitais (ex: cliente Shoulder).
- **Posicionamento no evento Wake Summit:** Parceiros de pagamento como PayPal podem participar como patrocinadores — alinhamento estratégico aumenta visibilidade mútua.

---

## Grupo: Mercado Pago + Wake
**JID:** `120363419259418672@g.us`

---

### TRB-003 — Erro "El payment method no viene en el request"

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Mercado Pago + Wake |
| **Data** | 2025-10-13 |
| **Categoria** | `troubleshooting` |
| **Tags** | `integração`, `audience-solutions`, `audience-cs` |

**Sintomas:** A plataforma Wake não estava enviando os campos `payment_method_id` e o ID de tokenização no JSON da requisição de pagamento ao Mercado Pago. Apenas o campo `installments` era enviado, causando o erro:

```
{"message": "El payment method no viene en el request y no es un pago con Tarjeta, no puedo deducirlo"}
```

**Exemplo do payload com erro:**
```json
{
  "installments": 1,
  "transaction_amount": 1.25,
  "notification_url": "...",
  "external_reference": "978521",
  "sponsor_id": 193313608,
  "payer": { ... }
}
// FALTANDO: "payment_method_id" e token de cartão
```

**Causa raiz:** Bug na integração Wake × Mercado Pago: os campos obrigatórios `payment_method_id` e token não eram incluídos no payload.

**Resolução:**
1. Lojista (Potiguar) compartilhou o log de erro do Mercado Pago com a Wake.
2. Time de Parcerias Wake acionado (canal WhatsApp + Slack).
3. Canal de suporte recomendado: abrir ticket com o time técnico Wake, incluindo payload completo, IDs dos pedidos e logs.
4. Para escalar com o Mercado Pago: incluir `douglas.cruz@mercadolivre.com` e `renata.rcandrade@mercadopago.com.br` em CC no ticket.
5. Também há canal de Slack entre Wake e Mercado Pago para acompanhamento.

**Erro secundário identificado:** `"El request ya fue posteado"` — indica requisição duplicada sendo bloqueada pelo Mercado Pago.

---

### TRB-004 — Erro "Bin not found" ao testar parcelamento em 18x

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Mercado Pago + Wake |
| **Data** | 2026-02-23 |
| **Categoria** | `troubleshooting` |
| **Tags** | `integração`, `audience-solutions` |

**Sintomas:** Ao testar parcelamento em 18x em ambiente Sandbox, retornava erro `"Bin not found"`, mesmo com cartão que o sistema reconhecia corretamente como Mastercard.

**Causa raiz:** O ambiente Sandbox não refletia o parcelamento em 18x, enquanto a conta do seller em produção já estava liberada para essa modalidade.

**Resolução:**
- Teste realizado em produção com conta real confirmou que o parcelamento em 18x funcionou corretamente nas duas pontas.
- Contas em Sandbox não suportam parcelamento acima de 12x por padrão — o ambiente de produção já estava configurado.
- Alternativa para testes: usar uma conta Mercado Pago TestUser (ambiente controlado do MP) para validar a jornada de ponta a ponta.

**Importante:** O parcelamento acima de 12x é liberado de forma controlada para alguns sellers. Para verificar o parcelamento disponível, utilizar o endpoint:
```
GET /v1/payment_methods/installments?bin={BIN}&public_key={CHAVE_PUBLICA}
```

---

### FAQ-003 — Parcelamento em 18x no Mercado Pago via Wake

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Mercado Pago + Wake |
| **Data** | 2025-10-29 |
| **Categoria** | `faq` |
| **Tags** | `integração`, `audience-solutions`, `audience-sales` |

**Pergunta:** A Wake suporta parcelamento em 18x com Mercado Pago?

**Resposta:** Tecnicamente, na plataforma Wake o lojista pode cadastrar quantas parcelas quiser no grupo de pagamento (sem limite fixo no admin). Porém, o parcelamento acima de 12x depende de liberação controlada do lado do Mercado Pago, conforme as particularidades dos emissores.

**Como configurar no admin Wake:**
1. Lojista acessa o admin e cria um **grupo de pagamento**.
2. Vincula uma **configuração de parcelamento** (com qtd de parcelas, juros, etc.) ao grupo.
3. Vincula um **conector de pagamentos** ao grupo.

Essa configuração é anterior e independente do conector utilizado. O lojista tem liberdade para cadastrar quantas parcelas desejar — a validação de disponibilidade real é feita pelo emissor/MP.

---

### HTO-002 — Documentação de integração Mercado Pago com Wake

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Mercado Pago + Wake |
| **Data** | 2026-04-29 |
| **Categoria** | `how_to` |
| **Tags** | `integração`, `audience-solutions` |

**Como integrar o Mercado Pago com a Wake (via conector customizado):**

- **Documentação principal:** https://wakecommerce.readme.io/docs/conector-de-pagamento-customizado
- **Documentação geral Wake:** https://wakecommerce.readme.io/

Esse é o documento enviado para provedores de pagamento que desejam construir integração com a Wake. Define o contrato de API e os fluxos esperados.

**Canal de dúvidas:** Grupo WhatsApp de parceria Wake + Mercado Pago.

---

### TRB-005 — Pedido travado em "Aguardando pagamento" (Sandbox MP)

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Mercado Pago + Wake |
| **Data** | 2026-05-21 |
| **Categoria** | `troubleshooting` |
| **Tags** | `homologação`, `audience-solutions`, `audience-cs` |

**Sintomas:** Cliente Venari (CNPJ: 41.907.917/0001-51) na fase final de testes para Go Live — ao realizar pedido, ele permanecia apenas com status "Aguardando pagamento", impedindo o avanço no fluxo de testes em Sandbox.

**Causa raiz:** [fonte original sem informação] — causa raiz não isolada nas conversas registradas; investigação iniciada via análise de payload.

**Resolução:** Solicitado ao Mercado Pago o payload/request enviado à sua API para análise. Para casos de integração em Sandbox, o parceiro Mercado Pago pode analisar via request compartilhado ou acompanhar via time de parcerias (canal Slack Wake × MP).

**Contatos Mercado Pago para suporte técnico em integrações:**
- douglas.cruz@mercadolivre.com
- renata.rcandrade@mercadopago.com.br

---

### BPR-002 — Monitoramento de incidentes Mercado Pago

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Mercado Pago + Wake |
| **Data** | 2025-10-16 |
| **Categoria** | `best_practices` |
| **Tags** | `suporte-escalação`, `audience-cs` |

**Recomendação:** Ao identificar queda repentina no volume de transações Mercado Pago, verificar a status page Wake/fbits (https://status.fbits.net/) antes de escalar; comunicar o parceiro MP pelo canal WhatsApp de parceria para confirmação rápida; documentar o incidente com horário de início e recuperação.

**Rationale:** Identificado volume baixo de transações a partir das 11h — possível afetação geral no Mercado Pago. O Mercado Pago (via gateway fbits) possui uma página de status para incidentes (ex: incidente documentado em https://status.fbits.net/incidents/56xh7dtcy8mt), o que permite confirmar incidentes amplos antes de iniciar diagnóstico cliente a cliente.

---

## Grupo: Wake + Givex
**JID:** `120363419883584104@g.us`

---

### HTO-003 — Processo de onboarding de nova integração (parceiro de pagamento)

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Wake + Givex |
| **Data** | 2025-08-19 |
| **Categoria** | `how_to` |
| **Tags** | `integração`, `audience-pam`, `audience-solutions` |

**Fluxo de onboarding de novo parceiro de integração de pagamento com a Wake:**

1. **Reunião inicial** com líder de parcerias Wake + envio de e-mail introdutório.
2. **NDA:** Jurídico Wake devolve comentários; parceiro assina NDA (pode ser em inglês).
3. **Integradores:** Wake sugere 4 parceiros integradores (SIs homologados como Vtrina, Leap to the Future, Wicomm).
4. **Orçamento:** SI parceira apresenta orçamento (ex: Vtrina estimou a partir de R$14.700 para integração Givex, sujeito a ajuste após análise técnica).
5. **NDA do integrador:** SI também assina NDA antes de ter acesso à documentação.
6. **Reunião de arquitetura:** Definição de escopo e desenho de arquitetura da integração.
7. **Desenvolvimento:** SI executa com base nos fluxos desenhados e documentação Wake.
8. **Homologação:** Validação ponta a ponta antes do Go Live.

**Documentação para integradores:**
- Vale-troca (Crédito em Conta Corrente): https://atendimento.wake.tech/hc/pt-br/articles/21406566351127-Conta-Corrente e https://wakecommerce.readme.io/docs/inserindo-lancamento-na-conta-corrente-do-cliente
- Gift Card (Produto Digital): https://atendimento.wake.tech/hc/pt-br/articles/24166558988695-Produto-Digital

---

### HTO-004 — Fluxo de integração Givex × Wake (gift card e vale-troca)

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Wake + Givex |
| **Data** | 2026-03-18 |
| **Categoria** | `how_to` |
| **Tags** | `integração`, `audience-solutions` |

**Funcionalidades mapeadas para integração Givex × Wake:**

- **Vale-troca:** Utiliza o recurso de **Crédito em Conta Corrente** da Wake.
- **Gift Card (produto digital):** Utiliza o recurso de **Produto Digital** da Wake.

**Fluxos técnicos desenhados (sem middleware — integração direta Givex × Wake):**
- Emissão de vale-troca
- Venda de gift card
- Resgate de gift card
- Resgate de vale-troca
- Devoluções
- Validação de saldo
- Chargeback
- Acompanhamento de pedidos (múltiplas opções de fluxo desenhadas)

**Arquitetura:** Integração direta Givex → Wake, sem intermediadores (middleware), conforme definido em reunião técnica com o time Givex/Shift4. Os fluxos foram desenhados em Miro e compartilhados por e-mail.

**Consideração de reutilização:** O projeto foi desenhado para ser genérico (não exclusivo para um cliente como Shoulder), permitindo uso por outros lojistas Wake.

---

### FAQ-004 — Custo de integração via parceiro integrador homologado

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Wake + Givex |
| **Data** | 2025-09-11 |
| **Categoria** | `faq` |
| **Tags** | `integração`, `audience-pam`, `audience-solutions` |

**Pergunta:** Qual o custo estimado para integração com a Wake via SI parceira?

**Resposta:** O orçamento de integração via Vtrina (SI parceira da Wake) para o caso Givex partiu de **R$14.700**, podendo variar após análise das APIs e complexidades do projeto. Esse valor foi apresentado antes de calls técnicas — o orçamento final é ajustado após reunião de arquitetura.

**Nota:** Para a integração Givex especificamente, Wake absorveu o custo do lado da Wake para o cliente Shoulder (projeto piloto), enquanto o custo do SI (Vtrina) ficou a cargo de negociação entre Givex e Vtrina.

---

### BPR-003 — Gestão do processo de parceria técnica (NDA, integradores, cronograma)

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Wake + Givex |
| **Data** | 2026-05-06 |
| **Categoria** | `best_practices` |
| **Tags** | `programa-parceiros`, `audience-pam` |

**Recomendação:** Em parcerias técnicas com SI intermediador, coordenar assinatura de NDA de todos os players antes da reunião técnica, validar e-mails de contato corretos, garantir acesso a ferramentas de colaboração (Miro), definir cronograma com marcos comunicados ao parceiro e separar integrações genéricas de customizações específicas de cliente.

**Rationale:** (lições aprendidas na parceria Givex × Wake)

- **NDA com múltiplos players:** Quando há um SI intermediador (ex: Vtrina, Wicomm), cada um deve assinar o NDA separadamente antes de ter acesso à documentação técnica. Coordenar a assinatura de todos antes da reunião técnica evita atrasos.
- **E-mail de contato correto:** Validar o e-mail certo do parceiro evita falhas de comunicação (ex: iago@givex.com.br vs iago.annes@givex.com vs iago.annes@shift4.com).
- **Ferramentas de colaboração:** Uso de Miro para desenho de fluxos — garantir que o acesso ao Miro funcione (domínios corporativos podem bloquear convites; usar e-mail pessoal como alternativa).
- **Cronograma com marcos:** Definir datas-alvo (ex: RM de Go Live em 02/06) e comunicar ao parceiro aumenta a pressão positiva no avanço.
- **Separação integrações genéricas de customizações:** Garantir que nenhum fluxo seja customizado exclusivamente para um cliente, para que a integração sirva a múltiplos lojistas.

---

### FAQ-005 — Recursos disponíveis na Wake para Gift Card e Vale-Troca

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Wake + Givex |
| **Data** | 2026-02-03 |
| **Categoria** | `faq` |
| **Tags** | `integração`, `audience-solutions`, `audience-sales` |

**Pergunta:** A Wake tem suporte nativo para gift card e vale-troca?

**Resposta:** Sim. A Wake possui dois recursos nativos relevantes:

1. **Vale-troca → Crédito em Conta Corrente**
   - Central de ajuda: https://atendimento.wake.tech/hc/pt-br/articles/21406566351127-Conta-Corrente
   - API (readme): https://wakecommerce.readme.io/docs/inserindo-lancamento-na-conta-corrente-do-cliente

2. **Gift Card → Produto Digital**
   - Central de ajuda: https://atendimento.wake.tech/hc/pt-br/articles/24166558988695-Produto-Digital

Esses recursos precisam ser integrados com o provedor externo de gift card/vale-troca (ex: Givex/Shift4) via API, usando a documentação de conector customizado da Wake.

---

## Grupo: Safrapay + Wake
**JID:** `120363423003129579@g.us`

---

### HTO-005 — Processo de integração Safrapay × Wake (conector de pagamento)

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Safrapay + Wake |
| **Data** | 2026-04-20 |
| **Categoria** | `how_to` |
| **Tags** | `integração`, `audience-solutions` |

**Fluxo de integração e homologação Safrapay × Wake:**

1. **Configuração inicial:** Acesso ao painel admin Wake via https://partnertechsafrapay.fbits.app/ (ambiente parceiro Safrapay).
2. **Frete fictício:** Wake cadastra frete fictício para viabilizar testes de checkout.
3. **Conector de pagamento:** Configurado diretamente no admin da loja.
4. **Testes de cartão:** Após resolver payload (incluindo `payment-details`), transações de cartão passam corretamente.
5. **PIX:** O QR Code do Pix é retornado pelo endpoint de `/confirmation` — a Wake chama esse endpoint após o checkout. O fluxo de confirmação é a forma correta de retornar o QR Code.
6. **Validação end-to-end:** Call conjunta Wake × Safrapay para validar: aprovação de transações, cancelamento total e parcial, e reflexo de cancelamento do portal Safrapay na plataforma Wake.
7. **Homologação formal:** Coordenada com o time de pagamentos Wake (Natali) — após homologação, evolui para RM de produção.

**Documentação:** https://wakecommerce.readme.io/docs/conector-de-pagamento-customizado

---

### TRB-006 — QR Code do Pix não carrega após checkout Safrapay

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Safrapay + Wake |
| **Data** | 2026-04-20 |
| **Categoria** | `troubleshooting` |
| **Tags** | `integração`, `audience-solutions`, `audience-cs` |

**Sintomas:** Após configurar o endpoint `/confirmation` para retornar o QR Code do Pix, o integrador não conseguia identificar chamada da Wake para esse endpoint após o checkout.

**Causa raiz:** Problema no front-end da loja ou na URL configurada para o endpoint de confirmação (causa provável conforme diagnóstico).

**Resolução:**
1. Confirmar que a URL do endpoint de confirmação está configurada corretamente no conector.
2. Verificar se o problema está no front da loja (não renderizando a resposta) ou na configuração da URL.
3. Acionar time de Suporte Tech Wake com: nome da loja, URL do pedido de teste, ID do pedido e logs.

**Canal de suporte tech Wake:** WhatsApp de suporte (bot inicial → informar que é parceiro realizando integração → direcionamento para squad de Pagamentos).

---

### TRB-007 — Acesso ao painel admin Wake (parceiro Safrapay)

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Safrapay + Wake |
| **Data** | 2026-03-17 |
| **Categoria** | `troubleshooting` |
| **Tags** | `integração`, `audience-cs`, `audience-solutions` |

**Sintomas:** Usuário do parceiro Safrapay (Danilo) não conseguia receber o e-mail de recuperação de senha para acessar o painel admin Wake.

**Causa raiz:** [fonte original sem informação] — falha de entrega do e-mail de recuperação não foi diagnosticada; resolução foi contornar o e-mail.

**Resolução:**
1. Wake gerou link direto de redefinição de senha (sem depender do e-mail): `https://idm.fbits.net/Account/AlterarSenha?Chave={CHAVE}`
2. Usuário acessa o link e redefine a senha diretamente pelo formulário.

**URL do painel admin Safrapay:** https://partnertechsafrapay.fbits.app/

**Nota:** Links de recuperação de senha têm prazo de expiração — se expirar, o time Wake gera um novo link.

---

### FAQ-006 — Versão do conector Safrapay e atualização

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Safrapay + Wake |
| **Data** | 2025-08-04 |
| **Categoria** | `faq` |
| **Tags** | `integração`, `audience-solutions` |

**Pergunta:** A versão 01.04 do conector Safrapay é a mais atual?

**Resposta (Safrapay):** A integração existente é antiga e a versão 01.04 não é a atual. Existe uma versão mais recente — confirmação dos endpoints em uso deve ser feita na agenda técnica. O time técnico do Safrapay (Edu) é responsável pela atualização e manutenção do conector.

---

### BPR-004 — Validação de cenários transacionais na integração

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Safrapay + Wake |
| **Data** | 2026-04-22 |
| **Categoria** | `best_practices` |
| **Tags** | `homologação`, `audience-solutions` |

**Recomendação:** Em call de homologação ponta a ponta (Safrapay × Wake), validar simultaneamente o checklist completo de cenários transacionais com ambos os times presentes.

**Rationale:** Call conjunta com ambos os times (Wake Pagamentos + Safrapay técnico) permite validação simultânea dos cenários, reduzindo ciclos de comunicação. Cenários a cobrir:

- [ ] Aprovação de transações (cartão de crédito, débito, Pix)
- [ ] Cancelamento total pelo portal Wake → verificar reflexo no portal Safrapay
- [ ] Cancelamento parcial pelo portal Wake → verificar reflexo no portal Safrapay
- [ ] Cancelamento pelo portal Safrapay → verificar se status reflete corretamente na plataforma Wake
- [ ] Fluxo de Pix: geração e exibição do QR Code no front da loja
- [ ] Reflexo de frete na cotação durante checkout

---

### BPR-005 — Antifraude e conformidade regulatória (CNPJ alfanumérico)

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Safrapay + Wake |
| **Data** | 2026-05-25 |
| **Categoria** | `best_practices` |
| **Tags** | `integração`, `audience-solutions` |

**Recomendação:** Em integrações de pagamento Safrapay × Wake, tratar conformidade regulatória (CNPJ Alfanumérico IN 2.229/2024) e coleta de dados para antifraude como tópicos críticos, alinhando diretamente com a equipe de Pagamentos Wake (Natali) desde o início do projeto.

**Rationale:**

1. **CNPJ Alfanumérico (IN 2.229/2024):** Instrução normativa federal que altera o formato do CNPJ para incluir letras. Wake estava avaliando a homologação deste requisito com o time de pagamentos (Natali). Status: em análise durante integração Safrapay.

2. **Coleta de dados para antifraude:** Ponto crítico durante a integração — o integrador (Danilo/Safrapay) tinha dúvidas sobre a forma correta de coletar dados para o módulo de antifraude. Recomendado alinhar diretamente com o time de pagamentos Wake (Natali).

**Contato para tópicos regulatórios/antifraude Wake:** Equipe de Pagamentos Wake (Natali).

---

## Grupo: Vindi + Wake
**JID:** `120363420795576168@g.us`

---

### HTO-006 — Processo de treinamento de parceiro de pagamento para time interno Wake

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Vindi + Wake |
| **Data** | 2025-08-22 |
| **Categoria** | `how_to` |
| **Tags** | `marketing-eventos`, `audience-pam` |

**Processo de treinamento Vindi para times de CS e Vendas Wake:**

1. **Solicitação:** Time de parceiros Vindi solicita treinamento via canal WhatsApp de parceria.
2. **Aprovação antecipada:** Todo material (PPT/PDF) deve ser enviado antecipadamente para aprovação dos coordenadores de CS e comercial Wake antes da apresentação.
   - E-mails de aprovação: `natasha.silva@wake.tech` / `leticia.cena@wake.tech`
3. **Agendamento:** Coordenadores de CS alinham com analistas antes do treinamento. Evitar agendar com menos de 1 semana de antecedência para garantir quórum.
4. **Formato:** 30 minutos é suficiente para um treinamento inicial — pode ser acoplado a agendas já existentes dos times.
5. **Lançamento de campanha:** O treinamento pode ser combinado com o lançamento de uma campanha de vendas (ver modelo Vindi abaixo).

**Contatos internos Wake para treinamentos:** Gestor de CS + Coordenadores de CS + Gerente Comercial.

---

### BPR-006 — Campanha de incentivo de vendas com parceiro de pagamento

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Vindi + Wake |
| **Data** | 2025-12-17 |
| **Categoria** | `best_practices` |
| **Tags** | `marketing-eventos`, `audience-pam`, `audience-cs` |

**Recomendação:** Manter campanhas de incentivo contínuas de indicação CS → parceiro de pagamento (sem hiatos longos) para preservar o ritmo de indicações. Ao encerrar uma campanha, iniciar planejamento da próxima imediatamente.

**Rationale:** Resultados da Campanha "Todos pela Vindi" (jun–dez 2025):**

| Métrica | Resultado |
|---|---|
| Total de clientes indicados pelos CS | 84 |
| Clientes fechados com a Vindi | 23 |
| Taxa de conversão da campanha | 23% |
| Taxa de conversão padrão Vindi | 20% |
| Oportunidades ainda ativas no funil | 17 |

**Estrutura da campanha:**
- CS Wake indicavam clientes para a Vindi e recebiam prêmios por indicações convertidas.
- Coordenadores de CS e gerente de onboarding foram fundamentais para o engajamento dos analistas.
- Campanha encerrada em 15/12/2025; ganhadores anunciados pela Vindi após contabilização.

---

### BPR-007 — Estratégia de co-marketing e eventos com parceiros de pagamento

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Vindi + Wake |
| **Data** | 2025-10-01 |
| **Categoria** | `best_practices` |
| **Tags** | `marketing-eventos`, `audience-pam` |

**Recomendação:** Estruturar co-marketing com parceiros de pagamento usando três frentes complementares — webinar conjunto temático, Wake Days com agências e patrocínio de Wake Summit — e oferecer pacotes combinados de patrocínio com desconto para reduzir resistência interna do parceiro.

**Rationale:** modelos de co-marketing testados:**

1. **Webinar conjunto:** Vindi foi convidada para participar de webinar de Black Friday com tema "como otimizar plataforma para picos de venda" — foco em otimização de checkout, escalabilidade, e volume alto de tráfego. Parceiro pode indicar speaker (diretor, product manager ou gerente de CS).

2. **Wake Days com agências:** Eventos presenciais com agências parceiras Wake — oportunidade para o parceiro de pagamento se apresentar ao ecossistema de agências.

3. **Patrocínio de evento:** Parceiro pode patrocinar o Wake Summit (maior evento do ecossistema Wake) com cotas de Almoço VIP e Jantar Pré-Summit. Permite exposição a C-LEVELs de grandes varejistas.

**Dica de negociação:** Pacotes combinados (ex: Almoço VIP + Jantar) têm desconto — Safrapay negociou R$80k para duas cotas que custariam R$110k separadas. Útil para fechar patrocínio com menor resistência interna do parceiro.

---

### FAQ-007 — Como funciona a parceria comercial Vindi × Wake (modelo de indicação)

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Vindi + Wake |
| **Data** | 2025-08-05 |
| **Categoria** | `faq` |
| **Tags** | `programa-parceiros`, `audience-pam` |

**Pergunta:** Como funciona o modelo de parceria comercial entre Vindi e Wake?

**Resposta:** O modelo principal é baseado em **indicações de clientes**: CS Wake identificam lojistas com perfil para Vindi e indicam formalmente via campanha. Do lado Vindi, há acompanhamento comercial dos leads indicados.

**Pontos de contato recorrentes:**
- Reuniões de pulso de performance (recorrentes semanais/quinzenais).
- Campanhas de vendas com premiação para CS Wake.
- Treinamentos periódicos do time Vindi para CS e vendas Wake.
- Participação em eventos Wake (Wake Days, Wake Summit).

**Segmentos indicados pela Vindi:** Clientes Vindi que buscam plataforma de e-commerce — parceria bidirecional (Wake indica clientes para Vindi e Vindi indica clientes para Wake).

---

### FAQ-008 — Entregáveis e prazos para expositor no Wake Summit

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Vindi + Wake |
| **Data** | 2025-08-19 |
| **Categoria** | `faq` |
| **Tags** | `marketing-eventos`, `audience-pam` |

**Pergunta:** Quais são os entregáveis para um parceiro expositor no Wake Summit?

**Resposta:** O time de eventos Wake (Partner Ops) disponibiliza um **manual do expositor** que detalha todos os entregáveis e prazos. Uma reunião de apresentação do manual pode ser agendada com o time de eventos.

**Entregáveis típicos confirmados nas conversas:**
- Logo em todas as versões + vetor + manual da marca (enviar ao time de eventos Wake).
- Arte/criativo do parceiro (aprovado pela agência e enviado ao time de eventos).
- Vídeo institucional (quando aplicável).
- Lista de staff com nome completo, CPF, e-mail, cargo e telefone para credenciamento.
- Definição de fornecedores que o parceiro levará ao evento.

**Convites:** Eventos Wake são fechados (sem venda de ingressos). Parceiros patrocinadores recebem cotas de convites para clientes e prospects. Convites são pessoais e intransferíveis.

**Agendamento do manual do expositor:** https://calendar.app.google/G3rQ3devJVJQzcHt8

---

# KB Extract — Projeto Serasa Pass <> Leanwork <> Wake
> JID: 120363422164488957@g.us | Extracted: 2026-06-04

---

## FAQ

---

### FAQ-001 — Qual é o campo `type` na criação de sessão e quando usar cada valor?

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Projeto Serasa Pass <> Leanwork <> Wake |
| **Data** | 2025-09-01 |
| **Categoria** | `faq` |
| **Tags** | `integração`, `audience-solutions` |

**Pergunta:** A API de criação de sessão tem um campo `type`. Quais são os tipos disponíveis e quando usar cada um?

**Resposta:** O campo `type` determina o comportamento da sessão:
- **Type 1 (webhook / Desktop):** Usado quando o usuário está em um dispositivo desktop. O fluxo exibe um QR Code para que o usuário leia com o celular e faça o compartilhamento no app. Ao finalizar, a API bate no endpoint webhook configurado com o `sessionId`.
- **Type 2 (url / Mobile):** Usado quando o usuário está em um dispositivo móvel. O link da sessão redireciona diretamente para o app Serasa, sem mostrar QR Code. Ao finalizar, redireciona o usuário de volta à URL configurada.

O retorno da criação de sessão inclui um deeplink para o app, além da `webUrl` com o `sessionId`.

---

### FAQ-002 — Qual URL usar para o fluxo de biometria facial (desktop x mobile)?

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Projeto Serasa Pass <> Leanwork <> Wake |
| **Data** | 2026-01-14 |
| **Categoria** | `faq` |
| **Tags** | `integração`, `audience-solutions` |

**Pergunta:** O link utilizado para biometria facial é `https://serasa.com.br/autorizacao/serasa-pass/verificacao/cpf?sessionId=###`. Esse link estava exibindo QR Code antes e agora manda direto para tela do CPF. O comportamento mudou?

**Resposta:** Sim, o comportamento mudou com a atualização de funcionalidades. A solução é **remover o `/cpf` da URL**:

```
INCORRETO: https://serasa.com.br/autorizacao/serasa-pass/verificacao/cpf?sessionId=###
CORRETO:   https://serasa.com.br/autorizacao/serasa-pass/verificacao?sessionId=###
```

Com a URL correta, no desktop aparece o QR Code para fazer o processo completo no celular antes do retorno do webhook. O `type` continua sendo enviado conforme o fluxo:
- Mobile: `type = url`
- Desktop: `type = webhook`

---

### FAQ-003 — Como funciona o formato dos webhooks para desktop e mobile?

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Projeto Serasa Pass <> Leanwork <> Wake |
| **Data** | 2025-09-24 |
| **Categoria** | `faq` |
| **Tags** | `integração`, `audience-solutions` |

**Pergunta:** Como configurar as URLs de webhook e redirect para desktop e mobile? As URLs devem ser diferentes?

**Resposta:** Sim, devem ser configuradas URLs distintas para desktop e mobile:

- **Desktop (webhook de notificação):** `POST` para a URL com `sessionId` no body — a Serasa bate nesse endpoint quando o cliente terminar de compartilhar no app.
  - Exemplo: `https://loja.com.br/webhook/token/{sessionId}`

- **Mobile (URL de redirecionamento):** Após o compartilhamento no app, o usuário é redirecionado para essa URL — não é um POST, é um redirect.
  - Formato: `/webhook-mobile?sessionId={id}&eventName=user_consent`
  - A loja identifica o parâmetro e fecha a aba.

**Importante:** A URL de redirecionamento mobile deve ser configurada pelo lado da Serasa como destino após o compartilhamento. No caso do projeto modelo Wake, as URLs foram:
```
Desktop webhook: https://partnerserasa.fbits.store/webhook/token
Mobile redirect: https://partnerserasa.fbits.store/webhook-mobile/
```

---

### FAQ-004 — Como o QR Code funciona na loja? O usuário precisa do app Serasa instalado?

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Projeto Serasa Pass <> Leanwork <> Wake |
| **Data** | 2025-09-03 |
| **Categoria** | `faq` |
| **Tags** | `integração`, `audience-solutions` |

**Pergunta:** Para ler o QR Code precisa de algum SDK ou funcionalidade especial no app? A câmera do celular padrão funciona?

**Resposta:** Sim, a câmera padrão do celular funciona para ler o QR Code, pois o código é um deeplink. Ao escanear o QR Code com qualquer câmera (ou leitor de QR padrão), o sistema operacional detecta que a URL é do app Serasa e redireciona automaticamente para o aplicativo, onde o usuário realiza o compartilhamento das credenciais. Não é necessário SDK adicional para a leitura do QR Code.

O app oficial da Serasa é o único necessário, disponível na App Store e Google Play.

---

### FAQ-005 — Quais eventos de monitoramento devem ser rastreados via JavaScript?

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Projeto Serasa Pass <> Leanwork <> Wake |
| **Data** | 2025-09-22 |
| **Categoria** | `faq` |
| **Tags** | `suporte-escalação`, `audience-solutions` |

**Pergunta:** Quais eventos de comportamento do usuário devem ser rastreados e enviados à API de monitoramento do Serasa?

**Resposta (confirmado pelo time Serasa):**

**Eventos essenciais:**
- Impressão (exibição do componente Serasa Pass)
- Clique no botão
- Visualização do QR Code
- Erro

**Metadados obrigatórios:**
- Tipo de dispositivo (desktop/mobile)
- Viewport
- User-Agent

**Eventos avançados (opcionais):**
- Tempo até o clique
- Porcentagem de visualização do botão

**Tecnologia recomendada:** `sendBeacon` (API nativa dos browsers), ideal para enviar dados de analytics sem bloquear a navegação.

A Serasa pode criar um endpoint específico para receber esses dados via `sendBeacon`.

---

### FAQ-006 — Quando deve ser disparado o endpoint de transação (API Transactions)?

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Projeto Serasa Pass <> Leanwork <> Wake |
| **Data** | 2025-09-22 |
| **Categoria** | `faq` |
| **Tags** | `integração`, `audience-solutions` |

**Pergunta:** A API de Transactions deve ser chamada após a finalização do pedido (fechamento) ou após a aprovação do pagamento (confirmação)?

**Resposta:** O disparo deve ocorrer **após a aprovação do pagamento**, para que o meio de pagamento utilizado seja identificado. Isso garante visibilidade completa sobre a transação, incluindo o método de pagamento.

**Observação adicional:** Em pedidos com frete, foi necessário criar um endpoint adicional para contemplar o valor do frete no envio. O time de desenvolvimento (Igor) já havia implementado essa variação.

---

### FAQ-007 — Como funciona o acesso ao painel administrativo da loja modelo Wake (ambiente sandbox)?

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Projeto Serasa Pass <> Leanwork <> Wake |
| **Data** | 2025-09-01 |
| **Categoria** | `faq` |
| **Tags** | `integração`, `audience-solutions` |

**Dados do ambiente de homologação:**

| Recurso | URL |
|---|---|
| Loja (Storefront) | `https://partnerserasa.fbits.store/` |
| Painel Admin | `https://partnerserasa.fbits.app/` |

**Formato de login:**
- Campo usuário: `partnerserasa\nome.sobrenome` (NÃO usar e-mail como login)
- Senha: Selecionar "Primeiro Acesso" para receber link de redefinição por e-mail

**Atenção:** O sistema bloqueia automaticamente usuários inativos após determinado prazo desde a criação ou reativação. Para reativar, contatar o administrador do painel.

---

### FAQ-008 — O campo `eventName` no redirect mobile retorna quais valores possíveis?

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Projeto Serasa Pass <> Leanwork <> Wake |
| **Data** | 2026-02-02 |
| **Categoria** | `faq` |
| **Tags** | `integração`, `audience-solutions` |

**Pergunta:** No redirect mobile, o campo `eventName` na query string pode retornar valores diferentes de `user_consent`? Como tratar erros no frontend?

**Resposta:** O mapeamento dos possíveis valores do `eventName` está disponível no Swagger da API (enviado como arquivo ZIP por questões de segurança, não há página pública do Swagger). O campo pode retornar diferentes valores em caso de erro ou situações específicas, permitindo que o frontend trate cada caso adequadamente.

Formato do redirect mobile com `eventName`:
```
/webhook-mobile?sessionId={id}&eventName=user_consent
```

Em caso de sucesso, `eventName = user_consent` e `sessionId` contém o ID da sessão do usuário.

---

## Troubleshooting

---

### TRB-001 — Erro de login rápido no checkout da loja (configuração desabilitada)

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Projeto Serasa Pass <> Leanwork <> Wake |
| **Data** | 2025-09-01 |
| **Categoria** | `troubleshooting` |
| **Tags** | `integração`, `audience-solutions`, `audience-cs` |

**Sintomas:** Erro no fluxo de login rápido no checkout da loja Wake.

**Causa raiz:** Uma configuração estava desabilitada no backend da plataforma.

**Resolução:** A configuração foi reabilitada pelo time Wake. Após a correção, o fluxo de login rápido passou a funcionar corretamente.

**Aprendizado:** Sempre verificar configurações de plataforma quando o login rápido não funcionar — o problema pode não ser de código, mas de configuração de feature no painel da plataforma.

---

### TRB-002 — IPs bloqueados impedindo chamadas à API Serasa

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Projeto Serasa Pass <> Leanwork <> Wake |
| **Data** | 2025-09-02 |
| **Categoria** | `troubleshooting` |
| **Tags** | `integração`, `audience-solutions` |

**Sintomas:** Requisições do middleware da Leanwork para a API Serasa retornando erros. Nenhuma request chegando do lado Serasa.

**Causa raiz:** Os IPs do servidor da Leanwork não estavam na whitelist do firewall/infra da Serasa.

**Resolução:**
1. IP específico `40.65.244.176` foi liberado como solução paliativa imediata
2. No dia seguinte, a liberação definitiva de todos os IPs foi realizada pelo time de infra cloud Serasa

**Processo de liberação:**
- Identificar os IPs de saída do ambiente do parceiro (Azure neste caso)
- Solicitar ao time de infra Serasa via canal interno
- Aguardar configuração — pode levar até 1 dia útil para liberação completa

---

### TRB-003 — Erro "Integration error" ao chamar `/v1/verifier/transaction/process`

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Projeto Serasa Pass <> Leanwork <> Wake |
| **Data** | 2025-10-28 |
| **Categoria** | `troubleshooting` |
| **Tags** | `integração`, `audience-solutions` |

**Sintomas:** Ao chamar o endpoint `/v1/verifier/transaction/process`, a API retorna:
```json
{
  "errors": "Integration error id: 964b26b1-8e71-4741-ba18-0f854f0060cd"
}
```

**Causa raiz:** Instabilidade intermitente na infraestrutura da API Serasa. O fluxo retornava a funcionar sozinho após alguns minutos.

**Resolução:** Aguardar normalização. O time Serasa foi notificado para monitorar as intermitências e investigar a causa raiz.

**Observação:** Este mesmo padrão de intermitência se repetiu em outras ocasiões. Recomenda-se implementar retry com backoff exponencial ao chamar este endpoint.

---

### TRB-004 — Tela branca após compartilhamento no fluxo desktop (biometria)

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Projeto Serasa Pass <> Leanwork <> Wake |
| **Data** | 2026-01-08 |
| **Categoria** | `troubleshooting` |
| **Tags** | `integração`, `audience-solutions` |

**Sintomas:** No fluxo Desktop de biometria facial — após o usuário ler o QR Code com o celular, realizar a biometria e compartilhar os dados — a aba de compartilhamento que estava aberta no desktop ficava em branco.

**Causa raiz:** Comportamento esperado era fechar ou notificar o usuário após o webhook ser recebido. A tela branca indicava que o webhook foi recebido mas o frontend não tratou o redirecionamento/fechamento adequadamente.

**Resolução:** Correção implementada pela Leanwork no frontend para tratar o retorno do webhook e redirecionar/fechar a aba corretamente. Os dados chegaram corretamente na Wake mesmo durante o período do bug.

---

### TRB-005 — Fluxo mobile não batia no webhook após compartilhamento

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Projeto Serasa Pass <> Leanwork <> Wake |
| **Data** | 2026-01-09 |
| **Categoria** | `troubleshooting` |
| **Tags** | `integração`, `audience-solutions` |

**Sintomas:** No fluxo mobile (biometria via URL), após o compartilhamento no app Serasa, o webhook de notificação não era acionado. Os dados não chegavam na loja Wake.

**Causa raiz:** Em mobile, a aba da Wake fica suspensa enquanto o usuário vai para o app Serasa. Mesmo que a Serasa bata no webhook, a aba suspensa não processa o evento.

**Resolução:** Redirect com `sessionId` e `eventName` via query string na URL de retorno mobile. A URL de redirect foi atualizada para o formato abaixo. Alternativas discutidas (não adotadas): (1) Polling no endpoint `/shared` — a loja fica consultando periodicamente se as credenciais foram compartilhadas; (2) Redirect com queryString ao finalizar o fluxo, sucesso ou falha.
```
/webhook-mobile?sessionId={id}&eventName=user_consent
```

---

### TRB-006 — Redirect mobile bloqueado pelo Safari como popup

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Projeto Serasa Pass <> Leanwork <> Wake |
| **Data** | 2025-10-14 |
| **Categoria** | `troubleshooting` |
| **Tags** | `integração`, `audience-solutions` |

**Sintomas:** No iPhone com Safari, o redirect após compartilhamento mobile era bloqueado como se fosse um popup. O usuário não era redirecionado de volta à loja.

**Causa raiz:** O componente de botão Serasa era implementado como `<button>` com `window.open()` via JavaScript, e o Safari por padrão bloqueia abertura de janelas/abas geradas por JavaScript.

**Resolução:** No Chrome o comportamento foi corrigido. Para resolver definitivamente no Safari, refatorar o componente para usar link HTML nativo (`<a href="" target="_blank">`) em vez de JavaScript para a abertura.

**Workaround:** Usar um link HTML nativo (`<a href="" target="_blank">`) em vez de JavaScript para a abertura. Links HTML nativos não são bloqueados pelo Safari como popup.

**Obstáculo:** O card do Serasa Pass era um `<button>` e as mudanças necessárias para desktop e mobile precisavam ser via JS. Refatoração foi necessária para resolver definitivamente.

**Lição aprendida:** Para links que precisam abrir em nova aba em ambientes móveis (especialmente iOS/Safari), sempre preferir `<a>` HTML nativo ao invés de `window.open()` via JavaScript.

---

### TRB-007 — E-mail de boas-vindas sendo disparado para clientes do evento

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Projeto Serasa Pass <> Leanwork <> Wake |
| **Data** | 2025-09-10 |
| **Categoria** | `troubleshooting` |
| **Tags** | `integração`, `audience-solutions` |

**Sintomas:** Durante o evento Make Together (demonstração presencial), visitantes que cadastravam-se na loja modelo recebiam e-mail de boas-vindas da plataforma Wake — comportamento indesejado durante uma demo.

**Causa raiz:** A loja modelo estava com os e-mails transacionais habilitados por padrão.

**Resolução:** O time Leanwork editou a referência do arquivo de template de e-mail no repositório Storefront. O ajuste foi publicado e verificado antes do evento. O e-mail de confirmação de compra continuou funcionando normalmente (comportamento desejado).

**Como fazer:** A desabilitação do e-mail de boas-vindas requer editar o arquivo HTML correspondente na pasta do Storefront no repositório Git — não é configuração do painel.

---

### TRB-008 — Erro na validação de e-mails no formulário de cadastro durante evento

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Projeto Serasa Pass <> Leanwork <> Wake |
| **Data** | 2025-09-10 |
| **Categoria** | `troubleshooting` |
| **Tags** | `integração`, `audience-solutions` |

**Sintomas:** Durante o evento, alguns e-mails de cadastro retornavam como inválidos mesmo sendo e-mails reais.

**Causa raiz:** O JavaScript de validação de e-mails na loja possivelmente rejeitava domínios desconhecidos ou de empresas (e.g., `usuario@empresa.com.br`). Domínios comuns (`@gmail.com`, `@hotmail.com`) funcionavam.

**Resolução:** [fonte original sem informação] — correção definitiva do regex de validação não documentada nas conversas registradas.

**Workaround durante o evento:** Usar e-mails de domínios conhecidos (Gmail, Hotmail) ao cadastrar durante a demonstração.

**Observação:** Pode existir uma política de bloqueio contra múltiplas tentativas de acesso do mesmo IP (proteção DDOS), que poderia ter contribuído com o comportamento.

---

## How-To

---

### HTO-001 — Como configurar o middleware Wake-Serasa Pass em nova infraestrutura

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Projeto Serasa Pass <> Leanwork <> Wake |
| **Data** | 2026-04-30 |
| **Categoria** | `how_to` |
| **Tags** | `integração`, `audience-solutions` |

**Contexto:** Ao integrar Serasa Pass em nova loja Wake (para um cliente), o middleware precisa ser configurado em infraestrutura dedicada.

**Passos:**

1. **Obter o código-fonte:**
   - O repositório do middleware e frontend é gerenciado pela Leanwork
   - Para transferência para infra do cliente/Serasa, solicitar acesso via gestão da Leanwork
   - Compartilhamento por repositório Git (acesso por e-mail) ou pacote seguro

2. **Setup de infraestrutura (estimativa Leanwork):**
   - 5 dias para configuração de infra (Azure)
   - 2 dias para deploy das aplicações (back + front)
   - 1 dia para testes integrados

3. **Configurar URLs na Serasa:**
   - Webhook desktop: `https://{loja}/webhook/token` (POST com `sessionId` no body)
   - Redirect mobile: `https://{loja}/webhook-mobile?sessionId={id}&eventName=user_consent`
   - Informar as URLs ao time Serasa para cadastro no sistema

4. **Liberar IPs:**
   - Identificar IPs de saída da infraestrutura do parceiro
   - Solicitar liberação ao time de infra Serasa
   - Aguardar confirmação antes de iniciar testes de integração

5. **Configuração no painel Wake (por agência implementadora):**
   - Instalar o plugin Serasa Pass na loja
   - Para lojas Storefront: a agência tem autonomia para editar o template
   - Para lojas Style Editor: configuração diferente — validar com time Wake

6. **Teste final:**
   - Testar fluxo completo desktop (QR Code + compartilhamento no app)
   - Testar fluxo mobile (redirect direto para app)
   - Verificar recebimento dos webhooks
   - Validar preenchimento do formulário com dados retornados

---

### HTO-002 — Como criar um usuário no painel administrativo Wake (ambiente parceiro)

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Projeto Serasa Pass <> Leanwork <> Wake |
| **Data** | 2026-05-04 |
| **Categoria** | `how_to` |
| **Tags** | `integração`, `audience-solutions`, `audience-cs` |

**Pré-requisito:** Ter acesso de administrador ao painel Wake.

**Passos:**
1. Acessar `https://partnerserasa.fbits.app/`
2. Qualquer usuário administrador pode criar novos usuários
3. No campo login, usar o formato: `partnerserasa\nome.sobrenome` (NÃO usar o e-mail)
4. Senha: o novo usuário deve clicar em "Primeiro Acesso" para receber link de redefinição por e-mail

**Formato de usuários:**
```
partnerserasa\nome.sobrenome
```

**Atenção:** Usuários inativos por muito tempo são bloqueados automaticamente. Para reativar, o administrador deve acessar o painel e reativar o usuário manualmente.

---

### HTO-003 — Como implementar a detecção de dispositivo e criação de sessão com type correto

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Projeto Serasa Pass <> Leanwork <> Wake |
| **Data** | 2025-09-24 |
| **Categoria** | `how_to` |
| **Tags** | `integração`, `audience-solutions` |

**Objetivo:** Criar a sessão Serasa Pass com o `type` correto baseado no dispositivo do usuário.

**Lógica de implementação:**

```javascript
// Detectar device
const isMobile = /Android|iPhone|iPad|iPod/i.test(navigator.userAgent);

// Criar sessão com type adequado
const sessionType = isMobile ? 2 : 1;

// Type 1 = Desktop (retorna QR Code + webhook ao final)
// Type 2 = Mobile (retorna redirect direto para app)
```

**Fluxo Desktop (type 1):**
- Criar sessão → receber `webUrl` com sessionId
- Exibir QR Code gerado a partir da URL
- Aguardar webhook POST em: `/{rota-configurada}` com `sessionId` no body
- Ao receber o webhook, buscar dados compartilhados via API de shared

**Fluxo Mobile (type 2):**
- Criar sessão → receber URL de redirect
- Redirecionar usuário para o app Serasa
- Aguardar redirect de volta para: `/webhook-mobile?sessionId={id}&eventName=user_consent`
- Ao receber o redirect, buscar dados compartilhados via API de shared

**Importante:** O fluxo deve ser executado sequencialmente. Não pular etapas — o app precisa validar a sessão antes de qualquer outra operação.

---

### HTO-004 — Como testar o blocklist (bloqueio de CPF/dispositivo) no Serasa Pass

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Projeto Serasa Pass <> Leanwork <> Wake |
| **Data** | 2026-02-09 |
| **Categoria** | `how_to` |
| **Tags** | `integração`, `audience-solutions` |

**Objetivo:** Testar o comportamento da integração quando um CPF ou dispositivo está na blocklist do Serasa Pass.

**Passos:**
1. Fornecer ao time Serasa os dados do dispositivo/CPF a ser bloqueado:
   - CPF do usuário de teste
   - Modelo do celular
2. Aguardar confirmação do bloqueio pelo time Serasa
3. Executar o fluxo completo de compartilhamento com o CPF/dispositivo bloqueado
4. Verificar se o comportamento de erro é tratado adequadamente pela integração

**Tempo de execução do bloqueio:** Geralmente realizado no mesmo dia pelo time Serasa.

---

### HTO-005 — Como solicitar a API Key para nova integração de cliente

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Projeto Serasa Pass <> Leanwork <> Wake |
| **Data** | 2026-06-01 |
| **Categoria** | `how_to` |
| **Tags** | `integração`, `audience-solutions` |

**Processo para obter API Key para um novo cliente:**

1. Enviar e-mail para o time Serasa com as seguintes informações:
   - IPs de saída da infraestrutura do cliente (para whitelist no firewall)
   - URLs de webhook (desktop e mobile redirect)
   - Domínio da loja

2. O time Serasa configura:
   - Whitelist de IPs
   - URLs de redirect/webhook
   - Gera a API Key para o cliente

3. Após liberação dos IPs, realizar teste de autenticação (chamada à API deve retornar 200, não 401 Unauthorized)

4. Com a API Key configurada, o ambiente está pronto para a agência implementadora configurar o plugin no painel Wake

**Para validar se a liberação foi feita:** Fazer uma requisição de teste via Postman contra qualquer endpoint da API Serasa e confirmar que não retorna `401 Unauthorized`.

---

## Best Practices

---

### BPR-001 — Separação de branches para demonstrações e desenvolvimento ativo

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Projeto Serasa Pass <> Leanwork <> Wake |
| **Data** | 2025-09-17 |
| **Categoria** | `best_practices` |
| **Tags** | `integração`, `audience-solutions` |

**Recomendação:** Manter uma branch de demonstração separada (com dados mock/fake estáveis e link de preview fixo) para apresentações a clientes/prospects, enquanto a branch principal evolui com novas funcionalidades. Avisar o time de desenvolvimento antes de apresentações e compartilhar sempre o link da branch de demonstração, nunca o da branch de desenvolvimento.

**Rationale:** O projeto Serasa Pass tinha uma loja modelo funcional com dados mock/fake (fase 1) que a Serasa queria usar para apresentações a clientes, enquanto o desenvolvimento da fase 2 continuava. Sem separação, clientes e prospects veriam funcionalidades em desenvolvimento ou parcialmente implementadas. Com a separação, eles sempre veem a versão estável e apresentável.

---

### BPR-002 — Gestão de ambientes de testes vs. produção em projetos piloto

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Projeto Serasa Pass <> Leanwork <> Wake |
| **Data** | 2025-09-26 |
| **Categoria** | `best_practices` |
| **Tags** | `homologação`, `audience-solutions` |

**Recomendação:** Em projetos piloto, sempre confirmar explicitamente em qual ambiente (HML ou PRD) as chamadas estão sendo feitas antes de iniciar testes, priorizar homologação mesmo quando demora mais, garantir que dados de teste em PRD sejam identificáveis e estabelecer períodos de freeze de deploy antes de eventos importantes.

**Rationale:** Durante o projeto piloto Serasa Pass, a integração foi desenvolvida diretamente em produção da API Serasa, pois havia dificuldades para liberar ambiente de homologação. Lições aprendidas:
1. **Confirmar sempre o ambiente:** Antes de iniciar testes, verificar explicitamente se as chamadas estão indo para HML ou PRD. Perguntar diretamente à equipe responsável.
2. **Priorizar homologação:** Mesmo que demore mais, usar ambiente de HML evita impactos em dados reais e clientes.
3. **Dados de teste em PRD:** Se precisar usar PRD para testes (como foi o caso), garantir que os dados de teste sejam claramente identificáveis e não poluam relatórios de produção.
4. **Freeze de deploy:** Estabelecer períodos de freeze antes de eventos ou apresentações importantes, como o time Wake fez antes do evento Make Together.

---

### BPR-003 — Tratamento de dados de usuário na integração Wake-Serasa Pass

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Projeto Serasa Pass <> Leanwork <> Wake |
| **Data** | 2026-03-12 |
| **Categoria** | `best_practices` |
| **Tags** | `integração`, `audience-solutions` |

**Recomendação:** Na integração Wake × Serasa Pass, não substituir o e-mail digitado pelo usuário pelo da Serasa, usar default `"outros"` para Gênero quando o campo for obrigatório, atualizar dados do usuário com a credencial após compartilhamento, e desmembrar a credencial de contato para evitar sobrescritas.

**Rationale:** (regras e decisões tomadas no projeto)

1. **Conflito de e-mail:** O e-mail digitado pelo usuário no início do login simplificado NÃO deve ser substituído pelo e-mail retornado pela Serasa (credencial). A substituição causava erros de duplicidade de e-mail na plataforma.

2. **Campo Gênero:** Como a Serasa Pass não retorna informação de gênero, a decisão foi configurar o valor default como `"outros"` quando o campo for obrigatório na loja.

3. **Atualização de dados pós-compartilhamento:** Após o compartilhamento das credenciais, os dados do usuário na loja DEVEM ser atualizados com as informações retornadas pelo Serasa Pass. Isso é requisito do produto — sem a atualização, há problemas com parceiros que dependem dos dados validados.

4. **Desmembramento da credencial de contato:** A credencial de contato deve ser desmembrada (tratada separadamente) para evitar sobrescrever dados que o usuário já inseriu na loja.

---

### BPR-004 — Arquitetura do middleware Wake-Serasa Pass

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Projeto Serasa Pass <> Leanwork <> Wake |
| **Data** | 2025-09-01 |
| **Categoria** | `best_practices` |
| **Tags** | `integração`, `audience-solutions` |

**Recomendação:** Implementar um middleware backend (hospedado em Azure ou equivalente) entre a loja Wake e a API Serasa, responsável por autenticação (API Key), gestão de sessões e recebimento de webhooks. Após o piloto, transferir o middleware para a infraestrutura do cliente/Serasa.

**Rationale:** arquitetura utilizada no projeto piloto:**

O projeto implementou um **middleware** (backend intermediário) entre a loja Wake e a API Serasa, hospedado na Azure. O middleware tem as seguintes responsabilidades:

- **Autenticação:** Guarda as credenciais de acesso à API Serasa (API Key), nunca expostas no frontend
- **Gestão de sessões:** Cria sessões Serasa Pass e retorna a URL/QR Code ao frontend
- **Recebimento de webhooks:** Endpoint para receber notificações de compartilhamento da Serasa
- **Endpoint de user consent:** `/api/serasa/webhook-user-consent` para notificações do fluxo de compartilhamento

**Endpoint de webhook do middleware:**
```
POST https://webapp-middleware-wake-serasa-pass.azurewebsites.net/api/serasa/webhook-user-consent
Content-Type: application/json
{
  "sessionId": "1233231"
}
```

**Motivo do middleware:** A API Serasa requer whitelist de IPs. Centralizar as chamadas no middleware facilita a gestão de IPs e evita expor credenciais no frontend.

**Transferência de propriedade:** O middleware foi desenvolvido pela Leanwork e deve ser transferido para a infraestrutura do cliente/Serasa ao produtizar a integração. O código-fonte é entregue via repositório Git (acesso controlado por segurança).

---

### BPR-005 — Modelo de sustentação e manutenção do plugin após entrega

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Projeto Serasa Pass <> Leanwork <> Wake |
| **Data** | 2026-03-23 |
| **Categoria** | `best_practices` |
| **Tags** | `suporte-escalação`, `audience-pam`, `audience-cs` |

**Recomendação:** Em projetos de plugin/integração com parceiro tech (Leanwork × Serasa × Wake), estruturar comercial em três camadas: (1) sustentação mensal obrigatória com horas/mês para melhorias e bugs quando a SI hospeda infra, (2) setup único por ambiente, (3) infra cobrada mensalmente em plano Básico ou Recomendado.

**Rationale:** Após a entrega do projeto piloto, a Serasa precisava de um modelo para continuar evoluindo e mantendo o plugin em produção para novos clientes Wake. O modelo adotado:
- **Proposta de sustentação mensal:** Inclui horas/mês para melhorias contínuas, correção de bugs e suporte — obrigatório quando a Leanwork hospeda a infraestrutura
- **Setup único:** Cobrado uma única vez por ambiente configurado (não recorrente)
- **Infra:** Cobrada mensalmente (plano Básico ou Recomendado); valores em moeda estrangeira sujeitos a variação cambial

**Processo de aprovação interna Serasa:**
- Proposta passa pelo time de compras (PR)
- Aprovação financeira necessária para setup e mensalidade
- Contrato/aditivo formal antes de qualquer trabalho adicional

**Para novos clientes Wake:**
- A Serasa é responsável por indicar o cliente à Leanwork
- A agência implementadora do e-commerce (ex.: b8one) é responsável pela configuração do plugin no painel Wake
- A Leanwork cuida do setup de infra e testes de integração
- Prazo estimado total (após aprovação): 8 dias (Leanwork) + prazo da agência

---

### BPR-006 — Padrão de comunicação e gestão de projeto em parceria tripartite

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Projeto Serasa Pass <> Leanwork <> Wake |
| **Data** | 2025-09-03 |
| **Categoria** | `best_practices` |
| **Tags** | `marketing-eventos`, `audience-pam` |

**Recomendação:** Em projetos tripartites (parceiro + integrador + Wake), estabelecer daily checkpoint diário, grupo WhatsApp para assíncrono, e-mail para formalização, e definir papéis claros sobre responsabilidade por cada item de infraestrutura para evitar ambiguidades.

**Rationale:** modelo de comunicação adotado no projeto:**

1. **Daily checkpoint:** Reunião diária (10h30) com representantes dos três lados (Serasa, Leanwork, Wake) para alinhamento de status e bloqueios
2. **Grupo WhatsApp:** Canal para comunicações assíncronas, impedimentos urgentes e compartilhamento de arquivos
3. **E-mail:** Formalização de atas, propostas comerciais e itens que precisam de rastreabilidade
4. **Frequência adaptativa:** Nos momentos de menor atividade (análise/estimativa), o checkpoint foi suspenso e substituído por updates assíncronos no grupo

**Papéis e responsabilidades:**
- **Serasa:** Define os requisitos do produto, valida entregas, fornece API Key e documentação
- **Leanwork:** Desenvolvimento do middleware e plugin Wake, gestão de projeto técnico
- **Wake:** Suporte à plataforma, configurações do painel, acesso aos repositórios do Storefront

**Lição aprendida:** Em projetos com três empresas distintas, é crítico que cada parte tenha clareza sobre quem é responsável por cada item de infraestrutura e configuração. Ambiguidades sobre "quem cuida de quê" causaram atrasos repetidos no projeto.

---

# KB Extract — Wave 6B (Serasa, Estoca, Konfidency, CRMBonus)
> Extracted: 2026-06-04

---

## Grupo: Serasa & Wake

**JID:** `120363401681102953@g.us`

---

### HTO-001 — Fluxo de integração do Serasa Pass na jornada de compra Wake

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Serasa & Wake |
| **Data** | 2025-08-18 |
| **Categoria** | `how_to` |
| **Tags** | `integração`, `audience-solutions` |

A integração do Serasa Pass na jornada de compra Wake foi definida em três fases:

- **Fase 1 (MVP):** Fluxo de preenchimento de dados com autenticação via aplicativo Serasa — sem cartão físico e sem emissão de link de biometria.
- **Fase 2:** Fluxo de biometria sem o app (link de biometria disponível como opção standalone).
- **V2 (futuro):** Inclusão do fluxo com cartão.

A integração exigiu envolvimento de um parceiro integrador (Leanwork) por não ser configurável apenas via painel admin — necessita desenvolvimento técnico. O perfil ideal do e-commerce para adoção inicial é aquele com time técnico próprio e porte médio a grande.

---

### FAQ-001 — Webhook vs. Polling para dados do Serasa Pass

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Serasa & Wake |
| **Data** | 2025-08-19 |
| **Categoria** | `faq` |
| **Tags** | `integração`, `audience-solutions` |

**Pergunta:** Para buscar dados do cliente no Serasa Pass, é obrigatório usar Webhook ou é possível fazer polling periódico para verificar disponibilidade?

**Resposta:** Ambos os modos são possíveis. A escolha do caminho mais adequado depende do caso de uso:
- Para o evento inicial de verificação (checar se o dado está disponível), o polling pode ser usado.
- Para o fluxo completo de dados após confirmação, o Webhook é o direcionamento padrão.

O time do Serasa Pass deve ser consultado para validar qual abordagem específica atende melhor o fluxo da loja.

---

### TRB-001 — Dificuldade de conexão real via API do Serasa Pass

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Serasa & Wake |
| **Data** | 2025-09-05 |
| **Categoria** | `troubleshooting` |
| **Tags** | `integração`, `audience-solutions` |

**Sintomas:** Durante o desenvolvimento da integração do Serasa Pass, o time da Leanwork encontrou dificuldades para estabelecer conexão real via API.

**Causa raiz:** [fonte original sem informação] — causa técnica específica da falha de conexão não foi documentada na conversa registrada.

**Resolução:** Foi adotado um plano B com dados fictícios enquanto o problema era resolvido. A previsão foi: finalizar o fluxo via API ainda naquele dia e deixar apenas ajustes de layout para a semana seguinte.

**Lição:** Ao integrar com parceiros externos como o Serasa Pass via integradores como a Leanwork, é recomendável ter um ambiente de demonstração paralelo (dados fictícios) para garantir continuidade no desenvolvimento e nas apresentações.

---

### HTO-002 — Processo de assinatura de contrato com a Serasa (cláusula de limitação de responsabilidade)

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Serasa & Wake |
| **Data** | 2026-03-24 |
| **Categoria** | `how_to` |
| **Tags** | `contrato-comercial`, `audience-pam` |

Um bloqueador recorrente na assinatura do contrato com a Serasa foi a cláusula de limitação de responsabilidade. O jurídico da Wake exige que o valor da cláusula seja determinado para prosseguir — não é possível assinar sem o valor e a cláusula não pode ser removida.

**Como resolver:**
1. O jurídico da Serasa deve propor um valor (pode ser simbólico: R$ 5k, R$ 10k, R$ 15k).
2. O jurídico da Wake aceita valores abaixo de R$ 50k.
3. Uma vez definido o valor, os demais pontos do contrato são tratados como detalhes menores e negociados diretamente entre os times.

**Resultado:** A Wake liberou a assinatura do contrato após o alinhamento desta cláusula em abril de 2026.

---

### FAQ-002 — Comprovação de poderes de representantes para assinatura contratual

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Serasa & Wake |
| **Data** | 2025-09-26 |
| **Categoria** | `faq` |
| **Tags** | `contrato-comercial`, `audience-pam` |

**Pergunta:** O jurídico da Wake solicitou documento comprovando que os representantes da Serasa/Experian são diretora e presidente. Como obter?

**Resposta:** O site da Experian (seção de Relações com Investidores - RI) contém a estrutura do conselho e diretores. Caso seja necessária uma procuração formal, o ponto de contato interno da Experian deve ser acionado. O próprio time de relações internas pode agilizar o documento.

---

### BPR-001 — Ativação do Serasa Pass no Wake Summit: estratégia de stand e palco

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Serasa & Wake |
| **Data** | 2025-09-05 |
| **Categoria** | `best_practices` |
| **Tags** | `marketing-eventos`, `audience-pam` |

**Recomendação:** Para parceiros patrocinadores de eventos Wake, combinar três ativações: (1) Palco Make Together de 20min para awareness, (2) Stand com demo ao vivo focada em capturar leads quentes, (3) Programação de lead capture com CS Wake indicando leads quentes ao stand.

**Rationale:** Combinação testada com Serasa Pass no Wake Summit. O palco principal do evento (keynote) tem lineup fechado com antecedência — parceiros ingressantes no programa não devem contar com esse espaço para o primeiro evento, então Make Together + stand é a alavanca disponível.

---

### HTO-003 — Engajamento de e-commerces Wake como primeiros clientes do Serasa Pass

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Serasa & Wake |
| **Data** | 2026-01-29 |
| **Categoria** | `how_to` |
| **Tags** | `integração`, `audience-pam`, `audience-sales` |

Para a ativação inicial da base de e-commerces Wake como clientes do Serasa Pass:

1. O time de CS Wake deve ser treinado no produto antes das indicações formais.
2. O gerente de parcerias Wake apresenta internamente o produto a gerentes de vendas em reuniões 1:1 antes de conectar com prospects.
3. A conexão com prospects é feita diretamente pelo gerente de parcerias (sem aguardar acionamento — "nem vou aguardar me acionarem").
4. A Konfidency mencionou a importância de ter ~20 contas em comum com a Wake para criar uma abordagem de casos de uso semanais para CS.
5. O primeiro "SIM" foi celebrado como marco simbólico — recomendável criar rituais de reconhecimento para o CS que indicou.

---

## Grupo: Estoca & Wake

**JID:** `120363420090271403@g.us`

---

### HTO-004 — Onboarding de parceiro de patrocínio no Wake Summit (cota Silver + Experiência)

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Estoca & Wake |
| **Data** | 2025-08-20 |
| **Categoria** | `how_to` |
| **Tags** | `marketing-eventos`, `audience-pam` |

Processo de onboarding de um novo patrocinador (Estoca, cota Silver + Experiência Wake Store):

1. **Negociação comercial:** Definir cota e tipo de ativação. A Estoca negociou budget de R$ 100k para cota Silver + Experiência de entrega (empacotamento de brindes com caixa personalizada).
2. **Alinhamento de marketing:** Confirmar detalhes da ativação com o time de marketing Wake antes de fechar.
3. **Contrato:** Jurídico Wake elabora a minuta e envia para o jurídico do parceiro (SLA de revisão: 3 dias úteis). O jurídico do parceiro pode solicitar proposta comercial/Mídia Kit.
4. **Envio de materiais:** O parceiro envia logos em todos os formatos (vetorial + guia de marca). Wake envia o Manual do Expositor com prazos de entrega de artes.
5. **Prazos do Manual do Expositor (exemplo Summit 2025):**
   - Arte da lona: até 01/09
   - Arte do balcão: até 01/09
   - Vídeo para telão: até 05/09
6. **Retirada de brindes:** Coordenar com o parceiro sobre retirada física no escritório Wake (ex: LWSA, Rua Itapaiúna, 2434 - Parque do Morumbi).
7. **Lista de participantes:** Parceiro deve enviar planilha com leads e staffs credenciados.

---

### FAQ-003 — Diferença entre número de participantes e quantidade de brindes no Wake Summit

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Estoca & Wake |
| **Data** | 2025-09-02 |
| **Categoria** | `faq` |
| **Tags** | `marketing-eventos`, `audience-pam` |

**Pergunta:** Por que a proposta indica empacotamento de 1.200 brindes se o evento tem ~700 participantes?

**Resposta:** O número de 700 é a expectativa de presença de patrocinadores, clientes e prospects. O estoque da loja virtual do site possui capacidade para 1.200 brindes (número de inscritos), mas há uma quebra natural (pessoas que não comparecem, saem antes etc.). Como os brindes são diversificados, não é possível estimar previamente quais sobrarão — por isso, o volume total disponível é entregue para empacotamento. O parceiro pode solicitar atualização da proposta para 700 unidades se preferir trabalhar com o número estimado de presença efetiva.

---

### BPR-002 — Follow-up pós-evento: como usar o mailing de participantes

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Estoca & Wake |
| **Data** | 2025-09-15 |
| **Categoria** | `best_practices` |
| **Tags** | `integração`, `audience-pam`, `audience-sales` |

**Recomendação:** Após o Wake Summit, patrocinadores devem solicitar o mailing de participantes ao time de eventos Wake (SLA até 5 dias úteis) e definir internamente antes do evento o responsável pelo follow-up e a cadência de abordagem para não perder o timing pós-evento.

**Rationale:** Priorizar o follow-up imediato aproveita o momento de alta receptividade dos leads que visitaram o stand — quanto mais tempo passa após o evento, menor a recall e taxa de resposta.

---

### HTO-005 — Processo para gravação de podcast no Wake Summit

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Estoca & Wake |
| **Data** | 2025-10-24 |
| **Categoria** | `how_to` |
| **Tags** | `marketing-eventos`, `audience-pam` |

Para parceiros que participaram do Wake Summit com gravação de podcast:

1. Os episódios são publicados em sequência — aproximadamente **1 por semana**.
2. A ordem de publicação considera Black Friday e prioridades de comunicação do calendário Wake.
3. O parceiro deve ser avisado com antecedência sobre a data de publicação para coordenar a divulgação nas próprias redes sociais.
4. Recomendável preparar a programação de divulgação do parceiro com base na previsão Wake.

---

### BPR-003 — ICP e expectativas da parceria Estoca & Wake (2026)

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Estoca & Wake |
| **Data** | 2026-03-17 |
| **Categoria** | `best_practices` |
| **Tags** | `lead-pipeline`, `audience-pam`, `audience-sales` |

**Recomendação:** Direcionar a parceria Estoca & Wake para o ICP combinado: GMV R$ 2M–4M/mês, 5.000–10.000 pedidos/mês, faturamento digital R$ 1M–2M/mês, segmentos prioritários Cosméticos e Suplementos. Usar a target list 2026 acordada com a Estoca como ponto de partida para indicações.

**Rationale:** Perfil ideal definido em conjunto entre Estoca e Wake. Target list 2026 da Estoca: Nutrata, Loccitane, Sallve, Puravida, Guday, AURA Beauty, +Mu, Beyoung, Principia, DUX, Linus, BOB, Braé, Liquidz, HeyMu, Soldiers Nutrition. Expectativas da Estoca para 2026: construir canal de parcerias sólido como um dos pilares estratégicos do ano; a Wake foi vista como solução complementar com tecnologia de ponta para entregar proposta de valor conjunta ao cliente.

---

### HTO-006 — Formalização de entrada no programa Elite Wake (parceiro tech)

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Estoca & Wake |
| **Data** | 2026-04-07 |
| **Categoria** | `how_to` |
| **Tags** | `contrato-comercial`, `audience-pam` |

Para formalizar a entrada de um parceiro tech no programa Elite Wake:

1. **Confirmação comercial:** O time de vendas/parcerias Wake confirma a decisão (ex: entrada no Elite + cota Gold de evento pós-VTEX).
2. **Responsável pelo contrato:** Um executivo de parcerias Wake é designado para cuidar do contrato.
3. **Dados cadastrais necessários do parceiro:**
   - Razão Social
   - Endereço completo
   - CNPJ e IE
   - Nome e e-mail do representante legal
   - Nome e e-mail de uma testemunha
4. **Envio de logos:** Parceiro envia logos (PNG fundo branco e principal) para uso em materiais Wake.
5. **Staffs para eventos:** Definir nomes completos dos participantes para credenciamento.

---

## Grupo: Konfidency & Wake

**JID:** `120363413122423249@g.us`

---

### TRB-002 — Script da Konfidency não renderiza em todas as páginas do site Wake

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Konfidency & Wake |
| **Data** | 2025-11-07 |
| **Categoria** | `troubleshooting` |
| **Tags** | `integração`, `audience-solutions`, `audience-cs` |

**Sintomas:** O script da Konfidency adicionado via API na plataforma Wake não é renderizado em todas as páginas do site do cliente — em alguns casos, apenas na home o script carregava. O lojista liberou o domínio nas políticas de Content Security Policy (CSP), mas o problema persistiu.

**Causa raiz:** A plataforma Wake possui uma política de segurança que bloqueia scripts não autorizados por padrão. A liberação precisa ser feita manualmente no admin para cada lojista.

**Resolução:** Liberação manual no admin Wake para cada lojista (estado atual). Em andamento (2025-09): o time de produto Wake avaliou a possibilidade de adicionar o subdomínio `reviews.konfidency.com.br` em uma whitelist geral, eliminando a necessidade de liberação manual cliente a cliente. A Konfidency confirmou que apenas esse subdomínio é utilizado pelos scripts.

**Status:** Em discussão no time de engenharia Wake. Enquanto não houver whitelist global, o processo de onboarding de cada novo cliente requer uma etapa manual de liberação no admin.

---

### BPR-004 — Gestão de parceiros concorrentes no ecossistema Wake: impacto nas indicações

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Konfidency & Wake |
| **Data** | 2025-08-22 |
| **Categoria** | `best_practices` |
| **Tags** | `ecossistema-parceiros`, `audience-pam` |

**Recomendação:** Quando houver parceiros concorrentes no programa Wake, definir critérios claros de direcionamento (por segmento, porte ou exclusividade), comunicar ao time de CS antes do lançamento de cada novo parceiro concorrente, e garantir que parceiros com mensalidade tenham vantagem de indicação proporcional ao investimento.

**Rationale:** A situação de ter dois parceiros concorrentes diretos no programa (Konfidency e Vurdere, ambos de reviews) gerou atrito significativo:

- CS Wake passou a indicar ambos os parceiros ao mesmo cliente sem critério claro, esvaziando o valor consultivo da indicação.
- Um lead enviado para ambos fechou com o concorrente, evidenciando falta de prioridade no direcionamento.
- A Wake não pode orientar nenhum dos parceiros a melhorar competitivamente sem gerar desigualdade.
- Parceiros que investem mensalidade no programa ficam em desvantagem em relação a parceiros sem mensalidade quando ambos recebem o mesmo número de indicações.

---

### FAQ-004 — O que configura uma "indicação qualificada" no programa de parceiros Wake

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Konfidency & Wake |
| **Data** | 2025-09-17 |
| **Categoria** | `faq` |
| **Tags** | `lead-pipeline`, `audience-pam` |

**Pergunta (parceiro):** O que foi prometido nas conversas iniciais sobre o tipo de indicação que a Wake geraria?

**Resposta consolidada das conversas:** A Wake não se posicionou como geradora de leads frios. O compromisso foi indicar empresas que:
1. Tenham dor mapeada pelo CS ou time de vendas
2. Tenham real interesse no produto do parceiro
3. Estejam abertas a pelo menos agendar uma reunião

**Problemas documentados com indicações:**
- Leads frios (sem interesse, sem abertura para reunião)
- Indicação de clientes já na base do parceiro (duplicatas)
- Re-indicação de clientes já indicados anteriormente pelo próprio CS Wake

**Boa prática:** O time de CS deve verificar na base do parceiro antes de fazer uma indicação para evitar duplicatas. Parceiros devem fornecer lista de clientes ativos para cruzamento.

---

### HTO-007 — Estratégia de ativação para Black November com parceiro de reviews

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Konfidency & Wake |
| **Data** | 2025-11-06 |
| **Categoria** | `how_to` |
| **Tags** | `marketing-eventos`, `audience-pam` |

Estratégia desenvolvida entre Konfidency e Wake CS para Black November:

**Próximos passos definidos:**
1. **Material Konfidency para CS (prazo: 10/11):** One-pager com CTA específico — "Garanta a reputação da sua marca e aproveite a Black Friday para reunir avaliações sobre o seu negócio + oportunidade de trial até 05/01/26".
2. **Lista de clientes potenciais (prazo: 12/11):** CS Commerce + CS Experience levantam relação segmentada de prospects por área.
3. **Integração Wake Experience + Konfidency:** Fase 1 — export/import de dados em CSV/sheets; Fase 2 — integração homologada.

**Tática de conteúdo:** Para os ~20 clientes em comum entre Wake e Konfidency, criar pílulas semanais de conteúdo mostrando como a Konfidency está atendendo clientes Wake específicos — case privado direcionado ao CS responsável pela conta.

---

### TRB-003 — Problema de faturamento retroativo no programa de parcerias

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Konfidency & Wake |
| **Data** | 2025-07-02 |
| **Categoria** | `troubleshooting` |
| **Tags** | `contrato-comercial`, `audience-pam` |

**Sintomas:** O financeiro da Wake enviou NFs de "Partner Fee" de meses anteriores ao kickoff e à assinatura do contrato:
- Partner Fee Março - R$ 4.167 (NF 52023)
- Partner Fee Abril - R$ 4.167 (NF 52681)
- Partner Fee Maio - R$ 4.167 (NF 53238)

O kickoff ocorreu em 09/04 e o contrato foi assinado em meados de abril. As NFs de março e abril não faziam sentido cronologicamente.

**Causa raiz:** Financeiro Wake não foi briefado com a data exata de início do contrato, gerando faturamento retroativo a meses anteriores à assinatura.

**Resolução:** Escalado para o time de parcerias para verificação e ajuste. A NF de maio era a máxima que poderia ser devida segundo o parceiro.

**Lição para o time:** Ao ativar o faturamento de um novo parceiro, o financeiro deve ser briefado com a data exata de início do contrato para evitar emissão retroativa de NFs.

---

### HTO-008 — Processo de integração Wake Experience com plataformas de parceiros

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Konfidency & Wake |
| **Data** | 2025-12-01 |
| **Categoria** | `how_to` |
| **Tags** | `integração`, `audience-pam`, `audience-solutions` |

Para parceiros SaaS interessados em oferecer seus produtos diretamente via painel Wake (modelo embedded):

1. A Wake está desenvolvendo o modelo de "embed de parceiros" no painel Wake — possibilidade de habilitação de oferta direta para a base via admin.
2. O primeiro parceiro testando esse modelo foi a Biso (dados/analytics).
3. Para a Konfidency, foi sinalizado que poderiam ser um dos próximos parceiros nesse modelo.
4. O processo passa por discussão com o time de produto Wake antes de iniciar o desenvolvimento.

**Documentação técnica da integração Wake Experience:** Disponível internamente. Para acesso, contatar o time técnico Wake. A integração tem dois caminhos possíveis que devem ser avaliados em call com o Tech Lead.

---

## Grupo: Wake + CRMBonus | Parcerias

**JID:** `120363356920429563@g.us`

---

### HTO-009 — Dois modelos de integração CRMBonus com Wake (V1 cupons vs. integração nativa)

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Wake + CRMBonus | Parcerias |
| **Data** | 2025-06-20 |
| **Categoria** | `how_to` |
| **Tags** | `integração`, `audience-solutions`, `audience-cs` |

A CRMBonus possui **dois modelos distintos de integração** com a Wake:

1. **V1 (cupons/resgate manual):** Modelo de resgate de cupom — o consumidor interage ativamente para aplicar o benefício. Este modelo foi depreciado/desativado para alguns clientes (ex: Blueman teve a integração V1 desativada).

2. **Integração Nativa V2 (automatizada):** Integração automática onde o cashback/bônus é aplicado sem ação explícita do consumidor. Este é o modelo preferencial e recomendado.

**Importante:** São duas experiências completamente diferentes para o lojista e para o consumidor final. Ao resolver um chamado ou integrar um novo cliente, sempre confirmar qual modelo está em uso antes de diagnosticar problemas.

**Documentação técnica:** https://crmbonus-api.readme.io/reference/fluxo-da-integra%C3%A7%C3%A3o-bi

---

### TRB-004 — Cashback aplicado junto a promoções gerando acúmulo indevido (Shoulder/Inbrands)

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Wake + CRMBonus | Parcerias |
| **Data** | 2025-10-16 |
| **Categoria** | `troubleshooting` |
| **Tags** | `integração`, `audience-cs`, `audience-solutions` |

**Sintomas:** O cashback CRMBonus estava sendo aplicado junto a outras promoções e cupons no carrinho, gerando acúmulo não previsto para lojistas como Shoulder e Inbrands.

**Causa raiz:** Por padrão, a CRMBonus calcula o bônus sobre o valor cheio do produto, sem considerar produtos com preço remarcado nem descontos de promoção/cupom aplicados.

**Resolução:** Solução técnica na CRMBonus —

Existe uma flag de configuração por loja: `permitirProdutosPromicionados: true | false`

- Quando `true`: produtos com preço remarcado (precoDe ≠ precoPor) são excluídos do cálculo do bônus.
- Para promoções/cupons: o valor da promoção é **subtraído** do valor base de cálculo do bônus. Exemplo: 2 produtos de R$10 com promoção de 10% → base de cálculo do bônus = R$18, não R$20.

Passo a passo:
1. Verificar no painel CRMBonus se a flag `permitirProdutosPromicionados` está ativa para o lojista.
2. Confirmar se a regra cobre também o cenário de desconto por promoção/cupom (subtração do valor).
3. Alinhar com a Shoulder/cliente os próximos passos de configuração e confirmar que a regra atende o caso de uso deles.

---

### FAQ-005 — Como funciona o Giftback transparente (sem digitação de telefone) no checkout Wake

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Wake + CRMBonus | Parcerias |
| **Data** | 2025-10-24 |
| **Categoria** | `faq` |
| **Tags** | `integração`, `audience-solutions` |

**Cenário:** Cliente quer implementar Giftback CRMBonus de forma transparente — sem exigir que o usuário digite o número de telefone durante a compra. Para isso, o telefone seria consultado diretamente no cadastro Wake via API.

**Perguntas e respostas:**
- **"A Wake suporta consulta de telefone do usuário autenticado via API?"** — Precisa de validação técnica com o time Wake. A CRMBonus precisa de confirmação se esse endpoint existe e está disponível.
- **"Na Wake, o usuário pode concluir compra sem estar autenticado?"** — Confirmação necessária com o time Wake (depende da configuração da loja).
- **"O desconto pode ser aplicado apenas na tela de pagamento, sem re-seleção de frete?"** — O gatilho padrão da alíquota Wake ocorre após a seleção do frete. Para lojas onde a tela de frete e pagamento são separadas, a aplicação na última tela exige análise técnica específica do fluxo da loja.

---

### HTO-010 — Integração Wake Experience com CRMBonus (agenda do vendedor / listas de audiência)

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Wake + CRMBonus | Parcerias |
| **Data** | 2025-08-18 |
| **Categoria** | `how_to` |
| **Tags** | `integração`, `audience-solutions` |

**Necessidade:** O time Wake Experience quer enviar dados de listas/audiências gerenciadas na plataforma (nomes, CPFs, e-mails ou telefones) para alimentar a Agenda do Vendedor da CRMBonus, gerando oportunidades automaticamente.

**Situação (agosto 2025):** Não existia integração pronta para esse fluxo na CRMBonus. A documentação disponível (`https://crmbonus-api.readme.io/reference/fluxo-da-integra%C3%A7%C3%A3o-bi`) era voltada a consulta/ingestão de dados para BI, não para envio de leads para a Agenda do Vendedor.

**Próximos passos recomendados:**
1. Contatar o Tech Lead da CRMBonus (Rodrigo Hidaka) para explorar possibilidades de desenvolvimento.
2. Marcar call técnica com as duas equipes para definir o formato da integração (API de envio, webhook, CSV/sheets etc.).
3. A CRMBonus sinalizou que existem "diversos caminhos" possíveis e recomenda alinhar o contexto antes de definir a arquitetura.

**Documentação de referência:** https://crmbonus-api.readme.io/reference/fluxo-da-integracao-ecomv2

---

### BPR-005 — Atualização de status do pipeline de indicações no grupo de parcerias

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Wake + CRMBonus | Parcerias |
| **Data** | 2025-06-23 |
| **Categoria** | `best_practices` |
| **Tags** | `lead-pipeline`, `audience-pam` |

**Recomendação:** Atualizar o pipeline de indicações no grupo de parcerias toda segunda-feira com os campos: data do contato inicial, nome do cliente indicado, CS responsável, data da reunião, temperatura do cliente e próxima reunião agendada. Usar o grupo WhatsApp diretamente para casos que precisam de apoio do CS.

**Rationale:** Visibilidade contínua do pipeline movimenta o grupo de parcerias e dá ao time de CS rastreabilidade direta sobre indicações em andamento, reduzindo perda de leads por falta de follow-up.

---

### TRB-005 — Ambiente de homologação Wake para desenvolvimento de integração

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Wake + CRMBonus | Parcerias |
| **Data** | 2025-08-06 |
| **Categoria** | `troubleshooting` |
| **Tags** | `homologação`, `audience-solutions` |

**Sintomas:** Durante o desenvolvimento da integração CRMBonus V1 com a Wake, o time técnico da CRMBonus precisou de acesso a um ambiente de homologação Wake e não tinha caminho de acesso direto.

**Causa raiz:** Ambiente de homologação Wake não é provisionado automaticamente a parceiros; depende de solicitação manual via gerente de parcerias.

**Resolução:** Como obter acesso —
1. O gerente de parcerias Wake solicita ao time de CS/tech o acesso ao ambiente de homologação.
2. O e-mail do desenvolvedor responsável deve ser fornecido para criação do acesso.
3. O acesso é vinculado a uma loja de teste — no caso da CRMBonus, foi direcionado para o responsável técnico (guilherme.sutto@crmbonus.com).

**Para dúvidas técnicas sobre a integração:** O time Wake disponibiliza agenda de call técnica com especialistas. Para integrações envolvendo gestão de cupons e promoções, usar o link de agenda da especialista designada:
`https://calendly.com/rafaela-teixeira/30min`

---

### FAQ-006 — Campo de cupom na segunda tela do checkout Wake: pode ser removido?

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Wake + CRMBonus | Parcerias |
| **Data** | 2026-02-06 |
| **Categoria** | `faq` |
| **Tags** | `integração`, `audience-cs`, `audience-solutions` |

**Pergunta (CRMBonus):** A Inbrands quer remover o campo de cupom na segunda tela do checkout para centralizar a aplicação de cupom apenas em uma tela — assim o Giftback CRMBonus não acumula indevidamente com cupons. O time Wake disse que o campo não pode ser removido. Por quê? A Shoulder não tem esse campo na segunda tela.

**Resposta:** O campo de cupom na segunda tela do checkout pode sim ser removido — a Shoulder é a prova disso. Não é uma limitação da plataforma, mas sim uma configuração. A Inbrands precisava ter esse ajuste feito pela agência responsável pelo checkout.

**Complicação reportada:** Após a agência da Inbrands alterar os campos, houve um conflito com o campo de "vendedor" que também estava sendo usado. O caso foi escalado para o time técnico Wake via ticket, e uma call foi agendada entre CRMBonus e o suporte Wake para resolver o conflito.

---

### BPR-006 — Uso de verbas de eventos não utilizados para outros patrocínios

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Wake + CRMBonus | Parcerias |
| **Data** | 2025-09-26 |
| **Categoria** | `best_practices` |
| **Tags** | `contrato-comercial`, `audience-pam` |

**Recomendação:** Quando um parceiro possui verba de evento Wake não utilizada, autorizar o redirecionamento desse valor para outros eventos de parceria mediante solicitação no grupo de parcerias e aprovação do gerente de parcerias via Slack com marketing Wake.

**Rationale:** A CRMBonus tinha verba do Wake on the Road Curitiba que não participou; essa verba foi proposta para cobrir o patrocínio de um evento de Farma organizado em parceria com a Wake. Redirecionar verba preserva o investimento do parceiro e mantém vivo o engajamento em eventos Wake mesmo quando o evento original não ocorre.

---

*Fim do documento — Batch 07*
