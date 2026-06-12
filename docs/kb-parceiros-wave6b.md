# KB Extract — Wave 6B (Serasa, Estoca, Konfidency, CRMBonus)
> Extracted: 2026-06-04

---

## Grupo: Serasa & Wake

**JID:** `120363401681102953@g.us`

---

### HOW_TO-1 — Fluxo de integração do Serasa Pass na jornada de compra Wake

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Serasa & Wake |
| **JID** | `120363401681102953@g.us` |
| **Data** | 2025-08-18 |
| **Categoria** | `how_to` |
| **Tags** | `serasa-pass`, `integração`, `checkout`, `autenticação` |

A integração do Serasa Pass na jornada de compra Wake foi definida em três fases:

- **Fase 1 (MVP):** Fluxo de preenchimento de dados com autenticação via aplicativo Serasa — sem cartão físico e sem emissão de link de biometria.
- **Fase 2:** Fluxo de biometria sem o app (link de biometria disponível como opção standalone).
- **V2 (futuro):** Inclusão do fluxo com cartão.

A integração exigiu envolvimento de um parceiro integrador (Leanwork) por não ser configurável apenas via painel admin — necessita desenvolvimento técnico. O perfil ideal do e-commerce para adoção inicial é aquele com time técnico próprio e porte médio a grande.

---

### FAQ-1 — Webhook vs. Polling para dados do Serasa Pass

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Serasa & Wake |
| **JID** | `120363401681102953@g.us` |
| **Data** | 2025-08-19 |
| **Categoria** | `faq` |
| **Tags** | `serasa-pass`, `webhook`, `polling`, `API`, `técnico` |

**Pergunta:** Para buscar dados do cliente no Serasa Pass, é obrigatório usar Webhook ou é possível fazer polling periódico para verificar disponibilidade?

**Resposta:** Ambos os modos são possíveis. A escolha do caminho mais adequado depende do caso de uso:
- Para o evento inicial de verificação (checar se o dado está disponível), o polling pode ser usado.
- Para o fluxo completo de dados após confirmação, o Webhook é o direcionamento padrão.

O time do Serasa Pass deve ser consultado para validar qual abordagem específica atende melhor o fluxo da loja.

---

### TROUBLESHOOTING-1 — Dificuldade de conexão real via API do Serasa Pass

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Serasa & Wake |
| **JID** | `120363401681102953@g.us` |
| **Data** | 2025-09-05 |
| **Categoria** | `troubleshooting` |
| **Tags** | `serasa-pass`, `API`, `Leanwork`, `integração`, `plano-B` |

**Problema:** Durante o desenvolvimento da integração do Serasa Pass, o time da Leanwork encontrou dificuldades para estabelecer conexão real via API.

**Resolução:** Foi adotado um plano B com dados fictícios enquanto o problema era resolvido. A previsão foi: finalizar o fluxo via API ainda naquele dia e deixar apenas ajustes de layout para a semana seguinte.

**Lição:** Ao integrar com parceiros externos como o Serasa Pass via integradores como a Leanwork, é recomendável ter um ambiente de demonstração paralelo (dados fictícios) para garantir continuidade no desenvolvimento e nas apresentações.

---

### HOW_TO-2 — Processo de assinatura de contrato com a Serasa (cláusula de limitação de responsabilidade)

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Serasa & Wake |
| **JID** | `120363401681102953@g.us` |
| **Data** | 2026-03-24 |
| **Categoria** | `how_to` |
| **Tags** | `contrato`, `jurídico`, `Serasa`, `assinatura`, `limitação-de-responsabilidade` |

Um bloqueador recorrente na assinatura do contrato com a Serasa foi a cláusula de limitação de responsabilidade. O jurídico da Wake exige que o valor da cláusula seja determinado para prosseguir — não é possível assinar sem o valor e a cláusula não pode ser removida.

**Como resolver:**
1. O jurídico da Serasa deve propor um valor (pode ser simbólico: R$ 5k, R$ 10k, R$ 15k).
2. O jurídico da Wake aceita valores abaixo de R$ 50k.
3. Uma vez definido o valor, os demais pontos do contrato são tratados como detalhes menores e negociados diretamente entre os times.

**Resultado:** A Wake liberou a assinatura do contrato após o alinhamento desta cláusula em abril de 2026.

---

### FAQ-2 — Comprovação de poderes de representantes para assinatura contratual

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Serasa & Wake |
| **JID** | `120363401681102953@g.us` |
| **Data** | 2025-09-26 |
| **Categoria** | `faq` |
| **Tags** | `contrato`, `jurídico`, `representação-legal`, `documentos`, `Experian` |

**Pergunta:** O jurídico da Wake solicitou documento comprovando que os representantes da Serasa/Experian são diretora e presidente. Como obter?

**Resposta:** O site da Experian (seção de Relações com Investidores - RI) contém a estrutura do conselho e diretores. Caso seja necessária uma procuração formal, o ponto de contato interno da Experian deve ser acionado. O próprio time de relações internas pode agilizar o documento.

---

### BEST_PRACTICES-1 — Ativação do Serasa Pass no Wake Summit: estratégia de stand e palco

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Serasa & Wake |
| **JID** | `120363401681102953@g.us` |
| **Data** | 2025-09-05 |
| **Categoria** | `best_practices` |
| **Tags** | `eventos`, `wake-summit`, `ativação`, `stand`, `palco`, `parceiro` |

Para parceiros que participam de eventos Wake como patrocinadores, a combinação de ativações recomendada é:

- **Palco Make Together:** Sessão de 20 minutos para apresentar o produto/parceria (ex: "Futuro do checkout na Wake com Serasa Pass"). Ideal para educação e awareness.
- **Stand:** Demonstração ao vivo da demo técnica/funcional do produto integrado. O foco deve ser atrair leads ativos para assistir à demo presencialmente.
- **Programação de lead capture:** O CS Wake indica leads quentes previamente para o stand durante o evento.

**Importante:** O palco principal do evento (keynote) tem lineup fechado com antecedência — parceiros ingressantes no programa não devem contar com esse espaço para o primeiro evento.

---

### HOW_TO-3 — Engajamento de e-commerces Wake como primeiros clientes do Serasa Pass

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Serasa & Wake |
| **JID** | `120363401681102953@g.us` |
| **Data** | 2026-01-29 |
| **Categoria** | `how_to` |
| **Tags** | `serasa-pass`, `vendas`, `indicação`, `primeiros-clientes`, `CS` |

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

### HOW_TO-4 — Onboarding de parceiro de patrocínio no Wake Summit (cota Silver + Experiência)

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Estoca & Wake |
| **JID** | `120363420090271403@g.us` |
| **Data** | 2025-08-20 |
| **Categoria** | `how_to` |
| **Tags** | `wake-summit`, `patrocínio`, `onboarding`, `contrato`, `stand`, `experiência` |

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

### FAQ-3 — Diferença entre número de participantes e quantidade de brindes no Wake Summit

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Estoca & Wake |
| **JID** | `120363420090271403@g.us` |
| **Data** | 2025-09-02 |
| **Categoria** | `faq` |
| **Tags** | `wake-summit`, `brindes`, `participantes`, `estoca`, `experiência` |

**Pergunta:** Por que a proposta indica empacotamento de 1.200 brindes se o evento tem ~700 participantes?

**Resposta:** O número de 700 é a expectativa de presença de patrocinadores, clientes e prospects. O estoque da loja virtual do site possui capacidade para 1.200 brindes (número de inscritos), mas há uma quebra natural (pessoas que não comparecem, saem antes etc.). Como os brindes são diversificados, não é possível estimar previamente quais sobrarão — por isso, o volume total disponível é entregue para empacotamento. O parceiro pode solicitar atualização da proposta para 700 unidades se preferir trabalhar com o número estimado de presença efetiva.

---

### BEST_PRACTICES-2 — Follow-up pós-evento: como usar o mailing de participantes

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Estoca & Wake |
| **JID** | `120363420090271403@g.us` |
| **Data** | 2025-09-15 |
| **Categoria** | `best_practices` |
| **Tags** | `pós-evento`, `mailing`, `follow-up`, `leads`, `wake-summit` |

Após o Wake Summit, patrocinadores devem solicitar o mailing de participantes para iniciar follow-up o mais rápido possível.

**Processo:**
- Solicitar o mailing ao time de eventos Wake logo após o evento.
- O SLA para envio é de **até 5 dias úteis**.
- Priorizar o follow-up imediato para aproveitar o momento de alta receptividade dos leads que visitaram o stand.

**Recomendação:** Definir internamente antes do evento o responsável pelo follow-up e a cadência de abordagem para não perder o timing pós-evento.

---

### HOW_TO-5 — Processo para gravação de podcast no Wake Summit

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Estoca & Wake |
| **JID** | `120363420090271403@g.us` |
| **Data** | 2025-10-24 |
| **Categoria** | `how_to` |
| **Tags** | `podcast`, `wake-summit`, `conteúdo`, `divulgação`, `parceiro` |

Para parceiros que participaram do Wake Summit com gravação de podcast:

1. Os episódios são publicados em sequência — aproximadamente **1 por semana**.
2. A ordem de publicação considera Black Friday e prioridades de comunicação do calendário Wake.
3. O parceiro deve ser avisado com antecedência sobre a data de publicação para coordenar a divulgação nas próprias redes sociais.
4. Recomendável preparar a programação de divulgação do parceiro com base na previsão Wake.

---

### BEST_PRACTICES-3 — ICP e expectativas da parceria Estoca & Wake (2026)

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Estoca & Wake |
| **JID** | `120363420090271403@g.us` |
| **Data** | 2026-03-17 |
| **Categoria** | `best_practices` |
| **Tags** | `ICP`, `parceria`, `Estoca`, `segmentação`, `meta-2026`, `target-list` |

Perfil ideal de cliente (ICP) para a parceria Estoca & Wake definido em conjunto:

- **GMV:** entre R$ 2M e R$ 4M/mês
- **Quantidade de pedidos:** entre 5.000 e 10.000/mês
- **Faturamento digital:** entre R$ 1M e R$ 2M/mês
- **Segmentos prioritários:** Cosméticos e suplementos

**Target list 2026 da Estoca:** Nutrata, Loccitane, Sallve, Puravida, Guday, AURA Beauty, +Mu, Beyoung, Principia, DUX, Linus, BOB, Braé, Liquidz, HeyMu, Soldiers Nutrition.

**Expectativas da Estoca para 2026:**
- Construir canal de parcerias sólido como um dos pilares estratégicos do ano.
- A Wake foi vista como solução complementar com tecnologia de ponta para entregar proposta de valor conjunta ao cliente.

---

### HOW_TO-6 — Formalização de entrada no programa Elite Wake (parceiro tech)

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Estoca & Wake |
| **JID** | `120363420090271403@g.us` |
| **Data** | 2026-04-07 |
| **Categoria** | `how_to` |
| **Tags** | `contrato`, `programa-elite`, `onboarding`, `dados-cadastrais`, `parceiro-tech` |

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

### TROUBLESHOOTING-2 — Script da Konfidency não renderiza em todas as páginas do site Wake

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Konfidency & Wake |
| **JID** | `120363413122423249@g.us` |
| **Data** | 2025-11-07 |
| **Categoria** | `troubleshooting` |
| **Tags** | `script`, `renderização`, `CSP`, `whitelist`, `API`, `técnico`, `Konfidency` |

**Problema:** O script da Konfidency adicionado via API na plataforma Wake não é renderizado em todas as páginas do site do cliente — em alguns casos, apenas na home o script carregava. O lojista liberou o domínio nas políticas de Content Security Policy (CSP), mas o problema persistiu.

**Causa raiz:** A plataforma Wake possui uma política de segurança que bloqueia scripts não autorizados por padrão. A liberação precisa ser feita manualmente no admin para cada lojista.

**Resolução em andamento (2025-09):** O time de produto Wake avaliou a possibilidade de adicionar o subdomínio `reviews.konfidency.com.br` em uma whitelist geral, eliminando a necessidade de liberação manual cliente a cliente. A Konfidency confirmou que apenas esse subdomínio é utilizado pelos scripts.

**Status:** Em discussão no time de engenharia Wake. Enquanto não houver whitelist global, o processo de onboarding de cada novo cliente requer uma etapa manual de liberação no admin.

---

### BEST_PRACTICES-4 — Gestão de parceiros concorrentes no ecossistema Wake: impacto nas indicações

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Konfidency & Wake |
| **JID** | `120363413122423249@g.us` |
| **Data** | 2025-08-22 |
| **Categoria** | `best_practices` |
| **Tags** | `parceiros-concorrentes`, `programa-de-parcerias`, `indicações`, `CS`, `exclusividade`, `reviews` |

A situação de ter dois parceiros concorrentes diretos no programa (Konfidency e Vurdere, ambos de reviews) gerou atrito significativo:

**Problemas identificados:**
- CS Wake passou a indicar ambos os parceiros ao mesmo cliente sem critério claro, esvaziando o valor consultivo da indicação.
- Um lead enviado para ambos fechou com o concorrente, evidenciando falta de prioridade no direcionamento.
- A Wake não pode orientar nenhum dos parceiros a melhorar competitivamente sem gerar desigualdade.
- Parceiros que investem mensalidade no programa ficam em desvantagem em relação a parceiros sem mensalidade quando ambos recebem o mesmo número de indicações.

**Recomendações para o time de parcerias:**
- Definir critérios claros para direcionamento quando há parceiros concorrentes no portfólio.
- Comunicar ao time de CS antes do lançamento de um novo parceiro concorrente.
- Considerar períodos de exclusividade ou diferenciação por segmento/porte de cliente.
- O modelo de parceria com mensalidade implica compromisso maior de geração de indicações qualificadas.

---

### FAQ-4 — O que configura uma "indicação qualificada" no programa de parceiros Wake

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Konfidency & Wake |
| **JID** | `120363413122423249@g.us` |
| **Data** | 2025-09-17 |
| **Categoria** | `faq` |
| **Tags** | `indicação`, `leads`, `programa-de-parcerias`, `CS`, `qualificação` |

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

### HOW_TO-7 — Estratégia de ativação para Black November com parceiro de reviews

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Konfidency & Wake |
| **JID** | `120363413122423249@g.us` |
| **Data** | 2025-11-06 |
| **Categoria** | `how_to` |
| **Tags** | `black-friday`, `reviews`, `CS`, `ativação`, `parceiro`, `Konfidency` |

Estratégia desenvolvida entre Konfidency e Wake CS para Black November:

**Próximos passos definidos:**
1. **Material Konfidency para CS (prazo: 10/11):** One-pager com CTA específico — "Garanta a reputação da sua marca e aproveite a Black Friday para reunir avaliações sobre o seu negócio + oportunidade de trial até 05/01/26".
2. **Lista de clientes potenciais (prazo: 12/11):** CS Commerce + CS Experience levantam relação segmentada de prospects por área.
3. **Integração Wake Experience + Konfidency:** Fase 1 — export/import de dados em CSV/sheets; Fase 2 — integração homologada.

**Tática de conteúdo:** Para os ~20 clientes em comum entre Wake e Konfidency, criar pílulas semanais de conteúdo mostrando como a Konfidency está atendendo clientes Wake específicos — case privado direcionado ao CS responsável pela conta.

---

### TROUBLESHOOTING-3 — Problema de faturamento retroativo no programa de parcerias

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Konfidency & Wake |
| **JID** | `120363413122423249@g.us` |
| **Data** | 2025-07-02 |
| **Categoria** | `troubleshooting` |
| **Tags** | `faturamento`, `NF`, `contrato`, `financeiro`, `mensalidade`, `partner-fee` |

**Problema:** O financeiro da Wake enviou NFs de "Partner Fee" de meses anteriores ao kickoff e à assinatura do contrato:
- Partner Fee Março - R$ 4.167 (NF 52023)
- Partner Fee Abril - R$ 4.167 (NF 52681)
- Partner Fee Maio - R$ 4.167 (NF 53238)

O kickoff ocorreu em 09/04 e o contrato foi assinado em meados de abril. As NFs de março e abril não faziam sentido cronologicamente.

**Resolução:** Escalado para o time de parcerias para verificação e ajuste. A NF de maio era a máxima que poderia ser devida segundo o parceiro.

**Lição para o time:** Ao ativar o faturamento de um novo parceiro, o financeiro deve ser briefado com a data exata de início do contrato para evitar emissão retroativa de NFs.

---

### HOW_TO-8 — Processo de integração Wake Experience com plataformas de parceiros

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Konfidency & Wake |
| **JID** | `120363413122423249@g.us` |
| **Data** | 2025-12-01 |
| **Categoria** | `how_to` |
| **Tags** | `wake-experience`, `integração`, `parceiro-saas`, `embed`, `painel-wake` |

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

### HOW_TO-9 — Dois modelos de integração CRMBonus com Wake (V1 cupons vs. integração nativa)

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Wake + CRMBonus | Parcerias |
| **JID** | `120363356920429563@g.us` |
| **Data** | 2025-06-20 |
| **Categoria** | `how_to` |
| **Tags** | `CRMBonus`, `integração`, `cupons`, `v1`, `nativa`, `cashback` |

A CRMBonus possui **dois modelos distintos de integração** com a Wake:

1. **V1 (cupons/resgate manual):** Modelo de resgate de cupom — o consumidor interage ativamente para aplicar o benefício. Este modelo foi depreciado/desativado para alguns clientes (ex: Blueman teve a integração V1 desativada).

2. **Integração Nativa V2 (automatizada):** Integração automática onde o cashback/bônus é aplicado sem ação explícita do consumidor. Este é o modelo preferencial e recomendado.

**Importante:** São duas experiências completamente diferentes para o lojista e para o consumidor final. Ao resolver um chamado ou integrar um novo cliente, sempre confirmar qual modelo está em uso antes de diagnosticar problemas.

**Documentação técnica:** https://crmbonus-api.readme.io/reference/fluxo-da-integra%C3%A7%C3%A3o-bi

---

### TROUBLESHOOTING-4 — Cashback aplicado junto a promoções gerando acúmulo indevido (Shoulder/Inbrands)

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Wake + CRMBonus | Parcerias |
| **JID** | `120363356920429563@g.us` |
| **Data** | 2025-10-16 |
| **Categoria** | `troubleshooting` |
| **Tags** | `cashback`, `promoções`, `cupons`, `acúmulo`, `Shoulder`, `Inbrands`, `flag` |

**Problema:** O cashback CRMBonus estava sendo aplicado junto a outras promoções e cupons no carrinho, gerando acúmulo não previsto para lojistas como Shoulder e Inbrands.

**Solução técnica na CRMBonus:**

Existe uma flag de configuração por loja: `permitirProdutosPromicionados: true | false`

- Quando `true`: produtos com preço remarcado (precoDe ≠ precoPor) são excluídos do cálculo do bônus.
- Para promoções/cupons: o valor da promoção é **subtraído** do valor base de cálculo do bônus. Exemplo: 2 produtos de R$10 com promoção de 10% → base de cálculo do bônus = R$18, não R$20.

**Passo a passo para resolver:**
1. Verificar no painel CRMBonus se a flag `permitirProdutosPromicionados` está ativa para o lojista.
2. Confirmar se a regra cobre também o cenário de desconto por promoção/cupom (subtração do valor).
3. Alinhar com a Shoulder/cliente os próximos passos de configuração e confirmar que a regra atende o caso de uso deles.

---

### FAQ-5 — Como funciona o Giftback transparente (sem digitação de telefone) no checkout Wake

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Wake + CRMBonus | Parcerias |
| **JID** | `120363356920429563@g.us` |
| **Data** | 2025-10-24 |
| **Categoria** | `faq` |
| **Tags** | `CRMBonus`, `giftback`, `autenticação`, `checkout`, `Wake`, `UX` |

**Cenário:** Cliente quer implementar Giftback CRMBonus de forma transparente — sem exigir que o usuário digite o número de telefone durante a compra. Para isso, o telefone seria consultado diretamente no cadastro Wake via API.

**Perguntas e respostas:**
- **"A Wake suporta consulta de telefone do usuário autenticado via API?"** — Precisa de validação técnica com o time Wake. A CRMBonus precisa de confirmação se esse endpoint existe e está disponível.
- **"Na Wake, o usuário pode concluir compra sem estar autenticado?"** — Confirmação necessária com o time Wake (depende da configuração da loja).
- **"O desconto pode ser aplicado apenas na tela de pagamento, sem re-seleção de frete?"** — O gatilho padrão da alíquota Wake ocorre após a seleção do frete. Para lojas onde a tela de frete e pagamento são separadas, a aplicação na última tela exige análise técnica específica do fluxo da loja.

---

### HOW_TO-10 — Integração Wake Experience com CRMBonus (agenda do vendedor / listas de audiência)

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Wake + CRMBonus | Parcerias |
| **JID** | `120363356920429563@g.us` |
| **Data** | 2025-08-18 |
| **Categoria** | `how_to` |
| **Tags** | `wake-experience`, `CRMBonus`, `agenda-do-vendedor`, `audiências`, `integração`, `API` |

**Necessidade:** O time Wake Experience quer enviar dados de listas/audiências gerenciadas na plataforma (nomes, CPFs, e-mails ou telefones) para alimentar a Agenda do Vendedor da CRMBonus, gerando oportunidades automaticamente.

**Situação (agosto 2025):** Não existia integração pronta para esse fluxo na CRMBonus. A documentação disponível (`https://crmbonus-api.readme.io/reference/fluxo-da-integra%C3%A7%C3%A3o-bi`) era voltada a consulta/ingestão de dados para BI, não para envio de leads para a Agenda do Vendedor.

**Próximos passos recomendados:**
1. Contatar o Tech Lead da CRMBonus (Rodrigo Hidaka) para explorar possibilidades de desenvolvimento.
2. Marcar call técnica com as duas equipes para definir o formato da integração (API de envio, webhook, CSV/sheets etc.).
3. A CRMBonus sinalizou que existem "diversos caminhos" possíveis e recomenda alinhar o contexto antes de definir a arquitetura.

**Documentação de referência:** https://crmbonus-api.readme.io/reference/fluxo-da-integracao-ecomv2

---

### BEST_PRACTICES-5 — Atualização de status do pipeline de indicações no grupo de parcerias

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Wake + CRMBonus | Parcerias |
| **JID** | `120363356920429563@g.us` |
| **Data** | 2025-06-23 |
| **Categoria** | `best_practices` |
| **Tags** | `pipeline`, `indicações`, `CS`, `follow-up`, `parcerias` |

Para movimentar o grupo de parcerias e manter visibilidade do pipeline de indicações, a boa prática estabelecida foi trazer atualizações com os seguintes campos:

- **Data do contato inicial**
- **Nome do cliente indicado**
- **CS responsável pela indicação**
- **Data da reunião (se realizada)**
- **Temperatura do cliente** (interesse, receio, em avaliação etc.)
- **Próxima reunião agendada**

A planilha de acompanhamento é atualizada **toda segunda-feira** e o time de CS tem acesso direto. Para casos que precisam de apoio do CS por falta de contato ou dificuldade de avanço, o grupo de WhatsApp deve ser acionado diretamente.

---

### TROUBLESHOOTING-5 — Ambiente de homologação Wake para desenvolvimento de integração

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Wake + CRMBonus | Parcerias |
| **JID** | `120363356920429563@g.us` |
| **Data** | 2025-08-06 |
| **Categoria** | `troubleshooting` |
| **Tags** | `homologação`, `ambiente-de-teste`, `integração`, `desenvolvedor`, `acesso` |

**Situação:** Durante o desenvolvimento da integração CRMBonus V1 com a Wake, o time técnico da CRMBonus precisou de acesso a um ambiente de homologação Wake.

**Como obter acesso:**
1. O gerente de parcerias Wake solicita ao time de CS/tech o acesso ao ambiente de homologação.
2. O e-mail do desenvolvedor responsável deve ser fornecido para criação do acesso.
3. O acesso é vinculado a uma loja de teste — no caso da CRMBonus, foi direcionado para o responsável técnico (guilherme.sutto@crmbonus.com).

**Para dúvidas técnicas sobre a integração:** O time Wake disponibiliza agenda de call técnica com especialistas. Para integrações envolvendo gestão de cupons e promoções, usar o link de agenda da especialista designada:
`https://calendly.com/rafaela-teixeira/30min`

---

### FAQ-6 — Campo de cupom na segunda tela do checkout Wake: pode ser removido?

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Wake + CRMBonus | Parcerias |
| **JID** | `120363356920429563@g.us` |
| **Data** | 2026-02-06 |
| **Categoria** | `faq` |
| **Tags** | `checkout`, `cupom`, `Wake`, `Inbrands`, `customização`, `campo` |

**Pergunta (CRMBonus):** A Inbrands quer remover o campo de cupom na segunda tela do checkout para centralizar a aplicação de cupom apenas em uma tela — assim o Giftback CRMBonus não acumula indevidamente com cupons. O time Wake disse que o campo não pode ser removido. Por quê? A Shoulder não tem esse campo na segunda tela.

**Resposta:** O campo de cupom na segunda tela do checkout pode sim ser removido — a Shoulder é a prova disso. Não é uma limitação da plataforma, mas sim uma configuração. A Inbrands precisava ter esse ajuste feito pela agência responsável pelo checkout.

**Complicação reportada:** Após a agência da Inbrands alterar os campos, houve um conflito com o campo de "vendedor" que também estava sendo usado. O caso foi escalado para o time técnico Wake via ticket, e uma call foi agendada entre CRMBonus e o suporte Wake para resolver o conflito.

---

### BEST_PRACTICES-6 — Uso de verbas de eventos não utilizados para outros patrocínios

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Wake + CRMBonus | Parcerias |
| **JID** | `120363356920429563@g.us` |
| **Data** | 2025-09-26 |
| **Categoria** | `best_practices` |
| **Tags** | `verba`, `eventos`, `patrocínio`, `Wake-on-the-Road`, `negociação` |

Quando um parceiro possui verba de evento Wake não utilizada (ex: participação em Wake on the Road que não ocorreu), é possível redirecionar esse valor para outros eventos de parceria.

**Exemplo:** A CRMBonus tinha verba do Wake on the Road Curitiba que não participou. Essa verba foi proposta para cobrir o patrocínio de um evento de Farma organizado em parceria com a Wake.

**Processo:** A solicitação deve ser feita no grupo de parcerias e tratada via Slack com o time de marketing Wake. A resposta/aprovação é dada pelo gerente de parcerias.

---

*Fim do documento — Wave 6B*
