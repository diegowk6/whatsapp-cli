# KB Delivery — Batch 05a
**Artigos**: 253–318 do registry  
**Fontes**: kb-parceiros-tech.md (completo), kb-parceiros-wave6c.md (completo)  
**Total**: 66 artigos  
**Status**: prontos para importação  
**Data**: 2026-06

---

# KB Extract — Parceiros Tech (Omnichat, FG, Leanwork, Eficaz)
> Extracted: 2026-06-04

---

## Grupo: Wake & Omnichat
> JID: `120363418663841968@g.us`

---

### FAQ-001 — Carrinho Abandonado: API Wake Commerce vs Wake Experience

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Wake & Omnichat |
| **Data** | 2025-06-23 |
| **Categoria** | `faq` |
| **Tags** | `integração`, `audience-solutions` |

**Pergunta:** A Omnichat pode recuperar carrinho abandonado via APIs do Wake Commerce?

**Resposta:** Nas APIs do Wake Commerce só existem APIs gerais de relatório — não há uma API específica de carrinho abandonado. O foco de evolução dessa feature está no produto Wake Experience. A integração de carrinho abandonado via Wake Commerce exigiria uma abordagem alternativa (pixel/tag/script no checkout). Clientes que possuem Wake Experience têm acesso a essa funcionalidade de forma nativa.

**Complemento técnico (2025-12-03):** O PM da Omnichat (Michel) esclareceu que o modelo atual usado na VTEX depende de capturar o `orderformId` no checkout e os dados do cliente (telefone). Se o cliente fecha o pedido posteriormente, o carrinho é anulado para não disparar campanha. Para funcionar na Wake seria necessário um checkout padronizado que permita capturar essas informações via script. O novo checkout da Wake pode viabilizar essa integração se tiver padronização de layers — ponto a ser validado com o time técnico.

---

### FAQ-002 — Checkout novo da Wake e impacto na integração Omnichat

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Wake & Omnichat |
| **Data** | 2025-12-02 |
| **Categoria** | `faq` |
| **Tags** | `integração`, `audience-solutions` |

**Contexto:** A Wake lançou um novo checkout. A Omnichat perguntou sobre o impacto na integração existente e se a recuperação de carrinho seria possível.

**Resposta da Wake:** O novo checkout utiliza a API Storefront. A busca de produtos pelo painel da Omnichat para montar carrinhos via SAC já usa essas novas APIs. A questão do carrinho abandonado não tem relação direta com o novo checkout — seria uma entrega separada. Quanto à URL de checkout: quando o cliente (Norte Refrigeração) migrar para o novo checkout, a URL muda, mas é apenas uma configuração no ambiente do cliente na Omnichat.

**Possibilidade via tag/pixel:** A abordagem de script no checkout (similar à VTEX) é viável se o novo checkout da Wake seguir um padrão de layers. Sem padronização, não é possível criar um script único agnóstico para captura de dados.

---

### TRB-001 — Problema de integração: cliente UncleK

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Wake & Omnichat |
| **Data** | 2025-10-24 |
| **Categoria** | `troubleshooting` |
| **Tags** | `integração`, `audience-cs`, `audience-solutions` |

**Sintomas:** Cliente UncleK acionou a Wake informando que o time técnico da Omnichat (Miguel) precisava validar algumas questões na integração.

**Causa raiz:** A Omnichat estava em fase de finalização de contrato de parceria, e o canal técnico ainda não estava completamente estruturado; demandas técnicas chegavam sem ponto de entrada formal.

**Resolução:** A parceira de negócios da Omnichat (Ana Aires) foi informada e assumiu o contato. Recomendou centralizar as comunicações técnicas com ela e/ou com o Kleber (comercial) para direcionar ao time correto. A resolução passou pela escalada comercial enquanto o canal técnico formal não estava pronto.

---

### HTO-001 — Como indicar oportunidade de lead Wake para a Omnichat

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Wake & Omnichat |
| **Data** | 2025-11-03 |
| **Categoria** | `how_to` |
| **Tags** | `lead-pipeline`, `audience-pam`, `audience-cs` |

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

### HTO-002 — Processo de onboarding e KickOff da parceria Omnichat x Wake

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Wake & Omnichat |
| **Data** | 2026-01-14 |
| **Categoria** | `how_to` |
| **Tags** | `onboarding`, `audience-pam` |

**Etapas do onboarding após assinatura de contrato:**

1. **KickOff:** Reunião conduzida pelo time de Partner Ops da Wake (Renato). Foco comercial — apresentação das mecânicas de co-sell, expectativas e responsabilidades.
2. **Business Plan:** Formulário compartilhado (Google Docs) que a parceira deve preencher com visão de negócios, metas e plano de divulgação.
3. **Canal Slack:** Wake cria canal no Slack inicialmente apenas entre times de Parcerias Wake e equipe Omnichat. Após treinamento, expande para liderança de CS.
4. **Treinamento:** Rodolfo (Omnichat) visita escritório da Wake (LWSA) para treinamento presencial com time de CS e comercial. Horário ideal: 14h (padrão do time de CS). Após 17h a adesão é menor.
5. **Corrida de Vendas:** Campanha de incentivo para o time de CS Wake indicar oportunidades à Omnichat, com regras de repasse (revenue share). Regras enviadas em PDF separado.
6. **Mídia Kit:** Omnichat deve fornecer logo (450x450px, variações preto/branco/transparente) e descrição atualizada do "Quem Somos" para publicação no site de parceiros Wake.

---

### BPR-001 — Construção de case de sucesso conjunto (Norte Refrigeração)

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Wake & Omnichat |
| **Data** | 2025-11-26 |
| **Categoria** | `best_practices` |
| **Tags** | `proposta`, `audience-pam`, `audience-sales` |

**Recomendação:** Construir cases de sucesso conjuntos com clientes em comum logo no primeiro faturamento da parceria, seguindo processo padronizado:
1. Omnichat (parceiro) produz o conteúdo do case internamente
2. Cliente (Eduardo da Norte Refrigeração) aprova o conteúdo
3. Case é enviado para aprovação da Wake: diego@wake.tech, rafael.correia@wake.tech, matheus.silva@wake.tech e leticia.cena@wake.tech
4. Publicação alinhada com time de PMM (Mari Pellegrino) e content marketing (Matheus Coneglian) da Wake
5. Calendário de publicação coordenado com o primeiro faturamento da parceria

**Rationale:** Resultado da integração Wake + Omnichat no cliente Norte Refrigeração: a Omnichat consegue recuperar em média 30% dos carrinhos abandonados, e e-commerces têm em média 70-80% de abandono de carrinho. O impacto em GMV é significativo — esse dado foi usado como argumento comercial com prospects durante negociação. Cases concretos com números reais geram credibilidade rapidamente e aceleram conversões de novos prospects.

---

### BPR-002 — Estrutura da campanha "Corrida de Vendas" para parceiro tech

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Wake & Omnichat |
| **Data** | 2026-02-06 |
| **Categoria** | `best_practices` |
| **Tags** | `marketing-eventos`, `audience-pam` |

**Recomendação:** Estruturar a campanha "Corrida de Vendas" para ativação do time de CS Wake na geração de leads para parceiros tech seguindo o modelo:
- Wake cria canal Slack exclusivo entre Parcerias Wake, liderança de CS (Tier 1/2 e Tier 3) e equipe do parceiro
- Fase inicial: canal restrito ao time de Parcerias e liderança para preparar a campanha sem criar expectativas antes do treinamento
- Após treinamento presencial do parceiro, canal é expandido para o time de CS
- Regras da corrida de vendas enviadas em PDF para aprovação do parceiro
- A campanha define incentivos para o time de CS por indicações que gerem: agendas com clientes, leads qualificados e suporte em conversão
- Preparar Landing Page para automatizar o cadastro de oportunidades indicadas pelo CS Wake (como a Omnichat fez)

**Rationale:** O time de CS Wake tem acesso direto à base de clientes e pode identificar oportunidades de cross-sell para soluções de parceiros tech, mas precisa de incentivos claros e de canal estruturado para reportar essas indicações. Sem campanha estruturada com regras escritas, a ativação fica dependente de boa vontade e produz volume errático.

---

### FAQ-003 — Integração Omnichat com Wake Experience

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Wake & Omnichat |
| **Data** | 2026-02-06 |
| **Categoria** | `faq` |
| **Tags** | `integração`, `audience-solutions` |

**Pergunta:** Como está a integração da Omnichat com o Wake Experience (CRM/DXP)?

**Resposta:** No momento (fev/2026), não há integração entre Omnichat e Wake Experience. Existe uma oportunidade identificada pelo time de CS que pode abrir uma frente para expandir as possibilidades — ponto a ser discutido entre os times técnicos.

---

### HTO-003 — Dia do Parceiro na LWSA (Wake Partner Day)

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Wake & Omnichat |
| **Data** | 2025-08-13 |
| **Categoria** | `how_to` |
| **Tags** | `marketing-eventos`, `audience-pam` |

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

### FAQ-004 — B2C e B2B no mesmo ambiente Wake: é possível?

| Campo | Valor |
|---|---|
| **Grupo/Chat** | FG + Wake |
| **Data** | 2025-06-26 |
| **Categoria** | `faq` |
| **Tags** | `homologação`, `audience-solutions`, `audience-sales` |

**Pergunta:** É possível criar uma loja B2C e B2B no mesmo admin da Wake, ou é necessário contratar dois ambientes separados?

**Resposta:** É possível ter B2C e B2B no mesmo ambiente Wake. Adicionalmente, é possível configurar login com senha para clientes B2B, de modo que a tabela de preços diferenciada não fique visível para clientes B2C.

---

### TRB-002 — Leads enviados com erro pelo Salesforce (carrossel)

| Campo | Valor |
|---|---|
| **Grupo/Chat** | FG + Wake |
| **Data** | 2025-07-07 |
| **Categoria** | `troubleshooting` |
| **Tags** | `integração`, `audience-pam` |

**Sintomas:** Leads chegando para a agência via carrossel de forma equivocada — clientes já existentes na base, erros do Salesforce enviando oportunidades que não deveriam ir para agências parceiras. Em julho/2025 ocorreram múltiplos casos (Farmácia Cruzeiro, Pássaro Marron, Fonseca Shop, A Dias, Casa Bergan, DotLiva) onde a Wake solicitou às agências que desconsiderassem os leads recebidos. Impacto para a agência: a FG reportou que isso atrapalha o fluxo de trabalho porque a equipe para para conversar com o lead, acionar o vendedor interno e cadastrar no CRM antes de receber o pedido de desconsiderar.

**Causa raiz:** Wake estava em fase de atualizações do sistema Salesforce, identificando melhorias; durante essa fase, regras de distribuição do carrossel disparavam para casos que não deveriam ser elegíveis.

**Resolução:** As desconsiderações foram garantidas como não afetando a fila/posição das agências no carrossel. Boa prática: quando a Wake pede para desconsiderar, explicitar que a posição na fila de carrossel não foi prejudicada.

---

### TRB-003 — Agência Elite não aparece no site de parceiros Wake

| Campo | Valor |
|---|---|
| **Grupo/Chat** | FG + Wake |
| **Data** | 2025-07-22 |
| **Categoria** | `troubleshooting` |
| **Tags** | `marketing-eventos`, `audience-pam` |

**Sintomas:** A FG (agência certificada Elite) não aparecia na listagem de parceiros Elite no site wake.tech, e o logo exibido estava desatualizado. URL afetada: `https://wake.tech/encontre-agencias-parceiras/?_sft_solucao=wake-commerce&_sft_certificacao_agencia=elite`. Problema recorrente — a FG já havia reportado e corrigido anteriormente, mas voltou a desaparecer.

**Causa raiz:** [fonte original sem informação] — falha recorrente na manutenção do diretório de parceiros pelo time de marketing Wake.

**Resolução:** Wake solicitou o logo atualizado em formato 450x450px, e o time de marketing foi acionado para incluir a FG na listagem Elite. A atualização foi direcionada ao time de marketing da Wake. Lição: agências Elite devem monitorar periodicamente se estão aparecendo corretamente no diretório. Ao atualizar logo, enviar em 450x450px.

---

### HTO-004 — Registro de leads no Salesforce para contabilizar na campanha de parceiros

| Campo | Valor |
|---|---|
| **Grupo/Chat** | FG + Wake |
| **Data** | 2025-07-23 |
| **Categoria** | `how_to` |
| **Tags** | `integração`, `audience-pam` |

**Regra importante para eventos e feiras (ex: ECBR):**

Se um cliente de interesse da agência aparecer no stand da Wake durante um evento sem estar previamente registrado no Salesforce com a agência como "PARCEIRO ORIGEM", a origem do lead será contabilizada como "Marketing" — não para a agência.

**Processo correto:**
1. Antes do evento, a agência deve passar à Wake os dados dos clientes que pretende levar ou que estão no radar.
2. A Wake cadastra esses leads no Salesforce com a agência como parceiro de origem.
3. Qualquer negócio fechado com esses leads será atribuído à agência na campanha.

**Dados necessários para cadastro:** Nome, e-mail e contato do prospect.

---

### HTO-005 — Alteração de e-mail de notificações de leads no Wake Community

| Campo | Valor |
|---|---|
| **Grupo/Chat** | FG + Wake |
| **Data** | 2025-07-15 |
| **Categoria** | `how_to` |
| **Tags** | `programa-parceiros`, `audience-pam` |

**Situação:** A FG solicitou alterar o e-mail que recebe notificações de novos leads para um endereço de grupo (comercial@agenciafg.com.br), para que múltiplos responsáveis recebessem cópia e agilizassem o atendimento.

**Processo:**
1. Solicitar ao time de Partner Ops da Wake (Renato) a alteração via mensagem no grupo ou WhatsApp.
2. Informar o novo e-mail e se deseja alterar apenas notificações ou também o e-mail de acesso ao Community.
3. Partner Ops realiza a alteração diretamente no sistema.

**Contato Partner Ops:** Renato (Wake) — responsável por demandas relacionadas ao Community.

---

### BPR-003 — Qualificação e ICP de leads para agências parceiras

| Campo | Valor |
|---|---|
| **Grupo/Chat** | FG + Wake |
| **Data** | 2025-07-14 |
| **Categoria** | `best_practices` |
| **Tags** | `lead-pipeline`, `audience-pam` |

**Recomendação:** Mapear o ICP de cada agência parceira via formulário (https://forms.gle/shdjyxJQUwJ26EaH9) e calibrar a distribuição de leads do carrossel com base nesse perfil:
- Wake deve entender o perfil comercial de cada agência para calibrar a distribuição de leads
- Leads focados em preço baixo devem ser direcionados para agências com modelo de negócios correspondente
- O formato ideal de lead para a FG: prospect já em avaliação de migração e em negociação, onde a agência pode entrar com proposta de projeto + plataforma

**Rationale:** A FG reportou dificuldade em fechar leads com foco exclusivo em preço — clientes que buscam projetos por R$ 15-20k. A agência não consegue competir nesse perfil e questiona se esses leads estão dentro do ICP Wake ou são perfil de plataformas de entrada (Tray, etc.). Distribuir leads sem calibrar pelo ICP da agência destrói NPS do parceiro e desperdiça leads que outras agências converteriam.

---

### BPR-004 — Gestão de pipeline de leads entre Wake e agências (planilha compartilhada)

| Campo | Valor |
|---|---|
| **Grupo/Chat** | FG + Wake |
| **Data** | 2025-06-25 |
| **Categoria** | `best_practices` |
| **Tags** | `lead-pipeline`, `audience-pam` |

**Recomendação:** Estabelecer pipeline compartilhado de leads entre Wake e agência usando Google Sheets com acesso compartilhado para ambos os times atualizarem status (alternativa às planilhas Excel bloqueadas para usuários externos). Campos recomendados na planilha: nome do lead, executivo responsável (Wake e agência), status atual, próximos passos, data de atualização. Fluxo paliativo quando a planilha não está disponível: Wake envia lista de oportunidades com status atual no WhatsApp do grupo; agência retorna com status atualizado do lado delas dentro de 24h.

**Rationale:** Acompanhar lead a lead em call ao final do mês é inviável pela correria das equipes. Tentativa inicial da FG de compartilhar planilha Excel foi bloqueada pelo sistema (usuários não-Wake não podiam receber o arquivo). Sem visibilidade compartilhada, leads são perdidos por falta de coordenação entre executivo Wake e agência.

---

### FAQ-005 — Wake Education: certificação e acesso para agências

| Campo | Valor |
|---|---|
| **Grupo/Chat** | FG + Wake |
| **Data** | 2025-06-30 |
| **Categoria** | `faq` |
| **Tags** | `marketing-eventos`, `audience-pam` |

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

### HTO-006 — Processo de integração técnica com novo parceiro tech via Wake

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Wake & Leanwork |
| **Data** | 2025-07-02 |
| **Categoria** | `how_to` |
| **Tags** | `integração`, `audience-solutions`, `audience-pam` |

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

### HTO-007 — Acesso ao Slack da Wake para parceiros tech

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Wake & Leanwork |
| **Data** | 2025-07-24 |
| **Categoria** | `how_to` |
| **Tags** | `integração`, `audience-solutions` |

**Situação:** Leanwork precisava de acesso ao canal de CS da Wake no Slack para coordenação de projetos.

**Processo:**
1. Parceiro tech solicita o convite de acesso ao canal Slack da Wake.
2. Wake (time de CS ou Parcerias) envia o convite por e-mail.
3. Requer plano pago do Slack do lado do parceiro (plano Pro ou superior).
4. Após aceite do convite, confirmar no grupo WhatsApp.

---

### TRB-004 — Integração APP Max pendente de homologação

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Wake & Leanwork |
| **Data** | 2025-07-23 |
| **Categoria** | `troubleshooting` |
| **Tags** | `homologação`, `audience-solutions`, `audience-pam` |

**Sintomas:** Integração da APP Max iniciada em setembro sem ter sido homologada até julho do ano seguinte (~10 meses).

**Causa raiz:** Integrações sem SLA de homologação definido e sem responsáveis claramente atribuídos ficam paradas por longos períodos.

**Resolução:** Leanwork solicitou ajuda do time de Partner Ops Wake (Renato) para desbloquear o processo de homologação. Partner Ops verificou internamente e retornou com plano de ação para seguir com o caso. Aprendizado: estabelecer prazos e responsáveis no início do processo e escalar via Partner Ops quando travado.

---

### BPR-005 — Lançamento de parceria com cobertura de mídia coordenada

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Wake & Leanwork |
| **Data** | 2025-08-06 |
| **Categoria** | `best_practices` |
| **Tags** | `marketing-eventos`, `audience-pam` |

**Recomendação:** Para anunciar uma parceria, planejar calendário coordenado de divulgação com aprovação bilateral dos materiais. Modelo aplicado ao lançamento Wake + Leanwork:

Calendário coordenado:
- **Imprensa:** 12/08
- **Blog:** 14/08
- **Redes Sociais:** 15/08

Materiais produzidos:
- Press release (nota de parceria) — aprovado por ambas as partes antes de publicar
- Blog post — enviado pelo time de marketing Wake para aprovação da Leanwork
- Inclusão do parceiro no site de parceiros Wake com logo correto

Quem aprova: a Leanwork faz revisão do conteúdo; o time de marketing Wake (Natasha) ajusta e publica conforme as datas acordadas. Ponto de atenção: enviar logo em alta resolução com fundo transparente para evitar problemas de exibição no site.

**Rationale:** Lançamento descoordenado fragmenta a mensagem e reduz o alcance. Calendário sequenciado (imprensa → blog → redes) constrói momentum e maximiza visibilidade da parceria.

---

### BPR-006 — Modelo de co-sell: como apresentar orçamento de integração a prospect

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Wake & Leanwork |
| **Data** | 2025-08-07 |
| **Categoria** | `best_practices` |
| **Tags** | `lead-pipeline`, `audience-pam`, `audience-sales` |

**Recomendação:** Em co-sell de parceiro tech com prospect, incluir o orçamento de integração já na proposta comercial inicial seguindo o fluxo:
1. Wake comunica à Leanwork o nome dos prospects e escopo básico
2. Leanwork compartilha link de agenda para call de escopo com o prospect
3. Após call, Leanwork prepara orçamento de integração
4. Para minimizar barreiras de aprovação: quanto menor o investimento inicial, mais fácil a liberação de verba. Usar como referência o modelo/blueprint de integrações anteriores (ex: Koin)

**Rationale:** Incluir integração no "pacotão" da proposta inicial aumenta o valor percebido e facilita o fechamento — o prospect não precisa negociar separadamente com o parceiro tech depois. Quando a integração é proposta em momento posterior, o cliente já tem o orçamento "fechado mentalmente" e qualquer adição é vista como surpresa de custo.

---

### FAQ-006 — Corrida de Vendas: como funciona e critérios de elegibilidade

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Wake & Leanwork |
| **Data** | 2025-08-07 |
| **Categoria** | `faq` |
| **Tags** | `marketing-eventos`, `audience-pam` |

**Pergunta:** Como funciona a corrida de vendas para parceiros tech da Wake?

**Resposta:** A corrida de vendas é uma campanha de incentivo onde o parceiro tech (ex: Leanwork) e os times da Wake trabalham em conjunto para fechar negócios. O parceiro deve:
1. Fazer adesão à campanha via formulário no prazo estipulado.
2. Obter certificação Wake até a data limite definida na campanha para ser elegível.
3. Ter o material de divulgação pronto para ser disparado junto com o lançamento.

As regras detalhadas são enviadas em PDF ("Regras Base Corrida de Vendas"). O material de divulgação da campanha é coordenado com o time de marketing Wake para coincidir com o lançamento oficial da parceria.

---

### TRB-005 — Logo do parceiro tech desatualizado ou pequeno no site Wake

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Wake & Leanwork |
| **Data** | 2025-08-13 |
| **Categoria** | `troubleshooting` |
| **Tags** | `marketing-eventos`, `audience-pam` |

**Sintomas:** O logo da Leanwork no site Wake estava pequeno e não ocupava todo o espaço disponível.

**Causa raiz:** Logo enviado sem dimensões adequadas / em baixa resolução, impedindo expansão ao tamanho-padrão do espaço reservado no site de parceiros.

**Resolução:** Solicitar ao time de marketing Wake (via Partner Ops ou grupo WhatsApp) que ajuste o tamanho do logo para ocupar todo o espaço reservado — expandindo a imagem sem perder qualidade. Como prevenir: enviar logo em alta resolução, idealmente em formato vetorial ou PNG com fundo transparente, com dimensões que permitam expansão. A Wake usa espaços padronizados no site de parceiros.

---

### HTO-008 — Eventos pós-fórum: processo de inscrição de leads como convidados

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Wake & Leanwork |
| **Data** | 2025-07-28 |
| **Categoria** | `how_to` |
| **Tags** | `marketing-eventos`, `audience-pam` |

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

### HTO-009 — Processo de acesso ao ambiente de demonstração Wake para agências

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Eficaz <> WAKE |
| **Data** | 2025-09-16 |
| **Categoria** | `how_to` |
| **Tags** | `integração`, `audience-pam` |

**Processo para criar acesso restrito ao painel Wake para agências:**

1. Agência solicita ao Partner Ops Wake (Renato) a criação de usuário.
2. Informar o e-mail no qual o acesso deve ser gerado (ex: parcerias@eficazmarketing.com).
3. Partner Ops cria o usuário mas não ativa até confirmação da agência.
4. Agência confirma ("Ok") e Partner Ops ativa o usuário e dispara o e-mail de acesso.

**Para acesso de clientes ao painel:** O cliente precisa acionar o suporte Wake diretamente. Apenas mediante pedido do cliente, os acessos de uma agência ao admin da conta do cliente são liberados. A agência deve fazer a ponte com o cliente para que ele abra o ticket de suporte.

**Acesso Loja Wake (ambiente de parceiro):** Parceiros têm acesso a um ambiente específico para demonstrações. O usuário parceiro com acesso restrito permite visualizar o painel sem permissão de edição em produção.

---

### TRB-006 — Perda de SEO orgânico após migração de plataforma

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Eficaz <> WAKE |
| **Data** | 2025-07-08 |
| **Categoria** | `troubleshooting` |
| **Tags** | `integração`, `audience-solutions`, `audience-cs` |

**Sintomas:** Cliente OG Brindes migrou da Tray para Wake há 90 dias e perdeu todo o tráfego orgânico construído ao longo de mais de 10 anos de e-commerce.

**Causa raiz:** Migração de plataforma sem estratégia de SEO — ausência de redirecionamentos 301, falta de migração consistente de URLs, perda de meta-dados e mapeamento incompleto de sitemap.

**Resolução:** Agência acionou o CS Wake (Janaina) para entender o que ocorreu e intervir junto ao cliente. Recomendação para agências: antes de qualquer migração para Wake, realizar auditoria de SEO da plataforma atual e mapear todos os redirecionamentos necessários. Incluir esse item como entregável no TAP (Termo de Abertura de Projeto). É fundamental incluir um plano de preservação de SEO no escopo de migração.

---

### TRB-007 — Evento Purchase não registrando receita no Meta Ads

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Eficaz <> WAKE |
| **Data** | 2025-12-15 |
| **Categoria** | `troubleshooting` |
| **Tags** | `marketing-eventos`, `audience-solutions`, `audience-cs` |

**Sintomas:** Para o cliente Voce Constrói (https://www.voceconstroi.com.br/), o evento Purchase disparava no Meta mas sem o parâmetro de receita (`value`). O Meta registrava a venda com valor R$ 0,00, destruindo métricas de ROAS e otimização de campanhas. Problema iniciou após ajuste no Front feito pela agência (Auaha) em novembro. Já havia retornado a funcionar temporariamente após retagueamento, mas voltou a falhar em 06/11/2025. Ticket 258709 aberto em dezembro sem resolução. Em fevereiro/2026 o ticket foi reaberto (269708) com urgência.

**Causa raiz:** O parâmetro `value` do evento de compra estava sendo perdido após um ajuste de front-end. O suporte Wake informou que a análise estava fora do escopo padrão (sendo atendida em caráter de exceção).

**Resolução:** Sempre após qualquer alteração no front-end do checkout, verificar se os eventos de conversão do Meta (e GA4) continuam disparando corretamente com todos os parâmetros (especialmente `value`, `currency` e `event_id`). Usar o Meta Events Manager e o Meta Pixel Helper para validação imediata. Em caso de persistência, reabrir ticket Wake e escalar via parceiro de canais.

---

### BPR-007 — Role Play: treinamento de pitch comercial e técnico para agências

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Eficaz <> WAKE |
| **Data** | 2025-10-02 |
| **Categoria** | `best_practices` |
| **Tags** | `proposta`, `audience-pam`, `audience-sales` |

**Recomendação:** Realizar sessões de Role Play Wake x Agência para fortalecer o pitch técnico e comercial da agência:

- **Objetivo:** Avaliar e fortalecer o pitch técnico e comercial da agência para venda de projetos Wake
- **Participantes:** Diego (Wake, papel de CEO da empresa cliente), Didier (Wake, papel de gerente de TI), equipe da agência (papel de si mesmos vendendo Wake)
- **Formato:** Simulação de venda de projeto de e-commerce + plataforma Wake para cliente fictício
- **Roteiro:** Wake fornece um case/roteiro com contexto do cliente para guiar a simulação
- **Duração:** 1-2 horas
- **Tom:** Não é julgamento — é desenvolvimento. O objetivo é ajudar a agência a ser mais forte no discurso
- **Orientações:** Enviadas em PDF com antecedência ("Orientações Role Play - Agências")

**Rationale:** Ideal para agências em início de parceria ou que precisam calibrar o discurso técnico e comercial. Bom para identificar lacunas no conhecimento da plataforma antes de reuniões com clientes reais — ensaiar antes em ambiente seguro reduz risco de perder deals por gaps de discurso.

---

### BPR-008 — Relatório semanal de pipeline de agência parceira

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Eficaz <> WAKE |
| **Data** | 2025-07-23 |
| **Categoria** | `best_practices` |
| **Tags** | `integração`, `audience-pam` |

**Recomendação:** Agências parceiras devem enviar relatório semanal de pipeline ao parceiro de canais Wake (WhatsApp) e relatório mensal completo por e-mail, complementados por forecast quinzenal por videoconferência. Estrutura recomendada:

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

Frequência: semanal no WhatsApp, mensal por e-mail (relatório completo). Complemento: forecast quinzenal entre agência e executivo Wake para alinhamento de pipeline, estratégia e próximos passos (agenda via Google Meet).

**Rationale:** Modelo adotado pela Eficaz. Sem ritmo formal de reporte, o parceiro de canais Wake fica sem visibilidade real do pipeline da agência, perde a chance de apoiar deals em momento crítico e não consegue prever resultado do trimestre. Cadência semanal + forecast quinzenal equilibra disciplina e leveza operacional.

---

### FAQ-007 — Programas Storefront 2.0 e Redesign de Checkout para agências parceiras

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Eficaz <> WAKE |
| **Data** | 2025-08-27 |
| **Categoria** | `faq` |
| **Tags** | `integração`, `audience-solutions`, `audience-pam` |

**O que são os programas StoreFront 2.0 e Redesign de Checkout:**

Programas lançados pela Wake para impulsionar agências parceiras com modelos pré-definidos e exclusivos (sem concorrência direta entre agências).

- **StoreFront 2.0:** Escopos Essencial, Básico, Intermediário e Personalizado.
- **Redesign de Checkout:** Escopos Essencial e Personalizado.

**Vantagens:**
- Preços e escopos criados em parceria com agências para agilizar o processo de vendas.
- Modelos exclusivos por agência — evita que duas agências Wake ofereçam o mesmo template para o mesmo mercado.

**Adesão:** Opcional. Agências que não aderirem continuam recebendo indicações padrão do carrossel de leads. Para aderir, responder "Ok" com interesse no pacote desejado.

---

### FAQ-008 — Leads de canais: responsabilidade da agência na apresentação da Wake

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Eficaz <> WAKE |
| **Data** | 2025-06-30 |
| **Categoria** | `faq` |
| **Tags** | `lead-pipeline`, `audience-pam`, `audience-sales` |

**O que são "leads de canais":**
Oportunidades onde a Wake não tem executivo comercial atuando diretamente. A agência recebe o lead e é responsável por toda a condução do processo — incluindo apresentar a plataforma Wake e passar a precificação ao prospect.

**Diferença do modelo padrão:**
- No modelo padrão: Wake tem executivo atuando em paralelo à agência.
- No modelo canais: agência conduz sozinha a venda da plataforma + escopo de projeto.

**Implicação prática:** A agência precisa estar certificada e preparada para responder perguntas técnicas e comerciais sobre a Wake sem apoio imediato do time de vendas Wake.

**Exemplo de leads de canais recebidos pela Eficaz:** Casafer, Fox Racing, Campinas Diesel, Festa Color, Cariñon Mix, Gangue do Caramelo, Pampulha Enxovais, Agroforte.

---

### HTO-010 — Aproximação com produto Wake Experience (CRM/DXP) para agências parceiras

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Eficaz <> WAKE |
| **Data** | 2025-07-15 |
| **Categoria** | `how_to` |
| **Tags** | `integração`, `audience-pam`, `audience-sales` |

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

### BPR-009 — Feedback de leads perdidos: como usar para calibrar parceria

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Eficaz <> WAKE |
| **Data** | 2025-08-19 |
| **Categoria** | `best_practices` |
| **Tags** | `lead-pipeline`, `audience-pam` |

**Recomendação:** Times de Parcerias Wake devem usar feedback de leads perdidos como insumo de calibração contínua, aplicando:
- Definir em qual estágio da negociação o lead é passado à agência (preferencialmente após qualificação mínima de interesse/fit)
- Criar SLA claro: agência deve abordar lead em X dias; se não houver progresso, Wake retoma contato
- Diferenciar leads "quentes" (em fase de decisão) de "frios" (prospecção inicial) na comunicação para a agência
- Estimular agências a construir relacionamento direto com clientes da base Wake (via CS) — agências que fazem isso têm mais sucesso do que as que dependem exclusivamente dos leads enviados pelo carrossel

**Rationale:** Diagnóstico compartilhado pela Eficaz com a Wake (agosto/2025):
- Leads da Wake chegavam "muito frios" — queriam falar com agência só depois de fechar com a plataforma
- Consultor Wake pedia para aguardar contato do cliente → negociação travava
- No mês referência, nenhuma indicação da Wake gerou proposta
- Prazo médio para mandar proposta aumentou de 2 semanas para quase 2 meses
- SDR em fase de qualificação → impacto na geração de oportunidades

Sem mecanismo de feedback estruturado, a Wake não consegue identificar gaps de processo e a agência se sente frustrada com leads desqualificados.

---

*Fim do documento kb-parceiros-tech.md.*

---

# KB Extract — Wave 6C (Eitri, O2, Social, Tuna)
> Extracted: 2026-06-04

---

## Grupo: Eitri & Wake
**JID:** `120363400732871223@g.us`

---

### HTO-001 — Como registrar indicações de eventos para contabilização no Salesforce

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Eitri & Wake |
| **Data** | 2025-06-23 |
| **Categoria** | `how_to` |
| **Tags** | `integração`, `audience-pam` |

Para que uma indicação de cliente feita por um parceiro seja contabilizada como originada pelo parceiro, é obrigatório que o Lead e a Oportunidade já existam no Salesforce com o campo **PARCEIRO ORIGEM** preenchido com o nome da agência **antes** de qualquer fechamento. Isso vale especialmente para indicações feitas no contexto de eventos. A criação retroativa da oportunidade após o fechamento não garante a contabilização.

---

### HTO-002 — Como solicitar acesso sandbox para novos parceiros tech

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Eitri & Wake |
| **Data** | 2025-06-24 |
| **Categoria** | `how_to` |
| **Tags** | `homologação`, `audience-solutions`, `audience-pam` |

Para criar acessos sandbox para um parceiro tech, o parceiro deve informar ao time de Parcerias da Wake os e-mails dos usuários que precisam de acesso. Os demais acessos podem ser criados diretamente pelo parceiro após o ambiente ser provisionado. Exemplo de fluxo: parceiro informa os e-mails principais → Wake cria o ambiente → parceiro gerencia usuários adicionais.

---

### BPR-001 — Modelo de engajamento embarcado com parceiros tech

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Eitri & Wake |
| **Data** | 2025-06-24 |
| **Categoria** | `best_practices` |
| **Tags** | `lead-pipeline`, `audience-pam` |

**Recomendação:** Discutir com parceiros tech que tenham produtos complementares à plataforma o modelo de entrega "embarcada" — solução do parceiro apresentada como parte do pacote Wake, com o faturamento consolidado pela Wake.

**Rationale:** A solução embarcada reduz o atrito na jornada do cliente: facilita a adoção (cliente não precisa fechar contrato separado com o parceiro tech) e aumenta a percepção de valor da parceria (uma proposta integrada vale mais que duas propostas avulsas).

---

### TRB-001 — reCAPTCHA ativado automaticamente no app bloqueando vendas

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Eitri & Wake |
| **Data** | 2026-05-11 |
| **Categoria** | `troubleshooting` |
| **Tags** | `integração`, `audience-solutions`, `audience-cs` |

**Problema:** A loja DLK Modas ficou sem vender no app porque o reCAPTCHA foi ativado automaticamente para o fluxo de autenticação. O parceiro tech (Eitri) não usa reCAPTCHA no app e não conseguia desativá-lo pelo painel, pois a opção de desativação manual foi removida da interface.

**Causa:** O reCAPTCHA pode ser ativado por algum evento ou trigger do sistema Wake, sem ação direta do lojista ou parceiro. Apps que usam tokens dedicados ficam vulneráveis a essa reativação.

**Resolução:**
1. Abrir ticket de suporte solicitando a desativação do reCAPTCHA para o token específico do app: https://atendimento.wake.tech/hc/pt-br/articles/37812329587991
2. Ticket de referência: https://atendimento.wake.tech/hc/pt-br/requests/279913
3. O partner manager pode acionar internamente para priorização.

**Solução de longo prazo em discussão:** Permitir que o parceiro tech informe sua própria chave de reCAPTCHA no painel Wake (similar ao que é feito na VTEX), o que tornaria a experiência mais segura e sem atrito. Lojas impactadas: DLK, Shoulder, Balaroti, Peahi.

---

### FAQ-001 — Posso adicionar minha solução de parceiro no app de evento da Wake Summit?

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Eitri & Wake |
| **Data** | 2025-06-25 |
| **Categoria** | `faq` |
| **Tags** | `marketing-eventos`, `audience-pam` |

**Pergunta:** A Eitri tem experiência com apps de eventos e gostaria de oferecer sua solução para o Wake Summit. Como proceder?

**Resposta:** A indicação correta é comunicar o interesse ao parceiro de parcerias (no caso, Rafael Correia) para que ele orquestre internamente junto ao time de marketing da Wake. O parceiro deve comunicar brevemente a solução e cases existentes para embasar a conversa interna.

---

### HTO-003 — Como obter o selo de parceiro oficial Wake para apresentações

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Eitri & Wake |
| **Data** | 2025-08-15 |
| **Categoria** | `how_to` |
| **Tags** | `marketing-eventos`, `audience-pam` |

Para obter o selo/badge de parceiro oficial Wake para uso em apresentações e materiais de marketing, o parceiro deve solicitar ao seu parceiro de parcerias (Account Manager) via WhatsApp ou e-mail. O material é enviado por e-mail. O selo é diferente do logo; é a badge específica de "Parceiro Oficial Wake".

---

### HTO-004 — Como funciona o Wake Partner Day (dia presencial semanal)

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Eitri & Wake |
| **Data** | 2025-08-13 |
| **Categoria** | `how_to` |
| **Tags** | `marketing-eventos`, `audience-pam` |

O Wake Partner Day é o dia oficial de presença de parceiros no escritório Wake. A partir de 27/08/2025, passou das quintas para as **quartas-feiras**. O evento ocorre no escritório LWSA (Rua Itapaiúna, 2434 – São Paulo/SP). O objetivo é aproximar parceiros dos times de CSM, vendas e marketing, com foco em construção de ofertas conjuntas e trabalho de base. Parceiros de fora de SP são bem-vindos a trabalhar do escritório nesses dias.

**Para participar presencialmente:** informar nome, e-mail, CPF e placa do carro (se vier de carro) ao parceiro de parcerias para liberação na portaria e acesso ao Wi-Fi.

---

### BPR-002 — Como preparar Press Release para anúncio de parceria Wake

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Eitri & Wake |
| **Data** | 2025-09-22 |
| **Categoria** | `best_practices` |
| **Tags** | `marketing-eventos`, `audience-pam` |

**Recomendação:** Para publicar o anúncio de parceria no blog da Wake, o parceiro deve preparar um Press Release contendo:
- Descrição de quem são e o que fazem
- Citação de um executivo da empresa

Exemplo de publicação final: https://wake.tech/blog/wake-parceria-biso-inteligencia-de-dados/

O documento é revisado pelo time de marketing da Wake antes da publicação. Após a aprovação, a Wake também publica um artigo no blog complementar.

**Rationale:** Press Release estruturado com citação executiva e descrição clara da empresa permite ao time de marketing Wake publicar com qualidade editorial e velocidade. Sem material padronizado, cada anúncio exige retrabalho de copy e demora mais para ir ao ar.

---

### HTO-005 — Como acionar suporte Wake durante a Black Friday

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Eitri & Wake |
| **Data** | 2025-11-24 |
| **Categoria** | `how_to` |
| **Tags** | `marketing-eventos`, `audience-cs`, `audience-solutions` |

Durante a Black Friday, a Wake oferece atendimento especial com canais e plantão dedicados para urgências. Os detalhes de horários e contatos estão disponíveis em:
https://atendimento.wake.tech/hc/pt-br/articles/35396146551447

Os parceiros tech devem também informar ao time Wake:
1. Se terão escala/plantão em horário especial
2. Qual o contato de urgência que a Wake deve acionar

Esse alinhamento prévio garante agilidade no atendimento conjunto durante o período crítico.

---

### HTO-006 — Como patrocinar o Wake On The Road

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Eitri & Wake |
| **Data** | 2026-03-04 |
| **Categoria** | `how_to` |
| **Tags** | `marketing-eventos`, `audience-pam` |

O Wake On The Road é um evento proprietário Wake que conecta clientes, prospects e parceiros em ambiente exclusivo (~40 convidados, decisores de e-commerce/tecnologia/digital). São oferecidas **3 cotas de patrocínio** (Ouro, Prata, Bronze) com exclusividade por categoria.

**Investimento:** R$10k a R$30k
**Entregáveis incluem:**
- Convites para clientes estratégicos
- Logo nos materiais oficiais
- Menção institucional durante o evento
- Cobertura nas redes sociais Wake
- (Cota Ouro) 100% da lista de leads ao final do evento

Para participar, contatar o parceiro de parcerias via WhatsApp para receber o prospecto completo e reservar a cota.

---

## Grupo: O2 Consultoria & Wake
**JID:** `120363420874830254@g.us`

---

### TRB-002 — Community (Salesforce) aparecendo vazio — sem leads visíveis

| Campo | Valor |
|---|---|
| **Grupo/Chat** | O2 Consultoria & Wake |
| **Data** | 2025-06-23 |
| **Categoria** | `troubleshooting` |
| **Tags** | `programa-parceiros`, `audience-pam`, `audience-cs` |

**Sintomas:** O Community Wake (Salesforce) aparecia vazio para a agência O2, impossibilitando atualização de status dos leads.

**Causa raiz:** Instabilidade geral no sistema Salesforce, não exclusiva da O2.

**Resolução:** O problema foi tratado pela equipe técnica Wake. Para demandas relacionadas ao Community, acionar o responsável de Partner Ops (Renato, a partir de 2025-07-15).

**Workaround:** Enquanto o problema não era regularizado, a Account Manager da Wake atualizava os leads manualmente mediante envio de listagem pela agência.

---

### HTO-007 — Processo de repasse de leads pelo canal (Canais)

| Campo | Valor |
|---|---|
| **Grupo/Chat** | O2 Consultoria & Wake |
| **Data** | 2025-07-11 |
| **Categoria** | `how_to` |
| **Tags** | `lead-pipeline`, `audience-pam`, `audience-sales` |

Quando um lead é repassado pelo canal de Canais Wake, a **agência deve apresentar a plataforma Wake E fazer a proposta comercial** — não há um executivo Wake atuando diretamente nessa modalidade. O processo é ponta a ponta pela agência.

**O que a agência precisa fazer:**
1. Entrar em contato com o lead
2. Apresentar a Wake como solução (plataforma + benefícios)
3. Montar proposta comercial incluindo escopo do projeto + mensalidade Wake
4. Conduzir a negociação até o fechamento

A Wake fornece materiais de apoio (PDF de proposta, decks) e pode apoiar em casos específicos com o executivo de produto.

---

### FAQ-002 — A campanha de vendas Wake vale para parceiros tech também?

| Campo | Valor |
|---|---|
| **Grupo/Chat** | O2 Consultoria & Wake |
| **Data** | 2025-07-22 |
| **Categoria** | `faq` |
| **Tags** | `marketing-eventos`, `audience-pam` |

**Pergunta:** A campanha de vendas da Wake vale para parceiros que não são agências de implementação (ex.: empresas de consultoria/social)?

**Resposta:** Sim, a campanha vale. Para ser elegível, é necessário:
1. Fazer adesão pelo link da campanha até o prazo indicado
2. Tirar a certificação Wake até a data definida

Sem a certificação, o parceiro não é elegível para os prêmios.

---

### HTO-008 — Programas StoreFront 2.0 e Redesign de Checkout para agências

| Campo | Valor |
|---|---|
| **Grupo/Chat** | O2 Consultoria & Wake |
| **Data** | 2025-08-27 |
| **Categoria** | `how_to` |
| **Tags** | `integração`, `audience-solutions`, `audience-pam` |

A Wake lançou dois programas para impulsionar as agências parceiras:
- **StoreFront 2.0**: Escopos Essencial, Básico, Intermediário e Personalizado
- **Redesign de Checkout**: Escopos Essencial e Personalizado

**Vantagens:**
- Modelos pré-definidos exclusivos para a agência (sem concorrência direta)
- Preços e escopos criados em parceria com agências para agilizar o processo de vendas
- Fomenta negócios com lojistas Wake sem agência vinculada

**Como aderir:** Responder "Ok" ao convite da Wake informando qual pacote tem interesse. A adesão é opcional; sem adesão, a agência continua recebendo indicações padrão.

---

### BPR-003 — Como lidar com leads que não respondem após primeira tentativa

| Campo | Valor |
|---|---|
| **Grupo/Chat** | O2 Consultoria & Wake |
| **Data** | 2025-10-24 |
| **Categoria** | `best_practices` |
| **Tags** | `lead-pipeline`, `audience-pam`, `audience-sales` |

**Recomendação:** Quando o lead não entra na agenda confirmada ou não responde após primeira tentativa de contato:
- Registrar a tentativa de contato no Community
- Fazer pelo menos 2 a 3 tentativas via WhatsApp e e-mail em dias diferentes
- Informar o status ao parceiro de parcerias Wake para que possam co-prospectar
- Se o lead não responder por período prolongado (ex.: 2 semanas), encerrar a oportunidade no Community e comunicar a Wake
- Se o lead retornar depois do encerramento, reabrir a oportunidade no Community

**Rationale:** Tentativas múltiplas em canais diferentes (WhatsApp + e-mail) maximizam a chance de resposta sem virar "stalking". Comunicar a Wake ativamente permite co-prospecção (Wake aciona outro canal) e mantém o lead no funil corretamente atribuído.

---

### FAQ-003 — Qual o padrão correto para o selo Wake no rodapé dos sites de clientes?

| Campo | Valor |
|---|---|
| **Grupo/Chat** | O2 Consultoria & Wake |
| **Data** | 2025-09-18 |
| **Categoria** | `faq` |
| **Tags** | `marketing-eventos`, `audience-pam` |

**Pergunta:** Como deve ser exibido o selo Wake no rodapé dos sites de clientes?

**Resposta:** A Wake definiu um novo padrão para a exibição da marca nos rodapés. O guia inclui:
- Formatação do logo Wake junto com o logo da agência
- O texto "Plataforma de E-commerce" deve ser o **único termo com link** no rodapé
- Especificações de fonte e tamanho do texto

Os detalhes completos foram enviados por e-mail às agências parceiras. Em caso de dúvida, consultar o parceiro de parcerias.

---

### BPR-004 — Como usar notas de reunião do Gemini para briefar parceiros em leads

| Campo | Valor |
|---|---|
| **Grupo/Chat** | O2 Consultoria & Wake |
| **Data** | 2026-02-05 |
| **Categoria** | `best_practices` |
| **Tags** | `integração`, `audience-pam` |

**Recomendação:** Ao repassar leads para agências parceiras, compartilhar as anotações automáticas do Gemini das reuniões de qualificação no momento do repasse via WhatsApp.

**Rationale:** A prática permite que a agência entre na negociação já contextualizada, sem necessidade de repetição das informações pelo lead. Acelera a qualificação e garante alinhamento entre todos os envolvidos. Especialmente útil em leads quentes ("quentíssimos") onde a velocidade de resposta é crítica.

---

### HTO-009 — Como funciona o Wake Score e a Certificação Wake

| Campo | Valor |
|---|---|
| **Grupo/Chat** | O2 Consultoria & Wake |
| **Data** | 2026-01-27 |
| **Categoria** | `how_to` |
| **Tags** | `programa-parceiros`, `audience-pam` |

A **Certificação Wake é um requisito obrigatório** do Wake Score para agências parceiras.

**Regras:**
- Turmas abrem toda **segunda-feira**
- Inscrição ao longo da semana → acesso liberado na segunda seguinte
- Prazo de conclusão após liberação: **30 dias**
- A certificação é **individual** por pessoa (cada membro da equipe precisa tirar a sua)
- **Link de inscrição:** https://forms.gle/AA6NR6vL3uEJXhDu7

Em caso de divergência sobre certificações já realizadas em anos anteriores, enviar os certificados de conclusão ao time de Education Wake para regularização.

---

### FAQ-004 — Como funciona a nova precificação Wake 2026?

| Campo | Valor |
|---|---|
| **Grupo/Chat** | O2 Consultoria & Wake |
| **Data** | 2026-01-12 |
| **Categoria** | `faq` |
| **Tags** | `contrato-comercial`, `audience-sales`, `audience-pam` |

**Pergunta:** Houve mudanças nos planos e preços Wake em 2026?

**Resposta:** Sim. Em janeiro de 2026, a Wake comunicou a Nova Precificação Wake 2026 com atualizações nos planos e novas regras. O comunicado completo foi enviado por e-mail às agências parceiras. Em caso de dúvida, consultar o Account Manager de parcerias.

Ponto de atenção: novas negociações (contratos novos) já entram na tabela de 2026. Contratos existentes seguem as condições anteriores conforme aditivo.

---

## Grupo: Wake <> Social
**JID:** `120363402518075591@g.us`

---

### HTO-010 — Como solicitar loja adicional Wake para cliente existente (guarda-chuva da agência)

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Wake <> Social |
| **Data** | 2026-02-20 |
| **Categoria** | `how_to` |
| **Tags** | `onboarding`, `audience-pam`, `audience-cs` |

**Situação:** Agência quer criar uma nova loja Wake para um cliente que ainda não faturava online (ex.: cliente que estava apenas em marketplace).

**Processo correto:**
1. A abertura de ticket de suporte **NÃO** é o caminho correto para loja adicional
2. O processo deve ser feito **via e-mail com o Account Manager** da parceria
3. Se a loja ficará atrelada ao contrato existente da agência (guarda-chuva), assinar um **aditivo contratual**
4. **Custo:** R$1.000/mês por loja adicional vinculada ao contrato
5. Se for negociação nova (cliente novo, contrato separado), segue a tabela de preços 2026

**Contato para loja adicional:** ana.vicente@wake.tech com cópia para amanda.araujo@wake.tech

---

### FAQ-005 — O parceiro Social pode participar de ações comerciais Wake conjuntas?

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Wake <> Social |
| **Data** | 2025-07-04 |
| **Categoria** | `faq` |
| **Tags** | `marketing-eventos`, `audience-pam` |

**Pergunta:** Como funciona a parceria comercial entre uma agência de Social Media e a Wake?

**Resposta:** O modelo recomendado é o **Account Planning cruzado**:
1. Cruzar as bases de clientes de ambos os lados
2. Selecionar ~10 contas target de cada lado com potencial de complementaridade
3. Realizar um Wake Day conjunto para apresentação das contas e sinergia

A agência de Social traz conhecimento de marcas e contatos; a Wake entra com a oferta de plataforma. O timing de fechamento dos leads pode não coincidir com o timing de eventos, então não é necessário forçar fechamentos para participar de eventos como o Wake Summit.

---

### BPR-005 — Como batalhar contra Shopify em negociações com clientes indecisos

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Wake <> Social |
| **Data** | 2026-04-28 |
| **Categoria** | `best_practices` |
| **Tags** | `competitivo`, `audience-sales` |

**Recomendação:** Em negociações contra Shopify com clientes indecisos, aplicar as seguintes táticas:
- Apresentar depoimentos de clientes Wake do mesmo segmento
- Solicitar ao time Wake um battlecard técnico com comparativo Wake vs Shopify (confere mais credibilidade)
- Usar cases publicados no blog Wake do segmento do lead
- Pedir ao time comercial Wake para co-participar na negociação quando a logo for estratégica

**Rationale:** Leads que circulam em grupos de e-commerce ouvem muito sobre VTEX e Shopify; por isso é importante ter materiais de referência de terceiros (cases, testemunhos de outros lojistas) além da argumentação própria. Caso ilustrativo: lead Brinquedos Estrelas com preferência inicial por Shopify, apesar da agência ter apresentado todos os argumentos para Wake.

---

### HTO-011 — Como configurar call quinzenal de acompanhamento de onboarding com Wake

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Wake <> Social |
| **Data** | 2026-05-14 |
| **Categoria** | `how_to` |
| **Tags** | `onboarding`, `audience-pam`, `audience-cs` |

Para agências com múltiplos projetos em implementação simultânea, a Wake pode estabelecer uma **call quinzenal com o time de onboarding** para acompanhamento dos projetos até o go-live.

**Para configurar:**
1. Acordar frequência e horário com o Account Manager (ex.: sextas-feiras às 14h30)
2. Fornecer a lista de e-mails de todos do time que tocam os novos projetos
3. A Wake inclui os perfis de onboarding relevantes

Isso é especialmente útil quando o onboarding é conduzido pela agência, mas a Wake precisa de visibilidade sobre o andamento.

---

### FAQ-006 — Qual o canal correto para tirar dúvidas técnicas sobre a plataforma B2B?

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Wake <> Social |
| **Data** | 2025-10-06 |
| **Categoria** | `faq` |
| **Tags** | `integração`, `audience-solutions` |

**Pergunta:** Qual o canal correto para tirar dúvidas técnicas sobre a plataforma B2B Wake?

**Resposta:** Para dúvidas técnicas, usar o canal privado (PV) com o Account Manager ou enviar um áudio no privado descrevendo a dúvida. Isso agiliza o retorno pois o AM pode encaminhar à squad técnica correspondente sem o ruído do grupo geral.

Para treinamento de novos colaboradores sobre cadastro na plataforma, o time de suporte da Wake pode ser acionado diretamente.

---

### HTO-012 — Como solicitar criação de ambiente Wake para novo projeto via guarda-chuva

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Wake <> Social |
| **Data** | 2026-03-31 |
| **Categoria** | `how_to` |
| **Tags** | `integração`, `audience-pam`, `audience-cs` |

**Processo correto para criação de nova instância/loja Wake por agência parceira:**

A solicitação **NÃO deve ser feita por ticket de suporte**. O fluxo correto é:

1. Enviar e-mail para **ana.vicente@wake.tech** com cópia para **amanda.araujo@wake.tech**
2. O e-mail deve conter confirmação formal do tipo: *"Informo que estamos de acordo com a liberação do ambiente na Wake através do contrato da Social para a [NOME DO CLIENTE]"*
3. Aguardar a formalização com assinatura do aditivo

Qualquer ticket aberto para essa finalidade será redirecionado para esse fluxo de e-mail.

---

## Grupo: Tuna & Wake
**JID:** `120363403383828730@g.us`

---

### HTO-013 — Como funciona o processo de homologação de conector de pagamento Wake

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Tuna & Wake |
| **Data** | 2025-10-08 |
| **Categoria** | `how_to` |
| **Tags** | `homologação`, `audience-solutions` |

O processo de homologação de conector de pagamento na Wake segue as seguintes etapas:

1. **Desenvolvimento da integração:** Parceiro implementa a integração técnica (conector)
2. **Preenchimento da planilha de homologação:** Parceiro preenche a planilha fornecida pela Wake com especificações técnicas e casos de teste
3. **Testes integrados:** Wake e parceiro executam testes em ambiente staging com lojistas beta
4. **Aprovação:** Squad de pagamentos Wake avalia e aprova a integração
5. **Liberação para produção:** Conector disponibilizado para o primeiro cliente em produção

**Pré-requisito:** O contrato de parceria deve estar assinado antes do início da homologação formal.

**Planilha de referência:** https://docs.google.com/spreadsheets/d/1jYSsktOj2SnC_6g1w6QclYf8DP67p1GjJTiNqSx1k0g/edit

**Suporte durante homologação:** Canal Slack dedicado com times técnicos de ambos os lados.

---

### TRB-003 — Problema de cache em rota de autenticação no ambiente de integração

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Tuna & Wake |
| **Data** | 2025-10-24 |
| **Categoria** | `troubleshooting` |
| **Tags** | `integração`, `audience-solutions`, `audience-cs` |

**Sintomas:** Após atualização de uma rota de autenticação no ambiente de integração, a rota antiga continuou sendo chamada por até 1 hora, mesmo após a mudança.

**Causa raiz:** Cache do checkout/gateway que não é atualizado imediatamente após mudança de configuração.

**Resolução:**
- Para cache de checkout, a Wake não consegue forçar a limpeza externamente
- O parceiro deve rodar a **indexação e limpeza de cache pelo painel admin** da Wake
- Aguardar a propagação natural (pode levar entre 30 minutos e 1 hora)
- Se o problema persistir, abrir ticket de suporte: suporte@wake.tech ou WhatsApp (11) 2853-0696

---

### TRB-004 — Ambiente da Tuna retornando "acesso negado" sem motivo aparente

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Tuna & Wake |
| **Data** | 2025-10-23 |
| **Categoria** | `troubleshooting` |
| **Tags** | `suporte-escalação`, `audience-cs`, `audience-solutions` |

**Sintomas:** O ambiente de sandbox da Tuna na Wake (partnertuna.fbits.app) retornou erro de "acesso negado" sem alteração de configuração.

**Causa raiz:** O bloqueio é um mecanismo de segurança quando o horário do PC do usuário não coincide com o horário do servidor Wake.

**Resolução:**
1. Verificar se o horário do computador está sincronizado corretamente
2. Limpar o cache do navegador
3. Fazer logout e login novamente
4. Se o problema persistir após essas etapas, reportar ao time de parceria Wake informando o erro específico

---

### FAQ-007 — O suporte técnico Wake pode ser adicionado em um grupo de WhatsApp para acompanhamento de integração?

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Tuna & Wake |
| **Data** | 2025-10-24 |
| **Categoria** | `faq` |
| **Tags** | `suporte-escalação`, `audience-solutions`, `audience-cs` |

**Pergunta:** É possível criar um grupo de WhatsApp com o número de suporte da Wake para acompanhamento próximo de uma integração?

**Resposta:** Não. O número de suporte Wake (WhatsApp: (11) 2853-0696) é para atendimento individual e não pode ser adicionado a grupos.

**Alternativa recomendada:** Canal **Slack** dedicado com os times técnicos de ambos os lados. A Wake pode criar o canal no workspace LWSA-Tech mediante autorização da liderança. Esse é o canal oficial para integração técnica complexa.

---

### TRB-005 — Formulário de cartão não aparece no checkout do app (incidente em produção)

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Tuna & Wake |
| **Data** | 2026-03-12 |
| **Categoria** | `troubleshooting` |
| **Tags** | `integração`, `audience-cs`, `audience-solutions` |

**Sintomas:** A loja Ugg (Grupo Aste) ficou sem apresentar o formulário de cartão no checkout a partir das 17h, fechando pedidos sem capturar dados de cartão.

**Causa raiz:** Problema no lado do conector da Tuna — identificado pelo CTO da Tuna durante a escalada.

**Resolução:**
1. Acionamento urgente via WhatsApp ao parceiro de parcerias Wake
2. Escalamento interno para o time técnico Wake via Slack: https://lwsa-tech.slack.com/archives/C09QWKZ9U9F
3. O CTO da Tuna verificou o problema no lado do conector
4. **Problema resolvido em ~1 hora**

Lição aprendida: para incidentes críticos em produção (loja sem vender), acionar simultaneamente WhatsApp do parceiro de parcerias Wake, canal Slack técnico dedicado e o time técnico do parceiro (Tuna) em paralelo.

---

### BPR-006 — Como estruturar o desenvolvimento de integração de pagamento com múltiplos CDs (split de pedido)

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Tuna & Wake |
| **Data** | 2025-11-05 |
| **Categoria** | `best_practices` |
| **Tags** | `integração`, `audience-solutions` |

**Recomendação:** Ao desenvolver integração de pagamento que suporte múltiplos centros de distribuição (CDs) com split de pedido, fornecer payload de exemplo que cubra cenários reais de uso. O payload deve incluir obrigatoriamente:
- Estrutura correta quando um pedido tem produtos de CDs **diferentes** (o item mais crítico)
- Valores reais de frete e outros custos (não zerados)
- Campos `pedido`, `chave`, `idTransacaoGateway`, `id`, `statusId`
- Suporte a Boleto (com `linhadigitavel` e `linkPagamento`) e PIX (com QR code e data de expiração)

**Modelo de response Wake para gateway de pagamento:**
```json
{
  "pedido": [id],
  "chave": "[uuid]",
  "idTransacaoGateway": [id],
  "id": [id],
  "statusId": [1=aprovado, 4=outros],
  "message": "...",
  "boleto": { "linhadigitavel": "...", "linkPagamento": "..." },
  "pix": { "pixQrCodeUrl": "...", "pixQrCode": "...", "dataExpiracaoQrCode": "..." }
}
```

**Rationale:** Um payload genérico com dados repetidos (todos os itens no mesmo CD, valores zerados) gera bloqueio no desenvolvimento do parceiro, que não consegue testar os cenários reais que aparecerão em produção (split com CDs diferentes, frete real, modalidades Boleto/PIX). Documentar cenários reais previne retrabalho e bugs em produção.

---

### HTO-014 — Como realizar de-para de conectores do parceiro tech com funcionalidades Wake

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Tuna & Wake |
| **Data** | 2026-02-26 |
| **Categoria** | `how_to` |
| **Tags** | `integração`, `audience-solutions` |

Para mapear quais conectores do parceiro tech complementam as funcionalidades Wake, o processo recomendado é:

1. O parceiro tech disponibiliza uma **apresentação completa** com todos os conectores, meios de pagamento e anti-fraudes suportados
2. O time de produto Wake cruza essas informações com a **Tabela de Funcionalidades** da plataforma
3. Os itens validados nas homologações vão compondo a lista oficial de recursos suportados

A apresentação da Tuna com meios de pagamento e anti-fraudes serviu de insumo para atualização da tabela de funcionalidades Wake para 2026.

---

*Fim do documento kb-parceiros-wave6c.md.*
