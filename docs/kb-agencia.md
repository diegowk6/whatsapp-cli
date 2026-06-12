# KB Extract — WAKE | Time Agência
> JID: 120363419712849778@g.us | Extracted: 2026-06-04

---

## FAQ

### FAQ-001 — Como classificar uma venda quando há atuação conjunta de agência e executivo direto?

| Campo | Valor |
|---|---|
| **Grupo/Chat** | WAKE \| Time Agência |
| **JID** | `120363419712849778@g.us` |
| **Data** | 2025-06-26 |
| **Categoria** | `faq` |
| **Tags** | `salesforce`, `classificação`, `venda-direta`, `venda-canal`, `origem` |

**Pergunta:** Uma oportunidade veio de indicação de agência, mas foi negociada e fechada em conjunto com um executivo de Vendas Diretas. Como ela deve ser classificada — como canal/indireto ou como direto?

**Resposta:** Se houver atuação de um executivo de Vendas Diretas (AE comercial), a oportunidade é contabilizada como **Venda Direta vinda de Agência**. A regra prática: se o campo "Responsável de Vendas" no Salesforce pertence a um AE de Diretas, é Diretas — porém com origem em canais/agência. Só é considerada Venda Indireta (Canal) quando toda a venda ocorre sem atuação do time de Diretas.

---

### FAQ-002 — O carrossel deve rodar para oportunidades de renovação de contrato?

| Campo | Valor |
|---|---|
| **Grupo/Chat** | WAKE \| Time Agência |
| **JID** | `120363419712849778@g.us` |
| **Data** | 2025-07-03 |
| **Categoria** | `faq` |
| **Tags** | `carrossel`, `renovação`, `salesforce`, `parceiro-operador` |

**Pergunta:** Quando uma oportunidade é de renovação de contrato, o carrossel de indicação de agências deve ser disparado?

**Resposta:** Não. Em caso de renovação, o correto é notificar apenas a agência que já opera aquela conta (parceiro operador atual), não rodar o carrossel para outras agências. Se o cliente pedir explicitamente outras opções, só então outras agências devem ser contatadas, e a agência atual deve ser informada dessa escolha do cliente. O sistema foi corrigido para que OPTs do tipo "renovação" não disparem o carrossel automaticamente.

---

### FAQ-003 — Qual modelo de indicação no Salesforce evita que o carrossel rode?

| Campo | Valor |
|---|---|
| **Grupo/Chat** | WAKE \| Time Agência |
| **JID** | `120363419712849778@g.us` |
| **Data** | 2025-07-11 |
| **Categoria** | `faq` |
| **Tags** | `salesforce`, `carrossel`, `indicação`, `venda-canal` |

**Pergunta:** Quais os tipos de indicação disponíveis no Salesforce e qual deles bloqueia o disparo do carrossel?

**Resposta:** Os tipos são: "Venda de Canal", "Indicação Agências" e "Indicação Parcerias". Para que o carrossel **não** rode, o campo "Parceiro Operador" deve estar preenchido com a agência responsável pela conta. O carrossel roda quando não há Parceiro Operador definido. Oportunidades do tipo "Renovação" também foram configuradas para não acionar o carrossel.

---

### FAQ-004 — Uma agência pode vender proativamente o pacote Storefront para clientes próprios dela?

| Campo | Valor |
|---|---|
| **Grupo/Chat** | WAKE \| Time Agência |
| **JID** | `120363419712849778@g.us` |
| **Data** | 2025-07-07 |
| **Categoria** | `faq` |
| **Tags** | `storefront`, `agência`, `prospecção`, `base-clientes` |

**Pergunta:** Uma agência parceira pode proativamente oferecer o pacote Storefront para clientes que já são dela, sem esperar as indicações via carrossel?

**Resposta:** Sim, se forem clientes da própria carteira da agência. O processo correto é: a agência aborda o cliente, envolve o CS da Wake (para estimular o time de CS também) e usa o escopo e preço tabelado do pacote. O cuidado necessário é que a agência não saia prospectando clientes de outras agências ou da base Wake indiscriminadamente — isso já gerou atritos no passado e pode resultar em penalizações.

---

### FAQ-005 — Qual o fluxo de pagamento do Revenue Share para agências?

| Campo | Valor |
|---|---|
| **Grupo/Chat** | WAKE \| Time Agência |
| **JID** | `120363419712849778@g.us` |
| **Data** | 2025-08-21 |
| **Categoria** | `faq` |
| **Tags** | `revenue-share`, `pagamento`, `nota-fiscal`, `contrato` |

**Pergunta:** Como funciona o processo de pagamento do Revenue Share para as agências parceiras? Precisa de contrato? Precisa de nota fiscal?

**Resposta:** Para qualquer pagamento de Revenue Share ou permuta de projeto, é obrigatório ter um contrato formal assinado. Sem contrato, o pagamento não pode ser processado. O fluxo é: (1) Wake envia os valores calculados via e-mail para a agência validar; (2) a agência valida e emite a Nota Fiscal; (3) Wake realiza o pagamento em até 15 dias após o recebimento da NF. Tratativas verbais ou acordos informais não garantem o pagamento.

---

### FAQ-006 — Quanto tempo depois de assinar o contrato a agência começa a receber leads do carrossel?

| Campo | Valor |
|---|---|
| **Grupo/Chat** | WAKE \| Time Agência |
| **JID** | `120363419712849778@g.us` |
| **Data** | 2025-08-15 |
| **Categoria** | `faq` |
| **Tags** | `carrossel`, `onboarding-agência`, `certificação`, `leads` |

**Pergunta:** Uma agência que acabou de assinar o contrato pode começar a receber leads pelo carrossel imediatamente?

**Resposta:** A prática recomendada (em discussão para formalização) é que a agência deve primeiro concluir a certificação mínima antes de receber leads do carrossel. A justificativa é que sem o treinamento mínimo, a agência não está preparada para converter os leads adequadamente, o que prejudica a taxa de conversão do ecossistema. O prazo sugerido é de 30 a 90 dias de maturidade antes de entrar no carrossel, que inclui: certificação + reunião de Wake Education comercial + construção de afinidade com o time de parcerias.

---

### FAQ-007 — Como funciona a lógica de distribuição do carrossel entre agências do mesmo nível?

| Campo | Valor |
|---|---|
| **Grupo/Chat** | WAKE \| Time Agência |
| **JID** | `120363419712849778@g.us` |
| **Data** | 2025-08-20 |
| **Categoria** | `faq` |
| **Tags** | `carrossel`, `score`, `plano`, `distribuição-leads` |

**Pergunta:** Como o carrossel decide qual agência recebe cada lead? É pelo plano pago ou pelo score?

**Resposta:** O carrossel usa dois fatores: (1) o **Score** define o tipo de lead que a agência pode receber (leads de Canais = Essential; leads do Comercial = Special/Elite); (2) a **frequência** de indicações é determinada pelo plano pago:
- Plano I → entra na fila 1x
- Plano II → entra na fila 1,5x
- Plano III → entra na fila 2x

Além disso, o algoritmo calcula um percentual de oportunidades recebidas vs. esperadas. Agências com percentual menor têm prioridade. Um bug histórico fazia agências novas receberem mais leads por entrar com percentual zero — esse bug foi corrigido em outubro/2025.

---

### FAQ-008 — Qual é a diferença entre os planos Essential, Special e Elite no programa de parceiros?

| Campo | Valor |
|---|---|
| **Grupo/Chat** | WAKE \| Time Agência |
| **JID** | `120363419712849778@g.us` |
| **Data** | 2025-07-23 |
| **Categoria** | `faq` |
| **Tags** | `plano-agência`, `essential`, `special`, `elite`, `score`, `wake-score` |

**Pergunta:** Qual a diferença entre os planos/níveis Essential, Special e Elite para agências parceiras Wake?

**Resposta:** Os níveis determinam o tipo de lead que a agência pode receber pelo carrossel:
- **Essential** → recebe leads vindos de Canais/Parcerias (indicações do time de parcerias)
- **Special** → recebe tanto leads de Canais quanto do time Comercial
- **Elite** → nível mais alto, agências estratégicas com maior volume e histórico de conversão

A progressão entre níveis é definida pelo **Wake Score**, que é calculado periodicamente (revisão prevista para agosto/2025 e Q1/2026) com base em métricas como: contratos fechados, certificações tiradas, role plays realizados, postagens em redes sociais, entre outros critérios. Agências que não cumprem os requisitos podem ser rebaixadas.

---

### FAQ-009 — Como registrar um lead no Salesforce quando a agência indicou mas o cliente entrou pelo formulário do site?

| Campo | Valor |
|---|---|
| **Grupo/Chat** | WAKE \| Time Agência |
| **JID** | `120363419712849778@g.us` |
| **Data** | 2025-09-19 |
| **Categoria** | `faq` |
| **Tags** | `salesforce`, `lead`, `origem`, `agência`, `parcerias-e-canais` |

**Pergunta:** Quando uma agência indicou um lead mas o cliente entrou pelo formulário do site (origem = "Formulário do site"), como deve ficar o registro no Salesforce para garantir crédito à agência?

**Resposta:** Neste caso, é necessário corrigir manualmente o registro: (1) trocar a origem de "Formulário do site" para "Parcerias e Canais"; (2) retirar os parceiros indicados que eventualmente foram disparados pelo carrossel; (3) vincular corretamente a agência que originou o lead como parceiro responsável. Para formalizar, abre-se um ticket no Salesforce com a solicitação de ajuste.

---

### FAQ-010 — Agências que aderem ao plano de Storefront precisam de um documento de aceite formal?

| Campo | Valor |
|---|---|
| **Grupo/Chat** | WAKE \| Time Agência |
| **JID** | `120363419712849778@g.us` |
| **Data** | 2025-07-23 |
| **Categoria** | `faq` |
| **Tags** | `storefront`, `contrato`, `aceite`, `pacote-migração` |

**Pergunta:** Para uma agência aderir ao pacote de Storefront e entrar no carrossel manual de Storefront, precisa assinar algum documento específico?

**Resposta:** Não há um documento de aceite separado formalizado. O processo atual é preencher a aba "full base - escopo" na planilha de controle do Storefront. O time de parcerias envia o material de escopo ("Storefront 2.0 - Empacotamento migração.pdf") para a agência, e o aceite é registrado na planilha. Um processo de formalização via documento assinado está sendo estruturado.

---

### FAQ-011 — Uma agência que não tem CNAE de serviços pode receber Revenue Share da Wake?

| Campo | Valor |
|---|---|
| **Grupo/Chat** | WAKE \| Time Agência |
| **JID** | `120363419712849778@g.us` |
| **Data** | 2025-11-27 |
| **Categoria** | `faq` |
| **Tags** | `revenue-share`, `cnpj`, `cnae`, `jurídico`, `nota-fiscal` |

**Pergunta:** Um cliente ou consultor que quer indicar leads para Wake e receber Revenue Share, mas não tem CNAE adequado para emitir NF de serviços de "intermediação" — como proceder?

**Resposta:** Este é um caso complexo que precisa de análise jurídica. A solução recomendada é que a pessoa abra um CNPJ com o CNAE correto (ex: abertura de uma agência) ou que assine o contrato de parceria padrão da Wake (sem mensalidade, apenas com cláusula de Rev Share), que o jurídico está avaliando como viabilizar. O caso foi escalado para uma thread no Slack com o jurídico para criar precedente e processo escalável.

---

### FAQ-012 — O carrossel dispara para oportunidades de produtos que não são Commerce (ex: OMS, Creators, DXP)?

| Campo | Valor |
|---|---|
| **Grupo/Chat** | WAKE \| Time Agência |
| **JID** | `120363419712849778@g.us` |
| **Data** | 2025-07-24 |
| **Categoria** | `faq` |
| **Tags** | `carrossel`, `oms`, `creators`, `dxp`, `divisão`, `salesforce` |

**Pergunta:** O carrossel de indicação de agências dispara para oportunidades de todas as divisões ou apenas para Wake Commerce?

**Resposta:** O carrossel deve ser configurado para rodar apenas para as divisões Time A, Time B e Canais (Wake Commerce). Não deve rodar para: OMS, Creators, DXP ou qualquer outra divisão que não seja de e-commerce Commerce. Casos de disparo indevido para essas divisões foram registrados como bugs e corrigidos pontualmente. A solicitação é que o Salesforce implemente uma trava por divisão.

---

## Troubleshooting

### TRB-001 — Carrossel disparando para agências em oportunidades de renovação

| Campo | Valor |
|---|---|
| **Grupo/Chat** | WAKE \| Time Agência |
| **JID** | `120363419712849778@g.us` |
| **Data** | 2025-07-03 |
| **Categoria** | `troubleshooting` |
| **Tags** | `carrossel`, `renovação`, `salesforce`, `bug`, `sales-ops` |

**Sintomas:** Oportunidades de renovação de contrato de clientes com agência operadora definida foram enviadas pelo carrossel para 3 agências aleatórias. As agências receberam a oportunidade e ficaram aguardando contato, mas o cliente já tinha agência.

**Causa raiz:** O carrossel não tinha trava para o tipo de oportunidade "Renovação" e para casos onde o campo "Parceiro Operador" já estava preenchido. Um executivo (Samires) abriu OPTs de renovação e o SF/Community disparou o carrossel normalmente.

**Resolução:** (1) Ticket aberto no Sales Ops para correção da regra; (2) Para as agências já indicadas erroneamente: a indicação foi "resetada" para não prejudicar o contador delas no carrossel; (3) O sistema foi corrigido — OPTs do tipo "Renovação" e OPTs com Parceiro Operador já preenchido não acionam mais o carrossel.

---

### TRB-002 — Agências novas recebendo mais leads que agências antigas no carrossel

| Campo | Valor |
|---|---|
| **Grupo/Chat** | WAKE \| Time Agência |
| **JID** | `120363419712849778@g.us` |
| **Data** | 2025-08-20 |
| **Categoria** | `troubleshooting` |
| **Tags** | `carrossel`, `bug`, `distribuição`, `salesforce`, `score` |

**Sintomas:** Agências recém-entradas (FZ Commerce, Simples Inovação, Nairuz) recebiam 3 oportunidades por semana, enquanto agências com planos superiores e mais tempo no programa (Eficaz, O2, Ge Commerce, Box Ideias) ficavam sem receber por semanas ou meses.

**Causa raiz:** Bug no algoritmo do carrossel: ao entrar no programa, as agências novas eram inseridas com percentual de oportunidades recebidas igual a zero. O algoritmo, ao calcular o delta entre "recebidas" e "esperadas", priorizava sempre as que tinham percentual menor — ou seja, as novas sempre ficavam na frente até atingirem o percentual dos pares.

**Resolução:** O desenvolvedor Lucas corrigiu o bug em outubro/2025. Agências novas agora entram com o mesmo número de indicações dos pares no sistema, indo para o fim da fila em vez de para o início. Os percentuais foram recalculados. Recomendação adicional: agências novas não devem entrar no carrossel antes de concluir a certificação mínima, para não repetir o desequilíbrio.

---

### TRB-003 — Lead convertido em OPT disparou carrossel para agências erradas

| Campo | Valor |
|---|---|
| **Grupo/Chat** | WAKE \| Time Agência |
| **JID** | `120363419712849778@g.us` |
| **Data** | 2025-07-11 |
| **Categoria** | `troubleshooting` |
| **Tags** | `salesforce`, `carrossel`, `lead`, `conversão`, `parceiro-origem` |

**Sintomas:** Ao converter um lead em OPT (onde o lead tinha origem em uma agência parceira), o carrossel disparou automaticamente para outras 3 agências, incluindo a própria agência que gerou o lead. O campo "Parceiro Origem" estava preenchido, mas o sistema não bloqueou o carrossel.

**Causa raiz:** O carrossel considerava apenas o campo "Parceiro Operador" para bloquear o disparo, não o "Parceiro Origem". Quando um lead de agência é convertido por um executivo direto, sem preencher corretamente o parceiro operador, o carrossel roda normalmente.

**Resolução:** Procedimento de correção manual: (1) remover as agências indicadas erroneamente via ticket no Salesforce; (2) garantir que a agência origem fique como "Parceiro Indicado 1"; (3) trocar a origem para "Parcerias e Canais". Para prevenção, o time de executivos deve ser treinado para preencher o campo Parceiro Operador antes de converter leads de agências.

---

### TRB-004 — Agência com plano ativo sem receber leads por mais de 30 dias

| Campo | Valor |
|---|---|
| **Grupo/Chat** | WAKE \| Time Agência |
| **JID** | `120363419712849778@g.us` |
| **Data** | 2025-09-23 |
| **Categoria** | `troubleshooting` |
| **Tags** | `carrossel`, `leads`, `plano`, `reequilíbrio`, `salesforce` |

**Sintomas:** Agência pagante (O2 Consultoria, Plano I) relatou não receber nenhuma oportunidade há mais de 50 dias, o que a levou a reconsiderar a permanência no programa.

**Causa raiz:** Combinação de fatores: (1) bug antigo do carrossel que priorizava agências novas; (2) competição de agências com planos maiores (que têm prioridade de frequência 2x); (3) possível configuração incorreta de score no Salesforce.

**Resolução:** O procedimento para reequilíbrio é: (1) puxar dados quantitativos — quantos leads foram gerados no total e qual o proporcional recebido pela agência nos últimos meses; (2) calcular a proporcionalidade esperada por plano; (3) o time de Salesforce pode fazer uma "priorização manual" da agência para que receba as próximas oportunidades do carrossel. O argumento para o parceiro deve ser apresentado com horizonte trimestral, não apenas o mês corrente.

---

### TRB-005 — Agência indicada em OPT que não deveria ter sido indicada (conflito de interesse)

| Campo | Valor |
|---|---|
| **Grupo/Chat** | WAKE \| Time Agência |
| **JID** | `120363419712849778@g.us` |
| **Data** | 2025-07-22 |
| **Categoria** | `troubleshooting` |
| **Tags** | `carrossel`, `indicação-indevida`, `relatório`, `quicksight` |

**Sintomas:** Diversas oportunidades aparecem nos relatórios de parceiros (QuickSight/dashboard) com agências vinculadas, mas que não deveriam estar ali — seja porque a oportunidade era teste, de renovação, de outro produto, ou porque a agência foi indicada erroneamente pelo carrossel.

**Causa raiz:** Falhas na configuração do carrossel (sem travas adequadas por tipo de OPT, divisão e tamanho de conta) e ausência de validação antes do disparo.

**Resolução:** Processo de limpeza manual periódica:
1. Identificar OPTs indevidas no relatório do QuickSight (dashboard de parceiros)
2. Para cada OPT: retirar o "Parceiro via Carrossel", mantendo apenas origem/operador correto
3. Resetar o contador de indicações das agências prejudicadas
4. Documentar as OPTs corrigidas em planilha compartilhada com o time de Salesforce
5. Registrar ticket formal para cada correção para rastreabilidade

---

### TRB-006 — Agência recomendando cliente a sair da Wake para concorrente

| Campo | Valor |
|---|---|
| **Grupo/Chat** | WAKE \| Time Agência |
| **JID** | `120363419712849778@g.us` |
| **Data** | 2025-07-04 |
| **Categoria** | `troubleshooting` |
| **Tags** | `agência`, `churn`, `cliente`, `relacionamento`, `parceiro-problemático` |

**Sintomas:** Cliente reclamou de demora na entrega do projeto à Wake. A Wake repassou o feedback para a agência (Benova). A agência, em vez de resolver, informou ao cliente que era melhor migrar para outra plataforma.

**Causa raiz:** Agência com relacionamento desgastado com a Wake (histórico de problemas de entrega em múltiplos projetos), sentiu-se acuada pelo feedback e reagiu de forma contrária aos interesses da Wake.

**Resolução:** (1) Reunião trilateral (Wake + agência + cliente) para entender os detalhes; (2) Separar os casos: o cliente ficou insatisfeito com a agência, não com a plataforma; (3) Avaliar a continuidade da parceria com a agência considerando o histórico de problemas; (4) O cliente migrou de agência após a entrega do projeto.

---

### TRB-007 — Agência prospectando clientes da base Wake com abordagem inadequada

| Campo | Valor |
|---|---|
| **Grupo/Chat** | WAKE \| Time Agência |
| **JID** | `120363419712849778@g.us` |
| **Data** | 2025-09-25 |
| **Categoria** | `troubleshooting` |
| **Tags** | `agência`, `prospecção-indevida`, `base-clientes`, `relacionamento` |

**Sintomas:** Agência parceira enviou mensagem a clientes da base Wake perguntando sobre a "satisfação com a plataforma" e oferecendo análise gratuita de site, passando a impressão de que a Wake havia enviado a mensagem ou estava insatisfeita com a performance dos clientes.

**Causa raiz:** Agência saiu prospectando a base de clientes Wake de forma independente e sem alinhamento, usando abordagem que gerou ambiguidade sobre o remetente e questionou indiretamente a satisfação com a plataforma.

**Resolução:** (1) Reunião com a agência para alinhar que prospectar clientes da base de outras agências é ineficiente e gera ruído; (2) Orientar que o correto é prospectar a própria carteira e novos negócios, não a base Wake de outras agências; (3) Agência pode ser penalizada no Wake Score por comportamento contrário ao ecossistema; (4) Deixar claro que a Wake pode ajudar a direcionar esforço de prospecção para o lugar certo (base de migração de concorrentes, campanhas específicas).

---

### TRB-008 — Campanha de vendas com participantes sem processo de pagamento definido

| Campo | Valor |
|---|---|
| **Grupo/Chat** | WAKE \| Time Agência |
| **JID** | `120363419712849778@g.us` |
| **Data** | 2025-11-25 |
| **Categoria** | `troubleshooting` |
| **Tags** | `campanha-vendas`, `pagamento`, `pessoa-física`, `nota-fiscal`, `jurídico` |

**Sintomas:** Vendedores (pessoas físicas de agências parceiras) se inscreveram na campanha de vendas e fecharam contratos elegíveis à bonificação, mas o processo de pagamento não estava definido — não era claro se o pagamento seria à agência ou à pessoa física, e se precisaria de NF.

**Causa raiz:** A campanha foi criada sem definir previamente o fluxo contábil/jurídico para pagamentos a pessoas físicas. O jurídico não havia sido consultado sobre o processo de emissão de NF e prazo de pagamento para esse perfil.

**Resolução:** Processo acordado: (1) se o pagamento for **via agência**: agência emite NF, prazo de até 30 dias; (2) se o pagamento for **direto à pessoa física**: é necessário um Termo de Prestação de Serviços assinado pela pessoa física, prazo de até 45 dias. Para campanhas futuras: enviar o termo de aceite no momento da inscrição, para ter tudo formalizado antes do fechamento.

---

## How-To

### HT-001 — Como fazer a indicação manual de agências para uma oportunidade (sem carrossel)

| Campo | Valor |
|---|---|
| **Grupo/Chat** | WAKE \| Time Agência |
| **JID** | `120363419712849778@g.us` |
| **Data** | 2025-07-24 |
| **Categoria** | `how_to` |
| **Tags** | `salesforce`, `indicação-manual`, `parceiro-indicado`, `carrossel` |

1. Abrir a oportunidade no Salesforce
2. Navegar até a aba **Parcerias** (não a aba Detalhes)
3. No campo **"Parceiro Indicado 1"** → selecionar a primeira agência
4. No campo **"Parceiro Indicado 2"** → selecionar a segunda agência (se houver)
5. No campo **"Parceiro Indicado 3"** → selecionar a terceira agência (se houver)
6. Se ainda não há agência operadora definida, **não preencher** o campo "Parceiro Operador" (ou deixar como "a definir")
7. Avisar as agências indicadas sobre a oportunidade (o Community pode não notificar automaticamente em indicações manuais — verificar)
8. Registrar no documento "Indicações Estratégicas" para controle do time

---

### HT-002 — Como registrar um lead gerado por agência parceira no Salesforce

| Campo | Valor |
|---|---|
| **Grupo/Chat** | WAKE \| Time Agência |
| **JID** | `120363419712849778@g.us` |
| **Data** | 2025-08-20 |
| **Categoria** | `how_to` |
| **Tags** | `salesforce`, `lead`, `agência`, `cadastro`, `origem` |

1. Solicitar à agência os dados mínimos do lead:
   - Nome e Sobrenome do contato
   - Telefone
   - E-mail corporativo
   - Nome da empresa
   - CNPJ
2. Acessar o Salesforce e criar um novo Lead
3. Preencher todos os campos obrigatórios (atenção: Sobrenome é obrigatório; Segmento deve ser preenchido — não pode ficar em branco)
4. No campo **"Origem"** → selecionar **"Parcerias e Canais"**
5. Vincular a agência parceira como indicadora (campo "Parceiro Origem")
6. Se o lead deve ir direto para uma OPT (o parceiro já identificou oportunidade concreta), converter o lead em OPT logo em seguida, garantindo que o parceiro fique como "Parceiro Indicado 1"
7. Notificar o executivo responsável sobre o novo lead/OPT

---

### HT-003 — Como retirar agências indicadas erroneamente de uma oportunidade

| Campo | Valor |
|---|---|
| **Grupo/Chat** | WAKE \| Time Agência |
| **JID** | `120363419712849778@g.us` |
| **Data** | 2025-07-22 |
| **Categoria** | `how_to` |
| **Tags** | `salesforce`, `carrossel`, `correção`, `indicação-indevida`, `ticket` |

1. Identificar as oportunidades com indicações indevidas (via relatório no QuickSight ou aviso do time de parcerias)
2. Para cada oportunidade: abrir o link no Salesforce
3. Abrir um ticket formal no canal do Salesforce com as informações:
   - Link da oportunidade
   - Agências a serem retiradas
   - Motivo (ex: "renovação", "OPT de teste", "divisão errada", "agência não atua neste segmento")
4. O time de Sales Ops/Salesforce remove os parceiros indicados e mantém apenas origem/operador correto
5. Verificar se o contador de indicações das agências foi resetado (elas não devem ser penalizadas no carrossel por uma indicação indevida)
6. Remover também do relatório de "Parceiro Operador" no dashboard do QuickSight se necessário
7. Documentar a lista de OPTs corrigidas para auditoria

---

### HT-004 — Como fazer onboarding de uma nova agência parceira Wake

| Campo | Valor |
|---|---|
| **Grupo/Chat** | WAKE \| Time Agência |
| **JID** | `120363419712849778@g.us` |
| **Data** | 2025-08-15 |
| **Categoria** | `how_to` |
| **Tags** | `onboarding`, `agência`, `parceiro`, `certificação`, `kickoff` |

1. **Captação**: Identificar agência interessada (evento, indicação, abordagem ativa do time de parcerias)
2. **Apresentação do Programa**: Fazer reunião de apresentação com o CEO/responsável comercial da agência, cobrindo: planos disponíveis (I, II, III), funcionamento do carrossel, Wake Score, Revenue Share
3. **Proposta e Contrato**: Enviar proposta comercial com o plano escolhido; aguardar assinatura (plataforma de assinatura digital)
4. **Kickoff**: Após assinatura, agendar kickoff com o time de parcerias para:
   - Apresentar os executivos comerciais Wake com quem vão trabalhar
   - Explicar o Community (portal de parceiros)
   - Compartilhar link da certificação Wake
5. **Certificação**: Cobrar que a agência conclua a certificação no prazo (30 dias recomendados)
6. **Wake Education Comercial**: Após certificação, realizar reunião de Education para treinar o time comercial da agência (não apenas o técnico) em como vender Wake
7. **Ativação no Carrossel**: Apenas após concluir a certificação e o Education, liberar a agência para receber leads via carrossel
8. **Grupo WhatsApp**: Criar grupo do WhatsApp com os contatos principais da agência e o time de parcerias Wake
9. **Divulgação**: Publicar nas redes sociais da Wake (LinkedIn/Instagram) anunciando a nova agência parceira

---

### HT-005 — Como calcular e apresentar o desempenho de leads para uma agência parceira

| Campo | Valor |
|---|---|
| **Grupo/Chat** | WAKE \| Time Agência |
| **JID** | `120363419712849778@g.us` |
| **Data** | 2025-10-02 |
| **Categoria** | `how_to` |
| **Tags** | `leads`, `relatório`, `quicksight`, `agência`, `forecast` |

1. Acessar o QuickSight no painel de indicadores: `https://us-east-1.quicksight.aws.amazon.com/sn/account/shared-services-wake/dashboards/6add377d-c600-459d-8a66-db1985409be4`
2. **Não usar** o filtro "Esta semana" — ele pode mostrar dados incompletos (snapshots antigos). Usar o filtro de período manual
3. Ao apresentar para um parceiro que reclama de baixo volume, **abrir o horizonte temporal** — mostrar Q1, Q2 e Q3 juntos, não apenas o mês corrente
4. Calcular: total de leads no período → proporção esperada pelo plano (1x, 1,5x ou 2x) → comparar com o recebido
5. Se a agência reclama da qualidade (não volume), isso é um problema diferente — investigar o perfil dos leads direcionados
6. Ter em mãos os dados de conversão separados: leads totais recebidos vs. convertidos em OPT vs. OPTs ganhas como parceiro operador
7. Documentar o forecast com a agência em planilha própria (não apenas no Salesforce), com colunas de status de cada oportunidade

---

### HT-006 — Como processar o pagamento de projetos para agências parceiras (permuta/bonificação)

| Campo | Valor |
|---|---|
| **Grupo/Chat** | WAKE \| Time Agência |
| **JID** | `120363419712849778@g.us` |
| **Data** | 2025-08-21 |
| **Categoria** | `how_to` |
| **Tags** | `pagamento`, `permuta`, `projeto`, `nota-fiscal`, `contrato` |

1. Verificar se há contrato formal assinado entre Wake e agência cobrindo o projeto em questão (sem contrato, não é possível prosseguir)
2. Enviar e-mail formal para a agência com os valores calculados para validação
3. Agência valida os valores e emite a **Nota Fiscal** contra a Wake (CNPJ da Wake)
4. NF deve contemplar os projetos acordados — caso sejam múltiplos projetos, todos devem estar na mesma NF
5. Após recebimento da NF, o pagamento é processado em até **15 dias úteis**
6. Registrar o PR (Purchase Request) interno para liberação do pagamento
7. Comunicar à agência o prazo esperado de depósito
8. **Atenção**: tratativas verbais ou mensagens de WhatsApp não substituem o contrato formal — todos os acordos de pagamento devem ser formalizados por escrito antes do início dos trabalhos

---

### HT-007 — Como fazer o reporte semanal de atividade para o gestor do time de agências

| Campo | Valor |
|---|---|
| **Grupo/Chat** | WAKE \| Time Agência |
| **JID** | `120363419712849778@g.us` |
| **Data** | 2025-07-11 |
| **Categoria** | `how_to` |
| **Tags** | `reporte`, `leads`, `opts`, `fechamentos`, `rotina` |

O reporte semanal deve ser enviado ao final de cada semana com o seguinte formato:

1. **OPTs geradas na semana** — listar cada OPT: nome do cliente, executivo responsável, agência que originou
2. **Leads gerados na semana** — listar cada lead: nome do cliente, agência que indicou, divisão (Time A/B/Canais)
3. **Fechamentos / Best Case / Commit do mês** — listar negociações em andamento com status:
   - T&C (termos e condições) assinados
   - Commit (alta probabilidade de fechar)
   - Best Case (provável mas com risco)
4. **Forecast por agência** — status das principais oportunidades de cada parceiro, com observações sobre próximas ações
5. **Agenda com novas agências em negociação** — agências prospectadas e em fase de apresentação/onboarding

---

### HT-008 — Como usar a campanha de migração de Mercado Shops para gerar leads

| Campo | Valor |
|---|---|
| **Grupo/Chat** | WAKE \| Time Agência |
| **JID** | `120363419712849778@g.us` |
| **Data** | 2025-09-22 |
| **Categoria** | `how_to` |
| **Tags** | `mercado-shops`, `migração`, `campanha`, `leads`, `incentivo` |

1. Identificar clientes que usam Mercado Shops na carteira das agências parceiras (e na base de prospecção)
2. Apresentar aos parceiros a campanha de incentivo: **as primeiras 12 mensalidades são isentas** para migrações de Mercado Shops
3. O MRR é calculado sobre o GMV do cliente (usando a alíquota padrão de TR de 2,3%)
4. Solicitar às agências que mapeiem na base delas quais clientes usam Mercado Shops e abram como leads no Salesforce
5. Quando indicar para executivos, incluir o argumento do incentivo de 12 meses isentos para acelerar o fechamento
6. Registrar os leads com tag de evento/campanha correspondente para rastreio de conversão
7. Caso o cliente já tenha assinado sem usar o incentivo, verificar se é possível aplicar retroativamente (requer validação interna)

---

## Best Practices

### BP-001 — Nunca vender Wake sem agência implementadora definida

| Campo | Valor |
|---|---|
| **Grupo/Chat** | WAKE \| Time Agência |
| **JID** | `120363419712849778@g.us` |
| **Data** | 2025-07-22 |
| **Categoria** | `best_practices` |
| **Tags** | `agência`, `implementação`, `qualificação`, `SDR`, `processo-de-venda` |

**Recomendação:** Toda venda de Wake Commerce deve ser condicionada à existência de uma agência implementadora. Historicamente, vendas sem agência geraram problemas graves na fase de implementação (cliente sem suporte técnico, projetos abandonados, churn). A SDR deve reforçar com todos os leads que a Wake só opera com agência parceira para a implementação.

**Rationale:** O ecossistema de agências é parte central da proposta de valor da Wake. Sem agência implementadora, o cliente fica desamparado e a taxa de sucesso do projeto cai significativamente. Isso também protege as agências parceiras, que dependem de novos projetos para seus negócios.

---

### BP-002 — Usar cadência controlada ao distribuir oportunidades do carrossel para agências

| Campo | Valor |
|---|---|
| **Grupo/Chat** | WAKE \| Time Agência |
| **JID** | `120363419712849778@g.us` |
| **Data** | 2025-06-20 |
| **Categoria** | `best_practices` |
| **Tags** | `carrossel`, `cadência`, `leads`, `agência`, `parceiro` |

**Recomendação:** Ao disponibilizar um lote de oportunidades para agências (ex: lista de 30 oportunidades não atribuídas), não distribuir todas de uma vez. Fazer a distribuição em cadência — um grupo por semana — para que as agências sempre tenham novidades e o relacionamento seja mantido ativo ao longo do mês.

**Rationale:** Distribuir todas as oportunidades de uma só vez não garante que as agências vão trabalhar efetivamente. A cadência semanal mantém o engajamento, gera senso de continuidade e dá tempo para as agências processarem cada oportunidade adequadamente.

---

### BP-003 — Apresentar o desempenho de parceiros com horizonte trimestral, não mensal

| Campo | Valor |
|---|---|
| **Grupo/Chat** | WAKE \| Time Agência |
| **JID** | `120363419712849778@g.us` |
| **Data** | 2025-10-02 |
| **Categoria** | `best_practices` |
| **Tags** | `relacionamento`, `agência`, `leads`, `relatório`, `forecast` |

**Recomendação:** Ao conversar com parceiros sobre volume de leads/oportunidades recebidos, sempre apresentar os dados em perspectiva trimestral (ex: Q1: 12 leads, Q2: 10, Q3: 16) — não apenas o mês corrente. Isso evita que um mês pontualmente fraco distorça a percepção de valor do programa.

**Rationale:** Um mês ruim em volume de leads pode ter várias causas conjunturais (bug no carrossel, entrada de agências novas, sazonalidade). O trimestre dá uma visão mais estável e justa do programa. Parceiros que focam apenas no mês tendem a reclamar e considerar cancelamento com base em dados incompletos.

---

### BP-004 — Exigir certificação antes de liberar agência para o carrossel de leads

| Campo | Valor |
|---|---|
| **Grupo/Chat** | WAKE \| Time Agência |
| **JID** | `120363419712849778@g.us` |
| **Data** | 2025-08-15 |
| **Categoria** | `best_practices` |
| **Tags** | `certificação`, `onboarding`, `carrossel`, `qualidade`, `conversão` |

**Recomendação:** Agências novas não devem entrar no carrossel de leads antes de concluir a certificação mínima da Wake. A trilha recomendada é: assinar contrato → certificar em 30 dias → Wake Education Comercial → entrar no carrossel. Para agências que já implementaram Wake anteriormente, a trilha pode ser encurtada.

**Rationale:** Agências sem certificação não conhecem bem a plataforma e não conseguem vender/converter os leads adequadamente. Isso resulta em baixa taxa de conversão, insatisfação do cliente (que não recebe apresentação de qualidade) e desperdício de oportunidades. A VTEX, por comparação, só libera leads após a agência concluir certificações e subir 3 projetos — a Wake tem um requisito menor, mas algum requisito é necessário.

---

### BP-005 — Usar role play com agências para melhorar a qualidade das vendas

| Campo | Valor |
|---|---|
| **Grupo/Chat** | WAKE \| Time Agência |
| **JID** | `120363419712849778@g.us` |
| **Data** | 2025-11-27 |
| **Categoria** | `best_practices` |
| **Tags** | `role-play`, `treinamento`, `agência`, `comercial`, `wake-score` |

**Recomendação:** Realizar sessões de role play com os comerciais das agências parceiras para prepará-los para apresentar Wake para prospects. O role play deve fazer parte do Wake Score (agências que participam ganham pontos). Prioridade para agências Elite e Special.

**Rationale:** O principal feedback recebido de uma grande agência parceira (B8One) foi que os comerciais das agências não se sentem tão preparados para defender Wake quanto defendem VTEX (plataforma com a qual têm mais histórico). O role play, combinado com acesso a um painel de demonstração e trilha de treinamento comercial robusta, é o caminho para reverter esse gap competitivo.

---

### BP-006 — Manter agenda semanal entre time de Parcerias e SDR

| Campo | Valor |
|---|---|
| **Grupo/Chat** | WAKE \| Time Agência |
| **JID** | `120363419712849778@g.us` |
| **Data** | 2025-07-22 |
| **Categoria** | `best_practices` |
| **Tags** | `SDR`, `qualificação`, `leads`, `agência`, `processo` |

**Recomendação:** Manter uma agenda recorrente semanal entre o time de Parcerias & Canais e o time de SDR para alinhar qualificação de leads. Esta agenda deve ser semanal (não quinzenal ou mensal) até que o processo esteja sob controle.

**Rationale:** Quando a SDR não tem alinhamento regular com Parcerias, leads de qualidade ruim chegam aos parceiros (sem CNPJ, sem histórico de GMV, sem agência definida), o que frustra as agências e prejudica a conversão. A agenda serve para ajustar critérios de qualificação, resolver casos específicos e garantir que a SDR reforce com os leads que a Wake exige agência implementadora.

---

### BP-007 — Garantir que todas as permutas e pagamentos de projetos tenham contrato antes do início

| Campo | Valor |
|---|---|
| **Grupo/Chat** | WAKE \| Time Agência |
| **JID** | `120363419712849778@g.us` |
| **Data** | 2025-08-21 |
| **Categoria** | `best_practices` |
| **Tags** | `contrato`, `permuta`, `projeto`, `agência`, `pagamento` |

**Recomendação:** Todo acordo de pagamento (permuta de projeto, bonificação, revenue share de projeto específico) entre Wake e agência parceira deve ser formalizado em contrato antes do início das atividades. Acordos verbais ou por WhatsApp não têm validade para fins de pagamento.

**Rationale:** Em múltiplos casos, projetos foram entregues pelas agências sem contrato formal, o que gerou disputas sobre os valores devidos, atrasos de meses no pagamento e desgaste no relacionamento. O jurídico não processa pagamentos sem contrato assinado. O processo contábil da Wake exige PR (Purchase Request) e NF para qualquer saída de caixa.

---

### BP-008 — Posicionamento da certificação Wake como diferencial em relação a VTEX

| Campo | Valor |
|---|---|
| **Grupo/Chat** | WAKE \| Time Agência |
| **JID** | `120363419712849778@g.us` |
| **Data** | 2025-08-14 |
| **Categoria** | `best_practices` |
| **Tags** | `certificação`, `diferencial`, `concorrência`, `vtex`, `onboarding` |

**Recomendação:** Ao fazer onboarding de novas agências (especialmente as que já são parceiras VTEX), destacar ativamente que a certificação Wake é **gratuita** e que as agências não perdem a validade com o tempo. Na VTEX, cada certificação custa em torno de R$2.000 e precisa ser renovada periodicamente.

**Rationale:** Esse é um argumento genuinamente relevante para agências que já arcam com custos de certificação em outras plataformas. Para uma agência com 5 desenvolvedores certificados na VTEX, isso representa R$10.000+ em custos de certificação. A Wake elimina esse custo, o que aumenta a margem da agência e reduz a barreira de entrada no ecossistema.

---

### BP-009 — Separar indicadores de conversão: leads recebidos vs. fechados com a agência como operadora

| Campo | Valor |
|---|---|
| **Grupo/Chat** | WAKE \| Time Agência |
| **JID** | `120363419712849778@g.us` |
| **Data** | 2025-09-23 |
| **Categoria** | `best_practices` |
| **Tags** | `relatório`, `conversão`, `parceiro-operador`, `indicadores`, `dashboard` |

**Recomendação:** Ao apresentar taxa de conversão para parceiros, usar dois indicadores distintos: (1) "OPT fechada onde a agência foi indicada" e (2) "OPT fechada onde a agência foi o parceiro operador (implementador)". O segundo é o mais relevante para a agência, pois representa receita direta para ela.

**Rationale:** Uma OPT pode ter sido indicada por uma agência, mas fechada com outra como implementadora. Se apresentarmos apenas "conversão de indicações", a agência vai achar que fechou um projeto quando na verdade apenas indicou e outra agência executou. Essa distinção é importante para gestão de expectativas e para calcular corretamente o Revenue Share devido.

---

### BP-010 — Criar grupos WhatsApp individuais por agência parceira para comunicação direta

| Campo | Valor |
|---|---|
| **Grupo/Chat** | WAKE \| Time Agência |
| **JID** | `120363419712849778@g.us` |
| **Data** | 2025-08-14 |
| **Categoria** | `best_practices` |
| **Tags** | `comunicação`, `whatsapp`, `agência`, `relacionamento`, `grupo` |

**Recomendação:** Manter um grupo de WhatsApp individual com cada agência parceira (executivos comerciais e gerente de parcerias Wake). Este canal é essencial mesmo para agências que ainda não são parceiras formais, mas estão em negociação.

**Rationale:** O WhatsApp é o canal mais rápido para alinhar oportunidades, passar feedback de clientes, tirar dúvidas operacionais e manter o relacionamento aquecido. Agências sem grupo tendem a se distanciar, e oportunidades se perdem por falta de comunicação ágil. Um grupo WhatsApp também facilita o compartilhamento de leads em tempo real durante eventos e visitas.

---

### BP-011 — Enviar comunicados personalizados por agência, não mensagens genéricas

| Campo | Valor |
|---|---|
| **Grupo/Chat** | WAKE \| Time Agência |
| **JID** | `120363419712849778@g.us` |
| **Data** | 2025-09-23 |
| **Categoria** | `best_practices` |
| **Tags** | `comunicação`, `personalização`, `agência`, `email`, `relacionamento` |

**Recomendação:** Ao comunicar eventos, campanhas ou atualizações às agências parceiras, sempre que possível usar o nome da agência ou da pessoa no início da mensagem (ex: "Olá, FG Agência" ou "Oi Rodrigo"). Manter uma lista oficial atualizada de contatos de cada parceiro (não depender apenas do grupo geral de parceiros, que está desatualizado).

**Rationale:** Comunicações genéricas ("Olá parceiros") tendem a ser ignoradas. Comunicação personalizada demonstra atenção e aumenta a taxa de abertura/engajamento. Além disso, uma lista oficial de contatos evita que parceiros ativos não recebam comunicados (como o caso da LiveSEO que não foi convidada para um evento).

---

### BP-012 — Usar eventos e feiras como canal principal de prospecção de novas agências

| Campo | Valor |
|---|---|
| **Grupo/Chat** | WAKE \| Time Agência |
| **JID** | `120363419712849778@g.us` |
| **Data** | 2025-08-14 |
| **Categoria** | `best_practices` |
| **Tags** | `prospecção`, `eventos`, `agência`, `fórum-ecbr`, `aws-summit` |

**Recomendação:** Mapear quais agências estarão presentes nos principais eventos de e-commerce (Fórum ECBR, AWS Summit, MeLi Xperience, etc.) e priorizar encontros presenciais para iniciar conversas sobre parceria. O primeiro contato em evento tem taxa de conversão muito superior ao outbound frio.

**Rationale:** Em agosto/2025, em uma semana pós-Fórum ECBR, foram iniciadas negociações com 9 novas agências (WeDigi, Avocado Tech, Tec3, Unlock, Driven, Haltex, Monks, Mocha, Octagência) — todas por contato em evento. O movimento de agências buscando a Wake proativamente (em vez de apenas serem buscadas) foi um sinal importante de posicionamento de mercado.

---

### BP-013 — Posicionar que a Wake não exige 3 projetos entregues para receber leads (diferencial vs. VTEX)

| Campo | Valor |
|---|---|
| **Grupo/Chat** | WAKE \| Time Agência |
| **JID** | `120363419712849778@g.us` |
| **Data** | 2025-08-15 |
| **Categoria** | `best_practices` |
| **Tags** | `diferencial`, `vtex`, `leads`, `onboarding`, `programa-parceiros` |

**Recomendação:** Durante apresentações do programa de parcerias para novas agências (especialmente aquelas que já são parceiras VTEX), destacar que a Wake não exige que a agência tenha entregue 3 projetos Wake antes de começar a receber leads — diferente da VTEX que tem esse requisito. Isso reduz significativamente a barreira de entrada e o tempo para a agência começar a ter retorno.

**Rationale:** Esse ponto foi explicitamente citado por uma agência (FZ Commerce) como um dos principais motivadores para assinar com a Wake. A VTEX tem um requisito que frustra agências menores que querem crescer — a Wake oferece um caminho mais rápido para validar a parceria e gerar receita.

---

### BP-014 — Fornecer instância de demonstração Wake para agências do programa pago

| Campo | Valor |
|---|---|
| **Grupo/Chat** | WAKE \| Time Agência |
| **JID** | `120363419712849778@g.us` |
| **Data** | 2025-08-19 |
| **Categoria** | `best_practices` |
| **Tags** | `demonstração`, `instância`, `agência`, `treinamento`, `diferencial` |

**Recomendação:** Para agências que pagam mensalidade (Plano I, II ou III), disponibilizar acesso a um ambiente de demonstração Wake (instância de homologação configurada) para que os comerciais e técnicos da agência possam fazer demos para prospects, treinar funcionalidades e experimentar a plataforma.

**Rationale:** Feedback direto de uma agência parceira (B8One, via seu sócio Rotondo): sem um painel de demonstração configurado, os comerciais das agências ficam restritos a apresentar slides, o que é menos convincente do que mostrar a plataforma ao vivo. A VTEX não disponibiliza painéis de demo para parceiros — se a Wake disponibilizar, esse se torna um diferencial concreto que melhora a taxa de conversão das agências.

---

### BP-015 — Ter no Wake Score penalização para agências que não entregam projetos ou desviam clientes

| Campo | Valor |
|---|---|
| **Grupo/Chat** | WAKE \| Time Agência |
| **JID** | `120363419712849778@g.us` |
| **Data** | 2025-09-18 |
| **Categoria** | `best_practices` |
| **Tags** | `wake-score`, `penalização`, `agência`, `qualidade-ecossistema`, `contrato` |

**Recomendação:** O Wake Score deve incluir não apenas critérios de progressão (subir de nível), mas também critérios de penalização que podem resultar em rebaixamento de nível ou suspensão do carrossel. Critérios de penalização recomendados: (1) atraso grave na entrega de projetos; (2) abandono de projeto no meio (largar o cliente sem solução); (3) indicação de cliente Wake para plataforma concorrente sem justificativa legítima; (4) uso abusivo de suporte (tickets que não são de escopo da agência).

**Rationale:** Um programa saudável precisa de consequências para comportamentos prejudiciais ao ecossistema. Sem penalização, agências problemáticas continuam recebendo leads e prejudicando a experiência dos clientes, o que impacta a reputação de toda a Wake no mercado.

---

### BP-016 — Desenvolver trilha de treinamento comercial separada da trilha técnica para agências

| Campo | Valor |
|---|---|
| **Grupo/Chat** | WAKE \| Time Agência |
| **JID** | `120363419712849778@g.us` |
| **Data** | 2025-08-19 |
| **Categoria** | `best_practices` |
| **Tags** | `educação`, `treinamento`, `comercial`, `agência`, `certificação` |

**Recomendação:** Criar uma trilha de certificação específica para o perfil **comercial** das agências, separada da trilha técnica de desenvolvedores. A trilha comercial deve incluir: material de leitura, vídeos sobre posicionamento de Wake, cases de uso, comparativos com concorrentes, e uma prova mais rigorosa (não apenas respostas óbvias).

**Rationale:** A certificação atual é voltada principalmente para o perfil técnico e considerada muito fácil. Um comercial de agência que não assistiu os vídeos consegue tirar 8/10 na prova. Isso significa que o certificado não garante que o comercial está preparado para brigar em uma negociação competitiva. Uma trilha comercial mais robusta prepara melhor os vendedores das agências para competir com VTEX e outros concorrentes.

---

### BP-017 — Usar eventos das agências parceiras como canal de captação de leads

| Campo | Valor |
|---|---|
| **Grupo/Chat** | WAKE \| Time Agência |
| **JID** | `120363419712849778@g.us` |
| **Data** | 2025-09-22 |
| **Categoria** | `best_practices` |
| **Tags** | `eventos`, `leads`, `agência`, `parceiro`, `prospecção` |

**Recomendação:** Mapear quais eventos as agências parceiras organizam ou participam e garantir presença do time de Parcerias Wake nesses eventos como forma de captar leads e fortalecer a parceria. Antes de cada evento, fornecer às agências o argumento de campanha vigente (ex: 12 meses isentos para migração de Mercado Shops) para que usem como gancho de prospecção.

**Rationale:** Eventos de agências reúnem exatamente o perfil de cliente que a Wake quer — empresas de médio porte com demanda de e-commerce, já engajadas com tecnologia. O custo de aquisição é muito menor que inbound puro, e a credibilidade da agência facilita a abertura das portas. A Oásis, por exemplo, gerou múltiplos leads e OPTs a partir de eventos e visitas presenciais em Minas Gerais.

---

### BP-018 — Regra de lead inativo: desqualificar automaticamente após 3 meses sem movimentação

| Campo | Valor |
|---|---|
| **Grupo/Chat** | WAKE \| Time Agência |
| **JID** | `120363419712849778@g.us` |
| **Data** | 2025-11-28 |
| **Categoria** | `best_practices` |
| **Tags** | `lead`, `salesforce`, `qualificação`, `pipeline`, `inativo` |

**Recomendação:** Estabelecer uma regra clara no processo de pipeline: leads registrados que não avançaram para oportunidade em 3 meses são automaticamente desqualificados/zerados. Isso mantém a base de dados limpa, evita que o carrossel considere leads fantasmas e incentiva as agências a trabalharem ativamente os leads que recebem.

**Rationale:** Leads parados por meses poluem o pipeline e distorcem as métricas de conversão. Agências que registram leads mas não trabalham passam a ocupar posição no carrossel sem contribuir para o ecossistema. A regra também serve como gatilho para o time de parcerias fazer follow-up proativo antes do prazo de desqualificação.

