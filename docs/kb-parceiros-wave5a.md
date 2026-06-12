# KB Extract — Wave 5A (SAVVI, Dotkom, A&EIGHT, Omni360)
> Extracted: 2026-06-04

---

## Grupo: SAVVI & Wake

**JID:** `120363420428455192@g.us`

---

### FAQ-1 — Comissionamento após rescisão de contrato de parceria

| Campo | Valor |
|---|---|
| **Grupo/Chat** | SAVVI & Wake |
| **JID** | `120363420428455192@g.us` |
| **Data** | 2025-09-05 |
| **Categoria** | `faq` |
| **Tags** | `comissionamento`, `rescisão`, `contrato`, `parceria`, `go-live` |

**Pergunta:** Caso o contrato de parceria seja rescindido, o comissionamento continua sendo pago?

**Resposta:** Sim. O comissionamento continua sendo pago por um período de **1 ano após o Go Live** do projeto indicado. Após esse prazo, o pagamento é encerrado. A Wake também aceita negociar cláusulas sobre prazo de rescisão — o prazo de 180 dias foi acordado como padrão neste caso.

**Contexto adicional:** A questão surgiu durante a negociação do contrato de adesão com a Savvi (plano I, R$1.500/mês). A agência parceira propôs uma redação que preservasse o comissionamento mesmo em rescisões onde o atraso fosse causado pelo cliente ou por circunstâncias alheias à agência.

---

### FAQ-2 — Como registrar uma indicação de cliente no Salesforce para garantir o comissionamento no estande de evento

| Campo | Valor |
|---|---|
| **Grupo/Chat** | SAVVI & Wake |
| **JID** | `120363420428455192@g.us` |
| **Data** | 2025-09-09 |
| **Categoria** | `faq` |
| **Tags** | `salesforce`, `lead`, `evento`, `fórum`, `comissionamento`, `SF` |

**Pergunta:** Se um cliente da agência bipar no stand da Wake em um evento sem estar cadastrado no Salesforce com a agência como origem, a indicação conta para o comissionamento da agência?

**Resposta:** Não. Se o cliente ainda não estiver registrado no Salesforce com a agência como "PARCEIRO ORIGEM" antes do evento, a fonte de origem será registrada como "Marketing", e não será possível creditar a indicação à agência. Por isso, é fundamental cadastrar o Lead e abrir a Oportunidade no SF com o nome da agência como origem **antes** de levar o cliente ao evento.

---

### HOW_TO-1 — Como cadastrar uma indicação de cliente no Salesforce (SF)

| Campo | Valor |
|---|---|
| **Grupo/Chat** | SAVVI & Wake |
| **JID** | `120363420428455192@g.us` |
| **Data** | 2025-09-09 |
| **Categoria** | `how_to` |
| **Tags** | `salesforce`, `lead`, `oportunidade`, `cadastro`, `parceiro-origem` |

**Processo:**

1. Acessar o Salesforce Community (SF).
2. Ir na aba **Leads** e clicar em **Criar**.
3. Preencher os dados do cliente indicado.
4. Garantir que o campo **PARCEIRO ORIGEM** esteja preenchido com o nome da agência.
5. Criar também uma **Oportunidade** vinculada ao Lead com a agência como origem.

**Importante:** Qualquer oportunidade sem a agência registrada como origem antes do contato com o time Wake será atribuída a Marketing e não gerará comissionamento para a agência.

---

### TROUBLESHOOTING-1 — Erro no Salesforce ao criar lead: "ID do proprietário não pode ficar em branco"

| Campo | Valor |
|---|---|
| **Grupo/Chat** | SAVVI & Wake |
| **JID** | `120363420428455192@g.us` |
| **Data** | 2025-06-25 |
| **Categoria** | `troubleshooting` |
| **Tags** | `salesforce`, `bug`, `erro`, `lead`, `SF`, `community` |

**Problema:** Ao tentar criar um Lead no Salesforce Community, aparece o erro:

> *"Não é possível salvar esse registro porque o processo '{Slack} {Community} {leads} redirecionar para o dono da conta quando for indicação de parceiros' falhou. ID do proprietário: proprietário não pode ficar em branco."*

**Causa:** Bug no fluxo de automação do Salesforce Community que tentava atribuir o registro a um proprietário com base em regras de parceria, mas o campo ficava em branco por falha no processo.

**Resolução:** O problema foi reportado ao fornecedor do Community. Lançamento de correções massivas foi previsto para os meses seguintes. Como solução temporária, as indicações devem ser enviadas diretamente à gestora de parcerias Wake via WhatsApp, que registra no sistema manualmente.

---

### BEST_PRACTICES-1 — Briefing estruturado de qualificação de lead para agências parceiras

| Campo | Valor |
|---|---|
| **Grupo/Chat** | SAVVI & Wake |
| **JID** | `120363420428455192@g.us` |
| **Data** | 2026-04-29 |
| **Categoria** | `best_practices` |
| **Tags** | `lead`, `qualificação`, `briefing`, `SDR`, `BANT`, `Gemini` |

A Wake passou a compartilhar briefings de leads qualificados com estrutura padronizada para as agências parceiras, incluindo:

- **Contexto do negócio**: segmento, canais, faturamento online e offline, ERP, plataforma atual.
- **Dores mapeadas**: problemas com checkout, conversão, integração de estoque, logística.
- **Requisitos técnicos**: integrações necessárias, funcionalidades demandadas.
- **Perfil do stakeholder**: cargo, perfil de decisão, forma de comunicação recomendada.
- **O que se espera da agência**: itens específicos que a agência deve endereçar na proposta.
- **BANT Summary**: Budget, Authority, Need, Timeline.
- **Notas do Gemini**: transcrições e insights gerados automaticamente a partir das reuniões de qualificação.

**Recomendação:** Agências devem ler o briefing completo antes do primeiro contato com o prospect e abordar os pontos de dor mapeados diretamente na proposta.

---

### BEST_PRACTICES-2 — Estratégia de abordagem para prospects com orçamento limitado

| Campo | Valor |
|---|---|
| **Grupo/Chat** | SAVVI & Wake |
| **JID** | `120363420428455192@g.us` |
| **Data** | 2026-04-29 |
| **Categoria** | `best_practices` |
| **Tags** | `proposta`, `objeção-preço`, `estratégia-comercial`, `prospect`, `decisor` |

**Situação:** Prospect (HDM Beauty) com orçamento limitado (~R$150-200/mês) e restrições para aprovar a solução Wake (~R$2.700/mês).

**Abordagem recomendada pela Wake:**
1. Explicar que a Wake não é a solução mais barata, mas que o valor se justifica pelo conjunto técnico e de conversão.
2. Focar na validação técnica do projeto com o contato operacional (não o decisor final).
3. Após a validação técnica, **ajudar o contato a apresentar a proposta para o decisor interno** (dono/diretor).
4. Não tentar fechar diretamente com quem não tem alçada de decisão — preparar o caminho para a conversa com o decisor.

---

### BEST_PRACTICES-3 — Seleção de parceiros de integração ERP para projetos Wake

| Campo | Valor |
|---|---|
| **Grupo/Chat** | SAVVI & Wake |
| **JID** | `120363420428455192@g.us` |
| **Data** | 2026-06-02 |
| **Categoria** | `best_practices` |
| **Tags** | `integração`, `ERP`, `middleware`, `TOTVS`, `parceiros-tech`, `ecossistema` |

A Wake mantém um ecossistema de parceiros de integração (middleware/iPaaS). As principais opções são:

| Parceiro | Fortaleza |
|---|---|
| **Argento Digital** | Hub de integrações, consultoria, projetos complexos com ERPs Linx, Senior, Omie, Abacos |
| **Integra.do (Kore Hub)** | TOTVS Protheus, Winthor, RM, SAP, Omie, Sankhya, Bling, Tiny |
| **Iota.Hub** | iPaaS para fluxos complexos, SAP, sistemas legados, B2B e B2C |
| **EBI (EBInterchange)** | TOTVS Protheus, Winthor, Linx Microvix, SAP |
| **Intergate** | Especialista em SAP Business One e TOTVS Protheus |
| **Floui** | Low-code, SAP S/4HANA, TOTVS Protheus, Sankhya |
| **Jitterbit** | iPaaS global, low-code com IA, conectores pré-construídos |
| **Vtrina** | Hub de marketplace + middleware para integrações complexas |
| **Leanwork Group** | Desenvolvimento sob demanda, integradores customizados |

**Recomendação para TOTVS RM:** Intergate é a mais especializada. Integra.do lista TOTVS RM no site, mas pode não ter a integração pronta — validar antes de indicar.

**Critério de seleção:** Avaliar o ERP do cliente + complexidade logística (múltiplos CDs, lojas físicas) + necessidade de consultoria vs. plataforma self-service.

---

## Grupo: Parcerias Dotkom & Wake

**JID:** `120363279257284512@g.us`

---

### FAQ-3 — Como funcionam as informações de contato nos leads enviados pelo Salesforce

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Parcerias Dotkom & Wake |
| **JID** | `120363279257284512@g.us` |
| **Data** | 2025-07-14 |
| **Categoria** | `faq` |
| **Tags** | `salesforce`, `leads`, `contato`, `executivo`, `SF` |

**Pergunta:** Os dados de contato do cliente (telefone, e-mail) chegam via Salesforce para as agências?

**Resposta:** Não. A partir de determinado momento, a informação do contato do cliente deixou de ser enviada diretamente no SF. O que as agências recebem é o **executivo Wake responsável** pela oportunidade. A agência deve acionar o executivo Wake para obter mais informações e entender se já é possível entrar na negociação com o cliente.

---

### FAQ-4 — Regras do processo de Storefront para clientes existentes

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Parcerias Dotkom & Wake |
| **JID** | `120363279257284512@g.us` |
| **Data** | 2025-11-04 |
| **Categoria** | `faq` |
| **Tags** | `storefront`, `CS`, `clientes-existentes`, `processo`, `oportunidade` |

**Pergunta:** Existe uma previsão de quando as oportunidades de Storefront serão direcionadas às agências?

**Resposta:** Não há uma previsão fechada. O processo segue a **demanda do time de Customer Success (CS)** da Wake. Quando o CS identifica que um cliente existente pode se beneficiar de uma atualização de Storefront e o cliente demonstra interesse em ouvir sobre a atualização, as agências participantes do programa são acionadas. Não há fila pré-definida.

---

### FAQ-5 — Como funciona a Campanha de Vendas para agências parceiras

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Parcerias Dotkom & Wake |
| **JID** | `120363279257284512@g.us` |
| **Data** | 2025-07-22 |
| **Categoria** | `faq` |
| **Tags** | `campanha`, `comissionamento`, `vendas`, `certificação`, `bônus` |

**Pergunta:** Como funciona a Campanha de Vendas para agências?

**Resposta:** A campanha oferece **R$500,00 por contrato fechado**, independentemente da faixa de faturamento do cliente, para cada oportunidade que a agência trouxer e se converter em contrato. A adesão à campanha é feita através de link específico incluído no material da campanha. Há também um componente de **Certificação Wake**, que é requisito para participação e pode ser acessada via formulário de inscrição.

---

### TROUBLESHOOTING-2 — Cliente não responde após indicação de SDR: o que fazer

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Parcerias Dotkom & Wake |
| **JID** | `120363279257284512@g.us` |
| **Data** | 2025-06-27 |
| **Categoria** | `troubleshooting` |
| **Tags** | `lead`, `sem-resposta`, `SDR`, `qualificação`, `desqualificação` |

**Problema:** A agência tentou contato com o lead por WhatsApp, e-mail e telefone sem retorno.

**Processo recomendado:**
1. Tentar contato por múltiplos canais (WhatsApp, e-mail, telefone).
2. Acionar a gestora de parcerias Wake para que o **executivo Wake responsável** tente contato diretamente.
3. Se após os esforços combinados não houver retorno, o lead deve ser **desqualificado** e marcado como "Lead nunca respondeu via WhatsApp e tel, desqualificado."
4. A oportunidade pode ser reaberta no futuro caso haja nova manifestação de interesse.

---

### HOW_TO-2 — Como conduzir o acompanhamento de forecast de oportunidades com a Wake

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Parcerias Dotkom & Wake |
| **JID** | `120363279257284512@g.us` |
| **Data** | 2025-07-23 |
| **Categoria** | `how_to` |
| **Tags** | `forecast`, `pipeline`, `oportunidade`, `funil`, `acompanhamento` |

O acompanhamento de pipeline (forecast) entre Wake e agência é feito em reuniões periódicas (quinzenal ou mensal), com o seguinte formato:

**Estrutura do relatório de forecast:**
- Nome da oportunidade
- Executivo Wake responsável
- Status atual: Apresentar / Qualificação / Negociação / Proposta / **Perdido Wake e Agência**
- Motivo de perda (quando aplicável): cliente não respondeu, foi para concorrente, desistiu do projeto
- Próximos passos: retomar em data futura / pausado / desqualificado

**Exemplo de registro de perda:**
- *"Antares Imagem - Cliente apontou que vai para a Tray"*
- *"Usee Brasil - Lead não respondeu e nem atendeu, desqualificado"*
- *"Goldko - Informou à Dotkom que não deseja Wake"*

---

### BEST_PRACTICES-4 — Como lidar com cliente que busca segunda agência durante negociação

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Parcerias Dotkom & Wake |
| **JID** | `120363279257284512@g.us` |
| **Data** | 2025-12-29 |
| **Categoria** | `best_practices` |
| **Tags** | `concorrência`, `agência`, `negociação`, `cliente-enterprise`, `maturidade` |

**Situação:** Cliente em negociação avançada com a agência solicita outra agência para comparação ou avança com concorrente sem aviso.

**Análise da Wake:**
- O pedido por uma segunda agência frequentemente indica **necessidade de validação e redução de risco percebido**, não falta de confiança na agência original.
- Pode refletir insegurança sobre a robustez da plataforma, prazo de implantação e curva de aprendizagem.

**Recomendações para a agência:**
1. **Fortalecer o discurso comercial** conectando claramente a solução Wake à realidade do cliente.
2. Reforçar expectativas realistas de implantação e evolução.
3. Demonstrar que a parceria Agência + Wake traz **proximidade e continuidade**, evitando trocas futuras.
4. Em casos Enterprise, solicitar apoio do executivo Wake para participar da reunião com o cliente — isso aumenta a confiança e melhora a taxa de conversão.

**Nota:** Em casos Enterprise, a comissão da agência pode ser reduzida para 1/3 quando há participação direta da Wake na venda. Avaliar esse trade-off antes de acionar o executivo.

---

### BEST_PRACTICES-5 — Perfil de lead ideal para Wake: critérios de qualificação por plataforma de origem

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Parcerias Dotkom & Wake |
| **JID** | `120363279257284512@g.us` |
| **Data** | 2026-01-20 |
| **Categoria** | `best_practices` |
| **Tags** | `qualificação`, `plataforma-origem`, `migração`, `ERP`, `GMV` |

Padrão observado nos leads qualificados enviados pela Wake via canal (Canais):

| Campo | O que informar |
|---|---|
| Plataforma atual | Shopify, VTEX, Tray, Wbuy, Magento, plataforma própria, etc. |
| ERP | Bling, Tiny, TOTVS, Sankhya, Omie, etc. |
| GMV online | Faturamento digital mensal/anual |
| GMV offline | Faturamento físico (quando disponível) |
| Canais de venda | E-commerce, marketplace, loja física, WhatsApp |
| Dor principal | Integração ruim, baixa conversão, checkout instável, falta de funcionalidades |
| Urgência | Prazo esperado para migração / go-live |
| Perfil do decisor | Proprietário, gestor de e-commerce, CTO, diretor |
| Concorrentes | VTEX, Nuvemshop, Tray, Magazord, UAPPI |

---

## Grupo: Wake & A&EIGHT (B8One)

**JID:** `120363399045808727@g.us`

---

### TROUBLESHOOTING-3 — FeedWake com erro em produtos com HTML na descrição

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Wake & A&EIGHT (B8One) |
| **JID** | `120363399045808727@g.us` |
| **Data** | 2025-07-08 |
| **Categoria** | `troubleshooting` |
| **Tags** | `FeedWake`, `feed`, `produto`, `HTML`, `bug`, `omnichannel`, `ticket` |

**Problema:** FeedWake (Balaroti #234109 e Guaporé #234732) apresentando erros, com produtos não aparecendo corretamente no feed de omnichannel.

**Causa identificada:** Produtos com **descrição em HTML** contendo caracteres especiais quebram os dados do feed. Esse problema é específico de cada loja (não é um bug global da plataforma), mas pode ocorrer em qualquer loja com esse tipo de conteúdo.

**Resolução:**
- O time de produto da Wake tratou o caso individualmente.
- Foram orientadas soluções para remoção/adequação do HTML nas descrições.
- Produtos **inativos no Omnichannel** por falta de estoque são automaticamente desativados — verificar se há ruptura de estoque antes de reportar como bug de feed.
- Abrir ticket no suporte Wake: https://wakecommerce.readme.io/docs/canais-de-atendimento-para-parceiros

**Lição aprendida:** Quando o go-live é feito, o feed pode estar funcionando na virada, mas erros surgem com o tempo caso haja conteúdo HTML mal formatado no catálogo. Recomenda-se auditoria das descrições de produto no pré-go-live.

---

### FAQ-6 — Como acionar escalation de suporte para clientes insatisfeitos

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Wake & A&EIGHT (B8One) |
| **JID** | `120363399045808727@g.us` |
| **Data** | 2025-07-08 |
| **Categoria** | `faq` |
| **Tags** | `suporte`, `escalation`, `ticket`, `CS`, `cliente-insatisfeito` |

**Pergunta:** Como acionar prioridade para um ticket quando o cliente está muito insatisfeito e não quer aguardar a agenda normal de atendimento?

**Processo:**
1. Acionar a gestora de parcerias Wake via WhatsApp solicitando **escalation** do ticket.
2. Informar o número do ticket (ex: #234732) e o nome do analista responsável (visível no próprio ticket).
3. A gestora aciona internamente para antecipar o atendimento.
4. Em casos críticos com cliente ameaçando saída de plataforma, o time de CS pode ser acionado diretamente.

**Limitação:** Nem sempre é possível antecipar agendas por restrições de disponibilidade interna. O CS pode manter a reunião original mas incluir um ponto de atenção sobre a urgência.

---

### BEST_PRACTICES-6 — Alinhamento comercial agência-Wake: Role Play e verticais prioritárias

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Wake & A&EIGHT (B8One) |
| **JID** | `120363399045808727@g.us` |
| **Data** | 2025-07-01 |
| **Categoria** | `best_practices` |
| **Tags** | `onboarding`, `role-play`, `verticais`, `alinhamento-comercial`, `executivos` |

**Processo recomendado para estreitar operação comercial entre agência e Wake:**

1. **Wake Day com Role Play**: reunião presencial no escritório Wake (Creators) com 1-2 executivos Wake + todos os executivos da agência. Objetivo: apresentação mútua, simulação de pitches de venda, alinhamento de discurso.
2. **Elencar verticais prioritárias**: B2B, Moda, Eletrônicos — focar a proximidade nos executivos dessas verticais primeiro, antes de expandir para todas.
3. **Certificação Wake**: obrigatória para todos os executivos da agência. Turmas abrem toda segunda-feira. Prazo de 30 dias após o acesso para conclusão. Link: https://forms.gle/AA6NR6vL3uEJXhDu7
4. **Guia de conhecimento interno**: usar o material do Wake Education como base para capacitação dos executivos da agência sobre a plataforma.

---

### HOW_TO-3 — Como registrar lead trazido de evento externo (ex: Fórum ECBR)

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Wake & A&EIGHT (B8One) |
| **JID** | `120363399045808727@g.us` |
| **Data** | 2025-08-04 |
| **Categoria** | `how_to` |
| **Tags** | `lead`, `evento`, `fórum`, `SF`, `cadastro`, `CNPJ` |

**Processo para registrar um lead coletado em evento externo:**

1. Acionar a gestora de parcerias Wake via WhatsApp informando:
   - Nome do cliente
   - CNPJ
   - Nome e telefone do contato
   - E-mail do contato
2. A gestora registra no Salesforce com a agência como origem.
3. A agência informa o estágio da conversa (ex: "Wake já vendida, aguardando aprovação do decisor").
4. A gestora abre a Oportunidade e aciona o executivo Wake do segmento correspondente.

**Importante:** Leads coletados em eventos sem registro prévio no SF não contam para o comissionamento se o cliente também for abordado pelo time de Marketing Wake no mesmo evento.

---

### FAQ-7 — Certificação Wake: requisito do Wake Score e como se inscrever

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Wake & A&EIGHT (B8One) |
| **JID** | `120363399045808727@g.us` |
| **Data** | 2026-01-27 |
| **Categoria** | `faq` |
| **Tags** | `certificação`, `wake-score`, `treinamento`, `obrigatório` |

**A Certificação Wake é requisito obrigatório do Wake Score para agências parceiras.**

- Turmas abrem toda **segunda-feira**.
- Após inscrição, o acesso é liberado na **segunda-feira seguinte**.
- Prazo para conclusão: **30 dias** após a liberação do acesso.
- Inscrição: https://forms.gle/AA6NR6vL3uEJXhDu7

Agências que não completarem a certificação terão impacto negativo no Wake Score, o que pode afetar o volume e qualidade das oportunidades recebidas.

---

### BEST_PRACTICES-7 — Uso de recursos nativos de IA da Wake em projetos de agência

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Wake & A&EIGHT (B8One) |
| **JID** | `120363399045808727@g.us` |
| **Data** | 2026-05-25 |
| **Categoria** | `best_practices` |
| **Tags** | `IA`, `MCP`, `LLM`, `Wake-Dev-Toolkit`, `plugin`, `agente-conversacional` |

A Wake disponibiliza recursos nativos de IA para uso em projetos de agências:

- **Wake MCP** (Model Context Protocol): usado por agências para desenvolver **agentes conversacionais de venda** integrados à plataforma.
- **Wake Dev Tool Kit (Wake Plugin)**: repositório de plugin de IA disponível em https://github.com/wake-engineering/ai-plugin
- **LLM**: integração nativa com modelos de linguagem para uso em automações e experiências de compra.

**Caso de uso identificado:** Agência B8One testou o MCP para construir um agente conversacional de venda integrado à loja Wake.

**Recomendação:** Agências que utilizem esses recursos em projetos devem sinalizar à Wake para documentação de casos de uso para marketing.

---

## Grupo: Wake + Omni360 + Affys

**JID:** `120363337557834516@g.us`

---

### HOW_TO-4 — Como estruturar proposta para cliente com operação B2B + B2C na Wake

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Wake + Omni360 + Affys |
| **JID** | `120363337557834516@g.us` |
| **Data** | 2026-04-14 |
| **Categoria** | `how_to` |
| **Tags** | `B2B`, `B2C`, `proposta`, `escopo`, `cliente`, `separação-projetos` |

**Situação:** Cliente interessado em e-commerce B2C com possibilidade futura de B2B.

**Abordagem recomendada:**
1. Priorizar o B2C como escopo principal da proposta inicial.
2. Tratar o B2B como um **segundo projeto**, com estrutura apartada (CNPJ, regras comerciais, catálogo distintos).
3. Incluir na proposta ambos os escopos para tentar fechar tudo em uma única negociação, apresentando o B2B como expansão natural.
4. Alinhar com o executivo Wake se o produto Wake B2B se aplica (ex: Wake B2B tem mensalidade própria, com custo de ~R$4.900/mês até 500 pedidos/mês).

**Nota:** Em projetos B2B onde o site não será público e servirá apenas para representantes emitirem pedidos, a agência ainda precisa desenvolver o front-end normalmente. O plano B2B Wake cobre a plataforma, não o desenvolvimento.

---

### FAQ-8 — Wake Gateway: adquirentes disponíveis e recursos de pagamento

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Wake + Omni360 + Affys |
| **JID** | `120363337557834516@g.us` |
| **Data** | 2026-02-04 |
| **Categoria** | `faq` |
| **Tags** | `gateway`, `pagamento`, `adquirente`, `Wake-Gateway`, `checkout`, `2-cartões`, `PIX` |

**Pergunta:** Quais adquirentes estão disponíveis no Wake Gateway? Todos suportam recursos como pagamento com 2 cartões, cartão + PIX etc.?

**Resposta:** O Wake Gateway possui múltiplos adquirentes disponíveis. Os novos contratos de plataforma já incluem o Gateway integrado. Para detalhes sobre quais adquirentes suportam recursos específicos (pagamento com 2 cartões, cartão+PIX), recomenda-se acionar o time comercial da Wake para uma agenda específica sobre o produto.

**Contato para agenda de Wake Gateway:** Acionar o responsável pelo produto via gestora de parcerias.

---

### TROUBLESHOOTING-4 — Dificuldade de integração com Meta (XML/Feed) após go-live

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Wake + Omni360 + Affys |
| **JID** | `120363337557834516@g.us` |
| **Data** | 2026-05-15 |
| **Categoria** | `troubleshooting` |
| **Tags** | `Meta`, `XML`, `feed`, `trackeamento`, `performance`, `go-live` |

**Problema:** Agência de performance do cliente Dumond teve dificuldades para subir XML para o Meta Ads após o go-live. O cliente ficou 2 dias sem vender pelo Meta.

**Causa:** Dificuldades de configuração de trackeamento (Pixel, feed XML, eventos de conversão) em ambiente Wake por parte da agência de performance.

**Resolução:**
1. A agência (B8One) assumiu as configurações de trackeamento e integração com Meta **em lugar da agência de performance** para resolver rapidamente.
2. O chamado foi aberto na Wake para suporte técnico.
3. A Wake fez os ajustes necessários no feed.

**Lição aprendida:** Agências de performance que não conhecem bem a plataforma Wake tendem a ter dificuldades com configurações de trackeamento. A agência implementadora deve estar preparada para dar suporte nessa etapa ou cobrir essa configuração no escopo de projeto.

---

### BEST_PRACTICES-8 — Campanha de vendas interna da agência com CS: estrutura e regras

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Wake + Omni360 + Affys |
| **JID** | `120363337557834516@g.us` |
| **Data** | 2026-02-09 |
| **Categoria** | `best_practices` |
| **Tags** | `campanha-vendas`, `CS`, `ranking`, `engajamento`, `incentivo`, `Affys` |

**Contexto:** A Affys estruturou uma campanha de vendas para o time de CS com premiação em dinheiro para contratos fechados.

**Pontos negociados com a Wake:**

1. **Duração:** Estendida de 3 para **6 meses** para dar mais fôlego ao amadurecimento das vendas.
2. **Regra do ranking:** O ranking de melhores indicações é ativado a partir de **3 fechamentos do time como um todo** (não individual). A Wake sugeriu que o ranking fosse ativado a partir do primeiro fechamento para não desmotivar no início — a agência manteve a regra de 3 para garantir qualidade das indicações.

**Recomendação:** Campanhas de vendas em agências com foco em indicações Wake devem ter duração mínima de 6 meses e regras de ranking que não desmotivem os primeiros esforços. Envolver a Wake na validação das regras evita conflitos operacionais.

---

### HOW_TO-5 — Dados do BTG360 Wake: disponibilidade para integração com sistemas externos

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Wake + Omni360 + Affys |
| **JID** | `120363337557834516@g.us` |
| **Data** | 2026-02-20 |
| **Categoria** | `how_to` |
| **Tags** | `BTG360`, `dados`, `integração`, `analytics`, `parceiro` |

**Pergunta:** É possível puxar dados do BTG360 da Wake para alimentar um sistema de tracking de interações de clientes em páginas de parceiros?

**Status registrado:** A questão foi levantada pela Affys que precisava de dados de interações dos clientes na página do parceiro para um mapeamento de comportamento. A consulta foi feita ao time Wake para verificar disponibilidade da API/dados do BTG360.

**Recomendação:** Acionar o time técnico Wake para verificar disponibilidade e escopo dos dados do BTG360 antes de incluir no roadmap do produto.

---

### BEST_PRACTICES-9 — Estratégia de ecossistema para clientes com múltiplos canais (Wake + Omni + Affys)

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Wake + Omni360 + Affys |
| **JID** | `120363337557834516@g.us` |
| **Data** | 2025-10-27 |
| **Categoria** | `best_practices` |
| **Tags** | `ecossistema`, `omnichannel`, `afiliados`, `comunidade`, `B2B`, `Yale`, `estratégia` |

**Caso:** Reunião de qualificação com cliente Keypartners (piscinas/spas).

**Estratégia de ecossistema proposta:**
1. **Comunidade de instaladores** ("piscineiros"): criar uma comunidade de prestadores de serviço de instalação, usando o modelo da Yale (como a Yale faz com distribuidores de fechaduras).
2. **Cadastro de prestadores como afiliados** via Affys: monetizar a comunidade e gerar canal de indicação com rastreabilidade.
3. **Solução unificada**: Wake Commerce (loja) + Omni360 (omnichannel/CRM) + Affys (afiliados/programa de fidelidade) formam um ecossistema completo para o cliente.

**Mensagem-chave para venda do ecossistema:** Cada solução fala com todas as pontas (cliente final, instaladores, representantes), criando um flywheel de vendas e serviços.

---

### FAQ-9 — Canais oficiais de suporte técnico para parceiros Wake

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Wake + Omni360 + Affys |
| **JID** | `120363337557834516@g.us` |
| **Data** | 2026-02-13 |
| **Categoria** | `faq` |
| **Tags** | `suporte`, `canal-oficial`, `parceiro`, `suporte-tecnico`, `atendimento` |

**Canal oficial de suporte técnico para parceiros Wake:**

- E-mail: suporte@wake.tech
- Documentação e canais: https://wakecommerce.readme.io/docs/canais-de-atendimento-para-parceiros

**Contatos do time de Parcerias Wake (referência 2026):**

| Papel | Responsável | Telefone |
|---|---|---|
| Account Manager – Agências | Ana Vicente | +55 14 99770-3611 |
| Account Manager – Checkout/Pagamento/Frete | Gabriel Cotrim | +55 11 96925-1769 |
| AWS Alliance / Tech Partnerships | Rafael Correia | +55 11 93800-6017 |
| Partner Development | Letícia Cena | +55 11 97433-3718 |

**Eventos recorrentes:**
- **Wake Partner Day**: toda quarta-feira no escritório LWSA (Rua Itapaiúna, 2434, São Paulo). Parceiros podem trabalhar presencialmente — enviar nome, e-mail e documento para liberar acesso ao Wi-Fi.

---
