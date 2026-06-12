# KB Extract — Parceiros Tech (Omnichat, FG, Leanwork, Eficaz)
> Extracted: 2026-06-04

---

## Grupo: Wake & Omnichat
> JID: `120363418663841968@g.us`

---

### FAQ-1 — Carrinho Abandonado: API Wake Commerce vs Wake Experience

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Wake & Omnichat |
| **JID** | `120363418663841968@g.us` |
| **Data** | 2025-06-23 |
| **Categoria** | `faq` |
| **Tags** | `carrinho-abandonado`, `wake-commerce`, `wake-experience`, `api` |

**Pergunta:** A Omnichat pode recuperar carrinho abandonado via APIs do Wake Commerce?

**Resposta:** Nas APIs do Wake Commerce só existem APIs gerais de relatório — não há uma API específica de carrinho abandonado. O foco de evolução dessa feature está no produto Wake Experience. A integração de carrinho abandonado via Wake Commerce exigiria uma abordagem alternativa (pixel/tag/script no checkout). Clientes que possuem Wake Experience têm acesso a essa funcionalidade de forma nativa.

**Complemento técnico (2025-12-03):** O PM da Omnichat (Michel) esclareceu que o modelo atual usado na VTEX depende de capturar o `orderformId` no checkout e os dados do cliente (telefone). Se o cliente fecha o pedido posteriormente, o carrinho é anulado para não disparar campanha. Para funcionar na Wake seria necessário um checkout padronizado que permita capturar essas informações via script. O novo checkout da Wake pode viabilizar essa integração se tiver padronização de layers — ponto a ser validado com o time técnico.

---

### FAQ-2 — Checkout novo da Wake e impacto na integração Omnichat

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Wake & Omnichat |
| **JID** | `120363418663841968@g.us` |
| **Data** | 2025-12-02 |
| **Categoria** | `faq` |
| **Tags** | `checkout`, `api-storefront`, `omnichat`, `integração` |

**Contexto:** A Wake lançou um novo checkout. A Omnichat perguntou sobre o impacto na integração existente e se a recuperação de carrinho seria possível.

**Resposta da Wake:** O novo checkout utiliza a API Storefront. A busca de produtos pelo painel da Omnichat para montar carrinhos via SAC já usa essas novas APIs. A questão do carrinho abandonado não tem relação direta com o novo checkout — seria uma entrega separada. Quanto à URL de checkout: quando o cliente (Norte Refrigeração) migrar para o novo checkout, a URL muda, mas é apenas uma configuração no ambiente do cliente na Omnichat.

**Possibilidade via tag/pixel:** A abordagem de script no checkout (similar à VTEX) é viável se o novo checkout da Wake seguir um padrão de layers. Sem padronização, não é possível criar um script único agnóstico para captura de dados.

---

### TROUBLESHOOTING-1 — Problema de integração: cliente UncleK

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Wake & Omnichat |
| **JID** | `120363418663841968@g.us` |
| **Data** | 2025-10-24 |
| **Categoria** | `troubleshooting` |
| **Tags** | `integração`, `acionamento`, `suporte`, `uncleck` |

**Situação:** Cliente UncleK acionou a Wake informando que o time técnico da Omnichat (Miguel) precisava validar algumas questões na integração.

**Resolução:** A parceira de negócios da Omnichat (Ana Aires) foi informada e assumiu o contato. Recomendou centralizar as comunicações técnicas com ela e/ou com o Kleber (comercial) para direcionar ao time correto. Lição: a Omnichat estava em fase de finalização de contrato de parceria, e o canal técnico ainda não estava completamente estruturado; a resolução passou pela escalada comercial.

---

### HOW_TO-1 — Como indicar oportunidade de lead Wake para a Omnichat

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Wake & Omnichat |
| **JID** | `120363418663841968@g.us` |
| **Data** | 2025-11-03 |
| **Categoria** | `how_to` |
| **Tags** | `indicação`, `lead`, `co-sell`, `omnichat`, `processo` |

**Processo para indicar um lead da base Wake para a Omnichat:**

1. Identificar oportunidade na base (CS ou Partner Account identifica cliente com potencial para solução de atendimento/WhatsApp).
2. Acionar a Ana Aires (ana.aires@omni.chat) ou o Kleber (kleber.albertini@omni.chat) diretamente, com o contexto do cliente.
3. Coletar os seguintes dados antes de encaminhar:
   - Site do cliente
   - Volume de mensagens ativas e receptivas (média mensal de conversas no WhatsApp ou outros canais)
   - Plataforma de e-commerce
   - CRM (se houver)
   - Nome, e-mail, telefone e cargo do contato
4. A Omnichat direciona para o executivo comercial correto.
5. Registrar a indicação no Salesforce para rastreabilidade.

**Threshold mínimo de volume:** A Omnichat foca em clientes com pelo menos 1.000 conversas/mês. Volumes menores são avaliados caso a caso.

**Canal de co-sell:** Após o onboarding de parceria, a Omnichat disponibilizou canal no Slack para coordenação de oportunidades entre os times de CS e comercial.

---

### HOW_TO-2 — Processo de onboarding e KickOff da parceria Omnichat x Wake

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Wake & Omnichat |
| **JID** | `120363418663841968@g.us` |
| **Data** | 2026-01-14 |
| **Categoria** | `how_to` |
| **Tags** | `onboarding`, `kickoff`, `parceria`, `business-plan`, `slack` |

**Etapas do onboarding após assinatura de contrato:**

1. **KickOff:** Reunião conduzida pelo time de Partner Ops da Wake (Renato). Foco comercial — apresentação das mecânicas de co-sell, expectativas e responsabilidades.
2. **Business Plan:** Formulário compartilhado (Google Docs) que a parceira deve preencher com visão de negócios, metas e plano de divulgação.
3. **Canal Slack:** Wake cria canal no Slack inicialmente apenas entre times de Parcerias Wake e equipe Omnichat. Após treinamento, expande para liderança de CS.
4. **Treinamento:** Rodolfo (Omnichat) visita escritório da Wake (LWSA) para treinamento presencial com time de CS e comercial. Horário ideal: 14h (padrão do time de CS). Após 17h a adesão é menor.
5. **Corrida de Vendas:** Campanha de incentivo para o time de CS Wake indicar oportunidades à Omnichat, com regras de repasse (revenue share). Regras enviadas em PDF separado.
6. **Mídia Kit:** Omnichat deve fornecer logo (450x450px, variações preto/branco/transparente) e descrição atualizada do "Quem Somos" para publicação no site de parceiros Wake.

---

### BEST_PRACTICES-1 — Construção de case de sucesso conjunto (Norte Refrigeração)

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Wake & Omnichat |
| **JID** | `120363418663841968@g.us` |
| **Data** | 2025-11-26 |
| **Categoria** | `best_practices` |
| **Tags** | `case-sucesso`, `norte-refrigeração`, `co-marketing`, `gmv` |

**Resultado da integração Wake + Omnichat no cliente Norte Refrigeração:**
- A Omnichat consegue recuperar em média 30% dos carrinhos abandonados.
- E-commerces têm em média 70-80% de abandono de carrinho.
- O impacto em GMV é significativo — esse dado foi usado como argumento comercial com prospects durante negociação.

**Processo de produção do case:**
1. Omnichat produz o conteúdo do case internamente.
2. Cliente (Eduardo da Norte Refrigeração) aprova o conteúdo.
3. Case é enviado para aprovação da Wake: enviado para diego@wake.tech, rafael.correia@wake.tech, matheus.silva@wake.tech e leticia.cena@wake.tech.
4. Publicação alinhada com time de PMM (Mari Pellegrino) e content marketing (Matheus Coneglian) da Wake.
5. Calendário de publicação deve ser coordenado com o primeiro faturamento da parceria.

---

### BEST_PRACTICES-2 — Estrutura da campanha "Corrida de Vendas" para parceiro tech

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Wake & Omnichat |
| **JID** | `120363418663841968@g.us` |
| **Data** | 2026-02-06 |
| **Categoria** | `best_practices` |
| **Tags** | `corrida-de-vendas`, `revenue-share`, `cs`, `incentivo` |

**Modelo de ativação de CS para geração de leads para parceiro tech:**

- Wake cria canal Slack exclusivo entre Parcerias Wake, liderança de CS (Tier 1/2 e Tier 3) e equipe do parceiro.
- Fase inicial: canal restrito ao time de Parcerias e liderança para preparar a campanha sem criar expectativas antes do treinamento.
- Após treinamento presencial do parceiro, canal é expandido para o time de CS.
- Regras da corrida de vendas enviadas em PDF para aprovação do parceiro.
- A campanha define incentivos para o time de CS por indicações que gerem: agendas com clientes, leads qualificados e suporte em conversão.
- A Omnichat estava preparando uma Landing Page para automatizar o cadastro de oportunidades indicadas pelo CS Wake.

---

### FAQ-3 — Integração Omnichat com Wake Experience

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Wake & Omnichat |
| **JID** | `120363418663841968@g.us` |
| **Data** | 2026-02-06 |
| **Categoria** | `faq` |
| **Tags** | `wake-experience`, `integração`, `roadmap` |

**Pergunta:** Como está a integração da Omnichat com o Wake Experience (CRM/DXP)?

**Resposta:** No momento (fev/2026), não há integração entre Omnichat e Wake Experience. Existe uma oportunidade identificada pelo time de CS que pode abrir uma frente para expandir as possibilidades — ponto a ser discutido entre os times técnicos.

---

### HOW_TO-3 — Dia do Parceiro na LWSA (Wake Partner Day)

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Wake & Omnichat |
| **JID** | `120363418663841968@g.us` |
| **Data** | 2025-08-13 |
| **Categoria** | `how_to` |
| **Tags** | `partner-day`, `visita-presencial`, `lwsa`, `treinamento` |

**Funcionamento do Wake Partner Day:**

- Ocorre toda quarta-feira no escritório da LWSA (R. Itapaiúna, 2434 – Parque do Morumbi, SP).
- A partir de 27/08/2025, o dia oficial do parceiro foi alterado de quinta para quarta.
- Objetivo: aproximar parceiros dos times de CSM responsáveis pelo atendimento de clientes na base Wake.
- Para visitas, o parceiro deve informar antecipadamente: nome, e-mail e documento de identidade (para liberação na portaria).
- Para parceiros de fora do estado, acessar pelo portão 4 (central de atendimento, paredes pretas à direita).
- Treinamentos de parceiros (ex: Snapfy, Omnichat) são realizados nesse formato presencial, geralmente às 14h.
- Em datas comemorativas (Carnaval, por exemplo), o Partner Day pode ser suspenso ou apenas online — comunicação enviada pelo time de Parcerias com antecedência.

---

## Grupo: FG + Wake
> JID: `120363163349772208@g.us`

---

### FAQ-4 — B2C e B2B no mesmo ambiente Wake: é possível?

| Campo | Valor |
|---|---|
| **Grupo/Chat** | FG + Wake |
| **JID** | `120363163349772208@g.us` |
| **Data** | 2025-06-26 |
| **Categoria** | `faq` |
| **Tags** | `b2b`, `b2c`, `ambiente`, `loja`, `wake-commerce` |

**Pergunta:** É possível criar uma loja B2C e B2B no mesmo admin da Wake, ou é necessário contratar dois ambientes separados?

**Resposta:** É possível ter B2C e B2B no mesmo ambiente Wake. Adicionalmente, é possível configurar login com senha para clientes B2B, de modo que a tabela de preços diferenciada não fique visível para clientes B2C.

---

### TROUBLESHOOTING-2 — Leads enviados com erro pelo Salesforce (carrossel)

| Campo | Valor |
|---|---|
| **Grupo/Chat** | FG + Wake |
| **JID** | `120363163349772208@g.us` |
| **Data** | 2025-07-07 |
| **Categoria** | `troubleshooting` |
| **Tags** | `salesforce`, `leads`, `carrossel`, `erro-sistema` |

**Problema:** Leads chegando para a agência via carrossel de forma equivocada — clientes já existentes na base, erros do Salesforce enviando oportunidades que não deveriam ir para agências parceiras.

**Contexto:** Em julho/2025 ocorreram múltiplos casos (Farmácia Cruzeiro, Pássaro Marron, Fonseca Shop, A Dias, Casa Bergan, DotLiva) onde a Wake solicitou às agências que desconsiderassem os leads recebidos.

**Impacto para a agência:** A FG reportou que isso atrapalha o fluxo de trabalho porque a equipe para para conversar com o lead, acionar o vendedor interno e cadastrar no CRM antes de receber o pedido de desconsiderar.

**Resolução:** Wake estava em fase de atualizações do sistema, identificando melhorias. As desconsiderações foram garantidas como não afetando a fila/posição das agências no carrossel.

**Boa prática:** Quando a Wake pede para desconsiderar, explicitar que a posição na fila de carrossel não foi prejudicada.

---

### TROUBLESHOOTING-3 — Agência Elite não aparece no site de parceiros Wake

| Campo | Valor |
|---|---|
| **Grupo/Chat** | FG + Wake |
| **JID** | `120363163349772208@g.us` |
| **Data** | 2025-07-22 |
| **Categoria** | `troubleshooting` |
| **Tags** | `site-parceiros`, `elite`, `logo`, `visibilidade` |

**Problema:** A FG (agência certificada Elite) não aparecia na listagem de parceiros Elite no site wake.tech, e o logo exibido estava desatualizado.

**URL afetada:** `https://wake.tech/encontre-agencias-parceiras/?_sft_solucao=wake-commerce&_sft_certificacao_agencia=elite`

**Histórico:** Problema recorrente — a FG já havia reportado e corrigido anteriormente, mas voltou a desaparecer.

**Resolução:** Wake solicitou o logo atualizado em formato 450x450px, e o time de marketing foi acionado para incluir a FG na listagem Elite. A atualização foi direcionada ao time de marketing da Wake.

**Lição:** Agências Elite devem monitorar periodicamente se estão aparecendo corretamente no diretório. Ao atualizar logo, enviar em 450x450px.

---

### HOW_TO-4 — Registro de leads no Salesforce para contabilizar na campanha de parceiros

| Campo | Valor |
|---|---|
| **Grupo/Chat** | FG + Wake |
| **JID** | `120363163349772208@g.us` |
| **Data** | 2025-07-23 |
| **Categoria** | `how_to` |
| **Tags** | `salesforce`, `lead`, `campanha`, `origem-parceiro`, `ecbr` |

**Regra importante para eventos e feiras (ex: ECBR):**

Se um cliente de interesse da agência aparecer no stand da Wake durante um evento sem estar previamente registrado no Salesforce com a agência como "PARCEIRO ORIGEM", a origem do lead será contabilizada como "Marketing" — não para a agência.

**Processo correto:**
1. Antes do evento, a agência deve passar à Wake os dados dos clientes que pretende levar ou que estão no radar.
2. A Wake cadastra esses leads no Salesforce com a agência como parceiro de origem.
3. Qualquer negócio fechado com esses leads será atribuído à agência na campanha.

**Dados necessários para cadastro:** Nome, e-mail e contato do prospect.

---

### HOW_TO-5 — Alteração de e-mail de notificações de leads no Wake Community

| Campo | Valor |
|---|---|
| **Grupo/Chat** | FG + Wake |
| **JID** | `120363163349772208@g.us` |
| **Data** | 2025-07-15 |
| **Categoria** | `how_to` |
| **Tags** | `community`, `notificações`, `e-mail`, `leads`, `partner-ops` |

**Situação:** A FG solicitou alterar o e-mail que recebe notificações de novos leads para um endereço de grupo (comercial@agenciafg.com.br), para que múltiplos responsáveis recebessem cópia e agilizassem o atendimento.

**Processo:**
1. Solicitar ao time de Partner Ops da Wake (Renato) a alteração via mensagem no grupo ou WhatsApp.
2. Informar o novo e-mail e se deseja alterar apenas notificações ou também o e-mail de acesso ao Community.
3. Partner Ops realiza a alteração diretamente no sistema.

**Contato Partner Ops:** Renato (Wake) — responsável por demandas relacionadas ao Community.

---

### BEST_PRACTICES-3 — Qualificação e ICP de leads para agências parceiras

| Campo | Valor |
|---|---|
| **Grupo/Chat** | FG + Wake |
| **JID** | `120363163349772208@g.us` |
| **Data** | 2025-07-14 |
| **Categoria** | `best_practices` |
| **Tags** | `qualificação`, `icp`, `leads`, `agência`, `perfil` |

**Feedback da FG sobre qualidade dos leads:**

A FG reportou dificuldade em fechar leads com foco exclusivo em preço — clientes que buscam projetos por R$ 15-20k. A agência não consegue competir nesse perfil e questiona se esses leads estão dentro do ICP Wake ou são perfil de plataformas de entrada (Tray, etc.).

**Recomendação prática:**
- Wake deve entender o perfil comercial de cada agência para calibrar a distribuição de leads.
- Leads focados em preço baixo devem ser direcionados para agências com modelo de negócios correspondente.
- O formato ideal de lead para a FG: prospect já em avaliação de migração e em negociação, onde a agência pode entrar com proposta de projeto + plataforma.

**Aprendizado para o time de Parcerias:** Mapear o ICP de cada agência parceira via formulário (https://forms.gle/shdjyxJQUwJ26EaH9) para direcionar leads com maior assertividade.

---

### BEST_PRACTICES-4 — Gestão de pipeline de leads entre Wake e agências (planilha compartilhada)

| Campo | Valor |
|---|---|
| **Grupo/Chat** | FG + Wake |
| **JID** | `120363163349772208@g.us` |
| **Data** | 2025-06-25 |
| **Categoria** | `best_practices` |
| **Tags** | `pipeline`, `leads`, `planilha`, `status`, `agencia` |

**Problema identificado:** Acompanhar lead a lead em call ao final do mês era inviável pela correria das equipes. Tentativa de compartilhar planilha Excel bloqueada pelo sistema (usuários não-Wake não podiam receber o arquivo).

**Solução adotada:**
1. A Wake enviou lista de oportunidades com status atual no WhatsApp do grupo.
2. A agência (FG) retornou com status atualizado do lado delas dentro de 24h.
3. Como alternativa de longo prazo: Google Sheets com acesso compartilhado para ambos os times atualizarem status.

**Campos recomendados na planilha:** Nome do lead, executivo responsável (Wake e agência), status atual, próximos passos, data de atualização.

---

### FAQ-5 — Wake Education: certificação e acesso para agências

| Campo | Valor |
|---|---|
| **Grupo/Chat** | FG + Wake |
| **JID** | `120363163349772208@g.us` |
| **Data** | 2025-06-30 |
| **Categoria** | `faq` |
| **Tags** | `wake-education`, `certificação`, `wake-day`, `treinamento` |

**O que é o Wake Day / Wake Education:**
- Evento anual presencial dedicado ao ecossistema de parceiros.
- Pauta inclui: nova Certificação Wake, Wakepédia (base de artigos + fórum), lançamentos da plataforma, roadmap e treinamento técnico.
- Abertura de inscrições via Google Forms enviado pelo time de Parcerias.

**Sobre a Certificação Wake:**
- Formulário de inscrição por pessoa (não por agência).
- Requisito para elegibilidade em campanhas de vendas: certificação concluída até a data limite da campanha.
- Certificação divide-se em módulos (ex: módulo de Migração).

**Wakepédia:** Nova base de conhecimento Wake com artigos, fórum e área de novidades — recurso de consulta para agências e parceiros.

---

## Grupo: Wake & Leanwork
> JID: `120363409546904446@g.us`

---

### HOW_TO-6 — Processo de integração técnica com novo parceiro tech via Wake

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Wake & Leanwork |
| **JID** | `120363409546904446@g.us` |
| **Data** | 2025-07-02 |
| **Categoria** | `how_to` |
| **Tags** | `integração`, `parceiro-tech`, `processo`, `onboarding-técnico` |

**Fluxo para novas integrações técnicas (ex: Signifyd, Sellbie, Konduto):**

1. Wake identifica parceiro tech com interesse em integração e comunica à Leanwork (responsável pelas integrações).
2. Leanwork compartilha link de agenda (Calendly) para que o parceiro tech agende diretamente: `https://calendly.com/diego-silva-leanwork/30min`.
3. Parceiro tech envia 3 slots de disponibilidade; Leanwork valida com time técnico interno (Michel e Diego Silva).
4. Leanwork precisa de: nome das empresas para familiarização prévia, e-mails dos participantes para envio de convite.
5. Após call inicial, Leanwork avalia viabilidade e prepara orçamento.
6. Para integrações que envolvem componentes de front-end com scripts conectados à API Storefront, a Leanwork atua também no desenvolvimento do componente.

**Documentações relevantes para integrações:**
- Frete customizado: `https://wakecommerce.readme.io/docs/frete-customizado-visaogeral`
- Sellers/Marketplace: `https://wakecommerce.readme.io/docs/consultando-todos-os-sellers-da-loja`
- Canal de atendimento a parceiros: `https://wakecommerce.readme.io/docs/canais-de-atendimento-para-parceiros`

---

### HOW_TO-7 — Acesso ao Slack da Wake para parceiros tech

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Wake & Leanwork |
| **JID** | `120363409546904446@g.us` |
| **Data** | 2025-07-24 |
| **Categoria** | `how_to` |
| **Tags** | `slack`, `acesso`, `cs`, `parceiro-tech` |

**Situação:** Leanwork precisava de acesso ao canal de CS da Wake no Slack para coordenação de projetos.

**Processo:**
1. Parceiro tech solicita o convite de acesso ao canal Slack da Wake.
2. Wake (time de CS ou Parcerias) envia o convite por e-mail.
3. Requer plano pago do Slack do lado do parceiro (plano Pro ou superior).
4. Após aceite do convite, confirmar no grupo WhatsApp.

---

### TROUBLESHOOTING-4 — Integração APP Max pendente de homologação

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Wake & Leanwork |
| **JID** | `120363409546904446@g.us` |
| **Data** | 2025-07-23 |
| **Categoria** | `troubleshooting` |
| **Tags** | `homologação`, `integração`, `app-max`, `app-mobile` |

**Problema:** Integração da APP Max iniciada em setembro sem ter sido homologada até julho do ano seguinte (~10 meses).

**Ação:** Leanwork solicitou ajuda do time de Partner Ops Wake (Renato) para desbloquear o processo de homologação.

**Resolução:** Partner Ops verificou internamente e retornou com plano de ação para seguir com o caso.

**Aprendizado:** Integrações sem SLA de homologação definido podem ficar paradas por longos períodos. É importante estabelecer prazos e responsáveis no início do processo e escalar via Partner Ops quando travado.

---

### BEST_PRACTICES-5 — Lançamento de parceria com cobertura de mídia coordenada

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Wake & Leanwork |
| **JID** | `120363409546904446@g.us` |
| **Data** | 2025-08-06 |
| **Categoria** | `best_practices` |
| **Tags** | `lançamento`, `press-release`, `marketing`, `parceria`, `datas` |

**Modelo de lançamento de parceria Wake + Leanwork:**

Calendário coordenado de divulgação:
- **Imprensa:** 12/08
- **Blog:** 14/08
- **Redes Sociais:** 15/08

**Materiais produzidos:**
- Press release (nota de parceria) — aprovado por ambas as partes antes de publicar.
- Blog post — enviado pelo time de marketing Wake para aprovação da Leanwork.
- Inclusão do parceiro no site de parceiros Wake com logo correto.

**Ponto de atenção:** O logo da Leanwork no site Wake estava pequeno e precisou de ajuste para ocupar o espaço total disponível (sem perda de qualidade). Recomenda-se enviar logo em alta resolução com fundo transparente.

**Quem aprova:** A Leanwork faz revisão do conteúdo; o time de marketing Wake (Natasha) ajusta e publica conforme as datas acordadas.

---

### BEST_PRACTICES-6 — Modelo de co-sell: como apresentar orçamento de integração a prospect

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Wake & Leanwork |
| **JID** | `120363409546904446@g.us` |
| **Data** | 2025-08-07 |
| **Categoria** | `best_practices` |
| **Tags** | `co-sell`, `orçamento`, `proposta`, `integração`, `prospect` |

**Contexto:** Wake tem dois parceiros tech em negociação e quer incluir orçamento de integração Leanwork na proposta comercial.

**Fluxo recomendado:**
1. Wake comunica à Leanwork o nome dos prospects e escopo básico.
2. Leanwork compartilha link de agenda para call de escopo com o prospect.
3. Após call, Leanwork prepara orçamento de integração.
4. Para minimizar barreiras de aprovação em 2025: quanto menor o investimento inicial, mais fácil a liberação de verba. Usar como referência o modelo/blueprint de integrações anteriores (ex: Koin).

**Aprendizado:** Incluir integração no "pacotão" da proposta inicial aumenta o valor percebido e facilita o fechamento — o prospect não precisa negociar separadamente com o parceiro tech depois.

---

### FAQ-6 — Corrida de Vendas: como funciona e critérios de elegibilidade

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Wake & Leanwork |
| **JID** | `120363409546904446@g.us` |
| **Data** | 2025-08-07 |
| **Categoria** | `faq` |
| **Tags** | `corrida-de-vendas`, `parceria`, `incentivo`, `elegibilidade` |

**Pergunta:** Como funciona a corrida de vendas para parceiros tech da Wake?

**Resposta:** A corrida de vendas é uma campanha de incentivo onde o parceiro tech (ex: Leanwork) e os times da Wake trabalham em conjunto para fechar negócios. O parceiro deve:
1. Fazer adesão à campanha via formulário no prazo estipulado.
2. Obter certificação Wake até a data limite definida na campanha para ser elegível.
3. Ter o material de divulgação pronto para ser disparado junto com o lançamento.

As regras detalhadas são enviadas em PDF ("Regras Base Corrida de Vendas"). O material de divulgação da campanha é coordenado com o time de marketing Wake para coincidir com o lançamento oficial da parceria.

---

### TROUBLESHOOTING-5 — Logo do parceiro tech desatualizado ou pequeno no site Wake

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Wake & Leanwork |
| **JID** | `120363409546904446@g.us` |
| **Data** | 2025-08-13 |
| **Categoria** | `troubleshooting` |
| **Tags** | `logo`, `site-parceiros`, `marketing`, `visibilidade` |

**Problema:** O logo da Leanwork no site Wake estava pequeno e não ocupava todo o espaço disponível.

**Solução:** Solicitar ao time de marketing Wake (via Partner Ops ou grupo WhatsApp) que ajuste o tamanho do logo para ocupar todo o espaço reservado — expandindo a imagem sem perder qualidade.

**Como prevenir:** Enviar logo em alta resolução, idealmente em formato vetorial ou PNG com fundo transparente, com dimensões que permitam expansão. A Wake usa espaços padronizados no site de parceiros.

---

### HOW_TO-8 — Eventos pós-fórum: processo de inscrição de leads como convidados

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Wake & Leanwork |
| **JID** | `120363409546904446@g.us` |
| **Data** | 2025-07-28 |
| **Categoria** | `how_to` |
| **Tags** | `eventos`, `ecbr`, `leads-convidados`, `ingressos` |

**Processo para inscrever leads/prospects em eventos Wake (ex: Pós-Fórum ECBR):**

1. Parceiro solicita ingressos para leads passando: nome completo, cargo, área de responsabilidade (para confirmar que são decisores).
2. Wake verifica disponibilidade (eventos podem lotar).
3. Entradas são confirmadas por lista — no dia, parceiros precisam apresentar pulseiras/credenciais do evento principal.
4. Para credenciais do evento principal: presença física na retirada é necessária, junto com RG.
5. Confirmar com a Wake com antecedência para evitar situações constrangedoras para os convidados.

**Atenção:** Links de formulário para eventos têm prazo de encerramento por lotação — ao receber o link, confirmar inscrição imediatamente.

---

## Grupo: Eficaz <> WAKE
> JID: `120363419946666756@g.us`

---

### HOW_TO-9 — Processo de acesso ao ambiente de demonstração Wake para agências

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Eficaz <> WAKE |
| **JID** | `120363419946666756@g.us` |
| **Data** | 2025-09-16 |
| **Categoria** | `how_to` |
| **Tags** | `acesso`, `ambiente-demo`, `painel`, `agência`, `parceria` |

**Processo para criar acesso restrito ao painel Wake para agências:**

1. Agência solicita ao Partner Ops Wake (Renato) a criação de usuário.
2. Informar o e-mail no qual o acesso deve ser gerado (ex: parcerias@eficazmarketing.com).
3. Partner Ops cria o usuário mas não ativa até confirmação da agência.
4. Agência confirma ("Ok") e Partner Ops ativa o usuário e dispara o e-mail de acesso.

**Para acesso de clientes ao painel:** O cliente precisa acionar o suporte Wake diretamente. Apenas mediante pedido do cliente, os acessos de uma agência ao admin da conta do cliente são liberados. A agência deve fazer a ponte com o cliente para que ele abra o ticket de suporte.

**Acesso Loja Wake (ambiente de parceiro):** Parceiros têm acesso a um ambiente específico para demonstrações. O usuário parceiro com acesso restrito permite visualizar o painel sem permissão de edição em produção.

---

### TROUBLESHOOTING-6 — Perda de SEO orgânico após migração de plataforma

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Eficaz <> WAKE |
| **JID** | `120363419946666756@g.us` |
| **Data** | 2025-07-08 |
| **Categoria** | `troubleshooting` |
| **Tags** | `seo`, `migração`, `orgânico`, `tray`, `wake` |

**Caso:** Cliente OG Brindes migrou da Tray para Wake há 90 dias e perdeu todo o tráfego orgânico construído ao longo de mais de 10 anos de e-commerce.

**Ação imediata:** Agência acionou o CS Wake (Janaina) para entender o que ocorreu e intervir junto ao cliente.

**Aprendizado crítico:** Migrações de plataforma sem estratégia de SEO (redirecionamentos 301, migração de URLs, preservação de meta-dados, mapeamento de sitemap) podem destruir anos de construção orgânica. É fundamental incluir um plano de preservação de SEO no escopo de migração.

**Recomendação para agências:** Antes de qualquer migração para Wake, realizar auditoria de SEO da plataforma atual e mapear todos os redirecionamentos necessários. Incluir esse item como entregável no TAP (Termo de Abertura de Projeto).

---

### TROUBLESHOOTING-7 — Evento Purchase não registrando receita no Meta Ads

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Eficaz <> WAKE |
| **JID** | `120363419946666756@g.us` |
| **Data** | 2025-12-15 |
| **Categoria** | `troubleshooting` |
| **Tags** | `meta-ads`, `tagueamento`, `pixel`, `purchase-event`, `roas` |

**Problema:** Para o cliente Voce Constrói (https://www.voceconstroi.com.br/), o evento Purchase disparava no Meta mas sem o parâmetro de receita (`value`). O Meta registrava a venda com valor R$ 0,00, destruindo métricas de ROAS e otimização de campanhas.

**Histórico:** Problema iniciou após ajuste no Front feito pela agência (Auaha) em novembro. Já havia retornado a funcionar temporariamente após retagueamento, mas voltou a falhar em 06/11/2025. Ticket 258709 aberto em dezembro sem resolução.

**Status:** Em fevereiro/2026 o ticket foi reaberto (269708) com urgência — o problema persistia mesmo após seguir a documentação oficial e refazer os ajustes solicitados.

**Causa provável:** O parâmetro `value` do evento de compra estava sendo perdido após um ajuste de front-end. O suporte Wake informou que a análise estava fora do escopo padrão (sendo atendida em caráter de exceção).

**Recomendação:** Sempre após qualquer alteração no front-end do checkout, verificar se os eventos de conversão do Meta (e GA4) continuam disparando corretamente com todos os parâmetros (especialmente `value`, `currency` e `event_id`). Usar o Meta Events Manager e o Meta Pixel Helper para validação imediata.

---

### BEST_PRACTICES-7 — Role Play: treinamento de pitch comercial e técnico para agências

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Eficaz <> WAKE |
| **JID** | `120363419946666756@g.us` |
| **Data** | 2025-10-02 |
| **Categoria** | `best_practices` |
| **Tags** | `role-play`, `pitch`, `treinamento`, `agência`, `comercial` |

**Formato do Role Play Wake x Agência:**

- **Objetivo:** Avaliar e fortalecer o pitch técnico e comercial da agência para venda de projetos Wake.
- **Participantes:** Diego (Wake, papel de CEO da empresa cliente), Didier (Wake, papel de gerente de TI), equipe da agência (papel de si mesmos vendendo Wake).
- **Formato:** Simulação de venda de projeto de e-commerce + plataforma Wake para cliente fictício.
- **Roteiro:** Wake fornece um case/roteiro com contexto do cliente para guiar a simulação.
- **Duração:** 1-2 horas.
- **Tom:** Não é julgamento — é desenvolvimento. O objetivo é ajudar a agência a ser mais forte no discurso.
- **Orientações:** Enviadas em PDF com antecedência ("Orientações Role Play - Agências").

**Quando usar:** Ideal para agências em início de parceria ou que precisam calibrar o discurso técnico e comercial. Bom para identificar lacunas no conhecimento da plataforma antes de reuniões com clientes reais.

---

### BEST_PRACTICES-8 — Relatório semanal de pipeline de agência parceira

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Eficaz <> WAKE |
| **JID** | `120363419946666756@g.us` |
| **Data** | 2025-07-23 |
| **Categoria** | `best_practices` |
| **Tags** | `relatório`, `pipeline`, `forecast`, `parceria`, `kpi` |

**Modelo de relatório semanal adotado pela Eficaz:**

A Eficaz passou a enviar relatórios semanais no grupo WhatsApp e mensais por e-mail. Estrutura recomendada:

```
📊 Relatório Leads Wake – Semana [X]

🪐 Negociações em andamento
[Cliente]: [status e próximo passo]

📜 Certificações
[Nome]: [status da certificação]

🚦 Geração de oportunidades
[Comentário sobre novas indicações da Wake para agência e vice-versa]

🔍 Leads indicados
[Lead]: [status e observação]
```

**Frequência:** Semanal no WhatsApp, mensal por e-mail (relatório completo).

**Complemento — Forecast quinzenal:** Reuniões de forecast entre agência e executivo Wake para alinhamento de pipeline, estratégia e próximos passos. Agenda via Google Meet, frequência quinzenal.

---

### FAQ-7 — Programas Storefront 2.0 e Redesign de Checkout para agências parceiras

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Eficaz <> WAKE |
| **JID** | `120363419946666756@g.us` |
| **Data** | 2025-08-27 |
| **Categoria** | `faq` |
| **Tags** | `storefront`, `checkout`, `redesign`, `programa-agência`, `proposta` |

**O que são os programas StoreFront 2.0 e Redesign de Checkout:**

Programas lançados pela Wake para impulsionar agências parceiras com modelos pré-definidos e exclusivos (sem concorrência direta entre agências).

- **StoreFront 2.0:** Escopos Essencial, Básico, Intermediário e Personalizado.
- **Redesign de Checkout:** Escopos Essencial e Personalizado.

**Vantagens:**
- Preços e escopos criados em parceria com agências para agilizar o processo de vendas.
- Modelos exclusivos por agência — evita que duas agências Wake ofereçam o mesmo template para o mesmo mercado.

**Adesão:** Opcional. Agências que não aderirem continuam recebendo indicações padrão do carrossel de leads. Para aderir, responder "Ok" com interesse no pacote desejado.

---

### FAQ-8 — Leads de canais: responsabilidade da agência na apresentação da Wake

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Eficaz <> WAKE |
| **JID** | `120363419946666756@g.us` |
| **Data** | 2025-06-30 |
| **Categoria** | `faq` |
| **Tags** | `leads-canais`, `apresentação`, `precificação`, `agência`, `wake-commerce` |

**O que são "leads de canais":**
Oportunidades onde a Wake não tem executivo comercial atuando diretamente. A agência recebe o lead e é responsável por toda a condução do processo — incluindo apresentar a plataforma Wake e passar a precificação ao prospect.

**Diferença do modelo padrão:**
- No modelo padrão: Wake tem executivo atuando em paralelo à agência.
- No modelo canais: agência conduz sozinha a venda da plataforma + escopo de projeto.

**Implicação prática:** A agência precisa estar certificada e preparada para responder perguntas técnicas e comerciais sobre a Wake sem apoio imediato do time de vendas Wake.

**Exemplo de leads de canais recebidos pela Eficaz:** Casafer, Fox Racing, Campinas Diesel, Festa Color, Cariñon Mix, Gangue do Caramelo, Pampulha Enxovais, Agroforte.

---

### HOW_TO-10 — Aproximação com produto Wake Experience (CRM/DXP) para agências parceiras

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Eficaz <> WAKE |
| **JID** | `120363419946666756@g.us` |
| **Data** | 2025-07-15 |
| **Categoria** | `how_to` |
| **Tags** | `wake-experience`, `crm`, `dxp`, `cross-sell`, `agência` |

**Contexto:** A Wake identificou que agências com hunters focados em Wake Commerce poderiam aumentar volume de negócios atuando também com Wake Experience (CRM/DXP).

**Proposta para agências:**
1. Wake agenda demonstração do Wake Experience com o executivo de CRM Marketing (Alessandro).
2. Executivo apresenta a ferramenta e o perfil ideal de cliente para esse produto.
3. Agência usa o Experience como forma de abrir portas em clientes que já são Wake Commerce ou que usam outra plataforma de e-commerce.
4. Com dois executivos Wake focados exclusivamente em vender DXP, a parceria pode ser mais assertiva.

**Quando faz sentido indicar Wake Experience:**
- Cliente já é Wake Commerce e precisa de CRM/automação de marketing.
- Prospect que usa outra plataforma de e-commerce mas tem necessidade clara de CRM.
- Clientes com modelos de assinatura/recorrência (Wake Experience tem forte aderência nesse perfil).

---

### BEST_PRACTICES-9 — Feedback de leads perdidos: como usar para calibrar parceria

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Eficaz <> WAKE |
| **JID** | `120363419946666756@g.us` |
| **Data** | 2025-08-19 |
| **Categoria** | `best_practices` |
| **Tags** | `leads-perdidos`, `feedback`, `qualidade`, `maturidade`, `processo` |

**Relatório de situação de leads da Eficaz (agosto/2025):**

Diagnóstico compartilhado pela agência com a Wake:

- Leads da Wake chegam "muito frios" — querem falar com agência só depois de fechar com a plataforma.
- Consultor Wake pede para aguardar contato do cliente → negociação trava.
- No mês referência, nenhuma indicação da Wake gerou proposta.
- Prazo médio para mandar proposta aumentou de 2 semanas para quase 2 meses.
- SDR em fase de qualificação → impacto na geração de oportunidades.

**Recomendação prática para times de Parcerias:**
- Definir em qual estágio da negociação o lead é passado à agência (preferencialmente após qualificação mínima de interesse/fit).
- Criar SLA claro: agência deve abordar lead em X dias; se não houver progresso, Wake retoma contato.
- Diferenciar leads "quentes" (em fase de decisão) de "frios" (prospecção inicial) na comunicação para a agência.
- Agências que constroem relacionamento direto com clientes da base Wake (via CS) têm mais sucesso do que as que dependem exclusivamente dos leads enviados pelo carrossel.

---

*Fim do documento.*
