# KB Extract — Parceiros de Pagamento (PayPal, Mercado Pago, Givex, Safrapay, Vindi)
> Extracted: 2026-06-04

---

## Grupo: WAKE & PAYPAL
**JID:** `120363401981841957@g.us`

---

### FAQ-1 — Wallets disponíveis no conector PayPal

| Campo | Valor |
|---|---|
| **Grupo/Chat** | WAKE & PAYPAL |
| **JID** | `120363401981841957@g.us` |
| **Data** | 2025-08-11 |
| **Categoria** | `faq` |
| **Tags** | `paypal`, `wallet`, `apple-pay`, `google-pay`, `roadmap` |

**Pergunta:** Quais wallets já estão disponíveis no conector PayPal da Wake?

**Resposta (PayPal):** A wallet do PayPal já está integrada, assim como o checkout transparente (BCDC). Apple Pay e Google Pay ainda não estão disponíveis no BCDC — ambos estão no roadmap com expectativa de entrega até final de H2 2025. Na solução de Braintree (BT), Apple Pay, Google Pay e 3DS já estão disponíveis.

**Detalhes adicionais:**
- Apple Pay e Google Pay via BCDC: previsão de entrega 08/09/2025 para início de homologação.
- Via Braintree (BT): Apple Pay, Google Pay e 3DS já funcionavam para clientes como Shoulder (que também usa crossborder).
- A Bycoders (SI parceira) era responsável pela implementação de wallets para Wake via Braintree.

---

### FAQ-2 — Diferença entre PayPal Plus (PP+) e BCDC

| Campo | Valor |
|---|---|
| **Grupo/Chat** | WAKE & PAYPAL |
| **JID** | `120363401981841957@g.us` |
| **Data** | 2025-08-06 |
| **Categoria** | `faq` |
| **Tags** | `paypal`, `paypal-plus`, `bcdc`, `migração` |

**Pergunta:** É possível migrar clientes do PayPal Plus para o BCDC? Existe alguma issue pendente?

**Resposta (PayPal):** Sim, é possível migrar. Havia um plano criado junto com o time Wake dividindo os clientes por sprints: na primeira onda entram clientes menos expressivos para acompanhamento, e com sucesso confirmado, aciona-se os demais.

**Contexto estratégico:** Quanto mais lojistas migrarem do PayPal Plus para o BCDC, maior o incentivo do PayPal para investimentos na parceria com Wake. A migração não é automática — é feita loja a loja, com abordagem comercial para comunicar os benefícios.

**Processo:**
1. Time comercial Wake aborda os lojistas explicando benefícios da migração.
2. Informações sobre configuração foram incluídas tanto no artigo antigo quanto em novo artigo de configuração.
3. Qualquer comunicação de massa só é enviada após confirmação de sucesso da primeira bateria de lojas.

---

### TROUBLESHOOTING-1 — Alto abandono de carrinho / User Dropoff com PayPal

| Campo | Valor |
|---|---|
| **Grupo/Chat** | WAKE & PAYPAL |
| **JID** | `120363401981841957@g.us` |
| **Data** | 2025-07-07 |
| **Categoria** | `troubleshooting` |
| **Tags** | `paypal`, `abandono-de-carrinho`, `checkout`, `diagnóstico` |

**Problema:** Cliente Mimeria com 34% de User Dropoff e alto abandono de carrinho no checkout com PayPal.

**Diagnóstico:** Pedido-teste realizado pelo time PayPal pareceu normal. Recomendação: verificar se há algo em comum nos pedidos abandonados (estado, cidade, perfil do pagador) para reproduzir o cenário.

**Boas práticas para diagnóstico:**
- Não acionar o cliente antes de investigar internamente.
- Solicitar ao time PayPal que verifique pelo lado deles sem envolver o cliente.
- Realizar pedido de teste próximo às condições dos pedidos abandonados (mesmo estado, cidade, método de pagamento).

---

### TROUBLESHOOTING-2 — Problema no checkout em múltiplos merchants PayPal

| Campo | Valor |
|---|---|
| **Grupo/Chat** | WAKE & PAYPAL |
| **JID** | `120363401981841957@g.us` |
| **Data** | 2025-07-10 |
| **Categoria** | `troubleshooting` |
| **Tags** | `paypal`, `checkout`, `suporte`, `escalação` |

**Problema:** 3 merchants reportaram problema no checkout. Time técnico PayPal encaminhou e-mail com dúvidas para o time Wake (Amanda).

**Resolução:** Amanda (Wake) tratou o caso diretamente após receber o e-mail. Escalação via canal técnico por e-mail foi o caminho correto.

**Canal de escalação:** Para problemas técnicos com múltiplos merchants, o time técnico do PayPal deve enviar e-mail diretamente para o responsável técnico na Wake (no período, Amanda).

---

### HOW_TO-1 — Processo de migração PayPal Plus → BCDC

| Campo | Valor |
|---|---|
| **Grupo/Chat** | WAKE & PAYPAL |
| **JID** | `120363401981841957@g.us` |
| **Data** | 2025-08-11 |
| **Categoria** | `how_to` |
| **Tags** | `paypal`, `bcdc`, `paypal-plus`, `migração`, `sprint` |

**Fluxo de migração PayPal Plus para BCDC:**

1. **Preparação:** Time Wake cria lista de lojistas segmentada por sprints (do menos ao mais expressivo).
2. **Sprint 1:** Entrar em contato com lojistas menos expressivos para acompanhar a experiência de migração.
3. **Validação:** Confirmar que tudo está OK com a primeira bateria antes de avançar.
4. **Sprint 2+:** Acionar lojistas de maior porte e expressividade após validação positiva.
5. **Comunicação:** Time comercial Wake aborda os lojistas; PayPal não comunica diretamente.
6. **Documentação:** Artigos atualizados na central de atendimento Wake com instruções de configuração do BCDC.

**Importante:** A migração não é automática — cada lojista precisa ser acionado individualmente. O plugin do PayPal Plus deve ser ocultado para novos lojistas durante o processo.

---

### BEST_PRACTICES-1 — Estratégia de parceria e posicionamento PayPal

| Campo | Valor |
|---|---|
| **Grupo/Chat** | WAKE & PAYPAL |
| **JID** | `120363401981841957@g.us` |
| **Data** | 2025-08-05 |
| **Categoria** | `best_practices` |
| **Tags** | `paypal`, `parceria`, `bcdc`, `braintree`, `estratégia` |

**Contexto estratégico da parceria Wake × PayPal:**

- **Alavanca de investimento:** Quanto mais lojistas Wake migrarem para BCDC (em detrimento do PayPal Plus), maior o ânimo do PayPal para investir na parceria.
- **Oportunidade de expansão:** Clientes que já usam PayPal via Braintree (VTEX) podem ter a parte nacional migrada para Wake — é uma oportunidade de ampliar o escopo da conta.
- **Internacionalização:** Tema crescente entre lojistas Wake — PayPal via Braintree suporta crossborder, o que pode ser argumento de venda para lojistas com ambição global.
- **Wallets como diferencial:** Apple Pay e Google Pay no roadmap do BCDC são um argumento forte para clientes com necessidade de carteiras digitais (ex: cliente Shoulder).
- **Posicionamento no evento Wake Summit:** Parceiros de pagamento como PayPal podem participar como patrocinadores — alinhamento estratégico aumenta visibilidade mútua.

---

## Grupo: Mercado Pago + Wake
**JID:** `120363419259418672@g.us`

---

### TROUBLESHOOTING-3 — Erro "El payment method no viene en el request"

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Mercado Pago + Wake |
| **JID** | `120363419259418672@g.us` |
| **Data** | 2025-10-13 |
| **Categoria** | `troubleshooting` |
| **Tags** | `mercado-pago`, `payment_method_id`, `tokenização`, `erro-api`, `json` |

**Problema:** A plataforma Wake não estava enviando os campos `payment_method_id` e o ID de tokenização no JSON da requisição de pagamento ao Mercado Pago. Apenas o campo `installments` era enviado, causando o erro:

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

**Causa:** Bug na integração Wake × Mercado Pago: os campos obrigatórios `payment_method_id` e token não eram incluídos no payload.

**Resolução:**
1. Lojista (Potiguar) compartilhou o log de erro do Mercado Pago com a Wake.
2. Time de Parcerias Wake acionado (canal WhatsApp + Slack).
3. Canal de suporte recomendado: abrir ticket com o time técnico Wake, incluindo payload completo, IDs dos pedidos e logs.
4. Para escalar com o Mercado Pago: incluir `douglas.cruz@mercadolivre.com` e `renata.rcandrade@mercadopago.com.br` em CC no ticket.
5. Também há canal de Slack entre Wake e Mercado Pago para acompanhamento.

**Erro secundário identificado:** `"El request ya fue posteado"` — indica requisição duplicada sendo bloqueada pelo Mercado Pago.

---

### TROUBLESHOOTING-4 — Erro "Bin not found" ao testar parcelamento em 18x

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Mercado Pago + Wake |
| **JID** | `120363419259418672@g.us` |
| **Data** | 2026-02-23 |
| **Categoria** | `troubleshooting` |
| **Tags** | `mercado-pago`, `parcelamento`, `18x`, `sandbox`, `bin` |

**Problema:** Ao testar parcelamento em 18x em ambiente Sandbox, retornava erro `"Bin not found"`, mesmo com cartão que o sistema reconhecia corretamente como Mastercard.

**Causa:** O ambiente Sandbox não refletia o parcelamento em 18x, enquanto a conta do seller em produção já estava liberada para essa modalidade.

**Resolução:**
- Teste realizado em produção com conta real confirmou que o parcelamento em 18x funcionou corretamente nas duas pontas.
- Contas em Sandbox não suportam parcelamento acima de 12x por padrão — o ambiente de produção já estava configurado.
- Alternativa para testes: usar uma conta Mercado Pago TestUser (ambiente controlado do MP) para validar a jornada de ponta a ponta.

**Importante:** O parcelamento acima de 12x é liberado de forma controlada para alguns sellers. Para verificar o parcelamento disponível, utilizar o endpoint:
```
GET /v1/payment_methods/installments?bin={BIN}&public_key={CHAVE_PUBLICA}
```

---

### FAQ-3 — Parcelamento em 18x no Mercado Pago via Wake

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Mercado Pago + Wake |
| **JID** | `120363419259418672@g.us` |
| **Data** | 2025-10-29 |
| **Categoria** | `faq` |
| **Tags** | `mercado-pago`, `parcelamento`, `18x`, `configuração` |

**Pergunta:** A Wake suporta parcelamento em 18x com Mercado Pago?

**Resposta:** Tecnicamente, na plataforma Wake o lojista pode cadastrar quantas parcelas quiser no grupo de pagamento (sem limite fixo no admin). Porém, o parcelamento acima de 12x depende de liberação controlada do lado do Mercado Pago, conforme as particularidades dos emissores.

**Como configurar no admin Wake:**
1. Lojista acessa o admin e cria um **grupo de pagamento**.
2. Vincula uma **configuração de parcelamento** (com qtd de parcelas, juros, etc.) ao grupo.
3. Vincula um **conector de pagamentos** ao grupo.

Essa configuração é anterior e independente do conector utilizado. O lojista tem liberdade para cadastrar quantas parcelas desejar — a validação de disponibilidade real é feita pelo emissor/MP.

---

### HOW_TO-2 — Documentação de integração Mercado Pago com Wake

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Mercado Pago + Wake |
| **JID** | `120363419259418672@g.us` |
| **Data** | 2026-04-29 |
| **Categoria** | `how_to` |
| **Tags** | `mercado-pago`, `integração`, `documentação`, `conector-customizado` |

**Como integrar o Mercado Pago com a Wake (via conector customizado):**

- **Documentação principal:** https://wakecommerce.readme.io/docs/conector-de-pagamento-customizado
- **Documentação geral Wake:** https://wakecommerce.readme.io/

Esse é o documento enviado para provedores de pagamento que desejam construir integração com a Wake. Define o contrato de API e os fluxos esperados.

**Canal de dúvidas:** Grupo WhatsApp de parceria Wake + Mercado Pago.

---

### TROUBLESHOOTING-5 — Pedido travado em "Aguardando pagamento" (Sandbox MP)

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Mercado Pago + Wake |
| **JID** | `120363419259418672@g.us` |
| **Data** | 2026-05-21 |
| **Categoria** | `troubleshooting` |
| **Tags** | `mercado-pago`, `sandbox`, `status-pagamento`, `go-live` |

**Problema:** Cliente Venari (CNPJ: 41.907.917/0001-51) na fase final de testes para Go Live — ao realizar pedido, ele permanecia apenas com status "Aguardando pagamento", impedindo o avanço no fluxo de testes em Sandbox.

**Diagnóstico inicial:** Solicitado ao Mercado Pago o payload/request enviado à sua API para análise.

**Canal de suporte:** Para casos de integração em Sandbox, o parceiro Mercado Pago pode analisar via request compartilhado ou acompanhar via time de parcerias (canal Slack Wake × MP).

**Contatos Mercado Pago para suporte técnico em integrações:**
- douglas.cruz@mercadolivre.com
- renata.rcandrade@mercadopago.com.br

---

### BEST_PRACTICES-2 — Monitoramento de incidentes Mercado Pago

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Mercado Pago + Wake |
| **JID** | `120363419259418672@g.us` |
| **Data** | 2025-10-16 |
| **Categoria** | `best_practices` |
| **Tags** | `mercado-pago`, `incidente`, `monitoramento`, `status-page` |

**Contexto:** Identificado volume baixo de transações a partir das 11h — possível afetação geral no Mercado Pago.

**Como verificar:** O Mercado Pago (via gateway fbits) possui uma página de status para incidentes:
- **Status page Wake/fbits:** https://status.fbits.net/ (ex: incidente documentado em https://status.fbits.net/incidents/56xh7dtcy8mt)

**Boas práticas:**
- Ao identificar queda repentina no volume de transações, verificar status page antes de escalar.
- Comunicar o parceiro MP pelo canal WhatsApp de parceria para confirmação rápida.
- Documentar o incidente com horário de início e recuperação para análise posterior.

---

## Grupo: Wake + Givex
**JID:** `120363419883584104@g.us`

---

### HOW_TO-3 — Processo de onboarding de nova integração (parceiro de pagamento)

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Wake + Givex |
| **JID** | `120363419883584104@g.us` |
| **Data** | 2025-08-19 |
| **Categoria** | `how_to` |
| **Tags** | `givex`, `integração`, `onboarding`, `nda`, `parceria` |

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

### HOW_TO-4 — Fluxo de integração Givex × Wake (gift card e vale-troca)

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Wake + Givex |
| **JID** | `120363419883584104@g.us` |
| **Data** | 2026-03-18 |
| **Categoria** | `how_to` |
| **Tags** | `givex`, `gift-card`, `vale-troca`, `conta-corrente`, `produto-digital`, `integração` |

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

### FAQ-4 — Custo de integração via parceiro integrador homologado

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Wake + Givex |
| **JID** | `120363419883584104@g.us` |
| **Data** | 2025-09-11 |
| **Categoria** | `faq` |
| **Tags** | `givex`, `integrador`, `custo`, `vtrina`, `nda` |

**Pergunta:** Qual o custo estimado para integração com a Wake via SI parceira?

**Resposta:** O orçamento de integração via Vtrina (SI parceira da Wake) para o caso Givex partiu de **R$14.700**, podendo variar após análise das APIs e complexidades do projeto. Esse valor foi apresentado antes de calls técnicas — o orçamento final é ajustado após reunião de arquitetura.

**Nota:** Para a integração Givex especificamente, Wake absorveu o custo do lado da Wake para o cliente Shoulder (projeto piloto), enquanto o custo do SI (Vtrina) ficou a cargo de negociação entre Givex e Vtrina.

---

### BEST_PRACTICES-3 — Gestão do processo de parceria técnica (NDA, integradores, cronograma)

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Wake + Givex |
| **JID** | `120363419883584104@g.us` |
| **Data** | 2026-05-06 |
| **Categoria** | `best_practices` |
| **Tags** | `givex`, `parceria`, `nda`, `integrador`, `gestão-de-projeto` |

**Lições aprendidas na gestão da parceria Givex × Wake:**

- **NDA com múltiplos players:** Quando há um SI intermediador (ex: Vtrina, Wicomm), cada um deve assinar o NDA separadamente antes de ter acesso à documentação técnica. Coordenar a assinatura de todos antes da reunião técnica evita atrasos.
- **E-mail de contato correto:** Validar o e-mail certo do parceiro evita falhas de comunicação (ex: iago@givex.com.br vs iago.annes@givex.com vs iago.annes@shift4.com).
- **Ferramentas de colaboração:** Uso de Miro para desenho de fluxos — garantir que o acesso ao Miro funcione (domínios corporativos podem bloquear convites; usar e-mail pessoal como alternativa).
- **Cronograma com marcos:** Definir datas-alvo (ex: RM de Go Live em 02/06) e comunicar ao parceiro aumenta a pressão positiva no avanço.
- **Separação integrações genéricas de customizações:** Garantir que nenhum fluxo seja customizado exclusivamente para um cliente, para que a integração sirva a múltiplos lojistas.

---

### FAQ-5 — Recursos disponíveis na Wake para Gift Card e Vale-Troca

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Wake + Givex |
| **JID** | `120363419883584104@g.us` |
| **Data** | 2026-02-03 |
| **Categoria** | `faq` |
| **Tags** | `givex`, `gift-card`, `vale-troca`, `conta-corrente`, `produto-digital` |

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

### HOW_TO-5 — Processo de integração Safrapay × Wake (conector de pagamento)

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Safrapay + Wake |
| **JID** | `120363423003129579@g.us` |
| **Data** | 2026-04-20 |
| **Categoria** | `how_to` |
| **Tags** | `safrapay`, `integração`, `pix`, `cartão`, `conector`, `homologação` |

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

### TROUBLESHOOTING-6 — QR Code do Pix não carrega após checkout Safrapay

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Safrapay + Wake |
| **JID** | `120363423003129579@g.us` |
| **Data** | 2026-04-20 |
| **Categoria** | `troubleshooting` |
| **Tags** | `safrapay`, `pix`, `qrcode`, `confirmation`, `webhook` |

**Problema:** Após configurar o endpoint `/confirmation` para retornar o QR Code do Pix, o integrador não conseguia identificar chamada da Wake para esse endpoint após o checkout.

**Causa provável:** Problema no front-end da loja ou na URL configurada para o endpoint de confirmação.

**Resolução:**
1. Confirmar que a URL do endpoint de confirmação está configurada corretamente no conector.
2. Verificar se o problema está no front da loja (não renderizando a resposta) ou na configuração da URL.
3. Acionar time de Suporte Tech Wake com: nome da loja, URL do pedido de teste, ID do pedido e logs.

**Canal de suporte tech Wake:** WhatsApp de suporte (bot inicial → informar que é parceiro realizando integração → direcionamento para squad de Pagamentos).

---

### TROUBLESHOOTING-7 — Acesso ao painel admin Wake (parceiro Safrapay)

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Safrapay + Wake |
| **JID** | `120363423003129579@g.us` |
| **Data** | 2026-03-17 |
| **Categoria** | `troubleshooting` |
| **Tags** | `safrapay`, `admin`, `acesso`, `senha`, `painel` |

**Problema:** Usuário do parceiro Safrapay (Danilo) não conseguia receber o e-mail de recuperação de senha para acessar o painel admin Wake.

**Resolução:**
1. Wake gerou link direto de redefinição de senha (sem depender do e-mail): `https://idm.fbits.net/Account/AlterarSenha?Chave={CHAVE}`
2. Usuário acessa o link e redefine a senha diretamente pelo formulário.

**URL do painel admin Safrapay:** https://partnertechsafrapay.fbits.app/

**Nota:** Links de recuperação de senha têm prazo de expiração — se expirar, o time Wake gera um novo link.

---

### FAQ-6 — Versão do conector Safrapay e atualização

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Safrapay + Wake |
| **JID** | `120363423003129579@g.us` |
| **Data** | 2025-08-04 |
| **Categoria** | `faq` |
| **Tags** | `safrapay`, `conector`, `versão`, `atualização` |

**Pergunta:** A versão 01.04 do conector Safrapay é a mais atual?

**Resposta (Safrapay):** A integração existente é antiga e a versão 01.04 não é a atual. Existe uma versão mais recente — confirmação dos endpoints em uso deve ser feita na agenda técnica. O time técnico do Safrapay (Edu) é responsável pela atualização e manutenção do conector.

---

### BEST_PRACTICES-4 — Validação de cenários transacionais na integração

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Safrapay + Wake |
| **JID** | `120363423003129579@g.us` |
| **Data** | 2026-04-22 |
| **Categoria** | `best_practices` |
| **Tags** | `safrapay`, `homologação`, `cancelamento`, `status-pagamento`, `validação` |

**Checklist de cenários a validar em call de homologação ponta a ponta (Safrapay × Wake):**

- [ ] Aprovação de transações (cartão de crédito, débito, Pix)
- [ ] Cancelamento total pelo portal Wake → verificar reflexo no portal Safrapay
- [ ] Cancelamento parcial pelo portal Wake → verificar reflexo no portal Safrapay
- [ ] Cancelamento pelo portal Safrapay → verificar se status reflete corretamente na plataforma Wake
- [ ] Fluxo de Pix: geração e exibição do QR Code no front da loja
- [ ] Reflexo de frete na cotação durante checkout

**Formato recomendado:** Call conjunta com ambos os times (Wake Pagamentos + Safrapay técnico) para validação simultânea dos cenários, reduzindo ciclos de comunicação.

---

### BEST_PRACTICES-5 — Antifraude e conformidade regulatória (CNPJ alfanumérico)

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Safrapay + Wake |
| **JID** | `120363423003129579@g.us` |
| **Data** | 2026-05-25 |
| **Categoria** | `best_practices` |
| **Tags** | `safrapay`, `antifraude`, `cnpj-alfanumerico`, `instrução-normativa`, `compliance` |

**Tópicos regulatórios e de antifraude relevantes na integração Safrapay × Wake:**

1. **CNPJ Alfanumérico (IN 2.229/2024):** Instrução normativa federal que altera o formato do CNPJ para incluir letras. Wake estava avaliando a homologação deste requisito com o time de pagamentos (Natali). Status: em análise durante integração Safrapay.

2. **Coleta de dados para antifraude:** Ponto crítico durante a integração — o integrador (Danilo/Safrapay) tinha dúvidas sobre a forma correta de coletar dados para o módulo de antifraude. Recomendado alinhar diretamente com o time de pagamentos Wake (Natali).

**Contato para tópicos regulatórios/antifraude Wake:** Equipe de Pagamentos Wake (Natali).

---

## Grupo: Vindi + Wake
**JID:** `120363420795576168@g.us`

---

### HOW_TO-6 — Processo de treinamento de parceiro de pagamento para time interno Wake

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Vindi + Wake |
| **JID** | `120363420795576168@g.us` |
| **Data** | 2025-08-22 |
| **Categoria** | `how_to` |
| **Tags** | `vindi`, `treinamento`, `cs`, `comercial`, `enablement` |

**Processo de treinamento Vindi para times de CS e Vendas Wake:**

1. **Solicitação:** Time de parceiros Vindi solicita treinamento via canal WhatsApp de parceria.
2. **Aprovação antecipada:** Todo material (PPT/PDF) deve ser enviado antecipadamente para aprovação dos coordenadores de CS e comercial Wake antes da apresentação.
   - E-mails de aprovação: `natasha.silva@wake.tech` / `leticia.cena@wake.tech`
3. **Agendamento:** Coordenadores de CS alinham com analistas antes do treinamento. Evitar agendar com menos de 1 semana de antecedência para garantir quórum.
4. **Formato:** 30 minutos é suficiente para um treinamento inicial — pode ser acoplado a agendas já existentes dos times.
5. **Lançamento de campanha:** O treinamento pode ser combinado com o lançamento de uma campanha de vendas (ver modelo Vindi abaixo).

**Contatos internos Wake para treinamentos:** Gestor de CS + Coordenadores de CS + Gerente Comercial.

---

### BEST_PRACTICES-6 — Campanha de incentivo de vendas com parceiro de pagamento

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Vindi + Wake |
| **JID** | `120363420795576168@g.us` |
| **Data** | 2025-12-17 |
| **Categoria** | `best_practices` |
| **Tags** | `vindi`, `campanha`, `cs`, `indicação`, `conversão`, `parceria-comercial` |

**Resultados da Campanha "Todos pela Vindi" (jun–dez 2025):**

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

**Recomendação:** Manter campanhas de incentivo contínuas (sem hiatos longos) para preservar o ritmo de indicações. Ao encerrar uma campanha, iniciar planejamento da próxima imediatamente.

---

### BEST_PRACTICES-7 — Estratégia de co-marketing e eventos com parceiros de pagamento

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Vindi + Wake |
| **JID** | `120363420795576168@g.us` |
| **Data** | 2025-10-01 |
| **Categoria** | `best_practices` |
| **Tags** | `vindi`, `co-marketing`, `webinar`, `black-friday`, `parceria-comercial` |

**Modelos de co-marketing testados com parceiros de pagamento:**

1. **Webinar conjunto:** Vindi foi convidada para participar de webinar de Black Friday com tema "como otimizar plataforma para picos de venda" — foco em otimização de checkout, escalabilidade, e volume alto de tráfego. Parceiro pode indicar speaker (diretor, product manager ou gerente de CS).

2. **Wake Days com agências:** Eventos presenciais com agências parceiras Wake — oportunidade para o parceiro de pagamento se apresentar ao ecossistema de agências.

3. **Patrocínio de evento:** Parceiro pode patrocinar o Wake Summit (maior evento do ecossistema Wake) com cotas de Almoço VIP e Jantar Pré-Summit. Permite exposição a C-LEVELs de grandes varejistas.

**Dica de negociação:** Pacotes combinados (ex: Almoço VIP + Jantar) têm desconto — Safrapay negociou R$80k para duas cotas que custariam R$110k separadas. Useful para fechar patrocínio com menor resistência interna do parceiro.

---

### FAQ-7 — Como funciona a parceria comercial Vindi × Wake (modelo de indicação)

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Vindi + Wake |
| **JID** | `120363420795576168@g.us` |
| **Data** | 2025-08-05 |
| **Categoria** | `faq` |
| **Tags** | `vindi`, `parceria-comercial`, `indicação`, `cs`, `recorrência` |

**Pergunta:** Como funciona o modelo de parceria comercial entre Vindi e Wake?

**Resposta:** O modelo principal é baseado em **indicações de clientes**: CS Wake identificam lojistas com perfil para Vindi e indicam formalmente via campanha. Do lado Vindi, há acompanhamento comercial dos leads indicados.

**Pontos de contato recorrentes:**
- Reuniões de pulso de performance (recorrentes semanais/quinzenais).
- Campanhas de vendas com premiação para CS Wake.
- Treinamentos periódicos do time Vindi para CS e vendas Wake.
- Participação em eventos Wake (Wake Days, Wake Summit).

**Segmentos indicados pela Vindi:** Clientes Vindi que buscam plataforma de e-commerce — parceria bidirecional (Wake indica clientes para Vindi e Vindi indica clientes para Wake).

---

### FAQ-8 — Entregáveis e prazos para expositor no Wake Summit

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Vindi + Wake |
| **JID** | `120363420795576168@g.us` |
| **Data** | 2025-08-19 |
| **Categoria** | `faq` |
| **Tags** | `wake-summit`, `expositor`, `manual`, `eventos`, `parceria` |

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
