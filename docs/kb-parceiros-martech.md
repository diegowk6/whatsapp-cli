# KB Extract — Parceiros Martech (Compra Rápida, After Click, Stayfilm, Revi, Vurdere)
> Extracted: 2026-06-04

---

## Wake & Compra Rápida (`120363406685999927@g.us`)

---

### HOW_TO-1 — Processo de Onboarding de Novo Parceiro Tech: Criação de Ambiente Trial

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Wake & Compra Rápida |
| **JID** | `120363406685999927@g.us` |
| **Data** | 2026-03-30 |
| **Categoria** | `how_to` |
| **Tags** | `onboarding`, `trial`, `ambiente`, `acesso`, `parceiro-tech` |

Ao iniciar a parceria técnica, a Wake provisiona um ambiente trial para o parceiro realizar a integração. O fluxo padrão é:

1. O time de Parcerias da Wake solicita a criação do ambiente trial (modelo padrão vinculado ao contrato essencial).
2. A instância é replicada — o processo pode levar algumas horas.
3. Após a replicação, o time Wake cria o usuário administrador com o e-mail indicado pelo parceiro.
4. O acesso ao admin é enviado por e-mail, junto com link de artigo explicando como gerar o token de API.
5. O parceiro acessa o admin em `https://[parceiro].fbits.app/` e a loja (hlog) em `https://[parceiro].fbits.store/`.

**Ponto de atenção:** O formato do e-mail deve ser conferido antes do envio do acesso. Erros como `.com` no lugar de `.co` impedem o recebimento do convite. Caso o e-mail não chegue, verificar spam e confirmar o domínio correto.

**Nome de usuário no login:** O campo "usuário" do admin segue o padrão `[parceiro]\[nomeUsuario]`. Ex: `partnercomprarapida\lucasescobar`.

**Token de API:** Dentro do admin, é possível gerar o token de acesso para consumir a API. O link do artigo de orientação é enviado no mesmo e-mail de acesso.

---

### TROUBLESHOOTING-1 — Usuário do Ambiente Trial Desativado Automaticamente

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Wake & Compra Rápida |
| **JID** | `120363406685999927@g.us` |
| **Data** | 2026-04-02 |
| **Categoria** | `troubleshooting` |
| **Tags** | `trial`, `usuário-desativado`, `acesso`, `ambiente` |

**Problema:** O usuário criado no ambiente trial retorna `Usuário desativado!` ao tentar fazer login.

**Causa:** O sistema desativa automaticamente usuários de trial após o período de inatividade ou expiração do prazo.

**Resolução:** Solicitar ao time Wake (via grupo WhatsApp ou pelo canal de suporte) a reativação do usuário. O time pode reativar diretamente pelo admin. Após a reativação, o parceiro deve tentar acessar pelo link atualizado da loja.

**Prevenção:** Manter comunicação ativa durante o período de trial. Se a parceria avançar (ex: após VTEX Day, férias do time), solicitar proativamente a reativação do ambiente antes de retomar a integração.

---

### FAQ-1 — Qual canal usar para dúvidas técnicas durante a integração?

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Wake & Compra Rápida |
| **JID** | `120363406685999927@g.us` |
| **Data** | 2026-05-12 |
| **Categoria** | `faq` |
| **Tags** | `suporte`, `canal-atendimento`, `ticket`, `integração-técnica` |

**Pergunta:** Qual o melhor canal para tirar dúvidas técnicas da integração com a Wake?

**Resposta:** Para dúvidas técnicas durante a integração, o caminho correto é abrir um ticket pelo canal oficial de atendimento aos parceiros:

- **Central de atendimento:** https://atendimento.wake.tech/hc/pt-br/requests/new
- **Documentação técnica para parceiros:** https://wakecommerce.readme.io/docs/canais-de-atendimento-para-parceiros

O time de Parcerias no grupo WhatsApp pode direcionar e fazer o apoio de triagem, mas demandas técnicas específicas exigem registro via ticket para que o time técnico seja acionado formalmente.

---

### TROUBLESHOOTING-2 — Erro de CSS no Ambiente Trial (Certificado de Segurança Pendente)

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Wake & Compra Rápida |
| **JID** | `120363406685999927@g.us` |
| **Data** | 2026-03-31 |
| **Categoria** | `troubleshooting` |
| **Tags** | `trial`, `css`, `certificado-ssl`, `front-end`, `ambiente` |

**Problema:** No ambiente trial recém-criado, o front da loja apresenta erro de CSS.

**Causa:** O certificado de segurança (SSL) ainda não foi aplicado no front do ambiente pelo time de infraestrutura da Wake.

**Impacto:** Visual comprometido, mas as funcionalidades de back-end estão disponíveis: configuração do token, criação de usuários adicionais e manipulação da API funcionam normalmente.

**Resolução:** O time de infra da Wake aplica o certificado em paralelo. O parceiro pode prosseguir com a configuração do token e testes de API enquanto aguarda a correção do front.

---

### HOW_TO-2 — Processo de Onboarding com Responsável Dedicado (Beatriz)

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Wake & Compra Rápida |
| **JID** | `120363406685999927@g.us` |
| **Data** | 2026-05-12 |
| **Categoria** | `how_to` |
| **Tags** | `onboarding`, `parceiro-tech`, `kickoff`, `entregáveis` |

Após a formalização da parceria, um responsável de onboarding do time de Parcerias Wake é designado e adicionado ao grupo WhatsApp. O processo padrão inclui:

1. **Boas-vindas:** O responsável se apresenta e agenda um kickoff.
2. **Kickoff:** Reunião de alinhamento sobre entregáveis da jornada de parceria.
3. **Materiais enviados por e-mail:** Apresentação + documentos complementares.
4. **Agenda técnica:** Reunião com o time técnico Wake para avançar a integração (agendada logo após o kickoff).
5. **Atualização do site Wake:** Todo dia 15 de cada mês, o site de parceiros é atualizado. Para inclusão, o parceiro deve enviar: descrição "quem somos", 1-2 especialidades e o manual de marca com o logo.

---

### TROUBLESHOOTING-3 — Ticket de Integração Travado — Como Desbloquear

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Wake & Compra Rápida |
| **JID** | `120363406685999927@g.us` |
| **Data** | 2026-05-18 |
| **Categoria** | `troubleshooting` |
| **Tags** | `ticket`, `suporte`, `integração`, `desbloqueio`, `escalonamento` |

**Problema:** Ticket de suporte aberto (ex: #280064) sem evolução — time ainda "analisando se pode fazer ou não".

**Contexto:** Na integração da Compra Rápida, o time enfrentou um bloqueio relacionado à listagem de formas de pagamento (ticket #282901), que o suporte classificou como necessitando de "solicitação de melhoria".

**Ação de desbloqueio:**
- Acionar o Account Manager responsável diretamente no grupo WhatsApp para dar visibilidade ao ticket.
- O AM pode solicitar internamente reanálise do caso para buscar caminhos alternativos à solicitação de melhoria.
- Se não houver retorno via grupo, o AM deve marcar uma agenda com o time técnico e o parceiro para alinhar em conjunto.

**Canal para escalonamento:** Grupo WhatsApp de parceria + menção direta ao AM responsável (ex: `@233216808083577`).

---

### HOW_TO-3 — Calendário e Canais de Atendimento do Programa de Parcerias Wake

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Wake & Compra Rápida |
| **JID** | `120363406685999927@g.us` |
| **Data** | 2026-03-03 |
| **Categoria** | `how_to` |
| **Tags** | `programa-parcerias`, `canais-atendimento`, `wake-partner-day`, `contatos` |

**Wake Partner Day:** Evento semanal presencial na sede Wake/LWSA (Rua Itapaiúna, 2434 - Parque do Morumbi, SP). A partir de agosto/2025, mudou de quintas para **quartas-feiras**. Objetivo: aproximar parceiros dos times de CS, vendas e marketing.

**Contatos do time de Parcerias Wake:**

| Papel | Nome | Contato |
|---|---|---|
| AM Agências | Ana Vicente | +55 14 99770-3611 |
| AM Checkout/Pagamento/Frete | Gabriel Cotrim | +55 11 96925-1769 |
| Tech Partnerships Lead / AWS Alliance | Rafael Correia | +55 11 93800-6017 |
| Partner Development | Letícia Cena | +55 11 97433-3718 |
| Suporte técnico parceiros | — | suporte@wake.tech |

**Documentação técnica:** https://wakecommerce.readme.io/docs/canais-de-atendimento-para-parceiros

---

## After Click <> Wake (`120363420659748080@g.us`)

---

### HOW_TO-4 — Processo de Inclusão de Parceiro no Site e Materiais de Marketing Wake

| Campo | Valor |
|---|---|
| **Grupo/Chat** | After Click <> Wake |
| **JID** | `120363420659748080@g.us` |
| **Data** | 2025-07-04 |
| **Categoria** | `how_to` |
| **Tags** | `marketing`, `logo`, `site-parceiros`, `brand-assets`, `onboarding` |

Para que um parceiro seja incluído em materiais de marketing da Wake (site de parceiros, convites de eventos, newsletter), o time de Parcerias solicita os seguintes ativos de marca:

1. **Logo em múltiplos formatos:** SVG, PNG com fundo transparente. Variações com fundos diferentes (preto, branco, verde, cinza) são bem-vindas.
2. **Preferência de uso:** A Wake prioriza o logo com as cores reais da marca (ex: check verde da After Click), não versão monocromática.
3. **Manual de marca:** Enviado via Google Drive compartilhado (e-mail de acesso fornecido pelo parceiro).

**Ponto de atenção de convites:** Para eventos com fundo preto, verificar se o logo tem versão totalmente branca (ícone + nome) para garantir visibilidade.

---

### HOW_TO-5 — Wake Partner Day: Funcionamento e Como Participar

| Campo | Valor |
|---|---|
| **Grupo/Chat** | After Click <> Wake |
| **JID** | `120363420659748080@g.us` |
| **Data** | 2025-09-02 |
| **Categoria** | `how_to` |
| **Tags** | `wake-partner-day`, `evento-presencial`, `LWSA`, `networking` |

O **Wake Partner Day** é um encontro semanal (toda quarta-feira) na sede da LWSA onde parceiros podem:

- Trabalhar de lá durante o dia (coworking informal).
- Interagir com os times de CS, vendas e marketing da Wake.
- Fazer networking com outros parceiros do ecossistema.

**Endereço:** R. Itapaiúna, 2434 - Parque do Morumbi, SP (2º Andar). Buscar "LWSA" no Waze — fica no interior de uma comunidade.

**Acesso ao Wi-Fi:** Enviar e-mail ao responsável de parcerias com antecedência para liberação do acesso.

**Participação remota:** Disponível quando há transmissão do evento, mediante inscrição no formulário específico.

**Cancelamentos e mudanças de data:** Comunicados via grupo WhatsApp e canal de parceiros. Casos recorrentes: feriados, eventos externos da ECBR, datas especiais.

---

### HOW_TO-6 — Modelo de Negócio para Indicações Cruzadas entre Parceiros Wake

| Campo | Valor |
|---|---|
| **Grupo/Chat** | After Click <> Wake |
| **JID** | `120363420659748080@g.us` |
| **Data** | 2025-12-12 |
| **Categoria** | `how_to` |
| **Tags** | `indicação`, `modelo-de-negócio`, `parceria-cruzada`, `fulfill`, `agência` |

Quando um parceiro recebe uma indicação de lead de outro parceiro ou da Wake, é necessário formalizar o modelo de negócio para que:

- A Wake permaneça no fluxo da negociação.
- A indicação seja interessante para todos os lados (Wake + parceiro indicador + parceiro receptor).

**Processo sugerido:**
1. A indicação é compartilhada no grupo WhatsApp de parceria.
2. O time de Parcerias alinha internamente o modelo (ex: comissão, co-sell, indicação simples).
3. Antes de seguir com o cliente, os parceiros envolvidos alinham o "modelo de negócio entre nós" para evitar ruído depois.

**Exemplo prático:** Indicação da agência "Simples Inovação" para a After Click (serviço de Fulfill) — Diego Santos (Wake) intermediou e destacou a necessidade de amarrar o modelo antes de avançar com o cliente.

---

### BEST_PRACTICES-1 — Gestão da Parceria After Click: Padrões de Relacionamento

| Campo | Valor |
|---|---|
| **Grupo/Chat** | After Click <> Wake |
| **JID** | `120363420659748080@g.us` |
| **Data** | 2025-07-04 |
| **Categoria** | `best_practices` |
| **Tags** | `gestão-parceria`, `contrato`, `relacionamento`, `programa-parcerias` |

Boas práticas observadas na gestão da parceria After Click <> Wake:

- **Formalizar o contrato cedo:** Mesmo que já haja atuação conjunta, iniciar o processo de contrato formal logo ao abrir o grupo de WhatsApp. A falta de contrato formal foi identificada como gap logo no início.
- **Apresentar o Programa de Parcerias formalmente:** Agendar uma reunião específica para apresentar o programa (entregáveis, categorias, benefícios) mesmo para parceiros que já têm relacionamento comercial.
- **Pesquisa de satisfação pós-evento:** Após grandes eventos (Wake Summit, Wake Creators), enviar formulário de satisfação para parceiros para coleta de feedback e melhoria das próximas edições.
- **Gestão de transição de responsáveis:** Ao sair da empresa, o responsável de parcerias deve comunicar no grupo, indicar o substituto e garantir a continuidade do relacionamento.

---

### FAQ-2 — O que é o Wake Trends Commerce e quando acontece?

| Campo | Valor |
|---|---|
| **Grupo/Chat** | After Click <> Wake |
| **JID** | `120363420659748080@g.us` |
| **Data** | 2025-08-12 |
| **Categoria** | `faq` |
| **Tags** | `wake-trends-commerce`, `treinamento`, `parceiros`, `produto` |

**O que é:** Workshop online exclusivo para agências e parceiros tech Wake. Apresenta as principais novidades do Wake Commerce do trimestre, com foco em:
- Funcionalidades lançadas e melhorias.
- Casos de uso práticos (times de Produto, CS e Product Marketing).
- Espaço para dúvidas e troca de experiências.

**Objetivo:** Alinhar o discurso dos parceiros com o time Wake para que consigam gerar valor real nas conversas com clientes e prospects.

**Periodicidade:** Realizado periodicamente (exemplo de edição: 14/08/2025, 10h–11h30).

**Acesso:** Link Google Meet enviado pelo time de Parcerias.

---

## Wake & Stayfilm (`120363408925526036@g.us`)

---

### HOW_TO-7 — Processo de PoC (Prova de Conceito) com Parceiro de Vídeo Personalizado

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Wake & Stayfilm |
| **JID** | `120363408925526036@g.us` |
| **Data** | 2025-06-23 |
| **Categoria** | `how_to` |
| **Tags** | `poc`, `prova-de-conceito`, `stayfilm`, `vídeo-personalizado`, `cs`, `clientes` |

Para viabilizar os primeiros cases de parceria entre Wake e Stayfilm (tecnologia de vídeo personalizado 1:1 em escala):

1. **Seleção de contas:** O time Wake (Experience/CS) seleciona 5 contas estratégicas com base em lista enviada pelo parceiro. Critério prioritário: empresas com alto faturamento capazes de pagar o mínimo de 10 mil vídeos + setup (aprox. R$ 55K de largada).
2. **Treinamento de CS Wake:** Realizar sessão de atualização/treinamento com o time de CS da Wake para que consigam abordar a solução Stayfilm com os clientes da base.
3. **Abordagem sniper:** Para garantir conversões, abordar um funil de 8-10 contas qualificadas, mirando fechar 4 clientes. Foco em clientes que possam gerar retorno de pelo menos 10% em conversão.
4. **Fee mensal:** Cobrado ao final do 1º mês trabalhado. Boleto gerado para pagamento no mês seguinte.
5. **Incentivos:** Avaliar modelo de incentivos por resultado para engajar o time de CS da Wake.

**Modelo de parceria:** Stayfilm paga fee mensal à Wake + comissões por resultado.

---

### BEST_PRACTICES-2 — Estratégia "Sniper" para Parceiros com Solução de Alto Ticket

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Wake & Stayfilm |
| **JID** | `120363408925526036@g.us` |
| **Data** | 2025-10-28 |
| **Categoria** | `best_practices` |
| **Tags** | `estratégia-comercial`, `sniper`, `alto-ticket`, `poc`, `pipeline` |

Para parceiros com soluções de alto ticket (acima de R$ 50K de contrato mínimo), a estratégia "sniper" é mais efetiva do que prospecção em volume:

- **Princípio:** Menos reuniões, mais qualificação. Focar em poucas contas com alta probabilidade de conversão e fit claro.
- **Funil orientativo:** Para fechar 4 clientes, abordar 8-10 contas qualificadas (sendo "bem sniper e otimista").
- **Critério de qualificação:** Empresas com volume de transações e faturamento suficiente para justificar o investimento mínimo da solução.
- **Papel do CS Wake:** CS identifica abertura com o cliente e faz a ponte para reunião com o parceiro. Não adianta abordar clientes sem slot disponível ou sem demanda identificada.
- **Resultado esperado:** Cases que demonstrem ao menos 10% de retorno para o cliente, viabilizando expansão da parceria.

---

### HOW_TO-8 — Entregáveis de Marketing para Parceiros Tech Wake (Press Release e Blog)

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Wake & Stayfilm |
| **JID** | `120363408925526036@g.us` |
| **Data** | 2025-11-07 |
| **Categoria** | `how_to` |
| **Tags** | `marketing`, `press-release`, `blog`, `plano-lançamento`, `parceiro-tech` |

O Plano de Lançamento de Parceiros Tech Wake inclui entregáveis de marketing com datas fixas. Para o Press Release e publicação no Blog:

**O que o parceiro precisa enviar:**
- Press release ou material institucional com: quem é a empresa, o que faz/principais serviços, benefícios para clientes, citação de executivo, dados de mercado ou pesquisa relacionados ao setor.
- Formato aceito: `.docx` simples (sem necessidade de diagramação).

**O que a Wake produz:**
- Publicação no Blog Wake com texto educativo agnóstico (sem propaganda direta), focado em autoridade e conteúdo.
- Publicação em redes sociais (LinkedIn, Instagram).
- Divulgação em imprensa especializada (ex: ECBR).
- Notificação em newsletter.

**Prazos típicos (exemplo real):**
- Publicação Blog + Redes Sociais: dia 19 do mês
- Imprensa: dia 23 do mês
- Newsletter: mês seguinte

**Ajuste de descrição no site:** O parceiro pode (e deve) revisar sua descrição no site Wake. Enviar sugestão de texto ao responsável de Parcerias Ops.

---

### HOW_TO-9 — Participação de Parceiros em Podcasts e Painéis de Eventos Wake

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Wake & Stayfilm |
| **JID** | `120363408925526036@g.us` |
| **Data** | 2025-08-12 |
| **Categoria** | `how_to` |
| **Tags** | `podcast`, `wake-summit`, `painel`, `visibilidade`, `parceiro-tech` |

Parceiros tech Wake têm oportunidade de participar de podcasts e painéis em eventos do ecossistema:

**Wake Summit Podcast:** Entrega premium do Wake Summit. Parceiros podem negociar participação mesmo sem cota de patrocínio no evento. Ideal para parceiros com posicionamento claro (ex: vídeo personalizado + IA + personalização).

**Formato de participação:**
1. Parceiro propõe tema ou o time Wake sugere pauta alinhada com tendências (ex: Gartner citou AI, vídeo e personalização como os 3 principais pontos do futuro da tecnologia).
2. Definir data, convidado do lado do parceiro (ex: co-founder/CSO) e representante Wake.
3. O link do estúdio de transmissão (StreamYard) é enviado somente aos participantes — não deve ser compartilhado publicamente.
4. Participantes entram ~30 min antes para teste de microfone e ajustes.

**Estratégia de conteúdo:** Alinhar tema com pautas em alta (personalização, IA, vídeo) para garantir relevância junto a veículos como CMO Summit, Resenha Company e ECBR.

---

### FAQ-3 — Como funciona o suporte técnico da Wake durante a Black Friday?

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Wake & Stayfilm |
| **JID** | `120363408925526036@g.us` |
| **Data** | 2025-11-24 |
| **Categoria** | `faq` |
| **Tags** | `black-friday`, `suporte`, `plantão`, `atendimento-especial` |

**Pergunta:** Como funciona o suporte da Wake durante a Black Friday?

**Resposta:** A Wake disponibiliza atendimento especial na semana da Black Friday (quinta, sexta e sábado), com canais e plantão dedicados para urgências.

**Detalhes completos:** https://atendimento.wake.tech/hc/pt-br/articles/35396146551447-Hor%C3%A1rio-de-atendimento-do-suporte-durante-a-Black-Friday

**O que a Wake solicita dos parceiros:**
1. Confirmação se haverá escala/plantão especial do lado do parceiro.
2. Contato de urgência do parceiro para acionamento durante o período.

Esse alinhamento bilateral é necessário para garantir agilidade na resolução de incidentes que envolvam a integração entre plataformas.

---

### BEST_PRACTICES-3 — Gestão de Parceria com Acompanhamento de Pipeline e Cases

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Wake & Stayfilm |
| **JID** | `120363408925526036@g.us` |
| **Data** | 2025-09-25 |
| **Categoria** | `best_practices` |
| **Tags** | `pipeline`, `gestão-parceria`, `case`, `cs`, `treinamento` |

Boas práticas para manter a parceria ativa e gerar resultados:

- **Revisão periódica do pipeline:** Quando a parceria completa meses sem nenhuma reunião com cliente ou proposta gerada, fazer uma reunião de revisão de estratégia — não apenas follow-up de status.
- **Pivotagem de estratégia:** Se a abordagem original não está gerando resultados (ex: focar em Experience mas sem resultados), pivotar para a base Commerce com apoio do time de CS.
- **Treinamento recorrente de CS:** Após cada período de inatividade ou entrada de novo CS, realizar nova sessão de treinamento online para garantir assimilação da oferta.
- **Descrição no site Wake:** Manter atualizada — verificar periodicamente se as especialidades e descrição refletem o momento atual da empresa.
- **Webinars de Black Friday:** Parceiros técnicos são convidados para webinars temáticos (ex: webinars Black Friday 2025). Participar como speaker posiciona a solução junto à base de clientes Wake.

---

## Revi + Wake (`120363422405368644@g.us`)

---

### HOW_TO-10 — Integração de Carrinho Abandonado via Pixel/Tag (Abordagem Paliativa)

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Revi + Wake |
| **JID** | `120363422405368644@g.us` |
| **Data** | 2025-09-15 |
| **Categoria** | `how_to` |
| **Tags** | `carrinho-abandonado`, `pixel`, `tag`, `gtm`, `integração`, `whatsapp` |

Enquanto a API nativa de eventos de carrinho abandonado da Wake não está disponível (prevista para Q4 2025), existe uma abordagem paliativa via pixel/tag:

**Como funciona:**
- Uma tag é adicionada via Gestor de Scripts da Wake ou via Google Tag Manager (GTM).
- A tag monitora eventos no checkout sem coletar dados sensíveis.
- Detecta quando uma sessão de checkout foi iniciada mas o pedido não foi finalizado (sinal de abandono).
- Dados de identidade do cliente (nome, telefone, produtos) podem ser obtidos via API do Storefront da Wake, desde que o gatilho correto seja definido.

**Pré-requisito:** Alinhar com a Revi qual será o gatilho de consulta na API do Storefront. A Wake pode adaptar o fluxo utilizado com outras plataformas.

**Lojistas já utilizando Wake com esse tipo de integração:** Canal Concept, Francesca, Assa Abloy, Yale (clientes em comum entre Revi e Wake que demandam o recurso).

**Limitação:** Não é um recurso nativo — requer configuração custom via GTM ou gestor de scripts para cada loja.

---

### FAQ-4 — Como consultar dados de checkout via API do Storefront Wake?

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Revi + Wake |
| **JID** | `120363422405368644@g.us` |
| **Data** | 2025-09-16 |
| **Categoria** | `faq` |
| **Tags** | `api-storefront`, `checkout`, `dados-cliente`, `integração-técnica` |

**Pergunta:** É possível coletar dados do cliente (nome, telefone, produtos) via API do Storefront da Wake para disparo de mensagens de recuperação de carrinho?

**Resposta:** Sim, via API do Storefront é possível consultar dados de checkout. O fluxo depende de:

1. Definir o **gatilho** para a consulta (ex: quando o cliente chega na tela de checkout — padrão usado pela Revi em outras plataformas).
2. A Wake pode ajudar a adequar o fluxo para o ambiente Wake, especialmente se a Revi já tiver implementação funcional em outra plataforma.

**Referência prática:** Lojas Canal Concept (`https://www.canal.com.br/`) e Yale (`https://loja.yale.com.br/`) são clientes em comum que podem ser usados como referência para mapeamento do fluxo de checkout.

**Nova tecnologia de checkout Wake:** A Wake possui novo checkout, mas o fluxo de telas permanece essencialmente o mesmo — não impacta a abordagem de integração via tag/API.

---

### HOW_TO-11 — Acesso de Parceiros ao Ambiente Wake (Credenciais para Integração)

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Revi + Wake |
| **JID** | `120363422405368644@g.us` |
| **Data** | 2026-02-25 |
| **Categoria** | `how_to` |
| **Tags** | `acesso`, `credenciais`, `wi-fi`, `partner-day`, `parceiro` |

Para parceiros que precisam de acesso ao ambiente físico ou sistemas Wake:

**Acesso Wi-Fi no Wake Partner Day (LWSA):**
- O responsável de parcerias libera o acesso ao Wi-Fi mediante envio prévio do e-mail.
- Credenciais de acesso do visitante são compartilhadas diretamente pelo grupo WhatsApp (ex: `Rede visitantes! Bruno@userevi.com >>> [senha]`).

**Acesso ao ambiente de integração:**
- O time de Parcerias Wake provisiona acesso direto ao admin/plataforma para parceiros técnicos conforme necessidade.
- Credenciais seguem o padrão de usuário do ambiente (ex: `loja\usuario`).

---

### FAQ-5 — Como avançar com a integração de API quando ela não está no roadmap do trimestre?

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Revi + Wake |
| **JID** | `120363422405368644@g.us` |
| **Data** | 2025-09-01 |
| **Categoria** | `faq` |
| **Tags** | `roadmap`, `api`, `priorização`, `q3`, `q4`, `integração` |

**Pergunta:** O desenvolvimento da API de integração não estava no Q3. Como pressionar para que entre no Q4?

**Resposta:** Quando uma demanda de integração não está no roadmap do trimestre, o caminho para priorização é:

1. **Evidenciar demanda de clientes:** Passar lista de clientes que já solicitaram ou que têm alta probabilidade de adesão. Isso "dá peso" à pauta internamente na Wake.
2. **Solução paliativa:** Enquanto a API não está pronta, avaliar abordagem alternativa (pixel/tag) para atender clientes urgentes.
3. **Transparência proativa:** O time Wake comunica o status real do roadmap. Se não está no Q3, informar claramente e indicar previsão para Q4.
4. **Follow-up regular:** Manter o assunto ativo no grupo WhatsApp para que não caia no esquecimento entre trimestres.

**Clientes que geraram pressão no caso Revi:** Canal Concept, Francesca, Rosie, Assa Abloy, Yale.

---

### BEST_PRACTICES-4 — Identificação de Clientes em Comum como Acelerador de Parceria

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Revi + Wake |
| **JID** | `120363422405368644@g.us` |
| **Data** | 2025-07-18 |
| **Categoria** | `best_practices` |
| **Tags** | `clientes-em-comum`, `acelerador`, `parceria`, `cross-sell` |

Ao iniciar uma parceria técnica, identificar clientes em comum é um acelerador relevante porque:

- Demonstra fit imediato entre as soluções.
- Facilita os primeiros cases (cliente já conhece ambas as plataformas).
- Gera urgência interna para desenvolver integrações (clientes já pedindo o recurso).
- Permite usar lojas reais como referência técnica para mapear fluxos de integração.

**Exemplo Revi + Wake:** Clientes em comum identificados logo na abertura do grupo — Canal Concept, Francesca, Assa Abloy, Yale — aceleraram a discussão técnica sobre carrinho abandonado e deram base para a equipe de produto da Wake priorizar o desenvolvimento.

**Ação recomendada:** No kickoff da parceria, cruzar as bases de clientes e identificar sobreposições. Usar esses clientes como ponto de partida para as primeiras PoCs.

---

## Vurdere & Wake (`120363420777714028@g.us`)

---

### HOW_TO-12 — Processo de Assinatura de Contrato de Parceria Elite Wake

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Vurdere & Wake |
| **JID** | `120363420777714028@g.us` |
| **Data** | 2025-07-17 |
| **Categoria** | `how_to` |
| **Tags** | `contrato`, `parceria-elite`, `assinatura-digital`, `jurídico` |

Fluxo de assinatura do contrato de Parceria Elite Wake:

1. O contrato é enviado pelo time de Parcerias Wake por e-mail e WhatsApp.
2. O parceiro revisa e solicita inclusão de termos da proposta se necessário.
3. A proposta e o contrato são documentos vinculados — condições comerciais ficam na proposta, não no contrato principal.
4. Após alinhamento das condições (ex: condição de pagamento), o time jurídico Wake sobe o documento para assinatura digital.
5. O parceiro informa: nome e e-mail do representante legal + nome e e-mail da testemunha.
6. Os assinantes recebem e-mail com link para assinatura.
7. Após assinatura, o responsável de Parcerias confirma no grupo.

**Tempo médio:** ~1-2 semanas do envio até a assinatura (considerando revisão com o board).

**Nível Elite:** A parceria Vurdere foi classificada como "Elite" desde o início, o que implica atendimento diferenciado e Account Manager dedicado.

---

### TROUBLESHOOTING-4 — Acesso ao Git da Wake para Instalação de Tema de Parceiro

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Vurdere & Wake |
| **JID** | `120363420777714028@g.us` |
| **Data** | 2025-12-15 |
| **Categoria** | `troubleshooting` |
| **Tags** | `git`, `tema`, `acesso`, `merge`, `instalação`, `wake-admin` |

**Problema:** Parceiro (Vurdere) precisa fazer merge do tema instalado no admin Wake com atualizações posteriores da loja (Lojas Lívia), mas não consegue acessar o repositório Git.

**Contexto:** A instalação foi feita via Admin Wake como Tema (não via Git diretamente). Alterações posteriores no código impedem sobrescrever sem merge.

**Resolução passo a passo:**

1. **Vincular conta Git:** No Admin Wake, existe opção de vincular a conta Git. O e-mail deve ser o mesmo usado no admin.
2. **Nome de usuário correto:** O campo de usuário no admin segue o padrão `[loja]\[usuario]`. Ex: `lojaslivia\tech.vurdere`.
3. **Problema de usuário vinculado errado:** Se o acesso foi vinculado a um usuário de outra loja (ex: Lojas Lívia ao invés de Carblue), o login no admin da nova loja exige o prefixo da loja original. O time técnico Wake pode verificar e corrigir esse vínculo.
4. **Se o Git não existir:** O time técnico Wake orienta via WhatsApp do suporte. Canal: acionar via portal de atendimento (bot inicial → time técnico).
5. **Se o acesso existir mas faltar permissão para merge:** Abrir chamado de suporte técnico.

**Contato técnico Wake:** Acessar portal de atendimento → bot → solicitar suporte técnico de parceiro em integração.

---

### TROUBLESHOOTING-5 — Evento "purchase" não disponível no DataLayer de Sucesso de Checkout

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Vurdere & Wake |
| **JID** | `120363420777714028@g.us` |
| **Data** | 2026-04-23 |
| **Categoria** | `troubleshooting` |
| **Tags** | `datalayer`, `gtm`, `purchase-event`, `checkout`, `e-mail-disparo`, `integração` |

**Problema:** O evento `purchase` é encontrado no Tag Assistant mas não persiste até a página de "sucesso do checkout", impossibilitando a coleta de dados para disparo de e-mails (avaliações de produto, recuperação pós-compra).

**Contexto (Vurdere — integração via GTM):**
A Vurdere precisa dos seguintes dados obrigatórios no datalayer da página de sucesso:
```
VurdereDataLayer.transactionId = {{transactionId}};
VurdereDataLayer.transactionProducts = {{transactionProducts}};
VurdereDataLayer.transactionEmail = {{transactionEmail}};
```
Referência: https://[loja].mais.social/install-guide/pt/wake#step6a

**Diagnóstico:** O evento `purchase` existe no datalayer durante o fluxo, mas não permanece disponível na página final de sucesso do checkout.

**Resolução:**
- O time de suporte Wake confirmou que a variável de e-mail (`transactionEmail`) estava disponível no datalayer após ajuste — mas o evento `purchase` não estava.
- Para eventos específicos não nativos: o time técnico Wake orienta que a agência do lojista (ou qualquer dev com permissão) pode editar o datalayer e adicionar o evento/variável necessária.
- Se a informação segue um padrão comum de parceiros, pode ser solicitado ao time técnico Wake para incluir na versão padrão/nativa da plataforma.

**Escalação:** Via WhatsApp do suporte técnico Wake (bot → time tech → orientação sobre o que é padrão vs. customização necessária).

**Manual de instalação Vurdere para Wake:** `https://[loja].mais.social/install-guide/pt/wake#step6a`

---

### HOW_TO-13 — Plano de Lançamento de Parceiro Elite: Entregáveis de Marketing

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Vurdere & Wake |
| **JID** | `120363420777714028@g.us` |
| **Data** | 2026-02-12 |
| **Categoria** | `how_to` |
| **Tags** | `plano-lançamento`, `marketing`, `blog`, `podcast`, `redes-sociais`, `press-release`, `parceiro-elite` |

O Plano de Lançamento para Parceiros Elite Wake inclui os seguintes entregáveis com datas coordenadas:

| Entregável | Responsável | Prazo típico |
|---|---|---|
| Popup no Admin Wake | Wake (produto) | A definir com o parceiro |
| Press Release | Wake (marketing) | ~dia 23 do mês de lançamento |
| Destaque Redes Sociais (LinkedIn, Instagram) | Wake (marketing) | ~dia 19 do mês de lançamento |
| Publicação no Blog Wake | Wake (marketing) | ~dia 19 do mês de lançamento |
| Divulgação em imprensa (ECBR e outros) | Wake (marketing) | ~dia 23 do mês de lançamento |
| Newsletter | Wake (marketing) | Mês seguinte |
| Podcast | Wake + Parceiro | 1ª ou 2ª semana do mês seguinte |
| Case Público | Wake + Parceiro | Após primeiro case real |

**O que o parceiro precisa entregar (inputs):**
1. **Perfil da Marca & Porta-voz:** Apresentação breve da marca + depoimento escrito do principal porta-voz.
2. **Conteúdo para Blog (educativo):** Texto agnóstico sobre o mercado, sem ser propaganda — foco em autoridade. Formato `.docx` simples.
3. **Assets de marca:** Logo, manual de marca.

**Exemplo real (Vurdere):** Publicação no LinkedIn, Instagram e Blog Wake em 19/02/2026 + imprensa no ECBR em 24/02/2026.

---

### FAQ-6 — Como registrar e acompanhar indicações de clientes feitas pela Vurdere?

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Vurdere & Wake |
| **JID** | `120363420777714028@g.us` |
| **Data** | 2025-07-24 |
| **Categoria** | `faq` |
| **Tags** | `indicação`, `pipeline`, `registro`, `parceiro-elite`, `cs` |

**Pergunta:** Como funciona o registro e acompanhamento de indicações de leads feitas pelo parceiro para a Wake?

**Resposta:** Para parceiros Elite como a Vurdere, o fluxo de indicação envolve:

1. O parceiro menciona o lead no grupo WhatsApp ou envia formalmente ao AM responsável.
2. O AM registra na planilha de indicações do programa de parcerias.
3. O CS Wake faz contato com o cliente para qualificar a oportunidade.
4. **Se o cliente não responde:** O CS escalona internamente e tenta outros canais. Se persistir, o lead é tratado como cold outbound.
5. **Para leads frios (sem retorno após e-mail e WhatsApp):** Colocar na régua de outbound do parceiro para não "estressar" o cliente com múltiplos contatos.

**Exemplo prático:** Vurdere indicou Casa dos Parafusos e Construmarques. Casa dos Parafusos fechou com concorrente (Konfidency). Construmarques: CS sem slot disponível — aguardando abertura de agenda para apresentar a Vurdere.

**Gestão de transição de AM:** Quando o AM responsável muda (ex: especialização da carteira), o novo AM deve ser apresentado ao parceiro com uma call de transição para colocar o contexto em dia.

---

### BEST_PRACTICES-5 — Co-marketing Vurdere + Wake: Cases Públicos e LinkedIn

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Vurdere & Wake |
| **JID** | `120363420777714028@g.us` |
| **Data** | 2025-12-18 |
| **Categoria** | `best_practices` |
| **Tags** | `co-marketing`, `case`, `linkedin`, `social-proof`, `parceiro-elite` |

Boas práticas de co-marketing observadas na parceria Vurdere & Wake:

- **Pipeline de cases:** Manter fila de cases para divulgar (ex: Puravida → Soneda → Inbrands). Divulgar em sequência mantém o momentum público da parceria.
- **LinkedIn co-branded:** Parceiro publica o case no LinkedIn e a Wake amplia o alcance pela sua rede. Ambos devem marcar um ao outro nas publicações.
- **Referência cruzada:** Ao divulgar um case, linkar para o post do parceiro aumenta o engajamento e o alcance orgânico de ambas as marcas.
- **Imprensa especializada:** ECBR (E-Commerce Brasil) é um canal relevante para divulgar parcerias martech. A Wake tem relacionamento para publicações editoriais (ex: notícia da parceria Wake + Vurdere publicada no ECBR).
- **Timing:** Publicar cases logo após o resultado ser validado, não esperar o "case perfeito". Cases em andamento com resultados parciais já geram credibilidade.

**Casos publicados Vurdere + Wake:**
- Inbrands: https://www.linkedin.com/posts/vurdere_vurdere-inbrands-activity-7394453155008360448
- Puravida: https://www.linkedin.com/posts/vurdere_vurdere-puravida-activity-7408884321006100482
- Wake + Vurdere (parceria oficial): https://wake.tech/blog/wake-vurdere-parceria-agentic-commerce/
- ECBR: https://ecommercebrasil.com.br/noticias/wake-e-vurdere-firmam-parceria-para-acelerar-o-agentic-commerce

---

### TROUBLESHOOTING-6 — Usuário de Integração Vinculado à Loja Errada no Admin Wake

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Vurdere & Wake |
| **JID** | `120363420777714028@g.us` |
| **Data** | 2026-03-11 |
| **Categoria** | `troubleshooting` |
| **Tags** | `acesso`, `admin', 'usuário`, `loja-errada`, `carblue`, `lojas-livia` |

**Problema:** O usuário criado para o parceiro (Vurdere) em uma nova loja (Carblue) foi vinculado com o prefixo de outra loja (Lojas Lívia), impossibilitando o acesso correto.

**Sintoma:** Ao tentar acessar o admin da Carblue (`https://carblue.fbits.app/`) com a senha padrão, o login falha. Nenhum e-mail de confirmação de novo acesso é recebido.

**Causa:** O time do lojista vinculou o novo acesso a um usuário já existente (mesmo e-mail) em outra loja. O sistema associou o prefixo da loja original.

**Resolução:**
- Para acessar o admin da Carblue, usar o campo de usuário com o prefixo da loja original: `lojaslivia\tech.vurdere`
- Se o acesso ainda não funcionar, o time técnico Wake verifica e corrige o vínculo.

**Lição aprendida:** Quando um mesmo e-mail é usado em múltiplas lojas Wake, o prefixo de usuário no login sempre referencia a primeira loja vinculada. Novos acessos podem ser criados com e-mails diferentes para evitar conflito.

---

*Fim do documento — 29 artigos extraídos de 5 grupos.*
