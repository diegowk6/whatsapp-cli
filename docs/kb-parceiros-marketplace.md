# KB Extract — Parceiros Marketplace (Blueman, Anjo Rosa, Duo)
> Extracted: 2026-06-04

---

## Grupo: Wake + Marketplace Blueman

### TROUBLESHOOTING-1 — Pedido não integrado por falta de nome de exibição na forma de pagamento

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Wake + Marketplace Blueman |
| **JID** | `120363402217486276@g.us` |
| **Data** | 2025-06-20 |
| **Categoria** | `troubleshooting` |
| **Tags** | `wake`, `anymarket`, `integração`, `forma-de-pagamento`, `pedido` |

**Problema:** Pedido criado no Mercado Livre não desceu para o sistema de faturamento (síntese/ERP).

**Causa raiz:** O sistema buscou os detalhes da forma de pagamento via API `/formasPagamento` da Wake para identificar o tipo (Pix, boleto, cartão). O campo "nome de exibição" da `formapagamentoid:200` estava em branco na Wake.

**Mensagem de erro:** `não será possivel importar o pedido: [ID]. o tipo pgto não foi informando, provavelmente o nome exibição da formapagamentoid:200 na wake está em branco`

**Resolução:** Preencher o campo "nome de exibição" da forma de pagamento correspondente no painel da Wake.

---

### TROUBLESHOOTING-2 — Bug no agrupamento de variantes bloqueando catálogo no Mercado Livre

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Wake + Marketplace Blueman |
| **JID** | `120363402217486276@g.us` |
| **Data** | 2025-06-20 |
| **Categoria** | `troubleshooting` |
| **Tags** | `anymarket`, `mercado-livre`, `variantes`, `catalogo`, `bug` |

**Problema:** Impossibilidade de subir novos anúncios no Mercado Livre via Anymarket por conta de bug no agrupamento de variantes.

**Contexto:** O bug impedia a integração do catálogo Wake na Anymarket, bloqueando a publicação de anúncios no Mercado Livre para a Blueman.

**Status/Resolução:** Bug foi para fase de homologação na Anymarket com previsão de correção no início da semana seguinte (segunda-feira). Após a correção confirmada, o catálogo Wake seria integrado à Anymarket e a Blueman poderia subir novos anúncios.

---

### TROUBLESHOOTING-3 — Pedidos de marketplace não aparecem no financeiro após mudança de JSON

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Wake + Marketplace Blueman |
| **JID** | `120363402217486276@g.us` |
| **Data** | 2025-10-29 |
| **Categoria** | `troubleshooting` |
| **Tags** | `anymarket`, `financeiro`, `json`, `de-x-para`, `integração` |

**Problema:** Pedidos de marketplaces (Riachuelo, TikTok Shop, Mercado Livre) pagos e integrados não apareciam no financeiro/ERP.

**Causa raiz:** A Anymarket mudou o padrão do JSON enviado, quebrou o DE x PARA configurado no ERP para a perna financeira. O sistema não gerava parcelas nem o movimento financeiro dos pedidos.

**Resolução:** O desenvolvedor do ERP revisou e remapeou o DE x PARA para o novo padrão de JSON. Cada pedido existente antes da melhoria precisou ser tratado manualmente. Para novos pedidos, a parametrização ficou automática.

**Aprendizado:** Mudanças de formato no JSON da Anymarket podem quebrar silenciosamente integrações financeiras. É necessário monitoramento de pedidos no financeiro após atualizações da plataforma de integração.

---

### TROUBLESHOOTING-4 — Pedido Mercado Livre sem CPF bloqueia emissão de NF-e

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Wake + Marketplace Blueman |
| **JID** | `120363402217486276@g.us` |
| **Data** | 2025-11-12 |
| **Categoria** | `troubleshooting` |
| **Tags** | `mercado-livre`, `cpf`, `nfe`, `anymarket`, `faturamento` |

**Problema:** Pedido do Mercado Livre chegou sem CPF do comprador. Como CPF é obrigatório para emissão de NF-e no Brasil, o pedido não pôde ser faturado automaticamente. Impactou o prazo de envio e a reputação da loja no ML.

**Causa raiz:** O Mercado Livre permite que compradores concluam a compra sem CPF. O CPF pode ser preenchido depois, mas a Anymarket não atualiza automaticamente na Wake um pedido já exportado.

**Resolução (caso a caso):** Identificar o cliente manualmente, inserir o CPF e faturar. Não há bloqueio automático disponível na Anymarket para pedidos sem CPF, pois cancelamentos automáticos prejudicam a reputação da loja.

**Workaround:** Monitorar pedidos que chegam sem CPF e tratar manualmente. Qualquer cancelamento deve ser ponderado contra o impacto na reputação.

---

### TROUBLESHOOTING-5 — Status "Sincronizando venda com marketplace" travado — Riachuelo

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Wake + Marketplace Blueman |
| **JID** | `120363402217486276@g.us` |
| **Data** | 2026-01-26 |
| **Categoria** | `troubleshooting` |
| **Tags** | `riachuelo`, `anymarket`, `status`, `entregue`, `chamado` |

**Problema:** Dois pedidos ficaram presos com status "Sincronizando venda com marketplace" desde o dia 14/01. Isso impedia a atualização para "entregue" na Riachuelo, o que bloqueava o recebimento do valor da venda.

**Resolução:** Abertura de chamado no suporte da Anymarket + chamado paralelo na Riachuelo. O time de suporte da Anymarket executou ação de voltar e avançar o status do pedido manualmente. Com isso o status "entregue" refletiu na Riachuelo e o recebimento foi liberado.

**Observação:** A plataforma da Riachuelo foi descontinuada em 30/01/2026, tornando esse marketplace inativo para a Blueman a partir dessa data.

---

### TROUBLESHOOTING-6 — Integração com Riachuelo exige token gerado via Seller Center

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Wake + Marketplace Blueman |
| **JID** | `120363402217486276@g.us` |
| **Data** | 2025-08-07 |
| **Categoria** | `troubleshooting` |
| **Tags** | `riachuelo`, `anymarket`, `token`, `seller-center`, `integração` |

**Problema:** Integração com a Riachuelo não estava puxando informações, mesmo com acesso ao portal disponível.

**Causa raiz:** O token necessário para a integração precisa ser solicitado diretamente pelo Seller Center da Riachuelo. As informações do portal padrão não são suficientes.

**Resolução:** Abrir solicitação no Seller Center da Riachuelo para geração do token de integração. Após obter o token, configurar na Anymarket.

---

### TROUBLESHOOTING-7 — SKUs com estrutura incompatível entre Wake e anúncios ML via Anymarket

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Wake + Marketplace Blueman |
| **JID** | `120363402217486276@g.us` |
| **Data** | 2025-10-20 |
| **Categoria** | `troubleshooting` |
| **Tags** | `mercado-livre`, `anymarket`, `sku`, `variantes`, `catalogo` |

**Problema:** Produtos cadastrados na Wake agrupados apenas por tamanho (sem agrupamento por cor) geraram incompatibilidade com anúncios no Mercado Livre que precisavam vincular múltiplos SKUs (por cor + tamanho) em um único anúncio.

**Causa raiz:** A estrutura do anúncio no ML era diferente do cadastro na Wake: o anúncio esperava agrupamento por cor e tamanho, mas o cadastro Wake tinha apenas tamanho único. A Anymarket não consegue vincular vários produtos em um único anúncio nessa situação.

**Resolução:** Pausar os SKUs incompatíveis no ML. Criar uma planilha de mapeamento com campo "SKU WAKE" para validação cruzada entre Anymarket e Wake antes de reativar anúncios.

---

### HOW_TO-1 — Como parametrizar formas de pagamento por marketplace no ERP

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Wake + Marketplace Blueman |
| **JID** | `120363402217486276@g.us` |
| **Data** | 2025-10-31 |
| **Categoria** | `how_to` |
| **Tags** | `erp`, `forma-de-pagamento`, `administradoras`, `marketplace`, `parametrização` |

Para suportar múltiplos marketplaces com controle financeiro separado no ERP (LINX/Protheus), o método recomendado é:

1. **Criar uma "Administradora" no ERP para cada marketplace.** Isso permite que o sistema gere parcelas corretamente e diferencie o fluxo financeiro por origem.
2. **Mapeamento por sigla de marketplace:**
   - `BB` → Babado Top
   - `IF` → iFood
   - `ML` → Mercado Livre
   - `RC` → Riachuelo
   - `TI` → TikTok Shop
   - `OP` → Off Premium
3. **Usar o campo `marketPlacePedidoId`** presente no JSON da Anymarket como o número do título financeiro. Esse campo está disponível em múltiplos lugares no JSON do pedido:
   ```json
   "marketPlacePedidoId": "1361270021"
   "metadados": [{ "chave": "Cód pedido mktp", "valor": "1361270021" }]
   ```
4. Pedidos que chegam com tipo de pagamento "J – Duplicatas" de marketplace devem ser reclassificados para o tipo correto via DE x PARA, pois o ERP não consegue gerar parcelas a partir de duplicatas de marketplace sem essa configuração.

---

### HOW_TO-2 — Fluxo de treinamento Anymarket para operadores de marketplace

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Wake + Marketplace Blueman |
| **JID** | `120363402217486276@g.us` |
| **Data** | 2025-08-06 |
| **Categoria** | `how_to` |
| **Tags** | `anymarket`, `treinamento`, `marketplace`, `operação` |

Após conclusão dos ajustes de integração Wake-Anymarket-Marketplace, é necessário treinar o time operacional da loja:

1. **Identificar participantes:** Definir quem será treinado (ex: responsável pelo marketplace + equipe de atendimento/SAC).
2. **Agendar sessão com a Anymarket:** O time da Anymarket conduz o treinamento sobre como operar a plataforma — incluindo gestão de anúncios, monitoramento de pedidos e resolução de erros comuns.
3. **Duração sugerida:** 1h30 a 2h por sessão; reservar sessão adicional caso seja necessário aprofundar temas.
4. **Tópicos prioritários:** gestão de catálogo, vínculo de SKUs, fluxo de pedidos, emissão de etiquetas, monitoramento de status de entrega.

---

### BEST_PRACTICES-1 — Controle de status de pedidos de marketplace — lições aprendidas (Blueman)

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Wake + Marketplace Blueman |
| **JID** | `120363402217486276@g.us` |
| **Data** | 2025-11-13 |
| **Categoria** | `best_practices` |
| **Tags** | `marketplace`, `operação`, `monitoramento`, `black-friday`, `anymarket` |

Principais aprendizados operacionais coletados ao longo da implementação da Blueman nos marketplaces:

- **Monitoramento contínuo é obrigatório:** Pedidos podem ficar presos em status intermediários (ex: "Sincronizando", "Aguardando faturamento") sem alerta automático. É necessário processo manual de verificação regular, especialmente antes de datas como Black Friday.
- **Não operar em produção sem integração 100% testada:** A Blueman operou quase um ano com integrações parciais (apenas Riachuelo funcionando corretamente), o que gerou retrabalho e risco operacional.
- **Etiqueta de envio ML deve ser gerada automaticamente:** A emissão manual de etiqueta do Mercado Livre gera risco de envio fora do prazo e impacto na reputação da loja. A API de emissão de etiquetas da Anymarket (`/api/v2/emissao-de-etiquetas`) deve ser integrada com o OMS para emissão automática após faturamento.
- **Manter múltiplos CDs habilitados:** A Riachuelo e o Mercado Livre inicialmente estavam configurados apenas com o CD principal. Os demais estoques precisaram ser habilitados separadamente na Anymarket.
- **Informar XML da NF por e-mail correto:** O e-mail de importação de XML da Anymarket (`importacaoxml.anymarket@db1.com.br`) deve estar configurado corretamente no sistema de emissão de NF-e. Verificar periodicamente a caixa de entrada e o espaço disponível.

---

### FAQ-1 — A Anymarket suporta emissão automática de etiquetas para Mercado Livre?

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Wake + Marketplace Blueman |
| **JID** | `120363402217486276@g.us` |
| **Data** | 2026-01-28 |
| **Categoria** | `faq` |
| **Tags** | `anymarket`, `mercado-livre`, `etiqueta`, `oms`, `integração` |

**Pergunta:** É possível que as etiquetas de envio do Mercado Livre sejam emitidas automaticamente ao faturar o pedido?

**Resposta:** A Anymarket possui uma API para emissão de etiquetas (`https://developers.anymarket.com.br/api/v2/b8c7fdc94a369-emissao-de-etiquetas`) que pode ser consumida pelo OMS ou ERP. Porém, essa integração precisa ser desenvolvida pelo time do OMS — a Anymarket fornece o endpoint, mas não faz a integração automaticamente com todos os sistemas de gestão. Verifique se o OMS utilizado já possui esse conector disponível antes de desenvolver uma solução customizada.

---

### FAQ-2 — Como funciona a cobrança de TR para lojas sem finalização de compra no site?

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Wake + Marketplace Blueman |
| **JID** | `120363402217486276@g.us` |
| **Data** | 2025-08-21 |
| **Categoria** | `faq` |
| **Tags** | `wake`, `tr`, `cobrança`, `checkout`, `venda-consultiva` |

**Pergunta:** Um cliente quer usar a Wake como vitrine com checkout, mas sem finalizar a compra no site — o pedido vai para um time de vendas consultivo. Como funciona a cobrança de TR nesse cenário?

**Resposta (Wake):** A cobrança de TR (Transaction Rate) é feita por pedido aprovado. Se a compra não finaliza no site, não há cobrança de TR padrão. No entanto, a Wake precisará de uma negociação diferente de cobrança — por exemplo, cobrança por ordem de pedido gerada ao invés de pedido aprovado. Esse cenário já foi implementado com outros clientes, mas exige um acordo comercial específico com o time da Wake.

---

## Grupo: Comercial Anjo Rosa

### FAQ-3 — Qual o fluxo de contratação de publicidade com exclusividade para influenciadores?

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Comercial Anjo Rosa |
| **JID** | `120363402691904210@g.us` |
| **Data** | 2026-03-11 |
| **Categoria** | `faq` |
| **Tags** | `influencer`, `exclusividade`, `contrato`, `precificação`, `publicidade` |

**Pergunta:** Como precificar e estruturar uma campanha com cláusula de exclusividade?

**Resposta:** Quando uma marca solicita exclusividade (ex.: sem menção a concorrentes durante o período de veiculação), aplica-se um acréscimo sobre o valor tabelado padrão. A prática de mercado é de **pelo menos 30% de acréscimo** sobre o valor da tabela base para cobrir o custo de oportunidade de recusar outras marcas do segmento.

**Exemplo prático:** Tabela base: R$ 9.200,00 → Com exclusividade de 3 meses: R$ 11.960,00 (mín.)

**Atenção:** Se a campanha ocorrer em meses de alta demanda (ex.: Setembro/Outubro), considerar acréscimo adicional de 20% por sazonalidade.

---

### FAQ-4 — Qual é o escopo padrão de uma campanha com presença em evento (meet point)?

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Comercial Anjo Rosa |
| **JID** | `120363402691904210@g.us` |
| **Data** | 2026-03-27 |
| **Categoria** | `faq` |
| **Tags** | `influencer`, `evento`, `meet-point`, `escopo`, `entregas` |

**Escopo padrão de campanha com presença em evento (formato "Meet Point" — ex.: Estadão/A.C. Camargo):**

- **Pré-evento:** 1 story de até 60" com inserção de link CTA, marcações e hashtag.
- **Durante o evento:** Cobertura com mínimo de 5 telas de stories não sequenciais, com link CTA, marcações e hashtag.
- **Pós-evento:** 1 Reels de até 60" captado no evento, postado no dia seguinte com os melhores momentos. Possibilidade de collab ou repost pelo contratante.
- **Direitos incluídos:** Replicação em todas as redes do grupo contratante; impulsionamento com mídia paga; uso de imagem em ambiente digital e impresso; assessoria de imprensa.
- **Produção:** Pelo influenciador (sem equipe de make/hair/stylist no local).
- **Logística:** Incluída no cachê (sem verba separada para deslocamento).

**Regras operacionais comuns:**
- Chegar com 30 min de antecedência.
- Gravar na vertical, com atenção à qualidade de iluminação e áudio.
- NÃO mencionar marcas concorrentes ou indicações de tratamento (quando aplicável).
- NF só pode ser emitida após liberação da marca contratante.

---

### FAQ-5 — Como funciona o programa de afiliadas (ex.: Desinchá)?

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Comercial Anjo Rosa |
| **JID** | `120363402691904210@g.us` |
| **Data** | 2026-04-13 |
| **Categoria** | `faq` |
| **Tags** | `afiliada`, `comissão`, `stories`, `cupom`, `link` |

**Modelo padrão de programa de afiliadas de marca (health/wellness):**

- **Envio de produtos:** Mensal (para uso e review).
- **Comissão:** 10% sobre vendas geradas via cupom e/ou link rastreável.
- **Cadência de conteúdo:** ~4 stories por mês (média de 1x por semana).
- **Onboarding:** Preenchimento de formulário de cadastro → inclusão em grupo de WhatsApp da marca → recebimento de briefings, direcionamentos e novidades.

**Fluxo de cadastro:** Formulário → sistema interno da marca → grupo WhatsApp → início das ativações.

---

### HOW_TO-3 — Como emitir Carta de Correção Eletrônica (CC-e) para ajuste em NF-e já autorizada

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Comercial Anjo Rosa |
| **JID** | `120363402691904210@g.us` |
| **Data** | 2026-04-29 |
| **Categoria** | `how_to` |
| **Tags** | `nfe`, `carta-correcao`, `fiscal`, `sefaz`, `faturamento` |

**Contexto:** O financeiro do contratante (agência de mídia/veículo) solicitou ajuste na descrição da NF após emissão. A NF estava correta conforme briefing original, mas houve troca de responsável no financeiro do cliente e foi solicitada uma nova descrição.

**O que é a CC-e:** A Carta de Correção Eletrônica (CC-e) é um documento fiscal digital usado para corrigir erros em uma NF-e já autorizada pela SEFAZ, sem cancelá-la.

**Quando usar:** Quando há erro em campos de texto/descrição da NF que não alteram valores, impostos ou dados principais. Não serve para corrigir valores, CNPJ do destinatário ou dados de produto.

**Atenção sobre Simples Nacional:** Empresas do Simples Nacional não aceitam NFs com retenção de impostos sob responsabilidade do tomador. Se o influenciador/prestador for optante do Simples, emitir a NF sem retenção de ISS/PIS/COFINS/CSLL pelo tomador.

---

### BEST_PRACTICES-2 — Estrutura de proposta comercial para campanhas de influência — formato mensal vs. pacote fechado

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Comercial Anjo Rosa |
| **JID** | `120363402691904210@g.us` |
| **Data** | 2025-09-06 |
| **Categoria** | `best_practices` |
| **Tags** | `proposta-comercial`, `influencer`, `negociação`, `mensal`, `pacote` |

**Lição aprendida:** Ao apresentar proposta de campanha de influência, estruturar como **preço mensal condicionado à contratação do período total**, não apenas mês a mês. Isso:

- Garante o budget completo da campanha para a marca.
- Deixa espaço para negociação de pagamento sem abrir mão do escopo total.
- Facilita renovação: ao final do período, apresentar nova tabela para o ciclo seguinte.

**Exemplo:** "O valor mensal de R$ X está condicionado ao fechamento do período completo de [N] meses. Para o próximo ciclo, apresentaremos nova proposta de acordo com o desempenho."

---

### BEST_PRACTICES-3 — Gestão de relacionamento com agências de mídia (intermediárias)

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Comercial Anjo Rosa |
| **JID** | `120363402691904210@g.us` |
| **Data** | 2026-04-24 |
| **Categoria** | `best_practices` |
| **Tags** | `agencia-de-midia`, `intermediação`, `briefing`, `nf`, `prazo-de-pagamento` |

**Boas práticas identificadas ao trabalhar com agências intermediárias (ex.: Pulso Digital, ID - Impulso Digital para o Estadão):**

1. **Prazo de pagamento:** Agências de mídia trabalham com prazos longos (ex.: 100 dias após recebimento da NF). Considerar no fluxo de caixa antes de fechar.
2. **NF só após liberação:** A NF deve ser emitida apenas após liberação expressa da marca contratante — nunca antes. Isso é cláusula padrão dos contratos.
3. **Confidencialidade do briefing:** As agências costumam solicitar que os valores e termos do briefing não sejam divulgados. Tratar como confidencial.
4. **Troca de contato:** É comum haver rotatividade de pessoas no cliente/agência. Sempre formalizar as condições por e-mail e confirmar o ponto de contato atual antes de emitir NF.
5. **Registro no CRM:** Registrar todos os projetos em CRM (ex.: Moskit) assim que fechados, com nome do cliente, valor e status da NF.

---

### BEST_PRACTICES-4 — Prospecção em eventos de saúde e inovação para influenciadores B2B

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Comercial Anjo Rosa |
| **JID** | `120363402691904210@g.us` |
| **Data** | 2026-05-05 |
| **Categoria** | `best_practices` |
| **Tags** | `prospecção`, `eventos`, `saude`, `networking`, `b2b` |

Participar de eventos de saúde e inovação é uma forma eficiente de prospectar marcas para campanhas de influência na área de saúde/bem-estar:

- **São Paulo Innovation Week:** Evento multissetorial de inovação em SP — útil para conectar com empresas de tecnologia, startups e marcas que investem em marketing de conteúdo.
- **Eventos hospitalares e oncológicos:** Eventos como Meet Point (A.C. Camargo/Estadão) são oportunidades de aparecer como referência no nicho de saúde feminina, oncologia e bem-estar.
- **Estratégia:** Participar como influenciador (cobertura de conteúdo) e simultaneamente prospectar novas marcas presentes no evento para campanhas futuras.

---

### TROUBLESHOOTING-8 — NF com retenção de impostos recusada por empresa do Simples Nacional

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Comercial Anjo Rosa |
| **JID** | `120363402691904210@g.us` |
| **Data** | 2026-05-04 |
| **Categoria** | `troubleshooting` |
| **Tags** | `nfe`, `simples-nacional`, `retencao-impostos`, `faturamento`, `fiscal` |

**Problema:** NF emitida com retenção de impostos (ISS/PIS/COFINS/CSLL) foi recusada pelo financeiro do cliente.

**Causa raiz:** A empresa contratante era optante do Simples Nacional, que não pode aceitar NFs com retenção de impostos sob responsabilidade do tomador.

**Resolução:** Reemitir a NF (ou emitir CC-e se aplicável) sem retenção de impostos. Para prestadores de serviço que emitem NF para empresas do Simples Nacional, configurar o sistema de emissão para não reter impostos na fonte.

---

## Grupo: Duo e Wake Parcerias

### HOW_TO-4 — Processo de qualificação e repasse de leads SDR (Wake → Agência Parceira)

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Duo e Wake Parcerias |
| **JID** | `120363402057862373@g.us` |
| **Data** | 2025-06-25 |
| **Categoria** | `how_to` |
| **Tags** | `sdr`, `lead`, `qualificação`, `parceiro`, `wake`, `repasse` |

**Fluxo padrão de repasse de lead qualificado pelo SDR da Wake para agência parceira:**

1. **SDR qualifica o lead** e cria o resumo com:
   - URL do site do cliente
   - Necessidade principal (ex.: migração de plataforma, novo e-commerce)
   - ERP utilizado
   - Faturamento online estimado
   - Canais de venda atuais (lojas físicas, marketplaces)
   - Observações operacionais relevantes (ex.: cliente fora do Brasil, responde via WhatsApp)
   - Cases relevantes para referência (ex.: Construmarques, Wig)

2. **SDR compartilha o lead no grupo** da parceria com a agência.

3. **Proposta Wake:** A agência pode incluir a proposta comercial da Wake junto com sua proposta de implementação — coordenada pelo Canal Wake.

4. **Proposta de plataforma:** O SDR pode enviar a proposta da Wake diretamente ao cliente ou solicitar que a agência inclua junto à sua proposta.

5. **Registro no Salesforce (SF):** Leads/OPTs devem ser registrados. Em caso de instabilidade no SF, o Canal Wake registra manualmente. A agência deve informar o nome do cliente para abertura do lead/OPT.

---

### HOW_TO-5 — Como responder perguntas técnicas de pré-venda B2B na Wake

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Duo e Wake Parcerias |
| **JID** | `120363402057862373@g.us` |
| **Data** | 2025-07-10 |
| **Categoria** | `how_to` |
| **Tags** | `b2b`, `pre-venda`, `wake`, `duvidas-tecnicas`, `agencia` |

**Perguntas técnicas comuns em pré-vendas B2B e respostas oficiais Wake:**

**1. DIFAL para CNPJ**
As lojas geralmente não pagam DIFAL pedido a pedido, mas sim mensalmente aos estados onde operam. Não é possível editar o valor do pedido posteriormente. A empresa deve considerar o DIFAL na precificação, distribuindo uma média para todos os produtos, ou utilizar o recurso de **Regionalização** da Wake.
Documentação: https://atendimento.wake.tech/hc/pt-br/articles/23908030467223-Regionaliza%C3%A7%C3%A3o

**2. Kit de produtos (agrupamento de SKUs individuais)**
É possível vender um kit no site e fazer baixa unitária de cada SKU no ERP.
Documentação: https://atendimento.wake.tech/hc/pt-br/articles/26616430401175-Kit-de-produtos

**3. Múltiplos acessos por CNPJ (árvore de acesso B2B)**
É possível através da configuração "Permitir Cadastro de CNPJ Duplicado". Atenção: estratégias de crédito e limite por CNPJ precisam ser alinhadas com essa configuração.
Documentação: https://atendimento.wake.tech/hc/pt-br/articles/21406545532823-B2B

---

### HOW_TO-6 — Linguagens de programação e integrações nativas da Wake

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Duo e Wake Parcerias |
| **JID** | `120363402057862373@g.us` |
| **Data** | 2026-02-24 |
| **Categoria** | `how_to` |
| **Tags** | `wake`, `tecnologia`, `integrações`, `linguagens`, `frete` |

**Linguagens de programação da plataforma Wake:** `HTML`, `CSS`, `JavaScript` e `Scriban`.

**Agendamento de ativação/inativação de produto:** Não é possível via painel administrativo. É possível via API.

**Integrações nativas de frete disponíveis:**
- Correios (nativo)
- Loggi (nativo)
- GO Fretes, Frenet, Intelipost, Freterápido, SisFrete, Data Frete, Mandaê, Box Delivery

**Não possui integração nativa com:** Melhor Envio (mas pode ser configurado via parceiros de frete listados acima).

**ERP com integração nativa:** Bling e Tiny. Demais ERPs (ex.: Compels, Protheus, MILL GESTI, Sisplan) integram via API aberta.

**Módulo de assinaturas:** A Wake possui módulo nativo de assinaturas — avaliar antes de desenvolver solução customizada.

**Clube de assinaturas (referência de caso):** Pura Vida Club — https://www.puravida.com.br/puravida-club

---

### FAQ-6 — A certificação de parceiro Wake é obrigatória para receber leads?

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Duo e Wake Parcerias |
| **JID** | `120363402057862373@g.us` |
| **Data** | 2025-07-22 |
| **Categoria** | `faq` |
| **Tags** | `certificação`, `parceiro`, `wake`, `programa-de-parcerias`, `lead` |

**Pergunta:** A certificação de parceiro Wake é obrigatória para continuar recebendo leads no carrossel? É a mesma que a campanha de vendas?

**Resposta (Wake):**
- A certificação é **obrigatória para o fluxo de carrossel de leads** — é uma regra do programa de parcerias.
- Esta é uma **nova certificação** (lançada em 2025), diferente da anterior. A antiga não vale mais.
- É obrigatória **pelo menos 1 certificação por agência** (para o plano I). Mais pessoas certificadas = melhor.
- É a **mesma certificação para vendas e implementação** — individual por pessoa, não por empresa.
- A campanha de vendas **também exige** a certificação como pré-requisito para ser elegível aos prêmios.
- Prazo da campanha: **até 11/08/2025**.
- O fluxo de recebimento de leads (contratual) e a campanha de vendas têm regras distintas — mas a certificação é necessária para ambos.

---

### FAQ-7 — Quais ERPs têm integração nativa com a Wake?

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Duo e Wake Parcerias |
| **JID** | `120363402057862373@g.us` |
| **Data** | 2025-07-07 |
| **Categoria** | `faq` |
| **Tags** | `erp`, `integração`, `wake`, `bling`, `tiny` |

**Pergunta:** O ERP "Compels" tem integração nativa com a Wake?

**Resposta:** Não. A Wake possui integração **nativa apenas com Bling e Tiny**. Todos os outros ERPs (Compels, Protheus, Sankhya, MILL GESTI, Sisplan, SAP, etc.) integram via **API aberta** da Wake. Para esses casos, é necessário desenvolvimento de integração customizado pela agência ou pelo parceiro de ERP.

---

### FAQ-8 — Wake suporta e-commerce B2B com múltiplos CNPJs por empresa?

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Duo e Wake Parcerias |
| **JID** | `120363402057862373@g.us` |
| **Data** | 2025-07-10 |
| **Categoria** | `faq` |
| **Tags** | `b2b`, `cnpj`, `multiplos-acessos`, `wake`, `configuração` |

**Pergunta:** É possível ter múltiplos logins de compra para o mesmo CNPJ na Wake?

**Resposta:** Sim, através da configuração **"Permitir Cadastro de CNPJ Duplicado"**. Uma empresa (CNPJ) pode ter vários usuários com login independente para compras.

**Atenção:** Estratégias como análise de crédito ou limite por CNPJ precisam ser alinhadas com essa configuração, pois podem conflitar. Veja a documentação completa de recursos B2B para verificar compatibilidade entre funcionalidades: https://atendimento.wake.tech/hc/pt-br/articles/21406545532823-B2B

---

### BEST_PRACTICES-5 — Estratégia de co-venda Wake + Agência: como posicionar em deals competitivos

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Duo e Wake Parcerias |
| **JID** | `120363402057862373@g.us` |
| **Data** | 2026-02-26 |
| **Categoria** | `best_practices` |
| **Tags** | `co-venda`, `wake`, `posicionamento`, `proposta`, `plataforma-enterprise` |

**Lições de deals perdidos e ganhos no ecossistema Wake + agências parceiras:**

**Deal GDom Malhas (perdido para Magazord):**
- Cliente valorizou a Wake mas foi voto vencido internamente — a maioria decidiu pela plataforma de entrada (menor custo).
- Mensagem do cliente: *"A plataforma Wake é robusta, mas para este início de operação, optamos por um modelo mais enxuto. Quando nosso volume crescer, podemos retomar."*
- **Aprendizado:** Para clientes em início de operação com faturamento online baixo, posicionar uma proposta escalonada (fase 1 simples + evolução posterior) pode reduzir a objeção de preço.

**Posicionamento correto da Wake vs. concorrentes de entrada (Tray, Magazord, WBUY):**
- Wake não se compara funcionalmente a plataformas de entrada — é enterprise.
- Reforçar os ganhos de escala, integrações nativas, performance e flexibilidade que as plataformas menores não oferecem.
- Para leads com expectativa de ir a mercado em 15-30 dias: realinhar expectativa (mínimo 30 dias, media 60-90 dias) e oferecer uma fase de "go-live simplificado" com evolução incremental.

**Fluxo de proposta em deals conjuntos:**
1. SDR Wake qualifica e realiza demo inicial sem falar de valores de plataforma.
2. Agência faz sua apresentação e proposta de implementação.
3. SDR Wake apresenta proposta de plataforma em reunião separada, após todas as agências terem apresentado.
4. Negociação de desconto: possível até 10% nos percentuais + 2 primeiras mensalidades isentas (exemplo de campanha).

---

### BEST_PRACTICES-6 — Estrutura de qualificação de leads para e-commerce B2B/B2C complexo

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Duo e Wake Parcerias |
| **JID** | `120363402057862373@g.us` |
| **Data** | 2026-05-05 |
| **Categoria** | `best_practices` |
| **Tags** | `bant`, `qualificação`, `lead`, `b2b`, `b2c`, `co-venda` |

**Relatório BANT — modelo aplicado em deal Off Rush (autopeças, R$400k/mês):**

- **Budget:** Validar faturamento mensal + percentual atual do site no total. Se o site representa <5% do faturamento, a urgência pode ser baixa — é necessário construir o business case de ROI.
- **Authority:** Identificar os decisores técnicos e comerciais separadamente. Perfis diferentes exigem argumentos diferentes (ROI vs. integração técnica).
- **Need:** Dor bem verbalizada = processo mais rápido. Documentar exatamente as dores (ex.: "site lento, baixa conversão, alto atrito operacional, dependência do WhatsApp para concluir compras").
- **Timeline:** Alinhar expectativas de prazo reais. Clientes tendem a querer 15-30 dias; mínimo real é 30 dias (average 60-90 dias). Oferecer "go-live mínimo em 30 dias + evoluções posteriores".

**Fatores de risco em deals:**
- Preço total (plataforma + agência) não claro na cabeça do lead.
- Concorrência com proposta já na mesa de plataforma atual.
- Operação de marketplaces já funcionando reduz urgência de site próprio.
- Ausência de um campeão interno comprado com a solução.

---

### FAQ-9 — Como funciona o Partner Day e os benefícios do programa de parceiros Wake?

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Duo e Wake Parcerias |
| **JID** | `120363402057862373@g.us` |
| **Data** | 2026-02-10 |
| **Categoria** | `faq` |
| **Tags** | `partner-day`, `wake`, `programa-de-parcerias`, `benefícios`, `escritório` |

**O que é o Partner Day Wake:**
Encontro presencial semanal (às quartas-feiras) no escritório da LWSA (R. Itapaiúna, 2434 – Parque do Morumbi, SP), onde agências parceiras podem trabalhar no espaço físico da Wake. Estacionamento gratuito. Estação de metrô mais próxima: Giovanni Gronchi.

**Para participar:** Enviar lista de participantes (Nome, E-mail, Documento) para o Canal Wake com antecedência para liberação na portaria.

**Eventos especiais:** Periodicamente a Wake realiza o **Partner Day em locais temáticos** (ex.: escritório do TikTok em março de 2026 — Infinity Tower, Itaim Bibi).

**Wake Trends Commerce:** Workshop trimestral exclusivo para agências parceiras apresentando novidades de produto Wake, cases de uso práticos e espaço para dúvidas e troca de experiências.

**Suporte técnico fora do horário:** Canal oficial: https://wakecommerce.readme.io/docs/canais-de-atendimento-para-parceiros

---

### TROUBLESHOOTING-9 — Lead rejeitado após proposta: feedback e análise de deal perdido

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Duo e Wake Parcerias |
| **JID** | `120363402057862373@g.us` |
| **Data** | 2026-03-09 |
| **Categoria** | `troubleshooting` |
| **Tags** | `deal-perdido`, `feedback`, `wake`, `agencia`, `plataforma-entrada` |

**Caso:** GDom Malhas — fechou com Magazord após avanço significativo na negociação com Wake.

**Feedback do cliente:**
> "No momento, decidimos migrar para uma solução de menor custo e estrutura, que se alinha melhor ao nosso fluxo atual de vendas. Entendemos que a plataforma WAKE é robusta, mas para este início de operação, optamos por um modelo mais enxuto. Quem sabe daqui a algum tempo, conforme nosso volume de vendas crescer e houver necessidade de uma estrutura maior, possamos retomar essa conversa para uma futura migração."

**Análise:** O decisor técnico (gerente de e-commerce) estava comprado com a Wake, mas foi voto vencido. A decisão foi tomada por critério de custo inicial, não de funcionalidade. Cliente com faturamento online baixo e plataforma atual funcional tem baixa urgência de upgrade.

**Ações recomendadas para nutrir:** Manter contato trimestral. Quando o GMV online atingir threshold de justificativa de upgrade (tipicamente R$ 50-100k/mês), reativar a conversa de migração.
