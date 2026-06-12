# Documentação e Suporte Wake para Parceiros

**Tipo**: how_to / best_practices  
**Tags**: `suporte`, `documentação`, `onboarding`, `sla`, `escalação`, `portal`, `parceiros`, `cs`, `audience-cs`, `audience-pam`  
**Atualizado**: 2026-06

---

## Visão Geral

A Wake suporta parceiros (agências e tech) por meio de uma combinação de canais: portal oficial de atendimento, grupos de WhatsApp dedicados por parceria, canal Slack interno (`@cscross`) e os rituais semanais de Co-Sell Acceleration e Co-Build Sessions. O onboarding formal dura 4 semanas (1 hora/semana) e o acompanhamento ongoing é mensal.

**Sobre SLA**: a Wake **não publica SLA formal partner-facing** fora de períodos sazonais. Os únicos compromissos de horário concretos encontrados no corpus são (a) a página oficial de horários de Black Friday (ver seção abaixo) e (b) prazos contratuais específicos por parceiro/integrador (ex.: Floui declarou em 2025-05-26 SLA contratual de até 15 dias para chamados de subida de produtos — `⚠ inferido` para o caso Wake-direto). Para tickets fora desses contextos, **não há comprometimento numérico publicado** — usar os gatilhos de escalação da matriz abaixo. ✓ verificado (`WAKE | Time Parcerias`, mensagem 2026-05-19 confirmando estrutura CS via `@cscross`; ausência de SLA documentado confirmada por buscas amplas).

---

## Documentação Oficial

### Onde Encontrar

| Recurso | URL |
|---|---|
| Portal de atendimento / Help Center | https://atendimento.wake.tech/hc/pt-br |
| Canais de atendimento (Wakepédia) | https://atendimento.wake.tech/hc/pt-br/articles/21406679910679-Canais-de-Atendimento-Wake-Commerce |
| Documentação técnica (API/Dev) | https://wakecommerce.readme.io/docs/canais-de-atendimento-para-parceiros |
| API REST (ReadMe) | https://wakecommerce.readme.io |
| Wake Dev Tool Kit (AI Plugin) | https://github.com/wake-engineering/ai-plugin |
| PRM de produção | https://partners.labs-wake.tech/admin/ |
| PRM staging (somente testes) | https://stg-partners.labs-wake.tech/admin |
| Listagem de agências parceiras | https://wake.tech/encontre-agencias-parceiras/ |
| Comparativo Storefront vs. Style Editor | https://wake-mkt.s3.us-east-1.amazonaws.com/branding/one_page/one-page-storefront.pdf |
| API ANYMARKET (NF-e) | https://developers.anymarket.com.br/api/v2/ky473o3qk52qb-orders-id-nfe |
| Portal de suporte ANYMARKET | https://suporte.anymarket.com.br |
| Escola de Marketplace (ANYMARKET) | https://escolademarketplace.com.br |

### Tipos de Documentação Disponível

- **Wakepédia**: base de artigos e fórum interno com novidades da plataforma, apresentada no Wake Education
- **API técnica**: documentação REST completa em `wakecommerce.readme.io` — cobre produtos, pedidos, conteúdo, banners, hotsites e integrações
- **Guia de canais de atendimento**: artigo específico no Help Center listando todos os canais e formas de contato
- **Material comercial**: one-pagers de produto (ex: Storefront vs. Style Editor), disponibilizados pelo parceiro de canais
- **Certificação Wake**: dois tipos — Vendas (comercial) e Técnica (devs); inscrição via https://forms.gle/AA6NR6vL3uEJXhDu7; turmas abertas toda segunda-feira
- **Documentação ANYMARKET**: portal próprio em `suporte.anymarket.com.br`, com artigos de how-to e release notes de API
- **Wake Dev Tool Kit**: repositório GitHub com AI Plugin para automação de desenvolvimento na plataforma via LLM/MCP

---

## Canais de Suporte

### Para Parceiros (Canal Principal)

| Canal | Contato / URL |
|---|---|
| E-mail | suporte@wake.tech |
| WhatsApp | (11) 2853-0696 |
| Portal de chamados | https://atendimento.wake.tech/hc/pt-br |
| Chat (painel de chamados) | Disponível via portal — atendimento similar ao WhatsApp |

**Uso correto por canal:**
- **Portal/e-mail**: demandas técnicas não urgentes; criar ticket com número rastreável
- **WhatsApp técnico**: dúvidas rápidas, complemento ao ticket aberto
- **Grupo de WhatsApp da parceria**: temas **comerciais** exclusivamente — não enviar demandas técnicas pelo grupo de parceria; times técnicos extras são removidos após resolução de incidentes sensíveis
- **Slack interno**: tópicos de maior duração devem sair do WhatsApp e ir para o Slack, onde as threads ficam documentadas

### Para Clientes Finais via Parceiros

O fluxo padrão é:
1. **Agência/parceiro tenta resolver** (N1) com base na documentação e no suporte técnico oficial
2. **Abre ticket** no portal com evidências do problema
3. **Escala pelo grupo WhatsApp** de parceria enviando número do ticket + impacto no cliente
4. O **Channel Manager / parceiro de canais** faz a ponte com o time técnico Wake

Para demandas ANYMARKET relacionadas a clientes finais, o suporte N1 Wake abre ticket em `suporte.anymarket.com.br` e aciona o coordenador ANYMARKET no grupo com o número do ticket quando necessário.

### Grupos de WhatsApp

Cada parceria tem um grupo dedicado (agência ↔ Wake, tech ↔ Wake). Uso esperado:
- Temas comerciais, atualizações de pipeline, alinhamento de co-sell
- Escalação urgente de tickets (informando número do ticket + impacto)
- Compartilhamento de comunicados e novidades da Wake

**Grupos de comunicação coletiva conhecidos:**
- `WAKE | Time Parcerias` — time interno de parcerias Wake
- `WAKE | Time Agência` — time interno de agências Wake
- `ANYMARKET & WAKE (TÁTICO / OPERACIONAL)` — operação conjunta Wake-ANYMARKET
- `Parcerias | Canais | Partnerships BR` — comunidade de profissionais de parcerias (não exclusivo Wake)
- Canal Slack `wake-parcerias-ingressoseventos` — pedidos e conferências de ingressos de eventos

**Regra de adição:** O grupo de WhatsApp com parceiros formais (contrato ativo + mensalidade em dia) é separado do grupo que inclui agências sem mensalidade. Apenas parceiros com contrato ativo são adicionados.

### Slack — canais e ownership

| Canal | Propósito | Owner / Quem mantém | Quando usar (vs. WhatsApp) | Status |
|---|---|---|---|---|
| `@cscross` | Mention/handle no Slack interno Wake que aciona **todo o time de CS, incluindo líderes**, para **todos os produtos Wake**. Usado para coordenar escalações cross-produto (Commerce + Anymarket + Experience). | Liderança CS Wake (estrutura criada em mai/2025) | Use quando o problema atravessa mais de um produto Wake ou quando precisa de visibilidade de liderança. Não use para 1:1 — use o canal direto do CSM. | ✓ verificado (`WAKE | Time Parcerias`, 2026-05-19: *"a estrutura de CS agora conta com o `@cscross` no Slack. Este `@` contempla todo o time, incluindo líderes, para todos os produtos Wake"*) |
| `wake-parcerias-ingressoseventos` | Canal **exclusivo** para pedidos, conferências e lista de participantes de eventos patrocinados (cotas de marketing). | Natasha (Time Marketing Wake) | Use apenas para ingressos/cotas de evento. Qualquer outro assunto vai para o canal da parceria. | ✓ verificado (`WAKE | Time Parcerias`, 2025-08-01: *"Foi criado o novo canal no Slack `wake-parcerias-ingressoseventos` onde realizaremos os pedidos, conferências e exclusivamente trataremos sobre lista de participantes dos eventos"*) |
| Canal por parceiro tech | Criado após onboarding de cada parceiro tech para co-sell e operação | PAM responsável pela parceria | Use para coordenação operacional contínua da parceria. | ⚠ inferido (padrão observado em múltiplos grupos; convenção de nome varia por parceiro) |
| Canal Salesforce | Acionado para abrir ticket de bug do Salesforce | Sales Ops | Use quando há bug no SF Community que bloqueia cadastro de lead. | ✓ verificado (`WAKE | Time Parcerias`, 2026-05-28: *"precisa abrir um ticket no canal do salesforce no slack"*) |

**Slack vs. WhatsApp — regra prática observada no corpus:**
- "O ideal é thread no Slack, detalhamento e aí sim acionamos os recursos" (`WAKE | Time Parcerias`, 2026-05-20). ✓ verificado
- "Leva no slack... aqui vai se perder" (`WAKE | Time Parcerias`, 2026-06-02). ✓ verificado
- Slack: tópicos longos com histórico, threads de resolução técnica multi-stakeholder, indicações comerciais formais
- WhatsApp: alinhamentos rápidos, comunicados urgentes, **escalação inicial com número de ticket**

**Regra geral:** comunicação interna da Wake prioriza Slack, não e-mail. E-mail é reservado para formalizações e comunicações externas (NDA, contratos, NF, Revenue Share).

### Portal de Tickets

- **Wake Commerce**: https://atendimento.wake.tech/hc/pt-br (portal principal)
- **ANYMARKET**: https://suporte.anymarket.com.br (para tickets específicos do hub de marketplace)
- Para escalar ticket urgente no ANYMARKET: enviar número do ticket + nome do cliente + data de abertura + impacto ao coordenador de suporte ANYMARKET diretamente no grupo WhatsApp
- Para escalar ticket Wake urgente: acionar o parceiro de canais/gestora de parcerias com número do ticket e impacto

---

## Fluxo de Escalação

### N1 → N2 → Especialista

```
N1 — Agência / parceiro tech
  ↓ (não consegue resolver)
N1 Wake — Abre ticket no portal de atendimento Wake
  ↓ (ticket sem retorno ou urgência)
N2 — Channel Manager / parceiro de canais aciona time técnico interno
  ↓ (problema de produto / bug / impacto alto)
Especialista — Product Owner ou time de Produto Wake
  ↓ (impacto crítico, risco de churn, falha global)
Gestão — Liderança Wake ou liderança do parceiro envolvido
```

**Exemplo prático (ANYMARKET):**
1. Wake abre ticket em `suporte.anymarket.com.br`
2. Sem retorno adequado → acionar coordenador de suporte ANYMARKET no grupo WhatsApp com número do ticket
3. Escalação para gestão ANYMARKET via liderança Wake (Diego) se necessário

**Exemplo prático (Wake Commerce — chamados Yamaha/Econverse):**
1. Agência abre ticket no portal Wake
2. Ticket sem retorno → agência envia número do chamado + impacto no grupo WhatsApp com Channel Manager
3. Channel Manager conecta diretamente com o responsável técnico Wake (mais rápido que o fluxo normal)

### Quando Escalar — matriz de severidade

A Wake não publica thresholds formais. A matriz abaixo é construída a partir de **casos reais de escalação observados no corpus** e deve ser usada como guia objetivo pelo CS/PAM.

| Severidade | Gatilho concreto (com exemplo do corpus) | Ação | Canal | Quem aciona |
|---|---|---|---|---|
| **Baixa** | Dúvida técnica de configuração / comportamento de API sem impacto em produção. Sem deadline. | Portal de chamados; aguardar fluxo normal. | `atendimento.wake.tech` | Agência / parceiro tech |
| **Média** | Ticket aberto **há ≥10 dias sem retorno adequado** (ex.: caso Cava, ticket de bug aberto 18/set sem resolução até 02/out — `WAKE | Time Agência`, 2025-10-02) **OU** prazo de cliente em risco em menos de 5 dias úteis. | Enviar número do ticket + impacto no grupo da parceria; PAM/Channel Manager aciona analista responsável. Se o ticket ainda envolve Anymarket, abrir **escalation flag** no portal. | WhatsApp grupo da parceria + portal | PAM / Channel Manager |
| **Alta** | Cliente "praticamente sem vender" (ex.: Hiper Equipamentos, `ANYMARKET & WAKE`, 2026-01-13: *"o cliente está praticamente sem vender. Conseguem me ajudar, por favor?"*) **OU** erro 500 generalizado em ambiente cliente **OU** bug afetando múltiplos clientes **OU** deadline contratual <24h **OU** apresentação para cliente já agendada. | Acionar diretamente Channel Manager / Coordenação CS no WhatsApp + Slack `@cscross`. Liderança Wake (Diego) acionada se for Anymarket. Não esperar fluxo normal do portal. | WhatsApp + Slack `@cscross` + ligação se aplicável | Channel Manager + Coordenação CS |

**Notas operacionais verificadas:**
- "Para suporte, teremos o acesso ao painel de tickets para cenários não urgentes e para cenários urgentes ou dúvidas usarão o próprio painel do ANY clicando no chat online" — regra Anymarket, ✓ verificado (`ANYMARKET & WAKE`, 2025-06-24).
- O Anymarket reforça que **mesmo para temas urgentes o fluxo via ticket no portal precisa ser mantido** em paralelo ao acionamento WhatsApp (✓ verificado, `ANYMARKET & WAKE`, 2026-02-04: *"Seron e Eliezer estão aqui para priorizar um tema urgente, mas o fluxo de solicitação de apoio precisa ser mantido por ticket no portal"*).
- Escalation flag no portal Wake pode gerar atrito com o analista — usar quando há justificativa documentada (✓ verificado, `WAKE | Time Agência`, 2025-10-02).

### Tempo de Resposta — o que existe e o que não existe

- ❌ **Não há SLA partner-facing publicado** para tickets fora de Black Friday. Buscas no corpus por "SLA", "prazo de resposta", "tempo de resposta", "respondemos em" não retornam compromisso numérico Wake-direto para suporte a parceiros.
- ✓ **Janela Black Friday** tem horário oficial documentado (ver seção Black Friday).
- ⚠ **Jurídico interno**: tarefas que envolvem Jurídico Wake podem levar "até 3 dias úteis devido às demandas internas" (✓ verificado, `Wake <> Avocado Tech`, 2025-12-11). Não é SLA de suporte técnico, mas é o único prazo numérico assumido por Wake no corpus.
- ⚠ **Floui (parceiro integrador)** declara SLA contratual próprio de **até 15 dias** para chamados de subida de produtos (✓ verificado, `Floui & Wake`, 2026-05-26) — aplica-se ao integrador, não à Wake.
- Para urgências com deadline crítico, comunicar o contexto ao abrir o ticket **e** acionar Channel Manager via WhatsApp da parceria com número do chamado + impacto. Caso histórico: erro 500 resolvido em ~10 minutos quando comunicado assim. ✓ verificado.

---

## Onboarding de Novos Parceiros

### Fase de Onboarding (4 semanas)

Estrutura padrão para agências e parceiros tech:

| Semana | Foco |
|---|---|
| 1 | KickOff: apresentação do programa, planos, carrossel, Wake Score, Revenue Share |
| 2 | Certificação: agência conclui trilha de certificação (prazo recomendado: 30 dias) |
| 3 | Wake Education Comercial: treinar o time comercial da agência em como vender Wake |
| 4 | Ativação: liberação para receber leads pelo carrossel e inclusão nos dashboards |

**Duração média completa do onboarding de cliente**: aproximadamente 90 dias.

### KickOff

O que é entregue/enviado no KickOff:
- Apresentação oficial do Programa de Parcerias (KickOff deck)
- Detalhes dos programas de Empacotamento de Storefront e Redesign de Checkout
- Liberação de acesso ao **SF Community** (Salesforce Community) — e-mail de redefinição de senha enviado ao parceiro
- Acesso ao **Ambiente de Homologação** — enviado ao contato técnico indicado
- **Badge de parceiro**: enviado após o KickOff para uso em site e materiais de marketing (validar uso com marketing Wake — contato: Matheus Coneglian)

**Duração recomendada do KickOff**: 30 minutos.

**Follow-up pós-KickOff**: verificar se o acesso ao Community foi recebido e se o ambiente de testes foi instanciado corretamente.

**Inclusão na listagem de parceiros** (https://wake.tech/encontre-agencias-parceiras/): ocorre após o KickOff; time de Partner Ops coleta logo, descrição e URL; atualização do site acontece mensalmente no **dia 15** de cada mês.

### Onboarding Técnico (parceiros tech)

Passos adicionais para parceiros tech:

1. **Burocracia em paralelo**: Julie (Ops) e liderança de Parcerias tocam documentação enquanto alinhamentos estratégicos acontecem
2. **Documentação para NDA**: contrato social, nome completo e e-mail dos responsáveis pela assinatura, nome completo + CPF + e-mail de uma testemunha
3. **Definir nível de entrada**: Essential ou Elite
4. **Verificar cláusulas de não-concorrência**: contratos de parceiros existentes (ex: ANYMARKET) têm cláusula de não-concorrência, mas aplica-se a produtos próprios — não ao ecossistema

**Prazo padrão**: primeira quinzena após conclusão dos alinhamentos internos.

### Ongoing (mensal)

- 1 hora por mês de acompanhamento
- **Wake Trends Commerce**: workshop recorrente para atualizar agências sobre novidades de produto, melhorias do trimestre e casos de uso práticos; convite disparado por WhatsApp nos grupos de agências
- **Wake Education**: evento anual (presencial + remoto) no Auditório LWSA (Rua Itapaiúna, 2434 — São Paulo/SP), das 9h às 17h; conteúdo: nova certificação, Wakepédia, atualizações e roadmap

---

## Rituais de Colaboração

### Co-Sell Acceleration (Sextas 9h–10h)

**Público**: SDRs, executivos de vendas, pré-vendas, parceiros comerciais  
**Foco**: geração de pipeline e evolução de abordagens comerciais — discussão de oportunidades reais, desenvolvimento de abordagens mais assertivas, exploração de dores dos prospects

### Co-Build Sessions (Sextas 10h–11h)

**Público**: arquitetos de solução, pré-vendas, especialistas técnicos, líderes comerciais  
**Foco**: solução, arquitetura e diferenciação — aprofundamento de cenários reais, exploração de arquitetura, produto e integrações, construção de soluções em conjunto

**Propósito declarado**: criar um modelo de colaboração contínua entre Wake e parceiros — mais próximo, mais estratégico, mais orientado a resultado. Menos teoria; mais construção e mais negócios.

### Partner Day (Quartas-feiras)

- **Local**: Escritório LWSA (Rua Itapaiúna, 2434 — Parque do Morumbi, São Paulo/SP); entrada pelo portão 4 para parceiros de fora do estado
- **Dia oficial**: quartas-feiras (confirmado a partir de ago/2025)
- **Objetivo**: aproximar parceiros dos times de CSM responsáveis pelo atendimento de clientes na base Wake
- **Agendamento**: informar ao time de Parcerias Wake com antecedência — nome, e-mail e documento de identidade para liberação na portaria
- **Treinamentos de parceiros** (ex: Omnichat): realizados presencialmente, geralmente às 14h; após 17h a adesão é menor
- **Em datas comemorativas** (Carnaval etc.): Partner Day pode ser suspenso ou apenas online — comunicação enviada com antecedência

### Pulso de Receita

- Frequência: quinzenal, com duas edições próximas ao fechamento do mês
- Participantes: liderança de Parcerias (Tech e Development) + CEO se disponível
- Documento de referência: planilha de acompanhamento com aba datada

---

## Black Friday e Períodos de Alto Volume

### Protocolo oficial de plantão (Wake → Parceiros)

Em **novembro/2025** a Wake disparou comunicado padronizado a **todos os grupos de parceiros tech e agências** (24/nov/2025, ≥30 grupos identificados — `WAKE | Time Parcerias` + grupos individuais). ✓ verificado.

**Texto-modelo do comunicado** (citado literalmente):

> "Time Wake: Teremos atendimento especial (quinta, sexta e sábado), com canais e plantão dedicados para urgências! Os detalhes de horários e contatos estão aqui: https://atendimento.wake.tech/hc/pt-br/articles/35396146551447-Horário-de-atendimento-do-suporte-durante-a-Black-Friday
>
> E a Operação de vocês? É super importante sabermos:
> 1. Terão escala/plantão em horário especial?
> 2. Qual é o contato de urgência que devemos acionar?"

**O que isso significa operacionalmente:**

| Item | Detalhe | Status |
|---|---|---|
| Janela com atendimento estendido | Quinta, sexta e sábado da semana de Black Friday | ✓ verificado |
| Documento oficial com horários e contatos | `https://atendimento.wake.tech/hc/pt-br/articles/35396146551447` (consultar antes de cada ciclo — conteúdo pode mudar ano a ano) | ✓ verificado (URL repetida em ≥30 mensagens, 2025-11-24) |
| Contato de urgência por parceiro | Cada parceiro informa nome + WhatsApp + horário de plantão ao PAM/Coordenação Wake | ✓ verificado |
| Protocolo se altera por ano | Comunicado é disparado pelo Time Parcerias na semana anterior; o texto pode variar e a URL do help center pode ser atualizada | ⚠ inferido — confirmar com Coordenação Parcerias a cada Q3 |

**Webinars temáticos pré-Black Friday** (co-produção Wake + agências parceiras):
- Wake envia convite pelo grupo WhatsApp com tema, data e contexto
- Agência confirma participação e indica o palestrante
- Wake solicita: nome, cargo, mini-bio, foto e LinkedIn do palestrante
- Transmissão via StreamYard (link restrito a participantes — não compartilhar publicamente)

**Exemplos de webinars realizados** (✓ verificado em mensagens 2025):
- Estratégias de Conteúdo e SEO para Black Friday (Econverse)
- OMS estratégico para Black Friday (Wicomm)
- Personalização e ticket médio na Black Friday (Hero Seguros)

**Plantão de fim de ano (dezembro–janeiro)**: cada parceiro informa ao time de onboarding Wake:
1. Período das férias coletivas (datas de início e fim)
2. Se haverá escala/plantão e em qual horário
3. Contato de urgência (nome + WhatsApp) para acionamento em situações críticas

---

## Problemas Comuns de Suporte e Como Resolver

### Acesso ao Treinamento Wake não chega após inscrição

**Causa**: turmas liberadas toda segunda-feira. Inscrição feita após a liberação da semana → acesso chega na próxima segunda.  
**Solução**: aguardar segunda-feira; verificar spam; se não chegar, acionar o parceiro de canais para liberação manual.

### Loop de login na plataforma de treinamento Wake (usuário/senha inválidos após reset)

**Causa**: cadastro com problema no sistema — e-mail de turma anterior ou cadastro não ativado corretamente.  
**Solução**: acionar o parceiro de canais para solicitar ativação manual do usuário na plataforma de treinamento. A Wake verifica internamente e envia novo e-mail de acesso com as credenciais corretas.

### Acesso ao admin da loja não liberado após assinatura de contrato

**Fluxo correto**: Assinatura → Sales Ops (revisão) → Onboarding (analista atribuído + liberação de acessos)  
**Solução**: verificar via CNPJ no sistema interno se o contrato passou pela revisão de Sales Ops; após a transição Sales Ops → Onboarding, acionar **Mayara Roberto** (coordenadora de Onboarding — ✓ verificado: `Wake <> Benova` 2025-10-07 *"Mayara (coordenadora de onboarding)"*; `Wake <> Ge Commerce` 2025-11-26 *"May (coordenadora de onboarding)"*; `Vindi + Wake` 2025-12-17 *"@Mayara Roberto [Wake - Commerce] coordenadora de onboarding"*).

**Quando Mayara não estiver disponível** (férias, ausência, indisponibilidade):
1. Acionar **Eduardo Manzoli** — gerente do time de Onboarding/CS Commerce (✓ verificado, `Vindi + Wake` 2025-12-17).
2. Em paralelo, contatar a Coordenação CS (qualquer um dos coordenadores de CS Commerce: **Andréa Fonseca**, **Math De Almeida**, **Daniel Guglielmello** — ✓ verificado, mesma mensagem) via `@cscross` no Slack para coordenar a liberação interna.
3. Se for crítico (cliente travado para go-live), abrir thread em `@cscross` mencionando líderes — o handle Slack contempla liderança CS (✓ verificado).

⚠ **Não há backup nomeado formal de "coordenadora de Onboarding" no corpus**. A escalada acima é construída a partir da estrutura hierárquica observada (gerente Eduardo + 3 coordenadores de CS) e não a partir de um runbook publicado pela Wake. Validar antes de comunicar ao parceiro.

### Botão de cadastro de lead não disponível no Community (SF)

**Solução**: enviar dados do lead diretamente ao parceiro de canais via WhatsApp — parceiro de canais realiza o cadastro no CRM e garante que a indicação fica registrada em nome da agência.  
**Dados mínimos**: nome da empresa, nome do contato, e-mail, telefone, site (se disponível), segmento/faturamento (se conhecido).

### Bug no Salesforce ao criar lead: "ID do proprietário não pode ficar em branco"

**Causa**: bug no fluxo de automação do SF Community que tentava atribuir o registro a um proprietário com base em regras de parceria, mas o campo ficava em branco.  
**Solução temporária**: indicações devem ser enviadas diretamente à gestora de parcerias Wake via WhatsApp para cadastro manual enquanto a correção não é aplicada.

### Erro 500 generalizado no ambiente de cliente

**Ação imediata**: comunicar via WhatsApp do grupo ao parceiro de canais, informar se há cliente/apresentação na iminência (acrescenta urgência), aguardar confirmação de que o time está analisando. Tempo médio de resolução observado: ~10 minutos.

### Chamados sem retorno comprometendo cronograma

**Escalação efetiva**: enviar número do chamado + impacto no cliente no grupo WhatsApp. O Channel Manager faz a ponte com o time técnico de forma mais rápida do que o fluxo normal de suporte (verificado em caso Yamaha com agência Econverse).

### Instabilidade global (AWS)

- AWS Health Dashboard: https://health.aws.amazon.com/health/status
- ANYMARKET Status Page: https://status.anymarket.com.br/

---

## Contatos e Responsáveis

Os seguintes responsáveis foram mencionados nas fontes com funções verificadas:

| Nome | Função |
|---|---|
| Julie (Letícia) | Ops — contratos, faturamento parceiros, inadimplência |
| Mayara Roberto | Coordenadora de Onboarding (Wake - Commerce). Backup: Eduardo Manzoli (gerente) + coordenadores CS (Andréa Fonseca / Math De Almeida / Daniel Guglielmello) via `@cscross`. ✓ verificado |
| Eduardo Manzoli | Gerente Wake - Commerce (CS + Onboarding). ✓ verificado |
| Matheus Coneglian | Content marketing Wake; ponto de contato para co-marketing e anúncios de parceria |
| Mari Pellegrino | PMM Wake |
| Renato (saiu, Nina e Mari assumiram) | Anteriormente Partner Ops; responsável por KickOff de parceiros tech |
| Geovanio Filho / Herbert Santos | Infra — configuração de acesso ao QuickSight |
| Henrique Maia | Onboarding ANYMARKET — booking: https://outlook.office.com/bookwithme/user/7a734b912e9d4e65ab41b9195fc72b96%40db1.com.br?anonymous&ismsaljsauthenabled — ✓ verificado ativo: agenda usada em onboardings em ago/set/out 2025 (`ANYMARKET & WAKE`, várias mensagens 2025-08 a 2025-10) |
| Lucas Rodrigues | Onboarding ANYMARKET — booking: https://outlook.office.com/bookwithme/user/0ace20555a9e4de9ac208c7d16c16a5d%40db1.com.br/meetingtype/9QPMQRf5_UyOvZdCylTahw2?anonymous&ismsaljsauthenabled — ✓ verificado ativo: link compartilhado em par com o de Henrique Maia em 2025-08-19 (`ANYMARKET & WAKE`); agenda usada em set/2025 |
| Ana Aires | Omnichat — contato comercial/técnico: ana.aires@omni.chat |
| Kleber Albertini | Omnichat — comercial: kleber.albertini@omni.chat |
| Natasha | Responsável por eventos e ingressos; centraliza canal Slack `wake-parcerias-ingressoseventos` — ✓ verificado (`WAKE | Time Parcerias`, 2025-08-01) |
| Andréa Fonseca / Math De Almeida / Daniel Guglielmello | Coordenadores CS (Wake - Commerce) — acionar via `@cscross` para escalações que precisam de coordenação ou backup do time de Onboarding. ✓ verificado |
| Angélica / Camila | RH — abertura de vagas e movimentações internas (Gupy) |

---

## Fontes

**Arquivos KB utilizados:**
- `/docs/kb-parcerias.md` — grupo WAKE | Time Parcerias
- `/docs/kb-parcerias-canais.md` — grupo Parcerias | Canais | Partnerships BR
- `/docs/kb-agencia.md` — grupo WAKE | Time Agência
- `/docs/kb-coordenacao.md` — grupo WAKE | Coordenação
- `/docs/kb-anymarket-wake.md` — grupo ANYMARKET & WAKE (TÁTICO / OPERACIONAL)
- `/docs/kb-comercial-cross.md` — grupo Comercial Cross WAKE
- `/docs/kb-parceiros-agencia.md` — grupos Wake ↔ Nairuz, Catus, Série//A
- `/docs/kb-parceiros-tech.md` — grupos Wake & Omnichat, FG + Wake, Leanwork, Eficaz
- `/docs/kb-parceiros-wave5a.md` — grupos SAVVI, Dotkom, A&EIGHT, Omni360
- `/docs/kb-parceiros-wave5b.md` — grupos BonifiQ, DRIVEN, Pagaleve, Mastercard
- `/docs/kb-parceiros-wave5c.md` — grupos Avocado Tech, Wicomm, Econverse, HERO
- `/docs/wake-partnership-faq.md` — FAQ consolidado extraído dos grupos de parcerias
- `/docs/kb-opus-analysis.md` — análise estratégica dos 668 artigos KB (referência: Cluster 14 — Suporte oficial e escalation)

---

## Verificação

| Claim | Status | Query (chave) | Evidência (grupo, data) |
|---|---|---|---|
| Wake não publica SLA partner-facing formal | ✓ verificado | `LIKE '%SLA%' OR LIKE '%prazo de resposta%' OR LIKE '%tempo de resposta%' OR LIKE '%respondemos em%'` (40 hits, nenhum compromisso numérico Wake-direto fora de Black Friday) | Múltiplos grupos, 2025–2026 — ausência de SLA |
| Black Friday tem horário oficial e canais dedicados (quinta/sexta/sábado) | ✓ verificado | `LIKE '%Black Friday%' AND (LIKE '%plantão%' OR LIKE '%escala%' OR LIKE '%suporte%')` (≥30 hits idênticos) | `Wake & Stayfilm` + ≥29 outros grupos, 2025-11-24 |
| URL oficial Black Friday: `atendimento.wake.tech/.../35396146551447` | ✓ verificado | `LIKE '%35396146551447%'` (≥30 hits) | `WAKE | Time Parcerias` + 29 grupos, 2025-11-24 |
| `@cscross` cobre todo o time CS Wake incluindo líderes, todos os produtos | ✓ verificado | `LIKE '%cscross%'` (1 hit primário) | `WAKE | Time Parcerias`, 2026-05-19 |
| `wake-parcerias-ingressoseventos` é exclusivo para ingressos/eventos (owner: Natasha) | ✓ verificado | `LIKE '%wake-parcerias-ingressoseventos%'` (2 hits) | `WAKE | Time Parcerias`, 2025-08-01 e 2025-09-05 |
| Mayara Roberto é coordenadora de Onboarding | ✓ verificado | `LIKE '%Mayara%' AND LIKE '%onboarding%'` (5+ hits) | `Wake <> Benova` 2025-10-07; `Wake <> Ge Commerce` 2025-11-26; `Vindi + Wake` 2025-12-17 |
| Backup para Mayara: Eduardo Manzoli (gerente) + 3 coordenadores CS | ⚠ inferido | mesma busca acima — estrutura hierárquica observada, não runbook publicado | `Vindi + Wake` 2025-12-17 (lista de coordenadores) |
| Henrique Maia booking link ativo | ✓ verificado | `LIKE '%Henrique Maia%' OR LIKE '%bookwithme%'` (20+ hits, agenda usada repetidamente ago–out/2025) | `ANYMARKET & WAKE`, 2025-08 a 2025-10 |
| Lucas Rodrigues booking link ativo | ✓ verificado | mesma busca; link compartilhado em par com o de Henrique | `ANYMARKET & WAKE`, 2025-08-19 e 2025-09-22 |
| Matriz de severidade — caso Cava (ticket 14 dias sem resposta) | ✓ verificado | revisão direta do thread | `WAKE | Time Agência`, 2025-10-02 |
| Matriz de severidade — caso Hiper Equipamentos ("cliente praticamente sem vender") | ✓ verificado | `LIKE '%urgência%' OR LIKE '%urgente%' AND LIKE '%ticket%'` | `ANYMARKET & WAKE`, 2026-01-13 |
| Anymarket exige ticket no portal mesmo para urgências | ✓ verificado | mesma busca | `ANYMARKET & WAKE`, 2026-02-04 e 2025-06-24 |
| Jurídico Wake — "até 3 dias úteis" | ✓ verificado | `LIKE '%dias úteis%' AND LIKE '%retorno%'` | `Wake <> Avocado Tech`, 2025-12-11 |
| Floui SLA contratual 15 dias (parceiro, não Wake) | ✓ verificado | mesma busca | `Floui & Wake`, 2026-05-26 |
| Slack > WhatsApp para threads longas (regra de uso) | ✓ verificado | revisão direta dos threads `WAKE | Time Parcerias` | `WAKE | Time Parcerias`, 2026-05-20 e 2026-06-02 |
| "Ticket sem retorno após 5 dias úteis" como gatilho (regra original do artigo) | ❌ sem fonte | Nenhum compromisso numérico de "5 dias" encontrado no corpus para Wake-direto | substituído pela matriz de severidade com gatilhos verificáveis |
| Comunicação interna prioriza Slack sobre e-mail | ⚠ inferido | observado em múltiplos threads, sem política escrita citada | inferido de uso recorrente |
