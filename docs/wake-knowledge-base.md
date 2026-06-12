# Wake Knowledge Base
> Gerado a partir de conversas reais no WhatsApp — grupos e chats internos Wake.
> Campos: `grupo/chat`, `JID`, `data`, `tipo de artigo`.

---

## Índice

1. [FAQ — Perguntas Frequentes](#faq)
2. [Troubleshooting — Problemas e Resoluções](#troubleshooting)
3. [How-To — Processos e Configurações](#how-to)
4. [Best Practices — Boas Práticas Comerciais](#best-practices)

---

## FAQ

---

### FAQ-001 — Como funciona o repasse de pagamentos no fluxo da Wake (D+1)?

| Campo | Valor |
|---|---|
| **Grupo/Chat** | WAKE \| Time Parcerias |
| **JID** | `120363321760797867@g.us` |
| **Data** | 2026-06-02 |
| **Categoria** | `faq` |
| **Tags** | `pagamento`, `gateway`, `d+1`, `repasse`, `fluxo` |

**Pergunta:** Como funciona o prazo de repasse para o lojista? O que é D+1?

**Resposta:**
D+1 é um adiantamento — significa que o repasse ocorre um dia útil após a aprovação da transação. **Na negociação padrão da Wake, o repasse é no fluxo (não antecipado).**

- Se o cliente comprou em **1x no crédito**, o lojista recebe em **30 dias**
- Se o cliente **parcelou em 12x**, o lojista recebe na mesma cadência das parcelas (30/60/90...)
- O D+1 como **antecipação** é um produto separado, disponível sob solicitação

> Ponto de atenção: a dúvida mais comum é confundir "D+1" com o prazo padrão. O padrão Wake é no fluxo, não antecipado. A antecipação tem custo adicional (3,49%/mês).

---

### FAQ-002 — O plano WakeUP tem fee adicional para loja extra?

| Campo | Valor |
|---|---|
| **Grupo/Chat** | WAKE \| Time Parcerias |
| **JID** | `120363321760797867@g.us` |
| **Data** | 2026-06-03 |
| **Categoria** | `faq` |
| **Tags** | `wakeup`, `plano`, `precificação`, `loja-adicional` |

**Pergunta:** No plano WakeUP, qual é o valor para uma loja adicional? O fee mensal é R$1.000, mas aparece como R$2.000.

**Resposta:**
Ainda em definição interna no momento desta conversa. Um FAQ específico foi criado para centralizar as dúvidas do WakeUP. O fee base do plano é R$1.000/mês. O valor de R$2.000 para loja adicional estava sendo questionado e não havia sido confirmado como definitivo. Aguardar alinhamento oficial pelo documento interno antes de passar essa informação ao cliente.

> Referência: doc interno criado pela equipe de Parcerias (Google Docs) para centralizar dúvidas do WakeUP.

---

### FAQ-003 — A Wake oferece ambiente sandbox para testar a integração com Anymarket?

| Campo | Valor |
|---|---|
| **Grupo/Chat** | ANYMARKET & WAKE (TÁTICO / OPERACIONAL) |
| **JID** | `120363418395948594@g.us` |
| **Data** | 2026-05-12 |
| **Categoria** | `faq` |
| **Tags** | `sandbox`, `ambiente-teste`, `anymarket`, `integração`, `pré-venda` |

**Pergunta:** É possível provisionar um ambiente sandbox Wake integrado com Anymarket para validar a integração antes de produção?

**Resposta:**
A Wake **não oferece ambiente de teste gratuito por padrão**. O sandbox do Anymarket existe, mas é para desenvolvimento de API com o Anymarket — não tem integração Wake disponível por padrão.

Se um sandbox for estritamente necessário:
1. Deve ser criado um painel em **ambiente produtivo** da Wake
2. Esse painel precisa ser vinculado ao contrato do cliente
3. O processo de solicitação é o mesmo de abertura de painel para novos sellers

> Ponto de atenção em pré-vendas: não prometer sandbox gratuito. Se o prospect exigir testes antes da assinatura, o custo de provisionar o painel precisa ser endereçado no negócio.

---

### FAQ-004 — Qual é o "De/Para" técnico entre os campos da Wake e do Anymarket?

| Campo | Valor |
|---|---|
| **Grupo/Chat** | ANYMARKET & WAKE (TÁTICO / OPERACIONAL) |
| **JID** | `120363418395948594@g.us` |
| **Data** | 2026-05-29 |
| **Categoria** | `faq` |
| **Tags** | `anymarket`, `integração`, `de-para`, `erp`, `campos`, `mapeamento` |

**Pergunta:** Qual campo da Wake a Anymarket consome para preencher cada campo do painel? Preciso passar o "gabarito" completo para o dev do ERP do cliente ajustar tudo de uma vez.

**Resposta:**
Existe um guia técnico de "De/Para" (field mapping) entre Wake e Anymarket. No momento desta conversa, o material estava sendo solicitado ao time Anymarket para envio formal. O time Wake abriu thread interna para estruturar essa documentação.

Para obter o mapeamento atualizado:
- Solicitar ao CS parceiro da Anymarket (`suporte.anymarket.com.br`)
- Ou acionar o time de parcerias Wake <> Anymarket via grupo WhatsApp tático

> Este documento é essencial para clientes que precisam ajustar a integração ERP → Wake → Anymarket de uma vez. Recomendado solicitar no início do onboarding, não durante.

---

### FAQ-005 — Para integrar com a Renner via Anymarket, como é o processo de habilitação?

| Campo | Valor |
|---|---|
| **Grupo/Chat** | ANYMARKET & WAKE (TÁTICO / OPERACIONAL) |
| **JID** | `120363418395948594@g.us` |
| **Data** | 2026-05-08 |
| **Categoria** | `faq` |
| **Tags** | `anymarket`, `renner`, `marketplace`, `integração`, `onboarding` |

**Pergunta:** Como um seller integra com a Renner via Anymarket?

**Resposta:**
O processo começa **na Renner**, não no Anymarket.

1. A Renner solicita três informações ao seller: **OI, TOKEN e IE**
2. Após o seller habilitar do lado da Renner, ele acessa o Anymarket para configurar o canal
3. Só então é possível seguir com a integração no painel Anymarket

> Erro comum: o seller tenta configurar no Anymarket antes de ter concluído a habilitação na Renner. Orientar o cliente a iniciar o processo pelo portal da Renner.

---

### FAQ-006 — A Wake faz B2B e B2C no mesmo ambiente?

| Campo | Valor |
|---|---|
| **Grupo/Chat** | WAKE \| Time Parcerias |
| **JID** | `120363321760797867@g.us` |
| **Data** | 2026-05 |
| **Categoria** | `faq` |
| **Tags** | `b2b`, `b2c`, `omnichannel`, `plataforma`, `pré-venda` |

**Pergunta:** É possível operar B2B e B2C no mesmo painel Wake?

**Resposta:**
Sim. A Wake suporta os dois modelos no mesmo ambiente:
- URL pública para B2C
- Área logada para B2B com tabelas de preço, produtos, banners e logotipos distintos
- Tudo gerenciado em um único admin e uma única integração ERP

> Diferencial competitivo relevante em deals com distribuidores ou indústrias que atendem tanto varejo quanto revendedores.

---

### FAQ-007 — Qual ERP integra com TOTVS Protheus + Wake?

| Campo | Valor |
|---|---|
| **Grupo/Chat** | WAKE \| Time Parcerias |
| **JID** | `120363321760797867@g.us` |
| **Data** | 2026-06-02 |
| **Categoria** | `faq` |
| **Tags** | `erp`, `totvs`, `protheus`, `integração`, `middleware` |

**Pergunta:** Qual parceiro de integração cobre TOTVS Protheus + Wake?

**Resposta:**
Vários parceiros no ecossistema Wake cobrem TOTVS Protheus:

| Parceiro | Cobertura TOTVS |
|---|---|
| Integra.do (Kore Hub) | Protheus, Winthor, RM |
| EBI (EBInterchange) | Protheus, Winthor |
| Intergate Consultoria | Protheus (API proprietária) |
| Floui | Protheus, Sankhya (low-code) |

> Atenção: a Integra.do tem Protheus listada no site, mas conforme relato de um parceiro (Savvi), o closer deles informou que a integração TOTVS RM ainda não estava pronta. Validar disponibilidade antes de indicar.

---

## Troubleshooting

---

### TRB-001 — Atualização de estoque parou após criação de segundo painel Wake

| Campo | Valor |
|---|---|
| **Grupo/Chat** | ANYMARKET & WAKE (TÁTICO / OPERACIONAL) |
| **JID** | `120363418395948594@g.us` |
| **Data** | 2026-05-27 |
| **Categoria** | `troubleshooting` |
| **Tags** | `estoque`, `anymarket`, `painel`, `integração`, `crítico` |
| **Ticket** | `#24249` |
| **Seller** | Backwash |

**Sintomas:**
- Atualizações de estoque da Wake não chegam ao Anymarket
- Problema afeta todos os canais de marketplace integrados (não apenas um)
- Iniciou após a criação de um segundo painel Wake para o mesmo seller

**Causa raiz (suspeita):**
Com a criação do segundo painel, o novo painel pode ter substituído o primeiro no mapeamento de IDs do Anymarket. O listener de atualizações de estoque estava vinculado ao painel original — transmissões do novo painel eram descartadas silenciosamente.

**Passos de resolução:**
1. Identificar qual Panel ID o Anymarket tem registrado para os produtos do seller
2. Confirmar se os produtos no Anymarket ainda apontam para VariantIDs do painel antigo
3. Acionar times Wake + Anymarket em conjunto (ticket cruzado)
4. Re-associar os produtos ao painel correto
5. Disparar sincronização manual de estoque para reconciliar saldos
6. Monitorar logs de transmissão por 24h para confirmar normalização

**Limitação conhecida:** A Wake **não permite alterar VariantIDs** de produtos existentes. Se os SKUs já estão vinculados aos VariantIDs do painel antigo, é necessária re-importação completa (deletar → reimportar), o que zera rankings e avaliações nos marketplaces.

---

### TRB-002 — Seller com pedidos não importados do Mercado Livre (erros de SKU e atributos)

| Campo | Valor |
|---|---|
| **Grupo/Chat** | ANYMARKET & WAKE (TÁTICO / OPERACIONAL) |
| **JID** | `120363418395948594@g.us` |
| **Data** | 2026-05-19 |
| **Categoria** | `troubleshooting` |
| **Tags** | `mercado-livre`, `importação`, `sku`, `atributos`, `pedidos`, `crítico` |
| **Tickets** | `#21735`, `#22160` |
| **Seller** | Adecil |

**Sintomas:**
- Pedidos do Mercado Livre não são importados para a Wake via Anymarket
- Erros de SKU inválido e conflitos de atributos nos logs de importação
- Seller sob risco de cancelar pedidos, gerando queda de nota no marketplace
- Cliente cogitando rollback para operação anterior (FOC)

**Causa raiz:**
Problemas de configuração na ponta do cliente (Anymarket) + possível inconsistência de mapeamento de atributos entre ML e o painel Anymarket.

**Passos de resolução:**
1. Abrir War Room com Wake CS + Anymarket CS + cliente
2. Analisar o painel do seller em conjunto para diagnóstico completo
3. Corrigir mapeamento de SKUs e atributos conflitantes
4. Reimportar pedidos retidos (validar se há janela de reprocessamento)
5. Comunicar cliente com SLA claro para não perder o deal

> Lição aprendida: não esperar o cliente ameaçar rollback para escalar. Quando o volume de erros de importação é contínuo, acionar War Room proativamente.

---

### TRB-003 — Atributo "modelo" configurado como variação impede listagem bulk no ML

| Campo | Valor |
|---|---|
| **Grupo/Chat** | ANYMARKET & WAKE (TÁTICO / OPERACIONAL) |
| **JID** | `120363418395948594@g.us` |
| **Data** | 2026-05-12 |
| **Categoria** | `troubleshooting` |
| **Tags** | `mercado-livre`, `atributo`, `variação`, `migração`, `anymarket` |
| **Seller** | LinhaConforto |

**Sintomas:**
- Seller não consegue criar anúncios em massa no Mercado Livre via Anymarket
- O atributo "modelo" foi configurado como variação no painel, mas o ML exige que seja atributo simples para listagem bulk
- Tentativa de corrigir (remover variação → deletar produtos → reimportar) não resolveu — erro persistiu

**Resolução aplicada:**
Criação de um **novo painel Wake** do zero, estruturado corretamente desde o início, sem herdar as inconsistências do painel anterior. Nova importação de produtos e criação de anúncios a partir do estado limpo.

**Passos:**
1. Criar novo painel Wake para o seller (processo padrão de abertura de painel)
2. Importar catálogo completo com atributos corretamente classificados (atributo vs variação)
3. Criar novos anúncios no ML via Anymarket a partir do novo painel
4. Monitorar listagem bulk no ML para confirmar ausência de erros

> Custo colateral: o seller abriu Reclame Aqui por conta dos problemas de migração. Transparência e SLA claro no início da migração reduzem esse risco.

---

### TRB-004 — Estoque reservado da Wake não é consumido pelo Anymarket

| Campo | Valor |
|---|---|
| **Grupo/Chat** | ANYMARKET & WAKE (TÁTICO / OPERACIONAL) |
| **JID** | `120363418395948594@g.us` |
| **Data** | 2026-06-02 |
| **Categoria** | `troubleshooting` |
| **Tags** | `estoque`, `reservado`, `anymarket`, `gap`, `integração` |

**Sintoma:**
O campo de "estoque reservado" da Wake não está sendo consumido pelo Anymarket. O saldo disponível exibido nos marketplaces não desconta itens já reservados no fluxo de pedidos da Wake, gerando risco de overselling.

**Status:** Gap técnico identificado e reportado — sem resolução disponível no momento. Ambos os times estavam investigando como a dinâmica de estoque reservado funciona no Anymarket.

**Impacto:** Sellers com alta rotatividade de estoque ou vendas simultâneas em múltiplos canais têm maior exposição a overselling.

**Workaround temporário:** Monitorar manualmente o saldo real vs. saldo exibido para sellers críticos enquanto aguarda resolução do gap pela integração.

---

### TRB-005 — Seller não consegue vincular anúncios existentes após migração para nova integração (Leroy Merlin)

| Campo | Valor |
|---|---|
| **Grupo/Chat** | ANYMARKET & WAKE (TÁTICO / OPERACIONAL) |
| **JID** | `120363418395948594@g.us` |
| **Data** | 2026-05-22 |
| **Categoria** | `troubleshooting` |
| **Tags** | `anymarket`, `leroy-merlin`, `migração`, `anúncios`, `vinculo` |

**Sintoma:**
Seller recebeu orientação de que para integrar com a Leroy Merlin via Anymarket seria necessário pausar a operação, excluir todos os anúncios existentes e reenviar os produtos pela nova integração. Seller ficou preocupado em perder anúncios já publicados.

**Resolução:**
Não é necessário excluir os anúncios. É possível **vincular** anúncios existentes da Leroy Merlin.

**Premissas para o vínculo funcionar:**
- O seller precisa ter o **mesmo SKU** tanto no Anymarket quanto na Leroy Merlin
- O produto e as promoções precisam estar criados previamente

> Orientação incorreta foi passada pelo canal da Leroy — validar com o time de implantação Anymarket antes de repassar instruções de migração que envolvam exclusão de anúncios.

---

## How-To

---

### HT-001 — Como configurar integração Wake com Anymarket (guia ERP → Wake → Marketplace)

| Campo | Valor |
|---|---|
| **Grupo/Chat** | ANYMARKET & WAKE (TÁTICO / OPERACIONAL) |
| **JID** | `120363418395948594@g.us` |
| **Data** | 2026-05-29 |
| **Categoria** | `how_to` |
| **Tags** | `anymarket`, `erp`, `integração`, `onboarding`, `de-para` |

**Contexto:** Para sellers com ERP customizado que precisam configurar a integração ERP → Wake → Anymarket do zero ou corrigir campos incorretos.

**Passos:**

1. **Identificar os campos obrigatórios** que a Anymarket consome da Wake (solicitar o "De/Para" técnico ao CS da Anymarket)
2. **Mapear os campos correspondentes na API Wake** para cada campo do painel Anymarket
3. **Ajustar o ERP do seller** para enviar os dados no formato correto para a Wake (categorias, atributos, variações)
4. **Validar o envio** via painel Wake antes de ativar a transmissão para Anymarket
5. **Testar ciclo completo**: produto → Wake → Anymarket → marketplace

> Dica: fazer o ajuste de todos os campos de uma vez só (solicitar o De/Para completo antes de começar). Correções incrementais multiplicam o tempo de onboarding.

---

### HT-002 — Como solicitar abertura de painel Wake para novo seller

| Campo | Valor |
|---|---|
| **Grupo/Chat** | ANYMARKET & WAKE (TÁTICO / OPERACIONAL) |
| **JID** | `120363418395948594@g.us` |
| **Data** | 2026-05-15 |
| **Categoria** | `how_to` |
| **Tags** | `painel`, `onboarding`, `novo-seller`, `provisioning` |

**Contexto:** Necessário para novos sellers, criação de painel sandbox (produtivo), migração para novo painel, ou integração com novos marketplaces como OQVestir.

**Passos:**

1. Acessar o formulário de solicitação de painel (link via time interno Wake)
2. Preencher os dados do seller (CNPJ, nome, tipo de operação)
3. Indicar na observação se é painel sandbox/teste (será vinculado ao contrato)
4. Aguardar criação e envio de credenciais de acesso
5. Para sandbox: informar no formulário que será usado para testes de integração (ex: "sandbox Wake para validação Anymarket")

---

### HT-003 — Como escalar um problema técnico entre Wake e Anymarket

| Campo | Valor |
|---|---|
| **Grupo/Chat** | ANYMARKET & WAKE (TÁTICO / OPERACIONAL) |
| **JID** | `120363418395948594@g.us` |
| **Data** | 2026-05 a 2026-06 |
| **Categoria** | `how_to` |
| **Tags** | `suporte`, `escalonamento`, `ticket`, `cross`, `anymarket` |

**Fluxo padrão de escalonamento Wake ↔ Anymarket:**

1. **Abrir ticket** no portal Anymarket: `suporte.anymarket.com.br`
2. **Compartilhar o número do ticket** no grupo WhatsApp tático ANYMARKET & WAKE
3. Marcar o CS responsável da contraparte no grupo (ex: `@CS_Anymarket` ou `@CS_Wake`)
4. Para casos críticos (operação parada, risco de cancelamento de pedidos): solicitar explicitamente **prioridade** e justificar impacto no negócio do seller
5. Para diagnóstico conjunto: solicitar **War Room** com ambos os times + cliente
6. Registros de acompanhamento ficam no ticket; alinhamentos rápidos no WhatsApp

> Tópicos de resolução estruturada devem ser movidos para **Slack** — conversas longas no WhatsApp se perdem. WhatsApp é para visibilidade rápida e acionamento; Slack para resolução documentada.

---

## Best Practices

---

### BP-001 — Prateleira Infinita (Connect): não assumir compatibilidade de maquininha Pagar.me

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Comercial Cross WAKE |
| **JID** | `120363043949073807@g.us` |
| **Data** | 2026-06-03 |
| **Categoria** | `best_practices` |
| **Tags** | `prateleira-infinita`, `pagar-me`, `pdv`, `maquininha`, `comercial` |

Ao prospectar um cliente com interesse em Prateleira Infinita (Connect), **não assuma que as maquininhas que o cliente já possui são compatíveis**.

Nem todos os terminais POS da Pagar.me funcionam com essa solução.

**Processo correto:**
1. Identificar o interesse do cliente em Prateleira Infinita
2. Informar que é necessária uma negociação comercial **específica com a Pagar.me**
3. No processo da Pagar.me: credenciamento de conta dedicada + envio de maquininhas compatíveis
4. Não avançar no projeto sem confirmar esse alinhamento

> Assumir compatibilidade antes da negociação Pagar.me gera expectativas incorretas, atraso em projetos e desgaste na experiência do cliente.

---

### BP-002 — Incluir Data Wake em todas as propostas comerciais

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Comercial Cross WAKE |
| **JID** | `120363043949073807@g.us` |
| **Data** | 2026-06-02 |
| **Categoria** | `best_practices` |
| **Tags** | `data-wake`, `proposta`, `comercial`, `produto` |

A partir de junho/2026, o **Data Wake deve ser incluído em todas as propostas comerciais** desde o primeiro contato.

> Instrução direta da liderança comercial. Dúvidas: acionar o responsável pelo produto no time de closing.

---

### BP-003 — Como precificar migração de checkout para parceiros

| Campo | Valor |
|---|---|
| **Grupo/Chat** | WAKE \| Time Parcerias |
| **JID** | `120363321760797867@g.us` |
| **Data** | 2026-06-02 |
| **Categoria** | `best_practices` |
| **Tags** | `checkout`, `migração`, `precificação`, `agência`, `parceiro` |

Ao responder perguntas de parceiros sobre custo de migração para o novo checkout Wake, usar as seguintes referências:

| Escopo | Faixa de preço |
|---|---|
| Checkout apenas | R$ 5K – R$ 15K |
| Redesign completo da loja | R$ 15K – R$ 40K |

A demanda de migração de checkout tipicamente **cai para as agências parceiras**, não para o time Wake direto.

> Atenção ao passar essa informação para o parceiro: com ~200 clientes aguardando migração para o novo checkout, o volume é alto. Avaliar como comunicar sem desanimar a agência parceira com o tamanho da fila.

---

### BP-004 — Parceiros Elite têm prioridade em eventos e painéis Wake

| Campo | Valor |
|---|---|
| **Grupo/Chat** | WAKE \| Time Parcerias |
| **JID** | `120363321760797867@g.us` |
| **Data** | 2026-06-02 |
| **Categoria** | `best_practices` |
| **Tags** | `elite`, `parceiro`, `evento`, `ecbr`, `fórum` |

Em painéis de eventos como ECBR e Wake Evolution, a prioridade de indicação de parceiros segue a ordem de tier:

1. **Elite** — prioridade máxima
2. Demais tiers em seguida, conforme sinergia com o tema do painel

Quando não for possível encaixar um Elite em determinado painel, escalar para a liderança de parcerias antes de indicar um parceiro de outro tier.

> Regra válida para seleção de parceiros em painéis, mentorias e qualquer visibilidade pública em nome da Wake.

---

### BP-005 — Corrida de vendas para parceiros: modelo unificado em construção

| Campo | Valor |
|---|---|
| **Grupo/Chat** | WAKE \| Time Parcerias |
| **JID** | `120363321760797867@g.us` |
| **Data** | 2026-06-02 |
| **Categoria** | `best_practices` |
| **Tags** | `corrida-vendas`, `incentivo`, `parceiro`, `crmbonus`, `modelo` |

Wake está unificando o modelo de campanhas de vendas ("corridas de vendas") para todos os parceiros. Anteriormente, a CRMBonus tinha estrutura de premiação diferente dos demais. O objetivo é um modelo único e padronizado.

Ao comunicar incentivos a parceiros, aguardar a publicação do novo modelo unificado antes de fazer promessas específicas de premiação — o modelo anterior pode não valer para todos.

---

*Última atualização: 2026-06-04 | Fonte: WhatsApp (grupos internos Wake e ANYMARKET & WAKE)*
