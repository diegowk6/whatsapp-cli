# KB Extract — Wave 6C (Eitri, O2, Social, Tuna)
> Extracted: 2026-06-04

---

## Grupo: Eitri & Wake
**JID:** `120363400732871223@g.us`

---

### HOW_TO-1 — Como registrar indicações de eventos para contabilização no Salesforce

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Eitri & Wake |
| **JID** | `120363400732871223@g.us` |
| **Data** | 2025-06-23 |
| **Categoria** | `how_to` |
| **Tags** | `salesforce`, `lead`, `indicação`, `parceiro-origem`, `evento` |

Para que uma indicação de cliente feita por um parceiro seja contabilizada como originada pelo parceiro, é obrigatório que o Lead e a Oportunidade já existam no Salesforce com o campo **PARCEIRO ORIGEM** preenchido com o nome da agência **antes** de qualquer fechamento. Isso vale especialmente para indicações feitas no contexto de eventos. A criação retroativa da oportunidade após o fechamento não garante a contabilização.

---

### HOW_TO-2 — Como solicitar acesso sandbox para novos parceiros tech

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Eitri & Wake |
| **JID** | `120363400732871223@g.us` |
| **Data** | 2025-06-24 |
| **Categoria** | `how_to` |
| **Tags** | `sandbox`, `acesso`, `parceiro-tech`, `onboarding` |

Para criar acessos sandbox para um parceiro tech, o parceiro deve informar ao time de Parcerias da Wake os e-mails dos usuários que precisam de acesso. Os demais acessos podem ser criados diretamente pelo parceiro após o ambiente ser provisionado. Exemplo de fluxo: parceiro informa os e-mails principais → Wake cria o ambiente → parceiro gerencia usuários adicionais.

---

### BEST_PRACTICES-1 — Modelo de engajamento embarcado com parceiros tech

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Eitri & Wake |
| **JID** | `120363400732871223@g.us` |
| **Data** | 2025-06-24 |
| **Categoria** | `best_practices` |
| **Tags** | `parceiro-tech`, `go-to-market`, `faturamento`, `embarcado` |

A entrega de soluções de parceiros tech de forma "embarcada" (com faturamento via Wake) representa um modelo com potencial para reduzir o atrito na jornada do cliente. Nesse modelo, a solução do parceiro é apresentada como parte do pacote Wake, com o faturamento consolidado. Isso facilita a adoção por parte dos clientes finais e aumenta a percepção de valor da parceria. Recomenda-se discutir esse modelo com parceiros tech que tenham produtos complementares à plataforma.

---

### TROUBLESHOOTING-1 — reCAPTCHA ativado automaticamente no app bloqueando vendas

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Eitri & Wake |
| **JID** | `120363400732871223@g.us` |
| **Data** | 2026-05-11 |
| **Categoria** | `troubleshooting` |
| **Tags** | `recaptcha`, `app`, `autenticação`, `token`, `DLK`, `Shoulder`, `Balaroti`, `Peahi` |

**Problema:** A loja DLK Modas ficou sem vender no app porque o reCAPTCHA foi ativado automaticamente para o fluxo de autenticação. O parceiro tech (Eitri) não usa reCAPTCHA no app e não conseguia desativá-lo pelo painel, pois a opção de desativação manual foi removida da interface.

**Causa:** O reCAPTCHA pode ser ativado por algum evento ou trigger do sistema Wake, sem ação direta do lojista ou parceiro. Apps que usam tokens dedicados ficam vulneráveis a essa reativação.

**Resolução:**
1. Abrir ticket de suporte solicitando a desativação do reCAPTCHA para o token específico do app: https://atendimento.wake.tech/hc/pt-br/articles/37812329587991
2. Ticket de referência: https://atendimento.wake.tech/hc/pt-br/requests/279913
3. O partner manager pode acionar internamente para priorização.

**Solução de longo prazo em discussão:** Permitir que o parceiro tech informe sua própria chave de reCAPTCHA no painel Wake (similar ao que é feito na VTEX), o que tornaria a experiência mais segura e sem atrito. Lojas impactadas: DLK, Shoulder, Balaroti, Peahi.

---

### FAQ-1 — Posso adicionar minha solução de parceiro no app de evento da Wake Summit?

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Eitri & Wake |
| **JID** | `120363400732871223@g.us` |
| **Data** | 2025-06-25 |
| **Categoria** | `faq` |
| **Tags** | `wake-summit`, `evento`, `app`, `parceria` |

**Pergunta:** A Eitri tem experiência com apps de eventos e gostaria de oferecer sua solução para o Wake Summit. Como proceder?

**Resposta:** A indicação correta é comunicar o interesse ao parceiro de parcerias (no caso, Rafael Correia) para que ele orquestre internamente junto ao time de marketing da Wake. O parceiro deve comunicar brevemente a solução e cases existentes para embasar a conversa interna.

---

### HOW_TO-3 — Como obter o selo de parceiro oficial Wake para apresentações

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Eitri & Wake |
| **JID** | `120363400732871223@g.us` |
| **Data** | 2025-08-15 |
| **Categoria** | `how_to` |
| **Tags** | `selo-parceiro`, `materiais`, `apresentação`, `branding` |

Para obter o selo/badge de parceiro oficial Wake para uso em apresentações e materiais de marketing, o parceiro deve solicitar ao seu parceiro de parcerias (Account Manager) via WhatsApp ou e-mail. O material é enviado por e-mail. O selo é diferente do logo; é a badge específica de "Parceiro Oficial Wake".

---

### HOW_TO-4 — Como funciona o Wake Partner Day (dia presencial semanal)

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Eitri & Wake |
| **JID** | `120363400732871223@g.us` |
| **Data** | 2025-08-13 |
| **Categoria** | `how_to` |
| **Tags** | `wake-partner-day`, `presencial`, `LWSA`, `rotina` |

O Wake Partner Day é o dia oficial de presença de parceiros no escritório Wake. A partir de 27/08/2025, passou das quintas para as **quartas-feiras**. O evento ocorre no escritório LWSA (Rua Itapaiúna, 2434 – São Paulo/SP). O objetivo é aproximar parceiros dos times de CSM, vendas e marketing, com foco em construção de ofertas conjuntas e trabalho de base. Parceiros de fora de SP são bem-vindos a trabalhar do escritório nesses dias.

**Para participar presencialmente:** informar nome, e-mail, CPF e placa do carro (se vier de carro) ao parceiro de parcerias para liberação na portaria e acesso ao Wi-Fi.

---

### BEST_PRACTICES-2 — Como preparar Press Release para anúncio de parceria Wake

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Eitri & Wake |
| **JID** | `120363400732871223@g.us` |
| **Data** | 2025-09-22 |
| **Categoria** | `best_practices` |
| **Tags** | `press-release`, `conteúdo`, `blog`, `parceria`, `comunicação` |

Para publicar o anúncio de parceria no blog da Wake, o parceiro deve preparar um Press Release com:
- Descrição de quem são e o que fazem
- Citação de um executivo da empresa

Exemplos de como ficam as publicações finais:
- https://wake.tech/blog/wake-parceria-biso-inteligencia-de-dados/

O documento é revisado pelo time de marketing da Wake antes da publicação. Após a aprovação, a Wake também publica um artigo no blog complementar.

---

### HOW_TO-5 — Como acionar suporte Wake durante a Black Friday

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Eitri & Wake |
| **JID** | `120363400732871223@g.us` |
| **Data** | 2025-11-24 |
| **Categoria** | `how_to` |
| **Tags** | `black-friday`, `suporte`, `plantão`, `urgência` |

Durante a Black Friday, a Wake oferece atendimento especial com canais e plantão dedicados para urgências. Os detalhes de horários e contatos estão disponíveis em:
https://atendimento.wake.tech/hc/pt-br/articles/35396146551447

Os parceiros tech devem também informar ao time Wake:
1. Se terão escala/plantão em horário especial
2. Qual o contato de urgência que a Wake deve acionar

Esse alinhamento prévio garante agilidade no atendimento conjunto durante o período crítico.

---

### HOW_TO-6 — Como patrocinar o Wake On The Road

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Eitri & Wake |
| **JID** | `120363400732871223@g.us` |
| **Data** | 2026-03-04 |
| **Categoria** | `how_to` |
| **Tags** | `wake-on-the-road`, `patrocínio`, `evento`, `leads` |

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

### TROUBLESHOOTING-2 — Community (Salesforce) aparecendo vazio — sem leads visíveis

| Campo | Valor |
|---|---|
| **Grupo/Chat** | O2 Consultoria & Wake |
| **JID** | `120363420874830254@g.us` |
| **Data** | 2025-06-23 |
| **Categoria** | `troubleshooting` |
| **Tags** | `community`, `salesforce`, `leads`, `bug`, `distribuição` |

**Problema:** O Community Wake (Salesforce) aparecia vazio para a agência O2, impossibilitando atualização de status dos leads.

**Causa:** Instabilidade geral no sistema, não exclusiva da O2.

**Solução temporária:** Enquanto o problema não era regularizado, a Account Manager da Wake atualizava os leads manualmente mediante envio de listagem pela agência.

**Resolução:** O problema foi tratado pela equipe técnica Wake. Para demandas relacionadas ao Community, acionar o responsável de Partner Ops (Renato, a partir de 2025-07-15).

---

### HOW_TO-7 — Processo de repasse de leads pelo canal (Canais)

| Campo | Valor |
|---|---|
| **Grupo/Chat** | O2 Consultoria & Wake |
| **JID** | `120363420874830254@g.us` |
| **Data** | 2025-07-11 |
| **Categoria** | `how_to` |
| **Tags** | `leads`, `canais`, `processo-de-venda`, `proposta`, `plataforma` |

Quando um lead é repassado pelo canal de Canais Wake, a **agência deve apresentar a plataforma Wake E fazer a proposta comercial** — não há um executivo Wake atuando diretamente nessa modalidade. O processo é ponta a ponta pela agência.

**O que a agência precisa fazer:**
1. Entrar em contato com o lead
2. Apresentar a Wake como solução (plataforma + benefícios)
3. Montar proposta comercial incluindo escopo do projeto + mensalidade Wake
4. Conduzir a negociação até o fechamento

A Wake fornece materiais de apoio (PDF de proposta, decks) e pode apoiar em casos específicos com o executivo de produto.

---

### FAQ-2 — A campanha de vendas Wake vale para parceiros tech também?

| Campo | Valor |
|---|---|
| **Grupo/Chat** | O2 Consultoria & Wake |
| **JID** | `120363420874830254@g.us` |
| **Data** | 2025-07-22 |
| **Categoria** | `faq` |
| **Tags** | `campanha-de-vendas`, `parceiro`, `agência`, `certificação` |

**Pergunta:** A campanha de vendas da Wake vale para parceiros que não são agências de implementação (ex.: empresas de consultoria/social)?

**Resposta:** Sim, a campanha vale. Para ser elegível, é necessário:
1. Fazer adesão pelo link da campanha até o prazo indicado
2. Tirar a certificação Wake até a data definida

Sem a certificação, o parceiro não é elegível para os prêmios.

---

### HOW_TO-8 — Programas StoreFront 2.0 e Redesign de Checkout para agências

| Campo | Valor |
|---|---|
| **Grupo/Chat** | O2 Consultoria & Wake |
| **JID** | `120363420874830254@g.us` |
| **Data** | 2025-08-27 |
| **Categoria** | `how_to` |
| **Tags** | `storefront`, `checkout`, `redesign`, `programa`, `agência` |

A Wake lançou dois programas para impulsionar as agências parceiras:
- **StoreFront 2.0**: Escopos Essencial, Básico, Intermediário e Personalizado
- **Redesign de Checkout**: Escopos Essencial e Personalizado

**Vantagens:**
- Modelos pré-definidos exclusivos para a agência (sem concorrência direta)
- Preços e escopos criados em parceria com agências para agilizar o processo de vendas
- Fomenta negócios com lojistas Wake sem agência vinculada

**Como aderir:** Responder "Ok" ao convite da Wake informando qual pacote tem interesse. A adesão é opcional; sem adesão, a agência continua recebendo indicações padrão.

---

### BEST_PRACTICES-3 — Como lidar com leads que não respondem após primeira tentativa

| Campo | Valor |
|---|---|
| **Grupo/Chat** | O2 Consultoria & Wake |
| **JID** | `120363420874830254@g.us` |
| **Data** | 2025-10-24 |
| **Categoria** | `best_practices` |
| **Tags** | `follow-up`, `lead`, `sem-retorno`, `community` |

**Situação comum:** Lead não entra na agenda confirmada ou não responde após primeira tentativa de contato.

**Boas práticas:**
- Registrar a tentativa de contato no Community
- Fazer pelo menos 2 a 3 tentativas via WhatsApp e e-mail em dias diferentes
- Informar o status ao parceiro de parcerias Wake para que possam co-prospectar
- Se o lead não responder por período prolongado (ex.: 2 semanas), encerrar a oportunidade no Community e comunicar a Wake

Se o lead retornar depois do encerramento, reabrir a oportunidade no Community.

---

### FAQ-3 — Qual o padrão correto para o selo Wake no rodapé dos sites de clientes?

| Campo | Valor |
|---|---|
| **Grupo/Chat** | O2 Consultoria & Wake |
| **JID** | `120363420874830254@g.us` |
| **Data** | 2025-09-18 |
| **Categoria** | `faq` |
| **Tags** | `branding`, `rodapé`, `site`, `logo`, `padrão` |

**Pergunta:** Como deve ser exibido o selo Wake no rodapé dos sites de clientes?

**Resposta:** A Wake definiu um novo padrão para a exibição da marca nos rodapés. O guia inclui:
- Formatação do logo Wake junto com o logo da agência
- O texto "Plataforma de E-commerce" deve ser o **único termo com link** no rodapé
- Especificações de fonte e tamanho do texto

Os detalhes completos foram enviados por e-mail às agências parceiras. Em caso de dúvida, consultar o parceiro de parcerias.

---

### BEST_PRACTICES-4 — Como usar notas de reunião do Gemini para briefar parceiros em leads

| Campo | Valor |
|---|---|
| **Grupo/Chat** | O2 Consultoria & Wake |
| **JID** | `120363420874830254@g.us` |
| **Data** | 2026-02-05 |
| **Categoria** | `best_practices` |
| **Tags** | `gemini`, `notas-de-reunião`, `briefing`, `lead`, `processo` |

Para acelerar a qualificação de leads e garantir alinhamento entre todos os envolvidos, a Wake passou a compartilhar as **anotações automáticas do Gemini** das reuniões de qualificação com as agências parceiras. Isso permite que a agência entre na negociação já contextualizada, sem necessidade de repetição das informações pelo lead.

O arquivo é enviado via WhatsApp no mesmo momento do repasse do lead. Isso é especialmente útil em leads quentes ("quentíssimos") onde a velocidade de resposta é crítica.

---

### HOW_TO-9 — Como funciona o Wake Score e a Certificação Wake

| Campo | Valor |
|---|---|
| **Grupo/Chat** | O2 Consultoria & Wake |
| **JID** | `120363420874830254@g.us` |
| **Data** | 2026-01-27 |
| **Categoria** | `how_to` |
| **Tags** | `wake-score`, `certificação`, `obrigatório`, `elegibilidade` |

A **Certificação Wake é um requisito obrigatório** do Wake Score para agências parceiras.

**Regras:**
- Turmas abrem toda **segunda-feira**
- Inscrição ao longo da semana → acesso liberado na segunda seguinte
- Prazo de conclusão após liberação: **30 dias**
- A certificação é **individual** por pessoa (cada membro da equipe precisa tirar a sua)
- **Link de inscrição:** https://forms.gle/AA6NR6vL3uEJXhDu7

Em caso de divergência sobre certificações já realizadas em anos anteriores, enviar os certificados de conclusão ao time de Education Wake para regularização.

---

### FAQ-4 — Como funciona a nova precificação Wake 2026?

| Campo | Valor |
|---|---|
| **Grupo/Chat** | O2 Consultoria & Wake |
| **JID** | `120363420874830254@g.us` |
| **Data** | 2026-01-12 |
| **Categoria** | `faq` |
| **Tags** | `precificação`, `2026`, `tabela-de-preços`, `planos` |

**Pergunta:** Houve mudanças nos planos e preços Wake em 2026?

**Resposta:** Sim. Em janeiro de 2026, a Wake comunicou a Nova Precificação Wake 2026 com atualizações nos planos e novas regras. O comunicado completo foi enviado por e-mail às agências parceiras. Em caso de dúvida, consultar o Account Manager de parcerias.

Ponto de atenção: novas negociações (contratos novos) já entram na tabela de 2026. Contratos existentes seguem as condições anteriores conforme aditivo.

---

## Grupo: Wake <> Social
**JID:** `120363402518075591@g.us`

---

### HOW_TO-10 — Como solicitar loja adicional Wake para cliente existente (guarda-chuva da agência)

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Wake <> Social |
| **JID** | `120363402518075591@g.us` |
| **Data** | 2026-02-20 |
| **Categoria** | `how_to` |
| **Tags** | `loja-adicional`, `contrato`, `aditivo`, `guarda-chuva`, `Social` |

**Situação:** Agência quer criar uma nova loja Wake para um cliente que ainda não faturava online (ex.: cliente que estava apenas em marketplace).

**Processo correto:**
1. A abertura de ticket de suporte **NÃO** é o caminho correto para loja adicional
2. O processo deve ser feito **via e-mail com o Account Manager** da parceria
3. Se a loja ficará atrelada ao contrato existente da agência (guarda-chuva), assinar um **aditivo contratual**
4. **Custo:** R$1.000/mês por loja adicional vinculada ao contrato
5. Se for negociação nova (cliente novo, contrato separado), segue a tabela de preços 2026

**Contato para loja adicional:** ana.vicente@wake.tech com cópia para amanda.araujo@wake.tech

---

### FAQ-5 — O parceiro Social pode participar de ações comerciais Wake conjuntas?

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Wake <> Social |
| **JID** | `120363402518075591@g.us` |
| **Data** | 2025-07-04 |
| **Categoria** | `faq` |
| **Tags** | `social`, `account-planning`, `co-sell`, `clientes-em-comum` |

**Pergunta:** Como funciona a parceria comercial entre uma agência de Social Media e a Wake?

**Resposta:** O modelo recomendado é o **Account Planning cruzado**:
1. Cruzar as bases de clientes de ambos os lados
2. Selecionar ~10 contas target de cada lado com potencial de complementaridade
3. Realizar um Wake Day conjunto para apresentação das contas e sinergia

A agência de Social traz conhecimento de marcas e contatos; a Wake entra com a oferta de plataforma. O timing de fechamento dos leads pode não coincidir com o timing de eventos, então não é necessário forçar fechamentos para participar de eventos como o Wake Summit.

---

### BEST_PRACTICES-5 — Como batalhar contra Shopify em negociações com clientes indecisos

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Wake <> Social |
| **JID** | `120363402518075591@g.us` |
| **Data** | 2026-04-28 |
| **Categoria** | `best_practices` |
| **Tags** | `battlecard`, `shopify`, `vtex`, `concorrência`, `argumentação` |

**Situação:** Lead (Brinquedos Estrelas) com preferência inicial por Shopify, apesar da agência ter apresentado todos os argumentos para Wake.

**O que funciona:**
- Apresentar **depoimentos de clientes Wake** do mesmo segmento
- Solicitar ao time Wake um **battlecard técnico** com comparativo Wake vs Shopify (confere mais credibilidade)
- Usar cases publicados no blog Wake do segmento do lead
- Pedir ao time comercial Wake para co-participar na negociação quando a logo for estratégica

**Observação:** Leads que circulam em grupos de e-commerce ouvem muito sobre VTEX e Shopify; por isso é importante ter materiais de referência de terceiros (cases, testemunhos de outros lojistas) além da argumentação própria.

---

### HOW_TO-11 — Como configurar call quinzenal de acompanhamento de onboarding com Wake

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Wake <> Social |
| **JID** | `120363402518075591@g.us` |
| **Data** | 2026-05-14 |
| **Categoria** | `how_to` |
| **Tags** | `onboarding`, `acompanhamento`, `call-quinzenal`, `go-live` |

Para agências com múltiplos projetos em implementação simultânea, a Wake pode estabelecer uma **call quinzenal com o time de onboarding** para acompanhamento dos projetos até o go-live.

**Para configurar:**
1. Acordar frequência e horário com o Account Manager (ex.: sextas-feiras às 14h30)
2. Fornecer a lista de e-mails de todos do time que tocam os novos projetos
3. A Wake inclui os perfis de onboarding relevantes

Isso é especialmente útil quando o onboarding é conduzido pela agência, mas a Wake precisa de visibilidade sobre o andamento.

---

### FAQ-6 — Qual o canal correto para tirar dúvidas técnicas sobre a plataforma B2B?

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Wake <> Social |
| **JID** | `120363402518075591@g.us` |
| **Data** | 2025-10-06 |
| **Categoria** | `faq` |
| **Tags** | `b2b`, `dúvidas-técnicas`, `suporte`, `canal` |

**Pergunta:** Qual o canal correto para tirar dúvidas técnicas sobre a plataforma B2B Wake?

**Resposta:** Para dúvidas técnicas, usar o canal privado (PV) com o Account Manager ou enviar um áudio no privado descrevendo a dúvida. Isso agiliza o retorno pois o AM pode encaminhar à squad técnica correspondente sem o ruído do grupo geral.

Para treinamento de novos colaboradores sobre cadastro na plataforma, o time de suporte da Wake pode ser acionado diretamente.

---

### HOW_TO-12 — Como solicitar criação de ambiente Wake para novo projeto via guarda-chuva

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Wake <> Social |
| **JID** | `120363402518075591@g.us` |
| **Data** | 2026-03-31 |
| **Categoria** | `how_to` |
| **Tags** | `instância`, `ambiente`, `novo-projeto`, `ticket`, `e-mail` |

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

### HOW_TO-13 — Como funciona o processo de homologação de conector de pagamento Wake

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Tuna & Wake |
| **JID** | `120363403383828730@g.us` |
| **Data** | 2025-10-08 |
| **Categoria** | `how_to` |
| **Tags** | `homologação`, `conector`, `pagamento`, `integração`, `planilha` |

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

### TROUBLESHOOTING-3 — Problema de cache em rota de autenticação no ambiente de integração

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Tuna & Wake |
| **JID** | `120363403383828730@g.us` |
| **Data** | 2025-10-24 |
| **Categoria** | `troubleshooting` |
| **Tags** | `cache`, `autenticação`, `rota`, `integração`, `staging` |

**Problema:** Após atualização de uma rota de autenticação no ambiente de integração, a rota antiga continuou sendo chamada por até 1 hora, mesmo após a mudança.

**Causa:** Cache do checkout/gateway que não é atualizado imediatamente após mudança de configuração.

**Solução:**
- Para cache de checkout, a Wake não consegue forçar a limpeza externamente
- O parceiro deve rodar a **indexação e limpeza de cache pelo painel admin** da Wake
- Aguardar a propagação natural (pode levar entre 30 minutos e 1 hora)
- Se o problema persistir, abrir ticket de suporte: suporte@wake.tech ou WhatsApp (11) 2853-0696

---

### TROUBLESHOOTING-4 — Ambiente da Tuna retornando "acesso negado" sem motivo aparente

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Tuna & Wake |
| **JID** | `120363403383828730@g.us` |
| **Data** | 2025-10-23 |
| **Categoria** | `troubleshooting` |
| **Tags** | `acesso-negado`, `ambiente`, `sandbox`, `parceiro-tech`, `login` |

**Problema:** O ambiente de sandbox da Tuna na Wake (partnertuna.fbits.app) retornou erro de "acesso negado" sem alteração de configuração.

**Causa:** O bloqueio é um mecanismo de segurança quando o horário do PC do usuário não coincide com o horário do servidor Wake.

**Solução:**
1. Verificar se o horário do computador está sincronizado corretamente
2. Limpar o cache do navegador
3. Fazer logout e login novamente

Se o problema persistir após essas etapas, reportar ao time de parceria Wake informando o erro específico.

---

### FAQ-7 — O suporte técnico Wake pode ser adicionado em um grupo de WhatsApp para acompanhamento de integração?

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Tuna & Wake |
| **JID** | `120363403383828730@g.us` |
| **Data** | 2025-10-24 |
| **Categoria** | `faq` |
| **Tags** | `suporte`, `whatsapp`, `canal`, `integração` |

**Pergunta:** É possível criar um grupo de WhatsApp com o número de suporte da Wake para acompanhamento próximo de uma integração?

**Resposta:** Não. O número de suporte Wake (WhatsApp: (11) 2853-0696) é para atendimento individual e não pode ser adicionado a grupos.

**Alternativa recomendada:** Canal **Slack** dedicado com os times técnicos de ambos os lados. A Wake pode criar o canal no workspace LWSA-Tech mediante autorização da liderança. Esse é o canal oficial para integração técnica complexa.

---

### TROUBLESHOOTING-5 — Formulário de cartão não aparece no checkout do app (incidente em produção)

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Tuna & Wake |
| **JID** | `120363403383828730@g.us` |
| **Data** | 2026-03-12 |
| **Categoria** | `troubleshooting` |
| **Tags** | `checkout`, `cartão`, `formulário`, `incidente`, `produção`, `Ugg`, `Tuna` |

**Problema:** A loja Ugg (Grupo Aste) ficou sem apresentar o formulário de cartão no checkout a partir das 17h, fechando pedidos sem capturar dados de cartão.

**Resolução:**
1. Acionamento urgente via WhatsApp ao parceiro de parcerias Wake
2. Escalamento interno para o time técnico Wake via Slack: https://lwsa-tech.slack.com/archives/C09QWKZ9U9F
3. O CTO da Tuna verificou o problema no lado do conector
4. **Problema resolvido em ~1 hora**

**Lição aprendida:** Para incidentes críticos em produção (loja sem vender), acionar simultaneamente:
- WhatsApp do parceiro de parcerias Wake
- Canal Slack técnico dedicado
- O time técnico do parceiro (Tuna) em paralelo

---

### BEST_PRACTICES-6 — Como estruturar o desenvolvimento de integração de pagamento com múltiplos CDs (split de pedido)

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Tuna & Wake |
| **JID** | `120363403383828730@g.us` |
| **Data** | 2025-11-05 |
| **Categoria** | `best_practices` |
| **Tags** | `split-de-pedido`, `múltiplos-CDs`, `json`, `payload`, `integração` |

Ao desenvolver integração de pagamento que suporte múltiplos centros de distribuição (CDs) com split de pedido, o payload de exemplo deve cobrir **cenários reais de uso**. Um payload genérico com dados repetidos gera bloqueio no desenvolvimento do parceiro.

**O que o payload deve incluir obrigatoriamente:**
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

---

### HOW_TO-14 — Como realizar de-para de conectores do parceiro tech com funcionalidades Wake

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Tuna & Wake |
| **JID** | `120363403383828730@g.us` |
| **Data** | 2026-02-26 |
| **Categoria** | `how_to` |
| **Tags** | `de-para`, `conectores`, `funcionalidades`, `tabela-de-funcionalidades`, `Wake` |

Para mapear quais conectores do parceiro tech complementam as funcionalidades Wake, o processo recomendado é:

1. O parceiro tech disponibiliza uma **apresentação completa** com todos os conectores, meios de pagamento e anti-fraudes suportados
2. O time de produto Wake cruza essas informações com a **Tabela de Funcionalidades** da plataforma
3. Os itens validados nas homologações vão compondo a lista oficial de recursos suportados

A apresentação da Tuna com meios de pagamento e anti-fraudes serviu de insumo para atualização da tabela de funcionalidades Wake para 2026.
