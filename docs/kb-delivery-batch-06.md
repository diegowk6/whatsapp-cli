# KB Delivery — Batch 06
**Artigos**: 349–436 do registry  
**Fontes**: kb-parceiros-marketplace.md (completo), kb-projeto-hero-leanwork.md (completo), kb-parceiros-wave7a.md (completo)  
**Total**: 88 artigos  
**Status**: prontos para importação  
**Data**: 2026-06

---
# KB Extract — Parceiros Marketplace (Blueman, Anjo Rosa, Duo)
> Extracted: 2026-06-04

---

## Grupo: Wake + Marketplace Blueman

### TRB-001 — Pedido não integrado por falta de nome de exibição na forma de pagamento

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Wake + Marketplace Blueman |
| **Data** | 2025-06-20 |
| **Categoria** | `troubleshooting` |
| **Tags** | `integração`, `audience-solutions`, `audience-cs` |

**Sintomas:** Pedido criado no Mercado Livre não desceu para o sistema de faturamento (síntese/ERP). Mensagem de erro: `não será possivel importar o pedido: [ID]. o tipo pgto não foi informando, provavelmente o nome exibição da formapagamentoid:200 na wake está em branco`

**Causa raiz:** O sistema buscou os detalhes da forma de pagamento via API `/formasPagamento` da Wake para identificar o tipo (Pix, boleto, cartão). O campo "nome de exibição" da `formapagamentoid:200` estava em branco na Wake.

**Resolução:** Preencher o campo "nome de exibição" da forma de pagamento correspondente no painel da Wake.

---

### TRB-002 — Bug no agrupamento de variantes bloqueando catálogo no Mercado Livre

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Wake + Marketplace Blueman |
| **Data** | 2025-06-20 |
| **Categoria** | `troubleshooting` |
| **Tags** | `integração`, `audience-solutions`, `audience-cs` |

**Sintomas:** Impossibilidade de subir novos anúncios no Mercado Livre via Anymarket por conta de bug no agrupamento de variantes. O bug impedia a integração do catálogo Wake na Anymarket, bloqueando a publicação de anúncios no Mercado Livre para a Blueman.

**Causa raiz:** Bug no agrupamento de variantes na Anymarket — [fonte original sem informação] sobre detalhes técnicos do bug.

**Resolução:** Bug foi para fase de homologação na Anymarket com previsão de correção no início da semana seguinte (segunda-feira). Após a correção confirmada, o catálogo Wake seria integrado à Anymarket e a Blueman poderia subir novos anúncios.

---

### TRB-003 — Pedidos de marketplace não aparecem no financeiro após mudança de JSON

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Wake + Marketplace Blueman |
| **Data** | 2025-10-29 |
| **Categoria** | `troubleshooting` |
| **Tags** | `contrato-comercial`, `audience-solutions`, `audience-cs` |

**Sintomas:** Pedidos de marketplaces (Riachuelo, TikTok Shop, Mercado Livre) pagos e integrados não apareciam no financeiro/ERP. O sistema não gerava parcelas nem o movimento financeiro dos pedidos.

**Causa raiz:** A Anymarket mudou o padrão do JSON enviado, quebrando o DE x PARA configurado no ERP para a perna financeira.

**Resolução:** O desenvolvedor do ERP revisou e remapeou o DE x PARA para o novo padrão de JSON. Cada pedido existente antes da melhoria precisou ser tratado manualmente. Para novos pedidos, a parametrização ficou automática. Aprendizado: mudanças de formato no JSON da Anymarket podem quebrar silenciosamente integrações financeiras. É necessário monitoramento de pedidos no financeiro após atualizações da plataforma de integração.

---

### TRB-004 — Pedido Mercado Livre sem CPF bloqueia emissão de NF-e

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Wake + Marketplace Blueman |
| **Data** | 2025-11-12 |
| **Categoria** | `troubleshooting` |
| **Tags** | `contrato-comercial`, `audience-cs`, `audience-solutions` |

**Sintomas:** Pedido do Mercado Livre chegou sem CPF do comprador. Como CPF é obrigatório para emissão de NF-e no Brasil, o pedido não pôde ser faturado automaticamente. Impactou o prazo de envio e a reputação da loja no ML.

**Causa raiz:** O Mercado Livre permite que compradores concluam a compra sem CPF. O CPF pode ser preenchido depois, mas a Anymarket não atualiza automaticamente na Wake um pedido já exportado.

**Resolução:** Identificar o cliente manualmente, inserir o CPF e faturar. Não há bloqueio automático disponível na Anymarket para pedidos sem CPF, pois cancelamentos automáticos prejudicam a reputação da loja.

**Workaround:** Monitorar pedidos que chegam sem CPF e tratar manualmente. Qualquer cancelamento deve ser ponderado contra o impacto na reputação.

---

### TRB-005 — Status "Sincronizando venda com marketplace" travado — Riachuelo

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Wake + Marketplace Blueman |
| **Data** | 2026-01-26 |
| **Categoria** | `troubleshooting` |
| **Tags** | `suporte-escalação`, `audience-cs`, `audience-solutions` |

**Sintomas:** Dois pedidos ficaram presos com status "Sincronizando venda com marketplace" desde o dia 14/01. Isso impedia a atualização para "entregue" na Riachuelo, o que bloqueava o recebimento do valor da venda.

**Causa raiz:** [fonte original sem informação] — falha de sincronização na Anymarket que exigiu intervenção manual do suporte para destravar o pedido.

**Resolução:** Abertura de chamado no suporte da Anymarket + chamado paralelo na Riachuelo. O time de suporte da Anymarket executou ação de voltar e avançar o status do pedido manualmente. Com isso o status "entregue" refletiu na Riachuelo e o recebimento foi liberado. Observação: a plataforma da Riachuelo foi descontinuada em 30/01/2026, tornando esse marketplace inativo para a Blueman a partir dessa data.

---

### TRB-006 — Integração com Riachuelo exige token gerado via Seller Center

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Wake + Marketplace Blueman |
| **Data** | 2025-08-07 |
| **Categoria** | `troubleshooting` |
| **Tags** | `integração`, `audience-solutions` |

**Sintomas:** Integração com a Riachuelo não estava puxando informações, mesmo com acesso ao portal disponível.

**Causa raiz:** O token necessário para a integração precisa ser solicitado diretamente pelo Seller Center da Riachuelo. As informações do portal padrão não são suficientes.

**Resolução:** Abrir solicitação no Seller Center da Riachuelo para geração do token de integração. Após obter o token, configurar na Anymarket.

---

### TRB-007 — SKUs com estrutura incompatível entre Wake e anúncios ML via Anymarket

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Wake + Marketplace Blueman |
| **Data** | 2025-10-20 |
| **Categoria** | `troubleshooting` |
| **Tags** | `integração`, `audience-solutions` |

**Sintomas:** Produtos cadastrados na Wake agrupados apenas por tamanho (sem agrupamento por cor) geraram incompatibilidade com anúncios no Mercado Livre que precisavam vincular múltiplos SKUs (por cor + tamanho) em um único anúncio.

**Causa raiz:** A estrutura do anúncio no ML era diferente do cadastro na Wake: o anúncio esperava agrupamento por cor e tamanho, mas o cadastro Wake tinha apenas tamanho único. A Anymarket não consegue vincular vários produtos em um único anúncio nessa situação.

**Resolução:** Pausar os SKUs incompatíveis no ML. Criar uma planilha de mapeamento com campo "SKU WAKE" para validação cruzada entre Anymarket e Wake antes de reativar anúncios.

---

### HTO-001 — Como parametrizar formas de pagamento por marketplace no ERP

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Wake + Marketplace Blueman |
| **Data** | 2025-10-31 |
| **Categoria** | `how_to` |
| **Tags** | `integração`, `audience-solutions` |

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

### HTO-002 — Fluxo de treinamento Anymarket para operadores de marketplace

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Wake + Marketplace Blueman |
| **Data** | 2025-08-06 |
| **Categoria** | `how_to` |
| **Tags** | `marketing-eventos`, `audience-cs`, `audience-solutions` |

Após conclusão dos ajustes de integração Wake-Anymarket-Marketplace, é necessário treinar o time operacional da loja:

1. **Identificar participantes:** Definir quem será treinado (ex: responsável pelo marketplace + equipe de atendimento/SAC).
2. **Agendar sessão com a Anymarket:** O time da Anymarket conduz o treinamento sobre como operar a plataforma — incluindo gestão de anúncios, monitoramento de pedidos e resolução de erros comuns.
3. **Duração sugerida:** 1h30 a 2h por sessão; reservar sessão adicional caso seja necessário aprofundar temas.
4. **Tópicos prioritários:** gestão de catálogo, vínculo de SKUs, fluxo de pedidos, emissão de etiquetas, monitoramento de status de entrega.

---

### BPR-001 — Controle de status de pedidos de marketplace — lições aprendidas (Blueman)

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Wake + Marketplace Blueman |
| **Data** | 2025-11-13 |
| **Categoria** | `best_practices` |
| **Tags** | `ecossistema-parceiros`, `audience-cs`, `audience-solutions` |

**Recomendação:** Adotar os seguintes controles operacionais ao integrar lojas em marketplaces via Anymarket:
- **Monitoramento contínuo é obrigatório:** pedidos podem ficar presos em status intermediários (ex: "Sincronizando", "Aguardando faturamento") sem alerta automático. É necessário processo manual de verificação regular, especialmente antes de datas como Black Friday
- **Não operar em produção sem integração 100% testada:** integrações parciais (ex: apenas um marketplace funcionando corretamente) geram retrabalho e risco operacional
- **Etiqueta de envio ML deve ser gerada automaticamente:** a emissão manual de etiqueta do Mercado Livre gera risco de envio fora do prazo e impacto na reputação da loja. A API de emissão de etiquetas da Anymarket (`/api/v2/emissao-de-etiquetas`) deve ser integrada com o OMS para emissão automática após faturamento
- **Manter múltiplos CDs habilitados:** Riachuelo e Mercado Livre podem vir configurados apenas com o CD principal — os demais estoques precisam ser habilitados separadamente na Anymarket
- **Informar XML da NF por e-mail correto:** o e-mail de importação de XML da Anymarket (`importacaoxml.anymarket@db1.com.br`) deve estar configurado corretamente no sistema de emissão de NF-e. Verificar periodicamente a caixa de entrada e o espaço disponível

**Rationale:** Principais aprendizados operacionais coletados ao longo da implementação da Blueman nos marketplaces. A Blueman operou quase um ano com integrações parciais (apenas Riachuelo funcionando corretamente), o que gerou retrabalho e risco operacional — esse padrão de "operar com integração incompleta" é a principal fonte de perda de receita e dano de reputação no canal de marketplaces.

---

### FAQ-001 — A Anymarket suporta emissão automática de etiquetas para Mercado Livre?

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Wake + Marketplace Blueman |
| **Data** | 2026-01-28 |
| **Categoria** | `faq` |
| **Tags** | `integração`, `audience-solutions` |

**Pergunta:** É possível que as etiquetas de envio do Mercado Livre sejam emitidas automaticamente ao faturar o pedido?

**Resposta:** A Anymarket possui uma API para emissão de etiquetas (`https://developers.anymarket.com.br/api/v2/b8c7fdc94a369-emissao-de-etiquetas`) que pode ser consumida pelo OMS ou ERP. Porém, essa integração precisa ser desenvolvida pelo time do OMS — a Anymarket fornece o endpoint, mas não faz a integração automaticamente com todos os sistemas de gestão. Verifique se o OMS utilizado já possui esse conector disponível antes de desenvolver uma solução customizada.

---

### FAQ-002 — Como funciona a cobrança de TR para lojas sem finalização de compra no site?

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Wake + Marketplace Blueman |
| **Data** | 2025-08-21 |
| **Categoria** | `faq` |
| **Tags** | `contrato-comercial`, `audience-sales`, `audience-solutions` |

**Pergunta:** Um cliente quer usar a Wake como vitrine com checkout, mas sem finalizar a compra no site — o pedido vai para um time de vendas consultivo. Como funciona a cobrança de TR nesse cenário?

**Resposta (Wake):** A cobrança de TR (Transaction Rate) é feita por pedido aprovado. Se a compra não finaliza no site, não há cobrança de TR padrão. No entanto, a Wake precisará de uma negociação diferente de cobrança — por exemplo, cobrança por ordem de pedido gerada ao invés de pedido aprovado. Esse cenário já foi implementado com outros clientes, mas exige um acordo comercial específico com o time da Wake.

---

## Grupo: Comercial Anjo Rosa

### FAQ-003 — Qual o fluxo de contratação de publicidade com exclusividade para influenciadores?

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Comercial Anjo Rosa |
| **Data** | 2026-03-11 |
| **Categoria** | `faq` |
| **Tags** | `lead-pipeline`, `audience-sales` |

**Pergunta:** Como precificar e estruturar uma campanha com cláusula de exclusividade?

**Resposta:** Quando uma marca solicita exclusividade (ex.: sem menção a concorrentes durante o período de veiculação), aplica-se um acréscimo sobre o valor tabelado padrão. A prática de mercado é de **pelo menos 30% de acréscimo** sobre o valor da tabela base para cobrir o custo de oportunidade de recusar outras marcas do segmento.

**Exemplo prático:** Tabela base: R$ 9.200,00 → Com exclusividade de 3 meses: R$ 11.960,00 (mín.)

**Atenção:** Se a campanha ocorrer em meses de alta demanda (ex.: Setembro/Outubro), considerar acréscimo adicional de 20% por sazonalidade.

---

### FAQ-004 — Qual é o escopo padrão de uma campanha com presença em evento (meet point)?

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Comercial Anjo Rosa |
| **Data** | 2026-03-27 |
| **Categoria** | `faq` |
| **Tags** | `lead-pipeline`, `audience-sales` |

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

### FAQ-005 — Como funciona o programa de afiliadas (ex.: Desinchá)?

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Comercial Anjo Rosa |
| **Data** | 2026-04-13 |
| **Categoria** | `faq` |
| **Tags** | `programa-parceiros`, `audience-pam` |

**Modelo padrão de programa de afiliadas de marca (health/wellness):**

- **Envio de produtos:** Mensal (para uso e review).
- **Comissão:** 10% sobre vendas geradas via cupom e/ou link rastreável.
- **Cadência de conteúdo:** ~4 stories por mês (média de 1x por semana).
- **Onboarding:** Preenchimento de formulário de cadastro → inclusão em grupo de WhatsApp da marca → recebimento de briefings, direcionamentos e novidades.

**Fluxo de cadastro:** Formulário → sistema interno da marca → grupo WhatsApp → início das ativações.

---

### HTO-003 — Como emitir Carta de Correção Eletrônica (CC-e) para ajuste em NF-e já autorizada

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Comercial Anjo Rosa |
| **Data** | 2026-04-29 |
| **Categoria** | `how_to` |
| **Tags** | `contrato-comercial`, `audience-sales` |

**Contexto:** O financeiro do contratante (agência de mídia/veículo) solicitou ajuste na descrição da NF após emissão. A NF estava correta conforme briefing original, mas houve troca de responsável no financeiro do cliente e foi solicitada uma nova descrição.

**O que é a CC-e:** A Carta de Correção Eletrônica (CC-e) é um documento fiscal digital usado para corrigir erros em uma NF-e já autorizada pela SEFAZ, sem cancelá-la.

**Quando usar:** Quando há erro em campos de texto/descrição da NF que não alteram valores, impostos ou dados principais. Não serve para corrigir valores, CNPJ do destinatário ou dados de produto.

**Atenção sobre Simples Nacional:** Empresas do Simples Nacional não aceitam NFs com retenção de impostos sob responsabilidade do tomador. Se o influenciador/prestador for optante do Simples, emitir a NF sem retenção de ISS/PIS/COFINS/CSLL pelo tomador.

---

### BPR-002 — Estrutura de proposta comercial para campanhas de influência — formato mensal vs. pacote fechado

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Comercial Anjo Rosa |
| **Data** | 2025-09-06 |
| **Categoria** | `best_practices` |
| **Tags** | `proposta`, `audience-sales` |

**Recomendação:** Ao apresentar proposta de campanha de influência, estruturar como preço mensal condicionado à contratação do período total, não apenas mês a mês. Exemplo: "O valor mensal de R$ X está condicionado ao fechamento do período completo de [N] meses. Para o próximo ciclo, apresentaremos nova proposta de acordo com o desempenho."

**Rationale:** Esse formato:
- Garante o budget completo da campanha para a marca
- Deixa espaço para negociação de pagamento sem abrir mão do escopo total
- Facilita renovação: ao final do período, apresentar nova tabela para o ciclo seguinte

Propostas mensais avulsas dão à marca a possibilidade de cancelar sem aviso e tornam difícil planejar produção de conteúdo e disponibilidade.

---

### BPR-003 — Gestão de relacionamento com agências de mídia (intermediárias)

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Comercial Anjo Rosa |
| **Data** | 2026-04-24 |
| **Categoria** | `best_practices` |
| **Tags** | `proposta`, `audience-sales` |

**Recomendação:** Ao trabalhar com agências de mídia intermediárias (ex.: Pulso Digital, ID - Impulso Digital para o Estadão), aplicar os seguintes controles:
1. **Prazo de pagamento:** agências de mídia trabalham com prazos longos (ex.: 100 dias após recebimento da NF). Considerar no fluxo de caixa antes de fechar
2. **NF só após liberação:** a NF deve ser emitida apenas após liberação expressa da marca contratante — nunca antes. Isso é cláusula padrão dos contratos
3. **Confidencialidade do briefing:** as agências costumam solicitar que os valores e termos do briefing não sejam divulgados. Tratar como confidencial
4. **Troca de contato:** é comum haver rotatividade de pessoas no cliente/agência. Sempre formalizar as condições por e-mail e confirmar o ponto de contato atual antes de emitir NF
5. **Registro no CRM:** registrar todos os projetos em CRM (ex.: Moskit) assim que fechados, com nome do cliente, valor e status da NF

**Rationale:** Trabalhar com agências de mídia exige disciplina contratual e financeira diferenciada — prazos longos exigem reserva de caixa, e emitir NF antes da liberação resulta em recusa do tomador. Sem formalizar trocas de contato, condições negociadas podem ser perdidas e cobranças ficam órfãs.

---

### BPR-004 — Prospecção em eventos de saúde e inovação para influenciadores B2B

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Comercial Anjo Rosa |
| **Data** | 2026-05-05 |
| **Categoria** | `best_practices` |
| **Tags** | `lead-pipeline`, `audience-sales`, `audience-pam` |

**Recomendação:** Participar de eventos de saúde e inovação como estratégia dupla de cobertura de conteúdo + prospecção B2B para campanhas de influência futuras:
- **São Paulo Innovation Week:** evento multissetorial de inovação em SP — útil para conectar com empresas de tecnologia, startups e marcas que investem em marketing de conteúdo
- **Eventos hospitalares e oncológicos:** eventos como Meet Point (A.C. Camargo/Estadão) são oportunidades de aparecer como referência no nicho de saúde feminina, oncologia e bem-estar

**Rationale:** Participar como influenciador (cobertura de conteúdo) e simultaneamente prospectar novas marcas presentes no evento maximiza o ROI de tempo investido — a presença gera autoridade no nicho e contato direto com decisores de marcas que valorizam esse posicionamento.

---

### TRB-008 — NF com retenção de impostos recusada por empresa do Simples Nacional

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Comercial Anjo Rosa |
| **Data** | 2026-05-04 |
| **Categoria** | `troubleshooting` |
| **Tags** | `contrato-comercial`, `audience-sales` |

**Sintomas:** NF emitida com retenção de impostos (ISS/PIS/COFINS/CSLL) foi recusada pelo financeiro do cliente.

**Causa raiz:** A empresa contratante era optante do Simples Nacional, que não pode aceitar NFs com retenção de impostos sob responsabilidade do tomador.

**Resolução:** Reemitir a NF (ou emitir CC-e se aplicável) sem retenção de impostos. Para prestadores de serviço que emitem NF para empresas do Simples Nacional, configurar o sistema de emissão para não reter impostos na fonte.

---

## Grupo: Duo e Wake Parcerias

### HTO-004 — Processo de qualificação e repasse de leads SDR (Wake → Agência Parceira)

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Duo e Wake Parcerias |
| **Data** | 2025-06-25 |
| **Categoria** | `how_to` |
| **Tags** | `lead-pipeline`, `audience-pam`, `audience-sales` |

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

### HTO-005 — Como responder perguntas técnicas de pré-venda B2B na Wake

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Duo e Wake Parcerias |
| **Data** | 2025-07-10 |
| **Categoria** | `how_to` |
| **Tags** | `integração`, `audience-solutions`, `audience-sales` |

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

### HTO-006 — Linguagens de programação e integrações nativas da Wake

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Duo e Wake Parcerias |
| **Data** | 2026-02-24 |
| **Categoria** | `how_to` |
| **Tags** | `integração`, `audience-solutions` |

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

### FAQ-006 — A certificação de parceiro Wake é obrigatória para receber leads?

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Duo e Wake Parcerias |
| **Data** | 2025-07-22 |
| **Categoria** | `faq` |
| **Tags** | `programa-parceiros`, `audience-pam` |

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

### FAQ-007 — Quais ERPs têm integração nativa com a Wake?

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Duo e Wake Parcerias |
| **Data** | 2025-07-07 |
| **Categoria** | `faq` |
| **Tags** | `integração`, `audience-solutions`, `audience-sales` |

**Pergunta:** O ERP "Compels" tem integração nativa com a Wake?

**Resposta:** Não. A Wake possui integração **nativa apenas com Bling e Tiny**. Todos os outros ERPs (Compels, Protheus, Sankhya, MILL GESTI, Sisplan, SAP, etc.) integram via **API aberta** da Wake. Para esses casos, é necessário desenvolvimento de integração customizado pela agência ou pelo parceiro de ERP.

---

### FAQ-008 — Wake suporta e-commerce B2B com múltiplos CNPJs por empresa?

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Duo e Wake Parcerias |
| **Data** | 2025-07-10 |
| **Categoria** | `faq` |
| **Tags** | `integração`, `audience-solutions`, `audience-sales` |

**Pergunta:** É possível ter múltiplos logins de compra para o mesmo CNPJ na Wake?

**Resposta:** Sim, através da configuração **"Permitir Cadastro de CNPJ Duplicado"**. Uma empresa (CNPJ) pode ter vários usuários com login independente para compras.

**Atenção:** Estratégias como análise de crédito ou limite por CNPJ precisam ser alinhadas com essa configuração, pois podem conflitar. Veja a documentação completa de recursos B2B para verificar compatibilidade entre funcionalidades: https://atendimento.wake.tech/hc/pt-br/articles/21406545532823-B2B

---

### BPR-005 — Estratégia de co-venda Wake + Agência: como posicionar em deals competitivos

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Duo e Wake Parcerias |
| **Data** | 2026-02-26 |
| **Categoria** | `best_practices` |
| **Tags** | `lead-pipeline`, `audience-pam`, `audience-sales` |

**Recomendação:** Em deals competitivos Wake + agência, aplicar posicionamento e fluxo de proposta estruturados:

Posicionamento correto da Wake vs. concorrentes de entrada (Tray, Magazord, WBUY):
- Wake não se compara funcionalmente a plataformas de entrada — é enterprise
- Reforçar os ganhos de escala, integrações nativas, performance e flexibilidade que as plataformas menores não oferecem
- Para leads com expectativa de ir a mercado em 15-30 dias: realinhar expectativa (mínimo 30 dias, média 60-90 dias) e oferecer uma fase de "go-live simplificado" com evolução incremental

Fluxo de proposta em deals conjuntos:
1. SDR Wake qualifica e realiza demo inicial sem falar de valores de plataforma
2. Agência faz sua apresentação e proposta de implementação
3. SDR Wake apresenta proposta de plataforma em reunião separada, após todas as agências terem apresentado
4. Negociação de desconto: possível até 10% nos percentuais + 2 primeiras mensalidades isentas (exemplo de campanha)

**Rationale:** Lições de deals perdidos no ecossistema Wake + agências parceiras. Deal GDom Malhas (perdido para Magazord): cliente valorizou a Wake mas foi voto vencido internamente — a maioria decidiu pela plataforma de entrada (menor custo). Mensagem do cliente: *"A plataforma Wake é robusta, mas para este início de operação, optamos por um modelo mais enxuto. Quando nosso volume crescer, podemos retomar."* Para clientes em início de operação com faturamento online baixo, posicionar uma proposta escalonada (fase 1 simples + evolução posterior) reduz a objeção de preço sem desvalorizar o produto.

---

### BPR-006 — Estrutura de qualificação de leads para e-commerce B2B/B2C complexo

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Duo e Wake Parcerias |
| **Data** | 2026-05-05 |
| **Categoria** | `best_practices` |
| **Tags** | `lead-pipeline`, `audience-sales`, `audience-pam` |

**Recomendação:** Aplicar análise BANT estruturada em qualificação de leads para e-commerce B2B/B2C complexo (modelo aplicado em deal Off Rush — autopeças, R$400k/mês):
- **Budget:** validar faturamento mensal + percentual atual do site no total. Se o site representa <5% do faturamento, a urgência pode ser baixa — é necessário construir o business case de ROI
- **Authority:** identificar os decisores técnicos e comerciais separadamente. Perfis diferentes exigem argumentos diferentes (ROI vs. integração técnica)
- **Need:** dor bem verbalizada = processo mais rápido. Documentar exatamente as dores (ex.: "site lento, baixa conversão, alto atrito operacional, dependência do WhatsApp para concluir compras")
- **Timeline:** alinhar expectativas de prazo reais. Clientes tendem a querer 15-30 dias; mínimo real é 30 dias (média 60-90 dias). Oferecer "go-live mínimo em 30 dias + evoluções posteriores"

Fatores de risco a monitorar em deals:
- Preço total (plataforma + agência) não claro na cabeça do lead
- Concorrência com proposta já na mesa de plataforma atual
- Operação de marketplaces já funcionando reduz urgência de site próprio
- Ausência de um campeão interno comprado com a solução

**Rationale:** Qualificação BANT estruturada permite previsibilidade de fechamento e priorização correta de esforço. Sem ela, agências e SDRs Wake gastam tempo em deals desqualificados ou perdem deals quentes por falta de identificação clara dos decisores.

---

### FAQ-009 — Como funciona o Partner Day e os benefícios do programa de parceiros Wake?

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Duo e Wake Parcerias |
| **Data** | 2026-02-10 |
| **Categoria** | `faq` |
| **Tags** | `marketing-eventos`, `audience-pam` |

**O que é o Partner Day Wake:**
Encontro presencial semanal (às quartas-feiras) no escritório da LWSA (R. Itapaiúna, 2434 – Parque do Morumbi, SP), onde agências parceiras podem trabalhar no espaço físico da Wake. Estacionamento gratuito. Estação de metrô mais próxima: Giovanni Gronchi.

**Para participar:** Enviar lista de participantes (Nome, E-mail, Documento) para o Canal Wake com antecedência para liberação na portaria.

**Eventos especiais:** Periodicamente a Wake realiza o **Partner Day em locais temáticos** (ex.: escritório do TikTok em março de 2026 — Infinity Tower, Itaim Bibi).

**Wake Trends Commerce:** Workshop trimestral exclusivo para agências parceiras apresentando novidades de produto Wake, cases de uso práticos e espaço para dúvidas e troca de experiências.

**Suporte técnico fora do horário:** Canal oficial: https://wakecommerce.readme.io/docs/canais-de-atendimento-para-parceiros

---

### TRB-009 — Lead rejeitado após proposta: feedback e análise de deal perdido

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Duo e Wake Parcerias |
| **Data** | 2026-03-09 |
| **Categoria** | `troubleshooting` |
| **Tags** | `lead-pipeline`, `audience-pam`, `audience-sales` |

**Sintomas:** GDom Malhas — lead fechou com Magazord após avanço significativo na negociação com Wake. Feedback do cliente: "No momento, decidimos migrar para uma solução de menor custo e estrutura, que se alinha melhor ao nosso fluxo atual de vendas. Entendemos que a plataforma WAKE é robusta, mas para este início de operação, optamos por um modelo mais enxuto. Quem sabe daqui a algum tempo, conforme nosso volume de vendas crescer e houver necessidade de uma estrutura maior, possamos retomar essa conversa para uma futura migração."

**Causa raiz:** O decisor técnico (gerente de e-commerce) estava comprado com a Wake, mas foi voto vencido. A decisão foi tomada por critério de custo inicial, não de funcionalidade. Cliente com faturamento online baixo e plataforma atual funcional tem baixa urgência de upgrade.

**Resolução:** Manter contato trimestral. Quando o GMV online atingir threshold de justificativa de upgrade (tipicamente R$ 50-100k/mês), reativar a conversa de migração.
# KB Extract — Projeto Hero <> Leanwork <> Wake
> JID: 120363417338422570@g.us | Extracted: 2026-06-04

---

## FAQ

---

### FAQ-001 — Diferença entre "bilhete" e "apólice" no contexto Hero Seguros

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Projeto Hero <> Leanwork <> Wake |
| **Data** | 2025-11-12 |
| **Categoria** | `faq` |
| **Tags** | `contrato-comercial`, `audience-solutions` |

**Pergunta:** O documento que o cliente recebe após contratar o seguro é chamado de apólice?

**Resposta:** Não. Na Hero Seguros, o documento emitido é chamado de **bilhete**, não de apólice. Esse é o termo correto e deve ser usado em toda a comunicação com o cliente e nas integrações técnicas.

---

### FAQ-002 — Ambiente de homologação da API Hero Seguros: necessidade de liberação de IP

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Projeto Hero <> Leanwork <> Wake |
| **Data** | 2025-09-17 |
| **Categoria** | `faq` |
| **Tags** | `integração`, `audience-solutions` |

**Pergunta:** O ambiente de homologação da API Hero Seguros exige liberação de IP, assim como o de produção?

**Resposta:** Não. Em homologação não é necessária liberação de IP. A restrição de whitelist de IP se aplica apenas ao ambiente de produção.

- **URL base de homologação:** `https://api.homologacao.heroseguros.com.br`
- Auth na raiz: `POST /oauth/token`
- Demais chamadas sob: `/api/..`

---

### FAQ-003 — Credenciais de autenticação da API Hero Seguros (ambiente de homologação)

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Projeto Hero <> Leanwork <> Wake |
| **Data** | 2025-09-17 |
| **Categoria** | `faq` |
| **Tags** | `integração`, `audience-solutions` |

**Pergunta:** Como autenticar na API de homologação da Hero Seguros?

**Resposta:** A autenticação utiliza OAuth2 com grant type `password`. Exemplo de requisição de token para o ambiente de homologação:

```bash
curl --location --request POST 'https://api.homologacao.heroseguros.com.br/oauth/token' \
  --form 'grant_type="password"' \
  --form 'client_id="266"' \
  --form 'client_secret="EAWvWG4BTn31v5Lp4TJzGVigVLs7j1WPkN8Tyshw"' \
  --form 'username="wake@heroseguros.com.br"' \
  --form 'password="vkw96eoN4a16"' \
  --form 'scope=""'
```

> **Atenção:** Para testes de Garantia Estendida vinculados à Karcher, usar as credenciais do usuário `hero@karcher.com.br` (client_id 278), pois o usuário `wake@heroseguros.com.br` (client_id 266) é exclusivo do Parcela Garantida.

---

### FAQ-004 — Payload da API de cotação do Parcela Garantida (Prestamista)

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Projeto Hero <> Leanwork <> Wake |
| **Data** | 2025-09-26 |
| **Categoria** | `faq` |
| **Tags** | `integração`, `audience-solutions` |

**Pergunta:** Qual o payload correto para chamar a API de cotação do Parcela Garantida?

**Resposta:**

```bash
curl --location --request POST 'https://api.homologacao.heroseguros.com.br/api/prestamista/quotation' \
  --header 'Content-Type: application/json' \
  --header 'Authorization: Bearer <token>' \
  --data-raw '{
    "debt_amount": 4000.00,
    "installments": 10,
    "type_of_product": 5,
    "last_installment_date": "2026-08-27",
    "customer": {
        "age": 30
    }
}'
```

Campos obrigatórios:
- `debt_amount`: valor da dívida (compras + frete, **sem incluir o valor do seguro**)
- `installments`: número de parcelas
- `type_of_product`: ID do tipo de produto. O ID varia por parceiro (ex: Karcher usa `5`, não `4`)
- `customer.age`: idade do cliente

---

### FAQ-005 — Payload da API de cotação da Garantia Estendida

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Projeto Hero <> Leanwork <> Wake |
| **Data** | 2026-02-11 |
| **Categoria** | `faq` |
| **Tags** | `integração`, `audience-solutions` |

**Pergunta:** Como chamar a API de cotação da Garantia Estendida?

**Resposta:** A busca do produto é feita por SKU (`external_type: "sku"`). Exemplo:

```bash
curl --location --request POST 'https://api.homologacao.heroseguros.com.br/api/extended-warranty/quotation' \
  --header 'Content-Type: application/json' \
  --header 'Authorization: Bearer <token>' \
  --data-raw '{
    "product_value": 2000.0,
    "searchProduct": {
        "external_type": "sku",
        "external_value": "30055"
    },
    "factory_warranty_end_date": "2026-12-01"
}'
```

> A `factory_warranty_end_date` é solicitada pela Hero, mas **não é usada como fator de cálculo** pela Hero Seguros. Para a Karcher, pode ser padronizada para 12 meses.

---

### FAQ-006 — Stack de infraestrutura recomendada para o middleware de integração

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Projeto Hero <> Leanwork <> Wake |
| **Data** | 2025-09-19 |
| **Categoria** | `faq` |
| **Tags** | `integração`, `audience-solutions` |

**Pergunta:** Qual stack de infraestrutura usar para hospedar o middleware de integração Hero/Wake?

**Resposta (Hero Seguros):**
- Sistema Operacional: **100% Linux**
- Banco de dados: **MySQL** ou, preferencialmente, **Aurora**

A stack de desenvolvimento utilizada pela Leanwork no projeto é **C# .NET**.

---

### FAQ-007 — Regras de elegibilidade do cliente para o Parcela Garantida (Prestamista)

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Projeto Hero <> Leanwork <> Wake |
| **Data** | 2025-09-08 |
| **Categoria** | `faq` |
| **Tags** | `integração`, `audience-solutions` |

**Questão levantada:** A Fase 1 do projeto não implementou as regras de negócio de elegibilidade. Identificadas para a Fase 2:

- Idade mínima do cliente: **18 anos**
- Idade máxima: **69 anos** (calculada como: idade na adesão + tempo de empréstimo)
- Outras regras a detalhar na Fase 2

---

### FAQ-008 — Como testar a integração na loja Karcher (ambiente de preview)

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Projeto Hero <> Leanwork <> Wake |
| **Data** | 2026-01-29 |
| **Categoria** | `faq` |
| **Tags** | `homologação`, `audience-solutions` |

**Pergunta:** Como testar a integração Hero na loja da Karcher sem impactar produção?

**Resposta:** Abrir uma aba anônima e seguir os passos nesta ordem:

1. Acessar primeiro: `https://karcher-center-altex.com.br/?preview-theme=d%2bvAEzPgMt3feu7mNxiD%2fQ%3d%3d`
2. Na mesma aba anônima, acessar: `https://karcher-center-altex.com.br/parceiros/wakegateway`

Esses dois passos ativam o cookie de parceiro e o grupo de pagamento correto com seguro.

> **Atenção:** Se a URL da barra de endereço começar com `checkout.`, o cookie do primeiro passo não foi capturado corretamente.

---

## Troubleshooting

---

### TRB-001 — API de cotação retorna `data: []` vazio sem erro

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Projeto Hero <> Leanwork <> Wake |
| **Data** | 2026-01-28 |
| **Categoria** | `troubleshooting` |
| **Tags** | `integração`, `audience-solutions` |

**Sintomas:** A chamada à API `/api/prestamista/quotation` retorna HTTP 200 mas com `"data": []` e `"success": true`.

**Causa raiz:** O `type_of_product` enviado não corresponde ao parceiro correto. Cada parceiro cadastrado na Hero tem um ID de produto diferente.

**Resolução:**
- Verificar com o time Hero qual o `type_of_product` correto para o parceiro em questão.
- Para a Karcher: usar `type_of_product: 5` (e credenciais `hero@karcher.com.br`).
- O usuário `wake@heroseguros.com.br` (client_id 266) usa `type_of_product: 4` para testes genéricos.

---

### TRB-002 — Erro `QUOTATION_NO_PRICING_RULES` na cotação

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Projeto Hero <> Leanwork <> Wake |
| **Data** | 2026-02-03 |
| **Categoria** | `troubleshooting` |
| **Tags** | `integração`, `audience-solutions` |

**Sintomas:** A API retorna:
```json
{
  "success": false,
  "message": "Erro na requisição externa",
  "error": {
    "code": "QUOTATION_NO_PRICING_RULES",
    "message": "The price quoted is not allowed for this product and plan."
  }
}
```

**Causa raiz:** O produto não possui regras de precificação cadastradas, ou o `type_of_product` está errado (ex: enviando `4` em vez de `5` para a Karcher).

**Resolução:**
1. Confirmar com o time Hero o ID correto do tipo de produto para o parceiro.
2. Verificar se o produto possui SKU cadastrado e associado ao parceiro na plataforma Hero.
3. Para a Karcher, após mudança de `wake@` para `hero@karcher.com.br`, o ID mudou de 4 para 5.

---

### TRB-003 — Pedidos sem seguro sendo alterados para status "Separado" indevidamente

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Projeto Hero <> Leanwork <> Wake |
| **Data** | 2026-01-27 |
| **Categoria** | `troubleshooting` |
| **Tags** | `suporte-escalação`, `audience-solutions`, `audience-cs` |

**Sintomas:** Pedidos comuns (sem seguro) na loja Karcher foram alterados automaticamente para o status "Separado" após a ativação do gatilho de integração Hero.

**Causa raiz:** O script de automação de status estava disparando para todos os pedidos pagos, sem verificar a flag de seguro contratado.

**Resolução implementada (paliativa — Wake):**
- Restringir transições de status: bloquear a mudança para "Separado" a partir de "Cancelado" e de "Enviado".

**Resolução definitiva (Leanwork):**
- Revisar o script para disparar o gatilho de "Separado" **apenas** para pedidos que contenham a flag de seguro ativo E estejam com status "Pago".

> **Lição aprendida:** Confirmar com o cliente se a loja já está em produção antes de ativar qualquer automação. Neste caso, a equipe não sabia que a Karcher já estava com pedidos reais chegando em produção.

---

### TRB-004 — `debt_amount` no bilhete incluindo erroneamente o valor do seguro

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Projeto Hero <> Leanwork <> Wake |
| **Data** | 2026-02-10 |
| **Categoria** | `troubleshooting` |
| **Tags** | `integração`, `audience-solutions` |

**Sintomas:** O bilhete do Parcela Garantida apresentava valores incorretos. O `debt_amount` enviado na proposta incluía compras + frete + valor do seguro, quando deveria ser apenas compras + frete.

**Causa raiz:** O middleware estava somando o valor do seguro ao subtotal antes de enviar à API de proposta Hero (`/api/prestamista/proposal`).

**Resolução:** O `debt_amount` deve ser o valor **original da compra + frete**, sem nenhuma adição do valor do seguro. O cálculo do percentual do seguro é feito pela própria API da Hero com base no `debt_amount` informado.

---

### TRB-005 — Alíquota configurada não aparecendo no checkout

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Projeto Hero <> Leanwork <> Wake |
| **Data** | 2025-11-05 |
| **Categoria** | `troubleshooting` |
| **Tags** | `contrato-comercial`, `audience-solutions`, `audience-cs` |

**Sintomas:** O grupo de produtos e alíquotas foram criados no painel Wake, mas o valor do seguro não aparecia no checkout.

**Causa raiz:** Configuração incompleta do grupo de produtos ou falta de associação correta dos produtos ao grupo de alíquota.

**Resolução:** Contato direto com o suporte técnico da Wake via WhatsApp para revisão das configurações. A documentação de referência é: [Alíquotas - Wake Help Center](https://atendimento.wake.tech/hc/pt-br/articles/21406631857687-Al%C3%ADquotas)

**Resolução alternativa adotada (Fase 2 — Karcher):** Em vez de usar alíquotas, adotar a estratégia de **configuração baseada em grupo de pagamentos e cartão de crédito** como solução temporária até finalizar as configurações de alíquota.

---

### TRB-006 — Script da integração para de funcionar após ativação do ETL na Wake

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Projeto Hero <> Leanwork <> Wake |
| **Data** | 2026-02-25 |
| **Categoria** | `troubleshooting` |
| **Tags** | `integração`, `audience-solutions` |

**Sintomas:** Após a ativação do ETL na Wake para a Karcher (que passou a preencher o campo `subtotal-somdesconto`), o script de integração parou de disparar: pedidos com seguro não alteravam status e não adicionavam o comentário com o link do bilhete.

**Causa raiz:** A estrutura de dados do pedido mudou com a ativação do ETL. O script estava se baseando em um campo ou estrutura que foi alterada pela nova configuração.

**Resolução (Leanwork/Proex):**
- Analisar qual campo/flag o script usava como trigger.
- Ajustar o script para reconhecer a nova estrutura de dados pós-ETL.
- Testar com pedidos com e sem seguro para validar que apenas os pedidos com seguro disparam o gatilho.

---

### TRB-007 — Pedido com PIX dispara gatilho de seguro indevidamente

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Projeto Hero <> Leanwork <> Wake |
| **Data** | 2026-03-23 |
| **Categoria** | `troubleshooting` |
| **Tags** | `integração`, `audience-solutions` |

**Sintomas:** Pedidos pagos via PIX (e potencialmente boleto) estavam disparando o gatilho do seguro, gerando bilhetes e e-mails indevidos para clientes que não contrataram o seguro.

**Causa raiz:** O script não estava validando o método de pagamento utilizado antes de disparar o gatilho. O seguro Parcela Garantida só é válido para pagamentos parcelados no cartão de crédito.

**Resolução (Leanwork):** Adicionar verificação do método de pagamento no script: o gatilho só deve ser disparado quando o ID do grupo de pagamento corresponder ao "Cartão com Seguro" (IDs configurados na Wake). Para a Karcher em produção, o ID correto é `16468` (GTW Cartão com Seguro). O grupo `15848` (GTW Cartão sem seguro) não deve gerar bilhete.

---

### TRB-008 — Pedido volta a status "Separado" após movimentação pela logística

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Projeto Hero <> Leanwork <> Wake |
| **Data** | 2026-02-10 |
| **Categoria** | `troubleshooting` |
| **Tags** | `suporte-escalação`, `audience-solutions` |

**Sintomas:** Após a logística alterar o status de um pedido de "Separado" para "Em preparação", o sistema voltava a mudar o pedido para "Separado" automaticamente.

**Causa raiz:** O script de integração não tinha controle de idempotência — disparava a mudança para "Separado" toda vez que o pedido fosse processado, sem verificar se já havia passado por esse status.

**Resolução implementada:**
1. **Wake (paliativo):** Bloquear transições de status para impedir que do status "Cancelado" ou "Enviado" se mova de volta para "Separado".
2. **Leanwork (definitivo):** Revisar o script para acionar a mudança para "Separado" apenas **uma vez**, somente quando o pedido estiver com status "Pago" pela primeira vez.

---

## How-To

---

### HTO-001 — Como configurar o onboarding de um novo lojista (Middleware + Scripts + Wake)

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Projeto Hero <> Leanwork <> Wake |
| **Data** | 2025-11-06 |
| **Categoria** | `how_to` |
| **Tags** | `onboarding`, `audience-solutions` |

**Visão geral do processo de implantação para um novo cliente:**

1. **Configuração do Middleware (Leanwork)**
   - Deploy do middleware na infraestrutura acordada
   - Configuração das credenciais do parceiro (client_id, secret, usuário)
   - Configuração do `type_of_product` correto para o parceiro

2. **Inclusão de Scripts (Leanwork/Proex)**
   - Incluir o script `guarantee_installment-hero.js` na loja Wake do cliente
   - Verificar se o script está apontando para as URLs corretas de bilhete

3. **Configuração da Loja Wake**
   - Criar grupos de pagamento: "GTW Cartão com Seguro" e "GTW Cartão" (sem seguro)
   - Configurar alíquotas (ou usar grupo de pagamentos como estratégia alternativa)
   - Mapear os IDs dos novos grupos de pagamento no ERP do cliente (ex: Tiny)
   - Liberar o cartão de crédito no checkout Storefront

4. **Ajustes de Front-end (Agência/Proex)**
   - Habilitar/desabilitar corretamente os grupos de pagamento no checkout
   - Aplicar o layout dos componentes Hero (box de seleção de seguro, labels, textos)

5. **Documentação de Onboarding**
   - A Leanwork elabora e envia por e-mail a documentação de onboarding para a agência responsável pela loja

> **Importante:** Deixar claro para a agência que após a configuração do backend, ela precisa ajustar o front-end.

---

### HTO-002 — Como adicionar as páginas "Saiba Mais" para os produtos de seguro na Wake

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Projeto Hero <> Leanwork <> Wake |
| **Data** | 2025-09-03 |
| **Categoria** | `how_to` |
| **Tags** | `integração`, `audience-solutions` |

**Contexto:** No box de seleção de seguro, existe um link "Saiba mais" que deve apontar para uma página com detalhes do produto.

**Passos:**
1. Criar as páginas de conteúdo na plataforma Wake para cada produto de seguro (Garantia Estendida e Parcela Garantida)
2. O conteúdo base deve usar o documento fornecido pela Hero Seguros (ex: "Regras e Benefícios: Garantia Estendida e Parcelamento com Seguro")
3. Ajustar o HTML e CSS da página conforme identidade visual do lojista
4. Configurações globais de CSS aplicadas refletirão tanto na URL pública da página quanto no popup exibido no checkout

**Documentação Hero de referência:** Verificar versão mais atualizada do documento de definições de produtos com a Hero.

---

### HTO-003 — Como realizar o deploy do Parcela Garantida em produção (checklist)

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Projeto Hero <> Leanwork <> Wake |
| **Data** | 2026-03-13 |
| **Categoria** | `how_to` |
| **Tags** | `integração`, `audience-solutions` |

**Checklist validado para o go-live do Parcela Garantida na Karcher:**

**Wake (Leo):**
- [ ] Remover conector de teste e vincular conector real da Vindi
- [ ] Confirmar novo ID do grupo de pagamento gerado (ex: `16468 | GTW Cartão com Seguro`)
- [ ] Repassar novos IDs de grupos de pagamento para mapeamento no ERP (Tiny) via cliente

**Agência/Proex (Yuri/Sabrina):**
- [ ] Habilitar grupos de pagamento corretos no painel da loja
- [ ] Configurar scripts nos templates corretos
- [ ] Remover mascaras indevidas nos tokens (ex: `tcs_karch_`)
- [ ] Validar que o checkout exibe corretamente a opção de seguro

**Hero Seguros:**
- [ ] Confirmar credenciais de produção do parceiro
- [ ] Validar que o parceiro está cadastrado com o `type_of_product` correto

**Processo de virada:**
1. Primeiro subir o checkout Storefront (SF) e validar funcionamento independente
2. Somente após SF estável, ativar os scripts Hero
3. Realizar pedido teste com seguro e pedido teste sem seguro
4. Confirmar recebimento de e-mail, mudança de status e link do bilhete no comentário do pedido

---

### HTO-004 — Como preparar uma loja modelo/demo para eventos e apresentações

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Projeto Hero <> Leanwork <> Wake |
| **Data** | 2025-09-05 |
| **Categoria** | `how_to` |
| **Tags** | `proposta`, `audience-solutions`, `audience-sales` |

**Contexto:** Para o Wake Summit 2025, foi necessário preparar uma loja modelo (`partnerhero.fbits.store`) para demonstração dos seguros.

**Recomendações:**
1. Usar uma URL que não apareça em buscas no Google e seja desconhecida do público (sem bloqueios técnicos, mas obscura)
2. Cadastrar produtos com preços próximos da realidade para não gerar percepção negativa do custo do seguro
3. Usar os percentuais reais das tabelas de precificação fornecidas pela Hero
4. Ajustar preços dos produtos para representar valores realistas no dia do evento
5. Gravar vídeos da jornada em formato `.mkv` e compartilhar como arquivo (o WhatsApp reduz a qualidade)
6. Para apresentação em tablet: resolução recomendada de **1024px** (4:3), testado com Google Chrome

---

### HTO-005 — Como mapear novos grupos de pagamento no ERP (Tiny) para a Karcher

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Projeto Hero <> Leanwork <> Wake |
| **Data** | 2026-02-23 |
| **Categoria** | `how_to` |
| **Tags** | `integração`, `audience-solutions` |

**Contexto:** Ao configurar novos grupos de pagamento no checkout Wake (Storefront), os IDs precisam ser mapeados no ERP do cliente para que os pedidos sejam identificados corretamente.

**Exemplo de lista de IDs para mapeamento — Karcher:**

| ID | Descrição |
|---|---|
| 1172 | Cartão |
| 1174 | Boleto |
| 5409 | Pix |
| 12644 | Link de pagamento |
| 16307 | GTW Cartão com Seguro |
| 15848 | GTW Cartão |
| 15846 | GTW Boleto |
| 15847 | GTW Pix |
| 16468 | GTW Cartão com Seguro (produção) |

**Processo:** A solicitação de mapeamento deve ser feita pelo próprio cliente (Karcher) para o ERP (Tiny), pois a agência não tem credenciais de acesso. A Wake pode acionar via ticket se necessário.

> **Atenção:** Ao trocar de conector de teste para conector real (Vindi), um novo ID de grupo é gerado e precisa ser remapeado na Tiny.

---

## Best Practices

---

### BPR-001 — Nunca fazer deploy em produção numa sexta-feira

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Projeto Hero <> Leanwork <> Wake |
| **Data** | 2026-02-20 |
| **Categoria** | `best_practices` |
| **Tags** | `integração`, `audience-solutions` |

**Recomendação:** Nunca fazer deploy em produção numa sexta-feira à tarde. Reservar janelas de deploy para segunda-feira cedo ou terça-feira pela manhã.

**Rationale:** Deploys em produção no final da semana reduzem a capacidade de resposta a incidentes — em caso de problema, a equipe estará menos disponível para resolver. Caso documentado: surgiu a oportunidade de fazer o deploy de uma configuração em produção numa sexta-feira à tarde; a decisão foi não realizar o deploy e reservar uma janela na segunda-feira cedo.

---

### BPR-002 — Confirmar o ambiente (teste vs. produção) com o cliente antes de qualquer automação

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Projeto Hero <> Leanwork <> Wake |
| **Data** | 2026-01-27 |
| **Categoria** | `best_practices` |
| **Tags** | `integração`, `audience-solutions` |

**Recomendação:** Antes de ativar qualquer automação, script ou webhook em um ambiente de cliente, confirmar explicitamente:
1. A loja já está em produção recebendo pedidos reais?
2. Qual a URL do ambiente de testes vs. produção?
3. Existe isolamento entre os ambientes (configurações separadas)?

Ao ativar integrações, sempre testar com um pedido com seguro e um sem seguro para validar o escopo do gatilho.

**Rationale:** A equipe de desenvolvimento ativou um gatilho de automação acreditando que a loja estava apenas em ambiente de teste. Na verdade, a loja Karcher já estava em produção com pedidos reais de clientes. Resultado: pedidos reais foram afetados (status alterado indevidamente para "Separado"). O custo de não confirmar o ambiente é alto — pedidos de clientes reais com status errado geram dano operacional e de reputação.

---

### BPR-003 — Estratégia de rollback: usar conector temporário para desacoplar dependências de gateway

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Projeto Hero <> Leanwork <> Wake |
| **Data** | 2026-02-04 |
| **Categoria** | `best_practices` |
| **Tags** | `suporte-escalação`, `audience-solutions` |

**Recomendação:** Sempre que uma dependência externa (gateway, ERP, etc.) bloqueie testes, avaliar se é possível desacoplar temporariamente sem comprometer a lógica de negócio. Modelo aplicado no projeto Hero: usar um "conector de teste" (sem Vindi) temporariamente, mantendo o mesmo grupo de pagamento para não precisar alterar nada no script de validação. Com isso:
- Os testes do seguro (Hero) podiam prosseguir
- A única diferença: o status de "Pago" precisava ser alterado manualmente durante os testes
- Os demais passos (status "Separado", comentário com bilhete) continuavam automáticos

Atenção: ao retornar para o conector real, um novo ID de grupo de pagamento é gerado e precisa ser remapeado no ERP.

**Rationale:** Caso concreto: a Vindi (gateway de pagamento) estava em período de alta demanda (Black Friday/pós-BF) e não podia ativar novos conectores de produção, bloqueando os testes completos da integração Hero. Sem a estratégia de desacoplamento, o projeto teria ficado parado semanas aguardando a Vindi liberar capacidade.

---

### BPR-004 — Documentar formalmente decisões técnicas tomadas durante war rooms

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Projeto Hero <> Leanwork <> Wake |
| **Data** | 2025-11-25 |
| **Categoria** | `best_practices` |
| **Tags** | `suporte-escalação`, `audience-solutions`, `audience-pam` |

**Recomendação:** Ao sair de uma war room, publicar no grupo um resumo estruturado com:
- **Decisão tomada** por cada parte
- **Responsável** por cada ação
- **Prazo** ou próximo checkpoint

Exemplo real do projeto Hero: "Lean/Wake: Adotará nova estratégia para envio das porcentagens de seguro, utilizando configurações baseadas em grupo de pagamentos e cartão de crédito. Proex: atualizações necessárias no front-end. Lean: ajustes no front e back para incluir a informação vinda da URL da policy."

**Rationale:** War rooms envolvem múltiplas equipes (no projeto Hero: Hero, Wake, Leanwork, Proex) tomando decisões rápidas sob pressão. Sem resumo estruturado por escrito, decisões verbais se perdem, gerando retrabalho, mal-entendidos e duplicação de esforço. O resumo serve como registro histórico para auditoria ou futuros projetos.

---

### BPR-005 — Separação de fases para go-live: SF primeiro, seguro depois

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Projeto Hero <> Leanwork <> Wake |
| **Data** | 2026-02-27 |
| **Categoria** | `best_practices` |
| **Tags** | `integração`, `audience-solutions` |

**Recomendação:** Adotar modelo de deploy faseado em todos os projetos de implantação de seguros em lojas já existentes — Storefront primeiro, seguro depois:
1. **Etapa 1:** deploy do novo checkout Wake Storefront (SF) — valida se o novo checkout funciona sem dependências do seguro
2. **Etapa 2:** após validação do SF, ativar os scripts e a integração Hero

**Rationale:** Modelo aplicado no deploy do Parcela Garantida para a Karcher. Benefícios:
- Se o novo checkout apresentar problema, não é possível imputar à integração Hero
- Janelas de rollback são mais claras para cada componente
- O cliente (Karcher) valida o checkout independentemente antes de acrescentar complexidade

Subir tudo de uma vez torna impossível identificar a origem do problema em caso de falha e dificulta rollback parcial.

---

### BPR-006 — Gestão de transição de PO durante o projeto

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Projeto Hero <> Leanwork <> Wake |
| **Data** | 2025-11-10 |
| **Categoria** | `best_practices` |
| **Tags** | `integração`, `audience-solutions`, `audience-pam` |

**Recomendação:** Iniciar a transição de gestores de projeto com pelo menos 2-3 semanas de antecedência, com sessões de repasse incluindo o time técnico, não apenas documentação. Estrutura de transição:
- **Repasse gradual** nas semanas anteriores à saída
- A nova PO assume acompanhando mais de perto o backlog de desenvolvimento
- Comunicado formal ao grupo com agradecimentos e clareza sobre continuidade

**Rationale:** No projeto Hero, o PO responsável saiu para uma nova oportunidade profissional durante a execução do projeto. A transição seguindo essa estrutura permitiu que o projeto ficasse sem impacto no escopo ou cronograma após a transição (nova PO Rebeca assumiu sem ruído). Documentação isolada não substitui repasse com o time técnico — contexto tácito só se transfere via conversas e participação em ritos.

---

### BPR-007 — Configurar e testar o produto de seguro para múltiplos SKUs antes do go-live

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Projeto Hero <> Leanwork <> Wake |
| **Data** | 2026-05-27 |
| **Categoria** | `best_practices` |
| **Tags** | `contrato-comercial`, `audience-solutions` |

**Recomendação:** Antes do go-live de produtos de seguro com múltiplos SKUs, garantir:
1. **Antes de iniciar testes:** definir com clareza quais SKUs estão cadastrados na Hero para o parceiro e quais produtos na loja Wake correspondem a esses SKUs
2. **Responsabilidade de cadastro:** a Wake/lojista/agência controla o painel da loja; a Hero controla a tabela de SKUs elegíveis. É necessário que ambas as listas estejam sincronizadas
3. **Lista de SKUs de teste em homologação Hero:** manter uma lista pré-cadastrada de SKUs de teste válidos para acelerar a fase de homologação de novos parceiros

SKUs de teste utilizados em homologação: `93986140`, `14287000`, `93984630`, `93984640`.

**Rationale:** Durante os testes finais da Garantia Estendida para a Karcher, o time ficou travado pois o produto teste inicial não estava cadastrado com um SKU elegível na plataforma Hero, e a equipe não tinha acesso ao painel da loja para cadastrar produtos. Sem sincronização prévia entre as duas listas (Hero ↔ loja), o go-live atrasa por bloqueios evitáveis de cadastro.

---

### BPR-008 — Uso da loja demo Hero para demonstrações comerciais e novos parceiros

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Projeto Hero <> Leanwork <> Wake |
| **Data** | 2025-09-11 |
| **Categoria** | `best_practices` |
| **Tags** | `onboarding`, `audience-pam`, `audience-sales` |

**Recomendação:** Manter a loja demo (`partnerhero.fbits.store`) atualizada com a versão mais recente da integração e com preços de produtos realistas, e usá-la como ferramenta principal de demonstração comercial para novos parceiros. Padrões de sucesso:
- Ter uma loja modelo funcional e com dados realistas acelera o ciclo de vendas
- A demonstração ao vivo de toda a jornada (seleção de GE + Parcela Garantida no checkout) é mais efetiva que slides
- Gravar vídeos da jornada para uso em eventos onde não é possível demonstrar ao vivo

**Rationale:** Após a apresentação no Wake Summit com a loja modelo `partnerhero.fbits.store`, a Hero fechou o primeiro contrato de parceiro (Karcher) usando exatamente essa demonstração. Preços realistas garantem que a relação custo-benefício do seguro seja percebida positivamente pelo cliente em potencial — preços fictícios distorcem a percepção e podem fazer o seguro parecer caro demais.
# KB Extract — Wave 7A (Vtrina, Fidelizar+, Visie, 3W)
> Extracted: 2026-06-04

---

## Grupo: Vtrina & Wake

### HTO-001 — Como registrar um lead indicado pelo parceiro via formulário

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Vtrina & Wake |
| **Data** | 2025-08-08 |
| **Categoria** | `how_to` |
| **Tags** | `lead-pipeline`, `audience-pam` |

Quando um parceiro identifica um lead para a Wake (por exemplo, após evento como o Fórum ECBR), o processo correto é registrar via formulário oficial antes de qualquer contato adicional, evitando perda de atribuição.

**Formulário de indicação de leads:**
`https://docs.google.com/forms/d/e/1FAIpQLSfcBuS74vXSfMOCUVwP5WoIjdUnc0Tip1mqS3_hQdvV3MjDpQ/viewform`

O parceiro deve preencher o formulário com dados do prospect e aguardar contato do time comercial Wake para dar sequência. Após o registro, a origem fica atribuída corretamente ao parceiro no CRM (Salesforce).

---

### FAQ-001 — Vtrina tem integração com Winthor (Totvs)?

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Vtrina & Wake |
| **Data** | 2025-08-14 |
| **Categoria** | `faq` |
| **Tags** | `integração`, `audience-solutions` |

**Pergunta:** A Vtrina tem integração com Winthor (ERP da Totvs)?

**Resposta:** Sim, a Vtrina tem integração com Winthor.

---

### FAQ-002 — Vtrina integra com Alpha7, Procfit, Vetor e Linx?

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Vtrina & Wake |
| **Data** | 2025-07-17 |
| **Categoria** | `faq` |
| **Tags** | `integração`, `audience-solutions` |

**Pergunta:** A Vtrina está integrada com Alpha7, Procfit, Vetor e Linx (Big/Itech)?

**Resposta:** Linx: sim. Alpha7 e Procfit: podem ser viabilizados. Vetor: desconhecido pelo time Vtrina naquele momento — necessita avaliação.

---

### FAQ-003 — Vtrina consegue atender clientes com multiCD envolvendo Correios Log+?

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Vtrina & Wake |
| **Data** | 2025-06-25 |
| **Categoria** | `faq` |
| **Tags** | `integração`, `audience-solutions` |

**Contexto:** Cliente (Artegift) queria configurar multiCD com um dos CDs sendo externo (Correios Log+). O ERP era Bling, que não suporta esse modelo nativamente sem middleware. O cliente cogitou usar dois ERPs (Bling + Tiny), mas ainda assim seria necessário um middleware.

**Resposta da Vtrina:** Acreditam conseguir atender o cenário, mas precisam de uma qualificação mais detalhada da operação dos Correios, por ser algo relativamente novo no mercado. A indicação é agendar uma reunião de qualificação antes de confirmar escopo.

---

### TRB-001 — Parceiro menciona concorrente como "principal integradora" — como lidar

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Vtrina & Wake |
| **Data** | 2025-08-29 |
| **Categoria** | `troubleshooting` |
| **Tags** | `ecossistema-parceiros`, `audience-pam` |

**Sintomas:** A Vtrina identificou que em dois eventos/reuniões a Wake citou a Anymarket como "principal parceira de integração", o que gerou questionamentos internos e sensação de preterimento da Vtrina no ecossistema de integrações.

**Causa raiz:** Comunicação genérica da Wake em eventos não diferenciou os blueprints por faixa de mercado e segmento, dando a impressão de exclusividade quando o desenho estratégico é multi-parceiro.

**Resolução:** Resposta da Wake: a relação com a Anymarket tem um cunho diferente de sustentação de negócios. A estratégia da Wake é construir blueprints por faixa de mercado (SMB, Enterprise) e por segmentos (Moda, CPG etc.), com cada parceiro aplicando sua especialidade. A intenção é que a Vtrina tenha cada vez mais espaço como integrador. O tema foi endereçado internamente. Orientação: ao falar com parceiros integradores, reforçar que o ecossistema Wake é multi-parceiro e que não há exclusividade formal. Cada integrador atua em suas especialidades e faixas de mercado.

---

### BPR-001 — Fluxo de atendimento a parceiro via CS durante negociação

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Vtrina & Wake |
| **Data** | 2025-08-04 |
| **Categoria** | `best_practices` |
| **Tags** | `suporte-escalação`, `audience-pam`, `audience-cs` |

**Recomendação:** Quando um lead indicado pelo parceiro está em processo de proposta e o contato direto não responde, o parceiro deve acionar a CS Wake responsável pela conta para fazer o follow-up, em vez de tentar acesso direto.

**Rationale:** Acionar a CS evita que o parceiro apareça como "cotovelo" para o cliente e mantém o relacionamento gerenciado pela CS, que já tem contexto e calendário com a conta. Caso documentado: Rafaela (Vtrina) tentou contato com o ArteGift sem sucesso e foi orientada pela Wake a usar a CS Nina para intermediar — a CS já estava em contato com o cliente.

---

### HTO-002 — Como funciona a presença do parceiro no escritório Wake (Partner Day)

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Vtrina & Wake |
| **Data** | 2025-08-13 |
| **Categoria** | `how_to` |
| **Tags** | `marketing-eventos`, `audience-pam` |

O **Dia do Parceiro** da Wake ocorre toda **quarta-feira** no escritório da LWSA (Rua Itapaiúna, 2434 – São Paulo/SP). O evento foi alterado de quintas para quartas a partir de 27/08/2025.

**Objetivo:** Aproximar a marca do parceiro dos times de CSM (atendimento e rentabilização de base) e time comercial Wake. O parceiro pode trabalhar do escritório o dia todo.

**Para parceiros de fora de SP:** É possível participar online. Basta confirmar presença com a equipe de parcerias informando e-mails para envio do invite.

---

### FAQ-004 — Preço padrão de integração da Vtrina para novas conexões

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Vtrina & Wake |
| **Data** | 2025-09-22 |
| **Categoria** | `faq` |
| **Tags** | `integração`, `audience-solutions`, `audience-sales` |

**Pergunta:** Qual é o preço padrão da Vtrina para uma nova integração?

**Resposta:** O preço de partida padrão é **R$ 14.300,00**. Para uma integração mais complexa (ex.: integração de cashback com módulo de pedidos e checkout), o valor informado foi de **R$ 14.700,00** com prazo estimado de **80 dias corridos** de implantação. Esses valores podem sofrer alterações conforme análise de APIs e complexidades durante o desenvolvimento.

---

### BPR-002 — Uso do NDA no fluxo de integração com parceiro externo

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Vtrina & Wake |
| **Data** | 2025-09-11 |
| **Categoria** | `best_practices` |
| **Tags** | `contrato-comercial`, `audience-pam` |

**Recomendação:** Antes de avançar no desenvolvimento de uma integração envolvendo um terceiro parceiro (ex.: plataforma de cashback), seguir o fluxo:
1. Qualificação inicial do escopo pelo time de parcerias Wake
2. Cotação prévia do integrador com base no escopo macro
3. Envio do NDA pela Wake para assinatura do integrador (após alinhamento inicial e antes de compartilhamento de documentação técnica sensível)
4. Avanço para detalhamento técnico e proposta formal

**Rationale:** Compartilhar documentação técnica sensível ou roadmap detalhado sem NDA assinado expõe a Wake a vazamento de informações estratégicas, especialmente em integrações onde o parceiro pode estar conversando com concorrentes. O NDA antes do detalhamento técnico é controle padrão para proteção de propriedade intelectual.

---

## Grupo: Fidelizar + e Wake

### HTO-003 — Como funciona a API de alíquota (cashback) na Wake

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Fidelizar + e Wake |
| **Data** | 2026-02-18 |
| **Categoria** | `how_to` |
| **Tags** | `integração`, `audience-solutions` |

A API de alíquota da Wake é acionada durante o checkout e recebe um payload com dados do carrinho (ID, produtos com `ProdutoVarianteId`, `Quantidade`, preços, `CentroDistribuicaoId`) e dados do usuário (CPF/CNPJ, e-mail).

**Payload de exemplo (campos mapeados):**
```json
{
  "Id": "255687d4-9cf9-442c-b1c0-20edfc36e183",
  "ValorTotal": 282,
  "Produtos": [
    {"ProdutoVarianteId": 266646, "Quantidade": 1, "PrecoPor": 12, ...}
  ],
  "Usuario": {"Cpf": "...", "Email": "..."}
}
```

**Resposta esperada:** Array com ajustes por produto:
```json
[{"ProdutoVarianteId": 266646, "Valor": 0.23, "Nome": "Cashback", "CentroDistribuicaoId": 25}]
```

**Timeout:** A documentação apresenta valores divergentes — 1s (central de atendimento) e 2s (readme). O time de integração Wake confirmou que em caso de dúvida, deve-se abrir ticket para esclarecimento. Valor seguro a adotar: **1s** (mais restritivo).

**Links de referência:**
- Vínculo de cupom a promoção: `https://wakecommerce.readme.io/docs/vincular-cupom-a-uma-promocao`
- Cadastro de alíquota via API: `https://wakecommerce.readme.io/docs/cadastro-de-aliquota-usando-api`
- Conceito de promoções (lojista): `https://atendimento.wake.tech/hc/pt-br/articles/21406597077399`

---

### FAQ-005 — Como correlacionar o retorno da API de alíquota com um pedido consolidado?

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Fidelizar + e Wake |
| **Data** | 2026-02-12 |
| **Categoria** | `faq` |
| **Tags** | `integração`, `audience-solutions` |

**Problema:** O corpo recebido na API de alíquota não possui identificador que correlacione diretamente ao pedido consolidado (o `carrinho-id` não serve para busca posterior de pedido).

**Solução confirmada pelo time Wake:** O caminho correto é usar **webhook**. O parceiro deve ser notificado de todos os pedidos via webhook e, ao receber cada notificação, fazer um `GET /Pedido/{id}` para verificar se o pedido contém ajuste de alíquota (indicando que a integração atuou). Se sim, processa; se não, ignora.

Esse foi o raciocínio que a equipe da Fidelizar+ havia levantado ("voltinha maior") e que o time Wake confirmou como o caminho correto.

---

### FAQ-006 — Qual é o limite de cupons vinculados a uma promoção na Wake?

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Fidelizar + e Wake |
| **Data** | 2026-02-03 |
| **Categoria** | `faq` |
| **Tags** | `integração`, `audience-solutions` |

**Pergunta:** Qual é o limite de cupons que podem ser vinculados a uma promoção?

**Resposta:**
- **Via painel admin:** O gerador de cupons permite no máximo **999 variações**.
- **Via API:** O limite não é tão rígido — é possível criar/vincular até **5.000 cupons**. Porém, acima de certo número pode haver impacto de performance na tela de detalhes da promoção no admin (ao carregar os cupons vinculados). Esse valor pode variar de acordo com a forma de manutenção pelo lojista.
- **Recomendação prática:** Trabalhar com o limite de **5.000 via API**, com atenção à performance do painel.

---

### TRB-002 — Usuário de loja teste desativado após primeiro acesso

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Fidelizar + e Wake |
| **Data** | 2026-02-05 |
| **Categoria** | `troubleshooting` |
| **Tags** | `onboarding`, `audience-solutions`, `audience-cs` |

**Sintomas:** Parceiro criou acesso à loja de teste e conseguiu entrar no primeiro login, mas ao tentar acessar novamente recebeu a mensagem "Usuário desativado!".

**Causa raiz:** O ambiente de homologação/teste pode desativar usuários automaticamente em certas condições (inatividade ou expiração de prazo de trial).

**Resolução:** Acionar o time de onboarding Wake para reativar o usuário. No caso documentado, a reativação foi feita pelo time de onboarding e o acesso foi restabelecido em minutos. Observação: o ambiente de homologação normalmente só é liberado após a finalização do contrato, mas pode ser disponibilizado antes como exceção mediante aprovação da coordenação de parcerias.

---

### FAQ-007 — A Wake tem webhook para carrinho abandonado?

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Fidelizar + e Wake |
| **Data** | 2026-04-27 |
| **Categoria** | `faq` |
| **Tags** | `integração`, `audience-solutions` |

**Pergunta:** Existe alguma forma de receber via webhook notificação de carrinho abandonado?

**Resposta (confirmada pelo time Wake):** Não. Hoje a Wake disponibiliza apenas o relatório interno de abandonos na plataforma, desde que o disparo de carrinho abandonado esteja ativo pelo admin. Não há webhook de saída para carrinho abandonado disponível no momento.

---

### HTO-004 — Como o parceiro pode solicitar ambiente de testes (homologação) antes de assinar contrato

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Fidelizar + e Wake |
| **Data** | 2026-02-03 |
| **Categoria** | `how_to` |
| **Tags** | `homologação`, `audience-pam`, `audience-solutions` |

O ambiente de homologação (loja de teste) normalmente é liberado após a finalização do contrato de parceria. No entanto, é possível abrir uma **exceção** para liberar antes, especialmente quando há necessidade técnica urgente (ex.: desenvolvimento de integração em andamento).

**Processo:**
1. O parceiro solicita a loja de teste via grupo de WhatsApp ou ao Account Manager
2. O AM valida a exceção com a coordenação de parcerias
3. O time de onboarding cria o ambiente — pode levar algumas horas para replicação
4. As instruções de acesso são enviadas por e-mail para o endereço informado pelo parceiro

**Dados necessários:** Nome da empresa e e-mail para o usuário administrador.

---

### BPR-003 — Limite de caracteres na descrição do parceiro no site Wake

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Fidelizar + e Wake |
| **Data** | 2026-03-27 |
| **Categoria** | `best_practices` |
| **Tags** | `onboarding`, `audience-pam` |

**Recomendação:** Ao criar o perfil do parceiro na página de parceiros do site Wake, respeitar os seguintes limites e processos:
- **Descrição:** máximo de 600 caracteres
- **Especialidades:** máximo de 30 caracteres por campo
- **Badge de parceiro Wake:** enviada pela Wake por e-mail (não disponível para download direto); a arte final é criada pelo time Wake após recebimento do media kit da marca do parceiro
- **Prazo para publicação do banner/pop-up:** normalmente 1 a 2 semanas após aprovação do material
- O parceiro deve enviar seu media kit (logo, cores, etc.) para o time de onboarding. A Wake produz a arte final seguindo os padrões de qualidade da marca

**Rationale:** Os limites garantem consistência visual e legibilidade no diretório de parceiros, evitando descrições truncadas ou inconsistências entre cards. Conhecer os limites antecipadamente permite ao parceiro preparar copy adequado e evita retrabalho.

---

### HTO-005 — Processo de onboarding de novo parceiro Wake

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Fidelizar + e Wake |
| **Data** | 2026-03-12 |
| **Categoria** | `how_to` |
| **Tags** | `onboarding`, `audience-pam` |

Após a assinatura do contrato de parceria, o processo de onboarding segue as etapas abaixo:

1. **Kickoff:** Reunião de apresentação com o time de parcerias Wake para alinhamento de próximos passos (material enviado: "Kickoff Onboarding")
2. **Reunião de marketing:** Apresentação das ações conjuntas disponíveis (webinar, case, post de blog, pop-up no admin Wake)
3. **Atualização no site Wake:** Envio de descrição e media kit pelo parceiro → Wake produz a arte e publica
4. **Certificação Wake:** Inscrição no treinamento de certificação (link fornecido pelo AM)
5. **Acesso à loja de teste (homologação):** Liberado mediante solicitação após contrato assinado
6. **Pop-up no admin Wake:** Agendado em janela disponível (~28 dias após onboarding); parceiro fornece texto e arte nas dimensões especificadas até a data definida

**Contatos de onboarding Wake:**
- Account Managers: Ana Vicente (agências), Gabriel Cotrim (checkout/meios de pagamento/frete)
- Coordenação: Rafael Correia (AWS Alliance/Tech Partnerships Lead), Letícia Cena (Partner Development)
- Suporte técnico: suporte@wake.tech | `https://wakecommerce.readme.io/docs/canais-de-atendimento-para-parceiros`

---

## Grupo: Visie & Wake

### FAQ-008 — Qual é a tabela de preços da Wake para clientes (B2C)?

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Visie & Wake |
| **Data** | 2026-03-17 |
| **Categoria** | `faq` |
| **Tags** | `contrato-comercial`, `audience-sales`, `audience-pam` |

**Contexto:** Agência parceira precisava de referência de preço para apresentar ao cliente enterprise antes de fechar o processo de parceria.

**Tabela de preços Wake (referência março/2026 — B2C):**

| Cenário | Faturamento/mês | Pedidos/mês | Plano | Fixo | Taxa de Recorrência |
|---|---|---|---|---|---|
| A | ~R$ 300 mil | ~500 | Start | R$ 2.900 | 2,30% TR |
| B | R$ 800k – R$ 1,2M | 2.000–3.000 | Business | R$ 10.000 | 1,00% TR |

**Observação:** Esses valores são referência para B2C. Para B2B ou enterprise, consultar o AM ou especialista comercial Wake. A precificação passou por atualização em janeiro/2026 (comunicado enviado por e-mail pela Wake).

---

### HTO-006 — Como registrar leads de clientes do parceiro no Salesforce (SF) com atribuição correta

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Visie & Wake |
| **Data** | 2026-03-11 |
| **Categoria** | `how_to` |
| **Tags** | `integração`, `audience-pam` |

Para que leads indicados pelo parceiro não percam a atribuição (e não gerem "carrossel" no CRM), o parceiro deve informar os dados do lead para o time Wake antes de qualquer contato via formulários ou eventos.

**Dados necessários para registro no SF:**
- Nome
- Cargo
- Telefone
- E-mail
- Empresa

**Processo:**
1. Parceiro identifica prospect com interesse em Wake
2. Passa os dados ao AM Wake para registro com a agência como origem
3. Se o lead preencher o formulário do site Wake, a origem já estará corretamente atribuída
4. Se o lead aparecer em evento (ex.: Fórum ECBR) sem estar registrado, a origem será "Marketing" e não a agência — perda de atribuição

**Atenção:** Leads com NDA assinado — verificar com o cliente se há permissão para compartilhar dados antes de registrar no CRM.

---

### HTO-007 — Como funciona a Certificação Wake para parceiros agência

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Visie & Wake |
| **Data** | 2026-04-13 |
| **Categoria** | `how_to` |
| **Tags** | `programa-parceiros`, `audience-pam` |

A Certificação Wake é um treinamento obrigatório para parceiros agência. Informações relevantes:

- **Acesso:** Buscar pela plataforma "WAKE LEARNING" (link enviado por e-mail após liberação de acesso)
- **Modalidades:** Certificação Completa (inclui parte comercial e de plataforma). Para perfis exclusivamente comerciais (não técnicos), a certificação completa ainda é recomendada pois tem camada operacional acessível
- **Prazo:** Normalmente 30 dias a partir da liberação do acesso
- **Inscrição de novos colaboradores:** Usar o formulário `https://docs.google.com/forms/d/e/1FAIpQLSeLjm_bZVT1WQ128YwywHPti5Gbs3OEaQmCmOM43AFpBp_TMw/viewform`
- **Novas turmas:** Liberadas toda segunda-feira da semana seguinte à inscrição
- **Problemas de acesso:** Erros na plataforma (ex.: certificado em andamento incorreto) devem ser reportados ao time de onboarding Wake para acionamento do time de Education

---

### BPR-004 — Sessões semanais de Co-Sell e Co-Build com parceiros Wake

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Visie & Wake |
| **Data** | 2026-05-07 |
| **Categoria** | `best_practices` |
| **Tags** | `lead-pipeline`, `audience-pam`, `audience-sales` |

**Recomendação:** Parceiros agência devem participar das duas sessões semanais oferecidas pela Wake às sextas-feiras:

**Co-Sell Acceleration (9h–10h)** — foco em pipeline qualificado:
- Discussão de leads e oportunidades reais
- Construção de abordagens assertivas de abordagem e descoberta
- Role play de abordagem inicial
- Revisão de objeções comuns
- Refinamento da conexão entre dor do cliente e proposta de valor Wake + parceiro

**Co-Build Sessions (10h–11h)** — foco em arquitetura e soluções:
- Exploração de funcionalidades, integrações e arquiteturas possíveis
- Co-criação de propostas para cenários reais de pipeline
- Troca de conhecimento técnico aplicado ao contexto comercial

Como participar: solicitar inclusão no invite ao AM Wake informando e-mails dos participantes.

**Rationale:** Ritual semanal entre Wake e parceiros materializa a colaboração contínua e converte o relacionamento em aprendizado prático compartilhado — separar pipeline (Co-Sell) e arquitetura (Co-Build) permite que pessoas certas participem do tema certo, sem misturar discussões e perder foco.

---

### HTO-008 — Como a Wake suporta o parceiro agência na negociação de cliente que considera Shopify

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Visie & Wake |
| **Data** | 2026-05-19 |
| **Categoria** | `how_to` |
| **Tags** | `competitivo`, `audience-sales`, `audience-pam` |

Quando um lead está inclinado para a Shopify, mas o parceiro quer defender a Wake, o processo recomendado é:

1. Parceiro avisa o AM Wake sobre a situação
2. Wake aciona o **gerente comercial** para apoiar a abordagem com argumentos de diferenciação
3. Parceiro recebe orientação sobre como abrir o assunto com o cliente ("entender melhor os diferenciais e como abrir para a Wake")
4. Se o cliente seguir com Shopify mesmo assim, o parceiro informa ao AM e o lead é desqualificado no CRM

**Dados do lead a informar:** nome, cargo, empresa, contato — verificar NDA antes de compartilhar.

---

### FAQ-009 — Como funciona o kit migração de storefront na Wake?

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Visie & Wake |
| **Data** | 2026-05-18 |
| **Categoria** | `faq` |
| **Tags** | `integração`, `audience-solutions`, `audience-pam` |

**Pergunta:** Como funciona o kit de migração de storefront e qual plano deve o parceiro agência aderir?

**Resposta:** O tema envolve análise dos planos disponíveis (com foco em migração de storefront e redesign de checkout) e o valor tabelado mais adequado para o perfil do parceiro e dos clientes que irão migrar. A Wake disponibiliza programas:

- **StoreFront 2.0:** Escopos Essencial, Básico, Intermediário e Personalizado — modelos pré-definidos exclusivos para o parceiro (sem concorrência direta)
- **Redesign de Checkout:** Escopos Essencial e Personalizado

Os preços e escopos foram criados em parceria com agências para agilizar o processo de vendas. A adesão é opcional; parceiros não aderentes continuam recebendo indicações padrão.

Para aprofundamento: agendar reunião com o AM Wake (link de agenda disponível via AM).

---

## Grupo: Wake <> 3W

### HTO-009 — Fluxo de distribuição de leads de canais para agências parceiras

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Wake <> 3W |
| **Data** | 2025-08-04 |
| **Categoria** | `how_to` |
| **Tags** | `lead-pipeline`, `audience-pam`, `audience-sales` |

Para leads classificados como "Commerce – Canais", a agência parceira deve atender o lead de **ponta a ponta**, sendo responsável por:

1. **Apresentação da plataforma Wake** ao lead (não há executivo Wake atuando diretamente)
2. **Precificação da plataforma** (usar tabela de preços Wake atualizada — consultar AM antes do envio da proposta)
3. **Proposta comercial do projeto** (escopo de implementação pela agência)

**Dados normalmente entregues pela Wake no repasse do lead:**
- URL do site atual e plataforma em uso
- ERP e volume de faturamento
- Perfil do negócio e histórico de contato
- Contato do decisor (nome, telefone, e-mail)
- Pontos de dor ou necessidades identificados pelo SDR Wake

**Importante:** A precificação Wake sofreu atualização em janeiro/2026 — sempre consultar o AM antes de enviar proposta ao lead.

---

### FAQ-010 — O que é o Wake Score e como funciona?

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Wake <> 3W |
| **Data** | 2026-01-23 |
| **Categoria** | `faq` |
| **Tags** | `programa-parceiros`, `audience-pam` |

O **Wake Score** é uma métrica utilizada no ecossistema de parcerias Wake. Em janeiro/2026 houve uma atualização da metodologia, apresentada em webinar específico para parceiros ("Desvendando o Novo Wake Score").

Para obter mais detalhes sobre o funcionamento do novo Wake Score, o parceiro deve acessar o canal de parcerias ou solicitar ao AM a gravação do webinar.

---

### BPR-005 — Registrar leads no SF antes de eventos para garantir atribuição ao parceiro

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Wake <> 3W |
| **Data** | 2025-07-23 |
| **Categoria** | `best_practices` |
| **Tags** | `integração`, `audience-pam` |

**Recomendação:** Antes de qualquer evento (especialmente Fórum ECBR), o parceiro deve informar ao AM Wake os dados dos clientes/prospects que pretende levar ao stand para que sejam previamente registrados no SF com a agência como origem.

**Rationale:** Se um cliente do parceiro bater no stand da Wake sem estar previamente registrado no SF com o parceiro como origem, a fonte de origem será atribuída ao Marketing Wake — e o parceiro não receberá crédito na campanha de vendas. Sem o registro prévio, leads "quentes" gerados pelo investimento do parceiro são perdidos para a atribuição correta.

---

### HTO-010 — Como funciona a Campanha de Vendas Wake para agências

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Wake <> 3W |
| **Data** | 2025-08-11 |
| **Categoria** | `how_to` |
| **Tags** | `marketing-eventos`, `audience-pam` |

A Wake realiza campanhas de vendas periódicas para agências parceiras, com incentivos por contrato fechado:

- **Valor fixo por contrato:** R$ 500,00 por contrato fechado (independente do GMV)
- **Prêmio adicional:** Concorre a R$ 50.000 (critérios disponíveis no material da campanha)
- **Pré-requisito:** Inscrição na campanha + certificação Wake ativa

**Processo de inscrição:** Link disponível dentro do documento da campanha enviado por e-mail pelo AM. Após inscrição, o parceiro deve manter os leads registrados corretamente no SF para que as oportunidades sejam contabilizadas.

---

### TRB-003 — Acesso a live/evento com Google Meet via e-mail corporativo

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Wake <> 3W |
| **Data** | 2025-11-18 |
| **Categoria** | `troubleshooting` |
| **Tags** | `integração`, `audience-pam`, `audience-cs` |

**Sintomas:** Parceiro recebeu convite para live da Wake no e-mail corporativo, mas ao tentar acessar o Google Meet, o sistema não permitia entrada com o e-mail corporativo — apenas com Gmail.

**Causa raiz:** A restrição acontece quando o convite é configurado para aceitar apenas contas Google pessoais (Gmail), não contas de Google Workspace com domínio próprio.

**Resolução:** Reportar o problema ao AM Wake para que seja verificado internamente e corrigido. No caso documentado, o time Wake ficou de analisar a configuração do evento. Prevenção: ao se inscrever em eventos Wake, utilizar o e-mail do Google Workspace ou verificar com o AM se o Meet aceita o domínio corporativo antes do evento.

---

### BPR-006 — Briefing de lead de canal com análise BANT para agências parceiras

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Wake <> 3W |
| **Data** | 2026-05-08 |
| **Categoria** | `best_practices` |
| **Tags** | `lead-pipeline`, `audience-pam`, `audience-sales` |

**Recomendação:** Ao repassar leads para agências parceiras, entregar briefing estruturado com análise BANT (Budget, Authority, Need, Timeline) e insights estratégicos gerados a partir da reunião do SDR com o prospect. Estrutura (exemplo aplicado em Ello Moda Fitness):
- **Contexto do negócio:** modelo, canais, fase atual
- **Desafio técnico:** ERP atual, pontos de dor, integrações necessárias
- **Perfil do decisor:** cargo, estilo de comunicação, como abordar
- **O que a agência deve fazer:** ações recomendadas para avançar o deal
- **BANT resumido:** budget (se informado), authority (quem decide), need (dores principais), timeline (urgência)
- **Observações estratégicas:** gatilhos emocionais, posicionamento recomendado

Orientação para agências: usar o briefing para preparar a abordagem antes do primeiro contato. Foco em redução de risco e execução acompanhada (não apenas features), especialmente quando o cliente tem histórico de frustração com fornecedor anterior.

**Rationale:** Briefing BANT antecipado evita que a agência refaça qualificação do zero e permite preparar abordagem focada nas dores reais do prospect. Sem esse nível de detalhe, primeiras reuniões viram "discovery duplicado" e perdem o momentum do interesse inicial gerado pelo SDR Wake.
