# Ambiente de Desenvolvimento Wake

**Tipo**: how_to / best_practices  
**Tags**: `desenvolvimento`, `sandbox`, `api`, `homologação`, `storefront`, `commerce-api`, `anymarket`, `headless`, `graphql`, `rate-limit`, `data-wake`, `audience-solutions`, `audience-cs`  
**Atualizado**: 2026-06

---

## Visão Geral

A Wake opera num modelo **totalmente headless**, onde o front-end é desacoplado do back-end e desenvolvido pelos parceiros/agências usando as APIs da plataforma. A documentação técnica central está em `https://wakecommerce.readme.io/`. O canal de suporte técnico para parceiros é `suporte@wake.tech`, com documentação de canais em `https://wakecommerce.readme.io/docs/canais-de-atendimento-para-parceiros`.

A Wake não é uma plataforma com ambiente de sandbox gratuito. Todo acesso para testes exige provisionamento via contrato.

---

## Ambiente de Homologação / Sandbox

### Disponibilidade

**A Wake não disponibiliza ambiente de sandbox gratuito.** Essa é uma limitação conhecida e confirmada pela própria equipe interna:

> "Como já dito, mesmo um sandbox requer um Tier Basic e partner fee de R$500. O time de vendas tem dificuldade para entender isso."  
> — WAKE | Time Tech

Se um sandbox for estritamente necessário para um parceiro tech ou integrador, é preciso provisionar um ambiente de produção vinculado ao contrato. O processo segue o mesmo fluxo de criação de nova loja.

### Como Provisionar um Ambiente de Teste

Para parceiros técnicos (integradores, fintechs, parceiros de pagamento):

1. O Partner Ops da Wake (Renato — ativo em Partner Ops, kickoffs confirmados de jul/2025 a fev/2026; coordenadora Mayara também atua no time) pode criar uma **loja de testes** vinculada à conta do parceiro. ✓ verificado
2. Informar o e-mail que será o primeiro usuário do ambiente.
3. O Partner Ops cria o usuário, mas só ativa após confirmação do parceiro ("Ok").
4. Após ativação, o e-mail de acesso é disparado automaticamente.

**Tempo médio de provisionamento**: nos casos observados (Uble, Pagaleve, Serasa Pass, Tuna, Quick Digital, SAVVI, CoreBiz, DRIVEN, Simples Inovação, Avocado Tech, Wake & Omnichat, bycoders_), o Partner Ops responde em **1–3 dias úteis** com opções de agenda de kickoff; a loja de testes em si fica disponível **após a kickoff call** (~1 semana corrida). ⚠ inferido a partir de padrão de mensagens de kickoff (Wake | Time Parcerias, ago/2025–fev/2026); não há SLA publicado.

**Quando escalar**:

- **Sem resposta do Partner Ops em 5 dias úteis** → escalar à Coordenação Parcerias (canal `WAKE | Coordenação`).
- **Loja de testes provisionada mas integração trava por >30 dias sem retorno técnico** → escalar via Partner Account Manager e abrir ticket formal em `suporte@wake.tech`.
- **Caso APP Max (referência de incidente extremo)**: o conector AppMax ficou em desenvolvimento/homologação por **mais de 1 ano** (primeira menção comercial em set/2025, homologação aprovada em 01/abr/2026, e ainda em 25/mai/2026 havia erros de checkout antigo exigindo war room — grupo Wake & Leanwork). Quando um projeto ultrapassa 6 meses sem avanço, a Leanwork escalou diretamente no grupo de projeto pedindo "war room" com a squad de pagamentos. ✓ verificado

Para parceiros de pagamento (ex: Lemoon, PayPal, Safrapay):
- É possível configurar **chaves de sandbox** separadas das chaves de produção.
- Cada ambiente (sandbox/produção) exige suas próprias chaves: `"Para cada ambiente sandbox ou produção tem que usar as chaves respectivas"` (grupo Wake <> Lemoon).
- Chaves de sandbox para o Wake Gateway devem ser solicitadas via ticket de suporte.

Para agências que precisam de acesso ao painel para demonstrações:
- Solicitar ao Partner Ops a criação de usuário com **acesso restrito** (visualização sem permissão de edição em produção).
- Informar e-mail do usuário a ser criado.
- Confirmação: o Partner Ops ativa e dispara o e-mail de acesso.

> Acesso de clientes ao painel do admin da conta deles: apenas mediante pedido do próprio cliente ao suporte Wake. A agência deve fazer a ponte para que o cliente abra o ticket.

### Limitações Conhecidas

- **Não existe sandbox gratuito**: qualquer ambiente de testes exige contrato ativo (mínimo Tier Basic + partner fee de R$ 500).
- **Chaves de sandbox do Wake Gateway**: não são geradas automaticamente; devem ser solicitadas via ticket de suporte.
- **Acesso ao painel de clientes**: a agência/integrador não recebe acesso ao admin do cliente automaticamente — o cliente precisa solicitar diretamente.
- **Integrações de homologação podem travar**: caso exemplificado com a integração APP Max, que ficou 10 meses sem ser homologada. Escalar via Partner Ops quando travado.

---

## APIs Wake

### Storefront API

A **Storefront API** é a API principal para experiências de front-end, usada no **checkout headless** e em integrações de componentes de loja.

- **Quando usar**: desenvolvimento de checkout headless, recuperação de carrinhos, busca de produtos, montagem de layouts de loja.
- **Autenticação**: via token de parceiro/loja.
- **Documentação**: `https://wakecommerce.readme.io/docs/storefront-ssr-visao-geral`
- **Integração de pagamentos via Storefront**: `https://wakecommerce.readme.io/docs/integra%C3%A7%C3%A3o-de-pagamentos-com-storefront-api`
- **Como identificar o parceiro no contexto do Scriban**: `https://wakecommerce.readme.io/docs/como-identificar-o-parceiro-no-contexto-do-scriban`

A Storefront API suporta consultas **GraphQL**. É o mecanismo padrão para componentes interativos no checkout.

**Divergência GraphQL × REST (API Pública) — caminho de resolução**: ✓ verificado

- **Sintoma**: campos presentes em uma das APIs e ausentes na outra, valores divergentes em pedidos (ex.: campos de preço no carrinho).
- **Caso de referência**: Ticket **#236794 — "Divergência entre GraphQL e RestAPI"** aberto pelo CRMBonus em ago/2025 (grupo `Wake + CRMBonus | Parcerias`). Em 07/jul/2025 o time CRMBonus foi instruído a "mandar amanhã aquela lista de todos os campos que pegam do GraphQL e não tem na API Pública" como pré-requisito para análise do time Wake.
- **Procedimento recomendado**:
  1. Abrir ticket em `suporte@wake.tech` referenciando GraphQL × REST com payload de exemplo.
  2. Anexar **lista comparativa de campos** (quais existem em GraphQL e não em REST e vice-versa) — sem isso o ticket fica parado.
  3. Em paralelo, escalar via Partner Account Manager se o caso bloquear go-live.
  4. Para campos consistentemente presentes apenas no GraphQL (ex.: kits — ver FAQ), tratar o GraphQL como fonte autoritativa para aquele domínio até resolução formal.

> **Importante**: O novo checkout (Storefront 2.0) usa a API Storefront. Parceiros que desenvolvem scripts de pagamento (ex: Pagaleve) precisam adaptar seus scripts para usar a API GraphQL do Storefront para extrair informações do pedido e abrir modais de pagamento.

### Commerce API (REST)

A **Commerce API** é a API REST pública da Wake, usada para operações de back-office: pedidos, produtos, catálogo, estoque, frete, rastreamento, clientes e marketplace.

- **Quando usar**: integrações com ERP, OMS, sistemas de frete, integrações de marketplace (via hub como Anymarket), automações de pedido.
- **Documentação central**: `https://wakecommerce.readme.io/`
- **Endpoints principais**:
  - Pedidos: `/pedidos/{pedidoId}`
  - Status de pedido: `/pedidos/{pedidoId}/status`
  - Rastreamento: `/pedidos/{pedidoId}/rastreamento`
  - Produtos: `/produtos/{identificador}`
  - Sellers (marketplace): `https://wakecommerce.readme.io/docs/consultando-todos-os-sellers-da-loja`
  - Frete customizado: `https://wakecommerce.readme.io/docs/frete-customizado-visaogeral`
  - Antifraude personalizado: `https://wakecommerce.readme.io/docs/antifraude-personalizado`
  - Kits: `https://wakecommerce.readme.io/docs/kits`
  - Chargeback: `https://wakecommerce.readme.io/docs/chargeback`
  - Checkout completo (validação): `https://wakecommerce.readme.io/docs/checkoutcomplete#valida%C3%A7%C3%A3o-checkout-online`
  - Conector de pagamento customizado: `https://wakecommerce.readme.io/docs/conector-de-pagamento-customizado`
  - Listas de compras: `https://wakecommerce.readme.io/reference/insere-dados-b%C3%A1sicos-de-uma-nova-lista-de-compras`

**Rate limit**:

- **Valor documentado**: 120 requisições por minuto (confirmado por mensagem da Uble em 24/nov/2025: "a API da Wake, que possui uma taxa limite de 120 requisições por minuto"). ✓ verificado
- **Escopo (por token / global / por endpoint)**: ⚠ **não publicado**. Em 15/set/2025 a Frete Rápido abriu a pergunta formal no grupo `Frete Rápido & Wake`: *"a API fala que é 120 por minuto, queria entender se o limite é por token ou global, se é por endpoint ou grupo de endpoints, e se temos como aumentar o limite se for o caso"* — não há mensagem subsequente no corpus com a resposta técnica definitiva. ⚠ inferido
- **Aumento de limite**: solicitado caso a caso via suporte técnico (`suporte@wake.tech`). Não há tabela pública de tiers.
- **Sintoma de estouro**: HTTP **429**. A Uble (25/nov/2025) reporta ter precisado adicionar "travas do nosso lado para não acabar retornando o http 429". A Equals atingiu o rate limit em 19/mai/2026 (grupo `Pedidos cancelados Icomm Vindi Wake`).
- **Cenários conhecidos de estouro**: integrações de frete com volume alto (40+ pedidos/min × 3+ requisições/pedido — Frete Rápido) e importações em massa (Uble: 6 meses de histórico × 2 requisições por pedido → vazão máxima ~60 pedidos/min).
- **Mitigação cliente-side**: implementar throttling/backoff; quando possível, usar endpoint de **listagem** em vez de N chamadas individuais (Uble usa `listagem de pedidos` + 2 chamadas complementares por pedido).

### Data Wake

O **Data Wake** é o produto de analytics/datalake da Wake.

- **Status (jun/2026)**: ✓ verificado — **incluído nas propostas comerciais a partir de 02/jun/2026** ("Time closing, na nossa semanal vamos falar mais sobre esse assunto, precisamos incluir desde hoje o Data Wake nas propostas comerciais" — Comercial Cross WAKE, 02/jun/2026). Em 26/mai/2026, o time discutia identificar clientes já consumindo Data Wake sem cobrança (grupo `The Rise of Skywalker`), sinalizando que o produto está em base ativa mas com modelo de cobrança em formalização.
- **Acesso**: ⚠ **modelo de acesso e pricing não publicados no corpus.** A direção é vender embedado em propostas novas; clientes existentes precisam ser revisados caso a caso. Em 28/jan/2026, time de Agência ainda perguntava internamente "Permissão pra acessar data wake, sabe quem seria?" — owner do acesso não estava claro.
- **Quando NÃO usar Data Wake**: ✓ verificado — em 22/dez/2025, a Biso descartou Data Wake após alinhamento com Engenharia: "*ele deve ser usado apenas para consultas em tempo real. Como a Biso precisa salvar os dados no servidor deles para cruzar com outras fontes, o caminho ideal é a API Pública + API Storefront*" (grupo `Wake & Biso`). **Regra prática**: para casos de extração/ETL ou cruzamento com fontes externas, preferir API Pública + Storefront; Data Wake é orientado a consulta em tempo real.
- **Diferenciação vs. Bisu**: em set/2025 havia confusão no time de CS sobre a relação entre Bisu e Data Wake após treinamento (grupo `Parcerias & Marketing Wake`). Verificar com produto antes de citar em proposta cliente. ⚠ inferido
- **Consulta de delta (não-Data-Wake)**: endpoints de catálogo retornam o delta de alterações de preço e estoque (grupo `IOTA & Wake`). Esse é um recurso da API REST, **não** do Data Wake.

### Wake Storefront MCP

A Wake disponibiliza um **MCP (Model Context Protocol)** para desenvolvimento de agentes de IA integrados à plataforma.

- **Documentação**: `https://wakecommerce.readme.io/docs/wake-storefront-mcp`
- **Wake Dev Tool Kit (Wake Plugin)**: repositório de plugin de IA disponível em `https://github.com/wake-engineering/ai-plugin`
- **Caso de uso confirmado**: agência B8One testou o MCP para construir um agente conversacional de venda integrado à loja Wake.
- Agências que utilizarem esses recursos em projetos devem sinalizar à Wake para documentação de casos de uso para marketing.

### Autenticação e Credenciais

**Commerce API (REST):**
- Autenticação via token de loja.
- Tokens são vinculados ao contrato do cliente (loja).
- Token de sandbox e token de produção são distintos — não intercambiáveis.
- Erros intermitentes de autenticação no Anymarket foram registrados; causa: formulário de vinculação não preenchido previamente (ver seção Anymarket).

**Storefront API:**
- Autenticação via token de parceiro no contexto do Scriban.
- Documentação: `https://wakecommerce.readme.io/docs/como-identificar-o-parceiro-no-contexto-do-scriban`.

**Wake Gateway (GTW):**
- Chaves separadas para sandbox e produção.
- Chaves de sandbox devem ser solicitadas via ticket de suporte.
- Ao trocar de conector de teste para conector real (ex: Vindi), um **novo ID de grupo de pagamento é gerado** — é necessário remapear no ERP do cliente.

---

## Processo de Homologação

### Passo a Passo

O processo de homologação de uma integração técnica com a Wake segue este fluxo geral:

1. **Identificação da integração**: Wake ou parceiro identifica a necessidade de integração e comunica ao time técnico responsável (geralmente via parceiro integrador como Leanwork).
2. **Agenda de scoping**: parceiro tech envia disponibilidade (3 slots); Wake/integrador valida com time técnico e confirma a call inicial.
3. **Avaliação de viabilidade**: time técnico avalia o escopo e prepara orçamento (quando aplicável).
4. **Desenvolvimento**: parceiro desenvolve a integração usando as APIs e documentação Wake.
5. **Testes em ambiente de homologação**: uso de loja de testes provisionada via Partner Ops.
6. **Homologação formal**: validação pelo time técnico Wake. Sem SLA definido publicamente.
7. **Go-live**: ativação em produção com checklist de validação.

**Para integrações de pagamento** (ex: conector customizado):
- Documentação base: `https://wakecommerce.readme.io/docs/conector-de-pagamento-customizado`
- Após go-live do checkout Storefront, ativar scripts/conectores de pagamento na etapa seguinte (deploy faseado — SF primeiro, pagamento depois).

**Para integrações com front-end (scripts/componentes)**:
- A Leanwork atua no desenvolvimento do componente de front-end quando envolve scripts conectados à API Storefront.
- Documentação de frete customizado: `https://wakecommerce.readme.io/docs/frete-customizado-visaogeral`

### Critérios de Aprovação

Não há publicação formal dos critérios de homologação. Pelos casos observados, a homologação exige:

- Testes funcionais com pedido com a feature ativa e pedido sem a feature (validação de escopo do gatilho).
- Validação de que eventos disparados (webhooks, status de pedido, automações) só afetam os pedidos corretos.
- Confirmação de que o ambiente em questão está em teste (não produção) antes de ativar automações.
- Para gateways de pagamento: confirmação de IDs de grupo de pagamento e mapeamento no ERP do cliente.

### Tempo Médio

- **Onboarding padrão de cliente**: aproximadamente 90 dias.
- **Homologação de integração de parceiro tech**: sem SLA definido. Caso reportado: integração APP Max ficou 10 meses sem homologação até escalada ao Partner Ops.
- **Provisão de painel Anymarket**: geralmente 1 dia útil após o formulário preenchido.

---

## Integrações Técnicas

### Integração com Anymarket

A integração Wake + Anymarket cobre: sincronização de catálogo de produtos, atualização de estoque, gestão de pedidos e publicação de anúncios em marketplaces.

**Provisionamento de painel:**
- Solicitar via formulário: `https://forms.office.com/r/23NsmfzG2B`
- Preencher com o máximo de detalhes no campo "Observações".
- Aguardar liberação pelo time Anymarket (geralmente 1 dia útil).
- Sem o formulário preenchido, o painel não é criado e a autenticação do token Wake falhará.

**Autenticação do token Wake no Anymarket:**
- Erro intermitente de autenticação pode ocorrer quando o formulário de solicitação de loja não foi preenchido previamente.
- Em alguns casos o erro se resolve sem ação adicional após a liberação do painel (possível fila de processamento interna).
- Se persistir após a liberação, abrir ticket no suporte Anymarket com evidências.

**Problema conhecido — Estoque reservado:**
O estoque reservado na Wake **não é consumido pelo Anymarket**. Isso significa que produtos com estoque reservado (ex: boletos pendentes que não foram baixados) podem parecer disponíveis no Anymarket. Verificar se o parceiro (ERP/OMS) está dando baixa correta no estoque reservado quando o boleto vence.

**De/Para (mapeamento de campos):**
- Existe um guia técnico de De/Para para que times de ERP configurem a integração corretamente. Solicitar ao time de CS Wake ou Anymarket.
- Configuração de atualização: verificar se está definida como `"Atualizar todos os dados dos produtos"` no módulo Wake dentro do Anymarket.
- **Imagens**: não são atualizadas automaticamente. Para atualizar, remover o produto do Anymarket e importar novamente, ou cadastrar via tela/planilha.
- **SKU Anymarket**: ao migrar seller que já possui anúncios no Magalu, preencher o campo SKU Anymarket com o SKU já existente no canal para evitar duplicação.

**Suporte Anymarket:**
- Portal de tickets: `https://suporte.anymarket.com.br`
- Chat online: disponível diretamente no painel do Anymarket (cenários urgentes).
- Escalação via grupo WhatsApp: informar número do ticket ao coordenador de suporte Anymarket.
- Status page: `https://status.anymarket.com.br/`

**Faturamento via NF-e no Anymarket:**
- Fluxo correto via e-mail: **primeiro** enviar o XML para o e-mail do Anymarket, **depois** atualizar o pedido na Wake com dados fiscais.
- Via ERP/API: verificar campo `NeedInvoicedXml` em `GET v2/orders/{id}`, depois `POST v2/orders/{id}/nfe`.
- Documentação: `https://developers.anymarket.com.br/api/v2/ky473o3qk52qb-orders-id-nfe`

### Integração com ERP (via Middleware)

A Wake não possui integração nativa com a maioria dos ERPs do mercado. A integração é feita via middleware/iPaaS certificado. Principais parceiros do ecossistema:

| Parceiro | Especialidade |
|---|---|
| Argento Digital | Hub de integrações; Linx, Senior, Omie, Abacos, WMS |
| Integra.do (Kore Hub) | TOTVS Protheus, Winthor, RM, SAP, Omie, Sankhya, Bling, Tiny |
| Iota.Hub | iPaaS; SAP, sistemas legados, B2B/B2C |
| EBI (EBInterchange) | TOTVS Protheus, Winthor, Linx Microvix, SAP |
| Intergate Consultoria | SAP Business One, TOTVS Protheus |
| Floui | SAP S/4HANA, TOTVS Protheus, Sankhya; low-code |
| Jitterbit | iPaaS global; cloud + on-premise |
| Vtrina | Hub marketplace + middleware para integrações complexas |
| Leanwork Group | Fábrica de software; conectores customizados |

Para indicar um integrador a um cliente, preencher o formulário oficial:  
`https://docs.google.com/forms/d/e/1FAIpQLSfDs4kbyGpLUZTX18Y1Y0VTH_akMpX5Y_mHbLvB9V0toO9SMw/viewform`

**Recomendação para TOTVS RM**: Intergate é a mais especializada. Integra.do lista TOTVS RM no site, mas validar disponibilidade antes de indicar.

### Webhooks

A Wake suporta webhooks para notificação de eventos.

- Usado para notificação de troca de status de pedido (ex: integração com Frete Rápido, Serasa Pass).
- No projeto Serasa Pass, o webhook de notificação de compartilhamento de dados usou o endpoint: `POST /api/serasa/webhook-user-consent` com `sessionId` no body.
- Rota de retorno após compartilhamento mobile: `/webhook-mobile/{sessionId}`.
- Para integradores de frete: receber notificação de troca de status via webhook, depois fazer GET em `/pedidos/{pedidoId}` e GET/POST em `/pedidos/{pedidoId}/rastreamento`.
- Polling como alternativa ao webhook: possível, mas webhook é o método recomendado para produção.

---

## Ferramentas de Desenvolvimento

### Painel Administrativo

O painel administrativo (admin) da Wake é o ponto central de gestão da loja para desenvolvedores e implementadores:

- Configuração de grupos de pagamento, alíquotas, conectores de gateway.
- Gestão de catálogo, produtos, variações e categorias.
- Configuração de scripts e componentes via Gestor de Scripts (posições como `Body - Última Linha`).
- Relatórios e pedidos.
- Configuração de frete customizado.

**Acesso para agências:**
- Usuário com acesso restrito (visualização): solicitar ao Partner Ops (Renato — ✓ ativo, conforme kickoffs registrados em `WAKE | Time Parcerias`, `bycoders_ & Wake`, `Tuna & Wake`, `Quick Digital & Wake Parcerias` entre jul/2025 e fev/2026).
- Acesso completo ao admin do cliente: apenas mediante solicitação do próprio cliente ao suporte Wake.

**Painel de demonstração para parceiros (Eficaz):**
A Wake mantém um ambiente específico para demonstrações de parceiros (`Acesso Loja Wake`). Usuários com acesso restrito podem demonstrar o painel sem permissão de edição em produção.

### Wake MCP e AI Plugin

- **Wake MCP (Model Context Protocol)**: permite desenvolvimento de agentes conversacionais de venda integrados à loja Wake.
  - Documentação: `https://wakecommerce.readme.io/docs/wake-storefront-mcp`
- **Wake Plugin (Wake Dev Tool Kit)**: repositório com plugin de IA, ferramentas e workflows para manipular a loja, melhores práticas de performance e segurança.
  - Repositório: `https://github.com/wake-engineering/ai-plugin`
- **LLM**: integração nativa com modelos de linguagem para automações e experiências de compra personalizadas.

### Portal de Treinamento Anymarket

Para sellers e times de CS que precisam entender a operação do Anymarket:
- Escola de Marketplace: `https://escolademarketplace.com.br`
- Anystart (portal de treinamento Anymarket): `https://escolademarketplace.com.br/anystart/`

### Postman / OpenAPI

Não há menção explícita de collection Postman ou spec OpenAPI oficial nas fontes analisadas. A documentação interativa e referência de API está centralizada em `https://wakecommerce.readme.io/` (interface Readme.io com possibilidade de testes inline).

---

## Problemas Comuns em Desenvolvimento

### Known Bugs

Problemas reportados nos grupos com status, workaround e ticket quando disponível.

#### KB-01 — Estoque reservado Wake não é consumido pelo Anymarket

- **Data identificada**: recorrente em 2025–2026 (grupo `ANYMARKET & WAKE (TÁTICO / OPERACIONAL)`).
- **Status**: ⚠ **aberto / comportamento de produto** (não classificado oficialmente como bug; tratado como gap de integração).
- **Sintoma**: produtos com boletos pendentes ou reservas não baixadas aparecem com estoque disponível no Anymarket → risco de overselling.
- **Causa raiz**: o consumo de "estoque reservado" não é refletido no contador exposto ao Anymarket.
- **Workaround**: garantir que o parceiro ERP/OMS dê baixa correta no estoque reservado quando o boleto vencer ou a reserva expirar.
- **Resolução**: abrir chamado no suporte Wake para zerar estoque reservado incorreto. Escalar ao CS da conta para priorização.

#### KB-02 — Divergência GraphQL × REST (API Pública)

- **Data identificada**: 04/ago/2025 (ticket aberto).
- **Status**: ⚠ **aberto** — sem evidência de resolução publicada no corpus.
- **Ticket**: **#236794** (grupo `Wake + CRMBonus | Parcerias`).
- **Sintoma**: campos retornados pelo GraphQL e pela API Pública divergem (presença, valor, ou nome).
- **Workaround**: tratar o GraphQL como fonte autoritativa para os domínios onde a divergência apareceu (kits, alguns campos de pedido) até a resolução do ticket. Documentar a lista comparativa de campos no próprio ticket.
- **Procedimento de escalação**: ver seção "Divergência GraphQL × REST" em **Storefront API**.

#### KB-03 — Erros de checkout antigo após homologação do conector AppMax

- **Data identificada**: 25/mai/2026 (grupo `Wake & Leanwork`).
- **Status**: ⚠ **em homologação / pós-homologação**. Conector aprovado em 01/abr/2026, mas em 25/mai/2026 surgiram erros de atualização de status e dados de cartão em lojas com **checkout antigo**.
- **Sintoma**: erros de atualização de status e cartão; logs do AppMax insuficientes para diagnóstico.
- **Workaround**: registrar ticket formal para o time Tech Wake (war room solicitado pela Leanwork). Lojas em checkout novo (Storefront 2.0) não exibiram o problema.

#### KB-04 — Autenticação Wake intermitente no Anymarket sem formulário prévio

- **Data identificada**: recorrente 2025 (grupo `ANYMARKET & WAKE (TÁTICO / OPERACIONAL)`).
- **Status**: ✅ **mitigado por processo** — não é bug de código; é dependência operacional.
- **Causa raiz**: painel Anymarket não foi criado por falta do formulário `https://forms.office.com/r/23NsmfzG2B`.
- **Resolução**: preencher o formulário, aguardar liberação (~1 dia útil); em alguns casos o erro se resolve sozinho após liberação. Se persistir, abrir ticket Anymarket.

#### KB-05 — Categoria duplicada na Wake gera vínculo errado no Anymarket

- **Data identificada**: relato recorrente 2025–2026.
- **Status**: ⚠ **aberto / corrigível manualmente**.
- **Sintoma**: produto importa categoria "errada" no Anymarket.
- **Resolução**: corrigir produto na Wake para categoria correta, excluir a categoria duplicada, aguardar sincronização.

#### KB-06 — Script de integração para de disparar após ativação de ETL

- **Data identificada**: projetos Karcher / Hero Seguros (2025).
- **Status**: ⚠ **comportamento esperado mal documentado** — ETL altera estrutura do pedido.
- **Causa raiz**: script depende de campo/flag alterado pela nova configuração ETL.
- **Workaround**: identificar trigger do script, ajustar para a nova estrutura. Testar em loja de testes antes de ativar em produção.

#### KB-07 — Evento Purchase Meta sem parâmetro `value` após mexer no checkout

- **Data identificada**: relato recorrente 2025–2026.
- **Status**: ⚠ **aberto / depende da agência** — não é bug Wake, é regressão introduzida em customizações.
- **Sintoma**: vendas registradas no Meta Ads com R$ 0,00 → ROAS destruído.
- **Workaround/validação**: após qualquer alteração no front-end do checkout, validar via Meta Events Manager e Meta Pixel Helper que `value`, `currency` e `event_id` continuam presentes.

#### KB-08 — HTML quebrado nas descrições de produto trava o feed

- **Data identificada**: relato recorrente (FeedWake).
- **Status**: ⚠ **aberto / corrigível pelo seller**.
- **Sintoma**: feed omnichannel (Anymarket, marketplaces) quebra.
- **Workaround**: auditar descrições de produto antes do go-live; remover/adequar HTML.

#### KB-09 — Equals atingindo rate limit por uso inadequado

- **Data identificada**: 19/mai/2026 (grupo `Pedidos cancelados Icomm Vindi Wake`).
- **Status**: ⚠ **aberto** (ticket Equals #442991).
- **Sintoma**: rate limit batido durante a noite.
- **Causa raiz**: padrão de chamadas mal dimensionado pelo parceiro.
- **Workaround**: implementar throttling/backoff cliente-side; revisar padrão de chamadas com o suporte Equals.

---

### Boas Práticas

Medidas preventivas que evitam classes inteiras de problemas.

#### BP-01 — Identificar plano Pagar.me (PSP × GTW) antes de prometer integração

| Plano | Descrição | Implicações |
|---|---|---|
| **PSP ("Pronto para Usar")** | Pacote fechado: adquirência + antifraude da própria Pagar.me | Antifraude só da Pagar.me; suporta split de pagamento; não é possível plugar adquirências externas |
| **GTW ("Customizado")** | Infraestrutura de gateway; o cliente escolhe adquirentes e antifraude | Pode plugar Konduto, ClearSale; adquirentes como Stone, Rede; mais flexível |

**Pergunta direta ao cliente**: `"Pronto para Usar (PSP)"` ou `"Customizado (GTW Pagar.me)"`.

Para clientes que usam Pagar.me para assinaturas, alinhar migração da base com o time técnico Wake antes de iniciar o projeto.

#### BP-02 — Validar integração de pagamento no checkout headless antes de vender

Parceiros de pagamento (Pagbank, Vindi, etc.) podem não ter integração 100% completa com o checkout headless. Caso de referência: Editora Fórum (Pagbank + checkout headless).

**Ação**: antes de fechar um deal com checkout headless + parceiro de pagamento específico, confirmar com o time técnico Wake se a integração está 100% validada.

#### BP-03 — Validar eventos Meta após qualquer mudança no checkout

Após mexer no front-end do checkout, abrir Meta Events Manager + Meta Pixel Helper e confirmar que `Purchase` dispara com `value`, `currency`, `event_id`.

#### BP-04 — Auditar HTML em descrições de produto antes do go-live

Evita quebra do feed omnichannel (ver KB-08).

#### BP-05 — Testar scripts de integração em loja de testes antes da produção

Especialmente após ativação de ETL ou alteração de configuração que mude estrutura de pedido (ver KB-06).

#### BP-06 — Implementar throttling cliente-side em qualquer integração de alto volume

Acima de ~60 pedidos/min com múltiplas chamadas por pedido, o rate limit (120 req/min) é estourado. Usar endpoints de listagem em vez de N chamadas individuais quando possível (ver seção "Rate limit").

#### BP-07 — Confirmar baixa de estoque reservado no ERP/OMS

Wake não consome estoque reservado para o Anymarket (ver KB-01). A baixa correta precisa ser feita pelo parceiro ERP/OMS.

---

## Perguntas Frequentes Técnicas

**P: A Wake tem uma coleção Postman pública?**  
R: Não foi confirmado em nenhuma fonte. A referência principal é a documentação em `https://wakecommerce.readme.io/`.

**P: O rate limit da API é por token ou global?**  
R: O valor (120 req/min) está ✓ confirmado, mas o **escopo** (por token, global, por endpoint ou grupo de endpoints) ⚠ **não foi publicamente respondido** — pergunta formal aberta pela Frete Rápido em 15/set/2025 segue sem resposta documentada no corpus. Solicitar esclarecimento ao `suporte@wake.tech` antes de dimensionar arquitetura. Ver seção "Rate limit" em Commerce API para detalhes operacionais e mitigação.

**P: É possível ter B2B e B2C no mesmo ambiente Wake?**  
R: Sim. É possível ter B2C e B2B no mesmo admin, com login com senha para clientes B2B para que a tabela de preços diferenciada não fique visível para B2C.

**P: Onde está a documentação do checkout headless?**  
R: `https://wakecommerce.readme.io/docs/storefront-ssr-visao-geral`. Há também um grupo de documentações dedicado ao Storefront no menu lateral da documentação.

**P: Como conseguir código-fonte de uma loja exemplo para o checkout headless?**  
R: Solicitar ao time técnico Wake via parceiro de parcerias. Não há repositório público de loja exemplo confirmado nas fontes.

**P: Tem integração com Loggi e Melhor Envio?**  
R: Sim, Wake tem integrações com Loggi e Melhor Envio.

**P: Os kits de produto podem ser manipulados via API?**  
R: Kits são construídos via GraphQL. Via GraphQL é possível listar, adicionar ao carrinho e remover. Não há uma chamada de API REST específica para criar/montar um kit no catálogo (grupo IOTA & Wake).

**P: Como funciona o Delta de alterações de preço e estoque via API?**  
R: Existem endpoints de catálogo que retornam o delta de produtos com alteração de preço e estoque. Detalhes devem ser validados com o time técnico Wake (grupo IOTA & Wake confirmou a disponibilidade).

**P: Como testar scripts de integração sem impactar produção?**  
R: Usar a loja de testes provisionada pelo Partner Ops. Para o checkout Storefront especificamente, é possível ativar via parâmetro de URL (`?preview-theme=...`) e cookie de parceiro (como no caso Karcher/Hero). Nunca ativar automações sem confirmar explicitamente com o cliente se a loja já está em produção.

**P: Onde está a documentação de suporte para parceiros?**  
R: `https://wakecommerce.readme.io/docs/canais-de-atendimento-para-parceiros`

**P: O painel da Wake precisa estar vinculado ao contrato para funcionar?**  
R: Sim. O painel é vinculado ao contrato do cliente. Um "sandbox" sem contrato não existe — qualquer ambiente de testes requer provisionamento via contrato.

---

## Fontes

- `/docs/kb-parceiros-tech.md` — grupos Wake & Leanwork, Eficaz <> WAKE
- `/docs/kb-parceiros-wave5a.md` — grupos SAVVI & Wake, Wake & A&EIGHT (B8One), Wake + Omni360 + Affys
- `/docs/kb-parceiros-wave6a.md` — grupo Wake <> Lemoon
- `/docs/kb-anymarket-wake.md` — grupo ANYMARKET & WAKE (TÁTICO / OPERACIONAL)
- `/docs/kb-projeto-hero-leanwork.md` — grupo Projeto Hero <> Leanwork <> Wake
- `/docs/wake-partnership-faq.md` — grupos WAKE | Time Parcerias, Comercial Cross WAKE, ANYMARKET & WAKE
- `/docs/kb-comercial-cross.md` — grupo Comercial Cross WAKE
- DB queries sobre: sandbox, homologação, API, token, credenciais, Storefront, Commerce API, Data Wake, webhook, headless, GraphQL, rate limit, Pagar.me PSP/GTW, estoque reservado, MCP, AI plugin
- Grupos adicionais consultados via DB: WAKE | Time Tech, WAKE & PAYPAL, Wake <> Lemoon, Frete Rápido & Wake, IOTA & Wake, Parcerias & Marketing Wake, Projeto Serasa Pass <> Leanwork <> Wake, Wake + CRMBonus, Pagaleve & Wake, [work] Wake & Stone, Safrapay + Wake, Eitri & Wake, NAPER.AI & Wake, Simples Inovação & Wake Parcerias

---

## Verificação

| Claim | Status | Query | Evidência |
|---|---|---|---|
| Rate limit = 120 req/min | ✓ verificado | `LIKE '%rate limit%' OR LIKE '%120 requisi%'` (8 hits) | Uble & Wake, 2025-11-24; Frete Rápido & Wake, 2025-09-15 |
| Escopo do rate limit (token/global/endpoint) não publicado | ⚠ inferido | mesma query — pergunta aberta sem resposta no corpus | Frete Rápido & Wake, 2025-09-15 |
| Estouro do rate limit gera HTTP 429 | ✓ verificado | `LIKE '%429%'` + contexto rate limit | Uble & Wake, 2025-11-26 |
| Data Wake incluído em propostas comerciais a partir de 02/jun/2026 | ✓ verificado | `LIKE '%Data Wake%' AND timestamp > '2026-01-01'` (3 hits relevantes) | Comercial Cross WAKE, 2026-06-02 |
| Modelo de cobrança/pricing do Data Wake | ❌ sem fonte | mesma query | clientes consumindo "sem cobrança" mencionados, formalização em curso |
| Data Wake é para consulta em tempo real (não ETL) | ✓ verificado | `LIKE '%Data Wake%'` | Wake & Biso, 2025-12-22 |
| Renato (Partner Ops) ativo na Wake | ✓ verificado | `LIKE '%Renato%'` (50+ hits, 30+ em contexto Partner Ops) | bycoders_ & Wake 2025-07-16; Quick Digital 2026-02-06; SAVVI 2025-11-17 |
| Coordenadora Mayara em Partner Ops | ✓ verificado | mesma query + cruzamento | FZ COMMERCE & Wake, 2025-07-23 |
| Caso APP Max: conector aprovado em 01/abr/2026 | ✓ verificado | `LIKE '%Appmax%' AND LIKE '%homolog%'` | Wake + Appmax, 2026-04-01 |
| APP Max: erros pós-homologação em checkout antigo (mai/2026) | ✓ verificado | `LIKE '%AppMax%'` (50+ hits) | Wake & Leanwork, 2026-05-25 |
| APP Max levou >1 ano da menção comercial até homologação | ✓ verificado | série temporal Wake + Appmax | set/2025 → abr/2026 (sustained thread) |
| Ticket #236794 — divergência GraphQL × REST | ✓ verificado | `LIKE '%236794%'` (1 hit) | Wake + CRMBonus, 2025-08-04 |
| Procedimento de resolução: lista comparativa de campos | ✓ verificado | `LIKE '%GraphQL%' AND grupo CRMBonus` | Wake + CRMBonus, 2025-07-07 |
| Ticket #236794 está fechado/resolvido | ❌ sem fonte | mesma query — nenhuma mensagem subsequente confirma resolução | tratar como aberto até nova evidência |
| Equals atingiu rate limit em 2026-05-19 | ✓ verificado | `LIKE '%rate limit%'` | Pedidos cancelados Icomm Vindi Wake, 2026-05-19 |
| Tempo médio de provisionamento de loja de testes (1-3 dias úteis kickoff) | ⚠ inferido | padrão observado em kickoffs do Renato | múltiplos grupos jul/2025–fev/2026 |
| Não existe sandbox gratuito | ✓ verificado | corpus + playbook §8 | WAKE \| Time Tech (citação direta) |
| Provisionamento Anymarket ~1 dia útil | ✓ verificado | citado em kb-anymarket-wake.md + grupo ANYMARKET & WAKE | (preserved from prior version) |

**Legenda**: ✓ verificado = confirmado por mensagem citada no `store/messages.db`; ⚠ inferido = padrão observado sem citação textual definitiva; ❌ sem fonte = claim não pôde ser substanciado e está marcado como tal no corpo do artigo.
